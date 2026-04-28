package github

import (
	"reflect"
	"testing"
)

func testGithubBranchProtectionStateDataV1() map[string]any {
	return map[string]any{
		"blocks_creations":  true,
		"push_restrictions": [...]string{"/example-user"},
	}
}

func testGithubBranchProtectionStateDataV2() map[string]any {
	restrictions := []any{map[string]any{
		"blocks_creations": true,
		"push_allowances":  [...]string{"/example-user"},
	}}
	return map[string]any{
		"restrict_pushes": restrictions,
	}
}

func Test_resourceGithubBranchProtectionStateUpgradeV1(t *testing.T) {
	expected := testGithubBranchProtectionStateDataV2()
	actual, err := resourceGithubBranchProtectionUpgradeV1(t.Context(), testGithubBranchProtectionStateDataV1(), nil)
	if err != nil {
		t.Fatalf("error migrating state: %s", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("\n\nexpected:\n\n%#v\n\ngot:\n\n%#v\n\n", expected, actual)
	}
}
