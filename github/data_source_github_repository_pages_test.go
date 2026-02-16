package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccGithubRepositoryPagesDataSource(t *testing.T) {
	baseRepoVisibility := "public"
	if testAccConf.authMode == enterprise {
		baseRepoVisibility = "private"
	}

	t.Run("reads_pages_configuration", func(t *testing.T) {
		randomID := acctest.RandString(5)
		repoName := fmt.Sprintf("%spages-ds-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name       = "%s"
				visibility = "%s"
				auto_init  = true
			}

			resource "github_repository_pages" "test" {
				owner      = "%s"
				repository = github_repository.test.name
				build_type = "legacy"
				source {
					branch = "main"
					path   = "/"
				}
			}

			data "github_repository_pages" "test" {
				owner           = "%s"
				repository = github_repository.test.name

				depends_on = [github_repository_pages.test]
			}
		`, repoName, baseRepoVisibility, testAccConf.owner, testAccConf.owner)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("data.github_repository_pages.test", "build_type", "legacy"),
						resource.TestCheckResourceAttr("data.github_repository_pages.test", "source.0.branch", "main"),
						resource.TestCheckResourceAttr("data.github_repository_pages.test", "source.0.path", "/"),
						resource.TestCheckResourceAttrSet("data.github_repository_pages.test", "url"),
					),
				},
			},
		})
	})
}
