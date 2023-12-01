package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccGithubRepositoryCodeScanningDataSource(t *testing.T) {

	randomId := acctest.RandStringFromCharSet(6, acctest.CharSetAlphaNum)
	t.Run("manages the code scanning setup for a repository", func(t *testing.T) {
		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name      = "tf-acc-test-cs-%s"
				auto_init = true
			}

			resource "github_repository_file" "test_py" {
				repository          = github_repository.test.name
				branch              = "main"
				file                = "main.py"
				content             = <<-EOT
				if __name__ == "__main__":
    				print ("This is a test")
				EOT
				commit_message      = "Managed by Terraform"
				commit_author       = "Terraform User"
				commit_email        = "terraform@example.com"
				overwrite_on_create = true
			}

			resource "github_repository_code_scanning" "test" {
				repository = github_repository.test.name
				owner      = "terraformgithubprovidertests"

				state       = "configured"
				query_suite = "default"
			}
		`, randomId)

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

	t.Run("manages the code scanning setup for a repository with multiple languages", func(t *testing.T) {
		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name      = "tf-acc-test-cs-%s"
				auto_init = true
			}

			resource "github_repository_file" "test_py" {
				repository          = github_repository.test.name
				branch              = "main"
				file                = "main.py"
				content             = <<-EOT
				if __name__ == "__main__":
    				print ("This is a test")
				EOT
				commit_message      = "Managed by Terraform"
				commit_author       = "Terraform User"
				commit_email        = "terraform@example.com"
				overwrite_on_create = true
			}

			resource "github_repository_file" "test_js" {
				repository          = github_repository.test.name
				branch              = "main"
				file                = "main.js"
				content             = <<-EOT
				function main() {
					console.log("This is a test");
				  }
				EOT
				commit_message      = "Managed by Terraform"
				commit_author       = "Terraform User"
				commit_email        = "terraform@example.com"
				overwrite_on_create = true
			}

			resource "github_repository_code_scanning" "test" {
				repository = github_repository.test.name
				owner      = "terraformgithubprovidertests"

				state      = "configured"
				query_suite = "extended"
			}
		`, randomId)

		config2 := config + `
			data "github_repository_code_scanning" "test" {
				repository = github_repository.test.name
				owner      = "terraformgithubprovidertests"
			}
		`

		const resourceName = "data.github_repository_code_scanning.test"
		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceName, "languages.0", "python"),
			resource.TestCheckResourceAttr(resourceName, "languages.1", "javascript-typescript"),
			resource.TestCheckResourceAttr(resourceName, "state", "configured"),
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
