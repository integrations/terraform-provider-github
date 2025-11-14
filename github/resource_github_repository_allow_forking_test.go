package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubRepositoryAllowForking(t *testing.T) {
	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

	t.Run("manages allow_forking setting for a repository", func(t *testing.T) {

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name          = "tf-acc-test-allow-forking-%[1]s"
				description   = "Terraform acceptance test for allow_forking (Issue #353)"
				visibility    = "private"
				allow_forking = false
				auto_init     = true
			}
		`, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_repository.test", "allow_forking",
				"false",
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

	t.Run("reads allow_forking setting from API", func(t *testing.T) {

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name          = "tf-acc-test-fork-read-%[1]s"
				description   = "Test allow_forking read from API"
				visibility    = "private"
				allow_forking = false
				auto_init     = true
			}
		`, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_repository.test", "allow_forking",
				"false",
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
}
