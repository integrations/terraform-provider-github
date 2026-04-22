package github

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const testAccEnterpriseForkingSettingResource = "github_enterprise_private_repository_forking_setting.test"

func testAccEnterpriseForkingSettingConfig(settingValue, policyValue string) string {
	if policyValue != "" {
		return fmt.Sprintf(`
		resource "github_enterprise_private_repository_forking_setting" "test" {
			enterprise_slug = "%s"
			setting_value   = "%s"
			policy_value    = "%s"
		}
		`, testAccConf.enterpriseSlug, settingValue, policyValue)
	}
	return fmt.Sprintf(`
	resource "github_enterprise_private_repository_forking_setting" "test" {
		enterprise_slug = "%s"
		setting_value   = "%s"
	}
	`, testAccConf.enterpriseSlug, settingValue)
}

func testAccEnterpriseForkingSettingCheck(settingValue, policyValue string) resource.TestCheckFunc {
	return resource.ComposeTestCheckFunc(
		resource.TestCheckResourceAttr(testAccEnterpriseForkingSettingResource, "enterprise_slug", testAccConf.enterpriseSlug),
		resource.TestCheckResourceAttr(testAccEnterpriseForkingSettingResource, "setting_value", settingValue),
		resource.TestCheckResourceAttr(testAccEnterpriseForkingSettingResource, "policy_value", policyValue),
	)
}

func TestAccGithubEnterprisePrivateRepositoryForkingSetting(t *testing.T) {
	t.Run("enables private repository forking with policy", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: testAccEnterpriseForkingSettingConfig("ENABLED", "SAME_ORGANIZATION"),
					Check:  testAccEnterpriseForkingSettingCheck("ENABLED", "SAME_ORGANIZATION"),
				},
			},
		})
	})

	t.Run("updates policy value", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: testAccEnterpriseForkingSettingConfig("ENABLED", "SAME_ORGANIZATION"),
					Check:  testAccEnterpriseForkingSettingCheck("ENABLED", "SAME_ORGANIZATION"),
				},
				{
					Config: testAccEnterpriseForkingSettingConfig("ENABLED", "ENTERPRISE_ORGANIZATIONS_USER_ACCOUNTS"),
					Check:  testAccEnterpriseForkingSettingCheck("ENABLED", "ENTERPRISE_ORGANIZATIONS_USER_ACCOUNTS"),
				},
			},
		})
	})

	t.Run("disables private repository forking", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: testAccEnterpriseForkingSettingConfig("DISABLED", ""),
					Check:  testAccEnterpriseForkingSettingCheck("DISABLED", ""),
				},
			},
		})
	})

	t.Run("sets no policy", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: testAccEnterpriseForkingSettingConfig("NO_POLICY", ""),
					Check:  testAccEnterpriseForkingSettingCheck("NO_POLICY", ""),
				},
			},
		})
	})

	t.Run("transitions from enabled to disabled", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: testAccEnterpriseForkingSettingConfig("ENABLED", "SAME_ORGANIZATION"),
					Check:  testAccEnterpriseForkingSettingCheck("ENABLED", "SAME_ORGANIZATION"),
				},
				{
					Config: testAccEnterpriseForkingSettingConfig("DISABLED", ""),
					Check:  testAccEnterpriseForkingSettingCheck("DISABLED", ""),
				},
			},
		})
	})

	t.Run("rejects policy_value when disabled", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      testAccEnterpriseForkingSettingConfig("DISABLED", "SAME_ORGANIZATION"),
					ExpectError: regexp.MustCompile(`policy_value must not be set when setting_value is DISABLED`),
				},
			},
		})
	})

	t.Run("requires policy_value when enabled", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      testAccEnterpriseForkingSettingConfig("ENABLED", ""),
					ExpectError: regexp.MustCompile(`policy_value is required when setting_value is ENABLED`),
				},
			},
		})
	})

	t.Run("imports without error", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: testAccEnterpriseForkingSettingConfig("ENABLED", "ENTERPRISE_ORGANIZATIONS_USER_ACCOUNTS"),
					Check:  testAccEnterpriseForkingSettingCheck("ENABLED", "ENTERPRISE_ORGANIZATIONS_USER_ACCOUNTS"),
				},
				{
					ResourceName:      testAccEnterpriseForkingSettingResource,
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})
}
