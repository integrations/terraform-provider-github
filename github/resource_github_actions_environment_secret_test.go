package github

import (
	"encoding/base64"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubActionsEnvironmentSecret(t *testing.T) {
	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

	t.Run("creates and updates secrets without error", func(t *testing.T) {
		secretValue := base64.StdEncoding.EncodeToString([]byte("super_secret_value"))
		updatedSecretValue := base64.StdEncoding.EncodeToString([]byte("updated_super_secret_value"))

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
			  name = "tf-acc-test-%s"
			}

			resource "github_repository_environment" "test" {
			  repository       = github_repository.test.name
			  environment      = "environment / test"
			}

			resource "github_actions_environment_secret" "plaintext_secret" {
			  repository       = github_repository.test.name
			  environment      = github_repository_environment.test.environment
			  secret_name      = "test_plaintext_secret_name"
			  plaintext_value  = "%s"
			}

			resource "github_actions_environment_secret" "encrypted_secret" {
			  repository       = github_repository.test.name
			  environment      = github_repository_environment.test.environment
			  secret_name      = "test_encrypted_secret_name"
			  encrypted_value  = "%s"
			}
		`, randomID, secretValue, secretValue)

		checks := map[string]resource.TestCheckFunc{
			"before": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_actions_environment_secret.plaintext_secret", "plaintext_value",
					secretValue,
				),
				resource.TestCheckResourceAttr(
					"github_actions_environment_secret.encrypted_secret", "encrypted_value",
					secretValue,
				),
				resource.TestCheckResourceAttrSet(
					"github_actions_environment_secret.plaintext_secret", "created_at",
				),
				resource.TestCheckResourceAttrSet(
					"github_actions_environment_secret.plaintext_secret", "updated_at",
				),
			),
			"after": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_actions_environment_secret.plaintext_secret", "plaintext_value",
					updatedSecretValue,
				),
				resource.TestCheckResourceAttr(
					"github_actions_environment_secret.encrypted_secret", "encrypted_value",
					updatedSecretValue,
				),
				resource.TestCheckResourceAttrSet(
					"github_actions_environment_secret.plaintext_secret", "created_at",
				),
				resource.TestCheckResourceAttrSet(
					"github_actions_environment_secret.plaintext_secret", "updated_at",
				),
			),
		}

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: config,
						Check:  checks["before"],
					},
					{
						Config: strings.Replace(config,
							secretValue,
							updatedSecretValue, 2),
						Check: checks["after"],
					},
				},
			})
		}

		t.Run("with an anonymous account", func(t *testing.T) {
			t.Skip("anonymous account not supported for this operation")
		})

		t.Run("with an individual account", func(t *testing.T) {
			testCase(t, individual)
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})
	})

	t.Run("deletes secrets without error", func(t *testing.T) {
		secretValue := base64.StdEncoding.EncodeToString([]byte("super_secret_value"))

		config := fmt.Sprintf(`
				resource "github_repository" "test" {
					name = "tf-acc-test-%s"
				}

				resource "github_repository_environment" "test" {
					repository       = github_repository.test.name
					environment      = "environment / test"
				}

				resource "github_actions_environment_secret" "plaintext_secret" {
					repository       = github_repository.test.name
					environment      = github_repository_environment.test.environment
					secret_name      = "test_plaintext_secret_name"
					plaintext_value  = "%s"
				}

				resource "github_actions_environment_secret" "encrypted_secret" {
					repository       = github_repository.test.name
					environment      = github_repository_environment.test.environment
					secret_name      = "test_encrypted_secret_name"
					encrypted_value  = "%s"
				}
			`, randomID, secretValue, secretValue)

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config:  config,
						Destroy: true,
					},
				},
			})
		}

		t.Run("with an anonymous account", func(t *testing.T) {
			t.Skip("anonymous account not supported for this operation")
		})

		t.Run("with an individual account", func(t *testing.T) {
			testCase(t, individual)
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})
	})
}

func TestAccGithubActionsEnvironmentSecretIgnoreChanges(t *testing.T) {
	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

	t.Run("creates environment secrets using lifecycle ignore_changes", func(t *testing.T) {
		secretValue := base64.StdEncoding.EncodeToString([]byte("super_secret_value"))
		modifiedSecretValue := base64.StdEncoding.EncodeToString([]byte("a_modified_super_secret_value"))

		configFmtStr := `
			resource "github_repository" "test" {
				name = "tf-acc-test-%s"

				# TODO: provider appears to have issues destroying repositories while running the tests.
				#
				# Even with Organization Admin an error is seen:
				# Error: DELETE https://api.<cut>/tf-acc-test-<id>: "403 Must have admin rights to Repository. []"
				#
				# Workaround to using 'archive_on_destroy' instead.
				archive_on_destroy = true

				visibility = "private"
			}

			resource "github_repository_environment" "test" {
				repository       = github_repository.test.name
				environment      = "environment / test"
			}

			resource "github_actions_environment_secret" "plaintext_secret" {
				repository       = github_repository.test.name
				environment      = github_repository_environment.test.environment
				secret_name      = "test_plaintext_secret_name"
				plaintext_value  = "%s"

				lifecycle {
					ignore_changes = [plaintext_value]
				}
			}

			resource "github_actions_environment_secret" "encrypted_secret" {
				repository       = github_repository.test.name
				environment      = github_repository_environment.test.environment
				secret_name      = "test_encrypted_secret_name"
				encrypted_value  = "%s"

				lifecycle {
					ignore_changes = [encrypted_value]
				}
			}
		`

		checks := map[string]resource.TestCheckFunc{
			"before": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_actions_environment_secret.plaintext_secret", "plaintext_value",
					secretValue,
				),
				resource.TestCheckResourceAttr(
					"github_actions_environment_secret.encrypted_secret", "encrypted_value",
					secretValue,
				),
				resource.TestCheckResourceAttrSet(
					"github_actions_environment_secret.plaintext_secret", "created_at",
				),
				resource.TestCheckResourceAttrSet(
					"github_actions_environment_secret.plaintext_secret", "updated_at",
				),
			),
			"after": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_actions_environment_secret.plaintext_secret", "plaintext_value",
					secretValue,
				),
				resource.TestCheckResourceAttr(
					"github_actions_environment_secret.encrypted_secret", "encrypted_value",
					secretValue,
				),
				resource.TestCheckResourceAttrSet(
					"github_actions_environment_secret.plaintext_secret", "created_at",
				),
				resource.TestCheckResourceAttrSet(
					"github_actions_environment_secret.plaintext_secret", "updated_at",
				),
			),
		}

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: fmt.Sprintf(configFmtStr, randomID, secretValue, secretValue),
						Check:  checks["before"],
					},
					{
						Config: fmt.Sprintf(configFmtStr, randomID, secretValue, secretValue),
						Check:  checks["after"],
					},
					{
						// In this case the values change in the config, but the lifecycle ignore_changes should
						// not cause the actual values to be updated. This would also be the case when a secret
						// is externally modified (when what is in state does not match what is given).
						Config: fmt.Sprintf(configFmtStr, randomID, modifiedSecretValue, modifiedSecretValue),
						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckResourceAttr(
								"github_actions_environment_secret.plaintext_secret", "plaintext_value",
								secretValue, // Should still have the original value in state.
							),
							resource.TestCheckResourceAttr(
								"github_actions_environment_secret.encrypted_secret", "encrypted_value",
								secretValue, // Should still have the original value in state.
							),
						),
					},
				},
			})
		}

		t.Run("with an anonymous account", func(t *testing.T) {
			t.Skip("anonymous account not supported for this operation")
		})

		t.Run("with an individual account", func(t *testing.T) {
			testCase(t, individual)
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})
	})
}
