package github

import (
	"reflect"
	"testing"
)

func testResourceGithubActionsSecretInstanceStateDataV0() map[string]any {
	return map[string]any{
		"id":              "test-secret",
		"repository":      "test-repo",
		"secret_name":     "test-secret",
		"created_at":      "2023-01-01T00:00:00Z",
		"updated_at":      "2023-01-01T00:00:00Z",
		"plaintext_value": "secret-value",
	}
}

func testResourceGithubActionsSecretInstanceStateDataV0_WithDrift() map[string]any {
	v0 := testResourceGithubActionsSecretInstanceStateDataV0()
	v0["destroy_on_drift"] = false
	return v0
}

func testResourceGithubActionsSecretInstanceStateDataV1() map[string]any {
	v0 := testResourceGithubActionsSecretInstanceStateDataV0()
	v0["destroy_on_drift"] = true
	return v0
}

func TestGithub_MigrateActionsSecretStateV0toV1(t *testing.T) {
	t.Run("without destroy_on_drift", func(t *testing.T) {
		expected := testResourceGithubActionsSecretInstanceStateDataV1()
		actual, err := resourceGithubActionsSecretInstanceStateUpgradeV0(t.Context(), testResourceGithubActionsSecretInstanceStateDataV0(), nil)
		if err != nil {
			t.Fatalf("error migrating state: %s", err)
		}

		if !reflect.DeepEqual(expected, actual) {
			t.Fatalf("\n\nexpected:\n\n%#v\n\ngot:\n\n%#v\n\n", expected, actual)
		}
	})

	t.Run("with destroy_on_drift", func(t *testing.T) {
		expected := testResourceGithubActionsSecretInstanceStateDataV0_WithDrift()
		actual, err := resourceGithubActionsSecretInstanceStateUpgradeV0(t.Context(), testResourceGithubActionsSecretInstanceStateDataV0_WithDrift(), nil)
		if err != nil {
			t.Fatalf("error migrating state: %s", err)
		}

		if !reflect.DeepEqual(expected, actual) {
			t.Fatalf("\n\nexpected:\n\n%#v\n\ngot:\n\n%#v\n\n", expected, actual)
		}
	})
}
