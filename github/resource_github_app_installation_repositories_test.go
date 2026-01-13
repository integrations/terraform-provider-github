package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubAppInstallationRepositories(t *testing.T) {
	t.Skip("TODO: Broken test")
	if testAccConf.testOrgAppInstallationId == 0 {
		t.Skip("No org app installation id provided")
	}

	t.Run("installs an app to multiple repositories", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName0 := fmt.Sprintf("%srepo-app-install-0-%s", testResourcePrefix, randomID)
		repoName1 := fmt.Sprintf("%srepo-app-install-1-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
		resource "github_repository" "test_0" {
			name      = "%s"
			auto_init = true
		}

		resource "github_repository" "test_1" {
			name      = "%s"
			auto_init = true
		}

		resource "github_app_installation_repositories" "test" {
			# The installation id of the app (in the organization).
			installation_id         = "%d"
			selected_repositories = [github_repository.test_0.name, github_repository.test_1.name]
		}

		`, repoName0, repoName1, testAccConf.testOrgAppInstallationId)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet(
				"github_app_installation_repositories.test", "installation_id",
			),
			resource.TestCheckResourceAttr(
				"github_app_installation_repositories.test", "selected_repositories.#", "2",
			),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  check,
				},
			},
		})
	})
}
