package github

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubOrganizationIpAllowListDataSource(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		config := `
data "github_organization_ip_allow_list" "test" {}
`

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_organization_ip_allow_list.test", tfjsonpath.New("ip_allow_list"), knownvalue.NotNull()),
					},
				},
			},
		})
	})
}
