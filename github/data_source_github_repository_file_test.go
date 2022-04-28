package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccGithubRepositoryFileDataSource(t *testing.T) {

	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

	t.Run("create and read a file with a branch name provided", func(t *testing.T) {

		config := fmt.Sprintf(`

			resource "github_repository" "test" {
				name      = "tf-acc-test-%s"
				auto_init = true
			}

			resource "github_repository_file" "test" {
				repository     = github_repository.test.name
				branch         = "main"
				file           = "test"
				content        = "bar"
				commit_message = "Managed by Terraform"
				commit_author  = "Terraform User"
				commit_email   = "terraform@example.com"
			}

			data "github_repository_file" "test" {
				repository     = github_repository.test.name
				branch         = "main"
				file           = github_repository_file.test.file
			}
		`, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"data.github_repository_file.test", "content",
				"bar",
			),
			resource.TestCheckResourceAttr(
				"data.github_repository_file.test", "sha",
				"ba0e162e1c47469e3fe4b393a8bf8c569f302116",
			),
			resource.TestCheckResourceAttr(
				"data.github_repository_file.test", "ref",
				"main",
			),
			resource.TestCheckResourceAttrSet(
				"data.github_repository_file.test", "commit_author",
			),
			resource.TestCheckResourceAttrSet(
				"data.github_repository_file.test", "commit_email",
			),
			resource.TestCheckResourceAttrSet(
				"data.github_repository_file.test", "commit_message",
			),
			resource.TestCheckResourceAttrSet(
				"data.github_repository_file.test", "commit_sha",
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
			testCase(t, anonymous)
		})

		t.Run("with an individual account", func(t *testing.T) {
			testCase(t, individual)
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})

	})

	t.Run("create and read a file without providing a branch name", func(t *testing.T) {

		config := fmt.Sprintf(`

			resource "github_repository" "test" {
				name      			= "tf-acc-test-%s"
				auto_init 			= true
			}

			resource "github_branch" "test" {
				repository = github_repository.test.name
				branch     = "test"
			}

			resource "github_branch_default" "default"{
				repository = github_repository.test.name
				branch     = github_branch.test.branch
			}

			resource "github_repository_file" "test" {
				repository     = github_repository.test.name
				branch         = github_branch_default.default.branch
				file           = "test"
				content        = "bar"
				commit_message = "Managed by Terraform"
				commit_author  = "Terraform User"
				commit_email   = "terraform@example.com"
			}

			data "github_repository_file" "test" {
				repository     = github_repository.test.name
				file           = github_repository_file.test.file
			}
		`, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"data.github_repository_file.test", "content",
				"bar",
			),
			resource.TestCheckResourceAttr(
				"data.github_repository_file.test", "sha",
				"ba0e162e1c47469e3fe4b393a8bf8c569f302116",
			),
			resource.TestCheckResourceAttr(
				"data.github_repository_file.test", "ref",
				"test",
			),
			resource.TestCheckResourceAttrSet(
				"data.github_repository_file.test", "commit_author",
			),
			resource.TestCheckResourceAttrSet(
				"data.github_repository_file.test", "commit_email",
			),
			resource.TestCheckResourceAttrSet(
				"data.github_repository_file.test", "commit_message",
			),
			resource.TestCheckResourceAttrSet(
				"data.github_repository_file.test", "commit_sha",
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
			testCase(t, anonymous)
		})

		t.Run("with an individual account", func(t *testing.T) {
			testCase(t, individual)
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})

	})

	t.Run("try reading a non-existent file without an error", func(t *testing.T) {

		config := fmt.Sprintf(`

			resource "github_repository" "test" {
				name      			= "tf-acc-test-%s"
				auto_init 			= true
			}

			data "github_repository_file" "test" {
				repository     = github_repository.test.name
				file           = "test"
			}
		`, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckNoResourceAttr(
				"data.github_repository_file.test", "id",
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
			testCase(t, anonymous)
		})

		t.Run("with an individual account", func(t *testing.T) {
			testCase(t, individual)
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})

	})
}
