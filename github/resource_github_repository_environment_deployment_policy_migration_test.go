package github

import (
	"maps"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_resourceGithubRepositoryEnvironmentDeploymentPolicyStateUpgradeV0_malformedID(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name     string
		rawState map[string]any
	}{
		{
			name:     "missing id",
			rawState: map[string]any{"repository": "my-repo"},
		},
		{
			name:     "nil id",
			rawState: map[string]any{"id": nil, "repository": "my-repo"},
		},
		{
			name:     "non-string id",
			rawState: map[string]any{"id": 123, "repository": "my-repo"},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var apiCalls atomic.Int32
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				apiCalls.Add(1)
				w.WriteHeader(http.StatusInternalServerError)
			}))
			t.Cleanup(srv.Close)

			meta := &Owner{
				name:     "test-org",
				v3client: mustCreateTestGitHubClient(t, srv.URL),
			}
			before := maps.Clone(tt.rawState)

			got, err := resourceGithubRepositoryEnvironmentDeploymentPolicyStateUpgradeV0(t.Context(), tt.rawState, meta)
			if err == nil {
				t.Fatal("expected error, got nil")
			}
			if !strings.Contains(err.Error(), "id not found or is not a string") {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != nil {
				t.Fatalf("expected nil state, got %#v", got)
			}
			if diff := cmp.Diff(before, tt.rawState); diff != "" {
				t.Fatalf("expected input state unchanged (-want +got):\n%s", diff)
			}
			if gotCalls := apiCalls.Load(); gotCalls != 0 {
				t.Fatalf("expected no GitHub API calls, got %d", gotCalls)
			}
		})
	}
}

// TODO: Enable this test once we have a pattern to create a mock client for the test.

// import (
// 	"testing"

// 	"github.com/google/go-cmp/cmp"
// )

// func Test_resourceGithubRepositoryEnvironmentDeploymentPolicyStateUpgradeV0(t *testing.T) {
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
// 				"branch_pattern": "main-*",
// 			},
// 			want: map[string]any{
// 				"id":             "my-repo:my-environment:123456",
// 				"repository":     "my-repo",
// 				"repository_id":  123456,
// 				"environment":    "my-environment",
// 				"branch_pattern": "main-*",
// 				"policy_id":      123456,
// 			},
// 			shouldError: false,
// 		},
// 		{
// 			testName: "invalid_resource_id",
// 			rawState: map[string]any{
// 				"id":             "my-repo:my-environment",
// 				"repository":     "my-repo",
// 				"environment":    "my-environment",
// 				"branch_pattern": "main-*",
// 			},
// 			want:        nil,
// 			shouldError: true,
// 		},
// 		{
// 			testName: "invalid_policy_id",
// 			rawState: map[string]any{
// 				"id":             "my-repo:my-environment:abcdef",
// 				"repository":     "my-repo",
// 				"environment":    "my-environment",
// 				"branch_pattern": "main-*",
// 			},
// 			want:        nil,
// 			shouldError: true,
// 		},
// 	} {
// 		t.Run(d.testName, func(t *testing.T) {
// 		t.Parallel()

// 			t.Parallel()

// 			got, err := resourceGithubRepositoryEnvironmentDeploymentPolicyStateUpgradeV0(t.Context(), d.rawState, nil)
// 			if (err != nil) != d.shouldError {
// 				t.Fatalf("unexpected error state: %v", err)
// 			}

// 			if !d.shouldError {
// 				if diff := cmp.Diff(d.want, got); diff != "" {
// 					t.Errorf("resourceGithubRepositoryEnvironmentDeploymentPolicyStateUpgradeV0() mismatch (-want +got):\n%s", diff)
// 				}
// 			}
// 		})
// 	}
// }
