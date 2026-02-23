package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubEnterpriseSCIMUsersDataSource(t *testing.T) {
	t.Run("lists users without error", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{{
				Config: fmt.Sprintf(`
					data "github_enterprise_scim_users" "test" {
						enterprise = "%s"
					}
				`, testAccConf.enterpriseSlug),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("data.github_enterprise_scim_users.test", tfjsonpath.New("total_results"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("data.github_enterprise_scim_users.test", tfjsonpath.New("schemas"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("data.github_enterprise_scim_users.test", tfjsonpath.New("resources"), knownvalue.NotNull()),
				},
			}},
		})
	})
}
