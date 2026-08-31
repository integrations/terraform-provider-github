package github

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/github/terraform-provider-tester/internal/redact"
	"github.com/github/terraform-provider-tester/provider"
)

// orgMux returns an http.ServeMux wired for a happy-path org preflight run.
// owner and templateRepo are the env values for GITHUB_OWNER / GH_TEST_ORG_TEMPLATE_REPOSITORY.
func orgMux(owner, templateRepo string, scopeHeader string, rateLimitRemaining int, isTemplate bool) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/user", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-OAuth-Scopes", scopeHeader)
		json.NewEncoder(w).Encode(map[string]string{"login": "testuser"})
	})
	mux.HandleFunc("/orgs/", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"login": owner})
	})
	if templateRepo != "" {
		mux.HandleFunc(fmt.Sprintf("/repos/%s/%s", owner, templateRepo), func(w http.ResponseWriter, _ *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]bool{"is_template": isTemplate})
		})
	}
	mux.HandleFunc("/rate_limit", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"resources":{"core":{"remaining":%d,"limit":5000}}}`, rateLimitRemaining)
	})
	return mux
}

// TestPreflightAnonymous verifies that anonymous mode returns OK with a
// single success check and no credential requirements.
func TestPreflightAnonymous(t *testing.T) {
	report := preflightWithBase(context.Background(), "anonymous", provider.TestRequirements{}, mapEnv(nil), "")
	if !report.OK() {
		t.Errorf("anonymous preflight should pass, got checks: %v", report.Checks)
	}
	if len(report.Checks) == 0 {
		t.Error("expected at least one check for anonymous mode")
	}
	if report.Checks[0].Status != provider.CheckOK {
		t.Errorf("check status = %v, want CheckOK", report.Checks[0].Status)
	}
}

// TestPreflightMissingEnvFails verifies that a missing GITHUB_OWNER causes CheckFail.
func TestPreflightMissingEnvFails(t *testing.T) {
	env := mapEnv(map[string]string{
		"GITHUB_TOKEN": "ghp_fake",
		// GITHUB_OWNER intentionally absent
	})
	report := preflightWithBase(context.Background(), "organization", provider.TestRequirements{}, env, "")
	if report.OK() {
		t.Error("expected preflight to fail when GITHUB_OWNER is not set")
	}
	if !hasCheckFail(report, "owner") {
		t.Errorf("expected CheckFail for 'owner', got checks: %+v", report.Checks)
	}
}

// TestPreflightCustomBaseURLWarnsNotBlocks verifies that a custom base URL (GHES),
// which TestMain accepts when it parses, produces a non-blocking warning and skips
// network probes rather than hard-failing a valid run -- even when the aggregated
// plan requests scopes and capabilities that would otherwise be checked.
func TestPreflightCustomBaseURLWarnsNotBlocks(t *testing.T) {
	env := mapEnv(map[string]string{
		"GITHUB_BASE_URL": "https://ghes.example.com",
		"GITHUB_TOKEN":    "ghp_fake",
		"GITHUB_OWNER":    "myorg",
	})
	requirements := provider.TestRequirements{
		Scopes:       []string{"repo", "delete_repo", "read:org", "admin:org"},
		Capabilities: []string{"organization", "template-repository"},
	}
	// Bogus apiBase and no httptest server: if any network probe ran it would error,
	// so reaching OK proves the early return happened before any network call.
	report := preflightWithBase(context.Background(), "organization", requirements, env, "http://127.0.0.1:0/")
	if !report.OK() {
		t.Errorf("custom base URL is a valid GHES config; expected OK (warn, not fail), got %+v", report.Checks)
	}
	baseCheck := findCheck(report, "base-url")
	if baseCheck == nil || baseCheck.Status != provider.CheckWarn {
		t.Errorf("expected a CheckWarn named 'base-url', got %+v", report.Checks)
	}
	// Proof no network probe ran: identity/capability/repository-root/rate-limit
	// checks must be absent, even though requirements ask for capabilities.
	for _, c := range report.Checks {
		if c.Name == "identity" || c.Name == "rate-limit" || c.Name == "repository-root" || strings.HasPrefix(c.Name, "capability/") {
			t.Errorf("expected no network probes for custom base URL, got check %q", c.Name)
		}
	}
}

// TestPreflightAmbiguousTokenAndApp verifies that setting both GITHUB_TOKEN and GITHUB_APP_ID fails immediately.
func TestPreflightAmbiguousTokenAndApp(t *testing.T) {
	env := mapEnv(map[string]string{
		"GITHUB_OWNER":  "myorg",
		"GITHUB_TOKEN":  "ghp_fake",
		"GITHUB_APP_ID": "42",
	})
	// No httptest server - must not make any network call.
	report := preflightWithBase(context.Background(), "organization", provider.TestRequirements{}, env, "http://127.0.0.1:0/")
	if report.OK() {
		t.Error("expected preflight to fail when both token and app auth are set")
	}
	if !hasCheckFail(report, "auth") {
		t.Errorf("expected CheckFail for 'auth', got %+v", report.Checks)
	}
	if !checkDetailContains(report, "auth", "Both token and app auth configured") {
		t.Errorf("expected 'auth' check detail to contain 'Both token and app auth configured', got %+v", report.Checks)
	}
}

// TestPreflightAppIncomplete verifies that an incomplete app auth trio causes CheckFail.
func TestPreflightAppIncomplete(t *testing.T) {
	env := mapEnv(map[string]string{
		"GITHUB_OWNER":  "myorg",
		"GITHUB_APP_ID": "42",
		// Missing GITHUB_APP_INSTALLATION_ID and GITHUB_APP_PEM_FILE
	})
	report := preflightWithBase(context.Background(), "organization", provider.TestRequirements{}, env, "http://127.0.0.1:0/")
	if report.OK() {
		t.Error("expected preflight to fail for incomplete app auth")
	}
	if !hasCheckFail(report, "auth") {
		t.Errorf("expected CheckFail for 'auth', got %+v", report.Checks)
	}
	if !checkDetailContains(report, "auth", "App auth configured without all required parameters") {
		t.Errorf("expected 'auth' check to mention missing params, got %+v", report.Checks)
	}
}

// TestPreflightHappyOrg verifies a fully healthy org mode preflight, driven by an
// aggregated plan that requires the organization and template-repository
// capabilities plus the classic org scope set.
func TestPreflightHappyOrg(t *testing.T) {
	const (
		owner        = "myorg"
		templateRepo = "my-template"
	)
	mux := orgMux(owner, templateRepo, "repo, delete_repo, read:org, admin:org", 5000, true)
	srv := httptest.NewServer(mux)
	defer srv.Close()

	env := mapEnv(map[string]string{
		"GITHUB_OWNER":                    owner,
		"GITHUB_TOKEN":                    "ghp_fake",
		"GH_TEST_ORG_TEMPLATE_REPOSITORY": templateRepo,
	})
	requirements := provider.TestRequirements{
		Scopes:       []string{"repo", "delete_repo", "read:org", "admin:org"},
		Capabilities: []string{"organization", "template-repository"},
	}
	report := preflightWithBase(context.Background(), "organization", requirements, env, srv.URL+"/")
	if !report.OK() {
		t.Errorf("expected all-OK preflight, got failing checks: %+v", failingChecks(report))
	}
	requireCheck(t, report, "identity", provider.CheckOK)
	requireCheck(t, report, "capability/organization", provider.CheckOK)
	requireCheck(t, report, "capability/template-repository", provider.CheckOK)
	requireCheck(t, report, "scopes", provider.CheckOK)
	requireCheck(t, report, "rate-limit", provider.CheckOK)
}

// TestPreflightTemplateNotTemplate verifies that a repo not marked as template causes
// CheckFail with Fix, when the aggregated plan requires the template-repository capability.
func TestPreflightTemplateNotTemplate(t *testing.T) {
	const (
		owner        = "myorg"
		templateRepo = "not-a-template"
	)
	mux := orgMux(owner, templateRepo, "repo, read:org, admin:org", 5000, false)
	srv := httptest.NewServer(mux)
	defer srv.Close()

	env := mapEnv(map[string]string{
		"GITHUB_OWNER":                    owner,
		"GITHUB_TOKEN":                    "ghp_fake",
		"GH_TEST_ORG_TEMPLATE_REPOSITORY": templateRepo,
	})
	requirements := provider.TestRequirements{Capabilities: []string{"template-repository"}}
	report := preflightWithBase(context.Background(), "organization", requirements, env, srv.URL+"/")
	if report.OK() {
		t.Error("expected preflight to fail when template repo is not marked as template")
	}
	c := findCheck(report, "capability/template-repository")
	if c == nil {
		t.Fatal("missing 'capability/template-repository' check")
	}
	if c.Status != provider.CheckFail {
		t.Errorf("capability/template-repository status = %v, want CheckFail", c.Status)
	}
	if !strings.Contains(c.Fix, "mark the repo as a template") {
		t.Errorf("Fix = %q, expected it to mention marking as template", c.Fix)
	}
}

// TestPreflightScopeGapDoctor verifies that a missing admin:org scope causes CheckFail with the right Fix.
func TestPreflightScopeGapDoctor(t *testing.T) {
	const owner = "myorg"
	// Scope "admin:org" is deliberately missing.
	mux := orgMux(owner, "", "repo, read:org", 5000, true)
	srv := httptest.NewServer(mux)
	defer srv.Close()

	env := mapEnv(map[string]string{
		"GITHUB_OWNER": owner,
		"GITHUB_TOKEN": "ghp_fake",
	})
	requirements := provider.TestRequirements{Scopes: []string{"repo", "delete_repo", "read:org", "admin:org"}}
	report := preflightWithBase(context.Background(), "organization", requirements, env, srv.URL+"/")
	if report.OK() {
		t.Error("expected preflight to fail due to missing admin:org scope")
	}
	c := findCheck(report, "scopes")
	if c == nil {
		t.Fatal("missing 'scopes' check")
	}
	if c.Status != provider.CheckFail {
		t.Errorf("scopes status = %v, want CheckFail", c.Status)
	}
	if !strings.Contains(c.Fix, "admin:org") {
		t.Errorf("Fix = %q, expected it to name 'admin:org'", c.Fix)
	}
}

// TestPreflightAdminOrgImpliesReadOrg verifies that a classic PAT granting the parent
// admin:org scope satisfies the read:org requirement. GitHub collapses hierarchical
// scopes: selecting admin:org auto-includes write:org and read:org, but the
// X-OAuth-Scopes header reports only the parent admin:org. A literal string match would
// wrongly flag read:org as missing and block every enterprise/org run.
func TestPreflightAdminOrgImpliesReadOrg(t *testing.T) {
	const owner = "myorg"
	// Real-world enterprise PAT header: admin:org present, read:org NOT listed.
	mux := orgMux(owner, "", "repo, delete_repo, admin:org, admin:enterprise, user, workflow", 5000, true)
	srv := httptest.NewServer(mux)
	defer srv.Close()

	env := mapEnv(map[string]string{
		"GITHUB_OWNER":           owner,
		"GITHUB_TOKEN":           "ghp_fake",
		"GITHUB_ENTERPRISE_SLUG": "my-enterprise",
	})
	requirements := provider.TestRequirements{
		Scopes: []string{"repo", "delete_repo", "read:org", "admin:org", "admin:enterprise"},
	}
	report := preflightWithBase(context.Background(), "enterprise", requirements, env, srv.URL+"/")
	c := findCheck(report, "scopes")
	if c == nil {
		t.Fatal("missing 'scopes' check")
	}
	if c.Status != provider.CheckOK {
		t.Errorf("scopes status = %v (detail %q), want CheckOK: admin:org must satisfy read:org", c.Status, c.Detail)
	}
}

// TestPreflightIndividualRequiresDeleteRepo verifies that an individual plan naming the
// delete_repo scope is enforced. Acceptance tests create AND destroy tf-acc-test
// repositories, and the teardown DELETE call returns 403 without delete_repo, leaking
// repos. A token with repo and user but no delete_repo must fail preflight with a Fix
// that names delete_repo.
func TestPreflightIndividualRequiresDeleteRepo(t *testing.T) {
	const login = "testuser"
	mux := http.NewServeMux()
	mux.HandleFunc("/user", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-OAuth-Scopes", "repo, user") // delete_repo deliberately missing
		json.NewEncoder(w).Encode(map[string]string{"login": login})
	})
	mux.HandleFunc("/rate_limit", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"resources":{"core":{"remaining":5000,"limit":5000}}}`)
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()

	env := mapEnv(map[string]string{
		"GITHUB_OWNER":    login,
		"GITHUB_USERNAME": login,
		"GITHUB_TOKEN":    "ghp_fake",
	})
	requirements := provider.TestRequirements{Scopes: []string{"repo", "delete_repo", "user"}}
	report := preflightWithBase(context.Background(), "individual", requirements, env, srv.URL+"/")
	if report.OK() {
		t.Error("expected preflight to fail: individual mode token lacks delete_repo")
	}
	c := findCheck(report, "scopes")
	if c == nil {
		t.Fatal("missing 'scopes' check")
	}
	if c.Status != provider.CheckFail {
		t.Errorf("scopes status = %v, want CheckFail", c.Status)
	}
	if !strings.Contains(c.Fix, "delete_repo") {
		t.Errorf("Fix = %q, expected it to name 'delete_repo'", c.Fix)
	}
}

