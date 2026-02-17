package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccGithubEnterpriseTeamsDataSource(t *testing.T) {
	t.Run("lists all enterprise teams without error", func(t *testing.T) {
		randomID := acctest.RandString(5)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessMode(t, enterprise) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(`
						data "github_enterprise" "enterprise" {
							slug = "%s"
						}

						resource "github_enterprise_team" "test" {
							enterprise_slug = data.github_enterprise.enterprise.slug
							name            = "%s%s"
						}

						data "github_enterprise_teams" "all" {
							enterprise_slug = data.github_enterprise.enterprise.slug
							depends_on      = [github_enterprise_team.test]
						}
					`, testAccConf.enterpriseSlug, testResourcePrefix, randomID),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttrSet("data.github_enterprise_teams.all", "id"),
						resource.TestCheckResourceAttrSet("data.github_enterprise_teams.all", "teams.0.team_id"),
						resource.TestCheckResourceAttrSet("data.github_enterprise_teams.all", "teams.0.slug"),
						resource.TestCheckResourceAttrSet("data.github_enterprise_teams.all", "teams.0.name"),
					),
				},
			},
		})
	})
}
