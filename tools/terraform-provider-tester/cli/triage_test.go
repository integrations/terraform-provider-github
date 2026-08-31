package cli

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"

	"github.com/github/terraform-provider-tester/engine"
	"github.com/github/terraform-provider-tester/fakeprovider"
	"github.com/github/terraform-provider-tester/provider"
	ghissues "github.com/github/terraform-provider-tester/provider/github"
)

type sequenceRunner struct {
	results []engine.RunResult
	calls   int
	specs   []engine.RunSpec
}

type cancelingTriageRunner struct {
	cancel context.CancelFunc
	result engine.RunResult
}

func (r *cancelingTriageRunner) Run(_ context.Context, _ engine.RunSpec, sink func(engine.TestResult)) (engine.RunResult, error) {
	for _, tr := range r.result.Tests {
		if sink != nil {
			sink(tr)
		}
	}
	r.cancel()
	return r.result, nil
}

type contextRecordingIssueRegistry struct {
	err error
}

func (r *contextRecordingIssueRegistry) LookupFingerprint(ctx context.Context, _, _ string) (engine.KnownIssueMatch, error) {
	r.err = ctx.Err()
	return engine.KnownIssueMatch{}, r.err
}

func (s *sequenceRunner) Run(_ context.Context, spec engine.RunSpec, sink func(engine.TestResult)) (engine.RunResult, error) {
	s.specs = append(s.specs, spec)
	idx := s.calls
	s.calls++
	if idx >= len(s.results) {
		return engine.RunResult{}, nil
	}
	res := s.results[idx]
	for _, tr := range res.Tests {
		if sink != nil {
			sink(tr)
		}
	}
	return res, nil
}

func TestTriageCommandJSONUsesPersistedFailuresAndRedactsDraft(t *testing.T) {
	root := t.TempDir()
	secret := "ENV_SECRET_VALUE_789"
	st := engine.State{Provider: "github", Mode: "organization", Failures: []engine.PersistFailure{{
		Package: "./github", Test: "TestAccThing", Status: "fail", Fingerprint: "sha256:" + strings.Repeat("a", 64), ShortFingerprint: strings.Repeat("a", 16),
		Class: "api/422-leftover-state", Canonical: "422 leftover state ***REDACTED***", Retryable: true, Classification: "real", Attempts: 3, Mode: "organization",
	}}}
	if err := st.Save(filepath.Join(root, ".pulsar-state.json")); err != nil {
		t.Fatal(err)
	}
	failDir := filepath.Join(root, ".pulsar-failures")
	if err := os.MkdirAll(failDir, 0o700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(failDir, "github_TestAccThing.log"), []byte("raw "+secret+"\n"), 0o600); err != nil {
		t.Fatal(err)
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"triage", "--json", "--repo-root", root}, &out, &errOut, deps{provider: triageFakeProvider()})
	if code != 0 {
		t.Fatalf("triage exit code = %d, want 0; stderr=%s", code, errOut.String())
	}
	if strings.Contains(out.String(), secret) {
		t.Fatalf("triage JSON leaked secret: %s", out.String())
	}
	var got map[string]any
	if err := json.Unmarshal(out.Bytes(), &got); err != nil {
		t.Fatalf("triage output is not one JSON object: %v\n%s", err, out.String())
	}
	if got["version"] != float64(1) || got["mode"] != "organization" || got["issues_repo"] != "integrations/terraform-provider-github" {
		t.Fatalf("bad triage envelope: %v", got)
	}
	if got["type"] != "triage" {
		t.Fatalf("type = %v, want triage", got["type"])
	}
	failures := got["failures"].([]any)
	if len(failures) != 1 {
		t.Fatalf("failures = %v, want one", failures)
	}
	failure := failures[0].(map[string]any)
	if failure["issue_action"] != "dry-run" || failure["known_issue"] != nil {
		t.Fatalf("bad failure issue fields: %v", failure)
	}
}

func TestRunTriageKeepsExistingJSONAndTextBehavior(t *testing.T) {
	root := seedTriageFailureState(t, "ghp_FAKEFAKEFAKEFAKEFAKEFAKE")

	var jsonOut, jsonErr bytes.Buffer
	jsonCode := runWithDeps([]string{"triage", "--repo-root", root, "--known-issues-offline", "--json"}, &jsonOut, &jsonErr, deps{})
	if jsonCode != 0 {
		t.Fatalf("triage --json exit code = %d, want 0; stderr=%s", jsonCode, jsonErr.String())
	}
	if jsonErr.Len() != 0 {
		t.Fatalf("triage --json stderr = %q, want empty", jsonErr.String())
	}
	wantJSON := fmt.Sprintf(`{"version":1,"type":"triage","mode":"organization","issues_repo":"integrations/terraform-provider-github","failures":[{"test":"TestAccThing","package":"./github","status":"fail","classification":"real","fingerprint":"sha256:%s","short_fingerprint":"%s","class":"api/422-leftover-state","canonical":"422 leftover state ***REDACTED***","retryable":true,"attempts":3,"known_issue":null,"issue_action":"dry-run"}]}
`, strings.Repeat("a", 64), strings.Repeat("a", 16))
	if jsonOut.String() != wantJSON {
		t.Fatalf("triage --json output mismatch\nwant:\n%s\ngot:\n%s", wantJSON, jsonOut.String())
	}

	var textOut, textErr bytes.Buffer
	textCode := runWithDeps([]string{"triage", "--repo-root", root, "--known-issues-offline"}, &textOut, &textErr, deps{})
	if textCode != 0 {
		t.Fatalf("triage text exit code = %d, want 0; stderr=%s", textCode, textErr.String())
	}
	if textErr.Len() != 0 {
		t.Fatalf("triage text stderr = %q, want empty", textErr.String())
	}
	for _, want := range []string{
		"triage: TestAccThing is real and unknown, would file issue in integrations/terraform-provider-github, fingerprint " + strings.Repeat("a", 16),
		"issue dry-run title:",
		"issue dry-run labels:",
		"issue dry-run body:",
		sweepSuggestion,
	} {
		if !strings.Contains(textOut.String(), want) {
			t.Fatalf("triage text missing %q:\n%s", want, textOut.String())
		}
	}
}

