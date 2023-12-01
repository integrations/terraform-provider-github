package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccGithubRepositoryCodeScanning(t *testing.T) {
	t.Run("enables the code scanning setup for a repository", func(t *testing.T) {
		repoName := fmt.Sprintf("tf-acc-test-code-scanning-%s", acctest.RandString(5))
		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name      = "%s"
				auto_init = true
			}

			resource "github_repository_code_scanning" "test" {
				repository = github_repository.test.name
				owner      = "terraformgithubprovidertests"

				state      = "configured"
				query_suite = "default"
				languages  = ["python"]
			}
		`, repoName)

		config2 := config + `
			data "github_repository_code_scanning" "test" {
				repository = github_repository.test.name
				owner      = "terraformgithubprovidertests"
			}
		`

		const resourceName = "data.github_repository_code_scanning.test"
		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceName, "languages.0", "python"),
			resource.TestCheckResourceAttr(resourceName, "state", "configured"),
			resource.TestCheckResourceAttr(resourceName, "query_suite", "default"),
		)

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: config,
						Check:  resource.ComposeTestCheckFunc(),
					},
					{
						Config: config2,
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

	t.Run("disables the code scanning setup for a repository", func(t *testing.T) {
		repoName := fmt.Sprintf("tf-acc-test-code-scanning-%s", acctest.RandString(5))
		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name      = "%s"
				auto_init = true
			}

			resource "github_repository_code_scanning" "test" {
				repository = github_repository.test.name
				owner      = "terraformgithubprovidertests"

				state      = "not-configured"
				query_suite = "extended"
				languages  = ["python", "javascript", "ruby"]
			}
		`, repoName)

		config2 := config + `
			data "github_repository_code_scanning" "test" {
				repository = github_repository.test.name
				owner      = "terraformgithubprovidertests"
			}
		`

		const resourceName = "data.github_repository_code_scanning.test"
		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceName, "languages.0", "python"),
			resource.TestCheckResourceAttr(resourceName, "languages.1", "javascript"),
			resource.TestCheckResourceAttr(resourceName, "languages.2", "ruby"),
			resource.TestCheckResourceAttr(resourceName, "state", "not-configured"),
			resource.TestCheckResourceAttr(resourceName, "query_suite", "extended"),
		)

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: config,
						Check:  resource.ComposeTestCheckFunc(),
					},
					{
						Config: config2,
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
