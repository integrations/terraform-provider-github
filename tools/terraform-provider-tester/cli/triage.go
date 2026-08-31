package cli

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/github/terraform-provider-tester/engine"
	"github.com/github/terraform-provider-tester/internal/redact"
	"github.com/github/terraform-provider-tester/provider"
	ghissues "github.com/github/terraform-provider-tester/provider/github"
)

const defaultIssuesRepo = "integrations/terraform-provider-github"
const sweepSuggestion = "leftover state suspected; run `terraform-provider-tester orphans` and then `terraform-provider-tester sweep --confirm` before retrying"

type triageOptions struct {
	Enabled            bool
	Retries            int
	RetryBackoff       time.Duration
	RetryMaxBackoff    time.Duration
	RetryPolicy        string
	RetryTimeouts      bool
	IssuesRepo         string
	KnownIssuesPath    string
	KnownIssuesOffline bool
	FileIssues         bool
	ConfirmFileIssues  bool
}

// issueRegistry is the Task 5 seam for known issue lookup and filing.
type issueRegistry interface {
	LookupFingerprint(ctx context.Context, fingerprint, mode string) (engine.KnownIssueMatch, error)
}

type issueFiler interface {
	FindIssueByFingerprint(ctx context.Context, fingerprint string) (*ghissues.IssueMatch, error)
	CreateFailureIssue(ctx context.Context, draft ghissues.IssueDraft) (*ghissues.IssueResult, error)
}

func addTriageRunFlags(fs *flag.FlagSet, opts *triageOptions) {
	opts.RetryBackoff = 30 * time.Second
	opts.RetryMaxBackoff = 5 * time.Minute
	opts.RetryPolicy = "retryable"
	opts.IssuesRepo = defaultIssuesRepo
	fs.BoolVar(&opts.Enabled, "triage", false, "classify failures after the run")
	fs.IntVar(&opts.Retries, "retries", 0, "retry eligible failed top-level tests, max 2")
	fs.DurationVar(&opts.RetryBackoff, "retry-backoff", opts.RetryBackoff, "initial retry backoff")
	fs.DurationVar(&opts.RetryMaxBackoff, "retry-max-backoff", opts.RetryMaxBackoff, "maximum retry backoff")
	fs.StringVar(&opts.RetryPolicy, "retry-policy", opts.RetryPolicy, "retry policy: retryable or none")
	fs.BoolVar(&opts.RetryTimeouts, "retry-timeouts", false, "allow one retry for bare go test timeouts")
	addKnownIssueTriageFlags(fs, opts)
}

func addKnownIssueTriageFlags(fs *flag.FlagSet, opts *triageOptions) {
	if opts.IssuesRepo == "" {
		opts.IssuesRepo = defaultIssuesRepo
	}
	fs.StringVar(&opts.KnownIssuesPath, "known-issues", "", "optional YAML known-issues registry")
	fs.BoolVar(&opts.KnownIssuesOffline, "known-issues-offline", false, "use only the local known-issues file")
	fs.StringVar(&opts.IssuesRepo, "issues-repo", opts.IssuesRepo, "GitHub issue repo owner/name")
	fs.BoolVar(&opts.FileIssues, "file-issues", false, "create issues for real unknown failures after dedup")
	fs.BoolVar(&opts.ConfirmFileIssues, "confirm-file-issues", false, "confirm non-interactive issue filing")
}

func (opts *triageOptions) finalize(fs *flag.FlagSet) error {
	visitedNew := false
	fs.Visit(func(f *flag.Flag) {
		switch f.Name {
		case "triage", "retries", "retry-backoff", "retry-max-backoff", "retry-policy", "retry-timeouts", "known-issues", "known-issues-offline", "issues-repo", "file-issues":
			visitedNew = true
		}
	})
	if opts.IssuesRepo == "" {
		opts.IssuesRepo = defaultIssuesRepo
	}
	if opts.Retries > 0 || opts.FileIssues || visitedNew {
		opts.Enabled = true
	}
	if opts.Retries > 2 {
		return fmt.Errorf("--retries max 2, got %d", opts.Retries)
	}
	if opts.Retries < 0 {
		return fmt.Errorf("--retries must be non-negative")
	}
	if opts.RetryPolicy != "retryable" && opts.RetryPolicy != "none" {
		return fmt.Errorf("--retry-policy must be retryable or none")
	}
	if _, _, err := ghissues.SplitIssueRepoForCLI(opts.IssuesRepo); err != nil {
		return err
	}
	return nil
}

