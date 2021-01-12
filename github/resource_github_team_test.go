package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccGithubTeam(t *testing.T) {
	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

	t.Run("creates teams without error", func(t *testing.T) {

		config := fmt.Sprintf(`
			resource "github_team" "parent" {
			  name        = "tf-acc-parent-%s"
			  description = "Terraform acc test parent team"
			  privacy     = "closed"
			}

			resource "github_team" "child" {
			  name           = "tf-acc-child-%[1]s"
			  description    = "Terraform acc test child team"
			  privacy        = "closed"
			  parent_team_id = "${github_team.parent.id}"
			}
		`, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet("github_team.parent", "id"),
			resource.TestCheckResourceAttrSet("github_team.child", "id"),
		)

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: config,
						Check:  check,
					},
				},
			})
		}

		t.Run("with an anonymous account", func(t *testing.T) {
			t.Skip("anonymous account not supported for this operation")
		})

		t.Run("with an individual account", func(t *testing.T) {
			t.Skip("individual account not supported for this operation")
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})

	})
}