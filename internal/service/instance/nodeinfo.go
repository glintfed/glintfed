package instance

import (
	"context"
	"strings"
	"time"

	"glintfed/internal/data"
	"glintfed/pkg/cache"
)

const (
	nodeinfoFeaturesCacheKey            = "api:nodeinfo:features"
	nodeinfoFeaturesCacheTTL            = 900 * time.Second
	nodeinfoUsersCacheKey               = "api:nodeinfo:users"
	nodeinfoActiveUsersMonthlyCacheKey  = "api:nodeinfo:active-users-monthly"
	nodeinfoActiveUsersHalfYearCacheKey = "api:nodeinfo:active-users-half-year"
	nodeinfoUsersCacheTTL               = 43200 * time.Second
	nodeinfoLocalPostsCacheKey          = "pf:services:instances:self:total-posts"
	nodeinfoLocalPostsCacheTTL          = time.Hour
)

type Nodeinfo struct {
	cfg *data.Config

	userModel   NodeinfoUserModel
	statusModel NodeinfoStatusModel

	now func() time.Time
}

//go:generate go tool moq -rm -out mock_nodeinfo_user_model.go . NodeinfoUserModel
type NodeinfoUserModel interface {
	CountAll(ctx context.Context) (int, error)
	CountActiveSince(ctx context.Context, since time.Time) (int, error)
}

//go:generate go tool moq -rm -out mock_nodeinfo_status_model.go . NodeinfoStatusModel
type NodeinfoStatusModel interface {
	GetLocalPostsCount(ctx context.Context) (int, error)
}

func NewNodeInfo(cfg *data.Config, userModel NodeinfoUserModel, statusModel NodeinfoStatusModel) *Nodeinfo {
	return &Nodeinfo{
		cfg: cfg,

		userModel:   userModel,
		statusModel: statusModel,

		now: time.Now,
	}
}

type NodeinfoFeatures struct {
	ActivityPub        bool                   `json:"activitypub,omitempty"`
	Timelines          NodeinfoTimelines      `json:"timelines"`
	MobileAPIs         bool                   `json:"mobile_apis"`
	MobileRegistration bool                   `json:"mobile_registration"`
	Stories            bool                   `json:"stories"`
	Video              bool                   `json:"video"`
	Import             NodeinfoImportFeatures `json:"import"`
	Label              NodeinfoLabelFeatures  `json:"label"`
	Groups             bool                   `json:"groups"`
}

type NodeinfoTimelines struct {
	Local   bool `json:"local"`
	Network bool `json:"network"`
}

type NodeinfoImportFeatures struct {
	Instagram bool `json:"instagram"`
	Mastodon  bool `json:"mastodon"`
	Pixelfed  bool `json:"pixelfed"`
}

type NodeinfoLabelFeatures struct {
	Covid NodeinfoCovidLabelFeature `json:"covid"`
}

type NodeinfoCovidLabelFeature struct {
	Enabled bool   `json:"enabled"`
	Org     string `json:"org"`
	URL     string `json:"url"`
}

type NodeinfoStats struct {
	Usage NodeinfoUsage
}

type NodeinfoUsage struct {
	LocalPosts    int               `json:"localPosts"`
	LocalComments int               `json:"localComments"`
	Users         NodeinfoUserUsage `json:"users"`
}

type NodeinfoUserUsage struct {
	Total          int `json:"total"`
	ActiveHalfyear int `json:"activeHalfyear"`
	ActiveMonth    int `json:"activeMonth"`
}

func (svc *Nodeinfo) NodeinfoStats(ctx context.Context) (*NodeinfoStats, error) {
	now := svc.now()

	activeHalfyear, err := cachedInt(ctx, nodeinfoActiveUsersHalfYearCacheKey, nodeinfoUsersCacheTTL, func() (int, error) {
		return svc.userModel.CountActiveSince(ctx, now.AddDate(0, -6, 0))
	})
	if err != nil {
		return nil, err
	}

	activeMonth, err := cachedInt(ctx, nodeinfoActiveUsersMonthlyCacheKey, nodeinfoUsersCacheTTL, func() (int, error) {
		return svc.userModel.CountActiveSince(ctx, now.AddDate(0, 0, -35))
	})
	if err != nil {
		return nil, err
	}

	users, err := cachedInt(ctx, nodeinfoUsersCacheKey, nodeinfoUsersCacheTTL, func() (int, error) {
		return svc.userModel.CountAll(ctx)
	})
	if err != nil {
		return nil, err
	}

	statuses, err := cachedInt(ctx, nodeinfoLocalPostsCacheKey, nodeinfoLocalPostsCacheTTL, func() (int, error) {
		return svc.statusModel.GetLocalPostsCount(ctx)
	})
	if err != nil {
		return nil, err
	}

	return &NodeinfoStats{
		Usage: NodeinfoUsage{
			LocalPosts:    statuses,
			LocalComments: 0,
			Users: NodeinfoUserUsage{
				Total:          users,
				ActiveHalfyear: activeHalfyear,
				ActiveMonth:    activeMonth,
			},
		},
	}, nil
}

