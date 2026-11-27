package ghclient

import (
	"context"
	"crypto"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/google/go-github/v93/github"
	"github.com/jferrl/go-githubauth"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
	"golang.org/x/sync/singleflight"
)

// newAppSource creates a [Source] that authenticates as a GitHub App using the client ID and signer in the options.
func newAppSource(opts sourceOptions) (*appSource, error) {
	ru, err := restURL(opts.baseURL, opts.isGHES)
	if err != nil {
		return nil, fmt.Errorf("failed to create rest url: %w", err)
	}

	return &appSource{
		restURL: ru.String(),
		groups:  make(map[string]*sourceClientGroup),
		opts:    opts,
	}, nil
}

// appSource is a [Source] that lazily creates and caches a client group for the app itself and for each owner's installation.
type appSource struct {
	sfg     singleflight.Group
	restURL string
	mu      sync.RWMutex
	groups  map[string]*sourceClientGroup
	opts    sourceOptions
}

// getGroup returns the client group stored under key, if any.
func (s *appSource) getGroup(key string) (*sourceClientGroup, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	sg, ok := s.groups[key]

	return sg, ok
}

// group returns the client group for the owner, creating it if needed. A nil owner returns the group authenticated as the app itself rather than as an installation.
//
// Creation is coalesced per owner, so concurrent callers share a single group and a single open cache store. The shared work runs on a context detached from the caller's cancellation so one caller giving up doesn't fail the others, while each caller can still stop waiting when its own context is done.
func (s *appSource) group(ctx context.Context, owner *string) (*sourceClientGroup, error) {
	var key string
	if owner == nil {
		key = "_"
	} else {
		key = strings.ToLower(*owner)
	}

	if sg, ok := s.getGroup(key); ok {
		return sg, nil
	}

	workCtx := context.WithoutCancel(ctx)

	ch := s.sfg.DoChan(key, func() (any, error) {
		// Re-check now that we're the only creator for this owner, in case a previous flight finished after the first lookup.
		if sg, ok := s.getGroup(key); ok {
			return sg, nil
		}

		var installationID *int64
		cacheDir := "app-" + s.opts.appClientID

		if owner != nil {
			id, err := s.installationID(workCtx, *owner)
			if err != nil {
				return nil, fmt.Errorf("failed to get installation id: %w", err)
			}

			installationID = &id
			cacheDir = cacheDir + "-" + key
		}

		ts, err := appTokenSource(s.restURL, s.opts.appClientID, s.opts.appSigner, installationID)
		if err != nil {
			return nil, fmt.Errorf("failed to create app token source: %w", err)
		}

		sg, err := newSourceClientGroup(s.opts, cacheDir, ts)
		if err != nil {
			return nil, fmt.Errorf("failed to create app client group: %w", err)
		}

		s.mu.Lock()
		s.groups[key] = sg
		s.mu.Unlock()

		return sg, nil
	})

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case res := <-ch:
		if res.Err != nil {
			return nil, res.Err
		}

		sg, _ := res.Val.(*sourceClientGroup)

		return sg, nil
	}
}

// RESTClient returns the default GitHub client for the app source, which is authenticated as the app and has access to the app level endpoints.
func (s *appSource) RESTClient(ctx context.Context) (*github.Client, error) {
	sg, err := s.group(ctx, nil)
	if err != nil {
		return nil, err
	}

	return sg.restClient, nil
}

// OwnerRESTClient returns a GitHub client authenticated as the app installation for the specified owner (which can be either a user or an organization). It returns an error if the app is not installed for the owner.
func (s *appSource) OwnerRESTClient(ctx context.Context, owner string) (*github.Client, error) {
	sg, err := s.group(ctx, &owner)
	if err != nil {
		return nil, err
	}

	return sg.restClient, nil
}

// GraphQLClient returns the default GitHub GraphQL client for the app source, which is authenticated as the app.
func (s *appSource) GraphQLClient(ctx context.Context) (*githubv4.Client, error) {
	sg, err := s.group(ctx, nil)
	if err != nil {
		return nil, err
	}

	return sg.graphqlClient, nil
}

// OwnerGraphQLClient returns a GitHub GraphQL client authenticated as the app installation for the specified owner (which can be either a user or an organization). It returns an error if the app is not installed for the owner.
func (s *appSource) OwnerGraphQLClient(ctx context.Context, owner string) (*githubv4.Client, error) {
	sg, err := s.group(ctx, &owner)
	if err != nil {
		return nil, err
	}

	return sg.graphqlClient, nil
}

// installationID retrieves the installation ID for the specified owner (which can be either a user or an organization). It first attempts to find an organization installation, and if that isn't found, it tries to find a user installation. If neither is found, it returns an error wrapping [ErrInstallationNotFound].
func (s *appSource) installationID(ctx context.Context, owner string) (int64, error) {
	appClient, err := s.RESTClient(ctx)
	if err != nil {
		return 0, fmt.Errorf("failed to get app client: %w", err)
	}

	installation, resp, err := appClient.Apps.GetOrganizationInstallation(ctx, owner)
	if err != nil {
		if resp == nil || resp.Response == nil || resp.StatusCode != http.StatusNotFound {
			return 0, fmt.Errorf("failed to get installation for owner %q: %w", owner, err)
		}

		installation, resp, err = appClient.Apps.GetUserInstallation(ctx, owner)
		if err != nil {
			if resp != nil && resp.Response != nil && resp.StatusCode == http.StatusNotFound {
				return 0, fmt.Errorf("%w for owner %q", ErrInstallationNotFound, owner)
			}

			return 0, fmt.Errorf("failed to get installation for owner %q: %w", owner, err)
		}
	}

	if installation == nil || installation.ID == nil {
		return 0, fmt.Errorf("installation for owner %q has no id", owner)
	}

	return installation.GetID(), nil
}

// appTokenSource creates an OAuth2 token source for the GitHub App using the provided client ID, signer, and optional installation ID. If an installation ID is provided, it returns a token source that generates tokens for that specific installation; otherwise, it returns a token source for the app itself.
//
// The installation token source uses go-githubauth's own HTTP client and retry rather than the source's transport stack. The request is authenticated with the app JWT and its response body is the installation token, so it mustn't go through the logging transport, and refreshes are already serialised per installation so they don't need throttling. go-githubauth sends it with a background context, so it isn't cancelled with the caller's request; its client's response header timeout bounds a stalled request instead.
func appTokenSource(restURL, clientID string, signer crypto.Signer, installationID *int64) (oauth2.TokenSource, error) {
	ts, err := githubauth.NewApplicationTokenSourceFromSigner(clientID, signer)
	if err != nil {
		return nil, fmt.Errorf("failed to create application token source: %w", err)
	}

	if installationID != nil {
		ts = githubauth.NewInstallationTokenSource(*installationID, ts, githubauth.WithBaseURL(restURL))
	}

	return ts, nil
}
