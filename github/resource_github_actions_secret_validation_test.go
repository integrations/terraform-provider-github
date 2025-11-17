package github

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestResourceGithubActionsSecretValidation(t *testing.T) {
	resource := resourceGithubActionsSecret()

	// Verify the resource does NOT have an Update function since all fields are ForceNew
	if resource.UpdateContext != nil || resource.UpdateWithoutTimeout != nil {
		t.Fatal("github_actions_secret resource should not have an Update function when all fields are ForceNew")
	}

	// Verify destroy_on_drift field exists and is configured correctly
	destroyOnDriftSchema, exists := resource.Schema["destroy_on_drift"]
	if !exists {
		t.Fatal("destroy_on_drift field should exist in schema")
	}

	if destroyOnDriftSchema.Type != schema.TypeBool {
		t.Error("destroy_on_drift should be TypeBool")
	}

	if !destroyOnDriftSchema.Optional {
		t.Error("destroy_on_drift should be Optional")
	}

	if !destroyOnDriftSchema.ForceNew {
		t.Error("destroy_on_drift should be ForceNew when no Update function exists")
	}

	// Verify all user-configurable fields are ForceNew (which is why Update is unnecessary)
	expectedForceNewFields := []string{"repository", "secret_name", "encrypted_value", "plaintext_value", "destroy_on_drift"}
	for _, fieldName := range expectedForceNewFields {
		field, exists := resource.Schema[fieldName]
		if !exists {
			continue // Skip fields that don't exist
		}
		if !field.Computed && (field.Required || field.Optional) && !field.ForceNew {
			t.Errorf("Field %s should have ForceNew: true since no Update function exists", fieldName)
		}
	}
}
