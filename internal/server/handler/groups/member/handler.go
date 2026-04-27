package member

import (
	"net/http"

	"glintfed/internal/server/handler/internal"
)

type Handler interface {
	GetGroupMembers(w http.ResponseWriter, r *http.Request)
	GetGroupMemberJoinRequests(w http.ResponseWriter, r *http.Request)
	HandleGroupMemberJoinRequest(w http.ResponseWriter, r *http.Request)
	GetGroupMember(w http.ResponseWriter, r *http.Request)
	GetGroupMemberCommonIntersections(w http.ResponseWriter, r *http.Request)
}

func New() Handler {
	return &handler{}
}

type handler struct{}

func (h *handler) GetGroupMembers(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsMember.GetGroupMembers")
	defer span.End()
	// TODO: Implement
}

func (h *handler) GetGroupMemberJoinRequests(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsMember.GetGroupMemberJoinRequests")
	defer span.End()
	// TODO: Implement
}

func (h *handler) HandleGroupMemberJoinRequest(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsMember.HandleGroupMemberJoinRequest")
	defer span.End()
	// TODO: Implement
}

func (h *handler) GetGroupMember(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsMember.GetGroupMember")
	defer span.End()
	// TODO: Implement
}

func (h *handler) GetGroupMemberCommonIntersections(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsMember.GetGroupMemberCommonIntersections")
	defer span.End()
	// TODO: Implement
}
