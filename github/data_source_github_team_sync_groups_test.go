package github

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccGithubTeamSyncGroupsDataSource_noMatchReturnsError(t *testing.T) {
	orgName, teamSlug := "ne", "non-existing"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckGithubTeamSyncGroupsDataSourceConfig(orgName, teamSlug),
				ExpectError: regexp.MustCompile(`Could not find team`),
			},
		},
	})
}

func testAccCheckGithubTeamSyncGroupsDataSourceConfig(orgName, teamSlug string) string {
	return fmt.Sprintf(`
data "github_team_sync_groups" "test" {
	org_name = "%s"
	team_slug = "%s"
}
`, orgName, teamSlug)
}
