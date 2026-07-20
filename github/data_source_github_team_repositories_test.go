package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccDataSourceGithubTeamRepositories(t *testing.T) {
	t.Parallel()

	skipUnlessHasOrgs(t)

	t.Run("queries_all_repositories_by_slug", func(t *testing.T) {
		t.Parallel()

		team := mustCreateTestTeam(t, nil)
		repo := mustCreateTestRepository(t)
		mustAddRepositoryTeam(t, repo, team)

		config := fmt.Sprintf(`
data "github_team_repositories" "test" {
  slug = "%v"
}
`, team.GetSlug())

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_team_repositories.test", tfjsonpath.New("repositories"), knownvalue.SetPartial([]knownvalue.Check{
							knownvalue.MapExact(map[string]knownvalue.Check{
								"id":         knownvalue.Int32Exact(int32(repo.GetID())),
								"node_id":    knownvalue.StringExact(repo.GetNodeID()),
								"name":       knownvalue.StringExact(repo.GetName()),
								"visibility": knownvalue.StringExact(repo.GetVisibility()),
								"archived":   knownvalue.Bool(repo.GetArchived()),
							}),
						})),
					},
				},
			},
		})
	})

	t.Run("queries_all_repositories_by_team_id", func(t *testing.T) {
		t.Parallel()

		team := mustCreateTestTeam(t, nil)
		repo := mustCreateTestRepository(t)
		mustAddRepositoryTeam(t, repo, team)

		config := fmt.Sprintf(`
data "github_team_repositories" "test" {
  team_id = %v
}
`, team.GetID())

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_team_repositories.test", tfjsonpath.New("repositories"), knownvalue.SetPartial([]knownvalue.Check{
							knownvalue.MapExact(map[string]knownvalue.Check{
								"id":         knownvalue.Int32Exact(int32(repo.GetID())),
								"node_id":    knownvalue.StringExact(repo.GetNodeID()),
								"name":       knownvalue.StringExact(repo.GetName()),
								"visibility": knownvalue.StringExact(repo.GetVisibility()),
								"archived":   knownvalue.Bool(repo.GetArchived()),
							}),
						})),
					},
				},
			},
		})
	})
}
