package instanceactor

import (
	"context"
	"encoding/json"
	"net/http"

	"glintfed/internal/data"
	"glintfed/internal/server/handler/internal"
)

type Handler interface {
	Profile(w http.ResponseWriter, r *http.Request)
	Inbox(w http.ResponseWriter, r *http.Request)
	Outbox(w http.ResponseWriter, r *http.Request)
}

//go:generate go tool moq -rm -out mock_instance_actor_model.go . InstanceActorModel
type InstanceActorModel interface {
	PublicKey(ctx context.Context) (*string, error)
}

func New(cfg *data.Config, instanceActorModel InstanceActorModel) Handler {
	return &handler{
		cfg: cfg,

		instanceActorModel: instanceActorModel,
	}
}

type handler struct {
	cfg *data.Config

	instanceActorModel InstanceActorModel
}

func (h *handler) Profile(w http.ResponseWriter, r *http.Request) {
	ctx, span := internal.T.Start(r.Context(), "InstanceActor.Profile")
	defer span.End()

	publicKey, err := h.instanceActorModel.PublicKey(ctx)
	if err != nil {
		internal.WriteError(w, err)
		return
	}

	internal.WriteActivityJSON(w, http.StatusOK, ActorResponse{
		Context:           activityContext(),
		ID:                h.appURL("/i/actor"),
		Type:              "Application",
		Inbox:             h.appURL("/i/actor/inbox"),
		Outbox:            h.appURL("/i/actor/outbox"),
		PreferredUsername: h.domain(),
		PublicKey: PublicKey{
			ID:           h.appURL("/i/actor#main-key"),
			Owner:        h.appURL("/i/actor"),
			PublicKeyPem: publicKey,
		},
		ManuallyApprovesFollowers: true,
		URL:                       h.appURL("/site/kb/instance-actor"),
	})
}

func (h *handler) Inbox(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "InstanceActor.Inbox")
	defer span.End()

	w.WriteHeader(http.StatusOK)
}

func (h *handler) Outbox(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "InstanceActor.Outbox")
	defer span.End()

	internal.WriteActivityJSON(w, http.StatusOK, OutboxResponse{
		Context:    activityContext(),
		ID:         h.appURL("/i/actor/outbox"),
		Type:       "OrderedCollection",
		TotalItems: 0,
		First:      h.appURL("/i/actor/outbox?page=true"),
		Last:       h.appURL("/i/actor/outbox?min_id=0page=true"),
	})
}

func (h *handler) domain() string {
	if h.cfg.App.URL != nil && h.cfg.App.URL.Host != "" {
		return h.cfg.App.URL.Host
	}
	return h.cfg.App.URLValue
}

func (h *handler) appURL(path string) string {
	return h.cfg.App.URL.JoinPath(path).String()
}

func activityContext() json.RawMessage {
	return json.RawMessage(`["https://www.w3.org/ns/activitystreams","https://w3id.org/security/v1",{"manuallyApprovesFollowers":"as:manuallyApprovesFollowers","toot":"http://joinmastodon.org/ns#","featured":{"@id":"toot:featured","@type":"@id"},"featuredTags":{"@id":"toot:featuredTags","@type":"@id"},"alsoKnownAs":{"@id":"as:alsoKnownAs","@type":"@id"},"movedTo":{"@id":"as:movedTo","@type":"@id"},"schema":"http://schema.org#","PropertyValue":"schema:PropertyValue","value":"schema:value","discoverable":"toot:discoverable","Device":"toot:Device","Ed25519Signature":"toot:Ed25519Signature","Ed25519Key":"toot:Ed25519Key","Curve25519Key":"toot:Curve25519Key","EncryptedMessage":"toot:EncryptedMessage","publicKeyBase64":"toot:publicKeyBase64","deviceId":"toot:deviceId","claim":{"@type":"@id","@id":"toot:claim"},"fingerprintKey":{"@type":"@id","@id":"toot:fingerprintKey"},"identityKey":{"@type":"@id","@id":"toot:identityKey"},"devices":{"@type":"@id","@id":"toot:devices"},"messageFranking":"toot:messageFranking","messageType":"toot:messageType","cipherText":"toot:cipherText","suspended":"toot:suspended"}]`)
}
