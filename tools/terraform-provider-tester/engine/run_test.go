package engine

import (
	"context"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/github/terraform-provider-tester/provider"
)

// TestRunSpecArgsNoShell asserts the exact argv slice produced by RunSpec.Args():
// starts with "test", contains packages, -json, -run <pattern as single element>,
// -timeout, and -count=1. Injection guard: even a pattern with shell metacharacters
// must appear as a single argv element after -run.
func TestRunSpecArgsNoShell(t *testing.T) {
	spec := RunSpec{
		Dir:      "/repo",
		Packages: []string{"./github/..."},
		Pattern:  "^TestAccOrg",
		Timeout:  30 * time.Minute,
	}
	args := spec.Args()

	if len(args) == 0 || args[0] != "test" {
		t.Errorf("args[0] = %q, want \"test\"", func() string {
			if len(args) > 0 {
				return args[0]
			}
			return "<empty>"
		}())
	}

	hasJSON, hasCount, hasPkg, hasRun, hasTimeout := false, false, false, false, false
	for i, a := range args {
		switch {
		case a == "-json":
			hasJSON = true
		case a == "-count=1":
			hasCount = true
		case a == "./github/...":
			hasPkg = true
		case a == "-run" && i+1 < len(args) && args[i+1] == "^TestAccOrg":
			hasRun = true
		case a == "-timeout":
			hasTimeout = true
		}
	}
	if !hasJSON {
		t.Error("args missing -json")
	}
	if !hasCount {
		t.Error("args missing -count=1")
	}
	if !hasPkg {
		t.Error("args missing package ./github/...")
	}
	if !hasRun {
		t.Error("args missing -run ^TestAccOrg as a single element")
	}
	if !hasTimeout {
		t.Error("args missing -timeout")
	}

	// Exact full-slice assertion: catches ordering and duplication regressions.
	// (30 * time.Minute).String() == "30m0s"
	want := []string{"test", "./github/...", "-json", "-run", "^TestAccOrg", "-timeout", "30m0s", "-count=1"}
	if !reflect.DeepEqual(args, want) {
		t.Errorf("Args() = %v, want %v", args, want)
	}

	// Injection guard: a Pattern containing shell metacharacters must be a
	// single argv element (never split across multiple args).
	inject := RunSpec{
		Packages: []string{"./github/..."},
		Pattern:  "^Test; rm -rf /",
		Timeout:  30 * time.Minute,
	}
	iArgs := inject.Args()
	foundSingle := false
	for i, a := range iArgs {
		if a == "-run" && i+1 < len(iArgs) {
			if iArgs[i+1] == "^Test; rm -rf /" {
				foundSingle = true
			}
		}
	}
	if !foundSingle {
		t.Error("injection pattern not passed as single argv element after -run")
	}
}

// TestRunSpecCommandPrinted checks that Command() produces a human-readable
// line containing "go test", the packages, and key flags, with NO env-var
// values embedded (secrets safe).
func TestRunSpecCommandPrinted(t *testing.T) {
	spec := RunSpec{
		Dir:      "/repo",
		Packages: []string{"./github/..."},
		Pattern:  "^TestAccOrg",
		Timeout:  30 * time.Minute,
	}
	cmd := spec.Command()

	if !strings.Contains(cmd, "go test") {
		t.Errorf("Command() = %q does not contain \"go test\"", cmd)
	}
	if !strings.Contains(cmd, "./github/...") {
		t.Errorf("Command() = %q does not contain packages", cmd)
	}
	if !strings.Contains(cmd, "-json") {
		t.Errorf("Command() = %q does not contain \"-json\"", cmd)
	}
	// Env-var names/values must NOT appear in the command string.
	for _, bad := range []string{"TF_ACC", "CGO_ENABLED", "GITHUB_TOKEN"} {
		if strings.Contains(cmd, bad) {
			t.Errorf("Command() = %q embeds env var %q", cmd, bad)
		}
	}
}

// TestRunnerRefusesEmptyPattern asserts that Run returns ErrNoTestsSelected
// immediately when spec.Pattern is empty, executing nothing.
func TestRunnerRefusesEmptyPattern(t *testing.T) {
	// Poison PATH so that any attempt to exec "go" would fail with a
	// go-not-found error instead of ErrNoTestsSelected. Also point Dir at a
	// path that does not exist, so any exec attempt would fail for THAT reason
	// too. Getting exactly ErrNoTestsSelected proves the guard short-circuited
	// before any subprocess was launched.
	t.Setenv("PATH", "")
	r := &Runner{}
	spec := RunSpec{
		Dir:     filepath.Join(t.TempDir(), "does-not-exist"),
		Pattern: "", // intentionally empty
	}
	_, err := r.Run(context.Background(), spec, func(TestResult) {})
	if !errors.Is(err, ErrNoTestsSelected) {
		t.Errorf("empty Pattern: got %v, want ErrNoTestsSelected", err)
	}
}

