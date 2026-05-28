package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccGithubActionsOrganizationPermissions(t *testing.T) {
	t.Run("test setting of basic actions organization permissions", func(t *testing.T) {
		allowedActions := "local_only"
		enabledRepositories := "all"

		config := fmt.Sprintf(`
			resource "github_actions_organization_permissions" "test" {
				allowed_actions = "%s"
				enabled_repositories = "%s"
			}
		`, allowedActions, enabledRepositories)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_actions_organization_permissions.test", "allowed_actions", allowedActions,
			),
			resource.TestCheckResourceAttr(
				"github_actions_organization_permissions.test", "enabled_repositories", enabledRepositories,
			),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  check,
				},
			},
		})
	})

	t.Run("imports entire set of github action organization permissions without error", func(t *testing.T) {
		allowedActions := "selected"
		enabledRepositories := "selected"
		githubOwnedAllowed := true
		verifiedAllowed := true
		shaPinningRequired := true
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-act-org-perm-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name        = "%[1]s"
				description = "Terraform acceptance tests %[1]s"
				topics			= ["terraform", "testing"]
			}

			resource "github_actions_organization_permissions" "test" {
				allowed_actions = "%s"
				enabled_repositories = "%s"
				allowed_actions_config {
					github_owned_allowed = %t
					patterns_allowed     = ["actions/cache@*", "actions/checkout@*"]
					verified_allowed     = %t
				}
				sha_pinning_required = %t
				enabled_repositories_config {
					repository_ids       = [github_repository.test.repo_id]
				}
			}
		`, repoName, allowedActions, enabledRepositories, githubOwnedAllowed, verifiedAllowed, shaPinningRequired)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_actions_organization_permissions.test", "allowed_actions", allowedActions,
			),
			resource.TestCheckResourceAttr(
				"github_actions_organization_permissions.test", "enabled_repositories", enabledRepositories,
			),
			resource.TestCheckResourceAttr(
				"github_actions_organization_permissions.test", "allowed_actions_config.#", "1",
			),
			resource.TestCheckResourceAttr(
				"github_actions_organization_permissions.test", "enabled_repositories_config.#", "1",
			),
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
					ResourceName:      "github_actions_organization_permissions.test",
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})

	t.Run("test setting of organization allowed actions", func(t *testing.T) {
		allowedActions := "selected"
		enabledRepositories := "all"
		githubOwnedAllowed := true
		verifiedAllowed := true

		config := fmt.Sprintf(`

			resource "github_actions_organization_permissions" "test" {
				allowed_actions = "%s"
				enabled_repositories = "%s"
				allowed_actions_config {
					github_owned_allowed = %t
					patterns_allowed     = ["actions/cache@*", "actions/checkout@*"]
					verified_allowed     = %t
				}
			}
		`, allowedActions, enabledRepositories, githubOwnedAllowed, verifiedAllowed)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_actions_organization_permissions.test", "allowed_actions", allowedActions,
			),
			resource.TestCheckResourceAttr(
				"github_actions_organization_permissions.test", "enabled_repositories", enabledRepositories,
			),
			resource.TestCheckResourceAttr(
				"github_actions_organization_permissions.test", "allowed_actions_config.#", "1",
			),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  check,
				},
			},
		})
	})

	t.Run("test not setting of organization allowed actions without error", func(t *testing.T) {
		allowedActions := "selected"
		enabledRepositories := "all"

		config := fmt.Sprintf(`

			resource "github_actions_organization_permissions" "test" {
				allowed_actions = "%s"
				enabled_repositories = "%s"
			}
		`, allowedActions, enabledRepositories)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_actions_organization_permissions.test", "allowed_actions", allowedActions,
			),
			resource.TestCheckResourceAttr(
				"github_actions_organization_permissions.test", "enabled_repositories", enabledRepositories,
			),
			resource.TestCheckResourceAttr(
				"github_actions_organization_permissions.test", "allowed_actions_config.#", "0",
			),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  check,
				},
			},
		})
	})

	t.Run("test setting of organization enabled repositories", func(t *testing.T) {
		allowedActions := "all"
		enabledRepositories := "selected"
		githubOwnedAllowed := true
		verifiedAllowed := true
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		randomID2 := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-act-org-perm-%s", testResourcePrefix, randomID)
		repoName2 := fmt.Sprintf("%srepo-act-org-perm-%s", testResourcePrefix, randomID2)

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name        = "%[1]s"
				description = "Terraform acceptance tests %[1]s"
				topics			= ["terraform", "testing"]
			}

			resource "github_repository" "test2" {
				name        = "%[2]s"
				description = "Terraform acceptance tests %[2]s"
				topics			= ["terraform", "testing"]
			}

			resource "github_actions_organization_permissions" "test" {
				allowed_actions = "%s"
				enabled_repositories = "%s"
				enabled_repositories_config {
					repository_ids       = [github_repository.test.repo_id, github_repository.test2.repo_id]
				}
			}
		`, repoName, repoName2, allowedActions, enabledRepositories, githubOwnedAllowed, verifiedAllowed)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_actions_organization_permissions.test", "allowed_actions", allowedActions,
			),
			resource.TestCheckResourceAttr(
				"github_actions_organization_permissions.test", "enabled_repositories", enabledRepositories,
			),
			resource.TestCheckResourceAttr(
				"github_actions_organization_permissions.test", "enabled_repositories_config.#", "1",
			),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  check,
				},
			},
		})
	})
}

