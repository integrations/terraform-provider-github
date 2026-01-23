package github

import (
	"reflect"
	"testing"
)

func testResourceGithubRepositoryInstanceStateDataV0() map[string]any {
	return map[string]any{
		"branches.#":           "1",
		"branches.0.name":      "foobar",
		"branches.0.protected": "false",
	}
}

func testResourceGithubRepositoryInstanceStateDataV1() map[string]any {
	return map[string]any{}
}

func TestGithub_MigrateRepositoryStateV0toV1(t *testing.T) {
	t.Run("migrates state without errors", func(t *testing.T) {
		expected := testResourceGithubRepositoryInstanceStateDataV1()
		actual, err := resourceGithubRepositoryInstanceStateUpgradeV0(t.Context(), testResourceGithubRepositoryInstanceStateDataV0(), nil)
		if err != nil {
			t.Fatalf("error migrating state: %s", err)
		}
		if !reflect.DeepEqual(expected, actual) {
			t.Fatalf("\n\nexpected:\n\n%#v\n\ngot:\n\n%#v\n\n", expected, actual)
		}
	})
}
