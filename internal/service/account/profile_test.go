package account

import (
	"context"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"glintfed/ent"
	"glintfed/internal/data"
	"glintfed/pkg/cache"
)

func TestProfile_GetProfileLocal(t *testing.T) {
	ctx := context.Background()
	createdAt := time.Date(2026, 4, 29, 12, 30, 0, 0, time.UTC)
	name := "Alice"
	bio := "<p>Hello <strong>world</strong></p>"
	p := profileFixture{
		id:             1,
		userID:         uint64Ptr(10),
		username:       "alice",
		name:           &name,
		bio:            &bio,
		privateKey:     stringPtr("secret"),
		followersCount: uintPtr(12),
		followingCount: uintPtr(8),
		statusCount:    uintPtr(3),
		createdAt:      &createdAt,
	}.profile()

	model := &ProfileModelMock{
		GetAccountProfileFunc: func(ctx context.Context, profileID uint64) (*ent.Profile, error) {
			assert.Equal(t, uint64(1), profileID)
			return p, nil
		},
		AccountHiddenCountsFunc: func(ctx context.Context, userID uint64) (bool, bool, error) {
			assert.Equal(t, uint64(10), userID)
			return false, true, nil
		},
		IsAdminAccountFunc: func(ctx context.Context, profileID uint64) (bool, error) {
			assert.Equal(t, uint64(1), profileID)
			return true, nil
		},
		AccountPronounsFunc: func(ctx context.Context, profileID uint64) (*ent.UserPronoun, error) {
			assert.Equal(t, uint64(1), profileID)
			return &ent.UserPronoun{Pronouns: stringPtr(`["she","they"]`)}, nil
		},
	}
	svc := newTestProfileService(t, model)

	res, err := svc.GetProfile(ctx, GetProfileParams{ProfileID: p.ID})
	require.NoError(t, err)

	assert.Equal(t, "1", res.ID)
	assert.Equal(t, "alice", res.Username)
	assert.Equal(t, "alice", res.Acct)
	assert.Equal(t, &name, res.DisplayName)
	assert.True(t, res.Discoverable)
	assert.True(t, res.Local)
	assert.True(t, res.IsAdmin)
	assert.Equal(t, 0, res.FollowersCount)
	assert.Equal(t, 8, res.FollowingCount)
	assert.Equal(t, 3, res.StatusesCount)
	assert.Equal(t, bio, res.Note)
	require.NotNil(t, res.NoteText)
	assert.Equal(t, "Hello world", *res.NoteText)
	assert.Equal(t, "https://example.com/alice", res.URL)
	assert.Equal(t, "https://example.com/storage/avatars/default.jpg", res.Avatar)
	assert.Equal(t, []string{"she", "they"}, res.Pronouns)
	require.NotNil(t, res.CreatedAt)
	assert.Equal(t, "2026-04-29T12:30:00Z", *res.CreatedAt)
	assert.Len(t, model.AccountHiddenCountsCalls(), 1)
	assert.Len(t, model.IsAdminAccountCalls(), 1)
}

func TestProfile_GetProfileUsesCache(t *testing.T) {
	ctx := context.Background()
	p := profileFixture{
		id:             5,
		username:       "cached",
		followersCount: uintPtr(1),
		followingCount: uintPtr(2),
		statusCount:    uintPtr(3),
	}.profile()
	model := &ProfileModelMock{
		GetAccountProfileFunc: func(ctx context.Context, profileID uint64) (*ent.Profile, error) {
			return p, nil
		},
		AccountPronounsFunc: func(ctx context.Context, profileID uint64) (*ent.UserPronoun, error) {
			return nil, new(ent.NotFoundError)
		},
	}
	svc := newTestProfileService(t, model)

	first, err := svc.GetProfile(ctx, GetProfileParams{ProfileID: p.ID})
	require.NoError(t, err)
	second, err := svc.GetProfile(ctx, GetProfileParams{ProfileID: p.ID})
	require.NoError(t, err)

	assert.Equal(t, first, second)
	assert.Len(t, model.GetAccountProfileCalls(), 1)
}

