package github

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"testing/iotest"
	"testing/synctest"
	"time"

	"github.com/google/go-github/v92/github"
)

func TestEtagTransport(t *testing.T) {
	// IMPORTANT: This test is not parallelized because it uses global state.

	ts := githubApiMock([]*mockResponse{
		{
			ExpectedUri: "/repos/test/blah",
			ExpectedHeaders: map[string]string{
				"If-None-Match": "something",
			},

			ResponseBody: `{"id": 1234}`,
			StatusCode:   200,
		},
	})
	defer ts.Close()

	client := mustCreateTestGitHubClient(t, ts.URL, github.WithTransport(NewEtagTransport(http.DefaultTransport)))
	ctx := context.WithValue(t.Context(), ctxEtag, "something")
	r, _, err := client.Repositories.Get(ctx, "test", "blah")
	if err != nil {
		t.Fatal(err)
	}

	if r.GetID() != 1234 {
		t.Fatalf("Expected ID to be 1234, got: %d", r.GetID())
	}
}

func githubApiMock(responseSequence []*mockResponse) *httptest.Server {
	position := new(0)
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("Server", "GitHub.com")

		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("[DEBUG] Error: %s", err)
		}
		log.Printf("[DEBUG] Mock server received %s request to %q; headers:\n%s\nrequest body: %q\n",
			r.Method, r.RequestURI, r.Header, string(bodyBytes))

		i := *position
		if len(responseSequence) < i+1 {
			w.WriteHeader(400)
			return
		}

		tc := responseSequence[i]

		headersMatch := func(h http.Header, expectedHeaders map[string]string) bool {
			for key, value := range expectedHeaders {
				if h.Get(key) != value {
					return false
				}
			}
			return true
		}

		if r.RequestURI != tc.ExpectedUri {
			log.Printf("[DEBUG] Error: expected URI: %q, given: %q", tc.ExpectedUri, r.RequestURI)
			w.WriteHeader(400)
			return
		}
		if !headersMatch(r.Header, tc.ExpectedHeaders) {
			log.Printf("[DEBUG] Error: expected headers: %q, given: %q", tc.ExpectedHeaders, r.Header)
			w.WriteHeader(400)
			return
		}
		if tc.ExpectedMethod != "" && r.Method != tc.ExpectedMethod {
			log.Printf("[DEBUG] Error: expected method: %q, given: %q", tc.ExpectedMethod, r.Method)
			w.WriteHeader(400)
			return
		}
		if len(tc.ExpectedBody) > 0 && string(bodyBytes) != string(tc.ExpectedBody) {
			log.Printf("[DEBUG] Error: expected body: %q, given: %q",
				string(tc.ExpectedBody), string(bodyBytes))
			w.WriteHeader(400)
			return
		}

		for key, value := range tc.ResponseHeaders {
			w.Header().Set(key, value)
		}
		w.WriteHeader(tc.StatusCode)
		fmt.Fprintln(w, tc.ResponseBody)

		// Treat response as disposable
		position = new(i + 1)
	}))
}

func TestRateLimitTransport_abuseLimit_get(t *testing.T) {
	// IMPORTANT: This test is not parallelized because it uses global state.

	ts := githubApiMock([]*mockResponse{
		{
			ExpectedUri: "/repos/test/blah",
			ResponseBody: `{
  "message": "You have triggered an abuse detection mechanism and have been temporarily blocked from content creation. Please retry your request again later.",
  "documentation_url": "https://developer.github.com/v3/#abuse-rate-limits"
}`,
			StatusCode: 403,
			ResponseHeaders: map[string]string{
				"Retry-After": "0.1",
			},
		},
		{
			ExpectedUri: "/repos/test/blah",
			ResponseBody: `{
  "message": "You have triggered an abuse detection mechanism and have been temporarily blocked from content creation. Please retry your request again later.",
  "documentation_url": "https://developer.github.com/v3/#abuse-rate-limits"
}`,
			StatusCode: 403,
			ResponseHeaders: map[string]string{
				"Retry-After": "0.1",
			},
		},
		{
			ExpectedUri:  "/repos/test/blah",
			ResponseBody: `{"id": 1234}`,
			StatusCode:   200,
		},
	})
	defer ts.Close()

	client := mustCreateTestGitHubClient(t, ts.URL, github.WithTransport(NewRateLimitTransport(http.DefaultTransport)))

	ctx := context.WithValue(t.Context(), ctxId, t.Name())
	r, _, err := client.Repositories.Get(ctx, "test", "blah")
	if err != nil {
		t.Fatal(err)
	}

	if r.GetID() != 1234 {
		t.Fatalf("Expected ID to be 1234, got: %d", r.GetID())
	}
}

