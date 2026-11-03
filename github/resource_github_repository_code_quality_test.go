package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/compare"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubRepositoryCodeQuality(t *testing.T) {
	t.Parallel()

	t.Run("creates_code_quality_with_defaults_without_error", func(t *testing.T) {
		t.Parallel()

		randomID := acctest.RandString(5)
		repoName := fmt.Sprintf("%scode-quality-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name       = "%s"
				visibility = "private"
				auto_init  = true
			}

			resource "github_repository_code_quality" "test" {
				repository = github_repository.test.name
			}
		`, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasPaidOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_code_quality.test", tfjsonpath.New("enabled"), knownvalue.Bool(true)),
						statecheck.CompareValuePairs("github_repository_code_quality.test", tfjsonpath.New("repository"), "github_repository.test", tfjsonpath.New("name"), compare.ValuesSame()),
						statecheck.ExpectKnownValue("github_repository_code_quality.test", tfjsonpath.New("repository_id"), knownvalue.NotNull()),
					},
				},
			},
		})
	})

	t.Run("updates_code_quality_without_error", func(t *testing.T) {
		t.Parallel()

		randomID := acctest.RandString(5)
		repoName := fmt.Sprintf("%scode-quality-%s", testResourcePrefix, randomID)

		config := `
			resource "github_repository" "test" {
				name       = "%s"
				visibility = "private"
				auto_init  = true
			}

			resource "github_repository_code_quality" "test" {
				repository = github_repository.test.name
				enabled    = %t
				languages  = ["python"]
			}
		`

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasPaidOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, repoName, true),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_code_quality.test", tfjsonpath.New("enabled"), knownvalue.Bool(true)),
						statecheck.ExpectKnownValue("github_repository_code_quality.test", tfjsonpath.New("languages"), knownvalue.SetExact([]knownvalue.Check{knownvalue.StringExact("python")})),
					},
				},
				{
					Config: fmt.Sprintf(config, repoName, false),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_code_quality.test", tfjsonpath.New("enabled"), knownvalue.Bool(false)),
					},
				},
			},
		})
	})

	t.Run("updates_renamed_repo", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)
		newRepoName := fmt.Sprintf("%s-updated", repo.GetName())

		config := `
			resource "github_repository_code_quality" "test" {
				repository = "%s"
				languages  = ["python"]
			}
		`

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasPaidOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, repo.GetName()),
				},
				{
					PreConfig: func() {
						mustRenameTestRepository(t, repo, newRepoName)
					},
					Config: fmt.Sprintf(config, newRepoName),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_repository_code_quality.test", plancheck.ResourceActionUpdate),
						},
					},
				},
			},
		})
	})

	t.Run("recreates_changed_repo", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)
		repo2 := mustCreateTestRepository(t)

		config := `
			resource "github_repository_code_quality" "test" {
				repository = "%s"
				languages  = ["python"]
			}
		`

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasPaidOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, repo.GetName()),
				},
				{
					Config: fmt.Sprintf(config, repo2.GetName()),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_repository_code_quality.test", plancheck.ResourceActionReplace),
						},
					},
				},
			},
		})
	})

	t.Run("imports_code_quality_without_error", func(t *testing.T) {
		t.Parallel()

		randomID := acctest.RandString(5)
		repoName := fmt.Sprintf("%scode-quality-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name       = "%s"
				visibility = "private"
				auto_init  = true
			}

			resource "github_repository_code_quality" "test" {
				repository = github_repository.test.name
				languages  = ["python"]
			}
		`, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasPaidOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
				},
				{
					ResourceName:      "github_repository_code_quality.test",
					ImportState:       true,
					ImportStateId:     repoName,
					ImportStateVerify: true,
				},
			},
		})
	})
}

