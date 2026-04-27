package discover

import (
	"net/http"

	"glintfed/internal/server/handler/internal"
)

type Handler interface {
	GetDiscoverPopular(w http.ResponseWriter, r *http.Request)
	GetDiscoverNew(w http.ResponseWriter, r *http.Request)
}

func New() Handler {
	return &handler{}
}

type handler struct{}

func (h *handler) GetDiscoverPopular(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsDiscover.GetDiscoverPopular")
	defer span.End()
	// TODO: Implement
}

func (h *handler) GetDiscoverNew(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsDiscover.GetDiscoverNew")
	defer span.End()
	// TODO: Implement
}
