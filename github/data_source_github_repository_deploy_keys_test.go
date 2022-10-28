package github

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccGithubRepositoryDeployKeysDataSource(t *testing.T) {
	t.Run("manages deploy keys", func(t *testing.T) {
		keyPath := filepath.Join("test-fixtures", "id_rsa.pub")

		repoName := fmt.Sprintf("tf-acc-test-deploy-keys-%s", acctest.RandString(5))
		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name      = "%s"
				auto_init = true
			}

			resource "github_repository_deploy_key" "test" {
				repository = github_repository.test.name
				key        = "${file("%s")}"
				title      = "title1"
			}
		 `, repoName, keyPath)

		config2 := config + `
			data "github_repository_deploy_keys" "test" {
				repository = github_repository.test.name
			}
		`

		const resourceName = "data.github_repository_deploy_keys.test"
		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceName, "keys.#", "1"),
			resource.TestCheckResourceAttr(resourceName, "keys.0.title", "title1"),
			resource.TestCheckResourceAttrSet(resourceName, "keys.0.id"),
			resource.TestCheckResourceAttr(resourceName, "keys.0.verified", "true"),
		)

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: config,
						Check:  resource.ComposeTestCheckFunc(),
					},
					{
						Config: config2,
						Check:  check,
					},
				},
			})
		}

		t.Run("with an anonymous account", func(t *testing.T) {
			t.Skip("anonymous account not supported for this operation")
		})

		t.Run("with an individual account", func(t *testing.T) {
			testCase(t, individual)
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})
	})
}
