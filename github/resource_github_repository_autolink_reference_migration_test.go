package github

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-github/v89/github"
)

func Test_resourceGithubRepositoryAutolinkReferenceStateUpgradeV1toV2(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		testName    string
		statusCode  int
		body        *github.Repository
		rawState    map[string]any
		want        map[string]any
		shouldError bool
	}{
		{
			testName:   "adds_repository_id",
			statusCode: 200,
			body: &github.Repository{
				ID:   new(int64(1234567890)),
				Name: new("test-repo"),
			},
			rawState: map[string]any{
				"id":                  "12345",
				"repository":          "test-repo",
				"key_prefix":          "TEST-",
				"target_url_template": "https://example.com/TEST-<num>",
				"is_alphanumeric":     true,
			},
			want: map[string]any{
				"id":                  "12345",
				"repository":          "test-repo",
				"repository_id":       1234567890,
				"key_prefix":          "TEST-",
				"target_url_template": "https://example.com/TEST-<num>",
				"is_alphanumeric":     true,
			},
			shouldError: false,
		},
	} {
		t.Run(tt.testName, func(t *testing.T) {
			t.Parallel()

			meta := &Owner{name: "test-org"}
			ts := githubApiMock([]*mockResponse{mustGetTestMockResponse(t, "/repos/"+meta.name+"/"+tt.body.GetName(), tt.statusCode, tt.body)})
			defer ts.Close()

			meta.v3client = mustCreateTestGitHubClient(t, ts.URL)

			got, err := resourceGithubRepositoryAutolinkReferenceStateUpgradeV1(t.Context(), tt.rawState, meta)
			if (err != nil) != tt.shouldError {
				t.Fatalf("unexpected error state: got error %v, shouldError %v", err, tt.shouldError)
			}

			if diff := cmp.Diff(got, tt.want); diff != "" && !tt.shouldError {
				t.Fatalf("got %+v, want %+v, diff %s", got, tt.want, diff)
			}
		})
	}
}
