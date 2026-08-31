package redact

import (
	"strings"
	"testing"
)

func TestRedactorScrubsCredentialShapedPatterns(t *testing.T) {
	prefixes := []string{
		"gh" + "p_",
		"gh" + "o_",
		"gh" + "u_",
		"gh" + "s_",
		"gh" + "r_",
	}
	for _, prefix := range prefixes {
		raw := prefix + strings.Repeat("A", 36)
		out := New(nil).String("token=" + raw)
		if strings.Contains(out, raw) {
			t.Fatalf("%s token leaked: %q", prefix, out)
		}
		if !strings.Contains(out, marker) {
			t.Fatalf("%s token was not replaced with marker: %q", prefix, out)
		}
	}

	pat := "github" + "_pat_" + strings.Repeat("A", 22) + "_" + strings.Repeat("B", 59)
	out := New(nil).String("token=" + pat)
	if strings.Contains(out, pat) {
		t.Fatalf("github_pat token leaked: %q", out)
	}
	if !strings.Contains(out, marker) {
		t.Fatalf("github_pat token was not replaced with marker: %q", out)
	}
}

func TestRedactorScrubsBearerJWTAndAuthorizationHeaders(t *testing.T) {
	jwt := "Bearer " + strings.Join([]string{
		strings.Repeat("A", 20),
		strings.Repeat("B", 24),
		strings.Repeat("C", 32),
	}, ".")
	for _, raw := range []string{
		jwt,
		"Authorization: token " + ("gh" + "p_") + strings.Repeat("D", 36),
		"Authorization: bearer " + strings.Join([]string{strings.Repeat("E", 20), strings.Repeat("F", 24), strings.Repeat("G", 32)}, "."),
	} {
		out := New(nil).String(raw)
		if strings.Contains(out, raw) || strings.Contains(out, strings.TrimPrefix(raw, "Bearer ")) {
			t.Fatalf("credential leaked: %q", out)
		}
		if !strings.Contains(out, marker) {
			t.Fatalf("credential was not replaced with marker: %q", out)
		}
	}
}

func TestRedactorScrubsPEMBody(t *testing.T) {
	body := strings.Repeat("A", 64) + "\n" + strings.Repeat("B", 64)
	raw := pemBeginLine() + body + "\n" + pemEnd()
	out := New(nil).String(raw)
	if strings.Contains(out, body) || strings.Contains(out, strings.Repeat("A", 64)) || strings.Contains(out, strings.Repeat("B", 64)) {
		t.Fatalf("PEM body leaked: %q", out)
	}
	if !strings.Contains(out, marker) {
		t.Fatalf("PEM body was not replaced with marker: %q", out)
	}
}

func TestRedactorLinesScrubsSplitPEMAndKnownMultilineSecret(t *testing.T) {
	bodyA := strings.Repeat("L", 64)
	bodyB := strings.Repeat("M", 64)
	lines := []string{
		"before\n",
		pemBeginLine(),
		bodyA + "\n",
		bodyB + "\n",
		pemEndLine(),
		"after\n",
	}
	joinedSecret := strings.Join(lines[1:5], "")
	out := strings.Join(New([]string{joinedSecret}).Lines(lines), "")
	for _, leaked := range []string{bodyA, bodyB, joinedSecret} {
		if strings.Contains(out, leaked) {
			t.Fatalf("split multiline secret leaked %q in %q", leaked, out)
		}
	}
	if !strings.Contains(out, marker) {
		t.Fatalf("split multiline secret was not replaced with marker: %q", out)
	}
}

func TestRedactorLinesScrubsScannerStyleSplitPEM(t *testing.T) {
	bodyA := strings.Repeat("P", 64)
	bodyB := strings.Repeat("Q", 64)
	lines := []string{
		pemBegin(),
		bodyA,
		bodyB,
		pemEnd(),
	}
	out := strings.Join(New(nil).Lines(lines), "")
	if strings.Contains(out, bodyA) || strings.Contains(out, bodyB) {
		t.Fatalf("scanner-style split PEM leaked: %q", out)
	}
	if !strings.Contains(out, marker) {
		t.Fatalf("scanner-style split PEM missing marker: %q", out)
	}
}

func pemBegin() string     { return "-----BEGIN " + "PRIVATE KEY-----" }
func pemEnd() string       { return "-----END " + "PRIVATE KEY-----" }
func pemBeginLine() string { return pemBegin() + "\n" }
func pemEndLine() string   { return pemEnd() + "\n" }
