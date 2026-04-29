package story

import "time"

type Profile struct {
	ID        uint64
	Permalink string
}

type Story struct {
	ID           string
	URL          string
	Type         string
	Duration     uint
	Mime         string
	MediaURL     string
	BearcapToken string
	CreatedAt    time.Time
	ExpiresAt    time.Time
	CanReply     bool
	CanReact     bool
}

type StoryActivityResponse struct {
	Context      string          `json:"@context"`
	ID           string          `json:"id"`
	Type         string          `json:"type"`
	To           []string        `json:"to"`
	CC           []string        `json:"cc"`
	AttributedTo string          `json:"attributedTo"`
	Published    string          `json:"published"`
	ExpiresAt    string          `json:"expiresAt"`
	Duration     uint            `json:"duration"`
	CanReply     bool            `json:"can_reply"`
	CanReact     bool            `json:"can_react"`
	Attachment   StoryAttachment `json:"attachment"`
}

type StoryAttachment struct {
	Type      string `json:"type"`
	URL       string `json:"url"`
	MediaType string `json:"mediaType"`
}
