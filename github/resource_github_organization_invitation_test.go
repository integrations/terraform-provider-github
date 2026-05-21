package github

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/google/go-github/v83/github"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccGithubOrganizationInvitation_byEmail(t *testing.T) {
	email := os.Getenv("GH_TEST_INVITATION_EMAIL")
	if email == "" {
		t.Skip("GH_TEST_INVITATION_EMAIL not set, skipping email invitation test")
	}

	config := fmt.Sprintf(`
		resource "github_organization_invitation" "test" {
			email = "%s"
		}
	`, email)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { skipUnlessHasOrgs(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckGithubOrganizationInvitationDestroy,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("github_organization_invitation.test", "email", email),
					resource.TestCheckResourceAttr("github_organization_invitation.test", "role", "direct_member"),
					resource.TestCheckResourceAttrSet("github_organization_invitation.test", "invitation_id"),
				),
			},
		},
	})
}

func TestAccGithubOrganizationInvitation_byInviteeId(t *testing.T) {
	config := fmt.Sprintf(`
		data "github_user" "test" {
			username = "%s"
		}

		resource "github_organization_invitation" "test" {
			invitee_id = data.github_user.test.id
		}
	`, testAccConf.testExternalUser)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			skipUnlessHasOrgs(t)
			if testAccConf.testExternalUser == "" {
				t.Skip("GH_TEST_EXTERNAL_USER not set, skipping invitee_id test")
			}
		},
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckGithubOrganizationInvitationDestroy,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("github_organization_invitation.test", "invitee_id"),
					resource.TestCheckResourceAttr("github_organization_invitation.test", "role", "direct_member"),
					resource.TestCheckResourceAttrSet("github_organization_invitation.test", "invitation_id"),
				),
			},
		},
	})
}

func TestAccGithubOrganizationInvitation_adminRole(t *testing.T) {
	email := os.Getenv("GH_TEST_INVITATION_EMAIL")
	if email == "" {
		t.Skip("GH_TEST_INVITATION_EMAIL not set, skipping admin role invitation test")
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
		CheckDestroy:      testAccCheckGithubOrganizationInvitationDestroy,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("github_organization_invitation.test", "email", email),
					resource.TestCheckResourceAttr("github_organization_invitation.test", "role", "admin"),
					resource.TestCheckResourceAttrSet("github_organization_invitation.test", "invitation_id"),
				),
			},
		},
	})
}

func testAccCheckGithubOrganizationInvitationDestroy(s *terraform.State) error {
	meta, err := getTestMeta()
	if err != nil {
		return err
	}
	client := meta.v3client
	orgName := meta.name

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "github_organization_invitation" {
			continue
		}

		invitationID, err := strconv.ParseInt(rs.Primary.ID, 10, 64)
		if err != nil {
			return err
		}

		opts := &github.ListOptions{PerPage: 100}
		for {
			invitations, resp, err := client.Organizations.ListPendingOrgInvitations(context.TODO(), orgName, opts)
			if err != nil {
				return err
			}

			for _, inv := range invitations {
				if inv.GetID() == invitationID {
					return fmt.Errorf("organization invitation %d still exists", invitationID)
				}
			}

			if resp.NextPage == 0 {
				break
			}
			opts.Page = resp.NextPage
		}
	}

	return nil
}
