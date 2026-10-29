package github

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccGithubMembershipDataSource(t *testing.T) {
	t.Parallel()
	skipUnlessHasOrgs(t)
	skipUnlessHasOrgUser1(t)

	t.Run("queries_the_membership_for_a_user_in_a_specified_organization", func(t *testing.T) {
		t.Parallel()

		config := fmt.Sprintf(`
			data "github_membership" "test" {
				username = "%s"
				organization = "%s"
			}
		`, testAccConf.testOrgUser1, testAccConf.owner)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("data.github_membership.test", "username", testAccConf.testOrgUser1),
						resource.TestCheckResourceAttrSet("data.github_membership.test", "role"),
						resource.TestCheckResourceAttrSet("data.github_membership.test", "etag"),
						resource.TestCheckResourceAttrSet("data.github_membership.test", "state"),
					),
				},
			},
		})
	})

	t.Run("errors_when_querying_with_a_nonexistent_user", func(t *testing.T) {
		t.Parallel()

		config := fmt.Sprintf(`
			data "github_membership" "test" {
				username = "!%s"
				organization = "%s"
			}
		`, testAccConf.testOrgUser1, testAccConf.owner)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      config,
					ExpectError: regexp.MustCompile(`Not Found`),
				},
			},
		})
	})
}
