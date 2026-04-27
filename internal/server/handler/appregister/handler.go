package appregister

import (
	"glintfed/internal/server/handler/internal"
	"net/http"
)

type Handler interface {
	VerifyCode(w http.ResponseWriter, r *http.Request)
	Onboarding(w http.ResponseWriter, r *http.Request)
}

func New() Handler {
	return &handler{}
}

type handler struct{}

func (h *handler) VerifyCode(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "AppRegister.VerifyCode")
	defer span.End()
	// TODO: Implement
}

func (h *handler) Onboarding(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "AppRegister.Onboarding")
	defer span.End()
	// TODO: Implement
}
