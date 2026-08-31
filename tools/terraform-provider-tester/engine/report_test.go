package engine

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/github/terraform-provider-tester/internal/redact"
	"github.com/github/terraform-provider-tester/provider"
)

func TestTerminalSummaryCounts(t *testing.T) {
	groups := []Group{
		{Name: "group1", Tests: []string{"TestA", "TestB"}},
		{Name: "group2", Tests: []string{"TestC", "TestD"}},
	}
	res := RunResult{
		Tests: []TestResult{
			{Name: "TestA", Sub: "", Status: provider.StatusPass},
			{Name: "TestB", Sub: "", Status: provider.StatusPass},
			{Name: "TestC", Sub: "", Status: provider.StatusFail},
			{Name: "TestD", Sub: "", Status: provider.StatusSkip},
		},
	}
	var buf bytes.Buffer
	TerminalSummary(&buf, groups, res)
	out := buf.String()

	for _, want := range []string{"2 passed", "1 failed", "1 skipped"} {
		if !strings.Contains(out, want) {
			t.Errorf("output missing %q; got:\n%s", want, out)
		}
	}
	for _, g := range groups {
		if !strings.Contains(out, g.Name) {
			t.Errorf("output missing group name %q; got:\n%s", g.Name, out)
		}
	}
}

func TestTerminalSummaryBuildFail(t *testing.T) {
	res := RunResult{BuildFailed: true}
	var buf bytes.Buffer
	TerminalSummary(&buf, nil, res)
	out := buf.String()

	if !strings.Contains(out, "BUILD FAILED") {
		t.Errorf("output missing BUILD FAILED; got:\n%s", out)
	}
	if strings.Contains(out, "passed") {
		t.Errorf("output must not contain success phrase 'passed'; got:\n%s", out)
	}
}

func TestTerminalSummaryPreRunFail(t *testing.T) {
	res := RunResult{
		PreRunFailed: true,
		PreRunOutput: []string{"missing GITHUB_OWNER"},
	}
	var buf bytes.Buffer
	TerminalSummary(&buf, nil, res)
	out := buf.String()

	for _, want := range []string{"SUITE EXITED BEFORE RUNNING TESTS", "missing GITHUB_OWNER"} {
		if !strings.Contains(out, want) {
			t.Errorf("output missing %q; got:\n%s", want, out)
		}
	}
	if strings.Contains(out, "BUILD FAILED") {
		t.Errorf("output must not contain BUILD FAILED; got:\n%s", out)
	}
	if strings.Contains(out, "passed") {
		t.Errorf("output must not contain success phrase 'passed'; got:\n%s", out)
	}
}

// TestSummaryBothFailFlags covers the multi-package case where one package
// fails to build and another exits before running tests. Both flags end up set
// on the run-level RunResult; neither condition may be hidden, and a success
// line must never appear, across all three report faces.
func TestSummaryBothFailFlags(t *testing.T) {
	res := RunResult{
		BuildFailed:  true,
		BuildOutput:  []string{"./github/x_test.go:1: undefined: Foo\n"},
		PreRunFailed: true,
		PreRunOutput: []string{"missing GITHUB_OWNER"},
	}

	var buf bytes.Buffer
	TerminalSummary(&buf, nil, res)
	term := buf.String()
	for _, want := range []string{"BUILD FAILED", "SUITE EXITED BEFORE RUNNING TESTS", "missing GITHUB_OWNER"} {
		if !strings.Contains(term, want) {
			t.Errorf("TerminalSummary missing %q; got:\n%s", want, term)
		}
	}
	if strings.Contains(term, "passed") {
		t.Errorf("TerminalSummary must not contain success phrase 'passed'; got:\n%s", term)
	}

	dir := t.TempDir()
	mdPath := filepath.Join(dir, "report.md")
	if err := ExportMarkdown(mdPath, nil, res, nil); err != nil {
		t.Fatalf("ExportMarkdown: %v", err)
	}
	mdData, err := os.ReadFile(mdPath)
	if err != nil {
		t.Fatalf("ReadFile: %v", err)
	}
	md := string(mdData)
	for _, want := range []string{"BUILD FAILED", "SUITE EXITED BEFORE RUNNING TESTS"} {
		if !strings.Contains(md, want) {
			t.Errorf("markdown missing %q; got:\n%s", want, md)
		}
	}

	htmlPath := filepath.Join(dir, "report.html")
	if err := ExportHTML(htmlPath, nil, res, nil); err != nil {
		t.Fatalf("ExportHTML: %v", err)
	}
	htmlData, err := os.ReadFile(htmlPath)
	if err != nil {
		t.Fatalf("ReadFile: %v", err)
	}
	html := string(htmlData)
	for _, want := range []string{"BUILD FAILED", "SUITE EXITED BEFORE RUNNING TESTS"} {
		if !strings.Contains(html, want) {
			t.Errorf("html missing %q; got:\n%s", want, html)
		}
	}
	if strings.Contains(html, "{{") {
		t.Errorf("html contains unrendered template directive; got:\n%s", html)
	}
}

