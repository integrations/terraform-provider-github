package github

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubActionsVariable(t *testing.T) {
	t.Parallel()

	skipUnauthenticated(t)

	t.Run("default", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)

		config := fmt.Sprintf(`
resource "github_actions_variable" "test" {
  repository    = "%s"
  variable_name = "TEST"
  value         = "%%s"
}
`, repo.GetName())

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, "my-value"),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_actions_variable.test", tfjsonpath.New("repository_id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_actions_variable.test", tfjsonpath.New("created_at"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_actions_variable.test", tfjsonpath.New("updated_at"), knownvalue.NotNull()),
					},
				},
				{
					Config: fmt.Sprintf(config, "my-value-2"),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_actions_variable.test", plancheck.ResourceActionUpdate),
						},
					},
				},
				{
					ResourceName:      "github_actions_variable.test",
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})

	t.Run("updates_renamed_repo", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)
		newRepoName := fmt.Sprintf("%s-updated", repo.GetName())

		config := `
resource "github_actions_variable" "test" {
  repository    = "%s"
  variable_name = "TEST"
  value         = "my-value"
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
							plancheck.ExpectResourceAction("github_actions_variable.test", plancheck.ResourceActionUpdate),
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
resource "github_actions_variable" "test" {
  repository    = "%s"
  variable_name = "TEST"
  value         = "my-value"
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
							plancheck.ExpectResourceAction("github_actions_variable.test", plancheck.ResourceActionReplace),
						},
					},
				},
			},
		})
	})

	t.Run("errors_if_variable_already_exists", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)
		varName := "TEST"

		mustCreateTestRepositoryVariable(t, repo, &varName, nil)

		config := fmt.Sprintf(`
resource "github_actions_variable" "test" {
  repository    = "%s"
  variable_name = "%s"
  value         = "my-value"
}
`, repo.GetName(), varName)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      config,
					ExpectError: regexp.MustCompile(`Variable already exists`),
				},
			},
		})
	})
}
