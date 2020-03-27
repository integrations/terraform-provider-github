package github

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccGithubTeamSyncGroupsDataSource_noMatchReturnsError(t *testing.T) {
	slug := "non-existing"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckGithubTeamSyncGroupsDataSourceConfig(slug),
				ExpectError: regexp.MustCompile(`Could not find team`),
			},
		},
	})
}

func testAccCheckGithubTeamSyncGroupsDataSourceConfig(slug string) string {
	return fmt.Sprintf(`
data "github_team_synv_groups" "test" {
  slug = "%s"
}
`, slug)
}
