package github

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// TestEtagDiffSuppressFunction tests that the etag diff suppress function
// always returns true, suppressing all etag differences.
func TestEtagDiffSuppressFunction(t *testing.T) {
	repositoryResource := resourceGithubRepository()
	etagField := repositoryResource.Schema["etag"]

	if etagField == nil {
		t.Fatal("etag field not found in repository schema")
		panic("unreachable") // This resolves https://github.com/golangci/golangci-lint/issues/5979
	}

	if etagField.DiffSuppressFunc == nil {
		t.Fatal("etag should have DiffSuppressFunc")
	}

	if !etagField.DiffSuppressOnRefresh {
		t.Fatal("etag should have DiffSuppressOnRefresh enabled")
	}

	testCases := []struct {
		name string
		key  string
		old  string
		new  string
	}{
		{
			name: "different etag values",
			key:  "etag",
			old:  `"abc123"`,
			new:  `"def456"`,
		},
		{
			name: "empty to non-empty etag",
			key:  "etag",
			old:  "",
			new:  `"abc123"`,
		},
		{
			name: "non-empty to empty etag",
			key:  "etag",
			old:  `"abc123"`,
			new:  "",
		},
		{
			name: "same etag values",
			key:  "etag",
			old:  `"abc123"`,
			new:  `"abc123"`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			d := schema.TestResourceDataRaw(t, repositoryResource.Schema, map[string]any{
				"name": "test-repo",
			})

			result := etagField.DiffSuppressFunc(tc.key, tc.old, tc.new, d)
			if !result {
				t.Errorf("DiffSuppressFunc should always return true for etag field")
			}
		})
	}
}

// TestEtagSchemaConsistency ensure DiffSuppressFunc and DiffSuppressOnRefresh are consistently applied.
func TestEtagSchemaConsistency(t *testing.T) {
	resourcesWithEtag := map[string]*schema.Resource{
		"github_repository":                          resourceGithubRepository(),
		"github_branch":                              resourceGithubBranch(),
		"github_branch_default":                      resourceGithubBranchDefault(),
		"github_issue_label":                         resourceGithubIssueLabel(),
		"github_repository_webhook":                  resourceGithubRepositoryWebhook(),
		"github_repository_deployment_branch_policy": resourceGithubRepositoryDeploymentBranchPolicy(),
		"github_repository_project":                  resourceGithubRepositoryProject(),
	}

	for resourceName, resource := range resourcesWithEtag {
		t.Run(resourceName, func(t *testing.T) {
			etagField, exists := resource.Schema["etag"]
			if !exists {
				t.Errorf("Resource %s should have etag field", resourceName)
				return
			}

			// Verify etag is optional and computed
			if !etagField.Optional {
				t.Errorf("etag should be optional in %s", resourceName)
			}
			if !etagField.Computed {
				t.Errorf("etag should be computed in %s", resourceName)
			}

			// Verify etag has DiffSuppressFunc
			if etagField.DiffSuppressFunc == nil {
				t.Errorf("etag should have DiffSuppressFunc in %s", resourceName)
			}

			// Verify DiffSuppressOnRefresh is enabled
			if !etagField.DiffSuppressOnRefresh {
				t.Errorf("etag should have DiffSuppressOnRefresh enabled in %s", resourceName)
			}

			// Verify the DiffSuppressFunc always returns true
			if etagField.DiffSuppressFunc != nil {
				d := schema.TestResourceDataRaw(t, resource.Schema, map[string]any{})
				result := etagField.DiffSuppressFunc("etag", "old", "new", d)
				if !result {
					t.Errorf("DiffSuppressFunc should return true in %s", resourceName)
				}
			}
		})
	}
}
