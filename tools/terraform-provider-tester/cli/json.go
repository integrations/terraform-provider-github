package cli

import (
	"encoding/json"
	"flag"
	"io"
	"strconv"

	"github.com/github/terraform-provider-tester/engine"
	"github.com/github/terraform-provider-tester/provider"
)

type outputFormat int

const (
	formatText outputFormat = iota
	formatJSON
)

type jsonFlag struct {
	format *outputFormat
}

func (f jsonFlag) IsBoolFlag() bool {
	return true
}

func (f jsonFlag) Set(value string) error {
	enabled, err := strconv.ParseBool(value)
	if err != nil {
		return err
	}
	if enabled {
		*f.format = formatJSON
		return nil
	}
	*f.format = formatText
	return nil
}

func (f jsonFlag) String() string {
	if f.format == nil {
		return "false"
	}
	return strconv.FormatBool(*f.format == formatJSON)
}

func addJSONFlag(fs *flag.FlagSet, format *outputFormat) {
	fs.Var(jsonFlag{format: format}, "json", "emit newline-delimited JSON")
}

type jsonWriter struct {
	enc *json.Encoder
}

func newJSONWriter(w io.Writer) *jsonWriter {
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	return &jsonWriter{enc: enc}
}

func (w *jsonWriter) Encode(v any) error {
	return w.enc.Encode(v)
}

type selectionInfo struct {
	Group string `json:"group"`
	Run   string `json:"run"`
	Tests int    `json:"tests"`
}

type runOptions struct {
	Format              outputFormat
	Command             string
	Selection           selectionInfo
	Groups              []engine.Group
	Triage              triageOptions
	IssueRegistry       func(triageOptions) (issueRegistry, error)
	IssueFiler          func(string) (issueFiler, error)
	Getenv              func(string) string
	CleanupStatus       string
	JSONSummaryObserver func(jsonRunSummary)
}

type jsonTotals struct {
	Total   int `json:"total"`
	Pass    int `json:"pass"`
	Fail    int `json:"fail"`
	Skip    int `json:"skip"`
	Panic   int `json:"panic"`
	Timeout int `json:"timeout"`
}

type jsonTestEvent struct {
	SchemaVersion int     `json:"schema_version"`
	Type          string  `json:"type"`
	Command       string  `json:"command"`
	Mode          string  `json:"mode"`
	Package       string  `json:"package"`
	Name          string  `json:"name"`
	Sub           string  `json:"sub"`
	Group         string  `json:"group"`
	Status        string  `json:"status"`
	Elapsed       float64 `json:"elapsed"`
	FailureLog    *string `json:"failure_log_path"`
}

type jsonFailure struct {
	Name       string  `json:"name"`
	Group      string  `json:"group"`
	Status     string  `json:"status"`
	FailureLog *string `json:"failure_log_path"`
}

type jsonGroupSummary struct {
	Name    string `json:"name"`
	Total   int    `json:"total"`
	Pass    int    `json:"pass"`
	Fail    int    `json:"fail"`
	Skip    int    `json:"skip"`
	Panic   int    `json:"panic"`
	Timeout int    `json:"timeout"`
}

type jsonRunSummary struct {
	SchemaVersion int                `json:"schema_version"`
	Type          string             `json:"type"`
	Command       string             `json:"command"`
	Mode          string             `json:"mode"`
	RepoRoot      string             `json:"repo_root"`
	StatePath     string             `json:"state_path"`
	FailureLogDir string             `json:"failure_log_dir"`
	GoTestCommand string             `json:"go_test_command"`
	Selection     selectionInfo      `json:"selection"`
	Totals        jsonTotals         `json:"totals"`
	Groups        []jsonGroupSummary `json:"groups"`
	BuildFailed   bool               `json:"build_failed"`
	PreRunFailed  bool               `json:"pre_run_failed"`
	BuildLog      *string            `json:"build_log_path"`
	PreRunLog     *string            `json:"pre_run_log_path"`
	Failures      []jsonFailure      `json:"failures"`
	CleanupStatus string             `json:"cleanup_status"`
	ExitCode      int                `json:"exit_code"`
}

type jsonGroupEvent struct {
	SchemaVersion int      `json:"schema_version"`
	Type          string   `json:"type"`
	Name          string   `json:"name"`
	Tests         []string `json:"tests"`
	Total         int      `json:"total"`
}

type jsonGroupsSummary struct {
	SchemaVersion int      `json:"schema_version"`
	Type          string   `json:"type"`
	Command       string   `json:"command"`
	Groups        int      `json:"groups"`
	Tests         int      `json:"tests"`
	Unmatched     int      `json:"unmatched"`
	UnmatchedList []string `json:"unmatched_tests,omitempty"`
	ExitCode      int      `json:"exit_code"`
}

type jsonCheckEvent struct {
	SchemaVersion int     `json:"schema_version"`
	Type          string  `json:"type"`
	Mode          string  `json:"mode"`
	Name          string  `json:"name"`
	Status        string  `json:"status"`
	Detail        string  `json:"detail"`
	Fix           *string `json:"fix"`
}

// jsonPlanEvent is the first event run/preflight emit in JSON mode: the
// resolved ExecutionPlan's counts and aggregated requirements. See
// cli/planning.go's emitPlan.
type jsonPlanEvent struct {
	SchemaVersion int      `json:"schema_version"`
	Type          string   `json:"type"`
	Mode          string   `json:"mode"`
	Selected      int      `json:"selected"`
	Eligible      int      `json:"eligible"`
	Excluded      int      `json:"excluded"`
	Unclassified  int      `json:"unclassified"`
	Scopes        []string `json:"scopes"`
	Capabilities  []string `json:"capabilities"`
	SideEffects   []string `json:"side_effects"`
}