func runWithSpecTriage(ctx context.Context, spec engine.RunSpec, statePath string, prev engine.State, prov provider.Provider, mode string, runner testRunner, out, errOut io.Writer, opts runOptions) int {
	cur := engine.State{
		Version:  engine.StateVersion,
		Provider: prov.Name(),
		Mode:     mode,
		RunAt:    time.Now(),
		History:  cloneHistory(prev.History),
		Plan:     prev.Plan,
		Orphans:  prev.Orphans,
	}
	getenv := opts.Getenv
	if getenv == nil {
		getenv = func(string) string { return "" }
	}
	red := redactorForProvider(prov, getenv)
	failDir := filepath.Join(filepath.Dir(statePath), ".pulsar-failures")
	groupsByTest := testGroupMap(opts.Groups)
	var jw *jsonWriter
	var streamErr error
	if opts.Format == formatJSON {
		jw = newJSONWriter(out)
	}

	retryResult, runErr := engine.RunWithRetries(ctx, runner, spec, engine.RetryOptions{
		Provider: prov.Name(), Mode: mode, Retries: opts.Triage.Retries, Backoff: opts.Triage.RetryBackoff,
		MaxBackoff: opts.Triage.RetryMaxBackoff, RetryPolicy: opts.Triage.RetryPolicy, RetryTimeouts: opts.Triage.RetryTimeouts,
		Redactor: red, PreviousHistory: prev.History, FailureLogDir: failDir,
	}, func(tr engine.TestResult) {
		if opts.Format == formatJSON {
			var failureLog *string
			if tr.Sub == "" && isFailureStatus(tr.Status) {
				failureLog = stringPtr(engine.FailureLogPath(failDir, tr.Package, tr.Name))
			}
			if encErr := jw.Encode(jsonTestEvent{SchemaVersion: 1, Type: "test", Command: opts.Command, Mode: mode, Package: tr.Package, Name: tr.Name, Sub: tr.Sub, Group: groupForTest(groupsByTest, tr.Name), Status: tr.Status.String(), Elapsed: tr.Elapsed, FailureLog: failureLog}); encErr != nil {
				fmt.Fprintln(errOut, "writing JSON:", encErr)
				if streamErr == nil {
					streamErr = encErr
				}
			}
			return
		}
		streamTextResult(out, tr, groupsByTest, failDir)
	})
	res := retryResult.Final
	cur.Results = mergeResults(prev.Results, persistResults(res.Tests))
	cur.History = historyWithFinalResults(prev.History, res.Tests)
	cur.Failures = retryResult.Failures
	persistSuiteFailureState(&cur, failDir, res, retryResult.Initial, red, func(action string, err error) {
		fmt.Fprintln(errOut, action+":", err)
	})

	if _, ferr := engine.WriteFailures(failDir, retryResult.Initial, red); ferr != nil {
		fmt.Fprintln(errOut, "writing failure logs:", ferr)
	}
	if saveErr := cur.Save(statePath); saveErr != nil {
		fmt.Fprintln(errOut, "saving state:", saveErr)
	}
	if runErr != nil {
		fmt.Fprintln(errOut, "run error:", runErr)
		if opts.Format == formatJSON {
			_ = emitRunJSONSummary(jw, opts, spec, statePath, failDir, cur, res, 1)
		} else {
			engine.TerminalSummary(out, opts.Groups, res, red)
		}
		return 1
	}
	updatedFailures, issueErr := processTriageIssues(ctx, cur.Failures, mode, filepath.Dir(statePath), opts.Triage, red, opts.IssueRegistry, opts.IssueFiler, getenv)
	if issueErr != nil {
		fmt.Fprintln(errOut, "triage issues:", issueErr)
		exitCode := exitFromResult(res)
		if exitCode == 0 {
			exitCode = 1
		}
		if streamErr != nil {
			exitCode = 1
		}
		if opts.Format == formatJSON {
			if err := emitRunJSONSummary(jw, opts, spec, statePath, failDir, cur, res, exitCode); err != nil {
				fmt.Fprintln(errOut, "writing JSON:", err)
				return 1
			}
			payload := triagePayloadFromFailures(mode, opts.Triage.IssuesRepo, cur.Failures)
			payload.Type = "triage"
			if err := jw.Encode(payload); err != nil {
				fmt.Fprintln(errOut, "writing JSON:", err)
				return 1
			}
		}
		return exitCode
	}
	cur.Failures = updatedFailures
	if saveErr := cur.Save(statePath); saveErr != nil {
		fmt.Fprintln(errOut, "saving state:", saveErr)
	}

	exitCode := exitFromResult(res)
	if streamErr != nil {
		exitCode = 1
	}
	if opts.Format == formatJSON {
		if err := emitRunJSONSummary(jw, opts, spec, statePath, failDir, cur, res, exitCode); err != nil {
			fmt.Fprintln(errOut, "writing JSON:", err)
			return 1
		}
		payload := triagePayloadFromFailures(mode, opts.Triage.IssuesRepo, cur.Failures)
		payload.Type = "triage"
		if err := jw.Encode(payload); err != nil {
			fmt.Fprintln(errOut, "writing JSON:", err)
			return 1
		}
		return exitCode
	}

	engine.TerminalSummary(out, opts.Groups, res, red)
	if len(cur.Failures) > 0 {
		fmt.Fprintln(out)
		printTriageSummary(out, opts.Triage.IssuesRepo, cur.Failures)
		printIssueDryRuns(out, cur.Failures, filepath.Dir(statePath), opts.Triage, red)
	}
	if matchesRetrySweepSuggestion(cur.Failures) {
		fmt.Fprintln(out, sweepSuggestion)
	}
	if matches, _ := filepath.Glob(filepath.Join(failDir, "*.log")); len(matches) > 0 {
		fmt.Fprintf(out, "\nfailure output written to %s (%d file(s)):\n", failDir, len(matches))
		for _, m := range matches {
			fmt.Fprintf(out, "  %s\n", m)
		}
	}
	return exitCode
}

