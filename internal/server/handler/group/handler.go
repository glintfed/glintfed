package group

import (
	"net/http"

	"glintfed/internal/server/handler/internal"
)

type Handler interface {
	GetGroup(w http.ResponseWriter, r *http.Request)
	ShowStatusLikes(w http.ResponseWriter, r *http.Request)
	UpdateGroup(w http.ResponseWriter, r *http.Request)
	GroupLeave(w http.ResponseWriter, r *http.Request)
	CancelJoinRequest(w http.ResponseWriter, r *http.Request)
	JoinGroup(w http.ResponseWriter, r *http.Request)
	MetaBlockSearch(w http.ResponseWriter, r *http.Request)
	ReportCreate(w http.ResponseWriter, r *http.Request)
	ReportAction(w http.ResponseWriter, r *http.Request)
	UpdateMemberInteractionLimits(w http.ResponseWriter, r *http.Request)
	GroupMemberInviteDecline(w http.ResponseWriter, r *http.Request)
	GroupMemberInviteAccept(w http.ResponseWriter, r *http.Request)
	GroupMemberInviteCheck(w http.ResponseWriter, r *http.Request)
	GetMemberInteractionLimits(w http.ResponseWriter, r *http.Request)
	LikePost(w http.ResponseWriter, r *http.Request) // In route but not in Controller search results? Wait.
}

func New() Handler {
	return &handler{}
}

type handler struct{}

func (h *handler) GetGroup(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Group.GetGroup")
	defer span.End()
	// TODO: Implement
}

func (h *handler) ShowStatusLikes(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Group.ShowStatusLikes")
	defer span.End()
	// TODO: Implement
}

func (h *handler) UpdateGroup(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Group.UpdateGroup")
	defer span.End()
	// TODO: Implement
}

func (h *handler) GroupLeave(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Group.GroupLeave")
	defer span.End()
	// TODO: Implement
}

func (h *handler) CancelJoinRequest(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Group.CancelJoinRequest")
	defer span.End()
	// TODO: Implement
}

func (h *handler) JoinGroup(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Group.JoinGroup")
	defer span.End()
	// TODO: Implement
}

func (h *handler) MetaBlockSearch(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Group.MetaBlockSearch")
	defer span.End()
	// TODO: Implement
}

func (h *handler) ReportCreate(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Group.ReportCreate")
	defer span.End()
	// TODO: Implement
}

func (h *handler) ReportAction(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Group.ReportAction")
	defer span.End()
	// TODO: Implement
}

func (h *handler) UpdateMemberInteractionLimits(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Group.UpdateMemberInteractionLimits")
	defer span.End()
	// TODO: Implement
}

func (h *handler) GroupMemberInviteDecline(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Group.GroupMemberInviteDecline")
	defer span.End()
	// TODO: Implement
}

func (h *handler) GroupMemberInviteAccept(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Group.GroupMemberInviteAccept")
	defer span.End()
	// TODO: Implement
}

func (h *handler) GroupMemberInviteCheck(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Group.GroupMemberInviteCheck")
	defer span.End()
	// TODO: Implement
}

func (h *handler) GetMemberInteractionLimits(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Group.GetMemberInteractionLimits")
	defer span.End()
	// TODO: Implement
}

func (h *handler) LikePost(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Group.LikePost")
	defer span.End()
	// TODO: Implement
}
