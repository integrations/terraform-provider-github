package github

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/compare"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestGithubActionsEnterpriseRunnerGroupNetworking(t *testing.T) {
	testRunnerGroupNetworking(t, resourceGithubActionsEnterpriseRunnerGroup(), "/enterprises/test-enterprise", map[string]any{
		"enterprise_slug":          "test-enterprise",
		"name":                     "test-group",
		"visibility":               "all",
		"network_configuration_id": "network-1",
	}, resourceGithubActionsEnterpriseRunnerGroupCreate, resourceGithubActionsEnterpriseRunnerGroupUpdate)
}

func TestGithubActionsEnterpriseRunnerGroupReadErrors(t *testing.T) {
	for _, tc := range []struct {
		name       string
		statusCode int
		wantID     string
		wantError  bool
	}{
		{"not modified", http.StatusNotModified, "42", false},
		{"not found", http.StatusNotFound, "", false},
		{"forbidden", http.StatusForbidden, "42", true},
		{"server error", http.StatusInternalServerError, "42", true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			server := githubApiMock([]*mockResponse{{
				ExpectedUri:    "/enterprises/test-enterprise/actions/runner-groups/42",
				ExpectedMethod: http.MethodGet,
				StatusCode:     tc.statusCode,
			}})
			defer server.Close()

			meta := &Owner{v3client: mustCreateTestGitHubClient(t, server.URL)}
			d := schema.TestResourceDataRaw(t, resourceGithubActionsEnterpriseRunnerGroup().Schema, map[string]any{
				"enterprise_slug":          "test-enterprise",
				"name":                     "test-group",
				"visibility":               "all",
				"network_configuration_id": "network-1",
			})
			d.SetId("42")

			err := resourceGithubActionsEnterpriseRunnerGroupRead(d, meta)
			if (err != nil) != tc.wantError {
				t.Fatalf("read error = %v, want error = %t", err, tc.wantError)
			}
			if d.Id() != tc.wantID {
				t.Errorf("ID = %q, want %q", d.Id(), tc.wantID)
			}
			if got := d.Get("network_configuration_id"); got != "network-1" {
				t.Errorf("network_configuration_id = %v, want network-1", got)
			}
		})
	}
}

func TestAccGithubActionsEnterpriseRunnerGroup(t *testing.T) {
	t.Parallel()

	t.Run("creates enterprise runner groups without error", func(t *testing.T) {
		t.Parallel()

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
			PreCheck:          func() { skipUnlessEnterprise(t) },
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
		t.Parallel()

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
			PreCheck:          func() { skipUnlessEnterprise(t) },
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
		t.Parallel()

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
			PreCheck:          func() { skipUnlessEnterprise(t) },
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
		t.Parallel()

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
			PreCheck:          func() { skipUnlessEnterprise(t) },
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
		networkConfiguration := mustCreateTestEnterpriseNetworkConfiguration(t)
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		resourceName := "github_enterprise_actions_runner_group.test"
		groupName := fmt.Sprintf("tf-acc-test-%s", randomID)
		sameID := statecheck.CompareValue(compare.ValuesSame())

		configWithoutNetworking := fmt.Sprintf(`
			resource "github_enterprise_actions_runner_group" "test" {
			  enterprise_slug = %q
			  name            = %q
			  visibility      = "all"
			}
		`, testAccConf.enterpriseSlug, groupName)

		configWithNetworking := fmt.Sprintf(`
			resource "github_enterprise_actions_runner_group" "test" {
			  enterprise_slug          = %q
			  name                     = %q
			  visibility               = "all"
			  network_configuration_id = %q
			}
		`, testAccConf.enterpriseSlug, groupName, networkConfiguration.GetID())

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: configWithoutNetworking,
					ConfigStateChecks: []statecheck.StateCheck{
						sameID.AddStateValue(resourceName, tfjsonpath.New("id")),
					},
				},
				{
					Config: configWithNetworking,
					ConfigStateChecks: []statecheck.StateCheck{
						sameID.AddStateValue(resourceName, tfjsonpath.New("id")),
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
		networkConfiguration := mustCreateTestEnterpriseNetworkConfiguration(t)
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		resourceName := "github_enterprise_actions_runner_group.test"
		groupName := fmt.Sprintf("tf-acc-test-create-%s", randomID)

		config := fmt.Sprintf(`
			resource "github_enterprise_actions_runner_group" "test" {
			  enterprise_slug          = %q
			  name                     = %q
			  visibility               = "all"
			  network_configuration_id = %q
			}
		`, testAccConf.enterpriseSlug, groupName, networkConfiguration.GetID())

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("runners_url"), knownvalue.NotNull()),
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
