package engine

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"

	"github.com/github/terraform-provider-tester/fakeprovider"
	"github.com/github/terraform-provider-tester/provider"
)

// TestBuildExecutionPlanInvalidModeFails verifies that supplying a mode that
// is not in the provider's supported list returns *PlanError code "plan/mode-unknown".
func TestBuildExecutionPlanInvalidModeFails(t *testing.T) {
	names := []string{"TestAccRepoA"}
	reqs := map[string]provider.TestRequirements{
		"TestAccRepoA": {Modes: []string{"organization"}},
	}
	classified := map[string]bool{"TestAccRepoA": true}
	p := plannerProvider(reqs, classified)

	_, err := BuildExecutionPlan(names, p, PlanOptions{Mode: "badmode"})
	if err == nil {
		t.Fatal("expected error for unknown mode, got nil")
	}
	pe, ok := err.(*PlanError)
	if !ok {
		t.Fatalf("expected *PlanError, got %T: %v", err, err)
	}
	if pe.Code != "plan/mode-unknown" {
		t.Errorf("code = %q; want plan/mode-unknown", pe.Code)
	}
	if pe.Detail == "" {
		t.Error("Detail must not be empty")
	}
	if pe.Fix == "" {
		t.Error("Fix must not be empty")
	}
}

// plannerProvider constructs a Fake configured with the given per-test
// requirements and classified set. A test name absent from classified is
// treated as unknown (RequirementsFor returns false). Any conservative
// requirements stored in req[name] are returned even when classified is false,
// so callers can exercise the conservative-aggregation path.
func plannerProvider(req map[string]provider.TestRequirements, classified map[string]bool) *fakeprovider.Fake {
	return &fakeprovider.Fake{
		ModesVal: []provider.Mode{
			{Name: "organization"},
			{Name: "individual"},
		},
		GroupFunc: func(name string) string {
			switch {
			case strings.HasPrefix(name, "TestAccRepo"):
				return "repositories"
			case strings.HasPrefix(name, "TestAccTeam"):
				return "teams"
			default:
				return "misc"
			}
		},
		RequirementsFn: func(name string) (provider.TestRequirements, bool) {
			r := req[name]
			return r, classified[name]
		},
	}
}

func requireJSONKeys(t *testing.T, got map[string]json.RawMessage, want []string) {
	t.Helper()
	if len(got) != len(want) {
		t.Fatalf("JSON keys = %v; want %v", keysOfRawMap(got), want)
	}
	for _, key := range want {
		if _, ok := got[key]; !ok {
			t.Fatalf("JSON keys = %v; missing %q", keysOfRawMap(got), key)
		}
	}
}

