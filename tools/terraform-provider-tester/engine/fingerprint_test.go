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

func TestSignaturesForRunClassifiesCanonicalClasses(t *testing.T) {
	cases := []struct {
		name      string
		result    RunResult
		wantClass string
		wantCanon string
		retryable bool
	}{
		{
			name:      "403 rate limit",
			result:    runWithFailedTest("GET https://api.github.com/rate_limit: 403 API rate limit exceeded; X-RateLimit-Reset: 1893456000\n"),
			wantClass: "api/403-rate-limit",
			wantCanon: "403 rate limit",
			retryable: true,
		},
		{
			name:      "403 secondary rate limit",
			result:    runWithFailedTest("403 You have exceeded a secondary rate limit. Retry-After: 60\n"),
			wantClass: "api/403-secondary-rate-limit",
			wantCanon: "403 secondary rate limit",
			retryable: true,
		},
		{
			name:      "403 permission",
			result:    runWithFailedTest("GET https://api.github.com/repos/o/r: 403 Resource not accessible by token []\n"),
			wantClass: "api/403-permission",
			wantCanon: "403 resource not accessible by token",
		},
		{
			name:      "409 conflict",
			result:    runWithFailedTest("PATCH https://api.github.com/repos/o/r: 409 Conflict: repository already exists\n"),
			wantClass: "api/409-conflict",
			wantCanon: "409 conflict repository already exists",
			retryable: true,
		},
		{
			name:      "422 validation",
			result:    runWithFailedTest("POST https://api.github.com/repos: 422 Validation Failed [{Resource:Repository Field:visibility Code:invalid Message:visibility is invalid}]\n"),
			wantClass: "api/422-validation",
			wantCanon: "422 validation visibility is invalid",
		},
		{
			name:      "422 bare already exists validation",
			result:    runWithFailedTest("POST https://api.github.com/repos: 422 Validation Failed: name already exists\n"),
			wantClass: "api/422-validation",
			wantCanon: "422 validation name already exists",
		},
		{
			name:      "422 leftover state",
			result:    runWithFailedTest("POST https://api.github.com/repos: 422 Validation Failed [{Resource:Repository Field:name Code:already_exists Message:name already exists}]\n"),
			wantClass: "api/422-leftover-state",
			wantCanon: "422 leftover state name already exists",
			retryable: true,
		},
		{
			name:      "transient network",
			result:    runWithFailedTest("Get \"https://api.github.com/repos/o/r\": EOF\n"),
			wantClass: "api/transient-network",
			wantCanon: "transient network eof",
			retryable: true,
		},
		{
			name:      "build",
			result:    RunResult{BuildFailed: true, BuildOutput: []string{"# github.com/example/pkg\n", "github/foo.go:17: undefined: nope\n"}},
			wantClass: "go/build",
			wantCanon: "github/foo.go:<line>: undefined: nope",
		},
		{
			name:      "pre run",
			result:    RunResult{PreRunFailed: true, PreRunOutput: []string{"provider TestMain failed: missing GITHUB_OWNER\n"}},
			wantClass: "provider/testmain",
			wantCanon: "provider TestMain failed: missing GITHUB_OWNER",
		},
		{
			name:      "terraform diagnostic",
			result:    runWithFailedTest("Error: expected visibility to be one of [public private]\n\n  with github_repository.repo,\n  on main.tf line 7, in resource \"github_repository\" \"repo\":\n"),
			wantClass: "terraform/diagnostic",
			wantCanon: "Error: expected visibility to be one of [public private]",
		},
		{
			name:      "panic",
			result:    RunResult{Tests: []TestResult{{Package: "./github", Name: "TestAccThing", Status: provider.StatusPanic, Output: []string{"panic: runtime error: invalid memory address\n", "github.com/example/provider.TestAccThing(0x1)\n", "\t/Users/me/provider/github/foo_test.go:44 +0x10\n"}}}},
			wantClass: "panic",
			wantCanon: "panic: runtime error: invalid memory address github.com/example/provider.TestAccThing",
		},
		{
			name:      "timeout",
			result:    RunResult{Tests: []TestResult{{Package: "./github", Name: "TestAccThing", Status: provider.StatusTimeout, Output: []string{"panic: test timed out after 10m0s\n"}}}},
			wantClass: "timeout",
			wantCanon: "test timed out",
		},
		{
			name:      "generic",
			result:    runWithFailedTest("assertion failed: got false want true\n"),
			wantClass: "generic",
			wantCanon: "assertion failed: got false want true",
		},
		{
			name:      "mode incompatible",
			result:    runWithFailedTestFixture(t, "mode-incompatible.txt"),
			wantClass: ClassModeIncompatible,
			wantCanon: "test requires organization mode but GH_TEST_AUTH_MODE=individual; switch to organization mode or skip this test",
			retryable: false,
		},
		{
			name:      "capability unavailable",
			result:    runWithFailedTestFixture(t, "capability-unavailable.txt"),
			wantClass: ClassCapabilityUnavailable,
			wantCanon: "endpoint /user/codespaces/secrets/public-key returned 404; feature unavailable or token lacks required access; enable the capability or verify the token has required access",
			retryable: false,
		},
		{
			name:      "fixture missing",
			result:    runWithFailedTestFixture(t, "fixture-missing.txt"),
			wantClass: ClassFixtureMissing,
			wantCanon: "required fixture GH_TEST_ORG_TEMPLATE_REPOSITORY is missing or inaccessible; set the environment variable to a valid value",
			retryable: false,
		},
		{
			name:      "api 404 permission or feature",
			result:    runWithFailedTestFixture(t, "api-404-permission-or-feature.txt"),
			wantClass: ClassAPINotFoundPermissionOrFeature,
			wantCanon: "GET /orgs/example/settings returned 404; endpoint requires permission or an enabled feature; verify permission, feature enablement, and endpoint availability",
			retryable: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			sigs := SignaturesForRun("github", tc.result, redact.New(nil))
			if len(sigs) != 1 {
				t.Fatalf("got %d signatures, want 1: %+v", len(sigs), sigs)
			}
			sig := sigs[0]
			if sig.Version != "pulsar-fp-v1" || sig.Provider != "github" {
				t.Fatalf("bad signature identity: %+v", sig)
			}
			if sig.Class != tc.wantClass {
				t.Fatalf("class = %q, want %q; canonical=%q sample=%q", sig.Class, tc.wantClass, sig.Canonical, sig.NormalizedSample)
			}
			if sig.Canonical != tc.wantCanon {
				t.Fatalf("canonical = %q, want %q", sig.Canonical, tc.wantCanon)
			}
			if sig.Retryable != tc.retryable {
				t.Fatalf("retryable = %v, want %v", sig.Retryable, tc.retryable)
			}
			if !strings.HasPrefix(sig.Fingerprint, "sha256:") || len(sig.Fingerprint) != len("sha256:")+64 {
				t.Fatalf("bad fingerprint: %q", sig.Fingerprint)
			}
			if sig.ShortFingerprint != strings.TrimPrefix(sig.Fingerprint, "sha256:")[:16] {
				t.Fatalf("short fingerprint = %q does not match %q", sig.ShortFingerprint, sig.Fingerprint)
			}
		})
	}
}

