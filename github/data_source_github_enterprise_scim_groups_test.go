package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubEnterpriseSCIMGroupsDataSource(t *testing.T) {
	t.Run("lists groups without error", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{{
				Config: fmt.Sprintf(`
					data "github_enterprise_scim_groups" "test" {
						enterprise = "%s"
					}
				`, testAccConf.enterpriseSlug),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("data.github_enterprise_scim_groups.test", tfjsonpath.New("total_results"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("data.github_enterprise_scim_groups.test", tfjsonpath.New("schemas"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("data.github_enterprise_scim_groups.test", tfjsonpath.New("resources"), knownvalue.NotNull()),
				},
			}},
		})
	})
}
