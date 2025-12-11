package github

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubOrganizationIpAllowListDataSource(t *testing.T) {
	t.Run("queries without error", func(t *testing.T) {
		config := `
			data "github_organization_ip_allow_list" "all" {}
		`

		check := resource.ComposeAggregateTestCheckFunc(
			resource.TestCheckResourceAttrSet("data.github_organization_ip_allow_list.all", "ip_allow_list.0.id"),
			resource.TestCheckResourceAttrSet("data.github_organization_ip_allow_list.all", "ip_allow_list.0.name"),
			resource.TestCheckResourceAttrSet("data.github_organization_ip_allow_list.all", "ip_allow_list.0.allow_list_value"),
			resource.TestCheckResourceAttrSet("data.github_organization_ip_allow_list.all", "ip_allow_list.0.is_active"),
			resource.TestCheckResourceAttrSet("data.github_organization_ip_allow_list.all", "ip_allow_list.0.created_at"),
			resource.TestCheckResourceAttrSet("data.github_organization_ip_allow_list.all", "ip_allow_list.0.updated_at"),
		)

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: config,
						Check:  check,
					},
				},
			})
		}

		t.Run("with an anonymous account", func(t *testing.T) {
			t.Skip("anonymous account not supported for this operation")
		})

		t.Run("with an individual account", func(t *testing.T) {
			t.Skip("individual account not supported for this operation")
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})
	})
}
