package github

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccDataSourceGithubOrganizationRepositories(t *testing.T) {
	t.Parallel()

	skipUnlessHasOrgs(t)

	t.Run("queries_all_repositories", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)

		config := `
data "github_organization_repositories" "test" {}
`

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_organization_repositories.test", tfjsonpath.New("repositories"), knownvalue.SetPartial([]knownvalue.Check{
							knownvalue.MapExact(map[string]knownvalue.Check{
								"id":         knownvalue.Int32Exact(int32(repo.GetID())),
								"node_id":    knownvalue.StringExact(repo.GetNodeID()),
								"name":       knownvalue.StringExact(repo.GetName()),
								"visibility": knownvalue.StringExact(repo.GetVisibility()),
								"archived":   knownvalue.Bool(repo.GetArchived()),
							}),
						})),
					},
				},
			},
		})
	})
}
