package topic

import (
	"net/http"

	"glintfed/internal/server/handler/internal"
)

type Handler interface {
	GroupTopics(w http.ResponseWriter, r *http.Request)
	GroupTopicTag(w http.ResponseWriter, r *http.Request)
}

func New() Handler {
	return &handler{}
}

type handler struct{}

func (h *handler) GroupTopics(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsTopic.GroupTopics")
	defer span.End()
	// TODO: Implement
}

func (h *handler) GroupTopicTag(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "GroupsTopic.GroupTopicTag")
	defer span.End()
	// TODO: Implement
}
