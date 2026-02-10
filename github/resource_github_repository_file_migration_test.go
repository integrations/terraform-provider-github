package github

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-github/v82/github"
)

func buildMockResponsesForRepositoryFileMigrationV0toV1(mockOwner, mockRepo string, wantRepoID int) []*mockResponse {
	responseBodyJson, err := json.Marshal(github.Repository{
		ID:            github.Ptr(int64(wantRepoID)),
		DefaultBranch: github.Ptr("main"),
		Name:          github.Ptr(mockRepo),
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
	} {
		t.Run(d.testName, func(t *testing.T) {
			t.Parallel()

			meta := &Owner{
				name: "test-org",
			}

			ts := githubApiMock(buildMockResponsesForRepositoryFileMigrationV0toV1(meta.name, d.want["repository"].(string), d.want["repository_id"].(int)))
			defer ts.Close()

			httpCl := http.DefaultClient
			httpCl.Transport = http.DefaultTransport

			client := github.NewClient(httpCl)
			u, _ := url.Parse(ts.URL + "/")
			client.BaseURL = u
			meta.v3client = client

			got, err := resourceGithubRepositoryFileStateUpgradeV0(context.Background(), d.rawState, meta)
			if (err != nil) != d.shouldError {
				t.Fatalf("unexpected error state: got error %v, shouldError %v", err, d.shouldError)
			}

			if diff := cmp.Diff(got, d.want); diff != "" && !d.shouldError {
				t.Fatalf("got %+v, want %+v, diff %s", got, d.want, diff)
			}
		})
	}
}
