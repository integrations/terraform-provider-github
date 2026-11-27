package httpx

import (
	"net/http"
	"testing"
	"testing/synctest"
	"time"
)

func Test_retryableResponse(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name      string
		method    string
		resp      *http.Response
		want      bool
		wantAfter *time.Duration
	}{
		{
			name:   "too_many_requests_can_retry",
			method: http.MethodGet,
			resp:   &http.Response{StatusCode: http.StatusTooManyRequests, Header: http.Header{}},
			want:   true,
		},
		{
			name:   "too_many_requests_can_retry_for_post",
			method: http.MethodPost,
			resp:   &http.Response{StatusCode: http.StatusTooManyRequests, Header: http.Header{}},
			want:   true,
		},
		{
			name:   "service_unavailable_can_retry",
			method: http.MethodGet,
			resp:   &http.Response{StatusCode: http.StatusServiceUnavailable, Header: http.Header{}},
			want:   true,
		},
		{
			name:   "service_unavailable_can_retry_for_post",
			method: http.MethodPost,
			resp:   &http.Response{StatusCode: http.StatusServiceUnavailable, Header: http.Header{}},
			want:   true,
		},
		{
			name:   "internal_server_error_can_retry",
			method: http.MethodGet,
			resp:   &http.Response{StatusCode: http.StatusInternalServerError, Header: http.Header{}},
			want:   true,
		},
		{
			name:   "internal_server_error_can_not_retry_for_post",
			method: http.MethodPost,
			resp:   &http.Response{StatusCode: http.StatusInternalServerError, Header: http.Header{}},
			want:   false,
		},
		{
			name:   "bad_gateway_can_retry",
			method: http.MethodGet,
			resp:   &http.Response{StatusCode: http.StatusBadGateway, Header: http.Header{}},
			want:   true,
		},
		{
			name:   "bad_gateway_can_not_retry_for_post",
			method: http.MethodPost,
			resp:   &http.Response{StatusCode: http.StatusBadGateway, Header: http.Header{}},
			want:   false,
		},
		{
			name:   "gateway_timeout_can_retry",
			method: http.MethodGet,
			resp:   &http.Response{StatusCode: http.StatusGatewayTimeout, Header: http.Header{}},
			want:   true,
		},
		{
			name:   "gateway_timeout_can_not_retry_for_post",
			method: http.MethodPost,
			resp:   &http.Response{StatusCode: http.StatusGatewayTimeout, Header: http.Header{}},
			want:   false,
		},
		{
			name:      "forbidden_can_retry_with_retry_after",
			method:    http.MethodGet,
			resp:      &http.Response{StatusCode: http.StatusForbidden, Header: http.Header{"Retry-After": []string{"5"}}},
			want:      true,
			wantAfter: new(5 * time.Second),
		},
		{
			name:      "forbidden_can_retry_with_retry_after_for_post",
			method:    http.MethodPost,
			resp:      &http.Response{StatusCode: http.StatusForbidden, Header: http.Header{"Retry-After": []string{"5"}}},
			want:      true,
			wantAfter: new(5 * time.Second),
		},
		{
			name:   "forbidden_can_not_retry_without_retry_after",
			method: http.MethodGet,
			resp:   &http.Response{StatusCode: http.StatusForbidden, Header: http.Header{}},
			want:   false,
		},
		{
			name:      "with_retry_after_header",
			method:    http.MethodGet,
			resp:      &http.Response{StatusCode: http.StatusInternalServerError, Header: http.Header{"Retry-After": []string{"5"}}},
			want:      true,
			wantAfter: new(5 * time.Second),
		},
		{
			name:   "not_implemented_can_not_retry",
			method: http.MethodGet,
			resp:   &http.Response{StatusCode: http.StatusNotImplemented, Header: http.Header{}},
			want:   false,
		},
		{
			name:   "other_status_codes_can_not_retry",
			method: http.MethodGet,
			resp:   &http.Response{StatusCode: http.StatusNotFound, Header: http.Header{}},
			want:   false,
		},
		{
			name:      "too_many_requests_can_retry_with_retry_after",
			method:    http.MethodGet,
			resp:      &http.Response{StatusCode: http.StatusTooManyRequests, Header: http.Header{"Retry-After": []string{"5"}}},
			want:      true,
			wantAfter: new(5 * time.Second),
		},
		{
			name:      "forbidden_can_retry_with_zero_retry_after",
			method:    http.MethodGet,
			resp:      &http.Response{StatusCode: http.StatusForbidden, Header: http.Header{"Retry-After": []string{"0"}}},
			want:      true,
			wantAfter: new(time.Duration(0)),
		},
		{
			name:   "forbidden_can_not_retry_with_invalid_retry_after",
			method: http.MethodGet,
			resp:   &http.Response{StatusCode: http.StatusForbidden, Header: http.Header{"Retry-After": []string{"soon"}}},
			want:   false,
		},
		{
			name:   "service_unavailable_can_retry_with_invalid_retry_after",
			method: http.MethodGet,
			resp:   &http.Response{StatusCode: http.StatusServiceUnavailable, Header: http.Header{"Retry-After": []string{"soon"}}},
			want:   true,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			req := mustCreateRequest(t, nil, tt.method, "https://example.com", nil)

			got, gotAfter := retryableResponse(req, tt.resp)
			if got != tt.want {
				t.Fatalf("retryable: got %v, want %v", got, tt.want)
			}

			if (gotAfter == nil && tt.wantAfter != nil) || (gotAfter != nil && tt.wantAfter == nil) || (tt.wantAfter != nil && *gotAfter != *tt.wantAfter) {
				t.Fatalf("after: got %v, want %v", gotAfter, tt.wantAfter)
			}
		})
	}
}

