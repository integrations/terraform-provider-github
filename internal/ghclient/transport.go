package ghclient

import (
	"net/http"
	"net/url"
	"strings"

	ghct "github.com/bored-engineer/github-conditional-http-transport"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/logging"
	"golang.org/x/oauth2"

	"github.com/integrations/terraform-provider-github/v6/internal/ghratelimit"
	"github.com/integrations/terraform-provider-github/v6/internal/httpx"
)

// newTransport creates a new HTTP RoundTripper that wraps the provided token source with OAuth2 authentication, adds conditional request caching, logging, and retry logic based on the provided options. The resulting RoundTripper is designed to be used with GitHub API clients to handle authentication, caching, rate limiting, and retries in a consistent manner.
func newTransport(opts clientOptions) http.RoundTripper {
	tr := httpx.DefaultHTTPTransport(nil)
	tr.MaxIdleConns = opts.maxIdleConns
	tr.MaxIdleConnsPerHost = opts.maxIdleConns
	tr.IdleConnTimeout = opts.idleConnTimeout
	tr.ResponseHeaderTimeout = opts.responseHeaderTimeout

	var rt http.RoundTripper = tr

	// The cache transport must be wrapped directly around the base transport, before the
	// OAuth2 transport is applied. The OAuth2 transport injects the Authorization header on a
	// cloned request immediately before invoking its Base RoundTripper, so any transport wrapping
	// it from the outside (i.e. added to tr afterwards) would only ever see requests without the
	// Authorization header. Since the cache transport partitions/validates its cache entries based
	// on the request's Authorization header (see the Vary handling in
	// github.com/bored-engineer/github-conditional-http-transport), placing it outside of the
	// OAuth2 transport silently breaks per-token cache validation for authenticated requests.
	//
	// The store is wrapped so cache failures never fail a request; ghct would otherwise fail the request and, on a failed write, return a response alongside the error.
	if opts.cacheEnabled {
		var store ghct.Storage
		if opts.cacheStore != nil {
			store = bestEffortStorage{inner: opts.cacheStore}
		}

		rt = ghct.NewTransport(store, rt)
	}

	if opts.tokenSource != nil {
		rt = &oauth2.Transport{Base: rt, Source: opts.tokenSource}
	}

	rt = logging.NewLoggingHTTPTransport(rt)

	rt = httpx.NewThrottleTransport(rt, opts.throttleLimiter)

	rt = ghratelimit.NewTransport(rt, ghratelimit.WithWaitMax(opts.retryWaitMax), ghratelimit.WithBasePath(rateLimitBasePath(opts.url)), ghratelimit.WithLimiter(opts.rateLimiter))

	rt = httpx.NewRetryTransport(rt, httpx.WithRetryMaxRetries(opts.retryMax), httpx.WithRetryWait(opts.retryWaitMin, opts.retryWaitMax), httpx.WithRetryJitter(opts.retryJitter), httpx.WithRetryNoRetryErrors(ghratelimit.ErrRateLimitWaitExceeded))

	rt = httpx.NewRewindableTransport(rt, httpx.WithRewindBufferBytes(10<<20))

	return rt
}

// rateLimitBasePath returns the path prefix of the API for the given client URL, such as "/api/v3" for GitHub Enterprise Server REST or "/api" for its GraphQL endpoint. REST URLs are the base path itself whereas GraphQL URLs also include the "/graphql" endpoint, which is not part of the prefix.
func rateLimitBasePath(u *string) string {
	if u == nil {
		return ""
	}

	pu, err := url.Parse(*u)
	if err != nil {
		return ""
	}

	return strings.TrimSuffix(strings.TrimRight(pu.Path, "/"), "/graphql")
}
