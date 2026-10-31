package ghclient

import (
	"bytes"
	"io"
	"net/http"
	"regexp"
	"testing"
)

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
