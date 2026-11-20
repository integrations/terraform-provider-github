package github

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccGithubRepositoryEnvironmentsDataSource(t *testing.T) {
	t.Parallel()

	t.Run("queries environments", func(t *testing.T) {
		t.Parallel()

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-envs-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
			  name = "%[1]s"
				auto_init = true
			}
			resource "github_repository_environment" "env1" {
				repository = github_repository.test.name
				environment = "env_x"
			}
		`, repoName)

		config2 := config + `
			data "github_repository_environments" "all" {
				repository = github_repository.test.name
			}
		`
		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("data.github_repository_environments.all", "environments.#", "1"),
			resource.TestCheckResourceAttr("data.github_repository_environments.all", "environments.0.name", "env_x"),
			resource.TestCheckResourceAttrSet("data.github_repository_environments.all", "environments.0.node_id"),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
				},
				{
					Config: config2,
					Check:  check,
				},
			},
		})
	})
}

func Test_dataSourceGithubRepositoryEnvironmentsRead(t *testing.T) {
	t.Parallel()

	const (
		owner    = "test-org"
		repoName = "test-repo"
	)

	envsURI := fmt.Sprintf("/repos/%s/%s/environments", owner, repoName)

	page1 := mustGetTestMockResponse(t, envsURI, http.StatusOK, &github.EnvResponse{
		TotalCount:   new(2),
		Environments: []*github.Environment{{Name: new("env_1"), NodeID: new("node_1")}},
	})
	page1.ResponseHeaders = map[string]string{
		"Link": fmt.Sprintf(`<https://api.github.com%s?page=2>; rel="next", <https://api.github.com%s?page=2>; rel="last"`, envsURI, envsURI),
	}
	page2 := mustGetTestMockResponse(t, envsURI+"?page=2", http.StatusOK, &github.EnvResponse{
		TotalCount:   new(2),
		Environments: []*github.Environment{{Name: new("env_2"), NodeID: new("node_2")}},
	})

	ts := githubApiMock([]*mockResponse{page1, page2})
	t.Cleanup(ts.Close)

	d := schema.TestResourceDataRaw(t, dataSourceGithubRepositoryEnvironments().Schema, map[string]any{
		"repository": repoName,
	})

	meta := &Owner{name: owner, v3client: mustCreateTestGitHubClient(t, ts.URL)}

	if diags := dataSourceGithubRepositoryEnvironmentsRead(t.Context(), d, meta); diags.HasError() {
		t.Fatalf("unexpected error: %v", diags)
	}

	envs, ok := d.Get("environments").([]any)
	if !ok {
		t.Fatalf("expected environments to be a list, got %T", d.Get("environments"))
	}
	if len(envs) != 2 {
		t.Fatalf("expected 2 environments across both pages, got %d", len(envs))
	}
	for i, want := range []string{"env_1", "env_2"} {
		env, ok := envs[i].(map[string]any)
		if !ok {
			t.Fatalf("expected environments.%d to be a map, got %T", i, envs[i])
		}
		if got := env["name"]; got != want {
			t.Errorf("expected environments.%d.name to be %q, got %q", i, want, got)
		}
	}
}
