package instanceactor

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"glintfed/internal/data"
)

func TestProfile(t *testing.T) {
	publicKey := "PUBLIC KEY"
	model := &InstanceActorModelMock{
		PublicKeyFunc: func(ctx context.Context) (*string, error) {
			return &publicKey, nil
		},
	}
	h := New(testConfig(t), model)

	req := httptest.NewRequest(http.MethodGet, "/i/actor", nil)
	rec := httptest.NewRecorder()

	h.Profile(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}
	if got := rec.Header().Get("Content-Type"); got != "application/activity+json" {
		t.Fatalf("Content-Type = %q, want application/activity+json", got)
	}

	var res ActorResponse
	if err := json.NewDecoder(rec.Body).Decode(&res); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if res.ID != "https://example.com/i/actor" {
		t.Fatalf("id = %q, want https://example.com/i/actor", res.ID)
	}
	if res.PublicKey.PublicKeyPem == nil || *res.PublicKey.PublicKeyPem != publicKey {
		t.Fatalf("publicKeyPem = %v, want %q", res.PublicKey.PublicKeyPem, publicKey)
	}
	if calls := model.PublicKeyCalls(); len(calls) != 1 {
		t.Fatalf("PublicKey calls = %d, want 1", len(calls))
	}
}

func TestInbox(t *testing.T) {
	h := New(testConfig(t), &InstanceActorModelMock{})

	req := httptest.NewRequest(http.MethodPost, "/i/actor/inbox", nil)
	rec := httptest.NewRecorder()

	h.Inbox(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}
}

func TestOutbox(t *testing.T) {
	h := New(testConfig(t), &InstanceActorModelMock{})

	req := httptest.NewRequest(http.MethodGet, "/i/actor/outbox", nil)
	rec := httptest.NewRecorder()

	h.Outbox(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	var res OutboxResponse
	if err := json.NewDecoder(rec.Body).Decode(&res); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if res.Type != "OrderedCollection" {
		t.Fatalf("type = %q, want OrderedCollection", res.Type)
	}
	if res.First != "https://example.com/i/actor/outbox%3Fpage=true" {
		t.Fatalf("first = %q, want generated outbox page URL", res.First)
	}
}

func testConfig(t *testing.T) *data.Config {
	t.Helper()

	u, err := url.Parse("https://example.com")
	if err != nil {
		t.Fatal(err)
	}
	return &data.Config{
		App: data.AppConfig{
			URL:      u,
			URLValue: u.String(),
		},
	}
}
