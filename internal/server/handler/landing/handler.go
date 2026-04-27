package landing

import (
	"net/http"

	"glintfed/internal/server/handler/internal"
)

type Handler interface {
	GetDirectoryApi(w http.ResponseWriter, r *http.Request)
}

func New() Handler {
	return &handler{}
}

type handler struct{}

func (h *handler) GetDirectoryApi(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Landing.GetDirectoryApi")
	defer span.End()
	// TODO: Implement
}
