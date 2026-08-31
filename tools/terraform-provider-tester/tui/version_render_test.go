package tui

import (
	"strings"
	"testing"
)

func TestVersionRenderingDoesNotDoubleV(t *testing.T) {
	const version = "v9.9.9"

	intro := renderIntro(introPeakFrame, 100, 30, version, true)
	if !strings.Contains(intro, "GH/TF Provider Acceptance v9.9.9") {
		t.Fatalf("intro should contain the versioned console identity\n%s", intro)
	}
	if strings.Contains(intro, "vv9.9.9") {
		t.Fatalf("intro should not contain doubled v version\n%s", intro)
	}

	banner := renderBanner(version, true)
	if !strings.Contains(banner, "GH/TF Provider Acceptance v9.9.9") {
		t.Fatalf("banner should contain the versioned console identity\n%s", banner)
	}
	if strings.Contains(banner, "vv9.9.9") {
		t.Fatalf("banner should not contain doubled v version\n%s", banner)
	}
}
