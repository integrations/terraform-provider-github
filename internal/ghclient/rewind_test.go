package ghclient

import (
	"io"
	"net/http"
	"regexp"
	"strings"
	"testing"
)

func Test_newRewindableTransport(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name      string
		maxBuffer int64
		wantErr   *string
	}{
		{
			name:      "no_buffer_limit",
			maxBuffer: 0,
		},
		{
			name:      "buffer_limit",
			maxBuffer: 1024,
		},
		{
			name:      "invalid_buffer_limit",
			maxBuffer: -1,
			wantErr:   new("max buffer bytes must be non-negative"),
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			inner := &testRoundTripper{}

			got, err := newRewindableTransport(inner, tt.maxBuffer)
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

			rw, ok := got.(*rewindableTransport)
			if !ok {
				t.Fatal("expected rewindableTransport")
			}

			if rw.inner != inner {
				t.Error("expected inner transport to be set")
			}

			if rw.maxBufferBytes != tt.maxBuffer {
				t.Errorf("expected maxBufferBytes to be %v, got %v", tt.maxBuffer, rw.maxBufferBytes)
			}
		})
	}
}

func Test_rewindableTransport_RoundTrip(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name    string
		req     *http.Request
		wantErr *string
	}{
		{
			name:    "errors_if_body_lookup_errors",
			req:     &http.Request{Body: &errorReadSeeker{}},
			wantErr: new("error seeking"),
		},
		{
			name: "calls_inner_transport_when_no_error",
			req:  &http.Request{},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			inner := &testRoundTripper{resp: &http.Response{StatusCode: http.StatusOK}}
			tr := &rewindableTransport{inner: inner}

			got, err := tr.RoundTrip(tt.req)
			if err != nil {
				if tt.wantErr == nil {
					t.Fatalf("expected no error, got %v", err)
				}

				if !regexp.MustCompile(regexp.QuoteMeta(*tt.wantErr)).MatchString(err.Error()) {
					t.Fatalf("expected error %q, got %q", *tt.wantErr, err.Error())
				}

				if inner.called.Load() != 0 {
					t.Fatalf("expected inner transport to not have been called")
				}

				return
			}

			if tt.wantErr != nil {
				t.Fatalf("expected error %q, got nil", *tt.wantErr)
				return
			}

			if inner.called.Load() != 1 {
				t.Errorf("expected inner transport to have been called")
			}

			if got == nil {
				t.Fatal("expected non-nil response")
			}
		})
	}
}

func Test_ensureRequestGetBody(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name           string
		req            *http.Request
		maxBufferBytes int64
		wantGetBody    bool
		wantErr        *string
	}{
		{
			name: "nil_request",
		},
		{
			name: "nil_body",
			req:  &http.Request{},
		},
		{
			name: "no_body",
			req:  &http.Request{Body: http.NoBody},
		},
		{
			name:        "body_and_get_body",
			req:         &http.Request{Body: io.NopCloser(strings.NewReader("TEST")), GetBody: func() (io.ReadCloser, error) { return io.NopCloser(strings.NewReader("TEST")), nil }},
			wantGetBody: true,
		},
		{
			name:        "body_read_seeker",
			req:         &http.Request{Body: io.NopCloser(strings.NewReader("TEST"))},
			wantGetBody: true,
		},
		{
			name:        "body_reader",
			req:         &http.Request{Body: io.NopCloser(&nonSeekableReader{r: strings.NewReader("TEST")})},
			wantGetBody: true,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := ensureRequestGetBody(tt.req, tt.maxBufferBytes)
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
				return
			}

			if tt.req != nil && (tt.req.GetBody != nil) != tt.wantGetBody {
				t.Errorf("expected req.GetBody to be nil as %v, got %v", tt.wantGetBody, tt.req.GetBody != nil)
			}
		})
	}
}
