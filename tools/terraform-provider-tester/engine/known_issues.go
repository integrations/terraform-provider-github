package engine

import (
	"context"
	"errors"
	"io"
	"os"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

const knownIssueStaleAfter = 7 * 24 * time.Hour

type KnownIssueFile struct {
	Version    int               `yaml:"version" json:"version"`
	Source     string            `yaml:"source,omitempty" json:"source,omitempty"`
	IssuesRepo string            `yaml:"issues_repo,omitempty" json:"issues_repo,omitempty"`
	Label      string            `yaml:"label,omitempty" json:"label,omitempty"`
	SyncedAt   *time.Time        `yaml:"synced_at,omitempty" json:"synced_at,omitempty"`
	Entries    []KnownIssueEntry `yaml:"entries,omitempty" json:"entries"`
}

type KnownIssueEntry struct {
	Fingerprint string     `yaml:"fingerprint" json:"fingerprint"`
	Short       string     `yaml:"short,omitempty" json:"short,omitempty"`
	Issue       int        `yaml:"issue,omitempty" json:"issue,omitempty"`
	IssueURL    string     `yaml:"issue_url,omitempty" json:"issue_url,omitempty"`
	Title       string     `yaml:"title,omitempty" json:"title,omitempty"`
	State       string     `yaml:"state,omitempty" json:"state,omitempty"`
	Mode        string     `yaml:"mode,omitempty" json:"mode,omitempty"`
	Tests       []string   `yaml:"tests,omitempty" json:"tests,omitempty"`
	Modes       []string   `yaml:"modes,omitempty" json:"modes,omitempty"`
	Classes     []string   `yaml:"classes,omitempty" json:"classes,omitempty"`
	Retry       string     `yaml:"retry,omitempty" json:"retry,omitempty"`
	ExpiresAt   *time.Time `yaml:"expires_at,omitempty" json:"expires_at,omitempty"`
	Note        string     `yaml:"note,omitempty" json:"note,omitempty"`
	UpdatedAt   *time.Time `yaml:"updated_at,omitempty" json:"updated_at,omitempty"`
	Origin      string     `yaml:"origin,omitempty" json:"origin,omitempty"`
	Stale       bool       `yaml:"-" json:"stale,omitempty"`
	Expired     bool       `yaml:"-" json:"expired,omitempty"`
	Source      string     `yaml:"-" json:"source,omitempty"`
}

type KnownIssueMatch struct {
	Found    bool
	Issue    int
	Mode     string
	Suppress bool
	Note     string
}

type KnownIssueSource interface {
	ListKnownIssues(context.Context) ([]KnownIssueEntry, error)
}

type KnownIssueRegistry struct {
	CachePath string
	Offline   bool
	Live      KnownIssueSource
	Now       func() time.Time

	loaded          bool
	entries         []KnownIssueEntry
	loadedCachePath string
	loadedOffline   bool
}

func LoadKnownIssuesFile(path string, now time.Time) (*KnownIssueFile, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var file KnownIssueFile
	if err := yaml.Unmarshal(data, &file); err != nil {
		return nil, err
	}
	if file.Version == 0 {
		file.Version = 1
	}
	markKnownIssueDerivedFields(&file, now)
	return &file, nil
}

func SaveKnownIssuesFile(path string, file KnownIssueFile) error {
	if file.Version == 0 {
		file.Version = 1
	}
	if file.Label == "" {
		file.Label = "acctest-failure"
	}
	if file.SyncedAt == nil {
		now := time.Now().UTC()
		file.SyncedAt = &now
	}
	data, err := yaml.Marshal(file)
	if err != nil {
		return err
	}
	return writeAtomic(path, 0o600, func(w io.Writer) error {
		_, err := w.Write(data)
		return err
	})
}

func (r *KnownIssueRegistry) Entries(ctx context.Context) ([]KnownIssueEntry, error) {
	if r.loaded && r.loadedCachePath == r.CachePath && r.loadedOffline == r.Offline {
		return append([]KnownIssueEntry{}, r.entries...), nil
	}
	now := r.now()
	merged := map[string]KnownIssueEntry{}
	ordered := map[string]bool{}
	var order []string
	if r.CachePath != "" {
		file, err := LoadKnownIssuesFile(r.CachePath, now)
		if err != nil {
			if !errors.Is(err, os.ErrNotExist) {
				return nil, err
			}
		} else {
			for _, entry := range file.Entries {
				entry.Source = "cache"
				if entry.Fingerprint == "" {
					continue
				}
				if !ordered[entry.Fingerprint] {
					order = append(order, entry.Fingerprint)
					ordered[entry.Fingerprint] = true
				}
				merged[entry.Fingerprint] = entry
			}
		}
	}
	if !r.Offline && r.Live != nil {
		liveEntries, err := r.Live.ListKnownIssues(ctx)
		if err != nil {
			return nil, err
		}
		for fingerprint, entry := range merged {
			if entry.Origin == "github" {
				delete(merged, fingerprint)
			}
		}
		for _, entry := range liveEntries {
			if entry.Fingerprint == "" {
				continue
			}
			if cached, ok := merged[entry.Fingerprint]; ok && cached.activeLocalOverride() {
				continue
			}
			entry.Source = "github"
			entry.Origin = "github"
			entry.Stale = false
			entry.Expired = entry.expiresBefore(now)
			if !ordered[entry.Fingerprint] {
				order = append(order, entry.Fingerprint)
				ordered[entry.Fingerprint] = true
			}
			if strings.EqualFold(entry.State, "open") || entry.State == "" {
				merged[entry.Fingerprint] = entry
			}
		}
	}
	out := make([]KnownIssueEntry, 0, len(order))
	for _, fp := range order {
		if entry, ok := merged[fp]; ok {
			out = append(out, entry)
		}
	}
	r.loaded = true
	r.loadedCachePath = r.CachePath
	r.loadedOffline = r.Offline
	r.entries = append([]KnownIssueEntry{}, out...)
	return out, nil
}

func (r *KnownIssueRegistry) LookupFingerprint(ctx context.Context, fingerprint, mode string) (KnownIssueMatch, error) {
	if fingerprint == "" {
		return KnownIssueMatch{}, nil
	}
	entries, err := r.Entries(ctx)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return KnownIssueMatch{}, nil
		}
		return KnownIssueMatch{}, err
	}
	for _, entry := range entries {
		if match, ok := entry.match(fingerprint, mode); ok {
			return match, nil
		}
	}
	return KnownIssueMatch{}, nil
}

