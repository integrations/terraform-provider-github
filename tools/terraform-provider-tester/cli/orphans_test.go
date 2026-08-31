package cli

import (
	"bytes"
	"context"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/github/terraform-provider-tester/engine"
	"github.com/github/terraform-provider-tester/fakeprovider"
	"github.com/github/terraform-provider-tester/provider"
)

func writeCleanupState(t *testing.T, root string, state engine.State) {
	t.Helper()
	if state.Provider == "" {
		state.Provider = "test"
	}
	if err := state.Save(filepath.Join(root, ".pulsar-state.json")); err != nil {
		t.Fatalf("write cleanup state: %v", err)
	}
}

func TestOrphansRefusesConcurrentRunLock(t *testing.T) {
	root := t.TempDir()
	release, err := engine.AcquireLock(filepath.Join(root, ".pulsar-state.json.lock"))
	if err != nil {
		t.Fatalf("AcquireLock: %v", err)
	}
	defer release() //nolint:errcheck

	orphansCalls := 0
	d := deps{
		provider: &fakeprovider.Fake{
			OrphansFn: func(_ context.Context, _ string) ([]provider.Resource, error) {
				orphansCalls++
				return nil, nil
			},
		},
		cwd: func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"orphans", "--repo-root", root, "--mode", "individual"}, &out, &errOut, d)

	if code != 1 {
		t.Fatalf("exit code = %d, want 1; stderr=%s", code, errOut.String())
	}
	if orphansCalls != 0 {
		t.Fatalf("Orphans calls = %d, want 0", orphansCalls)
	}
	if !strings.Contains(errOut.String(), "acquiring lock:") {
		t.Fatalf("stderr = %q, want lock acquisition failure", errOut.String())
	}
}

func TestSweepRefusesConcurrentRunLock(t *testing.T) {
	root := t.TempDir()
	release, err := engine.AcquireLock(filepath.Join(root, ".pulsar-state.json.lock"))
	if err != nil {
		t.Fatalf("AcquireLock: %v", err)
	}
	defer release() //nolint:errcheck

	sweepCalls := 0
	d := deps{
		provider: &fakeprovider.Fake{
			SweepFn: func(_ context.Context, _ string, _ provider.SweepOpts) error {
				sweepCalls++
				return nil
			},
		},
		cwd: func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"sweep", "--repo-root", root, "--mode", "individual", "--confirm"}, &out, &errOut, d)

	if code != 1 {
		t.Fatalf("exit code = %d, want 1; stderr=%s", code, errOut.String())
	}
	if sweepCalls != 0 {
		t.Fatalf("Sweep calls = %d, want 0", sweepCalls)
	}
	if !strings.Contains(errOut.String(), "acquiring lock:") {
		t.Fatalf("stderr = %q, want lock acquisition failure", errOut.String())
	}
}

func TestOrphansUsesPersistedModeOverAmbientMode(t *testing.T) {
	root := t.TempDir()
	writeCleanupState(t, root, engine.State{
		Mode: "organization",
		Plan: &engine.ExecutionPlan{Mode: "organization"},
		Orphans: &engine.OrphanAccounting{
			Mode: "organization",
		},
	})

	orphansCalls := 0
	d := deps{
		provider: &fakeprovider.Fake{
			OrphansFn: func(_ context.Context, mode string) ([]provider.Resource, error) {
				orphansCalls++
				if mode != "organization" {
					t.Fatalf("Orphans mode = %q, want organization", mode)
				}
				return nil, nil
			},
		},
		cwd:    func() (string, error) { return root, nil },
		getenv: getenvFromMap(map[string]string{"GH_TEST_AUTH_MODE": "individual"}),
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"orphans", "--repo-root", root}, &out, &errOut, d)

	if code != 0 {
		t.Fatalf("exit code = %d, want 0; stderr=%s", code, errOut.String())
	}
	if orphansCalls != 1 {
		t.Fatalf("Orphans calls = %d, want 1", orphansCalls)
	}
}

