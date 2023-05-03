package github

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccGithubActionsRequiredWorkflowDataSource(t *testing.T) {
	t.Run("create and read a required workflow", func(t *testing.T) {
		fileName := "required-workflow.yml"
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

		config := fmt.Sprintf(`
			resource "github_repository" "test_repo" {
				name = "tf-acc-test-required-workflow-%[1]s"
				auto_init = true
			}

			resource "github_repository_file" "required_workflow" {
			  repository          = github_repository.test_repo.name
			  branch              = "main"
			  file                = "%[2]s"
			  content             = <<EOT
name: Required Workflow
on: pull_request
jobs:
  hello:
    runs-on: ubuntu-latest
    steps:
      - name: Hello World
        run: echo "Hello, world!"
EOT
			}

			resource "github_actions_required_workflow" "test" {
				depends_on = [github_repository_file.required_workflow]
				repository_id        = github_repository.test_repo.repo_id
				scope 				 = "all"
				workflow_file_path	 = "%[2]s"
			}

			data "github_actions_required_workflow" "test" {
				required_workflow_id = github_actions_required_workflow.test.required_workflow_id
			}
		`, randomID, fileName)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"data.github_actions_required_workflow.test", "name",
				"Required Workflow",
			),
			resource.TestCheckResourceAttrSet(
				"data.github_actions_required_workflow.test", "repository",
			),
			resource.TestCheckResourceAttr(
				"data.github_actions_required_workflow.test", "scope",
				"all",
			),
			resource.TestCheckResourceAttrSet(
				"data.github_actions_required_workflow.test", "created_at",
			),
			resource.TestCheckResourceAttrSet(
				"data.github_actions_required_workflow.test", "path",
			),
			resource.TestCheckResourceAttrSet(
				"data.github_actions_required_workflow.test", "state",
			),
			resource.TestCheckResourceAttrSet(
				"data.github_actions_required_workflow.test", "updated_at",
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
