package ghclient

import (
	"fmt"
	"net/http"

	"github.com/jferrl/go-githubauth"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

// NewAnonymousGraphQLClient creates a new GitHub GraphQL client that is unauthenticated. This client will have limited access to public resources and will be subject to stricter rate limits compared to authenticated clients.
func NewAnonymousGraphQLClient(opts ClientOptions) (*githubv4.Client, error) {
	return newGraphQLClient(nil, opts)
}

// NewAppGraphQLClient creates a new GitHub GraphQL client authenticated as either the app itself (if installationID is nil) or as the specified installation (if installationID is provided), using the app's private key.
func NewAppGraphQLClient(clientID string, privateKey []byte, installationID *int64, opts ClientOptions) (*githubv4.Client, error) {
	tokenSource, err := githubauth.NewApplicationTokenSource(clientID, privateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to create app token source: %w", err)
	}

	if installationID != nil {
		authOpts := []githubauth.InstallationTokenSourceOpt{}
		if opts.BaseURL != "" {
			u, err := opts.getRESTURL()
			if err != nil {
				return nil, fmt.Errorf("failed to get rest url: %w", err)
			}
			authOpts = append(authOpts, githubauth.WithBaseURL(*u))
		}

		tokenSource = githubauth.NewInstallationTokenSource(*installationID, tokenSource, authOpts...)
	}

	return newGraphQLClient(tokenSource, opts)
}

// NewTokenGraphQLClient creates a new GitHub GraphQL client authenticated with the provided personal access token.
func NewTokenGraphQLClient(token string, opts ClientOptions) (*githubv4.Client, error) {
	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})

	return newGraphQLClient(tokenSource, opts)
}

// newGraphQLClient creates a new GitHub GraphQL client using the provided OAuth2 token source and options. It sets up the client's transport with caching and rate limit handling, and configures the client's API URL based on the provided options.
func newGraphQLClient(tokenSource oauth2.TokenSource, opts ClientOptions) (*githubv4.Client, error) {
	tr, err := newTransport(tokenSource, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to create transport: %w", err)
	}

	client := &http.Client{Transport: tr, Timeout: clientTimeout}

	if opts.BaseURL == "" {
		return githubv4.NewClient(client), nil
	}

	u, err := opts.getGraphQLURL()
	if err != nil {
		return nil, fmt.Errorf("failed to get graphql url: %w", err)
	}

	return githubv4.NewEnterpriseClient(*u, client), nil
}
