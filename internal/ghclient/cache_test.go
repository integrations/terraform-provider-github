package ghclient

import (
	"path/filepath"
	"regexp"
	"testing"
)

func Test_createCacheStore(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name      string
		path      string
		ref       string
		wantError string
	}{
		{
			name: "with_path",
			path: t.TempDir(),
		},
		{
			name: "with_path_and_ref",
			path: t.TempDir(),
			ref:  "test-ref",
		},
		{
			name:      "errors_without_path",
			wantError: "cache path cannot be empty",
		},
		{
			name:      "errors_with_invalid_path",
			path:      "\x00c",
			wantError: "failed to create cache directory",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			store, err := createCacheStore(tt.path, tt.ref)

			if tt.wantError != "" {
				if err == nil {
					t.Fatalf("expected error %q, got nil", tt.wantError)
				}
				if !regexp.MustCompile(regexp.QuoteMeta(tt.wantError)).MatchString(err.Error()) {
					t.Fatalf("expected error %q, got %q", tt.wantError, err.Error())
				}
				return
			}

			if err != nil {
				t.Fatalf("expected store to be created, got error: %v", err)
			}

			if store == nil {
				t.Fatal("expected store to be non-nil")
			}
			defer func() {
				_ = store.DB.Close()
			}()

			wantPath := filepath.Join(tt.path, "terraform-provider-github", tt.ref, "cache.db")
			if store.DB.Path() != wantPath {
				t.Fatalf("expected store path to be %q, got %q", wantPath, store.DB.Path())
			}
		})
	}
}
