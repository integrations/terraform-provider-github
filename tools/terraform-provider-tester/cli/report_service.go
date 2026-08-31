package cli

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/github/terraform-provider-tester/engine"
	"github.com/github/terraform-provider-tester/internal/redact"
	"github.com/github/terraform-provider-tester/provider"
)

func runResultFromState(st engine.State) engine.RunResult {
	res := engine.RunResult{
		BuildFailed:  st.BuildFailed,
		PreRunFailed: st.PreRunFailed,
	}
	if st.BuildLog != "" {
		res.BuildOutput = []string{"see " + st.BuildLog + "\n"}
	}
	if st.PreRunLog != "" {
		res.PreRunOutput = []string{"see " + st.PreRunLog + "\n"}
	}
	for _, r := range st.Results {
		status, _ := provider.ParseStatus(r.Status)
		res.Tests = append(res.Tests, engine.TestResult{
			Package: r.Package,
			Name:    r.Test,
			Sub:     r.Sub,
			Status:  status,
			Elapsed: r.Elapsed,
		})
	}
	return res
}

type reportExportResult struct {
	MarkdownPath string
	HTMLPath     string
	RunAt        time.Time
	TestCount    int
}

type reportExportService struct {
	statePath     string
	groups        []engine.Group
	red           *redact.Redactor
	userConfigDir func() (string, error)
	now           func() time.Time
}

func refuseExistingReportPath(path string) error {
	if _, err := os.Lstat(path); err == nil {
		return fmt.Errorf("report export refused: %s already exists", path)
	} else if !os.IsNotExist(err) {
		return fmt.Errorf("checking report destination %s: %w", path, err)
	}
	return nil
}

func (s reportExportService) Export(ctx context.Context) (result reportExportResult, err error) {
	if err = ctx.Err(); err != nil {
		return result, err
	}

	release, err := engine.AcquireLock(s.statePath + ".lock")
	if err != nil {
		return result, err
	}
	defer func() {
		err = errors.Join(err, release())
	}()

	if err = ctx.Err(); err != nil {
		return result, err
	}

	st, err := engine.Load(s.statePath)
	if err != nil {
		return result, err
	}
	if err = ctx.Err(); err != nil {
		return result, err
	}

	userConfigDir := s.userConfigDir
	if userConfigDir == nil {
		userConfigDir = os.UserConfigDir
	}
	base, err := userConfigDir()
	if err != nil {
		return result, err
	}
	providerRoot := filepath.Dir(s.statePath)
	dir, err := confinedArtifactPath(providerRoot, filepath.Join(base, "terraform-provider-tester", "reports"), "report")
	if err != nil {
		return result, err
	}
	if err = os.MkdirAll(dir, 0o700); err != nil {
		return result, err
	}
	if err = os.Chmod(dir, 0o700); err != nil {
		return result, err
	}
	now := s.now
	if now == nil {
		now = time.Now
	}
	reportBase := now().UTC().Format("20060102T150405Z") + "-report"
	reportLockPath := filepath.Join(dir, reportBase+".lock")
	reportRelease, err := engine.AcquireLock(reportLockPath)
	if err != nil {
		return result, err
	}
	defer func() {
		err = errors.Join(err, reportRelease())
	}()

	md := filepath.Join(dir, reportBase+".md")
	html := filepath.Join(dir, reportBase+".html")
	if err = refuseExistingReportPath(md); err != nil {
		return result, err
	}
	if err = refuseExistingReportPath(html); err != nil {
		return result, err
	}
	res := runResultFromState(st)

	if err = ctx.Err(); err != nil {
		return result, err
	}
	if err = engine.ExportMarkdown(md, s.groups, res, s.red); err != nil {
		return result, err
	}
	if err = ctx.Err(); err != nil {
		if removeErr := os.Remove(md); removeErr != nil && !os.IsNotExist(removeErr) {
			return result, errors.Join(err, removeErr)
		}
		return result, err
	}
	if err = engine.ExportHTML(html, s.groups, res, s.red); err != nil {
		if removeErr := os.Remove(md); removeErr != nil && !os.IsNotExist(removeErr) {
			return result, errors.Join(err, removeErr)
		}
		return result, err
	}

	return reportExportResult{
		MarkdownPath: md,
		HTMLPath:     html,
		RunAt:        st.RunAt,
		TestCount:    len(st.Results),
	}, nil
}
