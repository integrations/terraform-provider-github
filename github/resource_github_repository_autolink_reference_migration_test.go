package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-github/v88/github"
)

func buildMockResponsesForRepositoryAutolinkReferenceMigrationV0toV1(mockOwner, mockRepo string, wantRepoID int) []*mockResponse {
	responseBodyJSON, err := json.Marshal(github.Repository{
		ID:   new(int64(wantRepoID)),
		Name: new(mockRepo),
	})
	if err != nil {
		panic(fmt.Sprintf("error marshalling repository response: %s", err))
	}

	return []*mockResponse{{
		ExpectedUri: fmt.Sprintf("/repos/%s/%s", mockOwner, mockRepo),
		ExpectedHeaders: map[string]string{
			"Accept": "application/vnd.github.scarlet-witch-preview+json, application/vnd.github.mercy-preview+json, application/vnd.github.baptiste-preview+json, application/vnd.github.nebula-preview+json",
		},
		ResponseBody: string(responseBodyJSON),
		StatusCode:   http.StatusOK,
	}}
}

func Test_resourceGithubRepositoryAutolinkReferenceStateUpgradeV0toV1(t *testing.T) {
	t.Parallel()

	for _, d := range []struct {
		testName    string
		rawState    map[string]any
		want        map[string]any
		shouldError bool
	}{
		{
			testName: "adds_repository_id",
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
		t.Run(d.testName, func(t *testing.T) {
			t.Parallel()

			meta := &Owner{name: "test-org"}
			ts := githubApiMock(buildMockResponsesForRepositoryAutolinkReferenceMigrationV0toV1(meta.name, "test-repo", 1234567890))
			defer ts.Close()

			meta.v3client = mustGitHubClient(t, ts.URL)

			got, err := resourceGithubRepositoryAutolinkReferenceStateUpgradeV0(t.Context(), d.rawState, meta)
			if (err != nil) != d.shouldError {
				t.Fatalf("unexpected error state: got error %v, shouldError %v", err, d.shouldError)
			}

			if diff := cmp.Diff(got, d.want); diff != "" && !d.shouldError {
				t.Fatalf("got %+v, want %+v, diff %s", got, d.want, diff)
			}
		})
	}
}
