package instance

import (
	"context"
	"errors"
	"glintfed/ent"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDomainManager_GetBlockedDomains(t *testing.T) {
	setupTestCache(t)

	ctx := context.Background()
	model := &InstanceModelMock{
		GetBlockedDomainsFunc: func(ctx context.Context) ([]*ent.Instance, error) {
			return []*ent.Instance{
				{Domain: "blocked.example"},
				{Domain: "spam.example"},
			}, nil
		},
	}
	svc := NewDomainManager(model)

	blocked, err := svc.GetBlockedDomains(ctx)
	require.NoError(t, err)

	assert.Contains(t, blocked, "blocked.example")
	assert.Contains(t, blocked, "spam.example")
	assert.NotContains(t, blocked, "allowed.example")
	assert.Len(t, model.GetBlockedDomainsCalls(), 1)
}

func TestDomainManager_GetBlockedDomainsReturnsError(t *testing.T) {
	setupTestCache(t)

	ctx := context.Background()
	wantErr := errors.New("query failed")
	svc := NewDomainManager(&InstanceModelMock{
		GetBlockedDomainsFunc: func(ctx context.Context) ([]*ent.Instance, error) {
			return nil, wantErr
		},
	})

	blocked, err := svc.GetBlockedDomains(ctx)

	assert.Nil(t, blocked)
	assert.ErrorIs(t, err, wantErr)
}

func TestDomainManager_GetBlockedDomainsUsesCache(t *testing.T) {
	setupTestCache(t)

	ctx := context.Background()
	model := &InstanceModelMock{
		GetBlockedDomainsFunc: func(ctx context.Context) ([]*ent.Instance, error) {
			return []*ent.Instance{
				{Domain: "blocked.example"},
			}, nil
		},
	}
	svc := NewDomainManager(model)

	first, err := svc.GetBlockedDomains(ctx)
	require.NoError(t, err)
	second, err := svc.GetBlockedDomains(ctx)
	require.NoError(t, err)

	assert.Equal(t, first, second)
	assert.Len(t, model.GetBlockedDomainsCalls(), 1)
}
