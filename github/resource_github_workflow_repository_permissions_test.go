package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubWorkflowRepositoryPermissions(t *testing.T) {
	t.Run("test setting of basic workflow repository permissions", func(t *testing.T) {
		defaultWorkflowPermissions := "read"
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-workflow-perms-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name        = "%[1]s"
				description = "Terraform acceptance tests %[1]s"
				topics		= ["terraform", "testing"]
			}

			resource "github_workflow_repository_permissions" "test" {
				default_workflow_permissions = "%s"
				repository = github_repository.test.name
			}
		`, repoName, defaultWorkflowPermissions)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_workflow_repository_permissions.test", "default_workflow_permissions", defaultWorkflowPermissions,
			),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:  func() { skipUnauthenticated(t) },
			Providers: testAccProviders,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  check,
				},
			},
		})
	})

	t.Run("imports entire set of github workflow repository permissions without error", func(t *testing.T) {
		defaultWorkflowPermissions := "read"
		canApprovePullRequestReviews := "true"

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-workflow-perms-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name        = "%[1]s"
				description = "Terraform acceptance tests %[1]s"
				topics		= ["terraform", "testing"]
			}

			resource "github_workflow_repository_permissions" "test" {
				default_workflow_permissions = "%s"
				can_approve_pull_request_reviews = %s
				repository = github_repository.test.name
			}
		`, repoName, defaultWorkflowPermissions, canApprovePullRequestReviews)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_workflow_repository_permissions.test", "default_workflow_permissions", defaultWorkflowPermissions,
			),
			resource.TestCheckResourceAttr(
				"github_workflow_repository_permissions.test", "can_approve_pull_request_reviews", canApprovePullRequestReviews,
			),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:  func() { skipUnauthenticated(t) },
			Providers: testAccProviders,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  check,
				},
				{
					ResourceName:      "github_workflow_repository_permissions.test",
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})
}
