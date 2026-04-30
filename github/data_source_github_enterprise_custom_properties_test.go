package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
)

func TestAccGithubEnterpriseCustomPropertiesDataSource(t *testing.T) {
	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlpha)
	propertyName := fmt.Sprintf("%s-test-%s", testResourcePrefix, randomID)

	config := fmt.Sprintf(`
			resource "github_enterprise_custom_properties" "test" {
				enterprise_slug = "%s"
				property_name = "%s"
				value_type = "string"
				required = true
				default_value = "terraform"
				description = "A test property"
				values_editable_by = "org_and_repo_actors"
			}

			data "github_enterprise_custom_properties" "test" {
				enterprise_slug = "%s"
				property_name = github_enterprise_custom_properties.test.property_name
			}
		`,
		testAccConf.enterpriseSlug,
		propertyName,
		testAccConf.enterpriseSlug,
	)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { skipUnlessEnterprise(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: config,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("data.github_enterprise_custom_properties.test", tfjsonpath.New("enterprise_slug"), knownvalue.StringExact(testAccConf.enterpriseSlug)),
					statecheck.ExpectKnownValue("data.github_enterprise_custom_properties.test", tfjsonpath.New("property_name"), knownvalue.StringExact(propertyName)),
					statecheck.ExpectKnownValue("data.github_enterprise_custom_properties.test", tfjsonpath.New("value_type"), knownvalue.StringExact("string")),
					statecheck.ExpectKnownValue("data.github_enterprise_custom_properties.test", tfjsonpath.New("required"), knownvalue.BoolExact(true)),
					statecheck.ExpectKnownValue("data.github_enterprise_custom_properties.test", tfjsonpath.New("default_value"), knownvalue.StringExact("terraform")),
					statecheck.ExpectKnownValue("data.github_enterprise_custom_properties.test", tfjsonpath.New("description"), knownvalue.StringExact("A test property")),
					statecheck.ExpectKnownValue("data.github_enterprise_custom_properties.test", tfjsonpath.New("values_editable_by"), knownvalue.StringExact("org_and_repo_actors")),
				}
			},
		},
	},
	)
}
