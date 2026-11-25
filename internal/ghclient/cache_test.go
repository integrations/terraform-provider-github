package ghclient

import (
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
	"testing/iotest"

	"github.com/google/go-cmp/cmp"
	"go.etcd.io/bbolt"
	bbolterrs "go.etcd.io/bbolt/errors"
)

func Test_createCacheStore(t *testing.T) {
	t.Parallel()

	t.Run("creates_nested_directory_and_database", func(t *testing.T) {
		t.Parallel()

		dir := filepath.Join(t.TempDir(), "a", "b")

		store, closeStore, err := createCacheStore(dir)
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
		t.Cleanup(func() { _ = closeStore() })

		if store == nil {
			t.Fatal("store: got nil, want non-nil")
		}

		want := filepath.Join(dir, "cache.db")
		if got := store.DB.Path(); got != want {
			t.Fatalf("store path: got %q, want %q", got, want)
		}

		if _, err := os.Stat(want); err != nil {
			t.Fatalf("stat database file: got %v, want nil", err)
		}
	})

	t.Run("close_func_closes_database", func(t *testing.T) {
		t.Parallel()

		store, closeStore, err := createCacheStore(t.TempDir())
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		if err := closeStore(); err != nil {
			t.Fatalf("unexpected error closing store: %s", err)
		}

		err = store.DB.View(func(_ *bbolt.Tx) error { return nil })
		if !errors.Is(err, bbolterrs.ErrDatabaseNotOpen) {
			t.Fatalf("got %v, want %v", err, bbolterrs.ErrDatabaseNotOpen)
		}
	})

	t.Run("reopens_existing_database", func(t *testing.T) {
		t.Parallel()

		dir := t.TempDir()

		_, closeStore, err := createCacheStore(dir)
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		if err := closeStore(); err != nil {
			t.Fatalf("unexpected error closing store: %s", err)
		}

		_, closeStore, err = createCacheStore(dir)
		if err != nil {
			t.Fatalf("unexpected error reopening store: %s", err)
		}
		t.Cleanup(func() { _ = closeStore() })
	})

	t.Run("errors_when_directory_cannot_be_created", func(t *testing.T) {
		t.Parallel()

		store, closeStore, err := createCacheStore("\x00c")
		if err == nil {
			t.Cleanup(func() { _ = closeStore() })
			t.Fatal("got nil, want error")
		}

		if want := "failed to create cache directory"; !strings.Contains(err.Error(), want) {
			t.Fatalf("got %q, want it to contain %q", err.Error(), want)
		}

		if store != nil {
			t.Fatal("store: got non-nil, want nil")
		}

		if closeStore != nil {
			t.Fatal("close func: got non-nil, want nil")
		}
	})

	t.Run("errors_when_database_cannot_be_opened", func(t *testing.T) {
		t.Parallel()

		dir := t.TempDir()

		// A directory in place of the database file makes the open fail immediately, without waiting for the lock timeout.
		if err := os.Mkdir(filepath.Join(dir, "cache.db"), 0o700); err != nil {
			t.Fatalf("failed to create directory in place of database: %s", err)
		}

		store, closeStore, err := createCacheStore(dir)
		if err == nil {
			t.Cleanup(func() { _ = closeStore() })
			t.Fatal("got nil, want error")
		}

		if want := "failed to open cache storage"; !strings.Contains(err.Error(), want) {
			t.Fatalf("got %q, want it to contain %q", err.Error(), want)
		}

		if errors.Is(err, ErrCacheInUse) {
			t.Fatalf("got %v, want not %v", err, ErrCacheInUse)
		}

		if store != nil {
			t.Fatal("store: got non-nil, want nil")
		}

		if closeStore != nil {
			t.Fatal("close func: got non-nil, want nil")
		}
	})

	// Waits for the lock timeout, so takes about a second.
	t.Run("errors_when_cache_is_in_use", func(t *testing.T) {
		t.Parallel()

		dir := t.TempDir()

		_, closeStore, err := createCacheStore(dir)
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		store, closeSecond, err := createCacheStore(dir)
		if err == nil {
			t.Cleanup(func() { _ = closeSecond() })
			t.Fatal("got nil, want error")
		}

		if !errors.Is(err, ErrCacheInUse) {
			t.Fatalf("got %v, want %v", err, ErrCacheInUse)
		}

		if want := strconv.Quote(filepath.Join(dir, "cache.db")); !strings.Contains(err.Error(), want) {
			t.Fatalf("got %q, want it to contain %q", err.Error(), want)
		}

		if store != nil {
			t.Fatalf("store: got %v, want nil", store)
		}

		if closeSecond != nil {
			t.Fatal("close func: got non-nil, want nil")
		}

		if err := closeStore(); err != nil {
			t.Fatalf("unexpected error closing store: %s", err)
		}

		_, closeStore, err = createCacheStore(dir)
		if err != nil {
			t.Fatalf("unexpected error reopening store after close: %s", err)
		}
		t.Cleanup(func() { _ = closeStore() })
	})
}