// TestPreflightRateLimitLow verifies that a low remaining rate limit produces CheckWarn.
func TestPreflightRateLimitLow(t *testing.T) {
	const owner = "myorg"
	mux := orgMux(owner, "", "repo, read:org, admin:org", 5, true)
	srv := httptest.NewServer(mux)
	defer srv.Close()

	env := mapEnv(map[string]string{
		"GITHUB_OWNER": owner,
		"GITHUB_TOKEN": "ghp_fake",
	})
	report := preflightWithBase(context.Background(), "organization", provider.TestRequirements{}, env, srv.URL+"/")
	c := findCheck(report, "rate-limit")
	if c == nil {
		t.Fatal("missing 'rate-limit' check")
	}
	if c.Status != provider.CheckWarn {
		t.Errorf("rate-limit status = %v, want CheckWarn", c.Status)
	}
	if !strings.Contains(c.Detail, "low rate-limit headroom") {
		t.Errorf("Detail = %q, expected 'low rate-limit headroom'", c.Detail)
	}
}

// TestPreflightAppAuthHappy verifies a healthy app-auth preflight path, and that the
// scopes check names the exact requested scopes and capabilities for manual
// verification since a GitHub App has no classic scope header to inspect.
func TestPreflightAppAuthHappy(t *testing.T) {
	const owner = "myorg"
	mux := http.NewServeMux()
	mux.HandleFunc("POST /app/installations/99/access_tokens", func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") == "" {
			t.Error("installation token request missing Authorization header")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, `{"token":"app_installation_token"}`)
	})
	mux.HandleFunc("/orgs/", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"login": owner})
	})
	mux.HandleFunc("/rate_limit", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"resources":{"core":{"remaining":5000,"limit":5000}}}`)
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()

	env := mapEnv(map[string]string{
		"GITHUB_OWNER":               owner,
		"GITHUB_APP_ID":              "42",
		"GITHUB_APP_INSTALLATION_ID": "99",
		"GITHUB_APP_PEM_FILE":        testAppPEM(t),
	})
	requirements := provider.TestRequirements{
		Scopes:       []string{"repo", "delete_repo", "read:org", "admin:org"},
		Capabilities: []string{"organization"},
	}
	report := preflightWithBase(context.Background(), "organization", requirements, env, srv.URL+"/")

	requireCheck(t, report, "identity", provider.CheckOK)
	requireCheck(t, report, "capability/organization", provider.CheckOK)

	// Scopes must be UNVERIFIED (fine-grained/App) and must name the exact
	// requested scopes and capabilities so a human can verify them manually.
	c := findCheck(report, "scopes")
	if c == nil {
		t.Fatal("missing 'scopes' check")
	}
	if c.Status != provider.CheckWarn {
		t.Errorf("scopes status = %v, want CheckWarn", c.Status)
	}
	if !strings.Contains(c.Detail, "UNVERIFIED") {
		t.Errorf("Detail = %q, expected 'UNVERIFIED'", c.Detail)
	}
	for _, want := range []string{"repo", "delete_repo", "read:org", "admin:org", "organization"} {
		if !strings.Contains(c.Detail, want) {
			t.Errorf("Detail = %q, expected it to name requested %q for manual verification", c.Detail, want)
		}
	}
	if c.Fix == "" {
		t.Error("expected 'scopes' check to carry a non-empty Fix even when UNVERIFIED")
	}
}

// TestPreflightCompleteIndividualPlanRequiresAllScopes verifies that a complete
// individual plan -- one whose aggregated requirements name every individual-mode
// scope, including the SSH/GPG key and Codespaces extras -- is satisfied when the
// classic PAT actually grants all of them.
func TestPreflightCompleteIndividualPlanRequiresAllScopes(t *testing.T) {
	const login = "testuser"
	completeScopes := []string{"repo", "delete_repo", "user", "admin:public_key", "admin:gpg_key", "codespace"}

	mux := http.NewServeMux()
	mux.HandleFunc("/user", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-OAuth-Scopes", strings.Join(completeScopes, ", "))
		json.NewEncoder(w).Encode(map[string]string{"login": login})
	})
	mux.HandleFunc("/rate_limit", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"resources":{"core":{"remaining":5000,"limit":5000}}}`)
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()

	env := mapEnv(map[string]string{
		"GITHUB_OWNER":    login,
		"GITHUB_USERNAME": login,
		"GITHUB_TOKEN":    "ghp_fake",
	})
	requirements := provider.TestRequirements{Scopes: completeScopes}
	report := preflightWithBase(context.Background(), "individual", requirements, env, srv.URL+"/")

	requireCheck(t, report, "scopes", provider.CheckOK)
}