// TestRunnerStripsTFLogByDefault verifies that buildChildEnv strips all TF_LOG*
// variables from the child env when AllowSensitiveLogs is false, and forwards
// them when AllowSensitiveLogs is true.
func TestRunnerStripsTFLogByDefault(t *testing.T) {
	t.Setenv("TF_LOG", "DEBUG")
	t.Setenv("TF_LOG_PATH", "/home/user/tf.log")
	t.Setenv("TF_LOG_CORE", "INFO")
	t.Setenv("TF_LOG_PROVIDER", "TRACE")

	// Default: strip.
	env := buildChildEnv(RunSpec{AllowSensitiveLogs: false})
	for _, kv := range env {
		key := kv
		if idx := strings.IndexByte(kv, '='); idx >= 0 {
			key = kv[:idx]
		}
		if strings.HasPrefix(key, "TF_LOG") {
			t.Errorf("child env (strip) contains %q", kv)
		}
	}

	// AllowSensitiveLogs=true: ALL TF_LOG* vars must be forwarded.
	envAllow := buildChildEnv(RunSpec{AllowSensitiveLogs: true})
	forwardedKeys := make(map[string]bool)
	for _, kv := range envAllow {
		if idx := strings.IndexByte(kv, '='); idx >= 0 {
			forwardedKeys[kv[:idx]] = true
		}
	}
	for _, key := range []string{"TF_LOG", "TF_LOG_PATH", "TF_LOG_CORE", "TF_LOG_PROVIDER"} {
		if !forwardedKeys[key] {
			t.Errorf("AllowSensitiveLogs=true: %s not forwarded to child env", key)
		}
	}
}

// TestRunnerEnvPrecedence verifies that ExtraEnv overrides a parent-process
// variable (no duplicate entries) and that mandatory vars appear exactly once.
func TestRunnerEnvPrecedence(t *testing.T) {
	t.Setenv("GITHUB_OWNER", "parent-org")

	env := buildChildEnv(RunSpec{ExtraEnv: []string{"GITHUB_OWNER=extra-org"}})

	// collectValues returns every "value" for entries matching "key=" in env.
	collectValues := func(key string) []string {
		prefix := key + "="
		var vals []string
		for _, kv := range env {
			if strings.HasPrefix(kv, prefix) {
				vals = append(vals, strings.TrimPrefix(kv, prefix))
			}
		}
		return vals
	}

	// (a) Exactly one GITHUB_OWNER entry, and ExtraEnv value wins.
	ownerVals := collectValues("GITHUB_OWNER")
	if len(ownerVals) != 1 {
		t.Errorf("GITHUB_OWNER: want exactly 1 entry, got %d (%v)", len(ownerVals), ownerVals)
	} else if ownerVals[0] != "extra-org" {
		t.Errorf("GITHUB_OWNER value = %q, want \"extra-org\" (ExtraEnv must win)", ownerVals[0])
	}

	// (b) TF_ACC=1 and CGO_ENABLED=0 each present exactly once.
	for _, check := range []struct{ key, val string }{
		{"TF_ACC", "1"},
		{"CGO_ENABLED", "0"},
	} {
		vals := collectValues(check.key)
		if len(vals) != 1 {
			t.Errorf("%s: want exactly 1 entry, got %d (%v)", check.key, len(vals), vals)
		} else if vals[0] != check.val {
			t.Errorf("%s value = %q, want %q", check.key, vals[0], check.val)
		}
	}
}

