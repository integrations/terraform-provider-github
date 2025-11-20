package github

import (
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestMigrateGithubActionsSecretStateV0toV1(t *testing.T) {
	// Secret without destroy_on_drift should get default value
	oldAttributes := map[string]string{
		"id":              "test-secret",
		"repository":      "test-repo",
		"secret_name":     "test-secret",
		"created_at":      "2023-01-01T00:00:00Z",
		"updated_at":      "2023-01-01T00:00:00Z",
		"plaintext_value": "secret-value",
	}

	newState, err := migrateGithubActionsSecretStateV0toV1(&terraform.InstanceState{
		ID:         "test-secret",
		Attributes: oldAttributes,
	})
	if err != nil {
		t.Fatal(err)
	}

	expectedAttributes := map[string]string{
		"id":               "test-secret",
		"repository":       "test-repo",
		"secret_name":      "test-secret",
		"created_at":       "2023-01-01T00:00:00Z",
		"updated_at":       "2023-01-01T00:00:00Z",
		"plaintext_value":  "secret-value",
		"destroy_on_drift": "true",
	}
	if !reflect.DeepEqual(newState.Attributes, expectedAttributes) {
		t.Fatalf("Expected attributes:\n%#v\n\nGiven:\n%#v\n",
			expectedAttributes, newState.Attributes)
	}

	// Secret with existing destroy_on_drift should be preserved
	oldAttributesWithDrift := map[string]string{
		"id":               "test-secret",
		"repository":       "test-repo",
		"secret_name":      "test-secret",
		"destroy_on_drift": "false",
	}

	newState2, err := migrateGithubActionsSecretStateV0toV1(&terraform.InstanceState{
		ID:         "test-secret",
		Attributes: oldAttributesWithDrift,
	})
	if err != nil {
		t.Fatal(err)
	}

	expectedAttributesWithDrift := map[string]string{
		"id":               "test-secret",
		"repository":       "test-repo",
		"secret_name":      "test-secret",
		"destroy_on_drift": "false",
	}
	if !reflect.DeepEqual(newState2.Attributes, expectedAttributesWithDrift) {
		t.Fatalf("Expected attributes:\n%#v\n\nGiven:\n%#v\n",
			expectedAttributesWithDrift, newState2.Attributes)
	}
}
