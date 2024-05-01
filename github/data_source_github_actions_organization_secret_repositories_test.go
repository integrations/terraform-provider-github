package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubActionsOrganizationSecretRepositoriesDataSource(t *testing.T) {

	const ORG_NAME = "ORG_NAME"
	const ORG_SECRET_NAME = "ORG_SECRET_NAME"
	const SECRET_NAME = "SECRET_NAME"
	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

	t.Run("list repositories for a organization secret", func(t *testing.T) {
		config := fmt.Sprintf(`
			resource "github_repository" "test_repo_1" {
				name = "tf-acc-test-%s-1"
				visibility = "internal"
				vulnerability_alerts = "true"
			}

			resource "github_repository" "test_repo_2" {
				name = "tf-acc-test-%s-2"
				visibility = "internal"
				vulnerability_alerts = "true"
			}

			resource "github_actions_organization_secret_repositories" "test_secret_repos" {
				secret_name 		= "%s"
				selected_repository_ids = [
					github_repository.test_repo_1.repo_id,
					github_repository.test_repo_2.repo_id
				]
			}

			data "github_actions_organization_secret_reposities" "test_secret_repos" {
				secret_name = github_actions_organization_secret_repositories.test_secret_repos.secret_name
			}
	`, randomID, randomID, SECRET_NAME)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("data.github_actions_organization_secret_repositories.test", "repositories.#", "2"),
			resource.TestCheckResourceAttr("data.github_actions_organization_secret_repositories.test", "repositories.0.name", fmt.Sprintf("tf-acc-test-%s-1", randomID)),
			resource.TestCheckResourceAttr("data.github_actions_organization_secret_repositories.test", "repositories.0.full_name", fmt.Sprintf("%s/tf-acc-test-%s-1", ORG_NAME, randomID)),
			resource.TestCheckResourceAttr("data.github_actions_organization_secret_repositories.test", "repositories.1.name", fmt.Sprintf("tf-acc-test-%s-2", randomID)),
			resource.TestCheckResourceAttr("data.github_actions_organization_secret_repositories.test", "repositories.1.full_name", fmt.Sprintf("%s/tf-acc-test-%s-2", ORG_NAME, randomID)),
		)

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: config,
						Check:  check,
					},
				},
			})
		}

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})
	})
}
