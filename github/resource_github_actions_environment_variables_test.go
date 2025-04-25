package github

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubActionsEnvironmentVariables(t *testing.T) {

	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

	t.Run("creates and updates multiple environment variables without error", func(t *testing.T) {
		config := fmt.Sprintf(`
			resource "github_repository" "test" {
			  name = "tf-acc-test-%s"
			}

			resource "github_repository_environment" "test" {
			  repository       = github_repository.test.name
			  environment      = "environment / test"
			}

			resource "github_actions_environment_variables" "test_vars" {
			  repository       = github_repository.test.name
			  environment      = github_repository_environment.test.environment

			  variable {
			    name  = "test_variable_1"
			    value = "value_1"
			  }

			  variable {
			    name  = "test_variable_2"
			    value = "value_2"
			  }
			}
			`, randomID)

		updatedConfig := fmt.Sprintf(`
			resource "github_repository" "test" {
			  name = "tf-acc-test-%s"
			}

			resource "github_repository_environment" "test" {
			  repository       = github_repository.test.name
			  environment      = "environment / test"
			}

			resource "github_actions_environment_variables" "test_vars" {
			  repository       = github_repository.test.name
			  environment      = github_repository_environment.test.environment

			  variable {
			    name  = "test_variable_1"
			    value = "updated_value_1"
			  }

			  variable {
			    name  = "test_variable_3"
			    value = "value_3"
			  }
			}
			`, randomID)

		checks := map[string]resource.TestCheckFunc{
			"before": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_actions_environment_variables.test_vars", "variable.#", "2",
				),
				resource.TestCheckResourceAttrSet(
					"github_actions_environment_variables.test_vars", "variable.0.created_at",
				),
				resource.TestCheckResourceAttrSet(
					"github_actions_environment_variables.test_vars", "variable.0.updated_at",
				),
			),
			"after": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_actions_environment_variables.test_vars", "variable.#", "2",
				),
				resource.TestCheckResourceAttrSet(
					"github_actions_environment_variables.test_vars", "variable.0.created_at",
				),
				resource.TestCheckResourceAttrSet(
					"github_actions_environment_variables.test_vars", "variable.0.updated_at",
				),
			),
		}

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
						Config: updatedConfig,
						Check:  checks["after"],
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

	t.Run("deletes all environment variables without error", func(t *testing.T) {
		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name = "tf-acc-test-%s"
			}

			resource "github_repository_environment" "test" {
				repository       = github_repository.test.name
				environment      = "environment / test"
			}

			resource "github_actions_environment_variables" "test_vars" {
				repository 		= github_repository.test.name
				environment     = github_repository_environment.test.environment

				variable {
					name  = "test_variable_1"
					value = "value_1"
				}

				variable {
					name  = "test_variable_2"
					value = "value_2"
				}
			}
		`, randomID)

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config:  config,
						Destroy: true,
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

	t.Run("imports environment variables collection without error", func(t *testing.T) {
		envName := "environment / test"

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
			  name = "tf-acc-test-%s"
			}

			resource "github_repository_environment" "test" {
			  repository       = github_repository.test.name
			  environment      = "%s"
			}

			resource "github_actions_environment_variables" "test_vars" {
			  repository       = github_repository.test.name
			  environment      = github_repository_environment.test.environment

			  variable {
			    name  = "test_variable_1"
			    value = "value_1"
			  }

			  variable {
			    name  = "test_variable_2"
			    value = "value_2"
			  }
			}
		`, randomID, envName)

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: config,
					},
					{
						ResourceName:      "github_actions_environment_variables.test_vars",
						ImportStateId:     fmt.Sprintf(`tf-acc-test-%s:%s`, randomID, envName),
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
			testCase(t, individual)
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})
	})

	t.Run("enforces 100 variable limit", func(t *testing.T) {
		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name      = "tf-acc-test-%s"
				auto_init = true
			}

			resource "github_repository_environment" "test" {
				repository       = github_repository.test.name
				environment      = "environment / test"
			}

			resource "github_actions_environment_variables" "test_vars" {
				repository       = github_repository.test.name
				environment      = github_repository_environment.test.environment

				dynamic "variable" {
					for_each = range(101)
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
						ExpectError: regexp.MustCompile(`environment variable set cannot contain more than 100 items`),
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
