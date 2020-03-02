package github

import (
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestMigrateGithubWebhookStateV0toV1(t *testing.T) {
	oldAttributes := map[string]string{
		"configuration.%":            "4",
		"configuration.content_type": "form",
		"configuration.insecure_ssl": "0",
		"configuration.secret":       "blablah",
		"configuration.url":          "https://google.co.uk/",
	}

	newState, err := migrateGithubWebhookStateV0toV1(&terraform.InstanceState{
		ID:         "nonempty",
		Attributes: oldAttributes,
	})
	if err != nil {
		t.Fatal(err)
	}

	expectedAttributes := map[string]string{
		"configuration.#":              "1",
		"configuration.0.content_type": "form",
		"configuration.0.insecure_ssl": "0",
		"configuration.0.secret":       "blablah",
		"configuration.0.url":          "https://google.co.uk/",
	}
	if !reflect.DeepEqual(newState.Attributes, expectedAttributes) {
		t.Fatalf("Expected attributes:\n%#v\n\nGiven:\n%#v\n",
			expectedAttributes, newState.Attributes)
	}
}
