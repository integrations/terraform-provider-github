package github

// TODO: Enable this test once we have a way to mock the GitHub API

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-github/v88/github"
)

func buildMockResponsesForRepositoryFileMigrationV0toV1(mockOwner, mockRepo string, wantRepoID int) []*mockResponse {
	responseBodyJson, err := json.Marshal(github.Repository{
		ID:            new(int64(wantRepoID)),
		DefaultBranch: new("main"),
		Name:          new(mockRepo),
	})
	if err != nil {
		panic(fmt.Sprintf("error marshalling repository response: %s", err))
	}
	return []*mockResponse{{
		ExpectedUri: fmt.Sprintf("/repos/%s/%s", mockOwner, mockRepo),
		ExpectedHeaders: map[string]string{
			"Accept": "application/vnd.github.scarlet-witch-preview+json, application/vnd.github.mercy-preview+json, application/vnd.github.baptiste-preview+json, application/vnd.github.nebula-preview+json",
		},
		ResponseBody: string(responseBodyJson),
		StatusCode:   http.StatusOK,
	}}
}

func Test_resourceGithubRepositoryFileStateUpgradeV0toV1(t *testing.T) {
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
				"repository_id":       1234567890,
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
				"id":            "test-repo:README.md:develop",
				"repository":    "test-repo",
				"repository_id": 1234567890,
				"file":          "README.md",
				"content":       "# README",
				"branch":        "develop",
			},
			shouldError: false,
		},
		{
			testName: "migrates_with_missing_branch",
			rawState: map[string]any{
				"id":         "test-repo/path/to/file.txt",
				"repository": "test-repo",
				"file":       "path/to/file.txt",
				"content":    "file content",
			},
			want: map[string]any{
				"id":            "test-repo:path/to/file.txt:main",
				"repository":    "test-repo",
				"repository_id": 1234567890,
				"file":          "path/to/file.txt",
				"content":       "file content",
				"branch":        "main", // fetched from API
			},
			shouldError: false,
		},
		{
			testName: "migrates_with_empty_branch",
			rawState: map[string]any{
				"id":         "test-repo/path/to/file.txt",
				"repository": "test-repo",
				"file":       "path/to/file.txt",
				"content":    "file content",
				"branch":     "",
			},
			want: map[string]any{
				"id":            "test-repo:path/to/file.txt:main",
				"repository":    "test-repo",
				"repository_id": 1234567890,
				"file":          "path/to/file.txt",
				"content":       "file content",
				"branch":        "main", // fetched from API
			},
			shouldError: false,
		},
		{
			testName: "migrates_with_colon_in_file_path",
			rawState: map[string]any{
				"id":         "test-repo/path/to:file.txt",
				"repository": "test-repo",
				"file":       "path/to:file.txt",
				"content":    "file content",
				"branch":     "main",
			},
			want: map[string]any{
				"id":            "test-repo:path/to??file.txt:main",
				"repository":    "test-repo",
				"repository_id": 1234567890,
				"file":          "path/to:file.txt",
				"content":       "file content",
				"branch":        "main",
			},
			shouldError: false,
		},
	} {
		t.Run(d.testName, func(t *testing.T) {
			t.Parallel()

			meta := &Owner{
				name: "test-org",
			}

			wantRepositoryName := ""
			if v, ok := d.want["repository"]; ok {
				if s, ok := v.(string); ok {
					wantRepositoryName = s
				}
			}

			wantRepositoryID := -1
			if v, ok := d.want["repository_id"]; ok {
				if i, ok := v.(int); ok {
					wantRepositoryID = i
				}
			}

			ts := githubApiMock(buildMockResponsesForRepositoryFileMigrationV0toV1(meta.name, wantRepositoryName, wantRepositoryID))
			defer ts.Close()

			client := mustCreateTestGitHubClient(t, ts.URL)
			meta.v3client = client

			got, err := resourceGithubRepositoryFileStateUpgradeV0(t.Context(), d.rawState, meta)
			if (err != nil) != d.shouldError {
				t.Fatalf("unexpected error state: got error %v, shouldError %v", err, d.shouldError)
			}

			if diff := cmp.Diff(got, d.want); diff != "" && !d.shouldError {
				t.Fatalf("got %+v, want %+v, diff %s", got, d.want, diff)
			}
		})
	}
}
