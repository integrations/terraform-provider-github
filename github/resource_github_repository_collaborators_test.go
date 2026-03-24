package github

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubRepositoryCollaborators(t *testing.T) {
	t.Run("create", func(t *testing.T) {
		if len(testAccConf.testOrgUser) == 0 {
			t.Skip("No organization user provided")
		}

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)
		teamName0 := fmt.Sprintf("%s%s-0", testResourcePrefix, randomID)
		teamName1 := fmt.Sprintf("%s%s-1", testResourcePrefix, randomID)
		collaboratorUser := testAccConf.testOrgUser

		config := fmt.Sprintf(`
resource "github_repository" "test" {
	name       = "%s"
	visibility = "private"
}

resource "github_team" "test_0" {
	name = "%s"
}

resource "github_team" "test_1" {
	name = "%s"
}

resource "github_repository_collaborators" "test" {
	repository = github_repository.test.name

	user {
		username   = "%s"
		permission = "admin"
	}

	team {
		team_id    = github_team.test_0.id
		permission = "pull"
	}

	team {
		team_id = github_team.test_1.slug
	}
}
`, repoName, teamName0, teamName1, collaboratorUser)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_collaborators.test", tfjsonpath.New("user"), knownvalue.SetExact([]knownvalue.Check{
							knownvalue.MapExact(map[string]knownvalue.Check{
								"username":   knownvalue.StringExact(collaboratorUser),
								"permission": knownvalue.StringExact("admin"),
							}),
						})),
						statecheck.ExpectKnownValue("github_repository_collaborators.test", tfjsonpath.New("team"), knownvalue.SetExact([]knownvalue.Check{
							knownvalue.MapExact(map[string]knownvalue.Check{
								"team_id":    knownvalue.StringRegexp(regexp.MustCompile(`^\d+$`)),
								"permission": knownvalue.StringExact("pull"),
							}),
							knownvalue.MapExact(map[string]knownvalue.Check{
								"team_id":    knownvalue.StringExact(teamName1),
								"permission": knownvalue.StringExact("push"),
							}),
						})),
						statecheck.ExpectKnownValue("github_repository_collaborators.test", tfjsonpath.New("invitation_ids"), knownvalue.MapSizeExact(0)),
					},
				},
				{
					ResourceName:            "github_repository_collaborators.test",
					ImportState:             true,
					ImportStateId:           repoName,
					ImportStateVerify:       true,
					ImportStateVerifyIgnore: []string{"user", "team", "invitation_ids"},
				},
			},
		})
	})

	t.Run("add_external_user", func(t *testing.T) {
		if len(testAccConf.testExternalUser) == 0 {
			t.Skip("No external user provided")
		}

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
resource "github_repository" "test" {
	name       = "%s"
	visibility = "private"
}

resource "github_repository_collaborators" "test" {
	repository = github_repository.test.name

	user {
		username   = "%s"
		permission = "push"
	}
}
`, repoName, testAccConf.testExternalUser)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_collaborators.test", tfjsonpath.New("user"), knownvalue.SetSizeExact(0)),
						statecheck.ExpectKnownValue("github_repository_collaborators.test", tfjsonpath.New("team"), knownvalue.SetSizeExact(0)),
						statecheck.ExpectKnownValue("github_repository_collaborators.test", tfjsonpath.New("invitation_ids"), knownvalue.MapSizeExact(1)),
					},
				},
			},
		})
	})

	t.Run("update_teams", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)
		teamName0 := fmt.Sprintf("%s%s-0", testResourcePrefix, randomID)
		teamName1 := fmt.Sprintf("%s%s-1", testResourcePrefix, randomID)
		teamName2 := fmt.Sprintf("%s%s-2", testResourcePrefix, randomID)

		configPre := fmt.Sprintf(`
resource "github_repository" "test" {
	name       = "%s"
	visibility = "private"
}

resource "github_team" "test_0" {
	name = "%s"
}

resource "github_team" "test_1" {
	name = "%s"
}

resource "github_team" "test_2" {
	name = "%s"
}
`, repoName, teamName0, teamName1, teamName2)

		config0 := fmt.Sprintf(`
%s

resource "github_repository_collaborators" "test" {
	repository = github_repository.test.name

	team {
		team_id   = github_team.test_0.slug
		permission = "pull"
	}

	team {
		team_id   = github_team.test_1.slug
		permission = "pull"
	}
}
`, configPre)

		config1 := fmt.Sprintf(`
%s

resource "github_repository_collaborators" "test" {
	repository = github_repository.test.name

	team {
		team_id   = github_team.test_1.slug
		permission = "push"
	}

	team {
		team_id   = github_team.test_2.slug
		permission = "push"
	}
}
`, configPre)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config0,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_collaborators.test", tfjsonpath.New("team"), knownvalue.SetExact([]knownvalue.Check{
							knownvalue.MapExact(map[string]knownvalue.Check{
								"team_id":    knownvalue.StringExact(teamName0),
								"permission": knownvalue.StringExact("pull"),
							}),
							knownvalue.MapExact(map[string]knownvalue.Check{
								"team_id":    knownvalue.StringExact(teamName1),
								"permission": knownvalue.StringExact("pull"),
							}),
						})),
					},
				},
				{
					Config: config1,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_collaborators.test", tfjsonpath.New("team"), knownvalue.SetExact([]knownvalue.Check{
							knownvalue.MapExact(map[string]knownvalue.Check{
								"team_id":    knownvalue.StringExact(teamName1),
								"permission": knownvalue.StringExact("push"),
							}),
							knownvalue.MapExact(map[string]knownvalue.Check{
								"team_id":    knownvalue.StringExact(teamName2),
								"permission": knownvalue.StringExact("push"),
							}),
						})),
					},
				},
			},
		})
	})

	t.Run("remove_user", func(t *testing.T) {
		if len(testAccConf.testOrgUser) == 0 {
			t.Skip("No organization user provided")
		}

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)

		configPre := fmt.Sprintf(`
resource "github_repository" "test" {
	name       = "%s"
	visibility = "private"
}
`, repoName)

		config0 := fmt.Sprintf(`
%s

resource "github_repository_collaborators" "test" {
	repository = github_repository.test.name

	user {
		username   = "%s"
		permission = "admin"
	}
}
`, configPre, testAccConf.testOrgUser)

		config1 := fmt.Sprintf(`
%s

resource "github_repository_collaborators" "test" {
	repository = github_repository.test.name
}
`, configPre)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config0,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_collaborators.test", tfjsonpath.New("user"), knownvalue.SetSizeExact(1)),
						statecheck.ExpectKnownValue("github_repository_collaborators.test", tfjsonpath.New("team"), knownvalue.SetSizeExact(0)),
						statecheck.ExpectKnownValue("github_repository_collaborators.test", tfjsonpath.New("invitation_ids"), knownvalue.MapSizeExact(0)),
					},
				},
				{
					Config: config1,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_collaborators.test", tfjsonpath.New("user"), knownvalue.SetSizeExact(0)),
						statecheck.ExpectKnownValue("github_repository_collaborators.test", tfjsonpath.New("team"), knownvalue.SetSizeExact(0)),
						statecheck.ExpectKnownValue("github_repository_collaborators.test", tfjsonpath.New("invitation_ids"), knownvalue.MapSizeExact(0)),
					},
				},
			},
		})
	})

	t.Run("change_team_reference", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)
		teamName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)

		configPre := fmt.Sprintf(`
resource "github_repository" "test" {
	name       = "%s"
	visibility = "private"
}

resource "github_team" "test" {
	name = "%s"
}
`, repoName, teamName)

		config0 := fmt.Sprintf(`
%s

resource "github_repository_collaborators" "test" {
	repository = github_repository.test.name

	team {
		team_id    = github_team.test.id
		permission = "pull"
	}
}
`, configPre)

		config1 := fmt.Sprintf(`
%s

resource "github_repository_collaborators" "test" {
	repository = github_repository.test.name

	team {
		team_id    = github_team.test.slug
		permission = "pull"
	}
}
`, configPre)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config0,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_collaborators.test", tfjsonpath.New("user"), knownvalue.SetSizeExact(0)),
						statecheck.ExpectKnownValue("github_repository_collaborators.test", tfjsonpath.New("team"), knownvalue.SetSizeExact(1)),
						statecheck.ExpectKnownValue("github_repository_collaborators.test", tfjsonpath.New("invitation_ids"), knownvalue.MapSizeExact(0)),
					},
				},
				{
					Config: config1,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_collaborators.test", tfjsonpath.New("user"), knownvalue.SetSizeExact(0)),
						statecheck.ExpectKnownValue("github_repository_collaborators.test", tfjsonpath.New("team"), knownvalue.SetSizeExact(1)),
						statecheck.ExpectKnownValue("github_repository_collaborators.test", tfjsonpath.New("invitation_ids"), knownvalue.MapSizeExact(0)),
					},
				},
			},
		})
	})

	t.Run("remove_team", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)
		teamName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)

		configPre := fmt.Sprintf(`
resource "github_repository" "test" {
	name       = "%s"
	visibility = "private"
}

resource "github_team" "test" {
	name = "%s"
}
`, repoName, teamName)

		config0 := fmt.Sprintf(`
%s

resource "github_repository_collaborators" "test" {
	repository = github_repository.test.name

	team {
		team_id    = github_team.test.slug
		permission = "pull"
	}
}
`, configPre)

		config1 := fmt.Sprintf(`
%s

resource "github_repository_collaborators" "test" {
	repository = github_repository.test.name
}
`, configPre)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config0,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_collaborators.test", tfjsonpath.New("user"), knownvalue.SetSizeExact(0)),
						statecheck.ExpectKnownValue("github_repository_collaborators.test", tfjsonpath.New("team"), knownvalue.SetSizeExact(1)),
						statecheck.ExpectKnownValue("github_repository_collaborators.test", tfjsonpath.New("invitation_ids"), knownvalue.MapSizeExact(0)),
					},
				},
				{
					Config: config1,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_collaborators.test", tfjsonpath.New("user"), knownvalue.SetSizeExact(0)),
						statecheck.ExpectKnownValue("github_repository_collaborators.test", tfjsonpath.New("team"), knownvalue.SetSizeExact(0)),
						statecheck.ExpectKnownValue("github_repository_collaborators.test", tfjsonpath.New("invitation_ids"), knownvalue.MapSizeExact(0)),
					},
				},
			},
		})
	})

	t.Run("duplicate_teams_error", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)
		teamName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
resource "github_repository" "test" {
	name       = "%s"
	visibility = "private"
}

resource "github_team" "test" {
	name = "%s"
}

resource "github_repository_collaborators" "test" {
	repository = github_repository.test.name

	team {
		team_id    = github_team.test.slug
		permission = "pull"
	}

	team {
		team_id    = github_team.test.slug
		permission = "push"
	}
}
`, repoName, teamName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      config,
					ExpectError: regexp.MustCompile(`duplicate team .+ found`),
				},
			},
		})
	})
}
