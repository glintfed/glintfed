package media

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"glintfed/internal/data"

	"github.com/go-chi/chi/v5"
)

func TestFallbackRedirectCloudStorageDisabled(t *testing.T) {
	h := New(&data.Config{}, &MediaModelMock{})

	req := httptest.NewRequest(http.MethodGet, "/storage/m/_v2/1/a/b/c.jpg", nil)
	rec := httptest.NewRecorder()

	h.FallbackRedirect(rec, req)

	if rec.Code != http.StatusFound {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusFound)
	}
	if got := rec.Header().Get("Location"); got != "/storage/no-preview.png" {
		t.Fatalf("Location = %q, want %q", got, "/storage/no-preview.png")
	}
}

func TestFallbackRedirectUsesV2PathAndProfileID(t *testing.T) {
	cdnURL := "https://cdn.example/media.jpg"
	mediaModel := &MediaModelMock{
		CDNURLFunc: func(ctx context.Context, profileID uint64, path string) (*string, error) {
			if profileID != 42 {
				t.Fatalf("profileID = %d, want 42", profileID)
			}
			if path != "public/m/_v2/42/mhash/uhash/file.jpg" {
				t.Fatalf("path = %q, want public/m/_v2/42/mhash/uhash/file.jpg", path)
			}
			return &cdnURL, nil
		},
	}
	h := New(&data.Config{App: data.AppConfig{CloudStorage: true}}, mediaModel)

	req := httptest.NewRequest(http.MethodGet, "/storage/m/_v2/42/mhash/uhash/file.jpg", nil)
	req = withURLParams(req, map[string]string{
		"pid":   "42",
		"mhash": "mhash",
		"uhash": "uhash",
		"f":     "file.jpg",
	})
	rec := httptest.NewRecorder()

	h.FallbackRedirect(rec, req)

	if rec.Code != http.StatusFound {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusFound)
	}
	if got := rec.Header().Get("Location"); got != cdnURL {
		t.Fatalf("Location = %q, want %q", got, cdnURL)
	}
	if calls := mediaModel.CDNURLCalls(); len(calls) != 1 {
		t.Fatalf("CDNURL calls = %d, want 1", len(calls))
	}
}

func TestFallbackRedirectMissingMediaUsesNoPreview(t *testing.T) {
	mediaModel := &MediaModelMock{
		CDNURLFunc: func(ctx context.Context, profileID uint64, path string) (*string, error) {
			return nil, nil
		},
	}
	h := New(&data.Config{App: data.AppConfig{CloudStorage: true}}, mediaModel)

	req := httptest.NewRequest(http.MethodGet, "/storage/m/_v2/42/mhash/uhash/file.jpg", nil)
	req = withURLParams(req, map[string]string{
		"pid":   "42",
		"mhash": "mhash",
		"uhash": "uhash",
		"f":     "file.jpg",
	})
	rec := httptest.NewRecorder()

	h.FallbackRedirect(rec, req)

	if rec.Code != http.StatusFound {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusFound)
	}
	if got := rec.Header().Get("Location"); got != "/storage/no-preview.png" {
		t.Fatalf("Location = %q, want %q", got, "/storage/no-preview.png")
	}
}

func withURLParams(r *http.Request, params map[string]string) *http.Request {
	rctx := chi.NewRouteContext()
	for key, value := range params {
		rctx.URLParams.Add(key, value)
	}
	return r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))
}
