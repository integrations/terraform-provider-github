package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
)

func TestAccGithubEnterpriseCustomPropertiesDataSource(t *testing.T) {
	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlpha)
	propertyName := fmt.Sprintf("test-%s", randomID)

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

	check := resource.ComposeTestCheckFunc(
		resource.TestCheckResourceAttr("data.github_enterprise_custom_properties.test", "enterprise_slug", testAccConf.enterpriseSlug),
		resource.TestCheckResourceAttr("data.github_enterprise_custom_properties.test", "property_name", propertyName),
		resource.TestCheckResourceAttr("data.github_enterprise_custom_properties.test", "allowed_values.#", "0"),
		resource.TestCheckResourceAttr("data.github_enterprise_custom_properties.test", "value_type", "string"),
		resource.TestCheckResourceAttr("data.github_enterprise_custom_properties.test", "required", "true"),
		resource.TestCheckResourceAttr("data.github_enterprise_custom_properties.test", "default_value", "terraform"),
		resource.TestCheckResourceAttr("data.github_enterprise_custom_properties.test", "description", "A test property"),
		resource.TestCheckResourceAttr("data.github_enterprise_custom_properties.test", "values_editable_by", "org_and_repo_actors"),
	)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { skipUnlessEnterprise(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check:  check,
			},
		},
	},
	)
}
