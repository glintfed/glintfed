package federation

import (
	"glintfed/internal/server/handler/internal"
	"net/http"
)

func (h *handler) HostMeta(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Federation.HostMeta")
	defer span.End()

	if !h.cfg.App.Federation.Webfinger.Enabled {
		http.NotFound(w, r)
		return
	}

	path := h.appURL("/.well-known/webfinger")
	xml := `<?xml version="1.0" encoding="UTF-8"?><XRD xmlns="http://docs.oasis-open.org/ns/xri/xrd-1.0"><Link rel="lrdd" type="application/xrd+xml" template="` + path + `?resource={uri}"/></XRD>`
	w.Header().Set("Content-Type", "application/xrd+xml")
	_, _ = w.Write([]byte(xml))
}
