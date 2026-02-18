package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
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
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_enterprise_teams.all", tfjsonpath.New("id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_enterprise_teams.all", tfjsonpath.New("teams"), knownvalue.ListPartial(map[int]knownvalue.Check{
							0: knownvalue.ObjectPartial(map[string]knownvalue.Check{
								"team_id": knownvalue.NotNull(),
								"slug":    knownvalue.NotNull(),
								"name":    knownvalue.NotNull(),
							}),
						})),
					},
				},
			},
		})
	})
}