func (svc *Nodeinfo) NodeinfoFeatures(ctx context.Context) (*NodeinfoFeatures, error) {
	if features, ok := cachedNodeinfoFeatures(ctx, nodeinfoFeaturesCacheKey); ok {
		return features, nil
	}

	var features *NodeinfoFeatures
	err := cache.SetFunc(ctx, nodeinfoFeaturesCacheKey, func() any {
		features = svc.nodeinfoFeatures()
		return features
	}, nodeinfoFeaturesCacheTTL)
	if err != nil {
		return nil, err
	}
	return features, nil
}

func (svc *Nodeinfo) nodeinfoFeatures() *NodeinfoFeatures {
	return &NodeinfoFeatures{
		ActivityPub:        svc.cfg.App.Federation.Activitypub.Enabled,
		Timelines:          svc.nodeinfoTimelines(),
		MobileAPIs:         svc.cfg.App.Auth.EnableOAuth,
		MobileRegistration: svc.cfg.App.Auth.InAppRegistration,
		Stories:            svc.cfg.App.Instance.Stories.Enabled,
		Video:              strings.Contains(svc.cfg.App.MediaTypes, "video/mp4"),
		Import:             svc.nodeinfoImportFeatures(),
		Label:              svc.nodeinfoLabelFeatures(),
		Groups:             svc.cfg.App.Groups.Enabled,
	}
}

func (svc *Nodeinfo) nodeinfoTimelines() NodeinfoTimelines {
	return NodeinfoTimelines{
		Local:   true,
		Network: svc.cfg.App.Federation.NetworkTimeline,
	}
}

func (svc *Nodeinfo) nodeinfoImportFeatures() NodeinfoImportFeatures {
	return NodeinfoImportFeatures{
		Instagram: svc.cfg.App.Import.Instagram.Enabled,
		Mastodon:  false,
		Pixelfed:  false,
	}
}

func (svc *Nodeinfo) nodeinfoLabelFeatures() NodeinfoLabelFeatures {
	return NodeinfoLabelFeatures{
		Covid: NodeinfoCovidLabelFeature{
			Enabled: svc.cfg.App.Instance.Label.Covid.Enabled,
			Org:     svc.cfg.App.Instance.Label.Covid.Org,
			URL:     svc.covidLabelURL(),
		},
	}
}

func (svc *Nodeinfo) covidLabelURL() string {
	if svc.cfg.App.Instance.Label.Covid.URL != nil {
		return svc.cfg.App.Instance.Label.Covid.URL.String()
	}
	return svc.cfg.App.Instance.Label.Covid.URLValue
}

func cachedInt(ctx context.Context, key string, ttl time.Duration, load func() (int, error)) (int, error) {
	if val, ok := cachedIntValue(ctx, key); ok {
		return val, nil
	}

	var res int
	var loadErr error
	err := cache.SetFunc(ctx, key, func() any {
		res, loadErr = load()
		if loadErr != nil {
			return loadErr
		}
		return res
	}, ttl)
	if loadErr != nil {
		return 0, loadErr
	}
	if err != nil {
		return 0, err
	}
	return res, nil
}

func cachedIntValue(ctx context.Context, key string) (int, bool) {
	val := cache.Get(ctx, key)
	switch v := val.(type) {
	case int:
		return v, true
	case int64:
		return int(v), true
	case float64:
		return int(v), true
	default:
		return 0, false
	}
}

func cachedNodeinfoFeatures(ctx context.Context, key string) (*NodeinfoFeatures, bool) {
	val := cache.Get(ctx, key)
	switch v := val.(type) {
	case *NodeinfoFeatures:
		return v, v != nil
	case NodeinfoFeatures:
		return &v, true
	default:
		return nil, false
	}
}
