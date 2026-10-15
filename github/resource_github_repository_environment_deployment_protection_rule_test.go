package github

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubRepositoryEnvironmentDeploymentProtectionRule(t *testing.T) {
	t.Parallel()

	skipUnauthenticated(t)

	config := `
resource "github_repository_environment_deployment_protection_rule" "test" {
	repository     = "%s"
	environment    = "%s"
	integration_id = %d
}
`

	t.Run("creates_and_deletes", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)
		env := mustCreateTestRepositoryEnvironment(t, repo, withTestCreateName("environment/test"))
		appID := mustFindDeploymentProtectionRuleApp(t, repo, env)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			CheckDestroy:      testAccCheckGithubRepositoryEnvironmentDeploymentProtectionRuleDestroy(repo.GetName(), env.GetName()),
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, repo.GetName(), env.GetName(), appID),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_environment_deployment_protection_rule.test", tfjsonpath.New("repository_id"), knownvalue.Int64Exact(repo.GetID())),
						statecheck.ExpectKnownValue("github_repository_environment_deployment_protection_rule.test", tfjsonpath.New("protection_rule_id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_repository_environment_deployment_protection_rule.test", tfjsonpath.New("integration_id"), knownvalue.Int64Exact(appID)),
					},
				},
			},
		})
	})

	t.Run("imports", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)
		env := mustCreateTestRepositoryEnvironment(t, repo)
		appID := mustFindDeploymentProtectionRuleApp(t, repo, env)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, repo.GetName(), env.GetName(), appID),
				},
				{
					ResourceName:      "github_repository_environment_deployment_protection_rule.test",
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})

	t.Run("updates_renamed_repo", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)
		env := mustCreateTestRepositoryEnvironment(t, repo)
		newRepoName := fmt.Sprintf("%s-updated", repo.GetName())
		appID := mustFindDeploymentProtectionRuleApp(t, repo, env)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, repo.GetName(), env.GetName(), appID),
				},
				{
					PreConfig: func() {
						mustRenameTestRepository(t, repo, newRepoName)
					},
					Config: fmt.Sprintf(config, newRepoName, env.GetName(), appID),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_repository_environment_deployment_protection_rule.test", plancheck.ResourceActionUpdate),
						},
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_environment_deployment_protection_rule.test", tfjsonpath.New("repository_id"), knownvalue.Int64Exact(repo.GetID())),
					},
				},
			},
		})
	})

	t.Run("recreates_changed_repo", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)
		env := mustCreateTestRepositoryEnvironment(t, repo)
		repo2 := mustCreateTestRepository(t)
		_ = mustCreateTestRepositoryEnvironment(t, repo2, withTestCreateName(env.GetName()))
		appID := mustFindDeploymentProtectionRuleApp(t, repo, env)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, repo.GetName(), env.GetName(), appID),
				},
				{
					Config: fmt.Sprintf(config, repo2.GetName(), env.GetName(), appID),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_repository_environment_deployment_protection_rule.test", plancheck.ResourceActionReplace),
						},
					},
				},
			},
		})
	})
}

// mustFindDeploymentProtectionRuleApp returns the first app the test owner can use as a custom deployment
// protection rule on the environment, and skips the test when there is none.
func mustFindDeploymentProtectionRuleApp(t *testing.T, repo *github.Repository, env *github.Environment) int64 {
	t.Helper()

	meta := testAccConf.meta
	res, _, err := meta.v3client.Repositories.ListCustomDeploymentRuleIntegrations(t.Context(), meta.name, repo.GetName(), env.GetName(), nil)
	if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response.StatusCode == http.StatusForbidden {
		t.Skip("Skipping as the test credentials cannot list custom deployment protection rule apps")
	}
	if err != nil {
		t.Fatalf("failed to list custom deployment protection rule integrations: %v", err)
	}
	if len(res.AvailableIntegrations) == 0 {
		t.Skip("Skipping as no app is available as a custom deployment protection rule")
	}

	return res.AvailableIntegrations[0].GetID()
}