// jsonExcludedEvent reports one engine.PlanExclusion entry. Zero or more of
// these follow the jsonPlanEvent, one per plan.Excluded test.
type jsonExcludedEvent struct {
	SchemaVersion int      `json:"schema_version"`
	Type          string   `json:"type"`
	Mode          string   `json:"mode"`
	Test          string   `json:"test"`
	Code          string   `json:"code"`
	Detail        string   `json:"detail"`
	Fix           *string  `json:"fix"`
	RequiredModes []string `json:"required_modes"`
}

// jsonCapabilityEvent reports one provider.Check from a plan-aware preflight
// run inside `run --json` (see cli/planning.go's emitCapabilities). Standalone
// `preflight --json` continues to emit the existing jsonCheckEvent with
// type "check" instead; the two event types are otherwise identical in shape.
type jsonCapabilityEvent struct {
	SchemaVersion int     `json:"schema_version"`
	Type          string  `json:"type"`
	Mode          string  `json:"mode"`
	Name          string  `json:"name"`
	Status        string  `json:"status"`
	Detail        string  `json:"detail"`
	Fix           *string `json:"fix"`
}

type jsonCheckCounts struct {
	OK   int `json:"ok"`
	Warn int `json:"warn"`
	Fail int `json:"fail"`
}

type jsonPreflightSummary struct {
	SchemaVersion int             `json:"schema_version"`
	Type          string          `json:"type"`
	Command       string          `json:"command"`
	Mode          string          `json:"mode"`
	OK            bool            `json:"ok"`
	Checks        jsonCheckCounts `json:"checks"`
	ExitCode      int             `json:"exit_code"`
}

type jsonOrphanEvent struct {
	SchemaVersion int    `json:"schema_version"`
	Type          string `json:"type"`
	Kind          string `json:"kind"`
	Name          string `json:"name"`
	URL           string `json:"url"`
}

type jsonOrphanDeltaEvent struct {
	SchemaVersion  int    `json:"schema_version"`
	Type           string `json:"type"`
	Mode           string `json:"mode"`
	Baseline       int    `json:"baseline"`
	Final          int    `json:"final"`
	PreExisting    int    `json:"pre_existing"`
	New            int    `json:"new"`
	CleanupStatus  string `json:"cleanup_status"`
	PreviewCommand string `json:"preview_command"`
	CleanupCommand string `json:"cleanup_command"`
}

type jsonOrphansSummary struct {
	SchemaVersion int    `json:"schema_version"`
	Type          string `json:"type"`
	Command       string `json:"command"`
	Mode          string `json:"mode,omitempty"`
	RunDelta      bool   `json:"run_delta,omitempty"`
	Resources     int    `json:"resources"`
	ExitCode      int    `json:"exit_code"`
}

func stringPtr(value string) *string {
	if value == "" {
		return nil
	}
	return &value
}

func testGroupMap(groups []engine.Group) map[string]string {
	byTest := make(map[string]string)
	for _, group := range groups {
		for _, name := range group.Tests {
			byTest[name] = group.Name
		}
	}
	return byTest
}

func groupForTest(byTest map[string]string, name string) string {
	if group := byTest[name]; group != "" {
		return group
	}
	return "misc"
}

func addStatus(total *jsonTotals, status provider.Status) {
	total.Total++
	switch status {
	case provider.StatusPass:
		total.Pass++
	case provider.StatusFail:
		total.Fail++
	case provider.StatusSkip:
		total.Skip++
	case provider.StatusPanic:
		total.Panic++
	case provider.StatusTimeout:
		total.Timeout++
	}
}

func totalsForTests(tests []engine.TestResult, allow func(string) bool) jsonTotals {
	var total jsonTotals
	for _, test := range tests {
		if test.Sub != "" {
			continue
		}
		if allow != nil && !allow(test.Name) {
			continue
		}
		addStatus(&total, test.Status)
	}
	return total
}

func groupSummaries(groups []engine.Group, tests []engine.TestResult) []jsonGroupSummary {
	summaries := make([]jsonGroupSummary, 0, len(groups))
	for _, group := range groups {
		names := make(map[string]bool, len(group.Tests))
		for _, name := range group.Tests {
			names[name] = true
		}
		totals := totalsForTests(tests, func(name string) bool { return names[name] })
		summaries = append(summaries, jsonGroupSummary{
			Name:    group.Name,
			Total:   len(group.Tests),
			Pass:    totals.Pass,
			Fail:    totals.Fail,
			Skip:    totals.Skip,
			Panic:   totals.Panic,
			Timeout: totals.Timeout,
		})
	}
	return summaries
}

func failuresForSummary(tests []engine.TestResult, groups []engine.Group, failDir string) []jsonFailure {
	byTest := testGroupMap(groups)
	failures := []jsonFailure{}
	for _, test := range tests {
		if test.Sub != "" || !isFailureStatus(test.Status) {
			continue
		}
		logPath := engine.FailureLogPath(failDir, test.Package, test.Name)
		failures = append(failures, jsonFailure{
			Name:       test.Name,
			Group:      groupForTest(byTest, test.Name),
			Status:     test.Status.String(),
			FailureLog: stringPtr(logPath),
		})
	}
	return failures
}

func checkCounts(report provider.PreflightReport) jsonCheckCounts {
	var counts jsonCheckCounts
	for _, check := range report.Checks {
		switch check.Status {
		case provider.CheckOK:
			counts.OK++
		case provider.CheckWarn:
			counts.Warn++
		case provider.CheckFail:
			counts.Fail++
		}
	}
	return counts
}
