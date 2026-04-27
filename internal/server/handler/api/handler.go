package api

import (
	"net/http"

	"glintfed/internal/server/handler/internal"
)

type Handler interface {
	AvatarUpdate(w http.ResponseWriter, r *http.Request)
	Notifications(w http.ResponseWriter, r *http.Request)
	VerifyCredentials(w http.ResponseWriter, r *http.Request)
	AccountLikes(w http.ResponseWriter, r *http.Request)
	Archive(w http.ResponseWriter, r *http.Request)
	Unarchive(w http.ResponseWriter, r *http.Request)
	ArchivedPosts(w http.ResponseWriter, r *http.Request)
	SiteConfiguration(w http.ResponseWriter, r *http.Request)
	UserRecommendations(w http.ResponseWriter, r *http.Request)
}

func New() Handler {
	return &handler{}
}

type handler struct{}

func (h *handler) AvatarUpdate(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.AvatarUpdate")
	defer span.End()
	// TODO: Implement
}

func (h *handler) Notifications(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.Notifications")
	defer span.End()
	// TODO: Implement
}

func (h *handler) VerifyCredentials(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.VerifyCredentials")
	defer span.End()
	// TODO: Implement
}

func (h *handler) AccountLikes(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.AccountLikes")
	defer span.End()
	// TODO: Implement
}

func (h *handler) Archive(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.Archive")
	defer span.End()
	// TODO: Implement
}

func (h *handler) Unarchive(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.Unarchive")
	defer span.End()
	// TODO: Implement
}

func (h *handler) ArchivedPosts(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ArchivedPosts")
	defer span.End()
	// TODO: Implement
}

func (h *handler) SiteConfiguration(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.SiteConfiguration")
	defer span.End()
	// TODO: Implement
}

func (h *handler) UserRecommendations(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.UserRecommendations")
	defer span.End()
	// TODO: Implement
}
