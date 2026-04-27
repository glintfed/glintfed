package admininvite

import (
	"net/http"

	"glintfed/internal/server/handler/internal"
)

type Handler interface {
	ApiVerifyCheck(w http.ResponseWriter, r *http.Request)
	ApiUsernameCheck(w http.ResponseWriter, r *http.Request)
	ApiEmailCheck(w http.ResponseWriter, r *http.Request)
}

func New() Handler {
	return &handler{}
}

type handler struct{}

func (h *handler) stub(w http.ResponseWriter, r *http.Request, name string) {

}

func (h *handler) ApiVerifyCheck(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "AdminInvite.ApiVerifyCheck")
	defer span.End()
	// TODO: Implement
}
func (h *handler) ApiUsernameCheck(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "AdminInvite.ApiUsernameCheck")
	defer span.End()
	// TODO: Implement
}
func (h *handler) ApiEmailCheck(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "AdminInvite.ApiEmailCheck")
	defer span.End()
	// TODO: Implement
}
