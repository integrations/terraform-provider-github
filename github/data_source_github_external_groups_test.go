package github

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccGithubExternalGroupsDataSource(t *testing.T) {
	t.Run("queries_all", func(t *testing.T) {
		config := `
			data "github_external_groups" "all" {}
		`

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEMUEnterprise(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttrSet("data.github_external_groups.all", "external_groups.#"),
					),
				},
			},
		})
	})

	t.Run("queries_with_display_name_filter", func(t *testing.T) {
		config := `
			data "github_external_groups" "filtered" {
				display_name = "test"
			}
		`

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEMUEnterprise(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttrSet("data.github_external_groups.filtered", "external_groups.#"),
					),
				},
			},
		})
	})
}
