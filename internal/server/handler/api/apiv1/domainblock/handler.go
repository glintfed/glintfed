package domainblock

import (
	"net/http"

	"glintfed/internal/server/handler/internal"
)

type Handler interface {
	Index(w http.ResponseWriter, r *http.Request)
	Store(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func New() Handler {
	return &handler{}
}

type handler struct{}

func (h *handler) Index(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.DomainBlocks.Index")
	defer span.End()
	// TODO: Implement
}

func (h *handler) Store(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.DomainBlocks.Store")
	defer span.End()
	// TODO: Implement
}

func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.DomainBlocks.Delete")
	defer span.End()
	// TODO: Implement
}
