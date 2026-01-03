package github

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubIssueLabels(t *testing.T) {
	t.Run("authoritatively overtakes existing labels", func(t *testing.T) {
		repoName := fmt.Sprintf("%srepo-issue-labels-%s", testResourcePrefix, acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum))
		empty := []map[string]any{}

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				// 0. Check if some labels already exist (indicated by non-empty plan)
				{
					Config:             testAccGithubIssueLabelsConfig(repoName, empty),
					PlanOnly:           true,
					ExpectNonEmptyPlan: true,
				},
				// 1. Check if all the labels are destroyed when the resource is added
				{
					Config: testAccGithubIssueLabelsConfig(repoName, empty),
					Check:  resource.TestCheckResourceAttr("github_issue_labels.test", "label.#", "0"),
				},
				// 2. Check if a label can be created
				{
					Config: testAccGithubIssueLabelsConfig(repoName, append(empty, map[string]any{
						"name":        "foo",
						"color":       "000000",
						"description": "foo",
					})),
					Check: resource.TestCheckResourceAttr("github_issue_labels.test", "label.#", "1"),
				},
				// 3. Check if a label can be recreated
				{
					Config: testAccGithubIssueLabelsConfig(repoName, append(empty, map[string]any{
						"name":        "Foo",
						"color":       "000000",
						"description": "foo",
					})),
					Check: resource.TestCheckResourceAttr("github_issue_labels.test", "label.#", "1"),
				},
				// 4. Check if multiple labels can be created
				{
					Config: testAccGithubIssueLabelsConfig(repoName, append(empty,
						map[string]any{
							"name":        "Foo",
							"color":       "000000",
							"description": "foo",
						},
						map[string]any{
							"name":        "bar",
							"color":       "000000",
							"description": "bar",
						}, map[string]any{
							"name":        "baz",
							"color":       "000000",
							"description": "baz",
						})),
					Check: resource.TestCheckResourceAttr("github_issue_labels.test", "label.#", "3"),
				},
				// 5. Check if labels can be destroyed
				{
					Config: testAccGithubIssueLabelsConfig(repoName, nil),
				},
				// 6. Check if labels were actually destroyed
				{
					Config: testAccGithubIssueLabelsConfig(repoName, empty),
					Check:  resource.TestCheckResourceAttr("github_issue_labels.test", "label.#", "0"),
				},
			},
		})
	})
}

func testAccGithubIssueLabelsConfig(repoName string, labels []map[string]any) string {
	resource := ""
	if labels != nil {
		var dynamic strings.Builder
		for _, label := range labels {
			dynamic.WriteString(fmt.Sprintf(`
				label {
					name = "%s"
					color = "%s"
					description = "%s"
				}
			`, label["name"], label["color"], label["description"]))
		}

		resource = fmt.Sprintf(`
			resource "github_issue_labels" "test" {
				repository = github_repository.test.name

				%s
			}
		`, dynamic.String())
	}

	return fmt.Sprintf(`
		resource "github_repository" "test" {
			name = "%s"
			auto_init = true
		}

		%s
	`, repoName, resource)
}

func TestAccGithubIssueLabelsArchived(t *testing.T) {
	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

	t.Run("can delete labels from archived repositories without error", func(t *testing.T) {
		repoName := fmt.Sprintf("%srepo-labels-arch-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name = "%s"
				auto_init = true
			}

			resource "github_issue_labels" "test" {
				repository = github_repository.test.name
				label {
					name = "archived-label-1"
					color = "ff0000"
					description = "First test label"
				}
				label {
					name = "archived-label-2"
					color = "00ff00"
					description = "Second test label"
				}
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
							"github_issue_labels.test", "label.#",
							"2",
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
				// This step should succeed - the labels should be removed from state
				// without trying to actually delete them from the archived repo
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
