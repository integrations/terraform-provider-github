package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubEnterpriseCostCenter(t *testing.T) {
	t.Run("creates cost center without error", func(t *testing.T) {
		randomID := acctest.RandString(5)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(`
						data "github_enterprise" "enterprise" {
							slug = "%s"
						}

						resource "github_enterprise_cost_center" "test" {
							enterprise_slug = data.github_enterprise.enterprise.slug
							name            = "%s%s"
						}
					`, testAccConf.enterpriseSlug, testResourcePrefix, randomID),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_enterprise_cost_center.test", tfjsonpath.New("enterprise_slug"), knownvalue.StringExact(testAccConf.enterpriseSlug)),
						statecheck.ExpectKnownValue("github_enterprise_cost_center.test", tfjsonpath.New("name"), knownvalue.StringExact(testResourcePrefix+randomID)),
						statecheck.ExpectKnownValue("github_enterprise_cost_center.test", tfjsonpath.New("state"), knownvalue.StringExact("active")),
					},
				},
			},
		})
	})

	t.Run("updates cost center name without error", func(t *testing.T) {
		randomID := acctest.RandString(5)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(`
						data "github_enterprise" "enterprise" {
							slug = "%s"
						}

						resource "github_enterprise_cost_center" "test" {
							enterprise_slug = data.github_enterprise.enterprise.slug
							name            = "%s%s"
						}
					`, testAccConf.enterpriseSlug, testResourcePrefix, randomID),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_enterprise_cost_center.test", tfjsonpath.New("name"), knownvalue.StringExact(testResourcePrefix+randomID)),
					},
				},
				{
					Config: fmt.Sprintf(`
						data "github_enterprise" "enterprise" {
							slug = "%s"
						}

						resource "github_enterprise_cost_center" "test" {
							enterprise_slug = data.github_enterprise.enterprise.slug
							name            = "%supdated-%s"
						}
					`, testAccConf.enterpriseSlug, testResourcePrefix, randomID),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_enterprise_cost_center.test", tfjsonpath.New("name"), knownvalue.StringExact(testResourcePrefix+"updated-"+randomID)),
					},
				},
			},
		})
	})

	t.Run("imports cost center without error", func(t *testing.T) {
		randomID := acctest.RandString(5)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(`
						data "github_enterprise" "enterprise" {
							slug = "%s"
						}

						resource "github_enterprise_cost_center" "test" {
							enterprise_slug = data.github_enterprise.enterprise.slug
							name            = "%s%s"
						}
					`, testAccConf.enterpriseSlug, testResourcePrefix, randomID),
				},
				{
					ResourceName:        "github_enterprise_cost_center.test",
					ImportState:         true,
					ImportStateVerify:   true,
					ImportStateIdPrefix: testAccConf.enterpriseSlug + ":",
				},
			},
		})
	})
}
