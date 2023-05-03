package github

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccGithubActionsRequiredWorkflows(t *testing.T) {
	fileName := "required-workflow.yml"

	t.Run("creates a required workflow for all repositories", func(t *testing.T) {
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
		`, randomID, fileName)

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: config,
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

	t.Run("creates a required workflow with selected repositories", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("tf-acc-test-required-workflow-%[1]s", randomID)

		config := fmt.Sprintf(`
			resource "github_repository" "test_repo_1" {
				name = "%[1]s-1"
				auto_init = true
			}

			resource "github_repository" "test_repo_2" {
				name = "%[1]s-2"
			}

			resource "github_repository" "test_repo_3" {
				name = "%[1]s-3"
			}

			resource "github_repository_file" "required_workflow" {
			  repository          = github_repository.test_repo_1.name
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
				repository_id        = github_repository.test_repo_1.repo_id
				scope 				 = "selected"
				selected_repository_ids = [
						github_repository.test_repo_2.repo_id,
						github_repository.test_repo_3.repo_id
					]
				workflow_file_path	 = "%[2]s"
			}
		`, repoName, fileName)

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: config,
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

	t.Run("updates a required workflow to use a different workflow file", func(t *testing.T) {
		updatedFileName := "updated-required-workflow.yml"

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
		`, randomID, fileName)

		checks := map[string]resource.TestCheckFunc{
			"before": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_actions_required_workflow.test", "workflow_file_path",
					fileName,
				),
			),
			"after": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_actions_required_workflow.test", "workflow_file_path",
					updatedFileName,
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
							fileName,
							updatedFileName, 2),
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

	t.Run("deletes a required workflow", func(t *testing.T) {

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
		`, randomID, fileName)

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config:  config,
						Destroy: true,
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
