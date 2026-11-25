package ghclient

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	ghct "github.com/bored-engineer/github-conditional-http-transport"
	ghctbbolt "github.com/bored-engineer/github-conditional-http-transport/bbolt"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"go.etcd.io/bbolt"
	bbolterrs "go.etcd.io/bbolt/errors"
)

// cacheLockTimeout is the maximum time to wait for a lock on the cache database before returning an error.
const cacheLockTimeout = time.Second

// createCacheStore creates a new [ghctbbolt.Storage] for caching GitHub API responses. It returns an error wrapping [ErrCacheInUse] if the database is locked by another provider configuration or process.
func createCacheStore(dirPath string) (*ghctbbolt.Storage, func() error, error) {
	if err := os.MkdirAll(dirPath, 0o700); err != nil {
		return nil, nil, fmt.Errorf("failed to create cache directory: %w", err)
	}

	dbPath := filepath.Join(dirPath, "cache.db")

	store, err := ghctbbolt.Open(dbPath, 0o600, &bbolt.Options{Timeout: cacheLockTimeout}, nil)
	if err != nil {
		if errors.Is(err, bbolterrs.ErrTimeout) {
			return nil, nil, fmt.Errorf("%w: %q is locked by another provider configuration or process; each provider configuration must use a unique cache path", ErrCacheInUse, dbPath)
		}

		return nil, nil, fmt.Errorf("failed to open cache storage: %w", err)
	}

	return store, store.DB.Close, nil
}

// bestEffortStorage is a [ghct.Storage] that treats the cache as optional, so a cache read failure is a cache miss and a cache write failure is logged rather than failing the request.
type bestEffortStorage struct {
	inner ghct.Storage
}

// Get implements [ghct.Storage] for the [bestEffortStorage]. It returns a cache miss if the inner storage errors, so a corrupt entry can be replaced rather than failing every request for the URL.
func (s bestEffortStorage) Get(ctx context.Context, req *http.Request) (*http.Response, error) {
	resp, err := s.inner.Get(ctx, req)
	if err != nil {
		if resp != nil && resp.Body != nil {
			_ = resp.Body.Close()
		}

		tflog.Warn(ctx, "Failed to read response from cache, treating as a cache miss.", map[string]any{"url": req.URL.String(), "error": err.Error()})

		return nil, nil
	}

	return resp, nil
}

// Put implements [ghct.Storage] for the [bestEffortStorage]. It buffers and closes the response body so the body can always be restored, and only returns an error if the body itself can't be read.
func (s bestEffortStorage) Put(ctx context.Context, resp *http.Response) error {
	var body []byte
	if resp.Body != nil {
		b, err := io.ReadAll(resp.Body)
		_ = resp.Body.Close()
		if err != nil {
			return fmt.Errorf("failed to read response body: %w", err)
		}

		body = b
	}

	resp.Body = io.NopCloser(bytes.NewReader(body))
	if err := s.inner.Put(ctx, resp); err != nil {
		tflog.Warn(ctx, "Failed to write response to cache.", map[string]any{"url": resp.Request.URL.String(), "error": err.Error()})
	}

	// Restore the body regardless of what the inner storage did with it.
	resp.Body = io.NopCloser(bytes.NewReader(body))
	resp.ContentLength = int64(len(body))

	return nil
}
