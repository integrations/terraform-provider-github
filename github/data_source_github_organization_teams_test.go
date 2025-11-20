package github

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubOrganizationTeamsDataSource(t *testing.T) {
	t.Run("queries without error", func(t *testing.T) {
		config := `
			data "github_organization_teams" "all" {}
		`

		check := resource.ComposeAggregateTestCheckFunc(
			resource.TestCheckResourceAttrSet("data.github_organization_teams.all", "teams.0.id"),
			resource.TestCheckResourceAttrSet("data.github_organization_teams.all", "teams.0.node_id"),
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

	t.Run("queries root teams only without error", func(t *testing.T) {
		config := `
			data "github_organization_teams" "root_teams" {
				root_teams_only = true
			}
		`

		check := resource.ComposeAggregateTestCheckFunc(
			resource.TestCheckResourceAttrSet("data.github_organization_teams.root_teams", "teams.0.id"),
			resource.TestCheckResourceAttrSet("data.github_organization_teams.root_teams", "teams.0.node_id"),
			resource.TestCheckResourceAttr("data.github_organization_teams.root_teams", "teams.0.parent.id", ""),
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

	t.Run("queries summary only without error", func(t *testing.T) {
		config := `
			data "github_organization_teams" "all" {
				summary_only = true
			}
		`

		check := resource.ComposeAggregateTestCheckFunc(
			resource.TestCheckResourceAttrSet("data.github_organization_teams.all", "teams.0.id"),
			resource.TestCheckResourceAttrSet("data.github_organization_teams.all", "teams.0.node_id"),
			resource.TestCheckNoResourceAttr("data.github_organization_teams.all", "teams.0.members.0"),
			resource.TestCheckNoResourceAttr("data.github_organization_teams.all", "teams.0.repositories.0"),
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

	t.Run("queries results_per_page only without error", func(t *testing.T) {
		config := `
			data "github_organization_teams" "all" {
				results_per_page = 50
			}
		`

		check := resource.ComposeAggregateTestCheckFunc(
			resource.TestCheckResourceAttrSet("data.github_organization_teams.all", "teams.0.id"),
			resource.TestCheckResourceAttrSet("data.github_organization_teams.all", "teams.0.node_id"),
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
}
