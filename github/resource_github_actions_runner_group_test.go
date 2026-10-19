package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/google/go-github/v93/github"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/compare"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestGithubActionsRunnerGroupNetworking(t *testing.T) {
	testRunnerGroupNetworking(t, resourceGithubActionsRunnerGroup(), "/orgs/test-org", map[string]any{
		"name":                     "test-group",
		"visibility":               "all",
		"network_configuration_id": "network-1",
	})
}

func TestGithubActionsRunnerGroupContextCancellation(t *testing.T) {
	testRunnerGroupContextCancellation(t, resourceGithubActionsRunnerGroup(), map[string]any{
		"name":       "test-group",
		"visibility": "all",
	})
}

func TestGithubActionsRunnerGroupNetworkingLifecycle(t *testing.T) {
	testRunnerGroupNetworkingLifecycle(t, resourceGithubActionsRunnerGroup, "github_actions_runner_group", "/orgs/test-org", "")
}

func TestRunnerGroupNetworkConfigurationSchema(t *testing.T) {
	for name, r := range map[string]*schema.Resource{
		"organization": resourceGithubActionsRunnerGroup(),
		"enterprise":   resourceGithubActionsEnterpriseRunnerGroup(),
	} {
		t.Run(name, func(t *testing.T) {
			attribute := r.Schema["network_configuration_id"]
			if !attribute.Optional || !attribute.Computed {
				t.Fatal("network_configuration_id must be Optional and Computed")
			}
			if attribute.ValidateDiagFunc == nil {
				t.Fatal("network_configuration_id must use ValidateDiagFunc")
			}
			for _, tc := range []struct {
				value     string
				wantError bool
			}{
				{value: ""},
				{value: "network-1"},
				{value: " ", wantError: true},
				{value: "\t\n", wantError: true},
			} {
				if diags := attribute.ValidateDiagFunc(tc.value, nil); diags.HasError() != tc.wantError {
					t.Errorf("network_configuration_id %q: diagnostics = %v, want error = %t", tc.value, diags, tc.wantError)
				}
			}
		})
	}
}

func TestGithubActionsRunnerGroupReadErrors(t *testing.T) {
	for _, tc := range []struct {
		name       string
		statusCode int
		wantID     string
		wantError  bool
	}{
		{"not modified", http.StatusNotModified, "42", false},
		{"not found", http.StatusNotFound, "", false},
		{"forbidden", http.StatusForbidden, "42", true},
		{"server error", http.StatusInternalServerError, "42", true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			server := githubApiMock([]*mockResponse{{
				ExpectedUri:    "/orgs/test-org/actions/runner-groups/42",
				ExpectedMethod: http.MethodGet,
				StatusCode:     tc.statusCode,
			}})
			defer server.Close()

			meta := &Owner{
				name:           "test-org",
				IsOrganization: true,
				v3client:       mustCreateTestGitHubClient(t, server.URL),
			}
			d := schema.TestResourceDataRaw(t, resourceGithubActionsRunnerGroup().Schema, map[string]any{
				"name":                     "test-group",
				"visibility":               "all",
				"network_configuration_id": "network-1",
			})
			d.SetId("42")

			diags := resourceGithubActionsRunnerGroupRead(t.Context(), d, meta)
			if diags.HasError() != tc.wantError {
				t.Fatalf("read diagnostics = %v, want error = %t", diags, tc.wantError)
			}
			if d.Id() != tc.wantID {
				t.Errorf("ID = %q, want %q", d.Id(), tc.wantID)
			}
			if got := d.Get("network_configuration_id"); got != "network-1" {
				t.Errorf("network_configuration_id = %v, want network-1", got)
			}
		})
	}
}

