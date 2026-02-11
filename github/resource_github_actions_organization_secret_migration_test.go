package github

import (
	"context"
	"reflect"
	"testing"
)

func Test_resourceGithubActionsOrganizationSecretStateUpgradeV0(t *testing.T) {
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
				"id":              "test-secret",
				"secret_name":     "test-secret",
				"visibility":      "private",
				"created_at":      "2023-01-01T00:00:00Z",
				"updated_at":      "2023-01-01T00:00:00Z",
				"plaintext_value": "secret-value",
			},
			want: map[string]any{
				"id":               "test-secret",
				"secret_name":      "test-secret",
				"visibility":       "private",
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
				"id":               "test-secret",
				"secret_name":      "test-secret",
				"visibility":       "private",
				"created_at":       "2023-01-01T00:00:00Z",
				"updated_at":       "2023-01-01T00:00:00Z",
				"plaintext_value":  "secret-value",
				"destroy_on_drift": false,
			},
			want: map[string]any{
				"id":               "test-secret",
				"secret_name":      "test-secret",
				"visibility":       "private",
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

			got, err := resourceGithubActionsOrganizationSecretStateUpgradeV0(context.Background(), d.rawState, nil)
			if (err != nil) != d.shouldError {
				t.Fatalf("unexpected error state")
			}

			if !d.shouldError && !reflect.DeepEqual(got, d.want) {
				t.Fatalf("got %+v, want %+v", got, d.want)
			}
		})
	}
}
