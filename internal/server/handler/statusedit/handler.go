package statusedit

import (
	"net/http"

	"glintfed/internal/server/handler/internal"
)

type Handler interface {
	Store(w http.ResponseWriter, r *http.Request)
	History(w http.ResponseWriter, r *http.Request)
}

func New() Handler {
	return &handler{}
}

type handler struct{}

func (h *handler) Store(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "StatusEdit.Store")
	defer span.End()
	// TODO: Implement
}

func (h *handler) History(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "StatusEdit.History")
	defer span.End()
	// TODO: Implement
}
