package ghclient

import (
	"context"

	"github.com/google/go-github/v92/github"
	"github.com/shurcooL/githubv4"
)

// newTokenSource creates a [Source] that uses a single client group authenticated with the token source in the options, or unauthenticated if there isn't one.
func newTokenSource(opts sourceOptions) (*tokenSource, error) {
	cg, err := newSourceClientGroup(opts, "token", opts.tokenSource)
	if err != nil {
		return nil, err
	}

	return &tokenSource{
		restClient:    cg.restClient,
		graphqlClient: cg.graphqlClient,
	}, nil
}

// tokenSource is a [Source] backed by a single client group, which is used for every owner.
type tokenSource struct {
	restClient    *github.Client
	graphqlClient *githubv4.Client
}

// RESTClient returns the GitHub client for the token source, which is authenticated with the token source if one was configured and unauthenticated otherwise.
func (s *tokenSource) RESTClient(_ context.Context) (*github.Client, error) {
	return s.restClient, nil
}

// OwnerRESTClient returns the same GitHub client as [tokenSource.RESTClient] for any owner, as the token's permissions determine which resources it can access.
func (s *tokenSource) OwnerRESTClient(_ context.Context, _ string) (*github.Client, error) {
	return s.restClient, nil
}

// GraphQLClient returns the GitHub GraphQL client for the token source, which is authenticated with the token source if one was configured and unauthenticated otherwise.
func (s *tokenSource) GraphQLClient(_ context.Context) (*githubv4.Client, error) {
	return s.graphqlClient, nil
}

// OwnerGraphQLClient returns the same GitHub GraphQL client as [tokenSource.GraphQLClient] for any owner, as the token's permissions determine which resources it can access.
func (s *tokenSource) OwnerGraphQLClient(_ context.Context, _ string) (*githubv4.Client, error) {
	return s.graphqlClient, nil
}
