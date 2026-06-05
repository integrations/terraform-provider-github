package github

import (
	"regexp"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-github/v88/github"
)

func Test_resourceGithubCustomPropertyStateUpgradeV0(t *testing.T) {
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
				"id":             "my-org:my-repo:my-property",
				"repository":     "my-repo",
				"property_name":  "my-property",
				"property_value": "my-value",
			},
			want: map[string]any{
				"id":             "my-org:my-repo:my-property",
				"repository":     "my-repo",
				"repository_id":  123456,
				"property_name":  "my-property",
				"property_value": "my-value",
			},
		},
		{
			testName:   "fails_if_repo_not_found",
			statusCode: 404,
			body:       nil,
			rawState: map[string]any{
				"id":             "my-org:my-repo:my-property",
				"repository":     "my-repo",
				"property_name":  "my-property",
				"property_value": "my-value",
			},
			wantErr: new("failed to retrieve repository"),
		},
	} {
		t.Run(tt.testName, func(t *testing.T) {
			ts := githubApiMock([]*mockResponse{mustGetTestMockResponse(t, "/repos/my-org/my-repo", tt.statusCode, tt.body)})
			defer ts.Close()

			meta := &Owner{
				name:     "my-org",
				v3client: mustCreateTestGitHubClient(t, ts.URL),
			}

			got, err := resourceGithubRepositoryCustomPropertyStateUpgradeV0(t.Context(), tt.rawState, meta)
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
