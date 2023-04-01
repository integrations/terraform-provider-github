package github

import (
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestMigrateGithubRepositoryStateV0toV1(t *testing.T) {
	oldAttributes := map[string]string{
		"branches.#":           "1",
		"branches.0.name":      "foobar",
		"branches.0.protected": "false",
	}

	newState, err := migrateGithubRepositoryStateV0toV1(&terraform.InstanceState{
		ID:         "nonempty",
		Attributes: oldAttributes,
	})
	if err != nil {
		t.Fatal(err)
	}

	expectedAttributes := map[string]string{}
	if !reflect.DeepEqual(newState.Attributes, expectedAttributes) {
		t.Fatalf("Expected attributes:\n%#v\n\nGiven:\n%#v\n",
			expectedAttributes, newState.Attributes)
	}
}
