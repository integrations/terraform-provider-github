package github

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubReleaseAssetDataSource(t *testing.T) {

	testReleaseRepository := testAccConf.testPublicRepository

	// NOTE: the default repository, owner, asset ID, asset name, and asset content
	// values can be overridden with GH_TEST* environment variables to exercise
	// tests against different release assets in development.
	if os.Getenv("GH_TEST_REPOSITORY") != "" {
		testReleaseRepository = os.Getenv("GH_TEST_REPOSITORY")
	}

	// The terraform-provider-github_6.4.0_manifest.json asset ID from
	// https://github.com/integrations/terraform-provider-github/releases/tag/v6.4.0
	testReleaseAssetID := "207956097"
	if os.Getenv("GH_TEST_REPOSITORY_RELEASE_ASSET_ID") != "" {
		testReleaseAssetID = os.Getenv("GH_TEST_REPOSITORY_RELEASE_ASSET_ID")
	}

	testReleaseAssetName := "terraform-provider-github_6.4.0_manifest.json"
	if os.Getenv("GH_TEST_REPOSITORY_RELEASE_ASSET_NAME") != "" {
		testReleaseAssetName = os.Getenv("GH_TEST_REPOSITORY_RELEASE_ASSET_NAME")
	}

	testReleaseAssetContent := "{\n  \"version\": 1,\n  \"metadata\": {\n    \"protocol_versions\": [\n      \"5.0\"\n    ]\n  }\n}\n"
	if os.Getenv("GH_TEST_REPOSITORY_RELEASE_ASSET_CONTENT") != "" {
		testReleaseAssetContent = os.Getenv("GH_TEST_REPOSITORY_RELEASE_ASSET_CONTENT")
	}

	t.Run("queries specified asset ID", func(t *testing.T) {

		config := fmt.Sprintf(`
			data "github_release_asset" "test" {
				repository = "%s"
				owner = "%s"
				asset_id = "%s"
			}
		`, testReleaseRepository, testAccConf.testPublicRepositoryOwner, testReleaseAssetID)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Providers:         testAccProviders,
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
							"data.github_release_asset.test", "body", testReleaseAssetContent,
						),
					),
				},
			},
		})
	})
}
