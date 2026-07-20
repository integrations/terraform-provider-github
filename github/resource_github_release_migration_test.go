package github

import (
	"regexp"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-github/v89/github"
)

func Test_resourceGithubReleaseStateUpgradeV0(t *testing.T) {
	// IMPORTANT: This test is not parallelized because it uses global state.

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
				"id":                       "987654",
				"repository":               "my-repo",
				"tag_name":                 "v1.0.0",
				"target_commitish":         "main",
				"name":                     "My Release",
				"body":                     "Release notes",
				"draft":                    false,
				"prerelease":               false,
				"generate_release_notes":   true,
				"discussion_category_name": "General",
				"etag":                     "",
				"release_id":               123456,
				"node_id":                  "MDc6UmVsZWFzZTEyMzQ1Ng==",
				"created_at":               "2023-01-01T00:00:00Z",
				"published_at":             "2023-01-02T00:00:00Z",
				"url":                      "https://api.github.com/repos/my-org/my-repo/releases/987654",
				"html_url":                 "https://github.com/my-org/my-repo/releases/987654",
				"assets_url":               "https://api.github.com/repos/my-org/my-repo/releases/987654/assets",
				"upload_url":               "https://uploads.github.com/repos/my-org/my-repo/releases/987654/assets{?name,label}",
				"tarball_url":              "https://api.github.com/repos/my-org/my-repo/tarball/v1.0.0",
				"zipball_url":              "https://api.github.com/repos/my-org/my-repo/zipball/v1.0.0",
			},
			want: map[string]any{
				"id":                       "987654",
				"repository":               "my-repo",
				"repository_id":            123456,
				"tag_name":                 "v1.0.0",
				"target_commitish":         "main",
				"name":                     "My Release",
				"body":                     "Release notes",
				"draft":                    false,
				"prerelease":               false,
				"generate_release_notes":   true,
				"discussion_category_name": "General",
				"etag":                     "",
				"release_id":               123456,
				"node_id":                  "MDc6UmVsZWFzZTEyMzQ1Ng==",
				"created_at":               "2023-01-01T00:00:00Z",
				"published_at":             "2023-01-02T00:00:00Z",
				"url":                      "https://api.github.com/repos/my-org/my-repo/releases/987654",
				"html_url":                 "https://github.com/my-org/my-repo/releases/987654",
				"assets_url":               "https://api.github.com/repos/my-org/my-repo/releases/987654/assets",
				"upload_url":               "https://uploads.github.com/repos/my-org/my-repo/releases/987654/assets{?name,label}",
				"tarball_url":              "https://api.github.com/repos/my-org/my-repo/tarball/v1.0.0",
				"zipball_url":              "https://api.github.com/repos/my-org/my-repo/zipball/v1.0.0",
			},
		},
		{
			testName:   "fails_if_repo_not_found",
			statusCode: 404,
			body:       nil,
			rawState: map[string]any{
				"id":                       "987654",
				"repository":               "my-repo",
				"tag_name":                 "v1.0.0",
				"target_commitish":         "main",
				"name":                     "My Release",
				"body":                     "Release notes",
				"draft":                    false,
				"prerelease":               false,
				"generate_release_notes":   true,
				"discussion_category_name": "General",
				"etag":                     "",
				"release_id":               123456,
				"node_id":                  "MDc6UmVsZWFzZTEyMzQ1Ng==",
				"created_at":               "2023-01-01T00:00:00Z",
				"published_at":             "2023-01-02T00:00:00Z",
				"url":                      "https://api.github.com/repos/my-org/my-repo/releases/987654",
				"html_url":                 "https://github.com/my-org/my-repo/releases/987654",
				"assets_url":               "https://api.github.com/repos/my-org/my-repo/releases/987654/assets",
				"upload_url":               "https://uploads.github.com/repos/my-org/my-repo/releases/987654/assets{?name,label}",
				"tarball_url":              "https://api.github.com/repos/my-org/my-repo/tarball/v1.0.0",
				"zipball_url":              "https://api.github.com/repos/my-org/my-repo/zipball/v1.0.0",
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

			got, err := resourceGithubReleaseStateUpgradeV0(t.Context(), tt.rawState, meta)
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
