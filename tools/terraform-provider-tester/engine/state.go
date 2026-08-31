package engine

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/github/terraform-provider-tester/provider"
)

// PersistResult is the state-safe representation of one test result.
// It carries no Output, environment data, or secrets.
type PersistResult struct {
	Test    string  `json:"test"`
	Sub     string  `json:"sub,omitempty"`
	Status  string  `json:"status"`
	Elapsed float64 `json:"elapsed"`
	Package string  `json:"package"`
}

// PersistFailure is the state-safe triage record for one failure. It stores no raw output.
type PersistFailure struct {
	Package          string   `json:"package"`
	Test             string   `json:"test"`
	Sub              string   `json:"sub,omitempty"`
	Status           string   `json:"status"`
	Fingerprint      string   `json:"fingerprint"`
	ShortFingerprint string   `json:"short_fingerprint"`
	Class            string   `json:"class"`
	Canonical        string   `json:"canonical"`
	Retryable        bool     `json:"retryable"`
	Classification   string   `json:"classification"`
	Attempts         int      `json:"attempts"`
	KnownIssue       int      `json:"known_issue,omitempty"`
	IssueAction      string   `json:"issue_action,omitempty"`
	Mode             string   `json:"mode"`
	Reasons          []string `json:"reasons,omitempty"`
	LogPath          string   `json:"log_path,omitempty"`
}

const StateVersion = 2

// State is the on-disk run state used for resume, retry, and flaky detection.
type State struct {
	Version      int                 `json:"version"`
	Provider     string              `json:"provider"`
	Mode         string              `json:"mode"`
	RunAt        time.Time           `json:"run_at"`
	BuildFailed  bool                `json:"build_failed,omitempty"`
	BuildLog     string              `json:"build_log,omitempty"`
	PreRunFailed bool                `json:"pre_run_failed,omitempty"`
	PreRunLog    string              `json:"pre_run_log,omitempty"`
	Results      []PersistResult     `json:"results"`
	History      map[string][]string `json:"history"`
	Failures     []PersistFailure    `json:"failures,omitempty"`
	Plan         *ExecutionPlan      `json:"plan,omitempty"`
	Orphans      *OrphanAccounting   `json:"orphans,omitempty"`
}

// Load reads the state from path. A missing file returns a zero State and
// nil error. Other read or JSON errors are returned as-is.
func Load(path string) (State, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return State{}, nil
		}
		return State{}, err
	}
	var s State
	if err := json.Unmarshal(data, &s); err != nil {
		return State{}, err
	}
	return s, nil
}

// Save atomically writes the state to path. It creates a temp file in the
// same directory, writes indented JSON, then renames it over path. Any error
// before the rename causes the temp file to be removed.
func (s *State) Save(path string) error {
	s.Version = StateVersion

	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}

	dir := filepath.Dir(path)
	f, err := os.CreateTemp(dir, ".pulsar-state-*.tmp")
	if err != nil {
		return err
	}
	tmpPath := f.Name()

	_, writeErr := f.Write(data)
	closeErr := f.Close()

	if writeErr != nil {
		os.Remove(tmpPath)
		return writeErr
	}
	if closeErr != nil {
		os.Remove(tmpPath)
		return closeErr
	}

	if err := os.Rename(tmpPath, path); err != nil {
		os.Remove(tmpPath)
		return err
	}
	return nil
}

// Record maps r to a PersistResult (without Output) and appends it to
// Results. History is updated only for top-level tests (r.Sub == ""),
// capped to the last 10 entries per test name.
func (s *State) Record(r TestResult) {
	s.Results = append(s.Results, PersistResult{
		Test:    r.Name,
		Sub:     r.Sub,
		Status:  r.Status.String(),
		Elapsed: r.Elapsed,
		Package: r.Package,
	})

	if r.Sub == "" {
		if s.History == nil {
			s.History = make(map[string][]string)
		}
		h := append(s.History[r.Name], r.Status.String())
		if len(h) > 10 {
			h = h[len(h)-10:]
		}
		s.History[r.Name] = h
	}
}