func TestRunnerGroupNetworkConfigurationUpdate(t *testing.T) {
	for _, scope := range []struct {
		name      string
		path      string
		accessKey string
		resource  *schema.Resource
	}{
		{"organization", "/orgs/test-org/actions/runner-groups/42", "selected_repository_ids", resourceGithubActionsRunnerGroup()},
		{"enterprise", "/enterprises/test-enterprise/actions/runner-groups/42", "selected_organization_ids", resourceGithubActionsEnterpriseRunnerGroup()},
	} {
		for _, tc := range []struct {
			name         string
			config       cty.Value
			oldID        string
			patchStatus  int
			accessStatus int
			wantID       string
			wantJSON     string
		}{
			{"detach", cty.StringVal(""), "network-1", 200, 204, "", "null"},
			{"omit", cty.NilVal, "network-1", 200, 204, "network-1", `"network-1"`},
			{"null", cty.NullVal(cty.String), "network-1", 200, 204, "network-1", `"network-1"`},
			{"unassigned", cty.NilVal, "", 200, 204, "", ""},
			{"forbidden", cty.StringVal(""), "network-1", 403, 204, "network-1", "null"},
			{"not found", cty.StringVal(""), "network-1", 404, 204, "network-1", "null"},
			{"invalid", cty.StringVal(""), "network-1", 422, 204, "network-1", "null"},
			{"server error", cty.StringVal(""), "network-1", 500, 204, "network-1", "null"},
			{"access failure after detach", cty.StringVal(""), "network-1", 200, 403, "", "null"},
		} {
			t.Run(scope.name+"/"+tc.name, func(t *testing.T) {
				d := schema.TestResourceDataRaw(t, scope.resource.Schema, map[string]any{
					"name":                       "test-group",
					"enterprise_slug":            "test-enterprise",
					"visibility":                 "selected",
					"allows_public_repositories": true,
					"restricted_to_workflows":    true,
					"selected_workflows":         []any{"test-org/repo/.github/workflows/build.yml@refs/heads/main"},
					scope.accessKey:              []any{7, 8},
					"network_configuration_id":   tc.oldID,
				})
				d.SetId("42")
				config := map[string]cty.Value{}
				if tc.config != cty.NilVal {
					config["network_configuration_id"] = tc.config
				}
				state := d.State()
				state.RawConfig = cty.ObjectVal(config)
				d = scope.resource.Data(state)

				var calls []string
				client, err := github.NewClient(github.WithTransport(localRoundTripper{handler: http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
					calls = append(calls, req.Method)
					w.Header().Set("Content-Type", "application/json")
					var body map[string]json.RawMessage
					if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
						t.Fatal(err)
					}
					switch req.Method {
					case http.MethodPatch:
						if req.URL.Path != scope.path {
							t.Errorf("PATCH path = %q, want %q", req.URL.Path, scope.path)
						}
						if got := string(body["network_configuration_id"]); got != tc.wantJSON {
							t.Errorf("network_configuration_id JSON = %q, want %q", got, tc.wantJSON)
						}
						for key, want := range map[string]string{
							"name":                       `"test-group"`,
							"visibility":                 `"selected"`,
							"allows_public_repositories": "true",
							"restricted_to_workflows":    "true",
							"selected_workflows":         `["test-org/repo/.github/workflows/build.yml@refs/heads/main"]`,
						} {
							if got := string(body[key]); got != want {
								t.Errorf("%s JSON = %s, want %s", key, got, want)
							}
						}
						if _, exists := body["runners"]; exists {
							t.Error("network update must not modify runner membership")
						}
						w.WriteHeader(tc.patchStatus)
						mustWrite(w, `{"id":42,"message":"request rejected"}`)
					case http.MethodPut:
						wantPath := scope.path + "/repositories"
						if scope.name == "enterprise" {
							wantPath = scope.path + "/organizations"
						}
						if req.URL.Path != wantPath || string(body[scope.accessKey]) != "[7,8]" {
							t.Errorf("access request changed: %s %s", req.URL.Path, body[scope.accessKey])
						}
						w.WriteHeader(tc.accessStatus)
						if tc.accessStatus != http.StatusNoContent {
							mustWrite(w, `{"message":"access update rejected"}`)
						}
					default:
						t.Errorf("unexpected request %s %s", req.Method, req.URL.Path)
						w.WriteHeader(http.StatusMethodNotAllowed)
					}
				})}))
				if err != nil {
					t.Fatal(err)
				}
				diags := scope.resource.UpdateContext(t.Context(), d, &Owner{name: "test-org", IsOrganization: true, v3client: client})
				wantError := tc.patchStatus != http.StatusOK || tc.accessStatus != http.StatusNoContent
				if diags.HasError() != wantError {
					t.Fatalf("diagnostics = %v, want error = %t", diags, wantError)
				}
				if wantError && !strings.Contains(fmt.Sprint(diags), "rejected") {
					t.Errorf("API error was lost: %v", diags)
				}
				wantCalls := []string{http.MethodPatch, http.MethodPut}
				if tc.patchStatus != http.StatusOK {
					wantCalls = wantCalls[:1]
				}
				if !reflect.DeepEqual(calls, wantCalls) {
					t.Errorf("requests = %v, want %v (no read-after-write)", calls, wantCalls)
				}
				if d.Id() != "42" || d.Get("network_configuration_id") != tc.wantID {
					t.Errorf("unexpected state: ID=%s network_configuration_id=%v", d.Id(), d.Get("network_configuration_id"))
				}
			})
		}
	}
}

