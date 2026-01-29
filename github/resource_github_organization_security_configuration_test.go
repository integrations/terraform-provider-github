package github

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubOrganizationSecurityConfiguration(t *testing.T) {
	t.Run("creates organization security configuration without error", func(t *testing.T) {
		config := `
		resource "github_organization_security_configuration" "test" {
			name = "test-config"
			description = "Test configuration"
			advanced_security = "enabled"
			dependency_graph = "enabled"
			dependabot_alerts = "enabled"
			dependabot_security_updates = "enabled"
			code_scanning_default_setup = "enabled"
			secret_scanning = "enabled"
			secret_scanning_push_protection = "enabled"
			private_vulnerability_reporting = "enabled"
			enforcement = "enforced"
			target_type = "global"
		}`

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_organization_security_configuration.test",
				"name", "test-config",
			),
			resource.TestCheckResourceAttr(
				"github_organization_security_configuration.test",
				"description", "Test configuration",
			),
			resource.TestCheckResourceAttr(
				"github_organization_security_configuration.test",
				"advanced_security", "enabled",
			),
			resource.TestCheckResourceAttr(
				"github_organization_security_configuration.test",
				"enforcement", "enforced",
			),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  check,
				},
			},
		})
	})

	t.Run("updates organization security configuration without error", func(t *testing.T) {
		configBefore := `
		resource "github_organization_security_configuration" "test" {
			name = "test-config"
			description = "Test configuration"
			advanced_security = "disabled"
		}`

		configAfter := `
		resource "github_organization_security_configuration" "test" {
			name = "test-config-updated"
			description = "Test configuration updated"
			advanced_security = "enabled"
		}`

		checkBefore := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_organization_security_configuration.test",
				"name", "test-config",
			),
			resource.TestCheckResourceAttr(
				"github_organization_security_configuration.test",
				"advanced_security", "disabled",
			),
		)

		checkAfter := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_organization_security_configuration.test",
				"name", "test-config-updated",
			),
			resource.TestCheckResourceAttr(
				"github_organization_security_configuration.test",
				"advanced_security", "enabled",
			),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: configBefore,
					Check:  checkBefore,
				},
				{
					Config: configAfter,
					Check:  checkAfter,
				},
			},
		})
	})
}
