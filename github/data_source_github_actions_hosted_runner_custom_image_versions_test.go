package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccGithubActionsHostedRunnerCustomImageVersionsDataSource(t *testing.T) {
	t.Run("queries versions of a custom image for hosted runners", func(t *testing.T) {
		imageID := testAccGetCustomImageIDForVersions(t)

		config := fmt.Sprintf(`
			data "github_actions_hosted_runner_custom_image_versions" "test" {
				image_id = %s
			}
		`, imageID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet("data.github_actions_hosted_runner_custom_image_versions.test", "versions.#"),
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

func testAccGetCustomImageIDForVersions(t *testing.T) string {
	t.Helper()
	id := testOrganization + "_CUSTOM_IMAGE_ID"
	t.Skipf("Skipping: requires a custom image in the org (set %s)", id)
	return ""
}
