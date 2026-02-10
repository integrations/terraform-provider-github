package github

import (
	"context"
	"reflect"
	"testing"
)

func Test_resourceGithubActionsSecretStateUpgradeV0(t *testing.T) {
	t.Parallel()

	for _, d := range []struct {
		testName    string
		rawState    map[string]any
		want        map[string]any
		shouldError bool
	}{
		{
			testName: "migrates_v0_to_v1",
			rawState: map[string]any{
				"id":              "test-repo:test-secret",
				"repository":      "test-repo",
				"secret_name":     "test-secret",
				"created_at":      "2023-01-01T00:00:00Z",
				"updated_at":      "2023-01-01T00:00:00Z",
				"plaintext_value": "secret-value",
			},
			want: map[string]any{
				"id":               "test-repo:test-secret",
				"repository":       "test-repo",
				"secret_name":      "test-secret",
				"created_at":       "2023-01-01T00:00:00Z",
				"updated_at":       "2023-01-01T00:00:00Z",
				"plaintext_value":  "secret-value",
				"destroy_on_drift": true,
			},
			shouldError: false,
		},
		{
			testName: "migrates_v0_to_v1_with_existing_destroy_on_drift",
			rawState: map[string]any{
				"id":               "test-repo:test-secret",
				"repository":       "test-repo",
				"secret_name":      "test-secret",
				"created_at":       "2023-01-01T00:00:00Z",
				"updated_at":       "2023-01-01T00:00:00Z",
				"plaintext_value":  "secret-value",
				"destroy_on_drift": false,
			},
			want: map[string]any{
				"id":               "test-repo:test-secret",
				"repository":       "test-repo",
				"secret_name":      "test-secret",
				"created_at":       "2023-01-01T00:00:00Z",
				"updated_at":       "2023-01-01T00:00:00Z",
				"plaintext_value":  "secret-value",
				"destroy_on_drift": false,
			},
			shouldError: false,
		},
	} {
		t.Run(d.testName, func(t *testing.T) {
			t.Parallel()

			got, err := resourceGithubActionsSecretStateUpgradeV0(context.Background(), d.rawState, nil)
			if (err != nil) != d.shouldError {
				t.Fatalf("unexpected error state")
			}

			if !d.shouldError && !reflect.DeepEqual(got, d.want) {
				t.Fatalf("got %+v, want %+v", got, d.want)
			}
		})
	}
}

// TODO: Enable this test once we have a pattern to create a mock client for the test.
// func Test_resourceGithubActionsSecretStateUpgradeV1(t *testing.T) {
// 	t.Parallel()

// 	for _, d := range []struct {
// 		testName    string
// 		rawState    map[string]any
// 		want        map[string]any
// 		shouldError bool
// 	}{
// 		{
// 			testName: "migrates v1 to v2",
// 			rawState: map[string]any{
//				"id":               "test-repo:test-secret",
//				"repository":       "test-repo",
//				"secret_name":      "test-secret",
//				"created_at":       "2023-01-01T00:00:00Z",
//				"updated_at":       "2023-01-01T00:00:00Z",
//				"plaintext_value":  "secret-value",
//				"destroy_on_drift": true,
// 			},
// 			want: map[string]any{
//				"id":               "test-repo:test-secret",
//				"repository":       "test-repo",
//				"repository_id":    "123456",
//				"secret_name":      "test-secret",
//				"created_at":       "2023-01-01T00:00:00Z",
//				"updated_at":       "2023-01-01T00:00:00Z",
//				"plaintext_value":  "secret-value",
//				"destroy_on_drift": true,
// 			},
// 			shouldError: false,
// 		},
// 	} {
// 		t.Run(d.testName, func(t *testing.T) {
// 			t.Parallel()

// 			got, err := resourceGithubActionsSecretStateUpgradeV1(context.Background(), d.rawState, nil)
// 			if (err != nil) != d.shouldError {
// 				t.Fatalf("unexpected error state")
// 			}

// 			if !d.shouldError && !reflect.DeepEqual(got, d.want) {
// 				t.Fatalf("got %+v, want %+v", got, d.want)
// 			}
// 		})
// 	}
// }
