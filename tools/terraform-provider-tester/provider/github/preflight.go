package github

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/github/terraform-provider-tester/provider"
	gogithub "github.com/google/go-github/v88/github"
)

// preflightWithBase is the testable core of Preflight.
// apiBase is empty for GitHub.com (uses the go-github default);
// tests pass an httptest server URL to avoid live network calls.
// The env parameter mirrors the acc_test.go TestMain auth tree.
// requirements is the aggregated provider.TestRequirements for the tests that
// will actually run in mode (see Task 2/3): preflight checks exactly the
// scopes and capabilities the plan names instead of a hardcoded per-mode list.
func preflightWithBase(ctx context.Context, mode string, requirements provider.TestRequirements, env func(string) string, apiBase string) provider.PreflightReport {
	var checks []provider.Check
	add := func(c provider.Check) { checks = append(checks, c) }

	// Step 1: anonymous mode needs no credentials.
	if mode == "anonymous" {
		return provider.PreflightReport{
			Mode: mode,
			Checks: []provider.Check{
				{Name: "anonymous", Status: provider.CheckOK, Detail: "anonymous mode needs no credentials"},
			},
		}
	}

	// Step 3: GITHUB_OWNER must be set for all non-anonymous modes.
	owner := env("GITHUB_OWNER")
	if owner == "" {
		add(provider.Check{
			Name:   "owner",
			Status: provider.CheckFail,
			Detail: "GITHUB_OWNER environment variable not set",
			Fix:    "Set GITHUB_OWNER to the organization login or user login the tests should target",
		})
		return provider.PreflightReport{Mode: mode, Checks: checks}
	}

	// Step 4: individual mode requires GITHUB_USERNAME.
	if mode == "individual" {
		if env("GITHUB_USERNAME") == "" {
			add(provider.Check{
				Name:   "username",
				Status: provider.CheckFail,
				Detail: "GITHUB_USERNAME environment variable not set",
				Fix:    "Set GITHUB_USERNAME to the account login that owns GITHUB_TOKEN (required in individual mode)",
			})
			return provider.PreflightReport{Mode: mode, Checks: checks}
		}
	}

	// Step 5: detect auth method (no network call).
	method, err := detectAuth(env)
	if err != nil {
		add(provider.Check{Name: "auth", Status: provider.CheckFail, Detail: err.Error(), Fix: authFix(err)})
		return provider.PreflightReport{Mode: mode, Checks: checks}
	}

	// Custom base URL (GHES): TestMain accepts it when it parses, so do not block the run.
	// The probes below target api.github.com, so skip them and warn that GHES is unverified.
	if env("GITHUB_BASE_URL") != "" {
		add(provider.Check{
			Name:   "base-url",
			Status: provider.CheckWarn,
			Detail: "custom base URL (GHES) detected; preflight network probes skipped and targets unverified",
			Fix:    "Manually verify connectivity, permissions, and the requested scopes/capabilities against the custom GITHUB_BASE_URL; preflight cannot probe a custom GHES endpoint",
		})
		return provider.PreflightReport{Mode: mode, Checks: checks}
	}

	// Step 6: validate the chosen credential.
	var cred credInfo
	if method == authToken {
		token := env("GITHUB_TOKEN")
		cred, err = validateToken(ctx, apiBase, token)
		if err != nil {
			add(provider.Check{Name: "identity", Status: provider.CheckFail, Detail: err.Error()})
			return provider.PreflightReport{Mode: mode, Checks: checks}
		}
		identityStatus := provider.CheckOK
		identityDetail := fmt.Sprintf("authenticated as %s", cred.Login)
		if mode == "individual" && cred.Login != env("GITHUB_USERNAME") {
			identityStatus = provider.CheckWarn
			identityDetail = fmt.Sprintf("authenticated as %s but GITHUB_USERNAME is %s", cred.Login, env("GITHUB_USERNAME"))
		}
		add(provider.Check{Name: "identity", Status: identityStatus, Detail: identityDetail})
	} else {
		cred, err = validateApp(ctx, apiBase, env("GITHUB_APP_ID"), env("GITHUB_APP_INSTALLATION_ID"), env("GITHUB_APP_PEM_FILE"))
		if err != nil {
			add(provider.Check{Name: "identity", Status: provider.CheckFail, Detail: err.Error()})
			return provider.PreflightReport{Mode: mode, Checks: checks}
		}
		add(provider.Check{Name: "identity", Status: provider.CheckOK, Detail: "authenticated as GitHub App"})
	}

	// Build a client for the read-only probe checks.
	clientToken := env("GITHUB_TOKEN")
	if method == authApp {
		clientToken = cred.Token
	}
	client := newClient(apiBase, clientToken)

	// Step 7: capability probes. The aggregated plan (Task 2/3) names exactly
	// which read-only environment capabilities the selected tests need; run
	// only those, deduplicated, instead of a hardcoded per-mode list.
	//
	// The template-repository probe needs a concrete repo name to check. If
	// the plan requires that capability but no template repository is
	// configured, fail fast with a "repository-root" check instead of letting
	// the probe attempt (and fail) an HTTP call against an empty repo name.
	capabilityNames := requirements.Capabilities
	if sliceContains(capabilityNames, "template-repository") && env("GH_TEST_ORG_TEMPLATE_REPOSITORY") == "" {
		add(provider.Check{
			Name:   "repository-root",
			Status: provider.CheckFail,
			Detail: "template-repository capability requested but GH_TEST_ORG_TEMPLATE_REPOSITORY is not set",
			Fix:    "Set GH_TEST_ORG_TEMPLATE_REPOSITORY to an existing template repository under the target owner",
		})
		capabilityNames = sliceWithout(capabilityNames, "template-repository")
	}
	capChecks := probeCapabilities(ctx, client, capabilityNames, capabilityInput{
		Owner:        owner,
		TemplateRepo: env("GH_TEST_ORG_TEMPLATE_REPOSITORY"),
	})
	for _, c := range capChecks {
		add(c)
	}

	// Step 8: scope check, driven by the aggregated plan's exact requirements.Scopes
	// instead of a hardcoded per-mode list. A plan that never names
	// admin:public_key, admin:gpg_key, or codespace (no key or Codespaces tests
	// selected) does not require them.
	if !cred.FineGrained {
		needed := uniquePreserveOrder(requirements.Scopes)
		granted := expandScopes(cred.Scopes)
		var missing []string
		for _, s := range needed {
			if !granted[s] {
				missing = append(missing, s)
			}
		}
		sortByScopePresentationOrder(missing)
		if len(missing) > 0 {
			add(provider.Check{
				Name:   "scopes",
				Status: provider.CheckFail,
				Detail: fmt.Sprintf("missing scopes: %s", strings.Join(missing, ", ")),
				Fix:    fmt.Sprintf("Add these scopes at https://github.com/settings/tokens:\n%s", strings.Join(missing, ", ")),
			})
		} else {
			add(provider.Check{Name: "scopes", Status: provider.CheckOK, Detail: "scopes sufficient"})
		}
	} else {
		// Fine-grained PATs and GitHub Apps have no classic X-OAuth-Scopes header,
		// so the exact requested scopes and capabilities cannot be verified here;
		// name them so a human can check manually before running the tests.
		add(provider.Check{
			Name:   "scopes",
			Status: provider.CheckWarn,
			Detail: fmt.Sprintf(
				"scopes UNVERIFIED (fine-grained PAT or App): manual verification required for scopes (%s) and capabilities (%s)",
				strings.Join(requirements.Scopes, ", "), strings.Join(requirements.Capabilities, ", "),
			),
			Fix: "Manually verify the token/App grants the exact requested scopes and capabilities listed above",
		})
	}

	// Step 9: enterprise slug check.
	if mode == "enterprise" {
		slug := env("GITHUB_ENTERPRISE_SLUG")
		if slug == "" {
			add(provider.Check{Name: "enterprise", Status: provider.CheckFail, Detail: "GITHUB_ENTERPRISE_SLUG environment variable not set"})
		} else {
			add(provider.Check{Name: "enterprise", Status: provider.CheckWarn, Detail: "enterprise reachability UNVERIFIED"})
		}
	}

	// Step 10: rate-limit probe.
	rl, _, rlErr := client.RateLimit.Get(ctx)
	if rlErr != nil {
		add(provider.Check{Name: "rate-limit", Status: provider.CheckWarn, Detail: fmt.Sprintf("failed to get rate limit: %v", rlErr)})
	} else {
		remaining, limit := 0, 0
		if rl.Core != nil {
			remaining = rl.Core.Remaining
			limit = rl.Core.Limit
		}
		if remaining < 100 {
			add(provider.Check{Name: "rate-limit", Status: provider.CheckWarn, Detail: fmt.Sprintf("low rate-limit headroom: %d/%d remaining", remaining, limit)})
		} else {
			add(provider.Check{Name: "rate-limit", Status: provider.CheckOK, Detail: fmt.Sprintf("rate limit: %d/%d remaining", remaining, limit)})
		}
	}

	return provider.PreflightReport{Mode: mode, Checks: checks}
}