func keysOfRawMap(m map[string]json.RawMessage) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// TestBuildExecutionPlanDefaultExcludesModeIncompatible verifies that when no
// explicit --run regex is supplied, tests whose Modes list does not include the
// active mode are appended to Excluded with code "excluded/mode-incompatible"
// rather than returning an error. It also verifies plan.Mode is reflected.
func TestBuildExecutionPlanDefaultExcludesModeIncompatible(t *testing.T) {
	names := []string{"TestAccRepoA", "TestAccRepoB", "TestAccTeamA"}
	reqs := map[string]provider.TestRequirements{
		"TestAccRepoA": {Modes: []string{"organization", "individual"}},
		"TestAccRepoB": {Modes: []string{"individual"}}, // incompatible with organization
		"TestAccTeamA": {Modes: []string{"organization"}},
	}
	classified := map[string]bool{"TestAccRepoA": true, "TestAccRepoB": true, "TestAccTeamA": true}
	p := plannerProvider(reqs, classified)

	plan, err := BuildExecutionPlan(names, p, PlanOptions{Mode: "organization"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if plan.Mode != "organization" {
		t.Errorf("plan.Mode = %q; want organization", plan.Mode)
	}
	if len(plan.Eligible) != 2 {
		t.Errorf("eligible = %v; want [TestAccRepoA TestAccTeamA]", plan.Eligible)
	}
	if len(plan.Excluded) != 1 {
		t.Fatalf("excluded = %v; want 1 entry for TestAccRepoB", plan.Excluded)
	}
	if plan.Excluded[0].Test != "TestAccRepoB" {
		t.Errorf("excluded[0].Test = %q; want TestAccRepoB", plan.Excluded[0].Test)
	}
	if plan.Excluded[0].Code != "excluded/mode-incompatible" {
		t.Errorf("excluded[0].Code = %q; want excluded/mode-incompatible", plan.Excluded[0].Code)
	}
	if plan.Excluded[0].Detail == "" {
		t.Error("excluded[0].Detail must not be empty")
	}
	if plan.Excluded[0].Fix == "" {
		t.Error("excluded[0].Fix must not be empty")
	}
	// RequiredModes on exclusion must list the modes the test supports.
	if len(plan.Excluded[0].RequiredModes) == 0 {
		t.Error("excluded[0].RequiredModes must not be empty for mode-incompatible exclusion")
	}
}

// TestBuildExecutionPlanGroupFiltersBeforeEligibility verifies that the Group
// option restricts the candidate pool before eligibility is checked, and that
// discovery order is preserved in Selected. Also checks plan.Group is echoed.
func TestBuildExecutionPlanGroupFiltersBeforeEligibility(t *testing.T) {
	names := []string{"TestAccRepoA", "TestAccTeamA"}
	reqs := map[string]provider.TestRequirements{
		"TestAccRepoA": {Modes: []string{"organization"}},
		"TestAccTeamA": {Modes: []string{"organization"}},
	}
	classified := map[string]bool{"TestAccRepoA": true, "TestAccTeamA": true}
	p := plannerProvider(reqs, classified)

	plan, err := BuildExecutionPlan(names, p, PlanOptions{Mode: "organization", Group: "repositories"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if plan.Group != "repositories" {
		t.Errorf("plan.Group = %q; want repositories", plan.Group)
	}
	if len(plan.Selected) != 1 || plan.Selected[0] != "TestAccRepoA" {
		t.Errorf("selected = %v; want [TestAccRepoA]", plan.Selected)
	}
	if len(plan.Eligible) != 1 || plan.Eligible[0] != "TestAccRepoA" {
		t.Errorf("eligible = %v; want [TestAccRepoA]", plan.Eligible)
	}
}

// TestBuildExecutionPlanRegexSelectsMatchingNames verifies that the Run option
// filters names by regexp and that only matching names appear in Selected.
// Also checks plan.Run is echoed.
func TestBuildExecutionPlanRegexSelectsMatchingNames(t *testing.T) {
	names := []string{"TestAccRepoA", "TestAccRepoB", "TestAccTeamA"}
	reqs := map[string]provider.TestRequirements{
		"TestAccRepoA": {Modes: []string{"organization"}},
		"TestAccRepoB": {Modes: []string{"organization"}},
		"TestAccTeamA": {Modes: []string{"organization"}},
	}
	classified := map[string]bool{"TestAccRepoA": true, "TestAccRepoB": true, "TestAccTeamA": true}
	p := plannerProvider(reqs, classified)

	plan, err := BuildExecutionPlan(names, p, PlanOptions{Mode: "organization", Run: "TestAccRepo"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if plan.Run != "TestAccRepo" {
		t.Errorf("plan.Run = %q; want TestAccRepo", plan.Run)
	}
	if len(plan.Selected) != 2 {
		t.Errorf("selected = %v; want [TestAccRepoA TestAccRepoB]", plan.Selected)
	}
	for _, n := range plan.Selected {
		if !strings.HasPrefix(n, "TestAccRepo") {
			t.Errorf("selected contains non-matching name %q", n)
		}
	}
}

// TestBuildExecutionPlanUnknownFailsClosed verifies that a test with no
// classification causes BuildExecutionPlan to return a *PlanError with code
// "plan/unclassified" when AllowUnclassified is false.
// PlanError.Tests must contain the unclassified test name.
func TestBuildExecutionPlanUnknownFailsClosed(t *testing.T) {
	names := []string{"TestAccRepoA", "TestAccUnknown"}
	reqs := map[string]provider.TestRequirements{
		"TestAccRepoA": {Modes: []string{"organization"}},
	}
	classified := map[string]bool{"TestAccRepoA": true}
	p := plannerProvider(reqs, classified)

	_, err := BuildExecutionPlan(names, p, PlanOptions{Mode: "organization"})
	if err == nil {
		t.Fatal("expected error for unknown test, got nil")
	}
	pe, ok := err.(*PlanError)
	if !ok {
		t.Fatalf("expected *PlanError, got %T: %v", err, err)
	}
	if pe.Code != "plan/unclassified" {
		t.Errorf("code = %q; want plan/unclassified", pe.Code)
	}
	if pe.Detail == "" {
		t.Error("Detail must not be empty")
	}
	if len(pe.Tests) == 0 {
		t.Error("PlanError.Tests must list the unclassified test(s)")
	} else if pe.Tests[0] != "TestAccUnknown" {
		t.Errorf("PlanError.Tests[0] = %q; want TestAccUnknown", pe.Tests[0])
	}
}

// TestBuildExecutionPlanAllowUnclassifiedUsesConservativeRequirements verifies
// that when AllowUnclassified is true, the test is placed in Unclassified, its
// conservative requirements are aggregated into top-level plan slices, and
// side effect "unknown" is added.
func TestBuildExecutionPlanAllowUnclassifiedUsesConservativeRequirements(t *testing.T) {
	names := []string{"TestAccRepoA", "TestAccUnknown"}
	reqs := map[string]provider.TestRequirements{
		"TestAccRepoA":   {Modes: []string{"organization"}, Scopes: []string{"admin:org"}},
		"TestAccUnknown": {Scopes: []string{"repo"}}, // conservative, classified=false
	}
	classified := map[string]bool{"TestAccRepoA": true}
	p := plannerProvider(reqs, classified)

	plan, err := BuildExecutionPlan(names, p, PlanOptions{Mode: "organization", AllowUnclassified: true})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(plan.Unclassified) != 1 || plan.Unclassified[0] != "TestAccUnknown" {
		t.Errorf("unclassified = %v; want [TestAccUnknown]", plan.Unclassified)
	}

	// Conservative scope "repo" from the unknown test must be aggregated into plan.Scopes.
	foundRepo := false
	for _, s := range plan.Scopes {
		if s == "repo" {
			foundRepo = true
		}
	}
	if !foundRepo {
		t.Errorf("plan.Scopes %v missing conservative scope \"repo\"", plan.Scopes)
	}

	// Side effect "unknown" must be injected into plan.SideEffects.
	foundUnknown := false
	for _, se := range plan.SideEffects {
		if se == "unknown" {
			foundUnknown = true
		}
	}
	if !foundUnknown {
		t.Errorf("plan.SideEffects %v missing \"unknown\"", plan.SideEffects)
	}
}

// TestBuildExecutionPlanExplicitIncompatibleFails verifies that an explicit
// --run regex that selects a test incompatible with the active mode returns a
// *PlanError with code "plan/mode-incompatible".
// PlanError.Tests must name the incompatible test; RequiredModes must list what it supports.
func TestBuildExecutionPlanExplicitIncompatibleFails(t *testing.T) {
	names := []string{"TestAccRepoA", "TestAccRepoB"}
	reqs := map[string]provider.TestRequirements{
		"TestAccRepoA": {Modes: []string{"organization"}},
		"TestAccRepoB": {Modes: []string{"individual"}}, // incompatible with organization
	}
	classified := map[string]bool{"TestAccRepoA": true, "TestAccRepoB": true}
	p := plannerProvider(reqs, classified)

	_, err := BuildExecutionPlan(names, p, PlanOptions{Mode: "organization", Run: "TestAccRepoB"})
	if err == nil {
		t.Fatal("expected error for explicit incompatible test, got nil")
	}
	pe, ok := err.(*PlanError)
	if !ok {
		t.Fatalf("expected *PlanError, got %T: %v", err, err)
	}
	if pe.Code != "plan/mode-incompatible" {
		t.Errorf("code = %q; want plan/mode-incompatible", pe.Code)
	}
	if pe.Detail == "" {
		t.Error("Detail must not be empty")
	}
	if len(pe.Tests) == 0 {
		t.Error("PlanError.Tests must list the incompatible test(s)")
	} else if pe.Tests[0] != "TestAccRepoB" {
		t.Errorf("PlanError.Tests[0] = %q; want TestAccRepoB", pe.Tests[0])
	}
	if len(pe.RequiredModes) == 0 {
		t.Error("PlanError.RequiredModes must list the modes the test supports")
	} else if pe.RequiredModes[0] != "individual" {
		t.Errorf("PlanError.RequiredModes[0] = %q; want individual", pe.RequiredModes[0])
	}
}

// TestBuildExecutionPlanUnknownGroupFails verifies that specifying a Group that
// does not exist in the provider's classification returns a *PlanError with
// code "plan/group-unknown".
func TestBuildExecutionPlanUnknownGroupFails(t *testing.T) {
	names := []string{"TestAccRepoA"}
	reqs := map[string]provider.TestRequirements{
		"TestAccRepoA": {Modes: []string{"organization"}},
	}
	classified := map[string]bool{"TestAccRepoA": true}
	p := plannerProvider(reqs, classified)

	_, err := BuildExecutionPlan(names, p, PlanOptions{Mode: "organization", Group: "nonexistent"})
	if err == nil {
		t.Fatal("expected error for unknown group, got nil")
	}
	pe, ok := err.(*PlanError)
	if !ok {
		t.Fatalf("expected *PlanError, got %T: %v", err, err)
	}
	if pe.Code != "plan/group-unknown" {
		t.Errorf("code = %q; want plan/group-unknown", pe.Code)
	}
	if pe.Detail == "" {
		t.Error("Detail must not be empty")
	}
}

// TestBuildExecutionPlanInvalidRegexFails verifies that a syntactically invalid
// Run pattern returns a *PlanError with code "plan/regex-invalid".
func TestBuildExecutionPlanInvalidRegexFails(t *testing.T) {
	names := []string{"TestAccRepoA"}
	reqs := map[string]provider.TestRequirements{
		"TestAccRepoA": {Modes: []string{"organization"}},
	}
	classified := map[string]bool{"TestAccRepoA": true}
	p := plannerProvider(reqs, classified)

	_, err := BuildExecutionPlan(names, p, PlanOptions{Mode: "organization", Run: "["})
	if err == nil {
		t.Fatal("expected error for invalid regex, got nil")
	}
	pe, ok := err.(*PlanError)
	if !ok {
		t.Fatalf("expected *PlanError, got %T: %v", err, err)
	}
	if pe.Code != "plan/regex-invalid" {
		t.Errorf("code = %q; want plan/regex-invalid", pe.Code)
	}
	if pe.Detail == "" {
		t.Error("Detail must not be empty")
	}
}

// TestBuildExecutionPlanNoMatchesReturnsValidEmptyPlan verifies that when no
// names match the filter, BuildExecutionPlan returns a valid plan with
// intentional non-nil empty slices (not null) so JSON/state consumers get a
// valid empty plan.
func TestBuildExecutionPlanNoMatchesReturnsValidEmptyPlan(t *testing.T) {
	names := []string{"TestAccRepoA"}
	reqs := map[string]provider.TestRequirements{
		"TestAccRepoA": {Modes: []string{"organization"}},
	}
	classified := map[string]bool{"TestAccRepoA": true}
	p := plannerProvider(reqs, classified)

	plan, err := BuildExecutionPlan(names, p, PlanOptions{Mode: "organization", Run: "TestAccNothingMatches"})
	if err != nil {
		t.Fatalf("unexpected error for empty match: %v", err)
	}

	if plan.Selected == nil {
		t.Error("Selected must be non-nil empty slice, not nil")
	}
	if len(plan.Selected) != 0 {
		t.Errorf("selected = %v; want empty", plan.Selected)
	}
	if plan.Eligible == nil {
		t.Error("Eligible must be non-nil empty slice, not nil")
	}
	if len(plan.Eligible) != 0 {
		t.Errorf("eligible = %v; want empty", plan.Eligible)
	}
	if plan.Excluded == nil {
		t.Error("Excluded must be non-nil empty slice, not nil")
	}
	if len(plan.Excluded) != 0 {
		t.Errorf("excluded = %v; want empty", plan.Excluded)
	}
	if plan.Unclassified == nil {
		t.Error("Unclassified must be non-nil empty slice, not nil")
	}
	if len(plan.Unclassified) != 0 {
		t.Errorf("unclassified = %v; want empty", plan.Unclassified)
	}
	if plan.Scopes == nil {
		t.Error("Scopes must be non-nil empty slice, not nil")
	}
	if len(plan.Scopes) != 0 {
		t.Errorf("scopes = %v; want empty", plan.Scopes)
	}
	if plan.Capabilities == nil {
		t.Error("Capabilities must be non-nil empty slice, not nil")
	}
	if len(plan.Capabilities) != 0 {
		t.Errorf("capabilities = %v; want empty", plan.Capabilities)
	}
	if plan.SideEffects == nil {
		t.Error("SideEffects must be non-nil empty slice, not nil")
	}
	if len(plan.SideEffects) != 0 {
		t.Errorf("sideEffects = %v; want empty", plan.SideEffects)
	}
}

// TestBuildExecutionPlanAggregatesSortedUniqueRequirements verifies that
// top-level plan.Scopes, plan.Capabilities, and plan.SideEffects are built
// from the union of all eligible tests, deduplicated and sorted, with no
// duplicates across overlapping contributions.
func TestBuildExecutionPlanAggregatesSortedUniqueRequirements(t *testing.T) {
	names := []string{"TestAccRepoA", "TestAccRepoB", "TestAccTeamA"}
	reqs := map[string]provider.TestRequirements{
		"TestAccRepoA": {Modes: []string{"organization"}, Scopes: []string{"repo", "admin:org"}, Capabilities: []string{"enterprise"}},
		"TestAccRepoB": {Modes: []string{"organization"}, Scopes: []string{"admin:org"}},
		"TestAccTeamA": {Modes: []string{"organization"}, Scopes: []string{"repo"}, SideEffects: []string{"team"}},
	}
	classified := map[string]bool{"TestAccRepoA": true, "TestAccRepoB": true, "TestAccTeamA": true}
	p := plannerProvider(reqs, classified)

	plan, err := BuildExecutionPlan(names, p, PlanOptions{Mode: "organization"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	wantScopes := []string{"admin:org", "repo"}
	if len(plan.Scopes) != len(wantScopes) {
		t.Errorf("plan.Scopes = %v; want %v", plan.Scopes, wantScopes)
	} else {
		for i, s := range wantScopes {
			if plan.Scopes[i] != s {
				t.Errorf("plan.Scopes[%d] = %q; want %q", i, plan.Scopes[i], s)
			}
		}
	}

	wantCaps := []string{"enterprise"}
	if len(plan.Capabilities) != len(wantCaps) {
		t.Errorf("plan.Capabilities = %v; want %v", plan.Capabilities, wantCaps)
	} else if plan.Capabilities[0] != wantCaps[0] {
		t.Errorf("plan.Capabilities[0] = %q; want %q", plan.Capabilities[0], wantCaps[0])
	}

	wantSE := []string{"team"}
	if len(plan.SideEffects) != len(wantSE) {
		t.Errorf("plan.SideEffects = %v; want %v", plan.SideEffects, wantSE)
	} else if plan.SideEffects[0] != wantSE[0] {
		t.Errorf("plan.SideEffects[0] = %q; want %q", plan.SideEffects[0], wantSE[0])
	}
}

// TestExecutionPlanJSONRoundTripContract verifies the exported wire contract for
// a populated plan including nested PlanExclusion keys and round-trip values.
func TestExecutionPlanJSONRoundTripContract(t *testing.T) {
	plan := ExecutionPlan{
		Mode:  "organization",
		Group: "repositories",
		Run:   "TestAccRepo",
		Selected: []string{
			"TestAccRepoA",
		},
		Eligible: []string{
			"TestAccRepoA",
		},
		Excluded: []PlanExclusion{
			{
				Test:          "TestAccRepoB",
				Code:          "excluded/mode-incompatible",
				Detail:        "test \"TestAccRepoB\" does not support mode \"organization\"",
				Fix:           "run in one of: [individual]",
				RequiredModes: []string{"individual"},
			},
		},
		Unclassified: []string{
			"TestAccUnknown",
		},
		Scopes: []string{"admin:org", "repo"},
		Capabilities: []string{
			"enterprise",
		},
		SideEffects: []string{
			"team",
		},
	}

	blob, err := json.Marshal(plan)
	if err != nil {
		t.Fatalf("json.Marshal: %v", err)
	}

	var top map[string]json.RawMessage
	if err := json.Unmarshal(blob, &top); err != nil {
		t.Fatalf("json.Unmarshal top-level: %v", err)
	}
	requireJSONKeys(t, top, []string{"mode", "group", "run", "selected", "eligible", "excluded", "unclassified", "scopes", "capabilities", "side_effects"})
	for _, key := range []string{"name", "requirements"} {
		if _, ok := top[key]; ok {
			t.Fatalf("unexpected top-level key %q in %s", key, string(blob))
		}
	}

	var excluded []map[string]json.RawMessage
	if err := json.Unmarshal(top["excluded"], &excluded); err != nil {
		t.Fatalf("json.Unmarshal excluded: %v", err)
	}
	if len(excluded) != 1 {
		t.Fatalf("excluded = %v; want 1 item", excluded)
	}
	requireJSONKeys(t, excluded[0], []string{"test", "code", "detail", "fix", "required_modes"})
	for _, key := range []string{"name", "requirements"} {
		if _, ok := excluded[0][key]; ok {
			t.Fatalf("unexpected exclusion key %q in %s", key, string(blob))
		}
	}

	var got ExecutionPlan
	if err := json.Unmarshal(blob, &got); err != nil {
		t.Fatalf("json.Unmarshal plan: %v", err)
	}
	if !reflect.DeepEqual(got, plan) {
		t.Fatalf("round-trip plan = %#v; want %#v", got, plan)
	}
}

// TestExecutionPlanJSONOmitsEmptyOptionalFields verifies that empty optional
// fields stay omitted while the contract still emits non-nil empty arrays for
// the collection fields.
func TestExecutionPlanJSONOmitsEmptyOptionalFields(t *testing.T) {
	plan := ExecutionPlan{
		Mode:         "organization",
		Selected:     []string{},
		Eligible:     []string{},
		Excluded:     []PlanExclusion{},
		Unclassified: []string{},
		Scopes:       []string{},
		Capabilities: []string{},
		SideEffects:  []string{},
	}

	blob, err := json.Marshal(plan)
	if err != nil {
		t.Fatalf("json.Marshal: %v", err)
	}

	var top map[string]json.RawMessage
	if err := json.Unmarshal(blob, &top); err != nil {
		t.Fatalf("json.Unmarshal top-level: %v", err)
	}
	requireJSONKeys(t, top, []string{"mode", "selected", "eligible", "excluded", "unclassified", "scopes", "capabilities", "side_effects"})
	for _, key := range []string{"group", "run", "name", "requirements"} {
		if _, ok := top[key]; ok {
			t.Fatalf("unexpected optional key %q in %s", key, string(blob))
		}
	}

	var got ExecutionPlan
	if err := json.Unmarshal(blob, &got); err != nil {
		t.Fatalf("json.Unmarshal plan: %v", err)
	}
	if !reflect.DeepEqual(got, plan) {
		t.Fatalf("round-trip plan = %#v; want %#v", got, plan)
	}
}

// TestPlanErrorErrorReturnsDetail verifies that PlanError.Error() is the detail
// string used for the public error message.
func TestPlanErrorErrorReturnsDetail(t *testing.T) {
	pe := &PlanError{Detail: "planner detail text"}
	if got, want := pe.Error(), pe.Detail; got != want {
		t.Fatalf("PlanError.Error() = %q; want %q", got, want)
	}
}

// conservativeAuthenticated mirrors the provider catalog's fail-closed
// fallback for an unknown test: every authenticated mode, over-broad scopes,
// and side effect "unknown". Anonymous is deliberately absent, so an unknown
// test is mode-incompatible with anonymous.
func conservativeAuthenticated() provider.TestRequirements {
	return provider.TestRequirements{
		Modes:       []string{"individual", "organization", "team", "enterprise"},
		Scopes:      []string{"repo", "admin:org"},
		SideEffects: []string{"unknown"},
	}
}

// unknownPlannerProvider returns a provider whose modes cover anonymous plus
// the authenticated set, and whose RequirementsFor returns the conservative
// authenticated fallback with ok=false for every name in unknown.
func unknownPlannerProvider(unknown map[string]bool, known map[string]provider.TestRequirements) *fakeprovider.Fake {
	return &fakeprovider.Fake{
		ModesVal: []provider.Mode{
			{Name: "anonymous"},
			{Name: "individual"},
			{Name: "organization"},
			{Name: "team"},
			{Name: "enterprise"},
		},
		GroupFunc: func(string) string { return "misc" },
		RequirementsFn: func(name string) (provider.TestRequirements, bool) {
			if unknown[name] {
				return conservativeAuthenticated(), false
			}
			return known[name], true
		},
	}
}

// TestBuildExecutionPlanAllowedUnknownIsEligibleInCompatibleModes verifies
// that an explicitly allowed unknown test is actually runnable: it uses the
// conservative requirements RequirementsFor returned, and in every mode those
// requirements support it is added to Eligible (not only listed in
// Unclassified, which left it silently unrunnable) while still being reported
// as unclassified so the operator sees the warning.
func TestBuildExecutionPlanAllowedUnknownIsEligibleInCompatibleModes(t *testing.T) {
	for _, mode := range []string{"individual", "organization", "team", "enterprise"} {
		t.Run(mode, func(t *testing.T) {
			p := unknownPlannerProvider(map[string]bool{"TestAccMystery": true}, nil)

			plan, err := BuildExecutionPlan([]string{"TestAccMystery"}, p, PlanOptions{
				Mode:              mode,
				AllowUnclassified: true,
			})
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(plan.Eligible, []string{"TestAccMystery"}) {
				t.Errorf("eligible = %v; want [TestAccMystery] in mode %s", plan.Eligible, mode)
			}
			if !reflect.DeepEqual(plan.Unclassified, []string{"TestAccMystery"}) {
				t.Errorf("unclassified = %v; want [TestAccMystery]", plan.Unclassified)
			}
			if len(plan.Excluded) != 0 {
				t.Errorf("excluded = %v; want none in mode %s", plan.Excluded, mode)
			}
			if !reflect.DeepEqual(plan.Scopes, []string{"admin:org", "repo"}) {
				t.Errorf("scopes = %v; want the conservative scopes aggregated", plan.Scopes)
			}
			if !reflect.DeepEqual(plan.SideEffects, []string{"unknown"}) {
				t.Errorf("side_effects = %v; want [unknown]", plan.SideEffects)
			}
		})
	}
}

// TestBuildExecutionPlanAllowedUnknownExcludedInAnonymous verifies that the
// conservative requirements for an unknown test are authenticated-only, so a
// default (non-explicit) anonymous selection records it as
// excluded/mode-incompatible and never makes it eligible. Its scopes must not
// leak into the plan's aggregated requirements, which would make anonymous
// preflight demand credentials for a test that will not run.
func TestBuildExecutionPlanAllowedUnknownExcludedInAnonymous(t *testing.T) {
	p := unknownPlannerProvider(map[string]bool{"TestAccMystery": true}, nil)

	plan, err := BuildExecutionPlan([]string{"TestAccMystery"}, p, PlanOptions{
		Mode:              "anonymous",
		AllowUnclassified: true,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(plan.Eligible) != 0 {
		t.Errorf("eligible = %v; want none in anonymous mode", plan.Eligible)
	}
	if len(plan.Excluded) != 1 {
		t.Fatalf("excluded = %v; want one entry for TestAccMystery", plan.Excluded)
	}
	if plan.Excluded[0].Test != "TestAccMystery" || plan.Excluded[0].Code != "excluded/mode-incompatible" {
		t.Errorf("excluded[0] = %+v; want TestAccMystery excluded/mode-incompatible", plan.Excluded[0])
	}
	if !reflect.DeepEqual(plan.Excluded[0].RequiredModes, conservativeAuthenticated().Modes) {
		t.Errorf("excluded[0].RequiredModes = %v; want the conservative authenticated modes", plan.Excluded[0].RequiredModes)
	}
	if !reflect.DeepEqual(plan.Unclassified, []string{"TestAccMystery"}) {
		t.Errorf("unclassified = %v; want [TestAccMystery] still reported", plan.Unclassified)
	}
	if len(plan.Scopes) != 0 {
		t.Errorf("scopes = %v; want none aggregated from an excluded test", plan.Scopes)
	}
	if len(plan.SideEffects) != 0 {
		t.Errorf("side_effects = %v; want none aggregated from an excluded test", plan.SideEffects)
	}
}

// TestBuildExecutionPlanExplicitAllowedUnknownIncompatibleFails verifies that
// an explicit --run regex naming an unknown test whose conservative
// requirements do not support the mode fails closed with
// plan/mode-incompatible, exactly like an explicitly selected classified test
// that does not support the mode.
func TestBuildExecutionPlanExplicitAllowedUnknownIncompatibleFails(t *testing.T) {
	p := unknownPlannerProvider(map[string]bool{"TestAccMystery": true}, nil)

	_, err := BuildExecutionPlan([]string{"TestAccMystery"}, p, PlanOptions{
		Mode:              "anonymous",
		Run:               "^TestAccMystery$",
		AllowUnclassified: true,
	})
	if err == nil {
		t.Fatal("expected plan/mode-incompatible for an explicitly selected unknown test, got nil")
	}
	pe, ok := err.(*PlanError)
	if !ok {
		t.Fatalf("expected *PlanError, got %T: %v", err, err)
	}
	if pe.Code != "plan/mode-incompatible" {
		t.Errorf("code = %q; want plan/mode-incompatible", pe.Code)
	}
	if len(pe.Tests) != 1 || pe.Tests[0] != "TestAccMystery" {
		t.Errorf("PlanError.Tests = %v; want [TestAccMystery]", pe.Tests)
	}
	if !reflect.DeepEqual(pe.RequiredModes, conservativeAuthenticated().Modes) {
		t.Errorf("PlanError.RequiredModes = %v; want the conservative authenticated modes", pe.RequiredModes)
	}
}

// TestBuildExecutionPlanRunOverridesGroup verifies the documented contract of
// `--run` ("run regex (overrides --group)", see cli.go and
// docs/cli-reference.md): when Run is non-empty the Group filter is not
// applied at all, so a run match OUTSIDE the named group is still selected.
// The planner previously intersected the two filters, silently selecting
// nothing whenever the regex pointed outside the group.
func TestBuildExecutionPlanRunOverridesGroup(t *testing.T) {
	names := []string{"TestAccRepoA", "TestAccTeamA"}
	reqs := map[string]provider.TestRequirements{
		"TestAccRepoA": {Modes: []string{"organization"}},
		"TestAccTeamA": {Modes: []string{"organization"}},
	}
	classified := map[string]bool{"TestAccRepoA": true, "TestAccTeamA": true}
	p := plannerProvider(reqs, classified)

	plan, err := BuildExecutionPlan(names, p, PlanOptions{
		Mode:  "organization",
		Group: "repositories",
		Run:   "^TestAccTeamA$",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !reflect.DeepEqual(plan.Selected, []string{"TestAccTeamA"}) {
		t.Errorf("selected = %v; want [TestAccTeamA]: --run must override --group", plan.Selected)
	}
	if !reflect.DeepEqual(plan.Eligible, []string{"TestAccTeamA"}) {
		t.Errorf("eligible = %v; want [TestAccTeamA]", plan.Eligible)
	}
}

// TestBuildExecutionPlanUnknownGroupIgnoredWhenRunSet verifies that an
// overridden --group is not validated: once Run is non-empty the group name
// is inert, so an unknown group must NOT produce plan/group-unknown.
func TestBuildExecutionPlanUnknownGroupIgnoredWhenRunSet(t *testing.T) {
	names := []string{"TestAccRepoA", "TestAccRepoB"}
	reqs := map[string]provider.TestRequirements{
		"TestAccRepoA": {Modes: []string{"organization"}},
		"TestAccRepoB": {Modes: []string{"organization"}},
	}
	classified := map[string]bool{"TestAccRepoA": true, "TestAccRepoB": true}
	p := plannerProvider(reqs, classified)

	plan, err := BuildExecutionPlan(names, p, PlanOptions{
		Mode:  "organization",
		Group: "no-such-group",
		Run:   "^TestAccRepoB$",
	})
	if err != nil {
		t.Fatalf("unexpected error for an overridden unknown group: %v", err)
	}
	if !reflect.DeepEqual(plan.Selected, []string{"TestAccRepoB"}) {
		t.Errorf("selected = %v; want [TestAccRepoB]", plan.Selected)
	}
}

// TestBuildExecutionPlanUnknownGroupStillFailsWithoutRun verifies the
// override does not weaken group validation when Run is empty: an unknown
// group is still a fail-closed plan/group-unknown error.
func TestBuildExecutionPlanUnknownGroupStillFailsWithoutRun(t *testing.T) {
	names := []string{"TestAccRepoA"}
	reqs := map[string]provider.TestRequirements{
		"TestAccRepoA": {Modes: []string{"organization"}},
	}
	classified := map[string]bool{"TestAccRepoA": true}
	p := plannerProvider(reqs, classified)

	_, err := BuildExecutionPlan(names, p, PlanOptions{Mode: "organization", Group: "no-such-group"})
	pe, ok := err.(*PlanError)
	if !ok {
		t.Fatalf("expected *PlanError, got %T: %v", err, err)
	}
	if pe.Code != "plan/group-unknown" {
		t.Errorf("code = %q; want plan/group-unknown", pe.Code)
	}
}

// TestBuildExecutionPlanRegexStillValidatedWithGroup verifies that overriding
// --group does not skip --run compilation: an invalid regex is still
// plan/regex-invalid even when a (valid) group name is also supplied.
func TestBuildExecutionPlanRegexStillValidatedWithGroup(t *testing.T) {
	names := []string{"TestAccRepoA"}
	reqs := map[string]provider.TestRequirements{
		"TestAccRepoA": {Modes: []string{"organization"}},
	}
	classified := map[string]bool{"TestAccRepoA": true}
	p := plannerProvider(reqs, classified)

	_, err := BuildExecutionPlan(names, p, PlanOptions{
		Mode:  "organization",
		Group: "repositories",
		Run:   "TestAcc(",
	})
	pe, ok := err.(*PlanError)
	if !ok {
		t.Fatalf("expected *PlanError, got %T: %v", err, err)
	}
	if pe.Code != "plan/regex-invalid" {
		t.Errorf("code = %q; want plan/regex-invalid", pe.Code)
	}
}

// TestBuildExecutionPlanRunOverridingGroupStaysExplicitSelection verifies that
// overriding --group keeps the run selection EXPLICIT: a mode-incompatible
// test matched by the regex still fails closed with plan/mode-incompatible
// rather than being silently recorded as excluded.
func TestBuildExecutionPlanRunOverridingGroupStaysExplicitSelection(t *testing.T) {
	names := []string{"TestAccRepoA", "TestAccTeamA"}
	reqs := map[string]provider.TestRequirements{
		"TestAccRepoA": {Modes: []string{"organization"}},
		"TestAccTeamA": {Modes: []string{"organization"}},
	}
	classified := map[string]bool{"TestAccRepoA": true, "TestAccTeamA": true}
	p := plannerProvider(reqs, classified)

	_, err := BuildExecutionPlan(names, p, PlanOptions{
		Mode:  "individual",
		Group: "repositories",
		Run:   "^TestAccTeamA$",
	})
	pe, ok := err.(*PlanError)
	if !ok {
		t.Fatalf("expected *PlanError, got %T: %v", err, err)
	}
	if pe.Code != "plan/mode-incompatible" {
		t.Errorf("code = %q; want plan/mode-incompatible", pe.Code)
	}
	if !reflect.DeepEqual(pe.Tests, []string{"TestAccTeamA"}) {
		t.Errorf("PlanError.Tests = %v; want [TestAccTeamA]", pe.Tests)
	}
}
