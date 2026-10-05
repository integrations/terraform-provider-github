package ghclient

import (
	"regexp"
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
