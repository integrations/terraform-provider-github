package engine

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/github/terraform-provider-tester/internal/redact"
	"github.com/github/terraform-provider-tester/provider"
)

const FingerprintVersion = "pulsar-fp-v1"

const (
	ClassRateLimit          = "api/403-rate-limit"
	ClassSecondaryRateLimit = "api/403-secondary-rate-limit"
	ClassPermission         = "api/403-permission"
	ClassAuth               = "api/auth"
	ClassConflict           = "api/409-conflict"
	ClassValidation         = "api/422-validation"
	ClassLeftoverState      = "api/422-leftover-state"
	ClassTransientNetwork   = "api/transient-network"
	ClassBuild              = "go/build"
	ClassPreRun             = "pulsar/pre-run"
	ClassProviderTestMain   = "provider/testmain"
	ClassPanic              = "panic"
	ClassTimeout            = "timeout"
	ClassTerraformDiag      = "terraform/diagnostic"
	ClassGeneric            = "generic"

	ClassModeIncompatible               = "triage/mode-incompatible"
	ClassCapabilityUnavailable          = "triage/capability-unavailable"
	ClassFixtureMissing                 = "triage/fixture-missing"
	ClassAPINotFoundPermissionOrFeature = "triage/api-404-permission-or-feature"
)

// Signature is the stable, redacted failure identity stored in state and JSON.
type Signature struct {
	Version          string `json:"version"`
	Provider         string `json:"provider"`
	Package          string `json:"package"`
	Test             string `json:"test"`
	Sub              string `json:"sub,omitempty"`
	Status           string `json:"status"`
	Class            string `json:"class"`
	Canonical        string `json:"canonical"`
	NormalizedSample string `json:"normalized_sample,omitempty"`
	Fingerprint      string `json:"fingerprint"`
	ShortFingerprint string `json:"short_fingerprint"`
	Retryable        bool   `json:"retryable"`
}

// SignaturesForRun returns one safe fingerprint for each suite or top-level failure.
func SignaturesForRun(providerName string, res RunResult, red *redact.Redactor) []Signature {
	if providerName == "" {
		providerName = "github"
	}
	var sigs []Signature
	if res.BuildFailed {
		sigs = append(sigs, signatureForLines(providerName, "", "suite", "", "build", res.BuildOutput, red))
	}
	if res.PreRunFailed {
		sigs = append(sigs, signatureForLines(providerName, "", "suite", "", "pre-run", res.PreRunOutput, red))
	}
	if res.BuildFailed || res.PreRunFailed {
		return sigs
	}
	for _, tr := range res.Tests {
		if tr.Sub != "" || !isFailStatus(tr.Status) {
			continue
		}
		sigs = append(sigs, SignatureForTest(providerName, res, tr, red))
	}
	return sigs
}

// SignatureForTest fingerprints one failed top-level test and its failed subtests.
func SignatureForTest(providerName string, res RunResult, tr TestResult, red *redact.Redactor) Signature {
	if providerName == "" {
		providerName = "github"
	}
	lines := outputForTestFailure(res, tr)
	return signatureForLines(providerName, tr.Package, tr.Name, tr.Sub, tr.Status.String(), lines, red)
}

func outputForTestFailure(res RunResult, tr TestResult) []string {
	lines := append([]string{}, tr.Output...)
	for _, sub := range res.Tests {
		if sub.Package == tr.Package && sub.Name == tr.Name && sub.Sub != "" && isFailStatus(sub.Status) {
			lines = append(lines, sub.Output...)
		}
	}
	return lines
}

func signatureForLines(providerName, pkg, test, sub, status string, raw []string, red *redact.Redactor) Signature {
	lines := normalizeFailureLines(raw, red)
	class, canonical := classifyNormalized(status, lines)
	canonical = truncateBytes(canonical, 240)
	sample := truncateBytes(strings.Join(lines, "\n"), 4*1024)
	sig := Signature{
		Version:          FingerprintVersion,
		Provider:         providerName,
		Package:          pkg,
		Test:             test,
		Sub:              sub,
		Status:           status,
		Class:            class,
		Canonical:        canonical,
		NormalizedSample: sample,
		Retryable:        IsRetryableClass(class),
	}
	sig.Fingerprint, sig.ShortFingerprint = fingerprintForSignature(sig)
	return sig
}

