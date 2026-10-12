package github

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccGithubRepositoriesDataSource(t *testing.T) {
	t.Parallel()

	// FIXME: Find a way to reduce amount of `GET /search/repositories`
	// t.Skip("Skipping due to API rate limits exceeding")

	t.Run("queries_a_list_of_repositories_without_error", func(t *testing.T) {
		t.Parallel()

		config := fmt.Sprintf(`
			data "github_repositories" "test" {
				query = "org:%s"
			}
		`, testAccConf.owner)

		check := resource.ComposeTestCheckFunc(
			resource.TestMatchResourceAttr(
				"data.github_repositories.test", "full_names.0",
				regexp.MustCompile(`^`+testAccConf.owner),
			),
			resource.TestCheckResourceAttrSet(
				"data.github_repositories.test", "names.0",
			),
			resource.TestCheckNoResourceAttr(
				"data.github_repositories.test", "repo_ids.0",
			),
			resource.TestCheckResourceAttr(
				"data.github_repositories.test", "sort",
				"updated",
			),
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

	t.Run("queries_a_list_of_repositories_with_repo_ids_and_results_per_page_without_error", func(t *testing.T) {
		t.Parallel()

		config := fmt.Sprintf(`
			data "github_repositories" "test" {
				query = "org:%s"
				include_repo_id = true
				results_per_page = 20
			}
		`, testAccConf.owner)

		check := resource.ComposeTestCheckFunc(
			resource.TestMatchResourceAttr(
				"data.github_repositories.test", "full_names.0",
				regexp.MustCompile(`^`+testAccConf.owner),
			),
			resource.TestCheckResourceAttrSet(
				"data.github_repositories.test", "names.0",
			),
			resource.TestCheckResourceAttrSet(
				"data.github_repositories.test", "repo_ids.0",
			),
			resource.TestCheckResourceAttr(
				"data.github_repositories.test", "sort",
				"updated",
			),
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

	t.Run("returns_an_empty_list_given_an_invalid_query", func(t *testing.T) {
		t.Parallel()

		// FIXME: Find a way to reduce amount of `GET /search/repositories`
		// t.Skip("Skipping due to API rate limits exceeding")

		config := `
			data "github_repositories" "test" {
				query = "klsafj_23434_doesnt_exist"
			}
		`

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"data.github_repositories.test", "full_names.#",
				"0",
			),
			resource.TestCheckResourceAttr(
				"data.github_repositories.test", "names.#",
				"0",
			),
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
}
