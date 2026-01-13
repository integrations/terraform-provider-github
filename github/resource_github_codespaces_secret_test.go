package github

import (
	"encoding/base64"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubCodespacesSecret(t *testing.T) {
	t.Run("reads a repository public key without error", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-codespaces-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`

			resource "github_repository" "test" {
			  name = "%s"
			}

			data "github_codespaces_public_key" "test_pk" {
			  repository = github_repository.test.name
			}

		`, repoName)

		check := resource.ComposeAggregateTestCheckFunc(
			resource.TestCheckResourceAttrSet(
				"data.github_codespaces_public_key.test_pk", "key_id",
			),
			resource.TestCheckResourceAttrSet(
				"data.github_codespaces_public_key.test_pk", "key",
			),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  check,
				},
			},
		})
	})

	t.Run("creates and updates secrets without error", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-codespaces-%s", testResourcePrefix, randomID)
		secretValue := base64.StdEncoding.EncodeToString([]byte("super_secret_value"))
		updatedSecretValue := base64.StdEncoding.EncodeToString([]byte("updated_super_secret_value"))

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
			  name = "%s"
			}

			resource "github_codespaces_secret" "plaintext_secret" {
			  repository       = github_repository.test.name
			  secret_name      = "test_plaintext_secret"
			  plaintext_value  = "%s"
			}

			resource "github_codespaces_secret" "encrypted_secret" {
			  repository       = github_repository.test.name
			  secret_name      = "test_encrypted_secret"
			  encrypted_value  = "%s"
			}
			`, repoName, secretValue, secretValue)

		checks := map[string]resource.TestCheckFunc{
			"before": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_codespaces_secret.plaintext_secret", "plaintext_value",
					secretValue,
				),
				resource.TestCheckResourceAttr(
					"github_codespaces_secret.encrypted_secret", "encrypted_value",
					secretValue,
				),
				resource.TestCheckResourceAttrSet(
					"github_codespaces_secret.plaintext_secret", "created_at",
				),
				resource.TestCheckResourceAttrSet(
					"github_codespaces_secret.plaintext_secret", "updated_at",
				),
			),
			"after": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_codespaces_secret.plaintext_secret", "plaintext_value",
					updatedSecretValue,
				),
				resource.TestCheckResourceAttr(
					"github_codespaces_secret.encrypted_secret", "encrypted_value",
					updatedSecretValue,
				),
				resource.TestCheckResourceAttrSet(
					"github_codespaces_secret.plaintext_secret", "created_at",
				),
				resource.TestCheckResourceAttrSet(
					"github_codespaces_secret.plaintext_secret", "updated_at",
				),
			),
		}

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
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
	})

	t.Run("creates and updates repository name without error", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-codespaces-%s", testResourcePrefix, randomID)
		updatedRepoName := fmt.Sprintf("%s-updated", repoName)
		secretValue := base64.StdEncoding.EncodeToString([]byte("super_secret_value"))

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
			  name = "%s"
			}
			resource "github_codespaces_secret" "plaintext_secret" {
			  repository       = github_repository.test.name
			  secret_name      = "test_plaintext_secret"
			  plaintext_value  = "%s"
			}
			resource "github_codespaces_secret" "encrypted_secret" {
			  repository       = github_repository.test.name
			  secret_name      = "test_encrypted_secret"
			  encrypted_value  = "%s"
			}
			`, repoName, secretValue, secretValue)

		checks := map[string]resource.TestCheckFunc{
			"before": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_codespaces_secret.plaintext_secret", "repository",
					repoName,
				),
				resource.TestCheckResourceAttr(
					"github_codespaces_secret.plaintext_secret", "plaintext_value",
					secretValue,
				),
				resource.TestCheckResourceAttr(
					"github_codespaces_secret.encrypted_secret", "encrypted_value",
					secretValue,
				),
				resource.TestCheckResourceAttrSet(
					"github_codespaces_secret.plaintext_secret", "created_at",
				),
				resource.TestCheckResourceAttrSet(
					"github_codespaces_secret.plaintext_secret", "updated_at",
				),
			),
			"after": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_codespaces_secret.plaintext_secret", "repository",
					updatedRepoName,
				),
				resource.TestCheckResourceAttr(
					"github_codespaces_secret.plaintext_secret", "plaintext_value",
					secretValue,
				),
				resource.TestCheckResourceAttr(
					"github_codespaces_secret.encrypted_secret", "encrypted_value",
					secretValue,
				),
				resource.TestCheckResourceAttrSet(
					"github_codespaces_secret.plaintext_secret", "created_at",
				),
				resource.TestCheckResourceAttrSet(
					"github_codespaces_secret.plaintext_secret", "updated_at",
				),
			),
		}

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  checks["before"],
				},
				{
					Config: strings.Replace(config,
						repoName,
						updatedRepoName, 2),
					Check: checks["after"],
				},
			},
		})
	})

	t.Run("deletes secrets without error", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-codespaces-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`
				resource "github_repository" "test" {
					name = "%s"
				}

				resource "github_codespaces_secret" "plaintext_secret" {
					repository 	= github_repository.test.name
					secret_name	= "test_plaintext_secret"
				}

				resource "github_codespaces_secret" "encrypted_secret" {
					repository 	= github_repository.test.name
					secret_name	= "test_encrypted_secret"
				}
			`, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:  config,
					Destroy: true,
				},
			},
		})
	})
}
