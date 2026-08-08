package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccGithubCopilotTeamSeatAssignment(t *testing.T) {
	t.Run("assigns and removes Copilot seats for a team", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		teamName := fmt.Sprintf("%scopilot-team-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
resource "github_team" "test" {
  name = "%s"
}

resource "github_copilot_team_seat_assignment" "test" {
  team = github_team.test.slug
}
`, teamName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessCopilotSeatManagementEnabled(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttrSet("github_copilot_team_seat_assignment.test", "team"),
					),
				},
				{
					ResourceName:      "github_copilot_team_seat_assignment.test",
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})
}
