package github

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubRepositoryCollaborator(t *testing.T) {
	if len(testAccConf.testExternalUser) == 0 {
		t.Skip("No external user provided")
	}

	t.Run("creates invitations without error", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-collab-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name = "%s"
				auto_init = true
			}

			resource "github_repository_collaborator" "test_repo_collaborator" {
				repository = "${github_repository.test.name}"
				username   = "%s"
				permission = "triage"
			}
		`, repoName, testAccConf.testExternalUser)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_repository_collaborator.test_repo_collaborator", "permission",
				"triage",
			),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  check,
				},
			},
		})
	})

	t.Run("creates invitations when repository contains the org name", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-collab-%s", testResourcePrefix, randomID)
		configWithOwner := fmt.Sprintf(`
			resource "github_repository" "test" {
				name = "%s"
				auto_init = true
			}

			resource "github_repository_collaborator" "test_repo_collaborator_2" {
				repository = "%s/${github_repository.test.name}"
				username   = "%s"
				permission = "triage"
			}
		`, repoName, testAccConf.owner, testAccConf.testExternalUser)

		checkWithOwner := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_repository_collaborator.test_repo_collaborator_2", "permission",
				"triage",
			),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: configWithOwner,
					Check:  checkWithOwner,
				},
			},
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
	// Note: This test requires GH_TEST_COLLABORATOR to be set to a valid GitHub username and it won't work with `testExternalUser`
	testCollaborator := os.Getenv("GH_TEST_COLLABORATOR")
	if testCollaborator == "" {
		t.Skip("GH_TEST_COLLABORATOR not set, skipping archived repository collaborator test")
	}
	t.Run("can delete collaborators from archived repositories without error", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-collab-arch-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name = "%s"
				auto_init = true
			}

			resource "github_repository_collaborator" "test" {
				repository = github_repository.test.name
				username   = "%s"
				permission = "pull"
			}
		`, repoName, testCollaborator)

		archivedConfig := fmt.Sprintf(`
			resource "github_repository" "test" {
				name = "%s"
				auto_init = true
				archived = true
			}

			resource "github_repository_collaborator" "test" {
				repository = github_repository.test.name
				username   = "%s"
				permission = "pull"
			}
		`, repoName, testCollaborator)

		resource.Test(t, resource.TestCase{
			PreCheck:  func() { skipUnauthenticated(t) },
			Providers: testAccProviders,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr(
							"github_repository_collaborator.test", "username",
							testAccConf.testExternalUser,
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
								name = "%s"
								auto_init = true
								archived = true
							}
						`, repoName),
				},
			},
		})
	})
}
