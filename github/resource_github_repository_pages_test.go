package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccGithubRepositoryPages(t *testing.T) {
	baseRepoVisibility := "public"
	if testAccConf.authMode == enterprise {
		baseRepoVisibility = "private"
	}

	t.Run("creates_pages_with_legacy_build_type", func(t *testing.T) {
		randomID := acctest.RandString(5)
		repoName := fmt.Sprintf("%spages-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name       = "%s"
				visibility = "%s"
				auto_init  = true

			}

			resource "github_repository_pages" "test" {
				owner      = "%s"
				repository_name = github_repository.test.name
				build_type = "legacy"
				source {
					branch = "main"
					path   = "/"
				}
			}
		`, repoName, baseRepoVisibility, testAccConf.owner)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_repository_pages.test", "build_type", "legacy"),
						resource.TestCheckResourceAttr("github_repository_pages.test", "source.0.branch", "main"),
						resource.TestCheckResourceAttr("github_repository_pages.test", "source.0.path", "/"),
						resource.TestCheckResourceAttrSet("github_repository_pages.test", "api_url"),
					),
				},
			},
		})
	})

	t.Run("creates_pages_with_workflow_build_type", func(t *testing.T) {
		randomID := acctest.RandString(5)
		repoName := fmt.Sprintf("%spages-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name       = "%s"
				visibility = "%s"
				auto_init  = true

			}

			resource "github_repository_pages" "test" {
				owner      = "%s"
				repository_name = github_repository.test.name
				build_type = "workflow"
			}
		`, repoName, baseRepoVisibility, testAccConf.owner)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_repository_pages.test", "build_type", "workflow"),
					),
				},
			},
		})
	})

	t.Run("updates_pages_configuration", func(t *testing.T) {
		randomID := acctest.RandString(5)
		repoName := fmt.Sprintf("%spages-%s", testResourcePrefix, randomID)

		sourceConfig := `
source {
	branch = "main"
	path   = "/"
}
`
		config := `
			resource "github_repository" "test" {
				name       = "%s"
				visibility = "%s"
				auto_init  = true

			}

			resource "github_repository_pages" "test" {
				owner      = "%s"
				repository_name = github_repository.test.name
				build_type = "%s"
				%s
			}
		`

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, repoName, baseRepoVisibility, testAccConf.owner, "legacy", sourceConfig),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_repository_pages.test", "build_type", "legacy"),
					),
				},
				{
					Config: fmt.Sprintf(config, repoName, baseRepoVisibility, testAccConf.owner, "workflow", ""),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_repository_pages.test", "build_type", "workflow"),
					),
				},
			},
		})
	})

	t.Run("imports_pages_configuration", func(t *testing.T) {
		randomID := acctest.RandString(5)
		repoName := fmt.Sprintf("%spages-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name       = "%s"
				visibility = "%s"
				auto_init  = true

			}

			resource "github_repository_pages" "test" {
				owner      = "%s"
				repository_name = github_repository.test.name
				build_type = "legacy"
				source {
					branch = "main"
					path   = "/"
				}
			}
		`, repoName, baseRepoVisibility, testAccConf.owner)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_repository_pages.test", "build_type", "legacy"),
					),
				},
				{
					ResourceName:            "github_repository_pages.test",
					ImportState:             true,
					ImportStateVerify:       true,
					ImportStateVerifyIgnore: []string{"build_status"},
				},
			},
		})
	})
}
