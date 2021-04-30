package github

import (
	"fmt"
	"testing"

	"github.com/google/go-github/v32/github"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccGithubOrganizationMemberPrivileges_basic(t *testing.T) {
	if err := testAccCheckOrganization(); err != nil {
		t.Skipf("Skipping because %s.", err.Error())
	}

	var organization github.Organization

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubOrganizationMemberPrivilegesConfigBefore,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubOrganizationMemberPrivilegesAttributes(&organization, &testAccGithubOrganizationMemberPrivilegesExpectedAttributes{
						DefaultRepoPermission: "none",
						MembersCanCreateRepos: false,
					}),
				),
			},
			{
				Config: testAccGithubOrganizationMemberPrivilegesConfigAfter,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubOrganizationMemberPrivilegesAttributes(&organization, &testAccGithubOrganizationMemberPrivilegesExpectedAttributes{
						DefaultRepoPermission: "read",
						MembersCanCreateRepos: true,
					}),
				),
			},
			{
				ResourceName:      "github_organization_member_privileges.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

type testAccGithubOrganizationMemberPrivilegesExpectedAttributes struct {
	DefaultRepoPermission string
	MembersCanCreateRepos bool
}

func testAccCheckGithubOrganizationMemberPrivilegesAttributes(organization *github.Organization, want *testAccGithubOrganizationMemberPrivilegesExpectedAttributes) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if defaultRepoPermission := organization.GetDefaultRepoPermission(); defaultRepoPermission != want.DefaultRepoPermission {
			return fmt.Errorf("got defaultRepoPermission %q; want %q", defaultRepoPermission, want.DefaultRepoPermission)
		}
		if membersCanCreateRepos := organization.GetMembersCanCreateRepos(); membersCanCreateRepos != want.MembersCanCreateRepos {
			return fmt.Errorf("got membersCanCreateRepos %t; want %t", membersCanCreateRepos, want.MembersCanCreateRepos)
		}

		return nil
	}
}

const testAccGithubOrganizationMemberPrivilegesConfigBefore = `
resource "github_organization_member_privileges" "test" {
	default_repository_permission   = "none"
	members_can_create_repositories = false
}
`

const testAccGithubOrganizationMemberPrivilegesConfigAfter = `
resource "github_organization_member_privileges" "test" {
	default_repository_permission   = "read"
	members_can_create_repositories = true
}
`
