package github

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubRepositoryCollaborator(t *testing.T) {
	t.Skip("update <username> below to unskip this test run")

	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

	t.Run("creates invitations without error", func(t *testing.T) {
		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name = "tf-acc-test-%s"
				auto_init = true
			}

			resource "github_repository_collaborator" "test_repo_collaborator" {
				repository = "${github_repository.test.name}"
				username   = "<username>"
				permission = "triage"
			}
		`, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_repository_collaborator.test_repo_collaborator", "permission",
				"triage",
			),
		)

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: config,
						Check:  check,
					},
				},
			})
		}

		t.Run("with an anonymous account", func(t *testing.T) {
			t.Skip("anonymous account not supported for this operation")
		})

		t.Run("with an individual account", func(t *testing.T) {
			testCase(t, individual)
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})
	})

	t.Run("creates invitations when repository contains the org name", func(t *testing.T) {
		orgName := os.Getenv("GITHUB_ORGANIZATION")

		if orgName == "" {
			t.Skip("Set GITHUB_ORGANIZATION to unskip this test run")
		}

		configWithOwner := fmt.Sprintf(`
			resource "github_repository" "test" {
				name = "tf-acc-test-%s"
				auto_init = true
			}

			resource "github_repository_collaborator" "test_repo_collaborator_2" {
				repository = "%s/${github_repository.test.name}"
				username   = "<username>"
				permission = "triage"
			}
		`, randomID, orgName)

		checkWithOwner := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_repository_collaborator.test_repo_collaborator_2", "permission",
				"triage",
			),
		)

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: configWithOwner,
						Check:  checkWithOwner,
					},
				},
			})
		}

		t.Run("with an anonymous account", func(t *testing.T) {
			t.Skip("anonymous account not supported for this operation")
		})

		t.Run("with an individual account", func(t *testing.T) {
			testCase(t, individual)
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})
	})
}

func TestParseRepoName(t *testing.T) {
	tests := []struct {
		name         string
		repoName     string
		defaultOwner string
		wantOwner    string
		wantRepoName string
	}{
		{
			name:         "Repo name without owner",
			repoName:     "example-repo",
			defaultOwner: "default-owner",
			wantOwner:    "default-owner",
			wantRepoName: "example-repo",
		},
		{
			name:         "Repo name with owner",
			repoName:     "owner-name/example-repo",
			defaultOwner: "default-owner",
			wantOwner:    "owner-name",
			wantRepoName: "example-repo",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOwner, gotRepoName := parseRepoName(tt.repoName, tt.defaultOwner)
			if gotOwner != tt.wantOwner || gotRepoName != tt.wantRepoName {
				t.Errorf("parseRepoName(%q, %q) = %q, %q, want %q, %q",
					tt.repoName, tt.defaultOwner, gotOwner, gotRepoName, tt.wantOwner, tt.wantRepoName)
			}
		})
	}
}

func TestAccGithubRepositoryCollaboratorArchivedRepo(t *testing.T) {
	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

	t.Run("can delete collaborators from archived repositories without error", func(t *testing.T) {
		// Note: This test requires GITHUB_TEST_COLLABORATOR to be set to a valid GitHub username
		testCollaborator := os.Getenv("GITHUB_TEST_COLLABORATOR")
		if testCollaborator == "" {
			t.Skip("GITHUB_TEST_COLLABORATOR not set, skipping archived repository collaborator test")
		}

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name = "tf-acc-test-collab-archive-%s"
				auto_init = true
			}

			resource "github_repository_collaborator" "test" {
				repository = github_repository.test.name
				username   = "%s"
				permission = "pull"
			}
		`, randomID, testCollaborator)

		archivedConfig := fmt.Sprintf(`
			resource "github_repository" "test" {
				name = "tf-acc-test-collab-archive-%s"
				auto_init = true
				archived = true
			}

			resource "github_repository_collaborator" "test" {
				repository = github_repository.test.name
				username   = "%s"
				permission = "pull"
			}
		`, randomID, testCollaborator)

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: config,
						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckResourceAttr(
								"github_repository_collaborator.test", "username",
								testCollaborator,
							),
							resource.TestCheckResourceAttr(
								"github_repository_collaborator.test", "permission",
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
					// This step should succeed - the collaborator should be removed from state
					// without trying to actually delete it from the archived repo
					{
						Config: fmt.Sprintf(`
							resource "github_repository" "test" {
								name = "tf-acc-test-collab-archive-%s"
								auto_init = true
								archived = true
							}
						`, randomID),
					},
				},
			})
		}

		t.Run("with individual mode", func(t *testing.T) {
			testCase(t, individual)
		})

		t.Run("with organization mode", func(t *testing.T) {
			testCase(t, organization)
		})
	})
}
