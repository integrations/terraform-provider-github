package github

import (
	"context"
	"fmt"

	"github.com/github/terraform-provider-tester/provider"
	gogithub "github.com/google/go-github/v88/github"
)

// capabilityInput holds the identity context a capability probe needs to run
// its read-only check. Not every probe uses every field.
type capabilityInput struct {
	Owner        string // org/account login for the "organization" probe
	TemplateRepo string // repo name for the "template-repository" probe
}

// capabilityProbe performs one read-only GitHub API call and reports whether
// the aggregated test plan's environment capability is actually available.
// Implementations must not mutate any state.
type capabilityProbe func(ctx context.Context, client *gogithub.Client, input capabilityInput) provider.Check

// capabilityProbes is the registry of known environment capabilities. Keys are
// the exact capability names the planner (Task 2/3) emits in
// provider.TestRequirements.Capabilities.
var capabilityProbes = map[string]capabilityProbe{
	"codespaces-user-secrets": probeCodespacesUserSecrets,
	"organization":            probeOrganization,
	"template-repository":     probeTemplateRepository,
}

// probeCodespacesUserSecrets checks that the token/App identity can read its
// Codespaces user-secrets public key, a proxy for Codespaces access being
// usable by the tests that need it.
func probeCodespacesUserSecrets(ctx context.Context, client *gogithub.Client, _ capabilityInput) provider.Check {
	_, _, err := client.Codespaces.GetUserPublicKey(ctx)
	if err != nil {
		return provider.Check{
			Status: provider.CheckFail,
			Detail: fmt.Sprintf("failed to get Codespaces user public key: %v", err),
			Fix:    "grant the token/App the Codespaces user-secrets permission (or codespace scope for a classic PAT)",
		}
	}
	return provider.Check{Status: provider.CheckOK, Detail: "Codespaces user-secrets public key accessible"}
}

// probeOrganization checks that GET /orgs/{owner} succeeds, proving the owner
// is reachable as an organization (as opposed to a personal account, which
// 404s here).
func probeOrganization(ctx context.Context, client *gogithub.Client, input capabilityInput) provider.Check {
	_, _, err := client.Organizations.Get(ctx, input.Owner)
	if err != nil {
		return provider.Check{
			Status: provider.CheckFail,
			Detail: fmt.Sprintf("failed to get org %q: %v", input.Owner, err),
			Fix:    fmt.Sprintf("verify %q is an organization the token/App can access", input.Owner),
		}
	}
	return provider.Check{Status: provider.CheckOK, Detail: fmt.Sprintf("org %q accessible", input.Owner)}
}

// probeTemplateRepository checks that the configured template repository
// exists and is actually marked as a template.
func probeTemplateRepository(ctx context.Context, client *gogithub.Client, input capabilityInput) provider.Check {
	repo, _, err := client.Repositories.Get(ctx, input.Owner, input.TemplateRepo)
	if err != nil {
		return provider.Check{
			Status: provider.CheckFail,
			Detail: fmt.Sprintf("failed to get template repo %q: %v", input.TemplateRepo, err),
			Fix:    fmt.Sprintf("create %q or fix GH_TEST_ORG_TEMPLATE_REPOSITORY", input.TemplateRepo),
		}
	}
	if !repo.GetIsTemplate() {
		return provider.Check{
			Status: provider.CheckFail,
			Detail: fmt.Sprintf("repo %q exists but is not marked as a template", input.TemplateRepo),
			Fix:    "mark the repo as a template (Settings -> Template repository)",
		}
	}
	return provider.Check{Status: provider.CheckOK, Detail: fmt.Sprintf("repo %q is a template", input.TemplateRepo)}
}

// probeCapabilities runs the read-only probe for each requested capability
// name exactly once, in first-occurrence order, and returns one
// provider.Check per unique name. Names not present in the aggregated plan
// make zero HTTP requests. A name absent from capabilityProbes fails closed
// (CheckFail) rather than silently succeeding, since an unrecognized
// capability can never be verified as available.
func probeCapabilities(ctx context.Context, client *gogithub.Client, names []string, input capabilityInput) []provider.Check {
	var checks []provider.Check
	seen := make(map[string]bool, len(names))
	for _, name := range names {
		if name == "" || seen[name] {
			continue
		}
		seen[name] = true

		var c provider.Check
		if probe, ok := capabilityProbes[name]; ok {
			c = probe(ctx, client, input)
		} else {
			c = provider.Check{
				Status: provider.CheckFail,
				Detail: fmt.Sprintf("unknown capability %q: no probe registered", name),
				Fix:    fmt.Sprintf("remove %q from the test's required capabilities or add a probe for it", name),
			}
		}
		c.Name = "capability/" + name
		checks = append(checks, c)
	}
	return checks
}
