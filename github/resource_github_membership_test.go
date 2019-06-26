package github

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/google/go-github/v25/github"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccGithubMembership_basic(t *testing.T) {
	if testCollaborator == "" {
		t.Skip("Skipping because `GITHUB_TEST_COLLABORATOR` is not set")
	}

	var membership github.Membership

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGithubMembershipDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubMembershipConfig(testCollaborator),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubMembershipExists("github_membership.test_org_membership", &membership),
					testAccCheckGithubMembershipRoleState("github_membership.test_org_membership", &membership),
				),
			},
		},
	})
}

func TestAccGithubMembership_caseInsensitive(t *testing.T) {
	if testCollaborator == "" {
		t.Skip("Skipping because `GITHUB_TEST_COLLABORATOR` is not set")
	}

	var membership github.Membership
	var otherMembership github.Membership

	otherCase := flipUsernameCase(testCollaborator)

	if testCollaborator == otherCase {
		t.Skip("Skipping because `GITHUB_TEST_COLLABORATOR` has no letters to flip case")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGithubMembershipDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubMembershipConfig(testCollaborator),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubMembershipExists("github_membership.test_org_membership", &membership),
				),
			},
			{
				Config: testAccGithubMembershipConfig(otherCase),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubMembershipExists("github_membership.test_org_membership", &otherMembership),
					testAccGithubMembershipTheSame(&membership, &otherMembership),
				),
			},
		},
	})
}

func TestAccGithubMembership_importBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGithubMembershipDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubMembershipConfig(testCollaborator),
			},
			{
				ResourceName:      "github_membership.test_org_membership",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckGithubMembershipDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*Organization).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "github_membership" {
			continue
		}
		orgName, username, err := parseTwoPartID(rs.Primary.ID)
		if err != nil {
			return err
		}

		membership, resp, err := conn.Organizations.GetOrgMembership(context.TODO(), username, orgName)

		if err == nil {
			if membership != nil &&
				buildTwoPartID(membership.Organization.Login, membership.User.Login) == rs.Primary.ID {
				return fmt.Errorf("Organization membership %q still exists", rs.Primary.ID)
			}
		}
		if resp.StatusCode != 404 {
			return err
		}
		return nil
	}
	return nil
}

func testAccCheckGithubMembershipExists(n string, membership *github.Membership) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not Found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No membership ID is set")
		}

		conn := testAccProvider.Meta().(*Organization).client
		orgName, username, err := parseTwoPartID(rs.Primary.ID)
		if err != nil {
			return err
		}

		githubMembership, _, err := conn.Organizations.GetOrgMembership(context.TODO(), username, orgName)
		if err != nil {
			return err
		}
		*membership = *githubMembership
		return nil
	}
}

func testAccCheckGithubMembershipRoleState(n string, membership *github.Membership) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not Found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No membership ID is set")
		}

		conn := testAccProvider.Meta().(*Organization).client
		orgName, username, err := parseTwoPartID(rs.Primary.ID)
		if err != nil {
			return err
		}

		githubMembership, _, err := conn.Organizations.GetOrgMembership(context.TODO(), username, orgName)
		if err != nil {
			return err
		}

		resourceRole := membership.Role
		actualRole := githubMembership.Role

		if *resourceRole != *actualRole {
			return fmt.Errorf("Membership role %v in resource does match actual state of %v",
				*resourceRole, *actualRole)
		}
		return nil
	}
}

func testAccGithubMembershipConfig(username string) string {
	return fmt.Sprintf(`
  resource "github_membership" "test_org_membership" {
    username = "%s"
    role = "member"
  }
`, username)
}

func testAccGithubMembershipTheSame(orig, other *github.Membership) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if *orig.URL != *other.URL {
			return errors.New("users are different")
		}

		return nil
	}
}
