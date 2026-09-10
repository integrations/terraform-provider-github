package github

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

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

			err := resourceGithubActionsRunnerGroupRead(d, meta)
			if (err != nil) != tc.wantError {
				t.Fatalf("read error = %v, want error = %t", err, tc.wantError)
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

func TestAccGithubActionsRunnerGroup(t *testing.T) {
	t.Parallel()

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