func TestParseRetryAfter(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name       string
		raw        string
		timeOffset *time.Duration
		want       time.Duration
		wantOk     bool
	}{
		{
			name:   "seconds",
			raw:    "5",
			want:   5 * time.Second,
			wantOk: true,
		},
		{
			name:       "http_date",
			timeOffset: new(10 * time.Minute),
			want:       10 * time.Minute,
			wantOk:     true,
		},
		{
			name:   "invalid",
			raw:    "invalid",
			want:   0,
			wantOk: false,
		},
		{
			name:   "zero",
			raw:    "0",
			want:   0,
			wantOk: true,
		},
		{
			name:   "seconds_with_whitespace",
			raw:    " 5 ",
			want:   5 * time.Second,
			wantOk: true,
		},
		{
			name:   "leading_zeros",
			raw:    "007",
			want:   7 * time.Second,
			wantOk: true,
		},
		{
			name:       "past_http_date",
			timeOffset: new(-time.Hour),
			want:       0,
			wantOk:     true,
		},
		{
			name:   "seconds_at_max",
			raw:    "9223372036",
			want:   time.Duration(9223372036) * time.Second,
			wantOk: true,
		},
		{
			name:   "seconds_over_max",
			raw:    "9223372037",
			want:   0,
			wantOk: false,
		},
		{
			name:   "seconds_over_uint64",
			raw:    "99999999999999999999",
			want:   0,
			wantOk: false,
		},
		{
			name:   "empty",
			raw:    "",
			want:   0,
			wantOk: false,
		},
		{
			name:   "negative",
			raw:    "-5",
			want:   0,
			wantOk: false,
		},
		{
			name:   "decimal",
			raw:    "1.5",
			want:   0,
			wantOk: false,
		},
		{
			name:   "plus_sign",
			raw:    "+5",
			want:   0,
			wantOk: false,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			synctest.Test(t, func(t *testing.T) {
				in := tt.raw
				if tt.timeOffset != nil {
					in = time.Now().Add(*tt.timeOffset).UTC().Format(http.TimeFormat)
				}

				got, gotOk := ParseRetryAfter(in)

				if gotOk != tt.wantOk {
					t.Fatalf("ok: got %v, want %v", gotOk, tt.wantOk)
				}

				if got != tt.want {
					t.Fatalf("duration: got %v, want %v", got, tt.want)
				}
			})
		})
	}
}