func runTriage(args []string, out, errOut io.Writer, d deps) int {
	fs := flag.NewFlagSet("triage", flag.ContinueOnError)
	fs.SetOutput(errOut)
	var repoRoot string
	format := formatText
	opts := triageOptions{IssuesRepo: defaultIssuesRepo}
	fs.StringVar(&repoRoot, "repo-root", "", "repo root override")
	addKnownIssueTriageFlags(fs, &opts)
	addJSONFlag(fs, &format)
	if err := fs.Parse(args); err != nil {
		return 2
	}
	if err := validateIssueFilingFlags(opts, d, errOut); err != nil {
		fmt.Fprintln(errOut, err)
		return 2
	}
	root, err := d.resolveRoot(repoRoot)
	if err != nil {
		fmt.Fprintln(errOut, "resolving repo root:", err)
		return 1
	}
	statePath := filepath.Join(root, ".pulsar-state.json")
	getenv := getenvFunc(d)
	red := redact.New(nil)
	if d.provider != nil {
		red = redactorForProvider(d.provider, getenv)
	}
	svc := triageStateService{
		root:            root,
		statePath:       statePath,
		red:             red,
		registryFactory: d.newIssueRegistry,
		filerFactory:    d.newIssueFiler,
		getenv:          getenv,
	}
	result, err := svc.Refresh(context.Background(), opts, false)
	if err != nil {
		fmt.Fprintln(errOut, err)
		return 1
	}
	failures := result.Failures
	mode := result.Mode
	if format == formatJSON {
		payload := triagePayloadFromFailures(mode, opts.IssuesRepo, failures)
		payload.Type = "triage"
		enc := json.NewEncoder(out)
		enc.SetEscapeHTML(false)
		if err := enc.Encode(payload); err != nil {
			fmt.Fprintln(errOut, "writing JSON:", err)
			return 1
		}
		return 0
	}
	if len(failures) == 0 {
		fmt.Fprintln(out, "triage: no failures found")
		return 0
	}
	printTriageSummary(out, opts.IssuesRepo, failures)
	printIssueDryRuns(out, failures, root, opts, red)
	if matchesRetrySweepSuggestion(failures) {
		fmt.Fprintln(out, sweepSuggestion)
	}
	return 0
}

