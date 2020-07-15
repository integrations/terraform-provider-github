package github

import (
	"context"
	"fmt"
	"regexp"
	"sort"
	"testing"

	"github.com/google/go-github/v31/github"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/kylelemons/godebug/pretty"
)

func TestAccGithubBranchProtection_basic(t *testing.T) {
	if err := testAccCheckOrganization(); err != nil {
		t.Skipf("Skipping because %s.", err.Error())
	}

	var protection github.Protection

	rn := "github_branch_protection.master"
	rString := acctest.RandString(5)
	repoName := fmt.Sprintf("tf-acc-test-branch-prot-%s", rString)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccGithubBranchProtectionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubBranchProtectionConfig(repoName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubProtectedBranchExists(rn, repoName+":master", &protection),
					testAccCheckGithubBranchProtectionRequiredStatusChecks(&protection, true, []string{"github/foo"}),
					testAccCheckGithubBranchProtectionRestrictions(&protection, []string{testUser}, []string{}),
					testAccCheckGithubBranchProtectionPullRequestReviews(&protection, true, []string{testUser}, []string{}, true),
					resource.TestCheckResourceAttr(rn, "repository", repoName),
					resource.TestCheckResourceAttr(rn, "branch", "master"),
					resource.TestCheckResourceAttr(rn, "enforce_admins", "true"),
					resource.TestCheckResourceAttr(rn, "require_signed_commits", "true"),
					resource.TestCheckResourceAttr(rn, "required_status_checks.0.strict", "true"),
					resource.TestCheckResourceAttr(rn, "required_status_checks.0.contexts.#", "1"),
					resource.TestCheckResourceAttr(rn, "required_pull_request_reviews.0.dismiss_stale_reviews", "true"),
					resource.TestCheckResourceAttr(rn, "required_pull_request_reviews.0.dismissal_users.#", "1"),
					resource.TestCheckResourceAttr(rn, "required_pull_request_reviews.0.dismissal_teams.#", "0"),
					resource.TestCheckResourceAttr(rn, "required_pull_request_reviews.0.require_code_owner_reviews", "true"),
					resource.TestCheckResourceAttr(rn, "restrictions.0.users.#", "1"),
					resource.TestCheckResourceAttr(rn, "restrictions.0.teams.#", "0"),
				),
			},
			{
				Config: testAccGithubBranchProtectionUpdateConfig(repoName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubProtectedBranchExists(rn, repoName+":master", &protection),
					testAccCheckGithubBranchProtectionRequiredStatusChecks(&protection, false, []string{"github/bar"}),
					testAccCheckGithubBranchProtectionNoRestrictionsExist(&protection),
					testAccCheckGithubBranchProtectionNoPullRequestReviewsExist(&protection),
					resource.TestCheckResourceAttr(rn, "repository", repoName),
					resource.TestCheckResourceAttr(rn, "branch", "master"),
					resource.TestCheckResourceAttr(rn, "require_signed_commits", "false"),
					resource.TestCheckResourceAttr(rn, "required_status_checks.0.strict", "false"),
					resource.TestCheckResourceAttr(rn, "required_status_checks.0.contexts.#", "1"),
					resource.TestCheckResourceAttr(rn, "required_pull_request_reviews.#", "0"),
					resource.TestCheckResourceAttr(rn, "restrictions.#", "0"),
				),
			},
			{
				ResourceName:      rn,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccGithubBranchProtection_users(t *testing.T) {
	if err := testAccCheckOrganization(); err != nil {
		t.Skipf("Skipping because %s.", err.Error())
	}

	rn := "github_branch_protection.master"
	rString := acctest.RandString(5)
	repoName := fmt.Sprintf("tf-acc-test-branch-prot-%s", rString)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccGithubBranchProtectionDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccGithubBranchProtectionConfigUser(repoName, "user_with_underscore"),
				ExpectError: regexp.MustCompile("unable to add users in restrictions: user_with_underscore"),
			},
			{
				Config: testAccGithubBranchProtectionConfigUser(repoName, testUser),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(rn, "repository", repoName),
					resource.TestCheckResourceAttr(rn, "branch", "master"),
					resource.TestCheckResourceAttr(rn, "enforce_admins", "true"),
					resource.TestCheckResourceAttr(rn, "restrictions.0.users.#", "1"),
				),
			},
			{
				ResourceName:      rn,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccGithubBranchProtection_teams(t *testing.T) {
	if err := testAccCheckOrganization(); err != nil {
		t.Skipf("Skipping because %s.", err.Error())
	}

	var firstP, secondP, thirdP github.Protection

	rn := "github_branch_protection.master"
	rString := acctest.RandString(5)
	repoName := fmt.Sprintf("tf-acc-test-branch-prot-%s", rString)
	firstTeamName := fmt.Sprintf("team 1 %s", rString)
	firstTeamSlug := fmt.Sprintf("team-1-%s", rString)
	secondTeamName := fmt.Sprintf("team 2 %s", rString)
	secondTeamSlug := fmt.Sprintf("team-2-%s", rString)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccGithubBranchProtectionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubBranchProtectionConfigTeams(repoName, firstTeamName, secondTeamName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubProtectedBranchExists(rn, repoName+":master", &firstP),
					testAccCheckGithubBranchProtectionRequiredStatusChecks(&firstP, false, []string{}),
					testAccCheckGithubBranchProtectionRestrictions(&firstP, []string{}, []string{firstTeamSlug, secondTeamSlug}),
					testAccCheckGithubBranchProtectionPullRequestReviews(&firstP, true, []string{}, []string{firstTeamSlug, secondTeamSlug}, false),
					resource.TestCheckResourceAttr(rn, "repository", repoName),
					resource.TestCheckResourceAttr(rn, "branch", "master"),
					resource.TestCheckResourceAttr(rn, "enforce_admins", "true"),
					resource.TestCheckResourceAttr(rn, "required_status_checks.0.strict", "false"),
					resource.TestCheckResourceAttr(rn, "required_status_checks.0.contexts.#", "0"),
					resource.TestCheckResourceAttr(rn, "required_pull_request_reviews.0.dismiss_stale_reviews", "true"),
					resource.TestCheckResourceAttr(rn, "required_pull_request_reviews.0.dismissal_users.#", "0"),
					resource.TestCheckResourceAttr(rn, "required_pull_request_reviews.0.dismissal_teams.#", "2"),
					resource.TestCheckResourceAttr(rn, "required_pull_request_reviews.0.require_code_owner_reviews", "false"),
					resource.TestCheckResourceAttr(rn, "restrictions.0.users.#", "0"),
					resource.TestCheckResourceAttr(rn, "restrictions.0.teams.#", "2"),
				),
			},
			{
				Config: testAccGithubBranchProtectionConfigEmptyItems(repoName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubProtectedBranchExists(rn, repoName+":master", &secondP),
					testAccCheckGithubBranchProtectionRequiredStatusChecks(&secondP, false, []string{}),
					resource.TestCheckResourceAttr(rn, "repository", repoName),
					resource.TestCheckResourceAttr(rn, "branch", "master"),
					resource.TestCheckResourceAttr(rn, "enforce_admins", "true"),
					resource.TestCheckResourceAttr(rn, "require_signed_commits", "false"),
					resource.TestCheckResourceAttr(rn, "required_status_checks.#", "1"),
					resource.TestCheckResourceAttr(rn, "required_pull_request_reviews.#", "1"),
					resource.TestCheckResourceAttr(rn, "restrictions.#", "1"),
				),
			},
			{
				Config: testAccGithubBranchProtectionConfigTeams(repoName, firstTeamName, secondTeamName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubProtectedBranchExists(rn, repoName+":master", &thirdP),
					testAccCheckGithubBranchProtectionRequiredStatusChecks(&thirdP, false, []string{}),
					testAccCheckGithubBranchProtectionRestrictions(&thirdP, []string{}, []string{firstTeamSlug, secondTeamSlug}),
					testAccCheckGithubBranchProtectionPullRequestReviews(&thirdP, true, []string{}, []string{firstTeamSlug, secondTeamSlug}, false),
					resource.TestCheckResourceAttr(rn, "repository", repoName),
					resource.TestCheckResourceAttr(rn, "branch", "master"),
					resource.TestCheckResourceAttr(rn, "enforce_admins", "true"),
					resource.TestCheckResourceAttr(rn, "required_status_checks.0.strict", "false"),
					resource.TestCheckResourceAttr(rn, "required_status_checks.0.contexts.#", "0"),
					resource.TestCheckResourceAttr(rn, "required_pull_request_reviews.0.dismiss_stale_reviews", "true"),
					resource.TestCheckResourceAttr(rn, "required_pull_request_reviews.0.dismissal_users.#", "0"),
					resource.TestCheckResourceAttr(rn, "required_pull_request_reviews.0.dismissal_teams.#", "2"),
					resource.TestCheckResourceAttr(rn, "required_pull_request_reviews.0.require_code_owner_reviews", "false"),
					resource.TestCheckResourceAttr(rn, "restrictions.0.users.#", "0"),
					resource.TestCheckResourceAttr(rn, "restrictions.0.teams.#", "2"),
				),
			},
			{
				ResourceName:      rn,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

// See https://github.com/terraform-providers/terraform-provider-github/issues/8
func TestAccGithubBranchProtection_emptyItems(t *testing.T) {
	if err := testAccCheckOrganization(); err != nil {
		t.Skipf("Skipping because %s.", err.Error())
	}

	var protection github.Protection

	rn := "github_branch_protection.master"
	rString := acctest.RandString(5)
	repoName := fmt.Sprintf("tf-acc-test-branch-prot-%s", rString)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccGithubBranchProtectionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubBranchProtectionConfigEmptyItems(repoName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubProtectedBranchExists("github_branch_protection.master", repoName+":master", &protection),
					resource.TestCheckResourceAttr(rn, "repository", repoName),
					resource.TestCheckResourceAttr(rn, "branch", "master"),
					resource.TestCheckResourceAttr(rn, "enforce_admins", "true"),
					resource.TestCheckResourceAttr(rn, "require_signed_commits", "false"),
					resource.TestCheckResourceAttr(rn, "required_status_checks.#", "1"),
					resource.TestCheckResourceAttr(rn, "required_pull_request_reviews.#", "1"),
					resource.TestCheckResourceAttr(rn, "restrictions.#", "1"),
				),
			},
			{
				ResourceName:      rn,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccGithubBranchProtection_emptyDismissalRestrictions(t *testing.T) {
	if err := testAccCheckOrganization(); err != nil {
		t.Skipf("Skipping because %s.", err.Error())
	}

	var protection github.Protection
	rn := "github_branch_protection.master"
	repoName := acctest.RandomWithPrefix("tf-acc-test-branch-prot-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccGithubBranchProtectionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubBranchProtectionEmptyDismissalsConfig(repoName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubProtectedBranchExists("github_branch_protection.master", repoName+":master", &protection),
					testAccCheckGithubBranchProtectionPullRequestReviews(&protection, true, []string{}, []string{}, true),
					resource.TestCheckResourceAttr(rn, "required_pull_request_reviews.#", "1"),
					resource.TestCheckResourceAttr(rn, "required_pull_request_reviews.0.dismiss_stale_reviews", "true"),
					resource.TestCheckResourceAttr(rn, "required_pull_request_reviews.0.require_code_owner_reviews", "true"),
					resource.TestCheckResourceAttr(rn, "required_pull_request_reviews.0.dismissal_users.#", "0"),
					resource.TestCheckResourceAttr(rn, "required_pull_request_reviews.0.dismissal_teams.#", "0"),
				),
			},
			{
				ResourceName:      rn,
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

		conn := testAccProvider.Meta().(*Owner).v3client
		o := testAccProvider.Meta().(*Owner).name
		r, b, err := parseTwoPartID(rs.Primary.ID, "repository", "branch")
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
		rsc := protection.GetRequiredStatusChecks()
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
		restrictions := protection.GetRestrictions()
		if restrictions == nil {
			return fmt.Errorf("Expected Restrictions to be present, but was nil")
		}

		userLogins := []string{}
		for _, u := range restrictions.Users {
			userLogins = append(userLogins, u.GetLogin())
		}
		if diff := pretty.Compare(userLogins, expectedUserLogins); diff != "" {
			return fmt.Errorf("diff %q: (-got +want)\n%s", "restrictions.users", diff)
		}

		teamLogins := []string{}
		for _, t := range restrictions.Teams {
			teamLogins = append(teamLogins, t.GetSlug())
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
		reviews := protection.GetRequiredPullRequestReviews()
		if reviews == nil {
			return fmt.Errorf("Expected Pull Request Reviews to be present, but was nil")
		}

		if reviews.DismissStaleReviews != expectedStale {
			return fmt.Errorf("Expected `dismiss_state_reviews` to be %t, got %t", expectedStale, reviews.DismissStaleReviews)
		}

		var users, teams []string
		if reviews.DismissalRestrictions != nil {
			if len(expectedUsers) == 0 && len(expectedTeams) == 0 {
				return fmt.Errorf("Expected Dismissal Restrictions to be nil but was present")
			}
			for _, u := range reviews.GetDismissalRestrictions().Users {
				users = append(users, u.GetLogin())
			}

			for _, t := range reviews.GetDismissalRestrictions().Teams {
				teams = append(teams, t.GetSlug())
			}
		}

		if diff := pretty.Compare(users, expectedUsers); diff != "" {
			return fmt.Errorf("diff %q: (-got +want)\n%s", "required_pull_request_reviews.dismissal_users", diff)
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
		if restrictions := protection.GetRestrictions(); restrictions != nil {
			return fmt.Errorf("Expected Restrictions to be nil, but was %v", restrictions)
		}

		return nil

	}
}

func testAccCheckGithubBranchProtectionNoPullRequestReviewsExist(protection *github.Protection) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if requiredPullRequestReviews := protection.GetRequiredPullRequestReviews(); requiredPullRequestReviews != nil {
			return fmt.Errorf("Expected Pull Request reviews to be nil, but was %v", requiredPullRequestReviews)
		}

		return nil
	}
}

func testAccGithubBranchProtectionDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*Owner).v3client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "github_branch_protection" {
			continue
		}

		o := testAccProvider.Meta().(*Owner).name
		r, b, err := parseTwoPartID(rs.Primary.ID, "repository", "branch")
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
  repository     = "${github_repository.test.name}"
  branch         = "master"
  enforce_admins = true
  require_signed_commits = true

  required_status_checks {
    strict   = true
    contexts = ["github/foo"]
  }

  required_pull_request_reviews {
    dismiss_stale_reviews      = true
    dismissal_users            = ["%s"]
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
    strict   = false
    contexts = ["github/bar"]
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
  repository     = "${github_repository.test.name}"
  branch         = "master"
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
  permission = "push"
}

resource "github_team" "second" {
  name = "%s"
}

resource "github_team_repository" "second" {
  team_id    = "${github_team.second.id}"
  repository = "${github_repository.test.name}"
  permission = "push"
}

resource "github_branch_protection" "master" {
  depends_on     = ["github_team_repository.first", "github_team_repository.second"]
  repository     = "${github_repository.test.name}"
  branch         = "master"
  enforce_admins = true

  required_status_checks {
  }

  required_pull_request_reviews {
    dismiss_stale_reviews = true
    dismissal_teams       = ["${github_team.first.slug}", "${github_team.second.slug}"]
  }

  restrictions {
    teams = ["${github_team.first.slug}", "${github_team.second.slug}"]
  }
}
`, repoName, repoName, firstTeamName, secondTeamName)
}

func testAccGithubBranchProtectionConfigUser(repoName, user string) string {
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

  restrictions {
    users = ["%s"]
  }
}
`, repoName, repoName, user)
}

func testAccGithubBranchProtectionEmptyDismissalsConfig(repoName string) string {
	return fmt.Sprintf(`

resource "github_repository" "test" {
	name        = "%s"
	description = "Terraform Acceptance Test %s"
	auto_init   = true
}

resource "github_branch_protection" "master" {
	repository     = "${github_repository.test.name}"
	branch         = "master"
	enforce_admins = true

	required_pull_request_reviews {
		dismiss_stale_reviews = true
		require_code_owner_reviews = true
	}
}
`, repoName, repoName)
}
