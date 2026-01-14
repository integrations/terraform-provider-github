package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubRepositoryMilestone(t *testing.T) {
	t.Run("creates a repository milestone", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-milestone-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`

			resource "github_repository" "test" {
				name      = "%s"
			}

			resource "github_repository_milestone" "test" {
				owner = split("/", "${github_repository.test.full_name}")[0]
				repository = github_repository.test.name
		    title = "v1.0.0"
		    description = "General Availability"
		    due_date = "2020-11-22"
		    state = "closed"
			}

		`, repoName)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_repository_milestone.test", "state",
				"closed",
			),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  check,
				},
			},
		})
	})
}
