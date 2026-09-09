package github

import (
	"context"
	"reflect"
	"testing"
)

func Test_resourceGithubUserSshKeyStateUpgradeV0(t *testing.T) {
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
				"id":    "123",
				"title": "test-key",
				"key":   "test-key-data",
				"url":   "test-url",
			},
			want: map[string]any{
				"id":     "123",
				"key_id": int64(123),
				"title":  "test-key",
				"key":    "test-key-data",
				"url":    "test-url",
			},
			shouldError: false,
		},
	} {
		t.Run(d.testName, func(t *testing.T) {
			t.Parallel()

			got, err := resourceGithubUserSshKeyStateUpgradeV0(context.Background(), d.rawState, nil)
			if (err != nil) != d.shouldError {
				t.Fatalf("unexpected error state")
			}

			if !d.shouldError && !reflect.DeepEqual(got, d.want) {
				t.Fatalf("got %+v, want %+v", got, d.want)
			}
		})
	}
}
