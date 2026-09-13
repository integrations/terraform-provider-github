package ghclient

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	ghctbbolt "github.com/bored-engineer/github-conditional-http-transport/bbolt"
)

// createCacheStore creates a new bbolt storage for caching GitHub API responses.
func createCacheStore(path string) (*ghctbbolt.Storage, error) {
	if path == "" {
		return nil, errors.New("cache path cannot be empty")
	}

	if err := os.MkdirAll(path, 0o700); err != nil {
		return nil, fmt.Errorf("failed to create cache directory: %w", err)
	}

	store, err := ghctbbolt.Open(filepath.Join(path, "cache.db"), 0o600, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to open cache storage: %w", err)
	}

	return store, nil
}
