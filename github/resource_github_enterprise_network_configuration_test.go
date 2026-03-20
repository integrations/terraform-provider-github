package github

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-github/v84/github"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubEnterpriseNetworkConfiguration(t *testing.T) {
	t.Run("create", func(t *testing.T) {
		networkSettingsID := testAccEnterpriseNetworkConfigurationID(t)

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		resourceName := "github_enterprise_network_configuration.test"
		configurationName := fmt.Sprintf("%senterprise-network-config-%s", testResourcePrefix, randomID)

		config := testAccEnterpriseNetworkConfigurationConfig(configurationName, "actions", networkSettingsID)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessMode(t, enterprise) },
			ProviderFactories: providerFactories,
			CheckDestroy:      testAccCheckGithubEnterpriseNetworkConfigurationDestroy,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enterprise_slug"), knownvalue.StringExact(testAccConf.enterpriseSlug)),
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
		networkSettingsID := testAccEnterpriseNetworkConfigurationID(t)

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		resourceName := "github_enterprise_network_configuration.test"
		beforeName := fmt.Sprintf("%senterprise-network-config-%s-a", testResourcePrefix, randomID)
		afterName := fmt.Sprintf("%senterprise-network-config-%s-b", testResourcePrefix, randomID)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessMode(t, enterprise) },
			ProviderFactories: providerFactories,
			CheckDestroy:      testAccCheckGithubEnterpriseNetworkConfigurationDestroy,
			Steps: []resource.TestStep{
				{
					Config: testAccEnterpriseNetworkConfigurationConfig(beforeName, "actions", networkSettingsID),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(beforeName)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("compute_service"), knownvalue.StringExact("actions")),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("network_settings_ids"), knownvalue.ListExact([]knownvalue.Check{knownvalue.StringExact(networkSettingsID)})),
					},
				},
				{
					Config: testAccEnterpriseNetworkConfigurationConfig(afterName, "none", networkSettingsID),
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
		networkSettingsID := testAccEnterpriseNetworkConfigurationID(t)

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		configurationName := fmt.Sprintf("%senterprise-network-config-%s", testResourcePrefix, randomID)

		config := testAccEnterpriseNetworkConfigurationConfig(configurationName, "actions", networkSettingsID)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessMode(t, enterprise) },
			ProviderFactories: providerFactories,
			CheckDestroy:      testAccCheckGithubEnterpriseNetworkConfigurationDestroy,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_enterprise_network_configuration.test", tfjsonpath.New("id"), knownvalue.NotNull()),
					},
				},
				{
					ResourceName:        "github_enterprise_network_configuration.test",
					ImportState:         true,
					ImportStateVerify:   true,
					ImportStateIdPrefix: fmt.Sprintf(`%s/`, testAccConf.enterpriseSlug),
				},
			},
		})
	})
}

func testAccCheckGithubEnterpriseNetworkConfigurationDestroy(s *terraform.State) error {
	meta, err := getTestMeta()
	if err != nil {
		return err
	}

	client := meta.v3client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "github_enterprise_network_configuration" {
			continue
		}

		enterpriseSlug := rs.Primary.Attributes["enterprise_slug"]
		if enterpriseSlug == "" {
			enterpriseSlug = testAccConf.enterpriseSlug
		}

		_, _, err := client.Enterprise.GetEnterpriseNetworkConfiguration(context.Background(), enterpriseSlug, rs.Primary.ID)
		if err != nil {
			var ghErr *github.ErrorResponse
			if errors.As(err, &ghErr) && ghErr.Response != nil && ghErr.Response.StatusCode == http.StatusNotFound {
				continue
			}

			return err
		}

		return fmt.Errorf("enterprise network configuration still exists: %s", rs.Primary.ID)
	}

	return nil
}

func testAccEnterpriseNetworkConfigurationConfig(name, computeService, networkSettingsID string) string {
	return fmt.Sprintf(`
resource "github_enterprise_network_configuration" "test" {
  enterprise_slug      = %q
  name                 = %q
  compute_service      = %q
  network_settings_ids = [%q]
}
`, testAccConf.enterpriseSlug, name, computeService, networkSettingsID)
}
