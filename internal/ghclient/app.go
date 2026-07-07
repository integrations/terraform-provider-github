package ghclient

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/google/go-github/v88/github"
	lru "github.com/hashicorp/golang-lru/v2"
	"github.com/shurcooL/githubv4"
	"golang.org/x/sync/semaphore"
)

// appClientCacheSize defines the maximum number of app clients to cache in the appSource. This is used to limit memory usage while still providing efficient access to clients for different owners.
const appClientCacheSize = 8

// appSource is a concrete implementation of a [Source] that uses the provided app credentials to create GitHub clients.
type appSource struct {
	clientID           string
	privateKey         []byte
	semaCache          *lru.Cache[string, *semaphore.Weighted]
	restClientCache    *lru.Cache[string, *github.Client]
	graphQLClientCache *lru.Cache[string, *githubv4.Client]
	opts               SourceOptions
}

// NewAppSource creates a new appSource that provides GitHub clients authenticated as either the app itself or as an installation.
func NewAppSource(clientID string, privateKey []byte, opts SourceOptions) (*appSource, error) {
	semaCache, err := lru.New[string, *semaphore.Weighted](appClientCacheSize)
	if err != nil {
		return nil, err
	}

	restClientCache, err := lru.New[string, *github.Client](appClientCacheSize)
	if err != nil {
		return nil, err
	}

	graphQLClientCache, err := lru.New[string, *githubv4.Client](appClientCacheSize)
	if err != nil {
		return nil, err
	}

	if opts.Cache && opts.CacheBasePath == "" {
		s, err := os.MkdirTemp("", "*")
		if err != nil {
			return nil, fmt.Errorf("failed to create temporary cache directory: %w", err)
		}
		opts.CacheBasePath = s
	}

	return &appSource{
		clientID:           clientID,
		privateKey:         privateKey,
		semaCache:          semaCache,
		restClientCache:    restClientCache,
		graphQLClientCache: graphQLClientCache,
		opts:               opts,
	}, nil
}

// RESTClient returns the default GitHub client for the app source, which is an authenticated client with access to resources based on the app's permissions.
func (s *appSource) RESTClient() (*github.Client, error) {
	key := "_"
	if c, ok := s.restClientCache.Get(key); ok {
		return c, nil
	}

	sema, ok := s.semaCache.Get(key)
	if !ok {
		sema = semaphore.NewWeighted(maxConcurrentRequests)
		s.semaCache.Add(key, sema)
	}

	c, err := NewAppRESTClient(s.clientID, s.privateKey, nil, s.opts.getRESTClientOptions(sema, "app-rest-"+s.clientID))
	if err != nil {
		return nil, fmt.Errorf("failed to create app client: %w", err)
	}
	s.restClientCache.Add(key, c)

	return c, nil
}

// OwnerRESTClient returns a GitHub client authenticated to access resources owned by the specified owner. It creates a client for the installation associated with the owner, if available, or falls back to the default app client if no specific installation is found.
func (s *appSource) OwnerRESTClient(ctx context.Context, owner string) (*github.Client, error) {
	key := owner
	if c, ok := s.restClientCache.Get(key); ok {
		return c, nil
	}

	sema, ok := s.semaCache.Get(key)
	if !ok {
		sema = semaphore.NewWeighted(maxConcurrentRequests)
		s.semaCache.Add(key, sema)
	}

	installationID, err := s.GetInstallationID(ctx, owner)
	if err != nil {
		return nil, fmt.Errorf("failed to get installation id for owner %q: %w", owner, err)
	}

	c, err := NewAppRESTClient(s.clientID, s.privateKey, installationID, s.opts.getRESTClientOptions(sema, fmt.Sprintf("app-rest-%s-%s", s.clientID, owner)))
	if err != nil {
		return nil, fmt.Errorf("failed to create app client for owner %q: %w", owner, err)
	}
	s.restClientCache.Add(key, c)

	return c, nil
}

// GraphQLClient returns the default GitHub GraphQL client for the app source, which is an authenticated client with access to resources based on the app's permissions.
func (s *appSource) GraphQLClient() (*githubv4.Client, error) {
	key := "_"
	if c, ok := s.graphQLClientCache.Get(key); ok {
		return c, nil
	}

	sema, ok := s.semaCache.Get(key)
	if !ok {
		sema = semaphore.NewWeighted(maxConcurrentRequests)
		s.semaCache.Add(key, sema)
	}

	c, err := NewAppGraphQLClient(s.clientID, s.privateKey, nil, s.opts.getGraphQLClientOptions(sema))
	if err != nil {
		return nil, fmt.Errorf("failed to create app graphql client: %w", err)
	}
	s.graphQLClientCache.Add(key, c)

	return c, nil
}

// OwnerGraphQLClient returns a GitHub GraphQL client authenticated to access resources owned by the specified owner. It creates a client for the installation associated with the owner, if available, or falls back to the default app client if no specific installation is found.
func (s *appSource) OwnerGraphQLClient(ctx context.Context, owner string) (*githubv4.Client, error) {
	key := owner
	if c, ok := s.graphQLClientCache.Get(key); ok {
		return c, nil
	}

	sema, ok := s.semaCache.Get(key)
	if !ok {
		sema = semaphore.NewWeighted(maxConcurrentRequests)
		s.semaCache.Add(key, sema)
	}

	installationID, err := s.GetInstallationID(ctx, owner)
	if err != nil {
		return nil, fmt.Errorf("failed to get installation id for owner %q: %w", owner, err)
	}

	c, err := NewAppGraphQLClient(s.clientID, s.privateKey, installationID, s.opts.getGraphQLClientOptions(sema))
	if err != nil {
		return nil, fmt.Errorf("failed to create app graphql client for owner %q: %w", owner, err)
	}
	s.graphQLClientCache.Add(key, c)

	return c, nil
}

// GetInstallationID retrieves the installation ID for the specified owner (which can be either a user or an organization). It first attempts to find an organization installation, and if that fails, it tries to find a user installation. If neither is found, it returns an error.
func (s *appSource) GetInstallationID(ctx context.Context, owner string) (*int64, error) {
	appClient, err := s.RESTClient()
	if err != nil {
		return nil, fmt.Errorf("failed to get app client: %w", err)
	}

	installation, resp, err := appClient.Apps.GetOrganizationInstallation(ctx, owner)
	if err != nil && (resp == nil || resp.StatusCode != http.StatusNotFound) {
		return nil, fmt.Errorf("failed to get installation for owner %q: %w", owner, err)
	}

	if installation == nil {
		ui, _, err := appClient.Apps.GetUserInstallation(ctx, owner)
		if err != nil {
			return nil, fmt.Errorf("failed to get installation for owner %q: %w", owner, err)
		}
		installation = ui
	}

	if installation == nil || installation.ID == nil {
		return nil, fmt.Errorf("no installation found for owner %q", owner)
	}

	return installation.ID, nil
}
