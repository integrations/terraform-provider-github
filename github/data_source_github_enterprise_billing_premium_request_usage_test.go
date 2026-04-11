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

func TestAccGithubEnterpriseBillingPremiumRequestUsage(t *testing.T) {
	t.Run("reads premium request usage without error", func(t *testing.T) {
		config := fmt.Sprintf(`
			data "github_enterprise_billing_premium_request_usage" "test" {
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
						statecheck.ExpectKnownValue("data.github_enterprise_billing_premium_request_usage.test",
							tfjsonpath.New("enterprise_slug"),
							knownvalue.StringExact(testAccConf.enterpriseSlug),
						),
						statecheck.ExpectKnownValue("data.github_enterprise_billing_premium_request_usage.test",
							tfjsonpath.New("id"),
							knownvalue.StringRegexp(regexp.MustCompile(`^.+:billing-premium-request-usage$`)),
						),
						statecheck.ExpectKnownValue("data.github_enterprise_billing_premium_request_usage.test",
							tfjsonpath.New("usage_items"),
							knownvalue.NotNull(),
						),
						statecheck.ExpectKnownValue("data.github_enterprise_billing_premium_request_usage.test",
							tfjsonpath.New("time_period"),
							knownvalue.NotNull(),
						),
					},
				},
			},
		})
	})

	t.Run("reads premium request usage with filters without error", func(t *testing.T) {
		config := fmt.Sprintf(`
			data "github_enterprise_billing_premium_request_usage" "test" {
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
						statecheck.ExpectKnownValue("data.github_enterprise_billing_premium_request_usage.test",
							tfjsonpath.New("enterprise_slug"),
							knownvalue.StringExact(testAccConf.enterpriseSlug),
						),
						statecheck.ExpectKnownValue("data.github_enterprise_billing_premium_request_usage.test",
							tfjsonpath.New("usage_items"),
							knownvalue.NotNull(),
						),
						statecheck.ExpectKnownValue("data.github_enterprise_billing_premium_request_usage.test",
							tfjsonpath.New("time_period"),
							knownvalue.NotNull(),
						),
					},
				},
			},
		})
	})
}
