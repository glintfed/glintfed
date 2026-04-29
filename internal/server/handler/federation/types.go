package federation

import (
	"glintfed/ent"
	"glintfed/internal/lib/libstr"
	"glintfed/internal/service/instance"
)

type WebfingerResponse struct {
	Subject string          `json:"subject"`
	Aliases []string        `json:"aliases"`
	Links   []WebfingerLink `json:"links"`
}

type WebfingerLink struct {
	Rel      string `json:"rel"`
	Type     string `json:"type,omitempty"`
	Href     string `json:"href,omitempty"`
	Template string `json:"template,omitempty"`
}

type NodeinfoWellKnownResponse struct {
	Links []NodeinfoLink `json:"links"`
}

type NodeinfoLink struct {
	Href string `json:"href"`
	Rel  string `json:"rel"`
}

type NodeinfoResponse struct {
	Metadata          NodeinfoMetadata       `json:"metadata"`
	Protocols         []string               `json:"protocols"`
	Services          NodeinfoServices       `json:"services"`
	Software          NodeinfoSoftware       `json:"software"`
	Usage             instance.NodeinfoUsage `json:"usage"`
	Version           string                 `json:"version"`
	OpenRegistrations bool                   `json:"openRegistrations"`
}

type NodeinfoMetadata struct {
	NodeName string                   `json:"nodeName"`
	Software NodeinfoMetadataSoftware `json:"software"`
	Config   NodeinfoConfig           `json:"config"`
}

type NodeinfoMetadataSoftware struct {
	Homepage string `json:"homepage"`
	Repo     string `json:"repo"`
}

type NodeinfoConfig struct {
	Features instance.NodeinfoFeatures `json:"features"`
}

type NodeinfoServices struct {
	Inbound  []string `json:"inbound"`
	Outbound []string `json:"outbound"`
}

type NodeinfoSoftware struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func instanceActorWebfinger(domain string) WebfingerResponse {
	return WebfingerResponse{
		Subject: "acct:" + domain + "@" + domain,
		Aliases: []string{"https://" + domain + "/i/actor"},
		Links: []WebfingerLink{
			{
				Rel:  "http://webfinger.net/rel/profile-page",
				Type: "text/html",
				Href: "https://" + domain + "/site/kb/instance-actor",
			},
			{
				Rel:  "self",
				Type: "application/activity+json",
				Href: "https://" + domain + "/i/actor",
			},
			{
				Rel:      "http://ostatus.org/schema/1.0/subscribe",
				Template: "https://" + domain + "/authorize_interaction?uri={uri}",
			},
		},
	}
}

func profileWebfinger(baseURL string, domain string, profile *ent.Profile) WebfingerResponse {
	avatarURL := libstr.FromPtr(profile.AvatarURL)
	if avatarURL == "" {
		avatarURL = baseURL + "/storage/avatars/default.jpg"
	}
	permalink := profile.Permalink(baseURL)
	profileURL := profile.Url(baseURL)
	username := libstr.FromPtr(profile.Username)

	return WebfingerResponse{
		Subject: "acct:" + username + "@" + domain,
		Aliases: []string{profileURL, permalink},
		Links: []WebfingerLink{
			{
				Rel:  "http://webfinger.net/rel/profile-page",
				Type: "text/html",
				Href: profileURL,
			},
			{
				Rel:  "http://schemas.google.com/g/2010#updates-from",
				Type: "application/atom+xml",
				Href: permalink + ".atom",
			},
			{
				Rel:  "self",
				Type: "application/activity+json",
				Href: permalink,
			},
			{
				Rel:  "http://webfinger.net/rel/avatar",
				Type: avatarMimeType(avatarURL),
				Href: avatarURL,
			},
			{
				Rel:      "http://ostatus.org/schema/1.0/subscribe",
				Template: "https://" + domain + "/authorize_interaction?uri={uri}",
			},
		},
	}
}

func avatarMimeType(path string) string {
	if len(path) >= 4 && path[len(path)-4:] == ".png" {
		return "image/png"
	}
	if len(path) >= 4 && path[len(path)-4:] == ".gif" {
		return "image/gif"
	}
	if len(path) >= 4 && path[len(path)-4:] == ".svg" {
		return "image/svg"
	}
	if len(path) >= 5 && path[len(path)-5:] == ".jpeg" {
		return "image/jpeg"
	}
	if len(path) >= 4 && path[len(path)-4:] == ".jpg" {
		return "image/jpeg"
	}
	return "application/octet-stream"
}
