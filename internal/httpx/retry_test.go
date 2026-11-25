package httpx

import (
	"bytes"
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"sync/atomic"
	"testing"
	"testing/synctest"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestWithRetryMaxRetries(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name string
		m    int
		want int
	}{
		{
			name: "positive",
			m:    3,
			want: 3,
		},
		{
			name: "zero",
			m:    0,
			want: 0,
		},
		{
			name: "negative",
			m:    -1,
			want: 0,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			o := retryOptions{}
			opt := WithRetryMaxRetries(tt.m)
			opt(&o)

			if o.maxRetries != tt.want {
				t.Errorf("got %v, want %v", o.maxRetries, tt.want)
			}
		})
	}
}

func TestWithRetryWait(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name    string
		minWait time.Duration
		maxWait time.Duration
		wantMin time.Duration
		wantMax time.Duration
	}{
		{
			name:    "zero",
			minWait: 0,
			maxWait: 0,
			wantMin: 0,
			wantMax: 0,
		},
		{
			name:    "valid",
			minWait: 1 * time.Second,
			maxWait: 2 * time.Second,
			wantMin: 1 * time.Second,
			wantMax: 2 * time.Second,
		},
		{
			name:    "min_greater_than_max",
			minWait: 2 * time.Second,
			maxWait: 1 * time.Second,
			wantMin: 2 * time.Second,
			wantMax: 2 * time.Second,
		},
		{
			name:    "negative_min",
			minWait: -1 * time.Second,
			maxWait: 1 * time.Second,
			wantMin: 0,
			wantMax: 1 * time.Second,
		},
		{
			name:    "negative_max",
			minWait: 0,
			maxWait: -1 * time.Second,
			wantMin: 0,
			wantMax: 0,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			o := retryOptions{}
			opt := WithRetryWait(tt.minWait, tt.maxWait)
			opt(&o)

			if o.waitMin != tt.wantMin {
				t.Errorf("waitMin: got %v, want %v", o.waitMin, tt.wantMin)
			}

			if o.waitMax != tt.wantMax {
				t.Errorf("waitMax: got %v, want %v", o.waitMax, tt.wantMax)
			}
		})
	}
}

func TestWithRetryJitter(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name   string
		jitter time.Duration
		want   time.Duration
	}{
		{
			name:   "positive",
			jitter: 1 * time.Second,
			want:   1 * time.Second,
		},
		{
			name:   "zero",
			jitter: 0,
			want:   0,
		},
		{
			name:   "negative",
			jitter: -1 * time.Second,
			want:   0,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			o := retryOptions{}
			WithRetryJitter(tt.jitter)(&o)

			if o.jitter != tt.want {
				t.Errorf("got %v, want %v", o.jitter, tt.want)
			}
		})
	}
}

func TestWithRetryNoRetryErrors(t *testing.T) {
	t.Parallel()

	err1 := errors.New("error1")
	err2 := errors.New("error2")

	for _, tt := range []struct {
		name string
		errs []error
	}{
		{
			name: "nil_errors",
		},
		{
			name: "no_errors",
			errs: []error{},
		},
		{
			name: "single_error",
			errs: []error{err1},
		},
		{
			name: "multiple_errors",
			errs: []error{err1, err2},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			o := retryOptions{}
			WithRetryNoRetryErrors(tt.errs...)(&o)

			if diff := cmp.Diff(tt.errs, o.noRetryErrors, cmpopts.EquateErrors()); diff != "" {
				t.Fatalf("unexpected errors (-want +got):\n%s", diff)
			}
		})
	}
}

