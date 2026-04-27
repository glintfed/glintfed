package search

import (
	"net/http"

	"glintfed/internal/server/handler/internal"
)

type Handler interface {
	InviteFriendsToGroup(w http.ResponseWriter, r *http.Request)
	SearchFriendsToInvite(w http.ResponseWriter, r *http.Request)
	SearchGlobalResults(w http.ResponseWriter, r *http.Request)
	SearchLocalAutocomplete(w http.ResponseWriter, r *http.Request)
	SearchAddRecent(w http.ResponseWriter, r *http.Request)
	SearchGetRecent(w http.ResponseWriter, r *http.Request)
}

func New() Handler {
	return &handler{}
}

type handler struct{}

func (h *handler) InviteFriendsToGroup(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsSearch.InviteFriendsToGroup")
	defer span.End()
	// TODO: Implement
}

func (h *handler) SearchFriendsToInvite(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsSearch.SearchFriendsToInvite")
	defer span.End()
	// TODO: Implement
}

func (h *handler) SearchGlobalResults(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsSearch.SearchGlobalResults")
	defer span.End()
	// TODO: Implement
}

func (h *handler) SearchLocalAutocomplete(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsSearch.SearchLocalAutocomplete")
	defer span.End()
	// TODO: Implement
}

func (h *handler) SearchAddRecent(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsSearch.SearchAddRecent")
	defer span.End()
	// TODO: Implement
}

func (h *handler) SearchGetRecent(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsSearch.SearchGetRecent")
	defer span.End()
	// TODO: Implement
}
