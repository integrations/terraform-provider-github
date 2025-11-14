package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubEnterpriseActionsWorkflowPermissions(t *testing.T) {
	t.Run("creates enterprise workflow permissions without error", func(t *testing.T) {
		config := fmt.Sprintf(`
		resource "github_enterprise_actions_workflow_permissions" "test" {
			enterprise_slug = "%s"
			
			default_workflow_permissions = "read"
			can_approve_pull_request_reviews = false
		}
		`, testEnterprise)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("github_enterprise_actions_workflow_permissions.test", "enterprise_slug", testEnterprise),
			resource.TestCheckResourceAttr("github_enterprise_actions_workflow_permissions.test", "default_workflow_permissions", "read"),
			resource.TestCheckResourceAttr("github_enterprise_actions_workflow_permissions.test", "can_approve_pull_request_reviews", "false"),
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

		t.Run("with an enterprise account", func(t *testing.T) {
			if isEnterprise != "true" {
				t.Skip("Skipping because `ENTERPRISE_ACCOUNT` is not set or set to false")
			}
			if testEnterprise == "" {
				t.Skip("Skipping because `ENTERPRISE_SLUG` is not set")
			}
			testCase(t, enterprise)
		})
	})

	t.Run("updates enterprise workflow permissions without error", func(t *testing.T) {
		configs := map[string]string{
			"before": fmt.Sprintf(`
			resource "github_enterprise_actions_workflow_permissions" "test" {
				enterprise_slug = "%s"
				
				default_workflow_permissions = "read"
				can_approve_pull_request_reviews = false
			}
			`, testEnterprise),

			"after": fmt.Sprintf(`
			resource "github_enterprise_actions_workflow_permissions" "test" {
				enterprise_slug = "%s"
				
				default_workflow_permissions = "write"
				can_approve_pull_request_reviews = true
			}
			`, testEnterprise),
		}

		checks := map[string]resource.TestCheckFunc{
			"before": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr("github_enterprise_actions_workflow_permissions.test", "default_workflow_permissions", "read"),
				resource.TestCheckResourceAttr("github_enterprise_actions_workflow_permissions.test", "can_approve_pull_request_reviews", "false"),
			),
			"after": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr("github_enterprise_actions_workflow_permissions.test", "default_workflow_permissions", "write"),
				resource.TestCheckResourceAttr("github_enterprise_actions_workflow_permissions.test", "can_approve_pull_request_reviews", "true"),
			),
		}

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

		t.Run("with an enterprise account", func(t *testing.T) {
			if isEnterprise != "true" {
				t.Skip("Skipping because `ENTERPRISE_ACCOUNT` is not set or set to false")
			}
			if testEnterprise == "" {
				t.Skip("Skipping because `ENTERPRISE_SLUG` is not set")
			}
			testCase(t, enterprise)
		})
	})

	t.Run("imports enterprise workflow permissions without error", func(t *testing.T) {
		config := fmt.Sprintf(`
		resource "github_enterprise_actions_workflow_permissions" "test" {
			enterprise_slug = "%s"
			
			default_workflow_permissions = "read"
			can_approve_pull_request_reviews = false
		}
		`, testEnterprise)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("github_enterprise_actions_workflow_permissions.test", "enterprise_slug", testEnterprise),
			resource.TestCheckResourceAttr("github_enterprise_actions_workflow_permissions.test", "default_workflow_permissions", "read"),
			resource.TestCheckResourceAttr("github_enterprise_actions_workflow_permissions.test", "can_approve_pull_request_reviews", "false"),
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
						ResourceName:      "github_enterprise_actions_workflow_permissions.test",
						ImportState:       true,
						ImportStateVerify: true,
					},
				},
			})
		}

		t.Run("with an enterprise account", func(t *testing.T) {
			if isEnterprise != "true" {
				t.Skip("Skipping because `ENTERPRISE_ACCOUNT` is not set or set to false")
			}
			if testEnterprise == "" {
				t.Skip("Skipping because `ENTERPRISE_SLUG` is not set")
			}
			testCase(t, enterprise)
		})
	})
}
