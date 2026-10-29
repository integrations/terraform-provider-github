package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccGithubRepositoryBranchesDataSource(t *testing.T) {
	t.Parallel()

	t.Run("manages_branches_of_a_new_repository", func(t *testing.T) {
		t.Parallel()

		repoName := fmt.Sprintf("%srepo-branches-%s", testResourcePrefix, acctest.RandString(5))
		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name      = "%s"
				auto_init = true
			}

			data "github_repository_branches" "test" {
				repository = github_repository.test.name
			}
		`, repoName)

		const resourceName = "data.github_repository_branches.test"
		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceName, "branches.#", "1"),
			resource.TestCheckResourceAttr(resourceName, "branches.0.name", "main"),
			resource.TestCheckResourceAttr(resourceName, "branches.0.protected", "false"),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  check,
				},
			},
		})
	})

	t.Run("manages_branches_of_a_new_repository_with_filtering", func(t *testing.T) {
		t.Parallel()
		skipUnlessHasOrgs(t)

		repoName := fmt.Sprintf("%srepo-branches-%s", testResourcePrefix, acctest.RandString(5))
		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name      = "%s"
				auto_init = true
			}

			resource "github_branch" "test" {
				repository = github_repository.test.id
				branch     = "test"
			}

			resource "github_branch_protection_v3" "test" {
		  	repository = github_repository.test.name
		  	branch     = "test"
		  	depends_on = [github_branch.test]
		  }
		 `, repoName)

		config2 := config + `
			data "github_repository_branches" "test" {
				repository = github_repository.test.name
			}

			data "github_repository_branches" "protected" {
				repository              = github_repository.test.name
				only_protected_branches = true
			}

			data "github_repository_branches" "non_protected" {
				repository                  = github_repository.test.name
				only_non_protected_branches = true
			}
		`

		const resourceName = "data.github_repository_branches.test"
		const protectedResourceName = "data.github_repository_branches.protected"
		const nonProtectedResourceName = "data.github_repository_branches.non_protected"
		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceName, "branches.#", "2"),
			resource.TestCheckResourceAttr(protectedResourceName, "branches.#", "1"),
			resource.TestCheckResourceAttr(protectedResourceName, "branches.0.name", "test"),
			resource.TestCheckResourceAttr(protectedResourceName, "branches.0.protected", "true"),
			resource.TestCheckResourceAttr(nonProtectedResourceName, "branches.#", "1"),
			resource.TestCheckResourceAttr(nonProtectedResourceName, "branches.0.name", "main"),
			resource.TestCheckResourceAttr(nonProtectedResourceName, "branches.0.protected", "false"),
		)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
				},
				{
					Config: config2,
					Check:  check,
				},
			},
		})
	})
}
