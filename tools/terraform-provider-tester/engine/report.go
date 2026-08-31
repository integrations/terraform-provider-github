package engine

import (
	"fmt"
	"html/template"
	"io"
	"os"
	"sort"

	"github.com/github/terraform-provider-tester/internal/redact"
	"github.com/github/terraform-provider-tester/provider"
)

// redactLines applies r to each line; a nil r is a no-op that returns in unchanged.
func redactLines(r *redact.Redactor, in []string) []string {
	if r == nil {
		return in
	}
	return r.Lines(in)
}

// groupCounts tallies pass/fail/skip over top-level results (Sub == "").
// When nameSet is non-nil, only tests whose Name is in nameSet are counted.
func groupCounts(tests []TestResult, nameSet map[string]bool) (pass, fail, skip int) {
	for _, tr := range tests {
		if tr.Sub != "" {
			continue
		}
		if nameSet != nil && !nameSet[tr.Name] {
			continue
		}
		switch tr.Status {
		case provider.StatusPass:
			pass++
		case provider.StatusFail, provider.StatusPanic, provider.StatusTimeout:
			fail++
		case provider.StatusSkip:
			skip++
		}
	}
	return
}

// TerminalSummary writes a human-readable run summary to w.
//
// Build-fail trap: if res.BuildFailed, prints "BUILD FAILED" - never a success
// line. Pre-run-fail trap: if res.PreRunFailed, prints "SUITE EXITED BEFORE
// RUNNING TESTS" plus the captured output - never a success line and never
// mislabeled as a build failure. When BOTH flags are set (possible across a
// multi-package ./github/... run), both sections are printed so neither
// condition is hidden, then the function returns without a success line.
func TerminalSummary(w io.Writer, groups []Group, res RunResult, redactors ...*redact.Redactor) {
	var red *redact.Redactor
	if len(redactors) > 0 {
		red = redactors[0]
	}
	if res.BuildFailed || res.PreRunFailed {
		if res.BuildFailed {
			fmt.Fprintln(w, "BUILD FAILED")
			for _, line := range redactLines(red, res.BuildOutput) {
				fmt.Fprint(w, line)
			}
		}
		if res.PreRunFailed {
			fmt.Fprintln(w, "SUITE EXITED BEFORE RUNNING TESTS")
			for _, line := range redactLines(red, res.PreRunOutput) {
				fmt.Fprintln(w, line)
			}
		}
		return
	}

	pass, fail, skip := groupCounts(res.Tests, nil)
	fmt.Fprintf(w, "%d passed, %d failed, %d skipped\n", pass, fail, skip)

	for _, g := range groups {
		ns := make(map[string]bool, len(g.Tests))
		for _, name := range g.Tests {
			ns[name] = true
		}
		gp, gf, gs := groupCounts(res.Tests, ns)
		fmt.Fprintf(w, "%s: %d passed, %d failed, %d skipped\n", g.Name, gp, gf, gs)
	}
}

// SlowestN returns up to n TestResults ordered by Elapsed descending.
// n <= 0 returns nil; n > len(res.Tests) returns all. res.Tests is not mutated.
func SlowestN(res RunResult, n int) []TestResult {
	if n <= 0 {
		return nil
	}
	cp := make([]TestResult, len(res.Tests))
	copy(cp, res.Tests)
	sort.Slice(cp, func(i, j int) bool {
		return cp[i].Elapsed > cp[j].Elapsed
	})
	if n > len(cp) {
		return cp
	}
	return cp[:n]
}

// ExportMarkdown writes a markdown report to path.
// Redaction is applied to every output line via red before writing;
// red == nil is a no-op (no panic). Raw output is never written.
func ExportMarkdown(path string, groups []Group, res RunResult, red *redact.Redactor) error {
	return writeAtomic(path, 0o600, func(w io.Writer) error {
		return writeMarkdown(w, groups, res, red)
	})
}

