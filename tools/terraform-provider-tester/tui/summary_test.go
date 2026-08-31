package tui

import (
	"testing"

	"github.com/github/terraform-provider-tester/engine"
	"github.com/github/terraform-provider-tester/provider"
)

func makeResult(name, sub string, status provider.Status, elapsed float64) engine.TestResult {
	return engine.TestResult{
		Package: "pkg",
		Name:    name,
		Sub:     sub,
		Status:  status,
		Elapsed: elapsed,
	}
}

func TestSummarizeCountsByStatus(t *testing.T) {
	tests := []struct {
		name     string
		group    engine.Group
		results  []engine.TestResult
		wantPass int
		wantFail int
		wantRun  int
		wantNot  int
		wantTot  int
		wantDur  float64
		wantStat rowStatus
	}{
		{
			name:     "all pass",
			group:    engine.Group{Name: "g", Tests: []string{"TestA", "TestB"}},
			results:  []engine.TestResult{makeResult("TestA", "", provider.StatusPass, 1.0), makeResult("TestB", "", provider.StatusPass, 2.0)},
			wantPass: 2, wantFail: 0, wantRun: 0, wantNot: 0, wantTot: 2, wantDur: 3.0, wantStat: rowPass,
		},
		{
			name:     "one fail",
			group:    engine.Group{Name: "g", Tests: []string{"TestA", "TestB"}},
			results:  []engine.TestResult{makeResult("TestA", "", provider.StatusPass, 1.0), makeResult("TestB", "", provider.StatusFail, 2.0)},
			wantPass: 1, wantFail: 1, wantRun: 0, wantNot: 0, wantTot: 2, wantDur: 3.0, wantStat: rowFail,
		},
		{
			name:     "one running rest pass",
			group:    engine.Group{Name: "g", Tests: []string{"TestA", "TestB", "TestC"}},
			results:  []engine.TestResult{makeResult("TestA", "", provider.StatusPass, 0.5), makeResult("TestB", "", provider.StatusRunning, 0.0), makeResult("TestC", "", provider.StatusPass, 0.5)},
			wantPass: 2, wantFail: 0, wantRun: 1, wantNot: 0, wantTot: 3, wantDur: 1.0, wantStat: rowRunning,
		},
		{
			name:     "some missing not run",
			group:    engine.Group{Name: "g", Tests: []string{"TestA", "TestB", "TestC"}},
			results:  []engine.TestResult{makeResult("TestA", "", provider.StatusPass, 1.0)},
			wantPass: 1, wantFail: 0, wantRun: 0, wantNot: 2, wantTot: 3, wantDur: 1.0, wantStat: rowNotRun,
		},
		{
			name:     "panic counts as fail",
			group:    engine.Group{Name: "g", Tests: []string{"TestA"}},
			results:  []engine.TestResult{makeResult("TestA", "", provider.StatusPanic, 0.1)},
			wantPass: 0, wantFail: 1, wantRun: 0, wantNot: 0, wantTot: 1, wantDur: 0.1, wantStat: rowFail,
		},
		{
			name:     "timeout counts as fail",
			group:    engine.Group{Name: "g", Tests: []string{"TestA"}},
			results:  []engine.TestResult{makeResult("TestA", "", provider.StatusTimeout, 5.0)},
			wantPass: 0, wantFail: 1, wantRun: 0, wantNot: 0, wantTot: 1, wantDur: 5.0, wantStat: rowFail,
		},
		{
			name:     "empty results all not run",
			group:    engine.Group{Name: "g", Tests: []string{"TestA", "TestB"}},
			results:  nil,
			wantPass: 0, wantFail: 0, wantRun: 0, wantNot: 2, wantTot: 2, wantDur: 0, wantStat: rowNotRun,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := summarize(tc.group, tc.results)
			if got.Passed != tc.wantPass {
				t.Errorf("Passed: want %d, got %d", tc.wantPass, got.Passed)
			}
			if got.Failed != tc.wantFail {
				t.Errorf("Failed: want %d, got %d", tc.wantFail, got.Failed)
			}
			if got.Running != tc.wantRun {
				t.Errorf("Running: want %d, got %d", tc.wantRun, got.Running)
			}
			if got.NotRun != tc.wantNot {
				t.Errorf("NotRun: want %d, got %d", tc.wantNot, got.NotRun)
			}
			if got.Total != tc.wantTot {
				t.Errorf("Total: want %d, got %d", tc.wantTot, got.Total)
			}
			if got.Duration != tc.wantDur {
				t.Errorf("Duration: want %f, got %f", tc.wantDur, got.Duration)
			}
			if got.Status != tc.wantStat {
				t.Errorf("Status: want %d, got %d", tc.wantStat, got.Status)
			}
		})
	}
}

func TestSummarizeIgnoresSubtests(t *testing.T) {
	g := engine.Group{Name: "g", Tests: []string{"TestA", "TestB"}}
	results := []engine.TestResult{
		makeResult("TestA", "", provider.StatusPass, 1.0),
		makeResult("TestA", "Sub/1", provider.StatusFail, 0.1), // subtest - must not affect counts
		makeResult("TestA", "Sub/2", provider.StatusFail, 0.1), // subtest - must not affect counts
		makeResult("TestB", "", provider.StatusPass, 2.0),
	}
	got := summarize(g, results)
	if got.Passed != 2 {
		t.Errorf("Passed: want 2, got %d", got.Passed)
	}
	if got.Failed != 0 {
		t.Errorf("Failed: want 0 (subtests ignored), got %d", got.Failed)
	}
	if got.Status != rowPass {
		t.Errorf("Status: want rowPass, got %d", got.Status)
	}
	if got.Duration != 3.0 {
		t.Errorf("Duration: want 3.0 (only top-level), got %f", got.Duration)
	}
}

func TestSummarizeUnknownNamesExcluded(t *testing.T) {
	g := engine.Group{Name: "g", Tests: []string{"TestA"}}
	results := []engine.TestResult{
		makeResult("TestA", "", provider.StatusPass, 1.0),
		makeResult("TestX", "", provider.StatusFail, 0.5), // not in group
	}
	got := summarize(g, results)
	if got.Passed != 1 {
		t.Errorf("Passed: want 1, got %d", got.Passed)
	}
	if got.Failed != 0 {
		t.Errorf("Failed: want 0 (TestX excluded), got %d", got.Failed)
	}
	if got.Total != 1 {
		t.Errorf("Total: want 1, got %d", got.Total)
	}
}