func TestNewRetryTransport(t *testing.T) {
	t.Parallel()

	t.Run("creates_transport_with_defaults", func(t *testing.T) {
		t.Parallel()

		tr := &dummyRoundTripper{}
		got := NewRetryTransport(tr)

		if got == nil {
			t.Fatal("transport: got nil, want non-nil")
		}

		rt, ok := got.(*retryTransport)
		if !ok {
			t.Fatal("transport: got non-retryTransport, want retryTransport")
		}

		if rt.inner != tr {
			t.Fatal("inner: got other transport, want given transport")
		}

		if rt.opts.maxRetries != 3 {
			t.Errorf("maxRetries: got %v, want 3", rt.opts.maxRetries)
		}

		if rt.opts.waitMin != time.Second {
			t.Errorf("waitMin: got %v, want 1s", rt.opts.waitMin)
		}

		if rt.opts.waitMax != time.Minute {
			t.Errorf("waitMax: got %v, want 1m0s", rt.opts.waitMax)
		}

		if rt.opts.jitter != 100*time.Millisecond {
			t.Errorf("jitter: got %v, want 100ms", rt.opts.jitter)
		}
	})

	t.Run("creates_transport_with_options", func(t *testing.T) {
		t.Parallel()

		tr := &dummyRoundTripper{}
		got := NewRetryTransport(tr,
			WithRetryMaxRetries(5),
			WithRetryWait(2*time.Second, 20*time.Second),
			WithRetryJitter(500*time.Millisecond),
		)

		if got == nil {
			t.Fatal("transport: got nil, want non-nil")
		}

		rt, ok := got.(*retryTransport)
		if !ok {
			t.Fatal("transport: got non-retryTransport, want retryTransport")
		}

		if rt.inner != tr {
			t.Fatal("inner: got other transport, want given transport")
		}

		if rt.opts.maxRetries != 5 {
			t.Errorf("maxRetries: got %v, want 5", rt.opts.maxRetries)
		}

		if rt.opts.waitMin != 2*time.Second {
			t.Errorf("waitMin: got %v, want 2s", rt.opts.waitMin)
		}

		if rt.opts.waitMax != 20*time.Second {
			t.Errorf("waitMax: got %v, want 20s", rt.opts.waitMax)
		}

		if rt.opts.jitter != 500*time.Millisecond {
			t.Errorf("jitter: got %v, want 500ms", rt.opts.jitter)
		}
	})

	t.Run("returns_inner_if_max_retries_zero", func(t *testing.T) {
		t.Parallel()

		tr := &dummyRoundTripper{}
		got := NewRetryTransport(tr, WithRetryMaxRetries(0))

		if got != tr {
			t.Fatal("transport: got non-inner transport, want inner transport")
		}
	})
}

