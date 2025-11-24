package github

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubOrganizationCustomPropertiesValidation(t *testing.T) {
	t.Run("rejects invalid values_editable_by value", func(t *testing.T) {
		config := `
		resource "github_organization_custom_properties" "test" {
			property_name      = "TestInvalidValuesEditableBy"
			value_type         = "string"
			required           = false
			description        = "Test invalid values_editable_by"
			values_editable_by = "invalid_value"
		}`

		resource.Test(t, resource.TestCase{
			PreCheck:  func() { skipUnlessHasOrgs(t) },
			Providers: testAccProviders,
			Steps: []resource.TestStep{
				{
					Config:      config,
					ExpectError: regexp.MustCompile("invalid_value is an invalid value"),
				},
			},
		})
	})
}

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

	t.Run("creates custom property with values_editable_by without error", func(t *testing.T) {
		config := `
		resource "github_organization_custom_properties" "test" {
			property_name       = "TestValuesEditableBy"
			value_type          = "string"
			required            = false
			description         = "Test property for values_editable_by"
			values_editable_by  = "org_and_repo_actors"
		}`

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_organization_custom_properties.test",
				"property_name", "TestValuesEditableBy",
			),
			resource.TestCheckResourceAttr(
				"github_organization_custom_properties.test",
				"values_editable_by", "org_and_repo_actors",
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

	t.Run("backward compatibility - property without values_editable_by defaults correctly", func(t *testing.T) {
		config := `
		resource "github_organization_custom_properties" "test" {
			property_name = "TestBackwardCompat"
			value_type    = "string"
			required      = false
			description   = "Test property without values_editable_by"
		}`

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_organization_custom_properties.test",
				"property_name", "TestBackwardCompat",
			),
			// When not specified, API returns "org_actors" as the default
			resource.TestCheckResourceAttr(
				"github_organization_custom_properties.test",
				"values_editable_by", "org_actors",
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

	t.Run("update values_editable_by from org_actors to org_and_repo_actors", func(t *testing.T) {
		configBefore := `
		resource "github_organization_custom_properties" "test" {
			property_name      = "TestUpdateValuesEditableBy"
			value_type         = "string"
			required           = false
			description        = "Test updating values_editable_by"
			values_editable_by = "org_actors"
		}`

		configAfter := `
		resource "github_organization_custom_properties" "test" {
			property_name      = "TestUpdateValuesEditableBy"
			value_type         = "string"
			required           = false
			description        = "Test updating values_editable_by"
			values_editable_by = "org_and_repo_actors"
		}`

		const resourceName = "github_organization_custom_properties.test"

		checkBefore := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceName, "values_editable_by", "org_actors"),
		)
		checkAfter := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceName, "values_editable_by", "org_and_repo_actors"),
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

	t.Run("imports existing property with values_editable_by set via UI", func(t *testing.T) {
		// This test simulates a scenario where values_editable_by was set to
		// org_and_repo_actors in the GitHub UI before Terraform support was added.
		// The resource config intentionally omits values_editable_by to verify
		// Terraform can read and maintain the existing value from the API.

		configWithoutField := `
		resource "github_organization_custom_properties" "test" {
			property_name = "TestImportWithUISet"
			value_type    = "string"
			required      = false
			description   = "Test property set via UI"
		}`

		// After import, we explicitly set the value in config to match what's in the API
		configWithField := `
		resource "github_organization_custom_properties" "test" {
			property_name      = "TestImportWithUISet"
			value_type         = "string"
			required           = false
			description        = "Test property set via UI"
			values_editable_by = "org_and_repo_actors"
		}`

		const resourceName = "github_organization_custom_properties.test"

		resource.Test(t, resource.TestCase{
			PreCheck:  func() { skipUnlessHasOrgs(t) },
			Providers: testAccProviders,
			Steps: []resource.TestStep{
				{
					// First, create a property with values_editable_by set
					Config: configWithField,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr(resourceName, "values_editable_by", "org_and_repo_actors"),
					),
				},
				{
					// Simulate the scenario: config doesn't have values_editable_by
					// (as it would have been before Terraform support was added)
					// Terraform should read the existing value from the API
					Config: configWithoutField,
					Check: resource.ComposeTestCheckFunc(
						// Terraform should still see the value from the API
						resource.TestCheckResourceAttr(resourceName, "values_editable_by", "org_and_repo_actors"),
					),
				},
				{
					// Now add it back to the config - should be no changes needed
					Config: configWithField,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr(resourceName, "values_editable_by", "org_and_repo_actors"),
					),
				},
			},
		})
	})
}