func TestSweepUsesExplicitModeWhenItMatchesState(t *testing.T) {
	root := t.TempDir()
	writeCleanupState(t, root, engine.State{
		Mode: "organization",
		Plan: &engine.ExecutionPlan{Mode: "organization"},
		Orphans: &engine.OrphanAccounting{
			Mode: "organization",
		},
	})

	sweepCalls := 0
	d := deps{
		provider: &fakeprovider.Fake{
			SweepFn: func(_ context.Context, mode string, opts provider.SweepOpts) error {
				sweepCalls++
				if mode != "organization" {
					t.Fatalf("Sweep mode = %q, want organization", mode)
				}
				if !opts.Confirm {
					t.Fatal("Sweep Confirm = false, want true")
				}
				return nil
			},
		},
		cwd:    func() (string, error) { return root, nil },
		getenv: getenvFromMap(map[string]string{"GH_TEST_AUTH_MODE": "individual"}),
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"sweep", "--repo-root", root, "--mode", "organization", "--confirm"}, &out, &errOut, d)

	if code != 0 {
		t.Fatalf("exit code = %d, want 0; stderr=%s", code, errOut.String())
	}
	if sweepCalls != 1 {
		t.Fatalf("Sweep calls = %d, want 1", sweepCalls)
	}
}

func TestCleanupRejectsExplicitModeMismatch(t *testing.T) {
	root := t.TempDir()
	writeCleanupState(t, root, engine.State{
		Mode: "organization",
		Plan: &engine.ExecutionPlan{Mode: "organization"},
		Orphans: &engine.OrphanAccounting{
			Mode: "organization",
			New: []provider.Resource{
				{Kind: "repository", Name: "tf-acc-test-repo", URL: "https://example.test/repo"},
			},
		},
	})

	for _, tc := range []struct {
		name string
		args []string
	}{
		{
			name: "orphans",
			args: []string{"orphans", "--repo-root", root, "--mode", "individual", "--run-delta"},
		},
		{
			name: "sweep",
			args: []string{"sweep", "--repo-root", root, "--mode", "individual", "--run-delta", "--confirm"},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			orphansCalls := 0
			sweepCalls := 0
			d := deps{
				provider: &fakeprovider.Fake{
					OrphansFn: func(_ context.Context, _ string) ([]provider.Resource, error) {
						orphansCalls++
						return nil, nil
					},
					SweepFn: func(_ context.Context, _ string, _ provider.SweepOpts) error {
						sweepCalls++
						return nil
					},
				},
				cwd: func() (string, error) { return root, nil },
			}

			var out, errOut bytes.Buffer
			code := runWithDeps(tc.args, &out, &errOut, d)

			if code != 2 {
				t.Fatalf("exit code = %d, want 2; stderr=%s", code, errOut.String())
			}
			if !strings.Contains(errOut.String(), "mode mismatch") {
				t.Fatalf("stderr = %q, want mode mismatch", errOut.String())
			}
			if orphansCalls != 0 {
				t.Fatalf("Orphans calls = %d, want 0", orphansCalls)
			}
			if sweepCalls != 0 {
				t.Fatalf("Sweep calls = %d, want 0", sweepCalls)
			}
		})
	}
}

func TestCleanupRejectsInternallyMismatchedPersistedModesWithoutExplicitFlag(t *testing.T) {
	root := t.TempDir()
	writeCleanupState(t, root, engine.State{
		Mode: "organization",
		Plan: &engine.ExecutionPlan{Mode: "organization"},
		Orphans: &engine.OrphanAccounting{
			Mode:          "individual",
			CleanupStatus: engine.CleanupComplete,
			New: []provider.Resource{
				{Kind: "repository", Name: "tf-acc-test-repo", URL: "https://example.test/repo"},
			},
		},
	})

	for _, tc := range []struct {
		name string
		args []string
	}{
		{
			name: "orphans",
			args: []string{"orphans", "--repo-root", root, "--run-delta"},
		},
		{
			name: "sweep",
			args: []string{"sweep", "--repo-root", root, "--run-delta", "--confirm"},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			orphansCalls := 0
			sweepCalls := 0
			d := deps{
				provider: &fakeprovider.Fake{
					OrphansFn: func(_ context.Context, _ string) ([]provider.Resource, error) {
						orphansCalls++
						return nil, nil
					},
					SweepFn: func(_ context.Context, _ string, _ provider.SweepOpts) error {
						sweepCalls++
						return nil
					},
				},
				cwd: func() (string, error) { return root, nil },
			}

			var out, errOut bytes.Buffer
			code := runWithDeps(tc.args, &out, &errOut, d)

			if code != 2 {
				t.Fatalf("exit code = %d, want 2; stderr=%s", code, errOut.String())
			}
			if !strings.Contains(errOut.String(), "persisted") || !strings.Contains(errOut.String(), "conflicts") {
				t.Fatalf("stderr = %q, want persisted-mode mismatch guidance", errOut.String())
			}
			if orphansCalls != 0 {
				t.Fatalf("Orphans calls = %d, want 0", orphansCalls)
			}
			if sweepCalls != 0 {
				t.Fatalf("Sweep calls = %d, want 0", sweepCalls)
			}
		})
	}
}

