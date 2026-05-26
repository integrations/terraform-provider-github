package ghclient

import (
	"context"

	"github.com/google/go-github/v88/github"
	"github.com/shurcooL/githubv4"
)

// anonymousSource is a concrete implementation of a [Source] that creates GitHub clients without authentication.
type anonymousSource struct {
	restClient    *github.Client
	graphQLClient *githubv4.Client
}

// NewAnonymousSource creates a new anonymousSource that provides an unauthenticated GitHub client. This client will have limited access to public resources and will be subject to stricter rate limits compared to authenticated clients.
func NewAnonymousSource(options Options) (*anonymousSource, error) {
	client, err := NewAnonymousRESTClient(options)
	if err != nil {
		return nil, err
	}

	graphQLClient, err := NewAnonymousGraphQLClient(options)
	if err != nil {
		return nil, err
	}

	return &anonymousSource{
		restClient:    client,
		graphQLClient: graphQLClient,
	}, nil
}

// RESTClient returns the default GitHub client for the anonymous source, which is an unauthenticated client with limited access to public resources.
func (s *anonymousSource) RESTClient() (*github.Client, error) {
	return s.restClient, nil
}

// OwnerRESTClient returns a GitHub client authenticated to access resources owned by the specified owner. Since this is an anonymous source, it cannot provide an authenticated client for a specific owner, so it returns the same anonymous client for any owner.
func (s *anonymousSource) OwnerRESTClient(_ context.Context, _ string) (*github.Client, error) {
	return s.RESTClient()
}

// GraphQLClient returns the default GitHub GraphQL client for the anonymous source, which is an unauthenticated client with limited access to public resources.
func (s *anonymousSource) GraphQLClient() (*githubv4.Client, error) {
	return s.graphQLClient, nil
}

// OwnerGraphQLClient returns a GitHub GraphQL client authenticated to access resources owned by the specified owner. Since this is an anonymous source, it cannot provide an authenticated client for a specific owner, so it returns the same anonymous client for any owner.
func (s *anonymousSource) OwnerGraphQLClient(_ context.Context, _ string) (*githubv4.Client, error) {
	return s.GraphQLClient()
}
