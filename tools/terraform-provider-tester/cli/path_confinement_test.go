package cli

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/github/terraform-provider-tester/engine"
	"github.com/github/terraform-provider-tester/internal/redact"
)

func TestKnownIssuesPathRejectsDirectProviderRootDestination(t *testing.T) {
	scratch := cliScratchDir(t)
	root := filepath.Join(scratch, "provider")
	if err := os.MkdirAll(root, 0o700); err != nil {
		t.Fatalf("mkdir root: %v", err)
	}
	db := &dashboard{root: root, userConfigDir: func() (string, error) { return root, nil }}
	if _, err := db.knownIssuesPath(); err == nil || !strings.Contains(err.Error(), "provider root") {
		t.Fatalf("knownIssuesPath error = %v, want provider-root confinement error", err)
	}
}

func TestKnownIssuesPathRejectsSymlinkIntoProviderRoot(t *testing.T) {
	scratch := cliScratchDir(t)
	root := filepath.Join(scratch, "provider")
	if err := os.MkdirAll(filepath.Join(root, "cache-target"), 0o700); err != nil {
		t.Fatalf("mkdir root target: %v", err)
	}
	link := filepath.Join(scratch, "config-link")
	target, err := filepath.Abs(filepath.Join(root, "cache-target"))
	if err != nil {
		t.Fatalf("abs target: %v", err)
	}
	if err := os.Symlink(target, link); err != nil {
		t.Fatalf("symlink: %v", err)
	}
	db := &dashboard{root: root, userConfigDir: func() (string, error) { return link, nil }}
	if _, err := db.knownIssuesPath(); err == nil || !strings.Contains(err.Error(), "provider root") {
		t.Fatalf("knownIssuesPath symlink error = %v, want provider-root confinement error", err)
	}
}

func TestReportExportRejectsDirectProviderRootDestination(t *testing.T) {
	scratch := cliScratchDir(t)
	root := filepath.Join(scratch, "provider")
	if err := os.MkdirAll(root, 0o700); err != nil {
		t.Fatalf("mkdir root: %v", err)
	}
	statePath := filepath.Join(root, ".pulsar-state.json")
	st := engine.State{Provider: "github"}
	if err := st.Save(statePath); err != nil {
		t.Fatalf("save state: %v", err)
	}
	svc := reportExportService{statePath: statePath, red: redact.New(nil), now: func() time.Time { return time.Date(2026, 7, 9, 16, 30, 0, 0, time.UTC) }, userConfigDir: func() (string, error) { return root, nil }}
	if _, err := svc.Export(context.Background()); err == nil || !strings.Contains(err.Error(), "provider root") {
		t.Fatalf("Export direct-root error = %v, want provider-root confinement error", err)
	}
	if _, err := os.Stat(filepath.Join(root, "terraform-provider-tester")); !os.IsNotExist(err) {
		t.Fatalf("artifact directory was created before rejection: %v", err)
	}
}

func TestReportExportRejectsSymlinkIntoProviderRoot(t *testing.T) {
	scratch := cliScratchDir(t)
	root := filepath.Join(scratch, "provider")
	if err := os.MkdirAll(filepath.Join(root, "reports-target"), 0o700); err != nil {
		t.Fatalf("mkdir root target: %v", err)
	}
	statePath := filepath.Join(root, ".pulsar-state.json")
	st := engine.State{Provider: "github"}
	if err := st.Save(statePath); err != nil {
		t.Fatalf("save state: %v", err)
	}
	link := filepath.Join(scratch, "config-link")
	target, err := filepath.Abs(filepath.Join(root, "reports-target"))
	if err != nil {
		t.Fatalf("abs target: %v", err)
	}
	if err := os.Symlink(target, link); err != nil {
		t.Fatalf("symlink: %v", err)
	}
	svc := reportExportService{statePath: statePath, red: redact.New(nil), now: func() time.Time { return time.Date(2026, 7, 9, 16, 30, 0, 0, time.UTC) }, userConfigDir: func() (string, error) { return link, nil }}
	if _, err := svc.Export(context.Background()); err == nil || !strings.Contains(err.Error(), "provider root") {
		t.Fatalf("Export symlink error = %v, want provider-root confinement error", err)
	}
}

func TestKnownIssuesPathAllowsOutsideProviderRoot(t *testing.T) {
	scratch := cliScratchDir(t)
	root := filepath.Join(scratch, "provider")
	config := filepath.Join(scratch, "config")
	if err := os.MkdirAll(root, 0o700); err != nil {
		t.Fatalf("mkdir root: %v", err)
	}
	db := &dashboard{root: root, userConfigDir: func() (string, error) { return config, nil }}
	got, err := db.knownIssuesPath()
	if err != nil {
		t.Fatalf("knownIssuesPath outside root: %v", err)
	}
	want := filepath.Join(config, "terraform-provider-tester", "known-issues.yaml")
	if got != want {
		t.Fatalf("knownIssuesPath = %q, want %q", got, want)
	}
}
