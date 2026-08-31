package cli

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/mattn/go-isatty"

	"github.com/github/terraform-provider-tester/engine"
	"github.com/github/terraform-provider-tester/internal/redact"
	"github.com/github/terraform-provider-tester/provider"
	githubprovider "github.com/github/terraform-provider-tester/provider/github"
	"github.com/github/terraform-provider-tester/tui"
)

// version is overridable at build time via
// -ldflags "-X github.com/github/terraform-provider-tester/cli.version=<v>".
// It is a var (not a const) so release builds can stamp git describe output.
var version = "v0.1.0-dev"

// testRunner is an injectable seam for engine.Runner in tests.
type testRunner interface {
	Run(ctx context.Context, spec engine.RunSpec, sink func(engine.TestResult)) (engine.RunResult, error)
}

// lister is an injectable seam for engine.ListTests in tests.
type lister func(ctx context.Context, dir string, packages []string, pattern string) ([]string, error)

// deps holds injectable dependencies so CLI logic is testable without real
// network calls or go test execution.
type deps struct {
	provider            provider.Provider
	newRunner           func(io.Writer) testRunner
	list                lister
	newKnownIssueLister func(string) (knownIssueLister, error)
	newIssueRegistry    func(triageOptions) (issueRegistry, error)
	newIssueFiler       func(string) (issueFiler, error)
	// cwd returns the working directory; injectable for tests.
	cwd func() (string, error)
	// userConfigDir returns the user's config directory; injectable for tests.
	userConfigDir func() (string, error)
	// userHomeDir returns the user's home directory; injectable for tests.
	userHomeDir func() (string, error)
	// getenv reads an environment variable; injectable for tests.
	getenv func(string) string
	// setenv writes an environment variable; injectable for tests.
	setenv func(string, string) error
	// isTTY reports whether stdout is a terminal. nil is treated as non-TTY.
	isTTY func() bool
}

func defaultDeps() deps {
	return deps{
		provider:            githubprovider.New(),
		newRunner:           func(w io.Writer) testRunner { return &engine.Runner{Stdout: w} },
		list:                engine.ListTests,
		newKnownIssueLister: func(repo string) (knownIssueLister, error) { return githubprovider.NewIssueClientFromEnv(repo) },
		newIssueRegistry:    buildLiveIssueRegistry,
		newIssueFiler: func(repo string) (issueFiler, error) {
			return githubprovider.NewIssueClientFromEnv(repo)
		},
		cwd:           os.Getwd,
		userConfigDir: os.UserConfigDir,
		userHomeDir:   os.UserHomeDir,
		getenv:        os.Getenv,
		setenv:        os.Setenv,
		isTTY:         func() bool { return isatty.IsTerminal(os.Stdout.Fd()) },
	}
}

// Run is the process entrypoint. Returns the exit code.
func Run(args []string, out, errOut io.Writer) int {
	return runWithDeps(args, out, errOut, defaultDeps())
}

func runWithDeps(args []string, out, errOut io.Writer, d deps) int {
	// Pre-scan and strip global --env-file flag; peek at --mode for scoping.
	envFilePath, hadExplicitEnvFile, prescanMode, args := prescanArgs(args)
	if len(args) > 0 && args[0] == "e2e" && prescanMode == "" {
		prescanMode = e2eDefaultMode
	}
	baseGetenv := getenvFunc(d)
	env := newEnvOverlay(baseGetenv, setenvFunc(d))
	getenv := env.Get
	setenv := env.Set

	if hadExplicitEnvFile && envFilePath == "" {
		fmt.Fprintln(errOut, "env-file: --env-file requires a path")
		return 2
	}

	// Auto-detect .pulsar.env if no explicit flag.
	if !hadExplicitEnvFile {
		envFilePath = resolveAutoEnvFile(d)
	}

	var envFileVars map[string]string
	if envFilePath != "" {
		f, openErr := os.Open(envFilePath)
		if openErr != nil {
			fmt.Fprintln(errOut, "env-file:", openErr)
			return 2
		}
		vars, parseErr := parseEnvFile(f)
		f.Close()
		if parseErr != nil {
			fmt.Fprintln(errOut, "env-file:", parseErr)
			return 2
		}
		envFileVars = vars
	}

	if envFilePath != "" {
		effectiveMode, err := resolveEnvFileMode(args, prescanMode, envFileVars, d, getenv)
		if err != nil {
			fmt.Fprintln(errOut, err)
			return 2
		}
		var knownModes []string
		if envFileUsesPulsarKeys(envFileVars) {
			knownModes = make([]string, 0, len(d.provider.Modes()))
			for _, md := range d.provider.Modes() {
				knownModes = append(knownModes, md.Name)
			}
		}
		if _, err := resolveEnvFile(envFileVars, effectiveMode, knownModes, getenv, setenv); err != nil {
			fmt.Fprintln(errOut, "env-file:", err)
			return 2
		}
	}
	d.getenv = getenv
	d.setenv = setenv

	// Persisted resume modes are provider-independent. Reject corrupt orphan
	// accounting before TTY setup or later execution touches the provider, but
	// only after env loading so any required credentials are already scoped.
	if exitCode, handled := rejectResumeOrphanModeMismatchBeforeProvider(args, out, errOut, d, getenv); handled {
		return exitCode
	}

	if len(args) == 0 {
		noTUI := getenv("PULSAR_NO_TUI") != ""
		forceTTY := getenv("PULSAR_FORCE_TTY") != ""
		// Guard: nil isTTY is treated as non-TTY so existing tests that build
		// deps{} without the field continue to work without panicking.
		ttyFn := d.isTTY
		if ttyFn == nil {
			ttyFn = func() bool { return false }
		}
		if decideTUI(ttyFn(), forceTTY, noTUI) {
			mode := getenv("GH_TEST_AUTH_MODE")
			if mode == "" {
				mode = "anonymous"
			}
			owner := getenv("GITHUB_OWNER")
			ascii := getenv("NO_COLOR") != ""
			root, err := d.resolveRoot("")
			if err != nil {
				fmt.Fprintln(errOut, "resolving repo root:", err)
				return 1
			}
			// A bare invocation has no --allow-unclassified flag to honor,
			// so dashboard planning fails closed on an unclassified test
			// exactly like every non-TUI planning path.
			return runDashboard(d, root, mode, owner, ascii, false, out, errOut)
		}
		printHelp(out)
		return 0
	}

	if args[0] == "version" || args[0] == "--version" {
		fmt.Fprintf(out, "%s %s (terraform-provider-tester)\n", tui.AppName, version)
		fmt.Fprintf(out, "%s\n", tui.Description)
		fmt.Fprintf(out, "(c) %s · %s License\n", tui.Vendor, tui.License)
		fmt.Fprintf(out, "Maintainers: see %s\n", tui.MaintainersRef)
		return 0
	}

	subcmd := args[0]
	rest := args[1:]

	switch subcmd {
	case "e2e":
		return runE2E(rest, out, errOut, d)
	case "preflight":
		return runPreflight(rest, out, errOut, d)
	case "groups":
		return runGroups(rest, out, errOut, d)
	case "run":
		return runRun(rest, out, errOut, d)
	case "retry":
		return runRetry(rest, out, errOut, d)
	case "resume":
		return runResume(rest, out, errOut, d)
	case "report":
		return runReport(rest, out, errOut, d)
	case "triage":
		return runTriage(rest, out, errOut, d)
	case "known-issues":
		return runKnownIssues(rest, out, errOut, d)
	case "orphans":
		return runOrphans(rest, out, errOut, d)
	case "sweep":
		return runSweep(rest, out, errOut, d)
	case "unlock":
		return runUnlock(rest, out, errOut, d)
	case "discover":
		return runDiscover(rest, out, errOut, d)
	default:
		fmt.Fprintf(errOut, "unknown subcommand %q\n", subcmd)
		printHelp(errOut)
		return 2
	}
}

func getenvFunc(d deps) func(string) string {
	if d.getenv != nil {
		return d.getenv
	}
	return func(string) string { return "" }
}

func setenvFunc(d deps) func(string, string) error {
	if d.setenv != nil {
		return d.setenv
	}
	return func(string, string) error { return nil }
}

func resolveAutoEnvFile(d deps) string {
	if path := getenvFunc(d)("PULSAR_ENV_FILE"); path != "" {
		return path
	}
	for _, candidate := range []func() (string, error){
		func() (string, error) {
			if d.cwd == nil {
				return "", fmt.Errorf("cwd unavailable")
			}
			dir, err := d.cwd()
			if err != nil {
				return "", err
			}
			return filepath.Join(dir, ".pulsar.env"), nil
		},
		func() (string, error) {
			if d.userConfigDir == nil {
				return "", fmt.Errorf("user config dir unavailable")
			}
			dir, err := d.userConfigDir()
			if err != nil {
				return "", err
			}
			return filepath.Join(dir, "terraform-provider-tester", ".pulsar.env"), nil
		},
		func() (string, error) {
			if d.userHomeDir == nil {
				return "", fmt.Errorf("user home dir unavailable")
			}
			dir, err := d.userHomeDir()
			if err != nil {
				return "", err
			}
			return filepath.Join(dir, ".pulsar.env"), nil
		},
	} {
		path, err := candidate()
		if err != nil {
			continue
		}
		if _, statErr := os.Stat(path); statErr == nil {
			return path
		}
	}
	return ""
}