func TestRunTriageJSONUsesStateModeWhenFailureModeIsEmpty(t *testing.T) {
	root := cliScratchDir(t)
	st := engine.State{Provider: "github", Mode: "organization", Failures: []engine.PersistFailure{triageFailureFixture()}}
	st.Failures[0].Mode = ""
	if err := st.Save(filepath.Join(root, ".pulsar-state.json")); err != nil {
		t.Fatal(err)
	}
	failDir := filepath.Join(root, ".pulsar-failures")
	if err := os.MkdirAll(failDir, 0o700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(failDir, "github_TestAccThing.log"), []byte("failure log\n"), 0o600); err != nil {
		t.Fatal(err)
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"triage", "--repo-root", root, "--json"}, &out, &errOut, deps{})
	if code != 0 {
		t.Fatalf("triage exit code = %d, want 0; stderr=%s", code, errOut.String())
	}
	var got triagePayload
	if err := json.Unmarshal(out.Bytes(), &got); err != nil {
		t.Fatalf("triage output is not JSON: %v\n%s", err, out.String())
	}
	if got.Mode != "organization" {
		t.Fatalf("triage mode = %q, want organization", got.Mode)
	}
}

func TestRunTriageRefusesConcurrentStateMutation(t *testing.T) {
	root := seedTriageFailureState(t, "ghp_FAKEFAKEFAKEFAKEFAKEFAKE")
	statePath := filepath.Join(root, ".pulsar-state.json")
	release, err := engine.AcquireLock(statePath + ".lock")
	if err != nil {
		t.Fatalf("AcquireLock: %v", err)
	}
	defer func() {
		if err := release(); err != nil {
			t.Fatalf("release lock: %v", err)
		}
	}()

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"triage", "--repo-root", root}, &out, &errOut, deps{})
	if code != 1 {
		t.Fatalf("triage exit code = %d, want 1; stdout=%s stderr=%s", code, out.String(), errOut.String())
	}
	if out.Len() != 0 {
		t.Fatalf("triage stdout = %q, want empty when lock acquisition fails", out.String())
	}
	if !strings.Contains(errOut.String(), "acquiring lock: lock already held") {
		t.Fatalf("triage stderr missing lock failure: %s", errOut.String())
	}
}

func TestRunRetriesFlagImpliesTriageAndPersistsConfirmedFlake(t *testing.T) {
	root := t.TempDir()
	runner := &sequenceRunner{results: []engine.RunResult{
		{Tests: []engine.TestResult{{Package: "./github", Name: "TestAccThing", Status: provider.StatusFail, Output: []string{"403 API rate limit exceeded\n"}}}},
		{Tests: []engine.TestResult{{Package: "./github", Name: "TestAccThing", Status: provider.StatusPass}}},
	}}
	d := deps{provider: triageFakeProvider(), newRunner: func(_ io.Writer) testRunner { return runner }, list: stubList([]string{"TestAccThing"}), cwd: func() (string, error) { return root, nil }}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"run", "--repo-root", root, "--retries", "1", "--retry-backoff", "0s"}, &out, &errOut, d)
	if code != 0 {
		t.Fatalf("run exit code = %d, want 0 after passing retry; stderr=%s stdout=%s", code, errOut.String(), out.String())
	}
	if runner.calls != 2 {
		t.Fatalf("runner calls = %d, want 2", runner.calls)
	}
	loaded, err := engine.Load(filepath.Join(root, ".pulsar-state.json"))
	if err != nil {
		t.Fatal(err)
	}
	if len(loaded.Failures) != 1 {
		t.Fatalf("state failures = %+v, want one", loaded.Failures)
	}
	if loaded.Failures[0].Classification != "flake-confirmed" {
		t.Fatalf("classification = %q, want flake-confirmed", loaded.Failures[0].Classification)
	}
}