func Test_retryTransport(t *testing.T) {
	t.Parallel()

	t.Run("RoundTrip", func(t *testing.T) {
		for _, tt := range []struct {
			name           string
			opts           retryOptions
			method         string
			failures       []fakeHTTPResponse
			wantCalls      int
			wantStatusCode int
		}{
			{
				name:           "succeeds_if_no_error_rsponse",
				opts:           retryOptions{maxRetries: 1, waitMin: time.Millisecond, waitMax: time.Millisecond},
				method:         http.MethodGet,
				wantCalls:      1,
				wantStatusCode: http.StatusOK,
			},
			{
				name:           "succeeds_if_retry_successful",
				opts:           retryOptions{maxRetries: 1, waitMin: time.Millisecond, waitMax: time.Millisecond},
				method:         http.MethodGet,
				failures:       []fakeHTTPResponse{{statusCode: http.StatusServiceUnavailable, body: []byte("error")}},
				wantCalls:      2,
				wantStatusCode: http.StatusOK,
			},
			{
				name:           "succeeds_with_close_connection_if_retry_successful",
				opts:           retryOptions{maxRetries: 1, waitMin: time.Millisecond, waitMax: time.Millisecond},
				method:         http.MethodGet,
				failures:       []fakeHTTPResponse{{statusCode: http.StatusServiceUnavailable, headers: map[string]string{"Connection": "close"}, body: []byte("error")}},
				wantCalls:      2,
				wantStatusCode: http.StatusOK,
			},
			{
				name:           "fails_if_retry_unsuccessful",
				opts:           retryOptions{maxRetries: 1, waitMin: time.Millisecond, waitMax: time.Millisecond},
				method:         http.MethodGet,
				failures:       []fakeHTTPResponse{{statusCode: http.StatusServiceUnavailable, body: []byte("error")}, {statusCode: http.StatusServiceUnavailable, body: []byte("error")}},
				wantCalls:      2,
				wantStatusCode: http.StatusServiceUnavailable,
			},
			{
				name:           "fails_fast_on_non_retry_status_code",
				opts:           retryOptions{maxRetries: 1, waitMin: time.Millisecond, waitMax: time.Millisecond},
				method:         http.MethodGet,
				failures:       []fakeHTTPResponse{{statusCode: http.StatusBadRequest, body: []byte("error")}},
				wantCalls:      1,
				wantStatusCode: http.StatusBadRequest,
			},
			{
				name:           "retries_server_error_for_idempotent_method",
				method:         http.MethodPut,
				opts:           retryOptions{maxRetries: 1, waitMin: time.Millisecond, waitMax: time.Millisecond},
				failures:       []fakeHTTPResponse{{statusCode: http.StatusInternalServerError, body: []byte("error")}},
				wantCalls:      2,
				wantStatusCode: http.StatusOK,
			},
			{
				name:           "fails_fast_on_server_error_for_non_idempotent_method",
				method:         http.MethodPost,
				opts:           retryOptions{maxRetries: 1, waitMin: time.Millisecond, waitMax: time.Millisecond},
				failures:       []fakeHTTPResponse{{statusCode: http.StatusInternalServerError, body: []byte("error")}},
				wantCalls:      1,
				wantStatusCode: http.StatusInternalServerError,
			},
			{
				name:           "retries_forbidden_with_retry_after",
				opts:           retryOptions{maxRetries: 1, waitMin: time.Millisecond, waitMax: time.Millisecond},
				failures:       []fakeHTTPResponse{{statusCode: http.StatusForbidden, headers: map[string]string{"Retry-After": "0"}, body: []byte("error")}},
				wantCalls:      2,
				wantStatusCode: http.StatusOK,
			},
			{
				name:           "returns_response_if_retry_after_exceeds_wait_max",
				opts:           retryOptions{maxRetries: 1, waitMin: time.Millisecond, waitMax: time.Millisecond},
				failures:       []fakeHTTPResponse{{statusCode: http.StatusTooManyRequests, headers: map[string]string{"Retry-After": "5"}, body: []byte("error")}},
				wantCalls:      1,
				wantStatusCode: http.StatusTooManyRequests,
			},
		} {
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				reqBody := "REQUEST"
				wantBody := "PASS"
				called := atomic.Int32{}

				testHandler := func(req *http.Request) (*http.Response, error) {
					call := int(called.Add(1))

					by, err := io.ReadAll(req.Body)
					if err != nil {
						t.Fatalf("failed to read request body: %v", err)
					}

					if string(by) != reqBody {
						t.Fatalf("request body: got %q, want %q", string(by), reqBody)
					}

					resp := &http.Response{
						Header: make(http.Header),
					}

					if call <= len(tt.failures) {
						f := tt.failures[call-1]

						for k, v := range f.headers {
							resp.Header.Set(k, v)
						}

						resp.StatusCode = f.statusCode
						resp.Body = io.NopCloser(bytes.NewReader(f.body))

						return resp, nil
					}

					resp.StatusCode = http.StatusOK
					resp.Body = io.NopCloser(bytes.NewReader([]byte(wantBody)))

					return resp, nil
				}

				inner := &testRoundTripper{f: testHandler}
				tr := &retryTransport{
					inner: inner,
					opts:  tt.opts,
				}

				req := mustCreateRequest(t, nil, tt.method, "https://example.com", bytes.NewReader([]byte(reqBody)))

				resp, err := tr.RoundTrip(req)
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}

				if resp == nil {
					t.Fatal("response: got nil, want non-nil")
				}

				if resp.StatusCode != tt.wantStatusCode {
					t.Fatalf("status code: got %d, want %d", resp.StatusCode, tt.wantStatusCode)
				}

				gotCalls := int(called.Load())
				if gotCalls != tt.wantCalls {
					t.Fatalf("calls: got %d, want %d", called.Load(), tt.wantCalls)
				}
			})
		}

		t.Run("errors_if_context_canceled_early", func(t *testing.T) {
			t.Parallel()

			inner := &testRoundTripper{}
			tr := &retryTransport{
				inner: inner,
				opts: retryOptions{
					maxRetries: 3,
					waitMin:    time.Millisecond,
					waitMax:    time.Millisecond,
				},
			}

			ctx, cancel := context.WithCancel(t.Context())
			req := mustCreateRequest(t, ctx, http.MethodGet, "https://example.com", nil)
			cancel()

			_, err := tr.RoundTrip(req)
			if err == nil {
				t.Fatal("got nil, want error")
			}

			if !errors.Is(err, context.Canceled) {
				t.Fatalf("got %v, want %v", err, context.Canceled)
			}
		})

		t.Run("closes_request_body_if_context_canceled_early", func(t *testing.T) {
			t.Parallel()

			inner := &testRoundTripper{}
			tr := &retryTransport{
				inner: inner,
				opts: retryOptions{
					maxRetries: 3,
					waitMin:    time.Millisecond,
					waitMax:    time.Millisecond,
				},
			}

			ctx, cancel := context.WithCancel(t.Context())
			body := &closeRecorder{}
			req := mustCreateRequest(t, ctx, http.MethodPost, "https://example.com", body)
			cancel()

			_, err := tr.RoundTrip(req)
			if !errors.Is(err, context.Canceled) {
				t.Fatalf("got %v, want %v", err, context.Canceled)
			}

			if got := inner.called.Load(); got != 0 {
				t.Errorf("inner transport calls: got %d, want 0", got)
			}

			if !body.closed.Load() {
				t.Error("request body closed: got false, want true")
			}
		})

		t.Run("errors_if_context_canceled_in_flight", func(t *testing.T) {
			t.Parallel()

			synctest.Test(t, func(t *testing.T) {
				inner := &testRoundTripper{f: func(*http.Request) (*http.Response, error) {
					return &http.Response{StatusCode: http.StatusServiceUnavailable}, nil
				}}
				tr := &retryTransport{inner: inner, opts: retryOptions{maxRetries: 3, waitMin: time.Second, waitMax: time.Second}}

				ctx, cancel := context.WithCancel(t.Context())
				time.AfterFunc(500*time.Millisecond, cancel)

				_, err := tr.RoundTrip(mustCreateRequest(t, ctx, http.MethodGet, "https://example.com", nil))
				if !errors.Is(err, context.Canceled) {
					t.Fatalf("got %v, want %v", err, context.Canceled)
				}
			})
		})

		t.Run("closes_response_body_if_context_canceled_in_flight", func(t *testing.T) {
			t.Parallel()

			synctest.Test(t, func(t *testing.T) {
				rec := &closeRecorder{}
				inner := &testRoundTripper{f: func(*http.Request) (*http.Response, error) {
					return &http.Response{StatusCode: http.StatusServiceUnavailable, Body: rec}, nil
				}}
				tr := &retryTransport{inner: inner, opts: retryOptions{maxRetries: 3, waitMin: time.Second, waitMax: time.Second}}

				ctx, cancel := context.WithCancel(t.Context())
				time.AfterFunc(500*time.Millisecond, cancel)

				resp, err := tr.RoundTrip(mustCreateRequest(t, ctx, http.MethodGet, "https://example.com", nil))
				if !errors.Is(err, context.Canceled) {
					t.Fatalf("got %v, want %v", err, context.Canceled)
				}

				if resp != nil {
					t.Fatal("response: got non-nil, want nil")
				}

				if !rec.closed.Load() {
					t.Error("response body closed: got false, want true")
				}
			})
		})

		t.Run("closes_response_body_if_inner_returns_response_and_error", func(t *testing.T) {
			t.Parallel()

			errTest := errors.New("test error")

			rec := &closeRecorder{}
			inner := &testRoundTripper{f: func(*http.Request) (*http.Response, error) {
				return &http.Response{Body: rec}, errTest
			}}
			tr := &retryTransport{inner: inner, opts: retryOptions{maxRetries: 1, waitMin: time.Millisecond, waitMax: time.Millisecond, noRetryErrors: []error{errTest}}}

			resp, err := tr.RoundTrip(mustCreateRequest(t, nil, http.MethodGet, "https://example.com", nil))
			if !errors.Is(err, errTest) {
				t.Fatalf("got %v, want %v", err, errTest)
			}

			if resp != nil {
				t.Fatal("response: got non-nil, want nil")
			}

			if !rec.closed.Load() {
				t.Error("response body closed: got false, want true")
			}
		})

		t.Run("errors_if_rewinding_body_fails_on_retry", func(t *testing.T) {
			t.Parallel()

			inner := &testRoundTripper{resp: &http.Response{StatusCode: http.StatusServiceUnavailable}}
			tr := &retryTransport{inner: inner, opts: retryOptions{maxRetries: 1, waitMin: time.Millisecond, waitMax: time.Millisecond}}

			req := mustCreateRequest(t, nil, http.MethodPost, "https://example.com", strings.NewReader("TEST"))
			req.GetBody = func() (io.ReadCloser, error) {
				return nil, errors.New("boom")
			}

			resp, err := tr.RoundTrip(req)
			if err == nil {
				t.Fatal("got nil, want error")
			}

			if !strings.Contains(err.Error(), "failed to rewind request body") {
				t.Fatalf("got %v, want it to contain \"failed to rewind request body\"", err)
			}

			if resp != nil {
				t.Fatal("response: got non-nil, want nil")
			}

			if got := inner.called.Load(); got != 1 {
				t.Fatalf("calls: got %d, want 1", got)
			}
		})

		t.Run("errors_if_no_retry_error_returned_from_inner", func(t *testing.T) {
			t.Parallel()

			errTest := errors.New("test error")

			inner := &testRoundTripper{err: errTest}
			tr := &retryTransport{
				inner: inner,
				opts: retryOptions{
					maxRetries:    3,
					waitMin:       time.Millisecond,
					waitMax:       time.Millisecond,
					noRetryErrors: []error{errTest},
				},
			}

			req := mustCreateRequest(t, nil, http.MethodGet, "https://example.com", nil)

			_, err := tr.RoundTrip(req)
			if err == nil {
				t.Fatal("got nil, want error")
			}

			if !errors.Is(err, errTest) {
				t.Fatalf("got %v, want %v", err, errTest)
			}
		})

		t.Run("retries_idempotent_request_after_transport_timeout", func(t *testing.T) {
			t.Parallel()

			errTimeout := fmt.Errorf("timeout awaiting response headers: %w", context.DeadlineExceeded)

			var calls atomic.Int32
			inner := &testRoundTripper{f: func(*http.Request) (*http.Response, error) {
				if calls.Add(1) == 1 {
					return nil, errTimeout
				}

				return &http.Response{StatusCode: http.StatusOK, Body: http.NoBody}, nil
			}}
			tr := &retryTransport{inner: inner, opts: retryOptions{maxRetries: 1, waitMin: time.Millisecond, waitMax: time.Millisecond}}

			resp, err := tr.RoundTrip(mustCreateRequest(t, nil, http.MethodGet, "https://example.com", nil))
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				t.Fatalf("status code: got %d, want %d", resp.StatusCode, http.StatusOK)
			}

			if got := calls.Load(); got != 2 {
				t.Fatalf("calls: got %d, want 2", got)
			}
		})

		t.Run("does_not_retry_non_idempotent_request_after_transport_timeout", func(t *testing.T) {
			t.Parallel()

			errTimeout := fmt.Errorf("timeout awaiting response headers: %w", context.DeadlineExceeded)

			inner := &testRoundTripper{err: errTimeout}
			tr := &retryTransport{inner: inner, opts: retryOptions{maxRetries: 1, waitMin: time.Millisecond, waitMax: time.Millisecond}}

			_, err := tr.RoundTrip(mustCreateRequest(t, nil, http.MethodPost, "https://example.com", nil))
			if !errors.Is(err, errTimeout) {
				t.Fatalf("got %v, want %v", err, errTimeout)
			}

			if got := inner.called.Load(); got != 1 {
				t.Fatalf("calls: got %d, want 1", got)
			}
		})

		t.Run("returns_response_if_retry_after_exceeds_context_deadline", func(t *testing.T) {
			t.Parallel()

			inner := &testRoundTripper{resp: &http.Response{
				StatusCode: http.StatusTooManyRequests,
				Header:     http.Header{"Retry-After": []string{"30"}},
			}}
			tr := &retryTransport{
				inner: inner,
				opts: retryOptions{
					maxRetries: 3,
					waitMin:    time.Millisecond,
					waitMax:    time.Minute,
				},
			}

			ctx, cancel := context.WithTimeout(t.Context(), time.Second)
			defer cancel()

			req := mustCreateRequest(t, ctx, http.MethodGet, "https://example.com", nil)

			resp, err := tr.RoundTrip(req)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if resp == nil {
				t.Fatal("response: got nil, want non-nil")
			}

			if resp.StatusCode != http.StatusTooManyRequests {
				t.Fatalf("status code: got %d, want 429", resp.StatusCode)
			}

			if got := inner.called.Load(); got != 1 {
				t.Fatalf("calls: got %d, want 1", got)
			}
		})

		t.Run("returns_response_if_cant_rewind_body", func(t *testing.T) {
			t.Parallel()

			inner := &testRoundTripper{resp: &http.Response{StatusCode: http.StatusServiceUnavailable}}
			tr := &retryTransport{
				inner: inner,
				opts: retryOptions{
					maxRetries: 3,
					waitMin:    time.Millisecond,
					waitMax:    time.Millisecond,
				},
			}

			req := mustCreateRequest(t, nil, http.MethodPost, "https://example.com", io.NopCloser(&nonSeekableReader{r: strings.NewReader("TEST")}))

			resp, err := tr.RoundTrip(req)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if resp == nil {
				t.Fatal("response: got nil, want non-nil")
			}

			if resp.StatusCode != http.StatusServiceUnavailable {
				t.Fatalf("status code: got %d, want 503", resp.StatusCode)
			}

			if got := inner.called.Load(); got != 1 {
				t.Fatalf("calls: got %d, want 1", got)
			}
		})

		t.Run("retries_dial_error_for_non_idempotent_method", func(t *testing.T) {
			t.Parallel()

			var calls atomic.Int32
			inner := &testRoundTripper{f: func(*http.Request) (*http.Response, error) {
				if calls.Add(1) == 1 {
					return nil, &net.OpError{Op: "dial", Err: errors.New("refused")}
				}
				return &http.Response{StatusCode: http.StatusOK}, nil
			}}
			tr := &retryTransport{inner: inner, opts: retryOptions{maxRetries: 1, waitMin: time.Millisecond, waitMax: time.Millisecond}}

			resp, err := tr.RoundTrip(mustCreateRequest(t, nil, http.MethodPost, "https://example.com", nil))
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if resp.StatusCode != http.StatusOK || calls.Load() != 2 {
				t.Fatalf("got status %d after %d calls, want 200 after 2 calls", resp.StatusCode, calls.Load())
			}
		})
	})

	t.Run("retryableError", func(t *testing.T) {
		t.Parallel()

		errTest := errors.New("test error")
		errTimeout := fmt.Errorf("timeout awaiting response headers: %w", context.DeadlineExceeded)

		canceledCtx, cancel := context.WithCancel(t.Context())
		cancel()

		expiredCtx, cancelExpired := context.WithDeadline(t.Context(), time.Unix(0, 0))
		t.Cleanup(cancelExpired)

		for _, tt := range []struct {
			name          string
			noRetryErrors []error
			ctx           context.Context
			method        string
			err           error
			want          bool
		}{
			{
				name:   "nil",
				method: http.MethodGet,
				err:    nil,
				want:   false,
			},
			{
				name:   "context_canceled",
				ctx:    canceledCtx,
				method: http.MethodGet,
				err:    context.Canceled,
				want:   false,
			},
			{
				name:   "context_deadline_exceeded",
				ctx:    expiredCtx,
				method: http.MethodGet,
				err:    context.DeadlineExceeded,
				want:   false,
			},
			{
				name:   "context_done_with_other_error",
				ctx:    canceledCtx,
				method: http.MethodGet,
				err:    fmt.Errorf("test"),
				want:   false,
			},
			{
				name:   "context_done_with_dial_error",
				ctx:    canceledCtx,
				method: http.MethodPost,
				err:    &net.OpError{Op: "dial", Err: errors.New("refused")},
				want:   false,
			},
			{
				name:   "transport_timeout_idempotent",
				method: http.MethodGet,
				err:    errTimeout,
				want:   true,
			},
			{
				name:   "transport_timeout_not_idempotent",
				method: http.MethodPost,
				err:    errTimeout,
				want:   false,
			},
			{
				name:   "dial_timeout_not_idempotent",
				method: http.MethodPost,
				err:    &net.OpError{Op: "dial", Err: errTimeout},
				want:   true,
			},
			{
				name:          "rate_limit_wait_exceeded",
				noRetryErrors: []error{errTest},
				method:        http.MethodGet,
				err:           errTest,
				want:          false,
			},
			{
				name:          "wrapped_rate_limit_wait_exceeded",
				noRetryErrors: []error{errTest},
				method:        http.MethodGet,
				err:           fmt.Errorf("request blocked: %w", errTest),
				want:          false,
			},
			{
				name:   "tls_certificate_verification_error",
				method: http.MethodGet,
				err:    &tls.CertificateVerificationError{},
				want:   false,
			},
			{
				name:   "wrapped_tls_certificate_verification_error",
				method: http.MethodGet,
				err:    fmt.Errorf("tls error: %w", &tls.CertificateVerificationError{}),
				want:   false,
			},
			{
				name:   "other_error_not_idempotent",
				method: http.MethodPost,
				err:    fmt.Errorf("test"),
				want:   false,
			},
			{
				name:   "other_error_idempotent",
				method: http.MethodGet,
				err:    fmt.Errorf("test"),
				want:   true,
			},
			{
				name:   "dial_error_not_idempotent",
				method: http.MethodPost,
				err:    &net.OpError{Op: "dial", Err: errors.New("refused")},
				want:   true,
			},
			{
				name:   "wrapped_dial_error_not_idempotent",
				method: http.MethodPost,
				err:    fmt.Errorf("transport: %w", &net.OpError{Op: "dial", Err: errors.New("refused")}),
				want:   true,
			},
			{
				name:   "read_error_not_idempotent",
				method: http.MethodPost,
				err:    &net.OpError{Op: "read", Err: errors.New("reset")},
				want:   false,
			},
			{
				name:   "read_error_idempotent",
				method: http.MethodGet,
				err:    &net.OpError{Op: "read", Err: errors.New("reset")},
				want:   true,
			},
		} {
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				ctx := tt.ctx
				if ctx == nil {
					ctx = t.Context()
				}

				tr := &retryTransport{opts: retryOptions{noRetryErrors: tt.noRetryErrors}}
				req := mustCreateRequest(t, ctx, tt.method, "https://example.com", nil)

				if got := tr.retryableError(req, tt.err); got != tt.want {
					t.Errorf("got %v, want %v", got, tt.want)
				}
			})
		}
	})
}

