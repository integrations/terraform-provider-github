package ghclient

import (
	"context"
	"fmt"
	"net/http"

	ghct "github.com/bored-engineer/github-conditional-http-transport"
	ratelimit "github.com/gofri/go-github-ratelimit/v2/github_ratelimit"
	ratelimitp "github.com/gofri/go-github-ratelimit/v2/github_ratelimit/github_primary_ratelimit"
	ratelimits "github.com/gofri/go-github-ratelimit/v2/github_ratelimit/github_secondary_ratelimit"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/logging"
	"golang.org/x/oauth2"
)

// cloneTransport attempts to clone the given http.RoundTripper if it is an *http.Transport, otherwise it returns the original RoundTripper. Cloning the transport is important to avoid sharing state (such as idle connections) between different clients that use the same base transport.
func cloneTransport(tr http.RoundTripper) http.RoundTripper {
	if dtr, ok := tr.(*http.Transport); ok {
		return dtr.Clone()
	}

	return tr
}

// newTransport creates a new HTTP RoundTripper that wraps the provided token source with OAuth2 authentication, adds conditional request caching, logging, and retry logic based on the provided options. The resulting RoundTripper is designed to be used with GitHub API clients to handle authentication, caching, rate limiting, and retries in a consistent manner.
func newTransport(tokenSource oauth2.TokenSource, opts Options) (http.RoundTripper, error) {
	tr := cloneTransport(http.DefaultTransport)

	if tokenSource != nil {
		tr = &oauth2.Transport{
			Base:   tr,
			Source: tokenSource,
		}
	}

	// Create a cache store.
	store, err := createCacheStore(opts.CachePath, opts.cacheRef)
	if err != nil {
		return nil, fmt.Errorf("failed to create cache store: %w", err)
	}
	tr = ghct.NewTransport(store, tr)

	// Log each actual HTTP round-trip (including retries)
	tr = logging.NewLoggingHTTPTransport(tr)

	if opts.RetryMax > 0 {
		// Wrap with retry transport
		retryClient := retryablehttp.NewClient()
		retryClient.Logger = nil
		retryClient.HTTPClient = &http.Client{Transport: tr, Timeout: clientTimeout}
		retryClient.RetryMax = opts.RetryMax
		retryClient.RetryWaitMin = opts.RetryWaitMin
		retryClient.RetryWaitMax = opts.RetryWaitMax
		retryClient.CheckRetry = func(ctx context.Context, resp *http.Response, err error) (bool, error) {
			if err != nil {
				return retryablehttp.DefaultRetryPolicy(ctx, resp, err)
			}
			if resp.StatusCode == 0 || (resp.StatusCode >= 500 && resp.StatusCode != http.StatusNotImplemented) {
				return true, nil
			}
			return false, nil
		}

		// Use the RoundTripper adapter so it composes with other transports
		tr = &retryablehttp.RoundTripper{Client: retryClient}
	}

	// Wrap with rate limit transport
	tr = ratelimit.New(tr, ratelimitp.WithLimitDetectedCallback(primaryRateLimitCallback), ratelimits.WithLimitDetectedCallback(secondaryRateLimitCallback))

	return tr, nil
}

// primaryRateLimitCallback is a callback function that is called when the GitHub API primary rate limit is detected. It logs a warning message with the category of the rate limit and the reset time.
func primaryRateLimitCallback(cb *ratelimitp.CallbackContext) {
	tflog.Warn(cb.Request.Context(), "GitHub API primary rate limit detected.", map[string]any{"category": cb.Category, "reset_time": cb.ResetTime})
}

// secondaryRateLimitCallback is a callback function that is called when the GitHub API secondary rate limit is detected. It logs a warning message with the reset time.
func secondaryRateLimitCallback(cb *ratelimits.CallbackContext) {
	tflog.Warn(cb.Request.Context(), "GitHub API secondary rate limit detected.", map[string]any{"reset_time": cb.ResetTime})
}
