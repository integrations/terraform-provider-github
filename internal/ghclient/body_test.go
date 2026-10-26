package ghclient

import (
	"bytes"
	"io"
	"net/http"
	"regexp"
	"testing"
)

func Test_rewindRequestBody(t *testing.T) {
	t.Parallel()

	body := []byte("TEST")

	for _, tt := range []struct {
		name     string
		req      *http.Request
		wantBody bool
		wantErr  *string
	}{
		{
			name:    "nil_request_errors",
			req:     nil,
			wantErr: new("request is nil"),
		},
		{
			name:    "request_with_no_body_succeeds",
			req:     &http.Request{},
			wantErr: nil,
		},
		{
			name: "request_with_get_body_succeeds",
			req: &http.Request{
				Body:    io.NopCloser(bytes.NewBuffer(body)),
				GetBody: func() (io.ReadCloser, error) { return io.NopCloser(bytes.NewBuffer(body)), nil },
			},
			wantBody: true,
		},
		{
			name: "request_without_get_body_errors",
			req: &http.Request{
				Body: io.NopCloser(bytes.NewBuffer(body)),
			},
			wantErr: new("cannot rewind request body: get body is nil"),
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := rewindRequestBody(tt.req)
			if err != nil {
				if tt.wantErr == nil {
					t.Fatalf("unexpected error: %s", err)
				}
				if !regexp.MustCompile(regexp.QuoteMeta(*tt.wantErr)).MatchString(err.Error()) {
					t.Fatalf("unexpected error: %s", err)
				}
				return
			}

			if tt.wantErr != nil {
				t.Fatalf("expected error: %s, got: %v", *tt.wantErr, err)
			}

			if tt.req == nil {
				t.Fatalf("request is nil after rewind")
			}

			if !tt.wantBody {
				return
			}

			if tt.req.Body == nil {
				t.Fatalf("request body is nil after rewind")
			}

			if tt.req.GetBody == nil {
				t.Fatalf("request GetBody is nil after rewind")
			}

			got, err := io.ReadAll(tt.req.Body)
			if err != nil {
				t.Fatalf("failed to read request body: %s", err)
			}
			if !bytes.Equal(got, body) {
				t.Fatalf("unexpected request body: got %s, want %s", got, body)
			}
		})
	}
}

func Test_drainResponseBody(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name    string
		resp    *http.Response
		wantErr *string
	}{
		{
			name:    "nil_response_errors",
			resp:    nil,
			wantErr: new("response is nil"),
		},
		{
			name:    "response_with_no_body_seccedes",
			resp:    &http.Response{},
			wantErr: nil,
		},
		{
			name:    "response_with_body_seccedes",
			resp:    &http.Response{Body: io.NopCloser(bytes.NewBufferString("TEST"))},
			wantErr: nil,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := drainResponseBody(tt.resp)
			if err != nil {
				if tt.wantErr == nil {
					t.Fatalf("unexpected error: %s", err)
				}
				if !regexp.MustCompile(regexp.QuoteMeta(*tt.wantErr)).MatchString(err.Error()) {
					t.Fatalf("unexpected error: %s", err)
				}
				return
			}

			if tt.wantErr != nil {
				t.Fatalf("expected error: %s, got: %v", *tt.wantErr, err)
			}
		})
	}
}
