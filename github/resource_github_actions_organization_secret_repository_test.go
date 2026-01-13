package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubActionsOrganizationSecretRepository(t *testing.T) {
	t.Run("set repository allowlist for a organization secret", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-act-org-secret-%s", testResourcePrefix, randomID)
		secretName := testAccConf.testOrgSecretName
		if len(secretName) == 0 {
			t.Skip("test organization secret name is not set")
		}

		config := fmt.Sprintf(`
			resource "github_repository" "test_repo_1" {
				name = "%s"
				visibility = "internal"
				vulnerability_alerts = "true"
			}

			resource "github_actions_organization_secret_repository" "org_secret_repo" {
				secret_name = "%s"
				repository_id = github_repository.test_repo_1.repo_id
			}
		`, repoName, secretName)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet(
				"github_actions_organization_secret_repository.org_secret_repo", "secret_name",
			),
			resource.TestCheckResourceAttr(
				"github_actions_organization_secret_repository.org_secret_repo", "repository_id.#", "1",
			),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:  func() { skipUnlessHasOrgs(t) },
			Providers: testAccProviders,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  check,
				},
			},
		})
	})
}
