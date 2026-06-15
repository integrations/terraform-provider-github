package github

import (
	"regexp"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-github/v88/github"
)

func Test_migrateRepositoryWithID(t *testing.T) {
	for _, tt := range []struct {
		testName   string
		statusCode int
		body       *github.Repository
		rawState   map[string]any
		want       map[string]any
		wantErr    *string
	}{
		{
			testName:   "succeeds_if_repo_found",
			statusCode: 200,
			body: &github.Repository{
				ID: new(int64(123456)),
			},
			rawState: map[string]any{
				"repository": "my-repo",
			},
			want: map[string]any{
				"repository":    "my-repo",
				"repository_id": 123456,
			},
		},
		{
			testName:   "succeeds_if_repo_id_wrong_type_in_state",
			statusCode: 200,
			body: &github.Repository{
				ID: new(int64(123456)),
			},
			rawState: map[string]any{
				"repository":    "my-repo",
				"repository_id": "not-an-int",
			},
			want: map[string]any{
				"repository":    "my-repo",
				"repository_id": 123456,
			},
		},
		{
			testName:   "succeeds_if_repo_id_already_in_state",
			statusCode: 200,
			body: &github.Repository{
				ID: new(int64(123456)),
			},
			rawState: map[string]any{
				"repository":    "my-repo",
				"repository_id": 123456,
			},
			want: map[string]any{
				"repository":    "my-repo",
				"repository_id": 123456,
			},
		},
		{
			testName:   "fails_if_repo_not_found",
			statusCode: 404,
			body:       nil,
			rawState: map[string]any{
				"repository": "my-repo",
			},
			wantErr: new("failed to retrieve repository"),
		},
		{
			testName:   "fails_if_repo_not_in_state",
			statusCode: 404,
			body:       nil,
			rawState: map[string]any{
				"repo": "my-repo",
			},
			wantErr: new("repository name not found in state"),
		},
		{
			testName:   "fails_if_repo_wrong_type_in_state",
			statusCode: 404,
			body:       nil,
			rawState: map[string]any{
				"repository": 123,
			},
			wantErr: new("repository name is not a string"),
		},
	} {
		t.Run(tt.testName, func(t *testing.T) {
			ts := githubApiMock([]*mockResponse{mustGetTestMockResponse(t, "/repos/my-org/my-repo", tt.statusCode, tt.body)})
			defer ts.Close()

			got, err := migrateRepositoryWithID(t.Context(), mustCreateTestGitHubClient(t, ts.URL), "my-org", tt.rawState)
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
				t.Fatalf("got %+v, want %+v: %s", got, tt.want, diff)
			}
		})
	}
}
