package github

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubRepositoryEnvironment(t *testing.T) {
	t.Parallel()

	t.Run("create", func(t *testing.T) {
		t.Parallel()

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)
		envName := "test"

		config := fmt.Sprintf(`
resource "github_team" "test" {
	name        = "%[1]s"
	description = "test"
	privacy     = "closed"
}

resource "github_repository" "test" {
	name       = "%[1]s"
	visibility = "public"
}

resource "github_team_repository" "test" {
	team_id    = github_team.test.id
	repository = github_repository.test.name
	permission = "pull"
}

resource "github_repository_environment" "test" {
	repository  = github_repository.test.name
	environment = "%s"

	can_admins_bypass   = false
	wait_timer          = 10000
	prevent_self_review = true

	reviewers {
		teams = [github_team_repository.test.team_id]
	}

	deployment_branch_policy {
		protected_branches     = true
		custom_branch_policies = false
	}
}
`, repoName, envName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_environment.test", tfjsonpath.New("repository_id"), knownvalue.NotNull()),
					},
				},
			},
		})
	})

	t.Run("create_with_id_separator_in_name", func(t *testing.T) {
		t.Parallel()

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
resource "github_repository" "test" {
	name       = "%s"
	visibility = "public"
}

resource "github_repository_environment" "test" {
	repository  = github_repository.test.name
	environment = "environment:test"
}
`, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_environment.test", tfjsonpath.New("repository_id"), knownvalue.NotNull()),
					},
				},
			},
		})
	})

	t.Run("update_to_remove_reviewers", func(t *testing.T) {
		t.Parallel()

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)
		envName := "test"

		preConfig := fmt.Sprintf(`
resource "github_team" "test" {
	name        = "%[1]s"
	description = "test"
	privacy     = "closed"
}

resource "github_repository" "test" {
	name      = "%[1]s"
	visibility = "public"
}

resource "github_team_repository" "test" {
	team_id    = github_team.test.id
	repository = github_repository.test.name
	permission = "pull"
}
`, repoName)

		config := fmt.Sprintf(`
%s

resource "github_repository_environment" "test" {
	repository  = github_repository.test.name
	environment = "%s"

	reviewers {
		teams = [github_team_repository.test.team_id]
	}
}
`, preConfig, envName)

		configUpdated := fmt.Sprintf(`
%s

resource "github_repository_environment" "test" {
	repository  = github_repository.test.name
	environment = "%s"
}
`, preConfig, envName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_environment.test", tfjsonpath.New("repository_id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_repository_environment.test", tfjsonpath.New("reviewers"), knownvalue.ListSizeExact(1)),
					},
				},
				{
					Config: configUpdated,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_environment.test", tfjsonpath.New("reviewers"), knownvalue.ListSizeExact(0)),
					},
				},
			},
		})
	})

	t.Run("update_to_add_reviewers", func(t *testing.T) {
		t.Parallel()

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)
		envName := "test"

		preConfig := fmt.Sprintf(`
resource "github_team" "test" {
	name        = "%[1]s"
	description = "test"
	privacy     = "closed"
}

resource "github_repository" "test" {
	name      = "%[1]s"
	visibility = "public"
}

resource "github_team_repository" "test" {
	team_id    = github_team.test.id
	repository = github_repository.test.name
	permission = "pull"
}
`, repoName)

		config := fmt.Sprintf(`
%s

resource "github_repository_environment" "test" {
	repository  = github_repository.test.name
	environment = "%s"
}
`, preConfig, envName)

		configUpdated := fmt.Sprintf(`
%s

resource "github_repository_environment" "test" {
	repository  = github_repository.test.name
	environment = "%s"

	reviewers {
		teams = [github_team_repository.test.team_id]
	}
}
`, preConfig, envName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_environment.test", tfjsonpath.New("repository_id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_repository_environment.test", tfjsonpath.New("reviewers"), knownvalue.ListSizeExact(0)),
					},
				},
				{
					Config: configUpdated,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_environment.test", tfjsonpath.New("reviewers"), knownvalue.ListSizeExact(1)),
					},
				},
			},
		})
	})

	t.Run("detects_out_of_band_reviewer_removal", func(t *testing.T) {
		t.Parallel()
		skipUnlessHasOrgs(t)

		repo := mustCreateTestRepository(t)
		team := mustCreateTestTeam(t)
		mustAddRepositoryToTeam(t, team, repo)
		envName := "test"

		configWithReviewers := fmt.Sprintf(`
resource "github_repository_environment" "test" {
	repository  = "%s"
	environment = "%s"

	prevent_self_review = true

	reviewers {
		teams = [%d]
	}
}
`, repo.GetName(), envName, team.GetID())

		configWithoutReviewers := fmt.Sprintf(`
resource "github_repository_environment" "test" {
	repository  = "%s"
	environment = "%s"
}
`, repo.GetName(), envName)

		// Removing the last reviewer drops the "required_reviewers" protection
		// rule from the API response entirely.
		removeReviewers := func() {
			if _, _, err := testAccConf.meta.v3client.Repositories.CreateUpdateEnvironment(t.Context(), testAccConf.meta.name, repo.GetName(), envName, &github.CreateUpdateEnvironment{
				Reviewers:       []*github.EnvReviewers{},
				CanAdminsBypass: new(true),
			}); err != nil {
				t.Fatalf("failed to remove environment reviewers out-of-band: %v", err)
			}
		}

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: configWithReviewers,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_environment.test", tfjsonpath.New("reviewers"), knownvalue.ListSizeExact(1)),
						statecheck.ExpectKnownValue("github_repository_environment.test", tfjsonpath.New("prevent_self_review"), knownvalue.Bool(true)),
					},
				},
				{
					PreConfig: removeReviewers,
					Config:    configWithReviewers,
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_repository_environment.test", plancheck.ResourceActionUpdate),
						},
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_environment.test", tfjsonpath.New("reviewers"), knownvalue.ListSizeExact(1)),
						statecheck.ExpectKnownValue("github_repository_environment.test", tfjsonpath.New("prevent_self_review"), knownvalue.Bool(true)),
					},
				},
				{
					// A no-op plan against a config without reviewers is only possible
					// if refresh cleared both reviewers and prevent_self_review.
					PreConfig: removeReviewers,
					Config:    configWithoutReviewers,
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_repository_environment.test", plancheck.ResourceActionNoop),
						},
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_environment.test", tfjsonpath.New("reviewers"), knownvalue.ListSizeExact(0)),
						statecheck.ExpectKnownValue("github_repository_environment.test", tfjsonpath.New("prevent_self_review"), knownvalue.Bool(false)),
					},
				},
			},
		})
	})

	t.Run("import", func(t *testing.T) {
		t.Parallel()

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
resource "github_repository" "test" {
	name       = "%s"
	visibility = "public"
}

resource "github_repository_environment" "test" {
	repository 	= github_repository.test.name
	environment	= "test"
}
`, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_environment.test", tfjsonpath.New("repository_id"), knownvalue.NotNull()),
					},
				},
				{
					ResourceName:            "github_repository_environment.test",
					ImportState:             true,
					ImportStateVerify:       true,
					ImportStateVerifyIgnore: []string{"can_admins_bypass", "prevent_self_review", "reviewers", "wait_timer", "deployment_branch_policy"},
				},
			},
		})
	})

	t.Run("errors_with_more_than_six_reviewers", func(t *testing.T) {
		t.Parallel()

		skipUnlessHasOrgs(t)
		skipUnlessHasOrgUser1(t)

		repo := mustCreateTestRepository(t)
		user := mustGetUser(t, testAccConf.testOrgUser1)
		mustAddRepositoryCollaborator(t, repo, user.GetLogin())

		teamIDs := make([]string, 0, 6)
		for range 6 {
			team := mustCreateTestTeam(t)
			mustAddRepositoryToTeam(t, team, repo)
			teamIDs = append(teamIDs, strconv.FormatInt(team.GetID(), 10))
		}

		config := fmt.Sprintf(`
resource "github_repository_environment" "test" {
  repository  = "%s"
  environment = "test"

  reviewers {
    teams = [%s]
    users = [%d]
  }
}
`, repo.GetName(), strings.Join(teamIDs, ", "), user.GetID())

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      config,
					ExpectError: regexp.MustCompile(`reviewers can have at most 6 reviewers`),
				},
			},
		})
	})
}
