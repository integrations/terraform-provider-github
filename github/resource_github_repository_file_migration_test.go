package github

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_resourceGithubRepositoryFileStateUpgradeV0(t *testing.T) {
	t.Parallel()

	for _, d := range []struct {
		testName    string
		rawState    map[string]any
		want        map[string]any
		shouldError bool
	}{
		{
			testName: "preserves_existing_branch",
			rawState: map[string]any{
				"id":                  "test-repo/path/to/file.txt",
				"repository":          "test-repo",
				"file":                "path/to/file.txt",
				"content":             "file content",
				"branch":              "main",
				"commit_sha":          "abc123",
				"sha":                 "def456",
				"overwrite_on_create": false,
			},
			want: map[string]any{
				"id":                  "test-repo:path/to/file.txt:main",
				"repository":          "test-repo",
				"file":                "path/to/file.txt",
				"content":             "file content",
				"branch":              "main",
				"commit_sha":          "abc123",
				"sha":                 "def456",
				"overwrite_on_create": false,
			},
			shouldError: false,
		},
		{
			testName: "preserves_custom_branch",
			rawState: map[string]any{
				"id":         "test-repo/README.md",
				"repository": "test-repo",
				"file":       "README.md",
				"content":    "# README",
				"branch":     "develop",
			},
			want: map[string]any{
				"id":         "test-repo:README.md:develop",
				"repository": "test-repo",
				"file":       "README.md",
				"content":    "# README",
				"branch":     "develop",
			},
			shouldError: false,
		},
		// TODO: Enable this test once we have a pattern to create a mock client for the test.
		// 		{
		// 			testName: "migrates_with_missing_branch",
		// 			rawState: map[string]any{
		// 				"id":         "test-repo/path/to/file.txt",
		// 				"repository": "test-repo",
		// 				"file":       "path/to/file.txt",
		// 				"content":    "file content",
		// 			},
		// 			want: map[string]any{
		// 				"id":         "test-repo:path/to/file.txt:main",
		// 				"repository": "test-repo",
		// 				"file":       "path/to/file.txt",
		// 				"content":    "file content",
		// 				"branch":     "main", // fetched from API
		// 			},
		// 			shouldError: false,
		// 		},
		// TODO: Enable this test once we have a pattern to create a mock client for the test.
		// 		{
		// 			testName: "migrates_with_empty_branch",
		// 			rawState: map[string]any{
		// 				"id":         "test-repo/path/to/file.txt",
		// 				"repository": "test-repo",
		// 				"file":       "path/to/file.txt",
		// 				"content":    "file content",
		// 				"branch":     "",
		// 			},
		// 			want: map[string]any{
		// 				"id":         "test-repo:path/to/file.txt:main",
		// 				"repository": "test-repo",
		// 				"file":       "path/to/file.txt",
		// 				"content":    "file content",
		// 				"branch":     "main", // fetched from API
		// 			},
		// 			shouldError: false,
		// 		},
	} {
		t.Run(d.testName, func(t *testing.T) {
			t.Parallel()

			got, err := resourceGithubRepositoryFileStateUpgradeV0(context.Background(), d.rawState, nil)
			if (err != nil) != d.shouldError {
				t.Fatalf("unexpected error state: got error %v, shouldError %v", err, d.shouldError)
			}

			if diff := cmp.Diff(got, d.want); diff != "" && !d.shouldError {
				t.Fatalf("got %+v, want %+v, diff %s", got, d.want, diff)
			}
		})
	}
}
