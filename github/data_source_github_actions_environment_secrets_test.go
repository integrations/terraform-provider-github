package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubActionsEnvironmentSecretsDataSource(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)
		env := mustCreateTestRepositoryEnvironment(t, repo)
		secretName := mustCreateTestRepositoryEnvironmentSecret(t, repo, env, "super_secret_value")

		config := fmt.Sprintf(`
data "github_actions_environment_secrets" "test" {
  name        = "%s"
  environment = "%s"
}
`, repo.GetName(), env.GetName())

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_actions_environment_secrets.test", tfjsonpath.New("secrets"), knownvalue.ListExact([]knownvalue.Check{
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
