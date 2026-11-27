package github

import (
	"fmt"
	"testing"

	"github.com/google/go-github/v93/github"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccGithubDependabotOrganizationSecretRepositories(t *testing.T) {
	t.Parallel()
	skipUnlessHasOrgs(t)

	t.Run("create", func(t *testing.T) {
		t.Parallel()

		secretName := mustCreateTestDependabotOrganizationSecret(t, func(req *github.SecretOrgRequest) { req.Visibility = "selected" })
		repo1 := mustCreateTestRepository(t)
		repo2 := mustCreateTestRepository(t)

		config := fmt.Sprintf(`
resource "github_dependabot_organization_secret_repositories" "test" {
  secret_name = "%s"
  selected_repository_ids = [
    %d,
    %d
  ]
}
`, secretName, repo1.GetID(), repo2.GetID())

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
