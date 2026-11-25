package github

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccGithubRestApiDataSource(t *testing.T) {
	t.Parallel()

	t.Run("queries_an_existing_branch_without_error", func(t *testing.T) {
		t.Parallel()

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-rest-api-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
			  name = "%[1]s"
				auto_init = true
			}

			data "github_rest_api" "test" {
				endpoint = "repos/${github_repository.test.full_name}/git/refs/heads/${github_repository.test.default_branch}"
			}
		`, repoName)

		check := resource.ComposeTestCheckFunc(
			resource.TestMatchResourceAttr(
				"data.github_rest_api.test", "code", regexp.MustCompile("200"),
			),
			resource.TestMatchResourceAttr(
				"data.github_rest_api.test", "status", regexp.MustCompile("200 OK"),
			),
			resource.TestMatchResourceAttr("data.github_rest_api.test", "body", regexp.MustCompile(".*refs/heads/.*")),
			resource.TestCheckResourceAttrSet("data.github_rest_api.test", "headers"),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  check,
				},
			},
		})
	})

	t.Run("queries_a_collection_without_error", func(t *testing.T) {
		t.Parallel()

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-rest-api-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
			  name = "%[1]s"
				auto_init = true
			}

			data "github_rest_api" "test" {
				endpoint = "repos/${github_repository.test.full_name}/git/refs/heads/"
			}
		`, repoName)

		check := resource.ComposeTestCheckFunc(
			resource.TestMatchResourceAttr("data.github_rest_api.test", "body", regexp.MustCompile(`\[.*refs/heads/.*\]`)),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  check,
				},
			},
		})
	})

	t.Run("queries_an_invalid_branch_without_error", func(t *testing.T) {
		t.Parallel()

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-rest-api-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
			  name = "%[1]s"
				auto_init = true
			}

			data "github_rest_api" "test" {
				endpoint = "repos/${github_repository.test.full_name}/git/refs/heads/xxxxxx"
			}
		`, repoName)

		check := resource.ComposeTestCheckFunc(
			resource.TestMatchResourceAttr(
				"data.github_rest_api.test", "code", regexp.MustCompile("404"),
			),
			resource.TestMatchResourceAttr(
				"data.github_rest_api.test", "status", regexp.MustCompile("404 Not Found"),
			),
			resource.TestCheckResourceAttrSet("data.github_rest_api.test", "body"),
			resource.TestCheckResourceAttrSet("data.github_rest_api.test", "headers"),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  check,
				},
			},
		})
	})

	// t.Run("fails_for_invalid_endpoint", func(t *testing.T) {
	// 	t.Parallel()

	// 	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
	// 	repoName := fmt.Sprintf("%srepo-rest-api-%s", testResourcePrefix, randomID)
	// 	config := `
	// 		data "github_rest_api" "test" {
	// 			endpoint = "/xxx"
	// 		}
	// 	`

	// 	resource.Test(t, resource.TestCase{
	// 		PreCheck:          func() { skipUnauthenticated(t) },
	// 		ProviderFactories: providerFactories,
	// 		Steps: []resource.TestStep{
	// 			{
	// 				Config:      config,
	// 				ExpectError: regexp.MustCompile("Error: GET https://api.github.com/xx.*: 414"),
	// 			},
	// 		},
	// 	})
	// })
}
