package github

import (
	"context"
	"fmt"
	"sort"
	"testing"

	"github.com/google/go-github/v21/github"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/kylelemons/godebug/pretty"
)

func TestAccGithubBranchProtection_basic(t *testing.T) {
	var protection github.Protection

	rString := acctest.RandString(5)
	repoName := fmt.Sprintf("tf-acc-test-branch-prot-%s", rString)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccGithubBranchProtectionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubBranchProtectionConfig(repoName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubProtectedBranchExists("github_branch_protection.master", repoName+":master", &protection),
					testAccCheckGithubBranchProtectionRequiredStatusChecks(&protection, true, []string{"github/foo"}),
					testAccCheckGithubBranchProtectionRestrictions(&protection, []string{testUser}, []string{}),
					testAccCheckGithubBranchProtectionPullRequestReviews(&protection, true, []string{testUser}, []string{}, true),
					resource.TestCheckResourceAttr("github_branch_protection.master", "repository", repoName),
					resource.TestCheckResourceAttr("github_branch_protection.master", "branch", "master"),
					resource.TestCheckResourceAttr("github_branch_protection.master", "enforce_admins", "true"),
					resource.TestCheckResourceAttr("github_branch_protection.master", "required_status_checks.0.strict", "true"),
					resource.TestCheckResourceAttr("github_branch_protection.master", "required_status_checks.0.contexts.#", "1"),
					resource.TestCheckResourceAttr("github_branch_protection.master", "required_pull_request_reviews.0.dismiss_stale_reviews", "true"),
					resource.TestCheckResourceAttr("github_branch_protection.master", "required_pull_request_reviews.0.dismissal_users.#", "1"),
					resource.TestCheckResourceAttr("github_branch_protection.master", "required_pull_request_reviews.0.dismissal_teams.#", "0"),
					resource.TestCheckResourceAttr("github_branch_protection.master", "required_pull_request_reviews.0.require_code_owner_reviews", "true"),
					resource.TestCheckResourceAttr("github_branch_protection.master", "restrictions.0.users.#", "1"),
					resource.TestCheckResourceAttr("github_branch_protection.master", "restrictions.0.teams.#", "0"),
				),
			},
			{
				Config: testAccGithubBranchProtectionUpdateConfig(repoName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubProtectedBranchExists("github_branch_protection.master", repoName+":master", &protection),
					testAccCheckGithubBranchProtectionRequiredStatusChecks(&protection, false, []string{"github/bar"}),
					testAccCheckGithubBranchProtectionNoRestrictionsExist(&protection),
					testAccCheckGithubBranchProtectionNoPullRequestReviewsExist(&protection),
					resource.TestCheckResourceAttr("github_branch_protection.master", "repository", repoName),
					resource.TestCheckResourceAttr("github_branch_protection.master", "branch", "master"),
					resource.TestCheckResourceAttr("github_branch_protection.master", "required_status_checks.0.strict", "false"),
					resource.TestCheckResourceAttr("github_branch_protection.master", "required_status_checks.0.contexts.#", "1"),
					resource.TestCheckResourceAttr("github_branch_protection.master", "required_pull_request_reviews.#", "0"),
					resource.TestCheckResourceAttr("github_branch_protection.master", "restrictions.#", "0"),
				),
			},
		},
	})
}