func TestSignaturesForRunClassifiesDeterministicPreRunModeMismatch(t *testing.T) {
	result := parseRunFixture(t, "prerun_mode_incompatible.ndjson")
	if !result.PreRunFailed {
		t.Fatalf("PreRunFailed = false, want true: %+v", result)
	}
	if result.BuildFailed {
		t.Fatalf("BuildFailed = true, want false: %+v", result)
	}
	if len(result.Tests) != 0 {
		t.Fatalf("Tests = %d, want 0 for real pre-run failure", len(result.Tests))
	}

	sigs := SignaturesForRun("github", result, redact.New(nil))
	if len(sigs) != 1 {
		t.Fatalf("got %d signatures, want 1: %+v", len(sigs), sigs)
	}
	sig := sigs[0]
	if sig.Class != ClassModeIncompatible {
		t.Fatalf("class = %q, want %q; canonical=%q sample=%q", sig.Class, ClassModeIncompatible, sig.Canonical, sig.NormalizedSample)
	}
	if sig.Retryable {
		t.Fatalf("retryable = true, want false for %q", sig.Class)
	}
	wantCanon := "test requires organization mode but GH_TEST_AUTH_MODE=individual; switch to organization mode or skip this test"
	if sig.Canonical != wantCanon {
		t.Fatalf("canonical = %q, want %q", sig.Canonical, wantCanon)
	}
}

func TestSignaturesForRunKeepsUnrelatedPreRunTestMainFailuresAsProviderTestMain(t *testing.T) {
	result := RunResult{PreRunFailed: true, PreRunOutput: []string{"provider TestMain failed: missing GITHUB_OWNER\n"}}

	sigs := SignaturesForRun("github", result, redact.New(nil))
	if len(sigs) != 1 {
		t.Fatalf("got %d signatures, want 1: %+v", len(sigs), sigs)
	}
	sig := sigs[0]
	if sig.Class != ClassProviderTestMain {
		t.Fatalf("class = %q, want %q; canonical=%q", sig.Class, ClassProviderTestMain, sig.Canonical)
	}
	if sig.Retryable {
		t.Fatalf("retryable = true, want false for %q", sig.Class)
	}
	if sig.Canonical != "provider TestMain failed: missing GITHUB_OWNER" {
		t.Fatalf("canonical = %q, want provider TestMain failure", sig.Canonical)
	}
}