func (r *KnownIssueRegistry) now() time.Time {
	if r != nil && r.Now != nil {
		return r.Now()
	}
	return time.Now().UTC()
}

func markKnownIssueDerivedFields(file *KnownIssueFile, now time.Time) {
	if now.IsZero() {
		now = time.Now().UTC()
	}
	stale := file.SyncedAt != nil && now.Sub(*file.SyncedAt) > knownIssueStaleAfter
	for i := range file.Entries {
		if file.Entries[i].Origin == "" {
			file.Entries[i].Origin = inferKnownIssueOrigin(file.Source, file.Entries[i])
		}
		file.Entries[i].Stale = stale
		file.Entries[i].Expired = file.Entries[i].expiresBefore(now)
		file.Entries[i].Source = "cache"
		if file.Entries[i].Short == "" && strings.HasPrefix(file.Entries[i].Fingerprint, "sha256:") && len(file.Entries[i].Fingerprint) >= len("sha256:")+16 {
			file.Entries[i].Short = strings.TrimPrefix(file.Entries[i].Fingerprint, "sha256:")[:16]
		}
	}
}

func inferKnownIssueOrigin(fileSource string, entry KnownIssueEntry) string {
	if strings.EqualFold(fileSource, "github") &&
		entry.Issue != 0 &&
		(entry.IssueURL != "" || entry.Title != "" || entry.UpdatedAt != nil) {
		return "github"
	}
	return "local"
}

func (e KnownIssueEntry) activeLocalOverride() bool {
	return e.Origin == "local" &&
		!e.Expired &&
		(e.State == "" || strings.EqualFold(e.State, "open")) &&
		e.suppressesFiling()
}

func (e KnownIssueEntry) matches(fingerprint, mode string) bool {
	match, ok := e.match(fingerprint, mode)
	return ok && match.Suppress
}

func (e KnownIssueEntry) match(fingerprint, mode string) (KnownIssueMatch, bool) {
	if e.Fingerprint != fingerprint || e.Expired {
		return KnownIssueMatch{}, false
	}
	if e.State != "" && !strings.EqualFold(e.State, "open") {
		return KnownIssueMatch{}, false
	}
	if len(e.Modes) > 0 && !containsString(e.Modes, mode) {
		return KnownIssueMatch{}, false
	}
	return KnownIssueMatch{
		Found:    true,
		Issue:    e.Issue,
		Mode:     e.Mode,
		Suppress: e.suppressesFiling(),
		Note:     e.Note,
	}, true
}

func (e KnownIssueEntry) suppressesFiling() bool {
	switch e.Mode {
	case "known-real", "known-flake", "ignore", "investigating":
		return true
	case "":
		return e.Issue != 0
	default:
		return false
	}
}

func (e KnownIssueEntry) expiresBefore(now time.Time) bool {
	return e.ExpiresAt != nil && !e.ExpiresAt.After(now)
}

func containsString(values []string, want string) bool {
	for _, value := range values {
		if value == want {
			return true
		}
	}
	return false
}
