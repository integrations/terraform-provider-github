package httpx

import (
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestWithRewindBufferBytes(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name            string
		bufferBytes     int64
		wantBufferBytes int64
	}{
		{
			name:            "no_buffer_limit",
			bufferBytes:     0,
			wantBufferBytes: 0,
		},
		{
			name:            "buffer_limit",
			bufferBytes:     1024,
			wantBufferBytes: 1024,
		},
		{
			name:            "negative_buffer_limit",
			bufferBytes:     -1,
			wantBufferBytes: 0,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			opts := &rewindOptions{}
			WithRewindBufferBytes(tt.bufferBytes)(opts)

			if opts.bufferBytes != tt.wantBufferBytes {
				t.Errorf("got %v, want %v", opts.bufferBytes, tt.wantBufferBytes)
			}
		})
	}
}

func TestNewRewindableTransport(t *testing.T) {
	t.Parallel()

	t.Run("with_default_buffer_bytes", func(t *testing.T) {
		t.Parallel()

		inner := &dummyRoundTripper{}
		got := NewRewindableTransport(inner)

		if got == nil {
			t.Fatal("transport: got nil, want non-nil")
		}

		rw, ok := got.(*rewindableTransport)
		if !ok {
			t.Fatal("transport: got non-rewindableTransport, want rewindableTransport")
		}

		if rw.inner != inner {
			t.Error("inner: got other transport, want given transport")
		}

		if rw.bufferBytes != 0 {
			t.Errorf("bufferBytes: got %v, want 0", rw.bufferBytes)
		}
	})

	t.Run("with_custom_buffer_bytes", func(t *testing.T) {
		t.Parallel()

		bufferBytes := int64(1024)

		inner := &dummyRoundTripper{}
		got := NewRewindableTransport(inner, WithRewindBufferBytes(bufferBytes))

		if got == nil {
			t.Fatal("transport: got nil, want non-nil")
		}

		rw, ok := got.(*rewindableTransport)
		if !ok {
			t.Fatal("transport: got non-rewindableTransport, want rewindableTransport")
		}

		if rw.inner != inner {
			t.Error("inner: got other transport, want given transport")
		}

		if rw.bufferBytes != bufferBytes {
			t.Errorf("bufferBytes: got %v, want %v", rw.bufferBytes, bufferBytes)
		}
	})
}

