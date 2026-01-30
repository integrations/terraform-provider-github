package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubReleaseAssetDataSource(t *testing.T) {
	testRepositoryOwner := testAccConf.testPublicRepositoryOwner
	testReleaseRepository := testAccConf.testPublicRepository
	testReleaseAssetID := testAccConf.testPublicRelaseAssetId
	testReleaseAssetName := testAccConf.testPublicRelaseAssetName
	testReleaseAssetContent := testAccConf.testPublicReleaseAssetContent

	t.Run("queries and downloads specified asset ID", func(t *testing.T) {
		config := fmt.Sprintf(`
			data "github_release_asset" "test" {
				repository = "%s"
				owner = "%s"
				asset_id = "%s"
				download_file_contents = true
			}

			output "github_release_asset_contents" {
				value = base64decode(data.github_release_asset.test.file_contents)
			}
		`, testReleaseRepository, testRepositoryOwner, testReleaseAssetID)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr(
							"data.github_release_asset.test", "asset_id", testReleaseAssetID,
						),
						resource.TestCheckResourceAttr(
							"data.github_release_asset.test", "name", testReleaseAssetName,
						),
						resource.TestCheckOutput("github_release_asset_contents", testReleaseAssetContent),
					),
				},
			},
		})
	})

	t.Run("queries without downloading the specified asset ID", func(t *testing.T) {
		config := fmt.Sprintf(`
			data "github_release_asset" "test" {
				repository = "%s"
				owner = "%s"
				asset_id = "%s"
			}
		`, testReleaseRepository, testRepositoryOwner, testReleaseAssetID)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr(
							"data.github_release_asset.test", "asset_id", testReleaseAssetID,
						),
						resource.TestCheckResourceAttr(
							"data.github_release_asset.test", "name", testReleaseAssetName,
						),
						resource.TestCheckNoResourceAttr(
							"data.github_release_asset.test", "file",
						),
					),
				},
			},
		})
	})
}
