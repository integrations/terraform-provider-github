package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccGithubRepositoryEnvironment(t *testing.T) {
	t.Run("creates a repository environment", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)
		envName := "environment / test"
		config := fmt.Sprintf(`

			data "github_user" "current" {
				username = ""
			}

			resource "github_repository" "test" {
				name      = "%s"
				visibility = "public"
			}

			resource "github_repository_environment" "test" {
				repository  = github_repository.test.name
				environment = "%s"

				can_admins_bypass   = false
				wait_timer          = 10000
				prevent_self_review = true

				reviewers {
					users = [data.github_user.current.id]
				}

				deployment_branch_policy {
					protected_branches     = true
					custom_branch_policies = false
				}
			}

		`, repoName, envName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("github_repository_environment.test", "environment", envName),
						resource.TestCheckResourceAttr("github_repository_environment.test", "can_admins_bypass", "false"),
						resource.TestCheckResourceAttr("github_repository_environment.test", "prevent_self_review", "true"),
						resource.TestCheckResourceAttr("github_repository_environment.test", "wait_timer", "10000"),
					),
				},
			},
		})
	})

	t.Run("creates a repository environment with id separator", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)
		envName := "environment:test"
		config := fmt.Sprintf(`

			data "github_user" "current" {
				username = ""
			}

			resource "github_repository" "test" {
				name      = "%s"
				visibility = "public"
			}

			resource "github_repository_environment" "test" {
				repository  = github_repository.test.name
				environment = "%s"

				can_admins_bypass   = false
				wait_timer          = 10000
				prevent_self_review = true

				reviewers {
					users = [data.github_user.current.id]
				}

				deployment_branch_policy {
					protected_branches     = true
					custom_branch_policies = false
				}
			}

		`, repoName, envName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("github_repository_environment.test", "environment", envName),
						resource.TestCheckResourceAttr("github_repository_environment.test", "can_admins_bypass", "false"),
						resource.TestCheckResourceAttr("github_repository_environment.test", "prevent_self_review", "true"),
						resource.TestCheckResourceAttr("github_repository_environment.test", "wait_timer", "10000"),
					),
				},
			},
		})
	})

	t.Run("import", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)
		envName := "environment / test"
		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name      = "%s"
				visibility = "public"
			}

			resource "github_repository_environment" "test" {
				repository 	= github_repository.test.name
				environment	= "%s"
			}
		`, repoName, envName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
				},
				{
					ResourceName:            "github_repository_environment.test",
					ImportState:             true,
					ImportStateVerify:       true,
					ImportStateVerifyIgnore: []string{"can_admins_bypass", "prevent_self_review", "reviewers", "wait_timer", "deployment_branch_policy"},
				},
			},
		})
	})
}
