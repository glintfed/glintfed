package federation

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"glintfed/ent"
	"glintfed/internal/data"
	"glintfed/internal/service/instance"
	"glintfed/internal/service/worker"

	"github.com/go-chi/chi/v5"
)

func TestSharedInboxDispatchesInbox(t *testing.T) {
	body := `{"id":"https://remote.example/activities/1","type":"Create"}`
	inboxWorker := &InboxWorkerServiceMock{
		InboxFunc: func(ctx context.Context, params worker.InboxParams) error {
			if params.Payload.ID != "https://remote.example/activities/1" {
				t.Fatalf("payload ID = %q", params.Payload.ID)
			}
			if params.Payload.Raw != body {
				t.Fatalf("raw payload = %q, want %q", params.Payload.Raw, body)
			}
			return nil
		},
	}
	h := newTestHandler(t, &handlerDeps{inboxWorker: inboxWorker})

	req := httptest.NewRequest(http.MethodPost, "/f/inbox", strings.NewReader(body))
	rec := httptest.NewRecorder()

	h.SharedInbox(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}
	if calls := inboxWorker.InboxCalls(); len(calls) != 1 {
		t.Fatalf("Inbox calls = %d, want 1", len(calls))
	}
}

func TestUserInboxDispatchesValidateWithUsername(t *testing.T) {
	body := `{"id":"https://remote.example/activities/1","type":"Follow"}`
	inboxWorker := &InboxWorkerServiceMock{
		ValidateFunc: func(ctx context.Context, username string, params worker.InboxParams) error {
			if username != "alice" {
				t.Fatalf("username = %q, want alice", username)
			}
			if params.Payload.ID != "https://remote.example/activities/1" {
				t.Fatalf("payload ID = %q", params.Payload.ID)
			}
			return nil
		},
	}
	h := newTestHandler(t, &handlerDeps{inboxWorker: inboxWorker})

	req := httptest.NewRequest(http.MethodPost, "/users/alice/inbox", strings.NewReader(body))
	req = withRouteParam(req, "username", "alice")
	rec := httptest.NewRecorder()

	h.UserInbox(rec, req)

	if got := rec.Result().StatusCode; got != http.StatusOK {
		t.Fatalf("status = %d, want %d", got, http.StatusOK)
	}
	if calls := inboxWorker.ValidateCalls(); len(calls) != 1 {
		t.Fatalf("Validate calls = %d, want 1", len(calls))
	}
}

func TestWebfingerInvalidResourceReturnsBadRequest(t *testing.T) {
	profileModel := &ProfileModelMock{
		GetLocalByUsernameFunc: func(ctx context.Context, username string) (*ent.Profile, error) {
			t.Fatalf("GetLocalByUsername was called for invalid resource")
			return nil, nil
		},
	}
	h := newTestHandler(t, &handlerDeps{profileModel: profileModel})

	req := httptest.NewRequest(http.MethodGet, "/.well-known/webfinger?resource=acct:not-local@example.net", nil)
	rec := httptest.NewRecorder()

	h.Webfinger(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusBadRequest)
	}
	if calls := profileModel.GetLocalByUsernameCalls(); len(calls) != 0 {
		t.Fatalf("GetLocalByUsername calls = %d, want 0", len(calls))
	}
}

func TestNodeinfoWellKnown(t *testing.T) {
	h := newTestHandler(t, &handlerDeps{})

	req := httptest.NewRequest(http.MethodGet, "/.well-known/nodeinfo", nil)
	rec := httptest.NewRecorder()

	h.NodeinfoWellKnown(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}
	if got := rec.Header().Get("Access-Control-Allow-Origin"); got != "*" {
		t.Fatalf("CORS header = %q, want *", got)
	}

	var res NodeinfoWellKnownResponse
	if err := json.NewDecoder(rec.Body).Decode(&res); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if len(res.Links) != 1 {
		t.Fatalf("links = %d, want 1", len(res.Links))
	}
	if got := res.Links[0].Href; got != "https://example.com/api/nodeinfo/2.0.json" {
		t.Fatalf("href = %q, want nodeinfo URL", got)
	}
	if got := res.Links[0].Rel; got != "http://nodeinfo.diaspora.software/ns/schema/2.0" {
		t.Fatalf("rel = %q, want nodeinfo schema", got)
	}
}

func TestHostMeta(t *testing.T) {
	h := newTestHandler(t, &handlerDeps{})

	req := httptest.NewRequest(http.MethodGet, "/.well-known/host-meta", nil)
	rec := httptest.NewRecorder()

	h.HostMeta(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}
	if got := rec.Header().Get("Content-Type"); got != "application/xrd+xml" {
		t.Fatalf("content type = %q, want application/xrd+xml", got)
	}
	if body := rec.Body.String(); !strings.Contains(body, `template="https://example.com/.well-known/webfinger?resource={uri}"`) {
		t.Fatalf("body does not contain webfinger template: %s", body)
	}
}

