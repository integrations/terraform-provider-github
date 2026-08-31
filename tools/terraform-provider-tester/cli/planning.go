package cli

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/github/terraform-provider-tester/engine"
	"github.com/github/terraform-provider-tester/internal/redact"
	"github.com/github/terraform-provider-tester/provider"
)

// planningOptions selects and classifies tests for a single planning pass,
// shared by the run and preflight subcommands.
type planningOptions struct {
	Mode              string
	Group             string
	Run               string
	AllowUnclassified bool
}

// planForCommand discovers tests under root and builds the ExecutionPlan that
// governs everything downstream: which tests are eligible, which requirements
// preflight must verify, and the exact pattern the runner will use. It also
// returns the discovered names' group buckets (engine.GroupTests) for
// JSON/text selection reporting.
//
// Discovery always happens before any planning or preflight decision, so a
// build/listing failure is reported the same way regardless of mode or
// selection flags, and no preflight or runner call is ever reached once
// discovery or planning has failed.
func planForCommand(ctx context.Context, d deps, root string, opts planningOptions) (engine.ExecutionPlan, []engine.Group, error) {
	prov := d.provider
	names, err := d.list(ctx, root, prov.TestPackages(), prov.TestPattern())
	if err != nil {
		return engine.ExecutionPlan{}, nil, fmt.Errorf("listing tests: %w", err)
	}

	groups, _ := engine.GroupTests(names, prov)

	plan, err := engine.BuildExecutionPlan(names, prov, engine.PlanOptions{
		Mode:              opts.Mode,
		Group:             opts.Group,
		Run:               opts.Run,
		AllowUnclassified: opts.AllowUnclassified,
	})
	if err != nil {
		return engine.ExecutionPlan{}, nil, err
	}
	return plan, groups, nil
}

// planExitCode maps a planForCommand error to a process exit code. An
// invalid --run regex is a usage error (2, matching flag-parse failures);
// every other planning error (unknown mode/group, an unclassified test
// without --allow-unclassified, an explicit mode-incompatible selection) and
// any discovery failure is a runtime error (1).
func planExitCode(err error) int {
	var perr *engine.PlanError
	if errors.As(err, &perr) && perr.Code == "plan/regex-invalid" {
		return 2
	}
	return 1
}

// redactPlan returns a copy of plan with every Excluded entry's Detail and Fix
// passed through red.String. Excluded detail/fix text interpolates test names
// and required-mode lists (see engine.BuildExecutionPlan), so it is the only
// ExecutionPlan field that can ever carry operator-controlled text; this is
// the single point that redacts the plan before it is emitted or saved to
// state, so both destinations see identical, already-redacted text. Scopes,
// Capabilities, and SideEffects are static catalog vocabulary, never secrets,
// and are copied through unchanged.
func redactPlan(plan engine.ExecutionPlan, red *redact.Redactor) engine.ExecutionPlan {
	if len(plan.Excluded) == 0 {
		return plan
	}
	redacted := make([]engine.PlanExclusion, len(plan.Excluded))
	for i, ex := range plan.Excluded {
		redacted[i] = ex
		redacted[i].Detail = red.String(ex.Detail)
		redacted[i].Fix = red.String(ex.Fix)
	}
	plan.Excluded = redacted
	return plan
}

// emitPlan renders plan as the first phase of run/preflight output: in JSON
// mode a single jsonPlanEvent followed by one jsonExcludedEvent per
// plan.Excluded entry; in text mode the exact mode-summary line followed by
// one warning line per allowed-unclassified test. An unclassified test whose
// conservative requirements are incompatible with the mode is reported as
// excluded rather than as "using conservative requirements", because it will
// not run. Callers must pass an already redacted plan (see redactPlan) —
// emitPlan does not redact.
func emitPlan(w io.Writer, jw *jsonWriter, format outputFormat, plan engine.ExecutionPlan) error {
	if format == formatJSON {
		if err := jw.Encode(jsonPlanEvent{
			SchemaVersion: 1,
			Type:          "plan",
			Mode:          plan.Mode,
			Selected:      len(plan.Selected),
			Eligible:      len(plan.Eligible),
			Excluded:      len(plan.Excluded),
			Unclassified:  len(plan.Unclassified),
			Scopes:        plan.Scopes,
			Capabilities:  plan.Capabilities,
			SideEffects:   plan.SideEffects,
		}); err != nil {
			return err
		}
		for _, ex := range plan.Excluded {
			if err := jw.Encode(jsonExcludedEvent{
				SchemaVersion: 1,
				Type:          "excluded",
				Mode:          plan.Mode,
				Test:          ex.Test,
				Code:          ex.Code,
				Detail:        ex.Detail,
				Fix:           stringPtr(ex.Fix),
				RequiredModes: ex.RequiredModes,
			}); err != nil {
				return err
			}
		}
		return nil
	}

	fmt.Fprintf(w, "mode %s: %d selected, %d eligible, %d excluded, %d unclassified\n",
		plan.Mode, len(plan.Selected), len(plan.Eligible), len(plan.Excluded), len(plan.Unclassified))
	excluded := make(map[string]bool, len(plan.Excluded))
	for _, ex := range plan.Excluded {
		excluded[ex.Test] = true
	}
	for _, test := range plan.Unclassified {
		if excluded[test] {
			fmt.Fprintf(w, "warning: %s is unclassified and excluded from mode %s; its conservative authenticated requirements are not compatible\n", test, plan.Mode)
			continue
		}
		fmt.Fprintf(w, "warning: %s is unclassified; using conservative authenticated requirements and side effect unknown\n", test)
	}
	return nil
}

// emitCapabilities writes one jsonCapabilityEvent per report.Check, applying
// red.String to Detail and Fix inline. Unlike plan exclusions, checks are
// never persisted to state, so there is only one consumer of this text and no
// separate pre-redacted copy is needed.
func emitCapabilities(jw *jsonWriter, mode string, report provider.PreflightReport, red *redact.Redactor) error {
	for _, c := range report.Checks {
		if err := jw.Encode(jsonCapabilityEvent{
			SchemaVersion: 1,
			Type:          "capability",
			Mode:          mode,
			Name:          c.Name,
			Status:        c.Status.String(),
			Detail:        red.String(c.Detail),
			Fix:           stringPtr(red.String(c.Fix)),
		}); err != nil {
			return err
		}
	}
	return nil
}
