package engine

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/github/terraform-provider-tester/internal/redact"
	"github.com/github/terraform-provider-tester/provider"
)

func TestWriteFailuresWritesRedactedPerFailedTest(t *testing.T) {
	dir := t.TempDir()
	red := redact.New([]string{"ghp_SECRET"})
	res := RunResult{Tests: []TestResult{
		{
			Package: "github.com/integrations/terraform-provider-github/github",
			Name:    "TestAccPass",
			Status:  provider.StatusPass,
			Output:  []string{"passing ghp_SECRET\n"},
		},
		{
			Package: "github.com/integrations/terraform-provider-github/github",
			Name:    "TestAccFail/é",
			Status:  provider.StatusFail,
			Output:  []string{"failing ghp_SECRET\n"},
		},
	}}

	logs, err := WriteFailures(dir, res, red)
	if err != nil {
		t.Fatalf("WriteFailures: %v", err)
	}

	wantPath := filepath.Join(dir, "github_TestAccFail__.log")
	if len(logs) != 1 {
		t.Fatalf("len(logs) = %d; want 1 (%v)", len(logs), logs)
	}
	if logs[0].Test != "TestAccFail/é" || logs[0].Path != wantPath {
		t.Fatalf("logs[0] = %+v; want Test=TestAccFail/é Path=%s", logs[0], wantPath)
	}

	content := readFileString(t, wantPath)
	if !strings.Contains(content, "failing ***REDACTED***\n") {
		t.Fatalf("failure log missing redacted failure output: %q", content)
	}
	if strings.Contains(content, "ghp_SECRET") {
		t.Fatalf("failure log contains unredacted secret: %q", content)
	}
	if _, err := os.Stat(filepath.Join(dir, "github_TestAccPass.log")); !os.IsNotExist(err) {
		t.Fatalf("passing test log exists or stat failed unexpectedly: %v", err)
	}
}

func TestWriteFailuresAggregatesFailedSubtests(t *testing.T) {
	dir := t.TempDir()
	res := RunResult{Tests: []TestResult{
		{
			Package: "github.com/integrations/terraform-provider-github/github",
			Name:    "TestAccX",
			Status:  provider.StatusFail,
			Output:  []string{"parent\n"},
		},
		{
			Package: "github.com/integrations/terraform-provider-github/github",
			Name:    "TestAccX",
			Sub:     "case_a",
			Status:  provider.StatusFail,
			Output:  []string{"detail A\n"},
		},
		{
			Package: "github.com/integrations/terraform-provider-github/github",
			Name:    "TestAccX",
			Sub:     "case_b",
			Status:  provider.StatusPass,
			Output:  []string{"noise B\n"},
		},
	}}

	logs, err := WriteFailures(dir, res, nil)
	if err != nil {
		t.Fatalf("WriteFailures: %v", err)
	}
	if len(logs) != 1 {
		t.Fatalf("len(logs) = %d; want 1", len(logs))
	}

	content := readFileString(t, filepath.Join(dir, "github_TestAccX.log"))
	if !strings.Contains(content, "detail A") {
		t.Fatalf("failure log missing failed subtest output: %q", content)
	}
	if strings.Contains(content, "noise B") {
		t.Fatalf("failure log contains passed subtest output: %q", content)
	}
}

func TestWriteFailuresClearsStaleOnGreenRun(t *testing.T) {
	dir := t.TempDir()
	const pkg = "github.com/integrations/terraform-provider-github/github"
	stale := filepath.Join(dir, sanitizeName(pkg, "TestAccPass")+".log")
	if err := os.WriteFile(stale, []byte("old\n"), 0o600); err != nil {
		t.Fatalf("pre-create stale log: %v", err)
	}
	res := RunResult{Tests: []TestResult{{
		Package: pkg,
		Name:    "TestAccPass",
		Status:  provider.StatusPass,
		Output:  []string{"ok\n"},
	}}}

	logs, err := WriteFailures(dir, res, nil)
	if err != nil {
		t.Fatalf("WriteFailures: %v", err)
	}
	if logs != nil {
		t.Fatalf("logs = %#v; want nil", logs)
	}
	if _, err := os.Stat(stale); !os.IsNotExist(err) {
		t.Fatalf("stale log exists or stat failed unexpectedly: %v", err)
	}
	matches, err := filepath.Glob(filepath.Join(dir, "*.log"))
	if err != nil {
		t.Fatalf("Glob: %v", err)
	}
	if len(matches) != 0 {
		t.Fatalf("unexpected logs after green run: %v", matches)
	}
}

