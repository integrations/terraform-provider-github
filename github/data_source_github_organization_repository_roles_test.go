package github

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccDataSourceGithubOrganizationRepositoryRoles(t *testing.T) {
	t.Parallel()

	skipUnlessEnterprise(t)

	t.Run("queries_all_roles", func(t *testing.T) {
		t.Parallel()

		role := mustCreateTestOrganizationRepositoryRole(t)

		config := `
data "github_organization_repository_roles" "test" {}
`

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_organization_repository_roles.test", tfjsonpath.New("roles"), knownvalue.SetPartial([]knownvalue.Check{
							knownvalue.MapExact(map[string]knownvalue.Check{
								"id":          knownvalue.Int32Exact(int32(role.GetID())),
								"role_id":     knownvalue.Int32Exact(int32(role.GetID())),
								"name":        knownvalue.StringExact(role.GetName()),
								"description": knownvalue.StringExact(role.GetDescription()),
								"base_role":   knownvalue.StringExact(role.GetBaseRole()),
								"permissions": knownvalue.ListSizeExact(len(role.GetPermissions())),
							}),
						})),
					},
				},
			},
		})
	})
}
