package github

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-github/v82/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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

func buildMockResponsesForMigrationV0toV1() []*mockResponse {
	return []*mockResponse{
		{
			ExpectedUri: fmt.Sprintf("/orgs/%s/teams/%s/external-groups", "test-org", "test-team"),
			ExpectedHeaders: map[string]string{
				"Accept": "application/vnd.github.v3+json",
			},
			ResponseBody: fmt.Sprintf(`
{
	"groups": [
		{
			"group_id": %d,
			"group_name": "test-group",
			"updated_at": "2021-01-24T11:31:04-06:00"
		}
	]
}`, int64(testGroupID)),
			StatusCode: 201,
		},
		{
			ExpectedUri: fmt.Sprintf("/orgs/%s/teams/%s", "test-org", "test-team"),
			ExpectedHeaders: map[string]string{
				"Accept": "application/vnd.github.v3+json",
			},
			ResponseBody: fmt.Sprintf(`
{
	"id": %d
}
`, testTeamID),
			StatusCode: 200,
		},
	}
}

func TestGithub_MigrateEMUGroupMappingsState(t *testing.T) {
	t.Parallel()

	meta := &Owner{
		name: "test-org",
	}

	for _, d := range []struct {
		testName           string
		migrationFunc      schema.StateUpgradeFunc
		rawState           currentStateFunc
		want               expectedStateFunc
		buildMockResponses func() []*mockResponse
		shouldError        bool
	}{
		{
			testName:           "migrates v0 to v1",
			migrationFunc:      resourceGithubEMUGroupMappingStateUpgradeV0,
			rawState:           testResourceGithubEMUGroupMappingInstanceStateDataV0,
			want:               testResourceGithubEMUGroupMappingInstanceStateDataV1,
			buildMockResponses: buildMockResponsesForMigrationV0toV1,
			shouldError:        false,
		},
	} {
		t.Run(d.testName, func(t *testing.T) {
			t.Parallel()

			ts := githubApiMock(d.buildMockResponses())
			defer ts.Close()

			httpCl := http.DefaultClient
			httpCl.Transport = http.DefaultTransport

			client := github.NewClient(httpCl)
			u, _ := url.Parse(ts.URL + "/")
			client.BaseURL = u
			meta.v3client = client

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
