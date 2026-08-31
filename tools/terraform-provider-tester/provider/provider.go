package provider

import "context"

type Provider interface {
	Name() string
	TestPackages() []string // ["./github/..."]
	TestPattern() string    // "^TestAcc"
	Modes() []Mode
	EnvFor(mode string) []EnvVar
	GroupOf(testName string) string
	RequirementsFor(testName string) (TestRequirements, bool)
	Preflight(ctx context.Context, mode string, requirements TestRequirements) PreflightReport
	Discover(ctx context.Context, opts DiscoverOpts) (DiscoveryResult, error)
	Orphans(ctx context.Context, mode string) ([]Resource, error)
	Sweep(ctx context.Context, mode string, opts SweepOpts) error
	SecretEnvKeys() []string
}
