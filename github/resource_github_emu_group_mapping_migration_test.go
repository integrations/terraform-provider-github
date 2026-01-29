package github

import (
	"net/http"
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-github/v82/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/migueleliasweb/go-github-mock/src/mock"
)

type (
	currentStateFunc  func() map[string]any
	expectedStateFunc func(t *testing.T) map[string]any
)

var (
	testTeamID  = 432574718
	testGroupID = 1234567890
)

func testResourceGithubEMUGroupMappingInstanceStateDataV0() map[string]any {
	return map[string]any{
		"id":        "teams/test-team/external-groups",
		"team_slug": "test-team",
		"group_id":  testGroupID,
	}
}

func testResourceGithubEMUGroupMappingInstanceStateDataV1(t *testing.T) map[string]any {
	v0 := testResourceGithubEMUGroupMappingInstanceStateDataV0()
	v0["team_id"] = int64(testTeamID)
	resourceID, err := buildID(strconv.Itoa(testTeamID), v0["team_slug"].(string), strconv.Itoa(v0["group_id"].(int)))
	if err != nil {
		t.Fatalf("error building resource ID: %s", err)
	}
	v0["id"] = resourceID
	return v0
}

func buildMockClientForMigrationV0toV1() *http.Client {
	return mock.NewMockedHTTPClient(
		mock.WithRequestMatch(
			mock.GetOrgsTeamsExternalGroupsByOrgByTeamSlug,
			github.ExternalGroupList{
				Groups: []*github.ExternalGroup{
					{
						GroupID: github.Ptr(int64(testGroupID)),
						Teams: []*github.ExternalGroupTeam{
							{
								TeamID: github.Ptr(int64(testTeamID)),
							},
						},
					},
				},
			},
		),
		mock.WithRequestMatch(
			mock.GetOrgsTeamsByOrgByTeamSlug,
			github.Team{
				ID: github.Ptr(int64(testTeamID)),
			},
		),
	)
}

func TestGithub_MigrateEMUGroupMappingsState(t *testing.T) {
	t.Parallel()

	meta := &Owner{
		name: "test-org",
	}

	for _, d := range []struct {
		testName      string
		migrationFunc schema.StateUpgradeFunc
		rawState      currentStateFunc
		want          expectedStateFunc
		buildClient   func() *http.Client
		shouldError   bool
	}{
		{
			testName:      "migrates v0 to v1",
			migrationFunc: resourceGithubEMUGroupMappingInstanceStateUpgradeV0,
			rawState:      testResourceGithubEMUGroupMappingInstanceStateDataV0,
			want:          testResourceGithubEMUGroupMappingInstanceStateDataV1,
			buildClient:   buildMockClientForMigrationV0toV1,
			shouldError:   false,
		},
	} {
		t.Run(d.testName, func(t *testing.T) {
			t.Parallel()

			ghClient := github.NewClient(d.buildClient())
			meta.v3client = ghClient

			currentState := d.rawState()
			got, err := d.migrationFunc(t.Context(), currentState, meta)
			expectedState := d.want(t)
			if (err != nil) != d.shouldError {
				t.Fatalf("unexpected error state: %s", err.Error())
			}
			if diff := cmp.Diff(expectedState, got); !d.shouldError && diff != "" {
				t.Fatal(diff)
			}
		})
	}
}