func Test_newBackoff(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name string
		opts retryOptions
	}{
		{
			name: "with_options",
			opts: retryOptions{
				maxRetries: 3,
				waitMin:    10 * time.Millisecond,
				waitMax:    time.Second,
				jitter:     time.Millisecond,
			},
		},
		{
			name: "without_options",
			opts: retryOptions{},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := newBackoff(tt.opts)

			if got == nil {
				t.Fatal("got nil, want non-nil")
			}
		})
	}
}

func Test_retryAfterBackoff(t *testing.T) {
	t.Parallel()

	t.Run("overrideMin", func(t *testing.T) {
		t.Parallel()

		for _, tt := range []struct {
			name            string
			currentOverride time.Duration
			override        time.Duration
			want            time.Duration
		}{
			{
				name:     "updates_empty_min",
				override: time.Second,
				want:     time.Second,
			},
			{
				name:            "updates_set_min",
				currentOverride: time.Second,
				override:        2 * time.Second,
				want:            2 * time.Second,
			},
			{
				name:            "clears_min_if_negative_duration_passed_in",
				currentOverride: time.Second,
				override:        -2 * time.Second,
				want:            0,
			},
		} {
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				bo := retryAfterBackoff{override: tt.currentOverride}
				bo.overrideMin(tt.override)

				got := bo.override
				if got != tt.want {
					t.Fatalf("got %v, want %v", got, tt.want)
				}
			})
		}
	})

	t.Run("Next", func(t *testing.T) {
		t.Parallel()

		t.Run("stops_if_inner_backoff_stops", func(t *testing.T) {
			t.Parallel()

			inner := &testBackoff{next: time.Second, stop: true}
			bo := retryAfterBackoff{inner: inner}

			got, stop := bo.Next()
			if !stop {
				t.Fatal("stop: got false, want true")
			}

			if got != 0 {
				t.Fatalf("wait: got %v, want 0", got)
			}
		})

		t.Run("returns_override_if_greater_than_inner_next", func(t *testing.T) {
			t.Parallel()

			override := 2 * time.Second
			inner := &testBackoff{next: time.Second, stop: false}
			bo := retryAfterBackoff{inner: inner, override: override}

			got, stop := bo.Next()
			if stop {
				t.Fatal("stop: got true, want false")
			}

			if got != override {
				t.Fatalf("wait: got %v, want %v", got, override)
			}
		})

		t.Run("returns_inner_next_if_greater_than_override", func(t *testing.T) {
			t.Parallel()

			override := time.Second
			inner := &testBackoff{next: 2 * time.Second, stop: false}
			bo := retryAfterBackoff{inner: inner, override: override}

			got, stop := bo.Next()
			if stop {
				t.Fatal("stop: got true, want false")
			}

			if got != inner.next {
				t.Fatalf("wait: got %v, want %v", got, inner.next)
			}
		})

		t.Run("clears_override_after_use", func(t *testing.T) {
			t.Parallel()

			override := 2 * time.Second
			inner := &testBackoff{next: time.Second, stop: false}
			bo := retryAfterBackoff{inner: inner, override: override}

			got, stop := bo.Next()
			if stop {
				t.Fatal("stop: got true, want false")
			}

			if got != override {
				t.Fatalf("wait: got %v, want %v", got, override)
			}

			got, stop = bo.Next()
			if stop {
				t.Fatal("stop: got true, want false")
			}

			if got != inner.next {
				t.Fatalf("wait: got %v, want %v", got, inner.next)
			}
		})
	})
}

