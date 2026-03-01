package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccGithubOrganizationRepositoriesDataSource(t *testing.T) {
	t.Run("manages repositories", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repo1Name := fmt.Sprintf("%srepo-%s-1", testResourcePrefix, randomID)
		repo2Name := fmt.Sprintf("%srepo-%s-2", testResourcePrefix, randomID)

		config1 := fmt.Sprintf(`
			resource "github_repository" "test1" {
			  name       = "%s"
			  visibility = "private"
			}

			resource "github_repository" "test2" {
			  name       = "%s"
			  visibility = "public"
			  depends_on = [github_repository.test1]
			}
		`, repo1Name, repo2Name)

		config2 := fmt.Sprintf(`
			resource "github_repository" "test1" {
			  name       = "%s"
			  visibility = "private"
			}

			resource "github_repository" "test2" {
			  name       = "%s"
			  archived   = true
			  visibility = "public"
			  depends_on = [github_repository.test1]
			}
		`, repo1Name, repo2Name)

		configAll := config2 + `
			data "github_organization_repositories" "all" {}
		`

		configSkipArchived := config2 + `
			data "github_organization_repositories" "skip_archived" {
			  ignore_archived_repositories = true
			  depends_on = [github_repository.test2]
			}
		`

		const resourceAll = "data.github_organization_repositories.all"
		const resourceSkipArchived = "data.github_organization_repositories.skip_archived"

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config1,
				},
				{
					Config: config2,
				},
				{
					Config: configAll,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttrSet(resourceAll, "repositories.#"),
					),
				},
				{
					Config: configSkipArchived,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttrSet(resourceSkipArchived, "repositories.#"),
					),
				},
			},
		})
	})
}
