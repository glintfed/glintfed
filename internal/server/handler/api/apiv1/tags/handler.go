package tags

import (
	"net/http"

	"glintfed/internal/server/handler/internal"
)

type Handler interface {
	RelatedTags(w http.ResponseWriter, r *http.Request)
	FollowHashtag(w http.ResponseWriter, r *http.Request)
	UnfollowHashtag(w http.ResponseWriter, r *http.Request)
	GetHashtag(w http.ResponseWriter, r *http.Request)
	GetFollowedTags(w http.ResponseWriter, r *http.Request)
}

func New() Handler {
	return &handler{}
}

type handler struct{}

func (h *handler) RelatedTags(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.Tags.RelatedTags")
	defer span.End()
	// TODO: Implement
}

func (h *handler) FollowHashtag(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.Tags.FollowHashtag")
	defer span.End()
	// TODO: Implement
}

func (h *handler) UnfollowHashtag(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.Tags.UnfollowHashtag")
	defer span.End()
	// TODO: Implement
}

func (h *handler) GetHashtag(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.Tags.GetHashtag")
	defer span.End()
	// TODO: Implement
}

func (h *handler) GetFollowedTags(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.Tags.GetFollowedTags")
	defer span.End()
	// TODO: Implement
}
