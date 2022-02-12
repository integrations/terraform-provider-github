package github

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccGithubRepositoryDataSource(t *testing.T) {

	t.Run("anonymously queries a repository without error", func(t *testing.T) {

		config := fmt.Sprintf(`
			data "github_repositories" "test" {
				query = "org:%s"
			}

			data "github_repository" "test" {
				full_name = data.github_repositories.test.full_names.0
			}
		`, testOrganization)

		check := resource.ComposeTestCheckFunc(
			resource.TestMatchResourceAttr(
				"data.github_repositories.test", "full_names.0",
				regexp.MustCompile(`^`+testOrganization)),
			resource.TestMatchResourceAttr(
				"data.github_repository.test", "full_name",
				regexp.MustCompile(`^`+testOrganization)),
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
			testCase(t, anonymous)
		})

	})

	t.Run("queries a repository with pages configured", func(t *testing.T) {

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

		config := fmt.Sprintf(`
		
			resource "github_repository" "test" {
				name         = "tf-acc-%s"
				auto_init    = true
				pages {
					source {
						branch = "main"
					}
				}
			}

			data "github_repository" "test" {
				name = github_repository.test.name
			}
		`, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"data.github_repository.test", "pages.0.source.0.branch",
				"main",
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

	t.Run("checks defaults on a new repository", func(t *testing.T) {

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

		config := fmt.Sprintf(`

			resource "github_repository" "test" {
				name         = "tf-acc-%s"
				auto_init    = true
			}

			data "github_repository" "test" {
				name = github_repository.test.name
			}
		`, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"data.github_repository.test", "name", "tf-acc-"+randomID,
			),
			resource.TestCheckResourceAttrSet(
				"data.github_repository.test", "has_projects",
			),
			resource.TestCheckResourceAttr(
				"data.github_repository.test", "description", "",
			),
			resource.TestCheckResourceAttr(
				"data.github_repository.test", "homepage_url", "",
			),
			resource.TestCheckResourceAttr(
				"data.github_repository.test", "pages.#", "0",
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

		t.Run("with an individual account", func(t *testing.T) {
			testCase(t, individual)
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})

	})

}
