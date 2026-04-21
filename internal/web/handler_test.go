package web_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"glintfed/internal/web"
)

func TestNewHandler(t *testing.T) {
	t.Parallel()

	handler := web.NewHandler()

	tests := []struct {
		name       string
		target     string
		wantStatus int
	}{
		{
			name:       "index",
			target:     "/",
			wantStatus: http.StatusOK,
		},
		{
			name:       "healthz",
			target:     "/healthz",
			wantStatus: http.StatusOK,
		},
		{
			name:       "not found",
			target:     "/missing",
			wantStatus: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			req := httptest.NewRequest(http.MethodGet, tt.target, nil)
			rec := httptest.NewRecorder()

			handler.ServeHTTP(rec, req)

			if rec.Code != tt.wantStatus {
				t.Fatalf("status = %d, want %d", rec.Code, tt.wantStatus)
			}
		})
	}
}
