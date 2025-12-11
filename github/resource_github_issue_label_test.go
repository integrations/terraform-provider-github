package github

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubIssueLabel(t *testing.T) {
	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

	t.Run("creates and updates labels without error", func(t *testing.T) {
		description := "label_description"
		updatedDescription := "updated_label_description"

		config := fmt.Sprintf(`

			resource "github_repository" "test" {
			  name = "tf-acc-test-%s"
				auto_init = true
			}

			resource "github_issue_label" "test" {
			  repository  = github_repository.test.name
			  name        = "foo"
			  color       = "000000"
			  description = "%s"
			}
		`, randomID, description)

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
						Config: strings.Replace(config,
							description,
							updatedDescription, 1),
						Check: checks["after"],
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

	t.Run("can delete labels from archived repositories without error", func(t *testing.T) {
		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name = "tf-acc-test-archive-%s"
				auto_init = true
			}

			resource "github_issue_label" "test" {
				repository = github_repository.test.name
				name = "archived-test-label"
				color = "ff0000"
				description = "Test label for archived repo"
			}
		`, randomID)

		archivedConfig := strings.Replace(config,
			`auto_init = true`,
			`auto_init = true
				archived = true`, 1)

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
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
								name = "tf-acc-test-archive-%s"
								auto_init = true
								archived = true
							}
						`, randomID),
					},
				},
			})
		}

		t.Run("with an individual account", func(t *testing.T) {
			testCase(t, individual)
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})
	})
}
