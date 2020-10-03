package github

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccGithubRepositoryFile(t *testing.T) {

	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

	t.Run("create and manage files", func(t *testing.T) {

		config := fmt.Sprintf(`


			resource "github_repository" "test" {
				name      = "tf-acc-test-%s"
				auto_init = true
			}

			resource "github_branch" "test" {
				repository = github_repository.test.id
				branch     = "tf-acc-test-%[1]s"
			}

			resource "github_repository_file" "test" {
				repository     = github_repository.test.name
				branch         = github_branch.test.branch
				file           = "test"
				content        = "bar"
				commit_message = "Managed by Terraform"
				commit_author  = "Terraform User"
				commit_email   = "terraform@example.com"
				overwrite      = true
			}

			resource "github_repository_file" "gitignore" {
				repository = github_repository.test.name
				file       = ".gitignore"
				content    = "**/*.tfstate"
			}
		`, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestMatchResourceAttr(
				"github_repository_file.test", "url",
				regexp.MustCompile(randomID+"/projects/1"),
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

	// t.Run("can be configured to overwrite files on create", func(t *testing.T) {
	//
	// 	config := fmt.Sprintf(`
	//
	// 		resource "github_repository" "foo" {
	// 		  name      = "tf-acc-test-%s"
	// 		  auto_init = true
	// 		}
	//
	// 		resource "github_repository_file" "foo" {
	// 		  repository     = github_repository.foo.name
	// 		  branch         = "master"
	// 		  file           = "foo"
	// 		  content        = "bar"
	// 		  commit_message = "Managed by Terraform"
	// 		  commit_author  = "Terraform User"
	// 		  commit_email   = "terraform@example.com"
	// 		  overwrite      = true
	// 		}
	//
	// 		resource "github_repository_file" "gitignore" {
	// 		  repository = "example"
	// 		  file       = ".gitignore"
	// 		  content    = "**/*.tfstate"
	// 		}
	//
	// 	`, randomID)
	//
	// 	check := resource.ComposeTestCheckFunc(
	// 		resource.TestMatchResourceAttr(
	// 			"github_repository_file.test", "url",
	// 			regexp.MustCompile(randomID+"/projects/1"),
	// 		),
	// 	)
	//
	// 	testCase := func(t *testing.T, mode string) {
	// 		resource.Test(t, resource.TestCase{
	// 			PreCheck:  func() { skipUnlessMode(t, mode) },
	// 			Providers: testAccProviders,
	// 			Steps: []resource.TestStep{
	// 				{
	// 					Config: config,
	// 					Check:  check,
	// 				},
	// 			},
	// 		})
	// 	}
	//
	// 	t.Run("with an anonymous account", func(t *testing.T) {
	// 		t.Skip("anonymous account not supported for this operation")
	// 	})
	//
	// 	t.Run("with an individual account", func(t *testing.T) {
	// 		testCase(t, individual)
	// 	})
	//
	// 	t.Run("with an organization account", func(t *testing.T) {
	// 		testCase(t, organization)
	// 	})
	//
	// })
}
