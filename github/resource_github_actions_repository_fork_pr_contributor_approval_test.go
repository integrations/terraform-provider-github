package github

import (
	"fmt"
	"testing"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

// GitHub rejects this setting on private repositories unless the owner allows fork PR workflows on
// them, so the tests use public repositories.
func withPublicVisibility(r *github.Repository) {
	r.Visibility = new("public")
}

func TestAccGithubActionsRepositoryForkPRContributorApproval(t *testing.T) {
	const resourceName = "github_actions_repository_fork_pr_contributor_approval.test"

	config := `
		resource "github_actions_repository_fork_pr_contributor_approval" "test" {
			repository      = "%s"
			approval_policy = "%s"
		}
	`

	for _, policy := range []string{"first_time_contributors_new_to_github", "first_time_contributors", "all_external_contributors"} {
		t.Run(fmt.Sprintf("creates_with_%s", policy), func(t *testing.T) {
			t.Parallel()

			randomID := acctest.RandString(5)
			repoName := fmt.Sprintf("%srepo-fork-pr-approval-%s", testResourcePrefix, randomID)

			resource.Test(t, resource.TestCase{
				PreCheck:          func() { skipUnlessMode(t, individual, organization) },
				ProviderFactories: providerFactories,
				Steps: []resource.TestStep{
					{
						Config: fmt.Sprintf(`
							resource "github_repository" "test" {
								name       = "%s"
								visibility = "public"
							}

							resource "github_actions_repository_fork_pr_contributor_approval" "test" {
								repository      = github_repository.test.name
								approval_policy = "%s"
							}
						`, repoName, policy),
						ConfigStateChecks: []statecheck.StateCheck{
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("approval_policy"), knownvalue.StringExact(policy)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("repository"), knownvalue.StringExact(repoName)),
							statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("repository_id"), knownvalue.NotNull()),
						},
					},
				},
			})
		})
	}

	t.Run("updates_approval_policy", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t, withPublicVisibility)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessMode(t, individual, organization) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, repo.GetName(), "first_time_contributors"),
				},
				{
					Config: fmt.Sprintf(config, repo.GetName(), "all_external_contributors"),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
						},
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("approval_policy"), knownvalue.StringExact("all_external_contributors")),
					},
				},
			},
		})
	})

	t.Run("imports_by_repository_name", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t, withPublicVisibility)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessMode(t, individual, organization) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, repo.GetName(), "all_external_contributors"),
				},
				{
					ResourceName:      resourceName,
					ImportState:       true,
					ImportStateId:     repo.GetName(),
					ImportStateVerify: true,
				},
			},
		})
	})

	t.Run("updates_renamed_repo", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t, withPublicVisibility)
		newRepoName := fmt.Sprintf("%s-updated", repo.GetName())

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessMode(t, individual, organization) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, repo.GetName(), "all_external_contributors"),
				},
				{
					PreConfig: func() {
						mustRenameTestRepository(t, repo, newRepoName)
					},
					Config: fmt.Sprintf(config, newRepoName, "all_external_contributors"),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
						},
					},
				},
			},
		})
	})

	t.Run("recreates_changed_repo", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t, withPublicVisibility)
		repo2 := mustCreateTestRepository(t, withPublicVisibility)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessMode(t, individual, organization) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, repo.GetName(), "all_external_contributors"),
				},
				{
					Config: fmt.Sprintf(config, repo2.GetName(), "all_external_contributors"),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionReplace),
						},
					},
				},
			},
		})
	})
}
