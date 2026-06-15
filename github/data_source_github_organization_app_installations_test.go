package github

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccGithubOrganizationAppInstallations(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		config := `data "github_organization_app_installations" "test" {}`

		resource.Test(t, resource.TestCase{
			PreCheck: func() {
				skipUnlessHasOrgs(t)
				skipUnlessHasAppInstallations(t)
			},
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttrSet("data.github_organization_app_installations.test", "installations.0.id"),
						resource.TestCheckResourceAttrSet("data.github_organization_app_installations.test", "installations.0.app_slug"),
						resource.TestCheckResourceAttrSet("data.github_organization_app_installations.test", "installations.0.app_id"),
					),
				},
			},
		})
	})
}
