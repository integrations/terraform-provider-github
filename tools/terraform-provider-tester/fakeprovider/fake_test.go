package fakeprovider

import (
	"context"
	"reflect"
	"testing"

	"github.com/github/terraform-provider-tester/provider"
)

func TestFakeRequirementsFor(t *testing.T) {
	f := &Fake{RequirementsFn: func(name string) (provider.TestRequirements, bool) {
		return provider.TestRequirements{Modes: []string{"individual"}}, name == "TestAccA"
	}}
	got, ok := f.RequirementsFor("TestAccA")
	if !ok || !reflect.DeepEqual(got.Modes, []string{"individual"}) {
		t.Fatalf("RequirementsFor = %+v, %v", got, ok)
	}
}

func TestFakePreflightReceivesRequirements(t *testing.T) {
	want := provider.TestRequirements{Scopes: []string{"repo"}}
	f := &Fake{PreflightFn: func(_ context.Context, mode string, got provider.TestRequirements) provider.PreflightReport {
		if mode != "individual" || !reflect.DeepEqual(got, want) {
			t.Fatalf("Preflight(%q, %+v)", mode, got)
		}
		return provider.PreflightReport{Mode: mode}
	}}
	f.Preflight(context.Background(), "individual", want)
}

// TestFakeRequirementsForEmptyFallback verifies that an empty Fake returns no
// requirements and reports false when no callback is installed.
func TestFakeRequirementsForEmptyFallback(t *testing.T) {
	f := &Fake{}

	got, ok := f.RequirementsFor("TestAccX")
	if ok {
		t.Fatalf("RequirementsFor(TestAccX) ok = %v, want false", ok)
	}
	if !reflect.DeepEqual(got, provider.TestRequirements{}) {
		t.Fatalf("RequirementsFor(TestAccX) = %+v, want empty requirements", got)
	}
}

// TestFakePreflightDefaultMode verifies that the default Fake Preflight result
// preserves the mode passed by the caller.
func TestFakePreflightDefaultMode(t *testing.T) {
	f := &Fake{}

	got := f.Preflight(context.Background(), "individual", provider.TestRequirements{})
	if got.Mode != "individual" {
		t.Fatalf("Preflight mode = %q, want %q", got.Mode, "individual")
	}
}

func TestOrphansPassesExplicitMode(t *testing.T) {
	want := []provider.Resource{{Kind: "repository", Name: "tf-acc-test-a"}}
	f := &Fake{OrphansFn: func(_ context.Context, mode string) ([]provider.Resource, error) {
		if mode != "individual" {
			t.Fatalf("Orphans mode = %q, want %q", mode, "individual")
		}
		return want, nil
	}}

	got, err := f.Orphans(context.Background(), "individual")
	if err != nil {
		t.Fatalf("Orphans returned error: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Orphans resources = %+v, want %+v", got, want)
	}
}

func TestSweepPassesExplicitModeAndOpts(t *testing.T) {
	want := provider.SweepOpts{
		Targets:   []string{"teams"},
		Confirm:   true,
		Resources: []provider.Resource{{Kind: "team", Name: "tf-acc-test-team"}},
	}
	f := &Fake{SweepFn: func(_ context.Context, mode string, got provider.SweepOpts) error {
		if mode != "organization" {
			t.Fatalf("Sweep mode = %q, want %q", mode, "organization")
		}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("Sweep opts = %+v, want %+v", got, want)
		}
		return nil
	}}

	if err := f.Sweep(context.Background(), "organization", want); err != nil {
		t.Fatalf("Sweep returned error: %v", err)
	}
}
