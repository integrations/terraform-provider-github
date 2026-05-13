package github

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccGithubOrganizationPrivateRegistry_basic(t *testing.T) {
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
		resource.TestCheckResourceAttrSet("github_organization_private_registry.test", "name"),
		resource.TestCheckResourceAttr("github_organization_private_registry.test", "registry_type", "npm_registry"),
		resource.TestCheckResourceAttr("github_organization_private_registry.test", "url", "https://npm.pkg.github.com"),
		resource.TestCheckResourceAttr("github_organization_private_registry.test", "username", "github-actions"),
		resource.TestCheckResourceAttr("github_organization_private_registry.test", "visibility", "private"),
		resource.TestCheckResourceAttr("data.github_organization_private_registry.test", "registry_type", "npm_registry"),
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
