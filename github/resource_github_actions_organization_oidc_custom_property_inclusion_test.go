package github

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccGithubActionsOrganizationOIDCCustomPropertyInclusion(t *testing.T) {
	t.Run("creates and deletes an OIDC custom property inclusion without error", func(t *testing.T) {
		config := `
		resource "github_actions_organization_oidc_custom_property_inclusion" "test" {
			custom_property_name = "environment"
		}`

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_actions_organization_oidc_custom_property_inclusion.test",
				"custom_property_name", "environment",
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

	t.Run("imports an OIDC custom property inclusion without error", func(t *testing.T) {
		config := `
		resource "github_actions_organization_oidc_custom_property_inclusion" "test" {
			custom_property_name = "environment"
		}`

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_actions_organization_oidc_custom_property_inclusion.test",
				"custom_property_name", "environment",
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
					ResourceName:      "github_actions_organization_oidc_custom_property_inclusion.test",
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})

	t.Run("manages multiple OIDC custom property inclusions", func(t *testing.T) {
		config := `
		resource "github_actions_organization_oidc_custom_property_inclusion" "env" {
			custom_property_name = "environment"
		}

		resource "github_actions_organization_oidc_custom_property_inclusion" "team" {
			custom_property_name = "team"
		}`

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_actions_organization_oidc_custom_property_inclusion.env",
				"custom_property_name", "environment",
			),
			resource.TestCheckResourceAttr(
				"github_actions_organization_oidc_custom_property_inclusion.team",
				"custom_property_name", "team",
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
}

func TestAccGithubActionsOrganizationOIDCCustomPropertyInclusionsDataSource(t *testing.T) {
	t.Run("reads OIDC custom property inclusions without error", func(t *testing.T) {
		config := `
		resource "github_actions_organization_oidc_custom_property_inclusion" "test" {
			custom_property_name = "environment"
		}

		data "github_actions_organization_oidc_custom_property_inclusions" "test" {
			depends_on = [github_actions_organization_oidc_custom_property_inclusion.test]
		}`

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet(
				"data.github_actions_organization_oidc_custom_property_inclusions.test",
				"custom_property_names.#",
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
}
