package github

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubTeamDataSource(t *testing.T) {
	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

	t.Run("queries an existing team without error", func(t *testing.T) {
		config := fmt.Sprintf(`
			resource "github_team" "test" {
				name = "tf-acc-test-%s"
			}

			data "github_team" "test" {
				slug = github_team.test.slug
			}
		`, randomID)

		check := resource.ComposeAggregateTestCheckFunc(
			resource.TestCheckResourceAttrSet("data.github_team.test", "name"),
			resource.TestCheckResourceAttrSet("data.github_team.test", "node_id"),
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

	t.Run("queries an existing team without error with immediate membership", func(t *testing.T) {
		config := fmt.Sprintf(`
			resource "github_team" "test" {
				name = "tf-acc-test-%s"
			}

			data "github_team" "test" {
				slug            = github_team.test.slug
				membership_type = "immediate"
			}
		`, randomID)

		check := resource.ComposeAggregateTestCheckFunc(
			resource.TestCheckResourceAttrSet("data.github_team.test", "name"),
			resource.TestCheckResourceAttr("data.github_team.test", "name", fmt.Sprintf("tf-acc-test-%s", randomID)),
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

	t.Run("errors when querying a non-existing team", func(t *testing.T) {
		config := `
			data "github_team" "test" {
				slug = ""
			}
		`

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config:      config,
						ExpectError: regexp.MustCompile(`Not Found`),
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

	t.Run("queries an existing team without error in summary_only mode", func(t *testing.T) {
		config := fmt.Sprintf(`
			resource "github_team" "test" {
				name = "tf-acc-test-%s"
			}

			data "github_team" "test" {
				slug = github_team.test.slug
				summary_only = true
			}
		`, randomID)

		check := resource.ComposeAggregateTestCheckFunc(
			resource.TestCheckResourceAttrSet("data.github_team.test", "name"),
			resource.TestCheckResourceAttrSet("data.github_team.test", "node_id"),
			resource.TestCheckResourceAttr("data.github_team.test", "members.#", "0"),
			resource.TestCheckResourceAttr("data.github_team.test", "repositories.#", "0"),
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

	t.Run("queries an existing team without error with results_per_page reduced", func(t *testing.T) {
		config := fmt.Sprintf(`
			resource "github_team" "test" {
				name = "tf-acc-test-%s"
			}

			data "github_team" "test" {
				slug = github_team.test.slug
				results_per_page = 20
			}
		`, randomID)

		check := resource.ComposeAggregateTestCheckFunc(
			resource.TestCheckResourceAttrSet("data.github_team.test", "name"),
			resource.TestCheckResourceAttrSet("data.github_team.test", "node_id"),
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

	t.Run("queries an existing team with connected repositories", func(t *testing.T) {
		config := fmt.Sprintf(`
			resource "github_team" "test" {
				name = "tf-acc-test-%s"
			}
			resource "github_repository" "test" {
				name = "tf-acc-test"
			}
			resource "github_team_repository" "test" {
				team_id    = github_team.test.id
				repository = github_repository.test.name
				permission = "admin"
			}
		`, randomID)

		config2 := config + `
			data "github_team" "test" {
				slug = github_team.test.slug
			}
		`

		check := resource.ComposeAggregateTestCheckFunc(
			resource.TestCheckResourceAttrSet("data.github_team.test", "name"),
			resource.TestCheckResourceAttr("github_repository.test", "name", "tf-acc-test"),
			resource.TestCheckResourceAttr("data.github_team.test", "repositories_detailed.#", "1"),
			resource.TestCheckResourceAttrPair("data.github_team.test", "repositories_detailed.0.repo_id", "github_repository.test", "repo_id"),
			resource.TestCheckResourceAttrPair("data.github_team.test", "repositories_detailed.0.role_name", "github_team_repository.test", "permission"),
		)

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: config,
						Check:  resource.ComposeAggregateTestCheckFunc(),
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
			t.Skip("individual account not supported for this operation")
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})
	})
}
