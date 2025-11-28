package github

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubActionsVariables(t *testing.T) {

	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

	t.Run("creates and updates repository variables", func(t *testing.T) {
		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name      = "tf-acc-test-%s"
				auto_init = true
			}

			resource "github_actions_variables" "test" {
				repository = github_repository.test.name
				variable {
					name  = "test_variable_1"
					value = "test_value_1"
				}
				variable {
					name  = "test_variable_2"
					value = "test_value_2"
				}
			}
		`, randomID)

		checks := map[string]resource.TestCheckFunc{
			"before": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_actions_variables.test", "variable.#", "2",
				),
				resource.TestMatchResourceAttr(
					"github_actions_variables.test", "variable.0.created_at", regexp.MustCompile(`\d`),
				),
				resource.TestMatchResourceAttr(
					"github_actions_variables.test", "variable.0.updated_at", regexp.MustCompile(`\d`),
				),
			),
		}

		updateConfig := fmt.Sprintf(`
			resource "github_repository" "test" {
				name      = "tf-acc-test-%s"
				auto_init = true
			}

			resource "github_actions_variables" "test" {
				repository = github_repository.test.name
				variable {
					name  = "test_variable_1"
					value = "updated_value_1"
				}
				variable {
					name  = "test_variable_3"
					value = "test_value_3"
				}
			}
		`, randomID)

		checks["after"] = resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_actions_variables.test", "variable.#", "2",
			),
			resource.TestMatchResourceAttr(
				"github_actions_variables.test", "variable.0.updated_at", regexp.MustCompile(`\d`),
			),
		)

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: config,
						Check:  checks["before"],
					},
					{
						Config: updateConfig,
						Check:  checks["after"],
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

	t.Run("deletes all variables", func(t *testing.T) {
		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name      = "tf-acc-test-%s"
				auto_init = true
			}

			resource "github_actions_variables" "test" {
				repository = github_repository.test.name
				variable {
					name  = "test_variable_1"
					value = "test_value_1"
				}
				variable {
					name  = "test_variable_2"
					value = "test_value_2"
				}
			}
		`, randomID)

		emptyConfig := fmt.Sprintf(`
			resource "github_repository" "test" {
				name      = "tf-acc-test-%s"
				auto_init = true
			}

			resource "github_actions_variables" "test" {
				repository = github_repository.test.name
			}
		`, randomID)

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: config,
						Check: resource.TestCheckResourceAttr(
							"github_actions_variables.test", "variable.#", "2",
						),
					},
					{
						Config: emptyConfig,
						Check: resource.TestCheckResourceAttr(
							"github_actions_variables.test", "variable.#", "0",
						),
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

	t.Run("imports variables collection", func(t *testing.T) {
		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name      = "tf-acc-test-%s"
				auto_init = true
			}

			resource "github_actions_variables" "test" {
				repository = github_repository.test.name
				variable {
					name  = "test_variable_1"
					value = "test_value_1"
				}
				variable {
					name  = "test_variable_2"
					value = "test_value_2"
				}
			}
		`, randomID)

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: config,
						Check: resource.TestCheckResourceAttr(
							"github_actions_variables.test", "variable.#", "2",
						),
					},
					{
						ResourceName:      "github_actions_variables.test",
						ImportState:       true,
						ImportStateVerify: true,
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

	t.Run("enforces 500 variable limit", func(t *testing.T) {
		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name      = "tf-acc-test-%s"
				auto_init = true
			}

			resource "github_actions_variables" "test" {
				repository = github_repository.test.name

				dynamic "variable" {
					for_each = range(501)
					content {
						name  = "TEST_VAR_${variable.value + 1}"
						value = "test-value-${variable.value + 1}"
					}
				}
			}
		`, randomID)

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config:      config,
						ExpectError: regexp.MustCompile(`variable set cannot contain more than 500 items`),
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
