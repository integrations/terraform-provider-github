package github

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubRepositoryEnvironmentDeploymentCustomProtectionRules(t *testing.T) {
	const APP_INTEGRATION_ID = "APP_INTEGRATION_ID"
	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
	integration_id, exists := os.LookupEnv(APP_INTEGRATION_ID)

	t.Run("queries custom deployment protection rules", func(t *testing.T) {
		if !exists {
			t.Skipf("%s environment variable is missing", APP_INTEGRATION_ID)
		}
		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name      			 = "tf-acc-test-%s"
				vulnerability_alerts = "true"
			}
			resource "github_repository_environment" "test" {
				repository 	= github_repository.test.name
				environment	= "environment / test"
			}
		`, randomID)

		config2 := config + `
			data "github_repository_environment_custom_deployment_protection_rules" "test" {
				repository 	= github_repository.test.name
				environment = github_repository_environment.test.environment
			}
		`
		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"data.github_repository_environment_custom_deployment_protection_rules.test", "repository",
				fmt.Sprintf("tf-acc-test-%s", randomID),
			),
			resource.TestCheckResourceAttr(
				"data.github_repository_environment_custom_deployment_protection_rules.test", "environment",
				"environment / test",
			),
			resource.TestCheckResourceAttr(
				"data.github_repository_environment_custom_deployment_protection_rules.test", "custom_deployment_protection_rules.0.id",
				integration_id,
			),
		)

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: config,
					},
					{
						Config: config2,
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
