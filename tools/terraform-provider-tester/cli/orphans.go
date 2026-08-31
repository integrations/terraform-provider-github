package cli

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"path/filepath"

	"github.com/github/terraform-provider-tester/engine"
	"github.com/github/terraform-provider-tester/provider"
)

var (
	errCleanupModeUnknown  = errors.New("cleanup mode is unknown; pass --mode or start a plan-bearing run")
	errCleanupDeltaUnknown = errors.New(
		"persisted run delta is unavailable because cleanup status is unknown; rerun resume or use orphans without --run-delta for full read-only inspection",
	)
)

type cleanupSelection struct {
	Mode      string
	RunDelta  bool
	Resources []provider.Resource
}

func cloneResources(resources []provider.Resource) []provider.Resource {
	if resources == nil {
		return nil
	}
	cloned := make([]provider.Resource, len(resources))
	copy(cloned, resources)
	return cloned
}

func persistedCleanupMode(state engine.State) (string, error) {
	type persistedMode struct {
		source string
		mode   string
	}
	var persisted []persistedMode
	if state.Orphans != nil && state.Orphans.Mode != "" {
		persisted = append(persisted, persistedMode{source: "orphan accounting", mode: state.Orphans.Mode})
	}
	if state.Plan != nil && state.Plan.Mode != "" {
		persisted = append(persisted, persistedMode{source: "plan", mode: state.Plan.Mode})
	}
	if state.Mode != "" {
		persisted = append(persisted, persistedMode{source: "state", mode: state.Mode})
	}
	if len(persisted) == 0 {
		return "", nil
	}
	selected := persisted[0]
	for _, candidate := range persisted[1:] {
		if candidate.mode != selected.mode {
			return "", fmt.Errorf(
				"state/mode-mismatch: persisted %s mode %q conflicts with persisted %s mode %q; start a new guided `e2e` run or pass --mode explicitly",
				selected.source, selected.mode, candidate.source, candidate.mode,
			)
		}
	}
	return selected.mode, nil
}

func resolveCleanupSelection(
	state engine.State,
	explicitMode string,
	modeExplicit bool,
	runDelta bool,
) (cleanupSelection, error) {
	var selection cleanupSelection
	selection.RunDelta = runDelta

	persistedModes := make([]string, 0, 3)
	addPersistedMode := func(mode string) {
		if mode == "" {
			return
		}
		persistedModes = append(persistedModes, mode)
	}
	if state.Orphans != nil {
		addPersistedMode(state.Orphans.Mode)
	}
	if state.Plan != nil {
		addPersistedMode(state.Plan.Mode)
	}
	addPersistedMode(state.Mode)

	if modeExplicit {
		if runDelta {
			for _, persistedMode := range persistedModes {
				if explicitMode != persistedMode {
					return cleanupSelection{}, fmt.Errorf("mode mismatch: --mode %q conflicts with the persisted logical-run mode %q; omit --mode or start a new run", explicitMode, persistedMode)
				}
			}
		}
		selection.Mode = explicitMode
	} else {
		persistedMode, err := persistedCleanupMode(state)
		if err != nil {
			return cleanupSelection{}, err
		}
		if persistedMode == "" {
			return cleanupSelection{}, errCleanupModeUnknown
		}
		selection.Mode = persistedMode
	}

	if runDelta {
		if state.Orphans != nil && state.Orphans.CleanupStatus == engine.CleanupUnknown {
			return cleanupSelection{}, errCleanupDeltaUnknown
		}
		if state.Orphans == nil || state.Orphans.New == nil {
			selection.Resources = []provider.Resource{}
		} else {
			selection.Resources = cloneResources(state.Orphans.New)
		}
	}

	return selection, nil
}

func runOrphans(args []string, out, errOut io.Writer, d deps) int {
	return runOrphansContext(context.Background(), args, out, errOut, d)
}

