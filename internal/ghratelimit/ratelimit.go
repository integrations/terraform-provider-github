package ghratelimit

import (
	"context"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/integrations/terraform-provider-github/v6/internal/httpx"
)

// Limiter defines an interface for managing rate limit state, allowing the setting and retrieval of blocked until times for different rate limit categories.
type Limiter interface {
	// Block sets the blocked until time for the given rate limit category.
	Block(category string, until time.Time)

	// BlockedUntil retrieves the blocked until time for the given rate limit category or zero if no blocked until time is set.
	BlockedUntil(category string) time.Time
}

// NewLimiter creates a new instance of a Limiter that manages blocked until times for different rate limit categories.
func NewLimiter() Limiter {
	return &limiter{
		blockedUntil: make(map[string]time.Time),
	}
}

// limiter is a concrete implementation of the [Limiter] interface that uses a mutex to safely manage blocked until times for different rate limit categories in a concurrent environment.
type limiter struct {
	mu           sync.RWMutex
	blockedUntil map[string]time.Time
}

// Block sets the blocked until time for the given rate limit category. If the new blocked until time is later than the current one, it updates the value.
func (r *limiter) Block(category string, until time.Time) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if cur, exists := r.blockedUntil[category]; !exists || until.After(cur) {
		r.blockedUntil[category] = until
	}
}

// BlockedUntil retrieves the greater of the blocked until time for the rate limit category and the secondary rate limit category. If neither is set, it returns the zero value of [time.Time].
func (r *limiter) BlockedUntil(category string) time.Time {
	r.mu.RLock()
	defer r.mu.RUnlock()

	p := r.blockedUntil[category]
	s := r.blockedUntil[rateLimitCategorySecondary]

	if s.After(p) {
		return s
	}

	return p
}

// options defines the configuration for handling rate limits in the transport.
type options struct {
	defaultSecondaryWait time.Duration
	waitMax              time.Duration
	basePath             string
	limiter              Limiter
}

// Option defines a function type for configuring rate limit options.
type Option func(*options)

// WithDefaultSecondaryWait sets the default wait duration for secondary rate limits. If the provided duration is less than 0, it will be set to 0.
func WithDefaultSecondaryWait(d time.Duration) Option {
	return func(o *options) {
		if d < 0 {
			d = 0
		}

		o.defaultSecondaryWait = d
	}
}

// WithWaitMax sets the maximum wait duration for rate limits. If the provided duration is less than 0, it will be set to 0.
func WithWaitMax(d time.Duration) Option {
	return func(o *options) {
		if d < 0 {
			d = 0
		}

		o.waitMax = d
	}
}

// WithBasePath sets the path prefix of the API, such as "/api/v3" for GitHub Enterprise Server, which is removed from request paths before the rate limit category is looked up. A trailing slash is ignored.
func WithBasePath(p string) Option {
	return func(o *options) {
		p = strings.Trim(p, "/")
		if p != "" {
			p = "/" + p
		}

		o.basePath = p
	}
}

// WithLimiter sets a custom [Limiter] for managing rate limit state. If the provided Limiter is nil, the default implementation will be used.
func WithLimiter(limiter Limiter) Option {
	return func(o *options) {
		if limiter == nil {
			return
		}

		o.limiter = limiter
	}
}

// NewTransport returns a [http.RoundTripper] that wraps the given http.RoundTripper and handles rate limit responses by waiting for the specified duration before retrying. It also normalizes the rate limit headers in the response.
func NewTransport(inner http.RoundTripper, opts ...Option) http.RoundTripper {
	o := options{
		defaultSecondaryWait: 60 * time.Second,
		waitMax:              60 * time.Second,
		limiter:              NewLimiter(),
	}

	for _, opt := range opts {
		opt(&o)
	}

	return &rateLimitTransport{
		inner:                inner,
		defaultSecondaryWait: o.defaultSecondaryWait,
		waitMax:              o.waitMax,
		basePath:             o.basePath,
		limiter:              o.limiter,
	}
}

// rateLimitTransport is a [http.RoundTripper] that waits before sending a request while its rate limit category is blocked, and records the blocks reported by GitHub's rate limit responses.
type rateLimitTransport struct {
	inner                http.RoundTripper
	defaultSecondaryWait time.Duration
	waitMax              time.Duration
	basePath             string
	limiter              Limiter
}

