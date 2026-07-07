package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccDataSourceGithubOrganizationRepositoryRoles(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		name := fmt.Sprintf("%s%s", testResourcePrefix, acctest.RandString(5))

		config := fmt.Sprintf(`
resource "github_organization_repository_role" "test" {
  name        = "%s"
  description = "Test role description"
  base_role   = "read"
  permissions = [
     "reopen_issue",
    "reopen_pull_request",
  ]
}

data "github_organization_repository_roles" "test" {
  depends_on = [ github_organization_repository_role.test ]
}
`, name)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_organization_repository_roles.test", tfjsonpath.New("roles"), knownvalue.ListPartial(map[int]knownvalue.Check{
							0: knownvalue.MapExact(map[string]knownvalue.Check{
								"role_id":     knownvalue.NotNull(),
								"name":        knownvalue.NotNull(),
								"description": knownvalue.NotNull(),
								"base_role":   knownvalue.NotNull(),
								"permissions": knownvalue.NotNull(),
							}),
						})),
					},
				},
			},
		})
	})
}