func TestRateLimitTransport_abuseLimit_get_cancelled(t *testing.T) {
	// IMPORTANT: This test is not parallelized because it uses global state.

	ts := githubApiMock([]*mockResponse{
		{
			ExpectedUri: "/repos/test/blah",
			ResponseBody: `{
  "message": "You have triggered an abuse detection mechanism and have been temporarily blocked from content creation. Please retry your request again later.",
  "documentation_url": "https://developer.github.com/v3/#abuse-rate-limits"
}`,
			StatusCode: 403,
			ResponseHeaders: map[string]string{
				"Retry-After": "10",
			},
		},
	})
	defer ts.Close()

	client := mustCreateTestGitHubClient(t, ts.URL, github.WithTransport(NewRateLimitTransport(http.DefaultTransport)))

	ctx, cancel := context.WithTimeout(t.Context(), 100*time.Millisecond)
	defer cancel()

	start := time.Now()
	_, _, err := client.Repositories.Get(ctx, "test", "blah")
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("Expected context deadline exceeded, got: %v", err)
	}
	if time.Since(start) > time.Second {
		t.Fatalf("Waited for longer than expected: %s", time.Since(start))
	}
}

func TestRateLimitTransport_abuseLimit_post(t *testing.T) {
	// IMPORTANT: This test is not parallelized because it uses global state.

	ts := githubApiMock([]*mockResponse{
		{
			ExpectedUri:    "/orgs/tada/repos",
			ExpectedMethod: "POST",
			ExpectedBody: []byte(`{"name":"radek-example-48","description":""}
`),
			ResponseBody: `{
  "message": "You have triggered an abuse detection mechanism and have been temporarily blocked from content creation. Please retry your request again later.",
  "documentation_url": "https://developer.github.com/v3/#abuse-rate-limits"
}`,
			StatusCode: 403,
			ResponseHeaders: map[string]string{
				"Retry-After": "0.1",
			},
		},
		{
			ExpectedUri:    "/orgs/tada/repos",
			ExpectedMethod: "POST",
			ExpectedBody: []byte(`{"name":"radek-example-48","description":""}
`),
			ResponseBody: `{"id": 1234}`,
			StatusCode:   200,
		},
	})
	defer ts.Close()

	client := mustCreateTestGitHubClient(t, ts.URL, github.WithTransport(NewRateLimitTransport(http.DefaultTransport)))

	ctx := context.WithValue(t.Context(), ctxId, t.Name())
	r, _, err := client.Repositories.Create(ctx, "tada", &github.Repository{
		Name:        new("radek-example-48"),
		Description: new(""),
	})
	if err != nil {
		t.Fatal(err)
	}

	if r.GetID() != 1234 {
		t.Fatalf("Expected ID to be 1234, got: %d", r.GetID())
	}
}

func TestRateLimitTransport_abuseLimit_post_error(t *testing.T) {
	// IMPORTANT: This test is not parallelized because it uses global state.

	ts := githubApiMock([]*mockResponse{
		{
			ExpectedUri:    "/orgs/tada/repos",
			ExpectedMethod: "POST",
			ExpectedBody: []byte(`{"name":"radek-example-48","description":""}
`),
			ResponseBody: `{
  "message": "You have triggered an abuse detection mechanism and have been temporarily blocked from content creation. Please retry your request again later.",
  "documentation_url": "https://developer.github.com/v3/#abuse-rate-limits"
}`,
			StatusCode: 403,
			ResponseHeaders: map[string]string{
				"Retry-After": "0.1",
			},
		},
		{
			ExpectedUri:    "/orgs/tada/repos",
			ExpectedMethod: "POST",
			ExpectedBody: []byte(`{"name":"radek-example-48","description":""}
`),
			ResponseBody: `{
  "message": "Repository creation failed.",
  "errors": [
    {
      "resource": "Repository",
      "code": "custom",
      "field": "name",
      "message": "name already exists on this account"
    }
  ],
  "documentation_url": "https://developer.github.com/v3/repos/#create"
}
`,
			StatusCode: 422,
		},
	})
	defer ts.Close()

	client := mustCreateTestGitHubClient(t, ts.URL, github.WithTransport(NewRateLimitTransport(http.DefaultTransport)))

	ctx := context.WithValue(t.Context(), ctxId, t.Name())
	_, _, err := client.Repositories.Create(ctx, "tada", &github.Repository{
		Name:        new("radek-example-48"),
		Description: new(""),
	})
	if err == nil {
		t.Fatal("Expected 422 error, got nil")
	}

	var ghErr *github.ErrorResponse
	ok := errors.As(err, &ghErr)
	if !ok {
		t.Fatalf("Expected github.ErrorResponse, got: %#v", err)
	}

	expectedMessage := "Repository creation failed."
	if ghErr.Message != expectedMessage {
		t.Fatalf("Expected message %q, got: %q", expectedMessage, ghErr.Message)
	}
}