func TestRunTriageIssueErrorPersistsStateAndEmitsJSONSummary(t *testing.T) {
	root := cliScratchDir(t)
	runner := &sequenceRunner{results: []engine.RunResult{
		{Tests: []engine.TestResult{{Package: "./github", Name: "TestAccThing", Status: provider.StatusFail, Output: []string{"403 API rate limit exceeded\n"}}}},
	}}
	registryErr := errors.New("registry unavailable")
	d := deps{
		provider: triageFakeProvider(),
		newRunner: func(_ io.Writer) testRunner {
			return runner
		},
		list: stubList([]string{"TestAccThing"}),
		cwd:  func() (string, error) { return root, nil },
		newIssueRegistry: func(triageOptions) (issueRegistry, error) {
			return nil, registryErr
		},
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"run", "--repo-root", root, "--triage", "--json"}, &out, &errOut, d)
	if code == 0 {
		t.Fatalf("run exit code = %d, want non-zero; stderr=%s stdout=%s", code, errOut.String(), out.String())
	}
	if !strings.Contains(errOut.String(), "triage issues: registry unavailable") {
		t.Fatalf("stderr missing issue-processing error: %s", errOut.String())
	}
	loaded, err := engine.Load(filepath.Join(root, ".pulsar-state.json"))
	if err != nil {
		t.Fatalf("state was not persisted after issue error: %v", err)
	}
	if len(loaded.Results) != 1 || loaded.Results[0].Test != "TestAccThing" || loaded.Results[0].Status != "fail" {
		t.Fatalf("state results = %+v, want failed TestAccThing", loaded.Results)
	}
	if len(loaded.Failures) != 1 || loaded.Failures[0].Test != "TestAccThing" {
		t.Fatalf("state failures = %+v, want classified failure", loaded.Failures)
	}

	dec := json.NewDecoder(strings.NewReader(out.String()))
	seenSummary := false
	seenTriage := false
	for {
		var obj map[string]any
		err := dec.Decode(&obj)
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			t.Fatalf("decode JSON stream: %v\n%s", err, out.String())
		}
		switch obj["type"] {
		case "summary":
			seenSummary = true
			if obj["exit_code"] == float64(0) {
				t.Fatalf("summary exit_code = 0, want non-zero: %v", obj)
			}
		case "triage":
			seenTriage = true
			failures, _ := obj["failures"].([]any)
			if len(failures) != 1 {
				t.Fatalf("triage failures = %v, want one", obj["failures"])
			}
		}
	}
	if !seenSummary {
		t.Fatalf("JSON stream missing summary object:\n%s", out.String())
	}
	if !seenTriage {
		t.Fatalf("JSON stream missing terminal triage object:\n%s", out.String())
	}
}

func TestRunWithSpecTriagePropagatesCanceledContextToIssueProcessing(t *testing.T) {
	root := t.TempDir()
	statePath := filepath.Join(root, ".pulsar-state.json")
	ctx, cancel := context.WithCancel(context.Background())
	runner := &cancelingTriageRunner{
		cancel: cancel,
		result: engine.RunResult{Tests: []engine.TestResult{{
			Package: "./github",
			Name:    "TestAccThing",
			Status:  provider.StatusFail,
			Output:  []string{"403 API rate limit exceeded\n"},
		}}},
	}
	registry := &contextRecordingIssueRegistry{}
	var out, errOut bytes.Buffer

	code := runWithSpecTriage(
		ctx,
		engine.RunSpec{Dir: root},
		statePath,
		engine.State{},
		triageFakeProvider(),
		"organization",
		runner,
		&out,
		&errOut,
		runOptions{
			Format:  formatJSON,
			Command: "run",
			Triage:  triageOptions{Enabled: true},
			IssueRegistry: func(triageOptions) (issueRegistry, error) {
				return registry, nil
			},
		},
	)

	if code != 1 {
		t.Fatalf("exit code = %d, want 1 after canceled issue processing; stderr=%s stdout=%s",
			code, errOut.String(), out.String())
	}
	if !errors.Is(registry.err, context.Canceled) {
		t.Fatalf("issue registry context error = %v, want run context cancellation", registry.err)
	}
	if !strings.Contains(errOut.String(), "triage issues: context canceled") {
		t.Fatalf("stderr = %q, want canceled triage issue processing", errOut.String())
	}
}

func TestRunWithoutNewTriageFlagsKeepsDefaultPath(t *testing.T) {
	root := t.TempDir()
	runner := &sequenceRunner{results: []engine.RunResult{
		{Tests: []engine.TestResult{{Package: "./github", Name: "TestAccThing", Status: provider.StatusFail, Output: []string{"403 API rate limit exceeded\n"}}}},
	}}
	d := deps{provider: triageFakeProvider(), newRunner: func(_ io.Writer) testRunner { return runner }, list: stubList([]string{"TestAccThing"}), cwd: func() (string, error) { return root, nil }}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"run", "--repo-root", root}, &out, &errOut, d)
	if code != 1 {
		t.Fatalf("run exit code = %d, want 1; stderr=%s", code, errOut.String())
	}
	if runner.calls != 1 {
		t.Fatalf("runner calls = %d, want default one call", runner.calls)
	}
	loaded, err := engine.Load(filepath.Join(root, ".pulsar-state.json"))
	if err != nil {
		t.Fatal(err)
	}
	if len(loaded.Failures) != 0 {
		t.Fatalf("default run persisted triage failures: %+v", loaded.Failures)
	}
}

