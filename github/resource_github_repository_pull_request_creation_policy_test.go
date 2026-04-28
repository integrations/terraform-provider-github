package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccGithubRepositoryPullRequestCreationPolicy(t *testing.T) {
	t.Run("sets policy without error", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-pr-policy-%s", testResourcePrefix, randomID)
		initial := `policy = "collaborators_only"`
		updated := `policy = "all"`

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name       = "%s"
				visibility = "private"
				auto_init  = true
			}

			resource "github_repository_pull_request_creation_policy" "test" {
				repository = github_repository.test.name
				%%s
			}
		`, repoName)

		checks := map[string]resource.TestCheckFunc{
			"before": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_repository_pull_request_creation_policy.test", "policy",
					"collaborators_only",
				),
			),
			"after": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_repository_pull_request_creation_policy.test", "policy",
					"all",
				),
			),
		}

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, initial),
					Check:  checks["before"],
				},
				{
					Config: fmt.Sprintf(config, updated),
					Check:  checks["after"],
				},
			},
		})
	})

	t.Run("imports without error", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-pr-policy-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name       = "%s"
				visibility = "private"
				auto_init  = true
			}

			resource "github_repository_pull_request_creation_policy" "test" {
				repository = github_repository.test.name
				policy     = "collaborators_only"
			}
		`, repoName)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet("github_repository_pull_request_creation_policy.test", "repository"),
			resource.TestCheckResourceAttr("github_repository_pull_request_creation_policy.test", "policy", "collaborators_only"),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  check,
				},
				{
					ResourceName:      "github_repository_pull_request_creation_policy.test",
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})
}