func fingerprintForSignature(sig Signature) (string, string) {
	stable := strings.Join([]string{
		FingerprintVersion,
		"provider=" + sig.Provider,
		"package=" + sig.Package,
		"test=" + sig.Test,
		"sub=" + sig.Sub,
		"status=" + sig.Status,
		"class=" + sig.Class,
		"canonical=" + sig.Canonical,
	}, "\n")
	sum := sha256.Sum256([]byte(stable))
	hexsum := hex.EncodeToString(sum[:])
	return "sha256:" + hexsum, hexsum[:16]
}

func IsRetryableClass(class string) bool {
	switch class {
	case ClassRateLimit, ClassSecondaryRateLimit, ClassConflict, ClassLeftoverState, ClassTransientNetwork:
		return true
	default:
		return false
	}
}

func IsInfraRetryableClass(class string) bool {
	return IsRetryableClass(class)
}

var (
	ansiPattern               = regexp.MustCompile(`\x1b\[[0-9;?]*[ -/]*[@-~]`)
	tfAccDashPattern          = regexp.MustCompile(`tf-acc-test-[A-Za-z0-9_-]+`)
	tfAccUnderscorePattern    = regexp.MustCompile(`tf_acc_test_[A-Za-z0-9_-]+`)
	tfResourceNamePattern     = regexp.MustCompile(`\b([A-Za-z0-9_]+)\.tf-acc-test-<id>\b`)
	uuidPattern               = regexp.MustCompile(`\b[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}\b`)
	nodeIDPattern             = regexp.MustCompile(`\b(?:R_|U_|T_|I_|MD)[A-Za-z0-9_-]{12,}\b`)
	requestIDPattern          = regexp.MustCompile(`(?i)X-GitHub-Request-Id:\s*[^\s,;]+`)
	rfc3339Pattern            = regexp.MustCompile(`\b\d{4}-\d{2}-\d{2}[T ][0-9:.+-]+Z?\b`)
	epochPattern              = regexp.MustCompile(`\b\d{10,}\b`)
	durationPattern           = regexp.MustCompile(`\b(?:\d+(?:\.\d+)?(?:ns|us|µs|ms|s|m|h))+\b`)
	absolutePathPattern       = regexp.MustCompile(`(?:/Users|/home|/var|/private|/workspace|/github|/opt)/[^\s:]+`)
	normalizedPathLinePattern = regexp.MustCompile(`<path>:\d+\b`)
	goFileLinePattern         = regexp.MustCompile(`\b([A-Za-z0-9_./-]+\.go):(\d+)\b`)
	terraformLinePattern      = regexp.MustCompile(`\bon ([^\s]+\.tf) line \d+\b`)
	idEqualsPattern           = regexp.MustCompile(`\b(id=)\d+\b`)
	idWordPattern             = regexp.MustCompile(`\b(ID\s+)\d+\b`)
	databaseIDPattern         = regexp.MustCompile(`\b(databaseId:\s*)\d+\b`)
	slashIDPattern            = regexp.MustCompile(`/\d{4,}\b`)
	hexPattern                = regexp.MustCompile(`\b[0-9a-fA-F]{7,}\b`)
	failDurationPattern       = regexp.MustCompile(`--- FAIL: ([^\s]+) \([^)]*\)`)
	jsonErrorResourcePattern  = regexp.MustCompile(`Resource:([A-Za-z0-9_/-]+)\s+Field:([A-Za-z0-9_/-]+)\s+Code:([A-Za-z0-9_/-]+)\s+Message:([^}\]]+)`)
	alreadyExistsCodePattern  = regexp.MustCompile(`(?i)"?code"?\s*[:=]\s*"?already_exists"?`)
)

