package github

import (
	"regexp"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-github/v88/github"
)

func Test_resourceGithubTeamMembersStateUpgradeV0(t *testing.T) {
	for _, tt := range []struct {
		testName     string
		ghURI        string
		ghStatusCode int
		ghBody       *github.Team
		rawState     map[string]any
		want         map[string]any
		wantErr      *string
	}{
		{
			testName:     "succeeds_if_team_found_by_id",
			ghURI:        "/organizations/1000/team/123456",
			ghStatusCode: 200,
			ghBody: &github.Team{
				ID:   new(int64(123456)),
				Slug: new("my-team"),
			},
			rawState: map[string]any{
				"id":      "123456",
				"team_id": "123456",
			},
			want: map[string]any{
				"id":        "123456",
				"team_id":   "123456",
				"team_slug": "my-team",
			},
		},
		{
			testName:     "succeeds_if_team_found_by_slug",
			ghURI:        "/orgs/my-org/teams/my-team",
			ghStatusCode: 200,
			ghBody: &github.Team{
				ID:   new(int64(123456)),
				Slug: new("my-team"),
			},
			rawState: map[string]any{
				"id":      "my-team",
				"team_id": "my-team",
			},
			want: map[string]any{
				"id":        "123456",
				"team_id":   "my-team",
				"team_slug": "my-team",
			},
		},
		{
			testName:     "errors_if_team_not_found_by_id",
			ghURI:        "/organizations/1000/team/123456",
			ghStatusCode: 404,
			rawState: map[string]any{
				"id":      "123456",
				"team_id": "123456",
			},
			wantErr: new("failed to lookup team slug for ID"),
		},
		{
			testName:     "errors_if_team_not_found_by_slug",
			ghURI:        "/orgs/my-org/teams/my-team",
			ghStatusCode: 404,
			rawState: map[string]any{
				"id":      "my-team",
				"team_id": "my-team",
			},
			wantErr: new("failed to lookup team ID for slug"),
		},
		{
			testName: "errors_if_id_not_in_state",
			rawState: map[string]any{
				"team_id": "123456",
			},
			wantErr: new("missing id in raw state"),
		},
		{
			testName: "errors_if_id_not_string",
			rawState: map[string]any{
				"id":      123456,
				"team_id": "123456",
			},
			wantErr: new("id in raw state is not a string"),
		},
	} {
		t.Run(tt.testName, func(t *testing.T) {
			ts := githubApiMock([]*mockResponse{mustGetTestMockResponse(t, tt.ghURI, tt.ghStatusCode, tt.ghBody)})
			defer ts.Close()

			meta := &Owner{
				name:     "my-org",
				id:       1000,
				v3client: mustCreateTestGitHubClient(t, ts.URL),
			}

			got, err := resourceGithubTeamMembersStateUpgradeV0(t.Context(), tt.rawState, meta)
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
