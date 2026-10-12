package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubRefDataSource(t *testing.T) {
	t.Parallel()
	skipUnauthenticated(t)

	t.Run("queries_an_existing_branch_ref", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)

		config := fmt.Sprintf(`
data "github_ref" "test" {
  repository = "%s"
  ref        = "heads/main"
}
`, repo.GetName())

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_ref.test", tfjsonpath.New("id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_ref.test", tfjsonpath.New("sha"), knownvalue.NotNull()),
					},
				},
			},
		})
	})

	t.Run("queries_an_existing_branch_ref_with_owner", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)

		config := fmt.Sprintf(`
	data "github_ref" "test" {
  owner      = "%s"
  repository = "%s"
  ref        = "heads/main"
	}
	`, testAccConf.owner, repo.GetName())

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_ref.test", tfjsonpath.New("id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_ref.test", tfjsonpath.New("sha"), knownvalue.NotNull()),
					},
				},
			},
		})
	})

	t.Run("queries_an_invalid_ref_without_error", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)

		config := fmt.Sprintf(`
data "github_ref" "test" {
  repository = "%s"
  ref        = "heads/xxxxxx"
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