func Test_rewindableTransport(t *testing.T) {
	t.Parallel()

	t.Run("RoundTrip", func(t *testing.T) {
		t.Parallel()

		t.Run("success_if_get_body_ok", func(t *testing.T) {
			t.Parallel()

			inner := &testRoundTripper{resp: &http.Response{StatusCode: http.StatusOK}}
			tr := &rewindableTransport{
				inner:       inner,
				bufferBytes: 1024,
			}

			req := mustCreateRequest(t, nil, http.MethodGet, "https://example.com", nil)
			resp, err := tr.RoundTrip(req)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if resp == nil {
				t.Fatal("response: got nil, want non-nil")
			}
		})

		t.Run("sets_get_body_on_forwarded_request_only", func(t *testing.T) {
			t.Parallel()

			var got *http.Request
			inner := &testRoundTripper{f: func(req *http.Request) (*http.Response, error) {
				got = req
				return &http.Response{StatusCode: http.StatusOK}, nil
			}}
			tr := &rewindableTransport{inner: inner}

			req := mustCreateRequest(t, nil, http.MethodPost, "https://example.com", io.NopCloser(&nonSeekableReader{r: strings.NewReader("TEST")}))

			if _, err := tr.RoundTrip(req); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if got.GetBody == nil {
				t.Error("forwarded GetBody: got nil, want non-nil")
			}
			if req.GetBody != nil {
				t.Error("caller GetBody: got non-nil, want nil")
			}
		})

		t.Run("errors_if_get_body_fails", func(t *testing.T) {
			t.Parallel()

			inner := &testRoundTripper{resp: &http.Response{StatusCode: http.StatusOK}}
			tr := &rewindableTransport{
				inner:       inner,
				bufferBytes: 1024,
			}

			req := mustCreateRequest(t, nil, http.MethodGet, "https://example.com", nil)
			rec := &closeRecorder{}
			req.Body = &errorReadSeeker{Closer: rec}

			resp, err := tr.RoundTrip(req)
			if err == nil {
				t.Fatal("got nil, want error")
			}
			if want := "failed to seek request body"; !strings.Contains(err.Error(), want) {
				t.Errorf("got %q, want it to contain %q", err.Error(), want)
			}
			if resp != nil {
				t.Fatal("response: got non-nil, want nil")
			}
			if !rec.closed.Load() {
				t.Error("request body closed: got false, want true")
			}
		})

		t.Run("handles_inner_error", func(t *testing.T) {
			t.Parallel()

			inner := &testRoundTripper{err: errors.New("boom")}
			tr := &rewindableTransport{
				inner:       inner,
				bufferBytes: 1024,
			}

			req := mustCreateRequest(t, nil, http.MethodGet, "https://example.com", nil)
			resp, err := tr.RoundTrip(req)
			if err == nil {
				t.Fatal("got nil, want error")
			}
			if resp != nil {
				t.Fatal("response: got non-nil, want nil")
			}
		})

		t.Run("closes_read_seeker_body_after_response_body_closed", func(t *testing.T) {
			t.Parallel()

			body := &closeRecordingReadSeeker{Reader: strings.NewReader("TEST")}
			respBody := &closeRecorder{}
			inner := &testRoundTripper{f: func(req *http.Request) (*http.Response, error) {
				// Closing the forwarded body must not close the original so it can still be rewound.
				_ = req.Body.Close()
				return &http.Response{StatusCode: http.StatusOK, Body: respBody}, nil
			}}
			tr := &rewindableTransport{inner: inner}

			req := mustCreateRequest(t, nil, http.MethodPost, "https://example.com", nil)
			req.Body = body

			resp, err := tr.RoundTrip(req)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if got := body.closed.Load(); got != 0 {
				t.Fatalf("request body closes before response body close: got %d, want 0", got)
			}

			for range 2 {
				if err := resp.Body.Close(); err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
			}

			if !respBody.closed.Load() {
				t.Error("response body closed: got false, want true")
			}
			if got := body.closed.Load(); got != 1 {
				t.Errorf("request body closes: got %d, want 1", got)
			}
		})

		for _, tt := range []struct {
			name string
			resp *http.Response
			err  error
		}{
			{
				name: "closes_read_seeker_body_if_inner_errors",
				err:  errors.New("boom"),
			},
			{
				name: "closes_read_seeker_body_if_response_body_nil",
				resp: &http.Response{StatusCode: http.StatusOK},
			},
		} {
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				body := &closeRecordingReadSeeker{Reader: strings.NewReader("TEST")}
				tr := &rewindableTransport{inner: &testRoundTripper{resp: tt.resp, err: tt.err}}

				req := mustCreateRequest(t, nil, http.MethodPost, "https://example.com", nil)
				req.Body = body

				resp, err := tr.RoundTrip(req)
				if !errors.Is(err, tt.err) {
					t.Fatalf("got %v, want %v", err, tt.err)
				}
				if resp != tt.resp {
					t.Errorf("response: got %v, want %v", resp, tt.resp)
				}
				if got := body.closed.Load(); got != 1 {
					t.Errorf("request body closes: got %d, want 1", got)
				}
			})
		}
	})
}

