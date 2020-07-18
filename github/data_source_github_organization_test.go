package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccGithubOrganizationDataSource(t *testing.T) {

	t.Run("queries for an organization without error", func(t *testing.T) {

		organizationConfiguration := fmt.Sprintf(`
			provider "github" {
				organization = "%s"
				token = "%s"
			}
			data "github_organization" "test" { name = "%s" }
		`, testOrganization, testToken, testOrganization)

		organizationCheck := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet("data.github_organization.test", "login"),
			resource.TestCheckResourceAttrSet("data.github_organization.test", "name"),
			resource.TestCheckResourceAttrSet("data.github_organization.test", "description"),
			resource.TestCheckResourceAttrSet("data.github_organization.test", "plan"),
		)

		resource.Test(t, resource.TestCase{
			PreCheck: func() {
				requiredEnvironmentVariables := []string{
					"GITHUB_TOKEN",
					"GITHUB_ORGANIZATION",
				}
				testAccPreCheckEnvironment(t, requiredEnvironmentVariables)
			},
			Providers: testAccProviders,
			Steps: []resource.TestStep{
				{
					Config: organizationConfiguration,
					Check:  organizationCheck,
				},
			},
		})

	})
}