// newClient creates a go-github client, optionally with a base URL override.
// token is empty for anonymous/app-auth clients.
func newClient(apiBase, token string) *gogithub.Client {
	var opts []gogithub.ClientOptionsFunc
	if token != "" {
		opts = append(opts, gogithub.WithAuthToken(token))
	}
	if apiBase != "" {
		opts = append(opts, gogithub.WithURLs(&apiBase, nil))
	}
	c, _ := gogithub.NewClient(opts...)
	return c
}

// authFix returns actionable remediation text for a detectAuth error, keyed on
// its exact message so the "auth" check always carries a non-empty,
// method-specific Fix. The default case is a safe fallback for any message
// this function does not recognize.
func authFix(err error) string {
	switch err.Error() {
	case "authentication not configured":
		return "Set GITHUB_TOKEN, or set GITHUB_APP_ID, GITHUB_APP_INSTALLATION_ID, and GITHUB_APP_PEM_FILE together for GitHub App auth"
	case "Both token and app auth configured":
		return "Unset GITHUB_TOKEN or the GITHUB_APP_* trio; preflight cannot pick between token and App auth"
	case "App auth configured without all required parameters":
		return "Set all three of GITHUB_APP_ID, GITHUB_APP_INSTALLATION_ID, and GITHUB_APP_PEM_FILE, or remove them and use GITHUB_TOKEN instead"
	default:
		return "Fix the GITHUB_TOKEN or GITHUB_APP_* environment variables so exactly one auth method is fully configured"
	}
}

