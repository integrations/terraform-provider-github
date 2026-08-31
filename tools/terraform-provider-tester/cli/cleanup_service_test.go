package cli

import (
	"context"
	"errors"
	"reflect"
	"strings"
	"testing"

	"github.com/github/terraform-provider-tester/provider"
)

type cleanupOrphansResult struct {
	resources []provider.Resource
	err       error
}

type cleanupProviderSpy struct {
	orphansResults []cleanupOrphansResult
	sweepErr       error
	orphansCalls   int
	sweepCalls     int
	sweepOpts      []provider.SweepOpts
	callOrder      []string
	orphansAlias   bool
}

func (s *cleanupProviderSpy) Name() string                    { return "test" }
func (s *cleanupProviderSpy) TestPackages() []string          { return nil }
func (s *cleanupProviderSpy) TestPattern() string             { return "" }
func (s *cleanupProviderSpy) Modes() []provider.Mode          { return nil }
func (s *cleanupProviderSpy) EnvFor(string) []provider.EnvVar { return nil }
func (s *cleanupProviderSpy) GroupOf(string) string           { return "" }
func (s *cleanupProviderSpy) RequirementsFor(string) (provider.TestRequirements, bool) {
	return provider.TestRequirements{}, false
}
func (s *cleanupProviderSpy) Preflight(context.Context, string, provider.TestRequirements) provider.PreflightReport {
	return provider.PreflightReport{}
}
func (s *cleanupProviderSpy) Discover(context.Context, provider.DiscoverOpts) (provider.DiscoveryResult, error) {
	return provider.DiscoveryResult{}, nil
}
func (s *cleanupProviderSpy) SecretEnvKeys() []string { return nil }

func (s *cleanupProviderSpy) Orphans(context.Context, string) ([]provider.Resource, error) {
	s.orphansCalls++
	s.callOrder = append(s.callOrder, "orphans")
	idx := s.orphansCalls - 1
	if idx >= len(s.orphansResults) {
		return nil, nil
	}
	result := s.orphansResults[idx]
	if s.orphansAlias {
		return result.resources, result.err
	}
	return cleanupCloneResources(result.resources), result.err
}

func (s *cleanupProviderSpy) Sweep(_ context.Context, _ string, opts provider.SweepOpts) error {
	s.sweepCalls++
	s.callOrder = append(s.callOrder, "sweep")
	s.sweepOpts = append(s.sweepOpts, provider.SweepOpts{
		Targets:        append([]string(nil), opts.Targets...),
		Confirm:        opts.Confirm,
		ExactResources: cleanupCloneResources(opts.ExactResources),
	})
	if len(opts.ExactResources) > 0 {
		opts.ExactResources[0].Name = "mutated-by-provider"
	}
	return s.sweepErr
}

func cleanupCloneResources(in []provider.Resource) []provider.Resource {
	return append([]provider.Resource(nil), in...)
}

func cleanupResourcesEqual(a, b []provider.Resource) bool {
	return reflect.DeepEqual(a, b)
}

func cleanupSweepOptsEqual(a, b provider.SweepOpts) bool {
	return reflect.DeepEqual(a, b)
}

func cleanupEnv(owner string) func(string) string {
	return func(key string) string {
		if key == "GITHUB_OWNER" {
			return owner
		}
		return ""
	}
}

func cleanupTestResources() []provider.Resource {
	return []provider.Resource{
		{Kind: "team", Name: "tf-acc-team", URL: "https://github.com/orgs/acme/teams/tf-acc-team"},
		{Kind: "repository", Name: "tf-acc-repo", URL: "https://github.com/acme/tf-acc-repo"},
	}
}

func TestCloneResourcesPreservesEmptyExactSnapshot(t *testing.T) {
	got := cloneResources([]provider.Resource{})
	if got == nil {
		t.Fatal("cloneResources(non-nil empty) = nil, want non-nil empty")
	}
	if len(got) != 0 {
		t.Fatalf("len(cloneResources(non-nil empty)) = %d, want 0", len(got))
	}
	if cloneResources(nil) != nil {
		t.Fatal("cloneResources(nil) is non-nil, want nil")
	}
}

