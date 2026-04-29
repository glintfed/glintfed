package oauth

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"glintfed/internal/data"
	"glintfed/internal/data/client"
	"glintfed/internal/lib/fositestore"
)

func TestUsecase_CreateTokens(t *testing.T) {
	client, err := client.NewTestDatabase(t)
	require.NoError(t, err)

	cfg := &data.Config{}
	cfg.App.Auth.OAuth.HMACSecret = "test-hmac-secret-32-bytes-long!!"
	cfg.App.Auth.OAuth.PersonalClientID = "1"
	cfg.App.Auth.OAuth.AccessTokenLifespan = 30 * 24 * time.Hour

	store := fositestore.New(client, cfg)

	// Create the personal access client that CreateTokens will look up.
	_, err = client.Ent.OauthClient.Create().
		SetID(1).
		SetName("Personal Access Client").
		SetSecret("personal-secret").
		SetRedirect("").
		SetPersonalAccessClient(true).
		SetPasswordClient(false).
		SetRevoked(false).
		Save(context.Background())
	require.NoError(t, err)

	repo := NewRefreshToken(cfg, store)

	t.Run("success", func(t *testing.T) {
		result, err := repo.Create(context.Background(), RefreshTokenCreateParams{UserID: 42, Scopes: []string{"read", "write"}})
		require.NoError(t, err)
		assert.NotEmpty(t, result.AccessToken)
		assert.NotEmpty(t, result.RefreshToken)
		assert.Equal(t, "1", result.ClientID)
		assert.Equal(t, int64((30 * 24 * time.Hour).Seconds()), result.ExpiresIn)
		assert.Equal(t, "personal-secret", result.ClientSecret)
	})

	t.Run("different users get distinct tokens", func(t *testing.T) {
		r1, err := repo.Create(context.Background(), RefreshTokenCreateParams{UserID: 10, Scopes: []string{"read"}})
		require.NoError(t, err)
		r2, err := repo.Create(context.Background(), RefreshTokenCreateParams{UserID: 20, Scopes: []string{"read"}})
		require.NoError(t, err)
		assert.NotEqual(t, r1.AccessToken, r2.AccessToken)
		assert.NotEqual(t, r1.RefreshToken, r2.RefreshToken)
	})

	t.Run("unknown personal client returns error", func(t *testing.T) {
		cfg2 := *cfg
		cfg2.App.Auth.OAuth.PersonalClientID = "999"
		repo := NewRefreshToken(&cfg2, store)

		_, err := repo.Create(context.Background(), RefreshTokenCreateParams{UserID: 42, Scopes: []string{"read"}})
		assert.Error(t, err)
	})
}

func TestUsecase_DefaultTTL(t *testing.T) {
	client, err := client.NewTestDatabase(t)
	require.NoError(t, err)

	// Zero AccessTokenLifespan should fall back to 365 days.
	cfg := &data.Config{}
	cfg.App.Auth.OAuth.HMACSecret = "test-hmac-secret-32-bytes-long!!"
	cfg.App.Auth.OAuth.PersonalClientID = "1"
	cfg.App.Auth.OAuth.AccessTokenLifespan = 0

	store := fositestore.New(client, cfg)
	repo := NewRefreshToken(cfg, store)

	assert.Equal(t, 365*24*time.Hour, repo.tokenTTL)
}
