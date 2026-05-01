package appregister

import "glintfed/internal/service/account"

type VerifyCodeRequest struct {
	Email      string `json:"email" validate:"required,email"`
	VerifyCode string `json:"verify_code" validate:"required,numeric,len=6"`
}

type VerifyCodeResponse struct {
	Status string `json:"status"`
}

type OnboardingRequest struct {
	Email      string `json:"email" validate:"required,email"`
	VerifyCode string `json:"verify_code" validate:"required,numeric,len=6"`
	Username   string `json:"username" validate:"required,pixelfed_username"`
	Name       string `json:"name" validate:"omitempty"`
	Password   string `json:"password" validate:"required"`
}

type OnboardingResponse struct {
	Status       string          `json:"status"`
	TokenType    string          `json:"token_type"`
	Domain       string          `json:"domain"`
	ExpiresIn    int64           `json:"expires_in"`
	AccessToken  string          `json:"access_token"`
	RefreshToken string          `json:"refresh_token"`
	ClientID     string          `json:"client_id"`
	ClientSecret string          `json:"client_secret"`
	Scope        []string        `json:"scope"`
	User         OnboardingUser  `json:"user"`
	Account      AccountResponse `json:"account"`
}

type OnboardingUser struct {
	PID      string `json:"pid"`
	Username string `json:"username"`
}

type AccountResponse struct {
	*account.ProfileResult
}
