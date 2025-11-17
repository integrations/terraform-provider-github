package github

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubTeamRepository(t *testing.T) {
	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

	t.Run("manages team permissions to a repository", func(t *testing.T) {
		config := fmt.Sprintf(`
			resource "github_team" "test" {
				name        = "tf-acc-test-team-repo-%s"
				description = "test"
			}

			resource "github_repository" "test" {
				name = "tf-acc-test-%[1]s"
			}

			resource "github_team_repository" "test" {
				team_id    = "${github_team.test.id}"
				repository = "${github_repository.test.name}"
				permission = "pull"
			}
		`, randomID)

		checks := map[string]resource.TestCheckFunc{
			"pull": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_team_repository.test", "permission",
					"pull",
				),
			),
			"triage": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_team_repository.test", "permission",
					"triage",
				),
			),
			"push": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_team_repository.test", "permission",
					"push",
				),
			),
			"maintain": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_team_repository.test", "permission",
					"maintain",
				),
			),
			"admin": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_team_repository.test", "permission",
					"admin",
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
						Check:  checks["pull"],
					},
					{
						Config: strings.Replace(config,
							`permission = "pull"`,
							`permission = "triage"`, 1),
						Check: checks["triage"],
					},
					{
						Config: strings.Replace(config,
							`permission = "pull"`,
							`permission = "push"`, 1),
						Check: checks["push"],
					},
					{
						Config: strings.Replace(config,
							`permission = "pull"`,
							`permission = "maintain"`, 1),
						Check: checks["maintain"],
					},
					{
						Config: strings.Replace(config,
							`permission = "pull"`,
							`permission = "admin"`, 1),
						Check: checks["admin"],
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

	t.Run("accepts both team slug and team ID for `team_id`", func(t *testing.T) {
		config := fmt.Sprintf(`
			resource "github_team" "test" {
				name        = "tf-acc-test-team-repo-%s"
				description = "test"
			}

			resource "github_repository" "test" {
				name = "tf-acc-test-%[1]s"
			}

			resource "github_team_repository" "test" {
				team_id    = "${github_team.test.slug}"
				repository = "${github_repository.test.name}"
				permission = "pull"
			}
		`, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet("github_team_repository.test", "team_id"),
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
						Config: strings.Replace(config,
							`github_team.test.id`,
							`github_team.test.slug`, 1),
						Check: check,
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

func TestAccGithubTeamRepositoryArchivedRepo(t *testing.T) {
	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

	t.Run("can delete team repository access from archived repositories without error", func(t *testing.T) {
		config := fmt.Sprintf(`
			resource "github_team" "test" {
				name        = "tf-acc-test-team-archive-%s"
				description = "test team for archived repo"
			}

			resource "github_repository" "test" {
				name = "tf-acc-test-team-archive-%[1]s"
				auto_init = true
			}

			resource "github_team_repository" "test" {
				team_id    = github_team.test.id
				repository = github_repository.test.name
				permission = "pull"
			}
		`, randomID)

		archivedConfig := fmt.Sprintf(`
			resource "github_team" "test" {
				name        = "tf-acc-test-team-archive-%s"
				description = "test team for archived repo"
			}

			resource "github_repository" "test" {
				name = "tf-acc-test-team-archive-%[1]s"
				auto_init = true
				archived = true
			}

			resource "github_team_repository" "test" {
				team_id    = github_team.test.id
				repository = github_repository.test.name
				permission = "pull"
			}
		`, randomID)

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: config,
						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckResourceAttr(
								"github_team_repository.test", "permission",
								"pull",
							),
						),
					},
					{
						Config: archivedConfig,
						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckResourceAttr(
								"github_repository.test", "archived",
								"true",
							),
						),
					},
					{
						Config: fmt.Sprintf(`
							resource "github_team" "test" {
								name        = "tf-acc-test-team-archive-%s"
								description = "test team for archived repo"
							}

							resource "github_repository" "test" {
								name = "tf-acc-test-team-archive-%[1]s"
								auto_init = true
								archived = true
							}
						`, randomID),
					},
				},
			})
		}

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})
	})
}
