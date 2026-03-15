package github

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubIpRangesDataSource(t *testing.T) {
	t.Run("reads IP ranges without error", func(t *testing.T) {
		config := `data "github_ip_ranges" "test" {}`

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_ip_ranges.test", tfjsonpath.New("actions_ipv4"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_ip_ranges.test", tfjsonpath.New("actions_ipv6"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_ip_ranges.test", tfjsonpath.New("actions_macos_ipv4"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_ip_ranges.test", tfjsonpath.New("actions_macos_ipv6"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_ip_ranges.test", tfjsonpath.New("actions_macos"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_ip_ranges.test", tfjsonpath.New("actions"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_ip_ranges.test", tfjsonpath.New("api_ipv4"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_ip_ranges.test", tfjsonpath.New("api_ipv6"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_ip_ranges.test", tfjsonpath.New("api"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_ip_ranges.test", tfjsonpath.New("dependabot_ipv4"), knownvalue.Null()),
						statecheck.ExpectKnownValue("data.github_ip_ranges.test", tfjsonpath.New("dependabot_ipv6"), knownvalue.Null()),
						statecheck.ExpectKnownValue("data.github_ip_ranges.test", tfjsonpath.New("dependabot"), knownvalue.Null()),
						statecheck.ExpectKnownValue("data.github_ip_ranges.test", tfjsonpath.New("git_ipv4"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_ip_ranges.test", tfjsonpath.New("git_ipv6"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_ip_ranges.test", tfjsonpath.New("git"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_ip_ranges.test", tfjsonpath.New("github_enterprise_importer_ipv4"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_ip_ranges.test", tfjsonpath.New("github_enterprise_importer_ipv6"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_ip_ranges.test", tfjsonpath.New("github_enterprise_importer"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_ip_ranges.test", tfjsonpath.New("hooks_ipv4"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_ip_ranges.test", tfjsonpath.New("hooks_ipv6"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_ip_ranges.test", tfjsonpath.New("hooks"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_ip_ranges.test", tfjsonpath.New("importer_ipv4"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_ip_ranges.test", tfjsonpath.New("importer_ipv6"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_ip_ranges.test", tfjsonpath.New("importer"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_ip_ranges.test", tfjsonpath.New("packages_ipv4"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_ip_ranges.test", tfjsonpath.New("packages_ipv6"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_ip_ranges.test", tfjsonpath.New("packages"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_ip_ranges.test", tfjsonpath.New("pages_ipv4"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_ip_ranges.test", tfjsonpath.New("pages_ipv6"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_ip_ranges.test", tfjsonpath.New("pages"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_ip_ranges.test", tfjsonpath.New("web_ipv4"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_ip_ranges.test", tfjsonpath.New("web_ipv6"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_ip_ranges.test", tfjsonpath.New("web"), knownvalue.NotNull()),
					},
				},
			},
		})
	})
}
