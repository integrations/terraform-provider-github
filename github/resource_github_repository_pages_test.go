package github

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
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
				repository = github_repository.test.name
				build_type = "legacy"
				source {
					branch = "main"
					path   = "/"
				}
			}
		`, repoName, baseRepoVisibility)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_pages.test", tfjsonpath.New("build_type"), knownvalue.StringExact("legacy")),
						statecheck.ExpectKnownValue("github_repository_pages.test", tfjsonpath.New("source").AtSliceIndex(0).AtMapKey("branch"), knownvalue.StringExact("main")),
						statecheck.ExpectKnownValue("github_repository_pages.test", tfjsonpath.New("source").AtSliceIndex(0).AtMapKey("path"), knownvalue.StringExact("/")),
						statecheck.ExpectKnownValue("github_repository_pages.test", tfjsonpath.New("api_url"), knownvalue.StringRegexp(regexp.MustCompile("https://.*"))),
					},
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
				repository = github_repository.test.name
				build_type = "workflow"
			}
		`, repoName, baseRepoVisibility)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_pages.test", tfjsonpath.New("build_type"), knownvalue.StringExact("workflow")),
					},
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
				repository = github_repository.test.name
				build_type = "%s"
				%s
			}
		`

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, repoName, baseRepoVisibility, "legacy", sourceConfig),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_pages.test", tfjsonpath.New("build_type"), knownvalue.StringExact("legacy")),
					},
				},
				{
					Config: fmt.Sprintf(config, repoName, baseRepoVisibility, "workflow", ""),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_pages.test", tfjsonpath.New("build_type"), knownvalue.StringExact("workflow")),
					},
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
				repository = github_repository.test.name
				build_type = "legacy"
				source {
					branch = "main"
					path   = "/"
				}
			}
		`, repoName, baseRepoVisibility)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_pages.test", tfjsonpath.New("build_type"), knownvalue.StringExact("legacy")),
					},
				},
				{
					ResourceName:  "github_repository_pages.test",
					ImportState:   true,
					ImportStateId: repoName,
				},
			},
		})
	})
}
