package httpx

import (
	"io"
	"net/http"
	"strings"
	"testing"
)

func Test_isIdempotent(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name   string
		method string
		header map[string]string
		want   bool
	}{
		{
			name:   "get_is_idempotent",
			method: "GET",
			want:   true,
		},
		{
			name:   "head_is_idempotent",
			method: "HEAD",
			want:   true,
		},
		{
			name:   "options_is_idempotent",
			method: "OPTIONS",
			want:   true,
		},
		{
			name:   "put_is_idempotent",
			method: "PUT",
			want:   true,
		},
		{
			name:   "delete_is_idempotent",
			method: "DELETE",
			want:   true,
		},
		{
			name:   "trace_is_idempotent",
			method: "TRACE",
			want:   true,
		},
		{
			name:   "post_with_idempotency_key_is_idempotent",
			method: "POST",
			header: map[string]string{"Idempotency-Key": "abc123"},
			want:   true,
		},
		{
			name:   "post_without_idempotency_key_is_not_idempotent",
			method: "POST",
			want:   false,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			req := mustCreateRequest(t, nil, tt.method, "https://example.com/", nil)

			for k, v := range tt.header {
				req.Header.Set(k, v)
			}

			got := isIdempotent(req)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_canRewind(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name string
		req  *http.Request
		want bool
	}{
		{
			name: "nil_body_can_rewind",
			req:  &http.Request{Body: nil},
			want: true,
		},
		{
			name: "no_body_can_rewind",
			req:  &http.Request{Body: http.NoBody},
			want: true,
		},
		{
			name: "body_and_get_body_can_rewind",
			req:  &http.Request{Body: io.NopCloser(strings.NewReader("test")), GetBody: func() (io.ReadCloser, error) { return nil, nil }},
			want: true,
		},
		{
			name: "body_without_get_body_cannot_rewind",
			req:  &http.Request{Body: io.NopCloser(strings.NewReader("test")), GetBody: nil},
			want: false,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := canRewind(tt.req)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