func TestCleanupWithoutModeOrStateIsActionable(t *testing.T) {
	const want = "cleanup mode is unknown; pass --mode or start a plan-bearing run"

	for _, tc := range []struct {
		name string
		args []string
	}{
		{
			name: "orphans",
			args: []string{"orphans"},
		},
		{
			name: "sweep",
			args: []string{"sweep", "--confirm"},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := t.TempDir()
			orphansCalls := 0
			sweepCalls := 0
			d := deps{
				provider: &fakeprovider.Fake{
					OrphansFn: func(_ context.Context, _ string) ([]provider.Resource, error) {
						orphansCalls++
						return nil, nil
					},
					SweepFn: func(_ context.Context, _ string, _ provider.SweepOpts) error {
						sweepCalls++
						return nil
					},
				},
				cwd:    func() (string, error) { return root, nil },
				getenv: getenvFromMap(map[string]string{}),
			}

			var out, errOut bytes.Buffer
			code := runWithDeps(append(tc.args, "--repo-root", root), &out, &errOut, d)

			if code != 2 {
				t.Fatalf("exit code = %d, want 2; stderr=%s", code, errOut.String())
			}
			if strings.TrimSpace(errOut.String()) != want {
				t.Fatalf("stderr = %q, want %q", strings.TrimSpace(errOut.String()), want)
			}
			if orphansCalls != 0 {
				t.Fatalf("Orphans calls = %d, want 0", orphansCalls)
			}
			if sweepCalls != 0 {
				t.Fatalf("Sweep calls = %d, want 0", sweepCalls)
			}
		})
	}
}

func TestCleanupRunDeltaRejectsUnknownAccountingBeforeProviderCalls(t *testing.T) {
	root := t.TempDir()
	writeCleanupState(t, root, engine.State{
		Mode: "organization",
		Plan: &engine.ExecutionPlan{Mode: "organization"},
		Orphans: &engine.OrphanAccounting{
			Mode:          "organization",
			CleanupStatus: engine.CleanupUnknown,
			New: []provider.Resource{{
				Kind: "repository",
				Name: "tf-acc-test-stale",
				URL:  "https://example.test/stale",
			}},
		},
	})

	for _, tc := range []struct {
		name string
		args []string
	}{
		{
			name: "orphans",
			args: []string{"orphans", "--repo-root", root, "--run-delta"},
		},
		{
			name: "sweep",
			args: []string{"sweep", "--repo-root", root, "--run-delta", "--confirm"},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			orphansCalls := 0
			sweepCalls := 0
			d := deps{
				provider: &fakeprovider.Fake{
					OrphansFn: func(_ context.Context, _ string) ([]provider.Resource, error) {
						orphansCalls++
						return nil, nil
					},
					SweepFn: func(_ context.Context, _ string, _ provider.SweepOpts) error {
						sweepCalls++
						return nil
					},
				},
				cwd: func() (string, error) { return root, nil },
			}

			var out, errOut bytes.Buffer
			code := runWithDeps(tc.args, &out, &errOut, d)

			if code != 2 {
				t.Fatalf("exit code = %d, want 2; stderr=%s", code, errOut.String())
			}
			for _, want := range []string{"cleanup status is unknown", "without --run-delta", "read-only inspection"} {
				if !strings.Contains(errOut.String(), want) {
					t.Errorf("stderr = %q, want actionable guidance containing %q", errOut.String(), want)
				}
			}
			if orphansCalls != 0 {
				t.Fatalf("Orphans calls = %d, want 0", orphansCalls)
			}
			if sweepCalls != 0 {
				t.Fatalf("Sweep calls = %d, want 0", sweepCalls)
			}
		})
	}
}

