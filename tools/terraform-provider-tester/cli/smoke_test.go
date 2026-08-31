//go:build smoke

package cli

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/github/terraform-provider-tester/engine"
)

// TestSmokeAnonymousIpRanges drives the real, read-only acceptance test
// TestAccGithubIpRangesDataSource end-to-end through the assembled harness
// (preflight -> discover -> run -> parse -> persist) against live GitHub with
// NO credentials. The provider sources Terraform on demand via
// terraform-plugin-testing + hc-install, so no Terraform install is required.
// The test reads only the public api.github.com/meta endpoint and creates no
// resources, so it needs no Gate-D authorization.
//
// It is gated behind the `smoke` build tag so the default `go test ./...` never
// runs it (it needs network egress). Run it explicitly:
//
//	TPT_PROVIDER_ROOT=/path/to/terraform-provider-github go test -tags smoke -run TestSmokeAnonymousIpRanges ./cli/
func TestSmokeAnonymousIpRanges(t *testing.T) {
	// Force a clean anonymous environment so machine state cannot contaminate
	// the run: a stray GITHUB_TOKEN would flip auth mode and pull a credential
	// into preflight. These must be UNSET (not set empty): the provider rejects
	// a present-but-empty GITHUB_BASE_URL ("base url must not be empty").
	t.Setenv("GH_TEST_AUTH_MODE", "anonymous")
	clearEnv(t,
		"GITHUB_TOKEN",
		"GITHUB_APP_ID",
		"GITHUB_APP_INSTALLATION_ID",
		"GITHUB_APP_PEM_FILE",
		"GITHUB_OWNER",
		"GITHUB_BASE_URL",
	)

	providerRoot := os.Getenv("TPT_PROVIDER_ROOT")
	if providerRoot == "" {
		t.Fatal("TPT_PROVIDER_ROOT must point to a terraform-provider-github checkout")
	}
	root, err := engine.FindRepoRoot(providerRoot)
	if err != nil {
		t.Fatalf("finding provider repo root: %v", err)
	}

	// The run writes state at the repo root; start clean and do not leave the
	// transient files behind in the checkout.
	statePath := filepath.Join(root, ".pulsar-state.json")
	lockPath := statePath + ".lock"
	_ = os.Remove(statePath)
	_ = os.Remove(lockPath)
	t.Cleanup(func() {
		_ = os.Remove(statePath)
		_ = os.Remove(lockPath)
	})

	const target = "TestAccGithubIpRangesDataSource"
	var out, errOut bytes.Buffer
	code := Run([]string{
		"run",
		"--mode", "anonymous",
		"--run", "^" + target + "$",
		"--repo-root", root,
		"--timeout", "10m",
	}, &out, &errOut)

	// A zero exit implies anonymous preflight passed (runRun returns 1 when the
	// preflight report is not OK before any test runs).
	if code != 0 {
		t.Fatalf("harness run exited %d (preflight or run failed)\nstdout:\n%s\nstderr:\n%s",
			code, out.String(), errOut.String())
	}

	// The state file must be written and record the target as a passing
	// top-level test.
	st, err := engine.Load(statePath)
	if err != nil {
		t.Fatalf("loading persisted state: %v", err)
	}
	if len(st.Results) == 0 {
		t.Fatalf("state file recorded no results\nstdout:\n%s", out.String())
	}
	var found bool
	for _, r := range st.Results {
		if r.Test == target && r.Sub == "" {
			found = true
			if !strings.EqualFold(r.Status, "pass") {
				t.Errorf("%s status = %q, want pass", target, r.Status)
			}
		}
	}
	if !found {
		t.Fatalf("state did not record top-level %s; got %+v", target, st.Results)
	}

	// No credential markers should appear in operator-facing output or in the
	// persisted state (anonymous mode sets none; this guards the redaction path
	// end-to-end).
	for label, s := range map[string]string{"stdout": out.String(), "state file": readFile(t, statePath)} {
		for _, marker := range []string{"ghp_", "github_pat_"} {
			if strings.Contains(s, marker) {
				t.Errorf("possible secret marker %q leaked into %s", marker, label)
			}
		}
	}
}

func readFile(t *testing.T, path string) string {
	t.Helper()
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("reading %s: %v", path, err)
	}
	return string(b)
}

// clearEnv unsets the given variables for the duration of the test, restoring
// any prior value on cleanup. Unsetting (rather than setting empty) is required
// because the provider treats a present-but-empty GITHUB_BASE_URL as invalid.
func clearEnv(t *testing.T, keys ...string) {
	t.Helper()
	for _, k := range keys {
		if old, ok := os.LookupEnv(k); ok {
			t.Cleanup(func() { _ = os.Setenv(k, old) })
		}
		_ = os.Unsetenv(k)
	}
}
