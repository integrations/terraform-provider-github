package github

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubIssueLabel(t *testing.T) {
	t.Run("creates and updates labels without error", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-issue-label-%s", testResourcePrefix, randomID)
		description := "label_description"
		updatedDescription := "updated_label_description"

		config := fmt.Sprintf(`

			resource "github_repository" "test" {
			  name = "%s"
				auto_init = true
			}

			resource "github_issue_label" "test" {
			  repository  = github_repository.test.name
			  name        = "foo"
			  color       = "000000"
			  description = "%s"
			}
		`, repoName, description)

		checks := map[string]resource.TestCheckFunc{
			"before": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_issue_label.test", "description",
					description,
				),
			),
			"after": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_issue_label.test", "description",
					updatedDescription,
				),
			),
		}

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  checks["before"],
				},
				{
					Config: strings.Replace(config,
						description,
						updatedDescription, 1),
					Check: checks["after"],
				},
			},
		})
	})

	t.Run("can delete labels from archived repositories without error", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-issue-label-arch-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name = "%s"
				auto_init = true
			}

			resource "github_issue_label" "test" {
				repository = github_repository.test.name
				name = "archived-test-label"
				color = "ff0000"
				description = "Test label for archived repo"
			}
		`, repoName)

		archivedConfig := strings.Replace(config,
			`auto_init = true`,
			`auto_init = true
				archived = true`, 1)

		resource.Test(t, resource.TestCase{
			PreCheck:  func() { skipUnauthenticated(t) },
			Providers: testAccProviders,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr(
							"github_issue_label.test", "name",
							"archived-test-label",
						),
					),
				},
				{
					Config: archivedConfig,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr(
							"github_repository.test", "archived",
							"true",
						),
					),
				},
				// This step should succeed - the label should be removed from state
				// without trying to actually delete it from the archived repo
				{
					Config: fmt.Sprintf(`
							resource "github_repository" "test" {
								name = "%s"
								auto_init = true
								archived = true
							}
						`, repoName),
				},
			},
		})
	})
}
