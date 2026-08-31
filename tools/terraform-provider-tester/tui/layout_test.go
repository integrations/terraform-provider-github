package tui

import (
	"fmt"
	"strings"
	"testing"

	"github.com/charmbracelet/lipgloss"
)

func TestLayoutForWidthUsesDeterministicBreakpoints(t *testing.T) {
	tests := []struct {
		width int
		want  layoutMode
	}{{71, layoutCompact}, {72, layoutStandard}, {109, layoutStandard}, {110, layoutWide}}
	for _, tc := range tests {
		if got := layoutForWidth(tc.width); got != tc.want {
			t.Fatalf("layoutForWidth(%d) = %d, want %d", tc.width, got, tc.want)
		}
	}
}

func TestPanelAndMetricRowsFitWidth(t *testing.T) {
	for _, width := range []int{52, 80, 120} {
		got := renderPanel(width, "Ready", "Environment status", "All checks passed.", panelSuccess, true)
		assertRenderedWidth(t, got, width)
		metrics := renderMetricRow(width, []metricSpec{
			{Label: "READY", Value: "6", Tone: panelSuccess},
			{Label: "BLOCKED", Value: "2", Tone: panelDanger},
		}, true)
		assertRenderedWidth(t, metrics, width)
	}
}

func TestMetricRowsFitWidthAndKeepMetricsWhenTheyFitIndividually(t *testing.T) {
	allMetrics := []metricSpec{
		{Label: "OK", Value: "1", Tone: panelSuccess},
		{Label: "BLOCK", Value: "2", Tone: panelDanger},
		{Label: "WARN", Value: "3", Tone: panelAttention},
		{Label: "SKIP", Value: "4", Tone: panelNeutral},
		{Label: "FLAKE", Value: "5", Tone: panelFlaky},
		{Label: "TOTAL", Value: "6", Tone: panelAccent},
	}
	for _, width := range []int{1, 2, 8, 12, 24, 52, 80, 120} {
		for count := 1; count <= len(allMetrics); count++ {
			t.Run(fmt.Sprintf("width_%d_metrics_%d", width, count), func(t *testing.T) {
				metrics := allMetrics[:count]
				var got string
				func() {
					defer func() {
						if r := recover(); r != nil {
							t.Fatalf("renderMetricRow(%d, %d metrics) panicked: %v", width, count, r)
						}
					}()
					got = stripANSI(renderMetricRow(width, metrics, true))
				}()
				assertRenderedWidth(t, got, width)
				if width < 12 {
					return
				}
				for _, metric := range metrics {
					for _, want := range []string{metric.Label, metric.Value} {
						if !strings.Contains(got, want) {
							t.Fatalf("metric row at width %d missing %q:\n%s", width, want, got)
						}
					}
				}
			})
		}
	}
}

func assertRenderedWidth(t *testing.T, rendered string, width int) {
	t.Helper()
	for i, line := range strings.Split(stripANSI(rendered), "\n") {
		if got := lipgloss.Width(line); got > width {
			t.Fatalf("line %d width = %d, want <= %d:\n%s", i, got, width, rendered)
		}
	}
}