func TestCleanupListIsReadOnly(t *testing.T) {
	t.Run("missing owner rejects without provider calls", func(t *testing.T) {
		prov := &cleanupProviderSpy{}
		svc := cleanupService{prov: prov, getenv: cleanupEnv("")}

		_, err := svc.List(context.Background())
		if err == nil || !strings.Contains(err.Error(), "GITHUB_OWNER") {
			t.Fatalf("List() error = %v, want missing GITHUB_OWNER", err)
		}
		if prov.orphansCalls != 0 {
			t.Fatalf("Orphans calls = %d, want 0", prov.orphansCalls)
		}
		if prov.sweepCalls != 0 {
			t.Fatalf("Sweep calls = %d, want 0", prov.sweepCalls)
		}
	})

	t.Run("success calls orphans once and never sweeps", func(t *testing.T) {
		resources := cleanupTestResources()
		prov := &cleanupProviderSpy{orphansResults: []cleanupOrphansResult{{resources: resources}}}
		svc := cleanupService{prov: prov, getenv: cleanupEnv("acme")}

		got, err := svc.List(context.Background())
		if err != nil {
			t.Fatalf("List() error = %v", err)
		}
		if got.Owner != "acme" {
			t.Fatalf("Owner = %q, want acme", got.Owner)
		}
		if !cleanupResourcesEqual(got.Resources, resources) {
			t.Fatalf("Resources = %+v, want %+v", got.Resources, resources)
		}
		if prov.orphansCalls != 1 {
			t.Fatalf("Orphans calls = %d, want 1", prov.orphansCalls)
		}
		if prov.sweepCalls != 0 {
			t.Fatalf("Sweep calls = %d, want 0", prov.sweepCalls)
		}
		if !reflect.DeepEqual(prov.callOrder, []string{"orphans"}) {
			t.Fatalf("call order = %v, want [orphans]", prov.callOrder)
		}
	})
}

func TestCleanupSweepRejectsWrongPhrase(t *testing.T) {
	t.Run("phrase must exactly match current owner", func(t *testing.T) {
		resources := cleanupTestResources()
		prov := &cleanupProviderSpy{}
		svc := cleanupService{prov: prov, getenv: cleanupEnv("acme")}

		_, err := svc.Sweep(context.Background(), sweepRequest{
			Owner:     "acme",
			Phrase:    "sweep acme",
			Resources: resources,
		})
		if err == nil || !strings.Contains(err.Error(), "confirmation phrase") {
			t.Fatalf("Sweep() error = %v, want confirmation-phrase error", err)
		}
		if prov.orphansCalls != 0 || prov.sweepCalls != 0 {
			t.Fatalf("calls = orphans:%d sweep:%d, want 0/0", prov.orphansCalls, prov.sweepCalls)
		}
	})

	t.Run("empty displayed snapshot is rejected before provider calls", func(t *testing.T) {
		prov := &cleanupProviderSpy{}
		svc := cleanupService{prov: prov, getenv: cleanupEnv("acme")}

		_, err := svc.Sweep(context.Background(), sweepRequest{Owner: "acme", Phrase: "SWEEP acme"})
		if err == nil || !strings.Contains(err.Error(), "no orphaned resources") {
			t.Fatalf("Sweep() error = %v, want empty-snapshot error", err)
		}
		if prov.orphansCalls != 0 || prov.sweepCalls != 0 {
			t.Fatalf("calls = orphans:%d sweep:%d, want 0/0", prov.orphansCalls, prov.sweepCalls)
		}
	})
}

func TestCleanupSweepRejectsChangedOwner(t *testing.T) {
	resources := cleanupTestResources()

	for _, tc := range []struct {
		name         string
		currentOwner string
		requestOwner string
	}{
		{name: "owner mismatch", currentOwner: "acme", requestOwner: "other"},
		{name: "empty snapshot owner", currentOwner: "acme", requestOwner: ""},
		{name: "missing current owner", currentOwner: "", requestOwner: "acme"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			prov := &cleanupProviderSpy{}
			svc := cleanupService{prov: prov, getenv: cleanupEnv(tc.currentOwner)}

			_, err := svc.Sweep(context.Background(), sweepRequest{
				Owner:     tc.requestOwner,
				Phrase:    "SWEEP acme",
				Resources: resources,
			})
			if err == nil || !strings.Contains(err.Error(), "owner changed") {
				t.Fatalf("Sweep() error = %v, want owner-changed error", err)
			}
			if prov.orphansCalls != 0 || prov.sweepCalls != 0 {
				t.Fatalf("calls = orphans:%d sweep:%d, want 0/0", prov.orphansCalls, prov.sweepCalls)
			}
		})
	}
}