func TestSlowestN(t *testing.T) {
	res := RunResult{
		Tests: []TestResult{
			{Name: "T1", Elapsed: 0.1},
			{Name: "T2", Elapsed: 5.0},
			{Name: "T3", Elapsed: 2.0},
			{Name: "T4", Elapsed: 0.3},
		},
	}

	top2 := SlowestN(res, 2)
	if len(top2) != 2 {
		t.Fatalf("SlowestN(2): want 2, got %d", len(top2))
	}
	if top2[0].Elapsed != 5.0 || top2[1].Elapsed != 2.0 {
		t.Errorf("SlowestN(2): want [5.0, 2.0], got [%v, %v]", top2[0].Elapsed, top2[1].Elapsed)
	}

	zero := SlowestN(res, 0)
	if len(zero) != 0 {
		t.Errorf("SlowestN(0): want empty, got %d", len(zero))
	}

	all := SlowestN(res, 99)
	if len(all) != 4 {
		t.Errorf("SlowestN(99): want 4, got %d", len(all))
	}
}

func TestExportMarkdownRedacts(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "out.md")

	res := RunResult{
		Tests: []TestResult{
			{
				Name:   "TestSecret",
				Sub:    "",
				Status: provider.StatusPass,
				Output: []string{"token=ghp_SUPERSECRET\n"},
			},
		},
	}
	red := redact.New([]string{"ghp_SUPERSECRET"})
	if err := ExportMarkdown(path, nil, res, red); err != nil {
		t.Fatalf("ExportMarkdown: %v", err)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("ReadFile: %v", err)
	}
	content := string(data)
	info, err := os.Stat(path)
	if err != nil {
		t.Fatalf("Stat: %v", err)
	}
	if got := info.Mode().Perm(); got != 0o600 {
		t.Fatalf("mode = %v, want 0600", got)
	}
	matches, err := filepath.Glob(filepath.Join(dir, ".*.tmp"))
	if err != nil {
		t.Fatalf("Glob: %v", err)
	}
	if len(matches) != 0 {
		t.Fatalf("leftover temp files: %v", matches)
	}

	if strings.Contains(content, "ghp_SUPERSECRET") {
		t.Errorf("ExportMarkdown: secret not redacted; file contains ghp_SUPERSECRET")
	}
	if !strings.Contains(content, "***REDACTED***") {
		t.Errorf("ExportMarkdown: expected ***REDACTED*** in output; got:\n%s", content)
	}
}

func TestWriteMarkdownHelper(t *testing.T) {
	var buf bytes.Buffer
	groups := []Group{{Name: "group", Tests: []string{"TestA"}}}
	res := RunResult{
		Tests: []TestResult{{Name: "TestA", Sub: "", Status: provider.StatusPass}},
	}

	if err := writeMarkdown(&buf, groups, res, nil); err != nil {
		t.Fatalf("writeMarkdown: %v", err)
	}
	out := buf.String()
	for _, want := range []string{"# Test Results", "group", "TestA", "1 passed"} {
		if !strings.Contains(out, want) {
			t.Fatalf("writeMarkdown output missing %q; got:\n%s", want, out)
		}
	}
}