func reconstructFailures(root string, st engine.State) []engine.PersistFailure {
	red := redact.New(nil)
	failDir := filepath.Join(root, ".pulsar-failures")
	var res engine.RunResult
	res.BuildFailed = st.BuildFailed
	res.PreRunFailed = st.PreRunFailed
	if st.BuildLog != "" {
		if b, err := os.ReadFile(st.BuildLog); err == nil {
			res.BuildOutput = []string{string(b)}
		}
	}
	if st.PreRunLog != "" {
		if b, err := os.ReadFile(st.PreRunLog); err == nil {
			res.PreRunOutput = []string{string(b)}
		}
	}
	for _, r := range st.Results {
		status, ok := provider.ParseStatus(r.Status)
		if !ok {
			continue
		}
		tr := engine.TestResult{Package: r.Package, Name: r.Test, Sub: r.Sub, Status: status, Elapsed: r.Elapsed}
		if r.Sub == "" && isFailureStatus(status) {
			if b, err := os.ReadFile(engine.FailureLogPath(failDir, r.Package, r.Test)); err == nil {
				tr.Output = []string{string(b)}
			}
		}
		res.Tests = append(res.Tests, tr)
	}
	sigs := engine.SignaturesForRun(st.Provider, res, red)
	failures := make([]engine.PersistFailure, 0, len(sigs))
	for _, sig := range sigs {
		attempts := []engine.TriageAttempt{{Number: 1, Status: sig.Status, Signature: sig}}
		logPath := ""
		if sig.Package != "" && sig.Test != "" && sig.Test != "suite" {
			logPath = engine.FailureLogPath(failDir, sig.Package, sig.Test)
		}
		failures = append(failures, engine.ClassifyFailure(engine.TriageInput{Mode: st.Mode, Signature: sig, Attempts: attempts, PreviousHistory: st.History[sig.Test], LogPath: logPath}))
	}
	return failures
}

func persistResults(tests []engine.TestResult) []engine.PersistResult {
	out := make([]engine.PersistResult, 0, len(tests))
	for _, tr := range tests {
		out = append(out, engine.PersistResult{Test: tr.Name, Sub: tr.Sub, Status: tr.Status.String(), Elapsed: tr.Elapsed, Package: tr.Package})
	}
	return out
}

func cloneHistory(in map[string][]string) map[string][]string {
	if in == nil {
		return nil
	}
	out := make(map[string][]string, len(in))
	for k, v := range in {
		out[k] = append([]string{}, v...)
	}
	return out
}

func historyWithFinalResults(prev map[string][]string, tests []engine.TestResult) map[string][]string {
	out := cloneHistory(prev)
	for _, tr := range tests {
		if tr.Sub != "" {
			continue
		}
		if out == nil {
			out = make(map[string][]string)
		}
		h := append(out[tr.Name], tr.Status.String())
		if len(h) > 10 {
			h = h[len(h)-10:]
		}
		out[tr.Name] = h
	}
	return out
}

type triagePayload struct {
	Version    int                 `json:"version"`
	Type       string              `json:"type,omitempty"`
	Mode       string              `json:"mode"`
	IssuesRepo string              `json:"issues_repo"`
	Failures   []triageJSONFailure `json:"failures"`
}

type triageJSONFailure struct {
	Test             string `json:"test"`
	Package          string `json:"package"`
	Status           string `json:"status"`
	Classification   string `json:"classification"`
	Fingerprint      string `json:"fingerprint"`
	ShortFingerprint string `json:"short_fingerprint"`
	Class            string `json:"class"`
	Canonical        string `json:"canonical"`
	Retryable        bool   `json:"retryable"`
	Attempts         int    `json:"attempts"`
	KnownIssue       *int   `json:"known_issue"`
	IssueAction      string `json:"issue_action"`
}

