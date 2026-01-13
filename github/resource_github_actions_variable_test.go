package github

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubActionsVariable(t *testing.T) {
	t.Run("creates and updates repository variables without error", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-act-var-%s", testResourcePrefix, randomID)
		value := "my_variable_value"
		updatedValue := "my_updated_variable_value"

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
			  name = "%s"
			}

			resource "github_actions_variable" "variable" {
			  repository       = github_repository.test.name
			  variable_name    = "test_variable"
			  value  = "%s"
			}
			`, repoName, value)

		checks := map[string]resource.TestCheckFunc{
			"before": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_actions_variable.variable", "value",
					value,
				),
				resource.TestCheckResourceAttrSet(
					"github_actions_variable.variable", "created_at",
				),
				resource.TestCheckResourceAttrSet(
					"github_actions_variable.variable", "updated_at",
				),
			),
			"after": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_actions_variable.variable", "value",
					updatedValue,
				),
				resource.TestCheckResourceAttrSet(
					"github_actions_variable.variable", "created_at",
				),
				resource.TestCheckResourceAttrSet(
					"github_actions_variable.variable", "updated_at",
				),
			),
		}

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
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
	})

	t.Run("deletes repository variables without error", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-act-var-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`
				resource "github_repository" "test" {
					name = "%s"
				}

				resource "github_actions_variable" "variable" {
					repository 		= github_repository.test.name
					variable_name	= "test_variable"
					value			= "my_variable_value"
				}
			`, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:  config,
					Destroy: true,
				},
			},
		})
	})

	t.Run("imports repository variables without error", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-act-var-%s", testResourcePrefix, randomID)
		varName := "test_variable"
		value := "variable_value"

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
			  name = "%s"
			}

			resource "github_actions_variable" "variable" {
			  repository       = github_repository.test.name
			  variable_name    = "%s"
			  value  = "%s"
			}
			`, repoName, varName, value)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
				},
				{
					ResourceName:      "github_actions_variable.variable",
					ImportStateId:     fmt.Sprintf(`%s:%s`, repoName, varName),
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})
}
