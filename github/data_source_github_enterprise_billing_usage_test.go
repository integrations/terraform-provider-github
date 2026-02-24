package github

import (
	"fmt"
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
					},
				},
			},
		})
	})
}
