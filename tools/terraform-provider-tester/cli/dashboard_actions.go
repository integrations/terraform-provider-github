package cli

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"time"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/github/terraform-provider-tester/engine"
	"github.com/github/terraform-provider-tester/internal/redact"
	"github.com/github/terraform-provider-tester/tui"
)

func (db *dashboard) beginTrackedOperation() func(func()) {
	db.wg.Add(1)
	return func(f func()) {
		defer db.wg.Done()
		f()
	}
}

func (db *dashboard) waitForActiveOperations() {
	db.wg.Wait()
}

func (db *dashboard) beginOperation(ctx context.Context, name string) bool {
	if !db.opGate.CompareAndSwap(false, true) {
		db.sendIfActive(ctx, tui.OperationRejectedMsg{
			Op:  name,
			Err: errors.New("another operation is in progress"),
		})
		return false
	}
	db.sendIfActive(ctx, tui.OperationStartedMsg{Name: name})
	return true
}

func (db *dashboard) endOperation() {
	db.opGate.Store(false)
}

func (db *dashboard) sendOperationError(ctx context.Context, name string, err error) {
	db.cfgMu.Lock()
	red := db.red
	db.cfgMu.Unlock()
	if red != nil && err != nil {
		err = errors.New(red.String(err.Error()))
	}
	db.sendIfActive(ctx, tui.OperationErrMsg{Op: name, Err: err})
}

func (db *dashboard) knownIssuesPath() (string, error) {
	userConfigDir := db.userConfigDir
	if userConfigDir == nil {
		userConfigDir = os.UserConfigDir
	}
	base, err := userConfigDir()
	if err != nil {
		return "", err
	}
	return confinedArtifactPath(db.root, filepath.Join(base, "terraform-provider-tester", "known-issues.yaml"), "known-issues cache")
}

func (db *dashboard) triageService(red *redact.Redactor) triageStateService {
	return triageStateService{
		root:      db.root,
		statePath: db.statePath,
		red:       red,
		getenv:    db.getenv,
	}
}

func (db *dashboard) issueService(red *redact.Redactor) (issueService, error) {
	cachePath, err := db.knownIssuesPath()
	if err != nil {
		return issueService{}, err
	}
	return issueService{
		root:            db.root,
		statePath:       db.statePath,
		issuesRepo:      defaultIssuesRepo,
		knownIssuesPath: cachePath,
		red:             red,
		registryFactory: db.newIssueRegistry,
		filerFactory:    db.newIssueFiler,
		getenv:          dashboardGetenv(db),
	}, nil
}

func cacheFileExists(path string) bool {
	if path == "" {
		return false
	}
	_, err := os.Stat(path)
	return err == nil
}

func preserveIssueMetadata(prev, next []engine.PersistFailure) []engine.PersistFailure {
	byFingerprint := make(map[string]engine.PersistFailure, len(prev))
	for _, failure := range prev {
		if failure.Fingerprint != "" {
			byFingerprint[failure.Fingerprint] = failure
		}
	}
	for i := range next {
		previous, ok := byFingerprint[next[i].Fingerprint]
		if !ok {
			continue
		}
		if next[i].KnownIssue == 0 {
			next[i].KnownIssue = previous.KnownIssue
		}
		if next[i].IssueAction == "" {
			next[i].IssueAction = previous.IssueAction
		}
	}
	return next
}

func (db *dashboard) handleRefreshTriageIntent(ctx context.Context) {
	if !db.beginOperation(ctx, "triage") {
		return
	}
	defer db.endOperation()

	db.cfgMu.Lock()
	red := db.red
	db.cfgMu.Unlock()
	cachePath, err := db.knownIssuesPath()
	if err != nil {
		db.sendOperationError(ctx, "triage", err)
		return
	}
	result, err := db.triageService(red).Refresh(ctx, triageOptions{
		IssuesRepo:         defaultIssuesRepo,
		KnownIssuesPath:    cachePath,
		KnownIssuesOffline: true,
	}, false)
	if err != nil {
		db.sendOperationError(ctx, "triage", err)
		return
	}
	db.sendIfActive(ctx, tui.TriageLoadedMsg{Failures: result.Failures, CacheAvailable: result.CacheAvailable})
}

func (db *dashboard) handleSyncKnownIssuesIntent(ctx context.Context) {
	if !db.beginOperation(ctx, "known-issues-sync") {
		return
	}
	defer db.endOperation()

	db.cfgMu.Lock()
	red := db.red
	db.cfgMu.Unlock()
	cachePath, err := db.knownIssuesPath()
	if err != nil {
		db.sendOperationError(ctx, "known-issues-sync", err)
		return
	}
	syncResult, err := newKnownIssueSyncService(defaultIssuesRepo, cachePath, db.newKnownIssueLister).Sync(ctx)
	if err != nil {
		db.sendOperationError(ctx, "known-issues-sync", err)
		return
	}
	triageResult, err := db.triageService(red).Refresh(ctx, triageOptions{
		IssuesRepo:         defaultIssuesRepo,
		KnownIssuesPath:    cachePath,
		KnownIssuesOffline: true,
	}, false)
	if err != nil {
		db.sendOperationError(ctx, "known-issues-sync", err)
		return
	}
	db.sendIfActive(ctx, tui.KnownIssueSyncDoneMsg{
		Count:          syncResult.Count,
		CachePath:      syncResult.CachePath,
		Failures:       triageResult.Failures,
		CacheAvailable: triageResult.CacheAvailable,
	})
}

