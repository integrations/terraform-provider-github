package github

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubOrganizationNetworkConfiguration(t *testing.T) {
	t.Run("creates organization network configuration without error", func(t *testing.T) {
		networkSettingsID := os.Getenv("GITHUB_TEST_NETWORK_SETTINGS_ID")
		if networkSettingsID == "" {
			t.Skip("Skipping test: GITHUB_TEST_NETWORK_SETTINGS_ID not set")
		}

		config := fmt.Sprintf(`
		resource "github_organization_network_configuration" "test" {
			name                  = "test-network-configuration"
			compute_service       = "actions"
			network_settings_ids  = ["%s"]
		}`, networkSettingsID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_organization_network_configuration.test",
				"name", "test-network-configuration",
			),
			resource.TestCheckResourceAttr(
				"github_organization_network_configuration.test",
				"compute_service", "actions",
			),
			resource.TestCheckResourceAttr(
				"github_organization_network_configuration.test",
				"network_settings_ids.0", networkSettingsID,
			),
			resource.TestCheckResourceAttrSet(
				"github_organization_network_configuration.test",
				"id",
			),
			resource.TestCheckResourceAttrSet(
				"github_organization_network_configuration.test",
				"created_on",
			),
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

		t.Run("with an anonymous account", func(t *testing.T) {
			t.Skip("anonymous account not supported for this operation")
		})

		t.Run("with an individual account", func(t *testing.T) {
			t.Skip("individual account not supported for this operation")
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})
	})

	t.Run("updates organization network configuration without error", func(t *testing.T) {
		networkSettingsID := os.Getenv("GITHUB_TEST_NETWORK_SETTINGS_ID")
		if networkSettingsID == "" {
			t.Skip("Skipping test: GITHUB_TEST_NETWORK_SETTINGS_ID not set")
		}

		name := "test-network-config-one"
		computeService := "actions"

		updatedName := "test-network-config-two"
		updatedComputeService := "actions"
		updatedNetworkSettingsID := networkSettingsID

		configs := map[string]string{
			"before": fmt.Sprintf(`
			resource "github_organization_network_configuration" "test" {
				name                  = "%s"
				compute_service       = "%s"
				network_settings_ids  = ["%s"]
			}`, name, computeService, networkSettingsID),

			"after": fmt.Sprintf(`
			resource "github_organization_network_configuration" "test" {
				name                  = "%s"
				compute_service       = "%s"
				network_settings_ids  = ["%s"]
			}`, updatedName, updatedComputeService, updatedNetworkSettingsID),
		}

		checks := map[string]resource.TestCheckFunc{
			"before": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_organization_network_configuration.test",
					"name", name,
				),
				resource.TestCheckResourceAttr(
					"github_organization_network_configuration.test",
					"compute_service", computeService,
				),
				resource.TestCheckResourceAttr(
					"github_organization_network_configuration.test",
					"network_settings_ids.0", networkSettingsID,
				),
			),
			"after": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_organization_network_configuration.test",
					"name", updatedName,
				),
				resource.TestCheckResourceAttr(
					"github_organization_network_configuration.test",
					"compute_service", updatedComputeService,
				),
				resource.TestCheckResourceAttr(
					"github_organization_network_configuration.test",
					"network_settings_ids.0", updatedNetworkSettingsID,
				),
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

		t.Run("with an anonymous account", func(t *testing.T) {
			t.Skip("anonymous account not supported for this operation")
		})

		t.Run("with an individual account", func(t *testing.T) {
			t.Skip("individual account not supported for this operation")
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})
	})

	t.Run("imports organization network configuration without error", func(t *testing.T) {
		networkSettingsID := os.Getenv("GITHUB_TEST_NETWORK_SETTINGS_ID")
		if networkSettingsID == "" {
			t.Skip("Skipping test: GITHUB_TEST_NETWORK_SETTINGS_ID not set")
		}

		name := "test-network-config-import"
		computeService := "actions"

		config := fmt.Sprintf(`
		resource "github_organization_network_configuration" "test" {
			name                  = "%s"
			compute_service       = "%s"
			network_settings_ids  = ["%s"]
		}`, name, computeService, networkSettingsID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_organization_network_configuration.test",
				"name", name,
			),
			resource.TestCheckResourceAttr(
				"github_organization_network_configuration.test",
				"compute_service", computeService,
			),
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
						ResourceName:      "github_organization_network_configuration.test",
						ImportState:       true,
						ImportStateVerify: true,
					},
				},
			})
		}

		t.Run("with an anonymous account", func(t *testing.T) {
			t.Skip("anonymous account not supported for this operation")
		})

		t.Run("with an individual account", func(t *testing.T) {
			t.Skip("individual account not supported for this operation")
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})
	})
}
