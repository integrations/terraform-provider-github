package github

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubEnterpriseBillingUsage(t *testing.T) {
	t.Run("reads billing usage without error", func(t *testing.T) {
		config := fmt.Sprintf(`
			data "github_enterprise_billing_usage" "test" {
				enterprise_slug = "%s"
			}
		`, testAccConf.enterpriseSlug)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessMode(t, enterprise) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_enterprise_billing_usage.test",
							tfjsonpath.New("enterprise_slug"),
							knownvalue.StringExact(testAccConf.enterpriseSlug),
						),
						statecheck.ExpectKnownValue("data.github_enterprise_billing_usage.test",
							tfjsonpath.New("id"),
							knownvalue.StringRegexp(regexp.MustCompile(`^.+:billing-usage$`)),
						),
						statecheck.ExpectKnownValue("data.github_enterprise_billing_usage.test",
							tfjsonpath.New("usage_items"),
							knownvalue.NotNull(),
						),
					},
				},
			},
		})
	})

	t.Run("reads billing usage with year filter without error", func(t *testing.T) {
		config := fmt.Sprintf(`
			data "github_enterprise_billing_usage" "test" {
				enterprise_slug = "%s"
				year            = 2025
			}
		`, testAccConf.enterpriseSlug)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessMode(t, enterprise) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_enterprise_billing_usage.test",
							tfjsonpath.New("enterprise_slug"),
							knownvalue.StringExact(testAccConf.enterpriseSlug),
						),
						statecheck.ExpectKnownValue("data.github_enterprise_billing_usage.test",
							tfjsonpath.New("usage_items"),
							knownvalue.NotNull(),
						),
					},
				},
			},
		})
	})

	t.Run("reads billing usage with month filter without error", func(t *testing.T) {
		config := fmt.Sprintf(`
			data "github_enterprise_billing_usage" "test" {
				enterprise_slug = "%s"
				year            = 2025
				month           = 1
			}
		`, testAccConf.enterpriseSlug)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessMode(t, enterprise) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_enterprise_billing_usage.test",
							tfjsonpath.New("enterprise_slug"),
							knownvalue.StringExact(testAccConf.enterpriseSlug),
						),
						statecheck.ExpectKnownValue("data.github_enterprise_billing_usage.test",
							tfjsonpath.New("usage_items"),
							knownvalue.NotNull(),
						),
					},
				},
			},
		})
	})
}