func TestRateLimitTransport_smart_lock(t *testing.T) {
	// IMPORTANT: This test is not parallelized because it uses global state.

	t.Run("With parallelRequests true it does not lock the rate limit transport", func(t *testing.T) {
		rlt := NewRateLimitTransport(http.DefaultTransport, WithParallelRequests(true))

		isSuccess := make(chan bool)
		go func() {
			rlt.m.Lock()
			rlt.smartLock(true)
			rlt.m.Unlock()
			isSuccess <- true
		}()
		select {
		case <-isSuccess:
		case <-time.After(100 * time.Millisecond):
			t.Fatalf("Expected to succeed instantly, waited 100 milliseconds unsuccessfully")
		}
	})

	t.Run("With parallelRequests true it should not unlock the rate limit transport", func(t *testing.T) {
		rlt := NewRateLimitTransport(http.DefaultTransport, WithParallelRequests(true))

		isSuccess := make(chan bool)
		go func() {
			rlt.m.Lock()
			rlt.smartLock(false)
			rlt.m.Unlock()
			isSuccess <- true
		}()
		select {
		case <-isSuccess:
		case <-time.After(100 * time.Millisecond):
			t.Fatalf("Expected to succeed instantly, waited 100 milliseconds unsuccessfully")
		}
	})

	t.Run("With parallelRequests false with a lock present it should get stuck waiting", func(t *testing.T) {
		rlt := NewRateLimitTransport(http.DefaultTransport, WithParallelRequests(false))

		isSuccess := make(chan bool)
		go func() {
			rlt.m.Lock()
			rlt.smartLock(true)
			isSuccess <- true
		}()
		select {
		case <-isSuccess:
			t.Fatalf("Expected get stuck waiting but it acquired the lock successfully")
		case <-time.After(100 * time.Millisecond):
		}
	})

	t.Run("With parallelRequests false and a lock present it should be able to unlock the rate limit transport", func(t *testing.T) {
		rlt := NewRateLimitTransport(http.DefaultTransport, WithParallelRequests(false))

		isSuccess := make(chan bool)
		go func() {
			rlt.m.Lock()
			rlt.smartLock(false)
			rlt.m.Lock()
			defer rlt.m.Unlock()
			isSuccess <- true
		}()
		select {
		case <-isSuccess:
		case <-time.After(100 * time.Millisecond):
			t.Fatalf("Expected to succeed instantly, waited 100 milliseconds unsuccessfully")
		}
	})
}

func TestRetryTransport_retry_post_error(t *testing.T) {
	// IMPORTANT: This test is not parallelized because it uses global state.

	ts := githubApiMock([]*mockResponse{
		{
			ExpectedUri:    "/orgs/tada/repos",
			ExpectedMethod: "POST",
			ExpectedBody: []byte(`{"name":"radek-example-48","description":""}
`),
			ResponseBody: `{
  "message": "internal server error"
}`,
			StatusCode: 500,
		},
		{
			ExpectedUri:    "/orgs/tada/repos",
			ExpectedMethod: "POST",
			ExpectedBody: []byte(`{"name":"radek-example-48","description":""}
`),
			ResponseBody: `{
  "message": "internal server error"
}`,
			StatusCode: 500,
		},
		{
			ExpectedUri:    "/orgs/tada/repos",
			ExpectedMethod: "POST",
			ExpectedBody: []byte(`{"name":"radek-example-48","description":""}
`),
			ResponseBody: `{
  "message": "internal server error"
}`,
			StatusCode: 201,
		},
	})
	defer ts.Close()

	client := mustCreateTestGitHubClient(t, ts.URL, github.WithTransport(NewRetryTransport(http.DefaultTransport, WithMaxRetries(1))))

	ctx := context.WithValue(t.Context(), ctxId, t.Name())
	_, _, err := client.Repositories.Create(ctx, "tada", &github.Repository{
		Name:        new("radek-example-48"),
		Description: new(""),
	})
	if err == nil {
		t.Fatal("Expected error not to be nil")
	}

	var ghErr *github.ErrorResponse
	ok := errors.As(err, &ghErr)
	if !ok {
		t.Fatalf("Expected github.ErrorResponse, got: %#v", err)
	}

	expectedMessage := "internal server error"
	if ghErr.Message != expectedMessage {
		t.Fatalf("Expected message %q, got: %q", expectedMessage, ghErr.Message)
	}
}