func TestWriteFailuresPreservesLogsForTestsNotReRun(t *testing.T) {
	dir := t.TempDir()
	const pkg = "github.com/integrations/terraform-provider-github/github"
	otherLog := filepath.Join(dir, sanitizeName(pkg, "TestAccOther")+".log")
	if err := os.WriteFile(otherLog, []byte("previous failure\n"), 0o600); err != nil {
		t.Fatalf("pre-create other failure log: %v", err)
	}
	res := RunResult{Tests: []TestResult{{
		Package: pkg,
		Name:    "TestAccFail",
		Status:  provider.StatusFail,
		Output:  []string{"new failure\n"},
	}}}

	logs, err := WriteFailures(dir, res, nil)
	if err != nil {
		t.Fatalf("WriteFailures: %v", err)
	}
	if len(logs) != 1 {
		t.Fatalf("len(logs) = %d; want 1", len(logs))
	}
	if content := readFileString(t, otherLog); content != "previous failure\n" {
		t.Fatalf("other failure log = %q; want preserved previous failure", content)
	}
	failLog := filepath.Join(dir, sanitizeName(pkg, "TestAccFail")+".log")
	if content := readFileString(t, failLog); !strings.Contains(content, "new failure\n") {
		t.Fatalf("new failure log missing output: %q", content)
	}
}

func TestWriteFailuresIncludesPanicAndTimeout(t *testing.T) {
	dir := t.TempDir()
	res := RunResult{Tests: []TestResult{
		{
			Package: "github.com/integrations/terraform-provider-github/github",
			Name:    "TestAccPanic",
			Status:  provider.StatusPanic,
			Output:  []string{"panic\n"},
		},
		{
			Package: "github.com/integrations/terraform-provider-github/github",
			Name:    "TestAccTimeout",
			Status:  provider.StatusTimeout,
			Output:  []string{"timeout\n"},
		},
	}}

	logs, err := WriteFailures(dir, res, nil)
	if err != nil {
		t.Fatalf("WriteFailures: %v", err)
	}
	if len(logs) != 2 {
		t.Fatalf("len(logs) = %d; want 2 (%v)", len(logs), logs)
	}
	for _, name := range []string{"github_TestAccPanic.log", "github_TestAccTimeout.log"} {
		if _, err := os.Stat(filepath.Join(dir, name)); err != nil {
			t.Fatalf("expected %s to exist: %v", name, err)
		}
	}
}

func TestWriteFailuresNilRedactorNoPanic(t *testing.T) {
	dir := t.TempDir()
	res := RunResult{Tests: []TestResult{{
		Package: "github.com/integrations/terraform-provider-github/github",
		Name:    "TestAccFail",
		Status:  provider.StatusFail,
		Output:  []string{"raw ghp_SECRET\n"},
	}}}

	logs, err := WriteFailures(dir, res, nil)
	if err != nil {
		t.Fatalf("WriteFailures: %v", err)
	}
	if len(logs) != 1 {
		t.Fatalf("len(logs) = %d; want 1", len(logs))
	}

	content := readFileString(t, filepath.Join(dir, "github_TestAccFail.log"))
	if !strings.Contains(content, "raw ghp_SECRET\n") {
		t.Fatalf("failure log missing unredacted output with nil redactor: %q", content)
	}
}

