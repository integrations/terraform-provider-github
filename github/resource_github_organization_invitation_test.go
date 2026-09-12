package github

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccGithubOrganizationInvitation(t *testing.T) {
	t.Run("invite by email", func(t *testing.T) {
		email := os.Getenv("GH_TEST_INVITATION_EMAIL")
		if email == "" {
			t.Skip("GH_TEST_INVITATION_EMAIL not set")
		}

		config := fmt.Sprintf(`
resource "github_organization_invitation" "test" {
  email = "%s"
}
`, email)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_organization_invitation.test", "email", email),
						resource.TestCheckResourceAttr("github_organization_invitation.test", "role", "direct_member"),
						resource.TestCheckResourceAttrSet("github_organization_invitation.test", "id"),
					),
				},
			},
		})
	})

	t.Run("invite by invitee_id", func(t *testing.T) {
		if testAccConf.testExternalUser == "" {
			t.Skip("GH_TEST_EXTERNAL_USER not set")
		}

		config := fmt.Sprintf(`
data "github_user" "test" {
  username = "%s"
}

resource "github_organization_invitation" "test" {
  invitee_id = data.github_user.test.id
}
`, testAccConf.testExternalUser)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttrSet("github_organization_invitation.test", "invitee_id"),
						resource.TestCheckResourceAttr("github_organization_invitation.test", "role", "direct_member"),
						resource.TestCheckResourceAttrSet("github_organization_invitation.test", "login"),
					),
				},
			},
		})
	})

	t.Run("invite with admin role", func(t *testing.T) {
		email := os.Getenv("GH_TEST_INVITATION_EMAIL")
		if email == "" {
			t.Skip("GH_TEST_INVITATION_EMAIL not set")
		}

		config := fmt.Sprintf(`
resource "github_organization_invitation" "test" {
  email = "%s"
  role  = "admin"
}
`, email)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_organization_invitation.test", "email", email),
						resource.TestCheckResourceAttr("github_organization_invitation.test", "role", "admin"),
					),
				},
			},
		})
	})
}
