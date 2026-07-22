package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccDataSourceGithubTeamMembers(t *testing.T) {
	t.Parallel()

	skipUnlessHasOrgs(t)

	t.Run("queries_all_members_by_slug", func(t *testing.T) {
		t.Parallel()

		skipUnlessHasOrgUser1(t)
		skipUnlessHasOrgUser2(t)

		user1 := mustGetUser(t, testAccConf.testOrgUser1)
		team1 := mustCreateTestTeam(t, nil)
		mustAddTeamMaintainer(t, team1, user1.GetLogin())

		user2 := mustGetUser(t, testAccConf.testOrgUser2)
		team2 := mustCreateTestTeam(t, team1.ID)
		mustAddTeamMember(t, team2, user2.GetLogin())

		config := fmt.Sprintf(`
data "github_team_members" "test" {
  slug = "%v"
}
`, team1.GetSlug())

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_team_members.test", tfjsonpath.New("members"), knownvalue.SetPartial([]knownvalue.Check{
							knownvalue.MapExact(map[string]knownvalue.Check{
								"id":        knownvalue.Int32Exact(int32(user1.GetID())),
								"node_id":   knownvalue.StringExact(user1.GetNodeID()),
								"login":     knownvalue.StringExact(user1.GetLogin()),
								"email":     knownvalue.StringExact(user1.GetEmail()),
								"role":      knownvalue.StringExact("maintainer"),
								"inherited": knownvalue.Bool(false),
							}),
							knownvalue.MapExact(map[string]knownvalue.Check{
								"id":        knownvalue.Int32Exact(int32(user2.GetID())),
								"node_id":   knownvalue.StringExact(user2.GetNodeID()),
								"login":     knownvalue.StringExact(user2.GetLogin()),
								"email":     knownvalue.StringExact(user2.GetEmail()),
								"role":      knownvalue.StringExact("member"),
								"inherited": knownvalue.Bool(true),
							}),
						})),
					},
				},
			},
		})
	})

	t.Run("queries_all_members_by_team_id", func(t *testing.T) {
		t.Parallel()

		skipUnlessHasOrgUser1(t)
		skipUnlessHasOrgUser2(t)

		user1 := mustGetUser(t, testAccConf.testOrgUser1)
		team1 := mustCreateTestTeam(t, nil)
		mustAddTeamMaintainer(t, team1, user1.GetLogin())

		user2 := mustGetUser(t, testAccConf.testOrgUser2)
		team2 := mustCreateTestTeam(t, team1.ID)
		mustAddTeamMember(t, team2, user2.GetLogin())

		config := fmt.Sprintf(`
data "github_team_members" "test" {
  team_id = %v
}
`, team1.GetID())

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_team_members.test", tfjsonpath.New("members"), knownvalue.SetPartial([]knownvalue.Check{
							knownvalue.MapExact(map[string]knownvalue.Check{
								"id":        knownvalue.Int32Exact(int32(user1.GetID())),
								"node_id":   knownvalue.StringExact(user1.GetNodeID()),
								"login":     knownvalue.StringExact(user1.GetLogin()),
								"email":     knownvalue.StringExact(user1.GetEmail()),
								"role":      knownvalue.StringExact("maintainer"),
								"inherited": knownvalue.Bool(false),
							}),
							knownvalue.MapExact(map[string]knownvalue.Check{
								"id":        knownvalue.Int32Exact(int32(user2.GetID())),
								"node_id":   knownvalue.StringExact(user2.GetNodeID()),
								"login":     knownvalue.StringExact(user2.GetLogin()),
								"email":     knownvalue.StringExact(user2.GetEmail()),
								"role":      knownvalue.StringExact("member"),
								"inherited": knownvalue.Bool(true),
							}),
						})),
					},
				},
			},
		})
	})
}
