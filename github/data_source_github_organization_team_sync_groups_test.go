package github

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubOrganizationTeamSyncGroupsDataSource(t *testing.T) {
	t.Parallel()

	t.Run("all", func(t *testing.T) {
		t.Parallel()

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: `data "github_organization_team_sync_groups" "test" {}`,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_organization_team_sync_groups.test",
							tfjsonpath.New("groups"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_organization_team_sync_groups.test",
							tfjsonpath.New("groups").AtSliceIndex(0).AtMapKey("group_id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_organization_team_sync_groups.test",
							tfjsonpath.New("groups").AtSliceIndex(0).AtMapKey("group_name"), knownvalue.NotNull()),
					},
				},
			},
		})
	})

	t.Run("filtered", func(t *testing.T) {
		t.Parallel()

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: `data "github_organization_team_sync_groups" "test" { prefix_filter = "acctest-github-provider" }`,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_organization_team_sync_groups.test",
							tfjsonpath.New("prefix_filter"), knownvalue.StringExact("acctest-github-provider")),
						statecheck.ExpectKnownValue("data.github_organization_team_sync_groups.test",
							tfjsonpath.New("groups"), knownvalue.NotNull()),
					},
				},
				{
					Config: `data "github_organization_team_sync_groups" "test" { prefix_filter = "nonexistent_prefix_" }`,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_organization_team_sync_groups.test",
							tfjsonpath.New("prefix_filter"), knownvalue.StringExact("nonexistent_prefix_")),
						statecheck.ExpectKnownValue("data.github_organization_team_sync_groups.test",
							tfjsonpath.New("groups"), knownvalue.ListSizeExact(0)),
					},
				},
			},
		})
	})
}
