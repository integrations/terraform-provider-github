package provider

// Mode is a named run mode a provider supports.
type Mode struct {
	Name        string // "organization"
	Description string // human summary for help/preflight
}

// EnvVar describes one environment variable a provider needs in a mode.
type EnvVar struct {
	Key      string // "GITHUB_TOKEN"
	Required bool
	Secret   bool   // value must be redacted
	Doc      string // one-line help
}

// Status is the status of a test or group.
type Status int

const (
	StatusUnknown Status = iota
	StatusRunning
	StatusPass
	StatusFail
	StatusSkip
	StatusPanic
	StatusTimeout
)

// CheckStatus is the outcome of a single preflight check.
type CheckStatus int

const (
	CheckOK   CheckStatus = iota
	CheckWarn             // e.g. fine-grained PAT: scopes UNVERIFIED
	CheckFail
)

// Check is one read-only preflight result.
type Check struct {
	Name   string // stable check name, e.g. "anonymous", "identity", "owner", "scopes", "rate-limit", or "capability/organization"
	Status CheckStatus
	Detail string // redacted, human-readable
	Fix    string // doctor remediation when Status != CheckOK; "" otherwise
}

// PreflightReport is the read-only result of validating a mode can run.
type PreflightReport struct {
	Mode   string
	Checks []Check
}

// TestRequirements holds per-test metadata that the planner may supply.
// Fields use omitempty so that an empty struct marshals to {}.
type TestRequirements struct {
	Modes        []string `json:"modes,omitempty"`        // preferred run modes, e.g. ["individual"]
	Scopes       []string `json:"scopes,omitempty"`       // OAuth/token scopes required, e.g. ["repo"]
	Capabilities []string `json:"capabilities,omitempty"` // env capabilities required, e.g. ["organization"]
	SideEffects  []string `json:"side_effects,omitempty"` // stable resources created, e.g. ["repository"]
}

// Resource is a leaked/orphan resource (read-only listing).
type Resource struct {
	Kind string // "repository","team"
	Name string // "tf-acc-test-abc123"
	URL  string // display metadata only; identity is Kind+Name
}

// SweepOpts controls a confirmed destructive sweep.
type SweepOpts struct {
	Targets   []string   // ["repositories","teams"]
	Resources []Resource // nil = all prefixed resources in Targets; non-nil = exact Kind+Name subset (including empty)
	Confirm   bool       // engine sets true only after the explicit gate

	// ExactResources selects exact-snapshot mode when non-nil. Nil preserves the
	// legacy direct CLI behavior: Sweep discovers current tf-acc-test-* resources
	// matching Targets before deleting. Non-nil means Sweep must not discover or
	// list resources; it may delete only the provided resource entries. A non-nil
	// empty slice is an exact no-op.
	ExactResources []Resource
}

// DiscoveredEnterprise is one enterprise the viewer can access.
type DiscoveredEnterprise struct {
	Slug string
	Name string
}

// DiscoveryResult holds the orgs, enterprises, and template repos the token can reach.
type DiscoveryResult struct {
	Orgs          []string               // org logins (GET /user/orgs)
	Enterprises   []DiscoveredEnterprise // GraphQL viewer.enterprises
	TemplateRepos []string               // only when an org was queried
	Notes         []string               // human "unverified: ..." lines
}

// DiscoverOpts controls what Discover fetches.
type DiscoverOpts struct {
	Org string // empty = just orgs+enterprises; non-empty = also list template repos
}

func (s Status) String() string {
	switch s {
	case StatusRunning:
		return "running"
	case StatusPass:
		return "pass"
	case StatusFail:
		return "fail"
	case StatusSkip:
		return "skip"
	case StatusPanic:
		return "panic"
	case StatusTimeout:
		return "timeout"
	default:
		return "unknown"
	}
}

func ParseStatus(s string) (Status, bool) {
	switch s {
	case "running":
		return StatusRunning, true
	case "pass":
		return StatusPass, true
	case "fail":
		return StatusFail, true
	case "skip":
		return StatusSkip, true
	case "panic":
		return StatusPanic, true
	case "timeout":
		return StatusTimeout, true
	default:
		return StatusUnknown, false
	}
}

func (c CheckStatus) String() string {
	switch c {
	case CheckOK:
		return "ok"
	case CheckWarn:
		return "warn"
	case CheckFail:
		return "fail"
	default:
		return "unknown"
	}
}

func (r PreflightReport) OK() bool {
	for _, c := range r.Checks {
		if c.Status == CheckFail {
			return false
		}
	}
	return true
}
