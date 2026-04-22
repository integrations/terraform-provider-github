package github

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccGithubEnterprisePrivateRepositoryForkingSetting(t *testing.T) {
	t.Run("enables private repository forking with policy", func(t *testing.T) {
		config := fmt.Sprintf(`
		resource "github_enterprise_private_repository_forking_setting" "test" {
			enterprise_slug = "%s"
			setting_value   = "ENABLED"
			policy_value    = "SAME_ORGANIZATION"
		}
		`, testAccConf.enterpriseSlug)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("github_enterprise_private_repository_forking_setting.test", "enterprise_slug", testAccConf.enterpriseSlug),
			resource.TestCheckResourceAttr("github_enterprise_private_repository_forking_setting.test", "setting_value", "ENABLED"),
			resource.TestCheckResourceAttr("github_enterprise_private_repository_forking_setting.test", "policy_value", "SAME_ORGANIZATION"),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  check,
				},
			},
		})
	})

	t.Run("updates policy value", func(t *testing.T) {
		configs := map[string]string{
			"before": fmt.Sprintf(`
			resource "github_enterprise_private_repository_forking_setting" "test" {
				enterprise_slug = "%s"
				setting_value   = "ENABLED"
				policy_value    = "SAME_ORGANIZATION"
			}
			`, testAccConf.enterpriseSlug),

			"after": fmt.Sprintf(`
			resource "github_enterprise_private_repository_forking_setting" "test" {
				enterprise_slug = "%s"
				setting_value   = "ENABLED"
				policy_value    = "ENTERPRISE_ORGANIZATIONS_USER_ACCOUNTS"
			}
			`, testAccConf.enterpriseSlug),
		}

		checks := map[string]resource.TestCheckFunc{
			"before": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr("github_enterprise_private_repository_forking_setting.test", "setting_value", "ENABLED"),
				resource.TestCheckResourceAttr("github_enterprise_private_repository_forking_setting.test", "policy_value", "SAME_ORGANIZATION"),
			),
			"after": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr("github_enterprise_private_repository_forking_setting.test", "setting_value", "ENABLED"),
				resource.TestCheckResourceAttr("github_enterprise_private_repository_forking_setting.test", "policy_value", "ENTERPRISE_ORGANIZATIONS_USER_ACCOUNTS"),
			),
		}

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
			ProviderFactories: providerFactories,
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

	t.Run("disables private repository forking", func(t *testing.T) {
		config := fmt.Sprintf(`
		resource "github_enterprise_private_repository_forking_setting" "test" {
			enterprise_slug = "%s"
			setting_value   = "DISABLED"
		}
		`, testAccConf.enterpriseSlug)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("github_enterprise_private_repository_forking_setting.test", "enterprise_slug", testAccConf.enterpriseSlug),
			resource.TestCheckResourceAttr("github_enterprise_private_repository_forking_setting.test", "setting_value", "DISABLED"),
			resource.TestCheckResourceAttr("github_enterprise_private_repository_forking_setting.test", "policy_value", ""),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  check,
				},
			},
		})
	})

	t.Run("sets no policy", func(t *testing.T) {
		config := fmt.Sprintf(`
		resource "github_enterprise_private_repository_forking_setting" "test" {
			enterprise_slug = "%s"
			setting_value   = "NO_POLICY"
		}
		`, testAccConf.enterpriseSlug)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("github_enterprise_private_repository_forking_setting.test", "enterprise_slug", testAccConf.enterpriseSlug),
			resource.TestCheckResourceAttr("github_enterprise_private_repository_forking_setting.test", "setting_value", "NO_POLICY"),
			resource.TestCheckResourceAttr("github_enterprise_private_repository_forking_setting.test", "policy_value", ""),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  check,
				},
			},
		})
	})

	t.Run("transitions from enabled to disabled", func(t *testing.T) {
		configs := map[string]string{
			"enabled": fmt.Sprintf(`
			resource "github_enterprise_private_repository_forking_setting" "test" {
				enterprise_slug = "%s"
				setting_value   = "ENABLED"
				policy_value    = "SAME_ORGANIZATION"
			}
			`, testAccConf.enterpriseSlug),

			"disabled": fmt.Sprintf(`
			resource "github_enterprise_private_repository_forking_setting" "test" {
				enterprise_slug = "%s"
				setting_value   = "DISABLED"
			}
			`, testAccConf.enterpriseSlug),
		}

		checks := map[string]resource.TestCheckFunc{
			"enabled": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr("github_enterprise_private_repository_forking_setting.test", "setting_value", "ENABLED"),
				resource.TestCheckResourceAttr("github_enterprise_private_repository_forking_setting.test", "policy_value", "SAME_ORGANIZATION"),
			),
			"disabled": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr("github_enterprise_private_repository_forking_setting.test", "setting_value", "DISABLED"),
				resource.TestCheckResourceAttr("github_enterprise_private_repository_forking_setting.test", "policy_value", ""),
			),
		}

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: configs["enabled"],
					Check:  checks["enabled"],
				},
				{
					Config: configs["disabled"],
					Check:  checks["disabled"],
				},
			},
		})
	})

	t.Run("rejects policy_value when disabled", func(t *testing.T) {
		config := fmt.Sprintf(`
		resource "github_enterprise_private_repository_forking_setting" "test" {
			enterprise_slug = "%s"
			setting_value   = "DISABLED"
			policy_value    = "SAME_ORGANIZATION"
		}
		`, testAccConf.enterpriseSlug)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      config,
					ExpectError: regexp.MustCompile(`policy_value must not be set when setting_value is DISABLED`),
				},
			},
		})
	})

	t.Run("requires policy_value when enabled", func(t *testing.T) {
		config := fmt.Sprintf(`
		resource "github_enterprise_private_repository_forking_setting" "test" {
			enterprise_slug = "%s"
			setting_value   = "ENABLED"
		}
		`, testAccConf.enterpriseSlug)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      config,
					ExpectError: regexp.MustCompile(`policy_value is required when setting_value is ENABLED`),
				},
			},
		})
	})

	t.Run("imports without error", func(t *testing.T) {
		config := fmt.Sprintf(`
		resource "github_enterprise_private_repository_forking_setting" "test" {
			enterprise_slug = "%s"
			setting_value   = "ENABLED"
			policy_value    = "ENTERPRISE_ORGANIZATIONS_USER_ACCOUNTS"
		}
		`, testAccConf.enterpriseSlug)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("github_enterprise_private_repository_forking_setting.test", "enterprise_slug", testAccConf.enterpriseSlug),
			resource.TestCheckResourceAttr("github_enterprise_private_repository_forking_setting.test", "setting_value", "ENABLED"),
			resource.TestCheckResourceAttr("github_enterprise_private_repository_forking_setting.test", "policy_value", "ENTERPRISE_ORGANIZATIONS_USER_ACCOUNTS"),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  check,
				},
				{
					ResourceName:      "github_enterprise_private_repository_forking_setting.test",
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})
}
