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

func TestAccGithubEnterpriseAppInstallation(t *testing.T) {
	appClientID := os.Getenv("GH_TEST_ENTERPRISE_APP_CLIENT_ID")
	skipUnlessEnterpriseAppClientID := func(t *testing.T) {
		t.Helper()
		if appClientID == "" {
			t.Skip("Skipping because GH_TEST_ENTERPRISE_APP_CLIENT_ID is not set")
		}
	}

	t.Run("installs an app on all repositories", func(t *testing.T) {
		config := fmt.Sprintf(`
			resource "github_enterprise_app_installation" "test" {
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
						statecheck.ExpectKnownValue("github_enterprise_app_installation.test", tfjsonpath.New("enterprise_slug"), knownvalue.StringExact(testAccConf.enterpriseSlug)),
						statecheck.ExpectKnownValue("github_enterprise_app_installation.test", tfjsonpath.New("organization"), knownvalue.StringExact(testAccConf.owner)),
						statecheck.ExpectKnownValue("github_enterprise_app_installation.test", tfjsonpath.New("client_id"), knownvalue.StringExact(appClientID)),
						statecheck.ExpectKnownValue("github_enterprise_app_installation.test", tfjsonpath.New("repository_selection"), knownvalue.StringExact("all")),
						statecheck.ExpectKnownValue("github_enterprise_app_installation.test", tfjsonpath.New("installation_id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_enterprise_app_installation.test", tfjsonpath.New("app_slug"), knownvalue.NotNull()),
					},
				},
				{
					ResourceName:      "github_enterprise_app_installation.test",
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})

	t.Run("installs an app on selected repositories and toggles to all", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

		configSelected := fmt.Sprintf(`
			resource "github_repository" "test" {
				name      = "tf-acc-test-eai-%[4]s"
				auto_init = true
			}

			resource "github_enterprise_app_installation" "test" {
				enterprise_slug       = "%[1]s"
				organization          = "%[2]s"
				client_id             = "%[3]s"
				repository_selection  = "selected"
				selected_repositories = [github_repository.test.name]
			}
		`, testAccConf.enterpriseSlug, testAccConf.owner, appClientID, randomID)

		configAll := fmt.Sprintf(`
			resource "github_repository" "test" {
				name      = "tf-acc-test-eai-%[4]s"
				auto_init = true
			}

			resource "github_enterprise_app_installation" "test" {
				enterprise_slug      = "%[1]s"
				organization         = "%[2]s"
				client_id            = "%[3]s"
				repository_selection = "all"
			}
		`, testAccConf.enterpriseSlug, testAccConf.owner, appClientID, randomID)

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
						statecheck.ExpectKnownValue("github_enterprise_app_installation.test", tfjsonpath.New("repository_selection"), knownvalue.StringExact("selected")),
						statecheck.ExpectKnownValue("github_enterprise_app_installation.test", tfjsonpath.New("selected_repositories"), knownvalue.SetSizeExact(1)),
					},
				},
				{
					Config: configAll,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_enterprise_app_installation.test", tfjsonpath.New("repository_selection"), knownvalue.StringExact("all")),
						statecheck.ExpectKnownValue("github_enterprise_app_installation.test", tfjsonpath.New("selected_repositories"), knownvalue.SetSizeExact(0)),
					},
				},
			},
		})
	})
}
