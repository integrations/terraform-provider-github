package github

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/google/go-github/v89/github"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubTeamExternalGroupsDataSource(t *testing.T) {
	t.Run("queries_non_existent_team", func(t *testing.T) {
		config := `
			data "github_team_external_groups" "test" {
				slug = "non-existing-team-slug"
			}
		`

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEMUEnterprise(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      config,
					ExpectError: regexp.MustCompile(`Not Found`),
				},
			},
		})
	})

	t.Run("queries_team_without_external_groups", func(t *testing.T) {
		team := mustCreateTestTeam(t)
		config := fmt.Sprintf(`
			data "github_team_external_groups" "test" {
				slug = "%s"
			}
		`, team.GetSlug())

		mustRemoveAllMembersFromTeam(t, team)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEMUEnterprise(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_team_external_groups.test", tfjsonpath.New("external_groups"), knownvalue.ListSizeExact(0)),
					},
				},
			},
		})
	})

	t.Run("queries_team_with_external_groups", func(t *testing.T) {
		if testAccConf.testExternalGroup1ID == 0 {
			t.Skip("Skipping as no external groups are configured for the test organization")
		}
		groupID := int64(testAccConf.testExternalGroup1ID)
		team := mustCreateTestTeam(t)
		mustConnectTeamToExternalGroup(t, team, &github.ExternalGroup{
			GroupID: &groupID,
		})
		config := fmt.Sprintf(`
			data "github_team_external_groups" "test" {
				slug = "%s"
			}
		`, team.GetSlug())

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEMUEnterprise(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_team_external_groups.test", tfjsonpath.New("external_groups"), knownvalue.ListExact([]knownvalue.Check{
							knownvalue.MapPartial(map[string]knownvalue.Check{
								"group_id": knownvalue.Int64Exact(groupID),
							}),
						})),
					},
				},
			},
		})
	})
}
