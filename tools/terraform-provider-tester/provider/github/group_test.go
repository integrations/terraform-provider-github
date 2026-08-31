package github

import (
	"testing"
)

func TestGroupOfOddNames(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// The one explicitly hard-coded special case.
		{"TestAccOrganizationBlock_basic", "organization"},

		// TestAccDataSourceGithub* names (7 odd data-source names).
		{"TestAccDataSourceGithubRepository", "repositories"},
		{"TestAccDataSourceGithubOrganizationRole", "organization"},
		{"TestAccDataSourceGithubUser", "users"},
		{"TestAccDataSourceGithubTeam", "teams"},
		{"TestAccDataSourceGithubBranch", "branches"},
		{"TestAccDataSourceGithubApp", "apps"},
		{"TestAccDataSourceGithubIp", "ip-ranges"},

		// EMU token: "TestAccGithubEMUGroupMapping" -> stem "EMUGroupMapping" -> leading word "E" -> "enterprise".
		{"TestAccGithubEMUGroupMapping", "enterprise"},

		// Core hard requirement from brief.
		{"TestAccGithubRepository", "repositories"},

		// Standard names.
		{"TestAccGithubRepositoryBranch", "repositories"},
		{"TestAccGithubOrganizationMember", "organization"},
		{"TestAccGithubTeamMembership", "teams"},
		{"TestAccGithubUserGpgKey", "users"},
		{"TestAccGithubBranchProtection", "branches"},
		{"TestAccGithubIssueLabel", "issues"},
		{"TestAccGithubActionsSecret", "actions"},
		{"TestAccGithubAppInstallation", "apps"},
		{"TestAccGithubRelease", "releases"},
		{"TestAccGithubDependabotSecret", "dependabot"},
		{"TestAccGithubCodespacesSecret", "codespaces"},
		{"TestAccGithubEnterpriseOrganization", "enterprise"},

		// Names with no known prefix fall back to "misc".
		{"SomeRandomTest", "misc"},
		{"TestSomethingElse", "misc"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := groupOf(tt.name)
			if got != tt.want {
				t.Errorf("groupOf(%q) = %q, want %q", tt.name, got, tt.want)
			}
		})
	}
}
