package ghclient

import (
	"context"
	"fmt"
	"regexp"
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
