package ghclient

import (
	"fmt"
	"os"
	"path/filepath"

	ghctbbolt "github.com/bored-engineer/github-conditional-http-transport/bbolt"
)

// createCacheStore creates a new [ghctbbolt.Storage] for caching GitHub API responses.
func createCacheStore(opts CacheOptions) (*ghctbbolt.Storage, error) {
	if opts.BasePath == "" {
		return nil, fmt.Errorf("cache path cannot be empty")
	}

	dirPath := filepath.Join(opts.BasePath, opts.Ref)

	if dirPath == "" {
		return nil, fmt.Errorf("cache path cannot be empty")
	}

	if err := os.MkdirAll(dirPath, 0o700); err != nil {
		return nil, fmt.Errorf("failed to create cache directory: %w", err)
	}

	store, err := ghctbbolt.Open(filepath.Join(dirPath, "cache.db"), 0o600, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to open cache storage: %w", err)
	}

	return store, nil
}