func normalizeFailureLines(raw []string, red *redact.Redactor) []string {
	joined := joinOutputLines(raw)
	joined = strings.ReplaceAll(joined, "\r\n", "\n")
	joined = strings.ReplaceAll(joined, "\r", "\n")
	if red != nil {
		joined = red.String(joined)
	} else {
		joined = redact.New(nil).String(joined)
	}
	joined = ansiPattern.ReplaceAllString(joined, "")
	parts := strings.Split(joined, "\n")
	out := make([]string, 0, len(parts))
	blank := false
	for _, line := range parts {
		line = strings.TrimRight(line, " \t")
		line = normalizeFailureLine(line)
		if dropGoTestNoise(line) {
			continue
		}
		if line == "" {
			if blank {
				continue
			}
			blank = true
			out = append(out, line)
			continue
		}
		blank = false
		out = append(out, line)
	}
	for len(out) > 0 && out[0] == "" {
		out = out[1:]
	}
	for len(out) > 0 && out[len(out)-1] == "" {
		out = out[:len(out)-1]
	}
	return out
}

func joinOutputLines(raw []string) string {
	var b strings.Builder
	for i, line := range raw {
		b.WriteString(line)
		if i+1 < len(raw) && !strings.HasSuffix(line, "\n") && !strings.HasSuffix(line, "\r") {
			b.WriteByte('\n')
		}
	}
	return b.String()
}

func normalizeFailureLine(line string) string {
	line = failDurationPattern.ReplaceAllString(line, "--- FAIL: $1 (<duration>)")
	line = tfAccDashPattern.ReplaceAllString(line, "tf-acc-test-<id>")
	line = tfAccUnderscorePattern.ReplaceAllString(line, "tf_acc_test_<id>")
	line = tfResourceNamePattern.ReplaceAllString(line, "$1.<tf-acc-test>")
	line = requestIDPattern.ReplaceAllString(line, "X-GitHub-Request-Id: <request-id>")
	line = uuidPattern.ReplaceAllString(line, "<uuid>")
	line = nodeIDPattern.ReplaceAllString(line, "<node-id>")
	line = rfc3339Pattern.ReplaceAllString(line, "<timestamp>")
	line = epochPattern.ReplaceAllString(line, "<timestamp>")
	line = durationPattern.ReplaceAllString(line, "<duration>")
	line = absolutePathPattern.ReplaceAllString(line, "<path>")
	line = normalizedPathLinePattern.ReplaceAllString(line, "<path>:<line>")
	line = goFileLinePattern.ReplaceAllString(line, "$1:<line>")
	line = terraformLinePattern.ReplaceAllString(line, "on $1 line <line>")
	line = idEqualsPattern.ReplaceAllString(line, "${1}<id>")
	line = idWordPattern.ReplaceAllString(line, "${1}<id>")
	line = databaseIDPattern.ReplaceAllString(line, "${1}<id>")
	line = slashIDPattern.ReplaceAllString(line, "/<id>")
	line = hexPattern.ReplaceAllStringFunc(line, func(v string) string {
		if strings.HasPrefix(strings.ToLower(v), "sha256") {
			return v
		}
		return "<hex>"
	})
	return strings.TrimSpace(line)
}

func dropGoTestNoise(line string) bool {
	trim := strings.TrimSpace(line)
	if trim == "" {
		return false
	}
	if strings.HasPrefix(trim, "=== RUN") || strings.HasPrefix(trim, "=== PAUSE") || strings.HasPrefix(trim, "=== CONT") || strings.HasPrefix(trim, "--- PASS:") {
		return true
	}
	if trim == "PASS" || trim == "FAIL" || strings.HasPrefix(trim, "ok \t") || strings.HasPrefix(trim, "ok  ") {
		return true
	}
	return durationPattern.MatchString(trim) && durationPattern.ReplaceAllString(trim, "") == ""
}

