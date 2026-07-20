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

		team := mustCreateTestTeam(t, nil)
		mustAddTeamMember(t, team, testAccConf.testOrgUser1)

		config := fmt.Sprintf(`
data "github_team_members" "test" {
  slug = "%v"
}
`, team.GetSlug())

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_team_members.test", tfjsonpath.New("members"), knownvalue.SetPartial([]knownvalue.Check{
							knownvalue.MapPartial(map[string]knownvalue.Check{
								"login": knownvalue.StringExact(testAccConf.testOrgUser1),
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

		team := mustCreateTestTeam(t, nil)
		mustAddTeamMember(t, team, testAccConf.testOrgUser1)

		config := fmt.Sprintf(`
data "github_team_members" "test" {
  team_id = %v
}
`, team.GetID())

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_team_members.test", tfjsonpath.New("members"), knownvalue.SetPartial([]knownvalue.Check{
							knownvalue.MapPartial(map[string]knownvalue.Check{
								"login": knownvalue.StringExact(testAccConf.testOrgUser1),
							}),
						})),
					},
				},
			},
		})
	})
}
