package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubActionsEnterpriseRunnerGroup(t *testing.T) {
	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

	if isEnterprise != "true" {
		t.Skip("Skipping because `ENTERPRISE_ACCOUNT` is not set or set to false")
	}

	if testEnterprise == "" {
		t.Skip("Skipping because `ENTERPRISE_SLUG` is not set")
	}

	t.Run("creates enterprise runner groups without error", func(t *testing.T) {
		config := fmt.Sprintf(`
			data "github_enterprise" "enterprise" {
				slug = "%s"
			}

			resource "github_enterprise_actions_runner_group" "test" {
				enterprise_slug				= data.github_enterprise.enterprise.slug
				name						= "tf-acc-test-%s"
				visibility					= "all"
				allows_public_repositories	= true
			}
		`, testEnterprise, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet(
				"github_enterprise_actions_runner_group.test", "name",
			),
			resource.TestCheckResourceAttr(
				"github_enterprise_actions_runner_group.test", "name",
				fmt.Sprintf(`tf-acc-test-%s`, randomID),
			),
			resource.TestCheckResourceAttr(
				"github_enterprise_actions_runner_group.test", "visibility",
				"all",
			),
			resource.TestCheckResourceAttr(
				"github_enterprise_actions_runner_group.test", "allows_public_repositories",
				"true",
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

		t.Run("with an enterprise account", func(t *testing.T) {
			testCase(t, enterprise)
		})
	})

	t.Run("manages runner group visibility to selected orgs", func(t *testing.T) {

		config := fmt.Sprintf(`
			data "github_enterprise" "enterprise" {
				slug = "%s"
			}

			data "github_organization" "org" {
				name 			= "%s"
			}

			resource "github_enterprise_actions_runner_group" "test" {
				enterprise_slug				= data.github_enterprise.enterprise.slug
				name						= "tf-acc-test-%s"
				visibility					= "selected"
				selected_organization_ids	= [data.github_organization.org.id]
			}
		`, testEnterprise, testOrganization, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet(
				"github_enterprise_actions_runner_group.test", "name",
			),
			resource.TestCheckResourceAttr(
				"github_enterprise_actions_runner_group.test", "name",
				fmt.Sprintf(`tf-acc-test-%s`, randomID),
			),
			resource.TestCheckResourceAttr(
				"github_enterprise_actions_runner_group.test", "visibility",
				"selected",
			),
			resource.TestCheckResourceAttr(
				"github_enterprise_actions_runner_group.test", "selected_organization_ids.#",
				"1",
			),
			resource.TestCheckResourceAttrSet(
				"github_enterprise_actions_runner_group.test", "selected_organizations_url",
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

		t.Run("with an enterprise account", func(t *testing.T) {
			testCase(t, enterprise)
		})
	})

	t.Run("imports an all runner group without error", func(t *testing.T) {
		config := fmt.Sprintf(`
			data "github_enterprise" "enterprise" {
				slug = "%s"
			}

			resource "github_enterprise_actions_runner_group" "test" {
				enterprise_slug = data.github_enterprise.enterprise.slug
				name       		= "tf-acc-test-%s"
				visibility 		= "all"
			}
	`, testEnterprise, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet("github_enterprise_actions_runner_group.test", "name"),
			resource.TestCheckResourceAttrSet("github_enterprise_actions_runner_group.test", "visibility"),
			resource.TestCheckResourceAttr("github_enterprise_actions_runner_group.test", "visibility", "all"),
			resource.TestCheckResourceAttr("github_enterprise_actions_runner_group.test", "name", fmt.Sprintf(`tf-acc-test-%s`, randomID)),
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
						ResourceName:        "github_enterprise_actions_runner_group.test",
						ImportState:         true,
						ImportStateVerify:   true,
						ImportStateIdPrefix: fmt.Sprintf(`%s/`, testEnterprise),
					},
				},
			})
		}

		t.Run("with an enterprise account", func(t *testing.T) {
			testCase(t, enterprise)
		})
	})

	t.Run("imports a runner group with selected orgs without error", func(t *testing.T) {

		config := fmt.Sprintf(`
			data "github_enterprise" "enterprise" {
				slug = "%s"
			}

			data "github_organization" "org" {
				name 			= "%s"
			}

			resource "github_enterprise_actions_runner_group" "test" {
				enterprise_slug				= data.github_enterprise.enterprise.slug
				name						= "tf-acc-test-%s"
				visibility					= "selected"
				selected_organization_ids	= [data.github_organization.org.id]
			}
		`, testEnterprise, testOrganization, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet("github_enterprise_actions_runner_group.test", "name"),
			resource.TestCheckResourceAttr("github_enterprise_actions_runner_group.test", "name", fmt.Sprintf(`tf-acc-test-%s`, randomID)),
			resource.TestCheckResourceAttrSet("github_enterprise_actions_runner_group.test", "visibility"),
			resource.TestCheckResourceAttr("github_enterprise_actions_runner_group.test", "visibility", "selected"),
			resource.TestCheckResourceAttr(
				"github_enterprise_actions_runner_group.test", "selected_organization_ids.#",
				"1",
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
						ResourceName:        "github_enterprise_actions_runner_group.test",
						ImportState:         true,
						ImportStateVerify:   true,
						ImportStateIdPrefix: fmt.Sprintf(`%s/`, testEnterprise),
					},
				},
			})
		}

		t.Run("with an enterprise account", func(t *testing.T) {
			testCase(t, enterprise)
		})
	})
}
