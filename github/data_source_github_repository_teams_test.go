package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubRepositoryTeamsDataSource(t *testing.T) {
	t.Run("queries teams of an existing repository", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		teamName := fmt.Sprintf("%steam-%s", testResourcePrefix, randomID)
		repoName := fmt.Sprintf("%srepo-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name      = "%s"
				auto_init = true
			}

			resource "github_team" "test" {
				name      = "%s"
			}

			resource "github_team_repository" "test" {
				team_id    = github_team.test.id
				repository = github_repository.test.name
				permission = "push"
			}
		`, repoName, teamName)

		config2 := config + `
			data "github_repository_teams" "test" {
				name = github_repository.test.name
			}
		`

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("data.github_repository_teams.test", "name", repoName),
			resource.TestCheckResourceAttr("data.github_repository_teams.test", "teams.#", "1"),
			resource.TestCheckResourceAttr("data.github_repository_teams.test", "teams.0.slug", teamName),
			resource.TestCheckResourceAttr("data.github_repository_teams.test", "teams.0.permission", "push"),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  resource.ComposeTestCheckFunc(),
				},
				{
					Config: config2,
					Check:  check,
				},
			},
		})
	})
}
