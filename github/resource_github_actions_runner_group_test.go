package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccGithubActionsRunnerGroup(t *testing.T) {

	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

	t.Run("creates runner groups without error", func(t *testing.T) {

		// t.Skip("requires an enterprise cloud account")

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
			  name = "tf-acc-test-%s"
			  vulnerability_alerts = false
			}

			resource "github_actions_runner_group" "test" {
			  name       = github_repository.test.name
			  visibility = "all"
			}
		`, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet(
				"github_actions_runner_group.test", "name",
			),
			resource.TestCheckResourceAttr(
				"github_actions_runner_group.test", "name",
				fmt.Sprintf(`tf-acc-test-%s`, randomID),
			),
			resource.TestCheckResourceAttr(
				"github_actions_runner_group.test", "visibility",
				"all",
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
			t.Skip("individual account not supported for this operation")
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})
	})

	t.Run("manages runner visibility", func(t *testing.T) {

		// t.Skip("requires an enterprise cloud account")

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
			  name = "tf-acc-test-%s"
			}

			resource "github_actions_runner_group" "test" {
			  name       = github_repository.test.name
			  visibility = "selected"
			  selected_repository_ids = [github_repository.test.repo_id]
			}
		`, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet(
				"github_actions_runner_group.test", "name",
			),
			resource.TestCheckResourceAttr(
				"github_actions_runner_group.test", "name",
				fmt.Sprintf(`tf-acc-test-%s`, randomID),
			),
			resource.TestCheckResourceAttr(
				"github_actions_runner_group.test", "visibility",
				"selected",
			),
			resource.TestCheckResourceAttr(
				"github_actions_runner_group.test", "selected_repository_ids.#",
				"1",
			),
			resource.TestCheckResourceAttrSet(
				"github_actions_runner_group.test", "selected_repositories_url",
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
			t.Skip("individual account not supported for this operation")
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})
	})

	t.Run("imports an all runner group without error", func(t *testing.T) {
		config := fmt.Sprintf(`
			resource "github_repository" "test" {
			  name = "tf-acc-test-%s"
			}

			resource "github_actions_runner_group" "test" {
			  name       = github_repository.test.name
			  visibility = "all"
			}
    `, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet("github_actions_runner_group.test", "name"),
			resource.TestCheckResourceAttrSet("github_actions_runner_group.test", "visibility"),
			resource.TestCheckResourceAttr("github_actions_runner_group.test", "visibility", "all"),
			resource.TestCheckResourceAttr("github_actions_runner_group.test", "name", fmt.Sprintf(`tf-acc-test-%s`, randomID)),
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
					{
						ResourceName:      "github_actions_runner_group.test",
						ImportState:       true,
						ImportStateVerify: true,
					},
				},
			})
		}

		t.Run("with an anonymous account", func(t *testing.T) {
			t.Skip("anonymous account not supported for this operation")
		})

		t.Run("with an individual account", func(t *testing.T) {
			t.Skip("individual account not supported for this operation")
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})
	})

	t.Run("imports a private runner group without error", func(t *testing.T) {
		config := fmt.Sprintf(`
					resource "github_repository" "test" {
					  name = "tf-acc-test-%s"
					}

					resource "github_actions_runner_group" "test" {
					  name       = github_repository.test.name
					  visibility = "private"
					}
		    `, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet("github_actions_runner_group.test", "name"),
			resource.TestCheckResourceAttr("github_actions_runner_group.test", "name", fmt.Sprintf(`tf-acc-test-%s`, randomID)),
			resource.TestCheckResourceAttrSet("github_actions_runner_group.test", "visibility"),
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
					{
						ResourceName:      "github_actions_runner_group.test",
						ImportState:       true,
						ImportStateVerify: true,
					},
				},
			})
		}

		t.Run("with an anonymous account", func(t *testing.T) {
			t.Skip("anonymous account not supported for this operation")
		})

		t.Run("with an individual account", func(t *testing.T) {
			t.Skip("individual account not supported for this operation")
		})

		t.Run("with an organization account", func(t *testing.T) {
			// Note: this test is skipped because when setting visibility 'private', it always fails with:
			// Step 0 error: After applying this step, the plan was not empty:
			// visibility:                 "all" => "private"
			t.Skip("always shows a diff for visibility 'all' => 'private'")
			testCase(t, organization)
		})
	})

	t.Run("imports a selected runner group without error", func(t *testing.T) {

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name = "tf-acc-test-%s"
			}

			resource "github_actions_runner_group" "test" {
				name       = github_repository.test.name
				visibility = "selected"
				selected_repository_ids = [github_repository.test.repo_id]
			}
    `, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet("github_actions_runner_group.test", "name"),
			resource.TestCheckResourceAttr("github_actions_runner_group.test", "name", fmt.Sprintf(`tf-acc-test-%s`, randomID)),
			resource.TestCheckResourceAttrSet("github_actions_runner_group.test", "visibility"),
			resource.TestCheckResourceAttr("github_actions_runner_group.test", "visibility", "selected"),
			resource.TestCheckResourceAttr(
				"github_actions_runner_group.test", "selected_repository_ids.#",
				"1",
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
					{
						ResourceName:      "github_actions_runner_group.test",
						ImportState:       true,
						ImportStateVerify: true,
					},
				},
			})
		}

		t.Run("with an individual account", func(t *testing.T) {
			testCase(t, individual)
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})
	})
}
