package github

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"testing"
)

func TestAccGithubActionsRepositoryAccessLevel(t *testing.T) {
	t.Run("test setting of user action access level", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		accessLevel := "user"
		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name        = "tf-acc-test-topic-%[1]s"
				description = "Terraform acceptance tests %[1]s"
				topics		= ["terraform", "testing"]
				visibility  = "private"
			}

			resource "github_actions_repository_access_level" "test" {
				access_level = "%s"
				repository   = github_repository.test.name
			}
		`, randomID, accessLevel)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_actions_repository_access_level.test", "access_level", accessLevel,
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
			t.Skip("organization account not supported for this input")
		})
	})

	t.Run("test setting of organization action access level", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		accessLevel := "organization"
		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name        = "tf-acc-test-topic-%[1]s"
				description = "Terraform acceptance tests %[1]s"
				topics		= ["terraform", "testing"]
				visibility  = "private"
			}

			resource "github_actions_repository_access_level" "test" {
				access_level = "%s"
				repository   = github_repository.test.name
			}
		`, randomID, accessLevel)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_actions_repository_access_level.test", "access_level", accessLevel,
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
			t.Skip("individual account not supported for this input")
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})
	})
}
