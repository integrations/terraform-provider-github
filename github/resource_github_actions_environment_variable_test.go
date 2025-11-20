package github

import (
	"context"
	"fmt"
	"net/url"
	"strings"
	"testing"

	"github.com/google/go-github/v67/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccGithubActionsEnvironmentVariable(t *testing.T) {
	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

	t.Run("creates and updates environment variables without error", func(t *testing.T) {
		value := "my_variable_value"
		updatedValue := "my_updated_variable_value"

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
			  name = "tf-acc-test-%s"
			}

			resource "github_repository_environment" "test" {
			  repository       = github_repository.test.name
			  environment      = "environment / test"
			}

			resource "github_actions_environment_variable" "variable" {
			  repository       = github_repository.test.name
			  environment      = github_repository_environment.test.environment
			  variable_name    = "test_variable"
			  value  = "%s"
			}
			`, randomID, value)

		checks := map[string]resource.TestCheckFunc{
			"before": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_actions_environment_variable.variable", "value",
					value,
				),
				resource.TestCheckResourceAttrSet(
					"github_actions_environment_variable.variable", "created_at",
				),
				resource.TestCheckResourceAttrSet(
					"github_actions_environment_variable.variable", "updated_at",
				),
			),
			"after": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_actions_environment_variable.variable", "value",
					updatedValue,
				),
				resource.TestCheckResourceAttrSet(
					"github_actions_environment_variable.variable", "created_at",
				),
				resource.TestCheckResourceAttrSet(
					"github_actions_environment_variable.variable", "updated_at",
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
						Config: strings.Replace(config,
							value,
							updatedValue, 1),
						Check: checks["after"],
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

	t.Run("deletes environment variables without error", func(t *testing.T) {
		config := fmt.Sprintf(`
				resource "github_repository" "test" {
					name = "tf-acc-test-%s"
				}

				resource "github_repository_environment" "test" {
					repository       = github_repository.test.name
					environment      = "environment / test"
				}

				resource "github_actions_environment_variable" "variable" {
					repository 		= github_repository.test.name
					environment     = github_repository_environment.test.environment
					variable_name	= "test_variable"
					value 			= "my_variable_value"
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

	t.Run("imports environment variables without error", func(t *testing.T) {
		value := "my_variable_value"
		envName := "environment / test"
		varName := "test_variable"

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
			  name = "tf-acc-test-%s"
			}

			resource "github_repository_environment" "test" {
			  repository       = github_repository.test.name
			  environment      = "%s"
			}

			resource "github_actions_environment_variable" "variable" {
			  repository       = github_repository.test.name
			  environment      = github_repository_environment.test.environment
			  variable_name    = "%s"
			  value  = "%s"
			}
			`, randomID, envName, varName, value)

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: config,
					},
					{
						ResourceName:      "github_actions_environment_variable.variable",
						ImportStateId:     fmt.Sprintf(`tf-acc-test-%s:%s:%s`, randomID, envName, varName),
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
}

func TestAccGithubActionsEnvironmentVariable_alreadyExists(t *testing.T) {
	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
	repoName := fmt.Sprintf("tf-acc-test-%s", randomID)
	envName := "environment / test"
	varName := "test_variable"
	value := "my_variable_value"

	config := fmt.Sprintf(`
		resource "github_repository" "test" {
			name = "%s"
			vulnerability_alerts = true
		}

		resource "github_repository_environment" "test" {
			repository       = github_repository.test.name
			environment      = "%s"
		}

		resource "github_actions_environment_variable" "variable" {
			repository       = github_repository.test.name
			environment      = github_repository_environment.test.environment
			variable_name    = "%s"
			value  = "%s"
		}
	`, repoName, envName, varName, value)

	testCase := func(t *testing.T, mode string) {
		resource.Test(t, resource.TestCase{
			PreCheck: func() {
				skipUnlessMode(t, mode)
			},
			Providers: testAccProviders,
			Steps: []resource.TestStep{
				{
					// First, create the repository and environment.
					Config: fmt.Sprintf(`
						resource "github_repository" "test" {
							name = "%s"
							vulnerability_alerts = true
						}

						resource "github_repository_environment" "test" {
							repository       = github_repository.test.name
							environment      = "%s"
						}
					`, repoName, envName),
					Check: resource.ComposeTestCheckFunc(
						func(s *terraform.State) error {
							// Now that the repo and env are created, create the variable using the API.
							client := testAccProvider.Meta().(*Owner).v3client
							owner := testAccProvider.Meta().(*Owner).name
							ctx := context.Background()
							escapedEnvName := url.PathEscape(envName)

							variable := &github.ActionsVariable{
								Name:  varName,
								Value: value,
							}
							_, err := client.Actions.CreateEnvVariable(ctx, owner, repoName, escapedEnvName, variable)
							return err
						},
					),
				},
				{
					// Now, run the full config. Terraform should detect the existing variable and "adopt" it.
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_actions_environment_variable.variable", "value", value),
					),
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
}
