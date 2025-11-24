package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubOrganizationCustomProperties(t *testing.T) {
	t.Run("creates custom property without error", func(t *testing.T) {
		config := `
		resource "github_organization_custom_properties" "test" {
			allowed_values = [ "Test" ]
			description    = "Test Description"
			default_value  = "Test"
			property_name  = "Test"
			required       = true
			value_type     = "single_select"
		  }`

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_organization_custom_properties.test",
				"property_name", "Test",
			),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:  func() { skipUnlessHasOrgs(t) },
			Providers: testAccProviders,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  check,
				},
			},
		})
	})

	t.Run("create custom property and update them", func(t *testing.T) {
		configBefore := `
		resource "github_organization_custom_properties" "test" {
			allowed_values = ["one"]
			description    = "Test Description"
			property_name  = "Test"
			value_type     = "single_select"
		}`

		configAfter := `
		resource "github_organization_custom_properties" "test" {
			allowed_values = ["one", "two"]
			description    = "Test Description 2"
			property_name  = "Test"
			value_type     = "single_select"
		}`

		const resourceName = "github_organization_custom_properties.test"

		checkBefore := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceName, "allowed_values.#", "1"),
		)
		checkAfter := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceName, "allowed_values.#", "2"),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:  func() { skipUnlessHasOrgs(t) },
			Providers: testAccProviders,
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

	t.Run("imports organization custom property without error", func(t *testing.T) {
		description := "Test Description Import"
		propertyName := "Test"
		valueType := "string"

		config := fmt.Sprintf(`
		resource "github_organization_custom_properties" "test" {
			description = "%s"
			property_name = "%s"
			value_type = "%s"
			}`, description, propertyName, valueType)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_organization_custom_properties.test",
				"description", description,
			),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:  func() { skipUnlessHasOrgs(t) },
			Providers: testAccProviders,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  check,
				},
				{
					ResourceName:      "github_organization_custom_properties.test",
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})
}
