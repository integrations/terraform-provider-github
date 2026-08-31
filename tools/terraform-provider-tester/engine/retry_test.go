package engine

import (
	"context"
	"errors"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/github/terraform-provider-tester/internal/redact"
	"github.com/github/terraform-provider-tester/provider"
)

type retryFakeRunner struct {
	results []RunResult
	specs   []RunSpec
}

func (r *retryFakeRunner) Run(_ context.Context, spec RunSpec, sink func(TestResult)) (RunResult, error) {
	r.specs = append(r.specs, spec)
	idx := len(r.specs) - 1
	if idx >= len(r.results) {
		return RunResult{}, errors.New("unexpected runner call")
	}
	res := r.results[idx]
	for _, tr := range res.Tests {
		if sink != nil {
			sink(tr)
		}
	}
	return res, nil
}

func TestRunWithRetriesRejectsHardCap(t *testing.T) {
	_, err := RunWithRetries(context.Background(), &retryFakeRunner{}, RunSpec{Pattern: "^Test"}, RetryOptions{Retries: 3}, nil)
	if err == nil || !strings.Contains(err.Error(), "max 2") {
		t.Fatalf("expected max retry error, got %v", err)
	}
}

func TestRunWithRetriesDoesNotRetryNonRetryableFailures(t *testing.T) {
	cases := []struct {
		name   string
		result RunResult
	}{
		{name: "build", result: RunResult{BuildFailed: true, BuildOutput: []string{"x.go:10: broken\n"}}},
		{name: "pre run", result: RunResult{PreRunFailed: true, PreRunOutput: []string{"provider TestMain failed\n"}}},
		{name: "permission", result: RunResult{Tests: []TestResult{{Package: "./github", Name: "TestAccA", Status: provider.StatusFail, Output: []string{"403 Resource not accessible by token\n"}}}}},
		{name: "schema", result: RunResult{Tests: []TestResult{{Package: "./github", Name: "TestAccA", Status: provider.StatusFail, Output: []string{"Error: expected visibility to be one of [public private]\n"}}}}},
		{name: "mode incompatible", result: RunResult{Tests: []TestResult{{Package: "./github", Name: "TestAccA", Status: provider.StatusFail, Output: []string{"provider TestMain: selected test requires organization mode; GH_TEST_AUTH_MODE=individual\n"}}}}},
		{name: "capability unavailable", result: RunResult{Tests: []TestResult{{Package: "./github", Name: "TestAccA", Status: provider.StatusFail, Output: []string{"GET /user/codespaces/secrets/public-key: 404 feature unavailable or token lacks Codespaces access\n"}}}}},
		{name: "fixture missing", result: RunResult{Tests: []TestResult{{Package: "./github", Name: "TestAccA", Status: provider.StatusFail, Output: []string{"required fixture GH_TEST_ORG_TEMPLATE_REPOSITORY is missing or inaccessible\n"}}}}},
		{name: "api 404 permission or feature", result: RunResult{Tests: []TestResult{{Package: "./github", Name: "TestAccA", Status: provider.StatusFail, Output: []string{"GET /orgs/example/settings: 404 Not Found; endpoint may require permission or an enabled feature\n"}}}}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			runner := &retryFakeRunner{results: []RunResult{tc.result}}
			_, err := RunWithRetries(context.Background(), runner, RunSpec{Pattern: "^Test", Packages: []string{"./github/..."}}, RetryOptions{Provider: "github", Retries: 2, Backoff: time.Nanosecond, Redactor: redact.New(nil)}, nil)
			if err != nil {
				t.Fatalf("RunWithRetries error: %v", err)
			}
			if len(runner.specs) != 1 {
				t.Fatalf("runner calls = %d, want 1", len(runner.specs))
			}
		})
	}
}

