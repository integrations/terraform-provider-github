package engine

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

// ErrNoTestsSelected is returned by Run when spec.Pattern is empty, preventing
// an empty retry/resume set from degrading into `go test -run ""` which matches ALL tests.
var ErrNoTestsSelected = errors.New("no tests selected")

// RunSpec holds everything needed to execute a `go test` run.
type RunSpec struct {
	Dir                string        // provider repo root (from FindRepoRoot or --repo-root)
	Packages           []string      // e.g. ["./github/..."]
	Pattern            string        // -run regex; MUST be non-empty
	Timeout            time.Duration // default 120m when zero
	ExtraEnv           []string      // mode env from the adapter, e.g. ["GITHUB_OWNER=org"]
	AllowSensitiveLogs bool          // forward TF_LOG* to the child; default false
}

// Args builds the shell-free argv for go test:
// ["test", pkgs..., "-json", "-run", pattern, "-timeout", duration, "-count=1"]
// Pattern is always a single element (injection guard: never passed through sh -c).
func (s RunSpec) Args() []string {
	dur := s.Timeout
	if dur == 0 {
		dur = 120 * time.Minute
	}
	args := []string{"test"}
	args = append(args, s.Packages...)
	args = append(args, "-json", "-run", s.Pattern, "-timeout", dur.String(), "-count=1")
	return args
}

// Command renders a human-readable command line safe to print (no env values,
// so secrets are never embedded).
func (s RunSpec) Command() string {
	return "go " + strings.Join(s.Args(), " ")
}

// FindRepoRoot walks up from start until it finds a directory containing .git
// (either a directory for a normal clone, or a regular file for a git worktree).
// Returns an error if the filesystem root is reached without finding .git.
func FindRepoRoot(start string) (string, error) {
	dir, err := filepath.Abs(start)
	if err != nil {
		return "", fmt.Errorf("abs path of %q: %w", start, err)
	}
	for {
		// Accept .git as file OR directory (os.Stat succeeds for both).
		if _, statErr := os.Stat(filepath.Join(dir, ".git")); statErr == nil {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			// Reached filesystem root.
			return "", fmt.Errorf("no git repository root found starting from %q", start)
		}
		dir = parent
	}
}

// isTFLogVar reports whether the given env-var key is a TF_LOG* variable that
// should be stripped from the child env by default.
func isTFLogVar(key string) bool {
	return strings.HasPrefix(key, "TF_LOG")
}

// buildChildEnv constructs the environment slice for the child `go test` process:
//  1. Start from the parent process env.
//  2. Strip TF_LOG* variables unless AllowSensitiveLogs is true.
//  3. Remove any existing TF_ACC, CGO_ENABLED, or keys defined in ExtraEnv
//     so the explicitly set values below always win.
//  4. Append TF_ACC=1 and CGO_ENABLED=0.
//  5. Append spec.ExtraEnv (overrides win because duplicates come last and most
//     tools use the last matching entry on their side, but we also strip earlier).
func buildChildEnv(spec RunSpec) []string {
	// Build a set of keys that ExtraEnv (plus our mandatory overrides) will
	// supply so we can strip them from the parent env to avoid duplicates.
	overrideKeys := map[string]bool{
		"TF_ACC":      true,
		"CGO_ENABLED": true,
	}
	for _, kv := range spec.ExtraEnv {
		if idx := strings.IndexByte(kv, '='); idx >= 0 {
			overrideKeys[kv[:idx]] = true
		}
	}

	parent := os.Environ()
	env := make([]string, 0, len(parent)+2+len(spec.ExtraEnv))

	for _, kv := range parent {
		key := kv
		if idx := strings.IndexByte(kv, '='); idx >= 0 {
			key = kv[:idx]
		}
		// Strip TF_LOG* unless caller opted in.
		if !spec.AllowSensitiveLogs && isTFLogVar(key) {
			continue
		}
		// Strip keys that our explicit overrides will supply.
		if overrideKeys[key] {
			continue
		}
		env = append(env, kv)
	}

	// Mandatory: always on, always last so they cannot be shadowed by parent env.
	env = append(env, "TF_ACC=1", "CGO_ENABLED=0")
	env = append(env, spec.ExtraEnv...)
	return env
}

// Runner executes `go test` runs and streams results through Parse.
type Runner struct {
	Stdout io.Writer // receives the printed command line; may be nil
}

// Run executes `go test` with a curated child env (TF_ACC=1, CGO_ENABLED=0,
// plus ExtraEnv), streams stdout through Parse, calls sink per result, and
// returns the rollup. It:
//   - Refuses an empty spec.Pattern with ErrNoTestsSelected.
//   - Returns a typed exec.ErrNotFound error if `go` is not on PATH.
//   - Honours ctx cancellation: kills the child and returns ctx.Err().
//   - Treats a non-zero exit from `go test` (caused by test failures) as
//     normal; failures are captured in the RunResult, not returned as an error.
func (r *Runner) Run(ctx context.Context, spec RunSpec, sink func(TestResult)) (RunResult, error) {
	if spec.Pattern == "" {
		return RunResult{}, ErrNoTestsSelected
	}

	args := spec.Args()
	cmd := exec.CommandContext(ctx, "go", args...)
	cmd.Dir = spec.Dir
	cmd.Env = buildChildEnv(spec)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return RunResult{}, fmt.Errorf("creating stdout pipe: %w", err)
	}

	if r.Stdout != nil {
		fmt.Fprintln(r.Stdout, spec.Command())
	}

	if err := cmd.Start(); err != nil {
		if errors.Is(err, exec.ErrNotFound) {
			return RunResult{}, fmt.Errorf("go not found on PATH: %w", err)
		}
		return RunResult{}, fmt.Errorf("starting go test: %w", err)
	}

	// Parse streams stdout live, calling sink for each completed TestResult.
	result, parseErr := Parse(stdout, sink)

	// Wait must be called after all reads from the pipe are done.
	_ = cmd.Wait()

	// Context cancellation takes priority: the child was killed because ctx
	// was done, so return the context error rather than any parse/exit error.
	if ctx.Err() != nil {
		return RunResult{}, ctx.Err()
	}

	if parseErr != nil {
		return result, parseErr
	}

	// Non-zero exit (test failures) is reflected in result.Tests, not as an error.
	return result, nil
}
