package userappsettings

import (
	"net/http"

	"glintfed/internal/server/handler/internal"
)

type Handler interface {
	Get(w http.ResponseWriter, r *http.Request)
	Store(w http.ResponseWriter, r *http.Request)
}

func New() Handler {
	return &handler{}
}

type handler struct{}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "UserAppSettings.Get")
	defer span.End()
	// TODO: Implement
}

func (h *handler) Store(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "UserAppSettings.Store")
	defer span.End()
	// TODO: Implement
}