// FailedTopLevel returns distinct top-level test names (PersistResult.Test)
// whose Status is "fail", "panic", or "timeout". It scans all Results
// including subtests and returns names in first-seen order.
func (s State) FailedTopLevel() []string {
	var out []string
	seen := make(map[string]bool)
	for _, r := range s.Results {
		if r.Status == provider.StatusFail.String() ||
			r.Status == provider.StatusPanic.String() ||
			r.Status == provider.StatusTimeout.String() {
			if !seen[r.Test] {
				seen[r.Test] = true
				out = append(out, r.Test)
			}
		}
	}
	return out
}

// NotRun returns names from all that have no corresponding PersistResult in
// this run. It preserves the order of the all slice.
func (s State) NotRun(all []string) []string {
	ran := make(map[string]bool, len(s.Results))
	for _, r := range s.Results {
		ran[r.Test] = true
	}
	var out []string
	for _, name := range all {
		if !ran[name] {
			out = append(out, name)
		}
	}
	return out
}

// Flaky reports whether a test is flaky based on its run history. It looks at
// the last lastM entries (or all if fewer). Returns true iff the number of
// failures ("fail", "panic", "timeout") is >= failN and the window contains
// at least one non-failure (i.e. not all-fail and not all-pass).
func Flaky(history []string, failN, lastM int) bool {
	window := history
	if len(window) > lastM {
		window = window[len(window)-lastM:]
	}

	failures := 0
	for _, h := range window {
		if h == "fail" || h == "panic" || h == "timeout" {
			failures++
		}
	}

	return failures >= failN && failures < len(window) && failures > 0
}

// lockWriteGrace is how long an empty or unparseable lock is left alone before
// it is treated as abandoned. It covers the brief window between creating a
// lock file (O_EXCL) and stamping its owner info: a concurrent acquirer that
// reads the file mid-write must not mistake it for a crashed legacy lock and
// stomp a run that just started. A genuinely crashed run's lock ages past this
// quickly; `terraform-provider-tester unlock` is the instant override.
const lockWriteGrace = 5 * time.Second

// lockInfo records who owns a lock file. Stamping the owner lets a crashed
// run's abandoned lock be reclaimed safely instead of bricking every later run.
type lockInfo struct {
	PID        int       `json:"pid"`
	Host       string    `json:"host"`
	AcquiredAt time.Time `json:"acquired_at"`
}

// lockHost returns this machine's hostname, or "unknown" if it cannot be read.
func lockHost() string {
	h, err := os.Hostname()
	if err != nil || h == "" {
		return "unknown"
	}
	return h
}

// processAlive reports whether a process with the given pid currently exists.
// It uses signal 0, which performs error checking (permission and existence)
// without actually delivering a signal. EPERM means the process exists but is
// owned by another user, which still counts as alive.
func processAlive(pid int) bool {
	if pid <= 0 {
		return false
	}
	proc, err := os.FindProcess(pid)
	if err != nil {
		return false
	}
	err = proc.Signal(syscall.Signal(0))
	return err == nil || errors.Is(err, os.ErrPermission)
}

// AcquireLock creates an exclusive lock file at path using O_EXCL and stamps it
// with the current process's owner info. If the file already exists it tries to
// reclaim it when it is abandoned (empty/legacy, or owned by a dead process on
// this host); a lock held by a live owner, or written by another host whose
// liveness cannot be verified, is left in place and an error is returned. The
// returned release function closes and removes the lock file; calling it more
// than once is safe.
func AcquireLock(path string) (release func() error, err error) {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0o600)
	if err != nil {
		if !errors.Is(err, os.ErrExist) {
			return nil, fmt.Errorf("acquiring lock: %w", err)
		}
		reclaimed, rerr := reclaimStaleLock(path)
		if rerr != nil {
			return nil, rerr
		}
		if !reclaimed {
			return nil, fmt.Errorf("lock already held: %s is locked by a running terraform-provider-tester; "+
				"if no run is active, clear it with 'terraform-provider-tester unlock'", path)
		}
		if f, err = os.OpenFile(path, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0o600); err != nil {
			return nil, fmt.Errorf("acquiring lock after reclaiming a stale one: %w", err)
		}
	}

	if b, merr := json.Marshal(lockInfo{PID: os.Getpid(), Host: lockHost(), AcquiredAt: time.Now().UTC()}); merr == nil {
		_, _ = f.Write(b)
	}
	_ = f.Sync()

	var released bool
	release = func() error {
		if released {
			return nil
		}
		released = true
		f.Close()
		return os.Remove(path)
	}
	return release, nil
}