func TestRunRejectsRetriesAboveCap(t *testing.T) {
	root := t.TempDir()
	runner := &sequenceRunner{}
	d := deps{provider: triageFakeProvider(), newRunner: func(_ io.Writer) testRunner { return runner }, list: stubList([]string{"TestAccThing"}), cwd: func() (string, error) { return root, nil }}
	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"run", "--repo-root", root, "--retries", "3"}, &out, &errOut, d)
	if code != 2 {
		t.Fatalf("exit code = %d, want usage error 2; stderr=%s", code, errOut.String())
	}
	if runner.calls != 0 {
		t.Fatalf("runner calls = %d, want 0", runner.calls)
	}
	if !strings.Contains(errOut.String(), "max 2") {
		t.Fatalf("stderr missing cap message: %s", errOut.String())
	}
}

func triageFakeProvider() *fakeprovider.Fake {
	return &fakeprovider.Fake{
		NameVal: "github", Packages: []string{"./github/..."}, Pattern: "^TestAcc",
		ModesVal:       testModes,
		Secrets:        []string{"GITHUB_TOKEN"},
		RequirementsFn: allowAllRequirements,
		PreflightFn: func(_ context.Context, mode string, _ provider.TestRequirements) provider.PreflightReport {
			return provider.PreflightReport{Mode: mode}
		},
	}
}

func TestTriagePayloadDoesNotDraftIssuesForFlakes(t *testing.T) {
	payload := triagePayloadFromFailures("organization", "integrations/terraform-provider-github", []engine.PersistFailure{{
		Package: "./github", Test: "TestAccThing", Status: "fail", Classification: engine.ClassificationFlakeConfirmed,
		Fingerprint: "sha256:" + strings.Repeat("b", 64), ShortFingerprint: strings.Repeat("b", 16), Class: engine.ClassRateLimit,
		Canonical: "403 rate limit", Retryable: true, Attempts: 2, Mode: "organization",
	}})
	if len(payload.Failures) != 1 {
		t.Fatalf("failures = %v, want one", payload.Failures)
	}
	if payload.Failures[0].IssueAction != "none" {
		t.Fatalf("flake issue_action = %q, want none", payload.Failures[0].IssueAction)
	}
}

func TestShouldFileIssueUsesEngineEligibility(t *testing.T) {
	validA := "sha256:" + strings.Repeat("a", 64)
	cases := []engine.PersistFailure{
		{Classification: engine.ClassificationReal, Fingerprint: validA, ShortFingerprint: strings.Repeat("a", 16)},
		{Classification: engine.ClassificationRealUnstable, Fingerprint: validA, ShortFingerprint: strings.Repeat("a", 16)},
		{Classification: engine.ClassificationFlakeConfirmed, Fingerprint: validA, ShortFingerprint: strings.Repeat("a", 16)},
		{Classification: engine.ClassificationReal, Fingerprint: "", ShortFingerprint: ""},
		{Classification: engine.ClassificationReal, Fingerprint: "sha256:" + strings.Repeat("a", 63), ShortFingerprint: strings.Repeat("a", 16)},
		{Classification: engine.ClassificationReal, Fingerprint: "sha256:" + strings.Repeat("g", 64), ShortFingerprint: strings.Repeat("g", 16)},
		{Classification: engine.ClassificationReal, Fingerprint: validA, ShortFingerprint: strings.Repeat("b", 16)},
		{Classification: engine.ClassificationReal, Fingerprint: validA, ShortFingerprint: strings.Repeat("a", 16), KnownIssue: 99},
		{Classification: engine.ClassificationReal, Fingerprint: validA, ShortFingerprint: strings.Repeat("a", 16), IssueAction: "known"},
	}
	for _, tc := range cases {
		if got, want := shouldFileIssue(tc), engine.EligibleForIssueFiling(tc); got != want {
			t.Fatalf("shouldFileIssue(%+v) = %v, want engine eligibility %v", tc, got, want)
		}
	}
}

func TestTriageFileIssuesRequiresConfirmInNonInteractiveMode(t *testing.T) {
	root := seedTriageFailureState(t, "ghp_FAKEFAKEFAKEFAKEFAKEFAKE")
	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"triage", "--repo-root", root, "--file-issues"}, &out, &errOut, deps{})
	if code != 2 {
		t.Fatalf("exit = %d, want 2; stderr=%s", code, errOut.String())
	}
	if !strings.Contains(errOut.String(), "--file-issues requires --confirm-file-issues") {
		t.Fatalf("stderr missing confirmation gate: %s", errOut.String())
	}
}

