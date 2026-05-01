package worker

import (
	"bytes"
	"context"
	"encoding/gob"
	"errors"
	"fmt"
	"glintfed/internal/data/client"
	"net/http"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/google/uuid"
)

var (
	ErrMissingRequiredParams error = errors.New("missing required field")
)

type InboxWorker struct {
	event *client.Event
}

func NewInboxWorker(event *client.Event) *InboxWorker {
	return &InboxWorker{
		event: event,
	}
}

type InboxParams struct {
	HTTPHeader http.Header
	Payload    InboxPayload
}

func (p InboxParams) Encode() ([]byte, error) {
	var buf bytes.Buffer
	err := gob.NewEncoder(&buf).Encode(p)

	return buf.Bytes(), err
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
	if err := validateHeader(params.HTTPHeader); err != nil {
		return err
	}

	encoded, err := params.Encode()
	if err != nil {
		return err
	}

	return svc.event.Publisher.Publish("inbox.delete",
		message.NewMessageWithContext(ctx, uuid.Must(uuid.NewV7()).String(), encoded),
	)
}

func (svc *InboxWorker) Inbox(ctx context.Context, params InboxParams) error {
	if err := validateHeader(params.HTTPHeader); err != nil {
		return err
	}

	encoded, err := params.Encode()
	if err != nil {
		return err
	}

	return svc.event.Publisher.Publish("inbox.inbox",
		message.NewMessageWithContext(ctx, uuid.Must(uuid.NewV7()).String(), encoded),
	)
}

func (svc *InboxWorker) Validate(ctx context.Context, username string, params InboxParams) error {
	if err := validateHeader(params.HTTPHeader); err != nil {
		return err
	}

	encoded, err := params.Encode()
	if err != nil {
		return err
	}

	msg := message.NewMessageWithContext(ctx, uuid.Must(uuid.NewV7()).String(), encoded)
	msg.Metadata.Set("username", username)

	return svc.event.Publisher.Publish("inbox.validate", msg)
}

func validateHeader(h http.Header) error {
	required := []string{"signature", "date"}

	for _, field := range required {
		if h.Get(field) == "" {
			return fmt.Errorf("%w: %s", ErrMissingRequiredParams, field)
		}
	}

	return nil
}
