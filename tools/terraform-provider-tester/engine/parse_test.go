package engine

import (
	"os"
	"strings"
	"testing"

	"github.com/github/terraform-provider-tester/provider"
)

func mustOpenFixture(t *testing.T, name string) *os.File {
	t.Helper()
	f, err := os.Open("testdata/" + name)
	if err != nil {
		t.Fatalf("open fixture %s: %v", name, err)
	}
	t.Cleanup(func() { f.Close() })
	return f
}

func TestParseAllPass(t *testing.T) {
	f := mustOpenFixture(t, "ip_ranges.ndjson")
	var sinkCalls []TestResult
	result, err := Parse(f, func(tr TestResult) {
		sinkCalls = append(sinkCalls, tr)
	})
	if err != nil {
		t.Fatalf("Parse: %v", err)
	}
	if result.BuildFailed {
		t.Error("BuildFailed should be false")
	}
	if result.PreRunFailed {
		t.Error("PreRunFailed should be false")
	}
	if len(result.Tests) < 2 {
		t.Fatalf("expected at least 2 tests (parent+subtest), got %d", len(result.Tests))
	}

	var topLevel *TestResult
	for i := range result.Tests {
		if result.Tests[i].Name == "TestAccGithubIpRangesDataSource" && result.Tests[i].Sub == "" {
			topLevel = &result.Tests[i]
		}
	}
	if topLevel == nil {
		t.Fatal("TestAccGithubIpRangesDataSource not found in Tests")
	}
	if topLevel.Status != provider.StatusPass {
		t.Errorf("top-level status: want Pass, got %v", topLevel.Status)
	}

	var sub *TestResult
	for i := range result.Tests {
		if result.Tests[i].Name == "TestAccGithubIpRangesDataSource" && result.Tests[i].Sub != "" {
			sub = &result.Tests[i]
		}
	}
	if sub == nil {
		t.Fatal("subtest of TestAccGithubIpRangesDataSource not found")
	}
	if sub.Status != provider.StatusPass {
		t.Errorf("subtest status: want Pass, got %v", sub.Status)
	}

	if len(sinkCalls) != len(result.Tests) {
		t.Errorf("sink calls %d != tests %d", len(sinkCalls), len(result.Tests))
	}
}

func TestParseBuildFailureMarker(t *testing.T) {
	f := mustOpenFixture(t, "build_fail.ndjson")
	result, err := Parse(f, nil)
	if err != nil {
		t.Fatalf("Parse: %v", err)
	}
	if !result.BuildFailed {
		t.Error("BuildFailed should be true")
	}
	if result.PreRunFailed {
		t.Error("PreRunFailed should be false")
	}
	if len(result.Tests) != 0 {
		t.Errorf("Tests should be empty, got %d", len(result.Tests))
	}
	if len(result.BuildOutput) == 0 {
		t.Error("BuildOutput should be non-empty")
	}
}

