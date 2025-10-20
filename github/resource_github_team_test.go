package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubTeam(t *testing.T) {

	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

	t.Run("creates a team configured with defaults", func(t *testing.T) {

		config := fmt.Sprintf(`
			resource "github_team" "test" {
				name         = "tf-acc-%s"
			}
		`, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet("github_team.test", "slug"),
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

func TestAccGithubTeamHierarchical(t *testing.T) {

	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

	t.Run("creates a hierarchy of teams", func(t *testing.T) {

		config := fmt.Sprintf(`
			resource "github_team" "team01" {
				name        = "tf-acc-team01-%s"
				description = "Terraform acc test team01a"
				privacy     = "closed"
			}

			resource "github_team" "team02" {
				name           = "tf-acc-team02-%[1]s"
				description    = "Terraform acc test team02a"
				privacy        = "closed"
				parent_team_id = "${github_team.team01.id}"
			}

			resource "github_team" "team03" {
				name           = "tf-acc-team03-%[1]s"
				description    = "Terraform acc test team03a"
				privacy        = "closed"
				parent_team_id = "${github_team.team02.slug}"
			}
		`, randomID)

		config2 := fmt.Sprintf(`
			resource "github_team" "team01" {
				name        = "tf-acc-team01-%s"
				description = "Terraform acc test team01b"
				privacy     = "closed"
			}

			resource "github_team" "team02" {
				name           = "tf-acc-team02-%[1]s"
				description    = "Terraform acc test team02b"
				privacy        = "closed"
			}

			resource "github_team" "team03" {
				name           = "tf-acc-team03-%[1]s"
				description    = "Terraform acc test team03b"
				privacy        = "closed"
			}
		`, randomID)

		check := resource.ComposeAggregateTestCheckFunc(
			resource.TestCheckResourceAttrSet("github_team.team02", "parent_team_id"),
			resource.TestCheckResourceAttrSet("github_team.team03", "parent_team_id"),
		)

		check2 := resource.ComposeAggregateTestCheckFunc(
			resource.TestCheckResourceAttr("github_team.team02", "parent_team_id", ""),
			resource.TestCheckResourceAttr("github_team.team03", "parent_team_id", ""),
			resource.TestCheckResourceAttr("github_team.team02", "parent_team_read_id", ""),
			resource.TestCheckResourceAttr("github_team.team03", "parent_team_read_id", ""),
			resource.TestCheckResourceAttr("github_team.team02", "parent_team_read_slug", ""),
			resource.TestCheckResourceAttr("github_team.team03", "parent_team_read_slug", ""),
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
					{
						Config: config2,
						Check:  check2,
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
func TestAccGithubTeamRemovesDefaultMaintainer(t *testing.T) {

	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

	t.Run("creates a team and removes the default maintainer", func(t *testing.T) {

		config := fmt.Sprintf(`
			resource "github_team" "test" {
				name         = "tf-acc-%s"
				create_default_maintainer = false
			}
		`, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("github_team.test", "members_count", "0"),
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

func TestAccGithubTeamUpdateName(t *testing.T) {
	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

	t.Run("marks the slug as computed when the name changes", func(t *testing.T) {

		config := fmt.Sprintf(`
			resource "github_team" "test" {
				name         = "tf-acc-%s"
			}
		`, randomID)

		configUpdated := fmt.Sprintf(`
			resource "github_team" "test" {
				name         = "tf-acc-updated-%s"
			}

			resource "github_team" "other" {
				name         = "tf-acc-other-%s"
				description  = github_team.test.slug
			}
		`, randomID, randomID)

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: config,
						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckResourceAttr("github_team.test", "slug", fmt.Sprintf("tf-acc-%s", randomID)),
						),
					},
					{
						Config: configUpdated,
						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckResourceAttr("github_team.other", "description", fmt.Sprintf("tf-acc-updated-%s", randomID)),
						),
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
