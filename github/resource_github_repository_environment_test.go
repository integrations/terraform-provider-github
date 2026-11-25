package github

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
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

	t.Run("import_without_reviewers", func(t *testing.T) {
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
					// The API omits required_reviewers when none are configured, but
					// prevent_self_review must still be populated during import.
					ResourceName:            "github_repository_environment.test",
					ImportState:             true,
					ImportStateVerify:       true,
					ImportStateVerifyIgnore: []string{"can_admins_bypass", "reviewers", "wait_timer", "deployment_branch_policy"},
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
