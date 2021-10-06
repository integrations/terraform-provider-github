package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccGithubRepositoryAutolinkReference(t *testing.T) {

	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

	t.Run("creates repository autolink reference without error", func(t *testing.T) {

		config := fmt.Sprintf(`
			resource "github_repository" "oof" {
			  name         = "oof-%s"
			  description  = "Test autolink creation"
			}

			resource "github_repository_autolink_reference" "autolink" {
			  repository = github_repository.oof.name

			  key_prefix 		  = "OOF-"
			  target_url_template = "https://awesome.com/find/OOF-<num>"
			}
		`, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_repository_autolink_reference.autolink", "key_prefix", "OOF-",
			),
			resource.TestCheckResourceAttr(
				"github_repository_autolink_reference.autolink", "target_url_template", "https://awesome.com/find/OOF-<num>",
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

	t.Run("imports repository autolink reference without error", func(t *testing.T) {

		config := fmt.Sprintf(`
			resource "github_repository" "oof" {
			  name         = "oof-%s"
			  description  = "Test autolink creation"
			}

			resource "github_repository_autolink_reference" "autolink" {
			  repository = github_repository.oof.name

			  key_prefix 		  = "OOF-"
			  target_url_template = "https://awesome.com/find/OOF-<num>"
			}
		`, randomID)

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: config,
					},
					{
						ResourceName:        "github_repository_autolink_reference.autolink",
						ImportState:         true,
						ImportStateVerify:   true,
						ImportStateIdPrefix: fmt.Sprintf("oof-%s/", randomID),
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

	t.Run("deletes repository autolink reference without error", func(t *testing.T) {

		config := fmt.Sprintf(`
			resource "github_repository" "oof" {
			  name         = "oof-%s"
			  description  = "Test autolink creation"
			}

			resource "github_repository_autolink_reference" "autolink" {
			  repository = github_repository.oof.name

			  key_prefix 		  = "OOF-"
			  target_url_template = "https://awesome.com/find/OOF-<num>"
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