// TestFindRepoRootFromSubdir covers three cases:
//  1. .git is a directory (normal clone).
//  2. .git is a regular file (git worktree - the real case for this repo).
//  3. No .git anywhere → error.
func TestFindRepoRootFromSubdir(t *testing.T) {
	t.Run("git_dir", func(t *testing.T) {
		root := t.TempDir()
		if err := os.Mkdir(filepath.Join(root, ".git"), 0755); err != nil {
			t.Fatal(err)
		}
		sub := filepath.Join(root, "a", "b", "c")
		if err := os.MkdirAll(sub, 0755); err != nil {
			t.Fatal(err)
		}
		got, err := FindRepoRoot(sub)
		if err != nil {
			t.Fatalf("FindRepoRoot (git dir): %v", err)
		}
		if got != root {
			t.Errorf("FindRepoRoot (git dir) = %q, want %q", got, root)
		}
	})

	t.Run("git_file", func(t *testing.T) {
		// This is the worktree case: .git is a plain text file like
		// "gitdir: /parent/.git/worktrees/foo"
		root := t.TempDir()
		if err := os.WriteFile(filepath.Join(root, ".git"),
			[]byte("gitdir: /parent/.git/worktrees/foo\n"), 0644); err != nil {
			t.Fatal(err)
		}
		sub := filepath.Join(root, "a", "b")
		if err := os.MkdirAll(sub, 0755); err != nil {
			t.Fatal(err)
		}
		got, err := FindRepoRoot(sub)
		if err != nil {
			t.Fatalf("FindRepoRoot (git file): %v", err)
		}
		if got != root {
			t.Errorf("FindRepoRoot (git file) = %q, want %q", got, root)
		}
	})

	t.Run("no_git", func(t *testing.T) {
		root := t.TempDir()
		_, err := FindRepoRoot(root)
		if err == nil {
			t.Error("FindRepoRoot with no .git: expected error, got nil")
		}
	})
}

// TestRunnerStreamsThroughParse runs `go test -json` against the minimod
// fixture (one passing test, one failing test) and verifies that Parse is
// called live and that the RunResult contains both tests with correct statuses.
// This is the only test in the engine suite that invokes `go`; it requires no
// network access and no TF_ACC.
func TestRunnerStreamsThroughParse(t *testing.T) {
	minimod := filepath.Join("testdata", "minimod")

	r := &Runner{}
	spec := RunSpec{
		Dir:      minimod,
		Packages: []string{"./..."},
		Pattern:  "TestAlways",
		Timeout:  60 * time.Second,
	}

	var results []TestResult
	sink := func(tr TestResult) {
		results = append(results, tr)
	}

	// Run returns nil even when tests fail (failures appear in results).
	if _, err := r.Run(context.Background(), spec, sink); err != nil {
		t.Fatalf("Run: unexpected error: %v", err)
	}

	var passFound, failFound bool
	for _, tr := range results {
		switch tr.Name {
		case "TestAlwaysPasses":
			if tr.Status != provider.StatusPass {
				t.Errorf("TestAlwaysPasses status = %v, want StatusPass", tr.Status)
			}
			passFound = true
		case "TestAlwaysFails":
			if tr.Status != provider.StatusFail {
				t.Errorf("TestAlwaysFails status = %v, want StatusFail", tr.Status)
			}
			failFound = true
		}
	}
	if !passFound {
		t.Errorf("TestAlwaysPasses not found in results (got %d results)", len(results))
	}
	if !failFound {
		t.Errorf("TestAlwaysFails not found in results (got %d results)", len(results))
	}
}

// TestRunnerGoNotFound forces PATH to a nonexistent directory so that
// exec.LookPath("go") fails, then asserts Run returns a typed "not found"
// error wrapping exec.ErrNotFound - not a panic.
func TestRunnerGoNotFound(t *testing.T) {
	t.Setenv("PATH", "/nonexistent_path_for_testing")

	r := &Runner{}
	spec := RunSpec{
		Dir:      ".",
		Packages: []string{"./..."},
		Pattern:  "^TestSomething",
		Timeout:  10 * time.Second,
	}
	_, err := r.Run(context.Background(), spec, func(TestResult) {})
	if err == nil {
		t.Fatal("expected error when go not found on PATH, got nil")
	}
	if !errors.Is(err, exec.ErrNotFound) {
		t.Errorf("expected exec.ErrNotFound in error chain; got: %v", err)
	}
}

// TestRunnerContextCancel cancels a context mid-run (during a test that sleeps
// 60 s) and asserts that Run returns promptly with context.Canceled.
func TestRunnerContextCancel(t *testing.T) {
	minimod := filepath.Join("testdata", "minimod")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	r := &Runner{}
	spec := RunSpec{
		Dir:      minimod,
		Packages: []string{"./..."},
		Pattern:  "TestLongRunning",
		Timeout:  120 * time.Second,
	}

	// Cancel after a brief delay so the subprocess has time to start.
	go func() {
		time.Sleep(300 * time.Millisecond)
		cancel()
	}()

	start := time.Now()
	_, err := r.Run(ctx, spec, func(TestResult) {})
	elapsed := time.Since(start)

	if err == nil {
		t.Fatal("expected error from cancelled context, got nil")
	}
	if !errors.Is(err, context.Canceled) {
		t.Errorf("expected context.Canceled; got: %v", err)
	}
	if elapsed > 10*time.Second {
		t.Errorf("Run took %v after cancel; expected < 10s (timing concern on slow CI)", elapsed)
	}
}
