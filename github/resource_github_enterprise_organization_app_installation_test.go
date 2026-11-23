package github

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubEnterpriseOrganizationAppInstallation(t *testing.T) {
	appClientID := os.Getenv("GH_TEST_ENTERPRISE_APP_CLIENT_ID")
	skipUnlessEnterpriseAppClientID := func(t *testing.T) {
		t.Helper()
		if appClientID == "" {
			t.Skip("Skipping because GH_TEST_ENTERPRISE_APP_CLIENT_ID is not set")
		}
	}

	t.Run("installs an app on all repositories", func(t *testing.T) {
		config := fmt.Sprintf(`
			resource "github_enterprise_organization_app_installation" "test" {
				enterprise_slug      = "%s"
				organization         = "%s"
				client_id            = "%s"
				repository_selection = "all"
			}
		`, testAccConf.enterpriseSlug, testAccConf.owner, appClientID)

		resource.Test(t, resource.TestCase{
			PreCheck: func() {
				skipUnlessEnterprise(t)
				skipUnlessEnterpriseAppClientID(t)
			},
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_enterprise_organization_app_installation.test", tfjsonpath.New("enterprise_slug"), knownvalue.StringExact(testAccConf.enterpriseSlug)),
						statecheck.ExpectKnownValue("github_enterprise_organization_app_installation.test", tfjsonpath.New("organization"), knownvalue.StringExact(testAccConf.owner)),
						statecheck.ExpectKnownValue("github_enterprise_organization_app_installation.test", tfjsonpath.New("client_id"), knownvalue.StringExact(appClientID)),
						statecheck.ExpectKnownValue("github_enterprise_organization_app_installation.test", tfjsonpath.New("repository_selection"), knownvalue.StringExact("all")),
						statecheck.ExpectKnownValue("github_enterprise_organization_app_installation.test", tfjsonpath.New("installation_id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_enterprise_organization_app_installation.test", tfjsonpath.New("app_slug"), knownvalue.NotNull()),
					},
				},
				{
					ResourceName:      "github_enterprise_organization_app_installation.test",
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})

	t.Run("installs an app on selected repositories, replaces one and toggles to all", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

		repositories := fmt.Sprintf(`
			resource "github_repository" "test" {
				name      = "tf-acc-test-eai-%[1]s"
				auto_init = true
			}

			resource "github_repository" "other" {
				name      = "tf-acc-test-eai-other-%[1]s"
				auto_init = true
			}
		`, randomID)

		configSelected := repositories + fmt.Sprintf(`
			resource "github_enterprise_organization_app_installation" "test" {
				enterprise_slug       = "%[1]s"
				organization          = "%[2]s"
				client_id             = "%[3]s"
				repository_selection  = "selected"
				selected_repositories = [github_repository.test.name]
			}
		`, testAccConf.enterpriseSlug, testAccConf.owner, appClientID)

		configReplaced := repositories + fmt.Sprintf(`
			resource "github_enterprise_organization_app_installation" "test" {
				enterprise_slug       = "%[1]s"
				organization          = "%[2]s"
				client_id             = "%[3]s"
				repository_selection  = "selected"
				selected_repositories = [github_repository.other.name]
			}
		`, testAccConf.enterpriseSlug, testAccConf.owner, appClientID)

		configAll := repositories + fmt.Sprintf(`
			resource "github_enterprise_organization_app_installation" "test" {
				enterprise_slug      = "%[1]s"
				organization         = "%[2]s"
				client_id            = "%[3]s"
				repository_selection = "all"
			}
		`, testAccConf.enterpriseSlug, testAccConf.owner, appClientID)

		resource.Test(t, resource.TestCase{
			PreCheck: func() {
				skipUnlessEnterprise(t)
				skipUnlessEnterpriseAppClientID(t)
			},
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: configSelected,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_enterprise_organization_app_installation.test", tfjsonpath.New("repository_selection"), knownvalue.StringExact("selected")),
						statecheck.ExpectKnownValue("github_enterprise_organization_app_installation.test", tfjsonpath.New("selected_repositories"), knownvalue.SetExact([]knownvalue.Check{knownvalue.StringExact("tf-acc-test-eai-" + randomID)})),
					},
				},
				{
					Config: configReplaced,
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_enterprise_organization_app_installation.test", plancheck.ResourceActionUpdate),
						},
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_enterprise_organization_app_installation.test", tfjsonpath.New("repository_selection"), knownvalue.StringExact("selected")),
						statecheck.ExpectKnownValue("github_enterprise_organization_app_installation.test", tfjsonpath.New("selected_repositories"), knownvalue.SetExact([]knownvalue.Check{knownvalue.StringExact("tf-acc-test-eai-other-" + randomID)})),
					},
				},
				{
					Config: configAll,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_enterprise_organization_app_installation.test", tfjsonpath.New("repository_selection"), knownvalue.StringExact("all")),
						statecheck.ExpectKnownValue("github_enterprise_organization_app_installation.test", tfjsonpath.New("selected_repositories"), knownvalue.SetSizeExact(0)),
					},
				},
			},
		})
	})
}
