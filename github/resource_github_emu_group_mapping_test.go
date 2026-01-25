package github

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccGithubEMUGroupMapping(t *testing.T) {
	groupID := testAccConf.testEnterpriseEMUGroupId
	if groupID == 0 {
		t.Skip("Skipping EMU group mapping tests because testEnterpriseEMUGroupId is not set")
	}
	t.Run("creates and manages EMU group mapping", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		teamName := fmt.Sprintf("%steam-emu-%s", testResourcePrefix, randomID)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessMode(t, enterprise) },
			ProviderFactories: providerFactories,
			CheckDestroy:      testAccCheckGithubEMUGroupMappingDestroy,
			Steps: []resource.TestStep{
				{
					Config: testAccGithubEMUGroupMappingConfig(teamName, groupID),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttrSet("github_emu_group_mapping.test", "group_id"),
						resource.TestCheckResourceAttr("github_emu_group_mapping.test", "team_slug", teamName),
						resource.TestCheckResourceAttrSet("github_emu_group_mapping.test", "etag"),
					),
				},
			},
		})
	})

	t.Run("imports EMU group mapping", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		teamName := fmt.Sprintf("%steam-emu-%s", testResourcePrefix, randomID)
		rn := "github_emu_group_mapping.test"

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessMode(t, enterprise) },
			ProviderFactories: providerFactories,
			CheckDestroy:      testAccCheckGithubEMUGroupMappingDestroy,
			Steps: []resource.TestStep{
				{
					Config: testAccGithubEMUGroupMappingConfig(teamName, groupID),
				},
				{
					ResourceName:      rn,
					ImportState:       true,
					ImportStateIdFunc: testAccGithubEMUGroupMappingImportStateIdFunc(rn),
					ImportStateVerify: true,
				},
			},
		})
	})
	t.Run("imports EMU group mapping with multiple teams", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		teamName := fmt.Sprintf("%steam1-emu-%s", testResourcePrefix, randomID)
		teamName2 := fmt.Sprintf("%steam2-emu-%s", testResourcePrefix, randomID)
		rn := "github_emu_group_mapping.test1"
		config := fmt.Sprintf(`
	resource "github_team" "test1" {
		name        = "%s"
		description = "EMU group mapping test team 1"
	}
	resource "github_team" "test2" {
		name        = "%s"
		description = "EMU group mapping test team 2"
	}
	resource "github_emu_group_mapping" "test1" {
		team_slug = github_team.test1.slug
		group_id  = %d
	}
	resource "github_emu_group_mapping" "test2" {
		team_slug = github_team.test2.slug
		group_id  = %[3]d
	}
`, teamName, teamName2, groupID)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessMode(t, enterprise) },
			ProviderFactories: providerFactories,
			CheckDestroy:      testAccCheckGithubEMUGroupMappingDestroy,
			Steps: []resource.TestStep{
				{
					Config: config,
				},
				{
					ResourceName:      rn,
					ImportState:       true,
					ImportStateIdFunc: testAccGithubEMUGroupMappingImportStateIdFunc(rn),
					ImportStateVerify: true,
				},
			},
		})
	})

	t.Run("handles team slug update by recreating", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		teamName1 := fmt.Sprintf("%steam-emu-%s", testResourcePrefix, randomID)
		teamName2 := fmt.Sprintf("%s-upd", teamName1)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessMode(t, enterprise) },
			ProviderFactories: providerFactories,
			CheckDestroy:      testAccCheckGithubEMUGroupMappingDestroy,
			Steps: []resource.TestStep{
				{
					Config: testAccGithubEMUGroupMappingConfig(teamName1, groupID),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_emu_group_mapping.test", "team_slug", teamName1),
					),
				},
				{
					Config: testAccGithubEMUGroupMappingConfig(teamName2, groupID),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_emu_group_mapping.test", "team_slug", teamName2),
					),
				},
			},
		})
	})

	t.Run("forces new when switching to different team", func(t *testing.T) {
		t.Skip("Skipping this test because we don't have terraform-plugin-testing available yet.")
		randomID := acctest.RandString(5)
		teamName1 := fmt.Sprintf("%semu1-%s", testResourcePrefix, randomID)
		teamName2 := fmt.Sprintf("%semu2-%s", testResourcePrefix, randomID)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
			ProviderFactories: providerFactories,
			CheckDestroy:      testAccCheckGithubEMUGroupMappingDestroy,
			Steps: []resource.TestStep{
				{
					Config: testAccGithubEMUGroupMappingTwoTeamsConfig(teamName1, teamName2, groupID, "test1"),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_emu_group_mapping.test", "team_slug", teamName1),
					),
				},
				{
					Config: testAccGithubEMUGroupMappingTwoTeamsConfig(teamName1, teamName2, groupID, "test2"),
					// ConfigPlanChecks: resource.ConfigPlanChecks{
					// 	PreApply: []plancheck.PlanCheck{
					// 		plancheckExpectKnownValues("github_emu_group_mapping.test", "team_slug", teamName2),
					// 		plancheck.ExpectResourceAction("github_emu_group_mapping.test", plancheck.ResourceActionDestroyBeforeCreate), // Verify that ForceNew is triggered
					// 	},
					// },
				},
			},
		})
	})
}

func testAccCheckGithubEMUGroupMappingDestroy(s *terraform.State) error {
	meta, err := getTestMeta()
	if err != nil {
		return err
	}
	conn := meta.v3client
	orgName := meta.name
	ctx := context.Background()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "github_emu_group_mapping" {
			continue
		}

		groupIDStr := rs.Primary.Attributes["group_id"]
		groupID, err := strconv.ParseInt(groupIDStr, 10, 64)
		if err != nil {
			return fmt.Errorf("failed to parse group_id %q: %w", groupIDStr, err)
		}

		group, resp, err := conn.Teams.GetExternalGroup(ctx, orgName, groupID)
		if err == nil {
			if group != nil && len(group.Teams) > 0 {
				return fmt.Errorf("EMU group mapping still exists for group_id %d", groupID)
			}
		}
		if resp != nil && resp.StatusCode != 404 {
			return err
		}
	}
	return nil
}

func testAccGithubEMUGroupMappingImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("not found: %s", resourceName)
		}
		return buildTwoPartID(rs.Primary.Attributes["group_id"], rs.Primary.Attributes["team_slug"]), nil
	}
}

func testAccGithubEMUGroupMappingConfig(teamName string, groupID int) string {
	return fmt.Sprintf(`
	resource "github_team" "test" {
		name        = "%s"
		description = "EMU group mapping test team"
	}
	resource "github_emu_group_mapping" "test" {
		team_slug = github_team.test.slug
		group_id  = %d
	}
	`, teamName, groupID)
}

func testAccGithubEMUGroupMappingTwoTeamsConfig(teamName1, teamName2 string, groupID int, useTeam string) string {
	return fmt.Sprintf(`
	resource "github_team" "test1" {
		name        = "%s"
		description = "EMU group mapping test team 1"
	}
	resource "github_team" "test2" {
		name        = "%s"
		description = "EMU group mapping test team 2"
	}
	resource "github_emu_group_mapping" "test" {
		team_slug = github_team.%s.slug
		group_id  = %d
	}
	`, teamName1, teamName2, useTeam, groupID)
}
