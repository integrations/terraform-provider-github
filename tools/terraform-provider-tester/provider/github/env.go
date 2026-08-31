package github

import "github.com/github/terraform-provider-tester/provider"

func modes() []provider.Mode {
	return []provider.Mode{
		{Name: "anonymous", Description: "No credentials; public endpoints only"},
		{Name: "individual", Description: "Personal account tests requiring GITHUB_OWNER and credentials"},
		{Name: "organization", Description: "Organization tests requiring org owner and credentials"},
		{Name: "team", Description: "Team tests; adds external user fixtures"},
		{Name: "enterprise", Description: "Enterprise tests; adds GITHUB_ENTERPRISE_SLUG"},
	}
}

// credentialVars are the auth env vars shared by all non-anonymous modes (token XOR app trio).
var credentialVars = []provider.EnvVar{
	{Key: "GITHUB_TOKEN", Required: false, Secret: true, Doc: "PAT; or use the GITHUB_APP_* trio"},
	{Key: "GITHUB_APP_ID", Required: false, Secret: false, Doc: "GitHub App ID (alternative to GITHUB_TOKEN)"},
	{Key: "GITHUB_APP_INSTALLATION_ID", Required: false, Secret: false, Doc: "GitHub App Installation ID"},
	{Key: "GITHUB_APP_PEM_FILE", Required: false, Secret: true, Doc: "GitHub App PEM content; use \\n for newlines"},
}

// envFor returns the required/optional env vars for a given mode.
func envFor(mode string) []provider.EnvVar {
	switch mode {
	case "anonymous":
		return nil
	case "individual":
		vars := []provider.EnvVar{
			{Key: "GITHUB_OWNER", Required: true, Doc: "GitHub owner (user or org)"},
			{Key: "GITHUB_USERNAME", Required: true, Doc: "GitHub username for individual tests"},
		}
		vars = append(vars, credentialVars...)
		vars = append(vars,
			provider.EnvVar{Key: "GH_TEST_ORG_TEMPLATE_REPOSITORY", Required: false, Doc: "Template repository for selected tests (optional)"},
			provider.EnvVar{Key: "GH_TEST_USER_REPOSITORY", Required: false, Doc: "User repository for tests (optional)"},
		)
		return vars
	case "organization":
		return orgVars()
	case "team":
		vars := orgVars()
		vars = append(vars,
			provider.EnvVar{Key: "GH_TEST_ORG_USER2", Required: true, Doc: "Second org member"},
			provider.EnvVar{Key: "GH_TEST_EXTERNAL_USER1", Required: true, Doc: "External user 1"},
			provider.EnvVar{Key: "GH_TEST_EXTERNAL_USER1_TOKEN", Required: true, Secret: true, Doc: "Token for external user 1"},
			provider.EnvVar{Key: "GH_TEST_EXTERNAL_USER2", Required: true, Doc: "External user 2"},
		)
		return vars
	case "enterprise":
		vars := orgVars()
		vars = append(vars,
			provider.EnvVar{Key: "GITHUB_ENTERPRISE_SLUG", Required: true, Doc: "Enterprise slug"},
			provider.EnvVar{Key: "GH_TEST_ENTERPRISE_IS_EMU", Required: false, Doc: "Set to 'true' for EMU enterprise"},
			provider.EnvVar{Key: "GH_TEST_ENTERPRISE_EMU_GROUP_ID", Required: false, Doc: "EMU group ID (optional)"},
			provider.EnvVar{Key: "GH_TEST_ADVANCED_SECURITY", Required: false, Doc: "Set to 'true' to enable advanced security tests"},
		)
		return vars
	default:
		return nil
	}
}

// orgVars returns the env vars common to organization, team, and enterprise modes.
func orgVars() []provider.EnvVar {
	vars := []provider.EnvVar{
		{Key: "GITHUB_OWNER", Required: true, Doc: "GitHub org owner"},
	}
	vars = append(vars, credentialVars...)
	vars = append(vars,
		provider.EnvVar{Key: "GH_TEST_ORG_USER1", Required: true, Doc: "First org member"},
		provider.EnvVar{Key: "GH_TEST_ORG_REPOSITORY", Required: true, Doc: "Test org repository"},
		provider.EnvVar{Key: "GH_TEST_ORG_TEMPLATE_REPOSITORY", Required: true, Doc: "Template repository in org"},
		provider.EnvVar{Key: "GH_TEST_ORG_SECRET_NAME", Required: true, Doc: "Secret name for org secret tests"},
		provider.EnvVar{Key: "GH_TEST_ORG_USER3", Required: false, Doc: "Third org member (optional)"},
		provider.EnvVar{Key: "GH_TEST_ORG_APP_INSTALLATION_ID", Required: false, Doc: "App installation ID for org (optional)"},
	)
	return vars
}
