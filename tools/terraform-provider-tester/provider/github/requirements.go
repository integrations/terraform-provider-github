package github

import (
	"strings"

	"github.com/github/terraform-provider-tester/provider"
)

// requirementsFor implements the ordered rule catalog for GitHub acceptance
// test requirements. Rules are applied in the numbered order from the brief.
// It returns (requirements, true) for every known test and
// (conservativeRequirements(), false) for unknown tests so the planner can
// fail closed unless explicitly overridden.
func requirementsFor(testName string) (provider.TestRequirements, bool) {
	// Rule 1: IP ranges data source — all five modes, no auth requirements.
	if testName == "TestAccGithubIpRangesDataSource" {
		return provider.TestRequirements{
			Modes: []string{"anonymous", "individual", "organization", "team", "enterprise"},
		}, true
	}

	// Rule 2: Enterprise and EMU tests — enterprise only. A small exact-name
	// exception keeps the external-groups data source in this family even though
	// its name does not include Enterprise or EMU.
	if strings.Contains(testName, "Enterprise") ||
		strings.Contains(testName, "EMU") ||
		testName == "TestAccDataSourceGithubExternalGroups" {
		req := provider.TestRequirements{
			Modes:  []string{"enterprise"},
			Scopes: []string{"repo", "delete_repo", "read:org", "admin:org", "admin:enterprise"},
		}
		if !isDataSource(testName) {
			req.SideEffects = sideEffectsFor(testName)
		}
		return req, true
	}

	// Rule 3: Team and enterprise tests — names beginning with the three
	// known team-scope prefixes.
	if isTeamTest(testName) {
		req := provider.TestRequirements{
			Modes:        []string{"team", "enterprise"},
			Scopes:       []string{"repo", "delete_repo", "read:org", "admin:org"},
			Capabilities: []string{"organization"},
		}
		if !isDataSource(testName) {
			req.SideEffects = sideEffectsFor(testName)
		}
		return req, true
	}

	// Rule 4: Organization tests — names containing "Organization", repository
	// custom-property tests, and organization-only branch-protection/ruleset cases.
	if isOrgTest(testName) {
		req := provider.TestRequirements{
			Modes:        []string{"organization", "team", "enterprise"},
			Scopes:       []string{"repo", "delete_repo", "read:org", "admin:org"},
			Capabilities: []string{"organization"},
		}
		if !isDataSource(testName) {
			req.SideEffects = sideEffectsFor(testName)
		}
		return req, true
	}

	// Rule 5: User SSH/GPG and user Codespaces exceptions — individual only
	// with exact scopes. These use an exact-name map.
	if req, ok := userExceptions[testName]; ok {
		return cloneReqs(req), true
	}

	// Any test name that does not carry a known provider prefix is genuinely
	// unknown; return the conservative fallback so the planner fails closed.
	if !strings.HasPrefix(testName, "TestAccGithub") &&
		!strings.HasPrefix(testName, "TestAccDataSourceGithub") {
		return conservativeRequirements(), false
	}

	// Rule 6: Remaining repository/user tests — individual, organization, team,
	// enterprise with base scopes repo, delete_repo, and user.
	req := provider.TestRequirements{
		Modes:  []string{"individual", "organization", "team", "enterprise"},
		Scopes: []string{"repo", "delete_repo", "user"},
	}

	// Rule 7: Template-repository tests add the template-repository capability.
	if isTemplateRepositoryTest(testName) {
		req.Capabilities = []string{"template-repository"}
	}

	// Rule 8: Resource tests add stable side effects; data sources add none.
	if !isDataSource(testName) {
		req.SideEffects = sideEffectsFor(testName)
	}

	return req, true
}

// conservativeRequirements returns the fail-closed fallback for tests that are
// not present in the catalog. The returned value is intentionally over-broad so
// the planner can never skip a requirement due to a missing entry.
func conservativeRequirements() provider.TestRequirements {
	return provider.TestRequirements{
		Modes: []string{"individual", "organization", "team", "enterprise"},
		Scopes: []string{
			"repo", "delete_repo", "user", "admin:public_key",
			"admin:gpg_key", "codespace", "read:org", "admin:org",
			"admin:enterprise",
		},
		Capabilities: nil,
		SideEffects:  []string{"unknown"},
	}
}

// isDataSource reports whether a test name refers to a data-source acceptance
// test. Standard data sources contain "DataSource" in their name; tests starting
// with "TestAccDataSource" are also data sources. A small set of non-standard
// data-source tests is handled explicitly.
func isDataSource(name string) bool {
	if strings.Contains(name, "DataSource") {
		return true
	}
	if strings.HasPrefix(name, "TestAccDataSource") {
		return true
	}
	// Non-standard data-source test names that do not contain "DataSource":
	switch name {
	case "TestAccGithubOrganizationAppInstallations",
		"TestAccGithubOrganizationExternalIdentities",
		"TestAccGithubRepositoryDeploymentBranchPolicies",
		"TestAccGithubRepositoryEnvironmentDeploymentPolicies",
		"TestAccGithubUserExternalIdentity":
		return true
	}
	return false
}

