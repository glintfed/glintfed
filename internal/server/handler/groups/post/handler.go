package post

import (
	"net/http"

	"glintfed/internal/server/handler/internal"
)

type Handler interface {
	StorePost(w http.ResponseWriter, r *http.Request)
	DeletePost(w http.ResponseWriter, r *http.Request)
	LikePost(w http.ResponseWriter, r *http.Request)
	UnlikePost(w http.ResponseWriter, r *http.Request)
	GetGroupMedia(w http.ResponseWriter, r *http.Request)
	GetStatus(w http.ResponseWriter, r *http.Request)
}

func New() Handler {
	return &handler{}
}

type handler struct{}

func (h *handler) StorePost(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsPost.StorePost")
	defer span.End()
	// TODO: Implement
}

func (h *handler) DeletePost(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsPost.DeletePost")
	defer span.End()
	// TODO: Implement
}

func (h *handler) LikePost(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsPost.LikePost")
	defer span.End()
	// TODO: Implement
}

func (h *handler) UnlikePost(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsPost.UnlikePost")
	defer span.End()
	// TODO: Implement
}

func (h *handler) GetGroupMedia(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupPost.GetGroupMedia")
	defer span.End()
	// TODO: Implement
}

func (h *handler) GetStatus(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsPost.GetStatus")
	defer span.End()
	// TODO: Implement
}
