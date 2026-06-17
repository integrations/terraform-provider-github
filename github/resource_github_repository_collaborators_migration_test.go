package github

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// TODO: Enable this test once we have a pattern to create a mock client for the test.
// func Test_resourceGithubRepositoryCollaboratorsStateUpgradeV0(t *testing.T) {
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
//				"id":               "test-repo",
//				"repository":       "test-repo",
// 			},
// 			want: map[string]any{
//				"id":               "123456",
//				"repository":       "test-repo",
//				"repository_id":    "123456",
// 			},
// 			shouldError: false,
// 		},
// 	} {
// 		t.Run(d.testName, func(t *testing.T) {
// 			t.Parallel()

// 			got, err := resourceGithubRepositoryCollaboratorsStateUpgradeV0(context.Background(), d.rawState, nil)
// 			if (err != nil) != d.shouldError {
// 				t.Fatalf("unexpected error state")
// 			}

// 			if !d.shouldError && !reflect.DeepEqual(got, d.want) {
// 				t.Fatalf("got %+v, want %+v", got, d.want)
// 			}
// 		})
// 	}
// }

func Test_resourceGithubRepositoryCollaboratorsStateUpgradeV1(t *testing.T) {
	t.Parallel()

	for _, d := range []struct {
		testName string
		meta     *Owner
		rawState map[string]any
		want     map[string]any
	}{
		{
			testName: "organization_repository",
			meta:     &Owner{name: "test", IsOrganization: true},
			rawState: map[string]any{
				"id":            "test-repo",
				"repository":    "test-repo",
				"repository_id": "123456",
			},
			want: map[string]any{
				"id":               "test-repo",
				"repository":       "test-repo",
				"repository_id":    "123456",
				"owner_configured": false,
			},
		},
		{
			testName: "personal_repository_owner_configured",
			meta:     &Owner{name: "test", IsOrganization: false},
			rawState: map[string]any{
				"id":            "test-repo",
				"repository":    "test-repo",
				"repository_id": "123456",
				"user": []any{
					map[string]any{
						"username": "test",
					},
				},
			},
			want: map[string]any{
				"id":            "test-repo",
				"repository":    "test-repo",
				"repository_id": "123456",
				"user": []any{
					map[string]any{
						"username": "test",
					},
				},
				"owner_configured": true,
			},
		},
		{
			testName: "personal_repository_owner_not_configured",
			meta:     &Owner{name: "test", IsOrganization: false},
			rawState: map[string]any{
				"id":            "test-repo",
				"repository":    "test-repo",
				"repository_id": "123456",
				"user": []any{
					map[string]any{
						"username": "other-user",
					},
				},
			},
			want: map[string]any{
				"id":            "test-repo",
				"repository":    "test-repo",
				"repository_id": "123456",
				"user": []any{
					map[string]any{
						"username": "other-user",
					},
				},
				"owner_configured": false,
			},
		},
	} {
		t.Run(d.testName, func(t *testing.T) {
			t.Parallel()

			got, _ := resourceGithubRepositoryCollaboratorsStateUpgradeV1(t.Context(), d.rawState, d.meta)

			if diff := cmp.Diff(got, d.want); diff != "" {
				t.Fatalf("got %+v, want %+v, diff %s", got, d.want, diff)
			}
		})
	}
}
