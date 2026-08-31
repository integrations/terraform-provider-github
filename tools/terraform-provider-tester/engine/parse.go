package engine

import (
	"bufio"
	"encoding/json"
	"io"
	"strings"

	"github.com/github/terraform-provider-tester/provider"
)

// Event mirrors the test2json event structure.
// Actions: start|run|pause|cont|pass|fail|skip|output|build-output|build-fail
type Event struct {
	Action      string
	Package     string
	ImportPath  string // set on build-output/build-fail (e.g. "pkg [pkg.test]")
	Test        string // "" for package-level; "Parent/Sub" for subtests
	Elapsed     float64
	Output      string
	FailedBuild string // set on a package fail event when the cause was a build failure
}

// TestResult is the result of one test or subtest.
type TestResult struct {
	Package string
	Name    string // top-level parent name
	Sub     string // "" for top-level; "Sub" or "Sub/Sub2" for subtests
	Status  provider.Status
	Elapsed float64
	Output  []string // raw output lines
}

// RunResult is the roll-up for an entire go test run.
type RunResult struct {
	BuildFailed  bool
	BuildOutput  []string
	PreRunFailed bool
	PreRunOutput []string
	Tests        []TestResult
	RawLog       []string
	Elapsed      float64
}

// pkgState tracks per-package parsing state.
type pkgState struct {
	runSeen     bool
	buildFail   bool            // saw a build-fail action
	output      []string        // package-level output lines
	runOrder    []string        // test names in run order (for orphan flush)
	hasTerminal map[string]bool // tests that received a terminal event
}

