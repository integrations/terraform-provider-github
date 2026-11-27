package github

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
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

	skipUnlessHasOrgs(t)

	t.Run("manages_team_permissions_to_a_repository", func(t *testing.T) {
		t.Parallel()

		team := mustCreateTestTeam(t)
		repo := mustCreateTestRepository(t)

		config := fmt.Sprintf(`
resource "github_team_repository" "test" {
  team_id    = "%d"
  repository = "%s"
  permission = "%%s"
}
`, team.GetID(), repo.GetName())

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, "pull"),
				},
				{
					Config: fmt.Sprintf(config, "triage"),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_team_repository.test", plancheck.ResourceActionUpdate),
						},
					},
				},
				{
					Config: fmt.Sprintf(config, "push"),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_team_repository.test", plancheck.ResourceActionUpdate),
						},
					},
				},
				{
					Config: fmt.Sprintf(config, "maintain"),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_team_repository.test", plancheck.ResourceActionUpdate),
						},
					},
				},
				{
					Config: fmt.Sprintf(config, "admin"),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_team_repository.test", plancheck.ResourceActionUpdate),
						},
					},
				},
				{
					ResourceName:            "github_team_repository.test",
					ImportState:             true,
					ImportStateVerify:       true,
					ImportStateVerifyIgnore: []string{"etag"},
				},
			},
		})
	})

	t.Run("accepts_both_team_slug_and_team_id_for_team_id", func(t *testing.T) {
		t.Parallel()

		team := mustCreateTestTeam(t)
		repo := mustCreateTestRepository(t)

		config := fmt.Sprintf(`
resource "github_team_repository" "test" {
  team_id    = "%%s"
  repository = "%s"
  permission = "pull"
}
`, repo.GetName())

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, team.GetSlug()),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_team_repository.test", tfjsonpath.New("id"), knownvalue.StringExact(fmt.Sprintf("%d:%s", team.GetID(), repo.GetName()))),
					},
				},
				{
					Config: fmt.Sprintf(config, strconv.FormatInt(team.GetID(), 10)),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_team_repository.test", tfjsonpath.New("id"), knownvalue.StringExact(fmt.Sprintf("%d:%s", team.GetID(), repo.GetName()))),
					},
				},
			},
		})
	})

	t.Run("can_delete_team_repository_access_from_archived_repositories_without_error", func(t *testing.T) {
		t.Parallel()

		team := mustCreateTestTeam(t)
		repo := mustCreateTestRepository(t)

		config := fmt.Sprintf(`
resource "github_team_repository" "test" {
  team_id    = "%d"
  repository = "%s"
  permission = "pull"
}
`, team.GetID(), repo.GetName())

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
				},
				{
					PreConfig: func() {
						mustArchiveTestRepository(t, repo)
					},
					Config: config,
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
				{
					Config: "# no resources",
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_team_repository.test", plancheck.ResourceActionDestroy),
						},
					},
				},
			},
		})
	})
}
