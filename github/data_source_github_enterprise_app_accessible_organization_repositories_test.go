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

func TestAccDataSourceGithubEnterpriseAppAccessibleOrganizationRepositories(t *testing.T) {
	t.Parallel()

	skipUnlessEnterprise(t)

	t.Run("queries_accessible_repositories", func(t *testing.T) {
		t.Parallel()

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("tf-acc-test-eaaor-%s", randomID)

		config := fmt.Sprintf(`
resource "github_repository" "test" {
  name      = "%s"
  auto_init = true
}

data "github_enterprise_app_accessible_organization_repositories" "test" {
  enterprise_slug = "%s"
  organization    = "%s"

  depends_on = [github_repository.test]
}
`, repoName, testAccConf.enterpriseSlug, testAccConf.owner)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_enterprise_app_accessible_organization_repositories.test", tfjsonpath.New("repositories"), knownvalue.SetPartial([]knownvalue.Check{
							knownvalue.MapPartial(map[string]knownvalue.Check{
								"name": knownvalue.StringExact(repoName),
							}),
						})),
					},
				},
			},
		})
	})
}
