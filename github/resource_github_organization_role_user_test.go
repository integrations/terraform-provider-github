package github

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubOrganizationRoleUser(t *testing.T) {
	t.Run("adds user to an organization org role", func(t *testing.T) {
		roleId := 8134
		config := fmt.Sprintf(`
			resource "github_organization_role_user" "test" {
				role_id  = %d
				login = "%s"
			}
		`, roleId, testAccConf.testOrgUser)

		resource.Test(t, resource.TestCase{
			PreCheck:  func() { skipUnlessHasOrgs(t) },
			Providers: testAccProviders,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_organization_role_user.test", "role_id", strconv.Itoa(roleId)),
						resource.TestCheckResourceAttr("github_organization_role_user.test", "login", testAccConf.testOrgUser),
					),
				},
			},
		})
	})
}
