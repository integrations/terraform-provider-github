package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.CompareValuePairs("github_branch_default.test", tfjsonpath.New("repository"), "github_repository.test", tfjsonpath.New("name"), compare.ValuesSame()),
						statecheck.CompareValuePairs("github_branch_default.test", tfjsonpath.New("branch"), "github_repository.test", tfjsonpath.New("default_branch"), compare.ValuesSame()),
						statecheck.ExpectKnownValue("github_branch_default.test", tfjsonpath.New("repository_id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_branch_default.test", tfjsonpath.New("etag"), knownvalue.NotNull()),
					},
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
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.CompareValuePairs("github_branch_default.test", tfjsonpath.New("repository"), "github_repository.test", tfjsonpath.New("name"), compare.ValuesSame()),
						statecheck.CompareValuePairs("github_branch_default.test", tfjsonpath.New("branch"), "github_branch.test", tfjsonpath.New("branch"), compare.ValuesSame()),
						statecheck.ExpectKnownValue("github_branch_default.test", tfjsonpath.New("repository_id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_branch_default.test", tfjsonpath.New("etag"), knownvalue.NotNull()),
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
				wait_for_rename = true
			}
		`, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.CompareValuePairs("github_branch_default.test", tfjsonpath.New("repository"), "github_repository.test", tfjsonpath.New("name"), compare.ValuesSame()),
						statecheck.CompareValuePairs("github_branch_default.test", tfjsonpath.New("branch"), "github_repository.test", tfjsonpath.New("default_branch"), compare.ValuesSame()),
						statecheck.ExpectKnownValue("github_branch_default.test", tfjsonpath.New("repository_id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_branch_default.test", tfjsonpath.New("etag"), knownvalue.NotNull()),
					},
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
				repository      = github_repository.test.name
				branch          = "development"
				rename          = true
				wait_for_rename = true
			}
		`, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.CompareValuePairs("github_branch_default.test", tfjsonpath.New("repository"), "github_repository.test", tfjsonpath.New("name"), compare.ValuesSame()),
						statecheck.ExpectKnownValue("github_branch_default.test", tfjsonpath.New("branch"), knownvalue.StringExact("development")),
						statecheck.ExpectKnownValue("github_branch_default.test", tfjsonpath.New("repository_id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_branch_default.test", tfjsonpath.New("etag"), knownvalue.NotNull()),
					},
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
				repository      = github_repository.test.name
				branch          = "%s"
				rename          = true
				wait_for_rename = true
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
				repository      = github_repository.test.name
				branch          = "development"
				rename          = true
				wait_for_rename = true
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
					ImportStateVerifyIgnore: []string{"rename", "wait_for_rename", "etag"},
				},
			},
		})
	})
	t.Run("destroys_without_error", func(t *testing.T) {
		randomID := acctest.RandString(5)
		repoName := fmt.Sprintf("%sbranch-def-%s", testResourcePrefix, randomID)

		repoOnlyConfig := fmt.Sprintf(`
			resource "github_repository" "test" {
				name      = "%s"
				auto_init = true
			}
		`, repoName)

		fullConfig := fmt.Sprintf(`
			resource "github_repository" "test" {
				name      = "%s"
				auto_init = true
			}

			resource "github_branch_default" "test" {
				repository = github_repository.test.name
				branch     = "main"
			}
		`, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fullConfig,
				},
				{
					Config: repoOnlyConfig,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository.test", tfjsonpath.New("default_branch"), knownvalue.StringExact("main")),
					},
				},
			},
		})
	})

	t.Run("destroys_does_not_modify_remote_branch", func(t *testing.T) {
		// The Delete function is no-op since there is no way to "reset" the default branch via the API.
		// This test pins that behavior: the remote default branch must be unchanged
		// after the resource is removed from Terraform state.
		randomID := acctest.RandString(5)
		repoName := fmt.Sprintf("%sbranch-def-%s", testResourcePrefix, randomID)

		repoOnlyConfig := fmt.Sprintf(`
			resource "github_repository" "test" {
				name      = "%s"
				auto_init = true
			}
		`, repoName)

		fullConfig := fmt.Sprintf(`
			resource "github_repository" "test" {
				name      = "%s"
				auto_init = true
			}

			resource "github_branch_default" "test" {
				repository      = github_repository.test.name
				branch          = "development"
				rename          = true
				wait_for_rename = true
			}
		`, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fullConfig,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_branch_default.test", tfjsonpath.New("branch"), knownvalue.StringExact("development")),
					},
				},
				{
					Config: repoOnlyConfig,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository.test", tfjsonpath.New("default_branch"), knownvalue.StringExact("development")),
					},
				},
			},
		})
	})

	t.Run("gracefully_handles_repository_deleted_out_of_band", func(t *testing.T) {
		// This tests the Read 404 path: when the repository is deleted externally,
		// a state refresh discovers the 404, removes both resources from state
		// without error, and does not attempt to call Delete.
		randomID := acctest.RandString(5)
		repoName := fmt.Sprintf("%sbranch-def-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name      = "%s"
				auto_init = true
			}

			resource "github_branch_default" "test" {
				repository = github_repository.test.name
				branch     = "main"
			}
		`, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
				},
				{
					PreConfig: func() {
						meta, err := getTestMeta()
						if err != nil {
							t.Errorf("failed to get test meta: %s", err)
							return
						}
						if _, err := meta.v3client.Repositories.Delete(t.Context(), meta.name, repoName); err != nil {
							t.Errorf("failed to delete repository out-of-band: %s", err)
						}
					},
					RefreshState:       true,
					ExpectNonEmptyPlan: true,
				},
			},
		})
	})

	t.Run("regression_prevent_trying_rename_to_same_name", func(t *testing.T) {
		randomID := acctest.RandString(5)
		repoName := fmt.Sprintf("%sbranch-def-%s", testResourcePrefix, randomID)
		config := `
			resource "github_repository" "test" {
				name      = "%s"
				auto_init = true
			}

			resource "github_branch_default" "test"{
				repository      = github_repository.test.name
				branch          = "development"
				rename          = %t
				wait_for_rename = true
			}
		`

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, repoName, true),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_branch_default.test", tfjsonpath.New("branch"), knownvalue.StringExact("development")),
					},
				},
				{
					Config: fmt.Sprintf(config, repoName, false),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_branch_default.test", tfjsonpath.New("branch"), knownvalue.StringExact("development")),
					},
				},
				{
					Config: fmt.Sprintf(config, repoName, true),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_branch_default.test", tfjsonpath.New("branch"), knownvalue.StringExact("development")),
					},
				},
			},
		})
	})
}

func TestGithubBranchDefault(t *testing.T) {
	// Test verifies that waitForDefaultBranch
	// is not called when wait_for_rename is false (the default). The mock server
	// is set up with exactly two responses: the initial GET and the rename POST.
	// Any additional request (i.e. a polling GET from waitForDefaultBranch) would
	// receive a 400 from the mock, causing the resource function to return an error.
	t.Run("skips_wait_for_rename_when_not_configured", func(t *testing.T) {
		ts := githubApiMock([]*mockResponse{
			{
				ExpectedUri:    "/repos/owner/repo",
				ExpectedMethod: "GET",
				StatusCode:     200,
				ResponseBody:   `{"id": 42, "name": "repo", "default_branch": "main"}`,
			},
			{
				ExpectedUri:    "/repos/owner/repo/branches/main/rename",
				ExpectedMethod: "POST",
				StatusCode:     201,
				ResponseBody:   `{"name": "development"}`,
			},
		})
		defer ts.Close()

		client := mustCreateTestGitHubClient(t, ts.URL)
		meta := &Owner{name: "owner", v3client: client}

		d := schema.TestResourceDataRaw(t, resourceGithubBranchDefault().Schema, map[string]any{
			"repository":      "repo",
			"branch":          "development",
			"rename":          true,
			"wait_for_rename": false,
		})

		diags := resourceGithubBranchDefaultCreate(t.Context(), d, meta)
		if diags.HasError() {
			t.Fatalf("expected no error, got: %v", diags)
		}
	})
}
