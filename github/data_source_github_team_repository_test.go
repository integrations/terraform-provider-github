package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccGithubTeamRepositories(t *testing.T) {

	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

	t.Run("Get Repositories By Teams", func(t *testing.T) {

		config := fmt.Sprintf(`

		resource "github_repository" "test" {
			name      = "tf-acc-test-%s"
			auto_init = true
		  }

		  resource "github_team" "test" {
			name = "tf-acc-test-%[1]s"
		  }
			
		  resource "github_team_repository" "test" {
			team_id    = "${github_team.test.id}"
			repository = "${github_repository.test.name}"
		  }

		  data "github_team" "example" {
			depends_on = ["github_repository.test", "github_team.test", "github_team_repository.test"]
			slug = github_team.test.slug
		  }
		`, randomID)

		check := resource.ComposeAggregateTestCheckFunc(
			resource.TestCheckResourceAttr("data.github_team.example", "repositories.#", "1"),
		)

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config:             config,
						Check:              check,
						ExpectNonEmptyPlan: true,
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
