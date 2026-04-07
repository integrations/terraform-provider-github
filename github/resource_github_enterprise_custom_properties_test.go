package github

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubEnterpriseCustomPropertiesValidation(t *testing.T) {
	t.Run("rejects invalid values_editable_by value", func(t *testing.T) {
		config := fmt.Sprintf(`
		resource "github_enterprise_custom_property" "test" {
			enterprise_slug    = "%s"
			property_name      = "%senterprise-prop-invalid-editable-by"
			value_type         = "string"
			values_editable_by = "invalid_value"
		}`, testAccConf.enterpriseSlug, testResourcePrefix)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessMode(t, enterprise) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      config,
					ExpectError: regexp.MustCompile("invalid_value"),
				},
			},
		})
	})

	t.Run("rejects invalid value_type", func(t *testing.T) {
		config := fmt.Sprintf(`
		resource "github_enterprise_custom_property" "test" {
			enterprise_slug = "%s"
			property_name   = "%senterprise-prop-invalid-type"
			value_type      = "invalid_type"
		}`, testAccConf.enterpriseSlug, testResourcePrefix)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessMode(t, enterprise) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      config,
					ExpectError: regexp.MustCompile("invalid_type"),
				},
			},
		})
	})
}

func TestAccGithubEnterpriseCustomProperties(t *testing.T) {
	t.Run("creates custom property without error", func(t *testing.T) {
		config := fmt.Sprintf(`
		resource "github_enterprise_custom_property" "test" {
			enterprise_slug = "%s"
			allowed_values  = ["Test"]
			description     = "Test Description"
			default_value   = "Test"
			property_name   = "%senterprise-prop-create"
			required        = true
			value_type      = "single_select"
		}`, testAccConf.enterpriseSlug, testResourcePrefix)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessMode(t, enterprise) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_enterprise_custom_property.test", tfjsonpath.New("value_type"), knownvalue.StringExact("single_select")),
						statecheck.ExpectKnownValue("github_enterprise_custom_property.test", tfjsonpath.New("required"), knownvalue.Bool(true)),
					},
				},
			},
		})
	})

	t.Run("creates and updates a custom property", func(t *testing.T) {
		configBefore := fmt.Sprintf(`
		resource "github_enterprise_custom_property" "test" {
			enterprise_slug = "%s"
			allowed_values  = ["one"]
			description     = "Test Description"
			property_name   = "%senterprise-prop-update"
			value_type      = "single_select"
		}`, testAccConf.enterpriseSlug, testResourcePrefix)

		configAfter := fmt.Sprintf(`
		resource "github_enterprise_custom_property" "test" {
			enterprise_slug = "%s"
			allowed_values  = ["one", "two"]
			description     = "Test Description Updated"
			property_name   = "%senterprise-prop-update"
			value_type      = "single_select"
		}`, testAccConf.enterpriseSlug, testResourcePrefix)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessMode(t, enterprise) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: configBefore,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_enterprise_custom_property.test", tfjsonpath.New("allowed_values"), knownvalue.ListSizeExact(1)),
					},
				},
				{
					Config: configAfter,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_enterprise_custom_property.test", tfjsonpath.New("allowed_values"), knownvalue.ListSizeExact(2)),
						statecheck.ExpectKnownValue("github_enterprise_custom_property.test", tfjsonpath.New("description"), knownvalue.StringExact("Test Description Updated")),
					},
				},
			},
		})
	})

	t.Run("imports enterprise custom property without error", func(t *testing.T) {
		config := fmt.Sprintf(`
		resource "github_enterprise_custom_property" "test" {
			enterprise_slug = "%s"
			description     = "Test Description Import"
			property_name   = "%senterprise-prop-import"
			value_type      = "string"
		}`, testAccConf.enterpriseSlug, testResourcePrefix)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessMode(t, enterprise) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_enterprise_custom_property.test", tfjsonpath.New("description"), knownvalue.StringExact("Test Description Import")),
					},
				},
				{
					ResourceName:      "github_enterprise_custom_property.test",
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})

	t.Run("creates custom property with values_editable_by", func(t *testing.T) {
		config := fmt.Sprintf(`
		resource "github_enterprise_custom_property" "test" {
			enterprise_slug    = "%s"
			property_name      = "%senterprise-prop-editable-by"
			value_type         = "string"
			description        = "Test property for values_editable_by"
			values_editable_by = "org_and_repo_actors"
		}`, testAccConf.enterpriseSlug, testResourcePrefix)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessMode(t, enterprise) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_enterprise_custom_property.test", tfjsonpath.New("values_editable_by"), knownvalue.StringExact("org_and_repo_actors")),
					},
				},
			},
		})
	})

	t.Run("defaults values_editable_by to org_actors", func(t *testing.T) {
		config := fmt.Sprintf(`
		resource "github_enterprise_custom_property" "test" {
			enterprise_slug = "%s"
			property_name   = "%senterprise-prop-default-editable-by"
			value_type      = "string"
			description     = "Test property without values_editable_by"
		}`, testAccConf.enterpriseSlug, testResourcePrefix)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessMode(t, enterprise) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_enterprise_custom_property.test", tfjsonpath.New("values_editable_by"), knownvalue.StringExact("org_actors")),
					},
				},
			},
		})
	})
}
