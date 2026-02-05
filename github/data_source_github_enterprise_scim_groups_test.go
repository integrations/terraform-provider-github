package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubEnterpriseSCIMGroupsDataSource(t *testing.T) {
	t.Run("lists groups without error", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{{
				Config: fmt.Sprintf(`
					data "github_enterprise_scim_groups" "test" {
						enterprise = "%s"
					}
				`, testAccConf.enterpriseSlug),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.github_enterprise_scim_groups.test", "id"),
					resource.TestCheckResourceAttrSet("data.github_enterprise_scim_groups.test", "total_results"),
					resource.TestCheckResourceAttrSet("data.github_enterprise_scim_groups.test", "schemas.#"),
					resource.TestCheckResourceAttrSet("data.github_enterprise_scim_groups.test", "resources.#"),
				),
			}},
		})
	})
}
