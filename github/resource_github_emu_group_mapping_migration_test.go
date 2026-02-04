package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-github/v82/github"
)

func buildMockResponsesForMigrationV0toV1(mockResponsesOptions mockResponsesOptionsEMUGroupMappingMigrationV0V1) []*mockResponse {
	responseBodyJson, err := json.Marshal(mockResponsesOptions.ExternalGroupList)
	if err != nil {
		panic(fmt.Sprintf("error marshalling external groups response: %s", err))
	}

	mockTeamResponseJson, err := json.Marshal(mockResponsesOptions.Team)
	if err != nil {
		panic(fmt.Sprintf("error marshalling mock team response: %s", err))
	}
	return []*mockResponse{
		{
			ExpectedUri: fmt.Sprintf("/orgs/%s/teams/%s/external-groups", mockResponsesOptions.OrgSlug, mockResponsesOptions.TeamSlug),
			ExpectedHeaders: map[string]string{
				"Accept": "application/vnd.github.v3+json",
			},
			ResponseBody: string(responseBodyJson),
			StatusCode:   mockResponsesOptions.externalGroupsResponseStatusCode,
		},
		{
			ExpectedUri: fmt.Sprintf("/orgs/%s/teams/%s", mockResponsesOptions.OrgSlug, mockResponsesOptions.TeamSlug),
			ExpectedHeaders: map[string]string{
				"Accept": "application/vnd.github.v3+json",
			},
			ResponseBody: string(mockTeamResponseJson),
			StatusCode:   mockResponsesOptions.teamResponseStatusCode,
		},
	}
}

type mockResponsesOptionsEMUGroupMappingMigrationV0V1 struct {
	OrgSlug                          string
	TeamSlug                         string
	externalGroupsResponseStatusCode int
	teamResponseStatusCode           int
	ExternalGroupList                github.ExternalGroupList
	Team                             github.Team
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

	for _, d := range []struct {
		testName             string
		rawState             map[string]any
		want                 map[string]any
		shouldError          bool
		mockResponsesOptions mockResponsesOptionsEMUGroupMappingMigrationV0V1
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
				"team_id":   int64(testTeamID),
				"group_id":  testGroupID,
			},
			shouldError: false,
			mockResponsesOptions: mockResponsesOptionsEMUGroupMappingMigrationV0V1{
				OrgSlug:                          testOrgSlug,
				TeamSlug:                         testTeamSlug,
				externalGroupsResponseStatusCode: 201,
				teamResponseStatusCode:           200,
				ExternalGroupList: github.ExternalGroupList{
					Groups: []*github.ExternalGroup{{
						GroupID:   github.Ptr(int64(testGroupID)),
						GroupName: github.Ptr(testOrgSlug),
						UpdatedAt: github.Ptr(github.Timestamp{Time: time.Now()}),
					}},
				},
				Team: github.Team{
					ID: github.Ptr(int64(testTeamID)),
				},
			},
		},
	} {
		t.Run(d.testName, func(t *testing.T) {
			t.Parallel()

			ts := githubApiMock(buildMockResponsesForMigrationV0toV1(d.mockResponsesOptions))
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