func TestSignaturesForRunDoesNotClassifyUnrelated404AccessSignalsAsCapabilityUnavailable(t *testing.T) {
	cases := []struct {
		name      string
		output    string
		wantClass string
	}{
		{
			name:      "unrelated endpoint with codespaces access wording",
			output:    "GET /orgs/example/settings: 404 feature unavailable or token lacks Codespaces access\n",
			wantClass: ClassGeneric,
		},
		{
			name:      "codespaces endpoint without named capability signal",
			output:    "GET /user/codespaces/secrets: 404 feature unavailable or token lacks Codespaces access\n",
			wantClass: ClassGeneric,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			sig := SignaturesForRun("github", runWithFailedTest(tc.output), redact.New(nil))[0]
			if sig.Class == ClassCapabilityUnavailable {
				t.Fatalf("class = %q, want not %q; canonical=%q", sig.Class, ClassCapabilityUnavailable, sig.Canonical)
			}
			if sig.Class != tc.wantClass {
				t.Fatalf("class = %q, want %q; canonical=%q", sig.Class, tc.wantClass, sig.Canonical)
			}
		})
	}
}

func TestFingerprintNormalizesJitterToStableHash(t *testing.T) {
	resA := runWithFailedTest(strings.Join([]string{
		"POST https://api.github.com/repos: 422 Validation Failed [{Resource:Repository Field:name Code:already_exists Message:name tf-acc-test-abcd1234 already exists}]\n",
		"repo tf-acc-test-abcd1234 uuid 11111111-2222-3333-4444-555555555555 sha abcdef1234567890 node R_kgDOABCDEF123456 databaseId: 123456\n",
		"request X-GitHub-Request-Id: A1B2:C3D4:5E6F at 2026-07-08T15:04:05Z took 1.234s path /Users/alice/src/provider/github/resource.go:77 on main.tf line 8\n",
	}, ""))
	resB := runWithFailedTest(strings.Join([]string{
		"POST https://api.github.com/repos: 422 Validation Failed [{Resource:Repository Field:name Code:already_exists Message:name tf-acc-test-wxyz9876 already exists}]\n",
		"repo tf-acc-test-wxyz9876 uuid 99999999-aaaa-bbbb-cccc-dddddddddddd sha fedcba9876543210 node R_kgDOZYXWVU987654 databaseId: 999999\n",
		"request X-GitHub-Request-Id: F1E2:D3C4:B5A6 at 2027-01-02T03:04:05Z took 500ms path /home/bob/src/provider/github/resource.go:99 on main.tf line 42\n",
	}, ""))

	sigA := SignaturesForRun("github", resA, redact.New(nil))[0]
	sigB := SignaturesForRun("github", resB, redact.New(nil))[0]
	if sigA.Fingerprint != sigB.Fingerprint {
		t.Fatalf("fingerprints differ after normalization:\nA=%+v\nB=%+v", sigA, sigB)
	}
	for _, leaked := range []string{"abcd1234", "wxyz9876", "11111111-2222-3333-4444-555555555555", "fedcba9876543210", "123456", "999999", "/Users/alice", "/home/bob", "2026-07-08", "2027-01-02"} {
		if strings.Contains(sigA.NormalizedSample, leaked) || strings.Contains(sigB.NormalizedSample, leaked) || strings.Contains(sigA.Canonical, leaked) || strings.Contains(sigB.Canonical, leaked) {
			t.Fatalf("jitter %q leaked into signatures:\nA=%+v\nB=%+v", leaked, sigA, sigB)
		}
	}
}

func TestFingerprintStructuredErrorsAreOrderIndependent(t *testing.T) {
	first := "POST https://api.github.com/repos: 422 Validation Failed [{Resource:Repository Field:visibility Code:invalid Message:visibility is invalid} {Resource:Repository Field:description Code:missing Message:description is required}]\n"
	second := "POST https://api.github.com/repos: 422 Validation Failed [{Resource:Repository Field:description Code:missing Message:description is required} {Resource:Repository Field:visibility Code:invalid Message:visibility is invalid}]\n"

	sigA := SignaturesForRun("github", runWithFailedTest(first), redact.New(nil))[0]
	sigB := SignaturesForRun("github", runWithFailedTest(second), redact.New(nil))[0]

	if sigA.Fingerprint != sigB.Fingerprint {
		t.Fatalf("fingerprints differ for reordered structured errors:\nA=%+v\nB=%+v", sigA, sigB)
	}
	if sigA.Canonical != sigB.Canonical {
		t.Fatalf("canonicals differ for reordered structured errors:\nA=%q\nB=%q", sigA.Canonical, sigB.Canonical)
	}
}

