package github

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccGithubClientConfigDataSource(t *testing.T) {
	t.Run("reads provider configuration without error", func(t *testing.T) {
		config := `data "github_client_config" "test" {}`

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("data.github_client_config.test", "owner", testAccConf.owner),
			resource.TestCheckResourceAttr("data.github_client_config.test", "username", testAccConf.username),
			resource.TestCheckResourceAttr("data.github_client_config.test", "base_url", testAccConf.baseURL.String()),
			resource.TestCheckResourceAttrSet("data.github_client_config.test", "is_organization"),
			resource.TestCheckResourceAttrSet("data.github_client_config.test", "id"),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  check,
				},
			},
		})
	})

	t.Run("reads provider configuration in anonymous mode", func(t *testing.T) {
		if testAccConf.authMode != anonymous {
			t.Skip("Skipping as test mode is not anonymous")
		}

		config := `data "github_client_config" "test" {}`

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("data.github_client_config.test", "owner", ""),
			resource.TestCheckResourceAttr("data.github_client_config.test", "username", ""),
			resource.TestCheckResourceAttr("data.github_client_config.test", "is_organization", "false"),
			resource.TestCheckResourceAttrSet("data.github_client_config.test", "base_url"),
			resource.TestCheckResourceAttrSet("data.github_client_config.test", "id"),
		)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  check,
				},
			},
		})
	})
}