func TestCleanupSelectionDoesNotLeakPersistedOrphansBetweenCalls(t *testing.T) {
	leakedResources := []provider.Resource{
		{Kind: "repository", Name: "tf-acc-test-repo", URL: "https://example.test/repo"},
	}
	first, err := resolveCleanupSelection(engine.State{
		Mode: "organization",
		Orphans: &engine.OrphanAccounting{
			Mode: "organization",
			New:  leakedResources,
		},
	}, "", false, true)
	if err != nil {
		t.Fatalf("first resolveCleanupSelection error = %v", err)
	}
	if first.Mode != "organization" {
		t.Fatalf("first cleanup mode = %q, want organization", first.Mode)
	}
	if !reflect.DeepEqual(first.Resources, leakedResources) {
		t.Fatalf("first cleanup resources = %#v, want %#v", first.Resources, leakedResources)
	}

	second, err := resolveCleanupSelection(engine.State{Mode: "individual"}, "", false, true)
	if err != nil {
		t.Fatalf("second resolveCleanupSelection error = %v", err)
	}
	if second.Mode != "individual" {
		t.Fatalf("second cleanup mode = %q, want individual", second.Mode)
	}
	if len(second.Resources) != 0 {
		t.Fatalf("second cleanup resources = %#v, want empty slice", second.Resources)
	}
}

func TestOrphansRunDeltaPrintsOnlyPersistedNewResources(t *testing.T) {
	root := t.TempDir()
	writeCleanupState(t, root, engine.State{
		Mode: "organization",
		Plan: &engine.ExecutionPlan{Mode: "organization"},
		Orphans: &engine.OrphanAccounting{
			Mode: "organization",
			New: []provider.Resource{
				{Kind: "repository", Name: "tf-acc-test-repo", URL: "https://example.test/repo"},
				{Kind: "team", Name: "tf-acc-test-team", URL: "https://example.test/team"},
			},
		},
	})

	orphansCalls := 0
	d := deps{
		provider: &fakeprovider.Fake{
			OrphansFn: func(_ context.Context, _ string) ([]provider.Resource, error) {
				orphansCalls++
				return nil, nil
			},
		},
		cwd:    func() (string, error) { return root, nil },
		getenv: getenvFromMap(map[string]string{"GH_TEST_AUTH_MODE": "individual"}),
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"orphans", "--repo-root", root, "--run-delta"}, &out, &errOut, d)

	if code != 0 {
		t.Fatalf("exit code = %d, want 0; stderr=%s", code, errOut.String())
	}
	if orphansCalls != 0 {
		t.Fatalf("Orphans calls = %d, want 0", orphansCalls)
	}
	for _, want := range []string{"tf-acc-test-repo", "tf-acc-test-team"} {
		if !strings.Contains(out.String(), want) {
			t.Fatalf("stdout = %q, want %q", out.String(), want)
		}
	}
}

func TestOrphansRedactsLiveAndRunDeltaOutputWithoutChangingResources(t *testing.T) {
	const secret = "ghp_ORPHANOUTPUTSECRET1234567890"
	resource := provider.Resource{
		Kind: "repository",
		Name: "tf-acc-test-secret-url",
		URL:  "https://example.test/repos/tf-acc-test-secret-url?token=" + secret,
	}
	for _, source := range []string{"live", "run-delta"} {
		for _, format := range []string{"text", "json"} {
			t.Run(source+"/"+format, func(t *testing.T) {
				root := t.TempDir()
				args := []string{"orphans", "--repo-root", root}
				if source == "live" {
					args = append(args, "--mode", "organization")
				} else {
					writeCleanupState(t, root, engine.State{
						Mode: "organization",
						Plan: &engine.ExecutionPlan{Mode: "organization"},
						Orphans: &engine.OrphanAccounting{
							Mode:          "organization",
							New:           []provider.Resource{resource},
							CleanupStatus: engine.CleanupComplete,
						},
					})
					args = append(args, "--run-delta")
				}
				if format == "json" {
					args = append(args, "--json")
				}

				liveResources := []provider.Resource{resource}
				d := deps{
					provider: &fakeprovider.Fake{
						OrphansVal: liveResources,
						Secrets:    []string{"GITHUB_TOKEN"},
					},
					cwd:    func() (string, error) { return root, nil },
					getenv: getenvFromMap(map[string]string{"GITHUB_TOKEN": secret}),
				}

				var out, errOut bytes.Buffer
				code := runWithDeps(args, &out, &errOut, d)

				if code != 0 {
					t.Fatalf("exit code = %d, want 0; stderr=%s", code, errOut.String())
				}
				if strings.Contains(out.String(), secret) {
					t.Fatalf("%s %s output leaked secret-shaped URL:\n%s", source, format, out.String())
				}
				if !strings.Contains(out.String(), "***REDACTED***") {
					t.Fatalf("%s %s output missing redaction marker:\n%s", source, format, out.String())
				}
				if !reflect.DeepEqual(liveResources, []provider.Resource{resource}) {
					t.Fatalf("live provider resources mutated: %#v", liveResources)
				}
				if source == "run-delta" {
					persisted, err := engine.Load(filepath.Join(root, ".pulsar-state.json"))
					if err != nil {
						t.Fatal(err)
					}
					if persisted.Orphans == nil ||
						!reflect.DeepEqual(persisted.Orphans.New, []provider.Resource{resource}) {
						t.Fatalf("persisted run-delta identity changed: %+v", persisted.Orphans)
					}
				}
			})
		}
	}
}

