package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubActionsOrganizationPermissions(t *testing.T) {
	// IMPORTANT: Do not run these tests in parallel as they modify the organization state.

	t.Run("full_lifecycle", func(t *testing.T) {
		repo := mustCreateTestRepository(t)

		configMinimal := `
resource "github_actions_organization_permissions" "test" {
  allowed_actions      = "all"
  enabled_repositories = "all"
}
`

		configFull := fmt.Sprintf(`
resource "github_actions_organization_permissions" "test" {
  allowed_actions      = "selected"
  enabled_repositories = "selected"

  allowed_actions_config {
    github_owned_allowed = true
    patterns_allowed     = ["actions/cache@*", "actions/checkout@*"]
    verified_allowed     = true
  }

  enabled_repositories_config {
    repository_ids = [%d]
  }

  sha_pinning_required = true
}
`, repo.GetID())

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: configMinimal,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_actions_organization_permissions.test", tfjsonpath.New("allowed_actions_config"), knownvalue.ListSizeExact(0)),
						statecheck.ExpectKnownValue("github_actions_organization_permissions.test", tfjsonpath.New("enabled_repositories_config"), knownvalue.ListSizeExact(0)),
						statecheck.ExpectKnownValue("github_actions_organization_permissions.test", tfjsonpath.New("sha_pinning_required"), knownvalue.NotNull()),
					},
				},
				{
					Config: configFull,
				},
				{
					ResourceName:      "github_actions_organization_permissions.test",
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})
}
