package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccGithubOrganizationDataSource(t *testing.T) {

	t.Run("queries for an organization without error", func(t *testing.T) {

		config := fmt.Sprintf(`
			data "github_organization" "test" {
				name = "%s"
			}
		`, testOrganization)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("data.github_organization.test", "login", testOrganization),
			resource.TestCheckResourceAttrSet("data.github_organization.test", "name"),
			resource.TestCheckResourceAttrSet("data.github_organization.test", "orgname"),
			resource.TestCheckResourceAttrSet("data.github_organization.test", "node_id"),
			resource.TestCheckResourceAttrSet("data.github_organization.test", "description"),
			resource.TestCheckResourceAttrSet("data.github_organization.test", "plan"),
			resource.TestCheckResourceAttrSet("data.github_organization.test", "repositories.#"),
			resource.TestCheckResourceAttrSet("data.github_organization.test", "members.#"),
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

	t.Run("queries for an organization with archived repos", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

		config := fmt.Sprintf(`
			resource "github_repository" "archived" {
				name         = "tf-acc-archived-%s"
				archived     = true
		  	}

			data "github_organization" "skip_archived" {
				name = "%s"
				ignore_archived_repos = true
				depends_on = [
					github_repository.archived,
				]
			}
			data "github_organization" "all_repos" {
				name = "%s"
				ignore_archived_repos = false
				depends_on = [
					github_repository.archived,
				]
			}
			
			output "should_be_false" {
				value = contains(data.github_organization.skip_archived.repositories, github_repository.archived.full_name)
			}
			output "should_be_true" {
				value = contains(data.github_organization.all_repos.repositories, github_repository.archived.full_name)
			}
		`, randomID, testOrganization, testOrganization)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckOutput("should_be_false", "false"),
			resource.TestCheckOutput("should_be_true", "true"),
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

}
