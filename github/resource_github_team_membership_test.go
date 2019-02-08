package github

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/google/go-github/v21/github"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccGithubTeamMembership_basic(t *testing.T) {
	var membership github.Membership
	randString := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGithubTeamMembershipDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubTeamMembershipConfig(randString, testCollaborator, "member"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubTeamMembershipExists("github_team_membership.test_team_membership", &membership),
					testAccCheckGithubTeamMembershipRoleState("github_team_membership.test_team_membership", "member", &membership),
				),
			},
			{
				Config: testAccGithubTeamMembershipConfig(randString, testCollaborator, "maintainer"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubTeamMembershipExists("github_team_membership.test_team_membership", &membership),
					testAccCheckGithubTeamMembershipRoleState("github_team_membership.test_team_membership", "maintainer", &membership),
				),
			},
		},
	})
}

func TestAccGithubTeamMembership_importBasic(t *testing.T) {
	randString := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGithubTeamMembershipDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubTeamMembershipConfig(randString, testCollaborator, "member"),
			},
			{
				ResourceName:      "github_team_membership.test_team_membership",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckGithubTeamMembershipDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*Organization).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "github_team_membership" {
			continue
		}

		teamIdString, username, err := parseTwoPartID(rs.Primary.ID)
		if err != nil {
			return err
		}

		teamId, err := strconv.ParseInt(teamIdString, 10, 64)
		if err != nil {
			return unconvertibleIdErr(teamIdString, err)
		}

		membership, resp, err := conn.Teams.GetTeamMembership(context.TODO(),
			teamId, username)
		if err == nil {
			if membership != nil {
				return fmt.Errorf("Team membership still exists")
			}
		}
		if resp.StatusCode != 404 {
			return err
		}
		return nil
	}
	return nil
}

func testAccCheckGithubTeamMembershipExists(n string, membership *github.Membership) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not Found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No team membership ID is set")
		}

		conn := testAccProvider.Meta().(*Organization).client
		teamIdString, username, err := parseTwoPartID(rs.Primary.ID)
		if err != nil {
			return err
		}

		teamId, err := strconv.ParseInt(teamIdString, 10, 64)
		if err != nil {
			return unconvertibleIdErr(teamIdString, err)
		}

		teamMembership, _, err := conn.Teams.GetTeamMembership(context.TODO(), teamId, username)

		if err != nil {
			return err
		}
		*membership = *teamMembership
		return nil
	}
}

func testAccCheckGithubTeamMembershipRoleState(n, expected string, membership *github.Membership) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not Found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No team membership ID is set")
		}

		conn := testAccProvider.Meta().(*Organization).client
		teamIdString, username, err := parseTwoPartID(rs.Primary.ID)
		if err != nil {
			return err
		}
		teamId, err := strconv.ParseInt(teamIdString, 10, 64)
		if err != nil {
			return unconvertibleIdErr(teamIdString, err)
		}

		teamMembership, _, err := conn.Teams.GetTeamMembership(context.TODO(),
			teamId, username)
		if err != nil {
			return err
		}

		resourceRole := membership.Role
		actualRole := teamMembership.Role

		if *resourceRole != expected {
			return fmt.Errorf("Team membership role %v in resource does match expected state of %v", *resourceRole, expected)
		}

		if *resourceRole != *actualRole {
			return fmt.Errorf("Team membership role %v in resource does match actual state of %v", *resourceRole, *actualRole)
		}
		return nil
	}
}

func testAccGithubTeamMembershipConfig(randString, username, role string) string {
	return fmt.Sprintf(`
resource "github_membership" "test_org_membership" {
  username = "%s"
  role = "member"
}

resource "github_team" "test_team" {
  name = "tf-acc-test-team-membership-%s"
  description = "Terraform acc test group"
}

resource "github_team_membership" "test_team_membership" {
  team_id = "${github_team.test_team.id}"
  username = "%s"
  role = "%s"
}
`, username, randString, username, role)
}