func classifyNormalized(status string, lines []string) (string, string) {
	if status == "build" {
		return ClassBuild, firstBuildLine(lines)
	}
	if status == "pre-run" {
		if class, canon, ok := classifyDeterministic(lines); ok {
			return class, canon
		}
		canon := firstSignalLine(lines)
		if strings.Contains(strings.ToLower(canon), "testmain") {
			return ClassProviderTestMain, canon
		}
		return ClassPreRun, canon
	}
	if class, canon, ok := classifyDeterministic(lines); ok {
		return class, canon
	}
	if class, canon, ok := classifyAPI(lines); ok {
		return class, canon
	}
	if canon, ok := terraformDiagnostic(lines); ok {
		return ClassTerraformDiag, canon
	}
	if status == provider.StatusTimeout.String() {
		return ClassTimeout, timeoutCanonical(lines)
	}
	if status == provider.StatusPanic.String() {
		return ClassPanic, panicCanonical(lines)
	}
	return ClassGeneric, firstSignalLine(lines)
}

// classifyDeterministic recognises non-retryable deterministic failure patterns
// in priority order before the generic API / diagnostic fallbacks.
func classifyDeterministic(lines []string) (string, string, bool) {
	text := strings.Join(lines, "\n")
	lower := strings.ToLower(text)

	// 1. Mode / auth-mode mismatch.
	if (strings.Contains(lower, "requires organization mode") || strings.Contains(lower, "selected test requires organization mode")) &&
		strings.Contains(lower, "gh_test_auth_mode") {
		// Extract the auth mode value if present.
		authMode := "individual"
		if value, found := authModeValue(text); found {
			authMode = value
		}
		return ClassModeIncompatible,
			"test requires organization mode but GH_TEST_AUTH_MODE=" + authMode +
				"; switch to organization mode or skip this test",
			true
	}

	// 2. Codespaces / named-capability endpoint unavailable (404 with feature/access language).
	if strings.Contains(lower, "404") &&
		(strings.Contains(lower, "feature unavailable") || strings.Contains(lower, "lacks codespaces access") || strings.Contains(lower, "lacks required access")) {
		endpoint := extractEndpointPath(text)
		if isNamedCapabilityEndpoint(endpoint) {
			canon := "endpoint " + endpoint + " returned 404; feature unavailable or token lacks required access; enable the capability or verify the token has required access"
			return ClassCapabilityUnavailable, canon, true
		}
	}

	// 3. Required fixture missing or inaccessible.
	if strings.Contains(lower, "required fixture") &&
		(strings.Contains(lower, "missing") || strings.Contains(lower, "inaccessible")) {
		fixtureVar := extractFixtureVar(text)
		canon := "required fixture " + fixtureVar + " is missing or inaccessible; set the environment variable to a valid value"
		return ClassFixtureMissing, canon, true
	}

	// 4. 404 plus permission/feature ambiguity (generic endpoint).
	if strings.Contains(lower, "404") &&
		(strings.Contains(lower, "require permission") || strings.Contains(lower, "requires permission") ||
			strings.Contains(lower, "enabled feature") || strings.Contains(lower, "permission or an enabled feature")) {
		endpoint := extractEndpointPath(text)
		method := extractHTTPMethod(text)
		prefix := method
		if prefix != "" {
			prefix += " "
		}
		canon := prefix + endpoint + " returned 404; endpoint requires permission or an enabled feature; verify permission, feature enablement, and endpoint availability"
		return ClassAPINotFoundPermissionOrFeature, canon, true
	}

	return "", "", false
}

var (
	endpointPathPattern = regexp.MustCompile(`(?:GET|POST|PUT|PATCH|DELETE|HEAD)\s+(/[^\s:]+)`)
	httpMethodPattern   = regexp.MustCompile(`\b(GET|POST|PUT|PATCH|DELETE|HEAD)\b`)
	fixtureVarPattern   = regexp.MustCompile(`\b(GH_TEST_[A-Z0-9_]+)\b`)
)

func extractEndpointPath(text string) string {
	if m := endpointPathPattern.FindStringSubmatch(text); len(m) == 2 {
		return m[1]
	}
	return "unknown endpoint"
}

