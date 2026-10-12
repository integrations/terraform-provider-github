package github

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccGithubDependabotOrganizationPublicKeyDataSource(t *testing.T) {
	t.Parallel()
	skipUnlessHasOrgs(t)

	t.Run("queries_an_organization_public_key_without_error", func(t *testing.T) {
		t.Parallel()

		config := `
			data "github_dependabot_organization_public_key" "test" {}
		`

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  resource.TestCheckResourceAttrSet("data.github_dependabot_organization_public_key.test", "key"),
				},
			},
		})
	})
}