// resolveRoot returns override when non-empty, otherwise calls FindRepoRoot
// starting from d.cwd().
func (d deps) resolveRoot(override string) (string, error) {
	if override != "" {
		return override, nil
	}
	start, err := d.cwd()
	if err != nil {
		return "", fmt.Errorf("getting cwd: %w", err)
	}
	return engine.FindRepoRoot(start)
}

// prescanArgs pre-scans args for the global --env-file flag (stripped from the
// returned slice) and peeks at --mode (kept in the slice) for env-file scoping.
// Returns the env-file path (empty if absent), whether --env-file was explicit,
// the mode value (empty if absent), and the cleaned args slice.
func prescanArgs(args []string) (envFile string, hadExplicit bool, mode string, cleaned []string) {
	cleaned = make([]string, 0, len(args))
	for i := 0; i < len(args); i++ {
		arg := args[i]
		switch {
		case arg == "--env-file":
			if i+1 < len(args) {
				envFile = args[i+1]
				i++
			}
			hadExplicit = true
		case strings.HasPrefix(arg, "--env-file="):
			envFile = arg[len("--env-file="):]
			hadExplicit = true
		default:
			cleaned = append(cleaned, arg)
			if arg == "--mode" && i+1 < len(args) {
				mode = args[i+1]
			} else if strings.HasPrefix(arg, "--mode=") {
				mode = arg[len("--mode="):]
			}
		}
	}
	return
}

func scanArgValue(args []string, name string) (value string, set bool) {
	longFlag := "--" + name
	shortFlag := "-" + name
	for i := 0; i < len(args); i++ {
		arg := args[i]
		switch {
		case arg == longFlag || arg == shortFlag:
			if i+1 < len(args) {
				return args[i+1], true
			}
			return "", true
		case strings.HasPrefix(arg, longFlag+"="):
			return arg[len(longFlag)+1:], true
		case strings.HasPrefix(arg, shortFlag+"="):
			return arg[len(shortFlag)+1:], true
		}
	}
	return "", false
}

func loadStateForEnvMode(args []string, d deps) (engine.State, bool) {
	repoRoot, _ := scanArgValue(args, "repo-root")
	root, err := d.resolveRoot(repoRoot)
	if err != nil {
		return engine.State{}, false
	}
	state, err := engine.Load(filepath.Join(root, ".pulsar-state.json"))
	if err != nil {
		return engine.State{}, false
	}
	return state, true
}

func resolveEnvFileMode(args []string, explicitMode string, envFileVars map[string]string, d deps, getenv func(string) string) (string, error) {
	if explicitMode != "" {
		return explicitMode, nil
	}
	if len(args) > 0 && args[0] == "e2e" {
		return e2eDefaultMode, nil
	}
	if len(args) > 0 {
		if state, ok := loadStateForEnvMode(args, d); ok {
			switch args[0] {
			case "retry", "resume":
				mode, err := persistedRetryResumeMode(state)
				if err != nil {
					return "", err
				}
				if mode != "" {
					return mode, nil
				}
			case "orphans", "sweep":
				mode, err := persistedCleanupMode(state)
				if err != nil {
					return "", err
				}
				if mode != "" {
					return mode, nil
				}
			}
		}
	}
	if mode := getenv("GH_TEST_AUTH_MODE"); mode != "" {
		return mode, nil
	}
	if mode := envFileVars["GH_TEST_AUTH_MODE"]; mode != "" {
		return mode, nil
	}
	return "anonymous", nil
}

func envFileUsesPulsarKeys(vars map[string]string) bool {
	for key := range vars {
		if strings.HasPrefix(key, "PULSAR_") {
			return true
		}
	}
	return false
}

type resumeModeGuardOptions struct {
	repoRoot     string
	mode         string
	modeSet      bool
	format       outputFormat
	timeout      time.Duration
	optionsValid bool
}

func addResumeFlags(
	fs *flag.FlagSet,
	mode *string,
	timeout *time.Duration,
	repoRoot *string,
	allowSensitiveLogs *bool,
	format *outputFormat,
	tuiFlag, noTUIFlag *bool,
	triage *triageOptions,
) {
	addJSONFlag(fs, format)
	fs.StringVar(mode, "mode", "", "run mode (defaults to persisted plan mode)")
	fs.DurationVar(timeout, "timeout", 120*time.Minute, "test timeout")
	fs.StringVar(repoRoot, "repo-root", "", "repo root override")
	fs.BoolVar(allowSensitiveLogs, "allow-sensitive-logs", false, "forward TF_LOG* to child process")
	fs.BoolVar(tuiFlag, "tui", false, "open the interactive dashboard when possible")
	fs.BoolVar(noTUIFlag, "no-tui", false, "force non-interactive text output")
	addTriageRunFlags(fs, triage)
}

func scanResumeModeGuardOptions(args []string) (resumeModeGuardOptions, bool) {
	opts := resumeModeGuardOptions{
		format:       formatText,
		timeout:      120 * time.Minute,
		optionsValid: true,
	}
	if len(args) == 0 || args[0] != "resume" {
		return opts, false
	}

	fs := flag.NewFlagSet("resume-mode-guard", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	var (
		allowSensitiveLogs bool
		tuiFlag            bool
		noTUIFlag          bool
		triage             triageOptions
	)
	addResumeFlags(
		fs,
		&opts.mode,
		&opts.timeout,
		&opts.repoRoot,
		&allowSensitiveLogs,
		&opts.format,
		&tuiFlag,
		&noTUIFlag,
		&triage,
	)
	if err := fs.Parse(args[1:]); err != nil {
		opts.optionsValid = false
		return opts, true
	}
	fs.Visit(func(f *flag.Flag) {
		if f.Name == "mode" {
			opts.modeSet = true
		}
	})
	return opts, true
}

// printHelp writes usage to w.
func printHelp(w io.Writer) {
	fmt.Fprintf(w, "%s (terraform-provider-tester) - acceptance-test harness for terraform-provider-github\n", tui.AppName)
	fmt.Fprintln(w, "usage: terraform-provider-tester [--env-file <path>] <subcommand> [flags]")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "global flags:")
	fmt.Fprintln(w, "  --env-file <path>  load environment variables from file before running")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "subcommands:")
	fmt.Fprintln(w, "  e2e          run the guided acceptance-test workflow (anonymous, individual, or organization)")
	fmt.Fprintln(w, "  preflight    check provider configuration")
	fmt.Fprintln(w, "  groups       list test groups")
	fmt.Fprintln(w, "  run          run acceptance tests")
	fmt.Fprintln(w, "  retry        retry failed tests (--failed)")
	fmt.Fprintln(w, "  resume       resume incomplete run")
	fmt.Fprintln(w, "  report       generate test report")
	fmt.Fprintln(w, "  triage      classify last-run failures")
	fmt.Fprintln(w, "  known-issues list, sync, or add known failure fingerprints")
	fmt.Fprintln(w, "  orphans      list leaked resources")
	fmt.Fprintln(w, "  sweep        delete leaked resources (--confirm)")
	fmt.Fprintln(w, "  unlock       clear a stale run lock")
	fmt.Fprintln(w, "  discover     list accessible orgs, enterprises, and template repos")
	fmt.Fprintln(w, "  version      print version")
}

// printPreflight writes a human-readable checklist to w. Secret values are
// never printed - detail and fix text are passed through red before printing
// so a configured secret embedded in a Check's Detail or Fix (e.g. rejected
// by the provider's own API) never reaches text output, only the redaction
// marker.
func printPreflight(w io.Writer, report provider.PreflightReport, red *redact.Redactor) {
	for _, c := range report.Checks {
		sym := "✓"
		switch c.Status {
		case provider.CheckWarn:
			sym = "!"
		case provider.CheckFail:
			sym = "✗"
		}
		fmt.Fprintf(w, "[%s] %s: %s\n", sym, c.Name, red.String(c.Detail))
		if c.Fix != "" {
			fmt.Fprintf(w, "    fix: %s\n", red.String(c.Fix))
		}
	}
}

// exitFromResult returns 1 if res indicates any failure, 0 otherwise.
func exitFromResult(res engine.RunResult) int {
	if res.BuildFailed || res.PreRunFailed {
		return 1
	}
	for _, tr := range res.Tests {
		if tr.Sub != "" {
			continue
		}
		if tr.Status == provider.StatusFail ||
			tr.Status == provider.StatusPanic ||
			tr.Status == provider.StatusTimeout {
			return 1
		}
	}
	return 0
}

// buildExtraEnv converts provider EnvVar descriptors to KEY=VALUE strings by
// reading non-empty values from the resolved environment getter. Secret values
// are included so the child test process can authenticate; they are never printed.
//
// It always injects GH_TEST_AUTH_MODE=<mode> first so the provider's TestMain
// selects the same auth mode the operator chose. The engine appends ExtraEnv
// last when building the child environment, so this value overrides any stray
// GH_TEST_AUTH_MODE inherited from the parent process.
func buildExtraEnv(mode string, envVars []provider.EnvVar, getenv func(string) string) []string {
	extra := []string{"GH_TEST_AUTH_MODE=" + mode}
	for _, ev := range envVars {
		val := getenv(ev.Key)
		if val != "" {
			extra = append(extra, ev.Key+"="+val)
		}
	}
	return extra
}

func defaultMode(getenv func(string) string) string {
	if mode := getenv("GH_TEST_AUTH_MODE"); mode != "" {
		return mode
	}
	return "anonymous"
}

func redactorForProvider(prov provider.Provider, getenv func(string) string) *redact.Redactor {
	var secretVals []string
	for _, key := range prov.SecretEnvKeys() {
		if v := getenv(key); v != "" {
			secretVals = append(secretVals, v)
		}
	}
	return redact.New(secretVals)
}