func extractHTTPMethod(text string) string {
	if m := httpMethodPattern.FindString(text); m != "" {
		return m
	}
	return ""
}

func extractFixtureVar(text string) string {
	if m := fixtureVarPattern.FindString(text); m != "" {
		return m
	}
	return "unknown fixture"
}

func isNamedCapabilityEndpoint(endpoint string) bool {
	return strings.Contains(endpoint, "/codespaces/") && strings.HasSuffix(endpoint, "/public-key")
}

var authModeAssignmentPattern = regexp.MustCompile(`(?i)GH_TEST_AUTH_MODE=`)

// authModeValue extracts the value assigned to GH_TEST_AUTH_MODE in text,
// case-insensitively, and reports whether an assignment was present.
//
// Both indexing and slicing happen against the SAME string: a previous
// version searched a strings.ToLower copy and then sliced the original with
// that index, which is only safe for ASCII. strings.ToLower can change a
// string's byte length (U+023A "Ⱥ" is two bytes and lowercases to the
// three-byte U+2C65 "ⱥ"), so a non-ASCII prefix shifted every later index and
// the slice either read the wrong offset or panicked with a slice-bounds
// error.
//
// The value ends at the first delimiter. An assignment immediately followed
// by a delimiter (or by end of text) yields the empty string: index 0 is a
// found delimiter, not "no delimiter", so the rest of the failure output is
// never swallowed into the canonical string.
func authModeValue(text string) (string, bool) {
	loc := authModeAssignmentPattern.FindStringIndex(text)
	if loc == nil {
		return "", false
	}
	rest := text[loc[1]:]
	if sp := strings.IndexAny(rest, " \t\n;,"); sp >= 0 {
		return rest[:sp], true
	}
	return rest, true
}

// classifyAPI recognises retryable and permission-class API failures.
func classifyAPI(lines []string) (string, string, bool) {
	text := strings.Join(lines, "\n")
	lower := strings.ToLower(text)
	switch {
	case strings.Contains(lower, "403") && strings.Contains(lower, "secondary") && strings.Contains(lower, "rate limit"):
		return ClassSecondaryRateLimit, "403 secondary rate limit", true
	case strings.Contains(lower, "403") && strings.Contains(lower, "rate limit"):
		return ClassRateLimit, "403 rate limit", true
	case strings.Contains(lower, "403") && (strings.Contains(lower, "resource not accessible by token") || strings.Contains(lower, "permission") || strings.Contains(lower, "missing scope")):
		return ClassPermission, "403 resource not accessible by token", true
	case strings.Contains(lower, "401") || strings.Contains(lower, "bad credentials") || strings.Contains(lower, "requires authentication"):
		return ClassAuth, "auth failure", true
	case strings.Contains(lower, "409") && strings.Contains(lower, "conflict"):
		return ClassConflict, "409 conflict " + conflictReason(lines), true
	case strings.Contains(lower, "422") && (strings.Contains(lower, "already_exists") || strings.Contains(lower, "already exists")):
		if hasLeftoverStateSignal(text) {
			return ClassLeftoverState, "422 leftover state " + leftoverReason(lines), true
		}
		return ClassValidation, "422 validation " + validationReason(lines), true
	case strings.Contains(lower, "422") && strings.Contains(lower, "validation"):
		return ClassValidation, "422 validation " + validationReason(lines), true
	case strings.Contains(lower, "502") || strings.Contains(lower, "503") || strings.Contains(lower, "504") || strings.Contains(lower, "eof") || strings.Contains(lower, "connection reset"):
		return ClassTransientNetwork, "transient network " + transientReason(lower), true
	default:
		return "", "", false
	}
}

func hasLeftoverStateSignal(text string) bool {
	return alreadyExistsCodePattern.MatchString(text) || strings.Contains(strings.ToLower(text), "tf-acc-test-")
}

