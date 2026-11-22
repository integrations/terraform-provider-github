package github

import (
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccGithubActionsEnvironmentSecret(t *testing.T) {
	t.Parallel()

	skipUnauthenticated(t)

	t.Run("with_value", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)
		env := mustCreateTestRepositoryEnvironment(t, repo)

		config := fmt.Sprintf(`
resource "github_actions_environment_secret" "test" {
  repository  = "%s"
  environment = "%s"
  secret_name = "TEST"
  value       = "%%s"
}
`, repo.GetName(), env.GetName())

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, "super_secret_value"),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_actions_environment_secret.test", tfjsonpath.New("repository_id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_actions_environment_secret.test", tfjsonpath.New("key_id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_actions_environment_secret.test", tfjsonpath.New("created_at"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_actions_environment_secret.test", tfjsonpath.New("updated_at"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_actions_environment_secret.test", tfjsonpath.New("remote_updated_at"), knownvalue.NotNull()),
					},
				},
				{
					Config: fmt.Sprintf(config, "super_secret_value_2"),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_actions_environment_secret.test", plancheck.ResourceActionUpdate),
						},
					},
				},
				{
					ResourceName:            "github_actions_environment_secret.test",
					ImportState:             true,
					ImportStateVerify:       true,
					ImportStateVerifyIgnore: []string{"key_id", "value"},
				},
			},
		})
	})

	t.Run("with_value_encrypted", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)
		env := mustCreateTestRepositoryEnvironment(t, repo)
		key := mustGetTestRepositoryEnvironmentPublicKey(t, repo, env)

		config := fmt.Sprintf(`
resource "github_actions_environment_secret" "test" {
  repository      = "%s"
  environment     = "%s"
  secret_name     = "TEST"
	key_id          = "%s"
  value_encrypted = "%%s"
}
`, repo.GetName(), env.GetName(), key.GetKeyID())

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, base64.StdEncoding.EncodeToString([]byte("super_secret_value"))),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_actions_environment_secret.test", tfjsonpath.New("repository_id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_actions_environment_secret.test", tfjsonpath.New("created_at"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_actions_environment_secret.test", tfjsonpath.New("updated_at"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_actions_environment_secret.test", tfjsonpath.New("remote_updated_at"), knownvalue.NotNull()),
					},
				},
				{
					Config: fmt.Sprintf(config, base64.StdEncoding.EncodeToString([]byte("super_secret_value_2"))),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_actions_environment_secret.test", plancheck.ResourceActionUpdate),
						},
					},
				},
				{
					ResourceName:            "github_actions_environment_secret.test",
					ImportState:             true,
					ImportStateVerify:       true,
					ImportStateVerifyIgnore: []string{"key_id", "value_encrypted"},
				},
			},
		})
	})

	t.Run("with_plaintext_value", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)
		env := mustCreateTestRepositoryEnvironment(t, repo)

		config := fmt.Sprintf(`
resource "github_actions_environment_secret" "test" {
  repository      = "%s"
  environment     = "%s"
  secret_name     = "TEST"
  plaintext_value = "%%s"
}
`, repo.GetName(), env.GetName())

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, "super_secret_value"),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_actions_environment_secret.test", tfjsonpath.New("repository_id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_actions_environment_secret.test", tfjsonpath.New("key_id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_actions_environment_secret.test", tfjsonpath.New("created_at"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_actions_environment_secret.test", tfjsonpath.New("updated_at"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_actions_environment_secret.test", tfjsonpath.New("remote_updated_at"), knownvalue.NotNull()),
					},
				},
				{
					Config: fmt.Sprintf(config, "super_secret_value_2"),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_actions_environment_secret.test", plancheck.ResourceActionUpdate),
						},
					},
				},
				{
					ResourceName:            "github_actions_environment_secret.test",
					ImportState:             true,
					ImportStateVerify:       true,
					ImportStateVerifyIgnore: []string{"key_id", "plaintext_value"},
				},
			},
		})
	})

	t.Run("with_encrypted_value", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)
		env := mustCreateTestRepositoryEnvironment(t, repo)

		config := fmt.Sprintf(`
resource "github_actions_environment_secret" "test" {
  repository      = "%s"
  environment     = "%s"
  secret_name     = "TEST"
  encrypted_value = "%%s"
}
`, repo.GetName(), env.GetName())

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, base64.StdEncoding.EncodeToString([]byte("super_secret_value"))),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_actions_environment_secret.test", tfjsonpath.New("repository_id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_actions_environment_secret.test", tfjsonpath.New("key_id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_actions_environment_secret.test", tfjsonpath.New("created_at"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_actions_environment_secret.test", tfjsonpath.New("updated_at"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_actions_environment_secret.test", tfjsonpath.New("remote_updated_at"), knownvalue.NotNull()),
					},
				},
				{
					Config: fmt.Sprintf(config, base64.StdEncoding.EncodeToString([]byte("super_secret_value_2"))),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_actions_environment_secret.test", plancheck.ResourceActionUpdate),
						},
					},
				},
				{
					ResourceName:            "github_actions_environment_secret.test",
					ImportState:             true,
					ImportStateVerify:       true,
					ImportStateVerifyIgnore: []string{"key_id", "encrypted_value"},
				},
			},
		})
	})

	t.Run("with_env_name_id_separator_character", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)
		env := mustCreateTestRepositoryEnvironment(t, repo, withTestCreateName("env:test"))

		config := fmt.Sprintf(`
resource "github_actions_environment_secret" "test" {
  repository  = "%s"
  environment = "%s"
  secret_name = "TEST"
  value       = "super_secret_value"
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

	t.Run("updates_on_drift", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)
		env := mustCreateTestRepositoryEnvironment(t, repo)
		secretName := "TEST"

		config := fmt.Sprintf(`
resource "github_actions_environment_secret" "test" {
  repository  = "%s"
  environment = "%s"
  secret_name = "%s"
  value       = "super_secret_value"
}
`, repo.GetName(), env.GetName(), secretName)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
				},
				{
					PreConfig: func() {
						mustUpdateTestRepositoryEnvironmentSecret(t, repo, env, secretName, "super_secret_value_2")
					},
					Config: config,
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_actions_environment_secret.test", plancheck.ResourceActionUpdate),
						},
					},
				},
			},
		})
	})

	t.Run("lifecycle_ignore_suppresses_drift", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)
		env := mustCreateTestRepositoryEnvironment(t, repo)
		secretName := "TEST"

		config := fmt.Sprintf(`
resource "github_actions_environment_secret" "test" {
  repository  = "%s"
  environment = "%s"
  secret_name = "%s"
  value       = "super_secret_value"

  lifecycle {
    ignore_changes = [updated_at]
  }
}
`, repo.GetName(), env.GetName(), secretName)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
				},
				{
					PreConfig: func() {
						mustUpdateTestRepositoryEnvironmentSecret(t, repo, env, secretName, "super_secret_value_2")
					},
					Config: config,
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_actions_environment_secret.test", plancheck.ResourceActionNoop),
						},
					},
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
resource "github_actions_environment_secret" "test" {
  repository  = "%%s"
  environment = "%s"
  secret_name = "TEST"
  value       = "super_secret_value"
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
							plancheck.ExpectResourceAction("github_actions_environment_secret.test", plancheck.ResourceActionUpdate),
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
resource "github_actions_environment_secret" "test" {
  repository  = "%%s"
  environment = "%s"
  secret_name = "TEST"
  value       = "super_secret_value"
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
							plancheck.ExpectResourceAction("github_actions_environment_secret.test", plancheck.ResourceActionReplace),
						},
					},
				},
			},
		})
	})
}
