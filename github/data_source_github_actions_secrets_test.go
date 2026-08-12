package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubActionsSecretsDataSource(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)
		secretName := mustCreateTestRepositorySecret(t, repo)

		config := fmt.Sprintf(`
data "github_actions_secrets" "test" {
  name = "%s"
}
`, repo.GetName())
		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_actions_secrets.test", tfjsonpath.New("secrets"), knownvalue.ListExact([]knownvalue.Check{
							knownvalue.MapExact(map[string]knownvalue.Check{
								"name":       knownvalue.StringExact(secretName),
								"created_at": knownvalue.NotNull(),
								"updated_at": knownvalue.NotNull(),
							}),
						})),
					},
				},
			},
		})
	})
}
