package github

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccGithubRepositoryProject(t *testing.T) {
	t.Skip("Skipping test as the GitHub API no longer supports classic projects")

	t.Run("creates a repository project", func(t *testing.T) {
		randomID := acctest.RandString(5)
		repoName := fmt.Sprintf("%srepo-project-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`

			resource "github_repository" "test" {
			  name = "%s"
				has_projects = true
			}

			resource "github_repository_project" "test" {
			  name       = "test"
			  repository = github_repository.test.name
			  body       = "this is a test project"
			}
		`, repoName)

		check := resource.ComposeTestCheckFunc(
			resource.TestMatchResourceAttr(
				"github_repository_project.test", "url",
				regexp.MustCompile(repoName+"/projects/1"),
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
