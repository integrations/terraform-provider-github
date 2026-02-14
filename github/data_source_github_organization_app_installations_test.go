package github

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// TestAccGithubOrganizationAppInstallations requires at least one GitHub App
// to be installed in the test organization (GITHUB_OWNER). If no apps are
// installed, the attribute checks below will fail.
func TestAccGithubOrganizationAppInstallations(t *testing.T) {
	config := `data "github_organization_app_installations" "test" {}`

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { skipUnlessHasOrgs(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.github_organization_app_installations.test", "installations.0.id"),
					resource.TestCheckResourceAttrSet("data.github_organization_app_installations.test", "installations.0.app_slug"),
					resource.TestCheckResourceAttrSet("data.github_organization_app_installations.test", "installations.0.app_id"),
				),
			},
		},
	})
}