func TestAccGithubBranchProtection_teams(t *testing.T) {
	var firstP, secondP, thirdP github.Protection

	rString := acctest.RandString(5)
	repoName := fmt.Sprintf("tf-acc-test-branch-prot-%s", rString)
	firstTeamName := fmt.Sprintf("team 1 %s", rString)
	firstTeamSlug := fmt.Sprintf("team-1-%s", rString)
	secondTeamName := fmt.Sprintf("team 2 %s", rString)
	secondTeamSlug := fmt.Sprintf("team-2-%s", rString)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccGithubBranchProtectionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubBranchProtectionConfigTeams(repoName, firstTeamName, secondTeamName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubProtectedBranchExists("github_branch_protection.master", repoName+":master", &firstP),
					testAccCheckGithubBranchProtectionRequiredStatusChecks(&firstP, false, []string{}),
					testAccCheckGithubBranchProtectionRestrictions(&firstP, []string{}, []string{firstTeamSlug, secondTeamSlug}),
					testAccCheckGithubBranchProtectionPullRequestReviews(&firstP, true, []string{}, []string{firstTeamSlug, secondTeamSlug}, false),
					resource.TestCheckResourceAttr("github_branch_protection.master", "repository", repoName),
					resource.TestCheckResourceAttr("github_branch_protection.master", "branch", "master"),
					resource.TestCheckResourceAttr("github_branch_protection.master", "enforce_admins", "true"),
					resource.TestCheckResourceAttr("github_branch_protection.master", "required_status_checks.0.strict", "false"),
					resource.TestCheckResourceAttr("github_branch_protection.master", "required_status_checks.0.contexts.#", "0"),
					resource.TestCheckResourceAttr("github_branch_protection.master", "required_pull_request_reviews.0.dismiss_stale_reviews", "true"),
					resource.TestCheckResourceAttr("github_branch_protection.master", "required_pull_request_reviews.0.dismissal_users.#", "0"),
					resource.TestCheckResourceAttr("github_branch_protection.master", "required_pull_request_reviews.0.dismissal_teams.#", "2"),
					resource.TestCheckResourceAttr("github_branch_protection.master", "required_pull_request_reviews.0.require_code_owner_reviews", "false"),
					resource.TestCheckResourceAttr("github_branch_protection.master", "restrictions.0.users.#", "0"),
					resource.TestCheckResourceAttr("github_branch_protection.master", "restrictions.0.teams.#", "2"),
				),
			},
			{
				Config: testAccGithubBranchProtectionConfigEmptyItems(repoName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubProtectedBranchExists("github_branch_protection.master", repoName+":master", &secondP),
					testAccCheckGithubBranchProtectionRequiredStatusChecks(&secondP, false, []string{}),
					resource.TestCheckResourceAttr("github_branch_protection.master", "repository", repoName),
					resource.TestCheckResourceAttr("github_branch_protection.master", "branch", "master"),
					resource.TestCheckResourceAttr("github_branch_protection.master", "enforce_admins", "true"),
					resource.TestCheckResourceAttr("github_branch_protection.master", "required_status_checks.#", "1"),
					resource.TestCheckResourceAttr("github_branch_protection.master", "required_pull_request_reviews.#", "1"),
					resource.TestCheckResourceAttr("github_branch_protection.master", "restrictions.#", "1"),
				),
			},
			{
				Config: testAccGithubBranchProtectionConfigTeams(repoName, firstTeamName, secondTeamName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubProtectedBranchExists("github_branch_protection.master", repoName+":master", &thirdP),
					testAccCheckGithubBranchProtectionRequiredStatusChecks(&thirdP, false, []string{}),
					testAccCheckGithubBranchProtectionRestrictions(&thirdP, []string{}, []string{firstTeamSlug, secondTeamSlug}),
					testAccCheckGithubBranchProtectionPullRequestReviews(&thirdP, true, []string{}, []string{firstTeamSlug, secondTeamSlug}, false),
					resource.TestCheckResourceAttr("github_branch_protection.master", "repository", repoName),
					resource.TestCheckResourceAttr("github_branch_protection.master", "branch", "master"),
					resource.TestCheckResourceAttr("github_branch_protection.master", "enforce_admins", "true"),
					resource.TestCheckResourceAttr("github_branch_protection.master", "required_status_checks.0.strict", "false"),
					resource.TestCheckResourceAttr("github_branch_protection.master", "required_status_checks.0.contexts.#", "0"),
					resource.TestCheckResourceAttr("github_branch_protection.master", "required_pull_request_reviews.0.dismiss_stale_reviews", "true"),
					resource.TestCheckResourceAttr("github_branch_protection.master", "required_pull_request_reviews.0.dismissal_users.#", "0"),
					resource.TestCheckResourceAttr("github_branch_protection.master", "required_pull_request_reviews.0.dismissal_teams.#", "2"),
					resource.TestCheckResourceAttr("github_branch_protection.master", "required_pull_request_reviews.0.require_code_owner_reviews", "false"),
					resource.TestCheckResourceAttr("github_branch_protection.master", "restrictions.0.users.#", "0"),
					resource.TestCheckResourceAttr("github_branch_protection.master", "restrictions.0.teams.#", "2"),
				),
			},
		},
	})
}

