package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestWriteErrorValidationError(t *testing.T) {
	err := fmt.Errorf("validate request: %w", NewValidationError("Invalid request."))
	rec := httptest.NewRecorder()

	WriteError(rec, err)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusBadRequest)
	}
	if got := rec.Header().Get("Content-Type"); got != "application/json" {
		t.Fatalf("content type = %q, want application/json", got)
	}

	var res ValidationError
	if err := json.NewDecoder(rec.Body).Decode(&res); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if res.Status != "error" {
		t.Fatalf("status response = %q, want error", res.Status)
	}
	if res.Message != "Invalid request." {
		t.Fatalf("message response = %q, want Invalid request.", res.Message)
	}
}

func TestNewValidationErrorWithValidatorError(t *testing.T) {
	type request struct {
		Email      string `validate:"required,email"`
		VerifyCode string `validate:"required,len=6"`
	}

	err := validator.New(validator.WithRequiredStructEnabled()).Struct(request{
		Email:      "invalid",
		VerifyCode: "123",
	})

	validationErr := NewValidationError("validator", err)

	want := "validator: The email field must be a valid email address. The verify_code field must be 6 characters."
	if validationErr.Message != want {
		t.Fatalf("message = %q, want %q", validationErr.Message, want)
	}
}
