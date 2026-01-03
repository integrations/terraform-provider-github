package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubBranch(t *testing.T) {
	t.Run("creates a branch directly", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-branch-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`
		resource "github_repository" "test" {
			name = "%[1]s"
			auto_init = true
		}

		resource "github_branch" "test" {
			repository = github_repository.test.name
			branch     = "test"
		}
		`, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttrSet("github_branch.test", "id"),
						resource.TestCheckResourceAttr("github_branch.test", "ref", "refs/heads/test"),
						resource.TestCheckResourceAttrSet("github_branch.test", "sha"),
					),
				},
			},
		})
	})

	t.Run("creates a branch named main directly and a repository with a gitignore template", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-branch-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`
		resource "github_repository" "test" {
			name = "%[1]s"
			auto_init = true
			gitignore_template = "Python"
		}

		resource "github_branch" "test" {
			repository = github_repository.test.id
			branch     = "main"
		}
		`, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttrSet("github_branch.test", "id"),
						resource.TestCheckResourceAttr("github_branch.test", "ref", "refs/heads/main"),
						resource.TestCheckResourceAttrSet("github_branch.test", "sha"),
					),
				},
				{
					Config: `
					removed {
					  from = github_branch.test
					  lifecycle { destroy = false }
					}
					`,
				},
			},
		})
	})

	t.Run("creates a branch from a source branch", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-branch-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`
		resource "github_repository" "test" {
			name = "%[1]s"
			auto_init = true
		}

		resource "github_branch" "source" {
			repository = github_repository.test.id
			branch     = "source"
		}

		resource "github_branch" "test" {
			repository    = github_repository.test.id
			source_branch = github_branch.source.branch
			branch        = "test"
		}
		`, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttrSet("github_branch.test", "id"),
						resource.TestCheckResourceAttr("github_branch.test", "ref", "refs/heads/test"),
						resource.TestCheckResourceAttrPair("github_branch.test", "sha", "github_branch.source", "sha"),
					),
				},
			},
		})
	})

	t.Run("renames a branch without replacement", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-branch-%s", testResourcePrefix, randomID)
		initialConfig := fmt.Sprintf(`
			resource "github_repository" "test" {
			  name = "%[1]s"
			  auto_init = true
			}

			resource "github_branch" "test" {
			  repository = github_repository.test.id
			  branch     = "initial"
			}
		`, repoName)

		renamedConfig := fmt.Sprintf(`
			resource "github_repository" "test" {
			  name = "%[1]s"
			  auto_init = true
			}

			resource "github_branch" "test" {
			  repository = github_repository.test.id
			  branch     = "renamed"
			}
		`, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:  func() { skipUnauthenticated(t) },
			Providers: testAccProviders,
			Steps: []resource.TestStep{
				{
					Config: initialConfig,
				},
				{
					Config: renamedConfig,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr(
							"github_branch.test", "branch", "renamed",
						),
					),
				},
			},
		})
	})
}