// TestPreflightIndividualPlanMissingScopesExactRemediation verifies the exact
// classic-PAT remediation text when a complete individual plan requires the
// SSH/GPG key and Codespaces scopes but the token only grants the base
// repo/delete_repo/user trio.
func TestPreflightIndividualPlanMissingScopesExactRemediation(t *testing.T) {
	const login = "testuser"
	mux := http.NewServeMux()
	mux.HandleFunc("/user", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-OAuth-Scopes", "repo, delete_repo, user")
		json.NewEncoder(w).Encode(map[string]string{"login": login})
	})
	mux.HandleFunc("/rate_limit", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"resources":{"core":{"remaining":5000,"limit":5000}}}`)
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()

	env := mapEnv(map[string]string{
		"GITHUB_OWNER":    login,
		"GITHUB_USERNAME": login,
		"GITHUB_TOKEN":    "ghp_fake",
	})
	requirements := provider.TestRequirements{
		Scopes: []string{"repo", "delete_repo", "user", "admin:public_key", "admin:gpg_key", "codespace"},
	}
	report := preflightWithBase(context.Background(), "individual", requirements, env, srv.URL+"/")

	c := findCheck(report, "scopes")
	if c == nil {
		t.Fatal("missing 'scopes' check")
	}
	if c.Status != provider.CheckFail {
		t.Errorf("scopes status = %v, want CheckFail", c.Status)
	}
	wantFix := "Add these scopes at https://github.com/settings/tokens:\nadmin:public_key, admin:gpg_key, codespace"
	if c.Fix != wantFix {
		t.Errorf("Fix = %q, want %q", c.Fix, wantFix)
	}
}

// TestPreflightPlannerSortedScopesExactRemediation is a Task 6 regression test
// (Task 4 follow-up #2): the aggregated plan (engine.ExecutionPlan.Scopes)
// arrives alphabetically sorted, not in catalog order. This proves that
// feeding preflight an alphabetically-sorted requirements.Scopes still
// produces the exact required missing-scope remediation text in catalog
// order, not alphabetical order - the same wording asserted by
// TestPreflightIndividualPlanMissingScopesExactRemediation above, whose input
// happened to already be in catalog order and so could not by itself catch a
// presentation-order regression.
func TestPreflightPlannerSortedScopesExactRemediation(t *testing.T) {
	const login = "testuser"
	mux := http.NewServeMux()
	mux.HandleFunc("/user", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-OAuth-Scopes", "repo, delete_repo, user")
		json.NewEncoder(w).Encode(map[string]string{"login": login})
	})
	mux.HandleFunc("/rate_limit", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"resources":{"core":{"remaining":5000,"limit":5000}}}`)
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()

	env := mapEnv(map[string]string{
		"GITHUB_OWNER":    login,
		"GITHUB_USERNAME": login,
		"GITHUB_TOKEN":    "ghp_fake",
	})
	// Alphabetically sorted, as engine.ExecutionPlan.Scopes (sortedUniqueKeys)
	// would deliver it - deliberately not in catalog order.
	requirements := provider.TestRequirements{
		Scopes: []string{"admin:gpg_key", "admin:public_key", "codespace", "delete_repo", "repo", "user"},
	}
	report := preflightWithBase(context.Background(), "individual", requirements, env, srv.URL+"/")

	c := findCheck(report, "scopes")
	if c == nil {
		t.Fatal("missing 'scopes' check")
	}
	if c.Status != provider.CheckFail {
		t.Errorf("scopes status = %v, want CheckFail", c.Status)
	}
	wantFix := "Add these scopes at https://github.com/settings/tokens:\nadmin:public_key, admin:gpg_key, codespace"
	if c.Fix != wantFix {
		t.Errorf("Fix = %q, want %q (same required wording regardless of requirements.Scopes input order)", c.Fix, wantFix)
	}
}