func TestSweepRunDeltaPassesExactPersistedResources(t *testing.T) {
	root := t.TempDir()
	wantResources := []provider.Resource{
		{Kind: "repository", Name: "tf-acc-test-repo", URL: "https://example.test/repo"},
		{Kind: "team", Name: "tf-acc-test-team", URL: "https://example.test/team"},
	}
	writeCleanupState(t, root, engine.State{
		Mode: "organization",
		Plan: &engine.ExecutionPlan{Mode: "organization"},
		Orphans: &engine.OrphanAccounting{
			Mode: "organization",
			New:  wantResources,
		},
	})

	sweepCalls := 0
	d := deps{
		provider: &fakeprovider.Fake{
			SweepFn: func(_ context.Context, mode string, opts provider.SweepOpts) error {
				sweepCalls++
				if mode != "organization" {
					t.Fatalf("Sweep mode = %q, want organization", mode)
				}
				if !opts.Confirm {
					t.Fatal("Sweep Confirm = false, want true")
				}
				if opts.Resources == nil {
					t.Fatal("Sweep Resources = nil, want non-nil persisted subset")
				}
				if !reflect.DeepEqual(opts.Resources, wantResources) {
					t.Fatalf("Sweep Resources = %#v, want %#v", opts.Resources, wantResources)
				}
				return nil
			},
		},
		cwd:    func() (string, error) { return root, nil },
		getenv: getenvFromMap(map[string]string{"GH_TEST_AUTH_MODE": "individual"}),
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"sweep", "--repo-root", root, "--run-delta", "--confirm"}, &out, &errOut, d)

	if code != 0 {
		t.Fatalf("exit code = %d, want 0; stderr=%s", code, errOut.String())
	}
	if sweepCalls != 1 {
		t.Fatalf("Sweep calls = %d, want 1", sweepCalls)
	}
}

func TestSweepRunDeltaEmptyIsNoOp(t *testing.T) {
	root := t.TempDir()
	writeCleanupState(t, root, engine.State{
		Mode: "organization",
		Plan: &engine.ExecutionPlan{Mode: "organization"},
		Orphans: &engine.OrphanAccounting{
			Mode: "organization",
			New:  []provider.Resource{},
		},
	})

	sweepCalls := 0
	d := deps{
		provider: &fakeprovider.Fake{
			SweepFn: func(_ context.Context, _ string, opts provider.SweepOpts) error {
				sweepCalls++
				if opts.Resources == nil {
					t.Fatal("Sweep Resources = nil, want explicit empty slice")
				}
				if len(opts.Resources) != 0 {
					t.Fatalf("Sweep Resources len = %d, want 0", len(opts.Resources))
				}
				return nil
			},
		},
		cwd: func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"sweep", "--repo-root", root, "--run-delta", "--confirm"}, &out, &errOut, d)

	if code != 0 {
		t.Fatalf("exit code = %d, want 0; stderr=%s", code, errOut.String())
	}
	if sweepCalls != 1 {
		t.Fatalf("Sweep calls = %d, want 1", sweepCalls)
	}
}

func TestSweepChecksConfirmationBeforeProviderCall(t *testing.T) {
	root := t.TempDir()
	writeCleanupState(t, root, engine.State{
		Mode: "organization",
		Plan: &engine.ExecutionPlan{Mode: "organization"},
	})

	sweepCalls := 0
	d := deps{
		provider: &fakeprovider.Fake{
			SweepFn: func(_ context.Context, _ string, _ provider.SweepOpts) error {
				sweepCalls++
				return nil
			},
		},
		cwd: func() (string, error) { return root, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"sweep", "--repo-root", root, "--mode", "organization"}, &out, &errOut, d)

	if code != 1 {
		t.Fatalf("exit code = %d, want 1; stderr=%s", code, errOut.String())
	}
	if sweepCalls != 0 {
		t.Fatalf("Sweep calls = %d, want 0", sweepCalls)
	}
	if !strings.Contains(errOut.String(), "sweep requires --confirm") {
		t.Fatalf("stderr = %q, want confirmation error", errOut.String())
	}
}
