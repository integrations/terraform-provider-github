package httpx

import (
	"net/http"
	"strconv"
	"strings"
	"time"
)

// retryableResponse returns if the given response to the given request can be retried with an optional wait duration. Responses that mean the request was rejected before being processed (429, 503 and a 403 with Retry-After) are always retryable; other server errors may have been processed so are only retryable for idempotent requests.
func retryableResponse(req *http.Request, resp *http.Response) (bool, *time.Duration) {
	var wait *time.Duration

	if raw := resp.Header.Get(headerRetryAfter); raw != "" {
		if d, ok := ParseRetryAfter(raw); ok {
			wait = &d
		}
	}

	switch resp.StatusCode {
	case http.StatusTooManyRequests, http.StatusServiceUnavailable:
		return true, wait
	case http.StatusInternalServerError, http.StatusBadGateway, http.StatusGatewayTimeout:
		return isIdempotent(req), wait
	case http.StatusForbidden:
		return wait != nil, wait
	default:
		return false, nil
	}
}

// ParseRetryAfter parses the Retry-After header value and returns the duration to wait before retrying. It supports both seconds and HTTP-date formats. If the value cannot be parsed it returns 0 and false.
func ParseRetryAfter(raw string) (time.Duration, bool) {
	raw = strings.TrimSpace(raw)

	if raw != "" && strings.Trim(raw, "0123456789") == "" {
		seconds, err := strconv.ParseUint(raw, 10, 64)
		if err != nil || seconds > maxRetryAfterSeconds {
			return 0, false
		}
		return time.Duration(seconds) * time.Second, true
	}

	t, err := http.ParseTime(raw)
	if err != nil {
		return 0, false
	}
	return max(time.Until(t), 0), true
}
