package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccDataSourceGithubOrganizationRoleTeams(t *testing.T) {
	t.Parallel()

	skipUnlessHasOrgs(t)

	t.Run("queries_teams", func(t *testing.T) {
		t.Parallel()

		roleID := int64(138)
		team1 := mustCreateTestTeam(t)
		team2 := mustCreateTestTeam(t, withNewTeamParent(team1.GetID()))
		mustAssignOrganizationRoleToTeam(t, team1, roleID)

		config := fmt.Sprintf(`
data "github_organization_role_teams" "test" {
	role_id = %v
}
`, roleID)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_organization_role_teams.test", tfjsonpath.New("teams"), knownvalue.SetPartial([]knownvalue.Check{
							knownvalue.MapExact(map[string]knownvalue.Check{
								"id":          knownvalue.Int32Exact(int32(team1.GetID())),
								"team_id":     knownvalue.Int32Exact(int32(team1.GetID())),
								"slug":        knownvalue.StringExact(team1.GetSlug()),
								"name":        knownvalue.StringExact(team1.GetName()),
								"description": knownvalue.StringExact(team1.GetDescription()),
								"type":        knownvalue.StringExact(team1.GetType()),
								"privacy":     knownvalue.StringExact(team1.GetPrivacy()),
								"permission":  knownvalue.StringExact(team1.GetPermission()),
								"assignment":  knownvalue.StringExact("direct"),
								"parent_team": knownvalue.ListSizeExact(0),
							}),
							knownvalue.MapExact(map[string]knownvalue.Check{
								"id":          knownvalue.Int32Exact(int32(team2.GetID())),
								"team_id":     knownvalue.Int32Exact(int32(team2.GetID())),
								"slug":        knownvalue.StringExact(team2.GetSlug()),
								"name":        knownvalue.StringExact(team2.GetName()),
								"description": knownvalue.StringExact(team2.GetDescription()),
								"type":        knownvalue.StringExact(team2.GetType()),
								"privacy":     knownvalue.StringExact(team2.GetPrivacy()),
								"permission":  knownvalue.StringExact(team2.GetPermission()),
								"assignment":  knownvalue.StringExact("indirect"),
								"parent_team": knownvalue.ListExact([]knownvalue.Check{
									knownvalue.MapExact(map[string]knownvalue.Check{
										"id":   knownvalue.Int32Exact(int32(team1.GetID())),
										"slug": knownvalue.StringExact(team1.GetSlug()),
									}),
								}),
							}),
						})),
					},
				},
			},
		})
	})
}
