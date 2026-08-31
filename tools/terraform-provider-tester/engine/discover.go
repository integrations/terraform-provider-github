package engine

import (
	"bufio"
	"context"
	"fmt"
	"os/exec"
	"strings"
)

// ListTests runs "go test -list <pattern> <packages...>" with cwd set to dir
// and returns the top-level test names. On a non-zero exit it returns a
// descriptive error wrapping the combined output so the operator sees the
// compiler text.
func ListTests(ctx context.Context, dir string, packages []string, pattern string) ([]string, error) {
	args := append([]string{"test", "-list", pattern}, packages...)
	cmd := exec.CommandContext(ctx, "go", args...)
	cmd.Dir = dir
	out, err := cmd.CombinedOutput()
	return interpretListResult(string(out), err)
}

// interpretListResult is the pure, exec-free core of ListTests. Given the
// combined stdout/stderr and the exec error (nil on success), it returns the
// top-level test names or an error. A non-nil runErr (e.g. non-zero exit from
// a build failure) is turned into a descriptive error so the operator sees the
// compiler text; it is never silently reported as "0 tests found".
func interpretListResult(output string, runErr error) ([]string, error) {
	if runErr != nil {
		return nil, fmt.Errorf("go test -list failed: %w\n%s", runErr, output)
	}
	return parseListOutput(output), nil
}

// parseListOutput extracts top-level test names from the output of
// "go test -list". It accepts multi-package output (an "ok" line follows each
// package's names) and strips every non-name line.
func parseListOutput(output string) []string {
	var names []string
	sc := bufio.NewScanner(strings.NewReader(output))
	for sc.Scan() {
		line := sc.Text()
		if isTestName(line) {
			names = append(names, line)
		}
	}
	return names
}

// isTestName reports whether a line from "go test -list" output is a
// top-level test name. According to the Go spec, test names begin with
// Test, Benchmark, Example, or Fuzz and have no leading whitespace.
// Everything else (ok lines, FAIL lines, # lines, blank lines, indented
// context) is dropped.
func isTestName(line string) bool {
	if line == "" {
		return false
	}
	// Any leading whitespace means it is not a top-level name.
	if line[0] == ' ' || line[0] == '\t' {
		return false
	}
	return strings.HasPrefix(line, "Test") ||
		strings.HasPrefix(line, "Benchmark") ||
		strings.HasPrefix(line, "Example") ||
		strings.HasPrefix(line, "Fuzz")
}
