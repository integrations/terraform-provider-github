package github

import (
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccGithubUserInvitationAccepter_basic(t *testing.T) {
	resourceName := "github_repository_collaborator.test"
	repoName := fmt.Sprintf("tf-acc-test-collab-%s", acctest.RandString(5))

	inviteeToken := os.Getenv("GITHUB_TEST_COLLABORATOR_TOKEN")
	if inviteeToken == "" {
		t.Skip("GITHUB_TEST_COLLABORATOR_TOKEN was not provided, skipping test")
	}

	var providers []*schema.Provider

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories(&providers),
		CheckDestroy:      testAccCheckGithubUserInvitationAccepterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubUserInvitationAccepterConfig(inviteeToken, repoName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "permission", "push"),
					resource.TestMatchResourceAttr(resourceName, "invitation_id", regexp.MustCompile(`^[0-9]+$`)),
				),
			},
		},
	})
}

func testAccCheckGithubUserInvitationAccepterDestroy(s *terraform.State) error {
	return nil
}

func testAccGithubUserInvitationAccepterConfig(inviteeToken, repoName string) string {
	return fmt.Sprintf(`
provider "github" {
  alias = "main"
}

provider "github" {
  alias = "invitee"
  token = "%s"
}

resource "github_repository" "test" {
  provider = "github.main"
  name = "%s"
}

resource "github_repository_collaborator" "test" {
  provider = "github.main"
  repository = "${github_repository.test.name}"
  username = "%s"
  permission = "push"
}

resource "github_user_invitation_accepter" "test" {
  provider = "github.invitee"
  invitation_id = "${github_repository_collaborator.test.invitation_id}"
}
`, inviteeToken, repoName, testCollaborator)
}
