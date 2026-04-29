package model

import (
	"context"
	"glintfed/ent"
	"glintfed/internal/data"
	"glintfed/internal/data/client"
)

type InstanceActor struct {
	*ent.InstanceActorClient
}

func NewInstanceActor(client *client.Database, cfg *data.Config) *InstanceActor {
	return &InstanceActor{
		InstanceActorClient: client.Ent.InstanceActor,
	}
}

func (m *InstanceActor) PublicKey(ctx context.Context) (*string, error) {
	ia, err := m.Query().First(ctx)
	if err != nil {
		return nil, err
	}

	return ia.PublicKey, nil
}
