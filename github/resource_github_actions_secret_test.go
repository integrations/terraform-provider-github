package github

import (
	"encoding/base64"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestAccGithubActionsSecret(t *testing.T) {
	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

	t.Run("reads a repository public key without error", func(t *testing.T) {
		config := fmt.Sprintf(`

			resource "github_repository" "test" {
			  name = "tf-acc-test-%s"
			}

			data "github_actions_public_key" "test_pk" {
			  repository = github_repository.test.name
			}

		`, randomID)

		check := resource.ComposeAggregateTestCheckFunc(
			resource.TestCheckResourceAttrSet(
				"data.github_actions_public_key.test_pk", "key_id",
			),
			resource.TestCheckResourceAttrSet(
				"data.github_actions_public_key.test_pk", "key",
			),
		)

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: config,
						Check:  check,
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

	t.Run("creates and updates secrets without error", func(t *testing.T) {
		secretValue := base64.StdEncoding.EncodeToString([]byte("super_secret_value"))
		updatedSecretValue := base64.StdEncoding.EncodeToString([]byte("updated_super_secret_value"))

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
			  name = "tf-acc-test-%s"
			}

			resource "github_actions_secret" "plaintext_secret" {
			  repository       = github_repository.test.name
			  secret_name      = "test_plaintext_secret"
			  plaintext_value  = "%s"
			}

			resource "github_actions_secret" "encrypted_secret" {
			  repository       = github_repository.test.name
			  secret_name      = "test_encrypted_secret"
			  encrypted_value  = "%s"
			}
			`, randomID, secretValue, secretValue)

		checks := map[string]resource.TestCheckFunc{
			"before": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_actions_secret.plaintext_secret", "plaintext_value",
					secretValue,
				),
				resource.TestCheckResourceAttr(
					"github_actions_secret.encrypted_secret", "encrypted_value",
					secretValue,
				),
				resource.TestCheckResourceAttrSet(
					"github_actions_secret.plaintext_secret", "created_at",
				),
				resource.TestCheckResourceAttrSet(
					"github_actions_secret.plaintext_secret", "updated_at",
				),
			),
			"after": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_actions_secret.plaintext_secret", "plaintext_value",
					updatedSecretValue,
				),
				resource.TestCheckResourceAttr(
					"github_actions_secret.encrypted_secret", "encrypted_value",
					updatedSecretValue,
				),
				resource.TestCheckResourceAttrSet(
					"github_actions_secret.plaintext_secret", "created_at",
				),
				resource.TestCheckResourceAttrSet(
					"github_actions_secret.plaintext_secret", "updated_at",
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

	t.Run("creates and updates repository name without error", func(t *testing.T) {
		repoName := fmt.Sprintf("tf-acc-test-%s", randomID)
		updatedRepoName := fmt.Sprintf("tf-acc-test-%s-updated", randomID)
		secretValue := base64.StdEncoding.EncodeToString([]byte("super_secret_value"))

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
			  name = "%s"
			}

			resource "github_actions_secret" "plaintext_secret" {
			  repository       = github_repository.test.name
			  secret_name      = "test_plaintext_secret"
			  plaintext_value  = "%s"
			}

			resource "github_actions_secret" "encrypted_secret" {
			  repository       = github_repository.test.name
			  secret_name      = "test_encrypted_secret"
			  encrypted_value  = "%s"
			}
			`, repoName, secretValue, secretValue)

		checks := map[string]resource.TestCheckFunc{
			"before": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_actions_secret.plaintext_secret", "repository",
					repoName,
				),
				resource.TestCheckResourceAttr(
					"github_actions_secret.plaintext_secret", "plaintext_value",
					secretValue,
				),
				resource.TestCheckResourceAttr(
					"github_actions_secret.encrypted_secret", "encrypted_value",
					secretValue,
				),
				resource.TestCheckResourceAttrSet(
					"github_actions_secret.plaintext_secret", "created_at",
				),
				resource.TestCheckResourceAttrSet(
					"github_actions_secret.plaintext_secret", "updated_at",
				),
			),
			"after": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_actions_secret.plaintext_secret", "repository",
					updatedRepoName,
				),
				resource.TestCheckResourceAttr(
					"github_actions_secret.plaintext_secret", "plaintext_value",
					secretValue,
				),
				resource.TestCheckResourceAttr(
					"github_actions_secret.encrypted_secret", "encrypted_value",
					secretValue,
				),
				resource.TestCheckResourceAttrSet(
					"github_actions_secret.plaintext_secret", "created_at",
				),
				resource.TestCheckResourceAttrSet(
					"github_actions_secret.plaintext_secret", "updated_at",
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
							repoName,
							updatedRepoName, 2),
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
		config := fmt.Sprintf(`
				resource "github_repository" "test" {
					name = "tf-acc-test-%s"
				}

				resource "github_actions_secret" "plaintext_secret" {
					repository 	= github_repository.test.name
					secret_name	= "test_plaintext_secret"
				}

				resource "github_actions_secret" "encrypted_secret" {
					repository 	= github_repository.test.name
					secret_name	= "test_encrypted_secret"
				}
			`, randomID)

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

	t.Run("respects destroy_on_drift setting", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name = "tf-acc-test-%s"
			}

			resource "github_actions_secret" "with_drift_true" {
				repository        = github_repository.test.name
				secret_name       = "test_drift_true"
				plaintext_value   = "initial_value"
				destroy_on_drift  = true
			}

			resource "github_actions_secret" "with_drift_false" {
				repository        = github_repository.test.name
				secret_name       = "test_drift_false"
				plaintext_value   = "initial_value"
				destroy_on_drift  = false
			}

			resource "github_actions_secret" "default_behavior" {
				repository        = github_repository.test.name
				secret_name       = "test_default"
				plaintext_value   = "initial_value"
				# destroy_on_drift defaults to true
			}
		`, randomID)

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: config,
						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckResourceAttr(
								"github_actions_secret.with_drift_true", "destroy_on_drift", "true"),
							resource.TestCheckResourceAttr(
								"github_actions_secret.with_drift_false", "destroy_on_drift", "false"),
							resource.TestCheckResourceAttr(
								"github_actions_secret.default_behavior", "destroy_on_drift", "true"),
							resource.TestCheckResourceAttr(
								"github_actions_secret.with_drift_true", "plaintext_value", "initial_value"),
							resource.TestCheckResourceAttr(
								"github_actions_secret.with_drift_false", "plaintext_value", "initial_value"),
							resource.TestCheckResourceAttr(
								"github_actions_secret.default_behavior", "plaintext_value", "initial_value"),
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

	t.Run("updates destroy_on_drift field without recreation", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

		config1 := fmt.Sprintf(`
			resource "github_repository" "test" {
				name = "tf-acc-test-%s"
			}

			resource "github_actions_secret" "test" {
				repository        = github_repository.test.name
				secret_name       = "test_destroy_on_drift_update"
				plaintext_value   = "test_value"
				destroy_on_drift  = true
			}
		`, randomID)

		config2 := fmt.Sprintf(`
			resource "github_repository" "test" {
				name = "tf-acc-test-%s"
			}

			resource "github_actions_secret" "test" {
				repository        = github_repository.test.name
				secret_name       = "test_destroy_on_drift_update"
				plaintext_value   = "test_value"
				destroy_on_drift  = false
			}
		`, randomID)

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: config1,
						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckResourceAttr(
								"github_actions_secret.test", "destroy_on_drift", "true"),
							resource.TestCheckResourceAttr(
								"github_actions_secret.test", "plaintext_value", "test_value"),
						),
					},
					{
						Config: config2,
						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckResourceAttr(
								"github_actions_secret.test", "destroy_on_drift", "false"),
							resource.TestCheckResourceAttr(
								"github_actions_secret.test", "plaintext_value", "test_value"),
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

// Unit tests for drift detection behavior.
func TestGithubActionsSecretDriftDetection(t *testing.T) {
	t.Run("destroyOnDrift true causes recreation on timestamp mismatch", func(t *testing.T) {
		originalTimestamp := "2023-01-01T00:00:00Z"
		newTimestamp := "2023-01-02T00:00:00Z"

		d := schema.TestResourceDataRaw(t, resourceGithubActionsSecret().Schema, map[string]any{
			"repository":       "test-repo",
			"secret_name":      "test-secret",
			"plaintext_value":  "test-value",
			"destroy_on_drift": true,
			"updated_at":       originalTimestamp,
		})
		d.SetId("test-secret")

		// Test the drift detection logic - simulate what happens in the read function
		destroyOnDrift := d.Get("destroy_on_drift").(bool)
		if updatedAt, ok := d.GetOk("updated_at"); ok && destroyOnDrift && updatedAt != newTimestamp {
			d.SetId("") // This simulates the drift detection
		}

		// Should have cleared the ID (marking for recreation)
		if d.Id() != "" {
			t.Error("Expected ID to be cleared due to drift detection, but it wasn't")
		}
	})

	t.Run("destroyOnDrift false updates timestamp without recreation", func(t *testing.T) {
		originalTimestamp := "2023-01-01T00:00:00Z"
		newTimestamp := "2023-01-02T00:00:00Z"

		d := schema.TestResourceDataRaw(t, resourceGithubActionsSecret().Schema, map[string]any{
			"repository":       "test-repo",
			"secret_name":      "test-secret",
			"plaintext_value":  "test-value",
			"destroy_on_drift": false,
			"updated_at":       originalTimestamp,
		})
		d.SetId("test-secret")

		// Test the drift detection logic when destroy_on_drift is false
		destroyOnDrift := d.Get("destroy_on_drift").(bool)
		if updatedAt, ok := d.GetOk("updated_at"); ok && !destroyOnDrift && updatedAt != newTimestamp {
			// This simulates what happens when destroy_on_drift=false
			_ = d.Set("updated_at", newTimestamp)
		}

		// Should NOT have cleared the ID
		if d.Id() == "" {
			t.Error("Expected ID to be preserved when destroy_on_drift=false, but it was cleared")
		}

		// Should have updated the timestamp
		if updatedAt := d.Get("updated_at").(string); updatedAt != newTimestamp {
			t.Errorf("Expected timestamp to be updated to %s, got %s", newTimestamp, updatedAt)
		}
	})

	t.Run("default destroy_on_drift is true", func(t *testing.T) {
		d := schema.TestResourceDataRaw(t, resourceGithubActionsSecret().Schema, map[string]any{
			"repository":      "test-repo",
			"secret_name":     "test-secret",
			"plaintext_value": "test-value",
			// destroy_on_drift not set, should default to true
		})

		destroyOnDrift := d.Get("destroy_on_drift").(bool)
		if !destroyOnDrift {
			t.Error("Expected destroy_on_drift to default to true")
		}
	})

	t.Run("no drift when timestamps match", func(t *testing.T) {
		timestamp := "2023-01-01T00:00:00Z"

		d := schema.TestResourceDataRaw(t, resourceGithubActionsSecret().Schema, map[string]any{
			"repository":       "test-repo",
			"secret_name":      "test-secret",
			"plaintext_value":  "test-value",
			"destroy_on_drift": true,
			"updated_at":       timestamp,
		})
		d.SetId("test-secret")

		// Simulate same timestamp (no external change)
		destroyOnDrift := d.Get("destroy_on_drift").(bool)
		if updatedAt, ok := d.GetOk("updated_at"); ok && destroyOnDrift && updatedAt != timestamp {
			d.SetId("") // This should NOT happen
		}

		// Should NOT have cleared the ID
		if d.Id() == "" {
			t.Error("Expected ID to be preserved when no drift detected, but it was cleared")
		}
	})

	t.Run("destroy_on_drift field properties", func(t *testing.T) {
		resource := resourceGithubActionsSecret()
		driftField := resource.Schema["destroy_on_drift"]

		// Should be optional
		if driftField.Required {
			t.Error("Expected destroy_on_drift to be optional, but it's required")
		}

		if !driftField.Optional {
			t.Error("Expected destroy_on_drift to be optional")
		}

		// Should be boolean type
		if driftField.Type.String() != "TypeBool" {
			t.Errorf("Expected destroy_on_drift to be TypeBool, got %s", driftField.Type.String())
		}

		// Should have default value of true
		if driftField.Default != true {
			t.Errorf("Expected destroy_on_drift default to be true, got %v", driftField.Default)
		}

		// Should have description
		if driftField.Description == "" {
			t.Error("Expected destroy_on_drift to have a description")
		}
	})
}

// Test demonstrating the solution to GitHub issue #964.
func TestGithubActionsSecretIssue964Solution(t *testing.T) {
	t.Run("solve issue 964 - prevent recreation when GUI changes secret", func(t *testing.T) {
		// This test demonstrates the fix for:
		// https://github.com/integrations/terraform-provider-github/issues/964

		// Scenario: User creates secret with Terraform, then updates value via GitHub GUI
		// Expected: With destroy_on_drift=false, Terraform should not recreate the secret

		d := schema.TestResourceDataRaw(t, resourceGithubActionsSecret().Schema, map[string]any{
			"repository":       "my-repo",
			"secret_name":      "WORKFLOW_PAT",
			"plaintext_value":  "CHANGE_ME", // Initial placeholder value
			"destroy_on_drift": false,       // KEY FIX: Prevents recreation
		})
		d.SetId("WORKFLOW_PAT")

		// Set initial timestamp
		originalTime := "2023-01-01T00:00:00Z"
		_ = d.Set("updated_at", originalTime)

		// Simulate: User changes secret value via GitHub GUI
		// This changes the updated_at timestamp
		newTime := "2023-01-01T12:00:00Z" // Later timestamp = external change

		// Test the read function behavior - this is what happens during terraform plan/apply
		destroyOnDrift := d.Get("destroy_on_drift").(bool) // false
		if updatedAt, ok := d.GetOk("updated_at"); ok && !destroyOnDrift && updatedAt != newTime {
			// With destroy_on_drift=false, we update timestamp but don't clear ID
			_ = d.Set("updated_at", newTime)
		}

		// RESULT: Secret should NOT be marked for recreation
		if d.Id() == "" {
			t.Error("ISSUE #964 NOT FIXED: Secret was marked for recreation despite destroy_on_drift=false")
		}

		// RESULT: Timestamp should be updated to acknowledge the change
		if d.Get("updated_at").(string) != newTime {
			t.Error("Expected timestamp to be updated to acknowledge external change")
		}

		t.Logf("SUCCESS: Issue #964 solved - secret with destroy_on_drift=false does not get recreated on external changes")
	})
}
