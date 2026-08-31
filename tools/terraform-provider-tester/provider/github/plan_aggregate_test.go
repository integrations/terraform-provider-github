package github

import (
	"bufio"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/github/terraform-provider-tester/engine"
)

// fixtureTestNames reads the pinned acceptance-test catalog
// (engine/testdata/list_github.txt, generated at
// requirementsFixtureProviderSHA) in discovery order and fails if the pinned
// count ever drifts, so the aggregates asserted below always describe the
// same 173-name universe.
func fixtureTestNames(t *testing.T) []string {
	t.Helper()
	f, err := os.Open("../../engine/testdata/list_github.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	var names []string
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		if line := sc.Text(); strings.HasPrefix(line, "TestAcc") {
			names = append(names, line)
		}
	}
	if err := sc.Err(); err != nil {
		t.Fatal(err)
	}
	if len(names) != 173 {
		t.Fatalf("fixture has %d TestAcc names (SHA %s); want 173",
			len(names), requirementsFixtureProviderSHA)
	}
	return names
}

// TestIndividualPlanAggregatesRealCatalogRequirements pins what the COMPLETE
// individual suite actually demands, straight from the real catalog and the
// real planner. It is the ground truth behind every "individual mode
// prerequisites" claim in README.md, docs/, llms.txt and the Copilot skill:
// individual mode needs no dedicated organization, but it is not satisfied by
// a bare personal account and a generic PAT either.
//
// The authoritative design keeps the user SSH/GPG and Codespaces tests
// enabled in individual mode and expects preflight — not a narrowed catalog —
// to surface their missing permissions, so the aggregate classic scopes are
// repo, delete_repo, user, admin:public_key, admin:gpg_key and codespace, and
// the aggregate capabilities are codespaces-user-secrets plus
// template-repository (the repository resource and data source exercise a
// configured template repository under the same owner).
func TestIndividualPlanAggregatesRealCatalogRequirements(t *testing.T) {
	plan, err := engine.BuildExecutionPlan(fixtureTestNames(t), New(), engine.PlanOptions{Mode: "individual"})
	if err != nil {
		t.Fatalf("building the individual plan: %v", err)
	}

	// Sorted unique aggregates, per the planner contract (see
	// engine.BuildExecutionPlan's sortedUniqueKeys).
	wantScopes := []string{
		"admin:gpg_key",
		"admin:public_key",
		"codespace",
		"delete_repo",
		"repo",
		"user",
	}
	if !reflect.DeepEqual(plan.Scopes, wantScopes) {
		t.Errorf("individual plan scopes = %v; want %v", plan.Scopes, wantScopes)
	}

	wantCapabilities := []string{"codespaces-user-secrets", "template-repository"}
	if !reflect.DeepEqual(plan.Capabilities, wantCapabilities) {
		t.Errorf("individual plan capabilities = %v; want %v", plan.Capabilities, wantCapabilities)
	}

	// Individual mode still needs no dedicated organization: every
	// organization/team/enterprise test is excluded, so the "organization"
	// capability probe is never requested.
	for _, capability := range plan.Capabilities {
		if capability == "organization" {
			t.Errorf("individual plan capabilities = %v; must not require the organization capability", plan.Capabilities)
		}
	}

	// The user SSH/GPG and Codespaces tests really are part of this plan —
	// the scopes above are not an artifact of an unclassified fallback.
	eligible := make(map[string]bool, len(plan.Eligible))
	for _, name := range plan.Eligible {
		eligible[name] = true
	}
	for _, name := range []string{
		"TestAccGithubUserSshKey",
		"TestAccGithubUserGpgKey",
		"TestAccGithubCodespacesUserSecret",
		"TestAccGithubRepository",
	} {
		if !eligible[name] {
			t.Errorf("%s is not eligible in individual mode; the complete individual suite must keep it enabled", name)
		}
	}
	if len(plan.Unclassified) != 0 {
		t.Errorf("individual plan unclassified = %v; want none from the pinned catalog", plan.Unclassified)
	}
}
