package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubBranchProtectionRulesDataSource(t *testing.T) {
	t.Parallel()
	skipUnauthenticated(t)

	t.Run("gets_empty_branch_protection_rules", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)

		config := fmt.Sprintf(`
data "github_branch_protection_rules" "test" {
  repository = "%s"
}
`, repo.GetName())

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_branch_protection_rules.test", tfjsonpath.New("rules"), knownvalue.ListSizeExact(0)),
					},
				},
			},
		})
	})

	t.Run("queries_branch_protection", func(t *testing.T) {
		t.Parallel()

		pattern := "main"
		repo := mustCreateTestRepository(t)
		_ = mustCreateTestRepositoryBranchProtection(t, repo, pattern)

		config := fmt.Sprintf(`
data "github_branch_protection_rules" "test" {
  repository = "%s"
}
`, repo.GetName())

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_branch_protection_rules.test", tfjsonpath.New("rules"), knownvalue.ListExact([]knownvalue.Check{
							knownvalue.MapExact(map[string]knownvalue.Check{
								"pattern": knownvalue.StringExact(pattern),
							}),
						})),
					},
				},
			},
		})
	})
}
