package github

import (
	"fmt"
	"testing"

	"github.com/google/go-github/v93/github"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccGithubActionsOrganizationSecretRepository(t *testing.T) {
	t.Parallel()
	skipUnlessHasOrgs(t)

	t.Run("create", func(t *testing.T) {
		t.Parallel()

		secretName := mustCreateTestOrganizationSecret(t, func(req *github.SecretOrgRequest) { req.Visibility = "selected" })
		repo := mustCreateTestRepository(t)

		config := fmt.Sprintf(`
resource "github_actions_organization_secret_repository" "test" {
  secret_name   = "%s"
  repository_id = %d
}
`, secretName, repo.GetID())

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
				},
			},
		})
	})
}
