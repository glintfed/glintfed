package model

import (
	"context"
	"glintfed/ent"
	"glintfed/ent/media"
	"glintfed/internal/data/client"
)

type Media struct {
	*ent.MediaClient
}

func NewMedia(client *client.Database) *Media {
	return &Media{
		MediaClient: client.Ent.Media,
	}
}

func (m *Media) CDNURL(ctx context.Context, profileID uint64, path string) (*string, error) {
	res, err := m.Query().
		Select("cdn_url").
		Where(
			media.ProfileID(profileID),
			media.MediaPath(path),
			media.CdnURLNotNil(),
		).First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	}

	return res.CdnURL, nil
}
