package github

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/compare"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubEMUGroupMapping(t *testing.T) {
	groupID := testAccConf.testEnterpriseEMUGroupId
	if groupID == 0 {
		t.Skip("Skipping EMU group mapping tests because testEnterpriseEMUGroupId is not set")
	}
	t.Run("creates_without_error", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		teamName := fmt.Sprintf("%steam-emu-%s", testResourcePrefix, randomID)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEMUEnterprise(t) },
			ProviderFactories: providerFactories,
			CheckDestroy:      testAccCheckGithubEMUGroupMappingDestroy,
			Steps: []resource.TestStep{
				{
					Config: testAccGithubEMUGroupMappingConfig(teamName, groupID),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_emu_group_mapping.test", tfjsonpath.New("group_id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_emu_group_mapping.test", tfjsonpath.New("etag"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_emu_group_mapping.test", tfjsonpath.New("team_id"), knownvalue.NotNull()),
						statecheck.CompareValuePairs("github_emu_group_mapping.test", tfjsonpath.New("team_slug"), "github_team.test", tfjsonpath.New("slug"), compare.ValuesSame()),
					},
				},
			},
		})
	})

	t.Run("imports_without_error", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		teamName := fmt.Sprintf("%steam-emu-%s", testResourcePrefix, randomID)
		rn := "github_emu_group_mapping.test"

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEMUEnterprise(t) },
			ProviderFactories: providerFactories,
			CheckDestroy:      testAccCheckGithubEMUGroupMappingDestroy,
			Steps: []resource.TestStep{
				{
					Config: testAccGithubEMUGroupMappingConfig(teamName, groupID),
				},
				{
					ResourceName:            rn,
					ImportState:             true,
					ImportStateIdFunc:       testAccGithubEMUGroupMappingImportStateIdFunc(rn),
					ImportStateVerify:       true,
					ImportStateVerifyIgnore: []string{"etag"},
				},
			},
		})
	})
	t.Run("imports_multiple_teams_without_error", func(t *testing.T) {
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
			PreCheck:          func() { skipUnlessEMUEnterprise(t) },
			ProviderFactories: providerFactories,
			CheckDestroy:      testAccCheckGithubEMUGroupMappingDestroy,
			Steps: []resource.TestStep{
				{
					Config: config,
				},
				{
					ResourceName:            rn,
					ImportState:             true,
					ImportStateIdFunc:       testAccGithubEMUGroupMappingImportStateIdFunc(rn),
					ImportStateVerify:       true,
					ImportStateVerifyIgnore: []string{"etag"},
				},
			},
		})
	})

	t.Run("updates_team_slug_without_error", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		teamName1 := fmt.Sprintf("%steam-emu-%s", testResourcePrefix, randomID)
		teamName2 := fmt.Sprintf("%s-upd", teamName1)

		compareIDSame := statecheck.CompareValue(compare.ValuesSame())
		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEMUEnterprise(t) },
			ProviderFactories: providerFactories,
			CheckDestroy:      testAccCheckGithubEMUGroupMappingDestroy,
			Steps: []resource.TestStep{
				{
					Config: testAccGithubEMUGroupMappingConfig(teamName1, groupID),
					ConfigStateChecks: []statecheck.StateCheck{
						compareIDSame.AddStateValue("github_emu_group_mapping.test", tfjsonpath.New("id")),
						statecheck.ExpectKnownValue("github_emu_group_mapping.test", tfjsonpath.New("group_id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_emu_group_mapping.test", tfjsonpath.New("etag"), knownvalue.NotNull()),
						statecheck.CompareValuePairs("github_emu_group_mapping.test", tfjsonpath.New("team_slug"), "github_team.test", tfjsonpath.New("slug"), compare.ValuesSame()),
						statecheck.ExpectKnownValue("github_emu_group_mapping.test", tfjsonpath.New("team_id"), knownvalue.NotNull()),
					},
				},
				{
					Config: testAccGithubEMUGroupMappingConfig(teamName2, groupID),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectUnknownValue("github_emu_group_mapping.test", tfjsonpath.New("team_slug")),
							plancheck.ExpectResourceAction("github_emu_group_mapping.test", plancheck.ResourceActionUpdate),
						},
					},
					ConfigStateChecks: []statecheck.StateCheck{
						compareIDSame.AddStateValue("github_emu_group_mapping.test", tfjsonpath.New("id")),
						statecheck.CompareValuePairs("github_emu_group_mapping.test", tfjsonpath.New("team_slug"), "github_team.test", tfjsonpath.New("slug"), compare.ValuesSame()),
					},
				},
			},
		})
	})

	t.Run("recreates_when_switching_to_different_team_without_error", func(t *testing.T) {
		randomID := acctest.RandString(5)
		teamName1 := fmt.Sprintf("%semu1-%s", testResourcePrefix, randomID)
		teamName2 := fmt.Sprintf("%semu2-%s", testResourcePrefix, randomID)

		config := `
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
`

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEMUEnterprise(t) },
			ProviderFactories: providerFactories,
			CheckDestroy:      testAccCheckGithubEMUGroupMappingDestroy,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, teamName1, teamName2, "test1", groupID),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_emu_group_mapping.test", tfjsonpath.New("group_id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_emu_group_mapping.test", tfjsonpath.New("team_id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_emu_group_mapping.test", tfjsonpath.New("etag"), knownvalue.NotNull()),
						statecheck.CompareValuePairs("github_emu_group_mapping.test", tfjsonpath.New("team_slug"), "github_team.test1", tfjsonpath.New("slug"), compare.ValuesSame()),
					},
				},
				{
					Config: fmt.Sprintf(config, teamName1, teamName2, "test2", groupID),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectKnownValue("github_emu_group_mapping.test", tfjsonpath.New("team_slug"), knownvalue.StringExact(teamName2)),
							plancheck.ExpectResourceAction("github_emu_group_mapping.test", plancheck.ResourceActionReplace), // Verify that ForceNew is triggered
						},
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_emu_group_mapping.test", tfjsonpath.New("group_id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_emu_group_mapping.test", tfjsonpath.New("team_id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_emu_group_mapping.test", tfjsonpath.New("etag"), knownvalue.NotNull()),
						statecheck.CompareValuePairs("github_emu_group_mapping.test", tfjsonpath.New("team_slug"), "github_team.test2", tfjsonpath.New("slug"), compare.ValuesSame()),
					},
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
		resourceID, err := buildID(rs.Primary.Attributes["group_id"], rs.Primary.Attributes["team_slug"])
		if err != nil {
			return "", err
		}
		return resourceID, nil
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