// TestPreflightIndividualPlanWithoutExtrasDoesNotRequireKeyOrCodespaceScopes verifies
// that an individual plan without any SSH/GPG key or Codespaces tests does not require
// admin:public_key, admin:gpg_key, or codespace -- only the base repo/delete_repo/user
// trio the aggregated plan actually names.
func TestPreflightIndividualPlanWithoutExtrasDoesNotRequireKeyOrCodespaceScopes(t *testing.T) {
	const login = "testuser"
	mux := http.NewServeMux()
	mux.HandleFunc("/user", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-OAuth-Scopes", "repo, delete_repo, user")
		json.NewEncoder(w).Encode(map[string]string{"login": login})
	})
	mux.HandleFunc("/rate_limit", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"resources":{"core":{"remaining":5000,"limit":5000}}}`)
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()

	env := mapEnv(map[string]string{
		"GITHUB_OWNER":    login,
		"GITHUB_USERNAME": login,
		"GITHUB_TOKEN":    "ghp_fake",
	})
	requirements := provider.TestRequirements{Scopes: []string{"repo", "delete_repo", "user"}}
	report := preflightWithBase(context.Background(), "individual", requirements, env, srv.URL+"/")

	requireCheck(t, report, "scopes", provider.CheckOK)
}

// TestPreflightPlanScopesDeduplicate verifies that duplicate entries in
// requirements.Scopes collapse to a single mention in the Fix text, proving
// preflight treats requirements.Scopes as a set, not a literal list.
func TestPreflightPlanScopesDeduplicate(t *testing.T) {
	const login = "testuser"
	mux := http.NewServeMux()
	mux.HandleFunc("/user", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-OAuth-Scopes", "repo")
		json.NewEncoder(w).Encode(map[string]string{"login": login})
	})
	mux.HandleFunc("/rate_limit", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"resources":{"core":{"remaining":5000,"limit":5000}}}`)
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()

	env := mapEnv(map[string]string{
		"GITHUB_OWNER":    login,
		"GITHUB_USERNAME": login,
		"GITHUB_TOKEN":    "ghp_fake",
	})
	// delete_repo is duplicated three times, as an aggregator merging several
	// per-test requirement lists might produce before preflight deduplicates.
	requirements := provider.TestRequirements{Scopes: []string{"repo", "delete_repo", "delete_repo", "delete_repo"}}
	report := preflightWithBase(context.Background(), "individual", requirements, env, srv.URL+"/")

	c := findCheck(report, "scopes")
	if c == nil {
		t.Fatal("missing 'scopes' check")
	}
	if c.Status != provider.CheckFail {
		t.Errorf("scopes status = %v, want CheckFail", c.Status)
	}
	wantFix := "Add these scopes at https://github.com/settings/tokens:\ndelete_repo"
	if c.Fix != wantFix {
		t.Errorf("Fix = %q, want %q (duplicate scope entries must collapse to one)", c.Fix, wantFix)
	}
}

