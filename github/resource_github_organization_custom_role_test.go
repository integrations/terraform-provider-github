package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubOrganizationCustomRole(t *testing.T) {
	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

	t.Run("creates custom repo role without error", func(t *testing.T) {
		config := fmt.Sprintf(`
			resource "github_organization_custom_role" "test" {
			  name        = "tf-acc-test-%s"
			  description = "Test role description"
			  base_role   = "read"
			  permissions = [
					"reopen_issue",
					"reopen_pull_request",
				]
			}
		`, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet(
				"github_organization_custom_role.test", "name",
			),
			resource.TestCheckResourceAttr(
				"github_organization_custom_role.test", "name",
				fmt.Sprintf(`tf-acc-test-%s`, randomID),
			),
			resource.TestCheckResourceAttr(
				"github_organization_custom_role.test", "description",
				"Test role description",
			),
			resource.TestCheckResourceAttr(
				"github_organization_custom_role.test", "base_role",
				"read",
			),
			resource.TestCheckResourceAttr(
				"github_organization_custom_role.test", "permissions.#",
				"2",
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
			t.Skip("individual account not supported for this operation")
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})
	})

	// More tests can go here following the same format...
	t.Run("updates custom repo role without error", func(t *testing.T) {
		configs := map[string]string{
			"before": fmt.Sprintf(`
				resource "github_organization_custom_role" "test" {
					name        = "tf-acc-test-%s"
					description = "Updated test role description before"
					base_role   = "read"
					permissions = [
						"reopen_issue",
						"reopen_pull_request",
					]
				}
			`, randomID),
			"after": fmt.Sprintf(`
				resource "github_organization_custom_role" "test" {
					name        = "tf-acc-test-rename-%s"
					description = "Updated test role description after"
					base_role   = "write"
					permissions = [
						"reopen_issue",
						"read_code_scanning",
						"reopen_pull_request",
					]
				}
			`, randomID),
		}

		checks := map[string]resource.TestCheckFunc{
			"before": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttrSet(
					"github_organization_custom_role.test", "name",
				),
				resource.TestCheckResourceAttr(
					"github_organization_custom_role.test", "name",
					fmt.Sprintf(`tf-acc-test-%s`, randomID),
				),
				resource.TestCheckResourceAttr(
					"github_organization_custom_role.test", "description",
					"Updated test role description before",
				),
				resource.TestCheckResourceAttr(
					"github_organization_custom_role.test", "base_role",
					"read",
				),
				resource.TestCheckResourceAttr("github_organization_custom_role.test", "permissions.#", "2"),
			),
			"after": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttrSet(
					"github_organization_custom_role.test", "name",
				),
				resource.TestCheckResourceAttr(
					"github_organization_custom_role.test", "name",
					fmt.Sprintf(`tf-acc-test-rename-%s`, randomID),
				),
				resource.TestCheckResourceAttr(
					"github_organization_custom_role.test", "description",
					"Updated test role description after",
				),
				resource.TestCheckResourceAttr(
					"github_organization_custom_role.test", "base_role",
					"write",
				),
				resource.TestCheckResourceAttr("github_organization_custom_role.test", "permissions.#", "3"),
			),
		}

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: configs["before"],
						Check:  checks["before"],
					},
					{
						Config: configs["after"],
						Check:  checks["after"],
					},
				},
			})
		}

		t.Run("with an anonymous account", func(t *testing.T) {
			t.Skip("anonymous account not supported for this operation")
		})

		t.Run("with an individual account", func(t *testing.T) {
			t.Skip("individual account not supported for this operation")
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})
	})

	t.Run("imports custom repo role without error", func(t *testing.T) {
		config := fmt.Sprintf(`
			resource "github_organization_custom_role" "test" {
			  name        = "tf-acc-test-%s"
			  description = "Test role description"
			  base_role   = "read"
			  permissions = [
					"reopen_issue",
					"reopen_pull_request",
					"read_code_scanning"
				]
			}
    `, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet(
				"github_organization_custom_role.test", "name",
			),
			resource.TestCheckResourceAttr(
				"github_organization_custom_role.test", "name",
				fmt.Sprintf(`tf-acc-test-%s`, randomID),
			),
			resource.TestCheckResourceAttr(
				"github_organization_custom_role.test", "description",
				"Test role description",
			),
			resource.TestCheckResourceAttr(
				"github_organization_custom_role.test", "base_role",
				"read",
			),
			resource.TestCheckResourceAttr(
				"github_organization_custom_role.test", "permissions.#",
				"3",
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
						ResourceName:      "github_organization_custom_role.test",
						ImportState:       true,
						ImportStateVerify: true,
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
