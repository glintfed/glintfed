package feed

import (
	"net/http"

	"glintfed/internal/server/handler/internal"
)

type Handler interface {
	GetSelfFeed(w http.ResponseWriter, r *http.Request)
	GetGroupProfileFeed(w http.ResponseWriter, r *http.Request)
	GetGroupFeed(w http.ResponseWriter, r *http.Request)
}

func New() Handler {
	return &handler{}
}

type handler struct{}

func (h *handler) GetSelfFeed(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsFeed.GetSelfFeed")
	defer span.End()
	// TODO: Implement
}

func (h *handler) GetGroupProfileFeed(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsFeed.GetGroupProfileFeed")
	defer span.End()
	// TODO: Implement
}

func (h *handler) GetGroupFeed(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsFeed.GetGroupFeed")
	defer span.End()
	// TODO: Implement
}
