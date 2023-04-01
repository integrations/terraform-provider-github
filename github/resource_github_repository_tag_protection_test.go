package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccGithubRepositoryTagProtection(t *testing.T) {
	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
	t.Run("creates tag protection without error", func(t *testing.T) {

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
			  name         = "tf-acc-test-%s"
			  auto_init    = true
			}

			resource "github_repository_tag_protection" "test" {
				depends_on = ["github_repository.test"]
			  	repository = github_repository.test.name
				pattern    = "v*"
			}
		`, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_repository_tag_protection.test", "pattern", "v*",
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
		t.Run("run with an anonymous account", func(t *testing.T) {
			t.Skip("anonymous account not supported for this operation")
		})
		t.Run("run with an individual account", func(t *testing.T) {
			testCase(t, individual)
		})
		t.Run("run with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})

	})
	t.Run("imports tag protection without error", func(t *testing.T) {

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
			  name         = "tf-acc-test-%s"
			  auto_init    = true
			}

			resource "github_repository_tag_protection" "test" {
				depends_on = ["github_repository.test"]
			  	repository = github_repository.test.name
			  	pattern    = "v*"
			}
		`, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_repository_tag_protection.test", "pattern", "v*",
			),
			resource.TestCheckResourceAttrSet(
				"github_repository_tag_protection.test", "id",
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
					{
						ResourceName:        "github_repository_tag_protection.test",
						ImportState:         true,
						ImportStateIdPrefix: fmt.Sprintf("tf-acc-test-%s/", randomID),
						ImportStateVerify:   true,
					},
				},
			})
		}
		t.Run("run with an anonymous account", func(t *testing.T) {
			t.Skip("anonymous account not supported for this operation")
		})
		t.Run("run with an individual account", func(t *testing.T) {
			testCase(t, individual)
		})
		t.Run("run with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})
	})
	t.Run("deletes tag protection without error", func(t *testing.T) {

		config := fmt.Sprintf(`
				resource "github_repository" "test" {
					name = "tf-acc-test-%s"
					auto_init = true
				}

				resource "github_repository_tag_protection" "test" {
					depends_on = ["github_repository.test"]
					repository       = github_repository.test.name
					pattern          = "v*"
				}
			`, randomID)

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config:  config,
						Destroy: true,
					},
				},
			})
		}

		t.Run("run with an anonymous account", func(t *testing.T) {
			t.Skip("anonymous account not supported for this operation")
		})

		t.Run("run with an individual account", func(t *testing.T) {
			testCase(t, individual)
		})

		t.Run("run with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})

	})

}