func TestRetryTransport_retry_post_success(t *testing.T) {
	// IMPORTANT: This test is not parallelized because it uses global state.

	ts := githubApiMock([]*mockResponse{
		{
			ExpectedUri:    "/orgs/tada/repos",
			ExpectedMethod: "POST",
			ExpectedBody: []byte(`{"name":"radek-example-48","description":""}
`),
			ResponseBody: `{
  "message": "internal server error"
}`,
			StatusCode: 500,
		},
		{
			ExpectedUri:    "/orgs/tada/repos",
			ExpectedMethod: "POST",
			ExpectedBody: []byte(`{"name":"radek-example-48","description":""}
`),
			ResponseBody: `{
  "message": "internal server error"
}`,
			StatusCode: 500,
		},
		{
			ExpectedUri:    "/orgs/tada/repos",
			ExpectedMethod: "POST",
			ExpectedBody: []byte(`{"name":"radek-example-48","description":""}
`),
			ResponseBody: `{
  "message": "Resource created"
}`,
			StatusCode: 201,
		},
	})
	defer ts.Close()

	client := mustCreateTestGitHubClient(t, ts.URL, github.WithTransport(NewRetryTransport(http.DefaultTransport, WithMaxRetries(2), WithRetryDelay(time.Second))))

	ctx := context.WithValue(t.Context(), ctxId, t.Name())
	_, _, err := client.Repositories.Create(ctx, "tada", &github.Repository{
		Name:        new("radek-example-48"),
		Description: new(""),
	})
	if err != nil {
		t.Fatalf("Expected error to be nil, got %v", err)
	}

	var ghErr *github.ErrorResponse
	ok := errors.As(err, &ghErr)
	if ok {
		t.Fatalf("Expected successful github call, got: %q", ghErr.Message)
	}
}

func TestRetryTransport_cancelled(t *testing.T) {
	for _, deadline := range []bool{false, true} {
		t.Run(fmt.Sprintf("deadline_%t", deadline), func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				ctx, cancel := context.WithTimeout(t.Context(), time.Second)
				defer cancel()
				body := &recordingBody{Reader: strings.NewReader("retry")}
				calls := 0
				rt := NewRetryTransport(retryTestTransport(func(*http.Request) (*http.Response, error) {
					calls++
					return &http.Response{StatusCode: http.StatusInternalServerError, Body: body}, nil
				}), WithMaxRetries(2), WithRetryDelay(time.Hour))
				req := httptest.NewRequest(http.MethodGet, "http://example.com/", nil).WithContext(ctx)
				done := make(chan error, 1)
				go func() {
					resp, err := rt.RoundTrip(req)
					if resp != nil {
						_ = resp.Body.Close()
					}
					done <- err
				}()

				// Wait until the first response has been handled and the retry
				// goroutine is blocked in its delay, before cancelling it.
				synctest.Wait()
				if calls != 1 || body.closes != 1 || body.bytesRead != len("retry") {
					t.Fatalf("Before cancellation: calls = %d, closes = %d, bytes read = %d", calls, body.closes, body.bytesRead)
				}
				select {
				case err := <-done:
					t.Fatalf("RoundTrip returned before cancellation: %v", err)
				default:
				}
				wantErr := context.Canceled
				if deadline {
					wantErr = context.DeadlineExceeded
					time.Sleep(time.Second)
				} else {
					cancel()
				}
				synctest.Wait()
				select {
				case err := <-done:
					if !errors.Is(err, wantErr) {
						t.Fatalf("Got %v, want %v", err, wantErr)
					}
				default:
					t.Fatal("RoundTrip did not return after cancellation")
				}
				if calls != 1 || body.closes != 1 {
					t.Fatalf("After cancellation: calls = %d, closes = %d; want 1, 1", calls, body.closes)
				}
			})
		})
	}
}

