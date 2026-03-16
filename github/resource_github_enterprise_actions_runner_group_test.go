package github

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/compare"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func testCheckEnterpriseRunnerGroupNetworkConfigurationMatches(resourceName, networkConfigurationResourceName string) statecheck.StateCheck {
	return statecheck.CompareValuePairs(
		resourceName,
		tfjsonpath.New("network_configuration_id"),
		networkConfigurationResourceName,
		tfjsonpath.New("id"),
		compare.ValuesSame(),
	)
}

func TestAccGithubActionsEnterpriseRunnerGroup(t *testing.T) {
	t.Run("creates enterprise runner groups without error", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
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
		`, testAccConf.enterpriseSlug, randomID)

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

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessMode(t, enterprise) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  check,
				},
			},
		})
	})

	t.Run("manages runner group visibility to selected orgs", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
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
		`, testAccConf.enterpriseSlug, testAccConf.owner, randomID)

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

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessMode(t, enterprise) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  check,
				},
			},
		})
	})

	t.Run("imports an all runner group without error", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		config := fmt.Sprintf(`
			data "github_enterprise" "enterprise" {
				slug = "%s"
			}

			resource "github_enterprise_actions_runner_group" "test" {
				enterprise_slug = data.github_enterprise.enterprise.slug
				name       		= "tf-acc-test-%s"
				visibility 		= "all"
			}
	`, testAccConf.enterpriseSlug, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet("github_enterprise_actions_runner_group.test", "name"),
			resource.TestCheckResourceAttrSet("github_enterprise_actions_runner_group.test", "visibility"),
			resource.TestCheckResourceAttr("github_enterprise_actions_runner_group.test", "visibility", "all"),
			resource.TestCheckResourceAttr("github_enterprise_actions_runner_group.test", "name", fmt.Sprintf(`tf-acc-test-%s`, randomID)),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessMode(t, enterprise) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  check,
				},
				{
					ResourceName:        "github_enterprise_actions_runner_group.test",
					ImportState:         true,
					ImportStateVerify:   true,
					ImportStateIdPrefix: fmt.Sprintf(`%s/`, testAccConf.enterpriseSlug),
				},
			},
		})
	})

	t.Run("imports a runner group with selected orgs without error", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
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
		`, testAccConf.enterpriseSlug, testAccConf.owner, randomID)

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

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessMode(t, enterprise) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  check,
				},
				{
					ResourceName:        "github_enterprise_actions_runner_group.test",
					ImportState:         true,
					ImportStateVerify:   true,
					ImportStateIdPrefix: fmt.Sprintf(`%s/`, testAccConf.enterpriseSlug),
				},
			},
		})
	})

	t.Run("manages runner group network configuration", func(t *testing.T) {
		networkSettingsID := testAccEnterpriseNetworkConfigurationID(t)
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		resourceName := "github_enterprise_actions_runner_group.test"
		networkConfigurationResourceName := "github_enterprise_network_configuration.test"
		groupName := fmt.Sprintf("tf-acc-test-%s", randomID)
		networkConfigurationName := fmt.Sprintf("%senterprise-network-config-%s", testResourcePrefix, randomID)

		configWithoutNetworking := fmt.Sprintf(`
			resource "github_enterprise_network_configuration" "test" {
			  enterprise_slug      = %q
			  name                 = %q
			  compute_service      = "actions"
			  network_settings_ids = [%q]
			}

			resource "github_enterprise_actions_runner_group" "test" {
			  enterprise_slug = %q
			  name            = %q
			  visibility      = "all"
			}
		`, testAccConf.enterpriseSlug, networkConfigurationName, networkSettingsID, testAccConf.enterpriseSlug, groupName)

		configWithNetworking := fmt.Sprintf(`
			resource "github_enterprise_network_configuration" "test" {
			  enterprise_slug      = %q
			  name                 = %q
			  compute_service      = "actions"
			  network_settings_ids = [%q]
			}

			resource "github_enterprise_actions_runner_group" "test" {
			  enterprise_slug            = %q
			  name                       = %q
			  visibility                 = "all"
			  network_configuration_id   = github_enterprise_network_configuration.test.id
			}
		`, testAccConf.enterpriseSlug, networkConfigurationName, networkSettingsID, testAccConf.enterpriseSlug, groupName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessMode(t, enterprise) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: configWithoutNetworking,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(groupName)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("visibility"), knownvalue.StringExact("all")),
					},
				},
				{
					Config: configWithNetworking,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("network_configuration_id"), knownvalue.NotNull()),
						testCheckEnterpriseRunnerGroupNetworkConfigurationMatches(resourceName, networkConfigurationResourceName),
					},
				},
				{
					ResourceName:        resourceName,
					ImportState:         true,
					ImportStateVerify:   true,
					ImportStateIdPrefix: fmt.Sprintf(`%s/`, testAccConf.enterpriseSlug),
				},
				{
					Config: configWithoutNetworking,
				},
			},
		})
	})

	t.Run("creates runner group network configuration on create", func(t *testing.T) {
		networkSettingsID := testAccEnterpriseNetworkConfigurationID(t)
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		resourceName := "github_enterprise_actions_runner_group.test"
		networkConfigurationResourceName := "github_enterprise_network_configuration.test"
		groupName := fmt.Sprintf("tf-acc-test-create-%s", randomID)
		networkConfigurationName := fmt.Sprintf("%senterprise-network-config-create-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
			resource "github_enterprise_network_configuration" "test" {
			  enterprise_slug      = %q
			  name                 = %q
			  compute_service      = "actions"
			  network_settings_ids = [%q]
			}

			resource "github_enterprise_actions_runner_group" "test" {
			  enterprise_slug          = %q
			  name                     = %q
			  visibility               = "all"
			  network_configuration_id = github_enterprise_network_configuration.test.id
			}
		`, testAccConf.enterpriseSlug, networkConfigurationName, networkSettingsID, testAccConf.enterpriseSlug, groupName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessMode(t, enterprise) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("network_configuration_id"), knownvalue.NotNull()),
						testCheckEnterpriseRunnerGroupNetworkConfigurationMatches(resourceName, networkConfigurationResourceName),
					},
				},
				{
					ResourceName:        resourceName,
					ImportState:         true,
					ImportStateVerify:   true,
					ImportStateIdPrefix: fmt.Sprintf(`%s/`, testAccConf.enterpriseSlug),
				},
			},
		})
	})
}

func testAccEnterpriseNetworkConfigurationID(t *testing.T) string {
	t.Helper()

	networkSettingsID := os.Getenv("GITHUB_TEST_ENTERPRISE_NETWORK_SETTINGS_ID")
	if networkSettingsID == "" {
		t.Skip("GITHUB_TEST_ENTERPRISE_NETWORK_SETTINGS_ID not set")
	}

	return networkSettingsID
}
