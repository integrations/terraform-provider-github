package github

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccGithubActionsEnvironmentSecret(t *testing.T) {

	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

	t.Run("creates and updates secrets without error", func(t *testing.T) {

		secretValue := "super_secret_value"
		updatedSecretValue := "updated_super_secret_value"

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
			  name = "tf-acc-test-%s"
			}

			resource "github_repository_environment" "test" {
			  repository       = github_repository.test.name
			  environment      = "test_environment_name"
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

		secretValue := "super_secret_value"

		config := fmt.Sprintf(`
				resource "github_repository" "test" {
					name = "tf-acc-test-%s"
				}

				resource "github_repository_environment" "test" {
					repository       = github_repository.test.name
					environment      = "test_environment_name"
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
