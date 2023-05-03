package github

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccGithubActionsOrganizationRequiredWorkflowsDataSource(t *testing.T) {

	t.Run("creates and reads an org required workflow", func(t *testing.T) {
		fileName := "required-workflow.yml"
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

		config := fmt.Sprintf(`
			resource "github_repository" "test_repo" {
				name = "tf-acc-test-required-workflow-%[1]s"
				visibility = "private"
				auto_init = true
			}

			resource "github_actions_repository_access_level" "test" {
				access_level = "organization"
				repository   = github_repository.test_repo.name
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

			data "github_actions_organization_required_workflows" "test" {
				depends_on = [github_actions_required_workflow.test]
			}
		`, randomID, fileName)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"data.github_actions_organization_required_workflows.test", "total_count",
				"1",
			),
			resource.TestCheckResourceAttr("data.github_actions_organization_required_workflows.test", "required_workflows.#", "1"),
			resource.TestCheckResourceAttr("data.github_actions_organization_required_workflows.test", "required_workflows.0.name", "Required Workflow"),
			resource.TestCheckResourceAttr("data.github_actions_organization_required_workflows.test", "required_workflows.0.repository", fmt.Sprintf("tf-acc-test-required-workflow-%[1]s", randomID)),
			resource.TestCheckResourceAttrSet("data.github_actions_organization_required_workflows.test", "required_workflows.0.created_at"),
			resource.TestCheckResourceAttrSet("data.github_actions_organization_required_workflows.test", "required_workflows.0.updated_at"),
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
}
