package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubBranchDataSource(t *testing.T) {
	t.Parallel()
	skipUnauthenticated(t)

	t.Run("queries_an_existing_branch_without_error", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)

		config := fmt.Sprintf(`
data "github_branch" "test" {
  repository = "%s"
  branch     = "main"
}
`, repo.GetName())

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_branch.test", tfjsonpath.New("ref"), knownvalue.StringExact("refs/heads/main")),
					},
				},
			},
		})
	})

	t.Run("queries_an_invalid_branch_without_error", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)

		config := fmt.Sprintf(`
data "github_branch" "test" {
  repository = "%s"
  branch     = "xxxxxx"
}
`, repo.GetName())

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