func TestAccGithubActionsRunnerGroup(t *testing.T) {
	t.Parallel()

	t.Run("manages runner group network configuration", func(t *testing.T) {
		networkConfiguration := mustCreateTestOrganizationNetworkConfiguration(t)
		groupName := fmt.Sprintf("%srunner-group-%s", testResourcePrefix, acctest.RandString(testRandomIDLength))
		resourceName := "github_actions_runner_group.test"
		groupIDUnchanged := statecheck.CompareValue(compare.ValuesSame())

		configWithoutNetworking := fmt.Sprintf(`
			resource "github_actions_runner_group" "test" {
			  name       = %q
			  visibility = "all"
			}
		`, groupName)
		configWithNetworking := fmt.Sprintf(`
			resource "github_actions_runner_group" "test" {
			  name                     = %q
			  visibility               = "all"
			  network_configuration_id = %q
			}
		`, groupName, networkConfiguration.GetID())
		configWithNetworkingCleared := fmt.Sprintf(`
			resource "github_actions_runner_group" "test" {
			  name                     = %q
			  visibility               = "all"
			  network_configuration_id = ""
			}
		`, groupName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasPaidOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: configWithoutNetworking,
					ConfigStateChecks: []statecheck.StateCheck{
						groupIDUnchanged.AddStateValue(resourceName, tfjsonpath.New("id")),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("network_configuration_id"), knownvalue.StringExact("")),
					},
				},
				{
					Config: configWithNetworking,
					ConfigStateChecks: []statecheck.StateCheck{
						groupIDUnchanged.AddStateValue(resourceName, tfjsonpath.New("id")),
					},
				},
				{
					Config: configWithoutNetworking,
					ConfigStateChecks: []statecheck.StateCheck{
						groupIDUnchanged.AddStateValue(resourceName, tfjsonpath.New("id")),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("network_configuration_id"), knownvalue.StringExact(networkConfiguration.GetID())),
					},
				},
				{
					Config: configWithNetworkingCleared,
					ConfigStateChecks: []statecheck.StateCheck{
						groupIDUnchanged.AddStateValue(resourceName, tfjsonpath.New("id")),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("network_configuration_id"), knownvalue.StringExact("")),
					},
				},
			},
		})
	})

	t.Run("creates and imports a runner group with networking", func(t *testing.T) {
		networkConfiguration := mustCreateTestOrganizationNetworkConfiguration(t)
		groupName := fmt.Sprintf("%srunner-group-%s", testResourcePrefix, acctest.RandString(testRandomIDLength))
		resourceName := "github_actions_runner_group.test"

		config := fmt.Sprintf(`
			resource "github_actions_runner_group" "test" {
			  name                     = %q
			  visibility               = "all"
			  network_configuration_id = %q
			}
		`, groupName, networkConfiguration.GetID())

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasPaidOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("runners_url"), knownvalue.NotNull()),
					},
				},
				{
					ResourceName:      resourceName,
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})

	t.Run("creates runner groups without error", func(t *testing.T) {
		t.Parallel()

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-act-runner-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`
			resource "github_repository" "test" {
			  name = "%s"
			  vulnerability_alerts = false
			  auto_init = true
			}

			resource "github_repository_file" "workflow_file" {
			  repository          = github_repository.test.name
				branch              = "main"
			  file                = ".github/workflows/test.yml"
			  content             = ""
			  commit_message      = "Managed by Terraform"
			  commit_author       = "Terraform User"
			  commit_email        = "terraform@example.com"
			  overwrite_on_create = true
			}

			resource "github_actions_runner_group" "test" {
			  depends_on  = [github_repository_file.workflow_file]

			  name       = github_repository.test.name
			  visibility = "all"
			  restricted_to_workflows = true
			  selected_workflows = ["${github_repository.test.full_name}/.github/workflows/test.yml@refs/heads/main"]
			  allows_public_repositories = true
			}
		`, repoName)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet(
				"github_actions_runner_group.test", "name",
			),
			resource.TestCheckResourceAttr(
				"github_actions_runner_group.test", "name",
				repoName,
			),
			resource.TestCheckResourceAttr(
				"github_actions_runner_group.test", "visibility",
				"all",
			),
			resource.TestCheckResourceAttr(
				"github_actions_runner_group.test", "restricted_to_workflows",
				"true",
			),
			resource.TestCheckResourceAttr(
				"github_actions_runner_group.test", "selected_workflows.#",
				"1",
			),
			func(state *terraform.State) error {
				githubRepository := state.RootModule().Resources["github_repository.test"].Primary
				fullName := githubRepository.Attributes["full_name"]

				runnerGroup := state.RootModule().Resources["github_actions_runner_group.test"].Primary
				workflowActual := runnerGroup.Attributes["selected_workflows.0"]

				workflowExpected := fmt.Sprintf("%s/.github/workflows/test.yml@refs/heads/main", fullName)

				if workflowActual != workflowExpected {
					return fmt.Errorf("actual selected workflows %s not the same as expected selected workflows %s",
						workflowActual, workflowExpected)
				}
				return nil
			},
			resource.TestCheckResourceAttr(
				"github_actions_runner_group.test", "allows_public_repositories",
				"true",
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

	t.Run("manages runner visibility", func(t *testing.T) {
		t.Parallel()

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-act-runner-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`
			resource "github_repository" "test" {
			  name = "%s"
			}

			resource "github_actions_runner_group" "test" {
			  name       = github_repository.test.name
			  visibility = "selected"
			  selected_repository_ids = [github_repository.test.repo_id]
			}
		`, repoName)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet(
				"github_actions_runner_group.test", "name",
			),
			resource.TestCheckResourceAttr(
				"github_actions_runner_group.test", "name",
				repoName,
			),
			resource.TestCheckResourceAttr(
				"github_actions_runner_group.test", "visibility",
				"selected",
			),
			resource.TestCheckResourceAttr(
				"github_actions_runner_group.test", "selected_repository_ids.#",
				"1",
			),
			resource.TestCheckResourceAttrSet(
				"github_actions_runner_group.test", "selected_repositories_url",
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

	t.Run("imports an all runner group without error", func(t *testing.T) {
		t.Parallel()

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-act-runner-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`
			resource "github_repository" "test" {
			  name = "%s"
			}

			resource "github_actions_runner_group" "test" {
			  name       = github_repository.test.name
			  visibility = "all"
			}
    `, repoName)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet("github_actions_runner_group.test", "name"),
			resource.TestCheckResourceAttrSet("github_actions_runner_group.test", "visibility"),
			resource.TestCheckResourceAttr("github_actions_runner_group.test", "visibility", "all"),
			resource.TestCheckResourceAttr("github_actions_runner_group.test", "name", repoName),
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
					ResourceName:      "github_actions_runner_group.test",
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})

	t.Run("imports a private runner group without error", func(t *testing.T) {
		t.Parallel()

		// Note: this test is skipped because when setting visibility 'private', it always fails with:
		// Step 0 error: After applying this step, the plan was not empty:
		// visibility:                 "all" => "private"
		// Based on GitHub UI there is no way to create a private runner group
		t.Skip("This is not supported")

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-act-runner-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`
					resource "github_repository" "test" {
					  name = "%s"
					}

					resource "github_actions_runner_group" "test" {
					  name       = github_repository.test.name
					  visibility = "private"
					}
		    `, repoName)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet("github_actions_runner_group.test", "name"),
			resource.TestCheckResourceAttr("github_actions_runner_group.test", "name", repoName),
			resource.TestCheckResourceAttrSet("github_actions_runner_group.test", "visibility"),
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
					ResourceName:      "github_actions_runner_group.test",
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})

	t.Run("imports a selected runner group without error", func(t *testing.T) {
		t.Parallel()

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-act-runner-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name = "%s"
			}

			resource "github_actions_runner_group" "test" {
				name       = github_repository.test.name
				visibility = "selected"
				selected_repository_ids = [github_repository.test.repo_id]
			}
    `, repoName)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet("github_actions_runner_group.test", "name"),
			resource.TestCheckResourceAttr("github_actions_runner_group.test", "name", repoName),
			resource.TestCheckResourceAttrSet("github_actions_runner_group.test", "visibility"),
			resource.TestCheckResourceAttr("github_actions_runner_group.test", "visibility", "selected"),
			resource.TestCheckResourceAttr(
				"github_actions_runner_group.test", "selected_repository_ids.#",
				"1",
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
					ResourceName:      "github_actions_runner_group.test",
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})
}
