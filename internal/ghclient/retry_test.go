package ghclient

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"sync/atomic"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func Test_newRetryTransport(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name    string
		opts    RetryOptions
		wantErr *string
	}{
		{
			name: "empty_opts_returns_inner",
			opts: RetryOptions{},
		},
		{
			name:    "invalid_opts_errors",
			opts:    RetryOptions{Max: 1, WaitMin: 0},
			wantErr: new("wait min must be greater than 0"),
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			inner := &testRoundTripper{}

			got, err := newRetryTransport(inner, tt.opts)
			if err != nil {
				if tt.wantErr == nil {
					t.Fatalf("unexpected error: %s", err)
				}
				if !regexp.MustCompile(regexp.QuoteMeta(*tt.wantErr)).MatchString(err.Error()) {
					t.Fatalf("unexpected error: want %s, got: %s", *tt.wantErr, err)
				}
				return
			}

			if tt.wantErr != nil {
				t.Fatalf("expected error: %s, got: %v", *tt.wantErr, err)
			}

			if got == nil {
				t.Fatal("expected non-nil transport")
			}

			if tt.opts.Max == 0 {
				if got != inner {
					t.Error("expected transport to be the inner transport when max retries is 0")
				}
				return
			}

			retry, ok := got.(*retryTransport)
			if !ok {
				t.Fatal("expected retryTransport")
			}

			if retry.inner != inner {
				t.Error("expected inner transport to be set")
			}

			if diff := cmp.Diff(retry.opts, tt.opts); diff != "" {
				t.Fatalf("got opts %+v, want %+v: %s", retry.opts, tt.opts, diff)
			}
		})
	}
}

func Test_retryTransport_RoundTrip(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name           string
		retryMax       int
		failures       []fakeHTTPResponse
		cancel         bool
		wantStatusCode int
		wantErr        *string
	}{
		{
			name:    "errors_if_context_canceled",
			cancel:  true,
			wantErr: new("context canceled"),
		},
		{
			name:           "no_failures",
			retryMax:       3,
			wantStatusCode: http.StatusOK,
		},
		{
			name:           "no_failures_no_retries",
			retryMax:       0,
			wantStatusCode: http.StatusOK,
		},
		{
			name:           "failure_no_error",
			retryMax:       0,
			failures:       []fakeHTTPResponse{{statusCode: http.StatusInternalServerError}},
			wantStatusCode: http.StatusInternalServerError,
		},
		{
			name:           "retries_until_success",
			retryMax:       3,
			failures:       []fakeHTTPResponse{{statusCode: http.StatusInternalServerError}, {statusCode: http.StatusInternalServerError}},
			wantStatusCode: http.StatusOK,
		},
		{
			name:           "retries_until_failure",
			retryMax:       1,
			failures:       []fakeHTTPResponse{{statusCode: http.StatusInternalServerError}, {statusCode: http.StatusInternalServerError}},
			wantStatusCode: http.StatusInternalServerError,
		},
		{
			name:           "does_not_retry_on_4xx",
			retryMax:       3,
			failures:       []fakeHTTPResponse{{statusCode: http.StatusBadRequest}},
			wantStatusCode: http.StatusBadRequest,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			wantBody := "PASS"

			called := atomic.Int32{}
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				call := int(called.Add(1))

				if call <= len(tt.failures) {
					f := tt.failures[call-1]

					for k, v := range f.headers {
						w.Header().Set(k, v)
					}

					w.WriteHeader(f.statusCode)

					_, _ = w.Write(f.body)

					return
				}

				w.WriteHeader(http.StatusOK)
				_, _ = w.Write([]byte(wantBody))
			}))
			defer ts.Close()

			tr, err := newRetryTransport(ts.Client().Transport, RetryOptions{Max: tt.retryMax, WaitMin: time.Millisecond, WaitMax: time.Millisecond, Jitter: time.Microsecond})
			if err != nil || tr == nil {
				t.Fatalf("failed to create retry transport: %v", err)
			}

			client := &http.Client{Transport: tr}

			var ctx context.Context
			var cancel context.CancelFunc
			if tt.cancel {
				ctx, cancel = context.WithCancel(t.Context())
				defer cancel()
			} else {
				ctx = t.Context()
			}

			req, err := http.NewRequestWithContext(ctx, "GET", ts.URL, nil)
			if err != nil {
				t.Fatalf("failed to create request: %v", err)
			}

			if tt.cancel {
				cancel()
			}

			res, err := client.Do(req)
			if err != nil {
				if tt.wantErr == nil {
					t.Fatalf("expected no error, got %v", err)
				}

				if !regexp.MustCompile(regexp.QuoteMeta(*tt.wantErr)).MatchString(err.Error()) {
					t.Fatalf("expected error %q, got %q", *tt.wantErr, err.Error())
				}

				return
			}
			defer res.Body.Close()

			if tt.wantErr != nil {
				t.Fatalf("expected error %q, got nil", *tt.wantErr)
				return
			}

			if res.StatusCode != tt.wantStatusCode {
				t.Fatalf("expected status code %d, got %d", tt.wantStatusCode, res.StatusCode)
			}
		})
	}
}