// See https://github.com/terraform-providers/terraform-provider-github/issues/8
func TestAccGithubBranchProtection_emptyItems(t *testing.T) {
	var protection github.Protection

	rString := acctest.RandString(5)
	repoName := fmt.Sprintf("tf-acc-test-branch-prot-%s", rString)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccGithubBranchProtectionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubBranchProtectionConfigEmptyItems(repoName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubProtectedBranchExists("github_branch_protection.master", repoName+":master", &protection),
					resource.TestCheckResourceAttr("github_branch_protection.master", "repository", repoName),
					resource.TestCheckResourceAttr("github_branch_protection.master", "branch", "master"),
					resource.TestCheckResourceAttr("github_branch_protection.master", "enforce_admins", "true"),
					resource.TestCheckResourceAttr("github_branch_protection.master", "required_status_checks.#", "1"),
					resource.TestCheckResourceAttr("github_branch_protection.master", "required_pull_request_reviews.#", "1"),
					resource.TestCheckResourceAttr("github_branch_protection.master", "restrictions.#", "1"),
				),
			},
		},
	})
}

func TestAccGithubBranchProtection_importBasic(t *testing.T) {
	rString := acctest.RandString(5)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccGithubBranchProtectionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubBranchProtectionConfig(rString),
			},
			{
				ResourceName:      "github_branch_protection.master",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckGithubProtectedBranchExists(n, id string, protection *github.Protection) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not Found: %s", n)
		}

		if rs.Primary.ID != id {
			return fmt.Errorf("Expected ID to be %v, got %v", id, rs.Primary.ID)
		}

		conn := testAccProvider.Meta().(*Organization).client
		o := testAccProvider.Meta().(*Organization).name
		r, b, err := parseTwoPartID(rs.Primary.ID)
		if err != nil {
			return err
		}

		githubProtection, _, err := conn.Repositories.GetBranchProtection(context.TODO(), o, r, b)
		if err != nil {
			return err
		}

		*protection = *githubProtection
		return nil
	}
}

func testAccCheckGithubBranchProtectionRequiredStatusChecks(protection *github.Protection, expectedStrict bool, expectedContexts []string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rsc := protection.RequiredStatusChecks
		if rsc == nil {
			return fmt.Errorf("Expected RequiredStatusChecks to be present, but was nil")
		}

		if rsc.Strict != expectedStrict {
			return fmt.Errorf("Expected RequiredStatusChecks.Strict to be %v, got %v", expectedStrict, rsc.Strict)
		}

		if diff := pretty.Compare(rsc.Contexts, expectedContexts); diff != "" {
			return fmt.Errorf("diff %q: (-got +want)\n%s", "contexts", diff)
		}

		return nil
	}
}

func testAccCheckGithubBranchProtectionRestrictions(protection *github.Protection, expectedUserLogins []string, expectedTeamNames []string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		restrictions := protection.Restrictions
		if restrictions == nil {
			return fmt.Errorf("Expected Restrictions to be present, but was nil")
		}

		userLogins := []string{}
		for _, u := range restrictions.Users {
			userLogins = append(userLogins, *u.Login)
		}
		if diff := pretty.Compare(userLogins, expectedUserLogins); diff != "" {
			return fmt.Errorf("diff %q: (-got +want)\n%s", "restrictions.users", diff)
		}

		teamLogins := []string{}
		for _, t := range restrictions.Teams {
			teamLogins = append(teamLogins, *t.Slug)
		}
		sort.Strings(teamLogins)
		sort.Strings(expectedTeamNames)
		if diff := pretty.Compare(teamLogins, expectedTeamNames); diff != "" {
			return fmt.Errorf("diff %q: (-got +want)\n%s", "restrictions.teams", diff)
		}

		return nil
	}
}

func testAccCheckGithubBranchProtectionPullRequestReviews(protection *github.Protection, expectedStale bool, expectedUsers, expectedTeams []string, expectedCodeOwners bool) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		reviews := protection.RequiredPullRequestReviews
		if reviews == nil {
			return fmt.Errorf("Expected Pull Request Reviews to be present, but was nil")
		}

		if reviews.DismissStaleReviews != expectedStale {
			return fmt.Errorf("Expected `dismiss_state_reviews` to be %t, got %t", expectedStale, reviews.DismissStaleReviews)
		}

		users := []string{}
		for _, u := range reviews.DismissalRestrictions.Users {
			users = append(users, *u.Login)
		}
		if diff := pretty.Compare(users, expectedUsers); diff != "" {
			return fmt.Errorf("diff %q: (-got +want)\n%s", "required_pull_request_reviews.dismissal_users", diff)
		}

		teams := []string{}
		for _, t := range reviews.DismissalRestrictions.Teams {
			teams = append(teams, *t.Slug)
		}
		sort.Strings(teams)
		sort.Strings(expectedTeams)
		if diff := pretty.Compare(teams, expectedTeams); diff != "" {
			return fmt.Errorf("diff %q: (-got +want)\n%s", "required_pull_request_reviews.dismissal_teams", diff)
		}

		if reviews.RequireCodeOwnerReviews != expectedCodeOwners {
			return fmt.Errorf("Expected `require_code_owner_reviews` to be %t, got %t", expectedCodeOwners, reviews.RequireCodeOwnerReviews)
		}

		return nil
	}
}

