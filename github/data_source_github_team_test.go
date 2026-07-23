package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubTeamDataSource(t *testing.T) {
	t.Parallel()

	skipUnlessHasOrgs(t)

	t.Run("queries_root_team_by_slug_summary", func(t *testing.T) {
		t.Parallel()

		team := mustCreateTestTeam(t)

		config := fmt.Sprintf(`
data "github_team" "test" {
  slug         = "%s"
  summary_only = true
}
`, team.GetSlug())

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_team.test", tfjsonpath.New("team_id"), knownvalue.Int32Exact(int32(team.GetID()))),
						statecheck.ExpectKnownValue("data.github_team.test", tfjsonpath.New("slug"), knownvalue.StringExact(team.GetSlug())),
						statecheck.ExpectKnownValue("data.github_team.test", tfjsonpath.New("node_id"), knownvalue.StringExact(team.GetNodeID())),
						statecheck.ExpectKnownValue("data.github_team.test", tfjsonpath.New("name"), knownvalue.StringExact(team.GetName())),
						statecheck.ExpectKnownValue("data.github_team.test", tfjsonpath.New("description"), knownvalue.StringExact(team.GetDescription())),
						statecheck.ExpectKnownValue("data.github_team.test", tfjsonpath.New("type"), knownvalue.StringExact(team.GetType())),
						statecheck.ExpectKnownValue("data.github_team.test", tfjsonpath.New("privacy"), knownvalue.StringExact(team.GetPrivacy())),
						statecheck.ExpectKnownValue("data.github_team.test", tfjsonpath.New("notification_setting"), knownvalue.StringExact(team.GetNotificationSetting())),
						statecheck.ExpectKnownValue("data.github_team.test", tfjsonpath.New("permission"), knownvalue.StringExact(team.GetPermission())),
						statecheck.ExpectKnownValue("data.github_team.test", tfjsonpath.New("parent_team"), knownvalue.ListSizeExact(0)),
						statecheck.ExpectKnownValue("data.github_team.test", tfjsonpath.New("members"), knownvalue.ListSizeExact(0)),
						statecheck.ExpectKnownValue("data.github_team.test", tfjsonpath.New("repositories"), knownvalue.ListSizeExact(0)),
						statecheck.ExpectKnownValue("data.github_team.test", tfjsonpath.New("repositories_detailed"), knownvalue.ListSizeExact(0)),
					},
				},
			},
		})
	})

	t.Run("queries_root_team_by_id_summary", func(t *testing.T) {
		t.Parallel()

		team := mustCreateTestTeam(t)

		config := fmt.Sprintf(`
data "github_team" "test" {
  team_id      = "%d"
  summary_only = true
}
`, team.GetID())

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_team.test", tfjsonpath.New("team_id"), knownvalue.Int32Exact(int32(team.GetID()))),
						statecheck.ExpectKnownValue("data.github_team.test", tfjsonpath.New("slug"), knownvalue.StringExact(team.GetSlug())),
						statecheck.ExpectKnownValue("data.github_team.test", tfjsonpath.New("node_id"), knownvalue.StringExact(team.GetNodeID())),
						statecheck.ExpectKnownValue("data.github_team.test", tfjsonpath.New("name"), knownvalue.StringExact(team.GetName())),
						statecheck.ExpectKnownValue("data.github_team.test", tfjsonpath.New("description"), knownvalue.StringExact(team.GetDescription())),
						statecheck.ExpectKnownValue("data.github_team.test", tfjsonpath.New("type"), knownvalue.StringExact(team.GetType())),
						statecheck.ExpectKnownValue("data.github_team.test", tfjsonpath.New("privacy"), knownvalue.StringExact(team.GetPrivacy())),
						statecheck.ExpectKnownValue("data.github_team.test", tfjsonpath.New("notification_setting"), knownvalue.StringExact(team.GetNotificationSetting())),
						statecheck.ExpectKnownValue("data.github_team.test", tfjsonpath.New("permission"), knownvalue.StringExact(team.GetPermission())),
						statecheck.ExpectKnownValue("data.github_team.test", tfjsonpath.New("parent_team"), knownvalue.ListSizeExact(0)),
						statecheck.ExpectKnownValue("data.github_team.test", tfjsonpath.New("members"), knownvalue.ListSizeExact(0)),
						statecheck.ExpectKnownValue("data.github_team.test", tfjsonpath.New("repositories"), knownvalue.ListSizeExact(0)),
						statecheck.ExpectKnownValue("data.github_team.test", tfjsonpath.New("repositories_detailed"), knownvalue.ListSizeExact(0)),
					},
				},
			},
		})
	})

	t.Run("queries_team_details", func(t *testing.T) {
		t.Parallel()

		skipUnlessHasOrgUser1(t)
		skipUnlessHasOrgUser2(t)

		team := mustCreateTestTeam(t)
		mustAddTeamMember(t, team, testAccConf.testOrgUser1)
		childTeam := mustCreateTestTeam(t, withNewTeamParent(team.GetID()))
		mustAddTeamMember(t, childTeam, testAccConf.testOrgUser2)
		repo := mustCreateTestRepository(t)
		mustAddRepositoryToTeam(t, team, repo)

		config := fmt.Sprintf(`
data "github_team" "test" {
  slug            = "%s"
  summary_only    = false
  membership_type = "%%v"
}
`, team.GetSlug())

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, "all"),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_team.test", tfjsonpath.New("team_id"), knownvalue.Int32Exact(int32(team.GetID()))),
						statecheck.ExpectKnownValue("data.github_team.test", tfjsonpath.New("slug"), knownvalue.StringExact(team.GetSlug())),
						statecheck.ExpectKnownValue("data.github_team.test", tfjsonpath.New("node_id"), knownvalue.StringExact(team.GetNodeID())),
						statecheck.ExpectKnownValue("data.github_team.test", tfjsonpath.New("name"), knownvalue.StringExact(team.GetName())),
						statecheck.ExpectKnownValue("data.github_team.test", tfjsonpath.New("description"), knownvalue.StringExact(team.GetDescription())),
						statecheck.ExpectKnownValue("data.github_team.test", tfjsonpath.New("type"), knownvalue.StringExact(team.GetType())),
						statecheck.ExpectKnownValue("data.github_team.test", tfjsonpath.New("privacy"), knownvalue.StringExact(team.GetPrivacy())),
						statecheck.ExpectKnownValue("data.github_team.test", tfjsonpath.New("notification_setting"), knownvalue.StringExact(team.GetNotificationSetting())),
						statecheck.ExpectKnownValue("data.github_team.test", tfjsonpath.New("permission"), knownvalue.StringExact(team.GetPermission())),
						statecheck.ExpectKnownValue("data.github_team.test", tfjsonpath.New("parent_team"), knownvalue.ListSizeExact(0)),
						statecheck.ExpectKnownValue("data.github_team.test", tfjsonpath.New("members"), knownvalue.ListExact([]knownvalue.Check{
							knownvalue.StringExact(testAccConf.testOrgUser1),
							knownvalue.StringExact(testAccConf.testOrgUser2),
						})),
						statecheck.ExpectKnownValue("data.github_team.test", tfjsonpath.New("repositories"), knownvalue.ListExact([]knownvalue.Check{
							knownvalue.StringExact(repo.GetName()),
						})),
						statecheck.ExpectKnownValue("data.github_team.test", tfjsonpath.New("repositories_detailed"), knownvalue.ListExact([]knownvalue.Check{
							knownvalue.MapExact(map[string]knownvalue.Check{
								"repo_id":   knownvalue.Int32Exact(int32(repo.GetID())),
								"repo_name": knownvalue.StringExact(repo.GetName()),
								"role_name": knownvalue.StringExact("read"),
							}),
						})),
					},
				},
				{
					Config: fmt.Sprintf(config, "immediate"),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_team.test", tfjsonpath.New("team_id"), knownvalue.Int32Exact(int32(team.GetID()))),
						statecheck.ExpectKnownValue("data.github_team.test", tfjsonpath.New("slug"), knownvalue.StringExact(team.GetSlug())),
						statecheck.ExpectKnownValue("data.github_team.test", tfjsonpath.New("node_id"), knownvalue.StringExact(team.GetNodeID())),
						statecheck.ExpectKnownValue("data.github_team.test", tfjsonpath.New("name"), knownvalue.StringExact(team.GetName())),
						statecheck.ExpectKnownValue("data.github_team.test", tfjsonpath.New("description"), knownvalue.StringExact(team.GetDescription())),
						statecheck.ExpectKnownValue("data.github_team.test", tfjsonpath.New("type"), knownvalue.StringExact(team.GetType())),
						statecheck.ExpectKnownValue("data.github_team.test", tfjsonpath.New("privacy"), knownvalue.StringExact(team.GetPrivacy())),
						statecheck.ExpectKnownValue("data.github_team.test", tfjsonpath.New("notification_setting"), knownvalue.StringExact(team.GetNotificationSetting())),
						statecheck.ExpectKnownValue("data.github_team.test", tfjsonpath.New("permission"), knownvalue.StringExact(team.GetPermission())),
						statecheck.ExpectKnownValue("data.github_team.test", tfjsonpath.New("parent_team"), knownvalue.ListSizeExact(0)),
						statecheck.ExpectKnownValue("data.github_team.test", tfjsonpath.New("members"), knownvalue.ListExact([]knownvalue.Check{
							knownvalue.StringExact(testAccConf.testOrgUser1),
						})),
						statecheck.ExpectKnownValue("data.github_team.test", tfjsonpath.New("repositories"), knownvalue.ListExact([]knownvalue.Check{
							knownvalue.StringExact(repo.GetName()),
						})),
						statecheck.ExpectKnownValue("data.github_team.test", tfjsonpath.New("repositories_detailed"), knownvalue.ListExact([]knownvalue.Check{
							knownvalue.MapExact(map[string]knownvalue.Check{
								"repo_id":   knownvalue.Int32Exact(int32(repo.GetID())),
								"repo_name": knownvalue.StringExact(repo.GetName()),
								"role_name": knownvalue.StringExact("read"),
							}),
						})),
					},
				},
			},
		})
	})

	t.Run("queries_child_team", func(t *testing.T) {
		t.Parallel()

		parentTeam := mustCreateTestTeam(t)
		team := mustCreateTestTeam(t, withNewTeamParent(parentTeam.GetID()))

		config := fmt.Sprintf(`
data "github_team" "test" {
  slug         = "%s"
  summary_only = true
}
`, team.GetSlug())

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_team.test", tfjsonpath.New("team_id"), knownvalue.Int32Exact(int32(team.GetID()))),
						statecheck.ExpectKnownValue("data.github_team.test", tfjsonpath.New("slug"), knownvalue.StringExact(team.GetSlug())),
						statecheck.ExpectKnownValue("data.github_team.test", tfjsonpath.New("node_id"), knownvalue.StringExact(team.GetNodeID())),
						statecheck.ExpectKnownValue("data.github_team.test", tfjsonpath.New("name"), knownvalue.StringExact(team.GetName())),
						statecheck.ExpectKnownValue("data.github_team.test", tfjsonpath.New("description"), knownvalue.StringExact(team.GetDescription())),
						statecheck.ExpectKnownValue("data.github_team.test", tfjsonpath.New("type"), knownvalue.StringExact(team.GetType())),
						statecheck.ExpectKnownValue("data.github_team.test", tfjsonpath.New("privacy"), knownvalue.StringExact(team.GetPrivacy())),
						statecheck.ExpectKnownValue("data.github_team.test", tfjsonpath.New("notification_setting"), knownvalue.StringExact(team.GetNotificationSetting())),
						statecheck.ExpectKnownValue("data.github_team.test", tfjsonpath.New("permission"), knownvalue.StringExact(team.GetPermission())),
						statecheck.ExpectKnownValue("data.github_team.test", tfjsonpath.New("parent_team"), knownvalue.ListExact([]knownvalue.Check{
							knownvalue.MapExact(map[string]knownvalue.Check{
								"id":   knownvalue.Int32Exact(int32(parentTeam.GetID())),
								"slug": knownvalue.StringExact(parentTeam.GetSlug()),
							}),
						})),
						statecheck.ExpectKnownValue("data.github_team.test", tfjsonpath.New("members"), knownvalue.ListSizeExact(0)),
						statecheck.ExpectKnownValue("data.github_team.test", tfjsonpath.New("repositories"), knownvalue.ListSizeExact(0)),
						statecheck.ExpectKnownValue("data.github_team.test", tfjsonpath.New("repositories_detailed"), knownvalue.ListSizeExact(0)),
					},
				},
			},
		})
	})
}
