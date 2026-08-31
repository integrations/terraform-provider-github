package cli

import (
	"bytes"
	"context"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"
	"time"

	"github.com/github/terraform-provider-tester/engine"
	"github.com/github/terraform-provider-tester/internal/redact"
)

func reportServicePaths(t *testing.T) (root, configDir, statePath string) {
	t.Helper()
	scratch := cliScratchDir(t)
	root = filepath.Join(scratch, "provider")
	configDir = filepath.Join(scratch, "config")
	if err := os.MkdirAll(root, 0o700); err != nil {
		t.Fatalf("make root: %v", err)
	}
	if err := os.MkdirAll(configDir, 0o700); err != nil {
		t.Fatalf("make config: %v", err)
	}
	statePath = filepath.Join(root, ".pulsar-state.json")
	return root, configDir, statePath
}

func writeReportState(t *testing.T, statePath, buildLog string) engine.State {
	t.Helper()
	st := engine.State{
		Provider:    "test",
		Mode:        "organization",
		RunAt:       time.Date(2026, 7, 9, 16, 0, 0, 0, time.UTC),
		BuildFailed: true,
		BuildLog:    buildLog,
		Results: []engine.PersistResult{{
			Package: "./github",
			Test:    "TestAccSecret",
			Status:  "fail",
		}},
	}
	if err := st.Save(statePath); err != nil {
		t.Fatalf("save state: %v", err)
	}
	return st
}

func TestReportExportServiceWritesBothFormatsOutsideRoot(t *testing.T) {
	root, configDir, statePath := reportServicePaths(t)
	const secret = "ghp_REPORT_SECRET"
	st := writeReportState(t, statePath, filepath.Join(root, ".pulsar-failures", "build-"+secret+".log"))
	now := func() time.Time {
		return time.Date(2026, 7, 9, 16, 30, 0, 0, time.UTC)
	}

	svc := reportExportService{
		statePath: statePath,
		groups:    []engine.Group{{Name: "tests", Tests: []string{"TestAccSecret"}}},
		red:       redact.New([]string{secret}),
		userConfigDir: func() (string, error) {
			return configDir, nil
		},
		now: now,
	}

	got, err := svc.Export(context.Background())
	if err != nil {
		t.Fatalf("Export() error = %v", err)
	}

	expectedDir := filepath.Join(configDir, "terraform-provider-tester", "reports")
	if filepath.Dir(got.MarkdownPath) != expectedDir {
		t.Fatalf("MarkdownPath dir = %q, want %q", filepath.Dir(got.MarkdownPath), expectedDir)
	}
	if filepath.Dir(got.HTMLPath) != expectedDir {
		t.Fatalf("HTMLPath dir = %q, want %q", filepath.Dir(got.HTMLPath), expectedDir)
	}
	if rel, err := filepath.Rel(root, got.MarkdownPath); err != nil || !strings.HasPrefix(rel, "..") {
		t.Fatalf("MarkdownPath = %q, want outside repo root %q (rel=%q err=%v)", got.MarkdownPath, root, rel, err)
	}
	if rel, err := filepath.Rel(root, got.HTMLPath); err != nil || !strings.HasPrefix(rel, "..") {
		t.Fatalf("HTMLPath = %q, want outside repo root %q (rel=%q err=%v)", got.HTMLPath, root, rel, err)
	}
	if _, err := os.Stat(filepath.Join(root, filepath.Base(got.MarkdownPath))); !os.IsNotExist(err) {
		t.Fatalf("provider root markdown artifact exists: err=%v", err)
	}
	if _, err := os.Stat(filepath.Join(root, filepath.Base(got.HTMLPath))); !os.IsNotExist(err) {
		t.Fatalf("provider root html artifact exists: err=%v", err)
	}
	if got.RunAt != st.RunAt {
		t.Fatalf("RunAt = %v, want %v", got.RunAt, st.RunAt)
	}
	if got.TestCount != len(st.Results) {
		t.Fatalf("TestCount = %d, want %d", got.TestCount, len(st.Results))
	}

	for _, tc := range []struct {
		name string
		path string
		perm os.FileMode
	}{
		{name: "dir", path: expectedDir, perm: 0o700},
		{name: "markdown", path: got.MarkdownPath, perm: 0o600},
		{name: "html", path: got.HTMLPath, perm: 0o600},
	} {
		info, err := os.Stat(tc.path)
		if err != nil {
			t.Fatalf("stat %s: %v", tc.name, err)
		}
		if got := info.Mode().Perm(); got != tc.perm {
			t.Fatalf("%s perms = %#o, want %#o", tc.name, got, tc.perm)
		}
	}
}

