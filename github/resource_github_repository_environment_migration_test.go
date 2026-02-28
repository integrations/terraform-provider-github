package github

// TODO: Enable this test once we have a pattern to create a mock client for the test.

// import (
// 	"testing"

// 	"github.com/google/go-cmp/cmp"
// )

// func Test_resourceGithubRepositoryEnvironmenStateUpgradeV0(t *testing.T) {
// 	t.Parallel()

// 	for _, d := range []struct {
// 		testName    string
// 		rawState    map[string]any
// 		want        map[string]any
// 		shouldError bool
// 	}{
// 		{
// 			testName: "valid",
// 			rawState: map[string]any{
// 				"id":             "my-repo:my-environment:123456",
// 				"repository":     "my-repo",
// 				"environment":    "my-environment",
// 			},
// 			want: map[string]any{
// 				"id":             "my-repo:my-environment:123456",
// 				"repository":     "my-repo",
// 				"repository_id":  123456,
// 				"environment":    "my-environment",
// 			},
// 			shouldError: false,
// 		},
// 		{
// 			testName: "invalid_resource_id",
// 			rawState: map[string]any{
// 				"id":             "my-repo",
// 				"repository":     "my-repo",
// 				"environment":    "my-environment",
// 			},
// 			want:        nil,
// 			shouldError: true,
// 		},
// 	} {
// 		t.Run(d.testName, func(t *testing.T) {
// 			t.Parallel()

// 			got, err := resourceGithubRepositoryEnvironmentStateUpgradeV0(t.Context(), d.rawState, nil)
// 			if (err != nil) != d.shouldError {
// 				t.Fatalf("unexpected error state: %v", err)
// 			}

// 			if !d.shouldError {
// 				if diff := cmp.Diff(d.want, got); diff != "" {
// 					t.Errorf("resourceGithubRepositoryEnvironmentStateUpgradeV0() mismatch (-want +got):\n%s", diff)
// 				}
// 			}
// 		})
// 	}
// }
