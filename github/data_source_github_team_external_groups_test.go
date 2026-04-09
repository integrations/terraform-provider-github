package github

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubTeamExternalGroupsDataSource(t *testing.T) {

	t.Run("errors when querying a non-existing team", func(t *testing.T) {
		config := `
			data "github_team_external_groups" "test" {
				slug = "non-existing-team-slug"
			}
		`

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      config,
					ExpectError: regexp.MustCompile(`Not Found`),
				},
			},
		})
	})

	t.Run("returns empty list for team without external groups", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		teamName := fmt.Sprintf("%steam-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`
			resource "github_team" "test" {
				name = "%s"
			}

			data "github_team_external_groups" "test" {
				slug = github_team.test.slug
			}
		`, teamName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_team_external_groups.test", tfjsonpath.New("external_groups"), knownvalue.ListSizeExact(0)),
					},
				},
			},
		})
	})
}