func stdoutIsTTY(d deps) bool {
	if d.isTTY == nil {
		return false
	}
	return d.isTTY()
}

func shouldLaunchDashboard(format outputFormat, tuiFlag, noTUIFlag bool, d deps, errOut io.Writer) bool {
	if format == formatJSON || noTUIFlag {
		return false
	}
	isTTY := stdoutIsTTY(d)
	getenv := getenvFunc(d)
	forceTTY := getenv("PULSAR_FORCE_TTY") != ""
	if tuiFlag {
		if isTTY || forceTTY {
			return true
		}
		fmt.Fprintln(errOut, "--tui requested but stdout is not a TTY; falling back to text")
		return false
	}
	if getenv("PULSAR_NO_TUI") != "" {
		return false
	}
	return decideTUI(isTTY, forceTTY, false)
}

func runnerForFormat(d deps, format outputFormat, out io.Writer) testRunner {
	if format == formatJSON {
		return d.newRunner(io.Discard)
	}
	return d.newRunner(out)
}

func streamTextResult(w io.Writer, tr engine.TestResult, groupsByTest map[string]string, failDir string) {
	if tr.Sub != "" {
		return
	}
	fmt.Fprintf(w, "%s %s %s %.2fs", tr.Status.String(), groupForTest(groupsByTest, tr.Name), tr.Name, tr.Elapsed)
	if isFailureStatus(tr.Status) {
		fmt.Fprintf(w, " log=%s", engine.FailureLogPath(failDir, tr.Package, tr.Name))
	}
	fmt.Fprintln(w)
}

func countMatchingTests(names []string, re *regexp.Regexp) int {
	if re == nil {
		return len(names)
	}
	count := 0
	for _, name := range names {
		if re.MatchString(name) {
			count++
		}
	}
	return count
}

func emitEmptyRunJSON(out, errOut io.Writer, mode, statePath string, spec engine.RunSpec, opts runOptions) int {
	return emitEmptyRunJSONWithExit(out, errOut, mode, statePath, spec, opts, 0)
}

func emitEmptyRunJSONWithExit(out, errOut io.Writer, mode, statePath string, spec engine.RunSpec, opts runOptions, exitCode int) int {
	jw := newJSONWriter(out)
	cur := engine.State{Mode: mode}
	failDir := filepath.Join(filepath.Dir(statePath), ".pulsar-failures")
	if err := emitRunJSONSummary(jw, opts, spec, statePath, failDir, cur, engine.RunResult{}, exitCode); err != nil {
		fmt.Fprintln(errOut, "writing JSON:", err)
		return 1
	}
	return exitCode
}

// emitPreRunGateJSON writes a final run summary to stdout when a run is
// abandoned before any test executes because the provider preflight gate
// failed. Automation parsing the NDJSON stream still receives a terminal
// summary object (PreRunFailed=true, zero totals, exit 1) instead of an
// empty stream. Human-readable preflight detail goes to stderr separately.
func emitPreRunGateJSON(out, errOut io.Writer, root, mode string, prov provider.Provider) int {
	statePath := filepath.Join(root, ".pulsar-state.json")
	failDir := filepath.Join(root, ".pulsar-failures")
	jw := newJSONWriter(out)
	spec := engine.RunSpec{Dir: root, Packages: prov.TestPackages()}
	opts := runOptions{Format: formatJSON, Command: "run"}
	cur := engine.State{Mode: mode}
	res := engine.RunResult{PreRunFailed: true}
	if err := emitRunJSONSummary(jw, opts, spec, statePath, failDir, cur, res, 1); err != nil {
		fmt.Fprintln(errOut, "writing JSON:", err)
	}
	return 1
}

// emitPlanFailureRunJSON writes a terminal run summary to stdout when a run
// is abandoned during discovery or planning, before any plan exists. Without
// it, `run --json` returned an entirely empty stdout stream on a planning
// failure, leaving automation with nothing to parse. It deliberately emits
// only the summary: no plan or excluded events are fabricated for a plan that
// was never built. Callers pass the mapped planExitCode value and restrict
// this to runtime failures (1); usage errors (2) stay stderr-only.
func emitPlanFailureRunJSON(out, errOut io.Writer, root, mode string, prov provider.Provider, selection selectionInfo, exitCode int) {
	statePath := filepath.Join(root, ".pulsar-state.json")
	failDir := filepath.Join(root, ".pulsar-failures")
	jw := newJSONWriter(out)
	spec := engine.RunSpec{Dir: root, Packages: prov.TestPackages()}
	opts := runOptions{Format: formatJSON, Command: "run", Selection: selection}
	cur := engine.State{Mode: mode}
	if err := emitRunJSONSummary(jw, opts, spec, statePath, failDir, cur, engine.RunResult{}, exitCode); err != nil {
		fmt.Fprintln(errOut, "writing JSON:", err)
	}
}

func emitRunJSONSummary(w *jsonWriter, opts runOptions, spec engine.RunSpec, statePath, failDir string, cur engine.State, res engine.RunResult, exitCode int) error {
	summary := jsonRunSummary{
		SchemaVersion: 1,
		Type:          "summary",
		Command:       opts.Command,
		Mode:          cur.Mode,
		RepoRoot:      spec.Dir,
		StatePath:     statePath,
		FailureLogDir: failDir,
		GoTestCommand: spec.Command(),
		Selection:     opts.Selection,
		Totals:        totalsForTests(res.Tests, nil),
		Groups:        groupSummaries(opts.Groups, res.Tests),
		BuildFailed:   res.BuildFailed,
		PreRunFailed:  res.PreRunFailed,
		BuildLog:      stringPtr(cur.BuildLog),
		PreRunLog:     stringPtr(cur.PreRunLog),
		Failures:      failuresForSummary(res.Tests, opts.Groups, failDir),
		CleanupStatus: opts.CleanupStatus,
		ExitCode:      exitCode,
	}
	if opts.JSONSummaryObserver != nil {
		opts.JSONSummaryObserver(summary)
	}
	return w.Encode(summary)
}

// runWithSpec executes a RunSpec, streaming fresh results into a new State
// whose History is carried from prev. After the run it merges results: prior
// results for tests that produced no fresh result are preserved (see
// mergeResults), so filtered runs, retry, and resume keep previously-recorded
// results and a build/pre-run failure never wipes prior state. It saves under
// the caller-held lock, prints a summary, and returns the exit code.
func runWithSpec(ctx context.Context, spec engine.RunSpec, statePath string, prev engine.State, prov provider.Provider, mode string, runner testRunner, out, errOut io.Writer, opts runOptions) int {
	if opts.Triage.Enabled {
		return runWithSpecTriage(ctx, spec, statePath, prev, prov, mode, runner, out, errOut, opts)
	}
	getenv := opts.Getenv
	if getenv == nil {
		getenv = func(string) string { return "" }
	}
	cur := engine.State{
		Version:  engine.StateVersion,
		Provider: prov.Name(),
		Mode:     mode,
		RunAt:    time.Now(),
		History:  prev.History,
		Plan:     prev.Plan,
		Orphans:  prev.Orphans,
	}
	red := redactorForProvider(prov, getenv)
	failDir := filepath.Join(filepath.Dir(statePath), ".pulsar-failures")
	groupsByTest := testGroupMap(opts.Groups)
	var jw *jsonWriter
	var streamErr error
	if opts.Format == formatJSON {
		jw = newJSONWriter(out)
	}
	res, runErr := runner.Run(ctx, spec, func(tr engine.TestResult) {
		cur.Record(tr)
		if opts.Format == formatJSON {
			var failureLog *string
			if tr.Sub == "" && isFailureStatus(tr.Status) {
				failureLog = stringPtr(engine.FailureLogPath(failDir, tr.Package, tr.Name))
			}
			if encErr := jw.Encode(jsonTestEvent{
				SchemaVersion: 1,
				Type:          "test",
				Command:       opts.Command,
				Mode:          mode,
				Package:       tr.Package,
				Name:          tr.Name,
				Sub:           tr.Sub,
				Group:         groupForTest(groupsByTest, tr.Name),
				Status:        tr.Status.String(),
				Elapsed:       tr.Elapsed,
				FailureLog:    failureLog,
			}); encErr != nil {
				fmt.Fprintln(errOut, "writing JSON:", encErr)
				if streamErr == nil {
					streamErr = encErr
				}
			}
			return
		}
		streamTextResult(out, tr, groupsByTest, failDir)
	})
	cur.Results = mergeResults(prev.Results, cur.Results)
	persistSuiteFailureState(&cur, failDir, res, res, red, func(action string, err error) {
		fmt.Fprintln(errOut, action+":", err)
	})
	if saveErr := cur.Save(statePath); saveErr != nil {
		fmt.Fprintln(errOut, "saving state:", saveErr)
	}
	if runErr != nil {
		fmt.Fprintln(errOut, "run error:", runErr)
	}
	if opts.Format == formatJSON {
		exitCode := exitFromResult(res)
		if runErr != nil {
			exitCode = 1
		}
		if streamErr != nil {
			exitCode = 1
		}
		if _, ferr := engine.WriteFailures(failDir, res, red); ferr != nil {
			fmt.Fprintln(errOut, "writing failure logs:", ferr)
			exitCode = 1
		}
		if err := emitRunJSONSummary(jw, opts, spec, statePath, failDir, cur, res, exitCode); err != nil {
			fmt.Fprintln(errOut, "writing JSON:", err)
			return 1
		}
		return exitCode
	}
	engine.TerminalSummary(out, opts.Groups, res, red)
	logs, ferr := engine.WriteFailures(failDir, res, red)
	if ferr != nil {
		fmt.Fprintln(errOut, "writing failure logs:", ferr)
	} else if len(logs) > 0 {
		fmt.Fprintf(out, "\nfailure output written to %s (%d file(s)):\n", failDir, len(logs))
		for _, l := range logs {
			fmt.Fprintf(out, "  %s\n", l.Path)
		}
	}
	if runErr != nil {
		return 1
	}
	return exitFromResult(res)
}

