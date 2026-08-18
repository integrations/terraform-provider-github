package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccDataSourceGithubOrganizationRoleUsers(t *testing.T) {
	t.Parallel()

	skipUnlessHasOrgs(t)
	skipUnlessHasOrgUser1(t)
	skipUnlessHasOrgUser2(t)

	t.Run("queries_users", func(t *testing.T) {
		t.Parallel()

		roleID := int64(138)
		mustAssignOrganizationRoleToUser(t, testAccConf.testOrgUser1, roleID)
		team := mustCreateTestTeam(t)
		mustAddTeamMember(t, team, testAccConf.testOrgUser2)
		mustAssignOrganizationRoleToTeam(t, team, roleID)

		config := fmt.Sprintf(`
data "github_organization_role_users" "test" {
  role_id = %v
}
`, roleID)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_organization_role_users.test", tfjsonpath.New("users"), knownvalue.SetPartial([]knownvalue.Check{
							knownvalue.MapExact(map[string]knownvalue.Check{
								"user_id":    knownvalue.NotNull(),
								"login":      knownvalue.StringExact(testAccConf.testOrgUser1),
								"assignment": knownvalue.StringExact("direct"),
							}),
							knownvalue.MapExact(map[string]knownvalue.Check{
								"user_id":    knownvalue.NotNull(),
								"login":      knownvalue.StringExact(testAccConf.testOrgUser2),
								"assignment": knownvalue.StringExact("indirect"),
							}),
						})),
					},
				},
			},
		})
	})
}
