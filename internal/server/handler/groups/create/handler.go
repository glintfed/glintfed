package create

import (
	"net/http"

	"glintfed/internal/server/handler/internal"
)

type Handler interface {
	CheckCreatePermission(w http.ResponseWriter, r *http.Request)
	StoreGroup(w http.ResponseWriter, r *http.Request)
}

func New() Handler {
	return &handler{}
}

type handler struct{}

func (h *handler) CheckCreatePermission(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsCreate.CheckCreatePermission")
	defer span.End()
	// TODO: Implement
}

func (h *handler) StoreGroup(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsCreate.StoreGroup")
	defer span.End()
	// TODO: Implement
}
