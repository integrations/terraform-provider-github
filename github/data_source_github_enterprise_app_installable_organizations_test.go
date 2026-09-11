package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccDataSourceGithubEnterpriseAppInstallableOrganizations(t *testing.T) {
	t.Parallel()

	skipUnlessEnterprise(t)

	t.Run("queries_installable_organizations", func(t *testing.T) {
		t.Parallel()

		config := fmt.Sprintf(`
data "github_enterprise_app_installable_organizations" "test" {
  enterprise_slug = "%s"
}
`, testAccConf.enterpriseSlug)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_enterprise_app_installable_organizations.test", tfjsonpath.New("organizations"), knownvalue.SetPartial([]knownvalue.Check{
							knownvalue.MapPartial(map[string]knownvalue.Check{
								"login": knownvalue.StringExact(testAccConf.owner),
							}),
						})),
					},
				},
			},
		})
	})
}