// mergeResults returns prior results for tests that produced no fresh result
// this run, followed by all fresh results. Any test re-run this invocation is
// fully replaced (its prior top-level and subtest entries are dropped), so no
// stale or duplicate entries remain; untouched tests are preserved. Deriving
// the rerun set from actual results (not the -run pattern) keeps the merge
// correct for arbitrary regexes and subtests.
func mergeResults(prior, fresh []engine.PersistResult) []engine.PersistResult {
	reran := make(map[string]bool, len(fresh))
	for _, r := range fresh {
		reran[r.Test] = true
	}
	var merged []engine.PersistResult
	for _, r := range prior {
		if !reran[r.Test] {
			merged = append(merged, r)
		}
	}
	return append(merged, fresh...)
}

// ── subcommand implementations ───────────────────────────────────────────────

func runPreflight(args []string, out, errOut io.Writer, d deps) int {
	return runPreflightContext(context.Background(), args, out, errOut, d)
}

func runPreflightContext(ctx context.Context, args []string, out, errOut io.Writer, d deps) int {
	fs := flag.NewFlagSet("preflight", flag.ContinueOnError)
	fs.SetOutput(errOut)
	getenv := getenvFunc(d)
	mode := defaultMode(getenv)
	var (
		repoRoot          string
		group             string
		runRE             string
		allowUnclassified bool
	)
	format := formatText
	addJSONFlag(fs, &format)
	fs.StringVar(&mode, "mode", mode, "run mode")
	fs.StringVar(&group, "group", "", "check only tests in this group")
	fs.StringVar(&runRE, "run", "", "check only tests matching this run regex (overrides --group)")
	fs.BoolVar(&allowUnclassified, "allow-unclassified", false, "allow unclassified tests using conservative requirements instead of failing")
	fs.StringVar(&repoRoot, "repo-root", "", "repo root override")
	if err := fs.Parse(args); err != nil {
		return 2
	}

	root, err := d.resolveRoot(repoRoot)
	if err != nil {
		fmt.Fprintln(errOut, "resolving repo root:", err)
		return 1
	}

	plan, _, err := planForCommand(ctx, d, root, planningOptions{
		Mode:              mode,
		Group:             group,
		Run:               runRE,
		AllowUnclassified: allowUnclassified,
	})
	if err != nil {
		fmt.Fprintln(errOut, "planning:", err)
		exitCode := planExitCode(err)
		// Terminate the NDJSON stream with a summary for a RUNTIME planning
		// failure, so automation never sees an empty stdout stream where a
		// plan was attempted and failed. A usage error (exit 2, e.g. an
		// invalid --run regex) is deliberately excluded: the harness reports
		// argument mistakes on stderr only, matching `run --json`, which
		// rejects the same input before any planning happens.
		if format == formatJSON && exitCode == 1 {
			if encErr := newJSONWriter(out).Encode(jsonPreflightSummary{
				SchemaVersion: 1,
				Type:          "summary",
				Command:       "preflight",
				Mode:          mode,
				OK:            false,
				ExitCode:      exitCode,
			}); encErr != nil {
				fmt.Fprintln(errOut, "writing JSON:", encErr)
			}
		}
		return exitCode
	}

	prov := d.provider
	red := redactorForProvider(prov, getenv)
	plan = redactPlan(plan, red)

	var jw *jsonWriter
	if format == formatJSON {
		jw = newJSONWriter(out)
	}
	if err := emitPlan(out, jw, format, plan); err != nil {
		fmt.Fprintln(errOut, "writing JSON:", err)
		return 1
	}

	requirements := provider.TestRequirements{
		Scopes:       plan.Scopes,
		Capabilities: plan.Capabilities,
		SideEffects:  plan.SideEffects,
	}
	report := prov.Preflight(ctx, mode, requirements)
	if format == formatJSON {
		for _, check := range report.Checks {
			if err := jw.Encode(jsonCheckEvent{
				SchemaVersion: 1,
				Type:          "check",
				Mode:          mode,
				Name:          check.Name,
				Status:        check.Status.String(),
				Detail:        red.String(check.Detail),
				Fix:           stringPtr(red.String(check.Fix)),
			}); err != nil {
				fmt.Fprintln(errOut, "writing JSON:", err)
				return 1
			}
		}
		exitCode := 0
		if !report.OK() {
			exitCode = 1
		}
		if err := jw.Encode(jsonPreflightSummary{
			SchemaVersion: 1,
			Type:          "summary",
			Command:       "preflight",
			Mode:          mode,
			OK:            report.OK(),
			Checks:        checkCounts(report),
			ExitCode:      exitCode,
		}); err != nil {
			fmt.Fprintln(errOut, "writing JSON:", err)
			return 1
		}
		return exitCode
	}
	printPreflight(out, report, red)
	if !report.OK() {
		return 1
	}
	return 0
}

func runGroups(args []string, out, errOut io.Writer, d deps) int {
	fs := flag.NewFlagSet("groups", flag.ContinueOnError)
	fs.SetOutput(errOut)
	var unmatched bool
	var repoRoot string
	format := formatText
	addJSONFlag(fs, &format)
	fs.BoolVar(&unmatched, "unmatched", false, "list only unmatched (misc) tests, exit 1 if any found")
	fs.StringVar(&repoRoot, "repo-root", "", "repo root override")
	if err := fs.Parse(args); err != nil {
		return 2
	}

	root, err := d.resolveRoot(repoRoot)
	if err != nil {
		fmt.Fprintln(errOut, "resolving repo root:", err)
		return 1
	}

	prov := d.provider
	ctx := context.Background()
	names, err := d.list(ctx, root, prov.TestPackages(), prov.TestPattern())
	if err != nil {
		fmt.Fprintln(errOut, "listing tests:", err)
		return 1
	}

	groups, unm := engine.GroupTests(names, prov)
	if format == formatJSON {
		jw := newJSONWriter(out)
		for _, group := range groups {
			if err := jw.Encode(jsonGroupEvent{
				SchemaVersion: 1,
				Type:          "group",
				Name:          group.Name,
				Tests:         group.Tests,
				Total:         len(group.Tests),
			}); err != nil {
				fmt.Fprintln(errOut, "writing JSON:", err)
				return 1
			}
		}
		exitCode := 0
		if unmatched && len(unm) > 0 {
			exitCode = 1
		}
		summary := jsonGroupsSummary{
			SchemaVersion: 1,
			Type:          "summary",
			Command:       "groups",
			Groups:        len(groups),
			Tests:         len(names),
			Unmatched:     len(unm),
			ExitCode:      exitCode,
		}
		if unmatched && len(unm) > 0 {
			summary.UnmatchedList = unm
		}
		if err := jw.Encode(summary); err != nil {
			fmt.Fprintln(errOut, "writing JSON:", err)
			return 1
		}
		return exitCode
	}
	for _, g := range groups {
		fmt.Fprintf(out, "%s (%d)\n", g.Name, len(g.Tests))
		for _, t := range g.Tests {
			fmt.Fprintf(out, "  %s\n", t)
		}
	}

	if unmatched && len(unm) > 0 {
		fmt.Fprintln(out, "unmatched:")
		for _, t := range unm {
			fmt.Fprintf(out, "  %s\n", t)
		}
		return 1
	}
	return 0
}

func runRun(args []string, out, errOut io.Writer, d deps) int {
	return runRunContext(context.Background(), args, out, errOut, d, runContextOptions{
		Preflight:     true,
		PreserveState: true,
	})
}

type runContextOptions struct {
	Preflight     bool
	PreserveState bool
	LockHeld      bool
	Plan          *engine.ExecutionPlan
}

