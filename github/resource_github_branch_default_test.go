package github

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubBranchDefault(t *testing.T) {
	t.Parallel()

	skipUnauthenticated(t)

	t.Run("creates_as_import_without_error", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)

		config := fmt.Sprintf(`
resource "github_branch_default" "test" {
  repository = "%s"
  branch     = "%s"
}
`, repo.GetName(), repo.GetDefaultBranch())

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_branch_default.test", tfjsonpath.New("repository_id"), knownvalue.Int64Exact(repo.GetID())),
					},
				},
			},
		})
	})

	t.Run("creates_default_branch_without_error", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)
		branch := mustCreateTestRepositoryBranch(t, repo)

		config := fmt.Sprintf(`
resource "github_branch_default" "test" {
  repository = "%s"
  branch     = "%s"
}
`, repo.GetName(), branch)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_branch_default.test", tfjsonpath.New("repository_id"), knownvalue.Int64Exact(repo.GetID())),
					},
				},
			},
		})
	})

	t.Run("creates_as_import_with_rename_without_error", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)

		config := fmt.Sprintf(`
resource "github_branch_default" "test" {
  repository      = "%s"
  branch          = "%s"
  rename          = true
  wait_for_rename = true
}
`, repo.GetName(), repo.GetDefaultBranch())

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_branch_default.test", tfjsonpath.New("repository_id"), knownvalue.Int64Exact(repo.GetID())),
					},
				},
			},
		})
	})

	t.Run("creates_with_rename_without_error", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)

		config := fmt.Sprintf(`
resource "github_branch_default" "test" {
  repository      = "%s"
  branch          = "development"
  rename          = true
  wait_for_rename = true
}
`, repo.GetName())

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_branch_default.test", tfjsonpath.New("repository_id"), knownvalue.Int64Exact(repo.GetID())),
					},
				},
			},
		})
	})

	t.Run("updates_default_branch_without_error", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)
		branch := mustCreateTestRepositoryBranch(t, repo)

		config := `
resource "github_branch_default" "test" {
  repository = "%s"
  branch     = "%s"
}
`

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, repo.GetName(), repo.GetDefaultBranch()),
				},
				{
					Config: fmt.Sprintf(config, repo.GetName(), branch),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_branch_default.test", plancheck.ResourceActionUpdate),
						},
					},
				},
			},
		})
	})

	t.Run("updates_default_branch_with_rename_without_error", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)

		config := `
resource "github_branch_default" "test" {
  repository      = "%s"
  branch          = "%s"
  rename          = true
  wait_for_rename = true
}
`

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, repo.GetName(), repo.GetDefaultBranch()),
				},
				{
					Config: fmt.Sprintf(config, repo.GetName(), "development"),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_branch_default.test", plancheck.ResourceActionUpdate),
						},
					},
				},
			},
		})
	})

	t.Run("imports_with_rename_without_error", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)

		config := fmt.Sprintf(`
resource "github_branch_default" "test" {
  repository      = "%s"
  branch          = "development"
  rename          = true
  wait_for_rename = true
}
`, repo.GetName())

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
				},
				{
					ResourceName:            "github_branch_default.test",
					ImportState:             true,
					ImportStateVerify:       true,
					ImportStateVerifyIgnore: []string{"rename", "wait_for_rename"},
				},
			},
		})
	})

	t.Run("destroys_does_not_modify_remote_branch", func(t *testing.T) {
		t.Parallel()

		// The Delete function is no-op since there is no way to "reset" the default branch via the API.
		// This test pins that behavior: the remote default branch must be unchanged
		// after the resource is removed from Terraform state.
		repo := mustCreateTestRepository(t)

		config := fmt.Sprintf(`
resource "github_branch_default" "test" {
  repository      = "%s"
  branch          = "development"
  rename          = true
  wait_for_rename = true
}
`, repo.GetName())

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
				},
				{
					Config: `
removed {
  from = github_branch_default.test

  lifecycle {
    destroy = true
  }
}
`,
					ConfigStateChecks: []statecheck.StateCheck{
						ExpectFunc(func(ctx context.Context, _ statecheck.CheckStateRequest) error {
							r, _, err := testAccConf.meta.v3client.Repositories.Get(ctx, testAccConf.meta.name, repo.GetName())
							if err != nil {
								return err
							}

							if r.GetDefaultBranch() != "development" {
								return fmt.Errorf("expected default branch to remain development, got %s", r.GetDefaultBranch())
							}

							return nil
						}),
					},
				},
			},
		})
	})

	t.Run("gracefully_handles_repository_deleted_out_of_band", func(t *testing.T) {
		t.Parallel()

		// This tests the Read 404 path: when the repository is deleted externally,
		// a state refresh discovers the 404, removes the resource from state
		// without error, and does not attempt to call Delete.
		repo := mustCreateTestRepository(t)

		config := fmt.Sprintf(`
resource "github_branch_default" "test" {
  repository = "%s"
  branch     = "%s"
}
`, repo.GetName(), repo.GetDefaultBranch())

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
				},
				{
					PreConfig: func() {
						if _, err := testAccConf.meta.v3client.Repositories.Delete(t.Context(), testAccConf.meta.name, repo.GetName()); err != nil {
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
		t.Parallel()

		repo := mustCreateTestRepository(t)

		config := `
resource "github_branch_default" "test" {
  repository      = "%s"
  branch          = "development"
  rename          = %t
  wait_for_rename = true
}
`

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, repo.GetName(), true),
				},
				{
					Config: fmt.Sprintf(config, repo.GetName(), false),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_branch_default.test", plancheck.ResourceActionUpdate),
						},
					},
				},
				{
					Config: fmt.Sprintf(config, repo.GetName(), true),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_branch_default.test", plancheck.ResourceActionUpdate),
						},
					},
				},
			},
		})
	})
}

func TestGithubBranchDefault(t *testing.T) {
	// IMPORTANT: This test is not parallelized because it uses global state.

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
