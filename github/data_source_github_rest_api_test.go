package github

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccGithubRestApiDataSource(t *testing.T) {

	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

	t.Run("queries an existing branch without error", func(t *testing.T) {

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
			  name = "tf-acc-test-%[1]s"
				auto_init = true
			}

			data "github_rest_api" "test" {
				endpoint = "repos/${github_repository.test.full_name}/git/refs/heads/${github_repository.test.default_branch}"
			}
		`, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestMatchResourceAttr(
				"data.github_rest_api.test", "code", regexp.MustCompile("200"),
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
			t.Skip("anonymous account not supported for this operation")
		})

		t.Run("with an individual account", func(t *testing.T) {
			testCase(t, individual)
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})

	})

	t.Run("queries an invalid branch without error", func(t *testing.T) {

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
			  name = "tf-acc-test-%[1]s"
				auto_init = true
			}

			data "github_rest_api" "test" {
				endpoint = "repos/${github_repository.test.full_name}/git/refs/heads/xxxxxx"
			}
		`, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestMatchResourceAttr(
				"data.github_rest_api.test", "code", regexp.MustCompile("404"),
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
