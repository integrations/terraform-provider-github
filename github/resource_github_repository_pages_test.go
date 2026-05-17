package github

import (
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/compare"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubRepositoryPages(t *testing.T) {
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
		`, repoName, testAccConf.testRepositoryVisibility)

		resource.ParallelTest(t, resource.TestCase{
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
		`, repoName, testAccConf.testRepositoryVisibility)

		resource.ParallelTest(t, resource.TestCase{
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

		resource.ParallelTest(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, repoName, testAccConf.testRepositoryVisibility, "legacy", sourceConfig),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_pages.test", tfjsonpath.New("build_type"), knownvalue.StringExact("legacy")),
					},
				},
				{
					Config: fmt.Sprintf(config, repoName, testAccConf.testRepositoryVisibility, "workflow", ""),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_pages.test", tfjsonpath.New("build_type"), knownvalue.StringExact("workflow")),
					},
				},
			},
		})
	})

	t.Run("creates_pages_with_private_visibility", func(t *testing.T) {
		randomID := acctest.RandString(5)
		repoName := fmt.Sprintf("%spages-%s", testResourcePrefix, randomID)

		config := `
			resource "github_repository" "test" {
				name       = "%s"
				visibility = "%s"
				auto_init  = true

			}

			resource "github_repository_pages" "test" {
				repository = github_repository.test.name
				build_type = "workflow"
				
				public = false
			}
		`

		resource.ParallelTest(t, resource.TestCase{
			PreCheck: func() {
				skipUnlessEnterprise(t)
			},
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, repoName, testAccConf.testRepositoryVisibility),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_pages.test", tfjsonpath.New("public"), knownvalue.Bool(false)),
					},
				},
			},
		})
	})
	t.Run("updates_pages_visibility", func(t *testing.T) {
		randomID := acctest.RandString(5)
		repoName := fmt.Sprintf("%spages-%s", testResourcePrefix, randomID)

		config := `
			resource "github_repository" "test" {
				name       = "%s"
				visibility = "%s"
				auto_init  = true

			}

			resource "github_repository_pages" "test" {
				repository = github_repository.test.name
				build_type = "workflow"
				
				public = %t
			}
		`

		publicValuesDiffer := statecheck.CompareValue(compare.ValuesDiffer())

		resource.ParallelTest(t, resource.TestCase{
			PreCheck: func() {
				skipUnlessEnterprise(t)
				if os.Getenv("GH_TEST_ENTERPRISE_IS_EMU") == "true" {
					t.Skip("Skipping as enterprise test mode is EMU")
				}
			},
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, repoName, testAccConf.testRepositoryVisibility, true),
					ConfigStateChecks: []statecheck.StateCheck{
						publicValuesDiffer.AddStateValue("github_repository_pages.test", tfjsonpath.New("public")),
					},
				},
				{
					Config: fmt.Sprintf(config, repoName, testAccConf.testRepositoryVisibility, false),
					ConfigStateChecks: []statecheck.StateCheck{
						publicValuesDiffer.AddStateValue("github_repository_pages.test", tfjsonpath.New("public")),
					},
				},
			},
		})
	})

	t.Run("errors_when_https_enforced_without_cname", func(t *testing.T) {
		randomID := acctest.RandString(5)
		repoName := fmt.Sprintf("%spages-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name       = "%s"
				visibility = "%s"
				auto_init  = true
			}

			resource "github_repository_pages" "test" {
				repository     = github_repository.test.name
				build_type     = "workflow"
				https_enforced = true
			}
		`, repoName, testAccConf.testRepositoryVisibility)

		resource.ParallelTest(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      config,
					ExpectError: regexp.MustCompile(`all of .cname,https_enforced. must be specified`),
				},
			},
		})
	})

	t.Run("validates_that_source_is_not_set_for_workflow_build_type", func(t *testing.T) {
		randomID := acctest.RandString(5)
		repoName := fmt.Sprintf("%spages-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name       = "%s"
				visibility = "%s"
				auto_init  = true
			}

			resource "github_repository_pages" "test" {
				repository     = github_repository.test.name
				build_type     = "workflow"
				source {
					branch = "main"
					path   = "/"
				}
			}
		`, repoName, testAccConf.testRepositoryVisibility)

		resource.ParallelTest(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
					ExpectError: regexp.MustCompile(`'source' is not supported for workflow build type`),
				},
			},
		})
	})

	t.Run("validates_that_source_is_set_for_legacy_build_type", func(t *testing.T) {
		randomID := acctest.RandString(5)
		repoName := fmt.Sprintf("%spages-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name       = "%s"
				visibility = "%s"
				auto_init  = true
			}

			resource "github_repository_pages" "test" {
				repository     = github_repository.test.name
				build_type     = "legacy"
			}
		`, repoName, testAccConf.testRepositoryVisibility)

		resource.ParallelTest(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
					ExpectError: regexp.MustCompile(`'source' is required for legacy build type`),
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
		`, repoName, testAccConf.testRepositoryVisibility)

		resource.ParallelTest(t, resource.TestCase{
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
