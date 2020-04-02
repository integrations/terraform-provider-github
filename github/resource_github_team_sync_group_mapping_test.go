package github

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/go-github/v29/github"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccGithubTeamSyncGroupMapping_basic(t *testing.T) {
	if isEnterprise != "true" {
		t.Skip("Skipping because `ENTERPRISE_ACCOUNT` is not set or set to false")
	}
	var groupList github.IDPGroupList
	randString := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	teamName := fmt.Sprintf("tf-acc-test-%s", randString)
	description := fmt.Sprintf("tf-group-description-%s", randString)
	rn := "github_team_sync_group_mapping.test_team_group_mapping"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGithubTeamSyncGroupMappingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubTeamSyncGroupMappingConfig(teamName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubTeamSyncGroupMappingExists(rn, &groupList),
					testAccCheckGithubTeamSyncGroupMappingNum(1, &groupList),
				),
			},
			{
				Config: testAccGithubTeamSyncGroupMappingAddGroupAndUpdateConfig(teamName, description),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubTeamSyncGroupMappingExists(rn, &groupList),
					testAccCheckGithubTeamSyncGroupMappingDidIncreaseCount(&groupList),
					testAccCheckGithubTeamSyncGroupMappingUpdatedDescription(description, &groupList),
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
	var groupList github.IDPGroupList
	randString := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	teamName := fmt.Sprintf("tf-acc-test-%s", randString)
	rn := "github_team_sync_group_mapping.test_team_group_mapping"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGithubTeamSyncGroupMappingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubTeamSyncGroupMappingConfig(teamName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubTeamSyncGroupMappingExists(rn, &groupList),
					testAccCheckGithubTeamSyncGroupMappingNum(1, &groupList),
				),
			},
			{
				Config: testAccGithubTeamSyncGroupMappingEmptyConfig(teamName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubTeamSyncGroupMappingExists(rn, &groupList),
					testAccCheckGithubTeamSyncGroupMappingListEmpty(&groupList),
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

func testAccCheckGithubTeamSyncGroupMappingExists(n string, idpGroupList *github.IDPGroupList) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		client := testAccProvider.Meta().(*Organization).client
		orgName := testAccProvider.Meta().(*Organization).name
		slug := rs.Primary.Attributes["team_slug"]

		team, err := getGithubTeamBySlug(context.TODO(), client, orgName, slug)
		if err != nil {
			return err
		}
		found, _, err := client.Teams.ListIDPGroupsForTeam(context.TODO(), string(team.GetID()))

		if found == nil {
			return fmt.Errorf("Team-Sync Mapping not found in team '%s'", team.GetName())
		}

		*idpGroupList = *found

		return nil
	}
}

func testAccCheckGithubTeamSyncGroupMappingNum(v int, idpGroupList *github.IDPGroupList) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if len(idpGroupList.Groups) != v {
			return fmt.Errorf("should number of group connections to team set to %d", v)
		}

		return nil
	}
}

func testAccCheckGithubTeamSyncGroupMappingDidIncreaseCount(idpGroupList *github.IDPGroupList) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if len(idpGroupList.Groups) <= 1 {
			return fmt.Errorf("should number of group connections to team set to more than 1")
		}

		return nil
	}
}

func testAccCheckGithubTeamSyncGroupMappingUpdatedDescription(v string, idpGroupList *github.IDPGroupList) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		for _, group := range idpGroupList.Groups {
			if group.GetGroupDescription() != v {
				return fmt.Errorf("should description for group set to %s", v)
			}
		}

		return nil
	}
}

func testAccCheckGithubTeamSyncGroupMappingListEmpty(idpGroupList *github.IDPGroupList) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if len(idpGroupList.Groups) != 0 {
			return fmt.Errorf("should groups for team set to empty array")
		}

		return nil
	}
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

resource "github_team_sync_group_mapping" "test_team_group_mapping" {
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

resource "github_team_sync_group_mapping" "test_team_group_mapping" {
  team_slug  = github_team.test_team.slug
  
  dynamic "group" {
    for_each = data.test_groups.groups
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

resource "github_team_sync_group_mapping" "test_team_group_mapping" {
  team_slug  = github_team.test_team.slug
  
  group = []
}
`, teamName)
}