func testAccCheckGithubBranchProtectionNoRestrictionsExist(protection *github.Protection) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if protection.Restrictions != nil {
			return fmt.Errorf("Expected Restrictions to be nil, but was %v", protection.Restrictions)
		}

		return nil

	}
}

func testAccCheckGithubBranchProtectionNoPullRequestReviewsExist(protection *github.Protection) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if protection.RequiredPullRequestReviews != nil {
			return fmt.Errorf("Expected Pull Request reviews to be nil, but was %v", protection.RequiredPullRequestReviews)
		}

		return nil
	}
}

func testAccGithubBranchProtectionDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*Organization).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "github_branch_protection" {
			continue
		}

		o := testAccProvider.Meta().(*Organization).name
		r, b, err := parseTwoPartID(rs.Primary.ID)
		if err != nil {
			return err
		}

		protection, res, err := conn.Repositories.GetBranchProtection(context.TODO(), o, r, b)

		if err == nil {
			if protection != nil {
				return fmt.Errorf("Branch protection still exists")
			}
		}
		if res.StatusCode != 404 {
			return err
		}
		return nil
	}
	return nil
}

func testAccGithubBranchProtectionConfig(repoName string) string {
	return fmt.Sprintf(`
resource "github_repository" "test" {
  name        = "%s"
  description = "Terraform Acceptance Test %s"
  auto_init   = true
}

resource "github_branch_protection" "master" {
  repository = "${github_repository.test.name}"
  branch     = "master"
  enforce_admins = true

  required_status_checks {
    strict         = true
    contexts       = ["github/foo"]
  }

  required_pull_request_reviews {
    dismiss_stale_reviews = true
    dismissal_users = ["%s"]
    require_code_owner_reviews = true
  }

  restrictions {
    users = ["%s"]
  }
}
`, repoName, repoName, testUser, testUser)
}

func testAccGithubBranchProtectionUpdateConfig(repoName string) string {
	return fmt.Sprintf(`
resource "github_repository" "test" {
  name        = "%s"
  description = "Terraform Acceptance Test %s"
  auto_init   = true
}

resource "github_branch_protection" "master" {
  repository = "${github_repository.test.name}"
  branch     = "master"

  required_status_checks {
    strict         = false
    contexts       = ["github/bar"]
  }
}
`, repoName, repoName)
}

func testAccGithubBranchProtectionConfigEmptyItems(repoName string) string {
	return fmt.Sprintf(`
resource "github_repository" "test" {
  name        = "%s"
  description = "Terraform Acceptance Test %s"
  auto_init   = true
}

resource "github_branch_protection" "master" {
  repository = "${github_repository.test.name}"
  branch     = "master"
  enforce_admins = true

  required_status_checks {
  }

  required_pull_request_reviews {
  }

  restrictions {
  }
}
`, repoName, repoName)
}

func testAccGithubBranchProtectionConfigTeams(repoName, firstTeamName, secondTeamName string) string {
	return fmt.Sprintf(`
resource "github_repository" "test" {
  name        = "%s"
  description = "Terraform Acceptance Test %s"
  auto_init   = true
}

resource "github_team" "first" {
  name = "%s"
}

resource "github_team_repository" "first" {
  team_id    = "${github_team.first.id}"
  repository = "${github_repository.test.name}"
  permission = "pull"
}

resource "github_team" "second" {
  name = "%s"
}

resource "github_team_repository" "second" {
  team_id    = "${github_team.second.id}"
  repository = "${github_repository.test.name}"
  permission = "pull"
}

resource "github_branch_protection" "master" {
  depends_on = ["github_team_repository.first", "github_team_repository.second"]
  repository = "${github_repository.test.name}"
  branch     = "master"
  enforce_admins = true

  required_status_checks {
  }

  required_pull_request_reviews {
    dismiss_stale_reviews = true
    dismissal_teams = ["${github_team.first.slug}", "${github_team.second.slug}"]
  }

  restrictions {
    teams = ["${github_team.first.slug}", "${github_team.second.slug}"]
  }
}
`, repoName, repoName, firstTeamName, secondTeamName)
}
