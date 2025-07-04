package github

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"
)

// --- Unified Mock Refresh Function ---
func makeMockRefreshFunc(token string, expiry time.Time, fail bool) func(context.Context) (string, time.Time, error) {
	return func(ctx context.Context) (string, time.Time, error) {
		if fail {
			return "", time.Time{}, errors.New("mock refresh failure")
		}
		return token, expiry, nil
	}
}

// --- RefreshingTokenSource Tests ---

func TestRefreshingTokenSource_InitialValidToken(t *testing.T) {
	exp := time.Now().Add(5 * time.Minute)
	ts := NewRefreshingTokenSource("init-token", exp, makeMockRefreshFunc("new-token", time.Now().Add(10*time.Minute), false))

	token, err := ts.Token()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if token.AccessToken != "init-token" {
		t.Errorf("expected init-token, got %s", token.AccessToken)
	}
}

func TestRefreshingTokenSource_RefreshesAfterExpiry(t *testing.T) {
	exp := time.Now().Add(-1 * time.Minute)
	ts := NewRefreshingTokenSource("expired-token", exp, makeMockRefreshFunc("refreshed-token", time.Now().Add(10*time.Minute), false))

	token, err := ts.Token()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if token.AccessToken != "refreshed-token" {
		t.Errorf("expected refreshed-token, got %s", token.AccessToken)
	}
}

func TestRefreshingTokenSource_RefreshFails(t *testing.T) {
	exp := time.Now().Add(-1 * time.Minute)
	ts := NewRefreshingTokenSource("expired-token", exp, makeMockRefreshFunc("", time.Time{}, true))

	_, err := ts.Token()
	if err == nil {
		t.Fatal("expected error on refresh failure, got nil")
	}
}

func TestRefreshingTokenSource_Token(t *testing.T) {
	rt := NewRefreshingTokenSource("initial-token", time.Now().Add(-10*time.Minute), func(ctx context.Context) (string, time.Time, error) {
		return "fake-token", time.Now().Add(10 * time.Minute), nil
	})
	token, err := rt.Token()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if token.AccessToken != "fake-token" {
		t.Errorf("Expected token to be 'fake-token', got %s", token.AccessToken)
	}
}

// --- Config Behavior Tests ---

func TestConfig_Anonymous(t *testing.T) {
	cfg := &Config{Token: ""}
	if !cfg.Anonymous() {
		t.Error("expected anonymous to be true when token is empty")
	}
}

func TestConfig_NotAnonymous(t *testing.T) {
	cfg := &Config{Token: "abc"}
	if cfg.Anonymous() {
		t.Error("expected anonymous to be false when token is set")
	}
}

func TestAnonymousClient(t *testing.T) {
	config := &Config{}
	if !config.Anonymous() {
		t.Error("Expected config to be anonymous when no token is set")
	}
	client := config.AnonymousHTTPClient()
	if client == nil {
		t.Fatal("Expected a non-nil HTTP client")
	}
}

func TestAuthenticatedClientWithMock(t *testing.T) {
	os.Setenv("GITHUB_APP_ID", "123456")
	os.Setenv("GITHUB_APP_INSTALLATION_ID", "654321")
	os.Setenv("GITHUB_APP_PEM", "dummy-pem-content")

	cfg := &Config{Token: "initial", BaseURL: "https://api.github.com"}

	client := cfg.AuthenticatedHTTPClient()
	if client == nil {
		t.Fatal("Expected non-nil authenticated HTTP client")
	}
}