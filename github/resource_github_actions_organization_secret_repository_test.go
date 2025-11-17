package github

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubActionsOrganizationSecretRepository(t *testing.T) {
	const ORG_SECRET_NAME = "ORG_SECRET_NAME"
	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
	secret_name, exists := os.LookupEnv(ORG_SECRET_NAME)

	t.Run("set repository allowlist for a organization secret", func(t *testing.T) {
		if !exists {
			t.Skipf("%s environment variable is missing", ORG_SECRET_NAME)
		}

		config := fmt.Sprintf(`
			resource "github_repository" "test_repo_1" {
				name = "tf-acc-test-%s-1"
				visibility = "internal"
				vulnerability_alerts = "true"
			}

			resource "github_actions_organization_secret_repository" "org_secret_repo" {
				secret_name = "%s"
				repository_id = github_repository.test_repo_1.repo_id
			}
		`, randomID, secret_name)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet(
				"github_actions_organization_secret_repository.org_secret_repo", "secret_name",
			),
			resource.TestCheckResourceAttr(
				"github_actions_organization_secret_repository.org_secret_repo", "repository_id.#", "1",
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
