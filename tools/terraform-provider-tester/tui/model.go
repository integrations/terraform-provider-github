package tui

import (
	"os"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/github/terraform-provider-tester/engine"
	"github.com/github/terraform-provider-tester/provider"
)

// envField holds one env var entry displayed (and optionally edited) in the
// variable editor overlay. Secret fields are rendered as <set>/<unset> and
// their inputs are never activated.
type envField struct {
	Key      string
	Required bool
	Secret   bool
	input    textinput.Model
}

// buildEditorFields constructs the envField slice for the given env var list.
// Current non-secret values are pre-populated from getenv; secret fields carry
// no value so they can never be echoed.
func buildEditorFields(vars []provider.EnvVar, getenv func(string) string) []envField {
	fields := make([]envField, 0, len(vars))
	for _, ev := range vars {
		ti := textinput.New()
		ti.CharLimit = 512
		ti.Width = 40
		if !ev.Secret && getenv != nil {
			ti.SetValue(getenv(ev.Key))
		}
		fields = append(fields, envField{
			Key:      ev.Key,
			Required: ev.Required,
			Secret:   ev.Secret,
			input:    ti,
		})
	}
	return fields
}

type focusLevel int

const (
	focusGroups focusLevel = iota
	focusTests
	focusLog
)

// section is the active tab/section of the TUI.
type section int

const (
	sectionPreflight section = iota
	sectionGroups
	sectionRun
	sectionTriage
	numSections = 4
)

// resultKey is the upsert key for a test result.
type resultKey struct {
	Name string
	Sub  string
}

// Model holds the complete TUI state.
// It is kept as a value (not pointer) for idiomatic Bubble Tea.
type Model struct {
	// dimensions
	width  int
	height int

	// header context
	provider string
	mode     string
	owner    string
	// version is the build version shown in the header (e.g. "1.4.0" or a short
	// commit). Empty renders the header without a version segment.
	version string

	// rate-limit gauge
	rate RateMsg

	// preflight data
	checks []provider.Check

	// groups and run data
	groups          []engine.Group
	discoveryGroups []engine.Group
	results         []engine.TestResult
	// resultIndex is a lookup map used for upserts; indices into results slice.
	resultIndex map[resultKey]int

	// navigation
	section section
	cursor  int

	// drill-down state for the Groups section
	focus       focusLevel
	groupCursor int
	testCursor  int
	logVP       viewport.Model

	// transient UI state
	status                string
	err                   error
	errOperation          string
	errIsOperationFailure bool
	statusError           bool
	operation             string

	// triage state
	triageFailures       []engine.PersistFailure
	triageCursor         int
	triageDetailActive   bool
	triageCacheAvailable bool
	failuresFirst        bool

	// components
	keys keyMap
	help help.Model

	// glyph set selector: true = ASCII, false = TTY
	ascii bool

	// executor seam: the producer (CLI wiring, PR10d) injects the behavior that
	// turns an intent message into a tea.Cmd. The tui package never contains run
	// logic; it only holds and calls this seam. nil = intents are no-ops.
	exec func(tea.Msg) tea.Cmd

	// clipboard seam: copies a string to the clipboard. Default is OscCopy
	// (OSC 52). Tests inject a fake to avoid writing escape sequences.
	copyFn func(string) error

	// cmdFor seam: returns the human go test command for a test name.
	// Injected by the CLI wiring so the TUI stays free of run configuration.
	cmdFor func(string) string

	// live-run + resume-prompt state (pure; producer drives via messages)
	running bool
	// spinnerFrame is the Pulsar animation counter; advanced by spinnerTickMsg
	// while running, 0 when idle (and in golden tests, which never tick).
	spinnerFrame int
	// liveness layer: all timing derives from the time carried
	// in spinnerTickMsg, so these stay zero in golden tests that never tick.
	runLabel       string        // what is running, shown as "running <label>"
	runStartedAt   time.Time     // anchored on the first tick after RunStartedMsg
	runElapsed     time.Duration // now - runStartedAt, refreshed each tick
	lastActivityAt time.Time     // tick time of the most recent test result
	sawActivity    bool          // a TestUpdateMsg arrived since the last tick
	stalledFor     time.Duration // now - lastActivityAt, refreshed each tick
	resumePrompt   bool
	resumeWhen     time.Time
	resumeFailed   int
	resumeNotRun   int

	// intro "ignition" splash state (see intro.go). introActive gates the
	// splash; introFrame is advanced by introTickMsg. Both stay zero/false in
	// golden tests and ascii mode, so the dashboard renders normally there.
	introActive bool
	introFrame  int

	// mode picker overlay
	pickerActive bool
	pickerCursor int
	pickerModes  []provider.Mode

	// variable editor overlay
	editorActive  bool
	editorCursor  int
	editorFocused bool // true when a textinput inside editorFields is active
	editorFields  []envField

	// cleanup / confirmation / result overlays
	orphanOverlayActive bool
	orphanOwner         string
	orphans             []provider.Resource
	sweepConfirmActive  bool
	sweepInput          textinput.Model
	fileIssueActive     bool
	fileIssuePreview    IssuePreviewMsg
	fileIssueInput      textinput.Model
	resultOverlayActive bool
	resultTitle         string
	resultLines         []string
	resultTone          panelTone

	// env-var seams: envVars is the descriptor list for the current mode;
	// getenv reads the current value of a variable (default os.Getenv).
	// The TUI uses these only when building the editor overlay; it never
	// calls any provider method directly.
	envVars []provider.EnvVar
	getenv  func(string) string
}

