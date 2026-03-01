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

		config := `
			resource "github_repository" "test1" {
			  name       = "%s"
			  visibility = "private"
			}

			resource "github_repository" "test2" {
			  name       = "%s"
			  visibility = "public"
			  archived   = %t
			  depends_on = [github_repository.test1]
			}
		`
		configWithDS := config + `
			data "github_organization_repositories" "all" {
				ignore_archived_repositories = %t
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
					Config: fmt.Sprintf(config, repo1Name, repo2Name, false),
				},
				{
					Config: fmt.Sprintf(config, repo1Name, repo2Name, true),
				},
				{
					Config: fmt.Sprintf(configWithDS, repo1Name, repo2Name, true, false),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttrSet(resourceAll, "repositories.#"),
					),
				},
				{
					Config: fmt.Sprintf(configWithDS, repo1Name, repo2Name, true),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttrSet(resourceSkipArchived, "repositories.#"),
					),
				},
			},
		})
	})
}
