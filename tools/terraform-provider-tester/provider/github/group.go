package github

import (
	"regexp"
	"strings"
)

var leadingWordRE = regexp.MustCompile(`^[A-Z][a-z0-9]*`)

// groupOf returns the test group bucket for a TestAcc test name.
// It mirrors the logic needed to classify all 173 pinned acceptance test names.
func groupOf(name string) string {
	// Explicit special case - this name has no TestAccGithub* prefix.
	if name == "TestAccOrganizationBlock_basic" {
		return "organization"
	}

	// Strip one of the two known prefixes to get the resource stem.
	var res string
	var found bool
	if res, found = strings.CutPrefix(name, "TestAccDataSourceGithub"); !found {
		res, found = strings.CutPrefix(name, "TestAccGithub")
		if !found {
			return "misc"
		}
	}

	token := leadingWordRE.FindString(res)
	if token == "" {
		return "misc"
	}
	return bucketOf(token)
}

// bucketOf maps a leading CamelCase token to its group bucket.
// The 24 tokens listed here are the complete current set from the pinned 173-test catalog.
// Any unlisted token falls back to its lowercased form.
func bucketOf(token string) string {
	switch token {
	case "Repository", "Repositories", "Collaborators", "Ref", "Tree":
		return "repositories"
	case "User", "Users", "Ssh":
		return "users"
	case "Membership", "Organization":
		return "organization"
	case "Workflow", "Actions":
		return "actions"
	case "E", "Enterprise": // "E" comes from EMUGroupMapping -> leading word "E"
		return "enterprise"
	case "Ip":
		return "ip-ranges"
	case "Project":
		return "projects"
	case "Rest":
		return "rest-api"
	case "Team":
		return "teams"
	case "Branch":
		return "branches"
	case "Issue":
		return "issues"
	case "App":
		return "apps"
	case "Release":
		return "releases"
	case "Codespaces":
		return "codespaces"
	case "Dependabot":
		return "dependabot"
	default:
		return strings.ToLower(token)
	}
}