func triagePayloadFromFailures(mode, issuesRepo string, failures []engine.PersistFailure) triagePayload {
	items := make([]triageJSONFailure, 0, len(failures))
	for _, f := range failures {
		var known *int
		if f.KnownIssue != 0 {
			v := f.KnownIssue
			known = &v
		}
		items = append(items, triageJSONFailure{Test: f.Test, Package: f.Package, Status: f.Status, Classification: f.Classification, Fingerprint: f.Fingerprint, ShortFingerprint: f.ShortFingerprint, Class: f.Class, Canonical: f.Canonical, Retryable: f.Retryable, Attempts: f.Attempts, KnownIssue: known, IssueAction: issueActionForFailure(f)})
	}
	return triagePayload{Version: 1, Mode: mode, IssuesRepo: issuesRepo, Failures: items}
}

func issueActionForFailure(f engine.PersistFailure) string {
	if f.IssueAction != "" {
		return f.IssueAction
	}
	if f.KnownIssue != 0 {
		return "known"
	}
	switch f.Classification {
	case engine.ClassificationFlakeConfirmed, engine.ClassificationFlakeHistorical:
		return "none"
	default:
		return "dry-run"
	}
}

func printTriageSummary(w io.Writer, issuesRepo string, failures []engine.PersistFailure) {
	for _, f := range failures {
		if f.Classification == engine.ClassificationFlakeConfirmed {
			fmt.Fprintf(w, "triage: %s failed then passed on retry, flake-confirmed, fingerprint %s\n", f.Test, f.ShortFingerprint)
			continue
		}
		switch f.IssueAction {
		case "filed":
			fmt.Fprintf(w, "triage: %s is %s, filed issue #%d in %s, fingerprint %s\n", f.Test, f.Classification, f.KnownIssue, issuesRepo, f.ShortFingerprint)
			continue
		case "dedup":
			fmt.Fprintf(w, "triage: %s is %s, dedup matched issue #%d in %s, fingerprint %s\n", f.Test, f.Classification, f.KnownIssue, issuesRepo, f.ShortFingerprint)
			for _, reason := range f.Reasons {
				if strings.HasPrefix(reason, "matched closed issue") {
					fmt.Fprintf(w, "note: %s\n", reason)
				}
			}
			continue
		case "known":
			if f.KnownIssue != 0 {
				fmt.Fprintf(w, "triage: %s failed, %s, known issue #%d, fingerprint %s\n", f.Test, f.Classification, f.KnownIssue, f.ShortFingerprint)
			} else {
				fmt.Fprintf(w, "triage: %s failed, %s, known issue, fingerprint %s\n", f.Test, f.Classification, f.ShortFingerprint)
			}
			continue
		}
		if f.KnownIssue != 0 {
			fmt.Fprintf(w, "triage: %s failed, %s, known issue #%d, fingerprint %s\n", f.Test, f.Classification, f.KnownIssue, f.ShortFingerprint)
			continue
		}
		fmt.Fprintf(w, "triage: %s is %s and unknown, would file issue in %s, fingerprint %s\n", f.Test, f.Classification, issuesRepo, f.ShortFingerprint)
	}
}

func validateIssueFilingFlags(opts triageOptions, _ deps, errOut io.Writer) error {
	if opts.KnownIssuesOffline && opts.FileIssues {
		return fmt.Errorf("--known-issues-offline cannot be combined with --file-issues")
	}
	if opts.ConfirmFileIssues && !opts.FileIssues {
		fmt.Fprintln(errOut, "warning: --confirm-file-issues ignored without --file-issues")
	}
	if opts.FileIssues && !opts.ConfirmFileIssues {
		return fmt.Errorf("--file-issues requires --confirm-file-issues in non-interactive mode")
	}
	return nil
}

