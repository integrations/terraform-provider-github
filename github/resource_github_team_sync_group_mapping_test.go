package github

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccGithubTeamSyncGroupMapping_basic(t *testing.T) {
	if isEnterprise != "true" {
		t.Skip("Skipping because `ENTERPRISE_ACCOUNT` is not set or set to false")
	}
	randString := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	teamName := fmt.Sprintf("tf-acc-test-%s", randString)
	groupName := "test_team_group"
	description := fmt.Sprintf("tf-group-description-%s", randString)
	rn := "github_team_sync_group_mapping.test_mapping"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGithubTeamSyncGroupMappingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubTeamSyncGroupMappingConfig(teamName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(rn, "group.#"),
					resource.TestCheckResourceAttr(rn, "group.0.name", groupName),
				),
			},
			{
				Config: testAccGithubTeamSyncGroupMappingAddGroupAndUpdateConfig(teamName, description),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(rn, "group.#"),
					resource.TestCheckResourceAttr(rn, "group.0.description", description),
					resource.TestCheckResourceAttr(rn, "group.1.description", description),
				),
			},
			{
				ResourceName:      rn,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccGithubTeamSyncGroupMapping_empty(t *testing.T) {
	if isEnterprise != "true" {
		t.Skip("Skipping because `ENTERPRISE_ACCOUNT` is not set or set to false")
	}
	randString := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	teamName := fmt.Sprintf("tf-acc-test-%s", randString)
	groupName := "test_team_group"
	rn := "github_team_sync_group_mapping.test_mapping"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGithubTeamSyncGroupMappingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubTeamSyncGroupMappingConfig(teamName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(rn, "group.#"),
					resource.TestCheckResourceAttr(rn, "group.0.name", groupName),
				),
			},
			{
				Config: testAccGithubTeamSyncGroupMappingEmptyConfig(teamName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(rn, "group.#", "0"),
				),
			},
			{
				ResourceName:      rn,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckGithubTeamSyncGroupMappingDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*Organization).client
	orgName := testAccProvider.Meta().(*Organization).name
	ctx := context.TODO()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "github_team_sync_group_mapping" {
			continue
		}
		slug := rs.Primary.Attributes["team_slug"]
		team, err := getGithubTeamBySlug(ctx, conn, orgName, slug)
		if err != nil {
			return err
		}
		groupList, _, err := conn.Teams.ListIDPGroupsForTeam(ctx, string(team.GetID()))
		if err != nil {
			return err
		}

		if groupList != nil {
			if len(groupList.Groups) > 0 {
				return fmt.Errorf("Team Sync Group Mapping still exists for team %s", team.GetName())
			}
		}

		return nil
	}
	return nil
}

func testAccGithubTeamSyncGroupMappingConfig(teamName string) string {
	return fmt.Sprintf(`
data "github_organization_team_sync_groups" "test_groups" {}

resource "github_team" "test_team" {
  name        = "%s"
  description = "team for acc-tests"
}

resource "github_team_sync_group_mapping" "test_mapping" {
  team_slug  = github_team.test_team.slug
  
  dynamic "group" {
    for_each = [for g in data.github_organization_team_sync_groups.test_groups.groups : g if g.name == "test_team_group"]
    content {
      group_id          = each.value.group_id
      group_name        = each.value.group_name
      group_description = each.value.group_description
    }
  } 
}
`, teamName)
}

func testAccGithubTeamSyncGroupMappingAddGroupAndUpdateConfig(teamName, description string) string {
	return fmt.Sprintf(`
data "github_organization_team_sync_groups" "test_groups" {}

resource "github_team" "test_team" {
  name        = "%s"
  description = "team for acc-tests"
}

resource "github_team_sync_group_mapping" "test_mapping" {
  team_slug  = github_team.test_team.slug
  
  dynamic "group" {
    for_each = [for g in data.github_organization_team_sync_groups.test_groups.groups : g if "test_team_group" in g.name]
    content {
      group_id          = each.value.group_id
      group_name        = each.value.group_name
      group_description = "%s"
    }
  } 
}
`, teamName, description)
}

func testAccGithubTeamSyncGroupMappingEmptyConfig(teamName string) string {
	return fmt.Sprintf(`
data "github_organization_team_sync_groups" "test_groups" {}

resource "github_team" "test_team" {
  name        = "%s"
  description = "team for acc-tests"
}

resource "github_team_sync_group_mapping" "test_mapping" {
  team_slug  = github_team.test_team.slug
  
  group = []
}
`, teamName)
}