func conflictReason(lines []string) string {
	line := lowerFirstContaining(lines, "conflict")
	if idx := strings.LastIndex(line, "conflict:"); idx >= 0 {
		line = line[idx+len("conflict:"):]
	} else if idx := strings.LastIndex(line, "conflict"); idx >= 0 {
		line = line[idx+len("conflict"):]
	}
	line = strings.ReplaceAll(line, "409", "")
	line = strings.Trim(line, " :.-")
	if line == "" {
		return "conflict"
	}
	return compactReason(line)
}

func leftoverReason(lines []string) string {
	if errs := structuredAPIErrors(lines); len(errs) > 0 {
		if len(errs) == 1 && errs[0].code == "already_exists" && (errs[0].field == "name" || strings.Contains(errs[0].message, "already exists")) {
			return "name already exists"
		}
		return structuredErrorReason(errs)
	}
	text := strings.ToLower(strings.Join(lines, " "))
	if strings.Contains(text, "name") && strings.Contains(text, "already exists") {
		return "name already exists"
	}
	if strings.Contains(text, "already_exists") {
		return "already exists"
	}
	return "already exists"
}

func validationReason(lines []string) string {
	text := strings.Join(lines, " ")
	if errs := structuredAPIErrors(lines); len(errs) > 0 {
		if len(errs) == 1 {
			return errs[0].message
		}
		return structuredErrorReason(errs)
	}
	lowerText := strings.ToLower(text)
	if strings.Contains(lowerText, "name") && strings.Contains(lowerText, "already exists") {
		return "name already exists"
	}
	line := lowerFirstContaining(lines, "validation")
	line = strings.ReplaceAll(line, "422", "")
	line = strings.ReplaceAll(line, "validation failed", "")
	line = strings.ReplaceAll(line, "validation", "")
	line = strings.Trim(line, " :.-[]{}")
	if line == "" {
		return "failed"
	}
	return compactReason(line)
}

type structuredAPIError struct {
	resource string
	field    string
	code     string
	message  string
}

func structuredAPIErrors(lines []string) []structuredAPIError {
	matches := jsonErrorResourcePattern.FindAllStringSubmatch(strings.Join(lines, " "), -1)
	errs := make([]structuredAPIError, 0, len(matches))
	for _, m := range matches {
		if len(m) != 5 {
			continue
		}
		errs = append(errs, structuredAPIError{
			resource: compactReason(strings.ToLower(strings.TrimSpace(m[1]))),
			field:    compactReason(strings.ToLower(strings.TrimSpace(m[2]))),
			code:     compactReason(strings.ToLower(strings.TrimSpace(m[3]))),
			message:  compactReason(strings.ToLower(strings.TrimSpace(m[4]))),
		})
	}
	sort.Slice(errs, func(i, j int) bool {
		left := []string{errs[i].resource, errs[i].field, errs[i].code, errs[i].message}
		right := []string{errs[j].resource, errs[j].field, errs[j].code, errs[j].message}
		for idx := range left {
			if left[idx] != right[idx] {
				return left[idx] < right[idx]
			}
		}
		return false
	})
	return errs
}

func structuredErrorReason(errs []structuredAPIError) string {
	parts := make([]string, 0, len(errs))
	for _, err := range errs {
		key := strings.Trim(strings.Join([]string{err.resource, err.field, err.code}, "."), ".")
		if err.message == "" {
			parts = append(parts, key)
			continue
		}
		parts = append(parts, strings.TrimSpace(key+" "+err.message))
	}
	return strings.Join(parts, "; ")
}

func transientReason(lower string) string {
	for _, code := range []string{"502", "503", "504"} {
		if strings.Contains(lower, code) {
			return code
		}
	}
	if strings.Contains(lower, "connection reset") {
		return "connection reset"
	}
	return "eof"
}

func lowerFirstContaining(lines []string, needle string) string {
	for _, line := range lines {
		if strings.Contains(strings.ToLower(line), needle) {
			return strings.ToLower(line)
		}
	}
	return ""
}

func compactReason(s string) string {
	s = strings.ReplaceAll(s, "tf-acc-test-<id>", "")
	s = strings.ReplaceAll(s, "[]", "")
	s = strings.Join(strings.Fields(s), " ")
	return strings.Trim(s, " :.-")
}

