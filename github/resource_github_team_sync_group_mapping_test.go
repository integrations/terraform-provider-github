package github

import (
	"context"
	"errors"
	"fmt"
	"strconv"
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
					testAccCheckGithubTeamSyncGroupMappingMeta(rn),
				),
			},
			{
				Config: testAccGithubTeamSyncGroupMappingAddGroupAndUpdateConfig(teamName, description),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubTeamSyncGroupMappingDescriptionUpdateMeta(rn, description),
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
	rn := "github_team_sync_group_mapping.test_mapping"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGithubTeamSyncGroupMappingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubTeamSyncGroupMappingConfig(teamName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubTeamSyncGroupMappingMeta(rn),
				),
			},
			{
				Config: testAccGithubTeamSyncGroupMappingEmptyConfig(teamName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(rn, "groups.#", "0"),
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

func testAccCheckGithubTeamSyncGroupMappingMeta(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, groupCount, err := getResourceStateAndCount(s, n)
		if err != nil {
			return err
		}

		for i := 0; i < *groupCount; i++ {
			idx := "groups." + strconv.Itoa(i)
			if v, ok := rs.Primary.Attributes[idx+".group_id"]; !ok || v == "" {
				return fmt.Errorf("group %v is missing group_id", i)
			}
			if v, ok := rs.Primary.Attributes[idx+".group_name"]; !ok || v == "" {
				return fmt.Errorf("group %v is missing group_name", i)
			}
			if v, ok := rs.Primary.Attributes[idx+".group_description"]; !ok || v == "" {
				return fmt.Errorf("group %v is missing group_description", i)
			}
		}

		return nil
	}
}

func testAccCheckGithubTeamSyncGroupMappingDescriptionUpdateMeta(n, description string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, groupCount, err := getResourceStateAndCount(s, n)
		if err != nil {
			return err
		}

		for i := 0; i < *groupCount; i++ {
			idx := "groups." + strconv.Itoa(i)
			if v, ok := rs.Primary.Attributes[idx+".group_description"]; !ok || v != description {
				return fmt.Errorf("group %v group_description expected %s, actual %s", i, description, v)
			}
		}

		return nil
	}
}

func getResourceStateAndCount(s *terraform.State, rn string) (*terraform.ResourceState, *int, error) {
	rs, ok := s.RootModule().Resources[rn]
	if !ok {
		return nil, nil, fmt.Errorf("Can't find team-sync group-mappings resource: %s", rn)
	}

	groupCountStr, ok := rs.Primary.Attributes["groups.#"]
	if !ok {
		return rs, nil, errors.New("can't find 'groups' attribute")
	}

	groupCount, err := strconv.Atoi(groupCountStr)
	if err != nil {
		return rs, nil, errors.New("failed to read number of valid groups")
	}
	if groupCount < 1 {
		return rs, &groupCount, fmt.Errorf("expected at least 1 valid group, received %d, this is most likely a bug",
			groupCount)
	}
	return rs, &groupCount, nil
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
    for_each = [for g in data.github_organization_team_sync_groups.test_groups.groups : g if g.group_name == "test_team_group"]
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
    for_each = data.github_organization_team_sync_groups.test_groups.groups
    content {
      group_id          = group.value["group_id"]
      group_name        = group.value["group_name"]
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
  groups = []
}
`, teamName)
}
