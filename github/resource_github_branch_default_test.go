package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/compare"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubBranchDefault(t *testing.T) {
	t.Run("creates_as_import_without_error", func(t *testing.T) {
		randomID := acctest.RandString(5)
		repoName := fmt.Sprintf("%sbranch-def-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`

			resource "github_repository" "test" {
				name      = "%s"
				auto_init = true
			}

			resource "github_branch_default" "test" {
				repository     = github_repository.test.name
				branch         = "main"
			}
		`, repoName)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_branch_default.test", "branch",
				"main",
			),
			resource.TestCheckResourceAttr(
				"github_branch_default.test", "repository",
				repoName,
			),
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

	t.Run("creates_default_branch_without_error", func(t *testing.T) {
		randomID := acctest.RandString(5)
		repoName := fmt.Sprintf("%sbranch-def-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name      = "%s"
				auto_init = true
			}

			resource "github_branch" "test" {
				repository = github_repository.test.name
				branch     = "test"
			}

			resource "github_branch_default" "test"{
				repository = github_repository.test.name
				branch     = github_branch.test.branch
			}

		`, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_branch_default.test", "branch", "test"),
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

	t.Run("creates_as_import_with_rename_without_error", func(t *testing.T) {
		randomID := acctest.RandString(5)
		repoName := fmt.Sprintf("%sbranch-def-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name      = "%s"
				auto_init = true
			}

			resource "github_branch_default" "test" {
				repository     = github_repository.test.name
				branch         = "main"
				rename         = true
			}
		`, repoName)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_branch_default.test", "branch",
				"main",
			),
			resource.TestCheckResourceAttr(
				"github_branch_default.test", "repository",
				repoName,
			),
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

	t.Run("creates_with_rename_without_error", func(t *testing.T) {
		randomID := acctest.RandString(5)
		repoName := fmt.Sprintf("%sbranch-def-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name      = "%s"
				auto_init = true
			}

			resource "github_branch_default" "test"{
				repository = github_repository.test.name
				branch     = "development"
				rename     = true
			}
		`, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  resource.ComposeTestCheckFunc(resource.TestCheckResourceAttr("github_branch_default.test", "branch", "development")),
				},
			},
		})
	})

	t.Run("updates_default_branch_without_error", func(t *testing.T) {
		randomID := acctest.RandString(5)
		repoName := fmt.Sprintf("%sbranch-def-%s", testResourcePrefix, randomID)

		config := `
			resource "github_repository" "test" {
				name      = "%s"
				auto_init = true
			}
			resource "github_branch" "test" {
				repository = github_repository.test.name
				branch     = "test"
			}

			resource "github_branch_default" "test" {
				repository = github_repository.test.name
				branch     = "%s"
			}
		`

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, repoName, "main"),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.CompareValuePairs(
							"github_branch_default.test", tfjsonpath.New("branch"),
							"github_repository.test", tfjsonpath.New("default_branch"),
							compare.ValuesSame(),
						),
					},
				},
				{
					Config: fmt.Sprintf(config, repoName, "test"),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_branch_default.test", plancheck.ResourceActionUpdate),
						},
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_branch_default.test", tfjsonpath.New("branch"), knownvalue.StringExact("test")),
					},
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

	t.Run("updates_default_branch_with_rename_without_error", func(t *testing.T) {
		randomID := acctest.RandString(5)
		repoName := fmt.Sprintf("%sbranch-def-%s", testResourcePrefix, randomID)

		config := `
			resource "github_repository" "test" {
				name      = "%s"
				auto_init = true
			}

			resource "github_branch_default" "test" {
				repository = github_repository.test.name
				branch     = "%s"
				rename     = true
			}
		`

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, repoName, "main"),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.CompareValuePairs(
							"github_branch_default.test", tfjsonpath.New("branch"),
							"github_repository.test", tfjsonpath.New("default_branch"),
							compare.ValuesSame(),
						),
					},
				},
				{
					Config: fmt.Sprintf(config, repoName, "development"),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_branch_default.test", plancheck.ResourceActionUpdate),
						},
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_branch_default.test", tfjsonpath.New("branch"), knownvalue.StringExact("development")),
					},
				},
			},
		})
	})
	t.Run("imports_with_rename_without_error", func(t *testing.T) {
		randomID := acctest.RandString(5)
		repoName := fmt.Sprintf("%sbranch-def-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name      = "%s"
				auto_init = true
			}

			resource "github_branch_default" "test"{
				repository = github_repository.test.name
				branch     = "development"
				rename     = true
			}
		`, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_branch_default.test", tfjsonpath.New("branch"), knownvalue.StringExact("development")),
					},
				},
				{
					ResourceName:            "github_branch_default.test",
					ImportState:             true,
					ImportStateVerify:       true,
					ImportStateVerifyIgnore: []string{"rename", "etag"},
				},
			},
		})
	})
}
