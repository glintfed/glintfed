package healthcheck

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGet(t *testing.T) {
	h := New()

	req := httptest.NewRequest(http.MethodGet, "/api/service/health-check", nil)
	rec := httptest.NewRecorder()

	h.Get(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}
	if got := rec.Header().Get("Content-Type"); got != "text/plain" {
		t.Fatalf("content type = %q, want text/plain", got)
	}
	if got := rec.Header().Get("Cache-Control"); got != "max-age=0, must-revalidate, no-cache, no-store" {
		t.Fatalf("cache control = %q, want no-cache policy", got)
	}
	if got := rec.Body.String(); got != "OK" {
		t.Fatalf("body = %q, want OK", got)
	}
}
