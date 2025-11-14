package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubOrganizationProjectsDataSource(t *testing.T) {
	t.Run("queries organization projects", func(t *testing.T) {
		config := fmt.Sprintf(`
			data "github_organization_projects" "test" {
				organization = "%s"
			}
		`, testOrganization)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet("data.github_organization_projects.test", "projects.#"),
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
			testCase(t, anonymous)
		})

		t.Run("with an individual account", func(t *testing.T) {
			testCase(t, individual)
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})
	})

	t.Run("validates projects attributes", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

		config := fmt.Sprintf(`
			data "github_organization_projects" "test" {
				organization = "%s"
			}
		`, testOrganization)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet("data.github_organization_projects.test", "projects.#"),
			resource.TestCheckResourceAttrSet("data.github_organization_projects.test", "projects.0.id"),
			resource.TestCheckResourceAttrSet("data.github_organization_projects.test", "projects.0.node_id"),
			resource.TestCheckResourceAttrSet("data.github_organization_projects.test", "projects.0.number"),
			resource.TestCheckResourceAttrSet("data.github_organization_projects.test", "projects.0.title"),
			resource.TestCheckResourceAttrSet("data.github_organization_projects.test", "projects.0.url"),
			resource.TestCheckResourceAttrSet("data.github_organization_projects.test", "projects.0.created_at"),
			resource.TestCheckResourceAttrSet("data.github_organization_projects.test", "projects.0.updated_at"),
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

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})

		_ = randomID // Prevent unused variable error
	})
}