func TestProfile_GetProfileRemote(t *testing.T) {
	ctx := context.Background()
	remoteURL := "https://remote.example/@bob"
	p := profileFixture{
		id:             2,
		username:       "@bob@remote.example",
		remoteURL:      &remoteURL,
		followersCount: uintPtr(5),
		followingCount: uintPtr(6),
		statusCount:    uintPtr(7),
	}.profile()

	model := &ProfileModelMock{
		GetAccountProfileFunc: func(ctx context.Context, profileID uint64) (*ent.Profile, error) {
			assert.Equal(t, uint64(2), profileID)
			return p, nil
		},
		AccountPronounsFunc: func(ctx context.Context, profileID uint64) (*ent.UserPronoun, error) {
			assert.Equal(t, uint64(2), profileID)
			return nil, new(ent.NotFoundError)
		},
	}
	svc := newTestProfileService(t, model)

	res, err := svc.GetProfile(ctx, GetProfileParams{ProfileID: p.ID})
	require.NoError(t, err)

	assert.Equal(t, "bob", res.Username)
	assert.Equal(t, "bob@remote.example", res.Acct)
	assert.False(t, res.Local)
	assert.False(t, res.IsAdmin)
	assert.Equal(t, remoteURL, res.URL)
	assert.Equal(t, 5, res.FollowersCount)
	assert.Equal(t, 6, res.FollowingCount)
	assert.Equal(t, 7, res.StatusesCount)
	assert.Empty(t, res.Pronouns)
	assert.Empty(t, model.AccountHiddenCountsCalls())
	assert.Empty(t, model.IsAdminAccountCalls())
}

func TestProfile_GetProfileAvatarPath(t *testing.T) {
	ctx := context.Background()
	p := profileFixture{
		id:       3,
		username: "avatar",
		avatar: &ent.Avatar{
			MediaPath:   stringPtr("public/avatars/avatar.jpg"),
			ChangeCount: 2,
		},
	}.profile()

	model := &ProfileModelMock{
		GetAccountProfileFunc: func(ctx context.Context, profileID uint64) (*ent.Profile, error) {
			return p, nil
		},
		AccountPronounsFunc: func(ctx context.Context, profileID uint64) (*ent.UserPronoun, error) {
			return nil, new(ent.NotFoundError)
		},
	}
	svc := newTestProfileService(t, model)

	res, err := svc.GetProfile(ctx, GetProfileParams{ProfileID: p.ID})
	require.NoError(t, err)

	assert.Equal(t, "https://example.com/storage/avatars/avatar.jpg?v=2", res.Avatar)
}

func TestProfile_GetProfileNotFound(t *testing.T) {
	ctx := context.Background()
	svc := newTestProfileService(t, &ProfileModelMock{
		GetAccountProfileFunc: func(ctx context.Context, profileID uint64) (*ent.Profile, error) {
			return nil, new(ent.NotFoundError)
		},
	})

	_, err := svc.GetProfile(ctx, GetProfileParams{ProfileID: 999})
	assert.ErrorIs(t, err, ErrProfileNotFound)
}

func TestProfile_GetProfileDeleted(t *testing.T) {
	ctx := context.Background()
	status := "delete"
	p := profileFixture{id: 4, username: "deleted", status: &status}.profile()
	model := &ProfileModelMock{
		GetAccountProfileFunc: func(ctx context.Context, profileID uint64) (*ent.Profile, error) {
			return p, nil
		},
	}
	svc := newTestProfileService(t, model)

	_, err := svc.GetProfile(ctx, GetProfileParams{ProfileID: p.ID})
	assert.ErrorIs(t, err, ErrProfileNotFound)
	assert.Empty(t, model.AccountPronounsCalls())
	assert.Empty(t, model.AccountHiddenCountsCalls())
	assert.Empty(t, model.IsAdminAccountCalls())
}

type profileFixture struct {
	id             uint64
	userID         *uint64
	username       string
	status         *string
	name           *string
	bio            *string
	privateKey     *string
	remoteURL      *string
	followersCount *uint
	followingCount *uint
	statusCount    *uint
	createdAt      *time.Time
	avatar         *ent.Avatar
}

func (f profileFixture) profile() *ent.Profile {
	return &ent.Profile{
		ID:             f.id,
		UserID:         f.userID,
		Username:       &f.username,
		Status:         f.status,
		Name:           f.name,
		Bio:            f.bio,
		PrivateKey:     f.privateKey,
		RemoteURL:      f.remoteURL,
		FollowersCount: f.followersCount,
		FollowingCount: f.followingCount,
		StatusCount:    f.statusCount,
		CreatedAt:      f.createdAt,
		Edges: ent.ProfileEdges{
			Avatar: f.avatar,
		},
	}
}

func newTestProfileService(t *testing.T, model ProfileModel) *Profile {
	t.Helper()
	cache.Register(cache.NewMemoryDriver())

	u, err := url.Parse("https://example.com")
	require.NoError(t, err)

	cfg := &data.Config{}
	cfg.App.URL = u
	cfg.App.URLValue = u.String()
	return NewProfile(cfg, model)
}

func stringPtr(v string) *string {
	return &v
}

func uintPtr(v uint) *uint {
	return &v
}

func uint64Ptr(v uint64) *uint64 {
	return &v
}
