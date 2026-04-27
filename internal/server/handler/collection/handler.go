package collection

import (
	"net/http"

	"glintfed/internal/server/handler/internal"
)

type Handler interface {
	GetUserCollections(w http.ResponseWriter, r *http.Request)
	GetItems(w http.ResponseWriter, r *http.Request)
	GetCollection(w http.ResponseWriter, r *http.Request)
	StoreId(w http.ResponseWriter, r *http.Request)
	Store(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	DeleteId(w http.ResponseWriter, r *http.Request)
	GetSelfCollections(w http.ResponseWriter, r *http.Request)
}

func New() Handler {
	return &handler{}
}

type handler struct{}

func (h *handler) GetUserCollections(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Collection.AvatarUpdate")
	defer span.End()
	// TODO: Implement
}

func (h *handler) GetItems(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Collection.GetItems")
	defer span.End()
	// TODO: Implement
}

func (h *handler) GetCollection(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Collection.GetCollection")
	defer span.End()
	// TODO: Implement
}

func (h *handler) StoreId(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Collection.StoreId")
	defer span.End()
	// TODO: Implement
}

func (h *handler) Store(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Collection.Store")
	defer span.End()
	// TODO: Implement
}

func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Collection.Delete")
	defer span.End()
	// TODO: Implement
}

func (h *handler) DeleteId(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Collection.DeleteId")
	defer span.End()
	// TODO: Implement
}

func (h *handler) GetSelfCollections(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Collection.GetSelfCollections")
	defer span.End()
	// TODO: Implement
}
