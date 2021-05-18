package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccGithubTeamRepositorty(t *testing.T) {

	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

	t.Run("Get Repositories By Teams", func(t *testing.T) {

		config := fmt.Sprintf(`
			resource "github_repository" "repo-test" {
				name = "tf-acc-repo-%s"
				auto_init = true
			
		  	}

			resource "github_team" "team-test" {
				name = "tf-acc-test-team01"
			}

			resource "github_team_repository" "team-repo-test" {
				repository = "${github_repository.repo-test.id}"
				team_id = "${github_team.team-test.id}"
			}

			data "github_team" "example" {
				slug = "team-test-01"
			}

   		    output "team_repository_name" {
				value = data.github_team.example.repositories.0
			}
			
			output "team_repository_numbers" {
				value = data.github_team.example.repositories.#
			}
		`, randomID)

		check := resource.ComposeAggregateTestCheckFunc(
			resource.TestCheckResourceAttrSet("data.github_team.example", "name"),
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
