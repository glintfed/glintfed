package ent

import (
	"glintfed/internal/lib/libstr"
	"log/slog"
	"net/url"
	"strings"
)

func (p *Profile) Url(baseUrl string, suffixes ...string) string {
	if p.RemoteURL != nil {
		return *p.RemoteURL
	}

	res, err := url.JoinPath(baseUrl, libstr.FromPtr(p.Username), strings.Join(suffixes, ""))
	if err != nil {
		slog.Error("failed to join path",
			slog.String("baseUrl", baseUrl),
			slog.String("profile_username", *p.Username),
			slog.Any("suffixes", suffixes),
		)
	}

	return res
}

func (p *Profile) Permalink(baseUrl string, suffixes ...string) string {
	if p.RemoteURL != nil {
		return *p.RemoteURL
	}

	res, err := url.JoinPath(baseUrl, "users", libstr.FromPtr(p.Username))
	if err != nil {
		slog.Error("failed to join path",
			slog.String("baseUrl", baseUrl),
			slog.String("profile_username", libstr.FromPtr(p.Username)),
			slog.Any("suffixes", suffixes),
		)
	}

	return res + strings.Join(suffixes, "")
}