func TestGithubActionsOrganizationPermissionsOmitsAllowedActionsForNone(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.Method == http.MethodPut && r.URL.Path == "/orgs/test-org/actions/permissions":
			var body map[string]any
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatalf("decoding request body: %s", err)
			}
			if _, ok := body["allowed_actions"]; ok {
				t.Fatalf("allowed_actions should be omitted when enabled_repositories is none, got %#v", body)
			}
			if got := body["enabled_repositories"]; got != "none" {
				t.Fatalf("enabled_repositories = %#v, want none", got)
			}
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"enabled_repositories":"none","sha_pinning_required":false}`))
		case r.Method == http.MethodGet && r.URL.Path == "/orgs/test-org/actions/permissions":
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"enabled_repositories":"none","sha_pinning_required":false}`))
		default:
			t.Fatalf("unexpected request %s %s", r.Method, r.URL.String())
		}
	}))
	defer server.Close()

	resource := resourceGithubActionsOrganizationPermissions()
	data := schema.TestResourceDataRaw(t, resource.Schema, map[string]any{
		"enabled_repositories": "none",
	})
	meta := &Owner{
		name:           "test-org",
		IsOrganization: true,
		v3client:       mustGitHubClient(t, server.URL+"/"),
	}

	if err := resourceGithubActionsOrganizationPermissionsCreateOrUpdate(data, meta); err != nil {
		t.Fatalf("creating actions organization permissions: %s", err)
	}
}

func TestGithubActionsEnabledRepositoriesObjectAllowsEmptyConfig(t *testing.T) {
	resource := resourceGithubActionsOrganizationPermissions()
	data := schema.TestResourceDataRaw(t, resource.Schema, map[string]any{
		"enabled_repositories_config": []any{nil},
	})

	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("empty enabled_repositories_config should not panic: %v", r)
		}
	}()

	ids, err := resourceGithubActionsEnabledRepositoriesObject(data)
	if err != nil {
		t.Fatalf("empty enabled_repositories_config returned error: %s", err)
	}
	if len(ids) != 0 {
		t.Fatalf("repository ids = %#v, want empty", ids)
	}
}

func TestGithubActionsOrganizationPermissionsReadPreservesEmptySelectedRepositories(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.Method == http.MethodGet && r.URL.Path == "/orgs/test-org/actions/permissions":
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"enabled_repositories":"selected","allowed_actions":"local_only","sha_pinning_required":false}`))
		case r.Method == http.MethodGet && r.URL.Path == "/orgs/test-org/actions/permissions/repositories":
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"total_count":0,"repositories":[]}`))
		default:
			t.Fatalf("unexpected request %s %s", r.Method, r.URL.String())
		}
	}))
	defer server.Close()

	resource := resourceGithubActionsOrganizationPermissions()
	data := schema.TestResourceDataRaw(t, resource.Schema, map[string]any{
		"allowed_actions":      "local_only",
		"enabled_repositories": "selected",
		"enabled_repositories_config": []any{
			map[string]any{"repository_ids": []any{}},
		},
	})
	data.SetId("test-org")
	meta := &Owner{
		name:           "test-org",
		IsOrganization: true,
		v3client:       mustGitHubClient(t, server.URL+"/"),
	}

	if err := resourceGithubActionsOrganizationPermissionsRead(data, meta); err != nil {
		t.Fatalf("reading actions organization permissions: %s", err)
	}

	config := data.Get("enabled_repositories_config").([]any)
	if len(config) != 1 {
		t.Fatalf("enabled_repositories_config length = %d, want 1", len(config))
	}
	repositoryIDs := config[0].(map[string]any)["repository_ids"].(*schema.Set)
	if repositoryIDs.Len() != 0 {
		t.Fatalf("repository_ids length = %d, want 0", repositoryIDs.Len())
	}
}