func TestCleanupSweepAbortsWhenSnapshotChanges(t *testing.T) {
	t.Run("fresh list error blocks sweep", func(t *testing.T) {
		prov := &cleanupProviderSpy{orphansResults: []cleanupOrphansResult{{err: errors.New("fresh list failed")}}}
		svc := cleanupService{prov: prov, getenv: cleanupEnv("acme")}

		_, err := svc.Sweep(context.Background(), sweepRequest{
			Owner:     "acme",
			Phrase:    "SWEEP acme",
			Resources: cleanupTestResources(),
		})
		if err == nil || !strings.Contains(err.Error(), "fresh list failed") {
			t.Fatalf("Sweep() error = %v, want fresh-list error", err)
		}
		if prov.orphansCalls != 1 || prov.sweepCalls != 0 {
			t.Fatalf("calls = orphans:%d sweep:%d, want 1/0", prov.orphansCalls, prov.sweepCalls)
		}
	})

	t.Run("changed snapshot returns fresh resources without sweeping", func(t *testing.T) {
		request := []provider.Resource{
			{Kind: "repository", Name: "tf-acc-repo", URL: "https://github.com/acme/tf-acc-repo"},
			{Kind: "repository", Name: "tf-acc-repo", URL: "https://github.com/acme/tf-acc-repo"},
			{Kind: "team", Name: "tf-acc-team", URL: "https://github.com/orgs/acme/teams/tf-acc-team"},
		}
		fresh := []provider.Resource{
			{Kind: "team", Name: "tf-acc-team", URL: "https://github.com/orgs/acme/teams/tf-acc-team"},
			{Kind: "repository", Name: "tf-acc-repo", URL: "https://github.com/acme/tf-acc-repo"},
		}
		requestBefore := cleanupCloneResources(request)
		freshBefore := cleanupCloneResources(fresh)
		prov := &cleanupProviderSpy{orphansResults: []cleanupOrphansResult{{resources: fresh}}}
		svc := cleanupService{prov: prov, getenv: cleanupEnv("acme")}

		got, err := svc.Sweep(context.Background(), sweepRequest{
			Owner:     "acme",
			Phrase:    "SWEEP acme",
			Resources: request,
		})
		if err != nil {
			t.Fatalf("Sweep() error = %v", err)
		}
		if !got.SnapshotChanged {
			t.Fatal("SnapshotChanged = false, want true")
		}
		if !cleanupResourcesEqual(got.Remaining, freshBefore) {
			t.Fatalf("Remaining = %+v, want %+v", got.Remaining, freshBefore)
		}
		if prov.orphansCalls != 1 || prov.sweepCalls != 0 {
			t.Fatalf("calls = orphans:%d sweep:%d, want 1/0", prov.orphansCalls, prov.sweepCalls)
		}
		if !reflect.DeepEqual(prov.callOrder, []string{"orphans"}) {
			t.Fatalf("call order = %v, want [orphans]", prov.callOrder)
		}
		if !cleanupResourcesEqual(request, requestBefore) {
			t.Fatalf("request mutated to %+v, want %+v", request, requestBefore)
		}
		if !cleanupResourcesEqual(fresh, freshBefore) {
			t.Fatalf("provider fresh resources mutated to %+v, want %+v", fresh, freshBefore)
		}
	})
}

func TestCleanupSweepPassesFreshExactSnapshotAndReturnsResidual(t *testing.T) {
	request := []provider.Resource{
		{Kind: "team", Name: "tf-acc-test-team", URL: "https://github.com/orgs/acme/teams/tf-acc-test-team"},
		{Kind: "repository", Name: "tf-acc-test-repo", URL: "https://github.com/acme/tf-acc-test-repo"},
	}
	fresh := []provider.Resource{
		{Kind: "repository", Name: "tf-acc-test-repo", URL: "https://github.com/acme/tf-acc-test-repo"},
		{Kind: "team", Name: "tf-acc-test-team", URL: "https://github.com/orgs/acme/teams/tf-acc-test-team"},
	}
	residual := []provider.Resource{{Kind: "repository", Name: "tf-acc-test-repo-left", URL: "https://github.com/acme/tf-acc-test-repo-left"}}
	requestBefore := cleanupCloneResources(request)
	freshBefore := cleanupCloneResources(fresh)
	residualBefore := cleanupCloneResources(residual)
	prov := &cleanupProviderSpy{
		orphansResults: []cleanupOrphansResult{{resources: fresh}, {resources: residual}},
		orphansAlias:   true,
	}
	svc := cleanupService{prov: prov, getenv: cleanupEnv("acme")}

	got, err := svc.Sweep(context.Background(), sweepRequest{
		Owner:     "acme",
		Phrase:    "SWEEP acme",
		Resources: request,
	})
	if err != nil {
		t.Fatalf("Sweep() error = %v", err)
	}
	if got.SnapshotChanged {
		t.Fatal("SnapshotChanged = true, want false")
	}
	if !cleanupResourcesEqual(got.Remaining, residualBefore) {
		t.Fatalf("Remaining = %+v, want %+v", got.Remaining, residualBefore)
	}
	if prov.orphansCalls != 2 || prov.sweepCalls != 1 {
		t.Fatalf("calls = orphans:%d sweep:%d, want 2/1", prov.orphansCalls, prov.sweepCalls)
	}
	if !reflect.DeepEqual(prov.callOrder, []string{"orphans", "sweep", "orphans"}) {
		t.Fatalf("call order = %v, want [orphans sweep orphans]", prov.callOrder)
	}
	wantOpts := provider.SweepOpts{
		Targets:        []string{"repositories", "teams"},
		Confirm:        true,
		ExactResources: freshBefore,
	}
	if len(prov.sweepOpts) != 1 || !cleanupSweepOptsEqual(prov.sweepOpts[0], wantOpts) {
		t.Fatalf("Sweep opts = %+v, want %+v", prov.sweepOpts, wantOpts)
	}
	if !cleanupResourcesEqual(request, requestBefore) {
		t.Fatalf("request mutated to %+v, want %+v", request, requestBefore)
	}
	if !cleanupResourcesEqual(fresh, freshBefore) {
		t.Fatalf("fresh mutated to %+v, want %+v", fresh, freshBefore)
	}
	if !cleanupResourcesEqual(residual, residualBefore) {
		t.Fatalf("residual mutated to %+v, want %+v", residual, residualBefore)
	}
}