func TestGithubRepositoryCodeQuality(t *testing.T) {
	// IMPORTANT: These tests are not parallelized because they use global state.

	t.Run("create_sends_configured_setup_and_accepts_202", func(t *testing.T) {
		ts := githubApiMock([]*mockResponse{
			{
				ExpectedUri:    "/repos/owner/repo",
				ExpectedMethod: "GET",
				StatusCode:     200,
				ResponseBody:   `{"id": 42, "name": "repo"}`,
			},
			{
				ExpectedUri:    "/repos/owner/repo/code-quality/setup",
				ExpectedMethod: "PATCH",
				ExpectedBody: []byte(`{"state":"configured","runner_type":"labeled","runner_label":"code-quality","languages":["python"]}
`),
				StatusCode:   202,
				ResponseBody: `{"run_id": 1, "run_url": "https://api.github.com/repos/owner/repo/actions/runs/1"}`,
			},
		})
		defer ts.Close()

		meta := &Owner{name: "owner", v3client: mustCreateTestGitHubClient(t, ts.URL)}
		d := schema.TestResourceDataRaw(t, resourceGithubRepositoryCodeQuality().Schema, map[string]any{
			"repository":   "repo",
			"languages":    []any{"python"},
			"runner_type":  "labeled",
			"runner_label": "code-quality",
		})

		diags := resourceGithubRepositoryCodeQualityCreate(t.Context(), d, meta)
		if diags.HasError() {
			t.Fatalf("expected no error, got: %v", diags)
		}
		if d.Id() != "42" {
			t.Fatalf("expected ID 42, got %q", d.Id())
		}
		if got, _ := d.Get("repository_id").(int); got != 42 {
			t.Fatalf("expected repository_id 42, got %d", got)
		}
	})

	t.Run("create_fails_on_archived_repository", func(t *testing.T) {
		ts := githubApiMock([]*mockResponse{
			{
				ExpectedUri:    "/repos/owner/repo",
				ExpectedMethod: "GET",
				StatusCode:     200,
				ResponseBody:   `{"id": 42, "name": "repo", "archived": true}`,
			},
		})
		defer ts.Close()

		meta := &Owner{name: "owner", v3client: mustCreateTestGitHubClient(t, ts.URL)}
		d := schema.TestResourceDataRaw(t, resourceGithubRepositoryCodeQuality().Schema, map[string]any{
			"repository": "repo",
		})

		diags := resourceGithubRepositoryCodeQualityCreate(t.Context(), d, meta)
		if !diags.HasError() {
			t.Fatal("expected an error for an archived repository")
		}
	})

	t.Run("update_sends_not_configured_when_disabled", func(t *testing.T) {
		ts := githubApiMock([]*mockResponse{
			{
				ExpectedUri:    "/repos/owner/repo/code-quality/setup",
				ExpectedMethod: "PATCH",
				ExpectedBody: []byte(`{"state":"not-configured"}
`),
				StatusCode:   200,
				ResponseBody: `{}`,
			},
		})
		defer ts.Close()

		meta := &Owner{name: "owner", v3client: mustCreateTestGitHubClient(t, ts.URL)}
		d := schema.TestResourceDataRaw(t, resourceGithubRepositoryCodeQuality().Schema, map[string]any{
			"repository": "repo",
			"enabled":    false,
			"languages":  []any{"python"},
		})
		d.SetId("42")

		diags := resourceGithubRepositoryCodeQualityUpdate(t.Context(), d, meta)
		if diags.HasError() {
			t.Fatalf("expected no error, got: %v", diags)
		}
	})

	t.Run("read_sets_configured_setup", func(t *testing.T) {
		ts := githubApiMock([]*mockResponse{
			{
				ExpectedUri:    "/repos/owner/repo/code-quality/setup",
				ExpectedMethod: "GET",
				StatusCode:     200,
				ResponseBody:   `{"state": "configured", "languages": ["python", "javascript-typescript"], "runner_type": "standard", "runner_label": null}`,
			},
		})
		defer ts.Close()

		meta := &Owner{name: "owner", v3client: mustCreateTestGitHubClient(t, ts.URL)}
		d := schema.TestResourceDataRaw(t, resourceGithubRepositoryCodeQuality().Schema, map[string]any{
			"repository": "repo",
		})
		d.SetId("42")

		diags := resourceGithubRepositoryCodeQualityRead(t.Context(), d, meta)
		if diags.HasError() {
			t.Fatalf("expected no error, got: %v", diags)
		}
		if enabled, _ := d.Get("enabled").(bool); !enabled {
			t.Fatal("expected enabled to be true")
		}
		if languages, _ := d.Get("languages").(*schema.Set); languages.Len() != 2 {
			t.Fatalf("expected 2 languages, got %d", languages.Len())
		}
		if got, _ := d.Get("runner_type").(string); got != "standard" {
			t.Fatalf("expected runner_type standard, got %q", got)
		}
	})

	t.Run("read_keeps_settings_when_not_configured", func(t *testing.T) {
		ts := githubApiMock([]*mockResponse{
			{
				ExpectedUri:    "/repos/owner/repo/code-quality/setup",
				ExpectedMethod: "GET",
				StatusCode:     200,
				ResponseBody:   `{"state": "not-configured", "languages": []}`,
			},
		})
		defer ts.Close()

		meta := &Owner{name: "owner", v3client: mustCreateTestGitHubClient(t, ts.URL)}
		d := schema.TestResourceDataRaw(t, resourceGithubRepositoryCodeQuality().Schema, map[string]any{
			"repository": "repo",
			"enabled":    false,
			"languages":  []any{"python"},
		})
		d.SetId("42")

		diags := resourceGithubRepositoryCodeQualityRead(t.Context(), d, meta)
		if diags.HasError() {
			t.Fatalf("expected no error, got: %v", diags)
		}
		if enabled, _ := d.Get("enabled").(bool); enabled {
			t.Fatal("expected enabled to be false")
		}
		if languages, _ := d.Get("languages").(*schema.Set); languages.Len() != 1 {
			t.Fatalf("expected configured languages to be kept, got %d", languages.Len())
		}
	})

	t.Run("read_removes_from_state_on_404", func(t *testing.T) {
		ts := githubApiMock([]*mockResponse{
			{
				ExpectedUri:    "/repos/owner/repo/code-quality/setup",
				ExpectedMethod: "GET",
				StatusCode:     404,
				ResponseBody:   `{"message": "Not Found"}`,
			},
		})
		defer ts.Close()

		meta := &Owner{name: "owner", v3client: mustCreateTestGitHubClient(t, ts.URL)}
		d := schema.TestResourceDataRaw(t, resourceGithubRepositoryCodeQuality().Schema, map[string]any{
			"repository": "repo",
		})
		d.SetId("42")

		diags := resourceGithubRepositoryCodeQualityRead(t.Context(), d, meta)
		if diags.HasError() {
			t.Fatalf("expected no error, got: %v", diags)
		}
		if d.Id() != "" {
			t.Fatalf("expected resource to be removed from state, got ID %q", d.Id())
		}
	})

	t.Run("rejects_unsupported_language", func(t *testing.T) {
		elem, _ := resourceGithubRepositoryCodeQuality().Schema["languages"].Elem.(*schema.Schema)
		if diags := elem.ValidateDiagFunc("python", cty.Path{}); diags.HasError() {
			t.Fatalf("expected python to be accepted, got: %v", diags)
		}
		if diags := elem.ValidateDiagFunc("cobol", cty.Path{}); !diags.HasError() {
			t.Fatal("expected cobol to be rejected")
		}
	})

	t.Run("delete_ignores_missing_repository", func(t *testing.T) {
		ts := githubApiMock([]*mockResponse{
			{
				ExpectedUri:    "/repos/owner/repo/code-quality/setup",
				ExpectedMethod: "PATCH",
				ExpectedBody: []byte(`{"state":"not-configured"}
`),
				StatusCode:   404,
				ResponseBody: `{"message": "Not Found"}`,
			},
		})
		defer ts.Close()

		meta := &Owner{name: "owner", v3client: mustCreateTestGitHubClient(t, ts.URL)}
		d := schema.TestResourceDataRaw(t, resourceGithubRepositoryCodeQuality().Schema, map[string]any{
			"repository": "repo",
		})
		d.SetId("42")

		diags := resourceGithubRepositoryCodeQualityDelete(t.Context(), d, meta)
		if diags.HasError() {
			t.Fatalf("expected no error, got: %v", diags)
		}
	})
}
