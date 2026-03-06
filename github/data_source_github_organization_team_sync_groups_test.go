package github

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubOrganizationTeamSyncGroupsDataSource_existing(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { skipUnlessMode(t, enterprise) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: `data "github_organization_team_sync_groups" "test" {}`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("data.github_organization_team_sync_groups.test", tfjsonpath.New("groups"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("data.github_organization_team_sync_groups.test", tfjsonpath.New("groups").AtSliceIndex(0).AtMapKey("group_id"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("data.github_organization_team_sync_groups.test", tfjsonpath.New("groups").AtSliceIndex(0).AtMapKey("group_name"), knownvalue.NotNull()),
				},
			},
			{
				Config: `data "github_organization_team_sync_groups" "test" { prefix_filter = "nonexistent_prefix_" }`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("data.github_organization_team_sync_groups.test", tfjsonpath.New("prefix_filter"), knownvalue.StringExact("nonexistent_prefix_")),
					statecheck.ExpectKnownValue("data.github_organization_team_sync_groups.test", tfjsonpath.New("groups"), knownvalue.ListSizeExact(0)),
				},
			},
		},
	})
}
