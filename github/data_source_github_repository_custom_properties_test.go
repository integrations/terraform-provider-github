package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubRepositoryCustomPropertiesDataSource(t *testing.T) {
	t.Run("creates custom property of type single_select without error", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		propertyName := fmt.Sprintf("tf-acc-test-property-%s", randomID)

		config := fmt.Sprintf(`
			resource "github_organization_custom_properties" "test" {
				allowed_values = ["option1", "option2"]
				description    = "Test Description"
				property_name  = "%s"
				value_type     = "single_select"
			}
			resource "github_repository" "test" {
				name = "tf-acc-test-%s"
				auto_init = true
			}
			resource "github_repository_custom_property" "test" {
				repository    = github_repository.test.name
				property_name = github_organization_custom_properties.test.property_name
				property_type = github_organization_custom_properties.test.value_type
				property_value = ["option1"]
			}
			data "github_repository_custom_properties" "test" {
				repository    = github_repository_custom_property.test.repository
			}
		`, propertyName, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckTypeSetElemNestedAttrs("data.github_repository_custom_properties.test", "property.*", map[string]string{
				"property_name":    propertyName,
				"property_value.#": "1",
				"property_value.0": "option1",
			}),
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

	t.Run("creates custom property of type multi_select without error", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		propertyName := fmt.Sprintf("tf-acc-test-property-%s", randomID)

		config := fmt.Sprintf(`
			resource "github_organization_custom_properties" "test" {
				allowed_values = ["option1", "option2"]
				description    = "Test Description"
				property_name  = "%s"
				value_type     = "multi_select"
			}
			resource "github_repository" "test" {
				name = "tf-acc-test-%s"
				auto_init = true
			}
			resource "github_repository_custom_property" "test" {
				repository    = github_repository.test.name
				property_name = github_organization_custom_properties.test.property_name
				property_type = github_organization_custom_properties.test.value_type
				property_value = ["option1", "option2"]
			}
			data "github_repository_custom_properties" "test" {
				repository    = github_repository_custom_property.test.repository
			}
		`, propertyName, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckTypeSetElemNestedAttrs("data.github_repository_custom_properties.test", "property.*", map[string]string{
				"property_name":    propertyName,
				"property_value.#": "2",
				"property_value.0": "option1",
				"property_value.1": "option2",
			}),
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

	t.Run("creates custom property of type true_false without error", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		propertyName := fmt.Sprintf("tf-acc-test-property-%s", randomID)

		config := fmt.Sprintf(`
			resource "github_organization_custom_properties" "test" {
				description    = "Test Description"
				property_name  = "%s"
				value_type     = "true_false"
			}
			resource "github_repository" "test" {
				name = "tf-acc-test-%s"
				auto_init = true
			}
			resource "github_repository_custom_property" "test" {
				repository    = github_repository.test.name
				property_name = github_organization_custom_properties.test.property_name
				property_type = github_organization_custom_properties.test.value_type
				property_value = ["true"]
			}
			data "github_repository_custom_properties" "test" {
				repository    = github_repository_custom_property.test.repository
			}
		`, propertyName, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckTypeSetElemNestedAttrs("data.github_repository_custom_properties.test", "property.*", map[string]string{
				"property_name":    propertyName,
				"property_value.#": "1",
				"property_value.0": "true",
			}),
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

	t.Run("creates custom property of type string without error", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		propertyName := fmt.Sprintf("tf-acc-test-property-%s", randomID)

		config := fmt.Sprintf(`
			resource "github_organization_custom_properties" "test" {
				description    = "Test Description"
				property_name  = "%s"
				value_type     = "string"
			}
			resource "github_repository" "test" {
				name = "tf-acc-test-%s"
				auto_init = true
			}
			resource "github_repository_custom_property" "test" {
				repository    = github_repository.test.name
				property_name = github_organization_custom_properties.test.property_name
				property_type = github_organization_custom_properties.test.value_type
				property_value = ["text"]
			}
			data "github_repository_custom_properties" "test" {
				repository    = github_repository_custom_property.test.repository
			}
		`, propertyName, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckTypeSetElemNestedAttrs("data.github_repository_custom_properties.test", "property.*", map[string]string{
				"property_name":    propertyName,
				"property_value.#": "1",
				"property_value.0": "text",
			}),
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
}
