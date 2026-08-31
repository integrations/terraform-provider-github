package github

import (
	"context"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/github/terraform-provider-tester/provider"
)

// TestProviderInterface verifies that New() satisfies the expected interface values.
func TestProviderInterface(t *testing.T) {
	p := New()

	if p.Name() != "github" {
		t.Errorf("Name() = %q, want %q", p.Name(), "github")
	}
	if p.TestPattern() != "^TestAcc" {
		t.Errorf("TestPattern() = %q, want %q", p.TestPattern(), "^TestAcc")
	}
	pkgs := p.TestPackages()
	if len(pkgs) != 1 || pkgs[0] != "./github/..." {
		t.Errorf("TestPackages() = %v, want [\"./github/...\"]", pkgs)
	}
}

// TestModesCount verifies that all five modes are returned.
func TestModesCount(t *testing.T) {
	p := New()
	got := p.Modes()
	if len(got) != 5 {
		t.Errorf("Modes() returned %d modes, want 5", len(got))
	}
	names := map[string]bool{}
	for _, m := range got {
		names[m.Name] = true
	}
	for _, want := range []string{"anonymous", "individual", "organization", "team", "enterprise"} {
		if !names[want] {
			t.Errorf("missing mode %q", want)
		}
	}
}

// TestEnvForAnonymous verifies that anonymous mode has no required env vars.
func TestEnvForAnonymous(t *testing.T) {
	p := New()
	vars := p.EnvFor("anonymous")
	if len(vars) != 0 {
		t.Errorf("EnvFor(anonymous) returned %d vars, want 0", len(vars))
	}
}

// TestEnvForIndividualExposesConditionalTemplateRepository verifies the TUI
// can edit the template repository fixture when the selected catalog tests need
// it, without treating it as mandatory for every individual-mode selection.
func TestEnvForIndividualExposesConditionalTemplateRepository(t *testing.T) {
	p := New()

	var found bool
	for _, v := range p.EnvFor("individual") {
		if v.Key != "GH_TEST_ORG_TEMPLATE_REPOSITORY" {
			continue
		}
		found = true
		if v.Required {
			t.Error("GH_TEST_ORG_TEMPLATE_REPOSITORY should be optional in individual EnvFor")
		}
	}
	if !found {
		t.Fatal("individual EnvFor is missing GH_TEST_ORG_TEMPLATE_REPOSITORY")
	}

	templateReqs, ok := p.RequirementsFor("TestAccGithubRepository")
	if !ok || !containsString(templateReqs.Capabilities, "template-repository") {
		t.Fatalf("repository catalog requirements = %+v, %v; want template-repository capability", templateReqs, ok)
	}
	nonTemplateReqs, ok := p.RequirementsFor("TestAccGithubUserSshKey")
	if !ok {
		t.Fatal("user SSH key test is missing from the requirements catalog")
	}
	if containsString(nonTemplateReqs.Capabilities, "template-repository") {
		t.Fatalf("user SSH key catalog requirements = %+v; template repository must remain selection-dependent", nonTemplateReqs)
	}
}

// TestEnvForOrganizationHasRequiredVars verifies that organization mode has expected required vars.
func TestEnvForOrganizationHasRequiredVars(t *testing.T) {
	p := New()
	vars := p.EnvFor("organization")
	required := map[string]bool{}
	for _, v := range vars {
		if v.Required {
			required[v.Key] = true
		}
	}
	for _, want := range []string{"GITHUB_OWNER", "GH_TEST_ORG_USER1", "GH_TEST_ORG_REPOSITORY", "GH_TEST_ORG_TEMPLATE_REPOSITORY", "GH_TEST_ORG_SECRET_NAME"} {
		if !required[want] {
			t.Errorf("expected %q to be Required in organization mode", want)
		}
	}
}

// TestEnvForTeamHasExternalUserVars verifies that team mode extends organization vars.
func TestEnvForTeamHasExternalUserVars(t *testing.T) {
	p := New()
	vars := p.EnvFor("team")
	keys := map[string]bool{}
	for _, v := range vars {
		keys[v.Key] = true
	}
	for _, want := range []string{"GH_TEST_EXTERNAL_USER1", "GH_TEST_EXTERNAL_USER1_TOKEN", "GH_TEST_EXTERNAL_USER2", "GH_TEST_ORG_USER2"} {
		if !keys[want] {
			t.Errorf("expected %q in team mode env vars", want)
		}
	}
}

// TestEnvForEnterpriseHasSlug verifies that enterprise mode requires GITHUB_ENTERPRISE_SLUG.
func TestEnvForEnterpriseHasSlug(t *testing.T) {
	p := New()
	vars := p.EnvFor("enterprise")
	found := false
	for _, v := range vars {
		if v.Key == "GITHUB_ENTERPRISE_SLUG" && v.Required {
			found = true
		}
	}
	if !found {
		t.Error("expected GITHUB_ENTERPRISE_SLUG to be required in enterprise mode")
	}
}