// TestStateSummaryCoversPlanningAndRetryCounts verifies that, given a state
// with a plan and a mix of retried/recovered/remaining failures,
// TerminalStateSummary, ExportMarkdownState, and ExportHTMLState each report
// the existing pass/fail/skip totals plus the exact planning counts
// (selected, eligible, excluded, unclassified) and retry counts (retried,
// recovered, remaining) derived from the state.
func TestStateSummaryCoversPlanningAndRetryCounts(t *testing.T) {
	groups := []Group{
		{Name: "misc", Tests: []string{"TestAccKeep", "TestAccFlaky", "TestAccBroken"}},
	}
	state := State{
		Results: []PersistResult{
			{Test: "TestAccKeep", Status: "pass"},
			{Test: "TestAccFlaky", Status: "pass"},
			{Test: "TestAccBroken", Status: "fail"},
		},
		Failures: []PersistFailure{
			{Test: "TestAccFlaky", Classification: ClassificationFlakeConfirmed, Attempts: 2},
			{Test: "TestAccBroken", Classification: ClassificationReal, Attempts: 2},
			{Test: "TestAccOther", Classification: ClassificationReal, Attempts: 1},
		},
		Plan: &ExecutionPlan{
			Mode:         "individual",
			Selected:     []string{"TestAccKeep", "TestAccFlaky", "TestAccBroken", "TestAccExcludedOne", "TestAccMystery"},
			Eligible:     []string{"TestAccKeep", "TestAccFlaky", "TestAccBroken"},
			Excluded:     []PlanExclusion{{Test: "TestAccExcludedOne", Code: "excluded/mode-incompatible"}},
			Unclassified: []string{"TestAccMystery"},
		},
	}

	wantSubstrings := []string{
		"2 passed", "1 failed", "0 skipped",
		"5 selected", "3 eligible", "1 excluded", "1 unclassified",
		"2 retried", "1 recovered", "2 remaining",
	}

	var buf bytes.Buffer
	TerminalStateSummary(&buf, groups, state, nil)
	term := buf.String()
	for _, want := range wantSubstrings {
		if !strings.Contains(term, want) {
			t.Errorf("TerminalStateSummary missing %q; got:\n%s", want, term)
		}
	}

	dir := t.TempDir()
	mdPath := filepath.Join(dir, "state.md")
	if err := ExportMarkdownState(mdPath, groups, state, nil); err != nil {
		t.Fatalf("ExportMarkdownState: %v", err)
	}
	mdData, err := os.ReadFile(mdPath)
	if err != nil {
		t.Fatalf("ReadFile: %v", err)
	}
	md := string(mdData)
	for _, want := range wantSubstrings {
		if !strings.Contains(md, want) {
			t.Errorf("ExportMarkdownState missing %q; got:\n%s", want, md)
		}
	}

	htmlPath := filepath.Join(dir, "state.html")
	if err := ExportHTMLState(htmlPath, groups, state, nil); err != nil {
		t.Fatalf("ExportHTMLState: %v", err)
	}
	htmlData, err := os.ReadFile(htmlPath)
	if err != nil {
		t.Fatalf("ReadFile: %v", err)
	}
	html := string(htmlData)
	for _, want := range wantSubstrings {
		if !strings.Contains(html, want) {
			t.Errorf("ExportHTMLState missing %q; got:\n%s", want, html)
		}
	}
	if strings.Contains(html, "{{") {
		t.Errorf("ExportHTMLState: output contains unexpanded template directive {{")
	}
}

// TestStateSummaryLegacyPlanUnavailable verifies that a legacy state (no
// persisted plan - e.g. pre-state-v2 data, see engine.Load) still renders the
// existing pass/fail/skip totals, replacing the planning/retry section with a
// single "legacy state: planning counts unavailable" note instead of
// synthesizing zeroed planning counts.
func TestStateSummaryLegacyPlanUnavailable(t *testing.T) {
	state := State{
		Results: []PersistResult{
			{Test: "TestAccKeep", Status: "pass"},
		},
	}
	const legacyNote = "legacy state: planning counts unavailable"

	var buf bytes.Buffer
	TerminalStateSummary(&buf, nil, state, nil)
	term := buf.String()
	if !strings.Contains(term, "1 passed") {
		t.Errorf("TerminalStateSummary missing existing totals %q; got:\n%s", "1 passed", term)
	}
	if !strings.Contains(term, legacyNote) {
		t.Errorf("TerminalStateSummary missing %q; got:\n%s", legacyNote, term)
	}
	if strings.Contains(term, "selected") {
		t.Errorf("TerminalStateSummary must not synthesize planning counts for legacy state; got:\n%s", term)
	}

	dir := t.TempDir()
	mdPath := filepath.Join(dir, "legacy.md")
	if err := ExportMarkdownState(mdPath, nil, state, nil); err != nil {
		t.Fatalf("ExportMarkdownState: %v", err)
	}
	mdData, err := os.ReadFile(mdPath)
	if err != nil {
		t.Fatalf("ReadFile: %v", err)
	}
	md := string(mdData)
	if !strings.Contains(md, "1 passed") || !strings.Contains(md, legacyNote) {
		t.Errorf("ExportMarkdownState missing existing totals or legacy note; got:\n%s", md)
	}

	htmlPath := filepath.Join(dir, "legacy.html")
	if err := ExportHTMLState(htmlPath, nil, state, nil); err != nil {
		t.Fatalf("ExportHTMLState: %v", err)
	}
	htmlData, err := os.ReadFile(htmlPath)
	if err != nil {
		t.Fatalf("ReadFile: %v", err)
	}
	html := string(htmlData)
	if !strings.Contains(html, "1 passed") || !strings.Contains(html, legacyNote) {
		t.Errorf("ExportHTMLState missing existing totals or legacy note; got:\n%s", html)
	}
	if strings.Contains(html, "{{") {
		t.Errorf("ExportHTMLState: output contains unexpanded template directive {{")
	}
}

