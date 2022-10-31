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
			resource "github_repository" "test" {
				name		= "test-%s"
				description	= "Test autolink creation"
			}

			resource "github_repository_autolink_reference" "autolink_default" {
				repository = github_repository.test.name

				key_prefix			= "TEST1-"
				target_url_template = "https://example.com/TEST-<num>"
			}

			resource "github_repository_autolink_reference" "autolink_alphanumeric" {
				repository = github_repository.test.name

				key_prefix 		    = "TEST2-"
				target_url_template = "https://example.com/TEST-<num>"
				is_alphanumeric     = true
			}

			resource "github_repository_autolink_reference" "autolink_numeric" {
				repository = github_repository.test.name

				key_prefix 		    = "TEST3-"
				target_url_template = "https://example.com/TEST-<num>"
				is_alphanumeric     = false
			}
		`, randomID)

		check := resource.ComposeTestCheckFunc(
			// autolink_default
			resource.TestCheckResourceAttr(
				"github_repository_autolink_reference.autolink_default", "key_prefix", "TEST1-",
			),
			resource.TestCheckResourceAttr(
				"github_repository_autolink_reference.autolink_default", "target_url_template", "https://example.com/TEST-<num>",
			),
			resource.TestCheckResourceAttr(
				"github_repository_autolink_reference.autolink_default", "is_alphanumeric", "true",
			),
			// autolink_alphanumeric
			resource.TestCheckResourceAttr(
				"github_repository_autolink_reference.autolink_alphanumeric", "key_prefix", "TEST2-",
			),
			resource.TestCheckResourceAttr(
				"github_repository_autolink_reference.autolink_alphanumeric", "target_url_template", "https://example.com/TEST-<num>",
			),
			resource.TestCheckResourceAttr(
				"github_repository_autolink_reference.autolink_alphanumeric", "is_alphanumeric", "true",
			),
			// autolink_numeric
			resource.TestCheckResourceAttr(
				"github_repository_autolink_reference.autolink_numeric", "key_prefix", "TEST3-",
			),
			resource.TestCheckResourceAttr(
				"github_repository_autolink_reference.autolink_numeric", "target_url_template", "https://example.com/TEST-<num>",
			),
			resource.TestCheckResourceAttr(
				"github_repository_autolink_reference.autolink_numeric", "is_alphanumeric", "false",
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
			resource "github_repository" "test" {
				name		= "test-%s"
				description	= "Test autolink creation"
			}

			resource "github_repository_autolink_reference" "autolink_default" {
				repository = github_repository.test.name

				key_prefix			= "TEST1-"
				target_url_template = "https://example.com/TEST-<num>"
			}

			resource "github_repository_autolink_reference" "autolink_alphanumeric" {
				repository = github_repository.test.name

				key_prefix 		    = "TEST2-"
				target_url_template = "https://example.com/TEST-<num>"
				is_alphanumeric     = true
			}

			resource "github_repository_autolink_reference" "autolink_numeric" {
				repository = github_repository.test.name

				key_prefix 		    = "TEST3-"
				target_url_template = "https://example.com/TEST-<num>"
				is_alphanumeric     = false
			}
		`, randomID)

		check := resource.ComposeTestCheckFunc(
			// autolink_default
			resource.TestCheckResourceAttr(
				"github_repository_autolink_reference.autolink_default", "key_prefix", "TEST1-",
			),
			resource.TestCheckResourceAttr(
				"github_repository_autolink_reference.autolink_default", "target_url_template", "https://example.com/TEST-<num>",
			),
			resource.TestCheckResourceAttr(
				"github_repository_autolink_reference.autolink_default", "is_alphanumeric", "true",
			),
			// autolink_alphanumeric
			resource.TestCheckResourceAttr(
				"github_repository_autolink_reference.autolink_alphanumeric", "key_prefix", "TEST2-",
			),
			resource.TestCheckResourceAttr(
				"github_repository_autolink_reference.autolink_alphanumeric", "target_url_template", "https://example.com/TEST-<num>",
			),
			resource.TestCheckResourceAttr(
				"github_repository_autolink_reference.autolink_alphanumeric", "is_alphanumeric", "true",
			),
			// autolink_numeric
			resource.TestCheckResourceAttr(
				"github_repository_autolink_reference.autolink_numeric", "key_prefix", "TEST3-",
			),
			resource.TestCheckResourceAttr(
				"github_repository_autolink_reference.autolink_numeric", "target_url_template", "https://example.com/TEST-<num>",
			),
			resource.TestCheckResourceAttr(
				"github_repository_autolink_reference.autolink_numeric", "is_alphanumeric", "false",
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
					// autolink_default
					{
						ResourceName:        "github_repository_autolink_reference.autolink_default",
						ImportState:         true,
						ImportStateVerify:   true,
						ImportStateIdPrefix: fmt.Sprintf("test-%s/", randomID),
					},
					// autolink_alphanumeric
					{
						ResourceName:        "github_repository_autolink_reference.autolink_alphanumeric",
						ImportState:         true,
						ImportStateVerify:   true,
						ImportStateIdPrefix: fmt.Sprintf("test-%s/", randomID),
					},
					// autolink_numeric
					{
						ResourceName:        "github_repository_autolink_reference.autolink_numeric",
						ImportState:         true,
						ImportStateVerify:   true,
						ImportStateIdPrefix: fmt.Sprintf("test-%s/", randomID),
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
			resource "github_repository" "test" {
				name		= "test-%s"
				description	= "Test autolink creation"
			}

			resource "github_repository_autolink_reference" "autolink_default" {
				repository = github_repository.test.name

				key_prefix			= "TEST1-"
				target_url_template = "https://example.com/TEST-<num>"
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