func testAccCheckGithubRepositoryEnvironmentDeploymentProtectionRuleDestroy(repoName, envName string) resource.TestCheckFunc {
	return func(_ *terraform.State) error {
		meta := testAccConf.meta
		rules, _, err := meta.v3client.Repositories.GetAllDeploymentProtectionRules(context.Background(), meta.name, repoName, envName)
		if err != nil {
			return err
		}
		if n := len(rules.ProtectionRules); n != 0 {
			return fmt.Errorf("expected no custom deployment protection rules on %s/%s, found %d", repoName, envName, n)
		}
		return nil
	}
}

func Test_resourceGithubRepositoryEnvironmentDeploymentProtectionRuleRead(t *testing.T) {
	t.Parallel()

	const (
		owner    = "test-org"
		repoName = "test-repo"
		envName  = "environment/test"
		ruleID   = 42
	)

	ruleURI := fmt.Sprintf("/repos/%s/%s/environments/environment%%2Ftest/deployment_protection_rules/%d", owner, repoName, ruleID)

	for _, tt := range []struct {
		name              string
		status            int
		body              any
		wantErr           bool
		wantRemoved       bool
		wantIntegrationID int
	}{
		{
			name:              "sets integration_id from the rule's app",
			status:            http.StatusOK,
			body:              &github.CustomDeploymentProtectionRule{ID: new(int64(ruleID)), App: &github.CustomDeploymentProtectionRuleApp{ID: new(int64(1234))}},
			wantIntegrationID: 1234,
		},
		{
			name:              "replaces a stale integration_id when the rule has no app",
			status:            http.StatusOK,
			body:              &github.CustomDeploymentProtectionRule{ID: new(int64(ruleID))},
			wantIntegrationID: 0,
		},
		{
			name:        "removes the resource from state on 404",
			status:      http.StatusNotFound,
			wantRemoved: true,
		},
		{
			name:    "returns other errors",
			status:  http.StatusInternalServerError,
			wantErr: true,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ts := githubApiMock([]*mockResponse{
				mustGetTestMockResponse(t, ruleURI, tt.status, tt.body),
			})
			t.Cleanup(ts.Close)

			d := schema.TestResourceDataRaw(t, resourceGithubRepositoryEnvironmentDeploymentProtectionRule().Schema, map[string]any{
				"repository":         repoName,
				"environment":        envName,
				"integration_id":     9999,
				"protection_rule_id": ruleID,
			})
			d.SetId(fmt.Sprintf("%s:%s:%d", repoName, envName, ruleID))

			meta := &Owner{name: owner, v3client: mustCreateTestGitHubClient(t, ts.URL)}

			diags := resourceGithubRepositoryEnvironmentDeploymentProtectionRuleRead(t.Context(), d, meta)

			if diags.HasError() != tt.wantErr {
				t.Fatalf("expected error to be %v, got %v", tt.wantErr, diags)
			}
			if tt.wantErr {
				return
			}
			if removed := d.Id() == ""; removed != tt.wantRemoved {
				t.Fatalf("expected removed from state to be %v, got %v", tt.wantRemoved, removed)
			}
			if tt.wantRemoved {
				return
			}
			if got, _ := d.Get("integration_id").(int); got != tt.wantIntegrationID {
				t.Errorf("expected integration_id to be %d, got %d", tt.wantIntegrationID, got)
			}
		})
	}
}

