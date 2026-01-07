package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubProjectItem(t *testing.T) {
	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

	t.Run("creates a project item using an issue", func(t *testing.T) {
		config := fmt.Sprintf(`

			resource "github_repository" "test" {
			  name = "tf-acc-test-%s"
				has_issues   = true
			}

			resource "github_issue" "test" {
			  repository = github_repository.test.name
			  title      = "Test issue title"
			  body       = "Test issue body"
			}

			resource "github_organization_project" "test" {
			  name = "tf-acc-%s"
			  body = "This is a test project."
			}

			resource "github_project_item" "test" {
			  project_number = github_organization_project.test.project_number
			  content_id     = github_issue.test.issue_id
			  content_type   = "Issue"
			}

		`, randomID, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet(
				"github_project_item.test", "item_id",
			),
			resource.TestCheckResourceAttrSet(
				"github_project_item.test", "node_id",
			),
			resource.TestCheckResourceAttr(
				"github_project_item.test", "content_type", "Issue",
			),
			resource.TestCheckResourceAttr(
				"github_project_item.test", "archived", "false",
			),
			func(s *terraform.State) error {
				item := s.RootModule().Resources["github_project_item.test"]
				issue := s.RootModule().Resources["github_issue.test"]

				itemContentID := item.Primary.Attributes["content_id"]
				issueID := issue.Primary.Attributes["issue_id"]
				if itemContentID != issueID {
					return fmt.Errorf("project item content_id %s not the same as issue id %s",
						itemContentID, issueID)
				}
				return nil
			},
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

	t.Run("creates an archived project item", func(t *testing.T) {
		config := fmt.Sprintf(`

			resource "github_repository" "test" {
			  name = "tf-acc-test-%s"
				has_issues   = true
			}

			resource "github_issue" "test" {
			  repository = github_repository.test.name
			  title      = "Test issue title"
			  body       = "Test issue body"
			}

			resource "github_organization_project" "test" {
			  name = "tf-acc-%s"
			  body = "This is a test project."
			}

			resource "github_project_item" "test" {
			  project_number = github_organization_project.test.project_number
			  content_id     = github_issue.test.issue_id
			  content_type   = "Issue"
			  archived       = true
			}

		`, randomID, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_project_item.test", "archived", "true",
			),
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

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})
	})
}
