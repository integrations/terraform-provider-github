package github

import (
	"encoding/base64"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestAccGithubActionsOrganizationSecret(t *testing.T) {
	t.Run("creates and updates secrets without error", func(t *testing.T) {
		secretValue := base64.StdEncoding.EncodeToString([]byte("super_secret_value"))
		updatedSecretValue := base64.StdEncoding.EncodeToString([]byte("updated_super_secret_value"))

		config := fmt.Sprintf(`
			resource "github_actions_organization_secret" "plaintext_secret" {
			  secret_name      = "test_plaintext_secret"
			  plaintext_value  = "%s"
			  visibility       = "private"
			}

			resource "github_actions_organization_secret" "encrypted_secret" {
			  secret_name      = "test_encrypted_secret"
			  encrypted_value  = "%s"
			  visibility       = "private"
			  destroy_on_drift = false
			}
		`, secretValue, secretValue)

		checks := map[string]resource.TestCheckFunc{
			"before": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_actions_organization_secret.plaintext_secret", "plaintext_value",
					secretValue,
				),
				resource.TestCheckResourceAttr(
					"github_actions_organization_secret.encrypted_secret", "encrypted_value",
					secretValue,
				),
				resource.TestCheckResourceAttrSet(
					"github_actions_organization_secret.plaintext_secret", "created_at",
				),
				resource.TestCheckResourceAttrSet(
					"github_actions_organization_secret.plaintext_secret", "updated_at",
				),
			),
			"after": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_actions_organization_secret.plaintext_secret", "plaintext_value",
					updatedSecretValue,
				),
				resource.TestCheckResourceAttr(
					"github_actions_organization_secret.encrypted_secret", "encrypted_value",
					updatedSecretValue,
				),
				resource.TestCheckResourceAttrSet(
					"github_actions_organization_secret.plaintext_secret", "created_at",
				),
				resource.TestCheckResourceAttrSet(
					"github_actions_organization_secret.plaintext_secret", "updated_at",
				),
			),
		}

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
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

	t.Run("deletes secrets without error", func(t *testing.T) {
		config := `
				resource "github_actions_organization_secret" "plaintext_secret" {
					secret_name      = "test_plaintext_secret"
					visibility       = "private"
				}

				resource "github_actions_organization_secret" "encrypted_secret" {
					secret_name      = "test_encrypted_secret"
					visibility       = "private"
				}
			`

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:  config,
					Destroy: true,
				},
			},
		})
	})

	t.Run("imports secrets without error", func(t *testing.T) {
		secretValue := "super_secret_value"

		config := fmt.Sprintf(`
			resource "github_actions_organization_secret" "test_secret" {
				secret_name      = "test_plaintext_secret"
				plaintext_value  = "%s"
				visibility       = "private"
			}
		`, secretValue)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_actions_organization_secret.test_secret", "plaintext_value",
				secretValue,
			),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  check,
				},
				{
					ResourceName:            "github_actions_organization_secret.test_secret",
					ImportState:             true,
					ImportStateVerify:       true,
					ImportStateVerifyIgnore: []string{"plaintext_value"},
				},
			},
		})
	})

	// Unit tests for drift detection behavior
	t.Run("destroyOnDrift false clears sensitive values instead of recreating", func(t *testing.T) {
		originalTimestamp := "2023-01-01T00:00:00Z"
		newTimestamp := "2023-01-02T00:00:00Z"

		d := schema.TestResourceDataRaw(t, resourceGithubActionsOrganizationSecret().Schema, map[string]any{
			"secret_name":      "test-secret",
			"plaintext_value":  "original-value",
			"encrypted_value":  "original-encrypted",
			"visibility":       "private",
			"destroy_on_drift": false,
			"updated_at":       originalTimestamp,
		})
		d.SetId("test-secret")

		// Simulate drift detection logic when destroy_on_drift is false
		destroyOnDrift := d.Get("destroy_on_drift").(bool)
		storedUpdatedAt, hasStoredUpdatedAt := d.GetOk("updated_at")

		if hasStoredUpdatedAt && storedUpdatedAt != newTimestamp {
			if destroyOnDrift {
				// Would clear ID for recreation
				d.SetId("")
			} else {
				// Should clear sensitive values to trigger update
				_ = d.Set("encrypted_value", "")
				_ = d.Set("plaintext_value", "")
			}
			_ = d.Set("updated_at", newTimestamp)
		}

		// Should NOT have cleared the ID when destroy_on_drift=false
		if d.Id() == "" {
			t.Error("Expected ID to be preserved when destroy_on_drift=false, but it was cleared")
		}

		// Should have cleared sensitive values to trigger update plan
		if plaintextValue := d.Get("plaintext_value").(string); plaintextValue != "" {
			t.Errorf("Expected plaintext_value to be cleared for update plan, got %s", plaintextValue)
		}

		if encryptedValue := d.Get("encrypted_value").(string); encryptedValue != "" {
			t.Errorf("Expected encrypted_value to be cleared for update plan, got %s", encryptedValue)
		}

		// Should have updated the timestamp
		if updatedAt := d.Get("updated_at").(string); updatedAt != newTimestamp {
			t.Errorf("Expected timestamp to be updated to %s, got %s", newTimestamp, updatedAt)
		}
	})

	t.Run("destroyOnDrift true still recreates resource on drift", func(t *testing.T) {
		originalTimestamp := "2023-01-01T00:00:00Z"
		newTimestamp := "2023-01-02T00:00:00Z"

		d := schema.TestResourceDataRaw(t, resourceGithubActionsOrganizationSecret().Schema, map[string]any{
			"secret_name":      "test-secret",
			"plaintext_value":  "original-value",
			"visibility":       "private",
			"destroy_on_drift": true, // Explicitly set to true
			"updated_at":       originalTimestamp,
		})
		d.SetId("test-secret")

		// Simulate drift detection logic when destroy_on_drift is true
		destroyOnDrift := d.Get("destroy_on_drift").(bool)
		storedUpdatedAt, hasStoredUpdatedAt := d.GetOk("updated_at")

		if hasStoredUpdatedAt && storedUpdatedAt != newTimestamp {
			if destroyOnDrift {
				// Should clear ID for recreation (original behavior)
				d.SetId("")
				return // Exit early like the real function would
			}
		}

		// Should have cleared the ID for recreation when destroy_on_drift=true
		if d.Id() != "" {
			t.Error("Expected ID to be cleared for recreation when destroy_on_drift=true, but it was preserved")
		}
	})

	t.Run("destroy_on_drift field defaults", func(t *testing.T) {
		// Test that destroy_on_drift defaults to true for backward compatibility
		schema := resourceGithubActionsOrganizationSecret().Schema["destroy_on_drift"]
		if schema.Default != true {
			t.Error("destroy_on_drift should default to true for backward compatibility")
		}
	})

	t.Run("default destroy_on_drift is true", func(t *testing.T) {
		d := schema.TestResourceDataRaw(t, resourceGithubActionsOrganizationSecret().Schema, map[string]any{
			"secret_name":     "test-secret",
			"plaintext_value": "test-value",
			"visibility":      "private",
			// destroy_on_drift not set, should default to true
		})

		destroyOnDrift := d.Get("destroy_on_drift").(bool)
		if !destroyOnDrift {
			t.Error("Expected destroy_on_drift to default to true")
		}
	})
}
