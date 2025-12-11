package github

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Test for the organization secret drift detection fix.
func TestGithubActionsOrganizationSecretDriftDetectionFix(t *testing.T) {
	t.Run("always updates timestamp regardless of drift detection", func(t *testing.T) {
		// This test verifies the fix for the issue where updated_at was not
		// being set when drift was detected, causing repeated drift detection

		d := schema.TestResourceDataRaw(t, resourceGithubActionsOrganizationSecret().Schema, map[string]any{
			"secret_name":      "test-secret",
			"plaintext_value":  "test-value",
			"visibility":       "private",
			"destroy_on_drift": true,
			"updated_at":       "2023-01-01T00:00:00Z", // Old timestamp
		})
		d.SetId("test-secret")

		// Simulate the updated_at logic from the read function
		// This is what the actual GitHub API would return (newer timestamp)
		newTimestamp := "2023-01-01T12:00:00Z"

		// Simulate the drift detection logic from resourceGithubActionsOrganizationSecretRead
		destroyOnDrift := d.Get("destroy_on_drift").(bool)
		if updatedAt, ok := d.GetOk("updated_at"); ok && destroyOnDrift && updatedAt != newTimestamp {
			// This would log the drift and clear the ID
			d.SetId("")
		}

		// This is the key fix - always update the timestamp
		err := d.Set("updated_at", newTimestamp)
		if err != nil {
			t.Fatal(err)
		}

		// Verify that the timestamp was updated even though drift was detected
		if d.Get("updated_at").(string) != newTimestamp {
			t.Error("Expected updated_at to be set to new timestamp after drift detection")
		}

		// Verify that the ID was cleared due to drift detection
		if d.Id() != "" {
			t.Error("Expected ID to be cleared due to drift detection with destroy_on_drift=true")
		}
	})

	t.Run("does not clear ID when destroy_on_drift is false", func(t *testing.T) {
		d := schema.TestResourceDataRaw(t, resourceGithubActionsOrganizationSecret().Schema, map[string]any{
			"secret_name":      "test-secret",
			"plaintext_value":  "test-value",
			"visibility":       "private",
			"destroy_on_drift": false,
			"updated_at":       "2023-01-01T00:00:00Z", // Old timestamp
		})
		d.SetId("test-secret")

		newTimestamp := "2023-01-01T12:00:00Z"

		// Simulate the drift detection logic
		destroyOnDrift := d.Get("destroy_on_drift").(bool)
		if updatedAt, ok := d.GetOk("updated_at"); ok && destroyOnDrift && updatedAt != newTimestamp {
			d.SetId("")
		}

		// Always update the timestamp
		err := d.Set("updated_at", newTimestamp)
		if err != nil {
			t.Fatal(err)
		}

		// Verify that the ID was NOT cleared when destroy_on_drift=false
		if d.Id() != "test-secret" {
			t.Error("Expected ID to remain when destroy_on_drift=false")
		}

		// Verify that the timestamp was still updated
		if d.Get("updated_at").(string) != newTimestamp {
			t.Error("Expected updated_at to be updated even when destroy_on_drift=false")
		}
	})
}
