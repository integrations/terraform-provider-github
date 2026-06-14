package ghclient

import (
	"context"

	"github.com/google/go-github/v88/github"
	"github.com/shurcooL/githubv4"
)

// Source provides an interface for obtaining GitHub clients. It abstracts away the details of how the clients are created and authenticated, allowing different implementations (such as app-based, token-based, or anonymous) to be used interchangeably.
type Source interface {
	// RESTClient returns the default GitHub client for the source.
	RESTClient() (*github.Client, error)

	// OwnerRESTClient returns a GitHub client authenticated to access resources owned by the specified owner (which can be either a user or an organization). This method is only applicable for app and token sources.
	OwnerRESTClient(ctx context.Context, owner string) (*github.Client, error)

	// GraphQLClient returns the default GitHub GraphQL client for the source.
	GraphQLClient() (*githubv4.Client, error)

	// OwnerGraphQLClient returns a GitHub GraphQL client authenticated to access resources owned by the specified owner (which can be either a user or an organization). This method is only applicable for app and token sources.
	OwnerGraphQLClient(ctx context.Context, owner string) (*githubv4.Client, error)
}
