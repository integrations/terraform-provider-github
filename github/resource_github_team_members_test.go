package github

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/compare"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubTeamMembers(t *testing.T) {
	t.Parallel()

	skipUnlessHasOrgs(t)
	skipUnlessHasOrgUser1(t)
	flippedCaseUsername := flipUsernameCase(testAccConf.testOrgUser1)

	t.Run("team_by_slug", func(t *testing.T) {
		t.Parallel()

		team := mustCreateTestTeam(t, nil)

		config := fmt.Sprintf(`
resource "github_team_members" "test" {
  team_slug = "%s"

  members {
    username = "%s"
    role     = "maintainer"
  }
}
`, team.GetSlug(), flippedCaseUsername)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_team_members.test", tfjsonpath.New("team_id"), knownvalue.StringExact(strconv.FormatInt(team.GetID(), 10))),
						statecheck.ExpectKnownValue("github_team_members.test", tfjsonpath.New("members"), knownvalue.SetSizeExact(1)),
					},
				},
				{
					ResourceName: "github_team_members.test",
					ImportState:  true,
					ImportStateIdFunc: func(s *terraform.State) (string, error) {
						return s.RootModule().Resources["github_team_members.test"].Primary.Attributes["team_slug"], nil
					},
					ImportStateVerify: true,
				},
			},
		})
	})

	t.Run("updates_team_member_role", func(t *testing.T) {
		t.Parallel()

		team := mustCreateTestTeam(t, nil)

		config := fmt.Sprintf(`
resource "github_team_members" "test" {
  team_slug = "%s"

  members {
    username = "%s"
    role     = "%%s"
  }
}
`, team.GetSlug(), flippedCaseUsername)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, "maintainer"),
				},
				{
					Config: fmt.Sprintf(config, "member"),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_team_members.test", plancheck.ResourceActionUpdate),
							plancheck.ExpectKnownValue("github_team_members.test", tfjsonpath.New("members").AtSliceIndex(0).AtMapKey("role"), knownvalue.StringExact("member")),
						},
					},
				},
			},
		})
	})

	t.Run("team_by_id_as_slug", func(t *testing.T) {
		t.Parallel()

		team := mustCreateTestTeam(t, nil)

		config := fmt.Sprintf(`
resource "github_team_members" "test" {
  team_id = "%s"

  members {
    username = "%s"
    role     = "maintainer"
  }
}
`, team.GetSlug(), flippedCaseUsername)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_team_members.test", tfjsonpath.New("team_slug"), knownvalue.StringExact(team.GetSlug())),
						statecheck.ExpectKnownValue("github_team_members.test", tfjsonpath.New("members"), knownvalue.SetSizeExact(1)),
					},
				},
				{
					ResourceName: "github_team_members.test",
					ImportState:  true,
					ImportStateIdFunc: func(s *terraform.State) (string, error) {
						return s.RootModule().Resources["github_team_members.test"].Primary.Attributes["team_id"], nil
					},
					ImportStateVerify:       true,
					ImportStateVerifyIgnore: []string{"team_id"},
				},
			},
		})
	})

	t.Run("team_by_id", func(t *testing.T) {
		t.Parallel()

		team := mustCreateTestTeam(t, nil)

		config := fmt.Sprintf(`
resource "github_team_members" "test" {
  team_id = "%d"

  members {
    username = "%s"
    role     = "maintainer"
  }
}
`, team.GetID(), flippedCaseUsername)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_team_members.test", tfjsonpath.New("team_slug"), knownvalue.StringExact(team.GetSlug())),
						statecheck.ExpectKnownValue("github_team_members.test", tfjsonpath.New("members"), knownvalue.SetSizeExact(1)),
					},
				},
				{
					ResourceName: "github_team_members.test",
					ImportState:  true,
					ImportStateIdFunc: func(s *terraform.State) (string, error) {
						return s.RootModule().Resources["github_team_members.test"].Primary.Attributes["team_id"], nil
					},
					ImportStateVerify: true,
				},
			},
		})
	})

	t.Run("migrate_team_by_id_as_slug_to_slug", func(t *testing.T) {
		t.Parallel()

		team := mustCreateTestTeam(t, nil)

		config := fmt.Sprintf(`
resource "github_team_members" "test" {
  team_id = "%s"

  members {
    username = "%s"
    role     = "maintainer"
  }
}
`, team.GetSlug(), flippedCaseUsername)

		configMigrate := fmt.Sprintf(`
resource "github_team_members" "test" {
  team_slug = "%s"

  members {
    username = "%s"
    role     = "maintainer"
  }
}
`, team.GetSlug(), flippedCaseUsername)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_team_members.test", tfjsonpath.New("team_slug"), knownvalue.StringExact(team.GetSlug())),
						statecheck.ExpectKnownValue("github_team_members.test", tfjsonpath.New("members"), knownvalue.SetSizeExact(1)),
					},
				},
				{
					Config: configMigrate,
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_team_members.test", plancheck.ResourceActionUpdate),
						},
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_team_members.test", tfjsonpath.New("team_id"), knownvalue.StringExact(strconv.FormatInt(team.GetID(), 10))),
						statecheck.ExpectKnownValue("github_team_members.test", tfjsonpath.New("members"), knownvalue.SetSizeExact(1)),
					},
				},
			},
		})
	})

	t.Run("migrate_team_by_id_to_slug", func(t *testing.T) {
		t.Parallel()

		team := mustCreateTestTeam(t, nil)

		config := fmt.Sprintf(`
resource "github_team_members" "test" {
  team_id = "%d"

  members {
    username = "%s"
    role     = "maintainer"
  }
}
`, team.GetID(), flippedCaseUsername)

		configMigrate := fmt.Sprintf(`
resource "github_team_members" "test" {
  team_slug = "%s"

  members {
    username = "%s"
    role     = "maintainer"
  }
}
`, team.GetSlug(), flippedCaseUsername)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_team_members.test", tfjsonpath.New("team_slug"), knownvalue.StringExact(team.GetSlug())),
						statecheck.ExpectKnownValue("github_team_members.test", tfjsonpath.New("members"), knownvalue.SetSizeExact(1)),
					},
				},
				{
					Config: configMigrate,
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_team_members.test", plancheck.ResourceActionNoop),
						},
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_team_members.test", tfjsonpath.New("team_id"), knownvalue.StringExact(strconv.FormatInt(team.GetID(), 10))),
						statecheck.ExpectKnownValue("github_team_members.test", tfjsonpath.New("members"), knownvalue.SetSizeExact(1)),
					},
				},
			},
		})
	})

	t.Run("team_with_existing_members", func(t *testing.T) {
		t.Parallel()

		skipUnlessHasOrgUser2(t)

		team := mustCreateTestTeam(t, nil)
		mustAddTeamMember(t, team, testAccConf.testOrgUser2)

		config := fmt.Sprintf(`
resource "github_team_members" "test" {
  team_slug = "%s"

  members {
    username = "%s"
    role     = "maintainer"
  }
}
`, team.GetSlug(), flippedCaseUsername)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_team_members.test", tfjsonpath.New("members"), knownvalue.SetSizeExact(1)),
					},
				}, {
					PreConfig: func() { mustAddTeamMember(t, team, testAccConf.testOrgUser2) },
					Config:    config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_team_members.test", tfjsonpath.New("members"), knownvalue.SetSizeExact(1)),
					},
				},
			},
		})
	})

	t.Run("team_with_child_team", func(t *testing.T) {
		t.Parallel()

		skipUnlessHasOrgUser2(t)

		team := mustCreateTestTeam(t, nil)
		childTeam := mustCreateTestTeam(t, team.ID)
		mustAddTeamMember(t, childTeam, testAccConf.testOrgUser2)

		config := fmt.Sprintf(`
resource "github_team_members" "test" {
  team_slug = "%s"

  members {
    username = "%s"
    role     = "maintainer"
  }
}
`, team.GetSlug(), flippedCaseUsername)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_team_members.test", tfjsonpath.New("members"), knownvalue.SetSizeExact(1)),
					},
				},
			},
		})
	})

	t.Run("team_rename_handled", func(t *testing.T) {
		t.Parallel()

		team := mustCreateTestTeam(t, nil)
		newTeamName := fmt.Sprintf("%s-renamed", team.GetName())

		config := fmt.Sprintf(`
resource "github_team_members" "test" {
  team_slug = "%%s"

  members {
    username = "%s"
    role     = "maintainer"
  }
}
`, flippedCaseUsername)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, team.GetSlug()),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_team_members.test", tfjsonpath.New("team_id"), knownvalue.StringExact(strconv.FormatInt(team.GetID(), 10))),
						statecheck.ExpectKnownValue("github_team_members.test", tfjsonpath.New("members"), knownvalue.SetSizeExact(1)),
					},
				},
				{
					PreConfig: func() { mustRenameTestTeam(t, team, newTeamName) },
					Config:    fmt.Sprintf(config, newTeamName),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_team_members.test", tfjsonpath.New("team_id"), knownvalue.StringExact(strconv.FormatInt(team.GetID(), 10))),
						statecheck.ExpectKnownValue("github_team_members.test", tfjsonpath.New("members"), knownvalue.SetSizeExact(1)),
					},
				},
			},
		})
	})

	t.Run("team_delete_handled", func(t *testing.T) {
		t.Parallel()

		team := mustCreateTestTeam(t, nil)

		config := fmt.Sprintf(`
resource "github_team_members" "test" {
  team_slug = "%s"

  members {
    username = "%s"
    role     = "maintainer"
  }
}
`, team.GetSlug(), flippedCaseUsername)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_team_members.test", tfjsonpath.New("members"), knownvalue.SetSizeExact(1)),
					},
				},
				{
					PreConfig:          func() { mustDeleteTestTeam(t, team) },
					RefreshState:       true,
					ExpectNonEmptyPlan: true,
				},
				{
					Config:  config,
					Destroy: true,
				},
			},
		})
	})

	t.Run("team_changed_forces_recreate", func(t *testing.T) {
		t.Parallel()

		team := mustCreateTestTeam(t, nil)
		newTeam := mustCreateTestTeam(t, nil)

		config := fmt.Sprintf(`
resource "github_team_members" "test" {
  team_slug = "%%s"

  members {
    username = "%s"
    role     = "maintainer"
  }
}
`, flippedCaseUsername)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, team.GetSlug()),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_team_members.test", tfjsonpath.New("team_id"), knownvalue.StringExact(strconv.FormatInt(team.GetID(), 10))),
						statecheck.ExpectKnownValue("github_team_members.test", tfjsonpath.New("members"), knownvalue.SetSizeExact(1)),
					},
				},
				{
					Config: fmt.Sprintf(config, newTeam.GetSlug()),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_team_members.test", plancheck.ResourceActionReplace),
						},
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_team_members.test", tfjsonpath.New("team_id"), knownvalue.StringExact(strconv.FormatInt(newTeam.GetID(), 10))),
						statecheck.ExpectKnownValue("github_team_members.test", tfjsonpath.New("members"), knownvalue.SetSizeExact(1)),
					},
				},
			},
		})
	})

	t.Run("is_case_insensitive", func(t *testing.T) {
		t.Parallel()

		team := mustCreateTestTeam(t, nil)

		config := fmt.Sprintf(`
resource "github_team_members" "test" {
  team_slug = "%s"

  members {
    username = "%%s"
    role     = "maintainer"
  }
}
`, team.GetSlug())

		usernameIsSameComparer := statecheck.CompareValue(compare.ValuesSame())
		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, flippedCaseUsername),
					ConfigStateChecks: []statecheck.StateCheck{
						usernameIsSameComparer.AddStateValue("github_team_members.test", tfjsonpath.New("members").AtSliceIndex(0).AtMapKey("username")),
						statecheck.ExpectKnownValue("github_team_members.test", tfjsonpath.New("members"), knownvalue.SetSizeExact(1)),
						statecheck.ExpectKnownValue("github_team_members.test", tfjsonpath.New("members").AtSliceIndex(0).AtMapKey("username"), knownvalue.StringExact(strings.ToLower(testAccConf.testOrgUser1))),
					},
				},
				{
					Config: fmt.Sprintf(config, testAccConf.testOrgUser1),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_team_members.test", plancheck.ResourceActionNoop),
						},
					},
					ConfigStateChecks: []statecheck.StateCheck{
						usernameIsSameComparer.AddStateValue("github_team_members.test", tfjsonpath.New("members").AtSliceIndex(0).AtMapKey("username")),
						statecheck.ExpectKnownValue("github_team_members.test", tfjsonpath.New("members"), knownvalue.SetSizeExact(1)),
					},
				},
			},
		})
	})
}
