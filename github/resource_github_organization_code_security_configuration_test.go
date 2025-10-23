package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestGithubOrganizationCodeSecurityConfiguration(t *testing.T) {
	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

	t.Run("creates code security configuration without error", func(t *testing.T) {
		config := fmt.Sprintf(`
			resource "github_organization_code_security_configuration" "test" {
			  name        = "tf-acc-test-%s"
			  description = "Test code security configuration description"
			  advanced_security = "enabled"
			  dependency_graph = "enabled"
			  dependency_graph_autosubmit_action_options {
			  	labeled_runners = true
			  }
			  dependabot_alerts = "enabled"
			  dependabot_security_updates = "enabled"
			  secret_scanning = "enabled"
			  secret_scanning_push_protection = "enabled"
			  secret_scanning_validity_checks = "enabled"
			  secret_scanning_non_provider_patterns = "enabled"
			  private_vulnerability_reporting = "enabled"
			  enforcement = "unenforced"
			}
		`, randomID)
		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet(
				"github_organization_code_security_configuration.test", "name",
			),
			resource.TestCheckResourceAttr(
				"github_organization_code_security_configuration.test", "name",
				fmt.Sprintf(`tf-acc-test-%s`, randomID),
			),
			resource.TestCheckResourceAttr(
				"github_organization_code_security_configuration.test", "description",
				"Test code security configuration description",
			),
			resource.TestCheckResourceAttr(
				"github_organization_code_security_configuration.test", "advanced_security",
				"enabled",
			),
			resource.TestCheckResourceAttr(
				"github_organization_code_security_configuration.test", "dependency_graph",
				"enabled",
			),
			resource.TestCheckTypeSetElemNestedAttrs(
				"github_organization_code_security_configuration.test", "dependency_graph_autosubmit_action_options.*",
				map[string]string{
					"labeled_runners": "true",
				},
			),
			resource.TestCheckResourceAttr(
				"github_organization_code_security_configuration.test", "dependabot_alerts",
				"enabled",
			),
			resource.TestCheckResourceAttr(
				"github_organization_code_security_configuration.test", "dependabot_security_updates",
				"enabled",
			),
			resource.TestCheckResourceAttr(
				"github_organization_code_security_configuration.test", "secret_scanning",
				"enabled",
			),
			resource.TestCheckResourceAttr(
				"github_organization_code_security_configuration.test", "secret_scanning_push_protection",
				"enabled",
			),
			resource.TestCheckResourceAttr(
				"github_organization_code_security_configuration.test", "secret_scanning_validity_checks",
				"enabled",
			),
			resource.TestCheckResourceAttr(
				"github_organization_code_security_configuration.test", "secret_scanning_non_provider_patterns",
				"enabled",
			),
			resource.TestCheckResourceAttr(
				"github_organization_code_security_configuration.test", "private_vulnerability_reporting",
				"enabled",
			),
			resource.TestCheckResourceAttr(
				"github_organization_code_security_configuration.test", "enforcement",
				"unenforced",
			),
		)

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: config,
						Check:  check,
					},
				},
			})
		}

		t.Run("with an anonymous account", func(t *testing.T) {
			t.Skip("anonymous account not supported for this operation")
		})

		t.Run("with an individual account", func(t *testing.T) {
			t.Skip("individual account not supported for this operation")
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})
	})

	t.Run("updates code security configuration without error", func(t *testing.T) {

		configs := map[string]string{
			"before": fmt.Sprintf(`
				resource "github_organization_code_security_configuration" "test" {
					name        = "tf-acc-test-%s"
					description = "Test code security configuration description"
					enforcement = "unenforced"
				}
			`, randomID),
			"after": fmt.Sprintf(`
				resource "github_organization_code_security_configuration" "test" {
					name        = "tf-acc-test-%s"
					description = "Test code security configuration description"
					enforcement = "enforced"
				}
			`, randomID),
		}

		checks := map[string]resource.TestCheckFunc{
			"before": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttrSet(
					"github_organization_code_security_configuration.test", "name",
				),
				resource.TestCheckResourceAttr(
					"github_organization_code_security_configuration.test", "name",
					fmt.Sprintf(`tf-acc-test-%s`, randomID),
				),
				resource.TestCheckResourceAttr(
					"github_organization_code_security_configuration.test", "description",
					"Test code security configuration description",
				),
				resource.TestCheckResourceAttr(
					"github_organization_code_security_configuration.test", "enforcement",
					"unenforced",
				),
			),
			"after": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttrSet(
					"github_organization_code_security_configuration.test", "name",
				),
				resource.TestCheckResourceAttr(
					"github_organization_code_security_configuration.test", "name",
					fmt.Sprintf(`tf-acc-test-%s`, randomID),
				),
				resource.TestCheckResourceAttr(
					"github_organization_code_security_configuration.test", "description",
					"Test code security configuration description",
				),
				resource.TestCheckResourceAttr(
					"github_organization_code_security_configuration.test", "enforcement",
					"enforced",
				),
			)}

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
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
		}

		t.Run("with an anonymous account", func(t *testing.T) {
			t.Skip("anonymous account not supported for this operation")
		})

		t.Run("with an individual account", func(t *testing.T) {
			t.Skip("individual account not supported for this operation")
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})
	})

	t.Run("imports code security configuration without error", func(t *testing.T) {

		config := fmt.Sprintf(`
			resource "github_organization_code_security_configuration" "test" {
			  name        = "tf-acc-test-%s"
			  description = "Test code security configuration description"
			}
    `, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet(
				"github_organization_code_security_configuration.test", "name",
			),
			resource.TestCheckResourceAttr(
				"github_organization_code_security_configuration.test", "name",
				fmt.Sprintf(`tf-acc-test-%s`, randomID),
			),
			resource.TestCheckResourceAttr(
				"github_organization_code_security_configuration.test", "description",
				"Test code security configuration description",
			),
			resource.TestCheckResourceAttr(
				"github_organization_code_security_configuration.test", "advanced_security",
				"disabled",
			),
			resource.TestCheckResourceAttr(
				"github_organization_code_security_configuration.test", "dependency_graph",
				"enabled",
			),
			resource.TestCheckResourceAttr(
				"github_organization_code_security_configuration.test", "dependabot_alerts",
				"disabled",
			),
			resource.TestCheckResourceAttr(
				"github_organization_code_security_configuration.test", "dependabot_security_updates",
				"disabled",
			),
			resource.TestCheckResourceAttr(
				"github_organization_code_security_configuration.test", "secret_scanning",
				"disabled",
			),
			resource.TestCheckResourceAttr(
				"github_organization_code_security_configuration.test", "secret_scanning_push_protection",
				"disabled",
			),
			resource.TestCheckResourceAttr(
				"github_organization_code_security_configuration.test", "secret_scanning_validity_checks",
				"disabled",
			),
			resource.TestCheckResourceAttr(
				"github_organization_code_security_configuration.test", "secret_scanning_non_provider_patterns",
				"disabled",
			),
			resource.TestCheckResourceAttr(
				"github_organization_code_security_configuration.test", "private_vulnerability_reporting",
				"disabled",
			),
			resource.TestCheckResourceAttr(
				"github_organization_code_security_configuration.test", "enforcement",
				"enforced",
			),
		)

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: config,
						Check:  check,
					},
					{
						ResourceName:      "github_organization_code_security_configuration.test",
						ImportState:       true,
						ImportStateVerify: true,
					},
				},
			})
		}

		t.Run("with an anonymous account", func(t *testing.T) {
			t.Skip("anonymous account not supported for this operation")
		})

		t.Run("with an individual account", func(t *testing.T) {
			t.Skip("individual account not supported for this operation")
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})
	})
}
