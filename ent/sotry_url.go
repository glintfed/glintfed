package ent

import (
	"context"
	"glintfed/internal/lib/libstr"
	"log/slog"
	"net/url"
	"strconv"
)

// Url returns the URL of the story.
func (s *Story) URL(baseUrl string) string {
	s.mustProfile()

	res, err := url.JoinPath(baseUrl, "stories", libstr.FromPtr(s.Edges.Profile.Username), strconv.FormatUint(s.ID, 10))
	if err != nil {
		slog.Error("failed to join path",
			slog.String("baseUrl", baseUrl),
			slog.String("profile_username", libstr.FromPtr(s.Edges.Profile.Username)),
			slog.Uint64("story_id", s.ID),
		)
	}

	return res
}

// Permalink returns the permalink of the story.
func (s *Story) Permalink(baseUrl string) string {
	s.mustProfile()

	res, err := url.JoinPath(baseUrl, "stories", libstr.FromPtr(s.Edges.Profile.Username), strconv.FormatUint(s.ID, 10), "activity")
	if err != nil {
		slog.Error("failed to join path",
			slog.String("baseUrl", baseUrl),
			slog.String("profile_username", libstr.FromPtr(s.Edges.Profile.Username)),
			slog.Uint64("story_id", s.ID),
		)
	}

	return res
}

func (s *Story) MediaLink(baseUrl string) string {
	if s.MediaURL != nil {
		return *s.MediaURL
	}

	res, err := url.JoinPath(baseUrl, "storage", libstr.FromPtr(s.Path)) // TODO: storage url
	if err != nil {
		slog.Error("failed to join path",
			slog.String("baseUrl", baseUrl),
			slog.String("story_path", libstr.FromPtr(s.Path)),
		)
	}

	return res
}

func (s *Story) mustProfile() {
	if s.Edges.Profile != nil {
		return
	}

	slog.Warn("missing profile edge", slog.Any("story", s))
	s.Edges.Profile = s.QueryProfile().FirstX(context.Background())
}
