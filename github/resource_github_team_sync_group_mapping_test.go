package github

import (
	"context"
	"fmt"
	"regexp"
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
				Config:      testAccGithubTeamSyncGroupMappingConfig(teamName),
				ExpectError: regexp.MustCompile(`Not Found`),
			},
			{
				ResourceName:      rn,
				ImportState:       true,
				ImportStateVerify: true,
				ExpectError:       regexp.MustCompile(fmt.Sprintf("couldn't be found: %s", rn)),
			},
			{
				Config:      testAccGithubTeamSyncGroupMappingAddGroupAndUpdateConfig(teamName, description),
				ExpectError: regexp.MustCompile(`Not Found`),
			},
			{
				ResourceName:      rn,
				ImportState:       true,
				ImportStateVerify: true,
				ExpectError:       regexp.MustCompile(fmt.Sprintf("couldn't be found: %s", rn)),
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
				Config:      testAccGithubTeamSyncGroupMappingConfig(teamName),
				ExpectError: regexp.MustCompile(`Not Found`),
			},
			{
				ResourceName:      rn,
				ImportState:       true,
				ImportStateVerify: true,
				ExpectError:       regexp.MustCompile(fmt.Sprintf("couldn't be found: %s", rn)),
			},
			{
				Config:      testAccGithubTeamSyncGroupMappingEmptyConfig(teamName),
				ExpectError: regexp.MustCompile(`Not Found`),
			},
			{
				ResourceName:      rn,
				ImportState:       true,
				ImportStateVerify: true,
				ExpectError:       regexp.MustCompile(fmt.Sprintf("couldn't be found: %s", rn)),
			},
		},
	})
}

func testAccCheckGithubTeamSyncGroupMappingDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*Organization).v3client
	orgName := testAccProvider.Meta().(*Organization).name
	ctx := context.TODO()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "github_team_sync_group_mapping" {
			continue
		}
		slug := rs.Primary.Attributes["team_slug"]
		groupList, _, err := conn.Teams.ListIDPGroupsForTeamBySlug(ctx, orgName, slug)
		if err != nil {
			return err
		}

		if groupList != nil {
			if len(groupList.Groups) > 0 {
				return fmt.Errorf("Team Sync Group Mapping still exists for team slug %s", slug)
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
    for_each = [for g in data.github_organization_team_sync_groups.test_groups.groups : g if length(regexall("^acctest-github-provider", g.group_name)) > 0]
    content {
      group_id          = group.value.group_id
      group_name        = group.value.group_name
      group_description = group.value.group_description
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
    for_each = [for g in data.github_organization_team_sync_groups.test_groups.groups : g if length(regexall("^acctest-github-provider", g.group_name)) > 0]
    content {
      group_id          = group.value.group_id
      group_name        = group.value.group_name
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
}
`, teamName)
}