func Test_ensureRequestGetBody(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name        string
		req         *http.Request
		bufferBytes int64
		wantGetBody bool
		wantBody    *string
		wantCloser  bool
	}{
		{
			name: "with_nil_request",
		},
		{
			name: "with_nil_body",
			req:  &http.Request{},
		},
		{
			name: "with_no_body",
			req:  &http.Request{Body: http.NoBody},
		},
		{
			name:        "with_body_and_get_body",
			req:         &http.Request{Body: io.NopCloser(strings.NewReader(strings.Repeat("A", 4))), GetBody: func() (io.ReadCloser, error) { return io.NopCloser(strings.NewReader(strings.Repeat("A", 4))), nil }},
			wantGetBody: true,
			wantBody:    new(strings.Repeat("A", 4)),
		},
		{
			name:        "with_body_read_seeker",
			req:         &http.Request{Body: &testReadSeeker{strings.NewReader(strings.Repeat("A", 4))}},
			wantGetBody: true,
			wantBody:    new(strings.Repeat("A", 4)),
			wantCloser:  true,
		},
		{
			name:        "with_body_reader",
			req:         &http.Request{Body: io.NopCloser(&nonSeekableReader{r: strings.NewReader(strings.Repeat("A", 4))})},
			wantGetBody: true,
			wantBody:    new(strings.Repeat("A", 4)),
		},
		{
			name:        "with_body_reader_and_content_length",
			req:         &http.Request{Body: io.NopCloser(&nonSeekableReader{r: strings.NewReader(strings.Repeat("A", 4))}), ContentLength: 4},
			wantGetBody: true,
			wantBody:    new(strings.Repeat("A", 4)),
		},
		{
			name:        "with_body_reader_same_size_as_buffer",
			req:         &http.Request{Body: io.NopCloser(&nonSeekableReader{r: strings.NewReader(strings.Repeat("A", 16))})},
			bufferBytes: 16,
			wantGetBody: true,
			wantBody:    new(strings.Repeat("A", 16)),
		},
		{
			name:        "with_body_reader_same_size_as_buffer_and_content_length",
			req:         &http.Request{Body: io.NopCloser(&nonSeekableReader{r: strings.NewReader(strings.Repeat("A", 16))}), ContentLength: 16},
			bufferBytes: 16,
			wantGetBody: true,
			wantBody:    new(strings.Repeat("A", 16)),
		},
		{
			name:        "with_body_reader_larger_than_buffer",
			req:         &http.Request{Body: io.NopCloser(&nonSeekableReader{r: strings.NewReader(strings.Repeat("A", 16))})},
			bufferBytes: 8,
			wantBody:    new(strings.Repeat("A", 16)),
		},
		{
			name:        "with_body_reader_larger_than_buffer_and_content_length",
			req:         &http.Request{Body: io.NopCloser(&nonSeekableReader{r: strings.NewReader(strings.Repeat("A", 16))}), ContentLength: 16},
			bufferBytes: 8,
			wantBody:    new(strings.Repeat("A", 16)),
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			closer, err := ensureRequestGetBody(tt.req, tt.bufferBytes)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if (closer != nil) != tt.wantCloser {
				t.Errorf("closer non-nil: got %v, want %v", closer != nil, tt.wantCloser)
			}

			if tt.req != nil && (tt.req.GetBody != nil) != tt.wantGetBody {
				t.Errorf("GetBody non-nil: got %v, want %v", tt.req.GetBody != nil, tt.wantGetBody)
			}

			if tt.req == nil || tt.req.Body == nil || tt.req.Body == http.NoBody || tt.wantBody == nil {
				return
			}

			got, err := io.ReadAll(tt.req.Body)
			if err != nil {
				t.Fatalf("failed to read body: %v", err)
			}
			if string(got) != *tt.wantBody {
				t.Errorf("body: got %q, want %q", got, *tt.wantBody)
			}

			if tt.req.GetBody != nil {
				for range 2 {
					rc, err := tt.req.GetBody()
					if err != nil {
						t.Fatal(err)
					}
					again, _ := io.ReadAll(rc)
					if string(again) != *tt.wantBody {
						t.Errorf("GetBody body: got %q, want %q", again, *tt.wantBody)
					}
				}
			}
		})
	}

	t.Run("rewinds_read_seeker_body_from_start_offset", func(t *testing.T) {
		t.Parallel()

		r := &testReadSeeker{strings.NewReader("HEADERBODY")}
		if _, err := r.Seek(6, io.SeekStart); err != nil {
			t.Fatal(err)
		}

		req := &http.Request{Body: r}
		if _, err := ensureRequestGetBody(req, 0); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if req.ContentLength != 4 {
			t.Errorf("ContentLength: got %d, want 4", req.ContentLength)
		}

		for range 2 {
			rc, err := req.GetBody()
			if err != nil {
				t.Fatal(err)
			}
			got, _ := io.ReadAll(rc)
			if string(got) != "BODY" {
				t.Errorf("GetBody body: got %q, want %q", got, "BODY")
			}
		}
	})

	t.Run("errors_if_replaying_read_seeker_body_fails", func(t *testing.T) {
		t.Parallel()

		req := &http.Request{Body: &errorOnReplayReadSeeker{Reader: strings.NewReader("TEST")}, ContentLength: 4}
		if _, err := ensureRequestGetBody(req, 0); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if _, err := req.GetBody(); err == nil {
			t.Fatal("GetBody: got nil, want error")
		}
	})

	t.Run("rewinds_read_seeker_body_without_closing_it", func(t *testing.T) {
		t.Parallel()

		body := &closeRecordingReadSeeker{Reader: strings.NewReader("TEST")}
		req := &http.Request{Body: body}

		closer, err := ensureRequestGetBody(req, 0)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if closer != body {
			t.Fatalf("closer: got %v, want original body", closer)
		}

		if err := req.Body.Close(); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		for range 2 {
			rc, err := req.GetBody()
			if err != nil {
				t.Fatal(err)
			}
			got, _ := io.ReadAll(rc)
			if string(got) != "TEST" {
				t.Errorf("got %q, want %q", got, "TEST")
			}
			_ = rc.Close()
		}

		if got := body.closed.Load(); got != 0 {
			t.Errorf("original body closes: got %d, want 0", got)
		}
	})

	t.Run("errors_if_seeking_back_read_seeker_body_fails", func(t *testing.T) {
		t.Parallel()

		// Seeks are: current offset, end, back to start.
		req := &http.Request{Body: &failingSeekReadSeeker{Reader: strings.NewReader("TEST"), failOn: 3}}

		closer, err := ensureRequestGetBody(req, 0)
		if err == nil {
			t.Fatal("got nil, want error")
		}
		if want := "failed to seek request body"; !strings.Contains(err.Error(), want) {
			t.Errorf("got %q, want it to contain %q", err.Error(), want)
		}
		if closer != nil {
			t.Errorf("closer: got %v, want nil", closer)
		}
		if req.ContentLength != 0 {
			t.Errorf("ContentLength: got %d, want 0", req.ContentLength)
		}
		if req.GetBody != nil {
			t.Error("GetBody: got non-nil, want nil")
		}
	})

	t.Run("errors_if_reading_read_seeker_start_offset_fails", func(t *testing.T) {
		t.Parallel()

		req := &http.Request{Body: &failingSeekReadSeeker{Reader: strings.NewReader("TEST"), failOn: 1}}

		closer, err := ensureRequestGetBody(req, 0)
		if err == nil {
			t.Fatal("got nil, want error")
		}
		if want := "failed to seek request body"; !strings.Contains(err.Error(), want) {
			t.Errorf("got %q, want it to contain %q", err.Error(), want)
		}
		if closer != nil {
			t.Errorf("closer: got %v, want nil", closer)
		}
		if req.GetBody != nil {
			t.Error("GetBody: got non-nil, want nil")
		}
	})

	t.Run("rewinds_read_seeker_body_without_content_length_if_seeking_to_end_fails", func(t *testing.T) {
		t.Parallel()

		body := &failingSeekReadSeeker{Reader: strings.NewReader("TEST"), failOn: 2}
		req := &http.Request{Body: body}

		closer, err := ensureRequestGetBody(req, 0)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if closer != body {
			t.Errorf("closer: got %v, want original body", closer)
		}
		if req.ContentLength != 0 {
			t.Errorf("ContentLength: got %d, want 0", req.ContentLength)
		}

		got, err := io.ReadAll(req.Body)
		if err != nil {
			t.Fatalf("failed to read body: %v", err)
		}
		if string(got) != "TEST" {
			t.Errorf("body: got %q, want %q", got, "TEST")
		}

		for range 2 {
			rc, err := req.GetBody()
			if err != nil {
				t.Fatal(err)
			}
			again, _ := io.ReadAll(rc)
			if string(again) != "TEST" {
				t.Errorf("GetBody body: got %q, want %q", again, "TEST")
			}
		}
	})

	for _, tt := range []struct {
		name          string
		body          string
		contentLength int64
	}{
		{
			name:          "errors_if_body_shorter_than_content_length",
			body:          "TE",
			contentLength: 4,
		},
		{
			name:          "errors_if_body_empty_with_content_length",
			body:          "",
			contentLength: 4,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			req := &http.Request{Body: io.NopCloser(&nonSeekableReader{r: strings.NewReader(tt.body)}), ContentLength: tt.contentLength}

			_, err := ensureRequestGetBody(req, 0)
			if !errors.Is(err, io.ErrUnexpectedEOF) {
				t.Fatalf("got %v, want %v", err, io.ErrUnexpectedEOF)
			}
			if req.ContentLength != tt.contentLength {
				t.Errorf("ContentLength: got %d, want %d", req.ContentLength, tt.contentLength)
			}
			if req.GetBody != nil {
				t.Error("GetBody: got non-nil, want nil")
			}
		})
	}

	t.Run("uses_no_body_if_buffered_body_empty", func(t *testing.T) {
		t.Parallel()

		rec := &closeRecorder{}
		req := &http.Request{Body: rec}

		closer, err := ensureRequestGetBody(req, 0)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if closer != nil {
			t.Errorf("closer: got %v, want nil", closer)
		}
		if !rec.closed.Load() {
			t.Error("original body closed: got false, want true")
		}
		if req.Body != http.NoBody {
			t.Errorf("Body: got %T, want http.NoBody", req.Body)
		}
		if req.ContentLength != 0 {
			t.Errorf("ContentLength: got %d, want 0", req.ContentLength)
		}
		if req.GetBody == nil {
			t.Fatal("GetBody: got nil, want non-nil")
		}

		rc, err := req.GetBody()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if rc != http.NoBody {
			t.Errorf("GetBody: got %T, want http.NoBody", rc)
		}
	})

	for _, tt := range []struct {
		name          string
		contentLength int64
	}{
		{
			name:          "errors_if_reading_body_with_content_length_fails",
			contentLength: 4,
		},
		{
			name:          "errors_if_reading_body_without_content_length_fails",
			contentLength: 0,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			req := &http.Request{Body: io.NopCloser(errorReader{}), ContentLength: tt.contentLength}

			if _, err := ensureRequestGetBody(req, 0); err == nil {
				t.Fatal("got nil, want error")
			}
		})
	}
}

