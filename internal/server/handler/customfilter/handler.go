package customfilter

import (
	"net/http"

	"glintfed/internal/server/handler/internal"
)

type Handler interface {
	Index(w http.ResponseWriter, r *http.Request)
	Show(w http.ResponseWriter, r *http.Request)
	Store(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func New() Handler {
	return &handler{}
}

type handler struct{}

func (h *handler) Index(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "CustomFilter.Index")
	defer span.End()
	// TODO: Implement
}

func (h *handler) Show(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "CustomFilter.Show")
	defer span.End()
	// TODO: Implement
}

func (h *handler) Store(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "CustomFilter.Store")
	defer span.End()
	// TODO: Implement
}

func (h *handler) Update(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "CustomFilter.Update")
	defer span.End()
	// TODO: Implement
}

func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "CustomFilter.Delete")
	defer span.End()
	// TODO: Implement
}
