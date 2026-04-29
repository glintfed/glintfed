package instance

import (
	"context"
	"time"

	"glintfed/ent"
	"glintfed/pkg/cache"

	"github.com/samber/lo"
)

const (
	bannedDomainsCacheKey = "instances:banned:domains"
	bannedDomainsCacheTTL = 1209600 * time.Second
)

type DomainManager struct {
	instanceModel InstanceModel
}

//go:generate go tool moq -rm -out mock_instance_model.go . InstanceModel
type InstanceModel interface {
	GetBlockedDomains(ctx context.Context) ([]*ent.Instance, error)
}

func NewDomainManager(instanceModel InstanceModel) *DomainManager {
	return &DomainManager{
		instanceModel: instanceModel,
	}
}

func (svc *DomainManager) GetBlockedDomains(ctx context.Context) (map[string]struct{}, error) {
	if blocked, ok := cachedBlockedDomains(ctx, bannedDomainsCacheKey); ok {
		return blocked, nil
	}

	var blocked map[string]struct{}
	var loadErr error
	err := cache.SetFunc(ctx, bannedDomainsCacheKey, func() any {
		blocked, loadErr = svc.loadBlockedDomains(ctx)
		if loadErr != nil {
			return loadErr
		}
		return blocked
	}, bannedDomainsCacheTTL)
	if loadErr != nil {
		return nil, loadErr
	}
	if err != nil {
		return nil, err
	}
	return blocked, nil
}

func (svc *DomainManager) loadBlockedDomains(ctx context.Context) (map[string]struct{}, error) {
	domains, err := svc.instanceModel.GetBlockedDomains(ctx)
	if err != nil {
		return nil, err
	}

	return lo.Keyify(lo.Map(domains, func(d *ent.Instance, _ int) string { return d.Domain })), nil
}

func cachedBlockedDomains(ctx context.Context, key string) (map[string]struct{}, bool) {
	val := cache.Get(ctx, key)
	switch v := val.(type) {
	case map[string]struct{}:
		return v, true
	case []string:
		return lo.Keyify(v), true
	default:
		return nil, false
	}
}