func runRunContext(ctx context.Context, args []string, out, errOut io.Writer, d deps, contextOpts runContextOptions) int {
	fs := flag.NewFlagSet("run", flag.ContinueOnError)
	fs.SetOutput(errOut)
	getenv := getenvFunc(d)
	var (
		mode               = defaultMode(getenv)
		group              string
		runRE              string
		timeout            time.Duration
		repoRoot           string
		allowSensitiveLogs bool
		allowUnclassified  bool
		format             = formatText
		tuiFlag            bool
		noTUIFlag          bool
		triage             triageOptions
	)
	addJSONFlag(fs, &format)
	fs.StringVar(&mode, "mode", mode, "run mode")
	fs.StringVar(&group, "group", "", "run only tests in this group")
	fs.StringVar(&runRE, "run", "", "run regex (overrides --group)")
	fs.DurationVar(&timeout, "timeout", 120*time.Minute, "test timeout")
	fs.StringVar(&repoRoot, "repo-root", "", "repo root override")
	fs.BoolVar(&allowSensitiveLogs, "allow-sensitive-logs", false, "forward TF_LOG* to child process")
	fs.BoolVar(&allowUnclassified, "allow-unclassified", false, "run unclassified tests with conservative requirements")
	fs.BoolVar(&tuiFlag, "tui", false, "open the interactive dashboard when possible")
	fs.BoolVar(&noTUIFlag, "no-tui", false, "force non-interactive text output")
	addTriageRunFlags(fs, &triage)
	if err := fs.Parse(args); err != nil {
		return 2
	}
	if err := triage.finalize(fs); err != nil {
		fmt.Fprintln(errOut, err)
		return 2
	}
	if err := validateIssueFilingFlags(triage, d, errOut); err != nil {
		fmt.Fprintln(errOut, err)
		return 2
	}
	var compiledRun *regexp.Regexp
	if runRE != "" {
		re, reErr := regexp.Compile(runRE)
		if reErr != nil {
			fmt.Fprintf(errOut, "invalid --run regex: %v\n", reErr)
			return 2
		}
		compiledRun = re
	}

	root, err := d.resolveRoot(repoRoot)
	if err != nil {
		fmt.Fprintln(errOut, "resolving repo root:", err)
		return 1
	}
	// Caller-owned nested runs (for example guided e2e's internal run step)
	// must remain non-interactive: they may arrive with a preflighted plan,
	// caller-held state lock, or both, and auto-launching the dashboard would
	// hijack that batch flow on real TTYs.
	nestedInvocation := contextOpts.Plan != nil || contextOpts.LockHeld
	if !nestedInvocation && shouldLaunchDashboard(format, tuiFlag, noTUIFlag, d, errOut) {
		owner := getenv("GITHUB_OWNER")
		ascii := getenv("NO_COLOR") != ""
		return runDashboard(d, root, mode, owner, ascii, allowUnclassified, out, errOut)
	}

	prov := d.provider

	// planPtr carries the redacted plan built or supplied below into
	// prevState.Plan for state persistence. Guided e2e supplies its already
	// emitted/preflighted plan with Preflight false and a caller-held lock, so
	// the run step can reuse the same selection without rediscovery.
	var (
		groups    []engine.Group
		pattern   string
		selection selectionInfo
		planPtr   *engine.ExecutionPlan
	)
	if contextOpts.Plan != nil {
		plan := *contextOpts.Plan
		if plan.Mode != "" {
			mode = plan.Mode
		}
		red := redactorForProvider(prov, getenv)
		plan = redactPlan(plan, red)

		if contextOpts.Preflight {
			var jw *jsonWriter
			if format == formatJSON {
				jw = newJSONWriter(out)
			}
			if err := emitPlan(out, jw, format, plan); err != nil {
				fmt.Fprintln(errOut, "writing JSON:", err)
				return 1
			}

			requirements := provider.TestRequirements{
				Scopes:       plan.Scopes,
				Capabilities: plan.Capabilities,
				SideEffects:  plan.SideEffects,
			}
			report := prov.Preflight(ctx, mode, requirements)
			if format == formatJSON {
				if err := emitCapabilities(jw, mode, report, red); err != nil {
					fmt.Fprintln(errOut, "writing JSON:", err)
					return 1
				}
			}
			if !report.OK() {
				printPreflight(errOut, report, red)
				if format == formatJSON {
					return emitPreRunGateJSON(out, errOut, root, mode, prov)
				}
				return 1
			}
		}

		planGroups, _ := engine.GroupTests(plan.Selected, prov)
		groups = eligibleGroups(planGroups, plan.Eligible)
		pattern = engine.RunPattern(plan.Eligible)
		selection = selectionInfo{Group: plan.Group, Run: plan.Run, Tests: len(plan.Selected)}
		planPtr = &plan
	} else if contextOpts.Preflight {
		// Resolve root -> discover/plan -> emit redacted plan -> plan-aware
		// preflight. Discovery always happens before preflight, and a
		// planning failure calls neither preflight nor the runner.
		plan, planGroups, planErr := planForCommand(ctx, d, root, planningOptions{
			Mode:              mode,
			Group:             group,
			Run:               runRE,
			AllowUnclassified: allowUnclassified,
		})
		if planErr != nil {
			fmt.Fprintln(errOut, "planning:", planErr)
			exitCode := planExitCode(planErr)
			// Automation reading the NDJSON stream must still get a terminal
			// object instead of empty stdout. No plan event is fabricated:
			// planning failed, so there is no plan to report. As in
			// preflight, only a runtime planning failure gets a summary; a
			// usage error (exit 2) stays stderr-only, and `run` rejects its
			// one usage-level planning input - an invalid --run regex -
			// before planning is even reached.
			if format == formatJSON && exitCode == 1 {
				emitPlanFailureRunJSON(out, errOut, root, mode, prov,
					selectionInfo{Group: group, Run: runRE}, exitCode)
			}
			return exitCode
		}

		red := redactorForProvider(prov, getenv)
		plan = redactPlan(plan, red)

		var jw *jsonWriter
		if format == formatJSON {
			jw = newJSONWriter(out)
		}
		if err := emitPlan(out, jw, format, plan); err != nil {
			fmt.Fprintln(errOut, "writing JSON:", err)
			return 1
		}

		requirements := provider.TestRequirements{
			Scopes:       plan.Scopes,
			Capabilities: plan.Capabilities,
			SideEffects:  plan.SideEffects,
		}
		report := prov.Preflight(ctx, mode, requirements)
		if format == formatJSON {
			if err := emitCapabilities(jw, mode, report, red); err != nil {
				fmt.Fprintln(errOut, "writing JSON:", err)
				return 1
			}
		}
		if !report.OK() {
			printPreflight(errOut, report, red)
			if format == formatJSON {
				return emitPreRunGateJSON(out, errOut, root, mode, prov)
			}
			return 1
		}

		groups = eligibleGroups(planGroups, plan.Eligible)
		pattern = engine.RunPattern(plan.Eligible)
		selection = selectionInfo{Group: group, Run: runRE, Tests: len(plan.Selected)}
		planPtr = &plan
	} else {
		names, err := d.list(ctx, root, prov.TestPackages(), prov.TestPattern())
		if err != nil {
			fmt.Fprintln(errOut, "listing tests:", err)
			return 1
		}

		groups, _ = engine.GroupTests(names, prov)

		// Determine the selection pattern. The set of tests actually re-run
		// is derived from results after the run (see mergeResults), so no
		// pattern inference is needed here.
		selection = selectionInfo{Group: group, Run: runRE, Tests: len(names)}
		switch {
		case runRE != "":
			pattern = runRE
			selection.Tests = countMatchingTests(names, compiledRun)
		case group != "":
			for _, g := range groups {
				if g.Name == group {
					pattern = engine.RunPattern(g.Tests)
					selection.Tests = len(g.Tests)
					break
				}
			}
		default:
			pattern = engine.RunPattern(names)
		}
	}

	if pattern == "" {
		if format == formatJSON {
			statePath := filepath.Join(root, ".pulsar-state.json")
			spec := engine.RunSpec{
				Dir:                root,
				Packages:           prov.TestPackages(),
				Pattern:            pattern,
				Timeout:            timeout,
				ExtraEnv:           buildExtraEnv(mode, prov.EnvFor(mode), getenv),
				AllowSensitiveLogs: allowSensitiveLogs,
			}
			return emitEmptyRunJSON(out, errOut, mode, statePath, spec, runOptions{
				Format:    format,
				Command:   "run",
				Selection: selection,
				Groups:    groups,
			})
		}
		fmt.Fprintln(out, "nothing to run")
		return 0
	}

	spec := engine.RunSpec{
		Dir:                root,
		Packages:           prov.TestPackages(),
		Pattern:            pattern,
		Timeout:            timeout,
		ExtraEnv:           buildExtraEnv(mode, prov.EnvFor(mode), getenv),
		AllowSensitiveLogs: allowSensitiveLogs,
	}

	statePath := filepath.Join(root, ".pulsar-state.json")
	lockPath := statePath + ".lock"
	if !contextOpts.LockHeld {
		release, err := engine.AcquireLock(lockPath)
		if err != nil {
			fmt.Fprintln(errOut, "acquiring lock:", err)
			return 1
		}
		defer release() //nolint:errcheck
	}

	prevState, _ := engine.Load(statePath)
	if !nestedInvocation {
		// A top-level run starts a new logical run. Guided orphan accounting
		// belongs only to the e2e window that captured its baseline; carrying
		// it here would misattribute the prior run's completed delta. Nested
		// e2e runs retain the caller-owned baseline below.
		prevState.Orphans = nil
	}
	if !contextOpts.PreserveState {
		prevState = engine.State{Orphans: prevState.Orphans}
	}
	if planPtr != nil {
		prevState.Plan = planPtr
	}

	runner := runnerForFormat(d, format, out)
	return runWithSpec(ctx, spec, statePath, prevState, prov, mode, runner, out, errOut, runOptions{
		Format:        format,
		Command:       "run",
		Selection:     selection,
		Groups:        groups,
		Triage:        triage,
		IssueRegistry: d.newIssueRegistry,
		IssueFiler:    d.newIssueFiler,
		Getenv:        getenv,
	})
}

// planRequiredError is the exact user-facing message for a legacy run that has
// no persisted execution plan. It names both entry-points so the operator can
// start a fresh run immediately.
const planRequiredError = "state/plan-required: the last run predates execution plans; " +
	"start a new run with `terraform-provider-tester run --mode <mode>` or " +
	"`terraform-provider-tester e2e --mode <mode>`"

// intersectOrdered returns the elements of allowed that are also in the set
// built from names, preserving the order of allowed.
func intersectOrdered(names, allowed []string) []string {
	set := make(map[string]bool, len(names))
	for _, n := range names {
		set[n] = true
	}
	var out []string
	for _, a := range allowed {
		if set[a] {
			out = append(out, a)
		}
	}
	return out
}

