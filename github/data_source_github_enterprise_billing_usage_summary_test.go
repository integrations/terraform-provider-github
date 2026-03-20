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

func TestAccGithubEnterpriseBillingUsageSummary(t *testing.T) {
	t.Run("reads billing usage summary without error", func(t *testing.T) {
		config := fmt.Sprintf(`
			data "github_enterprise_billing_usage_summary" "test" {
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
						statecheck.ExpectKnownValue("data.github_enterprise_billing_usage_summary.test",
							tfjsonpath.New("enterprise_slug"),
							knownvalue.StringExact(testAccConf.enterpriseSlug),
						),
						statecheck.ExpectKnownValue("data.github_enterprise_billing_usage_summary.test",
							tfjsonpath.New("id"),
							knownvalue.StringRegexp(regexp.MustCompile(`^.+:billing-usage-summary$`)),
						),
						statecheck.ExpectKnownValue("data.github_enterprise_billing_usage_summary.test",
							tfjsonpath.New("usage_items"),
							knownvalue.NotNull(),
						),
						statecheck.ExpectKnownValue("data.github_enterprise_billing_usage_summary.test",
							tfjsonpath.New("time_period"),
							knownvalue.NotNull(),
						),
					},
				},
			},
		})
	})

	t.Run("reads billing usage summary with filters without error", func(t *testing.T) {
		config := fmt.Sprintf(`
			data "github_enterprise_billing_usage_summary" "test" {
				enterprise_slug = "%s"
				year            = 2025
				product         = "Actions"
			}
		`, testAccConf.enterpriseSlug)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessMode(t, enterprise) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_enterprise_billing_usage_summary.test",
							tfjsonpath.New("enterprise_slug"),
							knownvalue.StringExact(testAccConf.enterpriseSlug),
						),
						statecheck.ExpectKnownValue("data.github_enterprise_billing_usage_summary.test",
							tfjsonpath.New("usage_items"),
							knownvalue.NotNull(),
						),
						statecheck.ExpectKnownValue("data.github_enterprise_billing_usage_summary.test",
							tfjsonpath.New("time_period"),
							knownvalue.NotNull(),
						),
					},
				},
			},
		})
	})
}
