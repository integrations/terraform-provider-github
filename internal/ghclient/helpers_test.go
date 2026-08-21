package ghclient

import (
	"net/http"
	"os"
	"path/filepath"
	"sync/atomic"
	"testing"
	"time"
)

type testRoundTripper struct {
	delay  time.Duration
	called atomic.Int32
	resp   *http.Response
	err    error
}

func (tr *testRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if tr.delay > 0 {
		ctx := req.Context()
		select {
		case <-time.After(tr.delay):
		case <-ctx.Done():
		}
	}

	tr.called.Add(1)

	if tr.err != nil {
		return nil, tr.err
	}

	return tr.resp, nil
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

func mustCreateRequest(t *testing.T, method, url string) *http.Request {
	t.Helper()

	req, err := http.NewRequestWithContext(t.Context(), method, url, nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	return req
}
