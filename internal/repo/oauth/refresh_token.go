package oauth

import (
	"context"
	"fmt"
	"glintfed/internal/data"
	"glintfed/internal/lib/fositestore"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/ory/fosite"
)

type RefreshToken struct {
	store    *fositestore.Store
	clientID string
	tokenTTL time.Duration
}

func NewRefreshToken(cfg *data.Config, store *fositestore.Store) *RefreshToken {
	ttl := cfg.App.Auth.OAuth.AccessTokenLifespan
	if ttl <= 0 {
		ttl = 365 * 24 * time.Hour
	}
	return &RefreshToken{
		store:    store,
		clientID: cfg.App.Auth.OAuth.PersonalClientID,
		tokenTTL: ttl,
	}
}

type RefreshTokenCreateParams struct {
	UserID uint64
	Scopes []string
}

func (repo *RefreshToken) Create(ctx context.Context, params RefreshTokenCreateParams) (*TokenResult, error) {
	subject := strconv.FormatUint(params.UserID, 10)

	client, err := repo.store.GetClient(ctx, repo.clientID)
	if err != nil {
		return nil, fmt.Errorf("get personal access client: %w", err)
	}

	now := time.Now()

	session := &fosite.DefaultSession{
		Subject:  subject,
		Username: subject,
		ExpiresAt: map[fosite.TokenType]time.Time{
			fosite.AccessToken:  now.Add(repo.tokenTTL),
			fosite.RefreshToken: now.Add(repo.tokenTTL + 35*24*time.Hour),
		},
	}

	req := fosite.NewRequest()
	req.ID = uuid.Must(uuid.NewV7()).String()
	req.Client = client
	req.RequestedAt = now
	req.Session = session
	req.SetRequestedScopes(fosite.Arguments(params.Scopes))
	for _, s := range params.Scopes {
		req.GrantScope(s)
	}

	accessToken, refreshToken, err := repo.store.CreatePersonalAccessTokens(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("create tokens: %w", err)
	}

	fositeClient, ok := client.(*fositestore.FositeClient)
	if !ok {
		return nil, fmt.Errorf("unexpected client type")
	}

	return &TokenResult{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ClientID:     repo.clientID,
		ClientSecret: string(fositeClient.GetHashedSecret()),
		ExpiresIn:    int64(repo.tokenTTL.Seconds()),
	}, nil
}
