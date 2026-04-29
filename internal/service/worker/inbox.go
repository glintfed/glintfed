package worker

import (
	"context"
	"glintfed/internal/lib/liberrs"
	"net/http"
)

type InboxWorker struct{}

func NewInboxWorker() *InboxWorker {
	return &InboxWorker{}
}

type InboxParams struct {
	HTTPHeader http.Header
	Payload    InboxPayload
}

type InboxPayload struct {
	Raw   string
	ID    string  `json:"id"`
	Type  *string `json:"type,omitzero"`
	Actor *string `json:"actor,omitzero"`

	Object *InboxPayloadObject `json:"object,omitzero"`
}

type InboxPayloadObject struct {
	ID           string  `json:"id"`
	Type         string  `json:"type"`
	AttributedTo *string `json:"attributedTo"`
}

func (svc *InboxWorker) Delete(ctx context.Context, params InboxParams) error {
	return liberrs.Todo
}

func (svc *InboxWorker) Inbox(ctx context.Context, params InboxParams) error {
	return liberrs.Todo
}

func (svc *InboxWorker) Validate(ctx context.Context, username string, params InboxParams) error {
	return liberrs.Todo
}
