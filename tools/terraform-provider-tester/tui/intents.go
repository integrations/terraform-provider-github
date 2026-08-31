package tui

import "github.com/github/terraform-provider-tester/provider"

// RetryGroupIntent is emitted when the user presses 'r' on a selected group.
// The CLI layer (PR10c) turns this into an actual go test run - the TUI never
// executes go test itself.
type RetryGroupIntent struct {
	Group string // group name to retry
}

// RetryTestIntent is emitted when the user presses 'r' at focusTests or focusLog.
// The CLI layer (PR10c) turns this into an actual go test run for a single test.
type RetryTestIntent struct {
	Test string // top-level test name to retry
}

// RetryAllIntent is emitted when the user presses 'R' in Groups or Run.
// The CLI layer (PR10c) schedules a failures-first retry of all groups.
type RetryAllIntent struct{}

// ResumeIntent asks the producer to resume the last run: re-run all failed AND
// not-run top-level tests from persisted state. Emitted from the resume prompt.
type ResumeIntent struct{}

// SwitchModeIntent asks the producer to switch the active auth mode and re-run
// preflight. Emitted by the mode picker overlay.
type SwitchModeIntent struct {
	Mode string // one of anonymous/individual/organization/team/enterprise
}

// SetEnvVarIntent asks the producer to set an environment variable and re-run
// preflight. Emitted by the variable editor overlay for non-secret fields.
type SetEnvVarIntent struct {
	Key   string
	Value string
}

// PreflightIntent asks the producer to re-run preflight for the current mode.
// Emitted by pressing p from the main dashboard.
type PreflightIntent struct{}

type RefreshTriageIntent struct{}
type SyncKnownIssuesIntent struct{}
type PreviewFileIssueIntent struct{ Fingerprint string }
type ExportReportIntent struct{}
type ListOrphansIntent struct{}

type ConfirmSweepIntent struct {
	Owner     string
	Phrase    string
	Resources []provider.Resource
}

type ConfirmFileIssueIntent struct {
	Fingerprint string
	Phrase      string
}
