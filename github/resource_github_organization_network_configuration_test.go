package github

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubOrganizationNetworkConfiguration(t *testing.T) {
	t.Run("create", func(t *testing.T) {
		networkSettingsID := testAccOrganizationNetworkConfigurationID(t)

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		resourceName := "github_organization_network_configuration.test"
		configurationName := fmt.Sprintf("%snetwork-config-%s", testResourcePrefix, randomID)

		config := testAccOrganizationNetworkConfigurationConfig(configurationName, "actions", networkSettingsID)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasPaidOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(configurationName)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("compute_service"), knownvalue.StringExact("actions")),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("network_settings_ids"), knownvalue.ListExact([]knownvalue.Check{knownvalue.StringExact(networkSettingsID)})),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("created_on"), knownvalue.NotNull()),
					},
				},
			},
		})
	})

	t.Run("update", func(t *testing.T) {
		networkSettingsID := testAccOrganizationNetworkConfigurationID(t)

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		resourceName := "github_organization_network_configuration.test"
		beforeName := fmt.Sprintf("%snetwork-config-%s-a", testResourcePrefix, randomID)
		afterName := fmt.Sprintf("%snetwork-config-%s-b", testResourcePrefix, randomID)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasPaidOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: testAccOrganizationNetworkConfigurationConfig(beforeName, "actions", networkSettingsID),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(beforeName)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("compute_service"), knownvalue.StringExact("actions")),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("network_settings_ids"), knownvalue.ListExact([]knownvalue.Check{knownvalue.StringExact(networkSettingsID)})),
					},
				},
				{
					Config: testAccOrganizationNetworkConfigurationConfig(afterName, "none", networkSettingsID),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(afterName)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("compute_service"), knownvalue.StringExact("none")),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("network_settings_ids"), knownvalue.ListExact([]knownvalue.Check{knownvalue.StringExact(networkSettingsID)})),
					},
				},
			},
		})
	})

	t.Run("import", func(t *testing.T) {
		networkSettingsID := testAccOrganizationNetworkConfigurationID(t)

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		configurationName := fmt.Sprintf("%snetwork-config-%s", testResourcePrefix, randomID)

		config := testAccOrganizationNetworkConfigurationConfig(configurationName, "actions", networkSettingsID)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasPaidOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
				},
				{
					ResourceName:      "github_organization_network_configuration.test",
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})
}

func testAccOrganizationNetworkConfigurationID(t *testing.T) string {
	t.Helper()

	networkSettingsID := os.Getenv("GITHUB_TEST_NETWORK_SETTINGS_ID")
	if networkSettingsID == "" {
		t.Skip("GITHUB_TEST_NETWORK_SETTINGS_ID not set")
	}

	return networkSettingsID
}

func testAccOrganizationNetworkConfigurationConfig(name, computeService, networkSettingsID string) string {
	return fmt.Sprintf(`
resource "github_organization_network_configuration" "test" {
  name                 = %q
  compute_service      = %q
  network_settings_ids = [%q]
}
`, name, computeService, networkSettingsID)
}
