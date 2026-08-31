package tui

// Messages are the engine→TUI data boundary. The TUI receives engine/provider
// pure data values (PreflightReport, Group, TestResult, RunResult, rate, error)
// through these messages and renders them. The package references those data
// types directly for rendering, but it calls NO engine/provider functions and
// performs NO I/O, shelling out, or state mutation: that functional decoupling
// - not import avoidance - is the contract. Producers (the CLI wiring) own all
// behavior, including redacting TestResult.Output before it is sent here.

import (
	"time"

	"github.com/github/terraform-provider-tester/engine"
	"github.com/github/terraform-provider-tester/provider"
)

// PreflightMsg carries a completed preflight report from the engine to the TUI.
// EnvVars carries the environment variable descriptors for the (possibly new)
// mode; a nil slice leaves the model's stored env-var list unchanged.
type PreflightMsg struct {
	Report  provider.PreflightReport
	EnvVars []provider.EnvVar
}

// GroupsMsg carries the grouped test list from the engine to the TUI. Plan is
// the ExecutionPlan the CLI wiring built for the current mode: Groups is
// already filtered to Plan.Eligible before this message is sent, so the model
// never needs to consult Plan itself to know what is displayed or retryable.
type GroupsMsg struct {
	Groups []engine.Group
	Plan   *engine.ExecutionPlan
}

// TestUpdateMsg carries a single test result update (pass/fail/etc.) to the TUI.
type TestUpdateMsg struct{ Result engine.TestResult }

// RunDoneMsg signals that the current run has completed.
type RunDoneMsg struct{ Result engine.RunResult }

// RateMsg carries the current GitHub API rate-limit state to the TUI.
type RateMsg struct {
	Remaining int
	Limit     int
	Reset     time.Time
}

// ErrMsg carries a non-fatal error to the TUI (shown in the status line).
type ErrMsg struct{ Err error }

type OperationStartedMsg struct{ Name string }

type TriageLoadedMsg struct {
	Failures       []engine.PersistFailure
	CacheAvailable bool
}

type OperationErrMsg struct {
	Op  string
	Err error
}

type OperationRejectedMsg struct {
	Op  string
	Err error
}

type KnownIssueSyncDoneMsg struct {
	Count          int
	CachePath      string
	Failures       []engine.PersistFailure
	CacheAvailable bool
}

type ReportExportDoneMsg struct {
	MarkdownPath string
	HTMLPath     string
}

type OrphansListedMsg struct {
	Owner     string
	Resources []provider.Resource
}

type SweepDoneMsg struct {
	Remaining       []provider.Resource
	SnapshotChanged bool
	Failed          bool
	ResidualUnknown bool
}

type IssuePreviewMsg struct {
	Fingerprint      string
	ShortFingerprint string
	IssuesRepo       string
	Title            string
	Labels           []string
	Classification   string
}

type IssueFiledMsg struct {
	Fingerprint string
	IssueNumber int
	Action      string
	Failures    []engine.PersistFailure
}

// RunStartedMsg signals that a run has begun; Total is the number of top-level
// tests expected. Label is a short human description of what is running (a group
// name, "N selected tests", or "all tests") shown next to the spinner. The TUI
// shows a running indicator until RunDoneMsg.
type RunStartedMsg struct {
	Total int
	Label string
}

// ResumePromptMsg tells the TUI a prior run exists and offers to resume it.
// When is the prior run timestamp; Failed/NotRun summarize what resume would run.
type ResumePromptMsg struct {
	When   time.Time
	Failed int
	NotRun int
}
