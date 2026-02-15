package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccGithubUserExternalIdentity(t *testing.T) {
	t.Run("queries without error", func(t *testing.T) {
		config := `data "github_user_external_identity" "test" { username = "%s" }`

		check := resource.ComposeAggregateTestCheckFunc(
			resource.TestCheckResourceAttrSet("data.github_user_external_identity.test", "login"),
			resource.TestCheckResourceAttrSet("data.github_user_external_identity.test", "saml_identity.name_id"),
			resource.TestCheckResourceAttrSet("data.github_user_external_identity.test", "scim_identity.username"),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessMode(t, enterprise) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, testAccConf.testExternalUser),
					Check:  check,
				},
			},
		})
	})
}
