package github

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubOrganizationAppInstallations(t *testing.T) {
	config := `data "github_organization_app_installations" "test" {}`

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { skipUnlessHasOrgs(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.github_organization_app_installations.test", "installations.0.id"),
					resource.TestCheckResourceAttrSet("data.github_organization_app_installations.test", "installations.0.slug"),
					resource.TestCheckResourceAttrSet("data.github_organization_app_installations.test", "installations.0.app_id"),
				),
			},
		},
	})
}
