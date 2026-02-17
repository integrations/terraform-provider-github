package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccGithubEnterpriseSCIMUserDataSource(t *testing.T) {
	t.Run("reads user without error", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{{
				Config: fmt.Sprintf(`
					data "github_enterprise_scim_users" "all" {
						enterprise = "%s"
					}

					data "github_enterprise_scim_user" "test" {
						enterprise   = "%[1]s"
						scim_user_id = data.github_enterprise_scim_users.all.resources[0].id
					}
				`, testAccConf.enterpriseSlug),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.github_enterprise_scim_user.test", "id"),
					resource.TestCheckResourceAttrSet("data.github_enterprise_scim_user.test", "user_name"),
					resource.TestCheckResourceAttrSet("data.github_enterprise_scim_user.test", "display_name"),
					resource.TestCheckResourceAttrSet("data.github_enterprise_scim_user.test", "schemas.#"),
					resource.TestCheckResourceAttrSet("data.github_enterprise_scim_user.test", "emails.#"),
				),
			}},
		})
	})
}