func TestTriageOfflineRejectsConfirmedIssueFiling(t *testing.T) {
	root := seedTriageFailureState(t, "ghp_FAKEFAKEFAKEFAKEFAKEFAKE")
	var out, errOut bytes.Buffer
	d := deps{
		provider: triageFakeProvider(),
		newIssueFiler: func(string) (issueFiler, error) {
			t.Fatal("offline issue filing must not create a GitHub issue client")
			return nil, nil
		},
	}
	code := runWithDeps([]string{"triage", "--repo-root", root, "--known-issues-offline", "--file-issues", "--confirm-file-issues"}, &out, &errOut, d)
	if code != 2 {
		t.Fatalf("exit = %d, want usage error 2; stderr=%s", code, errOut.String())
	}
	if !strings.Contains(errOut.String(), "--known-issues-offline cannot be combined with --file-issues") {
		t.Fatalf("stderr missing offline filing rejection: %s", errOut.String())
	}
}

func TestTriageFilesIssueWithBothFlagsUsingFakeClient(t *testing.T) {
	secret := "ghp_FAKEFAKEFAKEFAKEFAKEFAKE"
	root := seedTriageFailureState(t, secret)
	filer := &fakeIssueFiler{created: &ghissues.IssueResult{Number: 55, URL: "https://github.com/integrations/terraform-provider-github/issues/55"}}
	d := deps{provider: triageFakeProvider(), newIssueFiler: func(repo string) (issueFiler, error) {
		if repo != defaultIssuesRepo {
			t.Fatalf("repo = %q, want default", repo)
		}
		return filer, nil
	}}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"triage", "--repo-root", root, "--file-issues", "--confirm-file-issues", "--json"}, &out, &errOut, d)
	if code != 0 {
		t.Fatalf("exit = %d, want 0; stderr=%s", code, errOut.String())
	}
	if filer.createCalls != 1 || filer.dedupCalls != 1 {
		t.Fatalf("dedup/create calls = %d/%d, want 1/1", filer.dedupCalls, filer.createCalls)
	}
	if strings.Contains(filer.lastDraft.Title+filer.lastDraft.Body+out.String(), secret) {
		t.Fatalf("secret leaked through draft or JSON:\n%s\n%s", filer.lastDraft.Body, out.String())
	}
	var got triagePayload
	if err := json.Unmarshal(out.Bytes(), &got); err != nil {
		t.Fatalf("triage output is not JSON: %v\n%s", err, out.String())
	}
	if got.Failures[0].IssueAction != "filed" {
		t.Fatalf("issue_action = %q, want filed", got.Failures[0].IssueAction)
	}
	if got.Failures[0].KnownIssue == nil || *got.Failures[0].KnownIssue != 55 {
		t.Fatalf("known_issue = %v, want 55", got.Failures[0].KnownIssue)
	}
}

func TestTriageDedupsBeforeCreate(t *testing.T) {
	root := seedTriageFailureState(t, "ghp_FAKEFAKEFAKEFAKEFAKEFAKE")
	filer := &fakeIssueFiler{dedup: &ghissues.IssueMatch{Number: 77, State: "open", URL: "https://github.com/integrations/terraform-provider-github/issues/77"}}
	d := deps{provider: triageFakeProvider(), newIssueFiler: func(string) (issueFiler, error) { return filer, nil }}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"triage", "--repo-root", root, "--file-issues", "--confirm-file-issues", "--json"}, &out, &errOut, d)
	if code != 0 {
		t.Fatalf("exit = %d, want 0; stderr=%s", code, errOut.String())
	}
	if filer.createCalls != 0 {
		t.Fatalf("create calls = %d, want 0 after dedup", filer.createCalls)
	}
	var got triagePayload
	if err := json.Unmarshal(out.Bytes(), &got); err != nil {
		t.Fatalf("triage output is not JSON: %v\n%s", err, out.String())
	}
	if got.Failures[0].IssueAction != "dedup" {
		t.Fatalf("issue_action = %q, want dedup", got.Failures[0].IssueAction)
	}
}

func TestTriageNotesClosedDedupWithoutCreate(t *testing.T) {
	root := seedTriageFailureState(t, "ghp_FAKEFAKEFAKEFAKEFAKEFAKE")
	filer := &fakeIssueFiler{dedup: &ghissues.IssueMatch{Number: 77, State: "closed", URL: "https://github.com/integrations/terraform-provider-github/issues/77"}}
	d := deps{provider: triageFakeProvider(), newIssueFiler: func(string) (issueFiler, error) { return filer, nil }}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"triage", "--repo-root", root, "--file-issues", "--confirm-file-issues"}, &out, &errOut, d)
	if code != 0 {
		t.Fatalf("exit = %d, want 0; stderr=%s", code, errOut.String())
	}
	if filer.createCalls != 0 {
		t.Fatalf("create calls = %d, want 0 for closed dedup", filer.createCalls)
	}
	if !strings.Contains(out.String(), "matched closed issue #77") {
		t.Fatalf("stdout missing closed issue note:\n%s", out.String())
	}
}

