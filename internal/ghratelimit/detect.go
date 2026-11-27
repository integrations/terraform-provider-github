package ghratelimit

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"strings"
)

// isPrimaryRateLimit checks if the response indicates a primary rate limit has been hit, based on the status code and X-Ratelimit-Remaining header. It returns true if a primary rate limit is detected, otherwise false.
func isPrimaryRateLimit(resp *http.Response) bool {
	if resp.StatusCode != http.StatusForbidden && resp.StatusCode != http.StatusTooManyRequests {
		return false
	}

	return resp.Header.Get(headerRateRemaining) == "0"
}

// isSecondaryRateLimit checks if the response indicates a secondary rate limit has been hit, based on the status code, Retry-After header, and response body content. It returns true if a secondary rate limit is detected, otherwise false.
func isSecondaryRateLimit(resp *http.Response) bool {
	if resp.StatusCode != http.StatusForbidden && resp.StatusCode != http.StatusTooManyRequests {
		return false
	}

	if resp.Header.Get(headerRetryAfter) != "" {
		return true
	}

	if resp.StatusCode == http.StatusTooManyRequests && resp.Header.Get(headerRateRemaining) != "0" {
		return true
	}

	return hasSecondaryRateLimitBody(resp)
}

// hasSecondaryRateLimitBody checks if the response body contains indications of a secondary rate limit being hit, by reading the first 1024 bytes of the body and looking for specific phrases. It returns true if such indications are found, otherwise false.
func hasSecondaryRateLimitBody(resp *http.Response) bool {
	if resp.Body == nil {
		return false
	}

	buf := make([]byte, 1024)
	n, err := io.ReadFull(resp.Body, buf)

	resp.Body = struct {
		io.Reader
		io.Closer
	}{
		Reader: io.MultiReader(bytes.NewReader(buf[:n]), resp.Body),
		Closer: resp.Body,
	}

	if err != nil && !errors.Is(err, io.EOF) && !errors.Is(err, io.ErrUnexpectedEOF) {
		return false
	}

	content := strings.ToLower(string(buf[:n]))
	return strings.Contains(content, "secondary rate limit") || strings.Contains(content, "abuse detection mechanism")
}
