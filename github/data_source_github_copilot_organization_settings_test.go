package github

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceGithubCopilotOrganizationSettings(t *testing.T) {
	t.Run("reads Copilot organization settings", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: `data "github_copilot_organization_settings" "test" {}`,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttrSet("data.github_copilot_organization_settings.test", "seat_management_setting"),
						resource.TestCheckResourceAttrSet("data.github_copilot_organization_settings.test", "public_code_suggestions"),
					),
				},
			},
		})
	})
}
