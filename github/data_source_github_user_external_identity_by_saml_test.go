package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccGithubUserExternalIdentityBySaml(t *testing.T) {
	t.Run("queries without error", func(t *testing.T) {
		config := `data "github_user_external_identity_by_saml" "test" { saml_name_id = "%s" }`

		check := resource.ComposeAggregateTestCheckFunc(
			resource.TestCheckResourceAttrSet("data.github_user_external_identity_by_saml.test", "login"),
			resource.TestCheckResourceAttrSet("data.github_user_external_identity_by_saml.test", "username"),
			resource.TestCheckResourceAttrSet("data.github_user_external_identity_by_saml.test", "saml_identity.name_id"),
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
