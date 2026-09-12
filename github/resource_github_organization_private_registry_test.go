package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubOrganizationPrivateRegistry(t *testing.T) {
	configTmpl := `
		resource "github_organization_private_registry" "test" {
			registry_type  = "npm_registry"
			url            = "%s"
			username       = "github-actions"
			value          = "super_secret_token_123"
			visibility     = "%s"
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
				Config: fmt.Sprintf(configTmpl, "https://npm.pkg.github.com", "private"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("github_organization_private_registry.test", tfjsonpath.New("name"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("github_organization_private_registry.test", tfjsonpath.New("registry_type"), knownvalue.StringExact("npm_registry")),
					statecheck.ExpectKnownValue("github_organization_private_registry.test", tfjsonpath.New("url"), knownvalue.StringExact("https://npm.pkg.github.com")),
					statecheck.ExpectKnownValue("github_organization_private_registry.test", tfjsonpath.New("username"), knownvalue.StringExact("github-actions")),
					statecheck.ExpectKnownValue("github_organization_private_registry.test", tfjsonpath.New("visibility"), knownvalue.StringExact("private")),
					statecheck.ExpectKnownValue("data.github_organization_private_registry.test", tfjsonpath.New("registry_type"), knownvalue.StringExact("npm_registry")),
				},
			},
			{
				ResourceName:            "github_organization_private_registry.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"value"},
			},
			{
				Config: fmt.Sprintf(configTmpl, "https://npm-registry.example.com", "all"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("github_organization_private_registry.test", tfjsonpath.New("name"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("github_organization_private_registry.test", tfjsonpath.New("registry_type"), knownvalue.StringExact("npm_registry")),
					statecheck.ExpectKnownValue("github_organization_private_registry.test", tfjsonpath.New("url"), knownvalue.StringExact("https://npm-registry.example.com")),
					statecheck.ExpectKnownValue("github_organization_private_registry.test", tfjsonpath.New("username"), knownvalue.StringExact("github-actions")),
					statecheck.ExpectKnownValue("github_organization_private_registry.test", tfjsonpath.New("visibility"), knownvalue.StringExact("all")),
					statecheck.ExpectKnownValue("data.github_organization_private_registry.test", tfjsonpath.New("registry_type"), knownvalue.StringExact("npm_registry")),
				},
			},
		},
	})
}
