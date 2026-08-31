package engine

import (
	"fmt"
	"regexp"
	"sort"

	"github.com/github/terraform-provider-tester/provider"
)

// PlanOptions controls how BuildExecutionPlan selects and classifies tests.
type PlanOptions struct {
	Mode              string
	Group             string
	Run               string // regexp filter applied to test names; empty = all
	AllowUnclassified bool   // when true, unknown tests are allowed and warned rather than rejected
}

// PlanExclusion records a single test that was excluded during planning.
type PlanExclusion struct {
	Test          string   `json:"test"`
	Code          string   `json:"code"`
	Detail        string   `json:"detail"`
	Fix           string   `json:"fix,omitempty"`
	RequiredModes []string `json:"required_modes,omitempty"`
}

// ExecutionPlan is the fully resolved output of BuildExecutionPlan.
// Selected, Eligible, Excluded, and Unclassified preserve discovery order;
// Scopes, Capabilities, and SideEffects are sorted unique aggregates stored on
// the plan.
type ExecutionPlan struct {
	Mode  string `json:"mode"`
	Group string `json:"group,omitempty"`
	Run   string `json:"run,omitempty"`
	// Selected contains every test that survived the Run/Group filter.
	Selected []string `json:"selected"`
	// Eligible contains the subset of Selected that is compatible with Mode.
	Eligible []string `json:"eligible"`
	// Excluded contains the subset of Selected that was filtered out,
	// each entry annotated with a stable code.
	Excluded []PlanExclusion `json:"excluded"`
	// Unclassified contains selected tests for which RequirementsFor returned
	// false, accepted only when AllowUnclassified is true. An entry here is
	// also in Eligible when its conservative requirements support Mode, and in
	// Excluded otherwise.
	Unclassified []string `json:"unclassified"`
	// Scopes, Capabilities, and SideEffects are aggregated from Eligible,
	// which includes any accepted Unclassified test, deduped and sorted.
	Scopes       []string `json:"scopes"`
	Capabilities []string `json:"capabilities"`
	SideEffects  []string `json:"side_effects"`
}

// PlanError is a structured planning-time error with a stable code.
type PlanError struct {
	Code          string
	Tests         []string
	RequiredModes []string
	Detail        string
	Fix           string
}

// Error implements the error interface.
func (e *PlanError) Error() string {
	return e.Detail
}

