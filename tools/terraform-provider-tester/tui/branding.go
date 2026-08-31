package tui

import "github.com/charmbracelet/lipgloss"

// AppName is the automation-facing product name shown by the CLI and
// documentation. The invocation command and binary are
// "terraform-provider-tester". The project was previously named Pulsar.
const AppName = "Terraform Provider Tester"

// ConsoleMark and ConsoleName form the dashboard-only identity. They connect
// the console directly to the GitHub Terraform provider without renaming the
// repository, binary, command, or compatibility surface.
const (
	ConsoleMark  = "GH/TF"
	ConsoleName  = "Provider Acceptance"
	ConsoleTitle = ConsoleMark + " " + ConsoleName
)

// TargetRepo identifies the provider this harness drives. It is shown in the
// Run console so the connection is explicit.
const TargetRepo = "integrations/terraform-provider-github"

// DocsPath points users at the bundled documentation tree.
const DocsPath = "./docs"

// Tagline is the one-line product promise shown in the welcome banner.
const Tagline = "Make the GitHub acceptance suite legible."

// Description is the longer one-line description of what Terraform Provider Tester is,
// shown by `terraform-provider-tester version` and in the documentation.
const Description = "Internal acceptance-test harness for integrations/terraform-provider-github."

// Vendor is the legal owner of Terraform Provider Tester. The harness is GitHub property;
// this is surfaced in the banner footer and `terraform-provider-tester version` so ownership is explicit.
const Vendor = "GitHub, Inc."

// License is the SPDX identifier the harness ships under, matching the
// repository's top-level LICENSE (MIT).
const License = "MIT"

// MaintainersRef points at the canonical maintainers list at the repository
// root. Names are not duplicated here so they cannot drift.
const MaintainersRef = "MAINTAINERS.md"

func renderConsoleMark() string {
	return lipgloss.NewStyle().Bold(true).Foreground(Accent).Render("GH") +
		lipgloss.NewStyle().Foreground(Muted).Render("/") +
		lipgloss.NewStyle().Bold(true).Foreground(TerraformPurple).Render("TF")
}

func renderConsoleTitle(includeName bool, version string) string {
	title := renderConsoleMark()
	if includeName {
		title += " " + lipgloss.NewStyle().Bold(true).Render(ConsoleName)
	}
	if version != "" {
		title += " " + lipgloss.NewStyle().Foreground(Muted).Render(version)
	}
	return title
}
