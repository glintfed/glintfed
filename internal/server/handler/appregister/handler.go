package appregister

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
	"time"

	"glintfed/ent"
	"glintfed/internal/data"
	"glintfed/internal/lib/liblogs"
	"glintfed/internal/model"
	"glintfed/internal/repo/oauth"
	"glintfed/internal/server/handler/internal"
	"glintfed/internal/service/account"

	"github.com/go-playground/validator/v10"
)

var onboardingScopes = [...]string{"read", "write", "follow", "push"}

type Handler interface {
	VerifyCode(w http.ResponseWriter, r *http.Request)
	Onboarding(w http.ResponseWriter, r *http.Request)
}

//go:generate go tool moq -rm -out mock_app_register_model.go . AppRegisterModel
type AppRegisterModel interface {
	// VerifyCodeExists checks whether an AppRegister record exists for the given email and
	// verify_code that was created within the past 90 days.
	VerifyCodeExists(ctx context.Context, email string, code string) (bool, error)
	// DeleteByEmail removes the AppRegister record for the given email after onboarding completes.
	DeleteByEmail(ctx context.Context, email string) error
}

//go:generate go tool moq -rm -out mock_user_model.go . UserModel
type UserModel interface {
	// Create inserts a new user with the given parameters. The implementation is responsible
	// for hashing the plaintext password before storing it.
	Create(ctx context.Context, params model.CreateUserParams) (*ent.User, error)
}

//go:generate go tool moq -rm -out mock_refresh_token_repository.go . RefreshTokenRepository
type RefreshTokenRepository interface {
	// CreateTokens creates an OAuth access token and a refresh token for the given user ID
	// with the specified scopes, and returns the resulting token details.
	Create(ctx context.Context, params oauth.RefreshTokenCreateParams) (*oauth.TokenResult, error)
}

//go:generate go tool moq -rm -out mock_account_service.go . AccountService
type AccountService interface {
	GetProfile(ctx context.Context, params account.GetProfileParams) (*account.ProfileResult, error)
}

func New(
	cfg *data.Config,

	appRegisterModel AppRegisterModel,
	userModel UserModel,

	refreshTokenRepo RefreshTokenRepository,

	accountService AccountService,
) Handler {
	v := validator.New(validator.WithRequiredStructEnabled())
	_ = v.RegisterValidation("pixelfed_username", validateUsername)

	return &handler{
		cfg: cfg,
		v:   v,

		appRegisterModel: appRegisterModel,
		userModel:        userModel,

		refreshTokenRepo: refreshTokenRepo,

		accountService: accountService,
	}
}

type handler struct {
	cfg *data.Config
	v   *validator.Validate

	appRegisterModel AppRegisterModel
	userModel        UserModel

	refreshTokenRepo RefreshTokenRepository

	accountService AccountService
}

func (h *handler) VerifyCode(w http.ResponseWriter, r *http.Request) {
	ctx, span := internal.T.Start(r.Context(), "AppRegister.VerifyCode")
	defer span.End()

	if !h.registrationEnabled(w, r) {
		return
	}

	var req VerifyCodeRequest
	if !h.decode(w, r, &req) {
		return
	}
	req.Email = strings.ToLower(req.Email)

	if err := h.v.Struct(req); err != nil {
		internal.WriteJSON(w, http.StatusUnprocessableEntity, ErrorResponse{Status: "error", Message: "Invalid request."})
		return
	}

	exists, err := h.appRegisterModel.VerifyCodeExists(ctx, req.Email, req.VerifyCode)
	if err != nil {
		slog.ErrorContext(ctx, "failed to verify code", liblogs.ErrAttr(err))
		internal.WriteError(w, err)
		return
	}

	status := "error"
	if exists {
		status = "success"
	}
	internal.WriteJSON(w, http.StatusOK, VerifyCodeResponse{Status: status})
}