func Test_closeAfterReadCloser(t *testing.T) {
	t.Parallel()

	t.Run("Close", func(t *testing.T) {
		t.Parallel()

		errBody := errors.New("body close error")
		errAfter := errors.New("after close error")

		for _, tt := range []struct {
			name     string
			bodyErr  error
			afterErr error
			wantErr  error
		}{
			{
				name: "returns_nil_if_both_close",
			},
			{
				name:    "returns_body_error_and_still_closes_after",
				bodyErr: errBody,
				wantErr: errBody,
			},
			{
				name:     "returns_after_error_if_body_closes",
				afterErr: errAfter,
				wantErr:  errAfter,
			},
			{
				name:     "returns_body_error_if_both_fail",
				bodyErr:  errBody,
				afterErr: errAfter,
				wantErr:  errBody,
			},
		} {
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				body := &errorCloser{err: tt.bodyErr}
				after := &errorCloser{err: tt.afterErr}
				c := &closeAfterReadCloser{ReadCloser: body, after: after}

				if err := c.Close(); !errors.Is(err, tt.wantErr) {
					t.Fatalf("got %v, want %v", err, tt.wantErr)
				}

				if got := body.closed.Load(); got != 1 {
					t.Errorf("body closes: got %d, want 1", got)
				}
				if got := after.closed.Load(); got != 1 {
					t.Errorf("after closes: got %d, want 1", got)
				}

				// A second close only reaches the wrapped body; the after closer is never closed again or reported again.
				if err := c.Close(); !errors.Is(err, tt.bodyErr) {
					t.Fatalf("second close: got %v, want %v", err, tt.bodyErr)
				}

				if got := after.closed.Load(); got != 1 {
					t.Errorf("after closes after second close: got %d, want 1", got)
				}
			})
		}
	})
}
