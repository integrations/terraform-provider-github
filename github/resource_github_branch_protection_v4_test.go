package github

import (
	"context"
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/kylelemons/godebug/pretty"
	"github.com/shurcooL/githubv4"
)

func TestAccGithubBranchProtection_basic(t *testing.T) {
	if err := testAccCheckOrganization(); err != nil {
		t.Skipf("Skipping because %s.", err.Error())
	}

	var protection BranchProtectionRule

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
					testAccCheckGithubProtectedBranchExists(rn, &protection),
					testAccCheckGithubBranchProtectionRequiredStatusChecks(&protection, true, []githubv4.String{"github/foo"}),
					testAccCheckGithubBranchProtectionRestrictions(&protection, []githubv4.String{githubv4.String(testUser)}),
					testAccCheckGithubBranchProtectionPullRequestReviews(&protection, true, []githubv4.String{githubv4.String(testUser)}, true),
					resource.TestCheckResourceAttr(rn, "repository_id", repoName),
					resource.TestCheckResourceAttr(rn, "pattern", "master"),
					resource.TestCheckResourceAttr(rn, "enforce_admins", "true"),
					resource.TestCheckResourceAttr(rn, "require_signed_commits", "true"),
					resource.TestCheckResourceAttr(rn, "required_status_checks.0.strict", "true"),
					resource.TestCheckResourceAttr(rn, "required_status_checks.0.contexts.#", "1"),
					resource.TestCheckResourceAttr(rn, "required_pull_request_reviews.0.dismiss_stale_reviews", "true"),
					resource.TestCheckResourceAttr(rn, "required_pull_request_reviews.0.dismissal_restrictions.0.actor_ids.#", "1"),
					resource.TestCheckResourceAttr(rn, "required_pull_request_reviews.0.require_code_owner_reviews", "true"),
					resource.TestCheckResourceAttr(rn, "push_restrictions.0.actor_ids.#", "1"),
				),
			},
			{
				Config: testAccGithubBranchProtectionUpdateConfig(repoName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubProtectedBranchExists(rn, &protection),
					testAccCheckGithubBranchProtectionRequiredStatusChecks(&protection, false, []githubv4.String{"github/bar"}),
					testAccCheckGithubBranchProtectionNoRestrictionsExist(&protection),
					resource.TestCheckResourceAttr(rn, "repository_id", repoName),
					resource.TestCheckResourceAttr(rn, "pattern", "master"),
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

	var firstP, secondP, thirdP BranchProtectionRule

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
					testAccCheckGithubProtectedBranchExists(rn, &firstP),
					testAccCheckGithubBranchProtectionRequiredStatusChecks(&firstP, false, []githubv4.String{}),
					testAccCheckGithubBranchProtectionRestrictions(&firstP, []githubv4.String{githubv4.String(firstTeamSlug), githubv4.String(secondTeamSlug)}),
					testAccCheckGithubBranchProtectionPullRequestReviews(&firstP, true, []githubv4.String{githubv4.String(firstTeamSlug), githubv4.String(secondTeamSlug)}, false),
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
					testAccCheckGithubProtectedBranchExists(rn, &secondP),
					testAccCheckGithubBranchProtectionRequiredStatusChecks(&secondP, false, []githubv4.String{}),
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
					testAccCheckGithubProtectedBranchExists(rn, &thirdP),
					testAccCheckGithubBranchProtectionRequiredStatusChecks(&thirdP, false, []githubv4.String{}),
					testAccCheckGithubBranchProtectionRestrictions(&thirdP, []githubv4.String{githubv4.String(firstTeamSlug), githubv4.String(secondTeamSlug)}),
					testAccCheckGithubBranchProtectionPullRequestReviews(&thirdP, true, []githubv4.String{githubv4.String(firstTeamSlug), githubv4.String(secondTeamSlug)}, false),
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

	var protection BranchProtectionRule
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
					testAccCheckGithubProtectedBranchExists("github_branch_protection.master", &protection),
					resource.TestCheckResourceAttr(rn, "repository_id", repoName),
					resource.TestCheckResourceAttr(rn, "pattern", "master"),
					resource.TestCheckResourceAttr(rn, "enforce_admins", "true"),
					resource.TestCheckResourceAttr(rn, "require_signed_commits", "false"),
					resource.TestCheckResourceAttr(rn, "required_status_checks.#", "1"),
					resource.TestCheckResourceAttr(rn, "required_pull_request_reviews.#", "1"),
					resource.TestCheckResourceAttr(rn, "push_restrictions.#", "1"),
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

	var protection BranchProtectionRule
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
					testAccCheckGithubProtectedBranchExists("github_branch_protection.master", &protection),
					testAccCheckGithubBranchProtectionPullRequestReviews(&protection, true, []githubv4.String{}, true),
					resource.TestCheckResourceAttr(rn, "required_pull_request_reviews.#", "1"),
					resource.TestCheckResourceAttr(rn, "required_pull_request_reviews.0.dismiss_stale_reviews", "true"),
					resource.TestCheckResourceAttr(rn, "required_pull_request_reviews.0.require_code_owner_reviews", "true"),
					resource.TestCheckResourceAttr(rn, "required_pull_request_reviews.0.dismissal_restrictions.0.actor_ids.#", "0"),
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

func testAccCheckGithubProtectedBranchExists(n string, protection *BranchProtectionRule) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not Found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Branch Protection ID is set")
		}

		var query struct {
			Node struct {
				Node BranchProtectionRule `graphql:"... on BranchProtectionRule"`
			} `graphql:"node(id: $id)"`
		}
		variables := map[string]interface{}{
			"id": rs.Primary.ID,
		}

		client := testAccProvider.Meta().(*Owner).v4client
		err := client.Query(context.TODO(), &query, variables)
		if err != nil {
			return err
		}

		*protection = query.Node.Node
		return nil
	}
}

