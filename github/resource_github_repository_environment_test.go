package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubRepositoryEnvironment(t *testing.T) {
	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

	t.Run("creates a repository environment", func(t *testing.T) {
		config := fmt.Sprintf(`

			data "github_user" "current" {
				username = ""
			}

			resource "github_repository" "test" {
				name      = "tf-acc-test-%s"
				visibility = "public"
			}

			resource "github_repository_environment" "test" {
				repository 	= github_repository.test.name
				environment	= "environment / test"
				can_admins_bypass = false
				wait_timer = 10000
                                prevent_self_review = true
				reviewers {
					users = [data.github_user.current.id]
				}
				deployment_branch_policy {
					protected_branches     = true
					custom_branch_policies = false
				}
			}

		`, randomID)

		check := resource.ComposeAggregateTestCheckFunc(
			resource.TestCheckResourceAttr("github_repository_environment.test", "environment", "environment / test"),
			resource.TestCheckResourceAttr("github_repository_environment.test", "can_admins_bypass", "false"),
			resource.TestCheckResourceAttr("github_repository_environment.test", "prevent_self_review", "true"),
			resource.TestCheckResourceAttr("github_repository_environment.test", "wait_timer", "10000"),
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
