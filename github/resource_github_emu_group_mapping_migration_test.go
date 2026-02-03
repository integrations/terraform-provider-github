package github

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-github/v82/github"
)

var (
	testTeamID  = 432574718
	testGroupID = 1234567890
)

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

func Test_resourceGithubEMUGroupMappingStateUpgradeV0(t *testing.T) {
	t.Parallel()

	meta := &Owner{
		name: "test-org",
	}

	for _, d := range []struct {
		testName           string
		rawState           map[string]any
		want               map[string]any
		buildMockResponses func() []*mockResponse
		shouldError        bool
	}{
		{
			testName: "migrates v0 to v1",
			rawState: map[string]any{
				"id":        "teams/test-team/external-groups",
				"team_slug": "test-team",
				"group_id":  testGroupID,
			},
			want: map[string]any{
				"id":        "432574718:test-team:1234567890",
				"team_slug": "test-team",
				"team_id":   int64(testTeamID),
				"group_id":  testGroupID,
			},
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

			currentState := d.rawState
			got, err := resourceGithubEMUGroupMappingStateUpgradeV0(t.Context(), currentState, meta)
			expectedState := d.want
			didError := err != nil
			if d.shouldError && !didError {
				t.Fatalf("state upgrade should have returned an error. Instead got: %#v", got)
			}
			if !d.shouldError && didError {
				t.Fatalf("state upgrade should not have returned an error. Instead got: %s", err.Error())
			}
			if diff := cmp.Diff(expectedState, got); diff != "" {
				t.Fatalf("state upgrade returned unexpected state. Diff: %s", diff)
			}
		})
	}
}