// retrySelection returns the ordered set of tests that should be retried:
// FailedTopLevel ∩ Plan.Eligible (in Eligible order). An error is returned when
// State.Plan is nil (legacy planless state).
func retrySelection(state engine.State) ([]string, error) {
	if state.Plan == nil {
		return nil, fmt.Errorf("%s", planRequiredError)
	}
	failed := state.FailedTopLevel()
	return intersectOrdered(failed, state.Plan.Eligible), nil
}

// resumeSelection returns the ordered set of tests that should be resumed:
// eligible failed tests followed by State.NotRun(Plan.Eligible), deduplicated.
// An error is returned when State.Plan is nil (legacy planless state).
func resumeSelection(state engine.State) ([]string, error) {
	if state.Plan == nil {
		return nil, fmt.Errorf("%s", planRequiredError)
	}
	failed := state.FailedTopLevel()
	eligibleFailed := intersectOrdered(failed, state.Plan.Eligible)
	notRun := state.NotRun(state.Plan.Eligible)
	// Deduplicate: notRun may overlap with eligibleFailed if a result is absent
	// for a failed entry (defensive; in practice they are disjoint).
	seen := make(map[string]bool, len(eligibleFailed))
	result := make([]string, 0, len(eligibleFailed)+len(notRun))
	for _, n := range eligibleFailed {
		if !seen[n] {
			seen[n] = true
			result = append(result, n)
		}
	}
	for _, n := range notRun {
		if !seen[n] {
			seen[n] = true
			result = append(result, n)
		}
	}
	return result, nil
}

func persistedRetryResumeMode(state engine.State) (string, error) {
	if state.Plan == nil {
		return "", nil
	}
	if state.Plan.Mode == "" {
		return "", fmt.Errorf("state/mode-mismatch: persisted execution plan is missing its mode; start a new run")
	}
	if state.Mode != "" && state.Mode != state.Plan.Mode {
		return "", fmt.Errorf("state/mode-mismatch: persisted state mode %q conflicts with persisted plan mode %q; start a new run", state.Mode, state.Plan.Mode)
	}
	return state.Plan.Mode, nil
}

// resolveRetryResumeMode resolves the effective mode for retry/resume.
// If modeExplicit is true, the caller supplied an explicit --mode value; if it
// differs from the plan's mode the call is an error. If modeExplicit is false,
// the plan's mode is used (falling back to the environment default when the
// plan is nil. The normal retry path already rejects a planless state via
// retrySelection, so a nil plan reaches here only through runRetryContext's
// internal selected-list parameter — exercised today only by tests, not by
// e2e or any other production caller).
func resolveRetryResumeMode(explicit bool, flagMode string, state engine.State, getenv func(string) string) (string, error) {
	persistedMode, err := persistedRetryResumeMode(state)
	if err != nil {
		return "", err
	}
	if explicit {
		// Explicit flag always wins. When a plan exists, reject a mismatch so
		// the operator is not silently running with the wrong mode.
		if persistedMode != "" && flagMode != persistedMode {
			return "", fmt.Errorf("mode mismatch: --mode %q conflicts with the persisted plan mode %q; "+
				"omit --mode to use the plan's mode, or start a new run", flagMode, persistedMode)
		}
		return flagMode, nil
	}
	if persistedMode != "" {
		return persistedMode, nil
	}
	return defaultMode(getenv), nil
}

func runRetry(args []string, out, errOut io.Writer, d deps) int {
	return runRetryContext(context.Background(), args, out, errOut, d, nil)
}

func runRetryContext(ctx context.Context, args []string, out, errOut io.Writer, d deps, selected []string) int {
	fs := flag.NewFlagSet("retry", flag.ContinueOnError)
	fs.SetOutput(errOut)
	getenv := getenvFunc(d)
	var (
		failed             bool
		modeFlag           string // resolved after state load; empty = not explicit
		timeout            time.Duration
		repoRoot           string
		allowSensitiveLogs bool
		format             = formatText
		tuiFlag            bool
		noTUIFlag          bool
		triage             triageOptions
	)
	addJSONFlag(fs, &format)
	fs.BoolVar(&failed, "failed", false, "retry only failed tests (required)")
	fs.StringVar(&modeFlag, "mode", "", "run mode (defaults to persisted plan mode)")
	fs.DurationVar(&timeout, "timeout", 120*time.Minute, "test timeout")
	fs.StringVar(&repoRoot, "repo-root", "", "repo root override")
	fs.BoolVar(&allowSensitiveLogs, "allow-sensitive-logs", false, "forward TF_LOG* to child process")
	fs.BoolVar(&tuiFlag, "tui", false, "open the interactive dashboard when possible")
	fs.BoolVar(&noTUIFlag, "no-tui", false, "force non-interactive text output")
	addTriageRunFlags(fs, &triage)
	if err := fs.Parse(args); err != nil {
		return 2
	}
	modeExplicit := false
	fs.Visit(func(f *flag.Flag) {
		if f.Name == "mode" {
			modeExplicit = true
		}
	})
	if err := triage.finalize(fs); err != nil {
		fmt.Fprintln(errOut, err)
		return 2
	}
	if err := validateIssueFilingFlags(triage, d, errOut); err != nil {
		fmt.Fprintln(errOut, err)
		return 2
	}
	if !failed {
		fmt.Fprintln(errOut, "usage: retry --failed [flags]")
		return 2
	}

	root, err := d.resolveRoot(repoRoot)
	if err != nil {
		fmt.Fprintln(errOut, "resolving repo root:", err)
		return 1
	}
	statePath := filepath.Join(root, ".pulsar-state.json")
	if shouldLaunchDashboard(format, tuiFlag, noTUIFlag, d, errOut) {
		state, loadErr := engine.Load(statePath)
		if loadErr != nil {
			fmt.Fprintln(errOut, "loading state:", loadErr)
			return 1
		}
		dashMode, modeErr := resolveRetryResumeMode(modeExplicit, modeFlag, state, getenv)
		if modeErr != nil {
			fmt.Fprintln(errOut, modeErr)
			return 2
		}
		owner := getenv("GITHUB_OWNER")
		ascii := getenv("NO_COLOR") != ""
		// retry/resume have no --allow-unclassified flag: their selection
		// comes from the persisted plan, so dashboard planning fails closed.
		return runDashboard(d, root, dashMode, owner, ascii, false, out, errOut)
	}
	lockPath := statePath + ".lock"
	release, err := engine.AcquireLock(lockPath)
	if err != nil {
		fmt.Fprintln(errOut, "acquiring lock:", err)
		return 1
	}
	defer release() //nolint:errcheck

	prevState, err := engine.Load(statePath)
	if err != nil {
		fmt.Fprintln(errOut, "loading state:", err)
		return 1
	}

	sel := selected
	if sel == nil {
		// Plan-required check: planless state is rejected on the direct retry path.
		var selErr error
		sel, selErr = retrySelection(prevState)
		if selErr != nil {
			fmt.Fprintln(errOut, selErr)
			return 2
		}
	}

	// Resolve the effective mode: explicit --mode must match the plan; omitting
	// --mode defaults to the plan's mode (or the environment default when the
	// plan is nil, reachable here only via the selected-list bypass above,
	// which today only a test exercises).
	mode, modeErr := resolveRetryResumeMode(modeExplicit, modeFlag, prevState, getenv)
	if modeErr != nil {
		fmt.Fprintln(errOut, modeErr)
		return 2
	}

	if len(sel) == 0 {
		if format == formatJSON {
			spec := engine.RunSpec{
				Dir:                root,
				Packages:           d.provider.TestPackages(),
				Pattern:            "",
				Timeout:            timeout,
				ExtraEnv:           buildExtraEnv(mode, d.provider.EnvFor(mode), getenv),
				AllowSensitiveLogs: allowSensitiveLogs,
			}
			return emitEmptyRunJSON(out, errOut, mode, statePath, spec, runOptions{
				Format:    format,
				Command:   "retry",
				Selection: selectionInfo{Tests: 0},
			})
		}
		fmt.Fprintln(out, "nothing to run")
		return 0
	}

	prov := d.provider
	var groups []engine.Group
	if format == formatJSON {
		var err error
		groups, _, err = discoverGroups(ctx, d, root)
		if err != nil {
			fmt.Fprintln(errOut, err)
			return 1
		}
	}
	spec := engine.RunSpec{
		Dir:                root,
		Packages:           prov.TestPackages(),
		Pattern:            engine.RunPattern(sel),
		Timeout:            timeout,
		ExtraEnv:           buildExtraEnv(mode, prov.EnvFor(mode), getenv),
		AllowSensitiveLogs: allowSensitiveLogs,
	}

	runner := runnerForFormat(d, format, out)
	return runWithSpec(ctx, spec, statePath, prevState, prov, mode, runner, out, errOut, runOptions{
		Format:        format,
		Command:       "retry",
		Selection:     selectionInfo{Tests: len(sel)},
		Groups:        groups,
		Triage:        triage,
		IssueRegistry: d.newIssueRegistry,
		IssueFiler:    d.newIssueFiler,
		Getenv:        getenv,
	})
}

