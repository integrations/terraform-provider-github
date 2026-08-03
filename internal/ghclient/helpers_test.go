package ghclient

import (
	"net/http"
	"os"
	"path/filepath"
	"sync/atomic"
	"testing"
)

type testRoundTripper struct {
	called atomic.Int32
	resp   *http.Response
	err    error
}

func (r *testRoundTripper) RoundTrip(_ *http.Request) (*http.Response, error) {
	r.called.Add(1)
	if r.err != nil {
		return nil, r.err
	}

	return r.resp, nil
}

func mustMkdirTemp(t *testing.T, dir, pattern string) string {
	t.Helper()

	dir, err := os.MkdirTemp(dir, pattern) //nolint:usetesting
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}

	return dir
}

func mustReadAppPrivateKey(t *testing.T) []byte {
	t.Helper()

	privateKeyData, err := os.ReadFile(filepath.Join("..", "..", "github", "test-fixtures", "github-app-key.pem"))
	if err != nil {
		t.Fatalf("failed to read app private key fixture: %v", err)
	}

	return privateKeyData
}
