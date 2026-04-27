package instanceactor

import (
	"net/http"

	"glintfed/internal/server/handler/internal"
)

type Handler interface {
	Profile(w http.ResponseWriter, r *http.Request)
	Inbox(w http.ResponseWriter, r *http.Request)
	Outbox(w http.ResponseWriter, r *http.Request)
}

func New() Handler {
	return &handler{}
}

type handler struct{}

func (h *handler) Profile(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "InstanceActor.Profile")
	defer span.End()

	// TODO: Implement
}

func (h *handler) Inbox(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "InstanceActor.Inbox")
	defer span.End()
	// TODO: Implement
}

func (h *handler) Outbox(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "InstanceActor.Outbox")
	defer span.End()
	// TODO: Implement
}
