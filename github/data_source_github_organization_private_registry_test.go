package github

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubDataSourceOrganizationPrivateRegistry(t *testing.T) {
	config := `
		resource "github_organization_private_registry" "test" {
			registry_type  = "npm_registry"
			url            = "https://npm.pkg.github.com"
			username       = "github-actions"
			value          = "super_secret_token_123"
			visibility     = "private"
		}

		data "github_organization_private_registry" "test" {
			name = github_organization_private_registry.test.name
		}
	`

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			skipUnlessMode(t, organization)
			skipUnlessHasOrgs(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: config,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("data.github_organization_private_registry.test", tfjsonpath.New("id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("data.github_organization_private_registry.test", tfjsonpath.New("registry_type"), knownvalue.StringExact("npm_registry")),
					statecheck.ExpectKnownValue("data.github_organization_private_registry.test", tfjsonpath.New("url"), knownvalue.StringExact("https://npm.pkg.github.com")),
					statecheck.ExpectKnownValue("data.github_organization_private_registry.test", tfjsonpath.New("username"), knownvalue.StringExact("github-actions")),
					statecheck.ExpectKnownValue("data.github_organization_private_registry.test", tfjsonpath.New("visibility"), knownvalue.StringExact("private")),
				},
			},
		},
	})
}
