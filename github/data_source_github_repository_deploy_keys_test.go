package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubRepositoryDeployKeysDataSource(t *testing.T) {
	t.Parallel()

	skipUnlessHasOrgs(t)

	t.Skip("Deploy keys aren't enabled in the testing enterprise.")

	t.Run("lists_deploy_keys", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)
		_ = mustCreateTestRepositoryDeployKey(t, repo)
		_ = mustCreateTestRepositoryDeployKey(t, repo)

		config := fmt.Sprintf(`
data "github_repository_deploy_keys" "test" {
  repository = "%s"
}
`, repo.GetName())

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_repository_deploy_keys.test", tfjsonpath.New("keys"), knownvalue.ListSizeExact(2)),
					},
				},
			},
		})
	})
}
