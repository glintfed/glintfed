package discover

import (
	"net/http"

	"glintfed/internal/server/handler/internal"
)

type Handler interface {
	TrendingApi(w http.ResponseWriter, r *http.Request)
	TrendingHashtags(w http.ResponseWriter, r *http.Request)
	DiscoverNetworkTrending(w http.ResponseWriter, r *http.Request)
}

func New() Handler {
	return &handler{}
}

type handler struct{}

func (h *handler) TrendingApi(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Discover.TrendingApi")
	defer span.End()
	// TODO: Implement
}

func (h *handler) TrendingHashtags(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Discover.TrendingHashtags")
	defer span.End()
	// TODO: Implement
}
func (h *handler) DiscoverNetworkTrending(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Discover.DiscoverNetworkTrending")
	defer span.End()
	// TODO: Implement
}