// TestStateSummaryBothFailFlags verifies that the state-aware report
// functions preserve the exact build-fail/pre-run-fail trap: when either flag
// is set, only that trap's section (plus, unlike the RunResult-based
// functions, a "see <path>" pointer to the redacted on-disk log recorded in
// State.BuildLog/PreRunLog) is ever printed - never a success line, and never
// the planning/retry section or the legacy-state note.
func TestStateSummaryBothFailFlags(t *testing.T) {
	state := State{
		BuildFailed:  true,
		BuildLog:     "/tmp/fake-failures/build.log",
		PreRunFailed: true,
		PreRunLog:    "/tmp/fake-failures/pre-run.log",
		Plan: &ExecutionPlan{
			Selected: []string{"TestAccA"},
			Eligible: []string{"TestAccA"},
		},
		Orphans: &OrphanAccounting{
			Mode:          "organization",
			CleanupStatus: CleanupUnknown,
		},
	}

	var buf bytes.Buffer
	TerminalStateSummary(&buf, nil, state, nil)
	term := buf.String()
	for _, want := range []string{
		"BUILD FAILED", "SUITE EXITED BEFORE RUNNING TESTS",
		"see /tmp/fake-failures/build.log", "see /tmp/fake-failures/pre-run.log",
	} {
		if !strings.Contains(term, want) {
			t.Errorf("TerminalStateSummary missing %q; got:\n%s", want, term)
		}
	}
	for _, mustNot := range []string{"passed", "selected", "legacy state", "orphan accounting"} {
		if strings.Contains(term, mustNot) {
			t.Errorf("TerminalStateSummary must not contain %q while build/pre-run failed; got:\n%s", mustNot, term)
		}
	}

	dir := t.TempDir()
	mdPath := filepath.Join(dir, "fail.md")
	if err := ExportMarkdownState(mdPath, nil, state, nil); err != nil {
		t.Fatalf("ExportMarkdownState: %v", err)
	}
	mdData, err := os.ReadFile(mdPath)
	if err != nil {
		t.Fatalf("ReadFile: %v", err)
	}
	md := string(mdData)
	for _, want := range []string{"BUILD FAILED", "SUITE EXITED BEFORE RUNNING TESTS"} {
		if !strings.Contains(md, want) {
			t.Errorf("ExportMarkdownState missing %q; got:\n%s", want, md)
		}
	}
	for _, mustNot := range []string{"selected", "legacy state", "orphan accounting"} {
		if strings.Contains(md, mustNot) {
			t.Errorf("ExportMarkdownState must not contain %q while build/pre-run failed; got:\n%s", mustNot, md)
		}
	}

	htmlPath := filepath.Join(dir, "fail.html")
	if err := ExportHTMLState(htmlPath, nil, state, nil); err != nil {
		t.Fatalf("ExportHTMLState: %v", err)
	}
	htmlData, err := os.ReadFile(htmlPath)
	if err != nil {
		t.Fatalf("ReadFile: %v", err)
	}
	html := string(htmlData)
	for _, want := range []string{"BUILD FAILED", "SUITE EXITED BEFORE RUNNING TESTS"} {
		if !strings.Contains(html, want) {
			t.Errorf("ExportHTMLState missing %q; got:\n%s", want, html)
		}
	}
	for _, mustNot := range []string{"selected", "legacy state", "orphan accounting"} {
		if strings.Contains(html, mustNot) {
			t.Errorf("ExportHTMLState must not contain %q while build/pre-run failed; got:\n%s", mustNot, html)
		}
	}
	if strings.Contains(html, "{{") {
		t.Errorf("ExportHTMLState: output contains unexpanded template directive {{")
	}
}

