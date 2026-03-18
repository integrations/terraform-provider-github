package github

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
