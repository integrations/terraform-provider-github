package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubEnterpriseSettings(t *testing.T) {
	
	t.Run("creates basic enterprise settings without error", func(t *testing.T) {
		
		config := fmt.Sprintf(`
		resource "github_enterprise_settings" "test" {
			enterprise_slug = "%s"
			
			actions_enabled_organizations = "all"
			actions_allowed_actions = "all"
			
			default_workflow_permissions = "read"
			can_approve_pull_request_reviews = false
		}
		`, testEnterprise)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("github_enterprise_settings.test", "enterprise_slug", testEnterprise),
			resource.TestCheckResourceAttr("github_enterprise_settings.test", "actions_enabled_organizations", "all"),
			resource.TestCheckResourceAttr("github_enterprise_settings.test", "actions_allowed_actions", "all"),
			resource.TestCheckResourceAttr("github_enterprise_settings.test", "default_workflow_permissions", "read"),
			resource.TestCheckResourceAttr("github_enterprise_settings.test", "can_approve_pull_request_reviews", "false"),
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
			if isEnterprise != "true" {
				t.Skip("Skipping because `ENTERPRISE_ACCOUNT` is not set or set to false")
			}
			if testEnterprise == "" {
				t.Skip("Skipping because `ENTERPRISE_SLUG` is not set")
			}
			testCase(t, enterprise)
		})
	})

	t.Run("creates enterprise settings with selected actions without error", func(t *testing.T) {
		
		config := fmt.Sprintf(`
		resource "github_enterprise_settings" "test" {
			enterprise_slug = "%s"
			
			actions_enabled_organizations = "all"
			actions_allowed_actions = "selected"
			actions_github_owned_allowed = true
			actions_verified_allowed = true
			actions_patterns_allowed = [
				"actions/cache@*",
				"actions/checkout@*"
			]
			
			default_workflow_permissions = "write"
			can_approve_pull_request_reviews = true
		}
		`, testEnterprise)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("github_enterprise_settings.test", "enterprise_slug", testEnterprise),
			resource.TestCheckResourceAttr("github_enterprise_settings.test", "actions_enabled_organizations", "all"),
			resource.TestCheckResourceAttr("github_enterprise_settings.test", "actions_allowed_actions", "selected"),
			resource.TestCheckResourceAttr("github_enterprise_settings.test", "actions_github_owned_allowed", "true"),
			resource.TestCheckResourceAttr("github_enterprise_settings.test", "actions_verified_allowed", "true"),
			resource.TestCheckResourceAttr("github_enterprise_settings.test", "actions_patterns_allowed.#", "2"),
			resource.TestCheckTypeSetElemAttr("github_enterprise_settings.test", "actions_patterns_allowed.*", "actions/cache@*"),
			resource.TestCheckTypeSetElemAttr("github_enterprise_settings.test", "actions_patterns_allowed.*", "actions/checkout@*"),
			resource.TestCheckResourceAttr("github_enterprise_settings.test", "default_workflow_permissions", "write"),
			resource.TestCheckResourceAttr("github_enterprise_settings.test", "can_approve_pull_request_reviews", "true"),
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
			if isEnterprise != "true" {
				t.Skip("Skipping because `ENTERPRISE_ACCOUNT` is not set or set to false")
			}
			if testEnterprise == "" {
				t.Skip("Skipping because `ENTERPRISE_SLUG` is not set")
			}
			testCase(t, enterprise)
		})
	})

	t.Run("updates enterprise settings without error", func(t *testing.T) {
		
		configs := map[string]string{
			"before": fmt.Sprintf(`
			resource "github_enterprise_settings" "test" {
				enterprise_slug = "%s"
				
				actions_enabled_organizations = "all"
				actions_allowed_actions = "all"
				
				default_workflow_permissions = "read"
				can_approve_pull_request_reviews = false
			}
			`, testEnterprise),

			"after": fmt.Sprintf(`
			resource "github_enterprise_settings" "test" {
				enterprise_slug = "%s"
				
				actions_enabled_organizations = "all"
				actions_allowed_actions = "local_only"
				
				default_workflow_permissions = "write"
				can_approve_pull_request_reviews = true
			}
			`, testEnterprise),
		}

		checks := map[string]resource.TestCheckFunc{
			"before": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr("github_enterprise_settings.test", "actions_allowed_actions", "all"),
				resource.TestCheckResourceAttr("github_enterprise_settings.test", "default_workflow_permissions", "read"),
				resource.TestCheckResourceAttr("github_enterprise_settings.test", "can_approve_pull_request_reviews", "false"),
			),
			"after": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr("github_enterprise_settings.test", "actions_allowed_actions", "local_only"),
				resource.TestCheckResourceAttr("github_enterprise_settings.test", "default_workflow_permissions", "write"),
				resource.TestCheckResourceAttr("github_enterprise_settings.test", "can_approve_pull_request_reviews", "true"),
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

		t.Run("with an enterprise account", func(t *testing.T) {
			if isEnterprise != "true" {
				t.Skip("Skipping because `ENTERPRISE_ACCOUNT` is not set or set to false")
			}
			if testEnterprise == "" {
				t.Skip("Skipping because `ENTERPRISE_SLUG` is not set")
			}
			testCase(t, enterprise)
		})
	})

	t.Run("updates from all to selected actions policy", func(t *testing.T) {
		
		configs := map[string]string{
			"before": fmt.Sprintf(`
			resource "github_enterprise_settings" "test" {
				enterprise_slug = "%s"
				
				actions_enabled_organizations = "all"
				actions_allowed_actions = "all"
				
				default_workflow_permissions = "read"
				can_approve_pull_request_reviews = false
			}
			`, testEnterprise),

			"after": fmt.Sprintf(`
			resource "github_enterprise_settings" "test" {
				enterprise_slug = "%s"
				
				actions_enabled_organizations = "all"
				actions_allowed_actions = "selected"
				actions_github_owned_allowed = true
				actions_verified_allowed = false
				actions_patterns_allowed = [
					"my-org/custom-action@v1"
				]
				
				default_workflow_permissions = "read"
				can_approve_pull_request_reviews = false
			}
			`, testEnterprise),
		}

		checks := map[string]resource.TestCheckFunc{
			"before": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr("github_enterprise_settings.test", "actions_allowed_actions", "all"),
			),
			"after": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr("github_enterprise_settings.test", "actions_allowed_actions", "selected"),
				resource.TestCheckResourceAttr("github_enterprise_settings.test", "actions_github_owned_allowed", "true"),
				resource.TestCheckResourceAttr("github_enterprise_settings.test", "actions_verified_allowed", "false"),
				resource.TestCheckResourceAttr("github_enterprise_settings.test", "actions_patterns_allowed.#", "1"),
				resource.TestCheckTypeSetElemAttr("github_enterprise_settings.test", "actions_patterns_allowed.*", "my-org/custom-action@v1"),
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

		t.Run("with an enterprise account", func(t *testing.T) {
			if isEnterprise != "true" {
				t.Skip("Skipping because `ENTERPRISE_ACCOUNT` is not set or set to false")
			}
			if testEnterprise == "" {
				t.Skip("Skipping because `ENTERPRISE_SLUG` is not set")
			}
			testCase(t, enterprise)
		})
	})

	t.Run("imports enterprise settings without error", func(t *testing.T) {
		
		config := fmt.Sprintf(`
		resource "github_enterprise_settings" "test" {
			enterprise_slug = "%s"
			
			actions_enabled_organizations = "all"
			actions_allowed_actions = "all"
			
			default_workflow_permissions = "read"
			can_approve_pull_request_reviews = false
		}
		`, testEnterprise)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("github_enterprise_settings.test", "enterprise_slug", testEnterprise),
			resource.TestCheckResourceAttr("github_enterprise_settings.test", "actions_enabled_organizations", "all"),
			resource.TestCheckResourceAttr("github_enterprise_settings.test", "actions_allowed_actions", "all"),
			resource.TestCheckResourceAttr("github_enterprise_settings.test", "default_workflow_permissions", "read"),
			resource.TestCheckResourceAttr("github_enterprise_settings.test", "can_approve_pull_request_reviews", "false"),
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
						ResourceName:      "github_enterprise_settings.test",
						ImportState:       true,
						ImportStateVerify: true,
					},
				},
			})
		}

		t.Run("with an enterprise account", func(t *testing.T) {
			if isEnterprise != "true" {
				t.Skip("Skipping because `ENTERPRISE_ACCOUNT` is not set or set to false")
			}
			if testEnterprise == "" {
				t.Skip("Skipping because `ENTERPRISE_SLUG` is not set")
			}
			testCase(t, enterprise)
		})
	})
}
