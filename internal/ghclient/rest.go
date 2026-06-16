package ghclient

import (
	"fmt"

	"github.com/google/go-github/v88/github"
	"github.com/jferrl/go-githubauth"
	"golang.org/x/oauth2"
)

// NewAnonymousRESTClient creates a new GitHub client that is unauthenticated. This client will have limited access to public resources and will be subject to stricter rate limits compared to authenticated clients.
func NewAnonymousRESTClient(opts Options) (*github.Client, error) {
	return newRESTClient(nil, opts)
}

// NewAppRESTClient creates a new GitHub client authenticated as either the app itself (if installationID is nil) or as the specified installation (if installationID is provided), using the app's private key.
func NewAppRESTClient(clientID string, privateKey []byte, installationID *int64, opts Options) (*github.Client, error) {
	tokenSource, err := githubauth.NewApplicationTokenSource(clientID, privateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to create app token source: %w", err)
	}

	if installationID != nil {
		tokenSource = githubauth.NewInstallationTokenSource(*installationID, tokenSource)
	}

	return newRESTClient(tokenSource, opts)
}

// NewTokenRESTClient creates a new GitHub client authenticated with the provided personal access token.
func NewTokenRESTClient(token string, opts Options) (*github.Client, error) {
	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})

	return newRESTClient(tokenSource, opts)
}

// newRESTClient creates a new GitHub client using the provided OAuth2 token source and options. It sets up the client's transport with caching and rate limit handling, and configures the client's API URLs based on the provided options.
func newRESTClient(tokenSource oauth2.TokenSource, opts Options) (*github.Client, error) {
	tr, err := newTransport(tokenSource, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to create transport: %w", err)
	}

	return github.NewClient(github.WithTransport(tr), github.WithTimeout(clientTimeout), github.WithDisableRateLimitCheck(), github.WithURLs(&opts.RESTAPIURL, nil))
}
