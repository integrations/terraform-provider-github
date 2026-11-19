package httpx

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"testing"
)

func Test_rewindRequestBody(t *testing.T) {
	t.Parallel()

	body := []byte("TEST")
	errGetBody := errors.New("get body failed")

	for _, tt := range []struct {
		name     string
		req      *http.Request
		wantBody bool
		wantErr  error
	}{
		{
			name:    "nil_request_errors",
			req:     nil,
			wantErr: errRequestNil,
		},
		{
			name: "request_with_no_body_succeeds",
			req:  &http.Request{},
		},
		{
			name: "request_with_get_body_succeeds",
			req: &http.Request{
				Body:    io.NopCloser(bytes.NewReader(nil)),
				GetBody: func() (io.ReadCloser, error) { return io.NopCloser(bytes.NewBuffer(body)), nil },
			},
			wantBody: true,
		},
		{
			name: "request_without_get_body_errors",
			req: &http.Request{
				Body: io.NopCloser(bytes.NewBuffer(body)),
			},
			wantErr: errRequestGetBodyNil,
		},
		{
			name: "request_with_no_body_sentinel_succeeds",
			req:  &http.Request{Body: http.NoBody},
		},
		{
			name: "request_with_failing_get_body_errors",
			req: &http.Request{
				Body:    io.NopCloser(bytes.NewBuffer(body)),
				GetBody: func() (io.ReadCloser, error) { return nil, errGetBody },
			},
			wantErr: errGetBody,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := rewindRequestBody(tt.req)
			if err != nil {
				if tt.wantErr == nil {
					t.Fatalf("unexpected error: %s", err)
				}
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("got %v, want %v", err, tt.wantErr)
				}
				return
			}

			if tt.wantErr != nil {
				t.Fatalf("got nil, want %v", tt.wantErr)
			}

			if tt.req == nil {
				t.Fatal("request: got nil, want non-nil")
			}

			if !tt.wantBody {
				return
			}

			if tt.req.Body == nil {
				t.Fatal("request body: got nil, want non-nil")
			}

			if tt.req.GetBody == nil {
				t.Fatal("GetBody: got nil, want non-nil")
			}

			got, err := io.ReadAll(tt.req.Body)
			if err != nil {
				t.Fatalf("failed to read request body: %s", err)
			}
			if !bytes.Equal(got, body) {
				t.Fatalf("request body: got %s, want %s", got, body)
			}
		})
	}
}
