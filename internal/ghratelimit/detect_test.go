package ghratelimit

import (
	"io"
	"net/http"
	"strings"
	"testing"
)

func Test_isPrimaryRateLimit(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name string
		resp *http.Response
		want bool
	}{
		{
			name: "primary_rate_limit_forbidden",
			resp: &http.Response{
				StatusCode: http.StatusForbidden,
				Header:     http.Header{headerRateRemaining: []string{"0"}},
			},
			want: true,
		},
		{
			name: "primary_rate_limit_too_many_requests",
			resp: &http.Response{
				StatusCode: http.StatusTooManyRequests,
				Header:     http.Header{headerRateRemaining: []string{"0"}},
			},
			want: true,
		},
		{
			name: "not_primary_rate_limit_by_code",
			resp: &http.Response{
				StatusCode: http.StatusNotFound,
			},
			want: false,
		},
		{
			name: "not_primary_rate_limit_forbidden_with_remaining",
			resp: &http.Response{
				StatusCode: http.StatusForbidden,
				Header:     http.Header{headerRateRemaining: []string{"5"}},
			},
			want: false,
		},
		{
			name: "not_primary_rate_limit_too_many_requests_with_remaining",
			resp: &http.Response{
				StatusCode: http.StatusTooManyRequests,
				Header:     http.Header{headerRateRemaining: []string{"5"}},
			},
			want: false,
		},
		{
			name: "not_primary_rate_limit_forbidden_with_remaining_not_set",
			resp: &http.Response{
				StatusCode: http.StatusForbidden,
			},
			want: false,
		},
		{
			name: "not_primary_rate_limit_too_many_requests_with_remaining_not_set",
			resp: &http.Response{
				StatusCode: http.StatusTooManyRequests,
			},
			want: false,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := isPrimaryRateLimit(tt.resp)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isSecondaryRateLimit(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name string
		resp *http.Response
		want bool
	}{
		{
			name: "secondary_rate_limit_forbidden_with_retry_after",
			resp: &http.Response{
				StatusCode: http.StatusForbidden,
				Header:     http.Header{headerRetryAfter: []string{"5"}},
			},
			want: true,
		},
		{
			name: "secondary_rate_limit_too_many_requests_with_retry_after",
			resp: &http.Response{
				StatusCode: http.StatusTooManyRequests,
				Header:     http.Header{headerRetryAfter: []string{"5"}},
			},
			want: true,
		},
		{
			name: "secondary_rate_limit_too_many_requests_with_rate_limit_remaining_not_zero",
			resp: &http.Response{
				StatusCode: http.StatusTooManyRequests,
				Header:     http.Header{headerRateRemaining: []string{"5"}},
			},
			want: true,
		},
		{
			name: "secondary_rate_limit_too_many_requests_with_rate_limit_remaining_not_set",
			resp: &http.Response{
				StatusCode: http.StatusTooManyRequests,
				Header:     http.Header{},
			},
			want: true,
		},
		{
			name: "secondary_rate_limit_too_many_requests_with_retry_after_and_rate_limit_remaining_zero",
			resp: &http.Response{
				StatusCode: http.StatusTooManyRequests,
				Header:     http.Header{headerRetryAfter: []string{"5"}, headerRateRemaining: []string{"0"}},
			},
			want: true,
		},
		{
			name: "secondary_rate_limit_forbidden_by_body",
			resp: &http.Response{
				StatusCode: http.StatusForbidden,
				Header:     http.Header{headerRateRemaining: []string{"0"}},
				Body:       io.NopCloser(strings.NewReader("You have hit the secondary rate limit.")),
			},
			want: true,
		},
		{
			name: "secondary_rate_limit_too_many_requests_by_body",
			resp: &http.Response{
				StatusCode: http.StatusTooManyRequests,
				Header:     http.Header{headerRateRemaining: []string{"0"}},
				Body:       io.NopCloser(strings.NewReader("You have hit the secondary rate limit.")),
			},
			want: true,
		},
		{
			name: "not_secondary_rate_limit",
			resp: &http.Response{
				StatusCode: http.StatusNotFound,
				Body:       io.NopCloser(strings.NewReader(`{"message":"foo bar"}`)),
			},
			want: false,
		},
		{
			name: "not_secondary_rate_limit_forbidden_with_missing_headers",
			resp: &http.Response{
				StatusCode: http.StatusForbidden,
			},
			want: false,
		},
		{
			name: "not_secondary_rate_limit_forbidden_with_rate_limit_remaining_not_zero",
			resp: &http.Response{
				StatusCode: http.StatusForbidden,
				Header:     http.Header{headerRateRemaining: []string{"5"}},
			},
			want: false,
		},
		{
			name: "not_secondary_rate_limit_too_many_requests_with_rate_limit_remaining_zero",
			resp: &http.Response{
				StatusCode: http.StatusTooManyRequests,
				Header:     http.Header{headerRateRemaining: []string{"0"}},
			},
			want: false,
		},
		{
			name: "not_secondary_rate_limit_too_many_requests_with_rate_limit_remaining_zero_and_primary_body",
			resp: &http.Response{
				StatusCode: http.StatusTooManyRequests,
				Header:     http.Header{headerRateRemaining: []string{"0"}},
				Body:       io.NopCloser(strings.NewReader(`{"message":"API rate limit exceeded for user ID 1."}`)),
			},
			want: false,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := isSecondaryRateLimit(tt.resp)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hasSecondaryRateLimitBody(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name string
		body *string
		want bool
	}{
		{
			name: "secondary_rate_limit",
			body: new("You have hit the secondary rate limit."),
			want: true,
		},
		{
			name: "secondary_rate_limit_by_abuse_detection",
			body: new("You have triggered the abuse detection mechanism."),
			want: true,
		},
		{
			name: "secondary_rate_limit_long_body",
			body: new("You have hit the secondary rate limit." + strings.Repeat("a", 2000)),
			want: true,
		},
		{
			name: "not_secondary_rate_limit_long_body_with_text_after_window",
			body: new(strings.Repeat("a", 2000) + "You have hit the secondary rate limit."),
			want: false,
		},
		{
			name: "not_secondary_rate_limit_body",
			body: new("some other error message."),
			want: false,
		},
		{
			name: "not_secondary_rate_limit_nil_body",
			body: nil,
			want: false,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var body io.ReadCloser
			if tt.body != nil {
				body = io.NopCloser(strings.NewReader(*tt.body))
			}

			resp := &http.Response{Body: body}

			got := hasSecondaryRateLimitBody(resp)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}

			if tt.body == nil {
				return
			}

			gotBody, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("unexpected error reading body: %v", err)
			}

			if string(gotBody) != *tt.body {
				t.Errorf("body length: got %d, want %d", len(gotBody), len(*tt.body))
			}
		})
	}

	t.Run("handles_body_read_error", func(t *testing.T) {
		t.Parallel()

		resp := &http.Response{
			StatusCode: http.StatusForbidden,
			Body:       &errorReadSeeker{},
		}

		got := hasSecondaryRateLimitBody(resp)
		if got {
			t.Error("got true, want false")
		}
	})

	t.Run("restores_consumed_bytes_on_body_read_error", func(t *testing.T) {
		t.Parallel()

		partial := "You have hit the secondary rate limit."
		resp := &http.Response{
			StatusCode: http.StatusForbidden,
			Body:       io.NopCloser(io.MultiReader(strings.NewReader(partial), &errorReadSeeker{})),
		}

		if hasSecondaryRateLimitBody(resp) {
			t.Error("got true, want false")
		}

		gotBody, err := io.ReadAll(resp.Body)
		if err == nil {
			t.Fatal("got nil, want read error after restored bytes")
		}
		if string(gotBody) != partial {
			t.Errorf("body: got %q, want %q", gotBody, partial)
		}
	})
}
