package httpx

import (
	"net/http"
	"testing"
)

func TestDefaultHTTPTransport(t *testing.T) {
	t.Parallel()

	t.Run("with_nil_input_clones_default", func(t *testing.T) {
		t.Parallel()

		got := DefaultHTTPTransport(nil)

		if got == nil {
			t.Fatal("transport: got nil, want non-nil")
		}

		if got == http.DefaultTransport {
			t.Fatal("transport: got http.DefaultTransport, want clone")
		}
	})

	t.Run("with_non_nil_http_transport_returns_clone", func(t *testing.T) {
		t.Parallel()

		in := &http.Transport{
			MaxIdleConns: 999,
		}

		got := DefaultHTTPTransport(in)

		if got == nil {
			t.Fatal("transport: got nil, want non-nil")
		}

		if got == in {
			t.Fatal("transport: got input transport, want clone")
		}

		if got.MaxIdleConns != in.MaxIdleConns {
			t.Fatalf("MaxIdleConns: got %d, want %d", got.MaxIdleConns, in.MaxIdleConns)
		}
	})

	t.Run("with_non_http_transport_returns_defaults", func(t *testing.T) {
		t.Parallel()

		in := &dummyRoundTripper{}

		got := DefaultHTTPTransport(in)

		if got == nil {
			t.Fatal("transport: got nil, want non-nil")
		}

		if got.Proxy == nil {
			t.Error("Proxy: got nil, want non-nil")
		}

		if got.DialContext == nil {
			t.Error("DialContext: got nil, want non-nil")
		}

		if !got.ForceAttemptHTTP2 {
			t.Error("ForceAttemptHTTP2: got false, want true")
		}

		if got.MaxIdleConns != 100 {
			t.Errorf("MaxIdleConns: got %d, want 100", got.MaxIdleConns)
		}
	})
}
