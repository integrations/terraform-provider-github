package github

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestResourceGithubActionsSecretValidation(t *testing.T) {
	resource := resourceGithubActionsSecret()
	
	// Verify the resource has an Update function
	if resource.Update == nil {
		t.Fatal("github_actions_secret resource must have an Update function to handle destroy_on_drift field changes")
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
	
	if destroyOnDriftSchema.ForceNew {
		t.Error("destroy_on_drift should not be ForceNew when Update function exists")
	}
}