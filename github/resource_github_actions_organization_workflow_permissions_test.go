package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubActionsOrganizationWorkflowPermissions(t *testing.T) {
	t.Run("creates organization workflow permissions without error", func(t *testing.T) {
		defaultPermission := "write"
		canApprovePRReviews := true

		config := fmt.Sprintf(`
resource "github_actions_organization_workflow_permissions" "test" {
	organization_slug = "%s"

	default_workflow_permissions     = "%s"
	can_approve_pull_request_reviews = %t
}
`, testAccConf.owner, defaultPermission, canApprovePRReviews)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_actions_organization_workflow_permissions.test", tfjsonpath.New("id"), knownvalue.StringExact(testAccConf.owner)),
						statecheck.ExpectKnownValue("github_actions_organization_workflow_permissions.test", tfjsonpath.New("organization_slug"), knownvalue.StringExact(testAccConf.owner)),
						statecheck.ExpectKnownValue("github_actions_organization_workflow_permissions.test", tfjsonpath.New("default_workflow_permissions"), knownvalue.StringExact(defaultPermission)),
						statecheck.ExpectKnownValue("github_actions_organization_workflow_permissions.test", tfjsonpath.New("can_approve_pull_request_reviews"), knownvalue.Bool(canApprovePRReviews)),
					},
				},
			},
		})
	})

	t.Run("updates organization workflow permissions without error", func(t *testing.T) {
		defaultPermission := "write"
		canApprovePRReviews := true
		defaultPermissionUpdated := "read"
		canApprovePRReviewsUpdated := false

		config := `
resource "github_actions_organization_workflow_permissions" "test" {
	organization_slug = "%s"

	default_workflow_permissions     = "%s"
	can_approve_pull_request_reviews = %t
}
`

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, testAccConf.owner, defaultPermission, canApprovePRReviews),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_actions_organization_workflow_permissions.test", tfjsonpath.New("id"), knownvalue.StringExact(testAccConf.owner)),
						statecheck.ExpectKnownValue("github_actions_organization_workflow_permissions.test", tfjsonpath.New("organization_slug"), knownvalue.StringExact(testAccConf.owner)),
						statecheck.ExpectKnownValue("github_actions_organization_workflow_permissions.test", tfjsonpath.New("default_workflow_permissions"), knownvalue.StringExact(defaultPermission)),
						statecheck.ExpectKnownValue("github_actions_organization_workflow_permissions.test", tfjsonpath.New("can_approve_pull_request_reviews"), knownvalue.Bool(canApprovePRReviews)),
					},
				},
				{
					Config: fmt.Sprintf(config, testAccConf.owner, defaultPermissionUpdated, canApprovePRReviewsUpdated),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_actions_organization_workflow_permissions.test", tfjsonpath.New("id"), knownvalue.StringExact(testAccConf.owner)),
						statecheck.ExpectKnownValue("github_actions_organization_workflow_permissions.test", tfjsonpath.New("organization_slug"), knownvalue.StringExact(testAccConf.owner)),
						statecheck.ExpectKnownValue("github_actions_organization_workflow_permissions.test", tfjsonpath.New("default_workflow_permissions"), knownvalue.StringExact(defaultPermissionUpdated)),
						statecheck.ExpectKnownValue("github_actions_organization_workflow_permissions.test", tfjsonpath.New("can_approve_pull_request_reviews"), knownvalue.Bool(canApprovePRReviewsUpdated)),
					},
				},
			},
		})
	})

	t.Run("imports organization workflow permissions without error", func(t *testing.T) {
		defaultPermission := "write"
		canApprovePRReviews := true

		config := fmt.Sprintf(`
resource "github_actions_organization_workflow_permissions" "test" {
	organization_slug = "%s"

	default_workflow_permissions     = "%s"
	can_approve_pull_request_reviews = %t
}
`, testAccConf.owner, defaultPermission, canApprovePRReviews)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_actions_organization_workflow_permissions.test", tfjsonpath.New("id"), knownvalue.StringExact(testAccConf.owner)),
						statecheck.ExpectKnownValue("github_actions_organization_workflow_permissions.test", tfjsonpath.New("organization_slug"), knownvalue.StringExact(testAccConf.owner)),
						statecheck.ExpectKnownValue("github_actions_organization_workflow_permissions.test", tfjsonpath.New("default_workflow_permissions"), knownvalue.StringExact(defaultPermission)),
						statecheck.ExpectKnownValue("github_actions_organization_workflow_permissions.test", tfjsonpath.New("can_approve_pull_request_reviews"), knownvalue.Bool(canApprovePRReviews)),
					},
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

	default_workflow_permissions     = "write"
	can_approve_pull_request_reviews = true
}
`, testAccConf.owner)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
				},
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

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_actions_organization_workflow_permissions.test", tfjsonpath.New("id"), knownvalue.StringExact(testAccConf.owner)),
						statecheck.ExpectKnownValue("github_actions_organization_workflow_permissions.test", tfjsonpath.New("organization_slug"), knownvalue.StringExact(testAccConf.owner)),
						statecheck.ExpectKnownValue("github_actions_organization_workflow_permissions.test", tfjsonpath.New("default_workflow_permissions"), knownvalue.StringExact("read")),
						statecheck.ExpectKnownValue("github_actions_organization_workflow_permissions.test", tfjsonpath.New("can_approve_pull_request_reviews"), knownvalue.Bool(false)),
					},
				},
			},
		})
	})
}
