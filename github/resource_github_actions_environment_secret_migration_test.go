package github

// TODO: Enable this test once we have a pattern to create a mock client for the test.

// import (
// 	"context"
// 	"reflect"
// 	"testing"
// )

// func Test_resourceGithubActionsEnvironmentSecretStateUpgradeV0(t *testing.T) {
// 	t.Parallel()

// 	for _, d := range []struct {
// 		testName    string
// 		rawState    map[string]any
// 		want        map[string]any
// 		shouldError bool
// 	}{
// 		{
// 			testName: "migrates v0 to v1",
// 			rawState: map[string]any{
// 				"id":              "my-repo:my-environment:MY_SECRET",
// 				"repository":      "my-repo",
// 				"environment":     "my-environment",
// 				"secret_name":     "MY_SECRET",
// 				"plaintext_value": "my-value",
// 			},
// 			want: map[string]any{
// 				"id":              "my-repo:my-environment:MY_SECRET",
// 				"repository":      "my-repo",
//				"repository_id":   123456,
// 				"environment":     "my-environment",
// 				"secret_name":     "MY_SECRET",
// 				"plaintext_value": "my-value",
// 			},
// 			shouldError: false,
// 		},
// 	} {
// 		t.Run(d.testName, func(t *testing.T) {
// 			t.Parallel()

// 			got, err := resourceGithubActionsEnvironmentSecretStateUpgradeV0(context.Background(), d.rawState, nil)
// 			if (err != nil) != d.shouldError {
// 				t.Fatalf("unexpected error state")
// 			}

// 			if !d.shouldError && !reflect.DeepEqual(got, d.want) {
// 				t.Fatalf("got %+v, want %+v", got, d.want)
// 			}
// 		})
// 	}
// }
