package ghclient

import (
	"path/filepath"
	"regexp"
	"testing"
)

func Test_createCacheStore(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name    string
		opts    CacheOptions
		wantErr *string
	}{
		{
			name: "with_base_path_no_ref",
			opts: CacheOptions{
				Enabled:  true,
				BasePath: t.TempDir(),
			},
		},
		{
			name: "with_base_path_and_ref",
			opts: CacheOptions{
				Enabled:  true,
				BasePath: t.TempDir(),
				Ref:      "test",
			},
		},
		{
			name:    "errors_without_path",
			wantErr: new("cache path cannot be empty"),
		},
		{
			name: "errors_with_invalid_path",
			opts: CacheOptions{
				Enabled:  true,
				BasePath: "\x00c",
			},
			wantErr: new("failed to create cache directory"),
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			store, err := createCacheStore(tt.opts)
			if err != nil {
				if tt.wantErr == nil {
					t.Fatalf("expected no error, got %v", err)
				}

				if !regexp.MustCompile(regexp.QuoteMeta(*tt.wantErr)).MatchString(err.Error()) {
					t.Fatalf("expected error %q, got %q", *tt.wantErr, err.Error())
				}

				return
			}

			if tt.wantErr != nil {
				t.Fatalf("expected error %q, got nil", *tt.wantErr)
			}

			if store == nil {
				t.Fatal("expected store to be non-nil")
			}
			defer func() {
				_ = store.DB.Close()
			}()

			wantPath := filepath.Join(tt.opts.BasePath, tt.opts.Ref, "cache.db")
			if store.DB.Path() != wantPath {
				t.Fatalf("expected store path to be %q, got %q", wantPath, store.DB.Path())
			}
		})
	}
}
