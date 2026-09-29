package ghclient

import (
	"fmt"
	"net/http"

	ghct "github.com/bored-engineer/github-conditional-http-transport"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/logging"
	"golang.org/x/oauth2"
)

// cloneTransport attempts to clone the given http.RoundTripper if it is an *http.Transport, otherwise it returns the original RoundTripper. Cloning the transport is important to avoid sharing state (such as idle connections) between different clients that use the same base transport.
func cloneTransport(tr http.RoundTripper, opts ClientOptions) http.RoundTripper {
	if dtr, ok := tr.(*http.Transport); ok {
		htr := dtr.Clone()
		htr.ForceAttemptHTTP2 = true
		htr.MaxIdleConns = opts.MaxIdleConns
		htr.MaxIdleConnsPerHost = opts.MaxIdleConns
		htr.IdleConnTimeout = opts.IdleConnTimeout
		return htr
	}

	return tr
}

// newTransport creates a new HTTP RoundTripper that wraps the provided token source with OAuth2 authentication, adds conditional request caching, logging, and retry logic based on the provided options. The resulting RoundTripper is designed to be used with GitHub API clients to handle authentication, caching, rate limiting, and retries in a consistent manner.
func newTransport(tokenSource oauth2.TokenSource, opts ClientOptions) (http.RoundTripper, error) {
	tr := cloneTransport(http.DefaultTransport, opts)

	// The cache transport must be wrapped directly around the base transport, before the
	// OAuth2 transport is applied. The OAuth2 transport injects the Authorization header on a
	// cloned request immediately before invoking its Base RoundTripper, so any transport wrapping
	// it from the outside (i.e. added to tr afterwards) would only ever see requests without the
	// Authorization header. Since the cache transport partitions/validates its cache entries based
	// on the request's Authorization header (see the Vary handling in
	// github.com/bored-engineer/github-conditional-http-transport), placing it outside of the
	// OAuth2 transport silently breaks per-token cache validation for authenticated requests.
	if opts.Cache.Enabled {
		store, err := createCacheStore(opts.Cache)
		if err != nil {
			return nil, fmt.Errorf("failed to create cache store: %w", err)
		}

		tr = ghct.NewTransport(store, tr)
	}

	if tokenSource != nil {
		tr = &oauth2.Transport{
			Base:   tr,
			Source: tokenSource,
		}
	}

	tr = logging.NewLoggingHTTPTransport(tr)

	if opts.Retry.Max > 0 {
		rtr, err := newRetryTransport(tr, opts.Retry)
		if err != nil {
			return nil, err
		}
		tr = rtr
	}

	tr = newRateLimitTransport(tr)

	if opts.Concurrency.Max > 0 {
		ctr, err := newThrottleTransport(tr, opts.Concurrency)
		if err != nil {
			return nil, err
		}
		tr = ctr
	}

	return tr, nil
}
