package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubActionsOrganizationSecretRepositories(t *testing.T) {
	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
	repoName1 := fmt.Sprintf("%srepo-act-org-secret-%s-1", testResourcePrefix, randomID)
	repoName2 := fmt.Sprintf("%srepo-act-org-secret-%s-2", testResourcePrefix, randomID)

	t.Run("set repository allowlist for a organization secret", func(t *testing.T) {
		if len(testAccConf.testOrgSecretName) == 0 {
			t.Skipf("'GH_TEST_ORG_SECRET_NAME' environment variable is missing")
		}

		config := fmt.Sprintf(`
			resource "github_repository" "test_repo_1" {
				name = "%s"
				visibility = "internal"
				vulnerability_alerts = "true"
			}

			resource "github_repository" "test_repo_2" {
				name = "%s"
				visibility = "internal"
				vulnerability_alerts = "true"
			}

			resource "github_actions_organization_secret_repositories" "org_secret_repos" {
				secret_name = "%s"
				selected_repository_ids = [
					github_repository.test_repo_1.repo_id,
					github_repository.test_repo_2.repo_id
				]
			}
		`, repoName1, repoName2, testAccConf.testOrgSecretName)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet(
				"github_actions_organization_secret_repositories.org_secret_repos", "secret_name",
			),
			resource.TestCheckResourceAttr(
				"github_actions_organization_secret_repositories.org_secret_repos", "selected_repository_ids.#", "2",
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
