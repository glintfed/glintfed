package federation

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"glintfed/internal/lib/liblogs"
	"glintfed/internal/server/handler/internal"
	"glintfed/internal/service/worker"
	"io"
	"log/slog"
	"net/http"
	"net/url"

	"github.com/go-chi/chi/v5"
)

type inboxObject struct {
	ID     string       `json:"id"`
	Type   string       `json:"type"`
	Object *inboxObject `json:"object"`
}

func (h *handler) SharedInbox(w http.ResponseWriter, r *http.Request) {
	ctx, span := internal.T.Start(r.Context(), "Federation.SharedInbox")
	defer span.End()

	if !h.cfg.App.Federation.Activitypub.Enabled ||
		!h.cfg.App.Federation.Activitypub.SharedInbox {
		http.NotFound(w, r)
		return
	}

	var payload worker.InboxPayload
	raw, err := io.ReadAll(r.Body)
	if err != nil {
		slog.ErrorContext(r.Context(), "failed to read request body", liblogs.ErrAttr(err))
		internal.WriteError(w, err)
		return
	}
	payload.Raw = string(raw)

	if err := json.Unmarshal(raw, &payload); err != nil {
		slog.ErrorContext(r.Context(), "failed to decode json payload", liblogs.ErrAttr(err))
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	if err := h.validInboxDomain(ctx, payload.ID); err != nil {
		slog.ErrorContext(ctx, "invalid domain", liblogs.ErrAttr(err))
		internal.WriteError(w, err)
		return
	}

	if payload.Type != nil {
		switch *payload.Type {
		case "Delete":
			if payload.Object == nil {
				slog.ErrorContext(ctx, "invalid payload", liblogs.ErrAttr(err))
				http.Error(w, "", http.StatusBadRequest)
				return
			}

			if err := h.handleDelete(ctx, r.Header, payload); err != nil {
				slog.ErrorContext(ctx, "failed to handle delete", liblogs.ErrAttr(err), slog.Any("payload", payload))
				switch {
				case errors.Is(err, ErrMissingUrl):
					w.WriteHeader(http.StatusOK)
				case errors.Is(err, ErrInvalidType):
					w.WriteHeader(http.StatusBadRequest)
				default:
					w.WriteHeader(http.StatusInternalServerError)
				}
			} else {
				w.WriteHeader(http.StatusOK)
			}
			return
		case "Follow", "Accept":
			if err := h.inboxWorkerService.Inbox(ctx, worker.InboxParams{
				HTTPHeader: r.Header.Clone(),
				Payload:    payload,
			}); err != nil {
				slog.ErrorContext(r.Context(), "failed to handle inbox", liblogs.ErrAttr(err))
			}
		default:
			if err := h.inboxWorkerService.Inbox(ctx, worker.InboxParams{
				HTTPHeader: r.Header.Clone(),
				Payload:    payload,
			}); err != nil {
				slog.ErrorContext(r.Context(), "failed to handle inbox", liblogs.ErrAttr(err))
			}
		}
	} else {
		if err := h.inboxWorkerService.Inbox(ctx, worker.InboxParams{
			HTTPHeader: r.Header.Clone(),
			Payload:    payload,
		}); err != nil {
			slog.ErrorContext(r.Context(), "failed to handle inbox", liblogs.ErrAttr(err))
		}
	}

	w.WriteHeader(http.StatusOK)
}

func (h *handler) UserInbox(w http.ResponseWriter, r *http.Request) {
	ctx, span := internal.T.Start(r.Context(), "Federation.UserInbox")
	defer span.End()

	if !h.cfg.App.Federation.Activitypub.Enabled || !h.cfg.App.Federation.Activitypub.Inbox {
		http.NotFound(w, r)
		return
	}

	var payload worker.InboxPayload
	raw, err := io.ReadAll(r.Body)
	if err != nil {
		slog.ErrorContext(r.Context(), "failed to read request body", liblogs.ErrAttr(err))
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	payload.Raw = string(raw)

	if err := json.Unmarshal(raw, &payload); err != nil {
		slog.ErrorContext(r.Context(), "failed to decode json payload", liblogs.ErrAttr(err))
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	if err := h.validInboxDomain(ctx, payload.ID); err != nil {
		slog.ErrorContext(r.Context(), "invalid domain", liblogs.ErrAttr(err))
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	username := chi.URLParam(r, "username")

	if payload.Type != nil {
		switch *payload.Type {
		case "Delete":
			if payload.Object == nil {
				slog.ErrorContext(ctx, "invalid payload", liblogs.ErrAttr(err))
				http.Error(w, "", http.StatusBadRequest)
				return
			}

			if err := h.handleDelete(ctx, r.Header, payload); err != nil {
				slog.ErrorContext(ctx, "failed to handle delete", liblogs.ErrAttr(err), slog.Any("payload", payload))
				switch {
				case errors.Is(err, ErrMissingUrl):
					w.WriteHeader(http.StatusOK)
				case errors.Is(err, ErrInvalidType):
					w.WriteHeader(http.StatusBadRequest)
				default:
					w.WriteHeader(http.StatusInternalServerError)
				}
			} else {
				w.WriteHeader(http.StatusOK)
			}
			return
		case "Follow", "Accept":
			if err := h.inboxWorkerService.Validate(ctx, username, worker.InboxParams{
				HTTPHeader: r.Header.Clone(),
				Payload:    payload,
			}); err != nil {
				slog.ErrorContext(ctx, "failed to handle validate inbox", liblogs.ErrAttr(err))
			}
		default:
			if err := h.inboxWorkerService.Validate(ctx, username, worker.InboxParams{
				HTTPHeader: r.Header.Clone(),
				Payload:    payload,
			}); err != nil {
				slog.ErrorContext(ctx, "failed to handle validate inbox", liblogs.ErrAttr(err))
			}
		}
	} else {
		if err := h.inboxWorkerService.Validate(ctx, username, worker.InboxParams{
			HTTPHeader: r.Header.Clone(),
			Payload:    payload,
		}); err != nil {
			slog.ErrorContext(ctx, "failed to handle validate inbox", liblogs.ErrAttr(err))
		}
	}
}

func (h *handler) validInboxDomain(ctx context.Context, domain string) error {
	parsed, err := url.Parse(domain)
	if err != nil {
		return err
	}

	blocked, err := h.instanceService.GetBlockedDomains(ctx)
	if err != nil {
		return err
	}

	if _, ok := blocked[parsed.Host]; ok {
		return fmt.Errorf("host %s has been blocked", parsed.Host)
	}

	return nil
}

func (h *handler) handleDelete(ctx context.Context, header http.Header, payload worker.InboxPayload) error {
	switch payload.Object.Type {
	case "Person":
		if exists, err := h.profileModel.RemoteURLExists(ctx, payload.Object.ID); err != nil {
			return err
		} else if !exists {
			return ErrMissingUrl
		}
	case "Tombstone":
		if exists, err := h.statusModel.ObjectURLExists(ctx, payload.Object.ID); err != nil {
			return err
		} else if !exists {
			return ErrMissingUrl
		}
	case "Story":
	default:
		return ErrInvalidType
	}

	return h.inboxWorkerService.Delete(ctx, worker.InboxParams{
		HTTPHeader: header.Clone(),
		Payload:    payload,
	})
}
