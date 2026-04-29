package model

import (
	"context"
	"glintfed/ent"
	"glintfed/ent/instance"
	"glintfed/internal/data/client"
)

type Instance struct {
	*ent.InstanceClient
}

func NewInstance(client *client.Database) *Instance {
	return &Instance{
		InstanceClient: client.Ent.Instance,
	}
}

// GetBlockedDomains
//
//	SELECT domain FROM instances WHERE banned = true
func (m *Instance) GetBlockedDomains(ctx context.Context) ([]*ent.Instance, error) {
	return m.Query().
		Where(instance.Banned(true)).
		Select(instance.FieldDomain).
		All(ctx)
}
