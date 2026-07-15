package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
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

		checks := map[string]resource.TestCheckFunc{
			"before": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr("github_code_security_configuration.test", "name", fmt.Sprintf("tf-acc-test-%s", randomID)),
				resource.TestCheckResourceAttr("github_code_security_configuration.test", "description", "Terraform acceptance test configuration"),
				resource.TestCheckResourceAttr("github_code_security_configuration.test", "dependency_graph", "enabled"),
				resource.TestCheckResourceAttr("github_code_security_configuration.test", "dependabot_alerts", "disabled"),
				resource.TestCheckResourceAttr("github_code_security_configuration.test", "private_vulnerability_reporting", "disabled"),
				resource.TestCheckResourceAttr("github_code_security_configuration.test", "enforcement", "unenforced"),
				resource.TestCheckResourceAttr("github_code_security_configuration.test", "target_type", "organization"),
				resource.TestCheckResourceAttrSet("github_code_security_configuration.test", "configuration_id"),
				resource.TestCheckResourceAttrSet("github_code_security_configuration.test", "html_url"),
			),
			"after": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr("github_code_security_configuration.test", "description", "Terraform acceptance test configuration (updated)"),
				resource.TestCheckResourceAttr("github_code_security_configuration.test", "dependabot_alerts", "enabled"),
				resource.TestCheckResourceAttr("github_code_security_configuration.test", "private_vulnerability_reporting", "enabled"),
				resource.TestCheckResourceAttr("github_code_security_configuration.test", "enforcement", "enforced"),
			),
		}

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: configs["before"],
					Check:  checks["before"],
				},
				{
					Config: configs["after"],
					Check:  checks["after"],
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
					Check:  resource.TestCheckResourceAttr("github_code_security_configuration.test", "name", fmt.Sprintf("tf-acc-test-%s", randomID)),
				},
				{
					ResourceName:      "github_code_security_configuration.test",
					ImportState:       true,
					ImportStateVerify: true,
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

		checks := map[string]resource.TestCheckFunc{
			"before": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr("github_code_security_configuration.test", "enterprise_slug", testAccConf.enterpriseSlug),
				resource.TestCheckResourceAttr("github_code_security_configuration.test", "name", fmt.Sprintf("tf-acc-test-%s", randomID)),
				resource.TestCheckResourceAttr("github_code_security_configuration.test", "dependabot_alerts", "disabled"),
				resource.TestCheckResourceAttr("github_code_security_configuration.test", "target_type", "enterprise"),
			),
			"after": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr("github_code_security_configuration.test", "dependabot_alerts", "enabled"),
			),
		}

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: configs["before"],
					Check:  checks["before"],
				},
				{
					Config: configs["after"],
					Check:  checks["after"],
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