func TestRunWithRetriesStopsAfterFirstPassingRetry(t *testing.T) {
	runner := &retryFakeRunner{results: []RunResult{
		retryableFailure("TestAccA", "403 API rate limit exceeded\n"),
		passingRun("TestAccA"),
	}}
	var slept []time.Duration
	got, err := RunWithRetries(context.Background(), runner, RunSpec{Pattern: "^Test", Packages: []string{"./github/..."}}, RetryOptions{
		Provider: "github", Retries: 2, Backoff: 5 * time.Millisecond, MaxBackoff: time.Second, Redactor: redact.New(nil),
		Sleep: func(_ context.Context, d time.Duration) error { slept = append(slept, d); return nil },
	}, nil)
	if err != nil {
		t.Fatalf("RunWithRetries error: %v", err)
	}
	if len(runner.specs) != 2 {
		t.Fatalf("runner calls = %d, want 2", len(runner.specs))
	}
	if runner.specs[1].Pattern != "^TestAccA$" {
		t.Fatalf("retry pattern = %q, want one top-level test", runner.specs[1].Pattern)
	}
	if !reflect.DeepEqual(slept, []time.Duration{5 * time.Millisecond}) {
		t.Fatalf("sleep durations = %v", slept)
	}
	if got.Failures[0].Classification != "flake-confirmed" {
		t.Fatalf("classification = %q, want flake-confirmed", got.Failures[0].Classification)
	}
}

func TestRunWithRetriesStopsWhenRetryReturnsNonRetryableFailure(t *testing.T) {
	runner := &retryFakeRunner{results: []RunResult{
		retryableFailure("TestAccA", "403 API rate limit exceeded\n"),
		{Tests: []TestResult{{Package: "./github", Name: "TestAccA", Status: provider.StatusFail, Output: []string{"403 Resource not accessible by token\n"}}}},
		passingRun("TestAccA"),
	}}
	got, err := RunWithRetries(context.Background(), runner, RunSpec{Pattern: "^Test", Packages: []string{"./github/..."}}, RetryOptions{
		Provider: "github", Retries: 2, Backoff: 0, MaxBackoff: time.Second, Redactor: redact.New(nil),
		Sleep: func(context.Context, time.Duration) error { return nil },
	}, nil)
	if err != nil {
		t.Fatalf("RunWithRetries error: %v", err)
	}
	if len(runner.specs) != 2 {
		t.Fatalf("runner calls = %d, want 2", len(runner.specs))
	}
	attempts := got.Attempts[retryAttemptKey("./github", "TestAccA")]
	if len(attempts) != 2 {
		t.Fatalf("attempts = %+v, want initial plus one retry", attempts)
	}
	if attempts[1].Signature.Class != ClassPermission {
		t.Fatalf("retry class = %q, want %q", attempts[1].Signature.Class, ClassPermission)
	}
}

func TestRunWithRetriesRetriesEachFailedTopLevelOneAtATime(t *testing.T) {
	runner := &retryFakeRunner{results: []RunResult{
		{Tests: []TestResult{
			{Package: "./github", Name: "TestAccA", Status: provider.StatusFail, Output: []string{"403 API rate limit exceeded\n"}},
			{Package: "./github", Name: "TestAccB", Status: provider.StatusFail, Output: []string{"409 Conflict: repository already exists\n"}},
		}},
		passingRun("TestAccA"),
		passingRun("TestAccB"),
	}}
	_, err := RunWithRetries(context.Background(), runner, RunSpec{Pattern: "^Test", Packages: []string{"./github/..."}}, RetryOptions{
		Provider: "github", Retries: 1, Backoff: 0, Redactor: redact.New(nil),
		Sleep: func(context.Context, time.Duration) error { return nil },
	}, nil)
	if err != nil {
		t.Fatalf("RunWithRetries error: %v", err)
	}
	gotPatterns := []string{runner.specs[1].Pattern, runner.specs[2].Pattern}
	wantPatterns := []string{"^TestAccA$", "^TestAccB$"}
	if !reflect.DeepEqual(gotPatterns, wantPatterns) {
		t.Fatalf("retry patterns = %v, want %v", gotPatterns, wantPatterns)
	}
}

