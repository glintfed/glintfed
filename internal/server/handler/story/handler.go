package story

import (
	"net/http"

	"glintfed/internal/server/handler/internal"
)

type Handler interface {
	GetActivityObject(w http.ResponseWriter, r *http.Request)
}

func New() Handler {
	return &handler{}
}

type handler struct {
}

func (h *handler) GetActivityObject(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Story.GetActivityObject")
	defer span.End()

	// TODO: Implement
}
