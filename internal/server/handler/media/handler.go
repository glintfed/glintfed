package media

import (
	"context"
	"net/http"
	"path"
	"strconv"

	"glintfed/internal/data"
	"glintfed/internal/server/handler/internal"

	"github.com/go-chi/chi/v5"
)

type Handler interface {
	FallbackRedirect(w http.ResponseWriter, r *http.Request)
}

//go:generate go tool moq -rm -out mock_media_model.go . MediaModel
type MediaModel interface {
	CDNURL(ctx context.Context, profileID uint64, path string) (*string, error)
}

func New(cfg *data.Config, mediaModel MediaModel) Handler {
	return &handler{
		cfg: cfg,

		mediaModel: mediaModel,
	}
}

type handler struct {
	cfg *data.Config

	mediaModel MediaModel
}

func (h *handler) FallbackRedirect(w http.ResponseWriter, r *http.Request) {
	ctx, span := internal.T.Start(r.Context(), "Media.FallbackRedirect")
	defer span.End()

	if !h.cfg.App.CloudStorage {
		http.Redirect(w, r, "/storage/no-preview.png", http.StatusFound)
		return
	}

	profileID, err := strconv.ParseUint(chi.URLParam(r, "pid"), 10, 64)
	if err != nil {
		http.Redirect(w, r, "/storage/no-preview.png", http.StatusFound)
		return
	}

	cdnURL, err := h.mediaModel.CDNURL(ctx, profileID, path.Join("public/m/_v2", chi.URLParam(r, "pid"), chi.URLParam(r, "mhash"), chi.URLParam(r, "uhash"), chi.URLParam(r, "f")))
	if err != nil {
		internal.WriteError(w, err)
		return
	}
	if cdnURL == nil {
		http.Redirect(w, r, "/storage/no-preview.png", http.StatusFound)
		return
	}

	http.Redirect(w, r, *cdnURL, http.StatusFound)
}
