package tui

import (
	"github.com/github/terraform-provider-tester/engine"
	"github.com/github/terraform-provider-tester/provider"
)

// groupSummary holds aggregated per-group statistics derived from the live
// results slice. All fields are computed purely from engine data - no I/O.
type groupSummary struct {
	Passed, Failed, Running, NotRun, Total int
	Duration                               float64   // sum of top-level Elapsed in seconds
	Status                                 rowStatus // aggregate display status
}

// summarize builds a groupSummary for group g using the provided results slice.
//
// Rules:
//   - Only top-level results (Sub == "") whose Name is in g.Tests are counted.
//   - Failed = Fail | Panic | Timeout.
//   - NotRun = Total − count of distinct g.Tests names that have any top-level result.
//   - Aggregate Status priority: Failed>0 → rowFail; Running>0 → rowRunning;
//     Total>0 && Passed==Total → rowPass; else rowNotRun.
func summarize(g engine.Group, results []engine.TestResult) groupSummary {
	total := len(g.Tests)

	// Build a set of names in the group for O(1) lookup.
	inGroup := make(map[string]bool, total)
	for _, name := range g.Tests {
		inGroup[name] = true
	}

	var passed, failed, running int
	var dur float64
	seen := make(map[string]bool)

	for _, r := range results {
		if r.Sub != "" {
			continue // subtests do not count
		}
		if !inGroup[r.Name] {
			continue // result outside this group
		}
		seen[r.Name] = true
		dur += r.Elapsed
		switch r.Status {
		case provider.StatusPass, provider.StatusSkip:
			passed++
		case provider.StatusFail, provider.StatusPanic, provider.StatusTimeout:
			failed++
		case provider.StatusRunning:
			running++
		}
	}

	notRun := total - len(seen)

	return groupSummary{
		Passed:   passed,
		Failed:   failed,
		Running:  running,
		NotRun:   notRun,
		Total:    total,
		Duration: dur,
		Status:   aggregateStatus(passed, failed, running, total),
	}
}

// aggregateStatus reduces a set of counts to a single display status using a
// fixed priority: any failure dominates, then any running test, then an
// all-passed group, otherwise not-run. Shared by per-group summaries and the
// Groups-table totals row so both agree on the aggregate color.
func aggregateStatus(passed, failed, running, total int) rowStatus {
	switch {
	case failed > 0:
		return rowFail
	case running > 0:
		return rowRunning
	case total > 0 && passed == total:
		return rowPass
	default:
		return rowNotRun
	}
}
