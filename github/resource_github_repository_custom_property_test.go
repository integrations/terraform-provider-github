package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccGithubRepositoryCustomProperty(t *testing.T) {
	t.Run("creates custom property of type single_select without error", func(t *testing.T) {
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

	t.Run("creates custom property of type multi_select without error", func(t *testing.T) {
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

	t.Run("creates custom property of type true-false without error", func(t *testing.T) {
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

	t.Run("updates custom property value in place without replacement", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-custom-prop-%s", testResourcePrefix, randomID)
		propertyName := fmt.Sprintf("tf-acc-test-property-%s", randomID)

		configTemplate := `
			resource "github_organization_custom_properties" "test" {
				allowed_values = ["alpha", "beta"]
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
				property_value = ["%s"]
			}
		`

		configAlpha := fmt.Sprintf(configTemplate, propertyName, repoName, "alpha")
		configBeta := fmt.Sprintf(configTemplate, propertyName, repoName, "beta")

		var firstID string

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: configAlpha,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_repository_custom_property.test", "property_value.#", "1"),
						resource.TestCheckResourceAttr("github_repository_custom_property.test", "property_value.0", "alpha"),
						func(s *terraform.State) error {
							rs, ok := s.RootModule().Resources["github_repository_custom_property.test"]
							if !ok {
								return fmt.Errorf("resource not found in state")
							}
							firstID = rs.Primary.ID
							return nil
						},
					),
				},
				{
					Config: configBeta,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_repository_custom_property.test", "property_value.#", "1"),
						resource.TestCheckResourceAttr("github_repository_custom_property.test", "property_value.0", "beta"),
						func(s *terraform.State) error {
							rs, ok := s.RootModule().Resources["github_repository_custom_property.test"]
							if !ok {
								return fmt.Errorf("resource not found in state")
							}
							if rs.Primary.ID != firstID {
								return fmt.Errorf("resource ID changed across update: %q -> %q (expected in-place update)", firstID, rs.Primary.ID)
							}
							return nil
						},
					),
				},
				{
					Config:   configBeta,
					PlanOnly: true,
				},
			},
		})
	})

	t.Run("creates custom property of type string without error", func(t *testing.T) {
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
}