func runResume(args []string, out, errOut io.Writer, d deps) int {
	fs := flag.NewFlagSet("resume", flag.ContinueOnError)
	fs.SetOutput(errOut)
	getenv := getenvFunc(d)
	var (
		modeFlag           string // resolved after state load; empty = not explicit
		timeout            time.Duration
		repoRoot           string
		allowSensitiveLogs bool
		format             = formatText
		tuiFlag            bool
		noTUIFlag          bool
		triage             triageOptions
	)
	addResumeFlags(
		fs,
		&modeFlag,
		&timeout,
		&repoRoot,
		&allowSensitiveLogs,
		&format,
		&tuiFlag,
		&noTUIFlag,
		&triage,
	)
	if err := fs.Parse(args); err != nil {
		return 2
	}
	modeExplicit := false
	fs.Visit(func(f *flag.Flag) {
		if f.Name == "mode" {
			modeExplicit = true
		}
	})
	if err := triage.finalize(fs); err != nil {
		fmt.Fprintln(errOut, err)
		return 2
	}
	if err := validateIssueFilingFlags(triage, d, errOut); err != nil {
		fmt.Fprintln(errOut, err)
		return 2
	}

	root, err := d.resolveRoot(repoRoot)
	if err != nil {
		fmt.Fprintln(errOut, "resolving repo root:", err)
		return 1
	}
	statePath := filepath.Join(root, ".pulsar-state.json")
	if modeExplicit {
		state, loadErr := engine.Load(statePath)
		if loadErr != nil {
			fmt.Fprintln(errOut, "loading state:", loadErr)
			return 1
		}
		if _, modeErr := resolveRetryResumeMode(true, modeFlag, state, getenv); modeErr != nil {
			fmt.Fprintln(errOut, modeErr)
			return 2
		}
	}
	if shouldLaunchDashboard(format, tuiFlag, noTUIFlag, d, errOut) {
		state, loadErr := engine.Load(statePath)
		if loadErr != nil {
			fmt.Fprintln(errOut, "loading state:", loadErr)
			return 1
		}
		dashMode, modeErr := resolveRetryResumeMode(modeExplicit, modeFlag, state, getenv)
		if modeErr != nil {
			fmt.Fprintln(errOut, modeErr)
			return 2
		}
		owner := getenv("GITHUB_OWNER")
		ascii := getenv("NO_COLOR") != ""
		// retry/resume have no --allow-unclassified flag: their selection
		// comes from the persisted plan, so dashboard planning fails closed.
		return runDashboard(d, root, dashMode, owner, ascii, false, out, errOut)
	}

	prov := d.provider
	ctx := context.Background()

	lockPath := statePath + ".lock"
	release, err := engine.AcquireLock(lockPath)
	if err != nil {
		fmt.Fprintln(errOut, "acquiring lock:", err)
		return 1
	}
	defer release() //nolint:errcheck

	prevState, err := engine.Load(statePath)
	if err != nil {
		fmt.Fprintln(errOut, "loading state:", err)
		return 1
	}

	sel, selErr := resumeSelection(prevState)
	if selErr != nil {
		fmt.Fprintln(errOut, selErr)
		return 2
	}

	mode, modeErr := resolveRetryResumeMode(modeExplicit, modeFlag, prevState, getenv)
	if modeErr != nil {
		fmt.Fprintln(errOut, modeErr)
		return 2
	}
	if resumeOrphanModeMismatch(prevState.Orphans, mode) {
		exitCode, accounting := finalizeResumeOrphanAccounting(
			prov, mode, statePath, prevState, nil, 0, errOut,
		)
		return emitResumeOrphanModeMismatchResult(
			out, errOut, format, root, statePath, mode, sel, timeout, accounting, exitCode,
		)
	}
	red := redactorForProvider(prov, getenv)
	finish := func(exitCode int) (int, *engine.OrphanAccounting) {
		return finalizeResumeOrphanAccounting(prov, mode, statePath, prevState, red, exitCode, errOut)
	}

	if len(sel) == 0 {
		exitCode, accounting := finish(0)
		if format != formatJSON {
			if err := emitResumeOrphanDelta(out, nil, format, root, accounting, red); err != nil {
				fmt.Fprintln(errOut, "writing orphan accounting:", err)
				exitCode = 1
			}
			fmt.Fprintln(out, "nothing to resume")
			return exitCode
		}
		groups, _, discErr := discoverGroups(ctx, d, root)
		if discErr != nil {
			fmt.Fprintln(errOut, discErr)
			exitCode = 1
		}
		spec := engine.RunSpec{
			Dir:                root,
			Packages:           prov.TestPackages(),
			Pattern:            "",
			Timeout:            timeout,
			ExtraEnv:           buildExtraEnv(mode, prov.EnvFor(mode), getenv),
			AllowSensitiveLogs: allowSensitiveLogs,
		}
		if err := emitResumeOrphanDelta(out, newJSONWriter(out), format, root, accounting, red); err != nil {
			fmt.Fprintln(errOut, "writing JSON:", err)
			exitCode = 1
		}
		cleanupStatus := ""
		if accounting != nil {
			cleanupStatus = accounting.CleanupStatus
		}
		return emitEmptyRunJSONWithExit(out, errOut, mode, statePath, spec, runOptions{
			Format:        format,
			Command:       "resume",
			Selection:     selectionInfo{Tests: 0},
			Groups:        groups,
			CleanupStatus: cleanupStatus,
		}, exitCode)
	}
	// Selection is plan-based, but grouping metadata still comes from
	// discovery. Route it through discoverGroups and fail loudly: a provider
	// package that no longer builds makes `go test -list` emit zero tests,
	// and resuming with silently empty grouping metadata reports every result
	// under "misc" and drops every group summary. This matches retry.
	groups, _, err := discoverGroups(ctx, d, root)
	if err != nil {
		fmt.Fprintln(errOut, err)
		return 1
	}

	spec := engine.RunSpec{
		Dir:                root,
		Packages:           prov.TestPackages(),
		Pattern:            engine.RunPattern(sel),
		Timeout:            timeout,
		ExtraEnv:           buildExtraEnv(mode, prov.EnvFor(mode), getenv),
		AllowSensitiveLogs: allowSensitiveLogs,
	}

	runner := runnerForFormat(d, format, out)
	var (
		lastSummary jsonRunSummary
		haveSummary bool
	)
	runExit := runWithSpec(ctx, spec, statePath, prevState, prov, mode, runner, out, errOut, runOptions{
		Format:        format,
		Command:       "resume",
		Selection:     selectionInfo{Tests: len(sel)},
		Groups:        groups,
		Triage:        triage,
		IssueRegistry: d.newIssueRegistry,
		IssueFiler:    d.newIssueFiler,
		Getenv:        getenv,
		JSONSummaryObserver: func(summary jsonRunSummary) {
			lastSummary = summary
			haveSummary = true
		},
	})
	finalExit, accounting := finish(runExit)
	if accounting == nil {
		return finalExit
	}

	var jw *jsonWriter
	if format == formatJSON {
		jw = newJSONWriter(out)
	}
	if err := emitResumeOrphanDelta(out, jw, format, root, accounting, red); err != nil {
		fmt.Fprintln(errOut, "writing orphan accounting:", err)
		finalExit = 1
	}
	if format == formatJSON && haveSummary {
		lastSummary.CleanupStatus = accounting.CleanupStatus
		lastSummary.ExitCode = finalExit
		if err := jw.Encode(lastSummary); err != nil {
			fmt.Fprintln(errOut, "writing JSON:", err)
			return 1
		}
	}
	return finalExit
}

func rejectResumeOrphanModeMismatchBeforeProvider(
	args []string,
	out, errOut io.Writer,
	d deps,
	getenv func(string) string,
) (int, bool) {
	opts, isResume := scanResumeModeGuardOptions(args)
	if !isResume || !opts.optionsValid {
		return 0, false
	}
	root, err := d.resolveRoot(opts.repoRoot)
	if err != nil {
		return 0, false
	}
	statePath := filepath.Join(root, ".pulsar-state.json")
	original, err := engine.Load(statePath)
	if err != nil {
		return 0, false
	}
	sel, selErr := resumeSelection(original)
	if selErr != nil {
		return 0, false
	}
	mode, modeErr := resolveRetryResumeMode(opts.modeSet, opts.mode, original, getenv)
	if modeErr != nil {
		fmt.Fprintln(errOut, modeErr)
		return 2, true
	}
	if !resumeOrphanModeMismatch(original.Orphans, mode) {
		return 0, false
	}

	release, lockErr := engine.AcquireLock(statePath + ".lock")
	if lockErr != nil {
		fmt.Fprintln(errOut, "acquiring lock:", lockErr)
		return 1, true
	}
	defer release() //nolint:errcheck

	original, err = engine.Load(statePath)
	if err != nil {
		fmt.Fprintln(errOut, "loading state:", err)
		return 1, true
	}
	sel, selErr = resumeSelection(original)
	if selErr != nil {
		fmt.Fprintln(errOut, selErr)
		return 2, true
	}
	mode, modeErr = resolveRetryResumeMode(opts.modeSet, opts.mode, original, getenv)
	if modeErr != nil {
		fmt.Fprintln(errOut, modeErr)
		return 2, true
	}
	if !resumeOrphanModeMismatch(original.Orphans, mode) {
		return 0, false
	}

	exitCode, accounting := finalizeResumeOrphanAccounting(
		d.provider, mode, statePath, original, nil, 0, errOut,
	)
	return emitResumeOrphanModeMismatchResult(
		out, errOut, opts.format, root, statePath, mode, sel, opts.timeout, accounting, exitCode,
	), true
}