func (h *handler) Onboarding(w http.ResponseWriter, r *http.Request) {
	ctx, span := internal.T.Start(r.Context(), "AppRegister.Onboarding")
	defer span.End()

	if !h.registrationEnabled(w, r) {
		return
	}

	var req OnboardingRequest
	if !h.decode(w, r, &req) {
		return
	}
	req.Email = strings.ToLower(req.Email)

	if err := h.v.Struct(req); err != nil {
		internal.WriteJSON(w, http.StatusUnprocessableEntity, ErrorResponse{Status: "error", Message: "Invalid request."})
		return
	}
	if err := h.v.Var(req.Name, fmt.Sprintf("omitempty,max=%d", h.cfg.App.MaxNameLength)); err != nil {
		internal.WriteJSON(w, http.StatusUnprocessableEntity, ErrorResponse{Status: "error", Message: "Invalid request."})
		return
	}
	if err := h.v.Var(req.Password, fmt.Sprintf("required,min=%d", h.cfg.App.MinPasswordLength)); err != nil {
		internal.WriteJSON(w, http.StatusUnprocessableEntity, ErrorResponse{Status: "error", Message: "Invalid request."})
		return
	}

	exists, err := h.appRegisterModel.VerifyCodeExists(ctx, req.Email, req.VerifyCode)
	if err != nil {
		slog.ErrorContext(ctx, "failed to verify code", liblogs.ErrAttr(err))
		internal.WriteError(w, err)
		return
	}
	if !exists {
		internal.WriteJSON(w, http.StatusBadRequest, ErrorResponse{
			Status:  "error",
			Message: "Invalid verification code, please try again later.",
		})
		return
	}

	user, err := h.userModel.Create(ctx, model.CreateUserParams{
		Name:            req.Name,
		Username:        req.Username,
		Email:           req.Email,
		Password:        req.Password,
		AppRegisterIP:   "",
		RegisterSource:  "app",
		EmailVerifiedAt: time.Now(),
	})
	if err != nil {
		slog.ErrorContext(ctx, "failed to create user", liblogs.ErrAttr(err))
		internal.WriteError(w, err)
		return
	}

	tokens, err := h.refreshTokenRepo.Create(ctx, oauth.RefreshTokenCreateParams{
		UserID: user.ID,
		Scopes: onboardingScopes[:],
	})
	if err != nil {
		slog.ErrorContext(ctx, "failed to create refresh token", liblogs.ErrAttr(err))
		internal.WriteError(w, err)
		return
	}

	if err := h.appRegisterModel.DeleteByEmail(ctx, req.Email); err != nil {
		slog.ErrorContext(ctx, "failed to delete app register record", liblogs.ErrAttr(err))
		// non-fatal: continue and return the successful response
	}

	profile, err := h.accountService.GetProfile(ctx, account.GetProfileParams{
		ProfileID: *user.ProfileID,
	})
	if err != nil {
		slog.ErrorContext(ctx, "failed to get profile", liblogs.ErrAttr(err))
	}

	internal.WriteJSON(w, http.StatusOK, OnboardingResponse{
		Status:       "success",
		TokenType:    "Bearer",
		Domain:       h.cfg.App.URL.Host,
		ExpiresIn:    tokens.ExpiresIn,
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
		ClientID:     tokens.ClientID,
		ClientSecret: tokens.ClientSecret,
		Scope:        onboardingScopes[:],
		User: OnboardingUser{
			PID:      strconv.FormatUint(*user.ProfileID, 10),
			Username: *user.Username,
		},
		Account: AccountResponse{
			ProfileResult: profile,
		},
	})
}

func (h *handler) registrationEnabled(w http.ResponseWriter, r *http.Request) bool {
	if !h.cfg.App.Auth.InAppRegistration {
		http.NotFound(w, r)
		return false
	}
	if !h.cfg.App.Auth.EnableRegistration {
		http.Redirect(w, r, "/", http.StatusFound)
		return false
	}
	return true
}

func (h *handler) decode(w http.ResponseWriter, r *http.Request, dst any) bool {
	if strings.Contains(r.Header.Get("Content-Type"), "application/json") {
		if err := json.NewDecoder(r.Body).Decode(dst); err != nil {
			internal.WriteJSON(w, http.StatusBadRequest, ErrorResponse{Status: "error", Message: "Invalid request."})
			return false
		}
		return true
	}

	if err := r.ParseForm(); err != nil {
		internal.WriteJSON(w, http.StatusBadRequest, ErrorResponse{Status: "error", Message: "Invalid request."})
		return false
	}

	switch req := dst.(type) {
	case *VerifyCodeRequest:
		req.Email = r.FormValue("email")
		req.VerifyCode = r.FormValue("verify_code")
	case *OnboardingRequest:
		req.Email = r.FormValue("email")
		req.VerifyCode = r.FormValue("verify_code")
		req.Username = r.FormValue("username")
		req.Name = r.FormValue("name")
		req.Password = r.FormValue("password")
	}
	return true
}

func validateUsername(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	if len(value) < 2 || len(value) > 30 {
		return false
	}
	if strings.HasSuffix(value, ".php") || strings.HasSuffix(value, ".js") || strings.HasSuffix(value, ".css") {
		return false
	}

	separators := strings.Count(value, "-") + strings.Count(value, "_") + strings.Count(value, ".")
	if separators > 1 {
		return false
	}

	if !isAlphaNum(value[0]) || !isAlphaNum(value[len(value)-1]) {
		return false
	}

	hasAlpha := false
	for i := 0; i < len(value); i++ {
		c := value[i]
		if isAlpha(c) {
			hasAlpha = true
			continue
		}
		if isDigit(c) || c == '-' || c == '_' || c == '.' {
			continue
		}
		return false
	}
	return hasAlpha
}

func isAlphaNum(c byte) bool {
	return isAlpha(c) || isDigit(c)
}

func isAlpha(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}
