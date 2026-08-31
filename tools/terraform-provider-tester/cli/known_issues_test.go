package cli

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/github/terraform-provider-tester/engine"
)

func TestKnownIssuesAddEditsOnlyNamedYAMLFile(t *testing.T) {
	root := cliScratchDir(t)
	path := filepath.Join(root, "known.yaml")
	fp := "sha256:" + strings.Repeat("1", 64)
	d := deps{
		newIssueFiler: func(string) (issueFiler, error) {
			t.Fatal("known-issues add must not create a GitHub issue client")
			return nil, nil
		},
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{
		"known-issues", "add",
		"--known-issues", path,
		"--fingerprint", fp,
		"--issue", "1234",
		"--mode", "known-real",
		"--test", "TestAccThing",
		"--class", "api/422-leftover-state",
	}, &out, &errOut, d)
	if code != 0 {
		t.Fatalf("exit = %d, want 0; stderr=%s", code, errOut.String())
	}
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read known issues file: %v", err)
	}
	got := string(data)
	for _, want := range []string{fp, "origin: local", "issue: 1234", "mode: known-real", "TestAccThing", "api/422-leftover-state"} {
		if !strings.Contains(got, want) {
			t.Fatalf("file missing %q:\n%s", want, got)
		}
	}
}

func TestKnownIssuesListJSONReadsFileOffline(t *testing.T) {
	root := cliScratchDir(t)
	path := filepath.Join(root, "known.yaml")
	fp := "sha256:" + strings.Repeat("2", 64)
	if err := os.WriteFile(path, []byte(`version: 1
synced_at: "2026-07-08T18:00:00Z"
entries:
  - fingerprint: "`+fp+`"
    issue: 5678
    state: open
    mode: known-real
    tests:
      - TestAccThing
`), 0o600); err != nil {
		t.Fatal(err)
	}

	var out, errOut bytes.Buffer
	code := runWithDeps([]string{"known-issues", "list", "--known-issues", path, "--known-issues-offline", "--json"}, &out, &errOut, deps{})
	if code != 0 {
		t.Fatalf("exit = %d, want 0; stderr=%s", code, errOut.String())
	}
	var payload struct {
		Entries []struct {
			Fingerprint string `json:"fingerprint"`
			Issue       int    `json:"issue"`
		} `json:"entries"`
	}
	if err := json.Unmarshal(out.Bytes(), &payload); err != nil {
		t.Fatalf("list output is not JSON: %v\n%s", err, out.String())
	}
	if len(payload.Entries) != 1 || payload.Entries[0].Fingerprint != fp || payload.Entries[0].Issue != 5678 {
		t.Fatalf("payload entries = %+v", payload.Entries)
	}
}

func TestKnownIssueSyncCommandKeepsExplicitOutBehavior(t *testing.T) {
	t.Run("requires explicit out", func(t *testing.T) {
		d := deps{
			newKnownIssueLister: func(string) (knownIssueLister, error) {
				t.Fatal("known-issues sync without --out must not create a lister")
				return nil, nil
			},
		}

		var out, errOut bytes.Buffer
		code := runKnownIssuesSync(nil, &out, &errOut, d)
		if code != 2 {
			t.Fatalf("exit = %d, want 2; stderr=%s", code, errOut.String())
		}
		if out.Len() != 0 {
			t.Fatalf("stdout = %q, want empty", out.String())
		}
		if !strings.Contains(errOut.String(), "known-issues sync requires --out") {
			t.Fatalf("stderr = %q, want missing --out error", errOut.String())
		}
	})

	t.Run("syncs explicit path and preserves output text", func(t *testing.T) {
		root := cliScratchDir(t)
		cachePath := filepath.Join(root, "known.yaml")
		fingerprint := "sha256:" + strings.Repeat("3", 64)
		lister := &fakeKnownIssueLister{entries: []engine.KnownIssueEntry{{
			Fingerprint: fingerprint,
			Issue:       777,
			State:       "open",
		}}}
		d := deps{
			newKnownIssueLister: func(repo string) (knownIssueLister, error) {
				if repo != defaultIssuesRepo {
					t.Fatalf("repo = %q, want %q", repo, defaultIssuesRepo)
				}
				return lister, nil
			},
		}

		var out, errOut bytes.Buffer
		code := runKnownIssuesSync([]string{"--out", cachePath}, &out, &errOut, d)
		if code != 0 {
			t.Fatalf("exit = %d, want 0; stderr=%s", code, errOut.String())
		}
		if errOut.Len() != 0 {
			t.Fatalf("stderr = %q, want empty", errOut.String())
		}
		wantOut := "synced 1 known issue(s) to " + cachePath + "\n"
		if out.String() != wantOut {
			t.Fatalf("stdout = %q, want %q", out.String(), wantOut)
		}
		data, err := os.ReadFile(cachePath)
		if err != nil {
			t.Fatalf("read synced file: %v", err)
		}
		if !strings.Contains(string(data), fingerprint) {
			t.Fatalf("synced file missing %q:\n%s", fingerprint, string(data))
		}
	})
}

func cliScratchDir(t *testing.T) string {
	t.Helper()
	dir := filepath.Join(".", ".test-scratch", strings.NewReplacer("/", "_", " ", "_").Replace(t.Name()))
	if err := os.RemoveAll(dir); err != nil {
		t.Fatalf("clean scratch: %v", err)
	}
	if err := os.MkdirAll(dir, 0o700); err != nil {
		t.Fatalf("make scratch: %v", err)
	}
	t.Cleanup(func() { _ = os.RemoveAll(dir) })
	return dir
}
