package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccGithubEnterpriseActionsHostedRunnersDataSource(t *testing.T) {
	t.Run("lists all enterprise hosted runners", func(t *testing.T) {
		config := fmt.Sprintf(`
			data "github_enterprise_actions_hosted_runners" "test" {
				enterprise_slug = "%s"
			}
		`, testAccConf.enterpriseSlug)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessMode(t, enterprise) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttrSet("data.github_enterprise_actions_hosted_runners.test", "runners.#"),
					),
				},
			},
		})
	})
}
