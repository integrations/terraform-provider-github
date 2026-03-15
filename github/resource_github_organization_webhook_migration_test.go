package github

import (
	"reflect"
	"testing"
)

func TestGithub_MigrateOrganizationWebhookStateV0toV1(t *testing.T) {
	t.Run("migrates state without errors", func(t *testing.T) {
		expected := testResourceGithubWebhookInstanceStateDataV1()
		actual, err := resourceGithubOrganizationWebhookInstanceStateUpgradeV0(t.Context(), testResourceGithubWebhookInstanceStateDataV0(), nil)
		if err != nil {
			t.Fatalf("error migrating state: %s", err)
		}

		if !reflect.DeepEqual(expected, actual) {
			t.Fatalf("\n\nexpected:\n\n%#v\n\ngot:\n\n%#v\n\n", expected, actual)
		}
	})
}