func (db *dashboard) handleExportReportIntent(ctx context.Context) {
	if !db.beginOperation(ctx, "export") {
		return
	}
	defer db.endOperation()

	db.cfgMu.Lock()
	red := db.red
	groups := append([]engine.Group(nil), db.groups...)
	db.cfgMu.Unlock()
	result, err := reportExportService{
		statePath:     db.statePath,
		groups:        groups,
		red:           red,
		userConfigDir: db.userConfigDir,
		now:           dashboardNow(db),
	}.Export(ctx)
	if err != nil {
		db.sendOperationError(ctx, "export", err)
		return
	}
	db.sendIfActive(ctx, tui.ReportExportDoneMsg{MarkdownPath: result.MarkdownPath, HTMLPath: result.HTMLPath})
}

func (db *dashboard) handleListOrphansIntent(ctx context.Context) {
	if !db.beginOperation(ctx, "orphans") {
		return
	}
	defer db.endOperation()

	db.cfgMu.Lock()
	mode := db.mode
	db.cfgMu.Unlock()
	snapshot, err := (cleanupService{prov: db.prov, mode: mode, getenv: dashboardGetenv(db)}).List(ctx)
	if err != nil {
		db.sendOperationError(ctx, "orphans", err)
		return
	}
	db.sendIfActive(ctx, tui.OrphansListedMsg{Owner: snapshot.Owner, Resources: snapshot.Resources})
}

func (db *dashboard) handleConfirmSweepIntent(ctx context.Context, msg tui.ConfirmSweepIntent) {
	if !db.beginOperation(ctx, "sweep") {
		return
	}
	defer db.endOperation()

	db.cfgMu.Lock()
	mode := db.mode
	db.cfgMu.Unlock()
	result, err := (cleanupService{prov: db.prov, mode: mode, getenv: dashboardGetenv(db)}).Sweep(ctx, sweepRequest{
		Owner: msg.Owner, Phrase: msg.Phrase, Resources: msg.Resources,
	})
	if err != nil {
		if result.attempted {
			db.sendIfActive(ctx, tui.SweepDoneMsg{
				Remaining: result.Remaining, SnapshotChanged: result.SnapshotChanged,
				Failed: true, ResidualUnknown: result.residualUnknown,
			})
		}
		db.sendOperationError(ctx, "sweep", err)
		return
	}
	db.sendIfActive(ctx, tui.SweepDoneMsg{Remaining: result.Remaining, SnapshotChanged: result.SnapshotChanged})
}

func (db *dashboard) handlePreviewFileIssueIntent(ctx context.Context, msg tui.PreviewFileIssueIntent) {
	if !db.beginOperation(ctx, "issue-preview") {
		return
	}
	defer db.endOperation()

	db.cfgMu.Lock()
	red := db.red
	db.cfgMu.Unlock()
	svc, err := db.issueService(red)
	if err != nil {
		db.sendOperationError(ctx, "issue-preview", err)
		return
	}
	result, err := svc.Preview(msg.Fingerprint)
	if err != nil {
		db.sendOperationError(ctx, "issue-preview", err)
		return
	}
	db.sendIfActive(ctx, tui.IssuePreviewMsg{
		Fingerprint: result.Fingerprint, ShortFingerprint: result.ShortFingerprint,
		IssuesRepo: result.IssuesRepo, Title: result.Title, Labels: result.Labels,
		Classification: result.Classification,
	})
}

func (db *dashboard) handleConfirmFileIssueIntent(ctx context.Context, msg tui.ConfirmFileIssueIntent) {
	if !db.beginOperation(ctx, "issue-file") {
		return
	}
	defer db.endOperation()

	db.cfgMu.Lock()
	red := db.red
	db.cfgMu.Unlock()
	svc, err := db.issueService(red)
	if err != nil {
		db.sendOperationError(ctx, "issue-file", err)
		return
	}
	result, err := svc.FileOne(ctx, msg.Fingerprint, msg.Phrase)
	if err != nil {
		db.sendOperationError(ctx, "issue-file", err)
		return
	}
	db.sendIfActive(ctx, tui.IssueFiledMsg{
		Fingerprint: result.Fingerprint, IssueNumber: result.IssueNumber,
		Action: result.Action, Failures: result.Failures,
	})
}

func dashboardNow(db *dashboard) func() time.Time {
	if db.now != nil {
		return db.now
	}
	return time.Now
}

func (db *dashboard) dispatchIntent(ctx context.Context, msg tea.Msg) {
	switch m := msg.(type) {
	case tui.SetEnvVarIntent, tui.PreflightIntent:
		db.handleConfigIntent(ctx, msg)
	case tui.RefreshTriageIntent:
		db.handleRefreshTriageIntent(ctx)
	case tui.SyncKnownIssuesIntent:
		db.handleSyncKnownIssuesIntent(ctx)
	case tui.ExportReportIntent:
		db.handleExportReportIntent(ctx)
	case tui.ListOrphansIntent:
		db.handleListOrphansIntent(ctx)
	case tui.ConfirmSweepIntent:
		db.handleConfirmSweepIntent(ctx, m)
	case tui.PreviewFileIssueIntent:
		db.handlePreviewFileIssueIntent(ctx, m)
	case tui.ConfirmFileIssueIntent:
		db.handleConfirmFileIssueIntent(ctx, m)
	default:
		if !db.beginOperation(ctx, "run") {
			return
		}
		defer db.endOperation()
		db.runIntent(ctx, msg)
	}
}
