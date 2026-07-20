package github

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// Test_resourceGithubTeamRepositoryDelete_nilResponseDoesNotPanic is a
// regression test for a nil-pointer panic in the delete path. When the
// RemoveTeamRepoByID call fails at the transport layer (for example a
// cancelled context or an exceeded deadline), go-github returns a nil
// *github.Response. Delete must surface that as an error diagnostic instead of
// dereferencing resp.StatusCode and panicking.
func Test_resourceGithubTeamRepositoryDelete_nilResponseDoesNotPanic(t *testing.T) {
	t.Parallel()

	ts := githubApiMock(nil)
	defer ts.Close()

	meta := &Owner{
		name:           "test-org",
		id:             12345,
		v3client:       mustCreateTestGitHubClient(t, ts.URL),
		IsOrganization: true,
	}

	d := schema.TestResourceDataRaw(t, resourceGithubTeamRepository().Schema, map[string]any{})
	d.SetId(buildTwoPartID("123", "test-repo"))

	// A cancelled context makes the underlying HTTP call return before it
	// reaches the server, so go-github yields a nil response with an error.
	ctx, cancel := context.WithCancel(t.Context())
	cancel()

	diags := resourceGithubTeamRepositoryDelete(ctx, d, meta)
	if !diags.HasError() {
		t.Fatal("expected an error diagnostic when the API response is nil, got none")
	}
}

func TestAccGithubTeamRepository(t *testing.T) {
	t.Parallel()

	t.Run("manages team permissions to a repository", func(t *testing.T) {
		t.Parallel()

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		teamName := fmt.Sprintf("%steam-repo-%s", testResourcePrefix, randomID)
		repoName := fmt.Sprintf("%srepo-team-repo-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
			resource "github_team" "test" {
				name        = "%s"
				description = "test"
			}

			resource "github_repository" "test" {
				name = "%s"
			}

			resource "github_team_repository" "test" {
				team_id    = "${github_team.test.id}"
				repository = "${github_repository.test.name}"
				permission = "pull"
			}
		`, teamName, repoName)

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

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
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
	})

	t.Run("accepts both team slug and team ID for `team_id`", func(t *testing.T) {
		t.Parallel()

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		teamName := fmt.Sprintf("%steam-repo-slug-%s", testResourcePrefix, randomID)
		repoName := fmt.Sprintf("%srepo-team-repo-slug-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
			resource "github_team" "test" {
				name        = "%s"
				description = "test"
			}

			resource "github_repository" "test" {
				name = "%s"
			}

			resource "github_team_repository" "test" {
				team_id    = "${github_team.test.slug}"
				repository = "${github_repository.test.name}"
				permission = "pull"
			}
		`, teamName, repoName)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet("github_team_repository.test", "team_id"),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
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
	})
}

func TestAccGithubTeamRepositoryArchivedRepo(t *testing.T) {
	t.Parallel()

	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
	teamName := fmt.Sprintf("%steam-archive-%s", testResourcePrefix, randomID)
	repoName := fmt.Sprintf("%srepo-team-archive-%s", testResourcePrefix, randomID)

	t.Run("can delete team repository access from archived repositories without error", func(t *testing.T) {
		t.Parallel()

		config := fmt.Sprintf(`
			resource "github_team" "test" {
				name        = "%s"
				description = "test team for archived repo"
			}

			resource "github_repository" "test" {
				name = "%s"
				auto_init = true
			}

			resource "github_team_repository" "test" {
				team_id    = github_team.test.id
				repository = github_repository.test.name
				permission = "pull"
			}
		`, teamName, repoName)

		archivedConfig := fmt.Sprintf(`
			resource "github_team" "test" {
				name        = "%s"
				description = "test team for archived repo"
			}

			resource "github_repository" "test" {
				name = "%s"
				auto_init = true
				archived = true
			}

			resource "github_team_repository" "test" {
				team_id    = github_team.test.id
				repository = github_repository.test.name
				permission = "pull"
			}
		`, teamName, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
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
								name        = "%s"
								description = "test team for archived repo"
							}

							resource "github_repository" "test" {
								name = "%s"
								auto_init = true
								archived = true
							}
						`, teamName, repoName),
				},
			},
		})
	})
}
