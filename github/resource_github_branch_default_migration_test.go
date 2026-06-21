package github

import (
	"regexp"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-github/v88/github"
)

func Test_resourceGithubBranchDefaultStateUpgradeV0toV1(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		testName   string
		statusCode int
		body       *github.Repository
		rawState   map[string]any
		want       map[string]any
		wantErr    *string
	}{
		{
			testName:   "adds_repository_id",
			statusCode: 200,
			body: &github.Repository{
				ID:   new(int64(1234567890)),
				Name: new("test-repo"),
			},
			rawState: map[string]any{
				"id":         "test-repo",
				"repository": "test-repo",
				"branch":     "main",
				"rename":     false,
				"etag":       `W/\"etag\"`,
			},
			want: map[string]any{
				"id":            "test-repo",
				"repository":    "test-repo",
				"repository_id": 1234567890,
				"branch":        "main",
				"rename":        false,
				"etag":          `W/\"etag\"`,
			},
		},
		{
			testName:   "falls_back_to_id_for_imported_state",
			statusCode: 200,
			body: &github.Repository{
				ID:   new(int64(1234567890)),
				Name: new("test-repo"),
			},
			rawState: map[string]any{
				"id":     "test-repo",
				"branch": "main",
				"rename": true,
			},
			want: map[string]any{
				"id":            "test-repo",
				"repository":    "test-repo",
				"repository_id": 1234567890,
				"branch":        "main",
				"rename":        true,
			},
		},
		{
			testName:   "fails_if_repo_not_found",
			statusCode: 404,
			body:       nil,
			rawState: map[string]any{
				"repository": "test-repo",
			},
			wantErr: new("failed to retrieve repository"),
		},
		{
			testName:   "fails_if_repository_empty_and_id_missing",
			statusCode: 404,
			body:       nil,
			rawState: map[string]any{
				"branch": "main",
			},
			wantErr: new("state upgrade v0: repository is not a string or not set"),
		},
	} {
		t.Run(tt.testName, func(t *testing.T) {
			t.Parallel()

			ts := githubApiMock([]*mockResponse{mustGetTestMockResponse(t, "/repos/test-org/test-repo", tt.statusCode, tt.body)})
			defer ts.Close()
			meta := &Owner{name: "test-org", v3client: mustCreateTestGitHubClient(t, ts.URL)}

			got, err := resourceGithubBranchDefaultStateUpgradeV0(t.Context(), tt.rawState, meta)
			if err != nil {
				if tt.wantErr == nil {
					t.Fatalf("unexpected error: %s", err)
				}
				if !regexp.MustCompile(regexp.QuoteMeta(*tt.wantErr)).MatchString(err.Error()) {
					t.Fatalf("unexpected error: %s", err)
				}
				return
			}

			if tt.wantErr != nil {
				t.Fatalf("expected error: %s", *tt.wantErr)
			}

			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Fatalf("got %+v, want %+v, diff %s", got, tt.want, diff)
			}
		})
	}
}
