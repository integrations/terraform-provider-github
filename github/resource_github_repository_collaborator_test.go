package github

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/google/go-github/v28/github"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const expectedPermission string = "admin"

func TestAccGithubRepositoryCollaborator_basic(t *testing.T) {
	if testCollaborator == "" {
		t.Skip("Skipping because `GITHUB_TEST_COLLABORATOR` is not set")
	}

	rn := "github_repository_collaborator.test_repo_collaborator"
	repoName := fmt.Sprintf("tf-acc-test-collab-%s", acctest.RandString(5))

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGithubRepositoryCollaboratorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubRepositoryCollaboratorConfig(repoName, testCollaborator),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubRepositoryCollaboratorExists(rn),
					testAccCheckGithubRepositoryCollaboratorPermission(rn),
					resource.TestCheckResourceAttr(rn, "permission", expectedPermission),
					resource.TestMatchResourceAttr(rn, "invitation_id", regexp.MustCompile(`^[0-9]+$`)),
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

func TestAccGithubRepositoryCollaborator_caseInsensitive(t *testing.T) {
	if testCollaborator == "" {
		t.Skip("Skipping because `GITHUB_TEST_COLLABORATOR` is not set")
	}

	rn := "github_repository_collaborator.test_repo_collaborator"
	repoName := fmt.Sprintf("tf-acc-test-collab-%s", acctest.RandString(5))

	var origInvitation github.RepositoryInvitation
	var otherInvitation github.RepositoryInvitation

	otherCase := flipUsernameCase(testCollaborator)

	if testCollaborator == otherCase {
		t.Skip("Skipping because `GITHUB_TEST_COLLABORATOR` has no letters to flip case")
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGithubRepositoryCollaboratorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubRepositoryCollaboratorConfig(repoName, testCollaborator),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubRepositoryCollaboratorInvited(repoName, testCollaborator, &origInvitation),
				),
			},
			{
				Config: testAccGithubRepositoryCollaboratorConfig(repoName, otherCase),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubRepositoryCollaboratorInvited(repoName, otherCase, &otherInvitation),
					resource.TestCheckResourceAttr(rn, "username", testCollaborator),
					testAccGithubRepositoryCollaboratorTheSame(&origInvitation, &otherInvitation),
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

func testAccCheckGithubRepositoryCollaboratorDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*Organization).v3client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "github_repository_collaborator" {
			continue
		}

		o := testAccProvider.Meta().(*Organization).name
		r, u, err := parseTwoPartID(rs.Primary.ID, "repository", "username")
		if err != nil {
			return err
		}

		isCollaborator, _, err := conn.Repositories.IsCollaborator(context.TODO(), o, r, u)

		if err != nil {
			return err
		}

		if isCollaborator {
			return fmt.Errorf("Repository collaborator still exists")
		}

		return nil
	}

	return nil
}

func testAccCheckGithubRepositoryCollaboratorExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not Found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No membership ID is set")
		}

		conn := testAccProvider.Meta().(*Organization).v3client
		orgName := testAccProvider.Meta().(*Organization).name
		repoName, username, err := parseTwoPartID(rs.Primary.ID, "repository", "username")
		if err != nil {
			return err
		}

		invitations, _, err := conn.Repositories.ListInvitations(context.TODO(),
			orgName, repoName, nil)
		if err != nil {
			return err
		}

		hasInvitation := false
		for _, i := range invitations {
			if *i.Invitee.Login == username {
				hasInvitation = true
				break
			}
		}

		if !hasInvitation {
			return fmt.Errorf("Repository collaboration invitation does not exist")
		}

		return nil
	}
}

func testAccCheckGithubRepositoryCollaboratorPermission(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not Found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No membership ID is set")
		}

		conn := testAccProvider.Meta().(*Organization).v3client
		orgName := testAccProvider.Meta().(*Organization).name
		repoName, username, err := parseTwoPartID(rs.Primary.ID, "repository", "username")
		if err != nil {
			return err
		}

		invitations, _, err := conn.Repositories.ListInvitations(context.TODO(),
			orgName, repoName, nil)
		if err != nil {
			return err
		}

		for _, i := range invitations {
			if *i.Invitee.Login == username {
				permName, err := getInvitationPermission(i)

				if err != nil {
					return err
				}

				if permName != expectedPermission {
					return fmt.Errorf("Expected permission %s on repository collaborator, actual permission %s", expectedPermission, permName)
				}

				return nil
			}
		}

		return fmt.Errorf("Repository collaborator did not appear in list of collaborators on repository")
	}
}

func testAccGithubRepositoryCollaboratorConfig(repoName, username string) string {
	return fmt.Sprintf(`
resource "github_repository" "test" {
  name = "%s"
}

resource "github_repository_collaborator" "test_repo_collaborator" {
  repository = "${github_repository.test.name}"
  username   = "%s"
  permission = "%s"
}
`, repoName, username, expectedPermission)
}

func testAccCheckGithubRepositoryCollaboratorInvited(repoName, username string, invitation *github.RepositoryInvitation) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		opt := &github.ListOptions{PerPage: maxPerPage}

		client := testAccProvider.Meta().(*Organization).v3client
		org := testAccProvider.Meta().(*Organization).name

		for {
			invitations, resp, err := client.Repositories.ListInvitations(context.TODO(), org, repoName, opt)
			if err != nil {
				return errors.New(err.Error())
			}

			if len(invitations) > 1 {
				return fmt.Errorf("multiple invitations have been sent for repository %s", repoName)
			}

			for _, i := range invitations {
				if strings.EqualFold(*i.Invitee.Login, username) {
					invitation = i
					return nil
				}
			}

			if resp.NextPage == 0 {
				break
			}
			opt.Page = resp.NextPage
		}

		return fmt.Errorf("no invitation found for %s", username)
	}
}

func testAccGithubRepositoryCollaboratorTheSame(orig, other *github.RepositoryInvitation) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if orig.ID != other.ID {
			return errors.New("collaborators are different")
		}

		return nil
	}
}
