package github

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubRepositoryAutolinkReference(t *testing.T) {
	t.Run("creates repository autolink reference without error", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-autolink-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name        = "%s"
				description = "Test autolink creation"
			}

			resource "github_repository_autolink_reference" "autolink_default" {
				repository = github_repository.test.name

				key_prefix          = "TEST1-"
				target_url_template = "https://example.com/TEST-<num>"
			}

			resource "github_repository_autolink_reference" "autolink_alphanumeric" {
				repository = github_repository.test.name

				key_prefix          = "TEST2-"
				target_url_template = "https://example.com/TEST-<num>"
				is_alphanumeric     = true
			}

			resource "github_repository_autolink_reference" "autolink_numeric" {
				repository = github_repository.test.name

				key_prefix          = "TEST3-"
				target_url_template = "https://example.com/TEST-<num>"
				is_alphanumeric     = false
			}

			resource "github_repository_autolink_reference" "autolink_with_port" {
				repository = github_repository.test.name

				key_prefix          = "TEST4-"
				target_url_template = "https://example.com:8443/TEST-<num>"
			}
		`, repoName)

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
			// autolink_with_port
			resource.TestCheckResourceAttr(
				"github_repository_autolink_reference.autolink_with_port", "key_prefix", "TEST4-",
			),
			resource.TestCheckResourceAttr(
				"github_repository_autolink_reference.autolink_with_port", "target_url_template", "https://example.com:8443/TEST-<num>",
			),
			resource.TestCheckResourceAttr(
				"github_repository_autolink_reference.autolink_with_port", "is_alphanumeric", "true",
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

	t.Run("imports repository autolink reference without error", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-autolink-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name        = "%s"
				description = "Test autolink creation"
			}

			resource "github_repository_autolink_reference" "autolink_default" {
				repository = github_repository.test.name

				key_prefix          = "TEST1-"
				target_url_template = "https://example.com/TEST-<num>"
			}

			resource "github_repository_autolink_reference" "autolink_alphanumeric" {
				repository = github_repository.test.name

				key_prefix          = "TEST2-"
				target_url_template = "https://example.com/TEST-<num>"
				is_alphanumeric     = true
			}

			resource "github_repository_autolink_reference" "autolink_numeric" {
				repository = github_repository.test.name

				key_prefix          = "TEST3-"
				target_url_template = "https://example.com/TEST-<num>"
				is_alphanumeric     = false
			}

			resource "github_repository_autolink_reference" "autolink_with_port" {
				repository = github_repository.test.name

				key_prefix          = "TEST4-"
				target_url_template = "https://example.com:8443/TEST-<num>"
			}
		`, repoName)

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
			// autolink_with_port
			resource.TestCheckResourceAttr(
				"github_repository_autolink_reference.autolink_with_port", "key_prefix", "TEST4-",
			),
			resource.TestCheckResourceAttr(
				"github_repository_autolink_reference.autolink_with_port", "target_url_template", "https://example.com:8443/TEST-<num>",
			),
			resource.TestCheckResourceAttr(
				"github_repository_autolink_reference.autolink_with_port", "is_alphanumeric", "true",
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
				// autolink_default
				{
					ResourceName:        "github_repository_autolink_reference.autolink_default",
					ImportState:         true,
					ImportStateVerify:   true,
					ImportStateIdPrefix: fmt.Sprintf("%s/", repoName),
				},
				// autolink_alphanumeric
				{
					ResourceName:        "github_repository_autolink_reference.autolink_alphanumeric",
					ImportState:         true,
					ImportStateVerify:   true,
					ImportStateIdPrefix: fmt.Sprintf("%s/", repoName),
				},
				// autolink_numeric
				{
					ResourceName:        "github_repository_autolink_reference.autolink_numeric",
					ImportState:         true,
					ImportStateVerify:   true,
					ImportStateIdPrefix: fmt.Sprintf("%s/", repoName),
				},
				// autolink_with_port
				{
					ResourceName:        "github_repository_autolink_reference.autolink_with_port",
					ImportState:         true,
					ImportStateVerify:   true,
					ImportStateIdPrefix: fmt.Sprintf("%s/", repoName),
				},
			},
		})
	})

	t.Run("imports repository autolink reference by key prefix without error", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-autolink-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`
			resource "github_repository" "oof" {
			  name         = "%s"
			  description  = "Test autolink creation"
			}

			resource "github_repository_autolink_reference" "autolink" {
			  repository = github_repository.oof.name

			  key_prefix 		  = "OOF-"
			  target_url_template = "https://awesome.com/find/OOF-<num>"
			}
		`, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
				},
				{
					ResourceName:      "github_repository_autolink_reference.autolink",
					ImportState:       true,
					ImportStateVerify: true,
					ImportStateId:     fmt.Sprintf("%s/OOF-", repoName),
				},
				{
					ResourceName:  "github_repository_autolink_reference.autolink",
					ImportState:   true,
					ImportStateId: fmt.Sprintf("%s/OCTOCAT-", repoName),
					ExpectError:   regexp.MustCompile(`cannot find autolink reference`),
				},
			},
		})
	})

	t.Run("deletes repository autolink reference without error", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-autolink-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name        = "%s"
				description = "Test autolink creation"
			}

			resource "github_repository_autolink_reference" "autolink_default" {
				repository = github_repository.test.name

				key_prefix          = "TEST1-"
				target_url_template = "https://example.com/TEST-<num>"
			}
		`, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:  config,
					Destroy: true,
				},
			},
		})
	})
}