func Test_bestEffortStorage(t *testing.T) {
	t.Parallel()

	t.Run("Get", func(t *testing.T) {
		t.Parallel()

		getErr := errors.New("get failed")

		for _, tt := range []struct {
			name       string
			body       *closeRecorder
			innerErr   error
			wantResp   bool
			wantClosed bool
		}{
			{
				name:     "returns_inner_response",
				body:     &closeRecorder{Reader: strings.NewReader("cached")},
				wantResp: true,
			},
			{
				name: "returns_miss_on_inner_miss",
			},
			{
				name:     "returns_miss_on_inner_error",
				innerErr: getErr,
			},
			{
				name:       "closes_inner_response_on_inner_error",
				body:       &closeRecorder{Reader: strings.NewReader("cached")},
				innerErr:   getErr,
				wantClosed: true,
			},
		} {
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				var innerResp *http.Response
				if tt.body != nil {
					innerResp = &http.Response{StatusCode: http.StatusOK, Body: tt.body}
				}

				s := bestEffortStorage{inner: &testStorage{getResp: innerResp, getErr: tt.innerErr}}

				got, err := s.Get(t.Context(), mustCreateRequest(t, http.MethodGet, "https://api.github.com/repos/o/r"))
				if err != nil {
					t.Fatalf("got %v, want nil", err)
				}

				var want *http.Response
				if tt.wantResp {
					want = innerResp
				}

				if got != want {
					t.Fatalf("response: got %v, want %v", got, want)
				}

				if tt.body != nil {
					if got := tt.body.closed.Load(); got != tt.wantClosed {
						t.Fatalf("inner body closed: got %t, want %t", got, tt.wantClosed)
					}
				}
			})
		}
	})

	t.Run("Put", func(t *testing.T) {
		t.Parallel()

		readErr := errors.New("read failed")

		for _, tt := range []struct {
			name          string
			body          *closeRecorder
			putErr        error
			wantErr       error
			wantPutBodies []string
			wantBody      string
		}{
			{
				name:          "stores_and_restores_body",
				body:          &closeRecorder{Reader: strings.NewReader("payload")},
				wantPutBodies: []string{"payload"},
				wantBody:      "payload",
			},
			{
				name:          "restores_body_on_inner_error",
				body:          &closeRecorder{Reader: strings.NewReader("payload")},
				putErr:        errors.New("put failed"),
				wantPutBodies: []string{"payload"},
				wantBody:      "payload",
			},
			{
				name:          "stores_and_restores_nil_body",
				wantPutBodies: []string{""},
				wantBody:      "",
			},
			{
				name:    "errors_and_closes_body_on_read_error",
				body:    &closeRecorder{Reader: io.MultiReader(strings.NewReader("pay"), iotest.ErrReader(readErr))},
				wantErr: readErr,
			},
		} {
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				resp := &http.Response{StatusCode: http.StatusOK, Request: mustCreateRequest(t, http.MethodGet, "https://api.github.com/repos/o/r")}
				if tt.body != nil {
					resp.Body = tt.body
				}

				inner := &testStorage{putErr: tt.putErr}

				err := bestEffortStorage{inner: inner}.Put(t.Context(), resp)
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("got %v, want %v", err, tt.wantErr)
				}

				if tt.body != nil && !tt.body.closed.Load() {
					t.Error("original body closed: got false, want true")
				}

				if diff := cmp.Diff(tt.wantPutBodies, inner.putBodies); diff != "" {
					t.Fatalf("put bodies (-want +got):\n%s", diff)
				}

				if tt.wantErr != nil {
					return
				}

				got, err := io.ReadAll(resp.Body)
				if err != nil {
					t.Fatalf("read restored body: got %v, want nil", err)
				}

				if string(got) != tt.wantBody {
					t.Fatalf("restored body: got %q, want %q", got, tt.wantBody)
				}

				if want := int64(len(tt.wantBody)); resp.ContentLength != want {
					t.Fatalf("content length: got %d, want %d", resp.ContentLength, want)
				}
			})
		}
	})
}
