package ghclient

import "testing"

func Test_newRateLimitTransport(t *testing.T) {
	t.Parallel()

	inner := &testRoundTripper{}
	tr := newRateLimitTransport(inner)

	if tr == nil {
		t.Fatal("expected non-nil transport")
	}
}