func TestTriageDryRunDedupsClosedIssueWithLiveFiler(t *testing.T) {
	root := seedTriageFailureState(t, "ghp_FAKEFAKEFAKEFAKEFAKEFAKE")
	filer := &fakeIssueFiler{dedup: &ghissues.IssueMatch{Number: 77, State: "closed", URL: "https://github.com/integrations/terraform-provider-github/issues/77"}}
	d := deps{
		provider:      triageFakeProvider(),
		newIssueFiler: func(string) (issueFiler, error) { return filer, nil },
		getenv:        getenvFromMap(map[string]string{"GITHUB_TOKEN": "ghp_FAKEFAKEFAKEFAKEFAKEFAKE"}),
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"triage", "--repo-root", root, "--json"}, &out, &errOut, d)
	if code != 0 {
		t.Fatalf("exit = %d, want 0; stderr=%s", code, errOut.String())
	}
	if filer.dedupCalls != 1 || filer.createCalls != 0 {
		t.Fatalf("dedup/create calls = %d/%d, want 1/0", filer.dedupCalls, filer.createCalls)
	}
	var got triagePayload
	if err := json.Unmarshal(out.Bytes(), &got); err != nil {
		t.Fatalf("triage output is not JSON: %v\n%s", err, out.String())
	}
	if got.Failures[0].IssueAction != "dedup" || got.Failures[0].KnownIssue == nil || *got.Failures[0].KnownIssue != 77 {
		t.Fatalf("closed dedup JSON = %+v, want dedup issue 77", got.Failures[0])
	}
	loaded, err := engine.Load(filepath.Join(root, ".pulsar-state.json"))
	if err != nil {
		t.Fatal(err)
	}
	wantNote := "matched closed issue #77; file with --file-issues --confirm-file-issues --reopen-known-issue or open a new issue manually"
	if len(loaded.Failures) != 1 || !slices.Contains(loaded.Failures[0].Reasons, wantNote) {
		t.Fatalf("state reasons = %+v, want closed issue note", loaded.Failures)
	}
}

func TestTriageDryRunDedupsOpenIssueWithLiveFiler(t *testing.T) {
	root := seedTriageFailureState(t, "ghp_FAKEFAKEFAKEFAKEFAKEFAKE")
	filer := &fakeIssueFiler{dedup: &ghissues.IssueMatch{Number: 88, State: "open", URL: "https://github.com/integrations/terraform-provider-github/issues/88"}}
	d := deps{
		provider:      triageFakeProvider(),
		newIssueFiler: func(string) (issueFiler, error) { return filer, nil },
		getenv:        getenvFromMap(map[string]string{"GITHUB_TOKEN": "ghp_FAKEFAKEFAKEFAKEFAKEFAKE"}),
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"triage", "--repo-root", root}, &out, &errOut, d)
	if code != 0 {
		t.Fatalf("exit = %d, want 0; stderr=%s", code, errOut.String())
	}
	if filer.dedupCalls != 1 || filer.createCalls != 0 {
		t.Fatalf("dedup/create calls = %d/%d, want 1/0", filer.dedupCalls, filer.createCalls)
	}
	if !strings.Contains(out.String(), "dedup matched issue #88") {
		t.Fatalf("stdout missing open dedup:\n%s", out.String())
	}
	for _, notWant := range []string{"would file issue", "issue dry-run title:"} {
		if strings.Contains(out.String(), notWant) {
			t.Fatalf("stdout should not contain %q after open dedup:\n%s", notWant, out.String())
		}
	}
}

func TestTriageOfflineDryRunKeepsWouldFilePreview(t *testing.T) {
	root := seedTriageFailureState(t, "ghp_FAKEFAKEFAKEFAKEFAKEFAKE")
	filer := &fakeIssueFiler{dedup: &ghissues.IssueMatch{Number: 88, State: "open"}}
	d := deps{provider: triageFakeProvider(), newIssueFiler: func(string) (issueFiler, error) { return filer, nil }}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"triage", "--repo-root", root, "--known-issues-offline"}, &out, &errOut, d)
	if code != 0 {
		t.Fatalf("exit = %d, want 0; stderr=%s", code, errOut.String())
	}
	if filer.dedupCalls != 0 || filer.createCalls != 0 {
		t.Fatalf("dedup/create calls = %d/%d, want 0/0 offline", filer.dedupCalls, filer.createCalls)
	}
	for _, want := range []string{"would file issue", "issue dry-run title:"} {
		if !strings.Contains(out.String(), want) {
			t.Fatalf("stdout missing %q in offline dry-run:\n%s", want, out.String())
		}
	}
}