func testAccCheckGithubBranchProtectionRequiredStatusChecks(protection *BranchProtectionRule, expectedStrict githubv4.Boolean, expectedContexts []githubv4.String) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if protection.RequiredStatusCheckContexts == nil {
			return fmt.Errorf("Expected RequiredStatusChecks to be present, but was nil")
		}

		if protection.RequiresStrictStatusChecks != expectedStrict {
			return fmt.Errorf("Expected RequiredStatusChecks.Strict to be %v, got %v", expectedStrict, bool(protection.RequiresStrictStatusChecks))
		}

		if diff := pretty.Compare(protection.RequiredStatusCheckContexts, expectedContexts); diff != "" {
			return fmt.Errorf("diff %q: (-got +want)\n%s", "contexts", diff)
		}

		return nil
	}
}

func testAccCheckGithubBranchProtectionRestrictions(protection *BranchProtectionRule, expectedActors []githubv4.String) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		restrictions := protection.PushAllowances.Nodes
		if restrictions == nil {
			return fmt.Errorf("Expected Restrictions to be present, but was nil")
		}
		actorIDs := setActorIDs(restrictions)
		if diff := pretty.Compare(actorIDs, expectedActors); diff != "" {
			return fmt.Errorf("diff %q: (-got +want)\n%s", "restrictions.users", diff)
		}

		return nil
	}
}

func testAccCheckGithubBranchProtectionPullRequestReviews(protection *BranchProtectionRule, expectedStale bool, expectedActors []githubv4.String, expectedCodeOwners bool) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		reviews := protection.RequiresApprovingReviews
		if reviews == false {
			return fmt.Errorf("Expected Pull Request Reviews to be true, but was false")
		}

		stale := bool(protection.DismissesStaleReviews)
		if stale != expectedStale {
			return fmt.Errorf("Expected `dismiss_state_reviews` to be %t, got %t", expectedStale, stale)
		}

		dismissalAllowances := protection.ReviewDismissalAllowances.Nodes
		actorIDs := setActorIDs(dismissalAllowances)
		if diff := pretty.Compare(actorIDs, expectedActors); diff != "" {
			return fmt.Errorf("diff %q: (-got +want)\n%s", "required_pull_request_reviews.dismissal_users", diff)
		}

		requiresCodeOwnerReviews := bool(protection.RequiresCodeOwnerReviews)
		if requiresCodeOwnerReviews != expectedCodeOwners {
			return fmt.Errorf("Expected `require_code_owner_reviews` to be %t, got %t", expectedCodeOwners, requiresCodeOwnerReviews)
		}

		return nil
	}
}

func testAccCheckGithubBranchProtectionNoRestrictionsExist(protection *BranchProtectionRule) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if restrictions := protection.PushAllowances.Nodes; restrictions != nil {
			return fmt.Errorf("Expected Restrictions to be nil, but was %v", restrictions)
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
  repository_id          = "${github_repository.test.node_id}"
  pattern                = "master"
  enforce_admins         = true
  require_signed_commits = true

  required_status_checks {
    strict   = true
    contexts = ["github/foo"]
  }

  required_pull_request_reviews {
    dismiss_stale_reviews      = true
    require_code_owner_reviews = true
    dismissal_restrictions {
	  actor_ids = ["%s"]
	}
  }

  push_restrictions {
    actor_ids = ["%s"]
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
  repository_id = "${github_repository.test.node_id}"
  pattern       = "master"

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
  repository_id  = "${github_repository.test.node_id}"
  pattern        = "master"
  enforce_admins = true

  required_status_checks {
  }

  required_pull_request_reviews {
  }

  push_restrictions {
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
  repository = "${github_repository.test.node_id}"
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
  repository_id  = "${github_repository.test.node_id}"
  pattern        = "master"
  enforce_admins = true

  required_status_checks {
  }

  required_pull_request_reviews {
    dismiss_stale_reviews = true
    dismissal_restrictions {
	  actor_ids = ["${github_team.first.slug}", "${github_team.second.slug}"]
	}
  }

  push_restrictions {
    actor_ids = ["${github_team.first.slug}", "${github_team.second.slug}"]
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
  repository_id  = "${github_repository.test.name}"
  pattern        = "master"
  enforce_admins = true

  push_restrictions {
    actor_ids = ["%s"]
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
	repository_id  = "${github_repository.test.name}"
	pattern        = "master"
	enforce_admins = true

	required_pull_request_reviews {
		dismiss_stale_reviews      = true
		require_code_owner_reviews = true
	}
}
`, repoName, repoName)
}