func processTriageIssues(ctx context.Context, failures []engine.PersistFailure, mode, root string, opts triageOptions, red *redact.Redactor, registryFactory func(triageOptions) (issueRegistry, error), filerFactory func(string) (issueFiler, error), getenv func(string) string) ([]engine.PersistFailure, error) {
	if opts.IssuesRepo == "" {
		opts.IssuesRepo = defaultIssuesRepo
	}
	updated := append([]engine.PersistFailure{}, failures...)
	for i := range updated {
		clearRefreshableKnownIssue(&updated[i])
		updated[i].Canonical = red.String(updated[i].Canonical)
	}
	reg, err := issueRegistryForFactory(opts, registryFactory)
	if err != nil {
		return nil, err
	}
	for i := range updated {
		if !shouldConsiderIssue(updated[i]) {
			continue
		}
		match, err := reg.LookupFingerprint(ctx, updated[i].Fingerprint, mode)
		if err != nil {
			return nil, err
		}
		if match.Found && match.Suppress {
			if match.Issue != 0 {
				updated[i].KnownIssue = match.Issue
			}
			updated[i].IssueAction = "known"
			if match.Note != "" {
				updated[i].Reasons = appendUniqueReason(updated[i].Reasons, match.Note)
			}
			continue
		}
	}
	filer, haveFiler, err := issueFilerForDedup(opts, filerFactory, getenv)
	if err != nil {
		return nil, err
	}
	if haveFiler {
		for i := range updated {
			if !shouldFileIssue(updated[i]) {
				continue
			}
			match, err := filer.FindIssueByFingerprint(ctx, updated[i].Fingerprint)
			if err != nil {
				return nil, err
			}
			if match != nil {
				applyIssueDedupMatch(&updated[i], match)
			}
		}
	}
	if opts.FileIssues {
		if filer == nil {
			filer, err = issueFilerForFactory(opts.IssuesRepo, filerFactory)
			if err != nil {
				return nil, err
			}
		}
		for i := range updated {
			if !shouldFileIssue(updated[i]) {
				continue
			}
			draft := ghissues.BuildFailureIssueDraft(updated[i], ghissues.IssueDraftOptions{
				IssuesRepo: opts.IssuesRepo,
				LogLines:   readFailureLogLines(root, updated[i]),
				Redactor:   red,
			})
			result, err := filer.CreateFailureIssue(ctx, draft)
			if err != nil {
				return nil, err
			}
			if result != nil {
				updated[i].KnownIssue = result.Number
			}
			updated[i].IssueAction = "filed"
		}
		return updated, nil
	}
	return updated, nil
}

func issueFilerForDedup(opts triageOptions, factory func(string) (issueFiler, error), getenv func(string) string) (issueFiler, bool, error) {
	if opts.KnownIssuesOffline {
		return nil, false, nil
	}
	if getenv == nil {
		getenv = func(string) string { return "" }
	}
	if opts.FileIssues {
		filer, err := issueFilerForFactory(opts.IssuesRepo, factory)
		if err != nil {
			return nil, false, err
		}
		return filer, filer != nil, nil
	}
	if factory == nil || getenv("GITHUB_TOKEN") == "" {
		return nil, false, nil
	}
	filer, err := factory(opts.IssuesRepo)
	if err != nil {
		return nil, false, nil
	}
	return filer, filer != nil, nil
}

func applyIssueDedupMatch(f *engine.PersistFailure, match *ghissues.IssueMatch) {
	f.KnownIssue = match.Number
	f.IssueAction = "dedup"
	if strings.EqualFold(match.State, "closed") {
		f.Reasons = append(f.Reasons, fmt.Sprintf("matched closed issue #%d; file with --file-issues --confirm-file-issues --reopen-known-issue or open a new issue manually", match.Number))
	}
}

func printIssueDryRuns(out io.Writer, failures []engine.PersistFailure, root string, opts triageOptions, red *redact.Redactor) {
	if opts.FileIssues {
		return
	}
	for _, failure := range failures {
		if !shouldFileIssue(failure) {
			continue
		}
		draft := ghissues.BuildFailureIssueDraft(failure, ghissues.IssueDraftOptions{
			IssuesRepo: opts.IssuesRepo,
			LogLines:   readFailureLogLines(root, failure),
			Redactor:   red,
		})
		fmt.Fprintf(out, "\nissue dry-run title: %s\n", draft.Title)
		fmt.Fprintf(out, "issue dry-run labels: %s\n", strings.Join(draft.Labels, ", "))
		fmt.Fprintf(out, "issue dry-run body:\n%s\n", draft.Body)
	}
}

func issueRegistryForFactory(opts triageOptions, factory func(triageOptions) (issueRegistry, error)) (issueRegistry, error) {
	if opts.KnownIssuesOffline {
		return &engine.KnownIssueRegistry{CachePath: opts.KnownIssuesPath, Offline: true}, nil
	}
	if factory != nil {
		return factory(opts)
	}
	return &engine.KnownIssueRegistry{CachePath: opts.KnownIssuesPath, Offline: true}, nil
}

