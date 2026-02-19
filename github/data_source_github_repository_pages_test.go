package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/compare"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
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
				repository = github_repository.test.name
				build_type = "legacy"
				source {
					branch = "main"
					path   = "/"
				}
			}

			data "github_repository_pages" "test" {
				repository = github_repository.test.name

				depends_on = [github_repository_pages.test]
			}
		`, repoName, baseRepoVisibility)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.CompareValuePairs("data.github_repository_pages.test", tfjsonpath.New("build_type"), "github_repository_pages.test", tfjsonpath.New("build_type"), compare.ValuesSame()),
						statecheck.CompareValuePairs("data.github_repository_pages.test", tfjsonpath.New("source").AtSliceIndex(0).AtMapKey("branch"), "github_repository_pages.test", tfjsonpath.New("source").AtSliceIndex(0).AtMapKey("branch"), compare.ValuesSame()),
						statecheck.CompareValuePairs("data.github_repository_pages.test", tfjsonpath.New("source").AtSliceIndex(0).AtMapKey("path"), "github_repository_pages.test", tfjsonpath.New("source").AtSliceIndex(0).AtMapKey("path"), compare.ValuesSame()),
						statecheck.CompareValuePairs("data.github_repository_pages.test", tfjsonpath.New("api_url"), "github_repository_pages.test", tfjsonpath.New("api_url"), compare.ValuesSame()),
					},
				},
			},
		})
	})
}
