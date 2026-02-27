package github

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/url"
	"testing"

	"github.com/google/go-github/v83/github"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccGithubActionsEnvironmentSecret(t *testing.T) {
	t.Run("create_plaintext", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)
		envName := "test"
		secretName := "test"
		value := base64.StdEncoding.EncodeToString([]byte("super_secret_value"))

		config := fmt.Sprintf(`
resource "github_repository" "test" {
	name = "%s"
}

resource "github_repository_environment" "test" {
	repository  = github_repository.test.name
	environment = "%s"
}

resource "github_actions_environment_secret" "test" {
	repository  = github_repository.test.name
	environment = github_repository_environment.test.environment
	secret_name = "%s"
	value       = "%s"
}
`, repoName, envName, secretName, value)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttrPair("github_actions_environment_secret.test", "repository", "github_repository.test", "name"),
						resource.TestCheckResourceAttr("github_actions_environment_secret.test", "environment", envName),
						resource.TestCheckResourceAttr("github_actions_environment_secret.test", "secret_name", secretName),
						resource.TestCheckResourceAttr("github_actions_environment_secret.test", "value", value),
						resource.TestCheckNoResourceAttr("github_actions_environment_secret.test", "value_encrypted"),
						resource.TestCheckResourceAttrSet("github_actions_environment_secret.test", "key_id"),
						resource.TestCheckResourceAttrSet("github_actions_environment_secret.test", "created_at"),
						resource.TestCheckResourceAttrSet("github_actions_environment_secret.test", "updated_at"),
					),
				},
			},
		})
	})

	t.Run("create_with_env_name_id_separator_character", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)
		envName := "env:test"
		secretName := "test"
		value := base64.StdEncoding.EncodeToString([]byte("super_secret_value"))

		config := fmt.Sprintf(`
resource "github_repository" "test" {
	name = "%s"
}

resource "github_repository_environment" "test" {
	repository  = github_repository.test.name
	environment = "%s"
}

resource "github_actions_environment_secret" "test" {
	repository  = github_repository.test.name
	environment = github_repository_environment.test.environment
	secret_name = "%s"
	value       = "%s"
}
`, repoName, envName, secretName, value)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttrPair("github_actions_environment_secret.test", "repository", "github_repository.test", "name"),
						resource.TestCheckResourceAttr("github_actions_environment_secret.test", "environment", envName),
						resource.TestCheckResourceAttr("github_actions_environment_secret.test", "secret_name", secretName),
						resource.TestCheckResourceAttr("github_actions_environment_secret.test", "value", value),
						resource.TestCheckNoResourceAttr("github_actions_environment_secret.test", "value_encrypted"),
						resource.TestCheckResourceAttrSet("github_actions_environment_secret.test", "key_id"),
						resource.TestCheckResourceAttrSet("github_actions_environment_secret.test", "created_at"),
						resource.TestCheckResourceAttrSet("github_actions_environment_secret.test", "updated_at"),
					),
				},
			},
		})
	})

	t.Run("create_update_plaintext", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)
		envName := "test"
		secretName := "test"
		value := base64.StdEncoding.EncodeToString([]byte("super_secret_value"))
		updatedValue := base64.StdEncoding.EncodeToString([]byte("updated_super_secret_value"))

		config := `
resource "github_repository" "test" {
	name = "%s"
}

resource "github_repository_environment" "test" {
	repository  = github_repository.test.name
	environment = "%s"
}

resource "github_actions_environment_secret" "test" {
	repository  = github_repository.test.name
	environment = github_repository_environment.test.environment
	secret_name = "%s"
	value       = "%s"
}
`

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, repoName, envName, secretName, value),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttrPair("github_actions_environment_secret.test", "repository", "github_repository.test", "name"),
						resource.TestCheckResourceAttr("github_actions_environment_secret.test", "environment", envName),
						resource.TestCheckResourceAttr("github_actions_environment_secret.test", "secret_name", secretName),
						resource.TestCheckResourceAttr("github_actions_environment_secret.test", "value", value),
						resource.TestCheckNoResourceAttr("github_actions_environment_secret.test", "value_encrypted"),
						resource.TestCheckResourceAttrSet("github_actions_environment_secret.test", "key_id"),
						resource.TestCheckResourceAttrSet("github_actions_environment_secret.test", "created_at"),
						resource.TestCheckResourceAttrSet("github_actions_environment_secret.test", "updated_at"),
					),
				},
				{
					Config: fmt.Sprintf(config, repoName, envName, secretName, updatedValue),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttrPair("github_actions_environment_secret.test", "repository", "github_repository.test", "name"),
						resource.TestCheckResourceAttr("github_actions_environment_secret.test", "environment", envName),
						resource.TestCheckResourceAttr("github_actions_environment_secret.test", "secret_name", secretName),
						resource.TestCheckResourceAttr("github_actions_environment_secret.test", "value", updatedValue),
						resource.TestCheckNoResourceAttr("github_actions_environment_secret.test", "value_encrypted"),
						resource.TestCheckResourceAttrSet("github_actions_environment_secret.test", "key_id"),
						resource.TestCheckResourceAttrSet("github_actions_environment_secret.test", "created_at"),
						resource.TestCheckResourceAttrSet("github_actions_environment_secret.test", "updated_at"),
					),
				},
			},
		})
	})

	t.Run("create_update_encrypted", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)
		envName := "test"
		secretName := "test"
		value := base64.StdEncoding.EncodeToString([]byte("super_secret_value"))
		updatedValue := base64.StdEncoding.EncodeToString([]byte("updated_super_secret_value"))

		config := `
resource "github_repository" "test" {
	name = "%s"
}

resource "github_repository_environment" "test" {
	repository  = github_repository.test.name
	environment = "%s"
}

resource "github_actions_environment_secret" "test" {
	repository      = github_repository.test.name
	environment     = github_repository_environment.test.environment
	secret_name     = "%s"
	value_encrypted = "%s"
}
`

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, repoName, envName, secretName, value),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttrPair("github_actions_environment_secret.test", "repository", "github_repository.test", "name"),
						resource.TestCheckResourceAttr("github_actions_environment_secret.test", "environment", envName),
						resource.TestCheckResourceAttr("github_actions_environment_secret.test", "secret_name", secretName),
						resource.TestCheckNoResourceAttr("github_actions_environment_secret.test", "value"),
						resource.TestCheckResourceAttr("github_actions_environment_secret.test", "value_encrypted", value),
						resource.TestCheckResourceAttrSet("github_actions_environment_secret.test", "key_id"),
						resource.TestCheckResourceAttrSet("github_actions_environment_secret.test", "created_at"),
						resource.TestCheckResourceAttrSet("github_actions_environment_secret.test", "updated_at"),
					),
				},
				{
					Config: fmt.Sprintf(config, repoName, envName, secretName, updatedValue),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttrPair("github_actions_environment_secret.test", "repository", "github_repository.test", "name"),
						resource.TestCheckResourceAttr("github_actions_environment_secret.test", "environment", envName),
						resource.TestCheckResourceAttr("github_actions_environment_secret.test", "secret_name", secretName),
						resource.TestCheckNoResourceAttr("github_actions_environment_secret.test", "value"),
						resource.TestCheckResourceAttr("github_actions_environment_secret.test", "value_encrypted", updatedValue),
						resource.TestCheckResourceAttrSet("github_actions_environment_secret.test", "key_id"),
						resource.TestCheckResourceAttrSet("github_actions_environment_secret.test", "created_at"),
						resource.TestCheckResourceAttrSet("github_actions_environment_secret.test", "updated_at"),
					),
				},
			},
		})
	})

	t.Run("create_update_encrypted_with_key", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)
		envName := "test"
		secretName := "test"
		value := base64.StdEncoding.EncodeToString([]byte("super_secret_value"))
		updatedValue := base64.StdEncoding.EncodeToString([]byte("updated_super_secret_value"))

		config := `
resource "github_repository" "test" {
	name = "%s"
}

resource "github_repository_environment" "test" {
	repository  = github_repository.test.name
	environment = "%s"
}

data "github_actions_environment_public_key" "test" {
	repository  = github_repository.test.name
	environment = github_repository_environment.test.environment
}

resource "github_actions_environment_secret" "test" {
	repository      = github_repository.test.name
	environment     = github_repository_environment.test.environment
	key_id          = data.github_actions_environment_public_key.test.key_id
	secret_name     = "%s"
	value_encrypted = "%s"
}
`

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, repoName, envName, secretName, value),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttrPair("github_actions_environment_secret.test", "repository", "github_repository.test", "name"),
						resource.TestCheckResourceAttr("github_actions_environment_secret.test", "environment", envName),
						resource.TestCheckResourceAttr("github_actions_environment_secret.test", "secret_name", secretName),
						resource.TestCheckNoResourceAttr("github_actions_environment_secret.test", "value"),
						resource.TestCheckResourceAttr("github_actions_environment_secret.test", "value_encrypted", value),
						resource.TestCheckResourceAttrSet("github_actions_environment_secret.test", "key_id"),
						resource.TestCheckResourceAttrSet("github_actions_environment_secret.test", "created_at"),
						resource.TestCheckResourceAttrSet("github_actions_environment_secret.test", "updated_at"),
					),
				},
				{
					Config: fmt.Sprintf(config, repoName, envName, secretName, updatedValue),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttrPair("github_actions_environment_secret.test", "repository", "github_repository.test", "name"),
						resource.TestCheckResourceAttr("github_actions_environment_secret.test", "environment", envName),
						resource.TestCheckResourceAttr("github_actions_environment_secret.test", "secret_name", secretName),
						resource.TestCheckNoResourceAttr("github_actions_environment_secret.test", "value"),
						resource.TestCheckResourceAttr("github_actions_environment_secret.test", "value_encrypted", updatedValue),
						resource.TestCheckResourceAttrSet("github_actions_environment_secret.test", "key_id"),
						resource.TestCheckResourceAttrSet("github_actions_environment_secret.test", "created_at"),
						resource.TestCheckResourceAttrSet("github_actions_environment_secret.test", "updated_at"),
					),
				},
			},
		})
	})

	t.Run("update_on_drift", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)
		envName := "test"
		secretName := "test"

		config := fmt.Sprintf(`
resource "github_repository" "test" {
	name = "%s"
}

resource "github_repository_environment" "test" {
	repository  = github_repository.test.name
	environment = "%s"
}

resource "github_actions_environment_secret" "test" {
	repository  = github_repository.test.name
	environment = github_repository_environment.test.environment
	secret_name = "%s"
	value       = "test"
}
`, repoName, envName, secretName)

		var beforeCreatedAt string
		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttrSet("github_actions_environment_secret.test", "created_at"),
						resource.TestCheckResourceAttrSet("github_actions_environment_secret.test", "updated_at"),
						resource.TestCheckResourceAttrSet("github_actions_environment_secret.test", "remote_updated_at"),
						func(s *terraform.State) error {
							beforeCreatedAt = s.RootModule().Resources["github_actions_environment_secret.test"].Primary.Attributes["created_at"]
							return nil
						},
					),
				},
				{
					PreConfig: func() {
						meta, err := getTestMeta()
						if err != nil {
							t.Fatal(err.Error())
						}
						client := meta.v3client
						owner := meta.name
						ctx := context.Background()

						escapedEnvName := url.PathEscape(envName)

						repo, _, err := client.Repositories.Get(ctx, owner, repoName)
						if err != nil {
							t.Fatal(err.Error())
						}
						repoID := int(repo.GetID())

						keyID, _, err := getEnvironmentPublicKeyDetails(ctx, meta, repoID, escapedEnvName)
						if err != nil {
							t.Fatal(err.Error())
						}

						_, err = client.Actions.CreateOrUpdateEnvSecret(ctx, repoID, escapedEnvName, &github.EncryptedSecret{
							Name:           secretName,
							EncryptedValue: base64.StdEncoding.EncodeToString([]byte("updated_super_secret_value")),
							KeyID:          keyID,
						})
						if err != nil {
							t.Fatal(err.Error())
						}
					},
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttrSet("github_actions_environment_secret.test", "created_at"),
						resource.TestCheckResourceAttrSet("github_actions_environment_secret.test", "updated_at"),
						resource.TestCheckResourceAttrSet("github_actions_environment_secret.test", "remote_updated_at"),
						func(s *terraform.State) error {
							afterCreatedAt := s.RootModule().Resources["github_actions_environment_secret.test"].Primary.Attributes["created_at"]

							if afterCreatedAt != beforeCreatedAt {
								return fmt.Errorf("expected resource to be updated, but created_at has been modified: %s", beforeCreatedAt)
							}
							return nil
						},
					),
				},
			},
		})
	})

	t.Run("lifecycle_can_ignore_drift", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)
		envName := "test"
		secretName := "test"

		config := fmt.Sprintf(`
resource "github_repository" "test" {
	name = "%s"
}

resource "github_repository_environment" "test" {
	repository  = github_repository.test.name
	environment = "%s"
}

resource "github_actions_environment_secret" "test" {
	repository  = github_repository.test.name
	environment = github_repository_environment.test.environment
	secret_name = "%s"
	value       = "test"

	lifecycle {
		ignore_changes = [remote_updated_at]
	}
}
`, repoName, envName, secretName)

		var beforeUpdatedAt string
		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttrSet("github_actions_environment_secret.test", "created_at"),
						resource.TestCheckResourceAttrSet("github_actions_environment_secret.test", "updated_at"),
						resource.TestCheckResourceAttrSet("github_actions_environment_secret.test", "remote_updated_at"),
						func(s *terraform.State) error {
							beforeUpdatedAt = s.RootModule().Resources["github_actions_environment_secret.test"].Primary.Attributes["updated_at"]
							return nil
						},
					),
				},
				{
					PreConfig: func() {
						meta, err := getTestMeta()
						if err != nil {
							t.Fatal(err.Error())
						}
						client := meta.v3client
						owner := meta.name
						ctx := context.Background()

						escapedEnvName := url.PathEscape(envName)

						repo, _, err := client.Repositories.Get(ctx, owner, repoName)
						if err != nil {
							t.Fatal(err.Error())
						}
						repoID := int(repo.GetID())

						keyID, _, err := getEnvironmentPublicKeyDetails(ctx, meta, repoID, escapedEnvName)
						if err != nil {
							t.Fatal(err.Error())
						}

						_, err = client.Actions.CreateOrUpdateEnvSecret(ctx, repoID, escapedEnvName, &github.EncryptedSecret{
							Name:           secretName,
							EncryptedValue: base64.StdEncoding.EncodeToString([]byte("updated_super_secret_value")),
							KeyID:          keyID,
						})
						if err != nil {
							t.Fatal(err.Error())
						}
					},
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttrSet("github_actions_environment_secret.test", "created_at"),
						resource.TestCheckResourceAttrSet("github_actions_environment_secret.test", "updated_at"),
						resource.TestCheckResourceAttrSet("github_actions_environment_secret.test", "remote_updated_at"),
						func(s *terraform.State) error {
							afterUpdatedAt := s.RootModule().Resources["github_actions_environment_secret.test"].Primary.Attributes["updated_at"]

							if afterUpdatedAt != beforeUpdatedAt {
								return fmt.Errorf("expected resource to ignore drift, but updated_at has been modified: %s", beforeUpdatedAt)
							}
							return nil
						},
					),
				},
			},
		})
	})

	t.Run("update_renamed_repo", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)
		updatedRepoName := fmt.Sprintf("%s%s-updated", testResourcePrefix, randomID)

		// TODO: Remove lifecycle ignore_changes block when repo rename is supported
		config := `
resource "github_repository" "test" {
	name = "%s"
}

resource "github_repository_environment" "test" {
	repository  = github_repository.test.name
	environment = "test"

	lifecycle {
		ignore_changes = all
	}
}

resource "github_actions_environment_secret" "test" {
	repository  = github_repository.test.name
	environment = github_repository_environment.test.environment
	secret_name = "test"
	value       = "test"
}
`

		var beforeCreatedAt string
		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, repoName),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttrSet("github_actions_environment_secret.test", "created_at"),
						resource.TestCheckResourceAttrSet("github_actions_environment_secret.test", "updated_at"),
						func(s *terraform.State) error {
							beforeCreatedAt = s.RootModule().Resources["github_actions_environment_secret.test"].Primary.Attributes["created_at"]
							return nil
						},
					),
				},
				{
					Config: fmt.Sprintf(config, updatedRepoName),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttrSet("github_actions_environment_secret.test", "created_at"),
						resource.TestCheckResourceAttrSet("github_actions_environment_secret.test", "updated_at"),
						func(s *terraform.State) error {
							afterCreatedAt := s.RootModule().Resources["github_actions_environment_secret.test"].Primary.Attributes["created_at"]

							if afterCreatedAt != beforeCreatedAt {
								return fmt.Errorf("expected resource to not be recreated, but created_at has been modified: %s", beforeCreatedAt)
							}
							return nil
						},
					),
				},
			},
		})
	})

	t.Run("recreate_changed_repo", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)
		repoName2 := fmt.Sprintf("%supdated-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
resource "github_repository" "test" {
	name = "%s"
}

resource "github_repository_environment" "test" {
	repository  = github_repository.test.name
	environment = "test"
}

resource "github_repository" "test2" {
	name = "%s"
}

resource "github_repository_environment" "test2" {
	repository  = github_repository.test2.name
	environment = "test"
}

resource "github_actions_environment_secret" "test" {
	repository  = github_repository.test.name
	environment = github_repository_environment.test.environment
	secret_name = "test"
	value       = "test"
}
`, repoName, repoName2)

		configUpdated := fmt.Sprintf(`
resource "github_repository" "test" {
	name = "%s"
}

resource "github_repository_environment" "test" {
	repository  = github_repository.test.name
	environment = "test"
}

resource "github_repository" "test2" {
	name = "%s"
}

resource "github_repository_environment" "test2" {
	repository  = github_repository.test2.name
	environment = "test"
}

resource "github_actions_environment_secret" "test" {
	repository  = github_repository.test2.name
	environment = github_repository_environment.test2.environment
	secret_name = "test"
	value       = "test"
}
`, repoName, repoName2)

		var beforeCreatedAt string
		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttrSet("github_actions_environment_secret.test", "created_at"),
						resource.TestCheckResourceAttrSet("github_actions_environment_secret.test", "updated_at"),
						func(s *terraform.State) error {
							beforeCreatedAt = s.RootModule().Resources["github_actions_environment_secret.test"].Primary.Attributes["created_at"]
							return nil
						},
					),
				},
				{
					Config: configUpdated,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttrSet("github_actions_environment_secret.test", "created_at"),
						resource.TestCheckResourceAttrSet("github_actions_environment_secret.test", "updated_at"),
						func(s *terraform.State) error {
							afterCreatedAt := s.RootModule().Resources["github_actions_environment_secret.test"].Primary.Attributes["created_at"]

							if afterCreatedAt == beforeCreatedAt {
								return fmt.Errorf("expected resource to be recreated, but created_at has not been modified: %s", beforeCreatedAt)
							}
							return nil
						},
					),
				},
			},
		})
	})

	t.Run("destroy", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
	resource "github_repository" "test" {
		name = "%s"
	}

	resource "github_repository_environment" "test" {
		repository  = github_repository.test.name
		environment = "test"
	}

	resource "github_actions_environment_secret" "test" {
		repository  = github_repository.test.name
		environment = github_repository_environment.test.environment
		secret_name = "test"
		value       = "test"
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
					Config:  config,
					Destroy: true,
				},
			},
		})
	})

	t.Run("import", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)
		envName := "test"
		secretName := "test"

		config := fmt.Sprintf(`
resource "github_repository" "test" {
	name = "%s"
}

resource "github_repository_environment" "test" {
	repository  = github_repository.test.name
	environment = "%s"
}

resource "github_actions_environment_secret" "test" {
	repository  = github_repository.test.name
	environment = github_repository_environment.test.environment
	secret_name = "%s"
	value       = "test"
}
`, repoName, envName, secretName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
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
}
