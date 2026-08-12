package github

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubRepositoryCollaborators(t *testing.T) {
	t.Parallel()

	t.Run("with_teams_and_users", func(t *testing.T) {
		t.Parallel()

		skipUnlessHasOrgs(t)
		skipUnlessHasOrgUser1(t)

		team0 := mustCreateTestTeam(t)
		mustAssignOrganizationRoleToTeam(t, team0, 138)

		repo := mustCreateTestRepository(t)
		team1 := mustCreateTestTeam(t)
		mustAddRepositoryToTeam(t, team1, repo)
		_ = mustCreateTestTeam(t, withNewTeamParent(team1.GetID()))

		config := fmt.Sprintf(`
resource "github_repository_collaborators" "test" {
  repository = "%v"

  team {
    team_id    = "%v"
    permission = "push"
  }

  user {
    username = "%v"
    permission = "pull"
  }
}
`, repo.GetName(), team1.GetSlug(), testAccConf.testOrgUser1)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_collaborators.test", tfjsonpath.New("repository_id"), knownvalue.Int32Exact(int32(repo.GetID()))),
						statecheck.ExpectKnownValue("github_repository_collaborators.test", tfjsonpath.New("owner_configured"), knownvalue.Bool(false)),
						statecheck.ExpectKnownValue("github_repository_collaborators.test", tfjsonpath.New("invitation_ids"), knownvalue.MapSizeExact(0)),
					},
				},
				{
					ResourceName:      "github_repository_collaborators.test",
					ImportState:       true,
					ImportStateId:     repo.GetName(),
					ImportStateVerify: true,
				},
			},
		})
	})

	t.Run("with_only_teams", func(t *testing.T) {
		t.Parallel()

		skipUnlessHasOrgs(t)

		team0 := mustCreateTestTeam(t)
		mustAssignOrganizationRoleToTeam(t, team0, 138)

		repo := mustCreateTestRepository(t)
		team1 := mustCreateTestTeam(t)
		mustAddRepositoryToTeam(t, team1, repo)
		team2 := mustCreateTestTeam(t)
		mustAddRepositoryToTeam(t, team2, repo)
		_ = mustCreateTestTeam(t, withNewTeamParent(team1.GetID()))

		config := fmt.Sprintf(`
resource "github_repository_collaborators" "test" {
  repository = "%v"

  team {
    team_id    = "%%v"
    permission = "%%v"
  }

  team {
    team_id = "%%v"
  }
}
`, repo.GetName())

		configRemoveTeam := fmt.Sprintf(`
resource "github_repository_collaborators" "test" {
  repository = "%v"

  team {
    team_id    = "%%v"
    permission = "%%v"
  }
}
`, repo.GetName())

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, team1.GetID(), "pull", team2.GetID()),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_collaborators.test", tfjsonpath.New("repository_id"), knownvalue.Int32Exact(int32(repo.GetID()))),
						statecheck.ExpectKnownValue("github_repository_collaborators.test", tfjsonpath.New("owner_configured"), knownvalue.Bool(false)),
						statecheck.ExpectKnownValue("github_repository_collaborators.test", tfjsonpath.New("invitation_ids"), knownvalue.MapSizeExact(0)),
					},
				},
				{
					Config: fmt.Sprintf(config, team1.GetSlug(), "pull", team2.GetSlug()),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_repository_collaborators.test", plancheck.ResourceActionUpdate),
						},
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_collaborators.test", tfjsonpath.New("repository_id"), knownvalue.Int32Exact(int32(repo.GetID()))),
						statecheck.ExpectKnownValue("github_repository_collaborators.test", tfjsonpath.New("owner_configured"), knownvalue.Bool(false)),
						statecheck.ExpectKnownValue("github_repository_collaborators.test", tfjsonpath.New("invitation_ids"), knownvalue.MapSizeExact(0)),
					},
				},
				{
					Config: fmt.Sprintf(config, team1.GetSlug(), "push", team2.GetSlug()),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_repository_collaborators.test", plancheck.ResourceActionUpdate),
						},
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_collaborators.test", tfjsonpath.New("repository_id"), knownvalue.Int32Exact(int32(repo.GetID()))),
						statecheck.ExpectKnownValue("github_repository_collaborators.test", tfjsonpath.New("owner_configured"), knownvalue.Bool(false)),
						statecheck.ExpectKnownValue("github_repository_collaborators.test", tfjsonpath.New("invitation_ids"), knownvalue.MapSizeExact(0)),
					},
				},
				{
					ResourceName:      "github_repository_collaborators.test",
					ImportState:       true,
					ImportStateId:     repo.GetName(),
					ImportStateVerify: true,
				},
				{
					Config: fmt.Sprintf(configRemoveTeam, team1.GetSlug(), "push"),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_repository_collaborators.test", plancheck.ResourceActionUpdate),
						},
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_collaborators.test", tfjsonpath.New("repository_id"), knownvalue.Int32Exact(int32(repo.GetID()))),
						statecheck.ExpectKnownValue("github_repository_collaborators.test", tfjsonpath.New("owner_configured"), knownvalue.Bool(false)),
						statecheck.ExpectKnownValue("github_repository_collaborators.test", tfjsonpath.New("invitation_ids"), knownvalue.MapSizeExact(0)),
					},
				},
			},
		})
	})

	t.Run("with_only_users", func(t *testing.T) {
		t.Parallel()

		skipUnlessHasOrgs(t)
		skipUnlessHasOrgUser1(t)
		skipUnlessHasOrgUser2(t)

		team0 := mustCreateTestTeam(t)
		mustAssignOrganizationRoleToTeam(t, team0, 138)

		repo := mustCreateTestRepository(t)

		config := fmt.Sprintf(`
resource "github_repository_collaborators" "test" {
  repository = "%v"

  user {
    username   = "%%v"
    permission = "%%v"
  }

  user {
    username = "%%v"
  }
}
`, repo.GetName())

		configRemoveUser := fmt.Sprintf(`
resource "github_repository_collaborators" "test" {
  repository = "%v"

  user {
    username   = "%%v"
    permission = "%%v"
  }
}
`, repo.GetName())

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, testAccConf.testOrgUser1, "pull", testAccConf.testOrgUser2),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_collaborators.test", tfjsonpath.New("repository_id"), knownvalue.Int32Exact(int32(repo.GetID()))),
						statecheck.ExpectKnownValue("github_repository_collaborators.test", tfjsonpath.New("owner_configured"), knownvalue.Bool(false)),
						statecheck.ExpectKnownValue("github_repository_collaborators.test", tfjsonpath.New("invitation_ids"), knownvalue.MapSizeExact(0)),
					},
				},
				{
					Config: fmt.Sprintf(config, testAccConf.testOrgUser1, "push", testAccConf.testOrgUser2),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_repository_collaborators.test", plancheck.ResourceActionUpdate),
						},
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_collaborators.test", tfjsonpath.New("repository_id"), knownvalue.Int32Exact(int32(repo.GetID()))),
						statecheck.ExpectKnownValue("github_repository_collaborators.test", tfjsonpath.New("owner_configured"), knownvalue.Bool(false)),
						statecheck.ExpectKnownValue("github_repository_collaborators.test", tfjsonpath.New("invitation_ids"), knownvalue.MapSizeExact(0)),
					},
				},
				{
					ResourceName:      "github_repository_collaborators.test",
					ImportState:       true,
					ImportStateId:     repo.GetName(),
					ImportStateVerify: true,
				},
				{
					Config: fmt.Sprintf(configRemoveUser, testAccConf.testOrgUser1, "push"),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_repository_collaborators.test", plancheck.ResourceActionUpdate),
						},
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_collaborators.test", tfjsonpath.New("repository_id"), knownvalue.Int32Exact(int32(repo.GetID()))),
						statecheck.ExpectKnownValue("github_repository_collaborators.test", tfjsonpath.New("owner_configured"), knownvalue.Bool(false)),
						statecheck.ExpectKnownValue("github_repository_collaborators.test", tfjsonpath.New("invitation_ids"), knownvalue.MapSizeExact(0)),
					},
				},
				{
					SkipFunc: func() (bool, error) {
						if len(testAccConf.testExternalUser1) == 0 {
							return true, nil
						}
						return false, nil
					},
					Config: fmt.Sprintf(config, testAccConf.testOrgUser1, "push", testAccConf.testExternalUser1),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_repository_collaborators.test", plancheck.ResourceActionUpdate),
						},
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_collaborators.test", tfjsonpath.New("repository_id"), knownvalue.Int32Exact(int32(repo.GetID()))),
						statecheck.ExpectKnownValue("github_repository_collaborators.test", tfjsonpath.New("owner_configured"), knownvalue.Bool(false)),
						statecheck.ExpectKnownValue("github_repository_collaborators.test", tfjsonpath.New("invitation_ids"), knownvalue.MapSizeExact(1)),
					},
				},
			},
		})
	})

	t.Run("with_user_repo", func(t *testing.T) {
		t.Parallel()

		skipUnlessMode(t, individual)

		repo := mustCreateTestRepository(t)

		configNoUser := fmt.Sprintf(`
resource "github_repository_collaborators" "test" {
  repository = "%v"
}
`, repo.GetName())

		configWithUser := fmt.Sprintf(`
resource "github_repository_collaborators" "test" {
  repository = "%v"

  user {
    username   = "%v"
    permission = "admin"
  }
}
`, repo.GetName(), testAccConf.owner)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: configNoUser,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_collaborators.test", tfjsonpath.New("repository_id"), knownvalue.Int32Exact(int32(repo.GetID()))),
						statecheck.ExpectKnownValue("github_repository_collaborators.test", tfjsonpath.New("owner_configured"), knownvalue.Bool(false)),
						statecheck.ExpectKnownValue("github_repository_collaborators.test", tfjsonpath.New("invitation_ids"), knownvalue.MapSizeExact(0)),
					},
				},
				{
					ResourceName:      "github_repository_collaborators.test",
					ImportState:       true,
					ImportStateId:     repo.GetName(),
					ImportStateVerify: true,
				},
				{
					Config: configWithUser,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_collaborators.test", tfjsonpath.New("repository_id"), knownvalue.Int32Exact(int32(repo.GetID()))),
						statecheck.ExpectKnownValue("github_repository_collaborators.test", tfjsonpath.New("owner_configured"), knownvalue.Bool(true)),
						statecheck.ExpectKnownValue("github_repository_collaborators.test", tfjsonpath.New("invitation_ids"), knownvalue.MapSizeExact(0)),
					},
				},
			},
		})
	})

	t.Run("errors_with_duplicate_teams", func(t *testing.T) {
		t.Parallel()

		skipUnlessHasOrgs(t)

		repo := mustCreateTestRepository(t)
		team1 := mustCreateTestTeam(t)

		config := fmt.Sprintf(`
resource "github_repository_collaborators" "test" {
  repository = "%v"

  team {
    team_id    = "%v"
    permission = "pull"
  }

  team {
    team_id    = "%v"
    permission = "push"
  }
}
`, repo.GetName(), team1.GetSlug(), team1.GetSlug())

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      config,
					ExpectError: regexp.MustCompile(`duplicate team .+ found`),
				},
			},
		})
	})

	t.Run("errors_with_duplicate_users", func(t *testing.T) {
		t.Parallel()

		skipUnlessHasOrgs(t)
		skipUnlessHasOrgUser1(t)

		repo := mustCreateTestRepository(t)

		config := fmt.Sprintf(`
resource "github_repository_collaborators" "test" {
  repository = "%v"

  user {
    username   = "%v"
    permission = "pull"
  }

  user {
    username   = "%v"
    permission = "push"
  }
}
`, repo.GetName(), testAccConf.testOrgUser1, testAccConf.testOrgUser1)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      config,
					ExpectError: regexp.MustCompile(`duplicate user .+ found`),
				},
			},
		})
	})
}
