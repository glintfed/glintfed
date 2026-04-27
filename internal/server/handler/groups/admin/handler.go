package admin

import (
	"net/http"

	"glintfed/internal/server/handler/internal"
)

type Handler interface {
	GetAdminTabs(w http.ResponseWriter, r *http.Request)
	GetInteractionLogs(w http.ResponseWriter, r *http.Request)
	GetBlocks(w http.ResponseWriter, r *http.Request)
	ExportBlocks(w http.ResponseWriter, r *http.Request)
	AddBlock(w http.ResponseWriter, r *http.Request)
	UndoBlock(w http.ResponseWriter, r *http.Request)
	GetReportList(w http.ResponseWriter, r *http.Request)
}

func New() Handler {
	return &handler{}
}

type handler struct{}

func (h *handler) GetAdminTabs(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsAdmin.GetAdminTabs")
	defer span.End()
	// TODO: Implement
}

func (h *handler) GetInteractionLogs(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsAdmin.GetInteractionLogs")
	defer span.End()
	// TODO: Implement
}

func (h *handler) GetBlocks(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsAdmin.GetBlocks")
	defer span.End()
	// TODO: Implement
}

func (h *handler) ExportBlocks(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsAdmin.ExportBlocks")
	defer span.End()
	// TODO: Implement
}

func (h *handler) AddBlock(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsAdmin.AddBlock")
	defer span.End()
	// TODO: Implement
}

func (h *handler) UndoBlock(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsAdmin.UndoBlock")
	defer span.End()
	// TODO: Implement
}

func (h *handler) GetReportList(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsAdmin.GetReportList")
	defer span.End()
	// TODO: Implement
}
