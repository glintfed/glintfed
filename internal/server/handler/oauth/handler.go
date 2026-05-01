package oauth

import (
	"context"
	"errors"
	"glintfed/ent"
	"glintfed/internal/data"
	"glintfed/internal/lib/fositestore"
	"glintfed/internal/lib/liblogs"
	"glintfed/internal/server/handler/internal"
	"log/slog"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/ory/fosite"
)

// Service defines the OAuth2 HTTP handlers.
type Handler interface {
	Authorize(w http.ResponseWriter, r *http.Request)
	Token(w http.ResponseWriter, r *http.Request)
	Revoke(w http.ResponseWriter, r *http.Request)
}

type UserModel interface {
	Authenticate(ctx context.Context, username, password string) (*ent.User, error)
}

type handler struct {
	cfg *data.Config

	provider fosite.OAuth2Provider
	store    *fositestore.Store

	userModel UserModel
}

// New creates a new OAuth service.
func New(cfg *data.Config, provider fosite.OAuth2Provider, store *fositestore.Store, userModel UserModel) Handler {
	return &handler{
		cfg: cfg,

		provider: provider,
		store:    store,

		userModel: userModel,
	}
}

// Authorize handles GET /oauth/authorize.
//
// It validates the authorize request via fosite, then redirects the user to
// the configured LoginUrl so the frontend can authenticate them. The full
// original query string is forwarded as a `next` parameter so the frontend
// can redirect back to this endpoint after login.
//
// If LoginUrl is not configured, the endpoint returns ErrAccessDenied.
func (h *handler) Authorize(w http.ResponseWriter, r *http.Request) {
	ctx, span := internal.T.Start(r.Context(), "OAuth.Authorize")
	defer span.End()

	authReq, err := h.provider.NewAuthorizeRequest(ctx, r)
	if err != nil {
		h.provider.WriteAuthorizeError(ctx, w, authReq, err)
		return
	}

	http.Redirect(w, r, h.buildLoginRedirect(r).String(), http.StatusSeeOther)
}

func (h *handler) buildLoginRedirect(r *http.Request) (redirect *url.URL) {
	nextURL := h.cfg.App.URL.JoinPath("/oauth/authorize")
	nextURL.RawQuery = r.URL.RawQuery

	redirect = h.cfg.App.Auth.LoginURL
	q := redirect.Query()
	q.Set("next", nextURL.String())
	redirect.RawQuery = q.Encode()

	return
}

// Token handles POST /oauth/token for all grant types.
func (h *handler) Token(w http.ResponseWriter, r *http.Request) {
	ctx, span := internal.T.Start(r.Context(), "OAuth.Token")
	defer span.End()

	if r.FormValue("grant_type") == "password" {
		h.handlePasswordGrant(ctx, w, r)
		return
	}

	accessReq, err := h.provider.NewAccessRequest(ctx, r, &fosite.DefaultSession{})
	if err != nil {
		h.provider.WriteAccessError(ctx, w, accessReq, err)
		return
	}

	accessResp, err := h.provider.NewAccessResponse(ctx, accessReq)
	if err != nil {
		h.provider.WriteAccessError(ctx, w, accessReq, err)
		return
	}

	h.provider.WriteAccessResponse(ctx, w, accessReq, accessResp)
}

// handlePasswordGrant processes the resource owner password credentials grant type.
// Credentials are validated manually; then tokens are issued via fositestore.
func (h *handler) handlePasswordGrant(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "" || password == "" {
		internal.WriteError(w, internal.NewValidationError("missing username or password"))
		return
	}

	user, err := h.userModel.Authenticate(ctx, username, password)
	if err != nil {
		slog.ErrorContext(ctx, "authenticate failed", liblogs.ErrAttr(err))
		internal.WriteError(w, err)
		return
	}

	// Parse requested scopes from form; fall back to default scopes.
	scopeStr := r.FormValue("scope")
	scopes := []string{"read", "write", "follow", "push"}
	if scopeStr != "" {
		scopes = strings.Split(scopeStr, " ")
	}

	// Look up the client to validate client_id.
	clientID := r.FormValue("client_id")
	if clientID == "" {
		internal.WriteError(w, internal.NewValidationError("missing client_id"))
		return
	}

	client, err := h.store.GetClient(ctx, clientID)
	if err != nil {
		slog.ErrorContext(ctx, "failed to get store client", liblogs.ErrAttr(err))
		internal.WriteError(w, errors.New("failed to get store client"))
		return
	}

	subject := strconv.FormatUint(user.ID, 10)
	now := time.Now()

	session := &fosite.DefaultSession{
		Subject:  subject,
		Username: username,
		ExpiresAt: map[fosite.TokenType]time.Time{
			fosite.AccessToken:  now.Add(h.cfg.App.Auth.OAuth.AccessTokenLifespan),
			fosite.RefreshToken: now.Add(h.cfg.App.Auth.OAuth.RefreshTokenLifespan),
		},
	}

	req := fosite.NewRequest()
	req.ID = uuid.Must(uuid.NewV7()).String()
	req.Client = client
	req.RequestedAt = now
	req.Session = session
	req.SetRequestedScopes(fosite.Arguments(scopes))
	for _, scope := range scopes {
		req.GrantScope(scope)
	}

	accessToken, refreshToken, err := h.store.CreatePersonalAccessTokens(ctx, req)
	if err != nil {
		slog.ErrorContext(ctx, "password grant: failed to create tokens", liblogs.ErrAttr(err))
		internal.WriteError(w, err)
		return
	}

	internal.WriteJSON(w, http.StatusOK, IssueTokenByPasswordGrantResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "Bearer",
		ExpiresIn:    int64(h.cfg.App.Auth.OAuth.AccessTokenLifespan.Seconds()),
		Scope:        scopeStr,
	})
}

// Revoke handles POST /oauth/revoke.
func (h *handler) Revoke(w http.ResponseWriter, r *http.Request) {
	ctx, span := internal.T.Start(r.Context(), "OAuth.Revoke")
	defer span.End()

	h.provider.WriteRevocationResponse(
		ctx, w,
		h.provider.NewRevocationRequest(ctx, r),
	)
}