func TestParsePreRunFailure(t *testing.T) {
	f := mustOpenFixture(t, "prerun_fail.ndjson")
	result, err := Parse(f, nil)
	if err != nil {
		t.Fatalf("Parse: %v", err)
	}
	if result.BuildFailed {
		t.Error("BuildFailed should be false (this is a pre-run failure, not a build failure)")
	}
	if !result.PreRunFailed {
		t.Error("PreRunFailed should be true")
	}
	var found bool
	for _, line := range result.PreRunOutput {
		if strings.Contains(line, "GITHUB_OWNER environment variable not set") {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("PreRunOutput should contain 'GITHUB_OWNER environment variable not set', got: %v", result.PreRunOutput)
	}
}

func TestParsePanic(t *testing.T) {
	f := mustOpenFixture(t, "panic.ndjson")
	result, err := Parse(f, nil)
	if err != nil {
		t.Fatalf("Parse: %v", err)
	}
	var found bool
	for _, tr := range result.Tests {
		if tr.Sub == "" && tr.Status == provider.StatusPanic {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("expected a test with StatusPanic, got: %+v", result.Tests)
	}
}

func TestParseTimeout(t *testing.T) {
	f := mustOpenFixture(t, "timeout.ndjson")
	var sinkCalls []TestResult
	result, err := Parse(f, func(tr TestResult) {
		sinkCalls = append(sinkCalls, tr)
	})
	if err != nil {
		t.Fatalf("Parse: %v", err)
	}
	if len(result.Tests) < 1 {
		t.Fatalf("expected at least 1 test result (run was reported as 0-tests-success), got %d", len(result.Tests))
	}
	var found bool
	for _, tr := range result.Tests {
		if tr.Name == "TestAccGithubBranchProtection" && tr.Sub == "" && tr.Status == provider.StatusTimeout {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("expected TestAccGithubBranchProtection with StatusTimeout, got: %+v", result.Tests)
	}
	if len(sinkCalls) == 0 {
		t.Error("sink should have been called for the synthesized timeout result")
	}
}

func TestParseNoTestsRunIsNotFailure(t *testing.T) {
	f := mustOpenFixture(t, "no_tests.ndjson")
	result, err := Parse(f, nil)
	if err != nil {
		t.Fatalf("Parse: %v", err)
	}
	if result.BuildFailed {
		t.Error("BuildFailed should be false for a package with no tests to run")
	}
	if result.PreRunFailed {
		t.Error("PreRunFailed should be false for a package with no tests to run")
	}
	if len(result.Tests) != 0 {
		t.Errorf("Tests should be empty for a package with no tests to run, got %d", len(result.Tests))
	}
}

func TestParseNonJSONTolerated(t *testing.T) {
	f := mustOpenFixture(t, "noise.ndjson")
	result, err := Parse(f, nil)
	if err != nil {
		t.Fatalf("Parse should not return error for non-JSON lines: %v", err)
	}
	if len(result.RawLog) == 0 {
		t.Error("RawLog should contain non-JSON lines")
	}
	var found bool
	for _, tr := range result.Tests {
		if tr.Status == provider.StatusPass {
			found = true
			break
		}
	}
	if !found {
		t.Error("expected at least one passing test parsed from noise fixture")
	}
}

func TestParseSubtestParenting(t *testing.T) {
	f := mustOpenFixture(t, "mixed.ndjson")
	result, err := Parse(f, nil)
	if err != nil {
		t.Fatalf("Parse: %v", err)
	}

	// All subtest Names must not contain "/"
	for _, tr := range result.Tests {
		if tr.Sub != "" && strings.Contains(tr.Name, "/") {
			t.Errorf("Name should not contain '/', got: %s", tr.Name)
		}
	}

	// Find the failed subtest (maintainer)
	var failedSub *TestResult
	for i := range result.Tests {
		if result.Tests[i].Sub != "" && result.Tests[i].Status == provider.StatusFail {
			failedSub = &result.Tests[i]
			break
		}
	}
	if failedSub == nil {
		t.Fatal("expected a failed subtest in results")
	}

	// The parent must also be in results (passed parent preserved alongside failed subtest)
	var parentFound bool
	for _, tr := range result.Tests {
		if tr.Sub == "" && tr.Name == failedSub.Name {
			parentFound = true
			break
		}
	}
	if !parentFound {
		t.Errorf("parent test %s not found in results", failedSub.Name)
	}
}

func TestParseEagerSink(t *testing.T) {
	f := mustOpenFixture(t, "ip_ranges.ndjson")
	var order []string
	result, err := Parse(f, func(tr TestResult) {
		name := tr.Name
		if tr.Sub != "" {
			name = tr.Name + "/" + tr.Sub
		}
		order = append(order, name)
	})
	if err != nil {
		t.Fatalf("Parse: %v", err)
	}
	if len(order) != len(result.Tests) {
		t.Errorf("sink called %d times, Tests has %d entries", len(order), len(result.Tests))
	}
	if len(order) < 2 {
		t.Fatal("expected at least 2 results in ip_ranges fixture")
	}
	// Subtest terminal event fires before parent: order[0] should contain "/"
	if !strings.Contains(order[0], "/") {
		t.Errorf("expected subtest (with /) first in stream, got order: %v", order)
	}
}
