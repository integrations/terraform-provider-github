package github

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccGithubActionsRepositoryVariable(t *testing.T) {
	t.Run("creates_updates_and_deletes", func(t *testing.T) {
		repoName := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		variableName := "a" + acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: fmt.Sprintf(`
							resource "github_repository" "test" {
								name = "tf-acc-test-%s"
							}

							resource "github_actions_repository_variable" "test" {
								repository = github_repository.test.name
								name       = "%s"
								value      = "testing_value"
							}
						`, repoName, variableName),
						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckResourceAttr(
								"github_actions_repository_variable.test", "name",
								strings.ToUpper(variableName),
							),
							resource.TestCheckResourceAttr(
								"github_actions_repository_variable.test", "value",
								"testing_value",
							),
							resource.TestCheckResourceAttrSet(
								"github_actions_repository_variable.test", "created_at",
							),
							resource.TestCheckResourceAttrSet(
								"github_actions_repository_variable.test", "updated_at",
							),
						),
					},

					// Update name
					{
						Config: fmt.Sprintf(`
							resource "github_repository" "test" {
								name = "tf-acc-test-%s"
							}

							resource "github_actions_repository_variable" "test" {
								repository = github_repository.test.name
								name       = "%s_2"
								value      = "testing_value"
							}
						`, repoName, variableName),
						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckResourceAttr(
								"github_actions_repository_variable.test", "name",
								strings.ToUpper(variableName+"_2"),
							),
						),
					},

					// Update value
					{
						Config: fmt.Sprintf(`
							resource "github_repository" "test" {
								name = "tf-acc-test-%s"
							}

							resource "github_actions_repository_variable" "test" {
								repository = github_repository.test.name
								name       = "%s_2"
								value      = "testing_value_2"
							}
						`, repoName, variableName),
						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckResourceAttr(
								"github_actions_repository_variable.test", "value",
								"testing_value_2",
							),
						),
					},
				},
			})
		}

		t.Run("with_an_individual_account", func(t *testing.T) {
			testCase(t, individual)
		})

		t.Run("with_an_organization_account", func(t *testing.T) {
			testCase(t, organization)
		})
	})

	t.Run("imports", func(t *testing.T) {
		repoName := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		variableName := "a" + acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					// Create
					{
						Config: fmt.Sprintf(`
							resource "github_repository" "test" {
								name = "tf-acc-test-%s"
							}

							resource "github_actions_repository_variable" "test" {
								repository = github_repository.test.name
								name       = "%s"
								value      = "testing_value"
							}
						`, repoName, variableName),
					},

					// Import
					{
						ResourceName:      "github_actions_repository_variable.test",
						ImportStateId:     fmt.Sprintf("tf-acc-test-%s:%s", repoName, variableName),
						ImportState:       true,
						ImportStateVerify: true,
						ImportStateVerifyIgnore: []string{
							"created_at",
							"updated_at",
						},
					},
				},
			})
		}

		t.Run("with_an_individual_account", func(t *testing.T) {
			testCase(t, individual)
		})

		t.Run("with_an_organization_account", func(t *testing.T) {
			testCase(t, organization)
		})
	})
}