// TestSecretEnvKeys verifies that SecretEnvKeys returns exactly the expected
// secret-bearing values.
func TestSecretEnvKeys(t *testing.T) {
	p := New()
	keys := p.SecretEnvKeys()
	want := []string{"GITHUB_TOKEN", "GH_TEST_EXTERNAL_USER1_TOKEN", "GITHUB_APP_PEM_FILE"}
	if len(keys) != len(want) {
		t.Fatalf("SecretEnvKeys = %v, want %v", keys, want)
	}
	for i, k := range want {
		if keys[i] != k {
			t.Errorf("SecretEnvKeys[%d] = %q, want %q", i, keys[i], k)
		}
	}
}

// TestSecretHygiene verifies that no Check.Detail or Check.Fix contains the raw token value.
func TestSecretHygiene(t *testing.T) {
	const fakeToken = "ghp_fake"

	mux := orgMux("myorg", "", "repo, read:org", 5000, true)
	srv := httptest.NewServer(mux)
	defer srv.Close()

	env := mapEnv(map[string]string{
		"GITHUB_OWNER": "myorg",
		"GITHUB_TOKEN": fakeToken,
	})
	// Deliberately require a scope the token lacks (admin:org) so the "scopes"
	// check's Fix text (which quotes scope names) is exercised too.
	requirements := provider.TestRequirements{Scopes: []string{"repo", "read:org", "admin:org"}}
	report := preflightWithBase(context.Background(), "organization", requirements, env, srv.URL+"/")

	for _, c := range report.Checks {
		if strings.Contains(c.Detail, fakeToken) {
			t.Errorf("check %q Detail contains the token value", c.Name)
		}
		if strings.Contains(c.Fix, fakeToken) {
			t.Errorf("check %q Fix contains the token value", c.Name)
		}
	}
}

