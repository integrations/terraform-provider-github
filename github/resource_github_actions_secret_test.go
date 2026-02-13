package github

import (
	"context"
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/google/go-github/v82/github"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccGithubActionsSecret(t *testing.T) {
	t.Run("create_plaintext", func(t *testing.T) {
		randomID := acctest.RandString(5)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)
		secretName := "test"
		value := base64.StdEncoding.EncodeToString([]byte("super_secret_value"))

		config := fmt.Sprintf(`
resource "github_repository" "test" {
	name = "%s"
}

resource "github_actions_secret" "test" {
	repository      = github_repository.test.name
	secret_name     = "%s"
	plaintext_value = "%s"
}
`, repoName, secretName, value)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttrPair("github_actions_secret.test", "repository", "github_repository.test", "name"),
						resource.TestCheckResourceAttr("github_actions_secret.test", "secret_name", secretName),
						resource.TestCheckResourceAttr("github_actions_secret.test", "plaintext_value", value),
						resource.TestCheckNoResourceAttr("github_actions_secret.test", "encrypted_value"),
						resource.TestCheckResourceAttrSet("github_actions_secret.test", "key_id"),
						resource.TestCheckResourceAttrSet("github_actions_secret.test", "created_at"),
						resource.TestCheckResourceAttrSet("github_actions_secret.test", "updated_at"),
					),
				},
			},
		})
	})

	t.Run("create_update_plaintext", func(t *testing.T) {
		randomID := acctest.RandString(5)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)
		secretName := "test"
		value := base64.StdEncoding.EncodeToString([]byte("super_secret_value"))
		updatedValue := base64.StdEncoding.EncodeToString([]byte("updated_super_secret_value"))

		config := `
resource "github_repository" "test" {
	name = "%s"
}

resource "github_actions_secret" "test" {
	repository       = github_repository.test.name
	secret_name      = "%s"
	plaintext_value  = "%s"
}
`

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, repoName, secretName, value),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttrPair("github_actions_secret.test", "repository", "github_repository.test", "name"),
						resource.TestCheckResourceAttr("github_actions_secret.test", "secret_name", secretName),
						resource.TestCheckResourceAttr("github_actions_secret.test", "plaintext_value", value),
						resource.TestCheckNoResourceAttr("github_actions_secret.test", "encrypted_value"),
						resource.TestCheckResourceAttrSet("github_actions_secret.test", "key_id"),
						resource.TestCheckResourceAttrSet("github_actions_secret.test", "created_at"),
						resource.TestCheckResourceAttrSet("github_actions_secret.test", "updated_at"),
					),
				},
				{
					Config: fmt.Sprintf(config, repoName, secretName, updatedValue),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttrPair("github_actions_secret.test", "repository", "github_repository.test", "name"),
						resource.TestCheckResourceAttr("github_actions_secret.test", "secret_name", secretName),
						resource.TestCheckResourceAttr("github_actions_secret.test", "plaintext_value", updatedValue),
						resource.TestCheckNoResourceAttr("github_actions_secret.test", "encrypted_value"),
						resource.TestCheckResourceAttrSet("github_actions_secret.test", "key_id"),
						resource.TestCheckResourceAttrSet("github_actions_secret.test", "created_at"),
						resource.TestCheckResourceAttrSet("github_actions_secret.test", "updated_at"),
					),
				},
			},
		})
	})

	t.Run("create_update_encrypted", func(t *testing.T) {
		randomID := acctest.RandString(5)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)
		secretName := "test"
		value := base64.StdEncoding.EncodeToString([]byte("super_secret_value"))
		updatedValue := base64.StdEncoding.EncodeToString([]byte("updated_super_secret_value"))

		config := `
resource "github_repository" "test" {
	name = "%s"
}

resource "github_actions_secret" "test" {
	repository      = github_repository.test.name
	secret_name     = "%s"
	encrypted_value = "%s"
}
`

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, repoName, secretName, value),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttrPair("github_actions_secret.test", "repository", "github_repository.test", "name"),
						resource.TestCheckResourceAttr("github_actions_secret.test", "secret_name", secretName),
						resource.TestCheckNoResourceAttr("github_actions_secret.test", "plaintext_value"),
						resource.TestCheckResourceAttr("github_actions_secret.test", "encrypted_value", value),
						resource.TestCheckResourceAttrSet("github_actions_secret.test", "key_id"),
						resource.TestCheckResourceAttrSet("github_actions_secret.test", "created_at"),
						resource.TestCheckResourceAttrSet("github_actions_secret.test", "updated_at"),
					),
				},
				{
					Config: fmt.Sprintf(config, repoName, secretName, updatedValue),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttrPair("github_actions_secret.test", "repository", "github_repository.test", "name"),
						resource.TestCheckResourceAttr("github_actions_secret.test", "secret_name", secretName),
						resource.TestCheckNoResourceAttr("github_actions_secret.test", "plaintext_value"),
						resource.TestCheckResourceAttr("github_actions_secret.test", "encrypted_value", updatedValue),
						resource.TestCheckResourceAttrSet("github_actions_secret.test", "key_id"),
						resource.TestCheckResourceAttrSet("github_actions_secret.test", "created_at"),
						resource.TestCheckResourceAttrSet("github_actions_secret.test", "updated_at"),
					),
				},
			},
		})
	})

	t.Run("create_update_encrypted_with_key", func(t *testing.T) {
		randomID := acctest.RandString(5)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)
		secretName := "test"
		value := base64.StdEncoding.EncodeToString([]byte("super_secret_value"))
		updatedValue := base64.StdEncoding.EncodeToString([]byte("updated_super_secret_value"))

		config := `
resource "github_repository" "test" {
	name = "%s"
}

data "github_actions_public_key" "test" {
	repository  = github_repository.test.name
}

resource "github_actions_secret" "test" {
	repository       = github_repository.test.name
	key_id           = data.github_actions_public_key.test.key_id
	secret_name      = "%s"
	encrypted_value  = "%s"
}
`

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, repoName, secretName, value),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttrPair("github_actions_secret.test", "repository", "github_repository.test", "name"),
						resource.TestCheckResourceAttr("github_actions_secret.test", "secret_name", secretName),
						resource.TestCheckNoResourceAttr("github_actions_secret.test", "plaintext_value"),
						resource.TestCheckResourceAttr("github_actions_secret.test", "encrypted_value", value),
						resource.TestCheckResourceAttrSet("github_actions_secret.test", "key_id"),
						resource.TestCheckResourceAttrSet("github_actions_secret.test", "created_at"),
						resource.TestCheckResourceAttrSet("github_actions_secret.test", "updated_at"),
					),
				},
				{
					Config: fmt.Sprintf(config, repoName, secretName, updatedValue),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttrPair("github_actions_secret.test", "repository", "github_repository.test", "name"),
						resource.TestCheckResourceAttr("github_actions_secret.test", "secret_name", secretName),
						resource.TestCheckNoResourceAttr("github_actions_secret.test", "plaintext_value"),
						resource.TestCheckResourceAttr("github_actions_secret.test", "encrypted_value", updatedValue),
						resource.TestCheckResourceAttrSet("github_actions_secret.test", "key_id"),
						resource.TestCheckResourceAttrSet("github_actions_secret.test", "created_at"),
						resource.TestCheckResourceAttrSet("github_actions_secret.test", "updated_at"),
					),
				},
			},
		})
	})

	t.Run("update_on_drift", func(t *testing.T) {
		randomID := acctest.RandString(5)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)
		secretName := "test"

		config := fmt.Sprintf(`
resource "github_repository" "test" {
	name = "%s"
}

resource "github_actions_secret" "test" {
	repository      = github_repository.test.name
	secret_name     = "%s"
	plaintext_value = "test"
}
`, repoName, secretName)

		var beforeCreatedAt string
		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttrSet("github_actions_secret.test", "created_at"),
						resource.TestCheckResourceAttrSet("github_actions_secret.test", "updated_at"),
						resource.TestCheckResourceAttrSet("github_actions_secret.test", "remote_updated_at"),
						func(s *terraform.State) error {
							beforeCreatedAt = s.RootModule().Resources["github_actions_secret.test"].Primary.Attributes["created_at"]
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

						keyID, _, err := getPublicKeyDetails(ctx, meta, repoName)
						if err != nil {
							t.Fatal(err.Error())
						}

						_, err = client.Actions.CreateOrUpdateRepoSecret(ctx, owner, repoName, &github.EncryptedSecret{
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
						resource.TestCheckResourceAttrSet("github_actions_secret.test", "created_at"),
						resource.TestCheckResourceAttrSet("github_actions_secret.test", "updated_at"),
						resource.TestCheckResourceAttrSet("github_actions_secret.test", "remote_updated_at"),
						func(s *terraform.State) error {
							afterCreatedAt := s.RootModule().Resources["github_actions_secret.test"].Primary.Attributes["created_at"]

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
		randomID := acctest.RandString(5)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)
		secretName := "test"

		config := fmt.Sprintf(`
resource "github_repository" "test" {
	name = "%s"
}

resource "github_actions_secret" "test" {
	repository      = github_repository.test.name
	secret_name     = "%s"
	plaintext_value = "test"

	lifecycle {
		ignore_changes = [remote_updated_at]
	}
}
`, repoName, secretName)

		var beforeUpdatedAt string
		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttrSet("github_actions_secret.test", "created_at"),
						resource.TestCheckResourceAttrSet("github_actions_secret.test", "updated_at"),
						resource.TestCheckResourceAttrSet("github_actions_secret.test", "remote_updated_at"),
						func(s *terraform.State) error {
							beforeUpdatedAt = s.RootModule().Resources["github_actions_secret.test"].Primary.Attributes["updated_at"]
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

						keyID, _, err := getPublicKeyDetails(ctx, meta, repoName)
						if err != nil {
							t.Fatal(err.Error())
						}

						_, err = client.Actions.CreateOrUpdateRepoSecret(ctx, owner, repoName, &github.EncryptedSecret{
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
						resource.TestCheckResourceAttrSet("github_actions_secret.test", "created_at"),
						resource.TestCheckResourceAttrSet("github_actions_secret.test", "updated_at"),
						resource.TestCheckResourceAttrSet("github_actions_secret.test", "remote_updated_at"),
						func(s *terraform.State) error {
							afterUpdatedAt := s.RootModule().Resources["github_actions_secret.test"].Primary.Attributes["updated_at"]

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
		randomID := acctest.RandString(5)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)
		updatedRepoName := fmt.Sprintf("%s%s-updated", testResourcePrefix, randomID)

		config := `
resource "github_repository" "test" {
	name = "%s"
}

resource "github_actions_secret" "test" {
	repository      = github_repository.test.name
	secret_name     = "test"
	plaintext_value = "test"
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
						resource.TestCheckResourceAttrSet("github_actions_secret.test", "created_at"),
						resource.TestCheckResourceAttrSet("github_actions_secret.test", "updated_at"),
						func(s *terraform.State) error {
							beforeCreatedAt = s.RootModule().Resources["github_actions_secret.test"].Primary.Attributes["created_at"]
							return nil
						},
					),
				},
				{
					Config: fmt.Sprintf(config, updatedRepoName),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttrSet("github_actions_secret.test", "created_at"),
						resource.TestCheckResourceAttrSet("github_actions_secret.test", "updated_at"),
						func(s *terraform.State) error {
							afterCreatedAt := s.RootModule().Resources["github_actions_secret.test"].Primary.Attributes["created_at"]

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
		randomID := acctest.RandString(5)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)
		repoName2 := fmt.Sprintf("%supdated-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
resource "github_repository" "test" {
	name = "%s"
}

resource "github_repository" "test2" {
	name = "%s"
}

resource "github_actions_secret" "test" {
	repository      = github_repository.test.name
	secret_name     = "test"
	plaintext_value = "test"
}
`, repoName, repoName2)

		configUpdated := fmt.Sprintf(`
resource "github_repository" "test" {
	name = "%s"
}

resource "github_repository" "test2" {
	name = "%s"
}

resource "github_actions_secret" "test" {
	repository      = github_repository.test2.name
	secret_name     = "test"
	plaintext_value = "test"
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
						resource.TestCheckResourceAttrSet("github_actions_secret.test", "created_at"),
						resource.TestCheckResourceAttrSet("github_actions_secret.test", "updated_at"),
						func(s *terraform.State) error {
							beforeCreatedAt = s.RootModule().Resources["github_actions_secret.test"].Primary.Attributes["created_at"]
							return nil
						},
					),
				},
				{
					Config: configUpdated,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttrSet("github_actions_secret.test", "created_at"),
						resource.TestCheckResourceAttrSet("github_actions_secret.test", "updated_at"),
						func(s *terraform.State) error {
							afterCreatedAt := s.RootModule().Resources["github_actions_secret.test"].Primary.Attributes["created_at"]

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
		randomID := acctest.RandString(5)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
	resource "github_repository" "test" {
		name = "%s"
	}

	resource "github_actions_secret" "test" {
		repository      = github_repository.test.name
		secret_name     = "test"
		plaintext_value = "test"
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
		randomID := acctest.RandString(5)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)
		secretName := "test"

		config := fmt.Sprintf(`
resource "github_repository" "test" {
	name = "%s"
}

resource "github_actions_secret" "test" {
	repository      = github_repository.test.name
	secret_name     = "%s"
	plaintext_value = "test"
}
`, repoName, secretName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
				},
				{
					ResourceName:            "github_actions_secret.test",
					ImportState:             true,
					ImportStateVerify:       true,
					ImportStateVerifyIgnore: []string{"key_id", "plaintext_value"},
				},
			},
		})
	})
}
