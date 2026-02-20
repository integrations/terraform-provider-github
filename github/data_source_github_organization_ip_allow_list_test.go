package github

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccGithubOrganizationIpAllowListDataSource(t *testing.T) {
	t.Run("queries without error", func(t *testing.T) {
		config := `
			data "github_organization_ip_allow_list" "all" {}
		`

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttrSet("data.github_organization_ip_allow_list.all", "ip_allow_list.#"),
					),
				},
			},
		})
	})
}
