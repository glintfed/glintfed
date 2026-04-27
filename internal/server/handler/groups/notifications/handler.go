package notifications

import (
	"net/http"

	"glintfed/internal/server/handler/internal"
)

type Handler interface {
	SelfGlobalNotifications(w http.ResponseWriter, r *http.Request)
}

func New() Handler {
	return &handler{}
}

type handler struct{}

func (h *handler) SelfGlobalNotifications(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsNotifications.SelfGlobalNotifications")
	defer span.End()
	// TODO: Implement
}