// sliceContains reports whether target is present in s.
func sliceContains(s []string, target string) bool {
	for _, v := range s {
		if v == target {
			return true
		}
	}
	return false
}

// sliceWithout returns s with every occurrence of target removed, preserving
// the order of the remaining elements.
func sliceWithout(s []string, target string) []string {
	out := make([]string, 0, len(s))
	for _, v := range s {
		if v != target {
			out = append(out, v)
		}
	}
	return out
}

// scopePresentationOrder fixes the display order for missing-scope
// remediation text, independent of the order requirements.Scopes arrives in.
// The aggregated plan (engine.ExecutionPlan.Scopes) is alphabetically sorted,
// which would otherwise reorder this catalog-ordered, human-reviewed
// remediation text every time the scope set driving it changes.
var scopePresentationOrder = []string{
	"repo", "delete_repo", "user",
	"admin:public_key", "admin:gpg_key", "codespace",
	"read:org", "admin:org", "admin:enterprise",
}

// scopePresentationRank maps each entry in scopePresentationOrder to its
// index; scopes absent from the catalog sort after all known scopes,
// preserving their relative order (stable sort).
var scopePresentationRank = func() map[string]int {
	rank := make(map[string]int, len(scopePresentationOrder))
	for i, s := range scopePresentationOrder {
		rank[s] = i
	}
	return rank
}()

// sortByScopePresentationOrder sorts scopes in place by
// scopePresentationOrder so remediation text reads in a fixed, catalog order
// regardless of the order the caller's requirements.Scopes arrived in (for
// example, alphabetically sorted from an aggregated engine.ExecutionPlan).
// Unknown scopes keep their relative order and sort after every known scope.
func sortByScopePresentationOrder(scopes []string) {
	rank := func(s string) int {
		if r, ok := scopePresentationRank[s]; ok {
			return r
		}
		return len(scopePresentationOrder)
	}
	sort.SliceStable(scopes, func(i, j int) bool {
		return rank(scopes[i]) < rank(scopes[j])
	})
}

// uniquePreserveOrder returns s with empty and duplicate entries removed,
// preserving the first-occurrence order of the remaining elements. Used to
// deduplicate an aggregated plan's requirements.Scopes before checking them,
// without imposing an alphabetical sort on the resulting remediation text.
func uniquePreserveOrder(s []string) []string {
	seen := make(map[string]bool, len(s))
	out := make([]string, 0, len(s))
	for _, v := range s {
		if v == "" || seen[v] {
			continue
		}
		seen[v] = true
		out = append(out, v)
	}
	return out
}

// scopeImplies maps a granted classic-PAT scope to the scopes it directly includes.
// GitHub's OAuth scopes are hierarchical: selecting a parent scope grants its children,
// but the X-OAuth-Scopes header lists only the parent (for example a PAT with admin:org
// reports "admin:org", never "write:org, read:org"). Preflight must expand the granted
// set through this hierarchy before checking requirements, or a token that is actually
// sufficient gets wrongly reported as missing scopes and blocks the run.
// Ref: https://docs.github.com/apps/oauth-apps/building-oauth-apps/scopes-for-oauth-apps
var scopeImplies = map[string][]string{
	"repo":                  {"repo:status", "repo_deployment", "public_repo", "repo:invite", "security_events"},
	"admin:org":             {"write:org", "read:org", "manage_runners:org"},
	"write:org":             {"read:org"},
	"admin:public_key":      {"write:public_key", "read:public_key"},
	"admin:repo_hook":       {"write:repo_hook", "read:repo_hook"},
	"user":                  {"read:user", "user:email", "user:follow"},
	"write:packages":        {"read:packages"},
	"write:discussion":      {"read:discussion"},
	"admin:gpg_key":         {"write:gpg_key", "read:gpg_key"},
	"admin:ssh_signing_key": {"write:ssh_signing_key", "read:ssh_signing_key"},
	"admin:enterprise":      {"manage_runners:enterprise", "manage_billing:enterprise", "read:enterprise"},
}

// expandScopes returns the set of granted scopes plus every scope they transitively
// imply per the GitHub scope hierarchy (see scopeImplies).
func expandScopes(granted []string) map[string]bool {
	set := make(map[string]bool, len(granted))
	var add func(string)
	add = func(s string) {
		if set[s] {
			return
		}
		set[s] = true
		for _, child := range scopeImplies[s] {
			add(child)
		}
	}
	for _, g := range granted {
		add(g)
	}
	return set
}
