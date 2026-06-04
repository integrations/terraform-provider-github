package github

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccGithubRepositoryCustomProperty(t *testing.T) {
	t.Run("creates_single_select", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-custom-prop-%s", testResourcePrefix, randomID)
		propertyName := fmt.Sprintf("tf-acc-test-property-%s", randomID)

		config := fmt.Sprintf(`
			resource "github_organization_custom_properties" "test" {
				allowed_values = ["option1", "option2"]
				description    = "Test Description"
				property_name  = "%s"
				value_type     = "single_select"
			}
			resource "github_repository" "test" {
				name = "%s"
				auto_init = true
			}
			resource "github_repository_custom_property" "test" {
				repository    = github_repository.test.name
				property_name = github_organization_custom_properties.test.property_name
				property_type = github_organization_custom_properties.test.value_type
				property_value = ["option1"]
			}
		`, propertyName, repoName)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("github_repository_custom_property.test", "property_name", propertyName),
			resource.TestCheckResourceAttr("github_repository_custom_property.test", "property_value.#", "1"),
			resource.TestCheckResourceAttr("github_repository_custom_property.test", "property_value.0", "option1"),
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

	t.Run("creates_multi_select", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-custom-prop-%s", testResourcePrefix, randomID)
		propertyName := fmt.Sprintf("tf-acc-test-property-%s", randomID)

		config := fmt.Sprintf(`
			resource "github_organization_custom_properties" "test" {
				allowed_values = ["option1", "option2"]
				description    = "Test Description"
				property_name  = "%s"
				value_type     = "multi_select"
			}
			resource "github_repository" "test" {
				name = "%s"
				auto_init = true
			}
			resource "github_repository_custom_property" "test" {
				repository    = github_repository.test.name
				property_name = github_organization_custom_properties.test.property_name
				property_type = github_organization_custom_properties.test.value_type
				property_value = ["option1", "option2"]
			}
		`, propertyName, repoName)

		checkWithOwner := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("github_repository_custom_property.test", "property_name", propertyName),
			resource.TestCheckResourceAttr("github_repository_custom_property.test", "property_value.#", "2"),
			resource.TestCheckResourceAttr("github_repository_custom_property.test", "property_value.0", "option1"),
			resource.TestCheckResourceAttr("github_repository_custom_property.test", "property_value.1", "option2"),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  checkWithOwner,
				},
			},
		})
	})

	t.Run("creates_true_false", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-custom-prop-%s", testResourcePrefix, randomID)
		propertyName := fmt.Sprintf("tf-acc-test-property-%s", randomID)

		config := fmt.Sprintf(`
			resource "github_organization_custom_properties" "test" {
				description    = "Test Description"
				property_name  = "%s"
				value_type     = "true_false"
			}
			resource "github_repository" "test" {
				name = "%s"
				auto_init = true
			}
			resource "github_repository_custom_property" "test" {
				repository    = github_repository.test.name
				property_name = github_organization_custom_properties.test.property_name
				property_type = github_organization_custom_properties.test.value_type
				property_value = ["true"]
			}
		`, propertyName, repoName)

		checkWithOwner := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("github_repository_custom_property.test", "property_name", propertyName),
			resource.TestCheckResourceAttr("github_repository_custom_property.test", "property_value.#", "1"),
			resource.TestCheckResourceAttr("github_repository_custom_property.test", "property_value.0", "true"),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  checkWithOwner,
				},
			},
		})
	})

	t.Run("creates_url", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)
		propertyName := fmt.Sprintf("tf-acc-test-property-%s", randomID)

		config := fmt.Sprintf(`
			resource "github_organization_custom_properties" "test" {
				description    = "Test Description"
				property_name  = "%s"
				value_type     = "url"
			}
			resource "github_repository" "test" {
				name = "%s"
				auto_init = true
			}
			resource "github_repository_custom_property" "test" {
				repository    = github_repository.test.name
				property_name = github_organization_custom_properties.test.property_name
				property_type = github_organization_custom_properties.test.value_type
				property_value = ["https://example.com"]
			}
		`, propertyName, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_repository_custom_property.test", "property_name", propertyName),
						resource.TestCheckResourceAttr("github_repository_custom_property.test", "property_value.#", "1"),
						resource.TestCheckResourceAttr("github_repository_custom_property.test", "property_value.0", "https://example.com"),
					),
				},
			},
		})
	})

	t.Run("creates_string", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-custom-prop-%s", testResourcePrefix, randomID)
		propertyName := fmt.Sprintf("tf-acc-test-property-%s", randomID)

		config := fmt.Sprintf(`
			resource "github_organization_custom_properties" "test" {
				description    = "Test Description"
				property_name  = "%s"
				value_type     = "string"
			}
			resource "github_repository" "test" {
				name = "%s"
				auto_init = true
			}
			resource "github_repository_custom_property" "test" {
				repository    = github_repository.test.name
				property_name = github_organization_custom_properties.test.property_name
				property_type = github_organization_custom_properties.test.value_type
				property_value = ["text"]
			}
		`, propertyName, repoName)

		checkWithOwner := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("github_repository_custom_property.test", "property_name", propertyName),
			resource.TestCheckResourceAttr("github_repository_custom_property.test", "property_value.#", "1"),
			resource.TestCheckResourceAttr("github_repository_custom_property.test", "property_value.0", "text"),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  checkWithOwner,
				},
			},
		})
	})

	t.Run("import", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)
		propertyName := fmt.Sprintf("tf-acc-test-property-%s", randomID)

		config := fmt.Sprintf(`
			resource "github_organization_custom_properties" "test" {
				description    = "Test Description"
				property_name  = "%s"
				value_type     = "string"
			}
			resource "github_repository" "test" {
				name = "%s"
				auto_init = true
			}
			resource "github_repository_custom_property" "test" {
				repository    = github_repository.test.name
				property_name = github_organization_custom_properties.test.property_name
				property_type = github_organization_custom_properties.test.value_type
				property_value = ["text"]
			}
		`, propertyName, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
				},
				{
					ResourceName:      "github_repository_custom_property.test",
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})

	t.Run("fails_when_property_value_contains_empty_string", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)
		propertyName := fmt.Sprintf("tf-acc-test-property-%s", randomID)

		config := fmt.Sprintf(`
			resource "github_organization_custom_properties" "test" {
				description    = "Test Description"
				property_name  = "%s"
				value_type     = "string"
			}
			resource "github_repository" "test" {
				name = "%s"
				auto_init = true
			}
			resource "github_repository_custom_property" "test" {
				repository    = github_repository.test.name
				property_name = github_organization_custom_properties.test.property_name
				property_type = github_organization_custom_properties.test.value_type
				property_value = [""]
			}
		`, propertyName, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      config,
					ExpectError: regexp.MustCompile(`to not be an empty string`),
				},
			},
		})
	})
}
