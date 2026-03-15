package github

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccGithubOrganizationNetworkConfiguration(t *testing.T) {
	t.Run("create", func(t *testing.T) {
		networkSettingsID := os.Getenv("GITHUB_TEST_NETWORK_SETTINGS_ID")
		if networkSettingsID == "" {
			t.Skip("GITHUB_TEST_NETWORK_SETTINGS_ID not set")
		}

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		resourceName := "github_organization_network_configuration.test"
		configurationName := fmt.Sprintf("%snetwork-config-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
resource "github_organization_network_configuration" "test" {
  name                 = %q
  compute_service      = "actions"
  network_settings_ids = [%q]
}
`, configurationName, networkSettingsID)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessMode(t, organization) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr(resourceName, "name", configurationName),
						resource.TestCheckResourceAttr(resourceName, "compute_service", "actions"),
						resource.TestCheckResourceAttr(resourceName, "network_settings_ids.0", networkSettingsID),
						resource.TestCheckResourceAttrSet(resourceName, "id"),
						resource.TestCheckResourceAttrSet(resourceName, "created_on"),
					),
				},
			},
		})
	})

	t.Run("update", func(t *testing.T) {
		networkSettingsID := os.Getenv("GITHUB_TEST_NETWORK_SETTINGS_ID")
		if networkSettingsID == "" {
			t.Skip("GITHUB_TEST_NETWORK_SETTINGS_ID not set")
		}

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		resourceName := "github_organization_network_configuration.test"
		beforeName := fmt.Sprintf("%snetwork-config-%s-a", testResourcePrefix, randomID)
		afterName := fmt.Sprintf("%snetwork-config-%s-b", testResourcePrefix, randomID)

		config := `
resource "github_organization_network_configuration" "test" {
  name                 = %q
  compute_service      = %q
  network_settings_ids = [%q]
}
`

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessMode(t, organization) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, beforeName, "actions", networkSettingsID),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr(resourceName, "name", beforeName),
						resource.TestCheckResourceAttr(resourceName, "compute_service", "actions"),
						resource.TestCheckResourceAttr(resourceName, "network_settings_ids.0", networkSettingsID),
					),
				},
				{
					Config: fmt.Sprintf(config, afterName, "none", networkSettingsID),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr(resourceName, "name", afterName),
						resource.TestCheckResourceAttr(resourceName, "compute_service", "none"),
						resource.TestCheckResourceAttr(resourceName, "network_settings_ids.0", networkSettingsID),
					),
				},
			},
		})
	})

	t.Run("import", func(t *testing.T) {
		networkSettingsID := os.Getenv("GITHUB_TEST_NETWORK_SETTINGS_ID")
		if networkSettingsID == "" {
			t.Skip("GITHUB_TEST_NETWORK_SETTINGS_ID not set")
		}

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		configurationName := fmt.Sprintf("%snetwork-config-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
resource "github_organization_network_configuration" "test" {
  name                 = %q
  compute_service      = "actions"
  network_settings_ids = [%q]
}
`, configurationName, networkSettingsID)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessMode(t, organization) },
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