func Test_deadlineBackoff(t *testing.T) {
	t.Parallel()

	t.Run("Next", func(t *testing.T) {
		t.Parallel()

		t.Run("stops_if_inner_backoff_stops", func(t *testing.T) {
			t.Parallel()

			inner := &testBackoff{next: time.Second, stop: true}
			bo := deadlineBackoff{inner: inner, deadline: time.Now().Add(time.Second)}

			got, stop := bo.Next()
			if !stop {
				t.Fatal("stop: got false, want true")
			}

			if got != 0 {
				t.Fatalf("wait: got %v, want 0", got)
			}
		})

		t.Run("stops_if_next_wait_exceeds_deadline", func(t *testing.T) {
			t.Parallel()

			inner := &testBackoff{next: time.Second, stop: false}
			bo := deadlineBackoff{inner: inner, deadline: time.Now().Add(500 * time.Millisecond)}

			got, stop := bo.Next()
			if !stop {
				t.Fatal("stop: got false, want true")
			}

			if got != 0 {
				t.Fatalf("wait: got %v, want 0", got)
			}
		})

		t.Run("returns_inner_next_if_within_deadline", func(t *testing.T) {
			t.Parallel()

			inner := &testBackoff{next: time.Second, stop: false}
			bo := deadlineBackoff{inner: inner, deadline: time.Now().Add(2 * time.Second)}

			got, stop := bo.Next()
			if stop {
				t.Fatal("stop: got true, want false")
			}

			if got != inner.next {
				t.Fatalf("wait: got %v, want %v", got, inner.next)
			}
		})
	})
}
