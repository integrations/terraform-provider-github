package github

import (
	"reflect"
	"testing"
)

func testResourceGithubWebhookInstanceStateDataV0() map[string]any {
	return map[string]any{
		"configuration.%":            "4",
		"configuration.content_type": "form",
		"configuration.insecure_ssl": "0",
		"configuration.secret":       "blablah",
		"configuration.url":          "https://google.co.uk/",
	}
}

func testResourceGithubWebhookInstanceStateDataV1() map[string]any {
	v0 := testResourceGithubWebhookInstanceStateDataV0()
	return map[string]any{
		"configuration.#":              "1",
		"configuration.0.content_type": v0["configuration.content_type"],
		"configuration.0.insecure_ssl": v0["configuration.insecure_ssl"],
		"configuration.0.secret":       v0["configuration.secret"],
		"configuration.0.url":          v0["configuration.url"],
	}
}

func TestGithub_MigrateRepositoryWebhookStateV0toV1(t *testing.T) {
	t.Run("migrates state without errors", func(t *testing.T) {
		expected := testResourceGithubWebhookInstanceStateDataV1()
		actual, err := resourceGithubRepositoryWebhookInstanceStateUpgradeV0(t.Context(), testResourceGithubWebhookInstanceStateDataV0(), nil)
		if err != nil {
			t.Fatalf("error migrating state: %s", err)
		}

		if !reflect.DeepEqual(expected, actual) {
			t.Fatalf("\n\nexpected:\n\n%#v\n\ngot:\n\n%#v\n\n", expected, actual)
		}
	})
}
