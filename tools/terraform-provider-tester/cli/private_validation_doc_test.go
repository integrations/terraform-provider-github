package cli

import (
	"os"
	"strings"
	"testing"
)

// The private-validation runbook is only useful if a reader can find it and if
// it stays pinned to the code. This guards that the doc exists and is linked
// from the docs hub. The existing docs-consistency scanners
// (TestDocsReferenceOnlyRealEnvVars / TestDocsReferenceOnlyRealSubcommands)
// then cover its env-var and subcommand references automatically, because
// harnessDocs globs ../docs/*.md.
func TestPrivateValidationRunbookLinkedFromHub(t *testing.T) {
	if _, err := os.Stat("../docs/private-validation.md"); err != nil {
		t.Fatalf("private-validation runbook missing: %v", err)
	}
	hub := readDoc(t, "../docs/README.md")
	if !strings.Contains(hub, "private-validation.md") {
		t.Error("docs/README.md hub does not link private-validation.md")
	}
}
