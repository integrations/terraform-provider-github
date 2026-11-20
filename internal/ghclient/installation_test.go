package ghclient

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"

	"github.com/jferrl/go-githubauth"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

func TestInstallationTokenStalledBody(t *testing.T) {
	t.Parallel()
	flushed := make(chan struct{})
	stopped := make(chan struct{})
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusCreated)
		writer.(http.Flusher).Flush()
		close(flushed)
		select {
		case <-request.Context().Done():
		case <-stopped:
		}
	}))
	defer server.Close()
	defer close(stopped)
	source := newInstallationTokenSource(42, oauth2.StaticTokenSource(&oauth2.Token{AccessToken: "test-jwt"}), ClientOptions{}, githubauth.WithBaseURL(server.URL))
	if source.client.Timeout != installationTokenTimeout || source.client.Timeout <= 0 {
		t.Fatal("installation token HTTP client must have a finite total timeout")
	}
	source.client.Timeout = 100 * time.Millisecond
	result := make(chan error, 1)
	go func() {
		_, err := source.Token()
		result <- err
	}()
	select {
	case <-flushed:
	case <-time.After(2 * time.Second):
		t.Fatal("token request did not reach flushed headers")
	}
	select {
	case err := <-result:
		if !errors.Is(err, context.DeadlineExceeded) {
			t.Fatalf("expected total timeout reading stalled body, got %v", err)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("token body read exceeded total timeout")
	}
}

func TestAppInstallationOperationCancellationAndRenewal(t *testing.T) {
	t.Parallel()
	for _, protocol := range []string{"REST", "GraphQL"} {
		for _, isGHES := range []bool{false, true} {
			t.Run(fmt.Sprintf("%s/GHES=%t", protocol, isGHES), func(t *testing.T) {
				t.Parallel()
				var tokenPosts atomic.Int32
				flushed := make(chan struct{})
				disconnected := make(chan struct{})
				stopped := make(chan struct{})
				tokenPath := "/app/installations/42/access_tokens"
				if isGHES {
					tokenPath = "/api/v3" + tokenPath
				}
				server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
					writer.Header().Set("Content-Type", "application/json")
					if request.URL.Path == tokenPath {
						if request.Method != http.MethodPost {
							t.Errorf("token method = %s, want POST", request.Method)
						}
						attempt := tokenPosts.Add(1)
						if attempt == 1 {
							writer.WriteHeader(http.StatusCreated)
							writer.(http.Flusher).Flush()
							close(flushed)
							select {
							case <-request.Context().Done():
								close(disconnected)
							case <-stopped:
							}
							return
						}
						expiry := time.Now().Add(time.Hour)
						if attempt == 2 {
							expiry = time.Now().Add(10 * time.Second)
						}
						fmt.Fprintf(writer, `{"token":"token-%d","expires_at":%q}`, attempt, expiry.UTC().Format(time.RFC3339))
						return
					}
					apiPath := "/user"
					if isGHES {
						apiPath = "/api/v3/user"
					}
					if protocol == "GraphQL" {
						apiPath = "/graphql"
						if isGHES {
							apiPath = "/api/graphql"
						}
					}
					if request.URL.Path != apiPath {
						t.Errorf("unexpected API path %s", request.URL.Path)
						http.NotFound(writer, request)
						return
					}
					if got, want := request.Header.Get("Authorization"), fmt.Sprintf("Bearer token-%d", tokenPosts.Load()); got != want {
						t.Errorf("API authorization = %q, want %q", got, want)
					}
					fmt.Fprint(writer, `{"data":{"viewer":{"login":"test-user"}}}`)
				}))
				defer server.Close()
				defer close(stopped)
				opts := ClientOptions{BaseURL: server.URL, IsGHES: isGHES}
				var call func(context.Context) error
				if protocol == "REST" {
					client, err := NewAppRESTClient("123456789", mustReadAppPrivateKey(t), new(int64(42)), opts)
					if err != nil {
						t.Fatal(err)
					}
					call = func(ctx context.Context) error {
						request, err := client.NewRequest(ctx, http.MethodGet, "user", nil)
						if err != nil {
							return err
						}
						_, err = client.Do(request, &struct{}{})
						return err
					}
				} else {
					client, err := NewAppGraphQLClient("123456789", mustReadAppPrivateKey(t), new(int64(42)), opts)
					if err != nil {
						t.Fatal(err)
					}
					call = func(ctx context.Context) error {
						query := struct {
							Viewer struct{ Login githubv4.String }
						}{}
						return client.Query(ctx, &query, nil)
					}
				}
				ctx, cancel := context.WithCancel(t.Context())
				defer cancel()
				result := make(chan error, 1)
				go func() { result <- call(ctx) }()
				select {
				case <-flushed:
				case <-time.After(2 * time.Second):
					t.Fatal("token request did not flush headers")
				}
				waitingCtx, cancelWaiting := context.WithTimeout(t.Context(), 100*time.Millisecond)
				defer cancelWaiting()
				waiting := make(chan error, 1)
				go func() { waiting <- call(waitingCtx) }()
				select {
				case err := <-waiting:
					if !errors.Is(err, context.DeadlineExceeded) {
						t.Fatalf("waiting caller: expected deadline, got %v", err)
					}
				case <-time.After(2 * time.Second):
					t.Fatal("waiting caller ignored its deadline")
				}
				cancel()
				select {
				case err := <-result:
					if !errors.Is(err, context.Canceled) {
						t.Fatalf("expected operation cancellation, got %v", err)
					}
				case <-time.After(2 * time.Second):
					t.Fatal("stalled token body ignored operation cancellation")
				}
				select {
				case <-disconnected:
				case <-time.After(2 * time.Second):
					t.Fatal("token HTTP request survived operation cancellation")
				}
				for _, wantPosts := range []int32{2, 3, 3} {
					if err := call(t.Context()); err != nil {
						t.Fatalf("reusable client failed after cancellation: %v", err)
					}
					if got := tokenPosts.Load(); got != wantPosts {
						t.Fatalf("token POSTs = %d, want %d (refresh then reuse)", got, wantPosts)
					}
				}
			})
		}
	}
}
