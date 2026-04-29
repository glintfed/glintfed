package story

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"testing/synctest"
	"time"

	"glintfed/ent"
	"glintfed/internal/data"

	"github.com/go-chi/chi/v5"
)

func TestGetActivityObjectWithBearerToken(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		token := "secret-token"
		username := "alice"
		storyType := "photo"
		mime := "image/jpeg"
		path := "stories/alice/photo.jpg"
		now := time.Now()

		profileModel := &ProfileModelMock{
			GetStoryFunc: func(ctx context.Context, gotUsername string, storyID uint64) (*ent.Story, error) {
				if gotUsername != username {
					t.Fatalf("username = %q, want %q", gotUsername, username)
				}
				if storyID != 99 {
					t.Fatalf("storyID = %d, want 99", storyID)
				}
				return &ent.Story{
					ID:           99,
					Type:         &storyType,
					Duration:     5,
					Mime:         &mime,
					Path:         &path,
					BearcapToken: &token,
					CreatedAt:    timePtr(now.Add(-time.Minute)),
					ExpiresAt:    timePtr(now.Add(time.Hour)),
					CanReply:     true,
					CanReact:     true,
					Edges: ent.StoryEdges{
						Profile: &ent.Profile{
							Username: &username,
						},
					},
				}, nil
			},
		}
		h := New(storyTestConfig(t), profileModel)

		req := httptest.NewRequest(http.MethodGet, "/stories/alice/99", nil)
		req.Header.Set("Accept", "application/activity+json")
		req.Header.Set("Authorization", "Bearer "+token)
		req = withStoryURLParams(req, map[string]string{"username": username, "id": "99"})
		rec := httptest.NewRecorder()

		h.GetActivityObject(rec, req)

		if rec.Code != http.StatusOK {
			t.Fatalf("status = %d, want %d; body=%q", rec.Code, http.StatusOK, rec.Body.String())
		}

		var res StoryActivityResponse
		if err := json.NewDecoder(rec.Body).Decode(&res); err != nil {
			t.Fatalf("decode response: %v", err)
		}
		if res.ID != "https://example.com/stories/alice/99" {
			t.Fatalf("id = %q, want story activity URL", res.ID)
		}
		if res.Attachment.Type != "Image" {
			t.Fatalf("attachment type = %q, want Image", res.Attachment.Type)
		}
		if calls := profileModel.GetStoryCalls(); len(calls) != 1 {
			t.Fatalf("GetStory calls = %d, want 1", len(calls))
		}
	})
}

func TestGetActivityObjectRedirectsWhenClientDoesNotWantJSON(t *testing.T) {
	h := New(storyTestConfig(t), &ProfileModelMock{})

	req := httptest.NewRequest(http.MethodGet, "/stories/alice/99", nil)
	req = withStoryURLParams(req, map[string]string{"username": "alice", "id": "99"})
	rec := httptest.NewRecorder()

	h.GetActivityObject(rec, req)

	if rec.Code != http.StatusFound {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusFound)
	}
	if got := rec.Header().Get("Location"); got != "/stories/alice" {
		t.Fatalf("Location = %q, want /stories/alice", got)
	}
}

func storyTestConfig(t *testing.T) *data.Config {
	t.Helper()

	u, err := url.Parse("https://example.com")
	if err != nil {
		t.Fatal(err)
	}
	return &data.Config{
		App: data.AppConfig{
			URL: u,
			Instance: data.InstanceConfig{
				Stories: data.StoriesConfig{Enabled: true},
			},
		},
	}
}

func withStoryURLParams(r *http.Request, params map[string]string) *http.Request {
	rctx := chi.NewRouteContext()
	for key, value := range params {
		rctx.URLParams.Add(key, value)
	}
	return r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))
}

func timePtr(t time.Time) *time.Time {
	return &t
}