func Test_isRateLimitStatusCodes(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name string
		code int
		want bool
	}{
		{
			name: "200",
			code: 200,
			want: false,
		},
		{
			name: "403",
			code: 403,
			want: true,
		},
		{
			name: "404",
			code: 404,
			want: false,
		},
		{
			name: "429",
			code: 429,
			want: true,
		},
		{
			name: "500",
			code: 500,
			want: false,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := isRateLimitStatusCode(tt.code); got != tt.want {
				t.Errorf("isRateLimitStatusCode(%d) = %v, want %v", tt.code, got, tt.want)
			}
		})
	}
}

func Test_newBackoff(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name    string
		opts    RetryOptions
		wantErr *string
	}{
		{
			name: "success",
			opts: RetryOptions{Max: 3, WaitMin: 100 * time.Millisecond, WaitMax: 10 * time.Second, Jitter: 100 * time.Millisecond},
		},
		{
			name:    "errors_with_invalid_config",
			opts:    RetryOptions{},
			wantErr: new("invalid backoff configuration"),
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := newBackoff(tt.opts)
			if err != nil {
				if tt.wantErr == nil {
					t.Fatalf("expected no error, got %v", err)
				}

				if !regexp.MustCompile(regexp.QuoteMeta(*tt.wantErr)).MatchString(err.Error()) {
					t.Fatalf("expected error %q, got %q", *tt.wantErr, err.Error())
				}

				return
			}

			if tt.wantErr != nil {
				t.Fatalf("expected error %q, got nil", *tt.wantErr)
			}

			if got == nil {
				t.Fatalf("expected non-nil backoff, got nil")
			}
		})
	}
}

func Test_isRetryableStatusCode(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name string
		code int
		want bool
	}{
		{
			name: "200",
			code: 200,
			want: false,
		},
		{
			name: "403",
			code: 403,
			want: false,
		},
		{
			name: "404",
			code: 404,
			want: false,
		},
		{
			name: "429",
			code: 429,
			want: false,
		},
		{
			name: "500",
			code: 500,
			want: true,
		},
		{
			name: "501",
			code: 501,
			want: false,
		},
		{
			name: "502",
			code: 502,
			want: true,
		},
		{
			name: "503",
			code: 503,
			want: true,
		},
		{
			name: "504",
			code: 504,
			want: true,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := isRetryableStatusCode(tt.code); got != tt.want {
				t.Errorf("isRetryableStatusCode(%d) = %v, want %v", tt.code, got, tt.want)
			}
		})
	}
}

func Test_isRetryableNetworkError(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name string
		err  error
		want bool
	}{
		{
			name: "nil",
			err:  nil,
			want: false,
		},
		{
			name: "context_canceled",
			err:  context.Canceled,
			want: false,
		},
		{
			name: "context_deadline_exceeded",
			err:  context.DeadlineExceeded,
			want: false,
		},
		{
			name: "new_error",
			err:  fmt.Errorf("test"),
			want: true,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := isRetryableNetworkError(tt.err); got != tt.want {
				t.Errorf("isRetryableNetworkError(%v) = %v, want %v", tt.err, got, tt.want)
			}
		})
	}
}
