package ghclient

import (
	"context"
	"crypto"
	"crypto/x509"
	"encoding/pem"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/google/go-cmp/cmp"
	"golang.org/x/oauth2"
)

var tokenSourceComparer = cmp.Comparer(func(a, b oauth2.TokenSource) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}

	ta, errA := a.Token()
	tb, errB := b.Token()
	if errA != nil || errB != nil {
		return false
	}

	return ta.AccessToken == tb.AccessToken
})

func mustReadAppSigner(t *testing.T) crypto.Signer {
	t.Helper()

	block, _ := pem.Decode(mustReadAppPrivateKey(t))
	if block == nil {
		t.Fatal("no decodeable PEM data found in app private key fixture")
	}

	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		t.Fatalf("failed to parse app private key fixture: %v", err)
	}

	return key
}

func mustParseURL(t *testing.T, raw string) *url.URL {
	t.Helper()

	u, err := url.Parse(raw)
	if err != nil {
		t.Fatalf("failed to parse url %q: %v", raw, err)
	}

	return u
}

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

type nonSeekableReader struct {
	r io.Reader
}

func (n *nonSeekableReader) Read(p []byte) (int, error) {
	return n.r.Read(p)
}

type closeRecorder struct {
	io.Reader
	closed atomic.Bool
}

func (c *closeRecorder) Close() error {
	c.closed.Store(true)
	return nil
}

// testStorage is a ghct.Storage that returns the configured results and, like a real storage, consumes the body passed to Put without restoring it.
type testStorage struct {
	getResp *http.Response
	getErr  error
	putErr  error

	mu        sync.Mutex
	putBodies []string
}

func (s *testStorage) Get(_ context.Context, _ *http.Request) (*http.Response, error) {
	return s.getResp, s.getErr
}

func (s *testStorage) Put(_ context.Context, resp *http.Response) error {
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	resp.Body = http.NoBody

	s.mu.Lock()
	defer s.mu.Unlock()

	s.putBodies = append(s.putBodies, string(b))

	return s.putErr
}
