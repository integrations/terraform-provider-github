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

func TestAccGithubActionsEnvironmentVariable(t *testing.T) {
	t.Parallel()

	skipUnauthenticated(t)

	t.Run("default", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)
		env := mustCreateTestRepositoryEnvironment(t, repo)

		config := fmt.Sprintf(`
resource "github_actions_environment_variable" "test" {
  repository    = "%s"
  environment   = "%s"
  variable_name = "TEST"
  value         = "%%s"
}
`, repo.GetName(), env.GetName())

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, "my-value"),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_actions_environment_variable.test", tfjsonpath.New("repository_id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_actions_environment_variable.test", tfjsonpath.New("created_at"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_actions_environment_variable.test", tfjsonpath.New("updated_at"), knownvalue.NotNull()),
					},
				},
				{
					Config: fmt.Sprintf(config, "my-value-2"),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_actions_environment_variable.test", plancheck.ResourceActionUpdate),
						},
					},
				},
				{
					ResourceName:      "github_actions_environment_variable.test",
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})

	t.Run("with_env_name_id_separator_character", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)
		env := mustCreateTestRepositoryEnvironment(t, repo, withTestCreateName("env:test"))

		config := fmt.Sprintf(`
resource "github_actions_environment_variable" "test" {
  repository    = "%s"
  environment   = "%s"
  variable_name = "TEST"
  value         = "my-value"
}
`, repo.GetName(), env.GetName())

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
				},
			},
		})
	})

	t.Run("updates_renamed_repo", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)
		env := mustCreateTestRepositoryEnvironment(t, repo)
		newRepoName := fmt.Sprintf("%s-updated", repo.GetName())

		config := fmt.Sprintf(`
resource "github_actions_environment_variable" "test" {
  repository    = "%%s"
  environment   = "%s"
  variable_name = "TEST"
  value         = "my-value"
}
`, env.GetName())

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
							plancheck.ExpectResourceAction("github_actions_environment_variable.test", plancheck.ResourceActionUpdate),
						},
					},
				},
			},
		})
	})

	t.Run("recreates_changed_repo", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)
		env := mustCreateTestRepositoryEnvironment(t, repo)
		repo2 := mustCreateTestRepository(t)
		_ = mustCreateTestRepositoryEnvironment(t, repo2, withTestCreateName(env.GetName()))

		config := fmt.Sprintf(`
resource "github_actions_environment_variable" "test" {
  repository    = "%%s"
  environment   = "%s"
  variable_name = "TEST"
  value         = "my-value"
}
`, env.GetName())

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
							plancheck.ExpectResourceAction("github_actions_environment_variable.test", plancheck.ResourceActionReplace),
						},
					},
				},
			},
		})
	})

	t.Run("errors_if_variable_already_exists", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)
		env := mustCreateTestRepositoryEnvironment(t, repo)
		varName := "TEST"

		mustCreateTestRepositoryEnvironmentVariable(t, repo, env, &varName, nil)

		config := fmt.Sprintf(`
resource "github_actions_environment_variable" "test" {
  repository    = "%s"
  environment   = "%s"
  variable_name = "%s"
  value         = "my-value"
}
`, repo.GetName(), env.GetName(), varName)

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
