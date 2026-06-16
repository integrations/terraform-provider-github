package ghclient

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

type staticRoundTripper struct{}

func (s *staticRoundTripper) RoundTrip(_ *http.Request) (*http.Response, error) {
	return nil, fmt.Errorf("not used")
}

func testOptions(t *testing.T) Options {
	t.Helper()

	return Options{
		RESTAPIURL: "https://api.github.com/",
		GraphQLURL: "https://api.github.com/graphql",
	}
}

func mustReadTestAppPrivateKey(t *testing.T) []byte {
	t.Helper()

	privateKeyData, err := os.ReadFile(filepath.Join("..", "..", "github", "test-fixtures", "github-app-key.pem"))
	if err != nil {
		t.Fatalf("failed to read app private key fixture: %v", err)
	}

	return privateKeyData
}

func mustTestAppSource(t *testing.T, handler http.Handler) *appSource {
	t.Helper()

	ts := httptest.NewServer(handler)
	t.Cleanup(ts.Close)

	privateKeyData := mustReadTestAppPrivateKey(t)

	apiURL := ts.URL + "/"
	opts := testOptions(t)
	opts.RESTAPIURL = apiURL

	source, err := NewAppSource("123456789", privateKeyData, opts)
	if err != nil {
		t.Fatalf("failed to create app source: %v", err)
	}

	return source
}
