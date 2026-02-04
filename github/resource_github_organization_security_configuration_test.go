package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubOrganizationSecurityConfiguration(t *testing.T) {
	t.Run("creates organization security configuration without error", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		configName := fmt.Sprintf("test-config-%s", randomID)

		config := fmt.Sprintf(`
		resource "github_organization_security_configuration" "test" {
			name = "%s"
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
		}`, configName)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_organization_security_configuration.test",
				"name", configName,
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
				{
					ResourceName:      "github_organization_security_configuration.test",
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})

	t.Run("updates organization security configuration without error", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		configName := fmt.Sprintf("test-config-%s", randomID)
		configNameUpdated := fmt.Sprintf("test-config-updated-%s", randomID)

		configBefore := fmt.Sprintf(`
		resource "github_organization_security_configuration" "test" {
			name = "%s"
			description = "Test configuration"
			advanced_security = "disabled"
		}`, configName)

		configAfter := fmt.Sprintf(`
		resource "github_organization_security_configuration" "test" {
			name = "%s"
			description = "Test configuration updated"
			advanced_security = "enabled"
		}`, configNameUpdated)

		checkBefore := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_organization_security_configuration.test",
				"name", configName,
			),
			resource.TestCheckResourceAttr(
				"github_organization_security_configuration.test",
				"advanced_security", "disabled",
			),
		)

		checkAfter := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_organization_security_configuration.test",
				"name", configNameUpdated,
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

	t.Run("creates organization security configuration with options", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		configName := fmt.Sprintf("test-config-options-%s", randomID)

		config := fmt.Sprintf(`
		resource "github_organization_security_configuration" "test" {
			name = "%s"
			description = "Test configuration with options"
			advanced_security = "enabled"
			dependency_graph = "enabled"
			dependency_graph_autosubmit_action = "enabled"
			dependency_graph_autosubmit_action_options {
				labeled_runners = true
			}
			code_scanning_default_setup = "enabled"
			code_scanning_default_setup_options {
				runner_type = "labeled"
				runner_label = "code-scanning"
			}
			secret_scanning = "enabled"
			secret_scanning_push_protection = "enabled"
		}`, configName)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_organization_security_configuration.test",
				"name", configName,
			),
			resource.TestCheckResourceAttr(
				"github_organization_security_configuration.test",
				"dependency_graph_autosubmit_action_options.0.labeled_runners", "true",
			),
			resource.TestCheckResourceAttr(
				"github_organization_security_configuration.test",
				"code_scanning_default_setup_options.0.runner_type", "labeled",
			),
			resource.TestCheckResourceAttr(
				"github_organization_security_configuration.test",
				"code_scanning_default_setup_options.0.runner_label", "code-scanning",
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
				{
					ResourceName:      "github_organization_security_configuration.test",
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})
}
