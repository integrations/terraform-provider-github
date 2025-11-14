package github

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubProjectFieldsDataSource(t *testing.T) {
	t.Run("queries organization project fields", func(t *testing.T) {
		config := fmt.Sprintf(`
			data "github_project_fields" "test" {
				project_number = 1
				organization   = "%s"
			}
		`, testOrganization)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet("data.github_project_fields.test", "fields.#"),
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

	t.Run("queries user project fields", func(t *testing.T) {
		config := fmt.Sprintf(`
			data "github_project_fields" "test" {
				project_number = 1
				username       = "%s"
			}
		`, testOwner)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet("data.github_project_fields.test", "fields.#"),
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

		t.Run("with an individual account", func(t *testing.T) {
			testCase(t, individual)
		})
	})

	t.Run("validates project fields attributes", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

		config := fmt.Sprintf(`
			data "github_project_fields" "test" {
				project_number = 1
				organization   = "%s"
			}
		`, testOrganization)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet("data.github_project_fields.test", "fields.#"),
			resource.TestCheckResourceAttrSet("data.github_project_fields.test", "fields.0.id"),
			resource.TestCheckResourceAttrSet("data.github_project_fields.test", "fields.0.node_id"),
			resource.TestCheckResourceAttrSet("data.github_project_fields.test", "fields.0.name"),
			resource.TestCheckResourceAttrSet("data.github_project_fields.test", "fields.0.data_type"),
			resource.TestCheckResourceAttrSet("data.github_project_fields.test", "fields.0.created_at"),
			resource.TestCheckResourceAttrSet("data.github_project_fields.test", "fields.0.updated_at"),
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

	t.Run("validates conflicting arguments", func(t *testing.T) {
		config := fmt.Sprintf(`
			data "github_project_fields" "test" {
				project_number = 1
				organization   = "%s"
				username       = "%s"
			}
		`, testOrganization, testOwner)

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config:      config,
						ExpectError: regexp.MustCompile("\"organization\": conflicts with username"),
					},
				},
			})
		}

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})
	})
}
