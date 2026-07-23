package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/compare"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubRepositoryImmutableReleases(t *testing.T) {
	t.Parallel()

	t.Run("creates_immutable_releases_as_enabled_without_error", func(t *testing.T) {
		t.Parallel()

		randomID := acctest.RandString(5)
		repoName := fmt.Sprintf("%simmutable-releases-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name       = "%s"
				visibility = "private"
				auto_init  = true
			}

			resource "github_repository_immutable_releases" "test" {
				repository = github_repository.test.name
				enabled    = true
			}
		`, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_immutable_releases.test", tfjsonpath.New("enabled"), knownvalue.Bool(true)),
						statecheck.CompareValuePairs("github_repository_immutable_releases.test", tfjsonpath.New("repository"), "github_repository.test", tfjsonpath.New("name"), compare.ValuesSame()),
						statecheck.ExpectKnownValue("github_repository_immutable_releases.test", tfjsonpath.New("repository_id"), knownvalue.NotNull()),
					},
				},
			},
		})
	})

	t.Run("updates_immutable_releases_without_error", func(t *testing.T) {
		t.Parallel()

		randomID := acctest.RandString(5)
		repoName := fmt.Sprintf("%simmutable-releases-%s", testResourcePrefix, randomID)

		config := `
			resource "github_repository" "test" {
				name       = "%s"
				visibility = "private"
				auto_init  = true
			}

			resource "github_repository_immutable_releases" "test" {
				repository = github_repository.test.name
				enabled    = %t
			}
		`

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, repoName, true),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_immutable_releases.test", tfjsonpath.New("enabled"), knownvalue.Bool(true)),
					},
				},
				{
					Config: fmt.Sprintf(config, repoName, false),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_immutable_releases.test", tfjsonpath.New("enabled"), knownvalue.Bool(false)),
					},
				},
			},
		})
	})
}
