package engine

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/github/terraform-provider-tester/internal/redact"
	"github.com/github/terraform-provider-tester/provider"
)

const maxSanitizedNameLen = 200

// FailureLog identifies one written failure artifact.
type FailureLog struct {
	Test string // top-level test name (TestResult.Name)
	Path string // full path to the written .log file
}

// FailureLogPath returns the path WriteFailures writes for the given pkg and test name.
func FailureLogPath(dir, pkg, test string) string {
	return filepath.Join(dir, sanitizeName(pkg, test)+".log")
}

// WriteFailures writes one redacted log file per failed top-level test in res.
func WriteFailures(dir string, res RunResult, red *redact.Redactor) ([]FailureLog, error) {
	if err := clearFailureLogsForRun(dir, res); err != nil {
		return nil, err
	}

	var failed []TestResult
	for _, tr := range res.Tests {
		if tr.Sub == "" && isFailStatus(tr.Status) {
			failed = append(failed, tr)
		}
	}
	if len(failed) == 0 {
		return nil, nil
	}

	if err := os.MkdirAll(dir, 0o700); err != nil {
		return nil, fmt.Errorf("creating failure log dir: %w", err)
	}

	logs := make([]FailureLog, 0, len(failed))
	for _, tr := range failed {
		lines := append([]string{}, tr.Output...)
		for _, sub := range res.Tests {
			if sub.Package == tr.Package && sub.Name == tr.Name && sub.Sub != "" && isFailStatus(sub.Status) {
				lines = append(lines, sub.Output...)
			}
		}

		logPath := FailureLogPath(dir, tr.Package, tr.Name)
		content := strings.Join(redactLines(red, lines), "")
		if err := os.WriteFile(logPath, []byte(content), 0o600); err != nil {
			return nil, fmt.Errorf("writing failure log %q: %w", logPath, err)
		}
		logs = append(logs, FailureLog{Test: tr.Name, Path: logPath})
	}
	return logs, nil
}

func clearFailureLogsForRun(dir string, res RunResult) error {
	for _, tr := range res.Tests {
		if tr.Sub != "" {
			continue
		}
		logPath := FailureLogPath(dir, tr.Package, tr.Name)
		if err := os.Remove(logPath); err != nil && !os.IsNotExist(err) {
			return fmt.Errorf("removing stale failure log %q: %w", logPath, err)
		}
	}
	return nil
}

func isFailStatus(status provider.Status) bool {
	switch status {
	case provider.StatusFail, provider.StatusPanic, provider.StatusTimeout:
		return true
	default:
		return false
	}
}

func sanitizeName(pkg, test string) string {
	name := path.Base(pkg) + "_" + test
	sanitized := strings.Map(func(r rune) rune {
		if r == '.' || r == '_' || r == '-' ||
			('A' <= r && r <= 'Z') ||
			('a' <= r && r <= 'z') ||
			('0' <= r && r <= '9') {
			return r
		}
		return '_'
	}, name)
	if len(sanitized) <= maxSanitizedNameLen {
		return sanitized
	}
	h := sha256.Sum256([]byte(name))
	hash8 := hex.EncodeToString(h[:])[:8]
	return sanitized[:maxSanitizedNameLen] + "-" + hash8
}
