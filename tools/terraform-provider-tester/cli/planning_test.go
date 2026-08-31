package cli

import (
	"bytes"
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/github/terraform-provider-tester/engine"
	"github.com/github/terraform-provider-tester/fakeprovider"
	"github.com/github/terraform-provider-tester/internal/redact"
	"github.com/github/terraform-provider-tester/provider"
)

// testModes is a shared Modes() fixture for Fakes used by planning-integrated
// run/preflight tests: both anonymous and organization are valid so tests can
// pick either mode without a mode-unknown planning error.
var testModes = []provider.Mode{{Name: "anonymous"}, {Name: "organization"}}

// allowAllRequirements is a shared RequirementsFn fixture that classifies any
// test name with no scope/capability/side-effect/mode requirements. An empty
// Modes list is unrestricted (see engine.modeCompatible), so every test using
// it is eligible in any mode. Fakes must set a RequirementsFn (even this
// permissive one) because BuildExecutionPlan drops any actually-unclassified
// test out of Eligible, so a bare zero-value Fake would silently run nothing.
func allowAllRequirements(string) (provider.TestRequirements, bool) {
	return provider.TestRequirements{}, true
}

// planningTestProvider returns a Fake configured with distinct per-test
// requirements so planForCommand's aggregation (Scopes/Capabilities/
// SideEffects/Eligible/Excluded) can be asserted precisely.
func planningTestProvider() *fakeprovider.Fake {
	return &fakeprovider.Fake{
		NameVal:  "test",
		Packages: []string{"./..."},
		Pattern:  "^TestAcc",
		ModesVal: testModes,
		RequirementsFn: func(name string) (provider.TestRequirements, bool) {
			switch name {
			case "TestAccRepo":
				return provider.TestRequirements{Scopes: []string{"repo"}}, true
			case "TestAccOrgOnly":
				return provider.TestRequirements{
					Modes:        []string{"organization"},
					Scopes:       []string{"read:org"},
					Capabilities: []string{"organization"},
				}, true
			default:
				return provider.TestRequirements{}, true
			}
		},
	}
}

func TestPlanForCommandBuildsEligiblePlanFromDiscoveredNames(t *testing.T) {
	root := t.TempDir()
	d := deps{
		provider: planningTestProvider(),
		list:     stubList([]string{"TestAccRepo", "TestAccOrgOnly"}),
	}

	plan, _, err := planForCommand(context.Background(), d, root, planningOptions{Mode: "anonymous"})
	if err != nil {
		t.Fatalf("planForCommand: unexpected error: %v", err)
	}
	if got := plan.Selected; len(got) != 2 {
		t.Fatalf("Selected = %v, want 2 entries", got)
	}
	if got := plan.Eligible; len(got) != 1 || got[0] != "TestAccRepo" {
		t.Fatalf("Eligible = %v, want [TestAccRepo] (TestAccOrgOnly requires organization mode)", got)
	}
	if len(plan.Excluded) != 1 || plan.Excluded[0].Test != "TestAccOrgOnly" {
		t.Fatalf("Excluded = %+v, want TestAccOrgOnly excluded for mode incompatibility", plan.Excluded)
	}
	if got := plan.Scopes; len(got) != 1 || got[0] != "repo" {
		t.Fatalf("Scopes = %v, want [repo] (only TestAccRepo is eligible and aggregated)", got)
	}
}

func TestPlanForCommandReturnsGroupsFromDiscoveredNames(t *testing.T) {
	root := t.TempDir()
	pf := planningTestProvider()
	pf.GroupFunc = func(name string) string {
		if name == "TestAccRepo" {
			return "repos"
		}
		return "misc"
	}
	d := deps{
		provider: pf,
		list:     stubList([]string{"TestAccRepo", "TestAccOrgOnly"}),
	}

	_, groups, err := planForCommand(context.Background(), d, root, planningOptions{Mode: "anonymous"})
	if err != nil {
		t.Fatalf("planForCommand: unexpected error: %v", err)
	}
	names := make(map[string][]string, len(groups))
	for _, g := range groups {
		names[g.Name] = g.Tests
	}
	if got := names["repos"]; len(got) != 1 || got[0] != "TestAccRepo" {
		t.Fatalf("group 'repos' = %v, want [TestAccRepo]", got)
	}
}

func TestPlanForCommandPropagatesDiscoveryError(t *testing.T) {
	root := t.TempDir()
	d := deps{
		provider: planningTestProvider(),
		list:     failList(errors.New("build failed: does not compile")),
	}

	plan, groups, err := planForCommand(context.Background(), d, root, planningOptions{Mode: "anonymous"})
	if err == nil {
		t.Fatal("expected discovery error, got nil")
	}
	if !strings.Contains(err.Error(), "build failed") {
		t.Errorf("expected wrapped discovery error, got: %v", err)
	}
	if groups != nil {
		t.Errorf("groups = %v, want nil on discovery error", groups)
	}
	if len(plan.Selected) != 0 || len(plan.Eligible) != 0 {
		t.Errorf("plan = %+v, want zero-value ExecutionPlan on discovery error", plan)
	}
}

