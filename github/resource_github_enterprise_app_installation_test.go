package github

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
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

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("github_enterprise_app_installation.test", "enterprise_slug", testAccConf.enterpriseSlug),
			resource.TestCheckResourceAttr("github_enterprise_app_installation.test", "organization", testAccConf.owner),
			resource.TestCheckResourceAttr("github_enterprise_app_installation.test", "client_id", appClientID),
			resource.TestCheckResourceAttr("github_enterprise_app_installation.test", "repository_selection", "all"),
			resource.TestCheckResourceAttrSet("github_enterprise_app_installation.test", "installation_id"),
			resource.TestCheckResourceAttrSet("github_enterprise_app_installation.test", "app_slug"),
		)

		resource.Test(t, resource.TestCase{
			PreCheck: func() {
				skipUnlessEnterprise(t)
				skipUnlessEnterpriseAppClientID(t)
			},
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  check,
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
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_enterprise_app_installation.test", "repository_selection", "selected"),
						resource.TestCheckResourceAttr("github_enterprise_app_installation.test", "selected_repositories.#", "1"),
					),
				},
				{
					Config: configAll,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_enterprise_app_installation.test", "repository_selection", "all"),
						resource.TestCheckResourceAttr("github_enterprise_app_installation.test", "selected_repositories.#", "0"),
					),
				},
			},
		})
	})
}