func writeMarkdown(w io.Writer, groups []Group, res RunResult, red *redact.Redactor) error {
	if res.BuildFailed || res.PreRunFailed {
		if res.BuildFailed {
			fmt.Fprintln(w, "# BUILD FAILED")
			for _, line := range redactLines(red, res.BuildOutput) {
				fmt.Fprint(w, line)
			}
		}
		if res.PreRunFailed {
			fmt.Fprintln(w, "# SUITE EXITED BEFORE RUNNING TESTS")
			for _, line := range redactLines(red, res.PreRunOutput) {
				fmt.Fprint(w, line)
			}
		}
		return nil
	}

	pass, fail, skip := groupCounts(res.Tests, nil)
	fmt.Fprintf(w, "# Test Results\n\n%d passed, %d failed, %d skipped\n\n", pass, fail, skip)

	for _, g := range groups {
		ns := make(map[string]bool, len(g.Tests))
		for _, name := range g.Tests {
			ns[name] = true
		}
		gp, gf, gs := groupCounts(res.Tests, ns)
		fmt.Fprintf(w, "## %s\n\n%d passed, %d failed, %d skipped\n\n", g.Name, gp, gf, gs)
	}

	for _, tr := range res.Tests {
		if tr.Sub != "" {
			continue
		}
		fmt.Fprintf(w, "### %s (%s)\n\n", tr.Name, tr.Status.String())
		for _, line := range redactLines(red, tr.Output) {
			fmt.Fprint(w, line)
		}
		fmt.Fprintln(w)
	}
	return nil
}

// htmlGroupRow is template data for one group row.
type htmlGroupRow struct {
	Name             string
	Pass, Fail, Skip int
}

// htmlTestRow is template data for one test row.
type htmlTestRow struct {
	Name   string
	Status string
	Output []string
}

// htmlPage is the top-level template data for ExportHTML.
type htmlPage struct {
	Title        string
	BuildFailed  bool
	PreRunFailed bool
	Normal       bool
	BuildOutput  []string
	PreRunOutput []string
	Summary      string
	Groups       []htmlGroupRow
	Tests        []htmlTestRow
}

const htmlTmpl = `<!DOCTYPE html>
<html>
<head><meta charset="utf-8"><title>Test Results</title></head>
<body>
<h1>{{.Title}}</h1>
{{if .BuildFailed}}<p class="build-fail">BUILD FAILED</p>
<pre>{{range .BuildOutput}}{{.}}{{end}}</pre>
{{end}}{{if .PreRunFailed}}<p class="pre-run-fail">SUITE EXITED BEFORE RUNNING TESTS</p>
<pre>{{range .PreRunOutput}}{{.}}{{end}}</pre>
{{end}}{{if .Normal}}<p>{{.Summary}}</p>
<table>
<tr><th>Group</th><th>Passed</th><th>Failed</th><th>Skipped</th></tr>
{{range .Groups}}<tr><td>{{.Name}}</td><td>{{.Pass}}</td><td>{{.Fail}}</td><td>{{.Skip}}</td></tr>
{{end}}</table>
{{range .Tests}}<div class="test">
<h3>{{.Name}} ({{.Status}})</h3>
<pre>{{range .Output}}{{.}}{{end}}</pre>
</div>
{{end}}{{end}}</body>
</html>
`

// ExportHTML writes an HTML report to path using html/template (auto-escaped).
// Redaction is applied to every output line via red before the template executes;
// red == nil is a no-op (no panic). Raw output is never written.
func ExportHTML(path string, groups []Group, res RunResult, red *redact.Redactor) error {
	return writeAtomic(path, 0o600, func(w io.Writer) error {
		return writeHTML(w, groups, res, red)
	})
}

func writeHTML(w io.Writer, groups []Group, res RunResult, red *redact.Redactor) error {
	tmpl, err := template.New("report").Parse(htmlTmpl)
	if err != nil {
		return err
	}

	page := htmlPage{Title: "Test Results"}
	if res.BuildFailed {
		page.Title = "BUILD FAILED"
		page.BuildFailed = true
		page.BuildOutput = redactLines(red, res.BuildOutput)
	}
	if res.PreRunFailed {
		if !res.BuildFailed {
			page.Title = "SUITE EXITED BEFORE RUNNING TESTS"
		}
		page.PreRunFailed = true
		page.PreRunOutput = redactLines(red, res.PreRunOutput)
	}
	if !res.BuildFailed && !res.PreRunFailed {
		page.Normal = true
		pass, fail, skip := groupCounts(res.Tests, nil)
		page.Summary = fmt.Sprintf("%d passed, %d failed, %d skipped", pass, fail, skip)

		for _, g := range groups {
			ns := make(map[string]bool, len(g.Tests))
			for _, name := range g.Tests {
				ns[name] = true
			}
			gp, gf, gs := groupCounts(res.Tests, ns)
			page.Groups = append(page.Groups, htmlGroupRow{
				Name: g.Name,
				Pass: gp,
				Fail: gf,
				Skip: gs,
			})
		}

		for _, tr := range res.Tests {
			if tr.Sub != "" {
				continue
			}
			page.Tests = append(page.Tests, htmlTestRow{
				Name:   tr.Name,
				Status: tr.Status.String(),
				Output: redactLines(red, tr.Output),
			})
		}
	}

	return tmpl.Execute(w, page)
}

