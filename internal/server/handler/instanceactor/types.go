package instanceactor

import "encoding/json"

type ActorResponse struct {
	Context                   json.RawMessage `json:"@context"`
	ID                        string          `json:"id"`
	Type                      string          `json:"type"`
	Inbox                     string          `json:"inbox"`
	Outbox                    string          `json:"outbox"`
	PreferredUsername         string          `json:"preferredUsername"`
	PublicKey                 PublicKey       `json:"publicKey"`
	ManuallyApprovesFollowers bool            `json:"manuallyApprovesFollowers"`
	URL                       string          `json:"url"`
}

type PublicKey struct {
	ID           string  `json:"id"`
	Owner        string  `json:"owner"`
	PublicKeyPem *string `json:"publicKeyPem"`
}

type OutboxResponse struct {
	Context    json.RawMessage `json:"@context"`
	ID         string          `json:"id"`
	Type       string          `json:"type"`
	TotalItems int             `json:"totalItems"`
	First      string          `json:"first"`
	Last       string          `json:"last"`
}