func TestRunWithRetriesHonorsBackoffAndRetryAfterCap(t *testing.T) {
	runner := &retryFakeRunner{results: []RunResult{
		retryableFailure("TestAccA", "403 API rate limit exceeded\nRetry-After: 90\n"),
		retryableFailure("TestAccA", "403 API rate limit exceeded\n"),
		passingRun("TestAccA"),
	}}
	var slept []time.Duration
	_, err := RunWithRetries(context.Background(), runner, RunSpec{Pattern: "^Test", Packages: []string{"./github/..."}}, RetryOptions{
		Provider: "github", Retries: 2, Backoff: 10 * time.Second, MaxBackoff: 30 * time.Second, Redactor: redact.New(nil),
		Sleep: func(_ context.Context, d time.Duration) error { slept = append(slept, d); return nil },
	}, nil)
	if err != nil {
		t.Fatalf("RunWithRetries error: %v", err)
	}
	want := []time.Duration{30 * time.Second, 20 * time.Second}
	if !reflect.DeepEqual(slept, want) {
		t.Fatalf("sleep durations = %v, want %v", slept, want)
	}
}

func TestRunWithRetriesHonorsRetryAfterFromFailedSubtestOutput(t *testing.T) {
	initial := RunResult{Tests: []TestResult{
		{Package: "./github", Name: "TestAccA", Status: provider.StatusFail, Output: []string{"403 API rate limit exceeded\n"}},
		{Package: "./github", Name: "TestAccA", Sub: "child", Status: provider.StatusFail, Output: []string{"Retry-After: 30\n"}},
	}}
	runner := &retryFakeRunner{results: []RunResult{
		initial,
		passingRun("TestAccA"),
	}}
	var slept []time.Duration
	_, err := RunWithRetries(context.Background(), runner, RunSpec{Pattern: "^Test", Packages: []string{"./github/..."}}, RetryOptions{
		Provider: "github", Retries: 1, Backoff: time.Second, MaxBackoff: 10 * time.Second, Redactor: redact.New(nil),
		Sleep: func(_ context.Context, d time.Duration) error { slept = append(slept, d); return nil },
	}, nil)
	if err != nil {
		t.Fatalf("RunWithRetries error: %v", err)
	}
	if !reflect.DeepEqual(slept, []time.Duration{10 * time.Second}) {
		t.Fatalf("sleep durations = %v, want capped Retry-After from failed subtest", slept)
	}
}

func TestRunWithRetriesTimeoutsNeedOptIn(t *testing.T) {
	runner := &retryFakeRunner{results: []RunResult{
		{Tests: []TestResult{{Package: "./github", Name: "TestAccA", Status: provider.StatusTimeout, Output: []string{"panic: test timed out after 10m0s\n"}}}},
	}}
	_, err := RunWithRetries(context.Background(), runner, RunSpec{Pattern: "^Test", Packages: []string{"./github/..."}}, RetryOptions{Provider: "github", Retries: 1, Redactor: redact.New(nil)}, nil)
	if err != nil {
		t.Fatalf("RunWithRetries error: %v", err)
	}
	if len(runner.specs) != 1 {
		t.Fatalf("runner calls without opt-in = %d, want 1", len(runner.specs))
	}

	runner = &retryFakeRunner{results: []RunResult{
		{Tests: []TestResult{{Package: "./github", Name: "TestAccA", Status: provider.StatusTimeout, Output: []string{"panic: test timed out after 10m0s\n"}}}},
		passingRun("TestAccA"),
	}}
	_, err = RunWithRetries(context.Background(), runner, RunSpec{Pattern: "^Test", Packages: []string{"./github/..."}}, RetryOptions{
		Provider: "github", Retries: 1, RetryTimeouts: true, Redactor: redact.New(nil), Sleep: func(context.Context, time.Duration) error { return nil },
	}, nil)
	if err != nil {
		t.Fatalf("RunWithRetries with opt-in error: %v", err)
	}
	if len(runner.specs) != 2 {
		t.Fatalf("runner calls with opt-in = %d, want 2", len(runner.specs))
	}
}

