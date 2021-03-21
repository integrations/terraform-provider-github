package github

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccGithubRepositoryEnvironments(t *testing.T) {

	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
	startWaitTimer := fmt.Sprintf("%d", acctest.RandIntRange(0, 43200))
	updatedWaitTimer := fmt.Sprintf("%d", acctest.RandIntRange(0, 43200))

	t.Run("creates and updates repository environment without error", func(t *testing.T) {
		config := fmt.Sprintf(`
			resource "github_repository" "test" {
			  name = "tf-acc-test-%s"
			}

			resource "github_repository_environment" "test" {
			  repository       		= github_repository.test.name
			  name      			= "tf-acc-test-%s"
              wait_timer			= %s
			}
		`, randomID, randomID, startWaitTimer)

		checks := map[string]resource.TestCheckFunc{
			"before": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_repository_environment.test", "name", fmt.Sprintf("tf-acc-test-%s", randomID),
				),
				resource.TestCheckResourceAttr(
					"github_repository_environment.test", "wait_timer", startWaitTimer,
				),
			),
			"after": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_repository_environment.test", "name", fmt.Sprintf("tf-acc-test-%s", randomID),
				),
				resource.TestCheckResourceAttr(
					"github_repository_environment.test", "wait_timer", updatedWaitTimer,
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
							startWaitTimer,
							updatedWaitTimer, 1),
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
	t.Run("deletes repository environment without error", func(t *testing.T) {

		config := fmt.Sprintf(`
				resource "github_repository" "test" {
			  name = "tf-acc-test-%s"
			}

			resource "github_repository_environment" "test" {
			  repository       		= github_repository.test.name
			  name      			= "tf-acc-test-%s"
              wait_timer			= %s
			}
			`, randomID, randomID, startWaitTimer)

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
}
