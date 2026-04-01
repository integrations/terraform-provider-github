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
		"github_actions_runner_group":                resourceGithubActionsRunnerGroup(),
		"github_branch":                              resourceGithubBranch(),
		"github_branch_default":                      resourceGithubBranchDefault(),
		"github_branch_protection_v3":                resourceGithubBranchProtectionV3(),
		"github_emu_group_mapping":                   resourceGithubEMUGroupMapping(),
		"github_enterprise_actions_runner_group":     resourceGithubActionsEnterpriseRunnerGroup(),
		"github_issue":                               resourceGithubIssue(),
		"github_issue_label":                         resourceGithubIssueLabel(),
		"github_membership":                          resourceGithubMembership(),
		"github_organization_project":                resourceGithubOrganizationProject(),
		"github_organization_ruleset":                resourceGithubOrganizationRuleset(),
		"github_organization_webhook":                resourceGithubOrganizationWebhook(),
		"github_project_card":                        resourceGithubProjectCard(),
		"github_project_column":                      resourceGithubProjectColumn(),
		"github_release":                             resourceGithubRelease(),
		"github_repository":                          resourceGithubRepository(),
		"github_repository_autolink_reference":       resourceGithubRepositoryAutolinkReference(),
		"github_repository_deploy_key":               resourceGithubRepositoryDeployKey(),
		"github_repository_deployment_branch_policy": resourceGithubRepositoryDeploymentBranchPolicy(),
		"github_repository_project":                  resourceGithubRepositoryProject(),
		"github_repository_ruleset":                  resourceGithubRepositoryRuleset(),
		"github_repository_webhook":                  resourceGithubRepositoryWebhook(),
		"github_team":                                resourceGithubTeam(),
		"github_team_membership":                     resourceGithubTeamMembership(),
		"github_team_repository":                     resourceGithubTeamRepository(),
		"github_team_sync_group_mapping":             resourceGithubTeamSyncGroupMapping(),
		"github_user_gpg_key":                        resourceGithubUserGpgKey(),
		"github_user_ssh_key":                        resourceGithubUserSshKey(),
		"organization_block":                         resourceOrganizationBlock(),
	}

	for resourceName, resource := range resourcesWithEtag {
		t.Run(resourceName, func(t *testing.T) {
			etagField, exists := resource.Schema["etag"]
			if !exists {
				t.Errorf("Resource %s should have etag field", resourceName)
				return
			}

			// Verify etag is computed
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
