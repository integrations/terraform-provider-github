package ghclient

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/v89/github"
	"github.com/shurcooL/githubv4"
	"golang.org/x/sync/semaphore"
)

// tokenSource is a concrete implementation of a [Source] that uses the provided token credentials to create GitHub clients.
type tokenSource struct {
	restClient    *github.Client
	graphQLClient *githubv4.Client
}

// NewTokenSource creates a new tokenSource that provides a GitHub client authenticated with the provided personal access token.
func NewTokenSource(token string, opts SourceOptions) (*tokenSource, error) {
	if opts.Cache && opts.CacheBasePath == "" {
		s, err := os.MkdirTemp("", "*")
		if err != nil {
			return nil, fmt.Errorf("failed to create temporary cache directory: %w", err)
		}
		opts.CacheBasePath = s
	}

	sema := semaphore.NewWeighted(maxConcurrentRequests)

	client, err := NewTokenRESTClient(token, opts.getRESTClientOptions(sema, "token-rest"))
	if err != nil {
		return nil, err
	}

	graphQLClient, err := NewTokenGraphQLClient(token, opts.getGraphQLClientOptions(sema))
	if err != nil {
		return nil, err
	}

	return &tokenSource{
		restClient:    client,
		graphQLClient: graphQLClient,
	}, nil
}

// RESTClient returns the default GitHub client for the token source, which is an authenticated client with access to resources based on the provided token.
func (s *tokenSource) RESTClient() (*github.Client, error) {
	return s.restClient, nil
}

// OwnerRESTClient returns a GitHub client authenticated to access resources owned by the specified owner. Since this is a token source, it can provide the same authenticated client for any owner, as the token's permissions will determine access to resources.
func (s *tokenSource) OwnerRESTClient(_ context.Context, _ string) (*github.Client, error) {
	return s.RESTClient()
}

// GraphQLClient returns the default GitHub GraphQL client for the token source, which is an authenticated client with access to resources based on the provided token.
func (s *tokenSource) GraphQLClient() (*githubv4.Client, error) {
	return s.graphQLClient, nil
}

// OwnerGraphQLClient returns a GitHub GraphQL client authenticated to access resources owned by the specified owner. Since this is a token source, it can provide the same authenticated client for any owner, as the token's permissions will determine access to resources.
func (s *tokenSource) OwnerGraphQLClient(_ context.Context, _ string) (*githubv4.Client, error) {
	return s.GraphQLClient()
}
