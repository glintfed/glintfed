package oauth

import (
	"glintfed/internal/server/handler/internal"
	"net/http"
)

// Service defines the OAuth2 HTTP handlers.
type Handler interface {
	Authorize(w http.ResponseWriter, r *http.Request)
	Token(w http.ResponseWriter, r *http.Request)
	Revoke(w http.ResponseWriter, r *http.Request)
}

type handler struct{}

// New creates a new OAuth service.
func New() Handler {
	return &handler{}
}

func (h *handler) Authorize(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "OAuth.Authorize")
	defer span.End()
	// TODO: Implement
}

func (h *handler) Token(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "OAuth.Token")
	defer span.End()
	// TODO: Implement
}

func (h *handler) Revoke(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "OAuth.Revoke")
	defer span.End()
	// TODO: Implement
}
