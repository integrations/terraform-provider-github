package github

import (
	"reflect"
	"testing"
)

func testResourceGithubActionsOrganizationSecretInstanceStateDataV0() map[string]any {
	return map[string]any{
		"id":              "test-secret",
		"secret_name":     "test-secret",
		"visibility":      "private",
		"created_at":      "2023-01-01T00:00:00Z",
		"updated_at":      "2023-01-01T00:00:00Z",
		"plaintext_value": "secret-value",
	}
}

func testResourceGithubActionsOrganizationSecretInstanceStateDataV0_WithDrift() map[string]any {
	v0 := testResourceGithubActionsOrganizationSecretInstanceStateDataV0()
	v0["destroy_on_drift"] = false
	return v0
}

func testResourceGithubActionsOrganizationSecretInstanceStateDataV1() map[string]any {
	v0 := testResourceGithubActionsOrganizationSecretInstanceStateDataV0()
	v0["destroy_on_drift"] = true
	return v0
}

func TestGithub_MigrateActionsOrganizationSecretStateV0toV1(t *testing.T) {
	t.Run("without destroy_on_drift", func(t *testing.T) {
		expected := testResourceGithubActionsOrganizationSecretInstanceStateDataV1()
		actual, err := resourceGithubActionsOrganizationSecretInstanceStateUpgradeV0(t.Context(), testResourceGithubActionsOrganizationSecretInstanceStateDataV0(), nil)
		if err != nil {
			t.Fatalf("error migrating state: %s", err)
		}

		if !reflect.DeepEqual(expected, actual) {
			t.Fatalf("\n\nexpected:\n\n%#v\n\ngot:\n\n%#v\n\n", expected, actual)
		}
	})

	t.Run("with destroy_on_drift", func(t *testing.T) {
		expected := testResourceGithubActionsOrganizationSecretInstanceStateDataV0_WithDrift()
		actual, err := resourceGithubActionsOrganizationSecretInstanceStateUpgradeV0(t.Context(), testResourceGithubActionsOrganizationSecretInstanceStateDataV0_WithDrift(), nil)
		if err != nil {
			t.Fatalf("error migrating state: %s", err)
		}

		if !reflect.DeepEqual(expected, actual) {
			t.Fatalf("\n\nexpected:\n\n%#v\n\ngot:\n\n%#v\n\n", expected, actual)
		}
	})
}