// BuildExecutionPlan resolves which tests from names are eligible to run in
// opts.Mode and returns an ExecutionPlan with aggregated requirements stored
// on the plan itself.
//
// Algorithm:
//  1. Validate opts.Mode against prov.Modes().
//  2. Compile opts.Run before selection.
//  3. Resolve opts.Group through GroupTests; reject unknown non-empty groups.
//     A non-empty opts.Run overrides opts.Group: the group is not resolved,
//     not validated, and not applied, so the two filters are never
//     intersected.
//  4. Preserve discovery order in Selected, Eligible, Excluded, Unclassified.
//  5. Look up every selected name via prov.RequirementsFor.
//  6. Fail closed on unknown names unless AllowUnclassified.
//  7. For explicit regex selection, return plan/mode-incompatible if any
//     selected test excludes the mode.
//  8. For default/group selection, append PlanExclusion{Code: "excluded/mode-incompatible"}.
//  9. Aggregate only eligible requirements using sorted unique slices.
//  10. An allowed unknown is classified with the conservative requirements
//     RequirementsFor returned alongside ok=false, then run through the SAME
//     mode-compatibility rules as a known test: compatible means eligible and
//     aggregated (plus side effect "unknown"), incompatible means
//     plan/mode-incompatible for an explicit selection and
//     excluded/mode-incompatible for a default/group selection. Either way it
//     is listed in Unclassified so the operator sees it was not classified.
func BuildExecutionPlan(names []string, prov provider.Provider, opts PlanOptions) (ExecutionPlan, error) {
	// Step 1: validate mode.
	if !providerHasMode(prov, opts.Mode) {
		return ExecutionPlan{}, &PlanError{
			Code:   "plan/mode-unknown",
			Detail: fmt.Sprintf("mode %q is not supported by this provider", opts.Mode),
			Fix:    "choose one of the provider's supported modes",
		}
	}

	// Step 2: compile Run regexp before touching names.
	var runRE *regexp.Regexp
	explicit := opts.Run != ""
	if explicit {
		var err error
		runRE, err = regexp.Compile(opts.Run)
		if err != nil {
			return ExecutionPlan{}, &PlanError{
				Code:   "plan/regex-invalid",
				Detail: fmt.Sprintf("invalid run pattern %q: %v", opts.Run, err),
				Fix:    "provide a valid Go regular expression",
			}
		}
	}

	// Step 3: resolve group filter; reject unknown non-empty group names.
	// A non-empty Run overrides Group entirely (the documented contract of
	// `--run`: "run regex (overrides --group)"), so the group is neither
	// resolved nor validated nor applied — the two filters are never
	// intersected. opts.Group is still echoed on the plan as a faithful
	// record of what the operator asked for.
	var groupSet map[string]bool
	if opts.Group != "" && !explicit {
		groups, _ := GroupTests(names, prov)
		found := false
		for _, g := range groups {
			if g.Name == opts.Group {
				found = true
				groupSet = make(map[string]bool, len(g.Tests))
				for _, n := range g.Tests {
					groupSet[n] = true
				}
				break
			}
		}
		if !found {
			return ExecutionPlan{}, &PlanError{
				Code:   "plan/group-unknown",
				Detail: fmt.Sprintf("group %q does not exist in the provider's classification", opts.Group),
				Fix:    "use a known group name returned by the provider",
			}
		}
	}

	// Step 4: select names in discovery order.
	selected := make([]string, 0)
	for _, n := range names {
		if groupSet != nil && !groupSet[n] {
			continue
		}
		if runRE != nil && !runRE.MatchString(n) {
			continue
		}
		selected = append(selected, n)
	}

	// Steps 5-10: classify each selected name.
	eligible := make([]string, 0)
	excluded := make([]PlanExclusion, 0)
	unclassified := make([]string, 0)

	scopeSet := make(map[string]bool)
	capSet := make(map[string]bool)
	seSet := make(map[string]bool)
	hasUnclassified := false

	for _, n := range selected {
		req, ok := prov.RequirementsFor(n)
		unknown := !ok
		if unknown {
			// Unknown test (step 6): fail closed unless explicitly allowed.
			if !opts.AllowUnclassified {
				return ExecutionPlan{}, &PlanError{
					Code:   "plan/unclassified",
					Tests:  []string{n},
					Detail: fmt.Sprintf("test %q has no requirements classification; pass --allow-unclassified to include it", n),
					Fix:    "classify the test in the provider catalog or pass --allow-unclassified",
				}
			}
			// Step 10: allowed unknown. Record it as unclassified for
			// reporting, then fall through to the SAME mode-compatibility and
			// aggregation rules a classified test gets, using the
			// conservative requirements RequirementsFor returned. Falling
			// through is what makes the test actually runnable: appending to
			// Unclassified alone left it out of Eligible, so nothing ever ran
			// it even though its requirements were aggregated into the plan
			// and verified by preflight.
			unclassified = append(unclassified, n)
		}

		// Step 7 / 8: check mode compatibility.
		if !modeCompatible(req, opts.Mode) {
			if explicit {
				// Explicit regex selection of an incompatible test is an error.
				return ExecutionPlan{}, &PlanError{
					Code:          "plan/mode-incompatible",
					Tests:         []string{n},
					RequiredModes: req.Modes,
					Detail:        fmt.Sprintf("test %q does not support mode %q (supports: %v)", n, opts.Mode, req.Modes),
					Fix:           "choose a mode the test supports, or remove it from the explicit selection",
				}
			}
			// Default/group selection: record as excluded, do not error.
			excluded = append(excluded, PlanExclusion{
				Test:          n,
				Code:          "excluded/mode-incompatible",
				Detail:        fmt.Sprintf("test %q does not support mode %q (supports: %v)", n, opts.Mode, req.Modes),
				Fix:           fmt.Sprintf("run in one of: %v", req.Modes),
				RequiredModes: req.Modes,
			})
			continue
		}

		// Step 9: eligible — aggregate requirements. An accepted unknown
		// contributes the extra "unknown" side effect below; an excluded one
		// contributes nothing, so an excluded unknown's over-broad
		// conservative scopes never inflate what preflight demands.
		eligible = append(eligible, n)
		aggregateInto(req, scopeSet, capSet, seSet)
		if unknown {
			hasUnclassified = true
		}
	}

	// Step 10 side effect: inject "unknown" when any unclassified test was accepted.
	if hasUnclassified {
		seSet["unknown"] = true
	}

	return ExecutionPlan{
		Mode:         opts.Mode,
		Group:        opts.Group,
		Run:          opts.Run,
		Selected:     selected,
		Eligible:     eligible,
		Excluded:     excluded,
		Unclassified: unclassified,
		Scopes:       sortedUniqueKeys(scopeSet),
		Capabilities: sortedUniqueKeys(capSet),
		SideEffects:  sortedUniqueKeys(seSet),
	}, nil
}

// providerHasMode reports whether mode is in prov.Modes().
func providerHasMode(prov provider.Provider, mode string) bool {
	for _, m := range prov.Modes() {
		if m.Name == mode {
			return true
		}
	}
	return false
}

// modeCompatible reports whether req is compatible with mode.
// A test with an empty Modes list has no restriction and is always compatible.
func modeCompatible(req provider.TestRequirements, mode string) bool {
	if len(req.Modes) == 0 {
		return true
	}
	for _, m := range req.Modes {
		if m == mode {
			return true
		}
	}
	return false
}

// aggregateInto merges req's Scopes, Capabilities, and SideEffects into the
// corresponding sets (deduplication via map membership).
func aggregateInto(req provider.TestRequirements, scopes, caps, effects map[string]bool) {
	for _, s := range req.Scopes {
		scopes[s] = true
	}
	for _, c := range req.Capabilities {
		caps[c] = true
	}
	for _, se := range req.SideEffects {
		effects[se] = true
	}
}

// buildRequirements converts the three accumulator sets into sorted, unique,
// non-nil empty slices (kept for future use; currently inlined into ExecutionPlan).
func buildRequirements(scopes, caps, effects map[string]bool) ([]string, []string, []string) {
	return sortedUniqueKeys(scopes), sortedUniqueKeys(caps), sortedUniqueKeys(effects)
}

// sortedUniqueKeys returns the keys of m in ascending order, or a non-nil empty
// slice if m is empty.
func sortedUniqueKeys(m map[string]bool) []string {
	if len(m) == 0 {
		return []string{}
	}
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