// TestSecretEnvVars verifies that credential-bearing env vars are marked Secret
// in env tables.
func TestSecretEnvVars(t *testing.T) {
	p := New()
	for _, mode := range []string{"individual", "organization", "team", "enterprise"} {
		vars := p.EnvFor(mode)
		secretByKey := map[string]bool{}
		for _, v := range vars {
			secretByKey[v.Key] = v.Secret
		}
		for _, key := range []string{"GITHUB_TOKEN", "GITHUB_APP_PEM_FILE"} {
			if !secretByKey[key] {
				t.Errorf("mode %q: %s should be Secret=true", mode, key)
			}
		}
	}
}

// TestSecretEnvKeysIncludesAppPEM verifies redactors include provider-compatible
// GitHub App PEM contents when app auth is used.
func TestSecretEnvKeysIncludesAppPEM(t *testing.T) {
	p := New()
	keys := p.SecretEnvKeys()
	want := map[string]bool{
		"GITHUB_TOKEN":                 true,
		"GH_TEST_EXTERNAL_USER1_TOKEN": true,
		"GITHUB_APP_PEM_FILE":          true,
	}
	for key := range want {
		if !containsString(keys, key) {
			t.Errorf("SecretEnvKeys() = %v, missing %s", keys, key)
		}
	}
}

func containsString(values []string, want string) bool {
	for _, v := range values {
		if v == want {
			return true
		}
	}
	return false
}

// TestGroupOfProvider verifies the GroupOf method on the provider.
func TestGroupOfProvider(t *testing.T) {
	p := New()
	if got := p.GroupOf("TestAccGithubRepository"); got != "repositories" {
		t.Errorf("GroupOf(TestAccGithubRepository) = %q, want %q", got, "repositories")
	}
}

func TestOrphansRequiresOwner(t *testing.T) {
	p := New()
	_, err := p.Orphans(context.Background(), "individual")
	if err == nil {
		t.Fatal("Orphans should require GITHUB_OWNER")
	}
}

func TestOrphansExplicitIndividualModeIgnoresAmbientOrganization(t *testing.T) {
	t.Setenv("GH_TEST_AUTH_MODE", "organization")

	mux := http.NewServeMux()
	mux.HandleFunc("GET /users/"+sweepTestOwner+"/repos", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, reposJSON("tf-acc-test-user-repo"))
	})
	mux.HandleFunc("/orgs/"+sweepTestOwner+"/repos", func(w http.ResponseWriter, r *http.Request) {
		t.Errorf("organization repos endpoint must not be called for explicit individual mode: %s %s", r.Method, r.URL.Path)
		http.Error(w, "unexpected organization repos call", http.StatusInternalServerError)
	})
	mux.HandleFunc("/orgs/"+sweepTestOwner+"/teams", func(w http.ResponseWriter, r *http.Request) {
		t.Errorf("teams endpoint must not be called for explicit individual mode: %s %s", r.Method, r.URL.Path)
		http.Error(w, "unexpected teams call", http.StatusInternalServerError)
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()

	got, err := orphansWithClient(context.Background(), newClient(srv.URL+"/", "t"), sweepTestOwner, isOrgMode("individual"))
	if err != nil {
		t.Fatalf("orphansWithClient error: %v", err)
	}
	want := []provider.Resource{{
		Kind: "repository",
		Name: "tf-acc-test-user-repo",
		URL:  "https://github.com/" + sweepTestOwner + "/tf-acc-test-user-repo",
	}}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("orphansWithClient(individual) = %+v, want %+v", got, want)
	}
}

