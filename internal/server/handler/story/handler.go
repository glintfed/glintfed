package story

import (
	"context"
	"crypto/subtle"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
	"time"

	"glintfed/ent"
	"glintfed/internal/data"
	"glintfed/internal/lib/liblogs"
	"glintfed/internal/lib/libstr"
	"glintfed/internal/server/handler/internal"

	"github.com/go-chi/chi/v5"
)

type Handler interface {
	GetActivityObject(w http.ResponseWriter, r *http.Request)
}

//go:generate go tool moq -rm -out mock_profile_model.go . ProfileModel
type ProfileModel interface {
	GetStory(ctx context.Context, username string, storyID uint64) (*ent.Story, error)
}

func New(cfg *data.Config, profileModel ProfileModel) Handler {
	return &handler{
		cfg: cfg,

		profileModel: profileModel,
	}
}

type handler struct {
	cfg *data.Config

	profileModel ProfileModel
}

func (h *handler) GetActivityObject(w http.ResponseWriter, r *http.Request) {
	ctx, span := internal.T.Start(r.Context(), "Story.GetActivityObject")
	defer span.End()

	username := chi.URLParam(r, "username")
	storyID, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		internal.WriteError(w, err)
		return
	}

	if !h.cfg.App.Instance.Stories.Enabled {
		http.NotFound(w, r)
		return
	}

	if !wantsJSON(r) {
		http.Redirect(w, r, "/stories/"+username, http.StatusFound)
		return
	}

	auth := r.Header.Get("Authorization")
	if auth == "" {
		http.NotFound(w, r)
		return
	}

	story, err := h.profileModel.GetStory(ctx, username, storyID)
	if err != nil {
		slog.ErrorContext(ctx, "failed to get story", liblogs.ErrAttr(err))
		internal.WriteError(w, err)
		return
	}

	if !validStory(story, strings.TrimPrefix(auth, "Bearer ")) {
		http.NotFound(w, r)
		return
	}

	internal.WriteJSON(w, http.StatusOK, StoryActivityResponse{
		Context:      "https://www.w3.org/ns/activitystreams",
		ID:           story.URL(h.cfg.App.URL.String()),
		Type:         "Story",
		To:           []string{story.Edges.Profile.Permalink(h.cfg.App.URL.String(), "/followers")},
		CC:           []string{},
		AttributedTo: story.Edges.Profile.Permalink(h.cfg.App.URL.String()),
		Published:    story.CreatedAt.Format(time.RFC3339),
		ExpiresAt:    story.ExpiresAt.Format(time.RFC3339),
		Duration:     story.Duration,
		CanReply:     story.CanReply,
		CanReact:     story.CanReact,
		Attachment: StoryAttachment{
			Type:      storyAttachmentType(libstr.FromPtr(story.Type)),
			URL:       story.MediaLink(h.cfg.App.URL.String()),
			MediaType: libstr.FromPtr(story.Mime),
		},
	})
}

func wantsJSON(r *http.Request) bool {
	accept := r.Header.Get("Accept")
	return strings.Contains(accept, "application/json") ||
		strings.Contains(accept, "application/activity+json") ||
		strings.Contains(accept, "application/ld+json")
}

func storyAttachmentType(storyType string) string {
	switch storyType {
	case "photo":
		return "Image"
	case "video":
		return "Video"
	default:
		return "Document"
	}
}

func validStory(story *ent.Story, token string) bool {
	if story.ExpiresAt.Before(time.Now()) {
		return false
	}
	if story.CreatedAt.Before(time.Now().Add(-20 * time.Minute)) {
		return false
	}

	if story.BearcapToken == nil {
		return false
	}
	if subtle.ConstantTimeCompare([]byte(libstr.FromPtr(story.BearcapToken)), []byte(token)) != 1 {
		return false
	}

	return true
}
