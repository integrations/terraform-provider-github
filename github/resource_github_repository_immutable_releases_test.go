package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubRepositoryImmutableReleases(t *testing.T) {
	t.Parallel()

	skipUnauthenticated(t)

	t.Run("default", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)

		config := fmt.Sprintf(`
resource "github_repository_immutable_releases" "test" {
  repository = "%s"
}
`, repo.GetName())

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_immutable_releases.test", tfjsonpath.New("enabled"), knownvalue.Bool(true)),
						statecheck.ExpectKnownValue("github_repository_immutable_releases.test", tfjsonpath.New("repository_id"), knownvalue.Int64Exact(repo.GetID())),
					},
				},
				{
					ResourceName:      "github_repository_immutable_releases.test",
					ImportState:       true,
					ImportStateId:     repo.GetName(),
					ImportStateVerify: true,
				},
			},
		})
	})

	t.Run("updates_enabled", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)

		config := `
resource "github_repository_immutable_releases" "test" {
  repository = "%s"
  enabled    = %t
}
`

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, repo.GetName(), false),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_immutable_releases.test", tfjsonpath.New("enabled"), knownvalue.Bool(false)),
					},
				},
				{
					Config: fmt.Sprintf(config, repo.GetName(), true),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_repository_immutable_releases.test", plancheck.ResourceActionUpdate),
						},
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_immutable_releases.test", tfjsonpath.New("enabled"), knownvalue.Bool(true)),
					},
				},
				{
					Config: fmt.Sprintf(config, repo.GetName(), false),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_repository_immutable_releases.test", plancheck.ResourceActionUpdate),
						},
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_immutable_releases.test", tfjsonpath.New("enabled"), knownvalue.Bool(false)),
					},
				},
			},
		})
	})

	t.Run("updates_renamed_repo", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)
		newRepoName := fmt.Sprintf("%s-updated", repo.GetName())

		config := `
resource "github_repository_immutable_releases" "test" {
  repository = "%s"
}
`

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, repo.GetName()),
				},
				{
					PreConfig: func() {
						mustRenameTestRepository(t, repo, newRepoName)
					},
					Config: fmt.Sprintf(config, newRepoName),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_repository_immutable_releases.test", plancheck.ResourceActionUpdate),
						},
					},
				},
			},
		})
	})

	t.Run("recreates_changed_repo", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)
		repo2 := mustCreateTestRepository(t)

		config := `
resource "github_repository_immutable_releases" "test" {
  repository = "%s"
}
`

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, repo.GetName()),
				},
				{
					Config: fmt.Sprintf(config, repo2.GetName()),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_repository_immutable_releases.test", plancheck.ResourceActionReplace),
						},
					},
				},
			},
		})
	})

	t.Run("removes_from_state_when_repository_deleted", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)

		config := fmt.Sprintf(`
resource "github_repository_immutable_releases" "test" {
  repository = "%s"
  enabled    = false
}
`, repo.GetName())

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
				},
				{
					PreConfig: func() {
						if _, err := testAccConf.meta.v3client.Repositories.Delete(t.Context(), testAccConf.meta.name, repo.GetName()); err != nil {
							t.Fatalf("failed to delete repository out-of-band: %s", err)
						}
					},
					RefreshState:       true,
					ExpectNonEmptyPlan: true,
					RefreshPlanChecks: resource.RefreshPlanChecks{
						PostRefresh: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_repository_immutable_releases.test", plancheck.ResourceActionCreate),
						},
					},
				},
			},
		})
	})
}
