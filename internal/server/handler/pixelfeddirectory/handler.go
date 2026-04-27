package pixelfeddirectory

import (
	"net/http"

	"glintfed/internal/server/handler/internal"
)

type Handler interface {
	Get(w http.ResponseWriter, r *http.Request)
}

func New() Handler {
	return &handler{}
}

type handler struct{}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "PixelfedDirectory.Get")
	defer span.End()
	// TODO: Implement
}
