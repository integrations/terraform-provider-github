package fakeprovider

import (
	"context"

	"github.com/github/terraform-provider-tester/provider"
)

// Fake is a configurable provider.Provider double for engine + CLI tests.
type Fake struct {
	NameVal        string
	Packages       []string
	Pattern        string
	ModesVal       []provider.Mode
	EnvByMode      map[string][]provider.EnvVar
	GroupFunc      func(string) string
	RequirementsFn func(string) (provider.TestRequirements, bool)
	PreflightFn    func(context.Context, string, provider.TestRequirements) provider.PreflightReport
	DiscoverFn     func(context.Context, provider.DiscoverOpts) (provider.DiscoveryResult, error)
	OrphansFn      func(context.Context, string) ([]provider.Resource, error)
	OrphansVal     []provider.Resource
	OrphansErr     error
	SweepFn        func(context.Context, string, provider.SweepOpts) error
	Secrets        []string
}

var _ provider.Provider = (*Fake)(nil)

func (f *Fake) Name() string                      { return f.NameVal }
func (f *Fake) TestPackages() []string            { return f.Packages }
func (f *Fake) TestPattern() string               { return f.Pattern }
func (f *Fake) Modes() []provider.Mode            { return f.ModesVal }
func (f *Fake) EnvFor(m string) []provider.EnvVar { return f.EnvByMode[m] }

func (f *Fake) GroupOf(name string) string {
	if f.GroupFunc != nil {
		return f.GroupFunc(name)
	}
	return "misc"
}

func (f *Fake) RequirementsFor(name string) (provider.TestRequirements, bool) {
	if f.RequirementsFn != nil {
		return f.RequirementsFn(name)
	}
	return provider.TestRequirements{}, false
}

func (f *Fake) Preflight(ctx context.Context, m string, requirements provider.TestRequirements) provider.PreflightReport {
	if f.PreflightFn != nil {
		return f.PreflightFn(ctx, m, requirements)
	}
	return provider.PreflightReport{Mode: m}
}

func (f *Fake) Discover(ctx context.Context, opts provider.DiscoverOpts) (provider.DiscoveryResult, error) {
	if f.DiscoverFn != nil {
		return f.DiscoverFn(ctx, opts)
	}
	return provider.DiscoveryResult{}, nil
}

func (f *Fake) Orphans(ctx context.Context, mode string) ([]provider.Resource, error) {
	if f.OrphansFn != nil {
		return f.OrphansFn(ctx, mode)
	}
	return append([]provider.Resource(nil), f.OrphansVal...), f.OrphansErr
}

func (f *Fake) Sweep(ctx context.Context, mode string, o provider.SweepOpts) error {
	if f.SweepFn != nil {
		return f.SweepFn(ctx, mode, o)
	}
	return nil
}

func (f *Fake) SecretEnvKeys() []string { return f.Secrets }