func TestPlanForCommandPropagatesPlanError(t *testing.T) {
	root := t.TempDir()
	d := deps{
		provider: planningTestProvider(),
		list:     stubList([]string{"TestAccRepo"}),
	}

	_, groups, err := planForCommand(context.Background(), d, root, planningOptions{Mode: "does-not-exist"})
	if err == nil {
		t.Fatal("expected plan/mode-unknown error, got nil")
	}
	var perr *engine.PlanError
	if !errors.As(err, &perr) {
		t.Fatalf("expected *engine.PlanError, got %T: %v", err, err)
	}
	if perr.Code != "plan/mode-unknown" {
		t.Errorf("PlanError.Code = %q, want plan/mode-unknown", perr.Code)
	}
	if groups != nil {
		t.Errorf("groups = %v, want nil on plan error", groups)
	}
}

func TestPlanExitCodeRegexInvalidIsUsageError(t *testing.T) {
	err := &engine.PlanError{Code: "plan/regex-invalid", Detail: "invalid run pattern"}
	if got := planExitCode(err); got != 2 {
		t.Errorf("planExitCode(regex-invalid) = %d, want 2", got)
	}
}

func TestPlanExitCodeOtherPlanErrorCodesAreRuntimeErrors(t *testing.T) {
	codes := []string{"plan/mode-unknown", "plan/group-unknown", "plan/unclassified", "plan/mode-incompatible"}
	for _, code := range codes {
		t.Run(code, func(t *testing.T) {
			err := &engine.PlanError{Code: code, Detail: "detail"}
			if got := planExitCode(err); got != 1 {
				t.Errorf("planExitCode(%s) = %d, want 1", code, got)
			}
		})
	}
}

func TestPlanExitCodeNonPlanErrorIsRuntimeError(t *testing.T) {
	if got := planExitCode(errors.New("discovering tests: boom")); got != 1 {
		t.Errorf("planExitCode(generic error) = %d, want 1", got)
	}
}

func TestRedactPlanRedactsExclusionDetailAndFix(t *testing.T) {
	red := redact.New([]string{"ghp_SECRETVALUE"})
	plan := engine.ExecutionPlan{
		Mode:     "organization",
		Selected: []string{"TestAcc_ghp_SECRETVALUE"},
		Excluded: []engine.PlanExclusion{{
			Test:          "TestAcc_ghp_SECRETVALUE",
			Code:          "excluded/mode-incompatible",
			Detail:        `test "TestAcc_ghp_SECRETVALUE" does not support mode "organization" (supports: [individual])`,
			Fix:           "run in one of: [individual] (token ghp_SECRETVALUE)",
			RequiredModes: []string{"individual"},
		}},
	}

	redacted := redactPlan(plan, red)

	if strings.Contains(redacted.Excluded[0].Detail, "ghp_SECRETVALUE") {
		t.Errorf("Detail leaked secret: %q", redacted.Excluded[0].Detail)
	}
	if strings.Contains(redacted.Excluded[0].Fix, "ghp_SECRETVALUE") {
		t.Errorf("Fix leaked secret: %q", redacted.Excluded[0].Fix)
	}
	if !strings.Contains(redacted.Excluded[0].Detail, "***REDACTED***") {
		t.Errorf("Detail not redacted: %q", redacted.Excluded[0].Detail)
	}
	if !strings.Contains(redacted.Excluded[0].Fix, "***REDACTED***") {
		t.Errorf("Fix not redacted: %q", redacted.Excluded[0].Fix)
	}
	// The original plan must be left untouched (redactPlan returns a copy).
	if strings.Contains(plan.Excluded[0].Detail, "***REDACTED***") {
		t.Errorf("redactPlan mutated the original plan's Detail: %q", plan.Excluded[0].Detail)
	}
	// Non-text fields must pass through unchanged.
	if redacted.Mode != "organization" || len(redacted.Selected) != 1 {
		t.Errorf("redactPlan altered non-text fields: %+v", redacted)
	}
}

func TestRedactPlanIsNoOpWithoutExclusions(t *testing.T) {
	red := redact.New(nil)
	plan := engine.ExecutionPlan{Mode: "anonymous", Selected: []string{"TestAccA"}, Eligible: []string{"TestAccA"}}
	if got := redactPlan(plan, red); len(got.Excluded) != 0 {
		t.Errorf("redactPlan(no exclusions) = %+v, want unchanged empty Excluded", got)
	}
}

func TestEmitPlanTextPrintsModeSummaryLine(t *testing.T) {
	var buf bytes.Buffer
	plan := engine.ExecutionPlan{
		Mode:         "organization",
		Selected:     []string{"TestAccA", "TestAccB"},
		Eligible:     []string{"TestAccA"},
		Excluded:     []engine.PlanExclusion{{Test: "TestAccB", Code: "excluded/mode-incompatible", Detail: "d", Fix: "f"}},
		Unclassified: nil,
	}
	if err := emitPlan(&buf, nil, formatText, plan); err != nil {
		t.Fatalf("emitPlan: %v", err)
	}
	want := "mode organization: 2 selected, 1 eligible, 1 excluded, 0 unclassified\n"
	if buf.String() != want {
		t.Errorf("emitPlan text = %q, want %q", buf.String(), want)
	}
}