func TestTriageRefreshRevalidatesPersistedKnownIssue(t *testing.T) {
	root := seedTriageFailureState(t, "ghp_FAKEFAKEFAKEFAKEFAKEFAKE")
	statePath := filepath.Join(root, ".pulsar-state.json")
	st, err := engine.Load(statePath)
	if err != nil {
		t.Fatal(err)
	}
	st.Failures[0].KnownIssue = 42
	st.Failures[0].IssueAction = "known"
	if err := st.Save(statePath); err != nil {
		t.Fatal(err)
	}

	registry := &fakeIssueRegistry{}
	filer := &fakeIssueFiler{}
	d := deps{
		newIssueRegistry: func(triageOptions) (issueRegistry, error) { return registry, nil },
		newIssueFiler:    func(string) (issueFiler, error) { return filer, nil },
		getenv:           getenvFromMap(map[string]string{"GITHUB_TOKEN": "ghp_FAKEFAKEFAKEFAKEFAKEFAKE"}),
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"triage", "--repo-root", root, "--json"}, &out, &errOut, d)
	if code != 0 {
		t.Fatalf("exit = %d, want 0; stderr=%s", code, errOut.String())
	}
	if registry.calls != 1 || filer.dedupCalls != 1 || filer.createCalls != 0 {
		t.Fatalf("registry/dedup/create calls = %d/%d/%d, want 1/1/0", registry.calls, filer.dedupCalls, filer.createCalls)
	}
	var payload triagePayload
	if err := json.Unmarshal(out.Bytes(), &payload); err != nil {
		t.Fatalf("triage output is not JSON: %v\n%s", err, out.String())
	}
	if payload.Failures[0].KnownIssue != nil || payload.Failures[0].IssueAction != "dry-run" {
		t.Fatalf("refreshed failure = %+v, want eligible dry-run without known issue", payload.Failures[0])
	}
	loaded, err := engine.Load(statePath)
	if err != nil {
		t.Fatal(err)
	}
	if loaded.Failures[0].KnownIssue != 0 || loaded.Failures[0].IssueAction != "" {
		t.Fatalf("persisted failure = %+v, want stale known annotation cleared", loaded.Failures[0])
	}
}

func TestTriageRefreshPreservesTerminalIssueActions(t *testing.T) {
	for _, action := range []string{"filed", "dedup"} {
		t.Run(action, func(t *testing.T) {
			root := seedTriageFailureState(t, "ghp_FAKEFAKEFAKEFAKEFAKEFAKE")
			statePath := filepath.Join(root, ".pulsar-state.json")
			st, err := engine.Load(statePath)
			if err != nil {
				t.Fatal(err)
			}
			st.Failures[0].KnownIssue = 42
			st.Failures[0].IssueAction = action
			if err := st.Save(statePath); err != nil {
				t.Fatal(err)
			}
			registry := &fakeIssueRegistry{match: engine.KnownIssueMatch{
				Found: true, Issue: 99, Mode: "known-real", Suppress: true,
			}}
			d := deps{newIssueRegistry: func(triageOptions) (issueRegistry, error) { return registry, nil }}

			var out, errOut bytes.Buffer
			code := runWithDeps([]string{"triage", "--repo-root", root, "--json"}, &out, &errOut, d)
			if code != 0 {
				t.Fatalf("exit = %d, want 0; stderr=%s", code, errOut.String())
			}
			if registry.calls != 0 {
				t.Fatalf("registry calls = %d, want no revalidation for terminal %q action", registry.calls, action)
			}
			loaded, err := engine.Load(statePath)
			if err != nil {
				t.Fatal(err)
			}
			if loaded.Failures[0].KnownIssue != 42 || loaded.Failures[0].IssueAction != action {
				t.Fatalf("persisted failure = %+v, want terminal %q issue #42 unchanged", loaded.Failures[0], action)
			}
		})
	}
}

func TestTriageLocalIgnoreKnownIssueSuppressesFiling(t *testing.T) {
	root := seedTriageFailureState(t, "ghp_FAKEFAKEFAKEFAKEFAKEFAKE")
	statePath := filepath.Join(root, ".pulsar-state.json")
	st, err := engine.Load(statePath)
	if err != nil {
		t.Fatal(err)
	}
	st.Failures[0].KnownIssue = 42
	st.Failures[0].IssueAction = "known"
	if err := st.Save(statePath); err != nil {
		t.Fatal(err)
	}
	registry := &fakeIssueRegistry{match: engine.KnownIssueMatch{Found: true, Mode: "ignore", Suppress: true, Note: "local suppress"}}
	filer := &fakeIssueFiler{created: &ghissues.IssueResult{Number: 99, URL: "https://github.com/integrations/terraform-provider-github/issues/99"}}
	d := deps{
		provider:         triageFakeProvider(),
		newIssueRegistry: func(triageOptions) (issueRegistry, error) { return registry, nil },
		newIssueFiler:    func(string) (issueFiler, error) { return filer, nil },
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"triage", "--repo-root", root, "--file-issues", "--confirm-file-issues", "--json"}, &out, &errOut, d)
	if code != 0 {
		t.Fatalf("exit = %d, want 0; stderr=%s", code, errOut.String())
	}
	if filer.createCalls != 0 || filer.dedupCalls != 0 {
		t.Fatalf("dedup/create calls = %d/%d, want 0/0 for local suppress", filer.dedupCalls, filer.createCalls)
	}
	var got triagePayload
	if err := json.Unmarshal(out.Bytes(), &got); err != nil {
		t.Fatalf("triage output is not JSON: %v\n%s", err, out.String())
	}
	if got.Failures[0].KnownIssue != nil || got.Failures[0].IssueAction != "known" {
		t.Fatalf("suppressed failure fields = %+v, want known action without issue number", got.Failures[0])
	}
}

