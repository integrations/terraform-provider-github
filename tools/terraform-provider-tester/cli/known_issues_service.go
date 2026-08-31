package cli

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/github/terraform-provider-tester/engine"
	githubprovider "github.com/github/terraform-provider-tester/provider/github"
)

type knownIssueLister interface {
	ListKnownIssues(context.Context) ([]engine.KnownIssueEntry, error)
}

type knownIssueSyncResult struct {
	Count     int
	CachePath string
	SyncedAt  time.Time
}

type knownIssueSyncService struct {
	issuesRepo string
	cachePath  string
	now        func() time.Time
	newLister  func(string) (knownIssueLister, error)
}

type knownIssueListerConfigError struct{ err error }

func (e knownIssueListerConfigError) Error() string { return e.err.Error() }
func (e knownIssueListerConfigError) Unwrap() error { return e.err }

func newKnownIssueSyncService(issuesRepo, cachePath string, newLister func(string) (knownIssueLister, error)) knownIssueSyncService {
	if newLister == nil {
		newLister = func(repo string) (knownIssueLister, error) {
			return githubprovider.NewIssueClientFromEnv(repo)
		}
	}
	return knownIssueSyncService{
		issuesRepo: issuesRepo,
		cachePath:  cachePath,
		now:        time.Now,
		newLister:  newLister,
	}
}

func (s knownIssueSyncService) Sync(ctx context.Context) (knownIssueSyncResult, error) {
	client, err := s.newLister(s.issuesRepo)
	if err != nil {
		return knownIssueSyncResult{}, knownIssueListerConfigError{err: err}
	}
	entries, err := client.ListKnownIssues(ctx)
	if err != nil {
		return knownIssueSyncResult{}, err
	}
	if err := ctx.Err(); err != nil {
		return knownIssueSyncResult{}, err
	}
	for i := range entries {
		entries[i].Origin = "github"
	}
	liveCount := len(entries)
	nowFn := s.now
	if nowFn == nil {
		nowFn = time.Now
	}
	now := nowFn().UTC()
	entries, err = preserveLocalKnownIssueEntries(s.cachePath, entries, now)
	if err != nil {
		return knownIssueSyncResult{}, err
	}
	file := engine.KnownIssueFile{
		Version:    1,
		Source:     "github",
		IssuesRepo: s.issuesRepo,
		Label:      "acctest-failure",
		SyncedAt:   &now,
		Entries:    entries,
	}
	if err := ctx.Err(); err != nil {
		return knownIssueSyncResult{}, err
	}
	if err := engine.SaveKnownIssuesFile(s.cachePath, file); err != nil {
		return knownIssueSyncResult{}, err
	}
	return knownIssueSyncResult{Count: liveCount, CachePath: s.cachePath, SyncedAt: now}, nil
}

func preserveLocalKnownIssueEntries(path string, live []engine.KnownIssueEntry, now time.Time) ([]engine.KnownIssueEntry, error) {
	existing, err := engine.LoadKnownIssuesFile(path, now)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return live, nil
		}
		return nil, err
	}
	index := make(map[string]int, len(live))
	for i, entry := range live {
		if entry.Fingerprint != "" {
			index[entry.Fingerprint] = i
		}
	}
	for _, entry := range existing.Entries {
		if entry.Origin != "local" || entry.Fingerprint == "" {
			continue
		}
		if i, ok := index[entry.Fingerprint]; ok {
			live[i] = entry
			continue
		}
		index[entry.Fingerprint] = len(live)
		live = append(live, entry)
	}
	return live, nil
}
