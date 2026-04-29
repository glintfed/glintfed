package federation

import (
	"context"
	"errors"
	"net/http"

	"glintfed/ent"
	"glintfed/internal/data"
	"glintfed/internal/service/instance"
	"glintfed/internal/service/worker"
)

var (
	ErrMissingUrl  = errors.New("missing url")
	ErrInvalidType = errors.New("invalid type")
)

type Handler interface {
	SharedInbox(w http.ResponseWriter, r *http.Request)
	UserInbox(w http.ResponseWriter, r *http.Request)
	Webfinger(w http.ResponseWriter, r *http.Request)
	NodeinfoWellKnown(w http.ResponseWriter, r *http.Request)
	HostMeta(w http.ResponseWriter, r *http.Request)
	Nodeinfo(w http.ResponseWriter, r *http.Request)
}

//go:generate go tool moq -rm -out mock_profile_model.go . ProfileModel
type ProfileModel interface {
	GetLocalByUsername(ctx context.Context, username string) (*ent.Profile, error)
	RemoteURLExists(ctx context.Context, remoteURL string) (bool, error)
}

//go:generate go tool moq -rm -out mock_status_model.go . StatusModel
type StatusModel interface {
	GetLocalPostsCount(ctx context.Context) (int, error)
	ObjectURLExists(ctx context.Context, objectURL string) (bool, error)
}

//go:generate go tool moq -rm -out mock_instance_service.go . InstanceService
type InstanceService interface {
	GetBlockedDomains(ctx context.Context) (map[string]struct{}, error)
	NodeinfoStats(ctx context.Context) (*instance.NodeinfoStats, error)
	NodeinfoFeatures(ctx context.Context) (*instance.NodeinfoFeatures, error)
}

//go:generate go tool moq -rm -out mock_inbox_worker_service.go . InboxWorkerService
type InboxWorkerService interface {
	Delete(ctx context.Context, params worker.InboxParams) error
	Inbox(ctx context.Context, params worker.InboxParams) error
	Validate(ctx context.Context, username string, params worker.InboxParams) error
}

func New(
	cfg *data.Config,

	profileModel ProfileModel,
	statusModel StatusModel,

	instanceService InstanceService,
	inboxWorkerService InboxWorkerService,
) Handler {
	return &handler{
		cfg: cfg,

		profileModel: profileModel,
		statusModel:  statusModel,

		instanceService:    instanceService,
		inboxWorkerService: inboxWorkerService,
	}
}

type handler struct {
	cfg *data.Config

	profileModel ProfileModel
	statusModel  StatusModel

	instanceService    InstanceService
	inboxWorkerService InboxWorkerService
}

func (h *handler) appURL(path string) string {
	return h.cfg.App.URL.JoinPath(path).String()
}