// isTeamTest reports whether a test name belongs to rule 3: team-scoped tests
// that run in team and enterprise modes only, including the exact refreshed
// team data sources whose names do not start with TestAccGithubTeam.
func isTeamTest(name string) bool {
	return strings.HasPrefix(name, "TestAccGithubTeam") ||
		name == "TestAccDataSourceGithubTeamMembers" ||
		name == "TestAccDataSourceGithubTeamRepositories" ||
		strings.HasPrefix(name, "TestAccDataSourceGithubOrganizationRoleTeam") ||
		strings.HasPrefix(name, "TestAccGithubOrganizationTeamSync")
}

// isOrgTest reports whether a test name belongs to rule 4: organization-scoped
// tests that run in organization, team, and enterprise modes.
func isOrgTest(name string) bool {
	return strings.Contains(name, "Organization") ||
		strings.Contains(name, "CustomPropert")
}

// isTemplateRepositoryTest reports whether a top-level acceptance test actually
// exercises configured template-repository behavior. The pinned provider source
// does that in the repository resource and repository data-source top-level
// tests; OIDC subject-claim customization template tests are unrelated.
func isTemplateRepositoryTest(name string) bool {
	switch name {
	case "TestAccGithubRepository", "TestAccDataSourceGithubRepository":
		return true
	default:
		return false
	}
}

// orgSettingResources is the set of resource test names whose primary side
// effect is changing an organization-level setting rather than a repository.
var orgSettingResources = map[string]bool{
	"TestAccGithubActionsOrganizationPermissions":         true,
	"TestAccGithubActionsOrganizationWorkflowPermissions": true,
	"TestAccGithubOrganizationSettings":                   true,
	"TestAccGithubActionsEnterprisePermissions":           true,
	"TestAccGithubEnterpriseActionsWorkflowPermissions":   true,
	"TestAccGithubEnterpriseSecurityAnalysisSettings":     true,
}

// sideEffectsFor returns the ordered set of stable side effects for a resource
// test. The caller must have already verified that testName is not a data source.
//
// Rule 8: every resource baseline includes "repository". Additional side
// effects are appended when the name matches a more specific resource kind:
//   - "team" for tests starting with TestAccGithubTeam
//   - "organization-setting" for known org-setting resources
//   - "organization-custom-property" for custom-property resources
func sideEffectsFor(name string) []string {
	effects := []string{"repository"}

	if strings.HasPrefix(name, "TestAccGithubTeam") {
		effects = append(effects, "team")
	}
	if orgSettingResources[name] {
		effects = append(effects, "organization-setting")
	}
	if strings.Contains(name, "CustomPropert") {
		effects = append(effects, "organization-custom-property")
	}
	return effects
}

// userExceptions contains the exact requirements for individual-only user
// SSH/GPG key and user Codespaces tests (rule 5). Values are stored without
// slice aliases so cloneReqs always returns fresh slices.
var userExceptions = map[string]provider.TestRequirements{
	// User SSH key — individual only, admin:public_key scope.
	"TestAccGithubUserSshKey": {
		Modes:       []string{"individual"},
		Scopes:      []string{"admin:public_key"},
		SideEffects: []string{"user-key"},
	},
	// User GPG key — individual only, admin:gpg_key scope.
	"TestAccGithubUserGpgKey": {
		Modes:       []string{"individual"},
		Scopes:      []string{"admin:gpg_key"},
		SideEffects: []string{"user-key"},
	},
	// User Codespaces secret resource.
	"TestAccGithubCodespacesUserSecret": {
		Modes:        []string{"individual"},
		Scopes:       []string{"codespace"},
		Capabilities: []string{"codespaces-user-secrets"},
		SideEffects:  []string{"codespaces-secret"},
	},
	// User Codespaces secrets data source (no side effects).
	"TestAccGithubCodespacesUserSecretsDataSource": {
		Modes:        []string{"individual"},
		Scopes:       []string{"codespace"},
		Capabilities: []string{"codespaces-user-secrets"},
	},
	// User Codespaces public-key data source (no side effects).
	"TestAccGithubCodespacesUserPublicKeyDataSource": {
		Modes:        []string{"individual"},
		Scopes:       []string{"codespace"},
		Capabilities: []string{"codespaces-user-secrets"},
	},
}

// cloneReqs returns a deep copy of r so callers cannot mutate catalog data.
func cloneReqs(r provider.TestRequirements) provider.TestRequirements {
	return provider.TestRequirements{
		Modes:        cloneSlice(r.Modes),
		Scopes:       cloneSlice(r.Scopes),
		Capabilities: cloneSlice(r.Capabilities),
		SideEffects:  cloneSlice(r.SideEffects),
	}
}

func cloneSlice(s []string) []string {
	if s == nil {
		return nil
	}
	out := make([]string, len(s))
	copy(out, s)
	return out
}