func TestNodeinfo(t *testing.T) {
	h := newTestHandler(t, &handlerDeps{
		instanceService: &InstanceServiceMock{
			GetBlockedDomainsFunc: func(ctx context.Context) (map[string]struct{}, error) {
				return map[string]struct{}{}, nil
			},
			NodeinfoStatsFunc: func(ctx context.Context) (*instance.NodeinfoStats, error) {
				return &instance.NodeinfoStats{
					Usage: instance.NodeinfoUsage{
						LocalPosts: 3,
						Users: instance.NodeinfoUserUsage{
							Total: 5,
						},
					},
				}, nil
			},
			NodeinfoFeaturesFunc: func(ctx context.Context) (*instance.NodeinfoFeatures, error) {
				return &instance.NodeinfoFeatures{ActivityPub: true}, nil
			},
		},
	})

	req := httptest.NewRequest(http.MethodGet, "/api/nodeinfo/2.0.json", nil)
	rec := httptest.NewRecorder()

	h.Nodeinfo(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}
	if got := rec.Header().Get("Access-Control-Allow-Origin"); got != "*" {
		t.Fatalf("CORS header = %q, want *", got)
	}

	var res NodeinfoResponse
	if err := json.NewDecoder(rec.Body).Decode(&res); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if res.Software.Name != "pixelfed" {
		t.Fatalf("software name = %q, want pixelfed", res.Software.Name)
	}
	if res.Usage.LocalPosts != 3 {
		t.Fatalf("localPosts = %d, want 3", res.Usage.LocalPosts)
	}
	if !res.OpenRegistrations {
		t.Fatalf("openRegistrations = false, want true")
	}
}

type handlerDeps struct {
	profileModel    ProfileModel
	statusModel     StatusModel
	instanceService InstanceService
	inboxWorker     InboxWorkerService
}

func newTestHandler(t *testing.T, deps *handlerDeps) Handler {
	t.Helper()

	if deps.profileModel == nil {
		deps.profileModel = &ProfileModelMock{
			GetLocalByUsernameFunc: func(ctx context.Context, username string) (*ent.Profile, error) {
				t.Fatalf("unexpected GetLocalByUsername call")
				return nil, nil
			},
			RemoteURLExistsFunc: func(ctx context.Context, remoteURL string) (bool, error) {
				t.Fatalf("unexpected RemoteURLExists call")
				return false, nil
			},
		}
	}
	if deps.statusModel == nil {
		deps.statusModel = &StatusModelMock{
			GetLocalPostsCountFunc: func(ctx context.Context) (int, error) {
				t.Fatalf("unexpected GetLocalPostsCount call")
				return 0, nil
			},
			ObjectURLExistsFunc: func(ctx context.Context, objectURL string) (bool, error) {
				t.Fatalf("unexpected ObjectURLExists call")
				return false, nil
			},
		}
	}
	if deps.instanceService == nil {
		deps.instanceService = &InstanceServiceMock{
			GetBlockedDomainsFunc: func(ctx context.Context) (map[string]struct{}, error) {
				return map[string]struct{}{}, nil
			},
			NodeinfoStatsFunc: func(ctx context.Context) (*instance.NodeinfoStats, error) {
				t.Fatalf("unexpected NodeinfoStats call")
				return nil, nil
			},
			NodeinfoFeaturesFunc: func(ctx context.Context) (*instance.NodeinfoFeatures, error) {
				t.Fatalf("unexpected NodeinfoFeatures call")
				return nil, nil
			},
		}
	}
	if deps.inboxWorker == nil {
		deps.inboxWorker = &InboxWorkerServiceMock{
			DeleteFunc: func(ctx context.Context, params worker.InboxParams) error {
				t.Fatalf("unexpected Delete call")
				return nil
			},
			InboxFunc: func(ctx context.Context, params worker.InboxParams) error {
				t.Fatalf("unexpected Inbox call")
				return nil
			},
			ValidateFunc: func(ctx context.Context, username string, params worker.InboxParams) error {
				t.Fatalf("unexpected Validate call")
				return nil
			},
		}
	}

	u, err := url.Parse("https://example.com")
	if err != nil {
		t.Fatal(err)
	}
	cfg := &data.Config{
		App: data.AppConfig{
			Name:     "Glintfed",
			Version:  "1.2.3",
			URL:      u,
			URLValue: u.String(),
			Auth: data.AuthConfig{
				EnableRegistration: true,
			},
			Federation: data.FederationConfig{
				NodeInfo:  data.NodeInfoConfig{Enabled: true},
				Webfinger: data.WebfingerConfig{Enabled: true},
				Activitypub: data.ActivitypubConfig{
					Enabled:     true,
					SharedInbox: true,
					Inbox:       true,
				},
			},
		},
	}
	return New(cfg, deps.profileModel, deps.statusModel, deps.instanceService, deps.inboxWorker)
}

func withRouteParam(r *http.Request, key string, value string) *http.Request {
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add(key, value)
	return r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))
}
