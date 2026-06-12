package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubRepositoryPullRequestCreationPolicy(t *testing.T) {
	t.Run("sets policy without error", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-pr-policy-%s", testResourcePrefix, randomID)
		initial := "collaborators_only"
		updated := "all"

		config := `
			resource "github_repository" "test" {
				name       = "%s"
				visibility = "private"
				auto_init  = true
			}

			resource "github_repository_pull_request_creation_policy" "test" {
				repository = github_repository.test.name
				policy     = "%s"
			}
		`

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, repoName, initial),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(
							"github_repository_pull_request_creation_policy.test",
							tfjsonpath.New("policy"),
							knownvalue.StringExact(initial),
						),
					},
				},
				{
					Config: fmt.Sprintf(config, repoName, updated),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(
							"github_repository_pull_request_creation_policy.test",
							tfjsonpath.New("policy"),
							knownvalue.StringExact(updated),
						),
					},
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

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(
							"github_repository_pull_request_creation_policy.test",
							tfjsonpath.New("repository"),
							knownvalue.StringExact(repoName),
						),
						statecheck.ExpectKnownValue(
							"github_repository_pull_request_creation_policy.test",
							tfjsonpath.New("policy"),
							knownvalue.StringExact("collaborators_only"),
						),
					},
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
