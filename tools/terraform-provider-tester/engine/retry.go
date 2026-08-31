package engine

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/github/terraform-provider-tester/internal/redact"
	"github.com/github/terraform-provider-tester/provider"
)

// RetryRunner is the runner seam used by retry orchestration.
type RetryRunner interface {
	Run(ctx context.Context, spec RunSpec, sink func(TestResult)) (RunResult, error)
}

// RateLimitProbe is a Task 5 seam for an authenticated, read-only rate probe.
type RateLimitProbe interface {
	RetryAllowed(ctx context.Context, maxWait time.Duration) (bool, string, error)
}

// RetryOptions controls conservative per-test retry.
type RetryOptions struct {
	Provider        string
	Mode            string
	Retries         int
	Backoff         time.Duration
	MaxBackoff      time.Duration
	RetryPolicy     string
	RetryTimeouts   bool
	Redactor        *redact.Redactor
	Sleep           func(context.Context, time.Duration) error
	Now             func() time.Time
	PreviousHistory map[string][]string
	FailN           int
	LastM           int
	FailureLogDir   string
	RateLimitProbe  RateLimitProbe
}

// RetryRunResult is the safe roll-up from an initial run plus optional retries.
type RetryRunResult struct {
	Initial  RunResult
	Final    RunResult
	Attempts map[string][]TriageAttempt
	Failures []PersistFailure
}

func RunWithRetries(ctx context.Context, runner RetryRunner, spec RunSpec, opts RetryOptions, sink func(TestResult)) (RetryRunResult, error) {
	if opts.Retries > 2 {
		return RetryRunResult{}, fmt.Errorf("--retries max 2, got %d", opts.Retries)
	}
	if opts.Retries < 0 {
		return RetryRunResult{}, fmt.Errorf("--retries must be non-negative")
	}
	if opts.Provider == "" {
		opts.Provider = "github"
	}
	if opts.RetryPolicy == "" {
		opts.RetryPolicy = "retryable"
	}
	if opts.MaxBackoff == 0 {
		opts.MaxBackoff = 5 * time.Minute
	}
	if opts.Sleep == nil {
		opts.Sleep = sleepContext
	}
	if opts.Now == nil {
		opts.Now = time.Now
	}

	initial, err := runner.Run(ctx, spec, sink)
	if err != nil {
		return RetryRunResult{Initial: initial, Final: initial}, err
	}
	out := RetryRunResult{Initial: initial, Final: initial, Attempts: map[string][]TriageAttempt{}}

	if initial.BuildFailed || initial.PreRunFailed || opts.Retries == 0 || opts.RetryPolicy == "none" {
		out.Failures = classifyInitialOnly(initial, opts)
		return out, nil
	}
	if opts.RetryPolicy != "retryable" {
		return out, fmt.Errorf("unsupported retry policy %q", opts.RetryPolicy)
	}

	for _, tr := range failedTopLevelResults(initial) {
		sig := SignatureForTest(opts.Provider, initial, tr, opts.Redactor)
		attempts := []TriageAttempt{{Number: 1, Status: tr.Status.String(), Signature: sig}}
		if !retryEligible(sig, opts) {
			out.Attempts[retryAttemptKey(tr.Package, tr.Name)] = attempts
			out.Failures = append(out.Failures, ClassifyFailure(triageInputFor(tr.Name, sig, attempts, opts)))
			continue
		}
		if opts.RateLimitProbe != nil {
			allowed, reason, probeErr := opts.RateLimitProbe.RetryAllowed(ctx, opts.MaxBackoff)
			if probeErr != nil {
				return out, probeErr
			}
			if !allowed {
				failure := ClassifyFailure(triageInputFor(tr.Name, sig, attempts, opts))
				failure.Classification = "blocked-rate-limit"
				failure.Reasons = append(failure.Reasons, reason)
				out.Attempts[retryAttemptKey(tr.Package, tr.Name)] = attempts
				out.Failures = append(out.Failures, failure)
				continue
			}
		}

		lastResult := initial
		maxRetries := opts.Retries
		if sig.Class == ClassTimeout && opts.RetryTimeouts && maxRetries > 1 {
			maxRetries = 1
		}
		for retry := 1; retry <= maxRetries; retry++ {
			wait := retryDelay(opts, retry, outputForTopLevel(lastResult, tr.Package, tr.Name))
			if wait > 0 {
				if err := opts.Sleep(ctx, wait); err != nil {
					return out, err
				}
			}
			retrySpec := spec
			retrySpec.Pattern = "^" + regexp.QuoteMeta(tr.Name) + "$"
			if tr.Package != "" {
				retrySpec.Packages = []string{tr.Package}
			}
			res, runErr := runner.Run(ctx, retrySpec, sink)
			if runErr != nil {
				return out, runErr
			}
			lastResult = res
			fresh, ok := topLevelResult(res, tr.Package, tr.Name)
			if !ok {
				fresh = TestResult{Package: tr.Package, Name: tr.Name, Status: provider.StatusFail, Output: res.RawLog}
			}
			out.Final = replaceRunResults(out.Final, tr.Package, tr.Name, res)
			if !isFailStatus(fresh.Status) {
				attempts = append(attempts, TriageAttempt{Number: retry + 1, Status: fresh.Status.String()})
				break
			}
			freshSig := SignatureForTest(opts.Provider, res, fresh, opts.Redactor)
			attempts = append(attempts, TriageAttempt{Number: retry + 1, Status: fresh.Status.String(), Signature: freshSig})
			if !retryEligible(freshSig, opts) {
				break
			}
		}
		out.Attempts[retryAttemptKey(tr.Package, tr.Name)] = attempts
		out.Failures = append(out.Failures, ClassifyFailure(triageInputFor(tr.Name, sig, attempts, opts)))
	}
	return out, nil
}

