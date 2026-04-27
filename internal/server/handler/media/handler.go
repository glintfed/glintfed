package media

import (
	"net/http"

	"glintfed/internal/server/handler/internal"
)

type Handler interface {
	FallbackRedirect(w http.ResponseWriter, r *http.Request)
}

func New() Handler {
	return &handler{}
}

type handler struct{}

func (h *handler) FallbackRedirect(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Media.FallbackRedirect")
	defer span.End()
	// TODO: Implement
}