func TestFailureLogPathRoundTrip(t *testing.T) {
	dir := t.TempDir()
	const (
		pkg  = "github.com/integrations/terraform-provider-github/github"
		test = "TestAccX"
	)

	want := FailureLogPath(dir, pkg, test)

	res := RunResult{Tests: []TestResult{{
		Package: pkg,
		Name:    test,
		Status:  provider.StatusFail,
		Output:  []string{"fail\n"},
	}}}
	logs, err := WriteFailures(dir, res, nil)
	if err != nil {
		t.Fatalf("WriteFailures: %v", err)
	}
	if len(logs) != 1 {
		t.Fatalf("len(logs) = %d; want 1", len(logs))
	}
	if logs[0].Path != want {
		t.Fatalf("logs[0].Path = %q; FailureLogPath = %q; mismatch", logs[0].Path, want)
	}
	if _, err := os.Stat(want); err != nil {
		t.Fatalf("FailureLogPath file does not exist: %v", err)
	}
}

func TestWriteFailuresLongNameCapped(t *testing.T) {
	dir := t.TempDir()
	const pkg = "github.com/integrations/terraform-provider-github/github"

	longName := strings.Repeat("A", 320)
	res := RunResult{Tests: []TestResult{{
		Package: pkg,
		Name:    longName,
		Status:  provider.StatusFail,
		Output:  []string{"fail\n"},
	}}}
	logs, err := WriteFailures(dir, res, nil)
	if err != nil {
		t.Fatalf("WriteFailures with 320-char name: %v", err)
	}
	if len(logs) != 1 {
		t.Fatalf("len(logs) = %d; want 1", len(logs))
	}
	if _, err := os.Stat(logs[0].Path); err != nil {
		t.Fatalf("log file not readable: %v", err)
	}
	base := filepath.Base(logs[0].Path)
	if len(base) > 255 {
		t.Fatalf("basename length %d > 255: %q", len(base), base)
	}

	longName2 := strings.Repeat("B", 320)
	res2 := RunResult{Tests: []TestResult{{
		Package: pkg,
		Name:    longName2,
		Status:  provider.StatusFail,
		Output:  []string{"fail2\n"},
	}}}
	logs2, err := WriteFailures(dir, res2, nil)
	if err != nil {
		t.Fatalf("WriteFailures with second 320-char name: %v", err)
	}
	if len(logs2) != 1 {
		t.Fatalf("len(logs2) = %d; want 1", len(logs2))
	}
	if logs[0].Path == logs2[0].Path {
		t.Fatalf("different long names produced same path: %q", logs[0].Path)
	}
}

func readFileString(t *testing.T, path string) string {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("ReadFile(%s): %v", path, err)
	}
	return string(data)
}

func TestWriteFailuresRedactsSplitPEMOutput(t *testing.T) {
	bodyA := strings.Repeat("N", 64)
	bodyB := strings.Repeat("O", 64)
	dir := t.TempDir()
	res := RunResult{Tests: []TestResult{{
		Package: "./github",
		Name:    "TestAccSecret",
		Status:  provider.StatusFail,
		Output: []string{
			fakePEMBeginLine(),
			bodyA + "\n",
			bodyB + "\n",
			fakePEMEndLine(),
		},
	}}}
	logs, err := WriteFailures(dir, res, redact.New(nil))
	if err != nil {
		t.Fatalf("WriteFailures: %v", err)
	}
	if len(logs) != 1 {
		t.Fatalf("logs = %v, want one", logs)
	}
	content, err := os.ReadFile(logs[0].Path)
	if err != nil {
		t.Fatal(err)
	}
	got := string(content)
	if strings.Contains(got, bodyA) || strings.Contains(got, bodyB) {
		t.Fatalf("failure log leaked split PEM body: %q", got)
	}
	if !strings.Contains(got, "***REDACTED***") {
		t.Fatalf("failure log missing redaction marker: %q", got)
	}
}
