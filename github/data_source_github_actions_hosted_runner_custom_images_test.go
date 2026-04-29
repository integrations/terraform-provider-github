package github

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccGithubActionsHostedRunnerCustomImagesDataSource(t *testing.T) {
	t.Run("queries custom images for hosted runners", func(t *testing.T) {
		config := `
			data "github_actions_hosted_runner_custom_images" "test" {
			}
		`

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet("data.github_actions_hosted_runner_custom_images.test", "images.#"),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasPaidOrgs(t) },
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