func TestRetryTransport_no_sleep_after_last_retry(t *testing.T) {
	for _, maxRetries := range []int{0, 2} {
		t.Run(fmt.Sprint(maxRetries), func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				var bodies []*recordingBody
				rt := NewRetryTransport(retryTestTransport(func(*http.Request) (*http.Response, error) {
					body := &recordingBody{Reader: strings.NewReader("final error")}
					bodies = append(bodies, body)
					return &http.Response{StatusCode: http.StatusInternalServerError, Body: body}, nil
				}), WithMaxRetries(maxRetries), WithRetryDelay(time.Second))
				start := time.Now()
				resp, err := rt.RoundTrip(httptest.NewRequest(http.MethodGet, "http://example.com/", nil))
				if err != nil || resp == nil || resp.StatusCode != http.StatusInternalServerError {
					t.Fatalf("Got (%v, %v), want final 500 response without transport error", resp, err)
				}
				defer resp.Body.Close()
				if elapsed := time.Since(start); elapsed != time.Duration(maxRetries)*time.Second {
					t.Fatalf("Retry delays = %s, want %s", elapsed, time.Duration(maxRetries)*time.Second)
				}
				if len(bodies) != maxRetries+1 {
					t.Fatalf("Attempts = %d, want %d", len(bodies), maxRetries+1)
				}
				for i, body := range bodies[:len(bodies)-1] {
					if body.closes != 1 || body.bytesRead != len("final error") {
						t.Fatalf("Intermediate body %d: closes = %d, bytes read = %d", i, body.closes, body.bytesRead)
					}
				}
				final := bodies[len(bodies)-1]
				if resp.Body != final || final.closed || final.reads != 0 {
					t.Fatal("Final error body was replaced, closed or consumed")
				}
				data, err := io.ReadAll(resp.Body)
				if err != nil || string(data) != "final error" {
					t.Fatalf("Final body = %q, err = %v", data, err)
				}
			})
		})
	}
}

type recordingBody struct {
	io.Reader
	closed    bool
	reads     int
	closes    int
	bytesRead int
}

func (b *recordingBody) Read(p []byte) (int, error) {
	b.reads++
	n, err := b.Reader.Read(p)
	b.bytesRead += n
	return n, err
}

func (b *recordingBody) Close() error {
	b.closed = true
	b.closes++
	return nil
}

type staticTransport struct {
	responses []*http.Response
	errs      []error
	calls     int
}

func (t *staticTransport) RoundTrip(*http.Request) (*http.Response, error) {
	resp := t.responses[t.calls]
	var err error
	if t.errs != nil {
		err = t.errs[t.calls]
	}
	t.calls++
	return resp, err
}

func TestRetryTransport_closes_intermediate_response_bodies(t *testing.T) {
	statuses := []int{500, 500, 200}
	bodies := make([]*recordingBody, len(statuses))
	responses := make([]*http.Response, len(statuses))
	for i, status := range statuses {
		bodies[i] = &recordingBody{Reader: strings.NewReader("{}")}
		responses[i] = &http.Response{StatusCode: status, Body: bodies[i]}
	}

	rt := NewRetryTransport(&staticTransport{responses: responses},
		WithMaxRetries(2),
		WithRetryDelay(time.Millisecond),
	)

	resp, err := rt.RoundTrip(httptest.NewRequest(http.MethodGet, "http://example.com/", nil))
	if err != nil {
		t.Fatalf("Expected error to be nil, got: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status %d, got: %d", http.StatusOK, resp.StatusCode)
	}

	for i, body := range bodies {
		// The final response is owned by the caller, intermediate ones must be closed.
		wantClosed := i < len(bodies)-1
		if body.closed != wantClosed {
			t.Fatalf("Response %d body closed = %t, want %t", i, body.closed, wantClosed)
		}
	}
}

func TestRetryTransport_closes_response_returned_with_error(t *testing.T) {
	transportErr := errors.New("transport error")
	firstBody := &recordingBody{Reader: strings.NewReader("{}")}
	secondBody := &recordingBody{Reader: strings.NewReader("{}")}
	rt := NewRetryTransport(&staticTransport{
		responses: []*http.Response{
			{StatusCode: http.StatusInternalServerError, Body: firstBody},
			{StatusCode: http.StatusOK, Body: secondBody},
		},
		errs: []error{transportErr, nil},
	},
		WithMaxRetries(1),
		WithRetryDelay(time.Millisecond),
	)

	resp, err := rt.RoundTrip(httptest.NewRequest(http.MethodGet, "http://example.com/", nil))
	if err != nil {
		t.Fatalf("Expected error to be nil, got: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status %d, got: %d", http.StatusOK, resp.StatusCode)
	}
	if !firstBody.closed {
		t.Fatal("Expected response body returned with an error to be closed")
	}
	if secondBody.closed {
		t.Fatal("Expected final response body not to be closed")
	}
}

