package compose

import (
	"net/http"

	"glintfed/internal/server/handler/internal"
)

type Handler interface {
	SearchLocation(w http.ResponseWriter, r *http.Request)
	ComposeSettings(w http.ResponseWriter, r *http.Request)
}

func New() Handler {
	return &handler{}
}

type handler struct{}

func (h *handler) SearchLocation(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Compose.SearchLocation")
	defer span.End()
	// TODO: Implement
}

func (h *handler) ComposeSettings(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Compose.ComposeSettings")
	defer span.End()
	// TODO: Implement
}
