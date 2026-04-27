package meta

import (
	"net/http"

	"glintfed/internal/server/handler/internal"
)

type Handler interface {
	DeleteGroup(w http.ResponseWriter, r *http.Request)
}

func New() Handler {
	return &handler{}
}

type handler struct{}

func (h *handler) DeleteGroup(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsMeta.DeleteGroup")
	defer span.End()
	// TODO: Implement
}
