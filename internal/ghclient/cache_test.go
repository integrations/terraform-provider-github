package ghclient

import (
	"os"
	"path/filepath"
	"testing"
)

func Test_createCacheStore(t *testing.T) {
	t.Parallel()

	tempPath, err := os.MkdirTemp("", "ghclient-cache-store-*")
	if err != nil {
		t.Fatalf("failed to create temporary directory: %v", err)
	}
	basePath := filepath.Join(tempPath, "cache-root")
	ref := "feature-branch"
	invalidPath := filepath.Join(tempPath, "not-a-dir")
	if err := os.WriteFile(invalidPath, []byte("x"), 0o600); err != nil {
		t.Fatalf("failed to create invalid path fixture: %v", err)
	}

	for _, tt := range []struct {
		name        string
		path        *string
		ref         *string
		expectError bool
		verify      func(t *testing.T)
	}{
		{
			name: "with_path_and_ref",
			path: &basePath,
			ref:  &ref,
			verify: func(t *testing.T) {
				t.Helper()
				cacheDBPath := filepath.Join(basePath, "terraform-provider-github", ref, "cache.db")
				if _, err := os.Stat(cacheDBPath); err != nil {
					t.Fatalf("expected cache db to exist at %q: %v", cacheDBPath, err)
				}
			},
		},
		{
			name: "without_path",
		},
		{
			name:        "invalid_path",
			path:        &invalidPath,
			expectError: true,
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			store, err := createCacheStore(tt.path, tt.ref)
			if tt.expectError {
				if err == nil {
					t.Fatal("expected error")
				}
				return
			}

			if err != nil {
				t.Fatalf("expected store to be created, got error: %v", err)
			}

			if tt.verify != nil {
				tt.verify(t)
			}

			if closer, ok := store.(interface{ Close() error }); ok {
				t.Cleanup(func() {
					if err := closer.Close(); err != nil {
						t.Fatalf("failed to close store: %v", err)
					}
				})
			}
		})
	}
}
