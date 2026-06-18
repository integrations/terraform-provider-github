package ghclient

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	ghctbbolt "github.com/bored-engineer/github-conditional-http-transport/bbolt"
)

// createCacheStore creates a new bbolt storage for caching GitHub API responses. It creates a temporary directory and returns a storage instance pointing to a file in that directory.
func createCacheStore(path, ref string) (*ghctbbolt.Storage, error) {
	if path == "" {
		return nil, errors.New("cache path cannot be empty")
	}

	cacheDir := filepath.Join(path, "terraform-provider-github")

	if ref != "" {
		cacheDir = filepath.Join(cacheDir, ref)
	}

	if err := os.MkdirAll(cacheDir, 0o700); err != nil {
		return nil, fmt.Errorf("failed to create cache directory: %w", err)
	}

	store, err := ghctbbolt.Open(filepath.Join(cacheDir, "cache.db"), 0o600, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to open cache storage: %w", err)
	}

	return store, nil
}
