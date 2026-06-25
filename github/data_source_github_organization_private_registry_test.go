package github

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceGithubOrganizationPrivateRegistry(t *testing.T) {
	skipUnlessMode(t, organization)

	config := `
		resource "github_organization_private_registry" "test" {
			registry_type  = "npm_registry"
			url            = "https://npm.pkg.github.com"
			username       = "github-actions"
			secret         = "super_secret_token_123"
			visibility     = "private"
		}

		data "github_organization_private_registry" "test" {
			name = github_organization_private_registry.test.name
		}
	`

	check := resource.ComposeTestCheckFunc(
		resource.TestCheckResourceAttrSet("data.github_organization_private_registry.test", "id"),
		resource.TestCheckResourceAttr("data.github_organization_private_registry.test", "registry_type", "npm_registry"),
		resource.TestCheckResourceAttr("data.github_organization_private_registry.test", "url", "https://npm.pkg.github.com"),
		resource.TestCheckResourceAttr("data.github_organization_private_registry.test", "username", "github-actions"),
		resource.TestCheckResourceAttr("data.github_organization_private_registry.test", "visibility", "private"),
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
}
