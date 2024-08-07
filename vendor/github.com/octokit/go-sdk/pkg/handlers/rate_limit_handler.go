package handlers

import (
	"context"
	"fmt"
	"log"
	netHttp "net/http"
	"strconv"
	"time"

	abs "github.com/microsoft/kiota-abstractions-go"
	kiotaHttp "github.com/microsoft/kiota-http-go"
	"github.com/octokit/go-sdk/pkg/headers"
)

// RateLimitHandler is a middleware that detects primary and secondary rate
// limits and retries requests after the appropriate time when necessary.
type RateLimitHandler struct {
	options RateLimitHandlerOptions
}

// RateLimitHandlerOptions is a struct that holds options for the RateLimitHandler.
// In the future, this could hold different strategies for handling rate limits:
// e.g. exponential backoff, jitter, throttling, etc.
type RateLimitHandlerOptions struct{}

// rateLimitHandlerOptions (lowercase, private) that RateLimitHandlerOptions
// (uppercase, public) implements.
type rateLimitHandlerOptions interface {
	abs.RequestOption
	IsRateLimited() func(req *netHttp.Request, res *netHttp.Response) RateLimitType
}

// RateLimitType is an enum that represents either None, Primary,
// or Secondary rate limiting
type RateLimitType int

const (
	None RateLimitType = iota
	Primary
	Secondary
)

var rateLimitKeyValue = abs.RequestOptionKey{
	Key: "RateLimitHandler",
}

// GetKey returns the unique RateLimitHandler key, used by Kiota to store
// request options.
func (options *RateLimitHandlerOptions) GetKey() abs.RequestOptionKey {
	return rateLimitKeyValue
}

// IsRateLimited returns a function that determines if an HTTP response was
// rate-limited, and if so, what type of rate limit was hit.
func (options *RateLimitHandlerOptions) IsRateLimited() func(req *netHttp.Request, resp *netHttp.Response) RateLimitType {
	return func(req *netHttp.Request, resp *netHttp.Response) RateLimitType {
		if resp.StatusCode != 429 && resp.StatusCode != 403 {
			return None
		}

		if resp.Header.Get(headers.RetryAfterKey) != "" {
			return Secondary // secondary rate limits are abuse limits
		}

		if resp.Header.Get(headers.XRateLimitRemainingKey) == "0" {
			return Primary
		}

		return None
	}
}

// NewRateLimitHandler creates a new RateLimitHandler with default options.
func NewRateLimitHandler() *RateLimitHandler {
	return &RateLimitHandler{}
}

// Intercept tries a request. If the response shows it was rate-limited, it
// retries the request after the appropriate period of time.
func (handler RateLimitHandler) Intercept(pipeline kiotaHttp.Pipeline, middlewareIndex int, request *netHttp.Request) (*netHttp.Response, error) {
	resp, err := pipeline.Next(request, middlewareIndex)
	if err != nil {
		return resp, err
	}

	rateLimit := handler.options.IsRateLimited()(request, resp)

	if rateLimit == Primary || rateLimit == Secondary {
		reqOption, ok := request.Context().Value(rateLimitKeyValue).(rateLimitHandlerOptions)
		if !ok {
			reqOption = &handler.options
		}
		return handler.retryRequest(request.Context(), pipeline, middlewareIndex, reqOption, rateLimit, request, resp)
	}
	return resp, nil
}

// retryRequest retries a request if it has been rate-limited.
func (handler RateLimitHandler) retryRequest(ctx context.Context, pipeline kiotaHttp.Pipeline, middlewareIndex int,
	options rateLimitHandlerOptions, rateLimitType RateLimitType, request *netHttp.Request, resp *netHttp.Response) (*netHttp.Response, error) {

	if rateLimitType == Secondary || rateLimitType == Primary {
		retryAfterDuration, err := parseRateLimit(resp)
		if err != nil {
			return nil, fmt.Errorf("failed to parse retry-after header into duration: %v", err)
		}
		if *retryAfterDuration < 0 {
			log.Printf("retry-after duration is negative: %s; sleeping until next request will be a no-op", *retryAfterDuration)
		}
		if rateLimitType == Secondary {
			log.Printf("Abuse detection mechanism (secondary rate limit) triggered, sleeping for %s before retrying\n", *retryAfterDuration)
		} else if rateLimitType == Primary {
			log.Printf("Primary rate limit (reset: %s) reached, sleeping for %s before retrying\n", resp.Header.Get(headers.XRateLimitResetKey), *retryAfterDuration)
			log.Printf("Rate limit information: %s: %s, %s: %s, %s: %s\n", headers.XRateLimitLimitKey, resp.Header.Get(headers.XRateLimitLimitKey), headers.XRateLimitUsedKey, resp.Header.Get(headers.XRateLimitUsedKey), headers.XRateLimitResourceKey, resp.Header.Get(headers.XRateLimitResourceKey))
		}
		time.Sleep(*retryAfterDuration)
		log.Printf("Retrying request after rate limit sleep\n")
		return handler.Intercept(pipeline, middlewareIndex, request)
	}

	return handler.retryRequest(ctx, pipeline, middlewareIndex, options, rateLimitType, request, resp)
}

// parseRateLimit parses rate-limit related headers and returns an appropriate
// time.Duration to retry the request after based on the header information.
// Much of this code was taken from the google/go-github library:
// see https://github.com/google/go-github/blob/0e3ab5807f0e9bc6ea690f1b49e94b78259f3681/github/github.go
// Note that "Retry-After" headers correspond to secondary rate limits and
// "x-ratelimit-reset" headers to primary rate limits.
// Docs for rate limit headers:
// https://docs.github.com/en/rest/using-the-rest-api/best-practices-for-using-the-rest-api?apiVersion=2022-11-28#handle-rate-limit-errors-appropriately
func parseRateLimit(r *netHttp.Response) (*time.Duration, error) {

	// "If the retry-after response header is present, you should not retry
	// your request until after that many seconds has elapsed."
	// (see docs link above)
	if v := r.Header.Get(headers.RetryAfterKey); v != "" {
		return parseRetryAfter(v)
	}

	// "If the x-ratelimit-remaining header is 0, you should not make another
	// request until after the time specified by the x-ratelimit-reset
	// header. The x-ratelimit-reset header is in UTC epoch seconds.""
	// (see docs link above)
	if v := r.Header.Get(headers.XRateLimitResetKey); v != "" {
		return parseXRateLimitReset(v)
	}

	return nil, nil
}

// parseRetryAfter parses the "Retry-After" header used for secondary
// rate limits.
func parseRetryAfter(retryAfterValue string) (*time.Duration, error) {
	if retryAfterValue == "" {
		return nil, fmt.Errorf("could not parse emtpy RetryAfter string")
	}

	retryAfterSeconds, err := strconv.ParseInt(retryAfterValue, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse retry-after header into duration: %v", err)
	}
	retryAfter := time.Duration(retryAfterSeconds) * time.Second
	return &retryAfter, nil
}

// parseXRateLimitReset parses the "x-ratelimit-reset" header used for primary
// rate limits
func parseXRateLimitReset(rateLimitResetValue string) (*time.Duration, error) {
	secondsSinceEpoch, err := strconv.ParseInt(rateLimitResetValue, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse x-ratelimit-reset header into duration: %v", err)
	}
	retryAfter := time.Until(time.Unix(secondsSinceEpoch, 0))
	return &retryAfter, nil
}
