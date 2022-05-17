package github

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccGithubDependabotOrganizationSecretRepositories(t *testing.T) {
	const ORG_SECRET_NAME = "ORG_SECRET_NAME"
	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
	secret_name, exists := os.LookupEnv(ORG_SECRET_NAME)

	t.Run("set repository allowlist for an organization secret", func(t *testing.T) {
		if !exists {
			t.Skipf("%s environment variable is missing", ORG_SECRET_NAME)
		}

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

			resource "github_dependabot_organization_secret_repositories" "org_secret_repos" {
				secret_name = "%s"
				selected_repository_ids = [
					github_repository.test_repo_1.repo_id,
					github_repository.test_repo_2.repo_id
				]
			}
		`, randomID, randomID, secret_name)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet(
				"github_dependabot_organization_secret_repositories.org_secret_repos", "secret_name",
			),
			resource.TestCheckResourceAttr(
				"github_dependabot_organization_secret_repositories.org_secret_repos", "selected_repository_ids.#", "2",
			),
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

		t.Run("with an anonymous account", func(t *testing.T) {
			t.Skip("anonymous account not supported for this operation")
		})

		t.Run("with an individual account", func(t *testing.T) {
			t.Skip("individual account not supported for this operation")
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})
	})
}