func TestRunWithRetriesTimeoutOptInOnlyAllowsOneRetry(t *testing.T) {
	runner := &retryFakeRunner{results: []RunResult{
		{Tests: []TestResult{{Package: "./github", Name: "TestAccA", Status: provider.StatusTimeout, Output: []string{"panic: test timed out after 10m0s\n"}}}},
		{Tests: []TestResult{{Package: "./github", Name: "TestAccA", Status: provider.StatusTimeout, Output: []string{"panic: test timed out after 10m0s\n"}}}},
		passingRun("TestAccA"),
	}}
	_, err := RunWithRetries(context.Background(), runner, RunSpec{Pattern: "^Test", Packages: []string{"./github/..."}}, RetryOptions{
		Provider: "github", Retries: 2, RetryTimeouts: true, Redactor: redact.New(nil), Sleep: func(context.Context, time.Duration) error { return nil },
	}, nil)
	if err != nil {
		t.Fatalf("RunWithRetries error: %v", err)
	}
	if len(runner.specs) != 2 {
		t.Fatalf("runner calls for timeout = %d, want exactly 2", len(runner.specs))
	}
}

func TestRunWithRetriesScopesCollidingTestNamesByPackage(t *testing.T) {
	runner := &retryFakeRunner{results: []RunResult{
		{Tests: []TestResult{
			{Package: "./pkgA", Name: "TestSame", Status: provider.StatusFail, Output: []string{"403 API rate limit exceeded\n"}},
			{Package: "./pkgB", Name: "TestSame", Status: provider.StatusFail, Output: []string{"409 Conflict: repository already exists\n"}},
		}},
		{Tests: []TestResult{{Package: "./pkgA", Name: "TestSame", Status: provider.StatusPass}}},
		{Tests: []TestResult{{Package: "./pkgB", Name: "TestSame", Status: provider.StatusFail, Output: []string{"409 Conflict: repository already exists\n"}}}},
	}}
	got, err := RunWithRetries(context.Background(), runner, RunSpec{Pattern: "^Test", Packages: []string{"./..."}}, RetryOptions{
		Provider: "github", Retries: 1, Redactor: redact.New(nil), Sleep: func(context.Context, time.Duration) error { return nil },
	}, nil)
	if err != nil {
		t.Fatalf("RunWithRetries error: %v", err)
	}
	if len(runner.specs) != 3 {
		t.Fatalf("runner calls = %d, want initial plus two package-scoped retries", len(runner.specs))
	}
	if !reflect.DeepEqual(runner.specs[1].Packages, []string{"./pkgA"}) || !reflect.DeepEqual(runner.specs[2].Packages, []string{"./pkgB"}) {
		t.Fatalf("retry packages = %v and %v, want ./pkgA then ./pkgB", runner.specs[1].Packages, runner.specs[2].Packages)
	}
	if len(got.Failures) != 2 {
		t.Fatalf("failures = %+v, want two", got.Failures)
	}
	byPkg := map[string]PersistFailure{}
	for _, f := range got.Failures {
		byPkg[f.Package] = f
	}
	if byPkg["./pkgA"].Classification != ClassificationFlakeConfirmed {
		t.Fatalf("pkgA classification = %+v, want flake-confirmed", byPkg["./pkgA"])
	}
	if byPkg["./pkgB"].Classification != ClassificationReal {
		t.Fatalf("pkgB classification = %+v, want real", byPkg["./pkgB"])
	}
}

func retryableFailure(test, line string) RunResult {
	return RunResult{Tests: []TestResult{{Package: "./github", Name: test, Status: provider.StatusFail, Output: []string{line}}}}
}

func passingRun(test string) RunResult {
	return RunResult{Tests: []TestResult{{Package: "./github", Name: test, Status: provider.StatusPass}}}
}