func stateReportOutputs(t *testing.T, state State) map[string]string {
	t.Helper()
	var terminal bytes.Buffer
	TerminalStateSummary(&terminal, nil, state, nil)

	dir := t.TempDir()
	markdownPath := filepath.Join(dir, "state.md")
	if err := ExportMarkdownState(markdownPath, nil, state, nil); err != nil {
		t.Fatalf("ExportMarkdownState: %v", err)
	}
	markdown, err := os.ReadFile(markdownPath)
	if err != nil {
		t.Fatal(err)
	}

	htmlPath := filepath.Join(dir, "state.html")
	if err := ExportHTMLState(htmlPath, nil, state, nil); err != nil {
		t.Fatalf("ExportHTMLState: %v", err)
	}
	html, err := os.ReadFile(htmlPath)
	if err != nil {
		t.Fatal(err)
	}

	return map[string]string{
		"terminal": terminal.String(),
		"markdown": string(markdown),
		"html":     string(html),
	}
}

func TestStateSummaryAndExportStateOrphanAccounting(t *testing.T) {
	state := State{
		Results: []PersistResult{{Test: "TestAccThing", Status: "pass"}},
		Plan:    &ExecutionPlan{Mode: "organization", Selected: []string{"TestAccThing"}, Eligible: []string{"TestAccThing"}},
		Orphans: &OrphanAccounting{
			Mode: "organization",
			Baseline: []provider.Resource{
				{Kind: "repository", Name: "tf-acc-test-existing-one"},
				{Kind: "repository", Name: "tf-acc-test-existing-two"},
			},
			Final: []provider.Resource{
				{Kind: "repository", Name: "tf-acc-test-existing-one"},
				{Kind: "issue_label", Name: "tf-acc-test-new-one"},
				{Kind: "issue_label", Name: "tf-acc-test-new-two"},
			},
			PreExisting:   []provider.Resource{{Kind: "repository", Name: "tf-acc-test-existing-one"}},
			New:           []provider.Resource{{Kind: "issue_label", Name: "tf-acc-test-new-one"}, {Kind: "issue_label", Name: "tf-acc-test-new-two"}},
			CleanupStatus: CleanupComplete,
		},
	}

	for format, output := range stateReportOutputs(t, state) {
		for _, want := range []string{
			"orphan accounting",
			"2 baseline",
			"1 pre-existing",
			"3 final",
			"2 new",
			"cleanup status complete",
		} {
			if !strings.Contains(output, want) {
				t.Errorf("%s report missing %q:\n%s", format, want, output)
			}
		}
	}
}

func TestStateSummaryAndExportStateOrphanUnavailable(t *testing.T) {
	state := State{
		Results: []PersistResult{{Test: "TestAccLegacy", Status: "pass"}},
	}
	for format, output := range stateReportOutputs(t, state) {
		if !strings.Contains(output, "orphan accounting unavailable") {
			t.Errorf("%s legacy report missing orphan-unavailable note:\n%s", format, output)
		}
	}
}

func TestStateSummaryAndExportStateOrphanAccountingNotApplicable(t *testing.T) {
	state := State{
		Results: []PersistResult{{Test: "TestAccPublic", Status: "pass"}},
		Plan:    &ExecutionPlan{Mode: "anonymous", Selected: []string{"TestAccPublic"}, Eligible: []string{"TestAccPublic"}},
		Orphans: &OrphanAccounting{
			Mode:          "anonymous",
			CleanupStatus: CleanupNotApplicable,
		},
	}

	for format, output := range stateReportOutputs(t, state) {
		if !strings.Contains(output, "orphan accounting not applicable") {
			t.Errorf("%s not-applicable report missing truthful accounting label:\n%s", format, output)
		}
		for _, unmeasured := range []string{"0 baseline", "0 pre-existing", "0 final", "0 new"} {
			if strings.Contains(output, unmeasured) {
				t.Errorf("%s not-applicable report presents unmeasured count as proven zero (%q):\n%s", format, unmeasured, output)
			}
		}
	}
}

