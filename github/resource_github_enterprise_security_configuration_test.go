package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubEnterpriseSecurityConfiguration(t *testing.T) {
	t.Run("creates enterprise security configuration without error", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		configName := fmt.Sprintf("test-config-%s", randomID)

		config := fmt.Sprintf(`
		resource "github_enterprise_security_configuration" "test" {
			enterprise_slug = "%s"
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
		}`, testAccConf.enterpriseSlug, configName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_enterprise_security_configuration.test",
							tfjsonpath.New("name"), knownvalue.StringExact(configName)),
						statecheck.ExpectKnownValue("github_enterprise_security_configuration.test",
							tfjsonpath.New("description"), knownvalue.StringExact("Test configuration")),
						statecheck.ExpectKnownValue("github_enterprise_security_configuration.test",
							tfjsonpath.New("advanced_security"), knownvalue.StringExact("enabled")),
						statecheck.ExpectKnownValue("github_enterprise_security_configuration.test",
							tfjsonpath.New("enforcement"), knownvalue.StringExact("enforced")),
					},
				},
				{
					ResourceName:      "github_enterprise_security_configuration.test",
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})

	t.Run("updates enterprise security configuration without error", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		configName := fmt.Sprintf("test-config-%s", randomID)
		configNameUpdated := fmt.Sprintf("test-config-updated-%s", randomID)

		tmpl := `
		resource "github_enterprise_security_configuration" "test" {
			enterprise_slug = "%s"
			name = "%s"
			description = "%s"
			advanced_security = "%s"
		}`
		configBefore := fmt.Sprintf(tmpl, testAccConf.enterpriseSlug, configName, "Test configuration", "disabled")
		configAfter := fmt.Sprintf(tmpl, testAccConf.enterpriseSlug, configNameUpdated, "Test configuration updated", "enabled")

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: configBefore,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_enterprise_security_configuration.test",
							tfjsonpath.New("name"), knownvalue.StringExact(configName)),
						statecheck.ExpectKnownValue("github_enterprise_security_configuration.test",
							tfjsonpath.New("advanced_security"), knownvalue.StringExact("disabled")),
					},
				},
				{
					Config: configAfter,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_enterprise_security_configuration.test",
							tfjsonpath.New("name"), knownvalue.StringExact(configNameUpdated)),
						statecheck.ExpectKnownValue("github_enterprise_security_configuration.test",
							tfjsonpath.New("advanced_security"), knownvalue.StringExact("enabled")),
					},
				},
			},
		})
	})

	t.Run("creates enterprise security configuration with options", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		configName := fmt.Sprintf("test-config-options-%s", randomID)

		config := fmt.Sprintf(`
		resource "github_enterprise_security_configuration" "test" {
			enterprise_slug = "%s"
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
		}`, testAccConf.enterpriseSlug, configName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_enterprise_security_configuration.test",
							tfjsonpath.New("name"), knownvalue.StringExact(configName)),
						statecheck.ExpectKnownValue("github_enterprise_security_configuration.test",
							tfjsonpath.New("dependency_graph_autosubmit_action_options").AtSliceIndex(0).AtMapKey("labeled_runners"), knownvalue.Bool(true)),
						statecheck.ExpectKnownValue("github_enterprise_security_configuration.test",
							tfjsonpath.New("code_scanning_default_setup_options").AtSliceIndex(0).AtMapKey("runner_type"), knownvalue.StringExact("labeled")),
						statecheck.ExpectKnownValue("github_enterprise_security_configuration.test",
							tfjsonpath.New("code_scanning_default_setup_options").AtSliceIndex(0).AtMapKey("runner_label"), knownvalue.StringExact("code-scanning")),
					},
				},
				{
					ResourceName:      "github_enterprise_security_configuration.test",
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})

	t.Run("creates enterprise security configuration with minimal config", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		configName := fmt.Sprintf("test-config-minimal-%s", randomID)

		config := fmt.Sprintf(`
		resource "github_enterprise_security_configuration" "test" {
			enterprise_slug = "%s"
			name = "%s"
			description = "Minimal test configuration"
		}`, testAccConf.enterpriseSlug, configName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_enterprise_security_configuration.test",
							tfjsonpath.New("name"), knownvalue.StringExact(configName)),
						statecheck.ExpectKnownValue("github_enterprise_security_configuration.test",
							tfjsonpath.New("target_type"), knownvalue.NotNull()),
					},
				},
			},
		})
	})
}
