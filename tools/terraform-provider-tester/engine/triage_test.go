package engine

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/github/terraform-provider-tester/internal/redact"
	"github.com/github/terraform-provider-tester/provider"
)

func TestClassifyFailureDecisions(t *testing.T) {
	retryable := testSignature("sha256:"+strings.Repeat("a", 64), "api/422-leftover-state", true)
	nonRetryable := testSignature("sha256:"+strings.Repeat("b", 64), "api/403-permission", false)
	different := testSignature("sha256:"+strings.Repeat("c", 64), "api/transient-network", true)

	cases := []struct {
		name string
		in   TriageInput
		want string
	}{
		{
			name: "flake confirmed",
			in: TriageInput{Mode: "organization", Signature: retryable, Attempts: []TriageAttempt{
				{Number: 1, Status: "fail", Signature: retryable},
				{Number: 2, Status: "pass"},
			}},
			want: "flake-confirmed",
		},
		{
			name: "flake historical",
			in: TriageInput{Mode: "organization", Signature: retryable, PreviousHistory: []string{"pass", "fail"}, Attempts: []TriageAttempt{
				{Number: 1, Status: "fail", Signature: retryable},
			}, FailN: 1, LastM: 10},
			want: "flake-historical",
		},
		{
			name: "real",
			in: TriageInput{Mode: "organization", Signature: nonRetryable, PreviousHistory: []string{"pass", "fail"}, Attempts: []TriageAttempt{
				{Number: 1, Status: "fail", Signature: nonRetryable},
			}, FailN: 1, LastM: 10},
			want: "real",
		},
		{
			name: "real unstable",
			in: TriageInput{Mode: "organization", Signature: retryable, Attempts: []TriageAttempt{
				{Number: 1, Status: "fail", Signature: retryable},
				{Number: 2, Status: "fail", Signature: different},
			}},
			want: "real-unstable",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := ClassifyFailure(tc.in)
			if got.Classification != tc.want {
				t.Fatalf("classification = %q, want %q; failure=%+v", got.Classification, tc.want, got)
			}
			if got.Fingerprint == "" || got.ShortFingerprint == "" || got.Class == "" || got.Canonical == "" {
				t.Fatalf("persisted failure missing signature metadata: %+v", got)
			}
			if got.Attempts != len(tc.in.Attempts) {
				t.Fatalf("attempts = %d, want %d", got.Attempts, len(tc.in.Attempts))
			}
		})
	}
}

func TestPersistFailureDoesNotStoreSamplesOrSecrets(t *testing.T) {
	secret := "ENV_SECRET_VALUE_456"
	sig := Signature{
		Version: "pulsar-fp-v1", Provider: "github", Package: "./github", Test: "TestAccThing", Status: "fail",
		Class: "generic", Canonical: "assertion failed ***REDACTED***", NormalizedSample: "raw " + secret,
		Fingerprint: "sha256:" + strings.Repeat("d", 64), ShortFingerprint: strings.Repeat("d", 16),
	}
	failure := ClassifyFailure(TriageInput{Mode: "organization", Signature: sig, Attempts: []TriageAttempt{{Number: 1, Status: "fail", Signature: sig}}})
	state := State{Provider: "github", Mode: "organization", Failures: []PersistFailure{failure}}
	blob, err := json.Marshal(state)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(blob), secret) || strings.Contains(string(blob), "NormalizedSample") || strings.Contains(string(blob), "normalized_sample") {
		t.Fatalf("state failure leaked raw sample or secret: %s", blob)
	}
}

func testSignature(fp, class string, retryable bool) Signature {
	short := strings.TrimPrefix(fp, "sha256:")[:16]
	return Signature{
		Version: "pulsar-fp-v1", Provider: "github", Package: "./github", Test: "TestAccThing", Status: "fail",
		Class: class, Canonical: class + " canonical", Fingerprint: fp, ShortFingerprint: short, Retryable: retryable,
	}
}

func TestDeterministicClassReasonStatements(t *testing.T) {
	cases := []struct {
		fixtureFile string
		wantClass   string
		wantProblem string
		wantCause   string
		wantFix     string
	}{
		{
			fixtureFile: "mode-incompatible.txt",
			wantClass:   ClassModeIncompatible,
			wantProblem: "organization mode",
			wantCause:   "GH_TEST_AUTH_MODE=individual",
			wantFix:     "switch to organization mode or skip this test",
		},
		{
			fixtureFile: "capability-unavailable.txt",
			wantClass:   ClassCapabilityUnavailable,
			wantProblem: "endpoint /user/codespaces/secrets/public-key returned 404",
			wantCause:   "feature unavailable or token lacks required access",
			wantFix:     "enable the capability or verify the token has required access",
		},
		{
			fixtureFile: "fixture-missing.txt",
			wantClass:   ClassFixtureMissing,
			wantProblem: "GH_TEST_ORG_TEMPLATE_REPOSITORY",
			wantCause:   "missing or inaccessible",
			wantFix:     "set the environment variable",
		},
		{
			fixtureFile: "api-404-permission-or-feature.txt",
			wantClass:   ClassAPINotFoundPermissionOrFeature,
			wantProblem: "GET /orgs/example/settings returned 404",
			wantCause:   "requires permission or an enabled feature",
			wantFix:     "verify permission, feature enablement, and endpoint availability",
		},
	}
	for _, tc := range cases {
		t.Run(tc.fixtureFile, func(t *testing.T) {
			data, err := os.ReadFile(filepath.Join("testdata", "triage", tc.fixtureFile))
			if err != nil {
				t.Fatalf("read fixture: %v", err)
			}
			res := RunResult{Tests: []TestResult{{Package: "./github", Name: "TestAccThing", Status: provider.StatusFail, Output: []string{string(data)}}}}
			sigs := SignaturesForRun("github", res, redact.New(nil))
			if len(sigs) != 1 {
				t.Fatalf("got %d signatures, want 1", len(sigs))
			}
			sig := sigs[0]
			if sig.Class != tc.wantClass {
				t.Fatalf("class=%q, want %q; canonical=%q", sig.Class, tc.wantClass, sig.Canonical)
			}
			if sig.Retryable {
				t.Fatalf("class %q must not be retryable", sig.Class)
			}
			canon := sig.Canonical
			for _, fragment := range []string{tc.wantProblem, tc.wantCause, tc.wantFix} {
				if !strings.Contains(canon, fragment) {
					t.Errorf("canonical %q missing fragment %q (problem/cause/fix check)", canon, fragment)
				}
			}
		})
	}
}
