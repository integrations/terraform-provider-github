package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccGithubBranchProtection(t *testing.T) {

	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

	t.Run("configures default settings when empty", func(t *testing.T) {

		config := fmt.Sprintf(`

		resource "github_repository" "test" {
		  name      = "tf-acc-test-%s"
		  auto_init = true
		}

		resource "github_branch_protection" "test" {

		  repository_id  = github_repository.test.node_id
		  pattern        = "main"

		}

	`, randomID)

		check := resource.ComposeAggregateTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_branch_protection.test", "pattern", "main",
			),
			resource.TestCheckResourceAttr(
				"github_branch_protection.test", "require_signed_commits", "false",
			),
			resource.TestCheckResourceAttr(
				"github_branch_protection.test", "required_status_checks.#", "0",
			),
			resource.TestCheckResourceAttr(
				"github_branch_protection.test", "required_pull_request_reviews.#", "0",
			),
			resource.TestCheckResourceAttr(
				"github_branch_protection.test", "push_restrictions.#", "0",
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
					{
						ResourceName:      "github_branch_protection.test",
						ImportState:       true,
						ImportStateVerify: true,
						ImportStateIdFunc: branchProtectionImportStateIdFunc(
							fmt.Sprintf("tf-acc-test-%s", randomID), "main",
						),
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

	t.Run("configures required status checks", func(t *testing.T) {

		config := fmt.Sprintf(`

			resource "github_repository" "test" {
			  name      = "tf-acc-test-%s"
			  auto_init = true
			}

			resource "github_branch_protection" "test" {

			  repository_id  = github_repository.test.node_id
			  pattern        = "main"

				required_status_checks {
			    strict   = true
			    contexts = ["github/foo"]
			  }

			}

	`, randomID)

		check := resource.ComposeAggregateTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_branch_protection.test", "required_status_checks.#", "1",
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

	t.Run("configures required pull request reviews", func(t *testing.T) {

		config := fmt.Sprintf(`

			resource "github_repository" "test" {
			  name      = "tf-acc-test-%s"
			  auto_init = true
			}

			resource "github_branch_protection" "test" {

			  repository_id  = github_repository.test.node_id
			  pattern        = "main"

				required_pull_request_reviews {
						dismiss_stale_reviews      = true
						require_code_owner_reviews = true
				}

			}

	`, randomID)

		check := resource.ComposeAggregateTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_branch_protection.test", "required_pull_request_reviews.#", "1",
			),
			resource.TestCheckResourceAttr(
				"github_branch_protection.test", "required_pull_request_reviews.0.dismiss_stale_reviews", "true",
			),
			resource.TestCheckResourceAttr(
				"github_branch_protection.test", "required_pull_request_reviews.0.require_code_owner_reviews", "true",
			),
			resource.TestCheckResourceAttr(
				"github_branch_protection.test", "required_pull_request_reviews.0.required_approving_review_count", "1",
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

	t.Run("configures branch push restrictions", func(t *testing.T) {

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
			  name      = "tf-acc-test-%s"
			  auto_init = true
			}

			data "github_user" "test" {
			  username = "%s"
			}

			resource "github_branch_protection" "test" {

			  repository_id   = github_repository.test.name
			  pattern       	= "main"

			  push_restrictions = [
			    data.github_user.test.node_id,
			  ]

			}
	`, randomID, testOwnerFunc())

		check := resource.ComposeAggregateTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_branch_protection.test", "push_restrictions.#", "1",
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
			t.Skip("individual account not supported for this operation")
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})

	})
}

func branchProtectionImportStateIdFunc(repo, pattern string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		return fmt.Sprintf("%s:%s", repo, pattern), nil
	}
}
