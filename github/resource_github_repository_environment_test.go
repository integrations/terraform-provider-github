package github

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubRepositoryEnvironment(t *testing.T) {
	t.Parallel()

	t.Run("create", func(t *testing.T) {
		t.Parallel()

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)
		envName := "test"

		config := fmt.Sprintf(`
resource "github_team" "test" {
	name        = "%[1]s"
	description = "test"
	privacy     = "closed"
}

resource "github_repository" "test" {
	name       = "%[1]s"
	visibility = "public"
}

resource "github_team_repository" "test" {
	team_id    = github_team.test.id
	repository = github_repository.test.name
	permission = "pull"
}

resource "github_repository_environment" "test" {
	repository  = github_repository.test.name
	environment = "%s"

	can_admins_bypass   = false
	wait_timer          = 10000
	prevent_self_review = true

	reviewers {
		teams = [github_team_repository.test.team_id]
	}

	deployment_branch_policy {
		protected_branches     = true
		custom_branch_policies = false
	}
}
`, repoName, envName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_environment.test", tfjsonpath.New("repository_id"), knownvalue.NotNull()),
					},
				},
			},
		})
	})

	t.Run("create_with_id_separator_in_name", func(t *testing.T) {
		t.Parallel()

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
resource "github_repository" "test" {
	name       = "%s"
	visibility = "public"
}

resource "github_repository_environment" "test" {
	repository  = github_repository.test.name
	environment = "environment:test"
}
`, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_environment.test", tfjsonpath.New("repository_id"), knownvalue.NotNull()),
					},
				},
			},
		})
	})

	t.Run("update_to_remove_reviewers", func(t *testing.T) {
		t.Parallel()

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)
		envName := "test"

		preConfig := fmt.Sprintf(`
resource "github_team" "test" {
	name        = "%[1]s"
	description = "test"
	privacy     = "closed"
}

resource "github_repository" "test" {
	name      = "%[1]s"
	visibility = "public"
}

resource "github_team_repository" "test" {
	team_id    = github_team.test.id
	repository = github_repository.test.name
	permission = "pull"
}
`, repoName)

		config := fmt.Sprintf(`
%s

resource "github_repository_environment" "test" {
	repository  = github_repository.test.name
	environment = "%s"

	reviewers {
		teams = [github_team_repository.test.team_id]
	}
}
`, preConfig, envName)

		configUpdated := fmt.Sprintf(`
%s

resource "github_repository_environment" "test" {
	repository  = github_repository.test.name
	environment = "%s"
}
`, preConfig, envName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_environment.test", tfjsonpath.New("repository_id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_repository_environment.test", tfjsonpath.New("reviewers"), knownvalue.ListSizeExact(1)),
					},
				},
				{
					Config: configUpdated,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_environment.test", tfjsonpath.New("reviewers"), knownvalue.ListSizeExact(0)),
					},
				},
			},
		})
	})

	t.Run("update_to_add_reviewers", func(t *testing.T) {
		t.Parallel()

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)
		envName := "test"

		preConfig := fmt.Sprintf(`
resource "github_team" "test" {
	name        = "%[1]s"
	description = "test"
	privacy     = "closed"
}

resource "github_repository" "test" {
	name      = "%[1]s"
	visibility = "public"
}

resource "github_team_repository" "test" {
	team_id    = github_team.test.id
	repository = github_repository.test.name
	permission = "pull"
}
`, repoName)

		config := fmt.Sprintf(`
%s

resource "github_repository_environment" "test" {
	repository  = github_repository.test.name
	environment = "%s"
}
`, preConfig, envName)

		configUpdated := fmt.Sprintf(`
%s

resource "github_repository_environment" "test" {
	repository  = github_repository.test.name
	environment = "%s"

	reviewers {
		teams = [github_team_repository.test.team_id]
	}
}
`, preConfig, envName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_environment.test", tfjsonpath.New("repository_id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_repository_environment.test", tfjsonpath.New("reviewers"), knownvalue.ListSizeExact(0)),
					},
				},
				{
					Config: configUpdated,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_environment.test", tfjsonpath.New("reviewers"), knownvalue.ListSizeExact(1)),
					},
				},
			},
		})
	})

	t.Run("import", func(t *testing.T) {
		t.Parallel()

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
resource "github_repository" "test" {
	name       = "%s"
	visibility = "public"
}

resource "github_repository_environment" "test" {
	repository 	= github_repository.test.name
	environment	= "test"
}
`, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_environment.test", tfjsonpath.New("repository_id"), knownvalue.NotNull()),
					},
				},
				{
					ResourceName:            "github_repository_environment.test",
					ImportState:             true,
					ImportStateVerify:       true,
					ImportStateVerifyIgnore: []string{"can_admins_bypass", "prevent_self_review", "reviewers", "wait_timer", "deployment_branch_policy"},
				},
			},
		})
	})

	t.Run("errors_with_more_than_six_reviewers", func(t *testing.T) {
		t.Parallel()

		skipUnlessHasOrgs(t)
		skipUnlessHasOrgUser1(t)

		repo := mustCreateTestRepository(t)
		user := mustGetUser(t, testAccConf.testOrgUser1)
		mustAddRepositoryCollaborator(t, repo, user.GetLogin())

		teamIDs := make([]string, 0, 6)
		for range 6 {
			team := mustCreateTestTeam(t)
			mustAddRepositoryToTeam(t, team, repo)
			teamIDs = append(teamIDs, strconv.FormatInt(team.GetID(), 10))
		}

		config := fmt.Sprintf(`
resource "github_repository_environment" "test" {
  repository  = "%s"
  environment = "test"

  reviewers {
    teams = [%s]
    users = [%d]
  }
}
`, repo.GetName(), strings.Join(teamIDs, ", "), user.GetID())

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      config,
					ExpectError: regexp.MustCompile(`reviewers can have at most 6 reviewers`),
				},
			},
		})
	})
}