func runOrphansContext(ctx context.Context, args []string, out, errOut io.Writer, d deps) int {
	fs := flag.NewFlagSet("orphans", flag.ContinueOnError)
	fs.SetOutput(errOut)
	var (
		repoRoot string
		modeFlag string
		runDelta bool
		format   = formatText
	)
	addJSONFlag(fs, &format)
	fs.StringVar(&repoRoot, "repo-root", "", "repo root override")
	fs.StringVar(&modeFlag, "mode", "", "cleanup mode")
	fs.BoolVar(&runDelta, "run-delta", false, "operate only on persisted new resources")
	if err := fs.Parse(args); err != nil {
		return 2
	}
	modeExplicit := false
	fs.Visit(func(f *flag.Flag) {
		if f.Name == "mode" {
			modeExplicit = true
		}
	})

	root, err := d.resolveRoot(repoRoot)
	if err != nil {
		fmt.Fprintln(errOut, "resolving repo root:", err)
		return 1
	}

	statePath := filepath.Join(root, ".pulsar-state.json")
	lockPath := statePath + ".lock"
	release, err := engine.AcquireLock(lockPath)
	if err != nil {
		fmt.Fprintln(errOut, "acquiring lock:", err)
		return 1
	}
	defer release() //nolint:errcheck

	state, err := engine.Load(statePath)
	if err != nil {
		fmt.Fprintln(errOut, "loading state:", err)
		return 1
	}

	selection, err := resolveCleanupSelection(state, modeFlag, modeExplicit, runDelta)
	if err != nil {
		fmt.Fprintln(errOut, err)
		return 2
	}

	resources := selection.Resources
	if !runDelta {
		resources, err = d.provider.Orphans(ctx, selection.Mode)
		if err != nil {
			fmt.Fprintln(errOut, "listing orphans:", err)
			if format == formatJSON {
				_ = newJSONWriter(out).Encode(jsonOrphansSummary{
					SchemaVersion: 1,
					Type:          "summary",
					Command:       "orphans",
					Mode:          selection.Mode,
					Resources:     0,
					ExitCode:      1,
				})
			}
			return 1
		}
	}
	resources = redactOrphanResources(resources, redactorForProvider(d.provider, getenvFunc(d)))

	if format == formatJSON {
		jw := newJSONWriter(out)
		for _, resource := range resources {
			if err := jw.Encode(jsonOrphanEvent{
				SchemaVersion: 1,
				Type:          "orphan",
				Kind:          resource.Kind,
				Name:          resource.Name,
				URL:           resource.URL,
			}); err != nil {
				fmt.Fprintln(errOut, "writing JSON:", err)
				return 1
			}
		}
		if err := jw.Encode(jsonOrphansSummary{
			SchemaVersion: 1,
			Type:          "summary",
			Command:       "orphans",
			Mode:          selection.Mode,
			RunDelta:      runDelta,
			Resources:     len(resources),
			ExitCode:      0,
		}); err != nil {
			fmt.Fprintln(errOut, "writing JSON:", err)
			return 1
		}
		return 0
	}

	if len(resources) == 0 {
		fmt.Fprintln(out, "no orphaned resources found")
		return 0
	}

	fmt.Fprintf(out, "%-20s %-40s %s\n", "KIND", "NAME", "URL")
	for _, r := range resources {
		fmt.Fprintf(out, "%-20s %-40s %s\n", r.Kind, r.Name, r.URL)
	}
	return 0
}

func runSweep(args []string, out, errOut io.Writer, d deps) int {
	fs := flag.NewFlagSet("sweep", flag.ContinueOnError)
	fs.SetOutput(errOut)
	var (
		confirm  bool
		repoRoot string
		modeFlag string
		runDelta bool
	)
	fs.BoolVar(&confirm, "confirm", false, "confirm destructive sweep")
	fs.StringVar(&repoRoot, "repo-root", "", "repo root override")
	fs.StringVar(&modeFlag, "mode", "", "cleanup mode")
	fs.BoolVar(&runDelta, "run-delta", false, "operate only on persisted new resources")
	if err := fs.Parse(args); err != nil {
		return 2
	}
	modeExplicit := false
	fs.Visit(func(f *flag.Flag) {
		if f.Name == "mode" {
			modeExplicit = true
		}
	})

	root, err := d.resolveRoot(repoRoot)
	if err != nil {
		fmt.Fprintln(errOut, "resolving repo root:", err)
		return 1
	}

	statePath := filepath.Join(root, ".pulsar-state.json")
	lockPath := statePath + ".lock"
	release, err := engine.AcquireLock(lockPath)
	if err != nil {
		fmt.Fprintln(errOut, "acquiring lock:", err)
		return 1
	}
	defer release() //nolint:errcheck

	state, err := engine.Load(statePath)
	if err != nil {
		fmt.Fprintln(errOut, "loading state:", err)
		return 1
	}

	selection, err := resolveCleanupSelection(state, modeFlag, modeExplicit, runDelta)
	if err != nil {
		fmt.Fprintln(errOut, err)
		return 2
	}

	if !confirm {
		fmt.Fprintln(errOut, "sweep requires --confirm to prevent accidental deletion of tf-acc-test-* resources")
		return 1
	}

	opts := provider.SweepOpts{
		Targets: []string{"repositories", "teams"},
		Confirm: true,
	}
	if runDelta {
		opts.Resources = selection.Resources
	}
	if err := d.provider.Sweep(context.Background(), selection.Mode, opts); err != nil {
		fmt.Fprintln(errOut, "sweep error:", err)
		return 1
	}
	fmt.Fprintln(out, "sweep completed")
	return 0
}