func buildLiveIssueRegistry(opts triageOptions) (issueRegistry, error) {
	reg := &engine.KnownIssueRegistry{CachePath: opts.KnownIssuesPath, Offline: opts.KnownIssuesOffline}
	if !opts.KnownIssuesOffline {
		client, err := ghissues.NewIssueClientFromEnv(opts.IssuesRepo)
		if err != nil {
			return nil, err
		}
		reg.Live = client
	}
	return reg, nil
}

func issueFilerForFactory(repo string, factory func(string) (issueFiler, error)) (issueFiler, error) {
	if factory != nil {
		return factory(repo)
	}
	return ghissues.NewIssueClientFromEnv(repo)
}

func shouldConsiderIssue(f engine.PersistFailure) bool {
	if f.IssueAction == "filed" || f.IssueAction == "dedup" {
		return false
	}
	return f.Fingerprint != "" && f.Classification != engine.ClassificationFlakeConfirmed && f.Classification != engine.ClassificationFlakeHistorical
}

func shouldFileIssue(f engine.PersistFailure) bool {
	return engine.EligibleForIssueFiling(f)
}

func clearRefreshableKnownIssue(f *engine.PersistFailure) {
	if f.IssueAction != "known" {
		return
	}
	f.KnownIssue = 0
	f.IssueAction = ""
}

func appendUniqueReason(reasons []string, reason string) []string {
	for _, existing := range reasons {
		if existing == reason {
			return reasons
		}
	}
	return append(reasons, reason)
}

func readFailureLogLines(root string, f engine.PersistFailure) []string {
	path := f.LogPath
	if path == "" && f.Package != "" && f.Test != "" && f.Test != "suite" {
		path = engine.FailureLogPath(filepath.Join(root, ".pulsar-failures"), f.Package, f.Test)
	}
	if path == "" {
		return nil
	}
	data, err := readFailureLogFile(root, path)
	if err != nil {
		return nil
	}
	return strings.Split(strings.TrimSuffix(string(data), "\n"), "\n")
}

func readFailureLogFile(root, path string) ([]byte, error) {
	absRoot, rel, ok := failureLogRelPath(root, path)
	if !ok {
		return nil, os.ErrPermission
	}
	rootDir, err := os.OpenRoot(absRoot)
	if err != nil {
		return nil, err
	}
	defer rootDir.Close()
	info, err := rootDir.Lstat(".pulsar-failures")
	if err != nil {
		return nil, err
	}
	if !info.IsDir() || info.Mode()&os.ModeSymlink != 0 {
		return nil, os.ErrPermission
	}
	failureDir, err := rootDir.OpenRoot(".pulsar-failures")
	if err != nil {
		return nil, err
	}
	defer failureDir.Close()
	return failureDir.ReadFile(rel)
}

func failureLogRelPath(root, path string) (string, string, bool) {
	absRoot, err := filepath.Abs(root)
	if err != nil {
		return "", "", false
	}
	failDir := filepath.Join(absRoot, ".pulsar-failures")
	for _, candidate := range failureLogAbsCandidates(absRoot, path) {
		rel, err := filepath.Rel(failDir, candidate)
		if err != nil || !safeRelativePath(rel) {
			continue
		}
		return absRoot, rel, true
	}
	return "", "", false
}

func failureLogAbsCandidates(absRoot, path string) []string {
	clean := filepath.Clean(path)
	if filepath.IsAbs(clean) {
		return []string{clean}
	}
	var out []string
	if abs, err := filepath.Abs(clean); err == nil {
		out = append(out, abs)
	}
	if abs, err := filepath.Abs(filepath.Join(absRoot, clean)); err == nil {
		out = append(out, abs)
	}
	return out
}

func safeRelativePath(path string) bool {
	if path == "" || path == "." || filepath.IsAbs(path) {
		return false
	}
	return path != ".." && !strings.HasPrefix(path, ".."+string(filepath.Separator))
}

func matchesRetrySweepSuggestion(failures []engine.PersistFailure) bool {
	for _, f := range failures {
		if f.Class == engine.ClassLeftoverState || f.Class == engine.ClassConflict {
			return true
		}
	}
	return false
}
