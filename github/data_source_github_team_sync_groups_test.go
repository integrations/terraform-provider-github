package github

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccGithubTeamSyncGroupsDataSource_noMatchReturnsError(t *testing.T) {
	teamSlug := "non-existing"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckGithubTeamSyncGroupsDataSourceConfig(teamSlug),
				ExpectError: regexp.MustCompile(`Could not find team`),
			},
		},
	})
}

func testAccCheckGithubTeamSyncGroupsDataSourceConfig(teamSlug string) string {
	return fmt.Sprintf(`
data "github_team_sync_groups" "test" {
	team_slug = "%s"
}
`, teamSlug)
}