// legacyPlanNote is printed by the state-aware report functions in place of
// planning/retry counts when state.Plan is nil - i.e. legacy, pre-state-v2
// data (see engine.Load, which never synthesizes a plan for such state).
const legacyPlanNote = "legacy state: planning counts unavailable"

// planningCounts holds the four ExecutionPlan-derived selection counts shown
// by the state-aware report functions.
type planningCounts struct {
	Selected, Eligible, Excluded, Unclassified int
}

// planningCountsFromPlan reports state's four selection counts and whether a
// plan was present at all. ok is false for legacy state (state.Plan == nil),
// in which case callers must render legacyPlanNote instead of a zeroed count.
func planningCountsFromPlan(state State) (pc planningCounts, ok bool) {
	if state.Plan == nil {
		return planningCounts{}, false
	}
	return planningCounts{
		Selected:     len(state.Plan.Selected),
		Eligible:     len(state.Plan.Eligible),
		Excluded:     len(state.Plan.Excluded),
		Unclassified: len(state.Plan.Unclassified),
	}, true
}

// planningLine formats pc for the text and Markdown state-aware reports.
func planningLine(pc planningCounts) string {
	return fmt.Sprintf("planning: %d selected, %d eligible, %d excluded, %d unclassified",
		pc.Selected, pc.Eligible, pc.Excluded, pc.Unclassified)
}

// retryCounts holds the three PersistFailure-derived retry/recovery counts
// shown by the state-aware report functions.
type retryCounts struct {
	Retried, Recovered, Remaining int
}

// retryCountsFromFailures derives retry/recovery counts from persisted
// failures. Retried counts failures retried at least once (Attempts > 1, an
// orthogonal tally). Recovered counts failures classified flake-confirmed by
// ClassifyFailure (a later retry passed). Remaining counts every other
// failure - still failing/unresolved - so Recovered+Remaining always equals
// len(failures).
func retryCountsFromFailures(failures []PersistFailure) retryCounts {
	var rc retryCounts
	for _, f := range failures {
		if f.Attempts > 1 {
			rc.Retried++
		}
		if f.Classification == ClassificationFlakeConfirmed {
			rc.Recovered++
		} else {
			rc.Remaining++
		}
	}
	return rc
}

// retryLine formats rc for the text and Markdown state-aware reports.
func retryLine(rc retryCounts) string {
	return fmt.Sprintf("retry: %d retried, %d recovered, %d remaining", rc.Retried, rc.Recovered, rc.Remaining)
}

const orphanAccountingUnavailable = "orphan accounting unavailable"

func orphanAccountingLine(state State) string {
	if state.Orphans == nil {
		return orphanAccountingUnavailable
	}
	if state.Orphans.CleanupStatus == CleanupNotApplicable {
		return "orphan accounting not applicable"
	}
	if state.Orphans.CleanupStatus == CleanupUnknown {
		return fmt.Sprintf(
			"orphan accounting: %d baseline, final/pre-existing/new counts unproven, cleanup status %s",
			len(state.Orphans.Baseline),
			state.Orphans.CleanupStatus,
		)
	}
	if state.Orphans.CleanupStatus == CleanupBaselineOnly {
		return fmt.Sprintf(
			"orphan accounting: %d baseline, pre-existing not captured, final not captured, 0 new (no run attempted; no attributed cleanup obligation), cleanup status %s",
			len(state.Orphans.Baseline),
			state.Orphans.CleanupStatus,
		)
	}
	return fmt.Sprintf(
		"orphan accounting: %d baseline, %d pre-existing, %d final, %d new, cleanup status %s",
		len(state.Orphans.Baseline),
		len(state.Orphans.PreExisting),
		len(state.Orphans.Final),
		len(state.Orphans.New),
		state.Orphans.CleanupStatus,
	)
}

