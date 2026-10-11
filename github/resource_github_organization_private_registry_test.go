package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubOrganizationPrivateRegistry(t *testing.T) {
	t.Parallel()

	skipUnlessMode(t, organization)
	skipUnlessHasOrgs(t)

	t.Run("creates_imports_and_updates", func(t *testing.T) {
		t.Parallel()

		host := acctest.RandString(testRandomIDLength)

		configTmpl := `
resource "github_organization_private_registry" "test" {
  registry_type = "npm_registry"
  url           = "%s"
  auth_type     = "username_password"
  username      = "github-actions"
  value         = "super_secret_token_123"
  visibility    = "%s"
}
`

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(configTmpl, fmt.Sprintf("https://npm-%s.example.com", host), "private"),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_organization_private_registry.test", tfjsonpath.New("name"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_organization_private_registry.test", tfjsonpath.New("key_id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_organization_private_registry.test", tfjsonpath.New("created_at"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_organization_private_registry.test", tfjsonpath.New("updated_at"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_organization_private_registry.test", tfjsonpath.New("registry_type"), knownvalue.StringExact("npm_registry")),
						statecheck.ExpectKnownValue("github_organization_private_registry.test", tfjsonpath.New("url"), knownvalue.StringExact(fmt.Sprintf("https://npm-%s.example.com", host))),
						statecheck.ExpectKnownValue("github_organization_private_registry.test", tfjsonpath.New("username"), knownvalue.StringExact("github-actions")),
						statecheck.ExpectKnownValue("github_organization_private_registry.test", tfjsonpath.New("visibility"), knownvalue.StringExact("private")),
					},
				},
				{
					ResourceName:            "github_organization_private_registry.test",
					ImportState:             true,
					ImportStateVerify:       true,
					ImportStateVerifyIgnore: []string{"value", "key_id"},
				},
				{
					Config: fmt.Sprintf(configTmpl, fmt.Sprintf("https://npm-%s-updated.example.com", host), "all"),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_organization_private_registry.test", plancheck.ResourceActionUpdate),
						},
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_organization_private_registry.test", tfjsonpath.New("url"), knownvalue.StringExact(fmt.Sprintf("https://npm-%s-updated.example.com", host))),
						statecheck.ExpectKnownValue("github_organization_private_registry.test", tfjsonpath.New("visibility"), knownvalue.StringExact("all")),
					},
				},
			},
		})
	})

	t.Run("with_selected_repositories", func(t *testing.T) {
		t.Parallel()

		repoName := fmt.Sprintf("%s%s", testResourcePrefix, acctest.RandString(testRandomIDLength))

		config := fmt.Sprintf(`
resource "github_repository" "test" {
  name = "%s"
}

resource "github_organization_private_registry" "test" {
  registry_type           = "maven_repository"
  url                     = "https://maven-%s.example.com/releases"
  auth_type               = "token"
  value                   = "super_secret_token_123"
  visibility              = "selected"
  selected_repository_ids = [github_repository.test.repo_id]
}
`, repoName, acctest.RandString(testRandomIDLength))

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_organization_private_registry.test", tfjsonpath.New("visibility"), knownvalue.StringExact("selected")),
						statecheck.ExpectKnownValue("github_organization_private_registry.test", tfjsonpath.New("selected_repository_ids"), knownvalue.SetSizeExact(1)),
					},
				},
			},
		})
	})

	t.Run("forces_new_when_registry_type_changes", func(t *testing.T) {
		t.Parallel()

		host := acctest.RandString(testRandomIDLength)

		configTmpl := `
resource "github_organization_private_registry" "test" {
  registry_type = "%s"
  url           = "%s"
  value         = "super_secret_token_123"
  visibility    = "private"
}
`

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(configTmpl, "npm_registry", fmt.Sprintf("https://npm-%s.example.com", host)),
				},
				{
					Config: fmt.Sprintf(configTmpl, "rubygems_server", fmt.Sprintf("https://gems-%s.example.com", host)),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_organization_private_registry.test", plancheck.ResourceActionDestroyBeforeCreate),
						},
					},
				},
			},
		})
	})

	t.Run("clears_username_when_removed", func(t *testing.T) {
		t.Parallel()

		host := acctest.RandString(testRandomIDLength)

		configTmpl := `
resource "github_organization_private_registry" "test" {
  registry_type = "npm_registry"
  url           = "https://npm-%s.example.com"
  auth_type     = "username_password"
  value         = "super_secret_token_123"
  visibility    = "private"
%s}
`

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(configTmpl, host, "  username      = \"deploy\"\n"),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_organization_private_registry.test", tfjsonpath.New("username"), knownvalue.StringExact("deploy")),
					},
				},
				{
					Config: fmt.Sprintf(configTmpl, host, ""),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_organization_private_registry.test", tfjsonpath.New("username"), knownvalue.StringExact("")),
					},
				},
			},
		})
	})
}
