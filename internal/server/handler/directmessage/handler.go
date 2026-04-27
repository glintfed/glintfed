package directmessage

import (
	"net/http"

	"glintfed/internal/server/handler/internal"
)

type Handler interface {
	Thread(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Mute(w http.ResponseWriter, r *http.Request)
	Unmute(w http.ResponseWriter, r *http.Request)
	MediaUpload(w http.ResponseWriter, r *http.Request)
	Read(w http.ResponseWriter, r *http.Request)
	ComposeLookup(w http.ResponseWriter, r *http.Request)
	ComposeMutuals(w http.ResponseWriter, r *http.Request)
}

func New() Handler {
	return &handler{}
}

type handler struct{}

func (h *handler) Thread(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "DirectMessage.Thread")
	defer span.End()
	// TODO: Implement
}

func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "DirectMessage.Create")
	defer span.End()
	// TODO: Implement
}

func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "DirectMessage.Delete")
	defer span.End()
	// TODO: Implement
}

func (h *handler) Mute(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "DirectMessage.Mute")
	defer span.End()
	// TODO: Implement
}

func (h *handler) Unmute(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "DirectMessage.Unmute")
	defer span.End()
	// TODO: Implement
}

func (h *handler) MediaUpload(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "DirectMessage.MediaUpload")
	defer span.End()
	// TODO: Implement
}

func (h *handler) Read(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "DirectMessage.Read")
	defer span.End()
	// TODO: Implement
}

func (h *handler) ComposeLookup(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "DirectMessage.ComposeLookup")
	defer span.End()
	// TODO: Implement
}

func (h *handler) ComposeMutuals(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "DirectMessage.ComposeMutuals")
	defer span.End()
	// TODO: Implement
}
