package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccDataSourceGithubEnterpriseAppInstallations(t *testing.T) {
	t.Parallel()

	skipUnlessEnterprise(t)

	t.Run("queries_app_installations", func(t *testing.T) {
		t.Parallel()

		config := fmt.Sprintf(`
data "github_enterprise_app_installations" "test" {
  enterprise_slug = "%s"
  organization    = "%s"
}
`, testAccConf.enterpriseSlug, testAccConf.owner)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_enterprise_app_installations.test", tfjsonpath.New("installations"), knownvalue.NotNull()),
					},
				},
			},
		})
	})
}
