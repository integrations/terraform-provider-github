package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubOrganizationTeamsDataSource(t *testing.T) {
	t.Run("queries", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		teamName := fmt.Sprintf("%steam-0-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
			resource "github_team" "test" {
				name = "%s"
			}

			data "github_organization_teams" "all" {
				depends_on = [github_team.test]
			}
		`, teamName)

		check := resource.ComposeAggregateTestCheckFunc(
			resource.TestCheckResourceAttrSet("data.github_organization_teams.all", "teams.#"),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  check,
				},
			},
		})
	})

	t.Run("queries results_per_page", func(t *testing.T) {
		config := `
		data "github_organization_teams" "all" {
			results_per_page = 50
		}
		`

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttrSet("data.github_organization_teams.all", "teams.#"),
					),
				},
			},
		})
	})
}
