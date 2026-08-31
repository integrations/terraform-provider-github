package tui

import (
	"strings"
	"testing"

	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"

	"github.com/github/terraform-provider-tester/engine"
	"github.com/github/terraform-provider-tester/provider"
)

func TestInteractiveSelectionCursorsUseTerraformPurple(t *testing.T) {
	lipgloss.SetColorProfile(termenv.TrueColor)
	lipgloss.SetHasDarkBackground(true)
	defer func() {
		lipgloss.SetColorProfile(termenv.Ascii)
		lipgloss.SetHasDarkBackground(false)
	}()

	wantCursor := lipgloss.NewStyle().
		Foreground(TerraformPurple).
		Bold(true).
		Render("❯")

	groupModel := fixedModel()
	groupModel.ascii = false
	groupModel.groups = []engine.Group{{Name: "repositories", Tests: []string{"TestAccRepository"}}}

	testModel := groupModel
	testModel.focus = focusTests

	pickerModel := fixedModel().WithModes([]provider.Mode{{Name: "anonymous", Description: "Public endpoints"}})
	pickerModel.ascii = false

	envVars := []provider.EnvVar{{Key: "GITHUB_OWNER", Required: true}}
	editorModel := fixedModel().WithEnvVars(envVars).WithGetenv(func(string) string { return "" })
	editorModel.ascii = false
	editorModel.editorFields = buildEditorFields(envVars, editorModel.getenv)

	triageModel := fixedModel()
	triageModel.ascii = false
	triageModel.triageFailures = sampleFailures()[:1]

	tests := []struct {
		name   string
		render func() string
	}{
		{name: "groups", render: func() string { return groupModel.renderGroups(100) }},
		{name: "tests", render: func() string { return testModel.renderTestsDrilldown(100) }},
		{name: "mode picker", render: func() string { return pickerModel.renderPickerOverlay(100) }},
		{name: "variable editor", render: func() string { return editorModel.renderEditorOverlay(100) }},
		{name: "triage", render: func() string { return triageModel.renderTriageList(100) }},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.render(); !strings.Contains(got, wantCursor) {
				t.Fatalf("selection cursor is not Terraform purple:\n%s", got)
			}
		})
	}
}