func TestFingerprintNormalizesAbsolutePathLineJitter(t *testing.T) {
	resA := runWithFailedTest("assertion failed at /Users/alice/src/provider/github/resource.go:77\n")
	resB := runWithFailedTest("assertion failed at /Users/alice/src/provider/github/resource.go:99\n")

	sigA := SignaturesForRun("github", resA, redact.New(nil))[0]
	sigB := SignaturesForRun("github", resB, redact.New(nil))[0]

	if sigA.Fingerprint != sigB.Fingerprint {
		t.Fatalf("fingerprints differ for path line jitter:\nA=%+v\nB=%+v", sigA, sigB)
	}
	if strings.Contains(sigA.Canonical, ":77") || strings.Contains(sigB.Canonical, ":99") {
		t.Fatalf("canonical retained path line jitter:\nA=%q\nB=%q", sigA.Canonical, sigB.Canonical)
	}
}

func TestFingerprintRedactsSecretShapesAndEnvValues(t *testing.T) {
	envSecret := "ENV_SECRET_VALUE_123"
	values := []string{
		("gh" + "p_") + strings.Repeat("A", 36),
		("gh" + "o_") + strings.Repeat("B", 36),
		("gh" + "u_") + strings.Repeat("C", 36),
		("gh" + "s_") + strings.Repeat("D", 36),
		("gh" + "r_") + strings.Repeat("E", 36),
		"github" + "_pat_" + strings.Repeat("F", 22) + "_" + strings.Repeat("G", 59),
		"Bearer " + strings.Join([]string{strings.Repeat("H", 20), strings.Repeat("I", 24), strings.Repeat("J", 32)}, "."),
		fakePEMBeginLine() + strings.Repeat("K", 64) + "\n" + fakePEMEnd(),
		envSecret,
	}
	res := runWithFailedTest("assertion failed with " + strings.Join(values, " ") + "\n")
	sig := SignaturesForRun("github", res, redact.New([]string{envSecret}))[0]
	blob, err := json.Marshal(sig)
	if err != nil {
		t.Fatal(err)
	}
	for _, raw := range values {
		raw = strings.TrimPrefix(raw, "Bearer ")
		if strings.Contains(string(blob), raw) {
			t.Fatalf("signature leaked secret %q in %s", raw, blob)
		}
	}
	if !strings.Contains(string(blob), "***REDACTED***") {
		t.Fatalf("signature did not include redaction marker: %s", blob)
	}
}

func runWithFailedTest(output string) RunResult {
	return RunResult{Tests: []TestResult{{Package: "./github", Name: "TestAccThing", Status: provider.StatusFail, Output: []string{output}}}}
}

func TestFingerprintRedactsScannerStyleSplitPEM(t *testing.T) {
	bodyA := strings.Repeat("R", 64)
	bodyB := strings.Repeat("S", 64)
	res := RunResult{Tests: []TestResult{{
		Package: "./github",
		Name:    "TestAccThing",
		Status:  provider.StatusFail,
		Output: []string{
			"assertion failed",
			fakePEMBegin(),
			bodyA,
			bodyB,
			fakePEMEnd(),
		},
	}}}
	sig := SignaturesForRun("github", res, redact.New(nil))[0]
	blob, err := json.Marshal(sig)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(blob), bodyA) || strings.Contains(string(blob), bodyB) {
		t.Fatalf("signature leaked scanner-style split PEM: %s", blob)
	}
}

func fakePEMBegin() string     { return "-----BEGIN " + "PRIVATE KEY-----" }
func fakePEMEnd() string       { return "-----END " + "PRIVATE KEY-----" }
func fakePEMBeginLine() string { return fakePEMBegin() + "\n" }
func fakePEMEndLine() string   { return fakePEMEnd() + "\n" }

func runWithFailedTestFixture(t *testing.T, name string) RunResult {
	t.Helper()
	data, err := os.ReadFile(filepath.Join("testdata", "triage", name))
	if err != nil {
		t.Fatalf("runWithFailedTestFixture: %v", err)
	}
	return runWithFailedTest(string(data))
}

