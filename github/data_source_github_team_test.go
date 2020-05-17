package github

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccGithubTeamDataSource_noMatchReturnsError(t *testing.T) {
	if err := testAccCheckOrganization(); err != nil {
		t.Skipf("Skipping because %s.", err.Error())
	}

	slug := "non-existing"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckGithubTeamDataSourceConfig(slug),
				ExpectError: regexp.MustCompile(`Could not find team`),
			},
		},
	})
}

func testAccCheckGithubTeamDataSourceConfig(slug string) string {
	return fmt.Sprintf(`
data "github_team" "test" {
  slug = "%s"
}
`, slug)
}
