package instance

import (
	"context"
	"errors"
	"net/url"
	"testing"
	"time"

	"glintfed/internal/data"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNodeinfo_NodeinfoStats(t *testing.T) {
	setupTestCache(t)

	ctx := context.Background()
	now := time.Date(2026, 4, 29, 12, 0, 0, 0, time.UTC)

	userModel := &NodeinfoUserModelMock{
		CountAllFunc: func(ctx context.Context) (int, error) {
			return 42, nil
		},
		CountActiveSinceFunc: func(ctx context.Context, since time.Time) (int, error) {
			switch since {
			case now.AddDate(0, -6, 0):
				return 21, nil
			case now.AddDate(0, 0, -35):
				return 7, nil
			default:
				t.Fatalf("unexpected active user cutoff: %s", since)
				return 0, nil
			}
		},
	}
	statusModel := &NodeinfoStatusModelMock{
		GetLocalPostsCountFunc: func(ctx context.Context) (int, error) {
			return 128, nil
		},
	}
	svc := NewNodeInfo(&data.Config{}, userModel, statusModel)
	svc.now = func() time.Time { return now }

	stats, err := svc.NodeinfoStats(ctx)
	require.NoError(t, err)

	assert.Equal(t, NodeinfoUsage{
		LocalPosts:    128,
		LocalComments: 0,
		Users: NodeinfoUserUsage{
			Total:          42,
			ActiveHalfyear: 21,
			ActiveMonth:    7,
		},
	}, stats.Usage)
	assert.Len(t, userModel.CountActiveSinceCalls(), 2)
	assert.Len(t, userModel.CountAllCalls(), 1)
	assert.Len(t, statusModel.GetLocalPostsCountCalls(), 1)
}

func TestNodeinfo_NodeinfoStatsReturnsError(t *testing.T) {
	setupTestCache(t)

	ctx := context.Background()
	wantErr := errors.New("query failed")
	userModel := &NodeinfoUserModelMock{
		CountActiveSinceFunc: func(ctx context.Context, since time.Time) (int, error) {
			return 0, wantErr
		},
	}
	statusModel := &NodeinfoStatusModelMock{}
	svc := NewNodeInfo(&data.Config{}, userModel, statusModel)

	stats, err := svc.NodeinfoStats(ctx)

	assert.Nil(t, stats)
	assert.ErrorIs(t, err, wantErr)
	assert.Empty(t, statusModel.GetLocalPostsCountCalls())
}

func TestNodeinfo_NodeinfoFeatures(t *testing.T) {
	setupTestCache(t)

	labelURL, err := url.Parse("https://example.test/label/covid")
	require.NoError(t, err)

	svc := NewNodeInfo(&data.Config{
		App: data.AppConfig{
			MediaTypes: "image/jpeg,video/mp4",
			Auth: data.AuthConfig{
				EnableOAuth:       true,
				InAppRegistration: true,
			},
			Instance: data.InstanceConfig{
				Stories: data.StoriesConfig{Enabled: true},
				Label: data.LabelConfig{
					Covid: data.LabelContentConfig{
						Enabled: true,
						Org:     "Health Org",
						URL:     labelURL,
					},
				},
			},
			Federation: data.FederationConfig{
				NetworkTimeline: true,
				Activitypub:     data.ActivitypubConfig{Enabled: true},
			},
			Import: data.ImportConfig{
				Instagram: data.InstagramConfig{Enabled: true},
			},
			Groups: data.GroupsConfig{
				Enabled: true,
			},
		},
	}, &NodeinfoUserModelMock{}, &NodeinfoStatusModelMock{})

	features, err := svc.NodeinfoFeatures(context.Background())
	require.NoError(t, err)

	assert.True(t, features.ActivityPub)
	assert.Equal(t, NodeinfoTimelines{Local: true, Network: true}, features.Timelines)
	assert.True(t, features.MobileAPIs)
	assert.True(t, features.MobileRegistration)
	assert.True(t, features.Stories)
	assert.True(t, features.Video)
	assert.Equal(t, NodeinfoImportFeatures{Instagram: true}, features.Import)
	assert.Equal(t, NodeinfoLabelFeatures{
		Covid: NodeinfoCovidLabelFeature{
			Enabled: true,
			Org:     "Health Org",
			URL:     "https://example.test/label/covid",
		},
	}, features.Label)
	assert.True(t, features.Groups)
}

func TestNodeinfo_NodeinfoStatsUsesCache(t *testing.T) {
	setupTestCache(t)

	ctx := context.Background()
	now := time.Date(2026, 4, 29, 12, 0, 0, 0, time.UTC)
	userModel := &NodeinfoUserModelMock{
		CountAllFunc: func(ctx context.Context) (int, error) {
			return 42, nil
		},
		CountActiveSinceFunc: func(ctx context.Context, since time.Time) (int, error) {
			switch since {
			case now.AddDate(0, -6, 0):
				return 21, nil
			case now.AddDate(0, 0, -35):
				return 7, nil
			default:
				t.Fatalf("unexpected active user cutoff: %s", since)
				return 0, nil
			}
		},
	}
	statusModel := &NodeinfoStatusModelMock{
		GetLocalPostsCountFunc: func(ctx context.Context) (int, error) {
			return 128, nil
		},
	}
	svc := NewNodeInfo(&data.Config{}, userModel, statusModel)
	svc.now = func() time.Time { return now }

	first, err := svc.NodeinfoStats(ctx)
	require.NoError(t, err)
	second, err := svc.NodeinfoStats(ctx)
	require.NoError(t, err)

	assert.Equal(t, first, second)
	assert.Len(t, userModel.CountActiveSinceCalls(), 2)
	assert.Len(t, userModel.CountAllCalls(), 1)
	assert.Len(t, statusModel.GetLocalPostsCountCalls(), 1)
}

func TestNodeinfo_NodeinfoFeaturesUsesCache(t *testing.T) {
	setupTestCache(t)

	cfg := &data.Config{
		App: data.AppConfig{
			Federation: data.FederationConfig{
				Activitypub: data.ActivitypubConfig{Enabled: true},
			},
		},
	}
	svc := NewNodeInfo(cfg, &NodeinfoUserModelMock{}, &NodeinfoStatusModelMock{})

	first, err := svc.NodeinfoFeatures(context.Background())
	require.NoError(t, err)

	cfg.App.Federation.Activitypub.Enabled = false
	second, err := svc.NodeinfoFeatures(context.Background())
	require.NoError(t, err)

	assert.Equal(t, first, second)
	assert.True(t, second.ActivityPub)
}
