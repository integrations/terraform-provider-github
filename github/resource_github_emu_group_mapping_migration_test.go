package github

import (
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-github/v89/github"
)

func Test_resourceGithubEMUGroupMappingStateUpgradeV1(t *testing.T) {
	t.Parallel()

	for _, d := range []struct {
		testName    string
		rawState    map[string]any
		want        map[string]any
		shouldError bool
	}{
		{
			testName: "migrates v1 to v2",
			rawState: map[string]any{
				"id": "123456:test-team:7765543",
			},
			want: map[string]any{
				"id": "7765543:123456",
			},
			shouldError: false,
		},
	} {
		t.Run(d.testName, func(t *testing.T) {
			t.Parallel()

			got, err := resourceGithubEMUGroupMappingStateUpgradeV1(t.Context(), d.rawState, nil)
			if (err != nil) != d.shouldError {
				t.Fatalf("unexpected error state")
			}

			if diff := cmp.Diff(got, d.want); !d.shouldError && diff != "" {
				t.Fatalf("got %+v, want %+v: %s", got, d.want, diff)
			}
		})
	}
}

type v0StatusCode struct {
	externalGroups int
	team           int
}
type v0Body struct {
	externalGroups *github.ExternalGroupList
	team           *github.Team
}

func Test_resourceGithubEMUGroupMappingStateUpgradeV0(t *testing.T) {
	t.Parallel()

	const testOrgSlug = "test-org"
	const testTeamSlug = "test-team"
	const testTeamID = 432574718
	const testGroupID = 1234567890

	meta := &Owner{
		name: testOrgSlug,
	}

	for _, tt := range []struct {
		testName   string
		rawState   map[string]any
		want       map[string]any
		wantErr    *string
		statusCode v0StatusCode
		body       v0Body
	}{
		{
			testName: "migrates v0 to v1",
			rawState: map[string]any{
				"id":        fmt.Sprintf("teams/%s/%d/external-groups", testTeamSlug, testGroupID),
				"team_slug": testTeamSlug,
				"group_id":  testGroupID,
			},
			want: map[string]any{
				"id":        fmt.Sprintf("%d:%s:%d", testTeamID, testTeamSlug, testGroupID),
				"team_slug": testTeamSlug,
				"team_id":   testTeamID,
				"group_id":  testGroupID,
			},
			wantErr: nil,
			statusCode: v0StatusCode{
				externalGroups: 201,
				team:           200,
			},
			body: v0Body{
				externalGroups: &github.ExternalGroupList{
					Groups: []*github.ExternalGroup{{
						GroupID:   new(int64(testGroupID)),
						GroupName: new(testOrgSlug),
						UpdatedAt: new(github.Timestamp{Time: time.Now()}),
					}},
				},
				team: &github.Team{
					ID: new(int64(testTeamID)),
				},
			},
		},
	} {
		t.Run(tt.testName, func(t *testing.T) {
			mockResponses := []*mockResponse{
				mustGetTestMockResponse(t, fmt.Sprintf("/orgs/%s/teams/%s/external-groups", testOrgSlug, testTeamSlug), tt.statusCode.externalGroups, tt.body.externalGroups),
				mustGetTestMockResponse(t, fmt.Sprintf("/orgs/%s/teams/%s", testOrgSlug, testTeamSlug), tt.statusCode.team, tt.body.team),
			}
			ts := githubApiMock(mockResponses)
			defer ts.Close()
			meta.v3client = mustCreateTestGitHubClient(t, ts.URL)

			got, err := resourceGithubEMUGroupMappingStateUpgradeV0(t.Context(), tt.rawState, meta)
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
