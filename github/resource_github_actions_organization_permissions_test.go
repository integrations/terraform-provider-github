package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	gogithub "github.com/google/go-github/v88/github"
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

func TestGithubActionsOrganizationPermissionsSendsEmptyPatternsAllowed(t *testing.T) {
	resource := resourceGithubActionsOrganizationPermissions()
	var selectedActionsPayload map[string]any

	mux := http.NewServeMux()
	mux.HandleFunc("/orgs/test-org/actions/permissions", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPut:
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"enabled_repositories":"all","allowed_actions":"selected","sha_pinning_required":true}`))
		case http.MethodGet:
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"enabled_repositories":"all","allowed_actions":"selected","sha_pinning_required":true}`))
		default:
			t.Fatalf("unexpected method %s for %s", r.Method, r.URL.Path)
		}
	})
	mux.HandleFunc("/orgs/test-org/actions/permissions/selected-actions", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPut:
			if err := json.NewDecoder(r.Body).Decode(&selectedActionsPayload); err != nil {
				t.Fatalf("decode selected actions payload: %v", err)
			}
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"github_owned_allowed":true,"patterns_allowed":[],"verified_allowed":true}`))
		case http.MethodGet:
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"github_owned_allowed":true,"patterns_allowed":[],"verified_allowed":true}`))
		default:
			t.Fatalf("unexpected method %s for %s", r.Method, r.URL.Path)
		}
	})

	server := httptest.NewServer(mux)
	defer server.Close()

	baseURL := server.URL + "/"
	client, err := gogithub.NewClient(gogithub.WithURLs(&baseURL, &baseURL))
	if err != nil {
		t.Fatalf("create GitHub client: %v", err)
	}

	d := schema.TestResourceDataRaw(t, resource.Schema, map[string]any{
		"allowed_actions":      "selected",
		"enabled_repositories": "all",
		"allowed_actions_config": []any{
			map[string]any{
				"github_owned_allowed": true,
				"patterns_allowed":     []any{},
				"verified_allowed":     true,
			},
		},
		"sha_pinning_required": true,
	})

	err = resourceGithubActionsOrganizationPermissionsCreateOrUpdate(d, &Owner{
		name:           "test-org",
		v3client:       client,
		IsOrganization: true,
	})
	if err != nil {
		t.Fatalf("create/update actions organization permissions: %v", err)
	}

	patterns, ok := selectedActionsPayload["patterns_allowed"]
	if !ok {
		t.Fatalf("selected actions payload omitted patterns_allowed: %#v", selectedActionsPayload)
	}
	if got := len(patterns.([]any)); got != 0 {
		t.Fatalf("patterns_allowed length = %d, want 0", got)
	}
	if got := selectedActionsPayload["verified_allowed"]; got != true {
		t.Fatalf("verified_allowed = %v, want true", got)
	}
}
