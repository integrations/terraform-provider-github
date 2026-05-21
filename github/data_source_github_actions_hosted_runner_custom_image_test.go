package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccGithubActionsHostedRunnerCustomImageDataSource(t *testing.T) {
	t.Run("queries a single custom image for hosted runners", func(t *testing.T) {
		// This test requires a custom image to exist in the org.
		// Set GITHUB_HOSTED_RUNNER_CUSTOM_IMAGE_ID env var to a valid image ID.
		imageID := testAccGetCustomImageID(t)

		config := fmt.Sprintf(`
			data "github_actions_hosted_runner_custom_image" "test" {
				image_id = %s
			}
		`, imageID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet("data.github_actions_hosted_runner_custom_image.test", "name"),
			resource.TestCheckResourceAttrSet("data.github_actions_hosted_runner_custom_image.test", "platform"),
			resource.TestCheckResourceAttrSet("data.github_actions_hosted_runner_custom_image.test", "state"),
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

func testAccGetCustomImageID(t *testing.T) string {
	t.Helper()
	id := testOrganization + "_CUSTOM_IMAGE_ID"
	// For CI, we'd look up from env. For now, skip if not available.
	t.Skipf("Skipping: requires a custom image in the org (set %s)", id)
	return ""
}
