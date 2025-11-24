package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubEnterpriseSecurityAnalysisSettings(t *testing.T) {
	t.Run("creates enterprise security analysis settings without error", func(t *testing.T) {
		config := fmt.Sprintf(`
		resource "github_enterprise_security_analysis_settings" "test" {
			enterprise_slug = "%s"

			advanced_security_enabled_for_new_repositories = true
			secret_scanning_enabled_for_new_repositories = true
			secret_scanning_push_protection_enabled_for_new_repositories = false
			secret_scanning_validity_checks_enabled = true
		}
		`, testAccConf.enterpriseSlug)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("github_enterprise_security_analysis_settings.test", "enterprise_slug", testAccConf.enterpriseSlug),
			resource.TestCheckResourceAttr("github_enterprise_security_analysis_settings.test", "advanced_security_enabled_for_new_repositories", "true"),
			resource.TestCheckResourceAttr("github_enterprise_security_analysis_settings.test", "secret_scanning_enabled_for_new_repositories", "true"),
			resource.TestCheckResourceAttr("github_enterprise_security_analysis_settings.test", "secret_scanning_push_protection_enabled_for_new_repositories", "false"),
			resource.TestCheckResourceAttr("github_enterprise_security_analysis_settings.test", "secret_scanning_validity_checks_enabled", "true"),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:  func() { skipUnlessEnterprise(t) },
			Providers: testAccProviders,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  check,
				},
			},
		})
	})

	t.Run("creates enterprise security analysis settings with custom link", func(t *testing.T) {
		config := fmt.Sprintf(`
		resource "github_enterprise_security_analysis_settings" "test" {
			enterprise_slug = "%s"

			advanced_security_enabled_for_new_repositories = true
			secret_scanning_enabled_for_new_repositories = true
			secret_scanning_push_protection_enabled_for_new_repositories = true
			secret_scanning_push_protection_custom_link = "https://example.com/security-help"
			secret_scanning_validity_checks_enabled = true
		}
		`, testAccConf.enterpriseSlug)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("github_enterprise_security_analysis_settings.test", "enterprise_slug", testAccConf.enterpriseSlug),
			resource.TestCheckResourceAttr("github_enterprise_security_analysis_settings.test", "advanced_security_enabled_for_new_repositories", "true"),
			resource.TestCheckResourceAttr("github_enterprise_security_analysis_settings.test", "secret_scanning_enabled_for_new_repositories", "true"),
			resource.TestCheckResourceAttr("github_enterprise_security_analysis_settings.test", "secret_scanning_push_protection_enabled_for_new_repositories", "true"),
			resource.TestCheckResourceAttr("github_enterprise_security_analysis_settings.test", "secret_scanning_push_protection_custom_link", "https://example.com/security-help"),
			resource.TestCheckResourceAttr("github_enterprise_security_analysis_settings.test", "secret_scanning_validity_checks_enabled", "true"),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:  func() { skipUnlessEnterprise(t) },
			Providers: testAccProviders,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  check,
				},
			},
		})
	})

	t.Run("updates enterprise security analysis settings without error", func(t *testing.T) {
		configs := map[string]string{
			"before": fmt.Sprintf(`
			resource "github_enterprise_security_analysis_settings" "test" {
				enterprise_slug = "%s"

				advanced_security_enabled_for_new_repositories = false
				secret_scanning_enabled_for_new_repositories = false
				secret_scanning_push_protection_enabled_for_new_repositories = false
				secret_scanning_validity_checks_enabled = false
			}
			`, testAccConf.enterpriseSlug),

			"after": fmt.Sprintf(`
			resource "github_enterprise_security_analysis_settings" "test" {
				enterprise_slug = "%s"

				advanced_security_enabled_for_new_repositories = true
				secret_scanning_enabled_for_new_repositories = true
				secret_scanning_push_protection_enabled_for_new_repositories = true
				secret_scanning_push_protection_custom_link = "https://updated.example.com/security"
				secret_scanning_validity_checks_enabled = true
			}
			`, testAccConf.enterpriseSlug),
		}

		checks := map[string]resource.TestCheckFunc{
			"before": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr("github_enterprise_security_analysis_settings.test", "advanced_security_enabled_for_new_repositories", "false"),
				resource.TestCheckResourceAttr("github_enterprise_security_analysis_settings.test", "secret_scanning_enabled_for_new_repositories", "false"),
				resource.TestCheckResourceAttr("github_enterprise_security_analysis_settings.test", "secret_scanning_push_protection_enabled_for_new_repositories", "false"),
				resource.TestCheckResourceAttr("github_enterprise_security_analysis_settings.test", "secret_scanning_validity_checks_enabled", "false"),
			),
			"after": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr("github_enterprise_security_analysis_settings.test", "advanced_security_enabled_for_new_repositories", "true"),
				resource.TestCheckResourceAttr("github_enterprise_security_analysis_settings.test", "secret_scanning_enabled_for_new_repositories", "true"),
				resource.TestCheckResourceAttr("github_enterprise_security_analysis_settings.test", "secret_scanning_push_protection_enabled_for_new_repositories", "true"),
				resource.TestCheckResourceAttr("github_enterprise_security_analysis_settings.test", "secret_scanning_push_protection_custom_link", "https://updated.example.com/security"),
				resource.TestCheckResourceAttr("github_enterprise_security_analysis_settings.test", "secret_scanning_validity_checks_enabled", "true"),
			),
		}

		resource.Test(t, resource.TestCase{
			PreCheck:  func() { skipUnlessEnterprise(t) },
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
	})

	t.Run("creates minimal enterprise security analysis settings", func(t *testing.T) {
		config := fmt.Sprintf(`
		resource "github_enterprise_security_analysis_settings" "test" {
			enterprise_slug = "%s"
		}
		`, testAccConf.enterpriseSlug)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("github_enterprise_security_analysis_settings.test", "enterprise_slug", testAccConf.enterpriseSlug),
			resource.TestCheckResourceAttr("github_enterprise_security_analysis_settings.test", "advanced_security_enabled_for_new_repositories", "false"),
			resource.TestCheckResourceAttr("github_enterprise_security_analysis_settings.test", "secret_scanning_enabled_for_new_repositories", "false"),
			resource.TestCheckResourceAttr("github_enterprise_security_analysis_settings.test", "secret_scanning_push_protection_enabled_for_new_repositories", "false"),
			resource.TestCheckResourceAttr("github_enterprise_security_analysis_settings.test", "secret_scanning_validity_checks_enabled", "false"),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:  func() { skipUnlessEnterprise(t) },
			Providers: testAccProviders,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  check,
				},
			},
		})
	})

	t.Run("imports enterprise security analysis settings without error", func(t *testing.T) {
		config := fmt.Sprintf(`
		resource "github_enterprise_security_analysis_settings" "test" {
			enterprise_slug = "%s"

			advanced_security_enabled_for_new_repositories = true
			secret_scanning_enabled_for_new_repositories = true
			secret_scanning_push_protection_enabled_for_new_repositories = false
			secret_scanning_validity_checks_enabled = true
		}
		`, testAccConf.enterpriseSlug)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("github_enterprise_security_analysis_settings.test", "enterprise_slug", testAccConf.enterpriseSlug),
			resource.TestCheckResourceAttr("github_enterprise_security_analysis_settings.test", "advanced_security_enabled_for_new_repositories", "true"),
			resource.TestCheckResourceAttr("github_enterprise_security_analysis_settings.test", "secret_scanning_enabled_for_new_repositories", "true"),
			resource.TestCheckResourceAttr("github_enterprise_security_analysis_settings.test", "secret_scanning_push_protection_enabled_for_new_repositories", "false"),
			resource.TestCheckResourceAttr("github_enterprise_security_analysis_settings.test", "secret_scanning_validity_checks_enabled", "true"),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:  func() { skipUnlessEnterprise(t) },
			Providers: testAccProviders,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  check,
				},
				{
					ResourceName:      "github_enterprise_security_analysis_settings.test",
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})

	t.Run("handles custom link removal", func(t *testing.T) {
		configs := map[string]string{
			"with_link": fmt.Sprintf(`
			resource "github_enterprise_security_analysis_settings" "test" {
				enterprise_slug = "%s"

				secret_scanning_push_protection_enabled_for_new_repositories = true
				secret_scanning_push_protection_custom_link = "https://example.com/help"
			}
			`, testAccConf.enterpriseSlug),

			"without_link": fmt.Sprintf(`
			resource "github_enterprise_security_analysis_settings" "test" {
				enterprise_slug = "%s"

				secret_scanning_push_protection_enabled_for_new_repositories = true
			}
			`, testAccConf.enterpriseSlug),
		}

		checks := map[string]resource.TestCheckFunc{
			"with_link": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr("github_enterprise_security_analysis_settings.test", "secret_scanning_push_protection_enabled_for_new_repositories", "true"),
				resource.TestCheckResourceAttr("github_enterprise_security_analysis_settings.test", "secret_scanning_push_protection_custom_link", "https://example.com/help"),
			),
			"without_link": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr("github_enterprise_security_analysis_settings.test", "secret_scanning_push_protection_enabled_for_new_repositories", "true"),
				resource.TestCheckResourceAttr("github_enterprise_security_analysis_settings.test", "secret_scanning_push_protection_custom_link", ""),
			),
		}

		resource.Test(t, resource.TestCase{
			PreCheck:  func() { skipUnlessEnterprise(t) },
			Providers: testAccProviders,
			Steps: []resource.TestStep{
				{
					Config: configs["with_link"],
					Check:  checks["with_link"],
				},
				{
					Config: configs["without_link"],
					Check:  checks["without_link"],
				},
			},
		})
	})
}
