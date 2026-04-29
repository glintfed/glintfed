package appregister

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"glintfed/internal/data"
)

func TestVerifyCode(t *testing.T) {
	appRegisterModel := &AppRegisterModelMock{
		VerifyCodeExistsFunc: func(ctx context.Context, email string, code string) (bool, error) {
			if email != "alice@example.com" {
				t.Fatalf("email = %q, want alice@example.com", email)
			}
			if code != "123456" {
				t.Fatalf("code = %q, want 123456", code)
			}
			return true, nil
		},
	}
	h := New(
		appRegisterTestConfig(t),
		appRegisterModel,
		&UserModelMock{},
		&RefreshTokenRepositoryMock{},
		&AccountServiceMock{},
	)

	req := httptest.NewRequest(http.MethodPost, "/api/auth/app-code-verify", strings.NewReader(`{"email":"Alice@Example.com","verify_code":"123456"}`))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	h.VerifyCode(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	var res VerifyCodeResponse
	if err := json.NewDecoder(rec.Body).Decode(&res); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if res.Status != "success" {
		t.Fatalf("status response = %q, want success", res.Status)
	}
	if calls := appRegisterModel.VerifyCodeExistsCalls(); len(calls) != 1 {
		t.Fatalf("VerifyCodeExists calls = %d, want 1", len(calls))
	}
}

func TestVerifyCodeDisabledRegistration(t *testing.T) {
	cfg := appRegisterTestConfig(t)
	cfg.App.Auth.InAppRegistration = false
	h := New(
		cfg,
		&AppRegisterModelMock{},
		&UserModelMock{},
		&RefreshTokenRepositoryMock{},
		&AccountServiceMock{},
	)

	req := httptest.NewRequest(http.MethodPost, "/api/auth/app-code-verify", strings.NewReader(`{"email":"alice@example.com","verify_code":"123456"}`))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	h.VerifyCode(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusNotFound)
	}
}

func TestOnboardingInvalidVerificationCode(t *testing.T) {
	appRegisterModel := &AppRegisterModelMock{
		VerifyCodeExistsFunc: func(ctx context.Context, email string, code string) (bool, error) {
			if email != "alice@example.com" {
				t.Fatalf("email = %q, want alice@example.com", email)
			}
			if code != "123456" {
				t.Fatalf("code = %q, want 123456", code)
			}
			return false, nil
		},
	}
	h := New(
		appRegisterTestConfig(t),
		appRegisterModel,
		&UserModelMock{},
		&RefreshTokenRepositoryMock{},
		&AccountServiceMock{},
	)

	body := `{"email":"Alice@Example.com","verify_code":"123456","username":"alice","name":"Alice","password":"password123"}`
	req := httptest.NewRequest(http.MethodPost, "/api/auth/onboarding", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	h.Onboarding(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusBadRequest)
	}

	var res ErrorResponse
	if err := json.NewDecoder(rec.Body).Decode(&res); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if res.Status != "error" {
		t.Fatalf("status response = %q, want error", res.Status)
	}
	if calls := appRegisterModel.VerifyCodeExistsCalls(); len(calls) != 1 {
		t.Fatalf("VerifyCodeExists calls = %d, want 1", len(calls))
	}
}

func appRegisterTestConfig(t *testing.T) *data.Config {
	t.Helper()

	u, err := url.Parse("https://example.com")
	if err != nil {
		t.Fatal(err)
	}
	return &data.Config{
		App: data.AppConfig{
			URL:               u,
			MaxNameLength:     30,
			MinPasswordLength: 8,
			Auth: data.AuthConfig{
				InAppRegistration:  true,
				EnableRegistration: true,
			},
		},
	}
}
