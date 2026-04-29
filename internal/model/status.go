package model

import (
	"context"

	"glintfed/ent"
	"glintfed/ent/status"
	"glintfed/internal/data/client"
)

type Status struct {
	*ent.StatusClient
}

func NewStatus(client *client.Database) *Status {
	return &Status{
		StatusClient: client.Ent.Status,
	}
}

// ObjectURLExists
//
//	SELECT EXISTS(
//	  SELECT 1 FROM statuses
//	  WHERE object_url = ?
//	)
func (m *Status) ObjectURLExists(ctx context.Context, objectURL string) (bool, error) {
	return m.Query().
		Where(status.ObjectURL(objectURL)).
		Exist(ctx)
}

// GetLocalPostCount
//
//	SELECT count(*)
//	FROM `statuses`
//	WHERE
//	  `deleted_at` IS NULL AND
//	  `local` = true
//	  `type` = "share"
func (m *Status) GetLocalPostsCount(ctx context.Context) (int, error) {
	return m.Query().
		Where(
			status.DeletedAtIsNil(),
			status.Local(true),
			status.TypeNEQ("share"),
		).
		Count(ctx)
}
