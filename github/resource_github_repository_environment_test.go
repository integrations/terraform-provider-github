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

func TestAccGithubRepositoryEnvironment(t *testing.T) {
	t.Run("create", func(t *testing.T) {
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

	t.Run("import", func(t *testing.T) {
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
		if len(testAccConf.testOrgUser) == 0 {
			t.Skip("skipping test that requires GH_TEST_ORG_USER env var to be set")
		}

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
locals {
	team_count = 6
}

data "github_user" "org" {
	username = "%s"
}

resource "github_team" "test" {
	count = local.team_count

	name        = "%[1]s-${count.index}"
	description = "test"
	privacy     = "closed"
}

resource "github_repository" "test" {
	name      = "%[1]s"
	visibility = "public"
}

resource "github_team_repository" "test" {
	count = local.team_count

	team_id    = github_team.test[count.index].id
	repository = github_repository.test.name
	permission = "pull"
}

resource "github_repository_collaborator" "test_repo_collaborator" {
	repository = github_repository.test.name
	username   = data.github_user.org.login
	permission = "push"
}

resource "github_repository_environment" "test" {
	repository  = github_repository.test.name
	environment = "test"

	reviewers {
		teams = github_team_repository.test[*].team_id
		users = [data.github_user.org.id]
	}
}
`, testAccConf.testOrgUser, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
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
