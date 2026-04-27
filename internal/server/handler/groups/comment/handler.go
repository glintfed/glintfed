package comment

import (
	"net/http"

	"glintfed/internal/server/handler/internal"
)

type Handler interface {
	GetComments(w http.ResponseWriter, r *http.Request)
	StoreComment(w http.ResponseWriter, r *http.Request)
	StoreCommentPhoto(w http.ResponseWriter, r *http.Request)
	DeleteComment(w http.ResponseWriter, r *http.Request)
	LikePost(w http.ResponseWriter, r *http.Request)
	UnlikePost(w http.ResponseWriter, r *http.Request)
}

func New() Handler {
	return &handler{}
}

type handler struct{}

func (h *handler) GetComments(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsComment.GetComments")
	defer span.End()
	// TODO: Implement
}

func (h *handler) StoreComment(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsComment.StoreComment")
	defer span.End()
	// TODO: Implement
}

func (h *handler) StoreCommentPhoto(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsComment.StoreCommentPhoto")
	defer span.End()
	// TODO: Implement
}

func (h *handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsComment.DeleteComment")
	defer span.End()
	// TODO: Implement
}

func (h *handler) LikePost(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsComment.LikePost")
	defer span.End()
	// TODO: Implement
}

func (h *handler) UnlikePost(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsComment.UnlikePost")
	defer span.End()
	// TODO: Implement
}