type retryTestTransport func(*http.Request) (*http.Response, error)

func (f retryTestTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func TestRetryTransport_cancelled_request_body(t *testing.T) {
	for _, kind := range []string{"nil", "no_body", "stream", "get_body"} {
		t.Run(kind, func(t *testing.T) {
			body := &recordingBody{Reader: strings.NewReader("request")}
			req := httptest.NewRequest(http.MethodPost, "http://example.com/", nil)
			switch kind {
			case "nil":
				req.Body = nil
			case "no_body":
				req.Body = http.NoBody
			default:
				req.Body = body
			}
			if kind == "get_body" {
				req.GetBody = func() (io.ReadCloser, error) {
					t.Fatal("GetBody called on a cancelled request")
					return nil, nil
				}
			}
			ctx, cancel := context.WithCancel(t.Context())
			cancel()
			req = req.WithContext(ctx)
			rt := NewRetryTransport(retryTestTransport(func(*http.Request) (*http.Response, error) {
				t.Fatal("Underlying transport called on a cancelled request")
				return nil, nil
			}), WithMaxRetries(1))

			resp, err := rt.RoundTrip(req)
			if resp != nil || !errors.Is(err, context.Canceled) {
				t.Fatalf("Got (%v, %v), want (nil, context.Canceled)", resp, err)
			}
			if kind == "stream" || kind == "get_body" {
				if body.closes != 1 || body.reads != 0 {
					t.Fatalf("Body closes = %d, reads = %d; want 1, 0", body.closes, body.reads)
				}
			}
		})
	}
}

func TestRetryTransport_cancelled_after_attempt(t *testing.T) {
	ctx, cancel := context.WithCancel(t.Context())
	defer cancel()
	body := &recordingBody{Reader: strings.NewReader("request")}
	t.Cleanup(func() { _ = body.Close() })
	req := httptest.NewRequest(http.MethodPost, "http://example.com/", nil).WithContext(ctx)
	req.GetBody = func() (io.ReadCloser, error) { return body, nil }
	calls := 0
	rt := NewRetryTransport(retryTestTransport(func(*http.Request) (*http.Response, error) {
		calls++
		// The underlying transport may close the body asynchronously, after returning.
		cancel()
		return nil, context.Canceled
	}), WithMaxRetries(1), WithRetryDelay(0))

	resp, err := rt.RoundTrip(req)
	if resp != nil || !errors.Is(err, context.Canceled) || calls != 1 {
		t.Fatalf("Got (%v, %v), calls = %d; want (nil, context.Canceled), 1", resp, err, calls)
	}
	if body.closes != 0 || body.reads != 0 {
		t.Fatalf("Sent body closes = %d, reads = %d; want 0, 0", body.closes, body.reads)
	}
}

func TestDrainResponseBody(t *testing.T) {
	for _, size := range []int{0, 100, 4095, 4096, 4097, 8192} {
		t.Run(fmt.Sprint(size), func(t *testing.T) {
			body := &recordingBody{Reader: strings.NewReader(strings.Repeat("x", size))}
			drainResponseBody(&http.Response{Body: body})
			if body.bytesRead != min(size, 4096) || body.closes != 1 {
				t.Fatalf("Read %d bytes, closed %d times; want %d, 1", body.bytesRead, body.closes, min(size, 4096))
			}
		})
	}
	t.Run("read_error", func(t *testing.T) {
		body := &recordingBody{Reader: iotest.ErrReader(errors.New("response read failed"))}
		drainResponseBody(&http.Response{Body: body})
		if body.closes != 1 {
			t.Fatal("Response body was not closed after a read error")
		}
	})
	t.Run("empty_response", func(t *testing.T) {
		drainResponseBody(nil)
		drainResponseBody(&http.Response{})
		drainResponseBody(&http.Response{Body: http.NoBody})
	})
}

type mockResponse struct {
	ExpectedUri     string
	ExpectedMethod  string
	ExpectedHeaders map[string]string
	ExpectedBody    []byte

	StatusCode      int
	ResponseHeaders map[string]string
	ResponseBody    string
}
