package github

import (
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubOrganizationTeamsDataSource(t *testing.T) {
	t.Parallel()

	skipUnlessHasOrgs(t)

	t.Run("queries_all_teams_summary", func(t *testing.T) {
		t.Parallel()

		team1 := mustCreateTestTeam(t, nil)
		team2 := mustCreateTestTeam(t, new(team1.GetID()))

		config := `
data "github_organization_teams" "test" {
  root_teams_only = false
  summary_only    = true
}
`

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_organization_teams.test", tfjsonpath.New("teams"), knownvalue.SetPartial([]knownvalue.Check{
							knownvalue.MapExact(map[string]knownvalue.Check{
								"id":                   knownvalue.Int32Exact(int32(team1.GetID())),
								"node_id":              knownvalue.StringExact(team1.GetNodeID()),
								"slug":                 knownvalue.StringExact(team1.GetSlug()),
								"name":                 knownvalue.StringExact(team1.GetName()),
								"description":          knownvalue.StringExact(team1.GetDescription()),
								"type":                 knownvalue.StringExact(team1.GetType()),
								"privacy":              knownvalue.StringExact(team1.GetPrivacy()),
								"notification_setting": knownvalue.StringExact(team1.GetNotificationSetting()),
								"permission":           knownvalue.StringExact(team1.GetPermission()),
								"parent_team":          knownvalue.ListSizeExact(0),
								"members":              knownvalue.ListSizeExact(0),
								"repositories":         knownvalue.ListSizeExact(0),
								"parent":               knownvalue.MapExact(map[string]knownvalue.Check{}),
								"parent_team_id":       knownvalue.StringExact(""),
								"parent_team_slug":     knownvalue.StringExact(""),
							}),
							knownvalue.MapExact(map[string]knownvalue.Check{
								"id":                   knownvalue.Int32Exact(int32(team2.GetID())),
								"node_id":              knownvalue.StringExact(team2.GetNodeID()),
								"slug":                 knownvalue.StringExact(team2.GetSlug()),
								"name":                 knownvalue.StringExact(team2.GetName()),
								"description":          knownvalue.StringExact(team2.GetDescription()),
								"type":                 knownvalue.StringExact(team2.GetType()),
								"privacy":              knownvalue.StringExact(team2.GetPrivacy()),
								"notification_setting": knownvalue.StringExact(team2.GetNotificationSetting()),
								"permission":           knownvalue.StringExact(team2.GetPermission()),
								"parent_team": knownvalue.ListExact([]knownvalue.Check{
									knownvalue.MapExact(map[string]knownvalue.Check{
										"id":   knownvalue.Int32Exact(int32(team1.GetID())),
										"slug": knownvalue.StringExact(team1.GetSlug()),
									}),
								}),
								"members":      knownvalue.ListSizeExact(0),
								"repositories": knownvalue.ListSizeExact(0),
								"parent": knownvalue.MapExact(map[string]knownvalue.Check{
									"id":   knownvalue.StringExact(strconv.FormatInt(team1.GetID(), 10)),
									"slug": knownvalue.StringExact(team1.GetSlug()),
									"name": knownvalue.StringExact(team1.GetName()),
								}),
								"parent_team_id":   knownvalue.StringExact(strconv.FormatInt(team1.GetID(), 10)),
								"parent_team_slug": knownvalue.StringExact(team1.GetSlug()),
							}),
						})),
					},
				},
			},
		})
	})

	t.Run("queries_root_teams_summary", func(t *testing.T) {
		t.Parallel()

		team1 := mustCreateTestTeam(t, nil)
		team2 := mustCreateTestTeam(t, new(team1.GetID()))

		config := `
data "github_organization_teams" "test" {
  root_teams_only = true
  summary_only    = true
}
`

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_organization_teams.test", tfjsonpath.New("teams"), knownvalue.SetPartial([]knownvalue.Check{
							knownvalue.MapExact(map[string]knownvalue.Check{
								"id":                   knownvalue.Int32Exact(int32(team1.GetID())),
								"node_id":              knownvalue.StringExact(team1.GetNodeID()),
								"slug":                 knownvalue.StringExact(team1.GetSlug()),
								"name":                 knownvalue.StringExact(team1.GetName()),
								"description":          knownvalue.StringExact(team1.GetDescription()),
								"type":                 knownvalue.StringExact(team1.GetType()),
								"privacy":              knownvalue.StringExact(team1.GetPrivacy()),
								"notification_setting": knownvalue.StringExact(team1.GetNotificationSetting()),
								"permission":           knownvalue.StringExact(team1.GetPermission()),
								"parent_team":          knownvalue.ListSizeExact(0),
								"members":              knownvalue.ListSizeExact(0),
								"repositories":         knownvalue.ListSizeExact(0),
								"parent":               knownvalue.MapExact(map[string]knownvalue.Check{}),
								"parent_team_id":       knownvalue.StringExact(""),
								"parent_team_slug":     knownvalue.StringExact(""),
							}),
						})),
					},
					Check: checkCollectionItemAbsent("data.github_organization_teams.test", "teams", "slug", team2.GetSlug()),
				},
			},
		})
	})

	t.Run("queries_all_teams_details", func(t *testing.T) {
		t.Parallel()

		skipUnlessHasOrgUser1(t)

		repo := mustCreateTestRepository(t)
		team1 := mustCreateTestTeam(t, nil)
		mustAddTeamMember(t, team1, testAccConf.testOrgUser1)
		mustAddRepositoryTeam(t, repo, team1)
		team2 := mustCreateTestTeam(t, nil)

		config := `
data "github_organization_teams" "test" {
  root_teams_only = false
  summary_only    = false
}
`

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_organization_teams.test", tfjsonpath.New("teams"), knownvalue.SetPartial([]knownvalue.Check{
							knownvalue.MapExact(map[string]knownvalue.Check{
								"id":                   knownvalue.Int32Exact(int32(team1.GetID())),
								"node_id":              knownvalue.StringExact(team1.GetNodeID()),
								"slug":                 knownvalue.StringExact(team1.GetSlug()),
								"name":                 knownvalue.StringExact(team1.GetName()),
								"description":          knownvalue.StringExact(team1.GetDescription()),
								"type":                 knownvalue.StringExact(team1.GetType()),
								"privacy":              knownvalue.StringExact(team1.GetPrivacy()),
								"notification_setting": knownvalue.StringExact(team1.GetNotificationSetting()),
								"permission":           knownvalue.StringExact(team1.GetPermission()),
								"parent_team":          knownvalue.ListSizeExact(0),
								"members": knownvalue.ListExact([]knownvalue.Check{
									knownvalue.StringExact(testAccConf.testOrgUser1),
								}),
								"repositories": knownvalue.ListExact([]knownvalue.Check{
									knownvalue.StringExact(repo.GetName()),
								}),
								"parent":           knownvalue.MapExact(map[string]knownvalue.Check{}),
								"parent_team_id":   knownvalue.StringExact(""),
								"parent_team_slug": knownvalue.StringExact(""),
							}),
							knownvalue.MapExact(map[string]knownvalue.Check{
								"id":                   knownvalue.Int32Exact(int32(team2.GetID())),
								"node_id":              knownvalue.StringExact(team2.GetNodeID()),
								"slug":                 knownvalue.StringExact(team2.GetSlug()),
								"name":                 knownvalue.StringExact(team2.GetName()),
								"description":          knownvalue.StringExact(team2.GetDescription()),
								"type":                 knownvalue.StringExact(team2.GetType()),
								"privacy":              knownvalue.StringExact(team2.GetPrivacy()),
								"notification_setting": knownvalue.StringExact(team2.GetNotificationSetting()),
								"permission":           knownvalue.StringExact(team2.GetPermission()),
								"parent_team":          knownvalue.ListSizeExact(0),
								"members":              knownvalue.ListSizeExact(0),
								"repositories":         knownvalue.ListSizeExact(0),
								"parent":               knownvalue.MapExact(map[string]knownvalue.Check{}),
								"parent_team_id":       knownvalue.StringExact(""),
								"parent_team_slug":     knownvalue.StringExact(""),
							}),
						})),
					},
				},
			},
		})
	})
}
