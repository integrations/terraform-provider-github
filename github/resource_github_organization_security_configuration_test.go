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

func TestAccGithubOrganizationSecurityConfiguration(t *testing.T) {
	t.Run("creates organization security configuration without error", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		configName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)

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

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_organization_security_configuration.test",
							tfjsonpath.New("name"), knownvalue.StringExact(configName)),
						statecheck.ExpectKnownValue("github_organization_security_configuration.test",
							tfjsonpath.New("description"), knownvalue.StringExact("Test configuration")),
						statecheck.ExpectKnownValue("github_organization_security_configuration.test",
							tfjsonpath.New("advanced_security"), knownvalue.StringExact("enabled")),
						statecheck.ExpectKnownValue("github_organization_security_configuration.test",
							tfjsonpath.New("enforcement"), knownvalue.StringExact("enforced")),
						statecheck.ExpectKnownValue("github_organization_security_configuration.test",
							tfjsonpath.New("configuration_id"), knownvalue.NotNull()),
					},
				},
			},
		})
	})

	t.Run("imports organization security configuration without error", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		configName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
		resource "github_organization_security_configuration" "test" {
			name = "%s"
			description = "Test configuration for import"
		}`, configName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
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
		configName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)
		configNameUpdated := fmt.Sprintf("%supdated-%s", testResourcePrefix, randomID)

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

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: configBefore,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_organization_security_configuration.test",
							tfjsonpath.New("name"), knownvalue.StringExact(configName)),
						statecheck.ExpectKnownValue("github_organization_security_configuration.test",
							tfjsonpath.New("advanced_security"), knownvalue.StringExact("disabled")),
					},
				},
				{
					Config: configAfter,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_organization_security_configuration.test",
							tfjsonpath.New("name"), knownvalue.StringExact(configNameUpdated)),
						statecheck.ExpectKnownValue("github_organization_security_configuration.test",
							tfjsonpath.New("advanced_security"), knownvalue.StringExact("enabled")),
					},
				},
			},
		})
	})

	t.Run("creates organization security configuration with options", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		configName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)

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

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_organization_security_configuration.test",
							tfjsonpath.New("name"), knownvalue.StringExact(configName)),
						statecheck.ExpectKnownValue("github_organization_security_configuration.test",
							tfjsonpath.New("dependency_graph_autosubmit_action_options").AtSliceIndex(0).AtMapKey("labeled_runners"), knownvalue.Bool(true)),
						statecheck.ExpectKnownValue("github_organization_security_configuration.test",
							tfjsonpath.New("code_scanning_default_setup_options").AtSliceIndex(0).AtMapKey("runner_type"), knownvalue.StringExact("labeled")),
						statecheck.ExpectKnownValue("github_organization_security_configuration.test",
							tfjsonpath.New("code_scanning_default_setup_options").AtSliceIndex(0).AtMapKey("runner_label"), knownvalue.StringExact("code-scanning")),
					},
				},
			},
		})
	})

	t.Run("creates organization security configuration with minimal config", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		configName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
		resource "github_organization_security_configuration" "test" {
			name = "%s"
			description = "Minimal test configuration"
		}`, configName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_organization_security_configuration.test",
							tfjsonpath.New("name"), knownvalue.StringExact(configName)),
						statecheck.ExpectKnownValue("github_organization_security_configuration.test",
							tfjsonpath.New("target_type"), knownvalue.NotNull()),
					},
				},
			},
		})
	})

	t.Run("creates organization security configuration with delegated bypass options", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		configName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
		resource "github_organization_security_configuration" "test" {
			name = "%s"
			description = "Test configuration with delegated bypass"
			advanced_security = "enabled"
			secret_scanning = "enabled"
			secret_scanning_push_protection = "enabled"
			secret_scanning_delegated_bypass = "enabled"
			secret_scanning_delegated_bypass_options {
				reviewers {
					reviewer_id   = 1
					reviewer_type = "TEAM"
				}
			}
		}`, configName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_organization_security_configuration.test",
							tfjsonpath.New("secret_scanning_delegated_bypass"), knownvalue.StringExact("enabled")),
						statecheck.ExpectKnownValue("github_organization_security_configuration.test",
							tfjsonpath.New("secret_scanning_delegated_bypass_options").AtSliceIndex(0).AtMapKey("reviewers").AtSliceIndex(0).AtMapKey("reviewer_type"), knownvalue.StringExact("TEAM")),
					},
				},
			},
		})
	})
}
