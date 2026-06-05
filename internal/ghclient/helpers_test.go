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
	uploadURL := ts.URL + "/"

	source, err := NewAppSource("123456789", privateKeyData, Options{RESTAPIURL: &apiURL, RESTUploadURL: &uploadURL})
	if err != nil {
		t.Fatalf("failed to create app source: %v", err)
	}

	return source
}
