package api

import (
	"net/http"

	"glintfed/internal/server/handler/internal"
)

type Handler interface {
	GetConfig(w http.ResponseWriter, r *http.Request)
	GetGroupAccount(w http.ResponseWriter, r *http.Request)
	GetGroupCategories(w http.ResponseWriter, r *http.Request)
	GetGroupsByCategory(w http.ResponseWriter, r *http.Request)
	GetRecommendedGroups(w http.ResponseWriter, r *http.Request)
	GetSelfGroups(w http.ResponseWriter, r *http.Request)
}

func New() Handler {
	return &handler{}
}

type handler struct{}

func (h *handler) GetConfig(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsAPI.GetConfig")
	defer span.End()
	// TODO: Implement
}

func (h *handler) GetGroupAccount(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsAPI.GetGroupAccount")
	defer span.End()
	// TODO: Implement
}

func (h *handler) GetGroupCategories(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsAPI.GetGroupCategories")
	defer span.End()
	// TODO: Implement
}

func (h *handler) GetGroupsByCategory(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsAPI.GetGroupsByCategory")
	defer span.End()
	// TODO: Implement
}

func (h *handler) GetRecommendedGroups(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsAPI.GetRecommendedGroups")
	defer span.End()
	// TODO: Implement
}

func (h *handler) GetSelfGroups(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsAPI.GetSelfGroups")
	defer span.End()
	// TODO: Implement
}