func parseRunFixture(t *testing.T, name string) RunResult {
	t.Helper()
	result, err := Parse(mustOpenFixture(t, name), nil)
	if err != nil {
		t.Fatalf("parseRunFixture: %v", err)
	}
	return result
}

// TestClassifyDeterministicModeExtractionEdgeCases pins the auth-mode value
// extraction in classifyDeterministic against three inputs that previously
// produced a wrong canonical string or crashed outright:
//
//   - an empty value ("GH_TEST_AUTH_MODE=" immediately followed by a
//     delimiter): the `sp > 0` guard treated index 0 as "no delimiter found"
//     and swallowed the whole remainder of the output as the mode value;
//   - a lowercase/mixed-case key, which must extract the same value;
//   - a non-ASCII prefix, where strings.ToLower changes the byte length
//     (U+023A "Ⱥ" is two bytes and lowercases to the three-byte U+2C65 "ⱥ"),
//     so an index computed on the lowercased text and applied to the original
//     text read past the wrong offset and panicked with a slice-bounds error.
func TestClassifyDeterministicModeExtractionEdgeCases(t *testing.T) {
	cases := []struct {
		name   string
		output string
		want   string
	}{
		{
			name:   "empty value followed by newline",
			output: "provider TestMain: selected test requires organization mode; GH_TEST_AUTH_MODE=\nmore output\n",
			want:   "test requires organization mode but GH_TEST_AUTH_MODE=; switch to organization mode or skip this test",
		},
		{
			name:   "lowercase key",
			output: "selected test requires organization mode; gh_test_auth_mode=individual\n",
			want:   "test requires organization mode but GH_TEST_AUTH_MODE=individual; switch to organization mode or skip this test",
		},
		{
			name:   "mixed case key",
			output: "selected test requires organization mode; Gh_Test_Auth_Mode=individual\n",
			want:   "test requires organization mode but GH_TEST_AUTH_MODE=individual; switch to organization mode or skip this test",
		},
		{
			name:   "non-ascii prefix that grows when lowercased",
			output: strings.Repeat("Ⱥ", 32) + " selected test requires organization mode; GH_TEST_AUTH_MODE=individual\n",
			want:   "test requires organization mode but GH_TEST_AUTH_MODE=individual; switch to organization mode or skip this test",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			sigs := SignaturesForRun("github", runWithFailedTest(tc.output), redact.New(nil))
			if len(sigs) != 1 {
				t.Fatalf("got %d signatures, want 1: %+v", len(sigs), sigs)
			}
			sig := sigs[0]
			if sig.Class != ClassModeIncompatible {
				t.Fatalf("class = %q, want %q; canonical=%q", sig.Class, ClassModeIncompatible, sig.Canonical)
			}
			if sig.Canonical != tc.want {
				t.Fatalf("canonical = %q, want %q", sig.Canonical, tc.want)
			}
			if sig.Retryable {
				t.Fatal("mode-incompatible failures must never be retryable")
			}
			// The canonical string feeds the fingerprint, so a stable
			// canonical must produce a stable fingerprint across runs.
			again := SignaturesForRun("github", runWithFailedTest(tc.output), redact.New(nil))
			if again[0].Fingerprint != sig.Fingerprint {
				t.Fatalf("fingerprint is not stable: %q vs %q", again[0].Fingerprint, sig.Fingerprint)
			}
		})
	}
}

// TestClassifyDeterministicModeExtractionPreservesRedaction verifies the
// extraction still operates on redacted text: a secret value appearing where
// the auth mode is read must never reach the canonical string or the
// normalized sample.
func TestClassifyDeterministicModeExtractionPreservesRedaction(t *testing.T) {
	const secret = "ghp_MODEEXTRACTIONSECRET"
	output := "selected test requires organization mode; GH_TEST_AUTH_MODE=" + secret + " token=" + secret + "\n"

	sigs := SignaturesForRun("github", runWithFailedTest(output), redact.New([]string{secret}))
	if len(sigs) != 1 {
		t.Fatalf("got %d signatures, want 1: %+v", len(sigs), sigs)
	}
	sig := sigs[0]
	if sig.Class != ClassModeIncompatible {
		t.Fatalf("class = %q, want %q", sig.Class, ClassModeIncompatible)
	}
	if strings.Contains(sig.Canonical, secret) {
		t.Fatalf("canonical leaks the secret: %q", sig.Canonical)
	}
	if strings.Contains(sig.NormalizedSample, secret) {
		t.Fatalf("normalized sample leaks the secret: %q", sig.NormalizedSample)
	}
}
