package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubIssueLabels(t *testing.T) {
	t.Run("authoritatively overtakes existing labels", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		empty := []map[string]any{}

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					// 0. Check if some labels already exist (indicated by non-empty plan)
					{
						Config:             testAccGithubIssueLabelsConfig(randomID, empty),
						ExpectNonEmptyPlan: true,
					},
					// 1. Check if all the labels are destroyed when the resource is added
					{
						Config: testAccGithubIssueLabelsConfig(randomID, empty),
						Check:  resource.TestCheckResourceAttr("github_issue_labels.test", "label.#", "0"),
					},
					// 2. Check if a label can be created
					{
						Config: testAccGithubIssueLabelsConfig(randomID, append(empty, map[string]any{
							"name":        "foo",
							"color":       "000000",
							"description": "foo",
						})),
						Check: resource.TestCheckResourceAttr("github_issue_labels.test", "label.#", "1"),
					},
					// 3. Check if a label can be recreated
					{
						Config: testAccGithubIssueLabelsConfig(randomID, append(empty, map[string]any{
							"name":        "Foo",
							"color":       "000000",
							"description": "foo",
						})),
						Check: resource.TestCheckResourceAttr("github_issue_labels.test", "label.#", "1"),
					},
					// 4. Check if multiple labels can be created
					{
						Config: testAccGithubIssueLabelsConfig(randomID, append(empty,
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
						Config: testAccGithubIssueLabelsConfig(randomID, nil),
					},
					// 6. Check if labels were actually destroyed
					{
						Config: testAccGithubIssueLabelsConfig(randomID, empty),
						Check:  resource.TestCheckResourceAttr("github_issue_labels.test", "label.#", "0"),
					},
				},
			})
		}

		t.Run("with an anonymous account", func(t *testing.T) {
			t.Skip("anonymous account not supported for this operation")
		})

		t.Run("with an individual account", func(t *testing.T) {
			testCase(t, individual)
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})
	})
}

func testAccGithubIssueLabelsConfig(randomId string, labels []map[string]any) string {
	resource := ""
	if labels != nil {
		dynamic := ""
		for _, label := range labels {
			dynamic += fmt.Sprintf(`
				label {
					name = "%s"
					color = "%s"
					description = "%s"
				}
			`, label["name"], label["color"], label["description"])
		}

		resource = fmt.Sprintf(`
			resource "github_issue_labels" "test" {
				repository = github_repository.test.name

				%s
			}
		`, dynamic)
	}

	return fmt.Sprintf(`
		resource "github_repository" "test" {
			name = "tf-acc-test-%s"
			auto_init = true
		}

		%s
	`, randomId, resource)
}