func TestGithubRepositoryEnvironmentReadMissingReviewerPayload(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet || r.URL.Path != "/repos/test-org/my-repo/environments/prod" {
			t.Errorf("unexpected request %s %s", r.Method, r.URL.Path)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{
			"can_admins_bypass": false,
			"protection_rules": [
				{
					"type": "required_reviewers",
					"prevent_self_review": true,
					"reviewers": [
						{"type": "Team", "reviewer": null},
						{"type": "Team", "reviewer": {"id": 101}},
						{"type": "User", "reviewer": null},
						{"type": "User", "reviewer": {"id": 202}}
					]
				}
			]
		}`))
	}))
	t.Cleanup(srv.Close)

	d := resourceGithubRepositoryEnvironment().TestResourceData()
	d.SetId("my-repo:prod")
	if err := d.Set("repository", "my-repo"); err != nil {
		t.Fatalf("failed setting repository: %v", err)
	}
	if err := d.Set("environment", "prod"); err != nil {
		t.Fatalf("failed setting environment: %v", err)
	}

	meta := &Owner{
		name:     "test-org",
		v3client: mustCreateTestGitHubClient(t, srv.URL),
	}

	diags := resourceGithubRepositoryEnvironmentRead(t.Context(), d, meta)
	if diags.HasError() {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}

	reviewersRaw, ok := d.Get("reviewers").([]any)
	if !ok || len(reviewersRaw) != 1 {
		t.Fatalf("expected 1 reviewers block, got %#v", d.Get("reviewers"))
	}
	reviewerBlock, ok := reviewersRaw[0].(map[string]any)
	if !ok {
		t.Fatalf("expected reviewers[0] to be map[string]any, got %T", reviewersRaw[0])
	}
	teamsSet, ok := reviewerBlock["teams"].(*schema.Set)
	if !ok {
		t.Fatalf("expected teams *schema.Set, got %T", reviewerBlock["teams"])
	}
	usersSet, ok := reviewerBlock["users"].(*schema.Set)
	if !ok {
		t.Fatalf("expected users *schema.Set, got %T", reviewerBlock["users"])
	}
	if teamsSet.Len() != 1 || !teamsSet.Contains(101) {
		t.Fatalf("got teams %v, want [101]", teamsSet.List())
	}
	if usersSet.Len() != 1 || !usersSet.Contains(202) {
		t.Fatalf("got users %v, want [202]", usersSet.List())
	}
}
