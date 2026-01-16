package github

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubTeamDataSource(t *testing.T) {
	t.Run("queries an existing team without error", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		teamName := fmt.Sprintf("%steam-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`
			resource "github_team" "test" {
				name = "%s"
			}

			data "github_team" "test" {
				slug = github_team.test.slug
			}
		`, teamName)

		check := resource.ComposeAggregateTestCheckFunc(
			resource.TestCheckResourceAttrSet("data.github_team.test", "name"),
			resource.TestCheckResourceAttrSet("data.github_team.test", "node_id"),
			resource.TestCheckResourceAttrSet("data.github_team.test", "privacy"),
			resource.TestCheckResourceAttrSet("data.github_team.test", "notification_setting"),
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

	t.Run("queries an existing team without error with immediate membership", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		teamName := fmt.Sprintf("%steam-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`
			resource "github_team" "test" {
				name = "%s"
			}

			data "github_team" "test" {
				slug            = github_team.test.slug
				membership_type = "immediate"
			}
		`, teamName)

		check := resource.ComposeAggregateTestCheckFunc(
			resource.TestCheckResourceAttrSet("data.github_team.test", "name"),
			resource.TestCheckResourceAttr("data.github_team.test", "name", teamName),
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

	t.Run("errors when querying a non-existing team", func(t *testing.T) {
		config := `
			data "github_team" "test" {
				slug = ""
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

	t.Run("queries an existing team without error in summary_only mode", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		teamName := fmt.Sprintf("%steam-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`
			resource "github_team" "test" {
				name = "%s"
			}

			data "github_team" "test" {
				slug = github_team.test.slug
				summary_only = true
			}
		`, teamName)

		check := resource.ComposeAggregateTestCheckFunc(
			resource.TestCheckResourceAttrSet("data.github_team.test", "name"),
			resource.TestCheckResourceAttrSet("data.github_team.test", "node_id"),
			resource.TestCheckResourceAttrSet("data.github_team.test", "privacy"),
			resource.TestCheckResourceAttrSet("data.github_team.test", "notification_setting"),
			resource.TestCheckResourceAttr("data.github_team.test", "members.#", "0"),
			resource.TestCheckResourceAttr("data.github_team.test", "repositories.#", "0"),
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

	t.Run("queries an existing team without error with results_per_page reduced", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		teamName := fmt.Sprintf("%steam-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`
			resource "github_team" "test" {
				name = "%s"
			}

			data "github_team" "test" {
				slug = github_team.test.slug
				results_per_page = 20
			}
		`, teamName)

		check := resource.ComposeAggregateTestCheckFunc(
			resource.TestCheckResourceAttrSet("data.github_team.test", "name"),
			resource.TestCheckResourceAttrSet("data.github_team.test", "node_id"),
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

	t.Run("get team with repositories without erroring", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-%s", testResourcePrefix, randomID)
		teamName := fmt.Sprintf("%steam-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`
		resource "github_repository" "test" {
			name      = "%s"
			auto_init = true
		}

		resource "github_team" "test" {
			name = "%s"
		}

		resource "github_team_repository" "test" {
			team_id    = github_team.test.id
			repository = github_repository.test.name
			permission = "admin"
		}

		data "github_team" "test" {
			slug = github_team.test.slug

			depends_on = [github_repository.test, github_team.test, github_team_repository.test]
		}
		`, repoName, teamName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttrPair("data.github_team.test", "id", "github_team.test", "id"),
						resource.TestCheckResourceAttrPair("data.github_team.test", "node_id", "github_team.test", "node_id"),
						resource.TestCheckResourceAttrPair("data.github_team.test", "slug", "github_team.test", "slug"),
						resource.TestCheckResourceAttrPair("data.github_team.test", "name", "github_team.test", "name"),
						resource.TestCheckResourceAttrPair("data.github_team.test", "description", "github_team.test", "description"),
						resource.TestCheckResourceAttrPair("data.github_team.test", "privacy", "github_team.test", "privacy"),
						resource.TestCheckResourceAttrSet("data.github_team.test", "repositories.#"),
						resource.TestCheckResourceAttr("data.github_team.test", "repositories.#", "1"),
						resource.TestCheckResourceAttrPair("data.github_team.test", "repositories.0", "github_repository.test", "name"),
						resource.TestCheckResourceAttrSet("data.github_team.test", "repositories_detailed.#"),
						resource.TestCheckResourceAttr("data.github_team.test", "repositories_detailed.#", "1"),
						resource.TestCheckResourceAttrPair("data.github_team.test", "repositories_detailed.0.repo_id", "github_repository.test", "repo_id"),
						resource.TestCheckResourceAttrPair("data.github_team.test", "repositories_detailed.0.repo_name", "github_repository.test", "name"),
						resource.TestCheckResourceAttrPair("data.github_team.test", "repositories_detailed.0.role_name", "github_team_repository.test", "permission"),
					),
				},
			},
		})
	})

	t.Run("queries an existing team with connected repositories", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		teamName := fmt.Sprintf("%steam-%s", testResourcePrefix, randomID)
		repoName := fmt.Sprintf("%srepo-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`
			resource "github_team" "test" {
        name = "%s"
			}
			resource "github_repository" "test" {
				name = "%s"
			}
			resource "github_team_repository" "test" {
				team_id    = github_team.test.id
				repository = github_repository.test.name
				permission = "admin"
			}
		`, teamName, repoName)

		config2 := config + `
			data "github_team" "test" {
				slug = github_team.test.slug
			}
		`

		check := resource.ComposeAggregateTestCheckFunc(
			resource.TestCheckResourceAttrSet("data.github_team.test", "name"),
			resource.TestCheckResourceAttr("github_repository.test", "name", repoName),
			resource.TestCheckResourceAttr("data.github_team.test", "repositories_detailed.#", "1"),
			resource.TestCheckResourceAttrPair("data.github_team.test", "repositories_detailed.0.repo_id", "github_repository.test", "repo_id"),
			resource.TestCheckResourceAttrPair("data.github_team.test", "repositories_detailed.0.repo_name", "github_repository.test", "name"),
			resource.TestCheckResourceAttrPair("data.github_team.test", "repositories_detailed.0.role_name", "github_team_repository.test", "permission"),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  resource.ComposeAggregateTestCheckFunc(),
				},
				{
					Config: config2,
					Check:  check,
				},
			},
		})
	})
}
