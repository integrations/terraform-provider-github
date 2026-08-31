package redact

import (
	"strings"
	"testing"
)

func TestRedactorScrubsValues(t *testing.T) {
	r := New([]string{"ghp_supersecret", "tok_two"})
	in := "using ghp_supersecret and tok_two here"
	out := r.String(in)
	if strings.Contains(out, "ghp_supersecret") || strings.Contains(out, "tok_two") {
		t.Fatalf("secret leaked: %q", out)
	}
	if !strings.Contains(out, "***") {
		t.Errorf("expected redaction marker, got %q", out)
	}
}

func TestRedactorIgnoresEmptyAndShort(t *testing.T) {
	// empty secret must never blanket-redact every character
	r := New([]string{"", "x"})
	if got := r.String("hello"); got != "hello" {
		t.Errorf("empty/1-char secrets must not alter output, got %q", got)
	}
}

func TestRedactorLines(t *testing.T) {
	r := New([]string{"SEKRET"})
	out := r.Lines([]string{"a SEKRET b", "clean"})
	if strings.Contains(strings.Join(out, "\n"), "SEKRET") {
		t.Fatal("secret leaked across lines")
	}
}

// TestRedactorHandlesOverlappingSecrets guards a partial-suffix leak: when one
// secret is a substring (here, a prefix) of another, masking the shorter one
// first would leave a fragment of the longer, real secret exposed. The redactor
// must always mask the longest match, regardless of input order.
func TestRedactorHandlesOverlappingSecrets(t *testing.T) {
	// "SEKRET" is a prefix of the longer secret "SEKRETPLUS".
	orders := [][]string{
		{"SEKRET", "SEKRETPLUS"}, // short-first: the order that used to leak
		{"SEKRETPLUS", "SEKRET"}, // long-first
	}
	for _, secrets := range orders {
		out := New(secrets).String("value=SEKRETPLUS end")
		if strings.Contains(out, "PLUS") {
			t.Errorf("partial secret suffix leaked for input order %v: %q", secrets, out)
		}
		if strings.Contains(out, "SEKRETPLUS") {
			t.Errorf("full secret leaked for input order %v: %q", secrets, out)
		}
	}
}
