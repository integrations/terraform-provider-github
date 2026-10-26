package httpx

import (
	"net"
	"net/http"
	"time"
)

// DefaultHTTPTransport returns a [http.Transport] that is a clone of the provided [http.RoundTripper] if it is already a [http.Transport], or a new [http.Transport] with default settings if the provided [http.RoundTripper] is nil or not a [http.Transport]. This ensures that the returned transport has sane defaults for timeouts and connection pooling.
func DefaultHTTPTransport(in http.RoundTripper) *http.Transport {
	if in == nil {
		in = http.DefaultTransport
	}

	if tr, ok := in.(*http.Transport); ok {
		return tr.Clone()
	}

	dialer := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}

	return &http.Transport{
		Proxy:                 http.ProxyFromEnvironment,
		DialContext:           dialer.DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
}
