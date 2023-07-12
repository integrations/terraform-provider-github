package github

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccGithubActionsDefaultWorkflowRepositoryPermissions(t *testing.T) {

	t.Run("test setting of write default workflow repository permissions", func(t *testing.T) {

		defaultWorkflowPermissions := "write"
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name        = "tf-acc-test-topic-%[1]s"
				description = "Terraform acceptance tests %[1]s"
				topics		= ["terraform", "testing"]
			}

			resource "github_actions_default_workflow_repository_permissions" "test" {
				default_workflow_permissions = "%s"
				repository                   = github_repository.test.name
			}
		`, randomID, defaultWorkflowPermissions)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_actions_default_workflow_repository_permissions.test", "default_workflow_permissions", defaultWorkflowPermissions,
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

	t.Run("test setting of all default workflow repository permissions params", func(t *testing.T) {

		defaultWorkflowPermissions := "write"
		canApprovePullRequestReviews := true
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name        = "tf-acc-test-topic-%[1]s"
				description = "Terraform acceptance tests %[1]s"
				topics		= ["terraform", "testing"]
			}

			resource "github_actions_default_workflow_repository_permissions" "test" {
				default_workflow_permissions     = "%s"
				can_approve_pull_request_reviews = %t
				repository                       = github_repository.test.name
			}
		`, randomID, defaultWorkflowPermissions, canApprovePullRequestReviews)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_actions_default_workflow_repository_permissions.test", "default_workflow_permissions", defaultWorkflowPermissions,
			),
			resource.TestCheckResourceAttr(
				"github_actions_default_workflow_repository_permissions.test", "can_approve_pull_request_reviews", strconv.FormatBool(canApprovePullRequestReviews),
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

	t.Run("test setting of github defaults for default workflow repository permissions", func(t *testing.T) {

		defaultWorkflowPermissions := "read"
		canApprovePullRequestReviews := false
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name        = "tf-acc-test-topic-%[1]s"
				description = "Terraform acceptance tests %[1]s"
				topics		= ["terraform", "testing"]
			}

			resource "github_actions_default_workflow_repository_permissions" "test" {
				default_workflow_permissions     = "%s"
				can_approve_pull_request_reviews = %t
				repository                       = github_repository.test.name
			}
		`, randomID, defaultWorkflowPermissions, canApprovePullRequestReviews)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_actions_default_workflow_repository_permissions.test", "default_workflow_permissions", defaultWorkflowPermissions,
			),
			resource.TestCheckResourceAttr(
				"github_actions_default_workflow_repository_permissions.test", "can_approve_pull_request_reviews", strconv.FormatBool(canApprovePullRequestReviews),
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
