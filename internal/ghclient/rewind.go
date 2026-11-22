package ghclient

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// newRewindableTransport returns a [rewindableTransport] that wraps the given [http.RoundTripper] and sets [request.GetBody] if it's unset and can either be rewound or fully buffered within the memory limit provided.
func newRewindableTransport(inner http.RoundTripper, maxBufferBytes int64) (http.RoundTripper, error) {
	if maxBufferBytes < 0 {
		return nil, fmt.Errorf("max buffer bytes must be non-negative")
	}

	return &rewindableTransport{
		inner:          inner,
		maxBufferBytes: maxBufferBytes,
	}, nil
}

// rewindableTransport is a [http.RoundTripper] that wraps another http.RoundTripper and sets [request.GetBody] if it's unset and can either be rewound or fully buffered within the memory limit provided.
type rewindableTransport struct {
	inner          http.RoundTripper
	maxBufferBytes int64
}

// RoundTrip implements the [http.RoundTripper] interface for the [rewindableTransport]. It sets [request.GetBody] if it's unset and can either be rewound or fully buffered within the memory limit provided.
func (t *rewindableTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if err := ensureRequestGetBody(req, t.maxBufferBytes); err != nil {
		return nil, err
	}

	return t.inner.RoundTrip(req)
}

// ensureRequestGetBody sets [request.GetBody] if it's unset and can either be rewound or fully buffered within the memory limit provided.
func ensureRequestGetBody(req *http.Request, maxBufferBytes int64) error {
	if req == nil || req.Body == nil || req.Body == http.NoBody || req.GetBody != nil {
		return nil
	}

	switch body := req.Body.(type) {
	case io.ReadSeeker:
		startOffset, err := body.Seek(0, io.SeekCurrent)
		if err != nil {
			return err
		}

		if req.ContentLength <= 0 {
			if endOffset, err := body.Seek(0, io.SeekEnd); err == nil {
				req.ContentLength = endOffset - startOffset
				_, _ = body.Seek(startOffset, io.SeekStart)
			}
		}

		req.Body = io.NopCloser(body)
		req.GetBody = func() (io.ReadCloser, error) {
			if _, err := body.Seek(startOffset, io.SeekStart); err != nil {
				return nil, err
			}
			return io.NopCloser(body), nil
		}
	default:
		var buf bytes.Buffer
		var err error

		if maxBufferBytes == 0 {
			_, err = io.Copy(&buf, req.Body)
		} else {
			_, err = io.CopyN(&buf, req.Body, maxBufferBytes)
		}

		if err != nil && !errors.Is(err, io.EOF) {
			return err
		}

		if maxBufferBytes == 0 || errors.Is(err, io.EOF) {
			_ = req.Body.Close()

			req.ContentLength = int64(buf.Len())
			req.Body = io.NopCloser(&buf)

			by := buf.Bytes()
			req.GetBody = func() (io.ReadCloser, error) {
				r := bytes.NewReader(by)
				return io.NopCloser(r), nil
			}
		} else {
			req.Body = struct {
				io.Reader
				io.Closer
			}{
				Reader: io.MultiReader(&buf, req.Body),
				Closer: req.Body,
			}
		}
	}

	return nil
}
