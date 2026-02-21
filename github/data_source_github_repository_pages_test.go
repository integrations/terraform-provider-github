package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/compare"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
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
			PreCheck: func() {
				skipUnauthenticated(t)
			},
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.CompareValuePairs("data.github_repository_pages.test", tfjsonpath.New("build_type"), "github_repository_pages.test", tfjsonpath.New("build_type"), compare.ValuesSame()),
						statecheck.CompareValuePairs("data.github_repository_pages.test", tfjsonpath.New("source").AtSliceIndex(0).AtMapKey("branch"), "github_repository_pages.test", tfjsonpath.New("source").AtSliceIndex(0).AtMapKey("branch"), compare.ValuesSame()),
						statecheck.CompareValuePairs("data.github_repository_pages.test", tfjsonpath.New("source").AtSliceIndex(0).AtMapKey("path"), "github_repository_pages.test", tfjsonpath.New("source").AtSliceIndex(0).AtMapKey("path"), compare.ValuesSame()),
						statecheck.CompareValuePairs("data.github_repository_pages.test", tfjsonpath.New("cname"), "github_repository_pages.test", tfjsonpath.New("cname"), compare.ValuesSame()),
						statecheck.CompareValuePairs("data.github_repository_pages.test", tfjsonpath.New("custom_404"), "github_repository_pages.test", tfjsonpath.New("custom_404"), compare.ValuesSame()),
						statecheck.CompareValuePairs("data.github_repository_pages.test", tfjsonpath.New("html_url"), "github_repository_pages.test", tfjsonpath.New("html_url"), compare.ValuesSame()),
						statecheck.ExpectKnownValue("data.github_repository_pages.test", tfjsonpath.New("build_status"), knownvalue.NotNull()),
						statecheck.CompareValuePairs("data.github_repository_pages.test", tfjsonpath.New("api_url"), "github_repository_pages.test", tfjsonpath.New("api_url"), compare.ValuesSame()),
						statecheck.ExpectKnownValue("data.github_repository_pages.test", tfjsonpath.New("public"), knownvalue.Bool(testAccConf.authMode != enterprise)),
						statecheck.CompareValuePairs("data.github_repository_pages.test", tfjsonpath.New("https_enforced"), "github_repository_pages.test", tfjsonpath.New("https_enforced"), compare.ValuesSame()),
					},
				},
			},
		})
	})
	t.Run("reads_pages_enterprise_configuration", func(t *testing.T) {
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
				public = false
			}

			data "github_repository_pages" "test" {
				repository = github_repository.test.name

				depends_on = [github_repository_pages.test]
			}
		`, repoName, baseRepoVisibility)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.CompareValuePairs("data.github_repository_pages.test", tfjsonpath.New("build_type"), "github_repository_pages.test", tfjsonpath.New("build_type"), compare.ValuesSame()),
						statecheck.CompareValuePairs("data.github_repository_pages.test", tfjsonpath.New("source").AtSliceIndex(0).AtMapKey("branch"), "github_repository_pages.test", tfjsonpath.New("source").AtSliceIndex(0).AtMapKey("branch"), compare.ValuesSame()),
						statecheck.CompareValuePairs("data.github_repository_pages.test", tfjsonpath.New("source").AtSliceIndex(0).AtMapKey("path"), "github_repository_pages.test", tfjsonpath.New("source").AtSliceIndex(0).AtMapKey("path"), compare.ValuesSame()),
						statecheck.CompareValuePairs("data.github_repository_pages.test", tfjsonpath.New("cname"), "github_repository_pages.test", tfjsonpath.New("cname"), compare.ValuesSame()),
						statecheck.CompareValuePairs("data.github_repository_pages.test", tfjsonpath.New("custom_404"), "github_repository_pages.test", tfjsonpath.New("custom_404"), compare.ValuesSame()),
						statecheck.CompareValuePairs("data.github_repository_pages.test", tfjsonpath.New("html_url"), "github_repository_pages.test", tfjsonpath.New("html_url"), compare.ValuesSame()),
						statecheck.ExpectKnownValue("data.github_repository_pages.test", tfjsonpath.New("build_status"), knownvalue.NotNull()),
						statecheck.CompareValuePairs("data.github_repository_pages.test", tfjsonpath.New("api_url"), "github_repository_pages.test", tfjsonpath.New("api_url"), compare.ValuesSame()),
						statecheck.CompareValuePairs("data.github_repository_pages.test", tfjsonpath.New("public"), "github_repository_pages.test", tfjsonpath.New("public"), compare.ValuesSame()),
						statecheck.CompareValuePairs("data.github_repository_pages.test", tfjsonpath.New("https_enforced"), "github_repository_pages.test", tfjsonpath.New("https_enforced"), compare.ValuesSame()),
					},
				},
			},
		})
	})
}
