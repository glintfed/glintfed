package instance

import (
	"context"
	"testing"

	"glintfed/pkg/cache"

	"github.com/stretchr/testify/require"
)

func setupTestCache(t *testing.T) {
	t.Helper()

	cache.Register(cache.NewMemoryDriver())
	t.Cleanup(func() {
		require.NoError(t, cache.Clear(context.Background()))
	})
}
