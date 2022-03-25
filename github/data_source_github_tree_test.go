package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccGithubTreeDataSource(t *testing.T) {
	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

	t.Run("get tree", func(t *testing.T) {
		config := fmt.Sprintf(`
			resource "github_repository" "this" {
				auto_init = true
				name      = "tf-acc-test-%s"
			}

			data "github_branch" "this" {
				branch     = "main"
				repository = github_repository.this.name
			}

			data "github_tree" "this" {
				recursive  = false
				repository = github_repository.this.name
				tree_sha   = data.github_branch.this.sha
			}
		`, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet(
				"data.github_tree.this", "entries.#",
			),
			resource.TestCheckResourceAttr(
				"data.github_tree.this", "entries.0.path",
				"README.md",
			),
			resource.TestCheckResourceAttr(
				"data.github_tree.this", "entries.0.type",
				"blob",
			),
		)

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: config,
						Check:  check,
					},
				},
			})
		}

		t.Run("with an anonymous account", func(t *testing.T) {
			testCase(t, anonymous)
		})

		t.Run("with an individual account", func(t *testing.T) {
			testCase(t, individual)
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})
	})
}