// runResultFromState reconstructs a RunResult from persisted, state-safe
// data for reporting. State never stores raw command output - PersistResult
// carries no Output field, and BuildLog/PreRunLog are paths to already
// redacted log files written by run/retry (see cli.go's runWithSpec*),  not
// raw text. BuildOutput/PreRunOutput here are therefore single "see <path>"
// pointer lines, matching exactly what runReport built by hand before this
// helper existed.
func runResultFromState(state State) RunResult {
	var res RunResult
	res.BuildFailed = state.BuildFailed
	res.PreRunFailed = state.PreRunFailed
	if state.BuildLog != "" {
		res.BuildOutput = []string{"see " + state.BuildLog + "\n"}
	}
	if state.PreRunLog != "" {
		res.PreRunOutput = []string{"see " + state.PreRunLog + "\n"}
	}
	for _, r := range state.Results {
		status, _ := provider.ParseStatus(r.Status)
		res.Tests = append(res.Tests, TestResult{
			Package: r.Package,
			Name:    r.Test,
			Sub:     r.Sub,
			Status:  status,
			Elapsed: r.Elapsed,
		})
	}
	return res
}

// TerminalStateSummary writes a state-aware, human-readable run summary to w.
//
// It first renders exactly what TerminalSummary renders from the state's
// persisted results, including TerminalSummary's own unchanged build-fail/
// pre-run-fail trap: if either fired, TerminalStateSummary returns
// immediately afterward, same as TerminalSummary, printing neither a success
// line nor a planning/retry section.
//
// Otherwise, it appends one line of ExecutionPlan-derived selection counts
// (selected, eligible, excluded, unclassified) and one line of
// PersistFailure-derived retry counts (retried, recovered, remaining). A nil
// state.Plan (legacy, pre-state-v2 data) never synthesizes planning counts:
// a single legacyPlanNote line replaces the planning/retry section entirely.
// The final line reports persisted orphan counts, or says orphan accounting is
// unavailable when reading state without the optional orphan fields.
func TerminalStateSummary(w io.Writer, groups []Group, state State, red *redact.Redactor) {
	res := runResultFromState(state)
	TerminalSummary(w, groups, res, red)
	if res.BuildFailed || res.PreRunFailed {
		return
	}
	pc, ok := planningCountsFromPlan(state)
	if !ok {
		fmt.Fprintln(w, legacyPlanNote)
	} else {
		fmt.Fprintln(w, planningLine(pc))
		fmt.Fprintln(w, retryLine(retryCountsFromFailures(state.Failures)))
	}
	fmt.Fprintln(w, orphanAccountingLine(state))
}

// ExportMarkdownState writes a state-aware Markdown report to path.
// It first writes exactly what ExportMarkdown writes - including
// ExportMarkdown's own unchanged build-fail/pre-run-fail trap and redaction -
// then, only in the non-trap case, appends a "Planning and retry" section
// (or, for legacy state.Plan == nil data, a single legacyPlanNote line)
// followed by persisted orphan counts or the unavailable note. See
// TerminalStateSummary for the exact section/trap/legacy rules this mirrors.
func ExportMarkdownState(path string, groups []Group, state State, red *redact.Redactor) error {
	res := runResultFromState(state)
	if err := ExportMarkdown(path, groups, res, red); err != nil {
		return err
	}
	if res.BuildFailed || res.PreRunFailed {
		return nil
	}
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}
	defer f.Close()

	pc, ok := planningCountsFromPlan(state)
	if !ok {
		fmt.Fprintf(f, "\n%s\n", legacyPlanNote)
	} else {
		fmt.Fprintf(f, "\n## Planning and retry\n\n%s\n\n%s\n", planningLine(pc), retryLine(retryCountsFromFailures(state.Failures)))
	}
	fmt.Fprintf(f, "\n## Orphan accounting\n\n%s\n", orphanAccountingLine(state))
	return nil
}

// htmlStatePage is the template data for ExportHTMLState. It embeds htmlPage
// so every field TerminalSummary/ExportHTML's shared template rules use
// (Title, BuildFailed, Summary, Groups, Tests, ...) is available unchanged,
// plus the planning/retry and orphan addendum fields rendered only in the
// non-trap case.
type htmlStatePage struct {
	htmlPage
	HasPlan      bool
	PlanningLine string
	RetryLine    string
	LegacyNote   string
	OrphanLine   string
}

