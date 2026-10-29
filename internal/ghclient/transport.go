package ghclient

import (
	"fmt"
	"net"
	"net/http"
	"time"

	ghct "github.com/bored-engineer/github-conditional-http-transport"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/logging"
	"golang.org/x/oauth2"
)

// getHTTPTransport returns an HTTP transport configured with the given options, including idle connection pooling and retry logic. If the [http.DefaultTransport] is a [*http.Transport], it is cloned and configured with the given options; otherwise, a new transport is created with the default settings and the given options are applied.
func getHTTPTransport(t http.RoundTripper, opts ClientOptions) *http.Transport {
	if dtr, ok := t.(*http.Transport); ok {
		tr := dtr.Clone()
		tr.ForceAttemptHTTP2 = true
		tr.MaxIdleConns = opts.MaxIdleConns
		tr.MaxIdleConnsPerHost = opts.MaxIdleConns
		tr.IdleConnTimeout = opts.IdleConnTimeout
		return tr
	}

	dialer := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}

	return &http.Transport{
		Proxy:                 http.ProxyFromEnvironment,
		DialContext:           dialer.DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          opts.MaxIdleConns,
		MaxIdleConnsPerHost:   opts.MaxIdleConns,
		IdleConnTimeout:       opts.IdleConnTimeout,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
}

// newTransport creates a new HTTP RoundTripper that wraps the provided token source with OAuth2 authentication, adds conditional request caching, logging, and retry logic based on the provided options. The resulting RoundTripper is designed to be used with GitHub API clients to handle authentication, caching, rate limiting, and retries in a consistent manner.
func newTransport(tokenSource oauth2.TokenSource, opts ClientOptions) (http.RoundTripper, error) {
	var tr http.RoundTripper

	tr = getHTTPTransport(http.DefaultTransport, opts)

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

	rtr, err := newRewindableTransport(tr, 0)
	if err != nil {
		return nil, err
	}
	tr = rtr

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