func TestCleanupSweepReturnsOriginalAndRefreshErrors(t *testing.T) {
	request := cleanupTestResources()
	fresh := []provider.Resource{
		{Kind: "repository", Name: "tf-acc-repo", URL: "https://github.com/acme/tf-acc-repo"},
		{Kind: "team", Name: "tf-acc-team", URL: "https://github.com/orgs/acme/teams/tf-acc-team"},
	}
	residual := []provider.Resource{{Kind: "team", Name: "tf-acc-team-left", URL: "https://github.com/orgs/acme/teams/tf-acc-team-left"}}
	sweepErr := errors.New("sweep failed")
	refreshErr := errors.New("refresh failed")
	prov := &cleanupProviderSpy{
		orphansResults: []cleanupOrphansResult{{resources: fresh}, {resources: residual, err: refreshErr}},
		sweepErr:       sweepErr,
	}
	svc := cleanupService{prov: prov, getenv: cleanupEnv("acme")}

	got, err := svc.Sweep(context.Background(), sweepRequest{
		Owner:     "acme",
		Phrase:    "SWEEP acme",
		Resources: request,
	})
	if !errors.Is(err, sweepErr) {
		t.Fatalf("error = %v, want errors.Is(_, sweepErr)", err)
	}
	if !errors.Is(err, refreshErr) {
		t.Fatalf("error = %v, want errors.Is(_, refreshErr)", err)
	}
	if !got.residualUnknown {
		t.Fatal("residualUnknown = false, want true after refresh error")
	}
	if len(got.Remaining) != 0 {
		t.Fatalf("Remaining = %+v, want empty when refresh failed", got.Remaining)
	}
	if got.SnapshotChanged {
		t.Fatal("SnapshotChanged = true, want false")
	}
	if prov.orphansCalls != 2 || prov.sweepCalls != 1 {
		t.Fatalf("calls = orphans:%d sweep:%d, want 2/1", prov.orphansCalls, prov.sweepCalls)
	}
	if !reflect.DeepEqual(prov.callOrder, []string{"orphans", "sweep", "orphans"}) {
		t.Fatalf("call order = %v, want [orphans sweep orphans]", prov.callOrder)
	}
}

func TestCleanupSweepMarksResidualUnknownWhenRefreshFails(t *testing.T) {
	request := cleanupTestResources()
	fresh := []provider.Resource{
		{Kind: "repository", Name: "tf-acc-repo", URL: "https://github.com/acme/tf-acc-repo"},
		{Kind: "team", Name: "tf-acc-team", URL: "https://github.com/orgs/acme/teams/tf-acc-team"},
	}
	refreshErr := errors.New("refresh failed")
	prov := &cleanupProviderSpy{
		orphansResults: []cleanupOrphansResult{{resources: fresh}, {err: refreshErr}},
	}
	svc := cleanupService{prov: prov, getenv: cleanupEnv("acme")}

	got, err := svc.Sweep(context.Background(), sweepRequest{Owner: "acme", Phrase: "SWEEP acme", Resources: request})
	if !errors.Is(err, refreshErr) {
		t.Fatalf("error = %v, want refresh error", err)
	}
	if !got.residualUnknown {
		t.Fatal("residualUnknown = false, want true")
	}
	if len(got.Remaining) != 0 {
		t.Fatalf("Remaining = %+v, want empty when residual state is unknown", got.Remaining)
	}
}