func TestTriageDryRunPrintsRedactedIssueDraft(t *testing.T) {
	secret := "ghp_FAKEFAKEFAKEFAKEFAKEFAKE"
	root := seedTriageFailureState(t, secret)
	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"triage", "--repo-root", root, "--known-issues-offline"}, &out, &errOut, deps{provider: triageFakeProvider()})
	if code != 0 {
		t.Fatalf("exit = %d, want 0; stderr=%s", code, errOut.String())
	}
	if strings.Contains(out.String(), secret) {
		t.Fatalf("dry-run output leaked token-shaped value:\n%s", out.String())
	}
	for _, want := range []string{"issue dry-run title:", "issue dry-run labels:", "issue dry-run body:", "<!-- pulsar-known-issue", "***REDACTED***"} {
		if !strings.Contains(out.String(), want) {
			t.Fatalf("dry-run output missing %q:\n%s", want, out.String())
		}
	}
	for _, want := range []string{"terraform-provider-tester orphans", "terraform-provider-tester sweep --confirm", "terraform-provider-tester run "} {
		if !strings.Contains(out.String(), want) {
			t.Fatalf("dry-run output missing current command %q:\n%s", want, out.String())
		}
	}
	for _, notWant := range []string{"pulsar orphans", "pulsar sweep", "pulsar run"} {
		if strings.Contains(out.String(), notWant) {
			t.Fatalf("dry-run output contains obsolete command %q:\n%s", notWant, out.String())
		}
	}
}

func TestTriageDryRunRedactsProviderSecretFromRawFailureLog(t *testing.T) {
	const (
		secretKey = "FAKE_PROVIDER_SECRET"
		secret    = "ENV_SECRET_VALUE_789"
	)
	root := cliScratchDir(t)
	st := engine.State{
		Provider: "github",
		Mode:     "organization",
		Results: []engine.PersistResult{{
			Package: "./github",
			Test:    "TestAccThing",
			Status:  "fail",
		}},
	}
	if err := st.Save(filepath.Join(root, ".pulsar-state.json")); err != nil {
		t.Fatal(err)
	}
	failDir := filepath.Join(root, ".pulsar-failures")
	if err := os.MkdirAll(failDir, 0o700); err != nil {
		t.Fatal(err)
	}
	logPath := engine.FailureLogPath(failDir, "./github", "TestAccThing")
	if err := os.WriteFile(logPath, []byte("assertion failed with "+secret+"\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	d := deps{
		provider: &fakeprovider.Fake{NameVal: "github", Secrets: []string{secretKey}},
		getenv:   getenvFromMap(map[string]string{secretKey: secret}),
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"triage", "--repo-root", root, "--known-issues-offline"}, &out, &errOut, d)
	if code != 0 {
		t.Fatalf("exit = %d, want 0; stderr=%s", code, errOut.String())
	}
	if strings.Contains(out.String(), secret) {
		t.Fatalf("dry-run output leaked provider secret from raw failure log:\n%s", out.String())
	}
	if !strings.Contains(out.String(), "***REDACTED***") {
		t.Fatalf("dry-run output missing redaction marker:\n%s", out.String())
	}
}

func seedTriageFailureState(t *testing.T, secret string) string {
	t.Helper()
	root := cliScratchDir(t)
	st := engine.State{Provider: "github", Mode: "organization", Failures: []engine.PersistFailure{{
		Package: "./github", Test: "TestAccThing", Status: "fail", Fingerprint: "sha256:" + strings.Repeat("a", 64), ShortFingerprint: strings.Repeat("a", 16),
		Class: engine.ClassLeftoverState, Canonical: "422 leftover state " + secret, Retryable: true, Classification: engine.ClassificationReal, Attempts: 3, Mode: "organization",
	}}}
	if err := st.Save(filepath.Join(root, ".pulsar-state.json")); err != nil {
		t.Fatal(err)
	}
	failDir := filepath.Join(root, ".pulsar-failures")
	if err := os.MkdirAll(failDir, 0o700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(failDir, "github_TestAccThing.log"), []byte("failure log "+secret+"\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	return root
}

type fakeIssueFiler struct {
	dedup       *ghissues.IssueMatch
	created     *ghissues.IssueResult
	dedupCalls  int
	createCalls int
	lastDraft   ghissues.IssueDraft
}

func (f *fakeIssueFiler) FindIssueByFingerprint(_ context.Context, _ string) (*ghissues.IssueMatch, error) {
	f.dedupCalls++
	return f.dedup, nil
}

func (f *fakeIssueFiler) CreateFailureIssue(_ context.Context, draft ghissues.IssueDraft) (*ghissues.IssueResult, error) {
	f.createCalls++
	f.lastDraft = draft
	return f.created, nil
}

type fakeIssueRegistry struct {
	match engine.KnownIssueMatch
	calls int
}

func (f *fakeIssueRegistry) LookupFingerprint(_ context.Context, _ string, _ string) (engine.KnownIssueMatch, error) {
	f.calls++
	return f.match, nil
}
