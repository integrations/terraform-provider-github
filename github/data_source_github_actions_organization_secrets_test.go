package github

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubActionsOrganizationSecretsDataSource(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		_ = mustCreateTestOrganizationSecret(t)

		config := `
data "github_actions_organization_secrets" "test" {}
`

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_actions_organization_secrets.test", tfjsonpath.New("secrets"), knownvalue.ListPartial(map[int]knownvalue.Check{
							0: knownvalue.MapPartial(map[string]knownvalue.Check{
								"name":       knownvalue.NotNull(),
								"visibility": knownvalue.NotNull(),
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