// New creates a fresh Model. ascii selects the ASCII glyph set.
// The section starts at Preflight. Init() is a no-op; the CLI wiring (PR10c)
// supplies the initial commands.
func New(providerName, mode, owner string, ascii bool) Model {
	sweepInput := textinput.New()
	sweepInput.CharLimit = 256
	sweepInput.Width = 40
	fileIssueInput := textinput.New()
	fileIssueInput.CharLimit = 256
	fileIssueInput.Width = 40
	return Model{
		provider:       providerName,
		mode:           mode,
		owner:          owner,
		ascii:          ascii,
		section:        sectionPreflight,
		keys:           defaultKeys(),
		help:           help.New(),
		resultIndex:    make(map[resultKey]int),
		copyFn:         OscCopy,
		getenv:         os.Getenv,
		sweepInput:     sweepInput,
		fileIssueInput: fileIssueInput,
		resultTone:     panelSuccess,
	}
}

// Init satisfies tea.Model. When the intro splash is active it kicks the first
// animation tick; otherwise it is a no-op and commands are injected by the CLI
// wiring in PR10c.
func (m Model) Init() tea.Cmd {
	if m.introActive {
		return introTickCmd()
	}
	return nil
}

// WithIntro activates the one-shot "ignition" splash shown when the dashboard
// opens on a color TTY. The CLI wiring calls it only when not in ascii mode.
// Returns a copy so it composes with New(...).WithIntro().
func (m Model) WithIntro() Model {
	m.introActive = true
	return m
}

// WithExec injects the executor seam used to handle intent messages. The CLI
// wiring (PR10d) supplies a closure that runs the engine and streams redacted
// result messages. Returns a copy so it composes with New(...).WithExec(fn).
func (m Model) WithExec(fn func(tea.Msg) tea.Cmd) Model {
	m.exec = fn
	return m
}

// WithVersion sets the build version shown in the header. Returns a copy so it
// composes with New(...).WithVersion(v).
func (m Model) WithVersion(v string) Model {
	m.version = v
	return m
}

// WithClipboard injects a copy function used by the y (copy log) and c (copy cmd)
// keys. Default is OscCopy. Tests inject a fake to avoid writing escape sequences.
func (m Model) WithClipboard(fn func(string) error) Model {
	m.copyFn = fn
	return m
}

// WithCmdFunc injects a function that builds the go test command for a test name.
// The CLI wiring supplies this so the TUI stays free of run configuration.
func (m Model) WithCmdFunc(fn func(string) string) Model {
	m.cmdFor = fn
	return m
}

// WithModes injects the list of available auth modes shown in the mode picker
// overlay (opened with 'm'). The CLI wiring passes prov.Modes(); tests inject
// a fixed slice. Nil or empty disables the mode picker.
func (m Model) WithModes(modes []provider.Mode) Model {
	m.pickerModes = modes
	return m
}

// WithEnvVars injects the env-var descriptors for the current mode. These are
// shown in the variable editor overlay (opened with 'v'). The CLI wiring passes
// prov.EnvFor(mode); tests inject a fixed slice.
func (m Model) WithEnvVars(vars []provider.EnvVar) Model {
	m.envVars = vars
	return m
}

// WithGetenv injects the function used to read current env-var values for
// non-secret fields in the editor overlay. Default is os.Getenv. Tests inject
// a hermetic function so no real environment variables are read.
func (m Model) WithGetenv(fn func(string) string) Model {
	m.getenv = fn
	return m
}
