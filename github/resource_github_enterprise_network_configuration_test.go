package github

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"testing"
	"time"

	"github.com/google/go-github/v89/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/compare"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubEnterpriseNetworkConfiguration(t *testing.T) {
	t.Run("create, import, and update in place", func(t *testing.T) {
		networkSettingsID := testAccEnterpriseNetworkConfigurationID(t)

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		resourceName := "github_enterprise_network_configuration.test"
		configurationName := fmt.Sprintf("%senterprise-network-config-%s", testResourcePrefix, randomID)

		sameID := statecheck.CompareValue(compare.ValuesSame())
		createdOn := knownvalue.StringFunc(func(value string) error {
			if value == "" {
				return nil
			}
			_, err := time.Parse(time.RFC3339, value)
			return err
		})

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
			ProviderFactories: providerFactories,
			CheckDestroy:      testAccCheckGithubEnterpriseNetworkConfigurationDestroy,
			Steps: []resource.TestStep{
				{
					Config: testAccEnterpriseNetworkConfigurationConfig(configurationName, "actions", networkSettingsID),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.StringRegexp(regexp.MustCompile(`^\S+$`))),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("created_on"), createdOn),
						sameID.AddStateValue(resourceName, tfjsonpath.New("id")),
					},
				},
				{
					ResourceName:        resourceName,
					ImportState:         true,
					ImportStateVerify:   true,
					ImportStateIdPrefix: testAccConf.enterpriseSlug + "/",
				},
				{
					Config: testAccEnterpriseNetworkConfigurationConfig(configurationName+"-updated", "none", networkSettingsID),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("created_on"), createdOn),
						sameID.AddStateValue(resourceName, tfjsonpath.New("id")),
					},
				},
				{
					ResourceName:        resourceName,
					ImportState:         true,
					ImportStateVerify:   true,
					ImportStateIdPrefix: testAccConf.enterpriseSlug + "/",
				},
			},
		})
	})
}

func TestGithubEnterpriseNetworkConfigurationImport(t *testing.T) {
	r := resourceGithubEnterpriseNetworkConfiguration()
	for _, id := range []string{"", "my-enterprise", "/NC_123", "my-enterprise/", "my-enterprise/NC_123/extra", " /NC_123", "my-enterprise/ "} {
		t.Run("invalid "+id, func(t *testing.T) {
			d := schema.TestResourceDataRaw(t, r.Schema, nil)
			d.SetId(id)
			if _, err := r.Importer.StateContext(context.Background(), d, nil); err == nil {
				t.Fatalf("import %q should fail", id)
			}
		})
	}

	d := schema.TestResourceDataRaw(t, r.Schema, nil)
	d.SetId("my-enterprise/NC_123")
	states, err := r.Importer.StateContext(context.Background(), d, nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(states) != 1 || states[0].Id() != "NC_123" || states[0].Get("enterprise_slug") != "my-enterprise" {
		t.Fatalf("import did not populate network configuration ID and enterprise slug: %v", states)
	}
}

func TestGithubEnterpriseNetworkConfigurationSlug(t *testing.T) {
	s := resourceGithubEnterpriseNetworkConfiguration().Schema["enterprise_slug"]
	if !s.ForceNew {
		t.Fatal("changing enterprise_slug must replace the network configuration")
	}
	for _, value := range []string{"", " \t"} {
		if !s.ValidateDiagFunc(value, nil).HasError() {
			t.Errorf("enterprise_slug %q should fail validation", value)
		}
	}
	if s.ValidateDiagFunc("my-enterprise", nil).HasError() {
		t.Error("valid enterprise_slug failed validation")
	}
}

func testAccCheckGithubEnterpriseNetworkConfigurationDestroy(s *terraform.State) error {
	client := testAccConf.meta.v3client

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
			if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response.StatusCode == http.StatusNotFound {
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
