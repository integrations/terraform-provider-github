package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubEnterpriseSCIMUserDataSource(t *testing.T) {
	t.Run("reads user without error", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{{
				Config: fmt.Sprintf(`
					data "github_enterprise_scim_users" "all" {
						enterprise = "%s"
					}

					data "github_enterprise_scim_user" "test" {
						enterprise   = "%[1]s"
						scim_user_id = data.github_enterprise_scim_users.all.resources[0].id
					}
				`, testAccConf.enterpriseSlug),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("data.github_enterprise_scim_user.test", tfjsonpath.New("user_name"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("data.github_enterprise_scim_user.test", tfjsonpath.New("display_name"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("data.github_enterprise_scim_user.test", tfjsonpath.New("schemas"), knownvalue.NotNull()),
					statecheck.ExpectKnownValue("data.github_enterprise_scim_user.test", tfjsonpath.New("emails"), knownvalue.NotNull()),
				},
			}},
		})
	})
}
