package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccGithubActionsOrganizationWorkflowPermissions(t *testing.T) {
	t.Run("creates organization workflow permissions without error", func(t *testing.T) {
		config := fmt.Sprintf(`
		resource "github_actions_organization_workflow_permissions" "test" {
			organization_slug = "%s"

			default_workflow_permissions = "read"
			can_approve_pull_request_reviews = false
		}
		`, testAccConf.owner)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("github_actions_organization_workflow_permissions.test", "organization_slug", testAccConf.owner),
			resource.TestCheckResourceAttr("github_actions_organization_workflow_permissions.test", "default_workflow_permissions", "read"),
			resource.TestCheckResourceAttr("github_actions_organization_workflow_permissions.test", "can_approve_pull_request_reviews", "false"),
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

	t.Run("updates organization workflow permissions without error", func(t *testing.T) {
		configs := map[string]string{
			"before": fmt.Sprintf(`
			resource "github_actions_organization_workflow_permissions" "test" {
				organization_slug = "%s"

				default_workflow_permissions = "read"
				can_approve_pull_request_reviews = false
			}
			`, testAccConf.owner),

			"after": fmt.Sprintf(`
			resource "github_actions_organization_workflow_permissions" "test" {
				organization_slug = "%s"

				default_workflow_permissions = "write" // This change might be restricted by the Enterprise's settings
				can_approve_pull_request_reviews = true
			}
			`, testAccConf.owner),
		}

		checks := map[string]resource.TestCheckFunc{
			"before": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr("github_actions_organization_workflow_permissions.test", "default_workflow_permissions", "read"),
				resource.TestCheckResourceAttr("github_actions_organization_workflow_permissions.test", "can_approve_pull_request_reviews", "false"),
			),
			"after": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr("github_actions_organization_workflow_permissions.test", "default_workflow_permissions", "write"),
				resource.TestCheckResourceAttr("github_actions_organization_workflow_permissions.test", "can_approve_pull_request_reviews", "true"),
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

	t.Run("imports organization workflow permissions without error", func(t *testing.T) {
		config := fmt.Sprintf(`
		resource "github_actions_organization_workflow_permissions" "test" {
			organization_slug = "%s"

			default_workflow_permissions = "read"
			can_approve_pull_request_reviews = false
		}
		`, testAccConf.owner)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("github_actions_organization_workflow_permissions.test", "organization_slug", testAccConf.owner),
			resource.TestCheckResourceAttr("github_actions_organization_workflow_permissions.test", "default_workflow_permissions", "read"),
			resource.TestCheckResourceAttr("github_actions_organization_workflow_permissions.test", "can_approve_pull_request_reviews", "false"),
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
					ResourceName:      "github_actions_organization_workflow_permissions.test",
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})

	t.Run("deletes organization workflow permissions without error", func(t *testing.T) {
		config := fmt.Sprintf(`
		resource "github_actions_organization_workflow_permissions" "test" {
			organization_slug = "%s"

			default_workflow_permissions = "write"
			can_approve_pull_request_reviews = true
		}
		`, testAccConf.owner)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:  config,
					Destroy: true,
				},
			},
		})
	})

	t.Run("creates with minimal config using defaults", func(t *testing.T) {
		config := fmt.Sprintf(`
		resource "github_actions_organization_workflow_permissions" "test" {
			organization_slug = "%s"
		}
		`, testAccConf.owner)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("github_actions_organization_workflow_permissions.test", "organization_slug", testAccConf.owner),
			resource.TestCheckResourceAttr("github_actions_organization_workflow_permissions.test", "default_workflow_permissions", "read"),
			resource.TestCheckResourceAttr("github_actions_organization_workflow_permissions.test", "can_approve_pull_request_reviews", "false"),
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
}
