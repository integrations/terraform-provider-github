package ghclient

import (
	"net/http"

	ratelimit "github.com/gofri/go-github-ratelimit/v2/github_ratelimit"
	ratelimitp "github.com/gofri/go-github-ratelimit/v2/github_ratelimit/github_primary_ratelimit"
	ratelimits "github.com/gofri/go-github-ratelimit/v2/github_ratelimit/github_secondary_ratelimit"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// newRateLimitTransport returns a [http.RoundTripper] that wraps the given [http.RoundTripper] and uses the rate limit callbacks to detect and log rate limit warnings.
func newRateLimitTransport(inner http.RoundTripper) http.RoundTripper {
	return ratelimit.New(inner, ratelimitp.WithLimitDetectedCallback(primaryRateLimitCallback), ratelimits.WithLimitDetectedCallback(secondaryRateLimitCallback))
}

// primaryRateLimitCallback is a callback function that is called when the GitHub API primary rate limit is detected. It logs a warning message with the category of the rate limit and the reset time.
func primaryRateLimitCallback(cb *ratelimitp.CallbackContext) {
	tflog.Warn(cb.Request.Context(), "GitHub API primary rate limit detected.", map[string]any{"category": cb.Category, "reset_time": cb.ResetTime})
}

// secondaryRateLimitCallback is a callback function that is called when the GitHub API secondary rate limit is detected. It logs a warning message with the reset time.
func secondaryRateLimitCallback(cb *ratelimits.CallbackContext) {
	tflog.Warn(cb.Request.Context(), "GitHub API secondary rate limit detected.", map[string]any{"reset_time": cb.ResetTime})
}
