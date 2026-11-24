package httpx

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sync"
)

// rewindOptions holds the configuration for the rewindableTransport.
type rewindOptions struct {
	bufferBytes int64
}

// RewindOption is a function to configure a new rewindableTransport.
type RewindOption func(*rewindOptions)

// WithRewindBufferBytes returns a [RewindOption] that sets the maximum number of bytes to buffer in memory for rewinding requests. If the request body exceeds this limit, it will not be buffered and cannot be rewound. A value of 0 will attempt to buffer all bytes, a negative value is treated as 0.
func WithRewindBufferBytes(by int64) RewindOption {
	return func(o *rewindOptions) {
		if by < 0 {
			by = 0
		}

		o.bufferBytes = by
	}
}

// NewRewindableTransport returns a [http.RoundTripper] that wraps the given [http.RoundTripper] and sets [request.GetBody] if it's unset and can either be rewound or fully buffered within the memory limit provided.
func NewRewindableTransport(inner http.RoundTripper, opts ...RewindOption) http.RoundTripper {
	o := rewindOptions{}

	for _, opt := range opts {
		opt(&o)
	}

	return &rewindableTransport{
		inner:       inner,
		bufferBytes: o.bufferBytes,
	}
}

// rewindableTransport is a [http.RoundTripper] that wraps another http.RoundTripper and sets [request.GetBody] if it's unset and can either be rewound or fully buffered within the memory limit provided.
type rewindableTransport struct {
	inner       http.RoundTripper
	bufferBytes int64
}

// RoundTrip implements the [http.RoundTripper] interface for the [rewindableTransport]. It sets [request.GetBody] if it's unset and can either be rewound or fully buffered within the memory limit provided.
func (t *rewindableTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Shallow copy so we never mutate the caller's Request
	reqCopy := new(*req)

	closer, err := ensureRequestGetBody(reqCopy, t.bufferBytes)
	if err != nil {
		if req.Body != nil {
			_ = req.Body.Close()
		}

		return nil, err
	}

	resp, err := t.inner.RoundTrip(reqCopy)
	if closer == nil {
		return resp, err
	}

	if err != nil || resp == nil || resp.Body == nil {
		_ = closer.Close()
		return resp, err
	}

	// The transport may still be reading the request body while the response body is being read (e.g. HTTP/2), so defer closing the original body until the response body is closed.
	resp.Body = &closeAfterReadCloser{ReadCloser: resp.Body, after: closer}

	return resp, nil
}

// ensureRequestGetBody sets [request.GetBody] if it's unset and can either be rewound or fully buffered within the memory limit provided. If the returned [io.Closer] is not nil, the request body has been replaced with a non-closing wrapper and the caller is responsible for closing the original body once the request is complete.
func ensureRequestGetBody(req *http.Request, bufferBytes int64) (io.Closer, error) {
	if req == nil || req.Body == nil || req.Body == http.NoBody || req.GetBody != nil {
		return nil, nil
	}

	switch body := req.Body.(type) {
	case io.ReadSeekCloser:
		startOffset, err := body.Seek(0, io.SeekCurrent)
		if err != nil {
			return nil, fmt.Errorf("failed to seek request body: %w", err)
		}

		if req.ContentLength <= 0 {
			if endOffset, err := body.Seek(0, io.SeekEnd); err == nil {
				if _, err := body.Seek(startOffset, io.SeekStart); err != nil {
					return nil, fmt.Errorf("failed to seek request body: %w", err)
				}

				req.ContentLength = endOffset - startOffset
			}
		}

		req.Body = io.NopCloser(body)
		req.GetBody = func() (io.ReadCloser, error) {
			if _, err := body.Seek(startOffset, io.SeekStart); err != nil {
				return nil, err
			}
			return io.NopCloser(body), nil
		}

		return body, nil
	default:
		var data []byte

		if req.ContentLength > 0 {
			// The size is known up front: skip buffering if it can't fit, otherwise allocate exactly what is needed.
			if bufferBytes > 0 && req.ContentLength > bufferBytes {
				return nil, nil
			}

			data = make([]byte, req.ContentLength)

			if _, err := io.ReadFull(req.Body, data); err != nil {
				if errors.Is(err, io.ErrUnexpectedEOF) || errors.Is(err, io.EOF) {
					return nil, fmt.Errorf("request body shorter than content length %d: %w", req.ContentLength, io.ErrUnexpectedEOF)
				}

				return nil, err
			}
		} else {
			var r io.Reader = req.Body
			if bufferBytes > 0 {
				// Read one byte past the limit to detect overflow without losing data.
				r = io.LimitReader(r, bufferBytes+1)
			}

			var err error
			if data, err = io.ReadAll(r); err != nil {
				return nil, err
			}

			if bufferBytes > 0 && int64(len(data)) > bufferBytes {
				// Too large to rewind; replay the bytes already consumed followed by the remainder.
				req.Body = struct {
					io.Reader
					io.Closer
				}{
					Reader: io.MultiReader(bytes.NewReader(data), req.Body),
					Closer: req.Body,
				}

				return nil, nil
			}
		}

		_ = req.Body.Close()

		if len(data) == 0 {
			req.ContentLength = 0
			req.Body = http.NoBody
			req.GetBody = func() (io.ReadCloser, error) {
				return http.NoBody, nil
			}

			return nil, nil
		}

		req.ContentLength = int64(len(data))
		req.Body = io.NopCloser(bytes.NewReader(data))
		req.GetBody = func() (io.ReadCloser, error) {
			return io.NopCloser(bytes.NewReader(data)), nil
		}
	}

	return nil, nil
}

// closeAfterReadCloser is an [io.ReadCloser] that closes an additional [io.Closer] after the wrapped [io.ReadCloser] has been closed.
type closeAfterReadCloser struct {
	io.ReadCloser
	after io.Closer
	once  sync.Once
}

// Close closes the wrapped [io.ReadCloser] and then the additional [io.Closer], returning the first error encountered.
func (c *closeAfterReadCloser) Close() error {
	err := c.ReadCloser.Close()

	c.once.Do(func() {
		if afterErr := c.after.Close(); err == nil {
			err = afterErr
		}
	})

	return err
}
