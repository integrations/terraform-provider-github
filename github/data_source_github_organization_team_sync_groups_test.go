package github

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccGithubOrganizationTeamSyncGroupsDataSource_existing(t *testing.T) {
	if isEnterprise != "true" {
		t.Skip("Skipping because `ENTERPRISE_ACCOUNT` is not set or set to false")
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckGithubOrganizationTeamSyncGroupsDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.github_organization_team_sync_groups.test", "groups.#"),
					resource.TestCheckResourceAttrSet("data.github_organization_team_sync_groups.test", "groups.0.group_id"),
					resource.TestCheckResourceAttrSet("data.github_organization_team_sync_groups.test", "groups.0.group_name"),
				),
			},
		},
	})
}

func testAccCheckGithubOrganizationTeamSyncGroupsDataSourceConfig() string {
	return `data "github_organization_team_sync_groups" "test" {}`
}
