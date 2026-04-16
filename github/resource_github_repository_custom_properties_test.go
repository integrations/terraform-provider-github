package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccGithubRepositoryCustomProperties(t *testing.T) {
	t.Run("sets custom_properties at creation time", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		propertyName := fmt.Sprintf("tfacc%s", randomID)
		repoName := fmt.Sprintf("%scustomprops-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
			resource "github_organization_custom_properties" "test" {
				allowed_values = ["alpha", "beta"]
				property_name  = "%[1]s"
				value_type     = "single_select"
			}

			resource "github_repository" "test" {
				name      = "%[2]s"
				auto_init = true

				custom_properties = {
					(github_organization_custom_properties.test.property_name) = "alpha"
				}
			}
		`, propertyName, repoName)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("github_repository.test", "custom_properties.%", "1"),
			resource.TestCheckResourceAttr(
				"github_repository.test",
				fmt.Sprintf("custom_properties.%s", propertyName),
				"alpha",
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

	t.Run("updates a custom property value", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		propertyName := fmt.Sprintf("tfacc%s", randomID)
		repoName := fmt.Sprintf("%scustomprops-%s", testResourcePrefix, randomID)

		configTemplate := `
			resource "github_organization_custom_properties" "test" {
				allowed_values = ["alpha", "beta"]
				property_name  = "%[1]s"
				value_type     = "single_select"
			}

			resource "github_repository" "test" {
				name      = "%[2]s"
				auto_init = true

				custom_properties = {
					(github_organization_custom_properties.test.property_name) = "%[3]s"
				}
			}
		`

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(configTemplate, propertyName, repoName, "alpha"),
					Check: resource.TestCheckResourceAttr(
						"github_repository.test",
						fmt.Sprintf("custom_properties.%s", propertyName),
						"alpha",
					),
				},
				{
					Config: fmt.Sprintf(configTemplate, propertyName, repoName, "beta"),
					Check: resource.TestCheckResourceAttr(
						"github_repository.test",
						fmt.Sprintf("custom_properties.%s", propertyName),
						"beta",
					),
				},
			},
		})
	})

	t.Run("removes a specific property key", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		propertyA := fmt.Sprintf("tfacca%s", randomID)
		propertyB := fmt.Sprintf("tfaccb%s", randomID)
		repoName := fmt.Sprintf("%scustomprops-%s", testResourcePrefix, randomID)

		propertiesConfig := fmt.Sprintf(`
			resource "github_organization_custom_properties" "a" {
				allowed_values = ["one"]
				property_name  = "%[1]s"
				value_type     = "single_select"
			}

			resource "github_organization_custom_properties" "b" {
				allowed_values = ["two"]
				property_name  = "%[2]s"
				value_type     = "single_select"
			}
		`, propertyA, propertyB)

		configBoth := propertiesConfig + fmt.Sprintf(`
			resource "github_repository" "test" {
				name      = "%[1]s"
				auto_init = true

				custom_properties = {
					(github_organization_custom_properties.a.property_name) = "one"
					(github_organization_custom_properties.b.property_name) = "two"
				}
			}
		`, repoName)

		configOnlyA := propertiesConfig + fmt.Sprintf(`
			resource "github_repository" "test" {
				name      = "%[1]s"
				auto_init = true

				custom_properties = {
					(github_organization_custom_properties.a.property_name) = "one"
				}
			}
		`, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: configBoth,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_repository.test", "custom_properties.%", "2"),
					),
				},
				{
					Config: configOnlyA,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_repository.test", "custom_properties.%", "1"),
						resource.TestCheckResourceAttr(
							"github_repository.test",
							fmt.Sprintf("custom_properties.%s", propertyA),
							"one",
						),
						resource.TestCheckNoResourceAttr(
							"github_repository.test",
							fmt.Sprintf("custom_properties.%s", propertyB),
						),
					),
				},
			},
		})
	})

	t.Run("clears all properties with empty map", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		propertyName := fmt.Sprintf("tfacc%s", randomID)
		repoName := fmt.Sprintf("%scustomprops-%s", testResourcePrefix, randomID)

		propertiesConfig := fmt.Sprintf(`
			resource "github_organization_custom_properties" "test" {
				allowed_values = ["alpha"]
				property_name  = "%[1]s"
				value_type     = "single_select"
			}
		`, propertyName)

		configWithValue := propertiesConfig + fmt.Sprintf(`
			resource "github_repository" "test" {
				name      = "%[1]s"
				auto_init = true

				custom_properties = {
					(github_organization_custom_properties.test.property_name) = "alpha"
				}
			}
		`, repoName)

		configEmpty := propertiesConfig + fmt.Sprintf(`
			resource "github_repository" "test" {
				name      = "%[1]s"
				auto_init = true

				custom_properties = {}
			}
		`, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: configWithValue,
					Check:  resource.TestCheckResourceAttr("github_repository.test", "custom_properties.%", "1"),
				},
				{
					Config: configEmpty,
					Check:  resource.TestCheckResourceAttr("github_repository.test", "custom_properties.%", "0"),
				},
			},
		})
	})

	t.Run("removing the block leaves values untouched", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		propertyName := fmt.Sprintf("tfacc%s", randomID)
		repoName := fmt.Sprintf("%scustomprops-%s", testResourcePrefix, randomID)

		propertiesConfig := fmt.Sprintf(`
			resource "github_organization_custom_properties" "test" {
				allowed_values = ["alpha"]
				property_name  = "%[1]s"
				value_type     = "single_select"
			}
		`, propertyName)

		configWithValue := propertiesConfig + fmt.Sprintf(`
			resource "github_repository" "test" {
				name      = "%[1]s"
				auto_init = true

				custom_properties = {
					(github_organization_custom_properties.test.property_name) = "alpha"
				}
			}
		`, repoName)

		configNoBlock := propertiesConfig + fmt.Sprintf(`
			resource "github_repository" "test" {
				name      = "%[1]s"
				auto_init = true
			}
		`, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: configWithValue,
					Check: resource.TestCheckResourceAttr(
						"github_repository.test",
						fmt.Sprintf("custom_properties.%s", propertyName),
						"alpha",
					),
				},
				{
					Config:             configNoBlock,
					PlanOnly:           true,
					ExpectNonEmptyPlan: false,
				},
			},
		})
	})
}
