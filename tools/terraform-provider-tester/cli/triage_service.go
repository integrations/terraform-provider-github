package cli

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/github/terraform-provider-tester/engine"
	"github.com/github/terraform-provider-tester/internal/redact"
)

type triageStateResult struct {
	Failures       []engine.PersistFailure
	CacheAvailable bool
	Mode           string
}

type triageStateService struct {
	root            string
	statePath       string
	red             *redact.Redactor
	registryFactory func(triageOptions) (issueRegistry, error)
	filerFactory    func(string) (issueFiler, error)
	getenv          func(string) string
}

func (s triageStateService) apply(ctx context.Context, st *engine.State, opts triageOptions, forceReconstruct bool) (triageStateResult, error) {
	root := s.root
	if root == "" {
		root = filepath.Dir(s.statePath)
	}
	failures := st.Failures
	if forceReconstruct || len(failures) == 0 {
		failures = reconstructFailures(root, *st)
	}

	cacheAvailable := false
	if opts.KnownIssuesPath != "" {
		cacheAvailable = true
		if _, err := os.Stat(opts.KnownIssuesPath); err != nil {
			if !os.IsNotExist(err) {
				return triageStateResult{}, err
			}
			cacheAvailable = false
		}
	}

	updated, err := processTriageIssues(ctx, failures, st.Mode, root, opts, s.red, s.registryFactory, s.filerFactory, s.getenv)
	if err != nil {
		return triageStateResult{}, err
	}
	st.Failures = updated
	return triageStateResult{Failures: updated, CacheAvailable: cacheAvailable, Mode: st.Mode}, nil
}

func (s triageStateService) Refresh(ctx context.Context, opts triageOptions, forceReconstruct bool) (result triageStateResult, err error) {
	release, err := engine.AcquireLock(s.statePath + ".lock")
	if err != nil {
		return triageStateResult{}, fmt.Errorf("acquiring lock: %w", err)
	}
	defer func() {
		if releaseErr := release(); releaseErr != nil {
			releaseErr = fmt.Errorf("releasing lock: %w", releaseErr)
			if err != nil {
				err = errors.Join(err, releaseErr)
				return
			}
			err = releaseErr
		}
	}()

	st, err := engine.Load(s.statePath)
	if err != nil {
		return triageStateResult{}, fmt.Errorf("loading state: %w", err)
	}
	result, err = s.apply(ctx, &st, opts, forceReconstruct)
	if err != nil {
		return triageStateResult{}, fmt.Errorf("triage issues: %w", err)
	}
	if err := st.Save(s.statePath); err != nil {
		return triageStateResult{}, fmt.Errorf("saving state: %w", err)
	}
	return result, nil
}

func persistSuiteFailureLogs(root string, st *engine.State, res engine.RunResult, red *redact.Redactor) error {
	if red == nil {
		red = redact.New(nil)
	}
	st.BuildFailed = res.BuildFailed
	st.PreRunFailed = res.PreRunFailed

	failDir := filepath.Join(root, ".pulsar-failures")
	if !res.BuildFailed {
		st.BuildLog = ""
	}
	if !res.PreRunFailed {
		st.PreRunLog = ""
	}
	if res.BuildFailed || res.PreRunFailed {
		if err := os.MkdirAll(failDir, 0o700); err != nil {
			return fmt.Errorf("creating failure log dir: %w", err)
		}
		var errs []error
		if res.BuildFailed {
			st.BuildLog = filepath.Join(failDir, "build.log")
			if err := os.WriteFile(st.BuildLog, []byte(strings.Join(red.Lines(res.BuildOutput), "")), 0o600); err != nil {
				errs = append(errs, fmt.Errorf("writing build failure log: %w", err))
			}
		}
		if res.PreRunFailed {
			st.PreRunLog = filepath.Join(failDir, "pre-run.log")
			if err := os.WriteFile(st.PreRunLog, []byte(strings.Join(red.Lines(res.PreRunOutput), "")), 0o600); err != nil {
				errs = append(errs, fmt.Errorf("writing pre-run failure log: %w", err))
			}
		}
		return errors.Join(errs...)
	}

	st.BuildLog = ""
	st.PreRunLog = ""
	var errs []error
	for _, stale := range []string{"build.log", "pre-run.log"} {
		if err := os.Remove(filepath.Join(failDir, stale)); err != nil && !os.IsNotExist(err) {
			errs = append(errs, fmt.Errorf("removing suite failure log: %w", err))
		}
	}
	return errors.Join(errs...)
}
