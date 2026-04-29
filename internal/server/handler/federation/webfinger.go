package federation

import (
	"fmt"
	"glintfed/internal/data"
	"glintfed/internal/lib/liblogs"
	"glintfed/internal/lib/libstr"
	"glintfed/internal/server/handler/internal"
	"log/slog"
	"net/http"
	"strings"
)

func (h *handler) Webfinger(w http.ResponseWriter, r *http.Request) {
	ctx, span := internal.T.Start(r.Context(), "Federation.Webfinger")
	defer span.End()

	resource := r.URL.Query().Get("resource")
	if !h.cfg.App.Federation.Webfinger.Enabled || resource == "" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	u := h.cfg.App.URL
	if isShareboxResource(h.cfg, resource) {
		internal.WriteJSON(w, http.StatusOK, WebfingerResponse{
			Subject: resource,
			Aliases: []string{"https://" + u.Host + "/i/actor"},
			Links: []WebfingerLink{
				{
					Rel:  "http://webfinger.net/rel/profile-page",
					Type: "text/html",
					Href: "https://" + u.Host + "/site/kb/instance-actor",
				},
				{
					Rel:  "self",
					Type: "application/activity+json",
					Href: "https://" + u.Host + "/i/actor",
				},
				{
					Rel:      "http://ostatus.org/schema/1.0/subscribe",
					Template: "https://" + u.Host + "/authorize_interaction?uri={uri}",
				},
			},
		})
		return
	}

	var username string
	if after, ok := strings.CutPrefix(resource, "https://"+u.Host+"/users/"); ok {
		username = after
	} else if after, ok := strings.CutPrefix(resource, "acct:"); ok {
		parts := strings.Split(after, "@")
		if len(parts) != 2 || parts[1] != u.Host {
			slog.ErrorContext(ctx, "invalid acct resource", slog.String("resource", resource))
			http.Error(w, "", http.StatusBadRequest)
			return
		}
		username = parts[0]
	}

	if username == "" {
		slog.ErrorContext(r.Context(), "invalid resource format", slog.String("resource", resource))
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	if len(username) > h.cfg.App.MaxNameLength {
		slog.ErrorContext(ctx, "username too long", slog.String("username", username))
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	if !isValidUsername(username) {
		slog.ErrorContext(ctx, "invalid username", slog.String("username", username))
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	profile, err := h.profileModel.GetLocalByUsername(ctx, username)
	if err != nil {
		slog.ErrorContext(ctx, "failed to get profile by username", liblogs.ErrAttr(err))
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	internal.WriteJSON(w, http.StatusOK, WebfingerResponse{
		Subject: fmt.Sprintf("acct:%s@%s", username, h.cfg.App.URL.Host),
		Aliases: []string{
			profile.Url(h.cfg.App.URL.String()),
			profile.Permalink(h.cfg.App.URL.String()),
		},
		Links: []WebfingerLink{
			{
				Rel:  "http://webfinger.net/rel/profile-page",
				Type: "text/html",
				Href: profile.Url(h.cfg.App.URL.String()),
			},
			{
				Rel:  "http://schemas.google.com/g/2010#updates-from",
				Type: "application/atom+xml",
				Href: profile.Permalink(h.cfg.App.URL.String(), ".atom"),
			},
			{
				Rel:  "self",
				Type: "application/activity+json",
				Href: profile.Permalink(h.cfg.App.URL.String()),
			},
			{
				Rel:  "http://webfinger.net/rel/avatar",
				Type: "image/webp",
				Href: libstr.FromPtr(profile.AvatarURL),
			},
			{
				Rel:      "http://ostatus.org/schema/1.0/subscribe",
				Template: "https://" + h.cfg.App.URL.Host + "/authorize_interaction?uri={uri}",
			},
		},
	})
}

func isShareboxResource(cfg *data.Config, resource string) bool {
	var sb strings.Builder
	sb.WriteString("acct:")
	sb.WriteString(cfg.App.URL.Host)
	sb.WriteRune('@')
	sb.WriteString(cfg.App.URL.Host)

	return cfg.App.Federation.Activitypub.SharedInbox &&
		resource == sb.String()
}

func isValidUsername(username string) bool {
	for _, c := range username {
		if c == '_' || c == '.' || c == '-' {
			continue
		}
		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') {
			continue
		}

		return false
	}

	return true
}
