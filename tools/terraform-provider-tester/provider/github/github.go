// Package github implements the provider.Provider interface for the GitHub Terraform provider.
// It provides environment tables, token validation, a read-only preflight, and test grouping.
package github

import (
	"context"
	"fmt"
	"os"

	"github.com/github/terraform-provider-tester/provider"
)

type ghProvider struct{}

var _ provider.Provider = (*ghProvider)(nil)

// New returns a new GitHub provider.Provider implementation.
func New() provider.Provider {
	return &ghProvider{}
}

func (p *ghProvider) Name() string                         { return "github" }
func (p *ghProvider) TestPackages() []string               { return []string{"./github/..."} }
func (p *ghProvider) TestPattern() string                  { return "^TestAcc" }
func (p *ghProvider) Modes() []provider.Mode               { return modes() }
func (p *ghProvider) EnvFor(mode string) []provider.EnvVar { return envFor(mode) }
func (p *ghProvider) GroupOf(testName string) string       { return groupOf(testName) }

func (p *ghProvider) RequirementsFor(testName string) (provider.TestRequirements, bool) {
	return requirementsFor(testName)
}

func (p *ghProvider) Preflight(ctx context.Context, mode string, requirements provider.TestRequirements) provider.PreflightReport {
	return preflightWithBase(ctx, mode, requirements, os.Getenv, "")
}

func (p *ghProvider) Discover(ctx context.Context, opts provider.DiscoverOpts) (provider.DiscoveryResult, error) {
	token := os.Getenv("GITHUB_TOKEN")
	apiBase := os.Getenv("GITHUB_BASE_URL")

	if token == "" {
		return provider.DiscoveryResult{Notes: []string{"no token configured; set GITHUB_TOKEN"}}, nil
	}

	cred, err := validateToken(ctx, apiBase, token)
	if err != nil {
		return provider.DiscoveryResult{}, fmt.Errorf("validating token: %w", err)
	}

	c := newClient(apiBase, token)
	return discoverWithClient(ctx, c, cred, apiBase, token, opts)
}

func (p *ghProvider) SecretEnvKeys() []string {
	return []string{"GITHUB_TOKEN", "GH_TEST_EXTERNAL_USER1_TOKEN", "GITHUB_APP_PEM_FILE"}
}