func emitResumeOrphanModeMismatchResult(
	out, errOut io.Writer,
	format outputFormat,
	root, statePath, mode string,
	sel []string,
	timeout time.Duration,
	accounting *engine.OrphanAccounting,
	exitCode int,
) int {
	if format != formatJSON {
		if err := emitResumeOrphanDelta(out, nil, format, root, accounting, nil); err != nil {
			fmt.Fprintln(errOut, "writing orphan accounting:", err)
			exitCode = 1
		}
		fmt.Fprintln(out, "resume stopped before execution")
		return exitCode
	}

	jw := newJSONWriter(out)
	if err := emitResumeOrphanDelta(out, jw, format, root, accounting, nil); err != nil {
		fmt.Fprintln(errOut, "writing JSON:", err)
		exitCode = 1
	}
	cleanupStatus := ""
	if accounting != nil {
		cleanupStatus = accounting.CleanupStatus
	}
	return emitEmptyRunJSONWithExit(out, errOut, mode, statePath, engine.RunSpec{
		Dir:     root,
		Pattern: engine.RunPattern(sel),
		Timeout: timeout,
	}, runOptions{
		Format:        format,
		Command:       "resume",
		Selection:     selectionInfo{Tests: len(sel)},
		CleanupStatus: cleanupStatus,
	}, exitCode)
}

func resumeOrphanModeMismatch(accounting *engine.OrphanAccounting, mode string) bool {
	return accounting != nil && accounting.Mode != "" && accounting.Mode != mode
}

func cloneOrphanAccounting(original *engine.OrphanAccounting) *engine.OrphanAccounting {
	if original == nil {
		return nil
	}
	accounting := *original
	if original.Baseline != nil {
		accounting.Baseline = append([]provider.Resource{}, original.Baseline...)
	}
	return &accounting
}

func finalizeResumeOrphanAccounting(prov provider.Provider, mode, statePath string, original engine.State, red *redact.Redactor, exitCode int, errOut io.Writer) (int, *engine.OrphanAccounting) {
	if resumeOrphanModeMismatch(original.Orphans, mode) {
		state := original
		state.Orphans = cloneOrphanAccounting(original.Orphans)
		state.Orphans.Final = nil
		state.Orphans.PreExisting = nil
		state.Orphans.New = nil
		state.Orphans.FinalCaptured = false
		state.Orphans.CleanupStatus = engine.CleanupUnknown
		fmt.Fprintf(errOut,
			"state/orphan-mode-mismatch: persisted orphan accounting mode %q conflicts with resolved resume mode %q; "+
				"start a new guided `e2e` run in the intended mode instead of resuming this state\n",
			original.Orphans.Mode, mode,
		)
		if err := state.Save(statePath); err != nil {
			fmt.Fprintln(errOut, "saving state:", err)
		}
		return 1, state.Orphans
	}
	if !e2eCredentialedModes[mode] {
		return exitCode, nil
	}
	// A nil record identifies a non-guided run, which never opened an orphan
	// accounting window. A present but uncaptured record remains fail-closed.
	if original.Orphans == nil {
		return exitCode, nil
	}

	state, err := engine.Load(statePath)
	if err != nil {
		fmt.Fprintln(errOut, "loading state for final orphan accounting:", err)
		state = original
		exitCode = 1
	}
	state.Orphans = cloneOrphanAccounting(original.Orphans)
	finalCtx, cancelFinal := freshFinalOrphanContext()
	defer cancelFinal()
	if err := finalizeOrphanAccountingWithRedactor(finalCtx, prov, mode, &state, red); err != nil {
		fmt.Fprintln(errOut, "capturing final orphan state:", err)
		exitCode = 1
	}
	if err := state.Save(statePath); err != nil {
		fmt.Fprintln(errOut, "saving state:", err)
		state.Orphans = invalidateFinalOrphanAccounting(state.Orphans, mode)
		exitCode = 1
	}
	return exitCode, state.Orphans
}

// discoverGroups lists acceptance tests via the injected lister and groups
// them. It surfaces discovery and grouping errors instead of silently yielding
// an empty set: a build failure in the provider package makes `go test -list`
// emit zero tests, and callers must fail loudly rather than render an empty
// report or open an empty dashboard.
func discoverGroups(ctx context.Context, d deps, root string) ([]engine.Group, []string, error) {
	prov := d.provider
	names, err := d.list(ctx, root, prov.TestPackages(), prov.TestPattern())
	if err != nil {
		return nil, nil, fmt.Errorf("discovering tests: %w", err)
	}
	groups, _ := engine.GroupTests(names, prov)
	var allNames []string
	for _, g := range groups {
		allNames = append(allNames, g.Tests...)
	}
	return groups, allNames, nil
}

func runReport(args []string, out, errOut io.Writer, d deps) int {
	fs := flag.NewFlagSet("report", flag.ContinueOnError)
	fs.SetOutput(errOut)
	getenv := getenvFunc(d)
	var htmlPath, mdPath, repoRoot string
	fs.StringVar(&htmlPath, "html", "", "HTML output path")
	fs.StringVar(&mdPath, "md", "", "markdown output path")
	fs.StringVar(&repoRoot, "repo-root", "", "repo root override")
	if err := fs.Parse(args); err != nil {
		return 2
	}

	root, err := d.resolveRoot(repoRoot)
	if err != nil {
		fmt.Fprintln(errOut, "resolving repo root:", err)
		return 1
	}

	statePath := filepath.Join(root, ".pulsar-state.json")
	prevState, err := engine.Load(statePath)
	if err != nil {
		fmt.Fprintln(errOut, "loading state:", err)
		return 1
	}

	prov := d.provider

	// Build redactor from secret values. Secret values are read from env but
	// never printed; only the redacted output reaches the report file.
	var secretVals []string
	for _, key := range prov.SecretEnvKeys() {
		if val := getenv(key); val != "" {
			secretVals = append(secretVals, val)
		}
	}
	red := redact.New(secretVals)

	// Build groups for report context.
	ctx := context.Background()
	groups, _, err := discoverGroups(ctx, d, root)
	if err != nil {
		fmt.Fprintln(errOut, err)
		return 1
	}

	if mdPath != "" {
		if err := engine.ExportMarkdownState(mdPath, groups, prevState, red); err != nil {
			fmt.Fprintln(errOut, "exporting markdown:", err)
			return 1
		}
	}
	if htmlPath != "" {
		if err := engine.ExportHTMLState(htmlPath, groups, prevState, red); err != nil {
			fmt.Fprintln(errOut, "exporting HTML:", err)
			return 1
		}
	}
	if mdPath == "" && htmlPath == "" {
		engine.TerminalStateSummary(out, groups, prevState, red)
		failDir := filepath.Join(root, ".pulsar-failures")
		if matches, _ := filepath.Glob(filepath.Join(failDir, "*.log")); len(matches) > 0 {
			fmt.Fprintf(out, "\nfailure logs from the last run (%d):\n", len(matches))
			for _, m := range matches {
				fmt.Fprintf(out, "  %s\n", m)
			}
		}
	}
	return 0
}

// runUnlock clears a stale run lock left behind by a crashed or killed run.
// It is the manual escape hatch for the cases the automatic reclaim in
// engine.AcquireLock deliberately leaves in place (for example a lock written
// by another host). Removing an absent lock succeeds with a no-op message.
func runUnlock(args []string, out, errOut io.Writer, d deps) int {
	fs := flag.NewFlagSet("unlock", flag.ContinueOnError)
	fs.SetOutput(errOut)
	var repoRoot string
	fs.StringVar(&repoRoot, "repo-root", "", "repo root override")
	if err := fs.Parse(args); err != nil {
		return 2
	}

	root, err := d.resolveRoot(repoRoot)
	if err != nil {
		fmt.Fprintln(errOut, "resolving repo root:", err)
		return 1
	}

	lockPath := filepath.Join(root, ".pulsar-state.json") + ".lock"
	msg, err := engine.RemoveLock(lockPath)
	if err != nil {
		fmt.Fprintln(errOut, "unlock:", err)
		return 1
	}
	fmt.Fprintln(out, msg)
	return 0
}

func runDiscover(args []string, out, errOut io.Writer, d deps) int {
	fs := flag.NewFlagSet("discover", flag.ContinueOnError)
	fs.SetOutput(errOut)
	var org string
	fs.StringVar(&org, "org", "", "org to list template repos for")
	if err := fs.Parse(args); err != nil {
		return 2
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	res, err := d.provider.Discover(ctx, provider.DiscoverOpts{Org: org})
	if err != nil {
		fmt.Fprintln(errOut, "discover error:", err)
		return 1
	}

	if len(res.Orgs) == 0 {
		fmt.Fprintln(out, "Organizations (0): none found")
	} else {
		fmt.Fprintf(out, "Organizations (%d):\n", len(res.Orgs))
		for _, o := range res.Orgs {
			fmt.Fprintf(out, "  %s\n", o)
		}
	}

	if len(res.Enterprises) == 0 {
		fmt.Fprintln(out, "Enterprises (0): none found")
	} else {
		fmt.Fprintf(out, "Enterprises (%d):\n", len(res.Enterprises))
		for _, e := range res.Enterprises {
			fmt.Fprintf(out, "  %-30s  %s\n", e.Slug, e.Name)
		}
	}

	if org != "" {
		if len(res.TemplateRepos) == 0 {
			fmt.Fprintf(out, "Template repos in %s (0): none found\n", org)
		} else {
			fmt.Fprintf(out, "Template repos in %s (%d):\n", org, len(res.TemplateRepos))
			for _, r := range res.TemplateRepos {
				fmt.Fprintf(out, "  %s\n", r)
			}
		}
	}

	for _, n := range res.Notes {
		fmt.Fprintf(out, "note: %s\n", n)
	}

	fmt.Fprintln(out, "hint: set GITHUB_OWNER / GITHUB_ENTERPRISE_SLUG / GH_TEST_ORG_TEMPLATE_REPOSITORY (or put them in your --env-file)")
	return 0
}