// TestPreflightRepositoryRootPreventsTemplateProbe verifies that when the
// aggregated plan requires the template-repository capability but no
// GH_TEST_ORG_TEMPLATE_REPOSITORY is configured, preflight fails fast with a
// "repository-root" check and makes zero HTTP requests for the
// template-repository probe, while still running any other requested
// capability (organization, here) normally.
func TestPreflightRepositoryRootPreventsTemplateProbe(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/user", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-OAuth-Scopes", "repo, read:org, admin:org")
		json.NewEncoder(w).Encode(map[string]string{"login": "myorg"})
	})
	mux.HandleFunc("/orgs/", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"login": "myorg"})
	})
	mux.HandleFunc("/rate_limit", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"resources":{"core":{"remaining":5000,"limit":5000}}}`)
	})
	mux.HandleFunc("/repos/", func(w http.ResponseWriter, r *http.Request) {
		t.Errorf("unexpected request to %s: repository-root precondition should have prevented the template-repository probe", r.URL.Path)
		http.Error(w, `{"message":"should not be called"}`, http.StatusInternalServerError)
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()

	env := mapEnv(map[string]string{"GITHUB_OWNER": "myorg", "GITHUB_TOKEN": "ghp_fake"})
	// GH_TEST_ORG_TEMPLATE_REPOSITORY deliberately unset.
	requirements := provider.TestRequirements{Capabilities: []string{"organization", "template-repository"}}
	report := preflightWithBase(context.Background(), "organization", requirements, env, srv.URL+"/")

	rr := findCheck(report, "repository-root")
	if rr == nil {
		t.Fatal("missing 'repository-root' check")
	}
	if rr.Status != provider.CheckFail {
		t.Errorf("repository-root status = %v, want CheckFail", rr.Status)
	}
	if rr.Detail == "" {
		t.Error("expected non-empty Detail for 'repository-root'")
	}
	if rr.Fix == "" {
		t.Error("expected non-empty Fix for 'repository-root'")
	}
	if findCheck(report, "capability/template-repository") != nil {
		t.Error("capability/template-repository must not run when repository-root failed")
	}
	requireCheck(t, report, "capability/organization", provider.CheckOK)
}

// TestPreflightFailuresIncludeProblemDetailAndFix is a table test asserting that
// every preflight failure path under Task 4's scope -- owner, username, auth,
// base URL, repository root, scopes, and capability -- carries a non-empty
// problem Detail and a non-empty, actionable Fix. identity, enterprise, and
// rate-limit are deliberately out of scope: they are not among the categories
// this task names, so no Fix text is added to them.
func TestPreflightFailuresIncludeProblemDetailAndFix(t *testing.T) {
	cases := []struct {
		name      string
		checkName string
		build     func(t *testing.T) (mode string, requirements provider.TestRequirements, env func(string) string, apiBase string)
	}{
		{
			name:      "owner",
			checkName: "owner",
			build: func(t *testing.T) (string, provider.TestRequirements, func(string) string, string) {
				env := mapEnv(map[string]string{"GITHUB_TOKEN": "ghp_fake"})
				return "organization", provider.TestRequirements{}, env, ""
			},
		},
		{
			name:      "username",
			checkName: "username",
			build: func(t *testing.T) (string, provider.TestRequirements, func(string) string, string) {
				env := mapEnv(map[string]string{"GITHUB_OWNER": "myuser", "GITHUB_TOKEN": "ghp_fake"})
				return "individual", provider.TestRequirements{}, env, ""
			},
		},
		{
			name:      "auth",
			checkName: "auth",
			build: func(t *testing.T) (string, provider.TestRequirements, func(string) string, string) {
				// Neither GITHUB_TOKEN nor any GITHUB_APP_* variable is set.
				env := mapEnv(map[string]string{"GITHUB_OWNER": "myorg"})
				return "organization", provider.TestRequirements{}, env, ""
			},
		},
		{
			name:      "base-url",
			checkName: "base-url",
			build: func(t *testing.T) (string, provider.TestRequirements, func(string) string, string) {
				env := mapEnv(map[string]string{
					"GITHUB_OWNER":    "myorg",
					"GITHUB_TOKEN":    "ghp_fake",
					"GITHUB_BASE_URL": "https://ghes.example.com",
				})
				return "organization", provider.TestRequirements{}, env, ""
			},
		},
		{
			name:      "repository-root",
			checkName: "repository-root",
			build: func(t *testing.T) (string, provider.TestRequirements, func(string) string, string) {
				mux := orgMux("myorg", "", "repo, read:org, admin:org", 5000, true)
				srv := httptest.NewServer(mux)
				t.Cleanup(srv.Close)
				env := mapEnv(map[string]string{"GITHUB_OWNER": "myorg", "GITHUB_TOKEN": "ghp_fake"})
				requirements := provider.TestRequirements{Capabilities: []string{"template-repository"}}
				return "organization", requirements, env, srv.URL + "/"
			},
		},
		{
			name:      "scopes",
			checkName: "scopes",
			build: func(t *testing.T) (string, provider.TestRequirements, func(string) string, string) {
				mux := orgMux("myorg", "", "repo", 5000, true)
				srv := httptest.NewServer(mux)
				t.Cleanup(srv.Close)
				env := mapEnv(map[string]string{"GITHUB_OWNER": "myorg", "GITHUB_TOKEN": "ghp_fake"})
				requirements := provider.TestRequirements{Scopes: []string{"repo", "admin:org"}}
				return "organization", requirements, env, srv.URL + "/"
			},
		},
		{
			name:      "capability",
			checkName: "capability/organization",
			build: func(t *testing.T) (string, provider.TestRequirements, func(string) string, string) {
				mux := http.NewServeMux()
				mux.HandleFunc("/user", func(w http.ResponseWriter, _ *http.Request) {
					w.Header().Set("Content-Type", "application/json")
					w.Header().Set("X-OAuth-Scopes", "repo")
					json.NewEncoder(w).Encode(map[string]string{"login": "myorg"})
				})
				mux.HandleFunc("/orgs/", func(w http.ResponseWriter, _ *http.Request) {
					http.Error(w, `{"message":"Not Found"}`, http.StatusNotFound)
				})
				mux.HandleFunc("/rate_limit", func(w http.ResponseWriter, _ *http.Request) {
					w.Header().Set("Content-Type", "application/json")
					fmt.Fprint(w, `{"resources":{"core":{"remaining":5000,"limit":5000}}}`)
				})
				srv := httptest.NewServer(mux)
				t.Cleanup(srv.Close)
				env := mapEnv(map[string]string{"GITHUB_OWNER": "myorg", "GITHUB_TOKEN": "ghp_fake"})
				requirements := provider.TestRequirements{Capabilities: []string{"organization"}}
				return "organization", requirements, env, srv.URL + "/"
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			mode, requirements, env, apiBase := tc.build(t)
			report := preflightWithBase(context.Background(), mode, requirements, env, apiBase)
			c := findCheck(report, tc.checkName)
			if c == nil {
				t.Fatalf("missing %q check, got checks: %+v", tc.checkName, report.Checks)
			}
			if c.Status == provider.CheckOK {
				t.Fatalf("%q check is CheckOK, want a non-OK status to exercise Detail/Fix", tc.checkName)
			}
			if c.Detail == "" {
				t.Errorf("%q check has empty Detail", tc.checkName)
			}
			if c.Fix == "" {
				t.Errorf("%q check has empty Fix", tc.checkName)
			}
		})
	}
}

// TestPreflightDetailAndFixSurviveRedaction is a characterization test proving
// that the existing internal/redact.Redactor -- the same redaction mechanism
// used at the CLI's output emission boundaries (cli.redactorForProvider,
// engine/report.go, etc.) -- would catch and mask a credential-shaped string
// if one ever appeared in a preflight Check's Detail or Fix, in both its plain
// text and JSON-marshaled forms. Task 4 does not add a parallel redactor to
// provider/github; preflight.go's own construction is responsible for never
// emitting a secret (see TestSecretHygiene), and this test guards the
// fallback net that exists elsewhere in the codebase.
func TestPreflightDetailAndFixSurviveRedaction(t *testing.T) {
	const fakeToken = "ghp_ABCDEFGHIJ0123456789abcdefghijklmnop" // gitleaks:allow -- synthetic fixture, not a real credential
	c := provider.Check{
		Name:   "scopes",
		Status: provider.CheckFail,
		Detail: fmt.Sprintf("missing scopes: repo (token %s)", fakeToken),
		Fix:    fmt.Sprintf("Add these scopes at https://github.com/settings/tokens (token %s):\nrepo", fakeToken),
	}

	red := redact.New(nil)
	redactedDetail := red.String(c.Detail)
	redactedFix := red.String(c.Fix)

	if strings.Contains(redactedDetail, fakeToken) {
		t.Errorf("redacted Detail still contains the fake token: %q", redactedDetail)
	}
	if strings.Contains(redactedFix, fakeToken) {
		t.Errorf("redacted Fix still contains the fake token: %q", redactedFix)
	}

	// Also prove this holds through JSON marshaling, since a Check may be
	// serialized (for example CLI --json preflight output) before any
	// redaction is applied at the emission boundary.
	blob, err := json.Marshal(c)
	if err != nil {
		t.Fatalf("json.Marshal: %v", err)
	}
	if strings.Contains(red.String(string(blob)), fakeToken) {
		t.Error("redacted JSON-marshaled Check still contains the fake token")
	}
}

// ---- helpers ----

func hasCheckFail(r provider.PreflightReport, name string) bool {
	c := findCheck(r, name)
	return c != nil && c.Status == provider.CheckFail
}

func checkDetailContains(r provider.PreflightReport, name, sub string) bool {
	c := findCheck(r, name)
	return c != nil && strings.Contains(c.Detail, sub)
}

func findCheck(r provider.PreflightReport, name string) *provider.Check {
	for i := range r.Checks {
		if r.Checks[i].Name == name {
			return &r.Checks[i]
		}
	}
	return nil
}

func requireCheck(t *testing.T, r provider.PreflightReport, name string, want provider.CheckStatus) {
	t.Helper()
	c := findCheck(r, name)
	if c == nil {
		t.Errorf("missing check %q", name)
		return
	}
	if c.Status != want {
		t.Errorf("check %q status = %v, want %v (detail: %s)", name, c.Status, want, c.Detail)
	}
}

func failingChecks(r provider.PreflightReport) []provider.Check {
	var out []provider.Check
	for _, c := range r.Checks {
		if c.Status == provider.CheckFail {
			out = append(out, c)
		}
	}
	return out
}

// TestPreflightIndividualSkipsOrgProbe verifies that individual mode does not probe
// the organization capability unless the aggregated plan actually requires it. Task
// 2's real catalog never grants the "organization" capability to individual-only
// tests, so a plan with no Capabilities must not synthesize an organization probe
// (which would 404 against a personal account and wrongly block the run).
func TestPreflightIndividualSkipsOrgProbe(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/user", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-OAuth-Scopes", "repo, delete_repo, user")
		json.NewEncoder(w).Encode(map[string]string{"login": "indyuser"})
	})
	mux.HandleFunc("/orgs/", func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, `{"message":"Not Found"}`, http.StatusNotFound)
	})
	mux.HandleFunc("/rate_limit", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"resources":{"core":{"remaining":5000,"limit":5000}}}`)
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()

	env := mapEnv(map[string]string{
		"GITHUB_OWNER":    "indyuser",
		"GITHUB_USERNAME": "indyuser",
		"GITHUB_TOKEN":    "ghp_fake",
	})
	requirements := provider.TestRequirements{Scopes: []string{"repo", "delete_repo", "user"}}
	report := preflightWithBase(context.Background(), "individual", requirements, env, srv.URL)

	if findCheck(report, "capability/organization") != nil {
		t.Errorf("individual plan without the organization capability must not run the org probe, got checks: %+v", report.Checks)
	}
	if !report.OK() {
		t.Errorf("individual preflight should pass when identity is valid, got checks: %+v", report.Checks)
	}
}