func TestOrphansExplicitOrganizationModeIgnoresAmbientIndividual(t *testing.T) {
	t.Setenv("GH_TEST_AUTH_MODE", "individual")

	mux := http.NewServeMux()
	mux.HandleFunc("GET /orgs/"+sweepTestOwner+"/repos", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, reposJSON("tf-acc-test-org-repo"))
	})
	mux.HandleFunc("GET /orgs/"+sweepTestOwner+"/teams", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, teamsJSON("tf-acc-test-org-team"))
	})
	mux.HandleFunc("/users/"+sweepTestOwner+"/repos", func(w http.ResponseWriter, r *http.Request) {
		t.Errorf("user repos endpoint must not be called for explicit organization mode: %s %s", r.Method, r.URL.Path)
		http.Error(w, "unexpected user repos call", http.StatusInternalServerError)
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()

	got, err := orphansWithClient(context.Background(), newClient(srv.URL+"/", "t"), sweepTestOwner, isOrgMode("organization"))
	if err != nil {
		t.Fatalf("orphansWithClient error: %v", err)
	}
	want := []provider.Resource{
		{
			Kind: "repository",
			Name: "tf-acc-test-org-repo",
			URL:  "https://github.com/" + sweepTestOwner + "/tf-acc-test-org-repo",
		},
		{
			Kind: "team",
			Name: "tf-acc-test-org-team",
			URL:  "https://github.com/orgs/" + sweepTestOwner + "/teams/tf-acc-test-org-team",
		},
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("orphansWithClient(organization) = %+v, want %+v", got, want)
	}
}

func TestSweepRequiresConfirm(t *testing.T) {
	p := New()
	err := p.Sweep(context.Background(), "individual", provider.SweepOpts{})
	if err == nil {
		t.Error("Sweep should require opts.Confirm")
	}
}

func TestOrphansUsesConfiguredBaseURL(t *testing.T) {
	t.Setenv("GITHUB_OWNER", sweepTestOwner)
	t.Setenv("GITHUB_TOKEN", "ghp_fake")

	mux := http.NewServeMux()
	mux.HandleFunc("GET /orgs/"+sweepTestOwner+"/repos", func(w http.ResponseWriter, r *http.Request) {
		if got := r.Header.Get("Authorization"); got == "" {
			t.Error("Orphans request missing Authorization header")
		}
		writeJSON(w, reposJSON("tf-acc-test-orphan"))
	})
	mux.HandleFunc("GET /orgs/"+sweepTestOwner+"/teams", func(w http.ResponseWriter, r *http.Request) {
		if got := r.Header.Get("Authorization"); got == "" {
			t.Error("Orphans request missing Authorization header")
		}
		writeJSON(w, teamsJSON("tf-acc-test-team"))
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()

	t.Setenv("GITHUB_BASE_URL", srv.URL+"/")

	got, err := New().Orphans(context.Background(), "organization")
	if err != nil {
		t.Fatalf("Orphans returned error: %v", err)
	}

	want := []provider.Resource{
		{Kind: "repository", Name: "tf-acc-test-orphan", URL: "https://github.com/" + sweepTestOwner + "/tf-acc-test-orphan"},
		{Kind: "team", Name: "tf-acc-test-team", URL: "https://github.com/orgs/" + sweepTestOwner + "/teams/tf-acc-test-team"},
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Orphans() = %+v, want %+v", got, want)
	}
}

func TestSweepUsesConfiguredBaseURL(t *testing.T) {
	t.Setenv("GITHUB_OWNER", sweepTestOwner)
	t.Setenv("GITHUB_TOKEN", "ghp_fake")

	var deleted []string
	mux := http.NewServeMux()
	mux.HandleFunc("GET /orgs/"+sweepTestOwner+"/repos", func(w http.ResponseWriter, r *http.Request) {
		if got := r.Header.Get("Authorization"); got == "" {
			t.Error("Sweep request missing Authorization header")
		}
		writeJSON(w, reposJSON("tf-acc-test-sweep"))
	})
	mux.HandleFunc("DELETE /repos/"+sweepTestOwner+"/{name}", func(w http.ResponseWriter, r *http.Request) {
		deleted = append(deleted, r.PathValue("name"))
		w.WriteHeader(http.StatusNoContent)
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()

	t.Setenv("GITHUB_BASE_URL", srv.URL+"/")

	err := New().Sweep(context.Background(), "organization", provider.SweepOpts{
		Targets: []string{"repositories"},
		Confirm: true,
	})
	if err != nil {
		t.Fatalf("Sweep returned error: %v", err)
	}
	if !reflect.DeepEqual(deleted, []string{"tf-acc-test-sweep"}) {
		t.Fatalf("deleted repos = %v, want [tf-acc-test-sweep]", deleted)
	}
}

// TestRequirementsForFallback verifies that the GitHub provider returns the
// conservative fallback requirements (and false) for a test name that does not
// match any catalog rule. The fallback is intentionally over-broad so the
// planner can fail closed rather than silently skipping requirements.
func TestRequirementsForFallback(t *testing.T) {
	p := New()

	got, ok := p.RequirementsFor("TestAccUnknownDoesNotExist")
	if ok {
		t.Fatalf("RequirementsFor(unknown) ok = %v, want false", ok)
	}
	want := conservativeRequirements()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("RequirementsFor(unknown) = %+v, want %+v", got, want)
	}
}

// TestPreflightAcceptsRequirements verifies that the concrete Preflight method
// still returns the existing anonymous-mode report when requirements are
// provided. Task 1 ignores requirements for GitHub, so the argument should be
// accepted without changing behavior.
func TestPreflightAcceptsRequirements(t *testing.T) {
	p := New().(*ghProvider)
	requirements := provider.TestRequirements{
		Modes:  []string{"individual"},
		Scopes: []string{"repo"},
	}

	got := p.Preflight(context.Background(), "anonymous", requirements)
	want := provider.PreflightReport{
		Mode: "anonymous",
		Checks: []provider.Check{
			{
				Name:   "anonymous",
				Status: provider.CheckOK,
				Detail: "anonymous mode needs no credentials",
			},
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Preflight(anonymous, %+v) = %+v, want %+v", requirements, got, want)
	}
}