// Parse streams r (a go test -json stream), invokes sink for each completed
// TestResult (eager writes), and returns the roll-up. It never returns a
// non-nil error for malformed lines; those go to RawLog. It classifies build
// vs. pre-run vs. normal failures.
func Parse(r io.Reader, sink func(TestResult)) (RunResult, error) {
	var result RunResult

	pkgs := make(map[string]*pkgState)
	testOutputs := make(map[string][]string) // testKey -> output lines
	buildOutputByImport := make(map[string][]string)
	buildOutputByPkg := make(map[string][]string)

	getPkg := func(pkg string) *pkgState {
		if pkgs[pkg] == nil {
			pkgs[pkg] = &pkgState{hasTerminal: make(map[string]bool)}
		}
		return pkgs[pkg]
	}

	scanner := bufio.NewScanner(r)
	scanner.Buffer(make([]byte, 1<<20), 1<<20)

	for scanner.Scan() {
		line := scanner.Text()
		var ev Event
		if err := json.Unmarshal([]byte(line), &ev); err != nil {
			result.RawLog = append(result.RawLog, line)
			continue
		}

		switch ev.Action {
		case "build-output":
			buildOutputByImport[ev.ImportPath] = append(buildOutputByImport[ev.ImportPath], ev.Output)
			base := importPathBase(ev.ImportPath)
			buildOutputByPkg[base] = append(buildOutputByPkg[base], ev.Output)

		case "build-fail":
			base := importPathBase(ev.ImportPath)
			getPkg(base).buildFail = true

		case "start":
			getPkg(ev.Package) // ensure state exists

		case "run":
			ps := getPkg(ev.Package)
			ps.runSeen = true
			ps.runOrder = append(ps.runOrder, ev.Test)

		case "output":
			if ev.Test == "" {
				ps := getPkg(ev.Package)
				ps.output = append(ps.output, ev.Output)
			} else {
				key := testKey(ev.Package, ev.Test)
				testOutputs[key] = append(testOutputs[key], ev.Output)
			}

		case "pass", "fail", "skip":
			if ev.Test == "" {
				// Package-level terminal event.
				ps := getPkg(ev.Package)
				result.Elapsed += ev.Elapsed

				// Bug 1 fix: only classify build/prerun failure on "fail".
				// A package "pass" or "skip" with zero run events is normal
				// (e.g. "no tests to run") and must not set any failure flag.
				if ev.Action == "fail" && !ps.runSeen {
					isBuild := ps.buildFail ||
						ev.FailedBuild != "" ||
						outputContainsBuildMarker(ps.output)
					if isBuild {
						result.BuildFailed = true
						if ev.FailedBuild != "" {
							result.BuildOutput = append(result.BuildOutput, buildOutputByImport[ev.FailedBuild]...)
						} else {
							result.BuildOutput = append(result.BuildOutput, buildOutputByPkg[ev.Package]...)
						}
						for _, o := range ps.output {
							if strings.Contains(o, "[build failed]") || strings.Contains(o, "[setup failed]") {
								result.BuildOutput = append(result.BuildOutput, o)
							}
						}
					} else {
						result.PreRunFailed = true
						result.PreRunOutput = append(result.PreRunOutput, ps.output...)
					}
				}

				// Bug 2 fix: synthesize results for in-flight tests that never
				// received a terminal event (e.g. tests orphaned by a timeout).
				if ev.Action == "fail" {
					for _, testName := range ps.runOrder {
						if ps.hasTerminal[testName] {
							continue
						}
						key := testKey(ev.Package, testName)
						output := testOutputs[key]
						status := synthesizeOrphanStatus(output, ps.output)
						name, sub := splitTest(testName)
						tr := TestResult{
							Package: ev.Package,
							Name:    name,
							Sub:     sub,
							Status:  status,
							Output:  output,
						}
						result.Tests = append(result.Tests, tr)
						if sink != nil {
							sink(tr)
						}
					}
				}
			} else {
				// Test-level terminal event.
				ps := getPkg(ev.Package)
				ps.hasTerminal[ev.Test] = true
				key := testKey(ev.Package, ev.Test)
				output := testOutputs[key]
				status := terminalStatus(ev.Action, output)
				name, sub := splitTest(ev.Test)
				tr := TestResult{
					Package: ev.Package,
					Name:    name,
					Sub:     sub,
					Status:  status,
					Elapsed: ev.Elapsed,
					Output:  output,
				}
				result.Tests = append(result.Tests, tr)
				if sink != nil {
					sink(tr)
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return result, err
	}

	return result, nil
}

// importPathBase extracts the base package name from an ImportPath like
// "pkg/foo [pkg/foo.test]" -> "pkg/foo".
func importPathBase(ip string) string {
	if i := strings.Index(ip, " ["); i >= 0 {
		return ip[:i]
	}
	return ip
}

// testKey returns a unique key for a test within its package.
func testKey(pkg, test string) string {
	return pkg + "\x00" + test
}

// outputContainsBuildMarker reports whether any line contains a build marker.
func outputContainsBuildMarker(lines []string) bool {
	for _, l := range lines {
		if strings.Contains(l, "[build failed]") || strings.Contains(l, "[setup failed]") {
			return true
		}
	}
	return false
}

// terminalStatus maps an action and test output to a provider.Status.
// Timeout is checked before panic because Go uses a panic internally for timeout.
func terminalStatus(action string, output []string) provider.Status {
	switch action {
	case "pass":
		return provider.StatusPass
	case "skip":
		return provider.StatusSkip
	case "fail":
		for _, line := range output {
			if strings.Contains(line, "test timed out") || strings.Contains(line, "*** Test killed") {
				return provider.StatusTimeout
			}
		}
		for _, line := range output {
			if strings.Contains(line, "panic:") {
				return provider.StatusPanic
			}
		}
		return provider.StatusFail
	}
	return provider.StatusFail
}

// synthesizeOrphanStatus determines the status for a test that never received
// a terminal event (orphaned by a timeout or hard kill). It checks the test's
// own output first, then the package-level output.
func synthesizeOrphanStatus(testOutput, pkgOutput []string) provider.Status {
	for _, line := range testOutput {
		if strings.Contains(line, "test timed out") || strings.Contains(line, "*** Test killed") {
			return provider.StatusTimeout
		}
	}
	for _, line := range pkgOutput {
		if strings.Contains(line, "test timed out") || strings.Contains(line, "*** Test killed") {
			return provider.StatusTimeout
		}
	}
	for _, line := range testOutput {
		if strings.Contains(line, "panic:") {
			return provider.StatusPanic
		}
	}
	return provider.StatusFail
}

// splitTest splits "Parent/Sub/Sub2" into ("Parent", "Sub/Sub2").
// Returns (name, "") for top-level tests.
func splitTest(test string) (name, sub string) {
	if i := strings.Index(test, "/"); i >= 0 {
		return test[:i], test[i+1:]
	}
	return test, ""
}
