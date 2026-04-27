package domainblocks

import (
	"net/http"

	"glintfed/internal/server/handler/internal"
)

type Handler interface {
	Index(w http.ResponseWriter, r *http.Request)
	Show(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func New() Handler {
	return &handler{}
}

type handler struct{}

func (h *handler) Index(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.Admin.DomainBlocks.Index")
	defer span.End()
	// TODO: Implement
}

func (h *handler) Show(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.Admin.DomainBlocks.Show")
	defer span.End()
	// TODO: Implement
}

func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.Admin.DomainBlocks.Create")
	defer span.End()
	// TODO: Implement
}

func (h *handler) Update(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.Admin.DomainBlocks.Update")
	defer span.End()
	// TODO: Implement
}

func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.Admin.DomainBlocks.Delete")
	defer span.End()
	// TODO: Implement
}