func Test_resourceGithubRepositoryEnvironmentDeploymentProtectionRuleDelete(t *testing.T) {
	t.Parallel()

	const (
		owner    = "test-org"
		repoName = "test-repo"
		envName  = "environment/test"
		ruleID   = 42
	)

	ruleURI := fmt.Sprintf("/repos/%s/%s/environments/environment%%2Ftest/deployment_protection_rules/%d", owner, repoName, ruleID)

	for _, tt := range []struct {
		name    string
		status  int
		wantErr bool
	}{
		{
			name:   "disables the rule",
			status: http.StatusNoContent,
		},
		{
			name:   "treats 404 as already deleted",
			status: http.StatusNotFound,
		},
		{
			name:    "returns other errors",
			status:  http.StatusInternalServerError,
			wantErr: true,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			resp := mustGetTestMockResponse(t, ruleURI, tt.status, nil)
			resp.ExpectedMethod = http.MethodDelete
			ts := githubApiMock([]*mockResponse{resp})
			t.Cleanup(ts.Close)

			d := schema.TestResourceDataRaw(t, resourceGithubRepositoryEnvironmentDeploymentProtectionRule().Schema, map[string]any{
				"repository":         repoName,
				"environment":        envName,
				"integration_id":     1234,
				"protection_rule_id": ruleID,
			})
			d.SetId(fmt.Sprintf("%s:%s:%d", repoName, envName, ruleID))

			meta := &Owner{name: owner, v3client: mustCreateTestGitHubClient(t, ts.URL)}

			diags := resourceGithubRepositoryEnvironmentDeploymentProtectionRuleDelete(t.Context(), d, meta)

			if diags.HasError() != tt.wantErr {
				t.Fatalf("expected error to be %v, got %v", tt.wantErr, diags)
			}
		})
	}
}

func Test_resourceGithubRepositoryEnvironmentDeploymentProtectionRuleUpdate(t *testing.T) {
	t.Parallel()

	const (
		owner       = "test-org"
		oldRepoName = "test-repo"
		newRepoName = "test-repo-renamed"
		envName     = "environment/test"
		repoID      = 100
		ruleID      = 42
	)

	repoURI := fmt.Sprintf("/repos/%s/%s", owner, newRepoName)

	for _, tt := range []struct {
		name    string
		status  int
		body    any
		wantErr bool
	}{
		{
			name:   "updates the ID after a rename",
			status: http.StatusOK,
			body:   &github.Repository{ID: new(int64(repoID)), Name: new(newRepoName)},
		},
		{
			name:    "errors when the name now belongs to another repository",
			status:  http.StatusOK,
			body:    &github.Repository{ID: new(int64(repoID + 1)), Name: new(newRepoName)},
			wantErr: true,
		},
		{
			name:    "errors when the repository does not exist",
			status:  http.StatusNotFound,
			wantErr: true,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ts := githubApiMock([]*mockResponse{
				mustGetTestMockResponse(t, repoURI, tt.status, tt.body),
			})
			t.Cleanup(ts.Close)

			d := schema.TestResourceDataRaw(t, resourceGithubRepositoryEnvironmentDeploymentProtectionRule().Schema, map[string]any{
				"repository":         newRepoName,
				"repository_id":      repoID,
				"environment":        envName,
				"integration_id":     1234,
				"protection_rule_id": ruleID,
			})
			oldID := fmt.Sprintf("%s:%s:%d", oldRepoName, envName, ruleID)
			d.SetId(oldID)

			meta := &Owner{name: owner, v3client: mustCreateTestGitHubClient(t, ts.URL)}

			diags := resourceGithubRepositoryEnvironmentDeploymentProtectionRuleUpdate(t.Context(), d, meta)

			if diags.HasError() != tt.wantErr {
				t.Fatalf("expected error to be %v, got %v", tt.wantErr, diags)
			}
			if tt.wantErr {
				if d.Id() != oldID {
					t.Errorf("expected ID to stay %q on error, got %q", oldID, d.Id())
				}
				return
			}
			if want := fmt.Sprintf("%s:%s:%d", newRepoName, escapeIDPart(envName), ruleID); d.Id() != want {
				t.Errorf("expected ID to be %q, got %q", want, d.Id())
			}
		})
	}
}
