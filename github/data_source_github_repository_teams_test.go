package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubRepositoryTeamsDataSource(t *testing.T) {
	t.Parallel()

	skipUnlessHasOrgs(t)

	t.Run("queries_all_teams", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)
		team1 := mustCreateTestTeam(t, nil)
		mustAddRepositoryToTeam(t, team1, repo)
		team2 := mustCreateTestTeam(t, nil)
		mustAddRepositoryToTeam(t, team2, repo)

		config := fmt.Sprintf(`
data "github_repository_teams" "test" {
  name = "%v"
}
`, repo.GetName())

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_repository_teams.test", tfjsonpath.New("teams"), knownvalue.SetPartial([]knownvalue.Check{
							knownvalue.MapExact(map[string]knownvalue.Check{
								"id":            knownvalue.Int32Exact(int32(team1.GetID())),
								"node_id":       knownvalue.StringExact(team1.GetNodeID()),
								"slug":          knownvalue.StringExact(team1.GetSlug()),
								"name":          knownvalue.StringExact(team1.GetName()),
								"description":   knownvalue.StringExact(team1.GetDescription()),
								"type":          knownvalue.StringExact(team1.GetType()),
								"privacy":       knownvalue.StringExact(team1.GetPrivacy()),
								"permission":    knownvalue.StringExact("pull"),
								"access_source": knownvalue.StringExact("direct"),
							}),
							knownvalue.MapExact(map[string]knownvalue.Check{
								"id":            knownvalue.Int32Exact(int32(team2.GetID())),
								"node_id":       knownvalue.StringExact(team2.GetNodeID()),
								"slug":          knownvalue.StringExact(team2.GetSlug()),
								"name":          knownvalue.StringExact(team2.GetName()),
								"description":   knownvalue.StringExact(team2.GetDescription()),
								"type":          knownvalue.StringExact(team2.GetType()),
								"privacy":       knownvalue.StringExact(team2.GetPrivacy()),
								"permission":    knownvalue.StringExact("pull"),
								"access_source": knownvalue.StringExact("direct"),
							}),
						})),
					},
				},
			},
		})
	})
}
