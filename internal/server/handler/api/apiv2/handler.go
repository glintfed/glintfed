package apiv2

import (
	"net/http"

	"glintfed/internal/server/handler/internal"
)

type Handler interface {
	Instance(w http.ResponseWriter, r *http.Request)
	Search(w http.ResponseWriter, r *http.Request)
	GetWebsocketConfig(w http.ResponseWriter, r *http.Request)
	MediaUploadV2(w http.ResponseWriter, r *http.Request)
	StatusContextV2(w http.ResponseWriter, r *http.Request)
	StatusDescendants(w http.ResponseWriter, r *http.Request)
	StatusAncestors(w http.ResponseWriter, r *http.Request)
}

func New() Handler {
	return &handler{}
}

type handler struct{}

func (h *handler) Instance(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV2.Instance")
	defer span.End()
	// TODO: Implement
}

func (h *handler) Search(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV2.Search")
	defer span.End()
	// TODO: Implement
}

func (h *handler) GetWebsocketConfig(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV2.GetWebsocketConfig")
	defer span.End()
	// TODO: Implement
}

func (h *handler) MediaUploadV2(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV2.MediaUploadV2")
	defer span.End()
	// TODO: Implement
}

func (h *handler) StatusContextV2(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV2.StatusContextV2")
	defer span.End()
	// TODO: Implement
}

func (h *handler) StatusDescendants(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV2.StatusDescendants")
	defer span.End()
	// TODO: Implement
}

func (h *handler) StatusAncestors(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV2.StatusAncestors")
	defer span.End()
	// TODO: Implement
}
