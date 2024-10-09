package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubRepositoryCustomProperty(t *testing.T) {

	t.Skip("update <property_name> below and make sure the org you're testing against have a custom property with that name to unskip this test run") // TOOD: Actually run this

	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

	t.Run("creates invitations without error", func(t *testing.T) {

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name = "tf-acc-test-%s"
				auto_init = true
			}
			resource "github_repository_custom_property" "test" {
				repository    = github_repository.test.name
				property_name = "<property_name>"
				property_value = ["test"]
			}
		`, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("github_repository_custom_property.test", "property_name", "<property_name>"),
			resource.TestCheckResourceAttr("github_repository_custom_property.test", "property_value.#", "1"),
			resource.TestCheckResourceAttr("github_repository_custom_property.test", "property_value.0", "test"),
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

	// t.Run("creates invitations when repository contains the org name", func(t *testing.T) {

	// 	orgName := os.Getenv("GITHUB_ORGANIZATION")

	// 	if orgName == "" {
	// 		t.Skip("Set GITHUB_ORGANIZATION to unskip this test run")
	// 	}

	// 	configWithOwner := fmt.Sprintf(`
	// 		resource "github_repository" "test" {
	// 			name = "tf-acc-test-%s"
	// 			auto_init = true
	// 		}

	// 		resource "github_repository_collaborator" "test_repo_collaborator_2" {
	// 			repository = "%s/${github_repository.test.name}"
	// 			username   = "<username>"
	// 			permission = "triage"
	// 		}
	// 	`, randomID, orgName)

	// 	checkWithOwner := resource.ComposeTestCheckFunc(
	// 		resource.TestCheckResourceAttr(
	// 			"github_repository_collaborator.test_repo_collaborator_2", "permission",
	// 			"triage",
	// 		),
	// 	)

	// 	testCase := func(t *testing.T, mode string) {
	// 		resource.Test(t, resource.TestCase{
	// 			PreCheck:  func() { skipUnlessMode(t, mode) },
	// 			Providers: testAccProviders,
	// 			Steps: []resource.TestStep{
	// 				{
	// 					Config: configWithOwner,
	// 					Check:  checkWithOwner,
	// 				},
	// 			},
	// 		})
	// 	}

	// 	t.Run("with an anonymous account", func(t *testing.T) {
	// 		t.Skip("anonymous account not supported for this operation")
	// 	})

	// 	t.Run("with an individual account", func(t *testing.T) {
	// 		testCase(t, individual)
	// 	})

	// 	t.Run("with an organization account", func(t *testing.T) {
	// 		testCase(t, organization)
	// 	})
	// })
}

