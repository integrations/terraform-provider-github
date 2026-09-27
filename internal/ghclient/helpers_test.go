package ghclient

import (
	"net/http"
	"os"
	"path/filepath"
	"sync/atomic"
	"testing"
	"time"

	"golang.org/x/oauth2"
)

// sequentialTokenSource is an oauth2.TokenSource that returns a different token on each call,
// following the order of the provided tokens. Once exhausted, it keeps returning the last token.
// This is used to simulate multiple distinct callers/tokens sharing the same underlying transport.
type sequentialTokenSource struct {
	tokens []string
	idx    atomic.Int32
}

func (s *sequentialTokenSource) Token() (*oauth2.Token, error) {
	i := int(s.idx.Add(1)) - 1
	if i >= len(s.tokens) {
		i = len(s.tokens) - 1
	}

	return &oauth2.Token{AccessToken: s.tokens[i]}, nil
}

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