// RoundTrip implements the [http.RoundTripper] interface for the [rateLimitTransport]. It checks if the request is blocked due to rate limits and waits if necessary. It also normalizes the rate limit headers in the response.
func (t *rateLimitTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	category := lookupRateLimitCategory(req.Method, t.trimBasePath(req.URL.Path))

	if err := t.waitIfBlocked(req.Context(), category); err != nil {
		if req.Body != nil {
			_ = req.Body.Close()
		}

		return nil, err
	}

	resp, err := t.inner.RoundTrip(req)
	if err != nil {
		// A RoundTripper shouldn't return a response with an error, but if one does its body still needs closing so the connection isn't leaked.
		if resp != nil && resp.Body != nil {
			_ = resp.Body.Close()
		}

		return nil, err
	}

	t.normalizeRateLimitHeaders(resp)
	t.blockIfExhausted(resp, category)

	return resp, nil
}

// trimBasePath removes the configured base path from the given request path, leaving a path that starts with a slash.
func (t *rateLimitTransport) trimBasePath(p string) string {
	if t.basePath == "" {
		return p
	}

	if rest, ok := strings.CutPrefix(p, t.basePath+"/"); ok {
		return "/" + rest
	}

	return p
}

// waitIfBlocked checks if the request is blocked due to rate limits for the given category. If it is blocked, it waits until the blocked time or until the context is done. If the wait exceeds the maximum allowed wait time, it returns a [WaitExceededError].
func (t *rateLimitTransport) waitIfBlocked(ctx context.Context, category string) error {
	start := time.Now()
	deadline := start.Add(t.waitMax)
	if d, ok := ctx.Deadline(); ok && d.Before(deadline) {
		deadline = d
	}

	for {
		if err := ctx.Err(); err != nil {
			return err
		}

		until := t.limiter.BlockedUntil(category)

		wait := time.Until(until)
		if wait <= 0 {
			return nil
		}

		if until.After(deadline) {
			blocked := category
			if !t.limiter.BlockedUntil(rateLimitCategorySecondary).Before(until) {
				blocked = rateLimitCategorySecondary
			}

			return &WaitExceededError{Category: blocked, Until: until, MaxWait: deadline.Sub(start)}
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(wait):
		}
	}
}

// normalizeRateLimitHeaders normalizes the rate limit headers in the response. If the response indicates a secondary rate limit, it sets the Retry-After header based on the default secondary wait duration or the Retry-After header in the response. If the response indicates a primary rate limit, it sets the Retry-After header based on the X-RateLimit-Reset header in the response.
func (t *rateLimitTransport) normalizeRateLimitHeaders(resp *http.Response) {
	if resp.StatusCode != http.StatusForbidden && resp.StatusCode != http.StatusTooManyRequests {
		return
	}

	if isSecondaryRateLimit(resp) {
		resp.Header.Del(headerRateRemaining)
		resp.Header.Del(headerRateReset)

		wait, ok := httpx.ParseRetryAfter(resp.Header.Get(headerRetryAfter))
		if !ok {
			d := t.defaultSecondaryWait

			resp.Header.Set(headerRetryAfter, strconv.FormatInt(ceilSeconds(d), 10))
			wait = d
		}

		t.limiter.Block(rateLimitCategorySecondary, time.Now().Add(wait))
		return
	}

	if isPrimaryRateLimit(resp) {
		if resetEpoch, err := strconv.ParseInt(resp.Header.Get(headerRateReset), 10, 64); err == nil {
			resetTime := time.Unix(resetEpoch, 0)
			wait := time.Until(resetTime)

			if wait <= 0 {
				resp.Header.Set(headerRetryAfter, "0")
				return
			}

			resp.Header.Set(headerRetryAfter, strconv.FormatInt(ceilSeconds(wait), 10))
			return
		}
	}
}

// blockIfExhausted blocks the request category until the reset time if the primary rate limit is exhausted. If the response reports a different rate limit category it is also blocked, so a mismatch between the local category lookup and GitHub cannot let requests through.
func (t *rateLimitTransport) blockIfExhausted(resp *http.Response, reqCategory string) {
	if resp.Header.Get(headerRateRemaining) != "0" {
		return
	}

	resetEpoch, err := strconv.ParseInt(resp.Header.Get(headerRateReset), 10, 64)
	if err != nil {
		return
	}

	reset := time.Unix(resetEpoch, 0)
	if !reset.After(time.Now()) {
		return
	}

	t.limiter.Block(reqCategory, reset)

	if category := resp.Header.Get(headerRateResource); category != "" && category != reqCategory {
		t.limiter.Block(category, reset)
	}
}

// ceilSeconds returns the duration in whole seconds, rounded up so a wait is never shortened. It divides rather than adding before dividing so it can't overflow.
func ceilSeconds(d time.Duration) int64 {
	seconds := int64(d / time.Second)
	if d%time.Second > 0 {
		seconds++
	}

	return seconds
}
