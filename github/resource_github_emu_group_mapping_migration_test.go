package github

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_resourceGithubEMUGroupMappingStateUpgradeV1(t *testing.T) {
	t.Parallel()

	for _, d := range []struct {
		testName    string
		rawState    map[string]any
		want        map[string]any
		shouldError bool
	}{
		{
			testName: "migrates v1 to v2",
			rawState: map[string]any{
				"id": "123456:test-team:7765543",
			},
			want: map[string]any{
				"id": "7765543:123456",
			},
			shouldError: false,
		},
	} {
		t.Run(d.testName, func(t *testing.T) {
			t.Parallel()

			got, err := resourceGithubEMUGroupMappingStateUpgradeV1(t.Context(), d.rawState, nil)
			if (err != nil) != d.shouldError {
				t.Fatalf("unexpected error state")
			}

			if diff := cmp.Diff(got, d.want); !d.shouldError && diff != "" {
				t.Fatalf("got %+v, want %+v: %s", got, d.want, diff)
			}
		})
	}
}
