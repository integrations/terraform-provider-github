package github

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/google/go-github/v44/github"
)

const (
	ctxEtag = ctxEtagType("etag")
	ctxId   = ctxIdType("id")
)

// ctxIdType is used to avoid collisions between packages using context
type ctxIdType string

// ctxEtagType is used to avoid collisions between packages using context
type ctxEtagType string

// etagTransport allows saving API quota by passing previously stored Etag
// available via context to request headers
type etagTransport struct {
	transport http.RoundTripper
}

func (ett *etagTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	ctx := req.Context()

	etag := ctx.Value(ctxEtag)
	if v, ok := etag.(string); ok && v != "" {
		req.Header.Set("If-None-Match", v)
	}

	return ett.transport.RoundTrip(req)
}

func NewEtagTransport(rt http.RoundTripper) *etagTransport {
	return &etagTransport{transport: rt}
}

// RateLimitTransport implements GitHub's best practices
// for avoiding rate limits
// https://developer.github.com/v3/guides/best-practices-for-integrators/#dealing-with-abuse-rate-limits
type RateLimitTransport struct {
	transport        http.RoundTripper
	nextRequestDelay time.Duration
	writeDelay       time.Duration
	readDelay        time.Duration

	m sync.Mutex
}

func (rlt *RateLimitTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Make requests for a single user or client ID serially
	// This is also necessary for safely saving
	// and restoring bodies between retries below
	rlt.lock(req)

	// Sleep for the delay that the last request defined. This delay might be different
	// for read and write requests. See isWriteMethod for the distinction between them.
	if rlt.nextRequestDelay > 0 {
		log.Printf("[DEBUG] Sleeping %s between operations", rlt.nextRequestDelay)
		time.Sleep(rlt.nextRequestDelay)
	}

	rlt.nextRequestDelay = rlt.calculateNextDelay(req.Method)

	resp, err := rlt.transport.RoundTrip(req)
	if err != nil {
		rlt.unlock(req)
		return resp, err
	}

	// Make response body accessible for retries & debugging
	// (work around bug in GitHub SDK)
	// See https://github.com/google/go-github/pull/986
	r1, r2, err := drainBody(resp.Body)
	if err != nil {
		return nil, err
	}
	resp.Body = r1
	ghErr := github.CheckResponse(resp)
	resp.Body = r2

	// When you have been limited, use the Retry-After response header to slow down.
	if arlErr, ok := ghErr.(*github.AbuseRateLimitError); ok {
		rlt.nextRequestDelay = 0
		retryAfter := arlErr.GetRetryAfter()
		log.Printf("[DEBUG] Abuse detection mechanism triggered, sleeping for %s before retrying",
			retryAfter)
		time.Sleep(retryAfter)
		rlt.unlock(req)
		return rlt.RoundTrip(req)
	}

	if rlErr, ok := ghErr.(*github.RateLimitError); ok {
		rlt.nextRequestDelay = 0
		retryAfter := time.Until(rlErr.Rate.Reset.Time)
		log.Printf("[DEBUG] Rate limit %d reached, sleeping for %s (until %s) before retrying",
			rlErr.Rate.Limit, retryAfter, time.Now().Add(retryAfter))
		time.Sleep(retryAfter)
		rlt.unlock(req)
		return rlt.RoundTrip(req)
	}

	rlt.unlock(req)

	return resp, nil
}

func (rlt *RateLimitTransport) lock(req *http.Request) {
	rlt.m.Lock()
}

func (rlt *RateLimitTransport) unlock(req *http.Request) {
	rlt.m.Unlock()
}

// calculateNextDelay returns a time.Duration specifying the backoff before the next request
// the actual value depends on the current method being a write or a read request
func (rlt *RateLimitTransport) calculateNextDelay(method string) time.Duration {
	if isWriteMethod(method) {
		return rlt.writeDelay
	}
	return rlt.readDelay
}

type RateLimitTransportOption func(*RateLimitTransport)

// NewRateLimitTransport takes in an http.RoundTripper and a variadic list of
// optional functions that modify the RateLimitTransport struct itself. This
// may be used to alter the write delay in between requests, for example.
func NewRateLimitTransport(rt http.RoundTripper, options ...RateLimitTransportOption) *RateLimitTransport {
	// Default to 1 second of write delay if none is provided
	// Default to no read delay if none is provided
	rlt := &RateLimitTransport{transport: rt, writeDelay: 1 * time.Second, readDelay: 0 * time.Second}

	for _, opt := range options {
		opt(rlt)
	}

	return rlt
}

// WithWriteDelay is used to set the write delay between requests
func WithWriteDelay(d time.Duration) RateLimitTransportOption {
	return func(rlt *RateLimitTransport) {
		rlt.writeDelay = d
	}
}

// WithReadDelay is used to set the delay between read requests
func WithReadDelay(d time.Duration) RateLimitTransportOption {
	return func(rlt *RateLimitTransport) {
		rlt.readDelay = d
	}
}

// drainBody reads all of b to memory and then returns two equivalent
// ReadClosers yielding the same bytes.
func drainBody(b io.ReadCloser) (r1, r2 io.ReadCloser, err error) {
	if b == http.NoBody {
		// No copying needed. Preserve the magic sentinel meaning of NoBody.
		return http.NoBody, http.NoBody, nil
	}
	var buf bytes.Buffer
	if _, err = buf.ReadFrom(b); err != nil {
		return nil, b, err
	}
	if err = b.Close(); err != nil {
		return nil, b, err
	}
	return ioutil.NopCloser(&buf), ioutil.NopCloser(bytes.NewReader(buf.Bytes())), nil
}

func isWriteMethod(method string) bool {
	switch method {
	case "POST", "PATCH", "PUT", "DELETE":
		return true
	}
	return false
}
