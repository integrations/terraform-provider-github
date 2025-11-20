package github

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubBranch(t *testing.T) {
	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

	t.Run("creates a branch directly", func(t *testing.T) {
		config := fmt.Sprintf(`
			resource "github_repository" "test" {
			  name = "tf-acc-test-%[1]s"
			  auto_init = true
			}

			resource "github_branch" "test" {
			  repository = github_repository.test.id
			  branch     = "test"
			}
		`, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestMatchResourceAttr(
				"github_branch.test", "id",
				regexp.MustCompile(fmt.Sprintf("tf-acc-test-%s:test", randomID)),
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
						ResourceName:            "github_branch.test",
						ImportState:             true,
						ImportStateId:           fmt.Sprintf("tf-acc-test-%s:test", randomID),
						ImportStateVerify:       true,
						ImportStateVerifyIgnore: []string{"source_sha"},
					},
					{
						ResourceName:  "github_branch.test",
						ImportState:   true,
						ImportStateId: fmt.Sprintf("tf-acc-test-%s:nonsense", randomID),
						ExpectError: regexp.MustCompile(
							"Repository tf-acc-test-[a-z0-9]* does not have a branch named nonsense.",
						),
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

	t.Run("creates a branch named main directly and a repository with a gitignore_template", func(t *testing.T) {
		config := fmt.Sprintf(`
			resource "github_repository" "test" {
			  name = "tf-acc-test-%[1]s"
			  auto_init = true
			  gitignore_template = "Python"
			}

			resource "github_branch" "test" {
			  repository = github_repository.test.id
			  branch     = "main"
			}
		`, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestMatchResourceAttr(
				"github_branch.test", "id",
				regexp.MustCompile(fmt.Sprintf("tf-acc-test-%s:test", randomID)),
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
						ResourceName:            "github_branch.test",
						ImportState:             true,
						ImportStateId:           fmt.Sprintf("tf-acc-test-%s:test", randomID),
						ImportStateVerify:       true,
						ImportStateVerifyIgnore: []string{"source_sha"},
					},
					{
						ResourceName:  "github_branch.test",
						ImportState:   true,
						ImportStateId: fmt.Sprintf("tf-acc-test-%s:nonsense", randomID),
						ExpectError: regexp.MustCompile(
							"Repository tf-acc-test-[a-z0-9]* does not have a branch named nonsense.",
						),
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

	t.Run("creates a branch from a source branch", func(t *testing.T) {
		config := fmt.Sprintf(`
			resource "github_repository" "test" {
			  name = "tf-acc-test-%[1]s"
			  auto_init = true
			}

			resource "github_branch" "source" {
			  repository = github_repository.test.id
			  branch     = "source"
			}

			resource "github_branch" "test" {
			  repository    = github_repository.test.id
			  source_branch = github_branch.source.branch
			  branch        = "test"
			}
		`, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestMatchResourceAttr(
				"github_branch.test", "id",
				regexp.MustCompile(fmt.Sprintf("tf-acc-test-%s:test", randomID)),
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
						ResourceName:            "github_branch.test",
						ImportState:             true,
						ImportStateId:           fmt.Sprintf("tf-acc-test-%s:test:source", randomID),
						ImportStateVerify:       true,
						ImportStateVerifyIgnore: []string{"source_sha"},
					},
					{
						ResourceName:  "github_branch.test",
						ImportState:   true,
						ImportStateId: fmt.Sprintf("tf-acc-test-%s:nonsense:source", randomID),
						ExpectError: regexp.MustCompile(
							"Repository tf-acc-test-[a-z0-9]* does not have a branch named nonsense.",
						),
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
