package federation

import (
	"glintfed/internal/lib/liblogs"
	"glintfed/internal/server/handler/internal"
	"log/slog"
	"net/http"
)

func (h *handler) NodeinfoWellKnown(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Federation.NodeinfoWellKnown")
	defer span.End()

	if !h.cfg.App.Federation.NodeInfo.Enabled {
		http.NotFound(w, r)
		return
	}

	internal.WriteJSONWithCORS(w, http.StatusOK, NodeinfoWellKnownResponse{
		Links: []NodeinfoLink{{
			Href: h.appURL("/api/nodeinfo/2.0.json"),
			Rel:  "http://nodeinfo.diaspora.software/ns/schema/2.0",
		}},
	})
}

func (h *handler) Nodeinfo(w http.ResponseWriter, r *http.Request) {
	ctx, span := internal.T.Start(r.Context(), "Federation.Nodeinfo")
	defer span.End()

	if !h.cfg.App.Federation.NodeInfo.Enabled {
		http.NotFound(w, r)
		return
	}

	stats, err := h.instanceService.NodeinfoStats(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "failed to get nodeinfo stats", liblogs.ErrAttr(err))
		internal.WriteError(w, err)
		return
	}

	features, err := h.instanceService.NodeinfoFeatures(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "failed to get nodeinfo features", liblogs.ErrAttr(err))
		internal.WriteError(w, err)
		return
	}

	internal.WriteJSONWithCORS(w, http.StatusOK, NodeinfoResponse{
		Metadata: NodeinfoMetadata{
			NodeName: h.cfg.App.Name,
			Software: NodeinfoMetadataSoftware{
				Homepage: "https://pixelfed.org",
				Repo:     "https://github.com/pixelfed/pixelfed",
			},
			Config: NodeinfoConfig{Features: *features},
		},
		Protocols:         []string{"activitypub"},
		Services:          NodeinfoServices{Inbound: []string{}, Outbound: []string{}},
		Software:          NodeinfoSoftware{Name: "pixelfed", Version: h.cfg.App.Version},
		Usage:             stats.Usage,
		Version:           "2.0",
		OpenRegistrations: h.cfg.App.Auth.EnableRegistration,
	})
}
