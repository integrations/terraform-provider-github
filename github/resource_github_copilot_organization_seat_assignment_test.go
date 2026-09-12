package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccGithubCopilotOrganizationSeatAssignment(t *testing.T) {
	if testAccConf.testOrgUser1 == "" {
		t.Skip("GH_TEST_ORG_USER1 not set")
	}

	username := testAccConf.testOrgUser1

	t.Run("assigns and removes a Copilot seat for a user", func(t *testing.T) {
		config := fmt.Sprintf(`
resource "github_copilot_organization_seat_assignment" "test" {
  username = "%s"
}
`, username)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessCopilotSeatManagementEnabled(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_copilot_organization_seat_assignment.test", "username", username),
					),
				},
				{
					ResourceName:      "github_copilot_organization_seat_assignment.test",
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})
}