func terraformDiagnostic(lines []string) (string, bool) {
	for i, line := range lines {
		if strings.HasPrefix(line, "Error:") {
			if strings.TrimSpace(line) != "Error:" {
				return line, true
			}
			for _, detail := range lines[i+1:] {
				if strings.TrimSpace(detail) != "" {
					return "Error: " + strings.TrimSpace(detail), true
				}
			}
			return "Error:", true
		}
	}
	return "", false
}

func timeoutCanonical(lines []string) string {
	for _, line := range lines {
		lower := strings.ToLower(line)
		if strings.Contains(lower, "test timed out") || strings.Contains(lower, "test killed") {
			return "test timed out"
		}
	}
	return firstSignalLine(lines)
}

func panicCanonical(lines []string) string {
	panicLine := "panic"
	for _, line := range lines {
		if strings.HasPrefix(strings.ToLower(line), "panic:") {
			panicLine = line
			break
		}
	}
	for _, line := range lines {
		trim := strings.TrimSpace(line)
		if strings.HasPrefix(trim, "github.com/") && !strings.Contains(trim, "/runtime") {
			if idx := strings.Index(trim, "("); idx >= 0 {
				trim = trim[:idx]
			}
			return panicLine + " " + trim
		}
	}
	return panicLine
}

func firstBuildLine(lines []string) string {
	for _, line := range lines {
		trim := strings.TrimSpace(line)
		if trim == "" || strings.HasPrefix(trim, "# ") {
			continue
		}
		return trim
	}
	return firstSignalLine(lines)
}

func firstSignalLine(lines []string) string {
	for _, line := range lines {
		trim := strings.TrimSpace(line)
		if trim == "" {
			continue
		}
		return trim
	}
	return "unknown failure"
}

func truncateBytes(s string, n int) string {
	if n <= 0 || len(s) <= n {
		return s
	}
	return string([]byte(s)[:n])
}

func SortSignatures(sigs []Signature) {
	sort.SliceStable(sigs, func(i, j int) bool {
		return fmt.Sprintf("%s/%s/%s", sigs[i].Package, sigs[i].Test, sigs[i].Class) < fmt.Sprintf("%s/%s/%s", sigs[j].Package, sigs[j].Test, sigs[j].Class)
	})
}

func retryAfterFromLines(lines []string, now time.Time) time.Duration {
	var best time.Duration
	for _, line := range lines {
		lower := strings.ToLower(line)
		if idx := strings.Index(lower, "retry-after:"); idx >= 0 {
			val := strings.TrimSpace(line[idx+len("retry-after:"):])
			fields := strings.Fields(val)
			if len(fields) > 0 {
				if dur, ok := parseRetrySeconds(fields[0]); ok && dur > best {
					best = dur
				}
			}
		}
		if idx := strings.Index(lower, "x-ratelimit-reset:"); idx >= 0 {
			val := strings.TrimSpace(line[idx+len("x-ratelimit-reset:"):])
			fields := strings.Fields(val)
			if len(fields) > 0 {
				if dur, ok := parseResetDuration(fields[0], now); ok && dur > best {
					best = dur
				}
			}
		}
	}
	return best
}

func parseRetrySeconds(raw string) (time.Duration, bool) {
	raw = strings.Trim(raw, " ,;.")
	var sec int64
	if _, err := fmt.Sscanf(raw, "%d", &sec); err != nil || sec < 0 {
		return 0, false
	}
	return time.Duration(sec) * time.Second, true
}

func parseResetDuration(raw string, now time.Time) (time.Duration, bool) {
	raw = strings.Trim(raw, " ,;.")
	var epoch int64
	if _, err := fmt.Sscanf(raw, "%d", &epoch); err != nil || epoch <= 0 {
		return 0, false
	}
	if now.IsZero() {
		now = time.Now()
	}
	d := time.Unix(epoch, 0).Sub(now)
	if d < 0 {
		return 0, true
	}
	return d, true
}
