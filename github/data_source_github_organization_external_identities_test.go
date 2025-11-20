package github

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubOrganizationExternalIdentities(t *testing.T) {
	if isEnterprise != "true" {
		t.Skip("Skipping because `ENTERPRISE_ACCOUNT` is not set or set to false")
	}

	t.Run("queries without error", func(t *testing.T) {
		config := `data "github_organization_external_identities" "test" {}`

		check := resource.ComposeAggregateTestCheckFunc(
			resource.TestCheckResourceAttrSet("data.github_organization_external_identities.test", "identities.#"),
			resource.TestCheckResourceAttrSet("data.github_organization_external_identities.test", "identities.0.login"),
			resource.TestCheckResourceAttrSet("data.github_organization_external_identities.test", "identities.0.saml_identity.name_id"),
		)

		testCase := func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: config,
						Check:  check,
					},
				},
			})
		}

		t.Run("with an anonymous account", func(t *testing.T) {
			t.Skip("anonymous account not supported for this operation")
		})

		t.Run("with an individual account", func(t *testing.T) {
			t.Skip("individual account not supported for this operation")
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t)
		})
	})
}
