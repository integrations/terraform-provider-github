package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubOrganizationRepositoryRoleDataSource(t *testing.T) {
	t.Parallel()

	skipUnlessEnterprise(t)

	t.Run("queries_role", func(t *testing.T) {
		t.Parallel()

		role := mustCreateTestOrganizationRepositoryRole(t)

		config := fmt.Sprintf(`
data "github_organization_repository_role" "test" {
  role_id = %v
}
`, role.GetID())

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_organization_repository_role.test", tfjsonpath.New("name"), knownvalue.StringExact(role.GetName())),
						statecheck.ExpectKnownValue("data.github_organization_repository_role.test", tfjsonpath.New("description"), knownvalue.StringExact(role.GetDescription())),
						statecheck.ExpectKnownValue("data.github_organization_repository_role.test", tfjsonpath.New("base_role"), knownvalue.StringExact(role.GetBaseRole())),
						statecheck.ExpectKnownValue("data.github_organization_repository_role.test", tfjsonpath.New("permissions"), knownvalue.ListSizeExact(len(role.GetPermissions()))),
					},
				},
			},
		})
	})
}
