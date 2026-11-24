package httpx

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

func mustCreateRequest(t *testing.T, ctx context.Context, method, url string, body io.Reader) *http.Request {
	t.Helper()

	if ctx == nil {
		ctx = t.Context()
	}

	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	return req
}

type fakeHTTPResponse struct {
	statusCode int
	headers    map[string]string
	body       []byte
}

type dummyRoundTripper struct {
	http.RoundTripper
}

type testRoundTripper struct {
	f      func(req *http.Request) (*http.Response, error)
	delay  time.Duration
	called atomic.Int32
	resp   *http.Response
	err    error
}

func (tr *testRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if tr.f != nil {
		return tr.f(req)
	}

	if tr.delay > 0 {
		ctx := req.Context()
		select {
		case <-time.After(tr.delay):
		case <-ctx.Done():
		}
	}

	tr.called.Add(1)

	if tr.err != nil {
		return nil, tr.err
	}

	return tr.resp, nil
}

type closeRecorder struct {
	closed atomic.Bool
}

func (c *closeRecorder) Read(p []byte) (int, error) { return 0, io.EOF }

func (c *closeRecorder) Close() error {
	c.closed.Store(true)
	return nil
}

type testReadSeeker struct {
	*strings.Reader
}

func (testReadSeeker) Close() error { return nil }

type errorReadSeeker struct {
	io.Closer
}

func (rs *errorReadSeeker) Read(p []byte) (n int, err error) {
	return 0, fmt.Errorf("error reading")
}

func (rs *errorReadSeeker) Seek(offset int64, whence int) (int64, error) {
	return 0, fmt.Errorf("error seeking")
}

// errorOnReplayReadSeeker is a read seeker where only the first Seek succeeds, which is how the start offset is read, so replaying the body fails.
type errorOnReplayReadSeeker struct {
	*strings.Reader
	seeks int
}

func (errorOnReplayReadSeeker) Close() error { return nil }

func (rs *errorOnReplayReadSeeker) Seek(offset int64, whence int) (int64, error) {
	rs.seeks++
	if rs.seeks > 1 {
		return 0, fmt.Errorf("error seeking")
	}

	return rs.Reader.Seek(offset, whence)
}

// failingSeekReadSeeker is a read seeker where the Seek call numbered failOn fails and all others succeed.
type failingSeekReadSeeker struct {
	*strings.Reader
	failOn int
	seeks  int
}

func (rs *failingSeekReadSeeker) Close() error { return nil }

func (rs *failingSeekReadSeeker) Seek(offset int64, whence int) (int64, error) {
	rs.seeks++
	if rs.seeks == rs.failOn {
		return 0, fmt.Errorf("error seeking")
	}

	return rs.Reader.Seek(offset, whence)
}

// closeRecordingReadSeeker is a read seeker that records how many times it has been closed.
type closeRecordingReadSeeker struct {
	*strings.Reader
	closed atomic.Int32
}

func (rs *closeRecordingReadSeeker) Close() error {
	rs.closed.Add(1)
	return nil
}

// errorCloser is a read closer that records how many times it has been closed and returns err from Close.
type errorCloser struct {
	err    error
	closed atomic.Int32
}

func (c *errorCloser) Read(p []byte) (int, error) { return 0, io.EOF }

func (c *errorCloser) Close() error {
	c.closed.Add(1)
	return c.err
}

type errorReader struct{}

func (errorReader) Read(p []byte) (int, error) {
	return 0, fmt.Errorf("error reading")
}

type nonSeekableReader struct {
	r io.Reader
}

func (n *nonSeekableReader) Read(p []byte) (int, error) {
	return n.r.Read(p)
}

type testBackoff struct {
	next time.Duration
	stop bool
}

func (b *testBackoff) Next() (time.Duration, bool) {
	return b.next, b.stop
}
