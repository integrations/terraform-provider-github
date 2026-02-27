package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccGithubRepositoryCustomProperties(t *testing.T) {
	t.Run("creates and reads multiple custom properties", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-custom-props-%s", testResourcePrefix, randomID)
		envPropName := fmt.Sprintf("tf-acc-env-%s", randomID)
		teamPropName := fmt.Sprintf("tf-acc-team-%s", randomID)

		config := fmt.Sprintf(`
			resource "github_organization_custom_properties" "environment" {
				allowed_values = ["production", "staging", "development"]
				description    = "Deployment environment"
				property_name  = "%s"
				value_type     = "single_select"
			}

			resource "github_organization_custom_properties" "team" {
				description   = "Team responsible"
				property_name = "%s"
				value_type    = "string"
			}

			resource "github_repository" "test" {
				name      = "%s"
				auto_init = true
			}

			resource "github_repository_custom_properties" "test" {
				repository_name = github_repository.test.name

				property {
					name  = github_organization_custom_properties.environment.property_name
					value = ["production"]
				}

				property {
					name  = github_organization_custom_properties.team.property_name
					value = ["platform-team"]
				}
			}
		`, envPropName, teamPropName, repoName)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("github_repository_custom_properties.test", "repository_name", repoName),
			resource.TestCheckResourceAttr("github_repository_custom_properties.test", "property.#", "2"),
			resource.TestCheckTypeSetElemNestedAttrs("github_repository_custom_properties.test", "property.*", map[string]string{
				"name": envPropName,
			}),
			resource.TestCheckTypeSetElemNestedAttrs("github_repository_custom_properties.test", "property.*", map[string]string{
				"name": teamPropName,
			}),
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

	t.Run("updates property value in place", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-custom-props-%s", testResourcePrefix, randomID)
		propName := fmt.Sprintf("tf-acc-env-%s", randomID)

		configCreate := fmt.Sprintf(`
			resource "github_organization_custom_properties" "test" {
				allowed_values = ["production", "staging", "development"]
				description    = "Deployment environment"
				property_name  = "%s"
				value_type     = "single_select"
			}

			resource "github_repository" "test" {
				name      = "%s"
				auto_init = true
			}

			resource "github_repository_custom_properties" "test" {
				repository_name = github_repository.test.name

				property {
					name  = github_organization_custom_properties.test.property_name
					value = ["production"]
				}
			}
		`, propName, repoName)

		configUpdate := fmt.Sprintf(`
			resource "github_organization_custom_properties" "test" {
				allowed_values = ["production", "staging", "development"]
				description    = "Deployment environment"
				property_name  = "%s"
				value_type     = "single_select"
			}

			resource "github_repository" "test" {
				name      = "%s"
				auto_init = true
			}

			resource "github_repository_custom_properties" "test" {
				repository_name = github_repository.test.name

				property {
					name  = github_organization_custom_properties.test.property_name
					value = ["staging"]
				}
			}
		`, propName, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: configCreate,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_repository_custom_properties.test", "property.#", "1"),
						resource.TestCheckTypeSetElemNestedAttrs("github_repository_custom_properties.test", "property.*", map[string]string{
							"name": propName,
						}),
					),
				},
				{
					Config: configUpdate,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_repository_custom_properties.test", "property.#", "1"),
						resource.TestCheckTypeSetElemNestedAttrs("github_repository_custom_properties.test", "property.*", map[string]string{
							"name": propName,
						}),
					),
				},
			},
		})
	})

	t.Run("imports all properties for a repository", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-custom-props-%s", testResourcePrefix, randomID)
		propName := fmt.Sprintf("tf-acc-env-%s", randomID)

		config := fmt.Sprintf(`
			resource "github_organization_custom_properties" "test" {
				allowed_values = ["production", "staging"]
				description    = "Deployment environment"
				property_name  = "%s"
				value_type     = "single_select"
			}

			resource "github_repository" "test" {
				name      = "%s"
				auto_init = true
			}

			resource "github_repository_custom_properties" "test" {
				repository_name = github_repository.test.name

				property {
					name  = github_organization_custom_properties.test.property_name
					value = ["production"]
				}
			}
		`, propName, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
				},
				{
					ResourceName:      "github_repository_custom_properties.test",
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})

	t.Run("creates multi_select property", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-custom-props-%s", testResourcePrefix, randomID)
		propName := fmt.Sprintf("tf-acc-tags-%s", randomID)

		config := fmt.Sprintf(`
			resource "github_organization_custom_properties" "test" {
				allowed_values = ["go", "python", "rust", "typescript"]
				description    = "Language tags"
				property_name  = "%s"
				value_type     = "multi_select"
			}

			resource "github_repository" "test" {
				name      = "%s"
				auto_init = true
			}

			resource "github_repository_custom_properties" "test" {
				repository_name = github_repository.test.name

				property {
					name  = github_organization_custom_properties.test.property_name
					value = ["go", "rust"]
				}
			}
		`, propName, repoName)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("github_repository_custom_properties.test", "property.#", "1"),
			resource.TestCheckTypeSetElemNestedAttrs("github_repository_custom_properties.test", "property.*", map[string]string{
				"name":    propName,
				"value.#": "2",
			}),
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
