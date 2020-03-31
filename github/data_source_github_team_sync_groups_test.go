package github

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccGithubTeamSyncGroupsDataSource_noMatchReturnsError(t *testing.T) {
	if isEnterprise != "true" {
		t.Skip()
	}
	teamID := "67890"
	teamSlug := "non-existing"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckGithubTeamSyncGroupsDataSourceByIDConfig(teamID),
				ExpectError: regexp.MustCompile(`Could not find team`),
			},
			{
				Config:      testAccCheckGithubTeamSyncGroupsDataSourceByOtherConfig(teamSlug),
				ExpectError: regexp.MustCompile(`Could not find team`),
			},
		},
	})
}

func testAccCheckGithubTeamSyncGroupsDataSourceByIDConfig(teamID string) string {
	return fmt.Sprintf(`
data "github_team_sync_groups" "test" {
	retrieve_by = "id"
	team_id = "%s"
}
`, teamID)
}

func testAccCheckGithubTeamSyncGroupsDataSourceByOtherConfig(teamSlug string) string {
	return fmt.Sprintf(`
data "github_team_sync_groups" "test" {
	retrieve_by = "slug"
	team_slug = "%s"
}
`, teamSlug)
}
