package github

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccGithubOrganizationCustomProperties(t *testing.T) {
	t.Parallel()

	t.Run("rejects invalid values_editable_by value", func(t *testing.T) {
		t.Parallel()

		name := fmt.Sprintf("%s%s", testResourcePrefix, acctest.RandString(5))

		config := fmt.Sprintf(`
resource "github_organization_custom_properties" "test" {
  property_name      = "%s"
  value_type         = "string"
  required           = false
  description        = "Test invalid values_editable_by"
  values_editable_by = "invalid_value"
}
`, name)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      config,
					ExpectError: regexp.MustCompile("invalid_value"),
				},
			},
		})
	})

	t.Run("creates custom property without error", func(t *testing.T) {
		t.Parallel()

		name := fmt.Sprintf("%s%s", testResourcePrefix, acctest.RandString(5))

		config := fmt.Sprintf(`
resource "github_organization_custom_properties" "test" {
  allowed_values = [ "Test" ]
  description    = "Test Description"
  default_value  = "Test"
  property_name  = "%s"
  required       = true
  value_type     = "single_select"
}
`, name)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_organization_custom_properties.test", "property_name", name),
					),
				},
			},
		})
	})

	t.Run("create custom property and update them", func(t *testing.T) {
		t.Parallel()

		name := fmt.Sprintf("%s%s", testResourcePrefix, acctest.RandString(5))

		configBefore := fmt.Sprintf(`
resource "github_organization_custom_properties" "test" {
  allowed_values = ["one"]
  description    = "Test Description"
  property_name  = "%s"
  value_type     = "single_select"
}
`, name)

		configAfter := fmt.Sprintf(`
resource "github_organization_custom_properties" "test" {
  allowed_values = ["one", "two"]
  description    = "Test Description 2"
  property_name  = "%s"
  value_type     = "single_select"
}
	`, name)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: configBefore,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_organization_custom_properties.test", "allowed_values.#", "1"),
					),
				},
				{
					Config: configAfter,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_organization_custom_properties.test", "allowed_values.#", "2"),
					),
				},
			},
		})
	})

	t.Run("imports organization custom property without error", func(t *testing.T) {
		t.Parallel()

		description := "Test Description Import"
		propertyName := fmt.Sprintf("%s%s", testResourcePrefix, acctest.RandString(5))
		valueType := "string"

		config := fmt.Sprintf(`
		resource "github_organization_custom_properties" "test" {
			description = "%s"
			property_name = "%s"
			value_type = "%s"
			}`, description, propertyName, valueType)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  resource.TestCheckResourceAttr("github_organization_custom_properties.test", "description", description),
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
		t.Parallel()

		name := fmt.Sprintf("%s%s", testResourcePrefix, acctest.RandString(5))

		config := fmt.Sprintf(`
		resource "github_organization_custom_properties" "test" {
			property_name       = "%s"
			value_type          = "string"
			required            = false
			description         = "Test property for values_editable_by"
			values_editable_by  = "org_and_repo_actors"
		}`, name)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_organization_custom_properties.test",
				"property_name", name,
			),
			resource.TestCheckResourceAttr(
				"github_organization_custom_properties.test",
				"values_editable_by", "org_and_repo_actors",
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

	t.Run("backward compatibility - property without values_editable_by defaults correctly", func(t *testing.T) {
		t.Parallel()

		name := fmt.Sprintf("%s%s", testResourcePrefix, acctest.RandString(5))

		config := fmt.Sprintf(`
		resource "github_organization_custom_properties" "test" {
			property_name = "%s"
			value_type    = "string"
			required      = false
			description   = "Test property without values_editable_by"
		}`, name)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_organization_custom_properties.test",
				"property_name", name,
			),
			// When not specified, API returns "org_actors" as the default
			resource.TestCheckResourceAttr(
				"github_organization_custom_properties.test",
				"values_editable_by", "org_actors",
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

	t.Run("update values_editable_by from org_actors to org_and_repo_actors", func(t *testing.T) {
		t.Parallel()

		name := fmt.Sprintf("%s%s", testResourcePrefix, acctest.RandString(5))

		configBefore := fmt.Sprintf(`
		resource "github_organization_custom_properties" "test" {
			property_name      = "%s"
			value_type         = "string"
			required           = false
			description        = "Test updating values_editable_by"
			values_editable_by = "org_actors"
		}`, name)

		configAfter := fmt.Sprintf(`
		resource "github_organization_custom_properties" "test" {
			property_name      = "%s"
			value_type         = "string"
			required           = false
			description        = "Test updating values_editable_by"
			values_editable_by = "org_and_repo_actors"
		}`, name)

		const resourceName = "github_organization_custom_properties.test"

		checkBefore := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceName, "values_editable_by", "org_actors"),
		)
		checkAfter := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceName, "values_editable_by", "org_and_repo_actors"),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
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

	t.Run("true_false property with default_value produces no drift", func(t *testing.T) {
		t.Parallel()

		name := fmt.Sprintf("%s%s", testResourcePrefix, acctest.RandString(5))

		config := fmt.Sprintf(`
resource "github_organization_custom_properties" "test" {
  property_name = "%s"
  value_type    = "true_false"
  required      = false
  description   = "Test true_false default_value"
  default_value = "true"
}
`, name)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_organization_custom_properties.test", "default_value", "true"),
					),
				},
				{
					// Read must round-trip the true_false default_value so the
					// second plan is empty (regression guard for perpetual drift).
					Config:             config,
					PlanOnly:           true,
					ExpectNonEmptyPlan: false,
				},
			},
		})
	})

	t.Run("imports existing property with values_editable_by set via UI", func(t *testing.T) {
		t.Parallel()

		// This test simulates a scenario where values_editable_by was set to
		// org_and_repo_actors in the GitHub UI before Terraform support was added.
		// The resource config intentionally omits values_editable_by to verify
		// Terraform can read and maintain the existing value from the API.

		name := fmt.Sprintf("%s%s", testResourcePrefix, acctest.RandString(5))

		configWithoutField := fmt.Sprintf(`
		resource "github_organization_custom_properties" "test" {
			property_name = "%s"
			value_type    = "string"
			required      = false
			description   = "Test property set via UI"
		}`, name)

		// After import, we explicitly set the value in config to match what's in the API
		configWithField := fmt.Sprintf(`
		resource "github_organization_custom_properties" "test" {
			property_name      = "%s"
			value_type         = "string"
			required           = false
			description        = "Test property set via UI"
			values_editable_by = "org_and_repo_actors"
		}`, name)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					// First, create a property with values_editable_by set
					Config: configWithField,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_organization_custom_properties.test", "values_editable_by", "org_and_repo_actors"),
					),
				},
				{
					// Simulate the scenario: config doesn't have values_editable_by
					// (as it would have been before Terraform support was added)
					// Terraform should read the existing value from the API
					Config: configWithoutField,
					Check: resource.ComposeTestCheckFunc(
						// Terraform should still see the value from the API
						resource.TestCheckResourceAttr("github_organization_custom_properties.test", "values_editable_by", "org_and_repo_actors"),
					),
				},
				{
					// Now add it back to the config - should be no changes needed
					Config: configWithField,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_organization_custom_properties.test", "values_editable_by", "org_and_repo_actors"),
					),
				},
			},
		})
	})
}
