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
		t.Skip()
	}
	var team github.Team

	rn := "github_team.foo"
	randStringTeam := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	teamName := fmt.Sprintf("tf-acc-test-%s", randStringTeam)
	randStringGroup := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	groupID := acctest.RandString(5)
	groupName := fmt.Sprintf("tf-acc-test-group-%s", randStringGroup)
	description := "Terraform acc test group"
	updatedDescription := "Terraform acc test group - updated"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGithubTeamSyncGroupMappingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGithubTeamSyncGroupMappingConfig(teamName, groupID, groupName, description),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubTeamExists(rn, &team),
					resource.TestCheckResourceAttrSet(rn, "group"),
					resource.TestCheckResourceAttr(rn, "group.0.group_id", groupID),
					resource.TestCheckResourceAttr(rn, "group.0.group_name", groupName),
					resource.TestCheckResourceAttr(rn, "group.0.group_description", description),
				),
			},
			{
				Config: testAccGithubTeamSyncGroupMappingConfig(teamName, groupID, groupName, updatedDescription),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGithubTeamExists(rn, &team),
					resource.TestCheckResourceAttrSet(rn, "group"),
					resource.TestCheckResourceAttr(rn, "group.0.group_description", updatedDescription),
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
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "github_team_sync_group_mapping" {
			continue
		}
		id := rs.Primary.Attributes["team_id"]
		groupList, resp, err := conn.Teams.ListIDPGroupsForTeam(context.TODO(), id)
		if err == nil {
			if len(groupList.Groups) > 0 {
				return fmt.Errorf("Team Sync Group Mapping still exists for team %s", id)
			}
		}
		if resp.StatusCode != 404 {
			return err
		}
		return nil
	}
	return nil
}

func testAccGithubTeamSyncGroupMappingConfig(teamName, groupID, groupName, description string) string {
	return fmt.Sprintf(`
resource "github_team" "foo" {
  name        = "%s"
  description = "Terraform acc test group"
  privacy     = "secret"
}

resource "github_team_sync_group_mapping" "test" {
	retrieve_by = "slug"
	team_slug = "${github_team.foo.slug}"
	group {
		group_id = "%s"
		group_name = "%s"
		group_description = "%s"
	}
}
`, teamName, groupID, groupName, description)
}
