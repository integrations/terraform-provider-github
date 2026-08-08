package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubCodeSecurityConfiguration(t *testing.T) {
	t.Run("creates and updates an organization configuration without error", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

		configs := map[string]string{
			"before": fmt.Sprintf(`
			resource "github_code_security_configuration" "test" {
				name        = "tf-acc-test-%s"
				description = "Terraform acceptance test configuration"

				dependency_graph                = "enabled"
				dependabot_alerts               = "disabled"
				private_vulnerability_reporting = "disabled"
				enforcement                     = "unenforced"
			}
			`, randomID),

			"after": fmt.Sprintf(`
			resource "github_code_security_configuration" "test" {
				name        = "tf-acc-test-%s"
				description = "Terraform acceptance test configuration (updated)"

				dependency_graph                = "enabled"
				dependabot_alerts               = "enabled"
				private_vulnerability_reporting = "enabled"
				enforcement                     = "enforced"
			}
			`, randomID),
		}

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: configs["before"],
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_code_security_configuration.test", tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("tf-acc-test-%s", randomID))),
						statecheck.ExpectKnownValue("github_code_security_configuration.test", tfjsonpath.New("description"), knownvalue.StringExact("Terraform acceptance test configuration")),
						statecheck.ExpectKnownValue("github_code_security_configuration.test", tfjsonpath.New("dependency_graph"), knownvalue.StringExact("enabled")),
						statecheck.ExpectKnownValue("github_code_security_configuration.test", tfjsonpath.New("dependabot_alerts"), knownvalue.StringExact("disabled")),
						statecheck.ExpectKnownValue("github_code_security_configuration.test", tfjsonpath.New("private_vulnerability_reporting"), knownvalue.StringExact("disabled")),
						statecheck.ExpectKnownValue("github_code_security_configuration.test", tfjsonpath.New("enforcement"), knownvalue.StringExact("unenforced")),
						statecheck.ExpectKnownValue("github_code_security_configuration.test", tfjsonpath.New("target_type"), knownvalue.StringExact("organization")),
						statecheck.ExpectKnownValue("github_code_security_configuration.test", tfjsonpath.New("configuration_id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_code_security_configuration.test", tfjsonpath.New("html_url"), knownvalue.NotNull()),
					},
				},
				{
					Config: configs["after"],
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_code_security_configuration.test", plancheck.ResourceActionUpdate),
						},
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_code_security_configuration.test", tfjsonpath.New("description"), knownvalue.StringExact("Terraform acceptance test configuration (updated)")),
						statecheck.ExpectKnownValue("github_code_security_configuration.test", tfjsonpath.New("dependabot_alerts"), knownvalue.StringExact("enabled")),
						statecheck.ExpectKnownValue("github_code_security_configuration.test", tfjsonpath.New("private_vulnerability_reporting"), knownvalue.StringExact("enabled")),
						statecheck.ExpectKnownValue("github_code_security_configuration.test", tfjsonpath.New("enforcement"), knownvalue.StringExact("enforced")),
					},
				},
			},
		})
	})

	t.Run("imports an organization configuration without error", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

		config := fmt.Sprintf(`
		resource "github_code_security_configuration" "test" {
			name        = "tf-acc-test-%s"
			description = "Terraform acceptance test import configuration"

			dependency_graph  = "enabled"
			dependabot_alerts = "enabled"
			enforcement       = "unenforced"
		}
		`, randomID)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_code_security_configuration.test", tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("tf-acc-test-%s", randomID))),
					},
				},
				{
					ResourceName:      "github_code_security_configuration.test",
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})

	t.Run("manages default_for_new_repos on an organization configuration without error", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

		configs := map[string]string{
			"set": fmt.Sprintf(`
			resource "github_code_security_configuration" "test" {
				name        = "tf-acc-test-%s"
				description = "Terraform acceptance test default configuration"

				dependency_graph  = "enabled"
				dependabot_alerts = "enabled"
				enforcement       = "unenforced"

				default_for_new_repos = "private_and_internal"
			}
			`, randomID),

			"changed": fmt.Sprintf(`
			resource "github_code_security_configuration" "test" {
				name        = "tf-acc-test-%s"
				description = "Terraform acceptance test default configuration"

				dependency_graph  = "enabled"
				dependabot_alerts = "enabled"
				enforcement       = "unenforced"

				default_for_new_repos = "all"
			}
			`, randomID),

			"removed": fmt.Sprintf(`
			resource "github_code_security_configuration" "test" {
				name        = "tf-acc-test-%s"
				description = "Terraform acceptance test default configuration"

				dependency_graph  = "enabled"
				dependabot_alerts = "enabled"
				enforcement       = "unenforced"
			}
			`, randomID),
		}

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: configs["set"],
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_code_security_configuration.test", tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("tf-acc-test-%s", randomID))),
						statecheck.ExpectKnownValue("github_code_security_configuration.test", tfjsonpath.New("default_for_new_repos"), knownvalue.StringExact("private_and_internal")),
					},
				},
				{
					Config: configs["changed"],
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_code_security_configuration.test", plancheck.ResourceActionUpdate),
						},
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_code_security_configuration.test", tfjsonpath.New("default_for_new_repos"), knownvalue.StringExact("all")),
					},
				},
				{
					Config: configs["removed"],
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_code_security_configuration.test", tfjsonpath.New("default_for_new_repos"), knownvalue.StringExact("")),
					},
				},
				{
					Config: configs["removed"],
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
			},
		})
	})

	t.Run("attaches an organization configuration to repositories by scope without error", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

		// attach_scope is write-only: the attachment cannot be read back from
		// the API, so these steps only assert that the configured value is
		// held in state and that re-planning the same config is a no-op.
		config := fmt.Sprintf(`
		resource "github_code_security_configuration" "test" {
			name        = "tf-acc-test-%s"
			description = "Terraform acceptance test attach configuration"

			dependency_graph  = "enabled"
			dependabot_alerts = "enabled"
			enforcement       = "unenforced"

			attach_scope = "all_without_configurations"
		}
		`, randomID)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_code_security_configuration.test", tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("tf-acc-test-%s", randomID))),
						statecheck.ExpectKnownValue("github_code_security_configuration.test", tfjsonpath.New("attach_scope"), knownvalue.StringExact("all_without_configurations")),
						statecheck.ExpectKnownValue("github_code_security_configuration.test", tfjsonpath.New("configuration_id"), knownvalue.NotNull()),
					},
				},
				{
					Config: config,
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
			},
		})
	})

	t.Run("creates, updates and imports an enterprise configuration without error", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

		configs := map[string]string{
			"before": fmt.Sprintf(`
			resource "github_code_security_configuration" "test" {
				enterprise_slug = "%s"
				name            = "tf-acc-test-%s"
				description     = "Terraform acceptance test enterprise configuration"

				dependency_graph  = "enabled"
				dependabot_alerts = "disabled"
				enforcement       = "unenforced"
			}
			`, testAccConf.enterpriseSlug, randomID),

			"after": fmt.Sprintf(`
			resource "github_code_security_configuration" "test" {
				enterprise_slug = "%s"
				name            = "tf-acc-test-%s"
				description     = "Terraform acceptance test enterprise configuration"

				dependency_graph  = "enabled"
				dependabot_alerts = "enabled"
				enforcement       = "unenforced"
			}
			`, testAccConf.enterpriseSlug, randomID),
		}

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: configs["before"],
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_code_security_configuration.test", tfjsonpath.New("enterprise_slug"), knownvalue.StringExact(testAccConf.enterpriseSlug)),
						statecheck.ExpectKnownValue("github_code_security_configuration.test", tfjsonpath.New("name"), knownvalue.StringExact(fmt.Sprintf("tf-acc-test-%s", randomID))),
						statecheck.ExpectKnownValue("github_code_security_configuration.test", tfjsonpath.New("dependabot_alerts"), knownvalue.StringExact("disabled")),
						statecheck.ExpectKnownValue("github_code_security_configuration.test", tfjsonpath.New("target_type"), knownvalue.StringExact("enterprise")),
					},
				},
				{
					Config: configs["after"],
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_code_security_configuration.test", plancheck.ResourceActionUpdate),
						},
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_code_security_configuration.test", tfjsonpath.New("dependabot_alerts"), knownvalue.StringExact("enabled")),
					},
				},
				{
					ResourceName:      "github_code_security_configuration.test",
					ImportState:       true,
					ImportStateVerify: true,
					ImportStateIdFunc: importCodeSecurityConfigurationByEnterprise("github_code_security_configuration.test"),
				},
			},
		})
	})
}

// importCodeSecurityConfigurationByEnterprise builds an import ID of the form
// <enterprise_slug>:<configuration_id> from the resource in state.
func importCodeSecurityConfigurationByEnterprise(logicalName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs := s.RootModule().Resources[logicalName]
		if rs == nil {
			return "", fmt.Errorf("cannot find %s in terraform state", logicalName)
		}
		return fmt.Sprintf("%s:%s", testAccConf.enterpriseSlug, rs.Primary.ID), nil
	}
}