// htmlStateTmpl is ExportHTMLState's own template - a copy of htmlTmpl with
// planning/retry and orphan-accounting addenda. htmlTmpl itself is untouched;
// ExportHTML keeps using it as-is.
const htmlStateTmpl = `<!DOCTYPE html>
<html>
<head><meta charset="utf-8"><title>Test Results</title></head>
<body>
<h1>{{.Title}}</h1>
{{if .BuildFailed}}<p class="build-fail">BUILD FAILED</p>
<pre>{{range .BuildOutput}}{{.}}{{end}}</pre>
{{end}}{{if .PreRunFailed}}<p class="pre-run-fail">SUITE EXITED BEFORE RUNNING TESTS</p>
<pre>{{range .PreRunOutput}}{{.}}{{end}}</pre>
{{end}}{{if .Normal}}<p>{{.Summary}}</p>
{{if .HasPlan}}<p class="planning">{{.PlanningLine}}</p>
<p class="retry">{{.RetryLine}}</p>
{{else}}<p class="legacy-note">{{.LegacyNote}}</p>
{{end}}<p class="orphan-accounting">{{.OrphanLine}}</p>
<table>
<tr><th>Group</th><th>Passed</th><th>Failed</th><th>Skipped</th></tr>
{{range .Groups}}<tr><td>{{.Name}}</td><td>{{.Pass}}</td><td>{{.Fail}}</td><td>{{.Skip}}</td></tr>
{{end}}</table>
{{range .Tests}}<div class="test">
<h3>{{.Name}} ({{.Status}})</h3>
<pre>{{range .Output}}{{.}}{{end}}</pre>
</div>
{{end}}{{end}}</body>
</html>
`

// ExportHTMLState writes a state-aware HTML report to path using its own
// template and page struct (htmlStateTmpl/htmlStatePage, distinct from
// ExportHTML's htmlTmpl/htmlPage), auto-escaped via html/template exactly
// like ExportHTML. It reproduces ExportHTML's exact build-fail/pre-run-fail
// rendering and redaction, then, only in the non-trap case, adds a
// planning/retry section (or, for legacy state.Plan == nil data, a single
// legacyPlanNote line) and orphan accounting.
func ExportHTMLState(path string, groups []Group, state State, red *redact.Redactor) error {
	tmpl, err := template.New("state-report").Parse(htmlStateTmpl)
	if err != nil {
		return err
	}

	res := runResultFromState(state)
	page := htmlStatePage{
		htmlPage:   htmlPage{Title: "Test Results"},
		OrphanLine: orphanAccountingLine(state),
	}
	if res.BuildFailed {
		page.Title = "BUILD FAILED"
		page.BuildFailed = true
		page.BuildOutput = redactLines(red, res.BuildOutput)
	}
	if res.PreRunFailed {
		if !res.BuildFailed {
			page.Title = "SUITE EXITED BEFORE RUNNING TESTS"
		}
		page.PreRunFailed = true
		page.PreRunOutput = redactLines(red, res.PreRunOutput)
	}
	if !res.BuildFailed && !res.PreRunFailed {
		page.Normal = true
		pass, fail, skip := groupCounts(res.Tests, nil)
		page.Summary = fmt.Sprintf("%d passed, %d failed, %d skipped", pass, fail, skip)

		for _, g := range groups {
			ns := make(map[string]bool, len(g.Tests))
			for _, name := range g.Tests {
				ns[name] = true
			}
			gp, gf, gs := groupCounts(res.Tests, ns)
			page.Groups = append(page.Groups, htmlGroupRow{
				Name: g.Name,
				Pass: gp,
				Fail: gf,
				Skip: gs,
			})
		}

		for _, tr := range res.Tests {
			if tr.Sub != "" {
				continue
			}
			page.Tests = append(page.Tests, htmlTestRow{
				Name:   tr.Name,
				Status: tr.Status.String(),
				Output: redactLines(red, tr.Output),
			})
		}

		if pc, ok := planningCountsFromPlan(state); ok {
			page.HasPlan = true
			page.PlanningLine = planningLine(pc)
			page.RetryLine = retryLine(retryCountsFromFailures(state.Failures))
		} else {
			page.LegacyNote = legacyPlanNote
		}
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	return tmpl.Execute(f, page)
}
