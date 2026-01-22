package github

import (
	"encoding/base64"
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
	base64EncodedAssetContent := base64.StdEncoding.EncodeToString([]byte(testReleaseAssetContent))

	t.Run("queries and downloads specified asset ID", func(t *testing.T) {
		config := fmt.Sprintf(`
			data "github_release_asset" "test" {
				repository = "%s"
				owner = "%s"
				asset_id = "%s"
				download_file_contents = true
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
						resource.TestCheckResourceAttr(
							"data.github_release_asset.test", "file_contents", base64EncodedAssetContent,
						),
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