// reclaimStaleLock decides whether an existing lock file is abandoned and, if
// so, removes it. It returns true when the lock was reclaimed (the caller may
// retry the acquire). A lock is stale when it is empty or unparseable (a legacy
// or partially written lock) or when its owner process is gone on this host. A
// lock owned by a live process, or written by a different host whose liveness
// cannot be checked, is left in place (returns false).
func reclaimStaleLock(path string) (bool, error) {
	identity, err := lockIdentityForPath(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return true, nil // vanished between OpenFile and now; let the caller retry
		}
		return false, nil // cannot assess; treat as held
	}

	data, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return true, nil
		}
		return false, nil
	}

	return reclaimStaleLockWithIdentity(path, data, identity)
}

func reclaimStaleLockWithIdentity(path string, data []byte, identity lockIdentity) (bool, error) {
	var info lockInfo
	if len(strings.TrimSpace(string(data))) == 0 || json.Unmarshal(data, &info) != nil {
		// Empty or unparseable: either a crashed legacy lock or one still being
		// stamped by a concurrent acquirer. Only reclaim once it is older than
		// the write grace, so we never stomp a lock created moments ago.
		if fi, statErr := os.Stat(path); statErr == nil && time.Since(fi.ModTime()) < lockWriteGrace {
			return false, nil
		}
		return removeLockFileIfSame(path, identity)
	}
	if info.Host == lockHost() && !processAlive(info.PID) {
		return removeLockFileIfSame(path, identity)
	}
	return false, nil
}

type lockIdentity struct {
	dev uint64
	ino uint64
	sum [sha256.Size]byte
}

func lockIdentityForPath(path string) (lockIdentity, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return lockIdentity{}, err
	}
	st, ok := fi.Sys().(*syscall.Stat_t)
	if !ok {
		return lockIdentity{}, fmt.Errorf("stat %s: unsupported stat type", path)
	}
	// A content fingerprint guards against inode reuse. On Linux, removing a
	// stale lock and creating a replacement at the same path commonly yields the
	// same dev+ino, so dev+ino alone cannot tell a replacement lock apart from
	// the file we judged stale. The bytes the lock holds are unique to it.
	data, err := os.ReadFile(path)
	if err != nil {
		return lockIdentity{}, err
	}
	return lockIdentity{dev: uint64(st.Dev), ino: uint64(st.Ino), sum: sha256.Sum256(data)}, nil
}

func removeLockFileIfSame(path string, want lockIdentity) (bool, error) {
	got, err := lockIdentityForPath(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return true, nil
		}
		return false, fmt.Errorf("checking stale lock %s: %w", path, err)
	}
	if got != want {
		return false, nil
	}
	if err := os.Remove(path); err != nil && !errors.Is(err, os.ErrNotExist) {
		return false, fmt.Errorf("removing stale lock %s: %w", path, err)
	}
	return true, nil
}

// RemoveLock unconditionally deletes a lock file and returns a human-readable
// description of what it removed. It is the implementation of `terraform-provider-tester unlock`:
// an explicit escape hatch for clearing a lock the automatic reclaim logic
// leaves in place (for example one written by another host). Removing an absent
// lock is not an error.
func RemoveLock(path string) (string, error) {
	data, readErr := os.ReadFile(path)
	if errors.Is(readErr, os.ErrNotExist) {
		return "no lock file present", nil
	}
	if err := os.Remove(path); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return "no lock file present", nil
		}
		return "", fmt.Errorf("removing lock %s: %w", path, err)
	}

	var info lockInfo
	if readErr == nil && json.Unmarshal(data, &info) == nil && info.PID != 0 {
		return fmt.Sprintf("removed lock held by pid %d on host %s (acquired %s)",
			info.PID, info.Host, info.AcquiredAt.Format(time.RFC3339)), nil
	}
	return "removed lock", nil
}