func TestStateSummaryAndExportStateUnknownOrphanDeltaIsUnproven(t *testing.T) {
	state := State{
		Results: []PersistResult{{Test: "TestAccThing", Status: "pass"}},
		Plan:    &ExecutionPlan{Mode: "organization", Selected: []string{"TestAccThing"}, Eligible: []string{"TestAccThing"}},
		Orphans: &OrphanAccounting{
			Mode:             "organization",
			Baseline:         []provider.Resource{{Kind: "repository", Name: "tf-acc-test-existing"}},
			BaselineCaptured: true,
			CleanupStatus:    CleanupUnknown,
		},
	}

	for format, output := range stateReportOutputs(t, state) {
		for _, want := range []string{
			"1 baseline",
			"final/pre-existing/new counts unproven",
			"cleanup status unknown",
		} {
			if !strings.Contains(output, want) {
				t.Errorf("%s unknown report missing %q:\n%s", format, want, output)
			}
		}
		if strings.Contains(output, "0 new") {
			t.Errorf("%s unknown report presents New as proven zero:\n%s", format, output)
		}
	}
}

func TestStateSummaryAndExportStateBaselineOnlyOrphanAccountingDoesNotProveUnmeasuredCounts(t *testing.T) {
	state := State{
		Plan: &ExecutionPlan{Mode: "organization"},
		Orphans: &OrphanAccounting{
			Mode: "organization",
			Baseline: []provider.Resource{
				{Kind: "repository", Name: "tf-acc-test-existing-one"},
				{Kind: "repository", Name: "tf-acc-test-existing-two"},
			},
			BaselineCaptured: true,
			CleanupStatus:    CleanupBaselineOnly,
		},
	}

	for format, output := range stateReportOutputs(t, state) {
		for _, want := range []string{
			"2 baseline",
			"pre-existing not captured",
			"final not captured",
			"0 new (no run attempted; no attributed cleanup obligation)",
			"cleanup status baseline-only",
		} {
			if !strings.Contains(output, want) {
				t.Errorf("%s baseline-only report missing %q:\n%s", format, want, output)
			}
		}
		for _, misleading := range []string{"0 pre-existing", "0 final"} {
			if strings.Contains(output, misleading) {
				t.Errorf("%s baseline-only report presents an unmeasured count as proven zero (%q):\n%s", format, misleading, output)
			}
		}
	}
}

func TestExportHTMLWellFormed(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "out.html")

	groups := []Group{
		{Name: "mygroup", Tests: []string{"TestHTMLSecret"}},
	}
	res := RunResult{
		Tests: []TestResult{
			{
				Name:   "TestHTMLSecret",
				Sub:    "",
				Status: provider.StatusPass,
				Output: []string{"key=html_SECRET_VALUE\n"},
			},
		},
	}
	red := redact.New([]string{"html_SECRET_VALUE"})
	if err := ExportHTML(path, groups, res, red); err != nil {
		t.Fatalf("ExportHTML: %v", err)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("ReadFile: %v", err)
	}
	content := string(data)
	info, err := os.Stat(path)
	if err != nil {
		t.Fatalf("Stat: %v", err)
	}
	if got := info.Mode().Perm(); got != 0o600 {
		t.Fatalf("mode = %v, want 0600", got)
	}
	matches, err := filepath.Glob(filepath.Join(dir, ".*.tmp"))
	if err != nil {
		t.Fatalf("Glob: %v", err)
	}
	if len(matches) != 0 {
		t.Fatalf("leftover temp files: %v", matches)
	}

	if strings.Contains(content, "{{") {
		t.Errorf("ExportHTML: output contains unexpanded template directive {{")
	}
	if !strings.Contains(content, "mygroup") {
		t.Errorf("ExportHTML: output missing group name 'mygroup'")
	}
	if !strings.Contains(content, "TestHTMLSecret") {
		t.Errorf("ExportHTML: output missing test name 'TestHTMLSecret'")
	}
	if strings.Contains(content, "html_SECRET_VALUE") {
		t.Errorf("ExportHTML: secret not redacted; file contains html_SECRET_VALUE")
	}
}

func TestWriteHTMLHelper(t *testing.T) {
	var buf bytes.Buffer
	groups := []Group{{Name: "group", Tests: []string{"TestA"}}}
	res := RunResult{
		Tests: []TestResult{{Name: "TestA", Sub: "", Status: provider.StatusPass}},
	}

	if err := writeHTML(&buf, groups, res, nil); err != nil {
		t.Fatalf("writeHTML: %v", err)
	}
	out := buf.String()
	for _, want := range []string{"Test Results", "group", "TestA", "passed"} {
		if !strings.Contains(out, want) {
			t.Fatalf("writeHTML output missing %q; got:\n%s", want, out)
		}
	}
	if strings.Contains(out, "{{") {
		t.Fatalf("writeHTML output contains template directive; got:\n%s", out)
	}
}
