package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubIssue(t *testing.T) {
	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

	t.Run("creates an issue without error", func(t *testing.T) {
		title := "issue_title"
		body := "issue_body"
		labels := "\"bug\", \"enhancement\""
		updatedTitle := "updated_issue_title"
		updatedBody := "update_issue_body"
		updatedLabels := "\"documentation\""

		issueHCL := `
			resource "github_repository" "test" {
			  name = "tf-acc-test-%s"
				auto_init  = true
                has_issues = true
			}

			resource "github_repository_milestone" "test" {
				owner = split("/", "${github_repository.test.full_name}")[0]
				repository = github_repository.test.name
		    	title = "v1.0.0"
				description = "General Availability"
		    	due_date = "2022-11-22"
		    	state = "open"
			}

			resource "github_issue" "test" {
			  repository       = github_repository.test.name
			  title            = "%s"
			  body             = "%s"
			  labels           = [%s]
 			  assignees        = ["%s"]
			  milestone_number = github_repository_milestone.test.number
			}
		`
		config := fmt.Sprintf(issueHCL, randomID, title, body, labels, testOwnerFunc())

		checks := map[string]resource.TestCheckFunc{
			"before": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_issue.test", "title",
					title,
				),
				resource.TestCheckResourceAttr(
					"github_issue.test", "body",
					body,
				),
				resource.TestCheckResourceAttr(
					"github_issue.test", "labels.#",
					"2",
				),
				func(state *terraform.State) error {
					issue := state.RootModule().Resources["github_issue.test"].Primary
					issueMilestone := issue.Attributes["milestone_number"]

					milestone := state.RootModule().Resources["github_repository_milestone.test"].Primary
					milestoneNumber := milestone.Attributes["number"]
					if issueMilestone != milestoneNumber {
						return fmt.Errorf("issue milestone number %s not the same as repository milestone number %s",
							issueMilestone, milestoneNumber)
					}
					return nil
				},
			),
			"after": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_issue.test", "title",
					updatedTitle,
				),
				resource.TestCheckResourceAttr(
					"github_issue.test", "body",
					updatedBody,
				), resource.TestCheckResourceAttr(
					"github_issue.test", "labels.#",
					"1",
				), resource.TestCheckResourceAttr(
					"github_issue.test", "assignees.#",
					"1",
				),
			),
		}

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: config,
						Check:  checks["before"],
					},
					{
						Config: fmt.Sprintf(issueHCL, randomID, updatedTitle, updatedBody, updatedLabels, testOwnerFunc()),
						Check:  checks["after"],
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

	t.Run("imports a issue without error", func(t *testing.T) {
		config := fmt.Sprintf(`
					resource "github_repository" "test" {
					  name       = "tf-acc-test-%s"
					  has_issues = true
					}

					resource "github_issue" "test" {
					  repository       = github_repository.test.name
					  title            = github_repository.test.name
					}
		    `, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet("github_issue.test", "title"),
			resource.TestCheckResourceAttr("github_issue.test", "title", fmt.Sprintf(`tf-acc-test-%s`, randomID)),
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
					{
						ResourceName:      "github_issue.test",
						ImportState:       true,
						ImportStateVerify: true,
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
