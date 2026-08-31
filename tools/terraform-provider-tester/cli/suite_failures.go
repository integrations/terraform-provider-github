package cli

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/github/terraform-provider-tester/engine"
	"github.com/github/terraform-provider-tester/internal/redact"
)

func persistSuiteFailureState(state *engine.State, failDir string, flagsFrom, logsFrom engine.RunResult, red *redact.Redactor, report func(action string, err error)) {
	state.BuildFailed = flagsFrom.BuildFailed
	state.PreRunFailed = flagsFrom.PreRunFailed
	state.BuildLog = ""
	state.PreRunLog = ""

	if logsFrom.BuildFailed || logsFrom.PreRunFailed {
		if err := os.MkdirAll(failDir, 0o700); err != nil {
			reportSuiteFailureError(report, "creating failure log dir", err)
			return
		}
		if logsFrom.BuildFailed {
			state.BuildLog = filepath.Join(failDir, "build.log")
			if err := os.WriteFile(state.BuildLog, []byte(strings.Join(redactSuiteFailureLines(red, logsFrom.BuildOutput), "")), 0o600); err != nil {
				reportSuiteFailureError(report, "writing build failure log", err)
			}
		}
		if logsFrom.PreRunFailed {
			state.PreRunLog = filepath.Join(failDir, "pre-run.log")
			if err := os.WriteFile(state.PreRunLog, []byte(strings.Join(redactSuiteFailureLines(red, logsFrom.PreRunOutput), "")), 0o600); err != nil {
				reportSuiteFailureError(report, "writing pre-run failure log", err)
			}
		}
		return
	}

	for _, stale := range []string{"build.log", "pre-run.log"} {
		if err := os.Remove(filepath.Join(failDir, stale)); err != nil && !os.IsNotExist(err) {
			reportSuiteFailureError(report, "removing suite failure log", err)
		}
	}
}

func redactSuiteFailureLines(red *redact.Redactor, lines []string) []string {
	if red == nil {
		return append([]string{}, lines...)
	}
	return red.Lines(lines)
}

func reportSuiteFailureError(report func(action string, err error), action string, err error) {
	if report != nil {
		report(action, err)
	}
}
