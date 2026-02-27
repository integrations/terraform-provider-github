package github

import (
	"context"
	"encoding/base64"
	"fmt"
	"regexp"
	"testing"

	"github.com/google/go-github/v83/github"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccGithubDependabotOrganizationSecret(t *testing.T) {
	t.Run("create_update_plaintext", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlpha)
		secretName := fmt.Sprintf("test_%s", randomID)
		value := base64.StdEncoding.EncodeToString([]byte("foo"))
		valueUpdated := base64.StdEncoding.EncodeToString([]byte("bar"))

		config := `
resource "github_dependabot_organization_secret" "test" {
	secret_name = "%s"
	value       = "%s"
	visibility  = "all"
}
`

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, secretName, value),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_dependabot_organization_secret.test", "secret_name", secretName),
						resource.TestCheckResourceAttrSet("github_dependabot_organization_secret.test", "key_id"),
						resource.TestCheckResourceAttr("github_dependabot_organization_secret.test", "value", value),
						resource.TestCheckNoResourceAttr("github_dependabot_organization_secret.test", "value_encrypted"),
						resource.TestCheckResourceAttr("github_dependabot_organization_secret.test", "visibility", "all"),
						resource.TestCheckResourceAttr("github_dependabot_organization_secret.test", "selected_repository_ids.#", "0"),
						resource.TestCheckResourceAttrSet("github_dependabot_organization_secret.test", "created_at"),
						resource.TestCheckResourceAttrSet("github_dependabot_organization_secret.test", "updated_at"),
					),
				},
				{
					Config: fmt.Sprintf(config, secretName, valueUpdated),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_dependabot_organization_secret.test", "secret_name", secretName),
						resource.TestCheckResourceAttrSet("github_dependabot_organization_secret.test", "key_id"),
						resource.TestCheckResourceAttr("github_dependabot_organization_secret.test", "value", valueUpdated),
						resource.TestCheckNoResourceAttr("github_dependabot_organization_secret.test", "value_encrypted"),
						resource.TestCheckResourceAttr("github_dependabot_organization_secret.test", "visibility", "all"),
						resource.TestCheckResourceAttr("github_dependabot_organization_secret.test", "selected_repository_ids.#", "0"),
						resource.TestCheckResourceAttrSet("github_dependabot_organization_secret.test", "created_at"),
						resource.TestCheckResourceAttrSet("github_dependabot_organization_secret.test", "updated_at"),
					),
				},
			},
		})
	})

	t.Run("create_update_encrypted", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlpha)
		secretName := fmt.Sprintf("test_%s", randomID)
		value := base64.StdEncoding.EncodeToString([]byte("foo"))
		valueUpdated := base64.StdEncoding.EncodeToString([]byte("bar"))

		config := `
resource "github_dependabot_organization_secret" "test" {
	secret_name     = "%s"
	value_encrypted = "%s"
	visibility      = "all"
}
`

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, secretName, value),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_dependabot_organization_secret.test", "secret_name", secretName),
						resource.TestCheckResourceAttrSet("github_dependabot_organization_secret.test", "key_id"),
						resource.TestCheckNoResourceAttr("github_dependabot_organization_secret.test", "value"),
						resource.TestCheckResourceAttr("github_dependabot_organization_secret.test", "value_encrypted", value),
						resource.TestCheckResourceAttr("github_dependabot_organization_secret.test", "visibility", "all"),
						resource.TestCheckResourceAttr("github_dependabot_organization_secret.test", "selected_repository_ids.#", "0"),
						resource.TestCheckResourceAttrSet("github_dependabot_organization_secret.test", "created_at"),
						resource.TestCheckResourceAttrSet("github_dependabot_organization_secret.test", "updated_at"),
					),
				},
				{
					Config: fmt.Sprintf(config, secretName, valueUpdated),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_dependabot_organization_secret.test", "secret_name", secretName),
						resource.TestCheckResourceAttrSet("github_dependabot_organization_secret.test", "key_id"),
						resource.TestCheckNoResourceAttr("github_dependabot_organization_secret.test", "value"),
						resource.TestCheckResourceAttr("github_dependabot_organization_secret.test", "value_encrypted", valueUpdated),
						resource.TestCheckResourceAttr("github_dependabot_organization_secret.test", "visibility", "all"),
						resource.TestCheckResourceAttr("github_dependabot_organization_secret.test", "selected_repository_ids.#", "0"),
						resource.TestCheckResourceAttrSet("github_dependabot_organization_secret.test", "created_at"),
						resource.TestCheckResourceAttrSet("github_dependabot_organization_secret.test", "updated_at"),
					),
				},
			},
		})
	})

	t.Run("create_update_encrypted_with_key", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlpha)
		secretName := fmt.Sprintf("test_%s", randomID)
		value := base64.StdEncoding.EncodeToString([]byte("foo"))
		valueUpdated := base64.StdEncoding.EncodeToString([]byte("bar"))

		config := `
data "github_dependabot_organization_public_key" "default" {}

resource "github_dependabot_organization_secret" "test" {
	secret_name     = "%s"
	key_id          = data.github_dependabot_organization_public_key.default.key_id
	value_encrypted = "%s"
	visibility      = "all"
}
`

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, secretName, value),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_dependabot_organization_secret.test", "secret_name", secretName),
						resource.TestCheckResourceAttrSet("github_dependabot_organization_secret.test", "key_id"),
						resource.TestCheckNoResourceAttr("github_dependabot_organization_secret.test", "value"),
						resource.TestCheckResourceAttr("github_dependabot_organization_secret.test", "value_encrypted", value),
						resource.TestCheckResourceAttr("github_dependabot_organization_secret.test", "visibility", "all"),
						resource.TestCheckResourceAttr("github_dependabot_organization_secret.test", "selected_repository_ids.#", "0"),
						resource.TestCheckResourceAttrSet("github_dependabot_organization_secret.test", "created_at"),
						resource.TestCheckResourceAttrSet("github_dependabot_organization_secret.test", "updated_at"),
					),
				},
				{
					Config: fmt.Sprintf(config, secretName, valueUpdated),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_dependabot_organization_secret.test", "secret_name", secretName),
						resource.TestCheckResourceAttrSet("github_dependabot_organization_secret.test", "key_id"),
						resource.TestCheckNoResourceAttr("github_dependabot_organization_secret.test", "value"),
						resource.TestCheckResourceAttr("github_dependabot_organization_secret.test", "value_encrypted", valueUpdated),
						resource.TestCheckResourceAttr("github_dependabot_organization_secret.test", "visibility", "all"),
						resource.TestCheckResourceAttr("github_dependabot_organization_secret.test", "selected_repository_ids.#", "0"),
						resource.TestCheckResourceAttrSet("github_dependabot_organization_secret.test", "created_at"),
						resource.TestCheckResourceAttrSet("github_dependabot_organization_secret.test", "updated_at"),
					),
				},
			},
		})
	})

	t.Run("create_update_visibility_all", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlpha)
		secretName := fmt.Sprintf("test_%s", randomID)
		value := base64.StdEncoding.EncodeToString([]byte("foo"))
		valueUpdated := base64.StdEncoding.EncodeToString([]byte("bar"))

		config := `
resource "github_dependabot_organization_secret" "test" {
	secret_name = "%s"
	value       = "%s"
	visibility  = "all"
}
`

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, secretName, value),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_dependabot_organization_secret.test", "secret_name", secretName),
						resource.TestCheckResourceAttrSet("github_dependabot_organization_secret.test", "key_id"),
						resource.TestCheckResourceAttr("github_dependabot_organization_secret.test", "value", value),
						resource.TestCheckNoResourceAttr("github_dependabot_organization_secret.test", "value_encrypted"),
						resource.TestCheckResourceAttr("github_dependabot_organization_secret.test", "visibility", "all"),
						resource.TestCheckResourceAttr("github_dependabot_organization_secret.test", "selected_repository_ids.#", "0"),
						resource.TestCheckResourceAttrSet("github_dependabot_organization_secret.test", "created_at"),
						resource.TestCheckResourceAttrSet("github_dependabot_organization_secret.test", "updated_at"),
					),
				},
				{
					Config: fmt.Sprintf(config, secretName, valueUpdated),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_dependabot_organization_secret.test", "secret_name", secretName),
						resource.TestCheckResourceAttrSet("github_dependabot_organization_secret.test", "key_id"),
						resource.TestCheckResourceAttr("github_dependabot_organization_secret.test", "value", valueUpdated),
						resource.TestCheckNoResourceAttr("github_dependabot_organization_secret.test", "value_encrypted"),
						resource.TestCheckResourceAttr("github_dependabot_organization_secret.test", "visibility", "all"),
						resource.TestCheckResourceAttr("github_dependabot_organization_secret.test", "selected_repository_ids.#", "0"),
						resource.TestCheckResourceAttrSet("github_dependabot_organization_secret.test", "created_at"),
						resource.TestCheckResourceAttrSet("github_dependabot_organization_secret.test", "updated_at"),
					),
				},
			},
		})
	})

	t.Run("create_update_visibility_private", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlpha)
		secretName := fmt.Sprintf("test_%s", randomID)
		value := base64.StdEncoding.EncodeToString([]byte("foo"))
		valueUpdated := base64.StdEncoding.EncodeToString([]byte("bar"))

		config := `
resource "github_dependabot_organization_secret" "test" {
	secret_name = "%s"
	value       = "%s"
	visibility  = "private"
}
`

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, secretName, value),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_dependabot_organization_secret.test", "secret_name", secretName),
						resource.TestCheckResourceAttrSet("github_dependabot_organization_secret.test", "key_id"),
						resource.TestCheckResourceAttr("github_dependabot_organization_secret.test", "value", value),
						resource.TestCheckNoResourceAttr("github_dependabot_organization_secret.test", "value_encrypted"),
						resource.TestCheckResourceAttr("github_dependabot_organization_secret.test", "visibility", "private"),
						resource.TestCheckResourceAttr("github_dependabot_organization_secret.test", "selected_repository_ids.#", "0"),
						resource.TestCheckResourceAttrSet("github_dependabot_organization_secret.test", "created_at"),
						resource.TestCheckResourceAttrSet("github_dependabot_organization_secret.test", "updated_at"),
					),
				},
				{
					Config: fmt.Sprintf(config, secretName, valueUpdated),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_dependabot_organization_secret.test", "secret_name", secretName),
						resource.TestCheckResourceAttrSet("github_dependabot_organization_secret.test", "key_id"),
						resource.TestCheckResourceAttr("github_dependabot_organization_secret.test", "value", valueUpdated),
						resource.TestCheckNoResourceAttr("github_dependabot_organization_secret.test", "value_encrypted"),
						resource.TestCheckResourceAttr("github_dependabot_organization_secret.test", "visibility", "private"),
						resource.TestCheckResourceAttr("github_dependabot_organization_secret.test", "selected_repository_ids.#", "0"),
						resource.TestCheckResourceAttrSet("github_dependabot_organization_secret.test", "created_at"),
						resource.TestCheckResourceAttrSet("github_dependabot_organization_secret.test", "updated_at"),
					),
				},
			},
		})
	})

	t.Run("create_update_visibility_selected", func(t *testing.T) {
		repoName0 := fmt.Sprintf("%s%s", testResourcePrefix, acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum))
		repoName1 := fmt.Sprintf("%s%s", testResourcePrefix, acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum))
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlpha)
		secretName := fmt.Sprintf("test_%s", randomID)
		value := base64.StdEncoding.EncodeToString([]byte("foo"))
		valueUpdated := base64.StdEncoding.EncodeToString([]byte("bar"))

		config := `
resource "github_repository" "test_0" {
	name = "%s"
}

resource "github_repository" "test_1" {
	name = "%s"
}

resource "github_dependabot_organization_secret" "test" {
	secret_name = "%s"
	value       = "%s"
	visibility  = "selected"

	selected_repository_ids = [github_repository.test_%s.repo_id]
}
`

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, repoName0, repoName1, secretName, value, "0"),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_dependabot_organization_secret.test", "secret_name", secretName),
						resource.TestCheckResourceAttrSet("github_dependabot_organization_secret.test", "key_id"),
						resource.TestCheckResourceAttr("github_dependabot_organization_secret.test", "value", value),
						resource.TestCheckNoResourceAttr("github_dependabot_organization_secret.test", "value_encrypted"),
						resource.TestCheckResourceAttr("github_dependabot_organization_secret.test", "visibility", "selected"),
						resource.TestCheckResourceAttr("github_dependabot_organization_secret.test", "selected_repository_ids.#", "1"),
						resource.TestCheckResourceAttrPair("github_dependabot_organization_secret.test", "selected_repository_ids.0", "github_repository.test_0", "repo_id"),
						resource.TestCheckResourceAttrSet("github_dependabot_organization_secret.test", "created_at"),
						resource.TestCheckResourceAttrSet("github_dependabot_organization_secret.test", "updated_at"),
					),
				},
				{
					Config: fmt.Sprintf(config, repoName0, repoName1, secretName, valueUpdated, "1"),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_dependabot_organization_secret.test", "secret_name", secretName),
						resource.TestCheckResourceAttrSet("github_dependabot_organization_secret.test", "key_id"),
						resource.TestCheckResourceAttr("github_dependabot_organization_secret.test", "value", valueUpdated),
						resource.TestCheckNoResourceAttr("github_dependabot_organization_secret.test", "value_encrypted"),
						resource.TestCheckResourceAttr("github_dependabot_organization_secret.test", "visibility", "selected"),
						resource.TestCheckResourceAttr("github_dependabot_organization_secret.test", "selected_repository_ids.#", "1"),
						resource.TestCheckResourceAttrPair("github_dependabot_organization_secret.test", "selected_repository_ids.0", "github_repository.test_1", "repo_id"),
						resource.TestCheckResourceAttrSet("github_dependabot_organization_secret.test", "created_at"),
						resource.TestCheckResourceAttrSet("github_dependabot_organization_secret.test", "updated_at"),
					),
				},
			},
		})
	})

	t.Run("create_update_visibility_selected_no_repo_ids", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlpha)
		secretName := fmt.Sprintf("test_%s", randomID)
		value := base64.StdEncoding.EncodeToString([]byte("foo"))
		valueUpdated := base64.StdEncoding.EncodeToString([]byte("bar"))

		config := `
resource "github_dependabot_organization_secret" "test" {
	secret_name = "%s"
	value       = "%s"
	visibility  = "selected"
}
`

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, secretName, value),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_dependabot_organization_secret.test", "secret_name", secretName),
						resource.TestCheckResourceAttrSet("github_dependabot_organization_secret.test", "key_id"),
						resource.TestCheckResourceAttr("github_dependabot_organization_secret.test", "value", value),
						resource.TestCheckNoResourceAttr("github_dependabot_organization_secret.test", "value_encrypted"),
						resource.TestCheckResourceAttr("github_dependabot_organization_secret.test", "visibility", "selected"),
						resource.TestCheckResourceAttr("github_dependabot_organization_secret.test", "selected_repository_ids.#", "0"),
						resource.TestCheckResourceAttrSet("github_dependabot_organization_secret.test", "created_at"),
						resource.TestCheckResourceAttrSet("github_dependabot_organization_secret.test", "updated_at"),
					),
				},
				{
					Config: fmt.Sprintf(config, secretName, valueUpdated),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_dependabot_organization_secret.test", "secret_name", secretName),
						resource.TestCheckResourceAttrSet("github_dependabot_organization_secret.test", "key_id"),
						resource.TestCheckResourceAttr("github_dependabot_organization_secret.test", "value", valueUpdated),
						resource.TestCheckNoResourceAttr("github_dependabot_organization_secret.test", "value_encrypted"),
						resource.TestCheckResourceAttr("github_dependabot_organization_secret.test", "visibility", "selected"),
						resource.TestCheckResourceAttr("github_dependabot_organization_secret.test", "selected_repository_ids.#", "0"),
						resource.TestCheckResourceAttrSet("github_dependabot_organization_secret.test", "created_at"),
						resource.TestCheckResourceAttrSet("github_dependabot_organization_secret.test", "updated_at"),
					),
				},
			},
		})
	})

	t.Run("create_update_change_visibility", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlpha)
		secretName := fmt.Sprintf("test_%s", randomID)
		value := base64.StdEncoding.EncodeToString([]byte("foo"))
		visibility := "all"
		valueUpdated := base64.StdEncoding.EncodeToString([]byte("bar"))
		visibilityUpdated := "private"

		config := `
resource "github_dependabot_organization_secret" "test" {
	secret_name = "%s"
	value       = "%s"
	visibility  = "%s"
}
`

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, secretName, value, visibility),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_dependabot_organization_secret.test", "secret_name", secretName),
						resource.TestCheckResourceAttrSet("github_dependabot_organization_secret.test", "key_id"),
						resource.TestCheckResourceAttr("github_dependabot_organization_secret.test", "value", value),
						resource.TestCheckNoResourceAttr("github_dependabot_organization_secret.test", "value_encrypted"),
						resource.TestCheckResourceAttr("github_dependabot_organization_secret.test", "visibility", visibility),
						resource.TestCheckResourceAttr("github_dependabot_organization_secret.test", "selected_repository_ids.#", "0"),
						resource.TestCheckResourceAttrSet("github_dependabot_organization_secret.test", "created_at"),
						resource.TestCheckResourceAttrSet("github_dependabot_organization_secret.test", "updated_at"),
					),
				},
				{
					Config: fmt.Sprintf(config, secretName, valueUpdated, visibilityUpdated),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_dependabot_organization_secret.test", "secret_name", secretName),
						resource.TestCheckResourceAttrSet("github_dependabot_organization_secret.test", "key_id"),
						resource.TestCheckResourceAttr("github_dependabot_organization_secret.test", "value", valueUpdated),
						resource.TestCheckNoResourceAttr("github_dependabot_organization_secret.test", "value_encrypted"),
						resource.TestCheckResourceAttr("github_dependabot_organization_secret.test", "visibility", visibilityUpdated),
						resource.TestCheckResourceAttr("github_dependabot_organization_secret.test", "selected_repository_ids.#", "0"),
						resource.TestCheckResourceAttrSet("github_dependabot_organization_secret.test", "created_at"),
						resource.TestCheckResourceAttrSet("github_dependabot_organization_secret.test", "updated_at"),
					),
				},
			},
		})
	})

	t.Run("update_on_drift", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlpha)
		secretName := fmt.Sprintf("test_%s", randomID)
		value := base64.StdEncoding.EncodeToString([]byte("foo"))

		config := fmt.Sprintf(`
resource "github_dependabot_organization_secret" "test" {
	secret_name = "%s"
	value       = "%s"
	visibility  = "all"
}
`, secretName, value)

		var beforeCreatedAt string
		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttrSet("github_dependabot_organization_secret.test", "created_at"),
						resource.TestCheckResourceAttrSet("github_dependabot_organization_secret.test", "updated_at"),
						resource.TestCheckResourceAttrSet("github_dependabot_organization_secret.test", "remote_updated_at"),
						func(s *terraform.State) error {
							beforeCreatedAt = s.RootModule().Resources["github_dependabot_organization_secret.test"].Primary.Attributes["created_at"]
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

						keyID, _, err := getOrganizationPublicKeyDetails(ctx, meta)
						if err != nil {
							t.Fatal(err.Error())
						}

						_, err = client.Actions.CreateOrUpdateOrgSecret(ctx, owner, &github.EncryptedSecret{
							Name:           secretName,
							EncryptedValue: base64.StdEncoding.EncodeToString([]byte("updated_super_secret_value")),
							KeyID:          keyID,
							Visibility:     "all",
						})
						if err != nil {
							t.Fatal(err.Error())
						}
					},
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttrSet("github_dependabot_organization_secret.test", "created_at"),
						resource.TestCheckResourceAttrSet("github_dependabot_organization_secret.test", "updated_at"),
						resource.TestCheckResourceAttrSet("github_dependabot_organization_secret.test", "remote_updated_at"),
						func(s *terraform.State) error {
							afterCreatedAt := s.RootModule().Resources["github_dependabot_organization_secret.test"].Primary.Attributes["created_at"]

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
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlpha)
		secretName := fmt.Sprintf("test_%s", randomID)
		value := base64.StdEncoding.EncodeToString([]byte("foo"))

		config := fmt.Sprintf(`
resource "github_dependabot_organization_secret" "test" {
	secret_name = "%s"
	value       = "%s"
	visibility  = "all"

	lifecycle {
		ignore_changes = [remote_updated_at]
	}
}
`, secretName, value)

		var beforeUpdatedAt string
		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttrSet("github_dependabot_organization_secret.test", "created_at"),
						resource.TestCheckResourceAttrSet("github_dependabot_organization_secret.test", "updated_at"),
						resource.TestCheckResourceAttrSet("github_dependabot_organization_secret.test", "remote_updated_at"),
						func(s *terraform.State) error {
							beforeUpdatedAt = s.RootModule().Resources["github_dependabot_organization_secret.test"].Primary.Attributes["updated_at"]
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

						keyID, _, err := getOrganizationPublicKeyDetails(ctx, meta)
						if err != nil {
							t.Fatal(err.Error())
						}

						_, err = client.Actions.CreateOrUpdateOrgSecret(ctx, owner, &github.EncryptedSecret{
							Name:           secretName,
							EncryptedValue: base64.StdEncoding.EncodeToString([]byte("updated_super_secret_value")),
							KeyID:          keyID,
							Visibility:     "all",
						})
						if err != nil {
							t.Fatal(err.Error())
						}
					},
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttrSet("github_dependabot_organization_secret.test", "created_at"),
						resource.TestCheckResourceAttrSet("github_dependabot_organization_secret.test", "updated_at"),
						resource.TestCheckResourceAttrSet("github_dependabot_organization_secret.test", "remote_updated_at"),
						func(s *terraform.State) error {
							afterUpdatedAt := s.RootModule().Resources["github_dependabot_organization_secret.test"].Primary.Attributes["updated_at"]
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

	t.Run("destroy", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlpha)
		secretName := fmt.Sprintf("test_%s", randomID)
		value := base64.StdEncoding.EncodeToString([]byte("foo"))

		config := fmt.Sprintf(`
resource "github_dependabot_organization_secret" "test" {
	secret_name = "%s"
	value       = "%s"
	visibility  = "all"
}
`, secretName, value)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
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
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlpha)
		secretName := fmt.Sprintf("test_%s", randomID)
		value := base64.StdEncoding.EncodeToString([]byte("foo"))

		config := fmt.Sprintf(`
resource "github_dependabot_organization_secret" "test" {
	secret_name = "%s"
	value       = "%s"
	visibility  = "all"
}
`, secretName, value)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
				},
				{
					ResourceName:            "github_dependabot_organization_secret.test",
					ImportState:             true,
					ImportStateVerify:       true,
					ImportStateVerifyIgnore: []string{"key_id", "value", "destroy_on_drift"},
				},
			},
		})
	})

	t.Run("error_on_invalid_selected_repository_ids", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlpha)
		secretName := fmt.Sprintf("test_%s", randomID)
		value := base64.StdEncoding.EncodeToString([]byte("foo"))

		config := fmt.Sprintf(`
resource "github_dependabot_organization_secret" "test" {
	secret_name = "%s"
	value       = "%s"
	visibility  = "all"

	selected_repository_ids = [123456]
}
`, secretName, value)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      config,
					ExpectError: regexp.MustCompile("cannot use selected_repository_ids without visibility being set to selected"),
				},
			},
		})
	})
}