func retryEligible(sig Signature, opts RetryOptions) bool {
	return sig.Retryable || (sig.Class == ClassTimeout && opts.RetryTimeouts)
}

func classifyInitialOnly(res RunResult, opts RetryOptions) []PersistFailure {
	sigs := SignaturesForRun(opts.Provider, res, opts.Redactor)
	failures := make([]PersistFailure, 0, len(sigs))
	for _, sig := range sigs {
		attempts := []TriageAttempt{{Number: 1, Status: sig.Status, Signature: sig}}
		failures = append(failures, ClassifyFailure(triageInputFor(sig.Test, sig, attempts, opts)))
	}
	return failures
}

func triageInputFor(test string, sig Signature, attempts []TriageAttempt, opts RetryOptions) TriageInput {
	logPath := ""
	if opts.FailureLogDir != "" && sig.Package != "" && test != "" && test != "suite" {
		logPath = FailureLogPath(opts.FailureLogDir, sig.Package, test)
	}
	return TriageInput{Mode: opts.Mode, Signature: sig, Attempts: attempts, PreviousHistory: opts.PreviousHistory[test], FailN: opts.FailN, LastM: opts.LastM, LogPath: logPath}
}

func failedTopLevelResults(res RunResult) []TestResult {
	var failed []TestResult
	for _, tr := range res.Tests {
		if tr.Sub == "" && isFailStatus(tr.Status) {
			failed = append(failed, tr)
		}
	}
	return failed
}

func topLevelResult(res RunResult, pkg, test string) (TestResult, bool) {
	for _, tr := range res.Tests {
		if tr.Sub == "" && tr.Name == test && (pkg == "" || tr.Package == pkg) {
			return tr, true
		}
	}
	return TestResult{}, false
}

func outputForTopLevel(res RunResult, pkg, test string) []string {
	if tr, ok := topLevelResult(res, pkg, test); ok {
		return outputForTestFailure(res, tr)
	}
	return append(append([]string{}, res.BuildOutput...), res.PreRunOutput...)
}

func replaceRunResults(base RunResult, pkg, test string, fresh RunResult) RunResult {
	merged := base
	var kept []TestResult
	for _, tr := range merged.Tests {
		if tr.Name != test || (pkg != "" && tr.Package != pkg) {
			kept = append(kept, tr)
		}
	}
	for _, tr := range fresh.Tests {
		if tr.Name == test && (pkg == "" || tr.Package == pkg) {
			kept = append(kept, tr)
		}
	}
	merged.Tests = kept
	merged.BuildFailed = fresh.BuildFailed
	merged.PreRunFailed = fresh.PreRunFailed
	if len(fresh.BuildOutput) > 0 {
		merged.BuildOutput = fresh.BuildOutput
	}
	if len(fresh.PreRunOutput) > 0 {
		merged.PreRunOutput = fresh.PreRunOutput
	}
	return merged
}

func retryAttemptKey(pkg, test string) string {
	return pkg + "\x00" + test
}

func retryDelay(opts RetryOptions, retry int, previousOutput []string) time.Duration {
	base := opts.Backoff
	for i := 1; i < retry; i++ {
		base *= 2
	}
	if header := retryAfterFromLines(previousOutput, opts.Now()); header > base {
		base = header
	}
	if opts.MaxBackoff > 0 && base > opts.MaxBackoff {
		base = opts.MaxBackoff
	}
	return base
}

func sleepContext(ctx context.Context, d time.Duration) error {
	t := time.NewTimer(d)
	defer t.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-t.C:
		return nil
	}
}