func TestEmitPlanTextPrintsUnclassifiedWarningPerTest(t *testing.T) {
	var buf bytes.Buffer
	plan := engine.ExecutionPlan{
		Mode:         "anonymous",
		Selected:     []string{"TestAccMystery"},
		Unclassified: []string{"TestAccMystery"},
	}
	if err := emitPlan(&buf, nil, formatText, plan); err != nil {
		t.Fatalf("emitPlan: %v", err)
	}
	wantWarning := "warning: TestAccMystery is unclassified; using conservative authenticated requirements and side effect unknown\n"
	if !strings.Contains(buf.String(), wantWarning) {
		t.Errorf("emitPlan text = %q, want to contain %q", buf.String(), wantWarning)
	}
}

func TestEmitPlanJSONWritesPlanEventBeforeExcludedEvents(t *testing.T) {
	var buf bytes.Buffer
	jw := newJSONWriter(&buf)
	plan := engine.ExecutionPlan{
		Mode:         "organization",
		Selected:     []string{"TestAccA", "TestAccB"},
		Eligible:     []string{"TestAccA"},
		Excluded:     []engine.PlanExclusion{{Test: "TestAccB", Code: "excluded/mode-incompatible", Detail: "detail", Fix: "fix", RequiredModes: []string{"individual"}}},
		Unclassified: []string{},
		Scopes:       []string{"repo"},
		Capabilities: []string{"organization"},
		SideEffects:  []string{"repository"},
	}
	if err := emitPlan(&buf, jw, formatJSON, plan); err != nil {
		t.Fatalf("emitPlan: %v", err)
	}
	events := decodeNDJSON(t, buf.String())
	if len(events) != 2 {
		t.Fatalf("events = %d, want 2 (plan + excluded); got %v", len(events), events)
	}
	planEvent := events[0]
	if planEvent["type"] != "plan" || planEvent["mode"] != "organization" {
		t.Fatalf("plan event = %v", planEvent)
	}
	if planEvent["selected"] != float64(2) || planEvent["eligible"] != float64(1) ||
		planEvent["excluded"] != float64(1) || planEvent["unclassified"] != float64(0) {
		t.Fatalf("plan event counts = %v", planEvent)
	}
	scopes, _ := planEvent["scopes"].([]any)
	if len(scopes) != 1 || scopes[0] != "repo" {
		t.Fatalf("plan event scopes = %v, want [repo]", planEvent["scopes"])
	}
	excludedEvent := events[1]
	if excludedEvent["type"] != "excluded" || excludedEvent["test"] != "TestAccB" ||
		excludedEvent["code"] != "excluded/mode-incompatible" || excludedEvent["fix"] != "fix" {
		t.Fatalf("excluded event = %v", excludedEvent)
	}
}

func TestEmitPlanJSONOmitsFixWhenEmpty(t *testing.T) {
	var buf bytes.Buffer
	jw := newJSONWriter(&buf)
	plan := engine.ExecutionPlan{
		Mode:     "anonymous",
		Selected: []string{"TestAccB"},
		Excluded: []engine.PlanExclusion{{Test: "TestAccB", Code: "excluded/mode-incompatible", Detail: "detail", Fix: ""}},
	}
	if err := emitPlan(&buf, jw, formatJSON, plan); err != nil {
		t.Fatalf("emitPlan: %v", err)
	}
	events := decodeNDJSON(t, buf.String())
	if fix, ok := events[1]["fix"]; ok && fix != nil {
		t.Errorf("excluded event fix = %v, want nil/omitted for empty Fix", fix)
	}
}

func TestEmitCapabilitiesRedactsDetailAndFix(t *testing.T) {
	var buf bytes.Buffer
	jw := newJSONWriter(&buf)
	red := redact.New([]string{"ghp_SECRETVALUE"})
	report := provider.PreflightReport{
		Mode: "organization",
		Checks: []provider.Check{
			{Name: "identity", Status: provider.CheckFail, Detail: "token ghp_SECRETVALUE invalid", Fix: "rotate ghp_SECRETVALUE"},
		},
	}
	if err := emitCapabilities(jw, "organization", report, red); err != nil {
		t.Fatalf("emitCapabilities: %v", err)
	}
	events := decodeNDJSON(t, buf.String())
	if len(events) != 1 {
		t.Fatalf("events = %d, want 1", len(events))
	}
	ev := events[0]
	if ev["type"] != "capability" || ev["name"] != "identity" || ev["status"] != "fail" {
		t.Fatalf("capability event = %v", ev)
	}
	if strings.Contains(ev["detail"].(string), "ghp_SECRETVALUE") || strings.Contains(ev["fix"].(string), "ghp_SECRETVALUE") {
		t.Fatalf("capability event leaked secret: %v", ev)
	}
}