func TestReportExportServiceUsesUTCNamesAndRedacts(t *testing.T) {
	root, configDir, statePath := reportServicePaths(t)
	const secret = "ghp_REPORT_SECRET"
	writeReportState(t, statePath, filepath.Join(root, ".pulsar-failures", "build-"+secret+".log"))
	now := func() time.Time {
		return time.Date(2026, 7, 9, 16, 30, 0, 0, time.UTC)
	}

	svc := reportExportService{
		statePath: statePath,
		groups:    []engine.Group{{Name: "tests", Tests: []string{"TestAccSecret"}}},
		red:       redact.New([]string{secret}),
		userConfigDir: func() (string, error) {
			return configDir, nil
		},
		now: now,
	}

	got, err := svc.Export(context.Background())
	if err != nil {
		t.Fatalf("Export() error = %v", err)
	}
	if filepath.Base(got.MarkdownPath) != "20260709T163000Z-report.md" {
		t.Fatalf("MarkdownPath base = %q", filepath.Base(got.MarkdownPath))
	}
	if filepath.Base(got.HTMLPath) != "20260709T163000Z-report.html" {
		t.Fatalf("HTMLPath base = %q", filepath.Base(got.HTMLPath))
	}

	for _, path := range []string{got.MarkdownPath, got.HTMLPath} {
		data, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("read %s: %v", path, err)
		}
		text := string(data)
		if strings.Contains(text, secret) {
			t.Fatalf("report leaked secret in %s: %s", path, text)
		}
		if !strings.Contains(text, "***REDACTED***") {
			t.Fatalf("report missing redaction marker in %s: %s", path, text)
		}
	}
}

func TestReportExportServicePreservesExistingPairOnTimestampCollision(t *testing.T) {
	root, configDir, statePath := reportServicePaths(t)
	writeReportState(t, statePath, filepath.Join(root, ".pulsar-failures", "build-first.log"))
	now := func() time.Time {
		return time.Date(2026, 7, 9, 16, 30, 0, 0, time.UTC)
	}
	groups := []engine.Group{{Name: "tests", Tests: []string{"TestAccSecret"}}}

	first := reportExportService{
		statePath: statePath,
		groups:    groups,
		red:       redact.New(nil),
		userConfigDir: func() (string, error) {
			return configDir, nil
		},
		now: now,
	}

	got, err := first.Export(context.Background())
	if err != nil {
		t.Fatalf("first Export() error = %v", err)
	}
	originalMarkdown, err := os.ReadFile(got.MarkdownPath)
	if err != nil {
		t.Fatalf("read original markdown: %v", err)
	}
	originalHTML, err := os.ReadFile(got.HTMLPath)
	if err != nil {
		t.Fatalf("read original html: %v", err)
	}

	scratch := filepath.Dir(root)
	secondRoot := filepath.Join(scratch, "provider-other")
	if err := os.MkdirAll(secondRoot, 0o700); err != nil {
		t.Fatalf("make second root: %v", err)
	}
	secondStatePath := filepath.Join(secondRoot, ".pulsar-state.json")
	writeReportState(t, secondStatePath, filepath.Join(secondRoot, ".pulsar-failures", "build-second.log"))

	second := reportExportService{
		statePath: secondStatePath,
		groups:    groups,
		red:       redact.New(nil),
		userConfigDir: func() (string, error) {
			return configDir, nil
		},
		now: now,
	}

	_, err = second.Export(context.Background())
	if err == nil {
		t.Fatal("second Export() error = nil, want collision error")
	}
	if !strings.Contains(err.Error(), filepath.Base(got.MarkdownPath)) {
		t.Fatalf("second Export() error = %q, want mention of %q", err, filepath.Base(got.MarkdownPath))
	}

	currentMarkdown, err := os.ReadFile(got.MarkdownPath)
	if err != nil {
		t.Fatalf("read current markdown: %v", err)
	}
	if !bytes.Equal(currentMarkdown, originalMarkdown) {
		t.Fatalf("markdown changed on collision\noriginal:\n%s\ncurrent:\n%s", originalMarkdown, currentMarkdown)
	}
	currentHTML, err := os.ReadFile(got.HTMLPath)
	if err != nil {
		t.Fatalf("read current html: %v", err)
	}
	if !bytes.Equal(currentHTML, originalHTML) {
		t.Fatalf("html changed on collision\noriginal:\n%s\ncurrent:\n%s", originalHTML, currentHTML)
	}

	reportDir := filepath.Join(configDir, "terraform-provider-tester", "reports")
	entries, err := os.ReadDir(reportDir)
	if err != nil {
		t.Fatalf("ReadDir(%q) error = %v", reportDir, err)
	}
	var gotNames []string
	for _, entry := range entries {
		gotNames = append(gotNames, entry.Name())
	}
	slices.Sort(gotNames)
	wantNames := []string{
		filepath.Base(got.MarkdownPath),
		filepath.Base(got.HTMLPath),
	}
	slices.Sort(wantNames)
	if !slices.Equal(gotNames, wantNames) {
		t.Fatalf("report dir entries = %v, want %v", gotNames, wantNames)
	}
}
