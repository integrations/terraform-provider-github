package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubEnterpriseSCIMGroupDataSource(t *testing.T) {
	t.Run("reads group without error", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{{
				Config: fmt.Sprintf(`
					data "github_enterprise_scim_groups" "all" {
						enterprise = "%s"
					}

					data "github_enterprise_scim_group" "test" {
						enterprise    = "%[1]s"
						scim_group_id = data.github_enterprise_scim_groups.all.resources[0].id
					}
				`, testAccConf.enterpriseSlug),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.github_enterprise_scim_group.test", "id"),
					resource.TestCheckResourceAttrSet("data.github_enterprise_scim_group.test", "display_name"),
					resource.TestCheckResourceAttrSet("data.github_enterprise_scim_group.test", "schemas.#"),
				),
			}},
		})
	})
}
