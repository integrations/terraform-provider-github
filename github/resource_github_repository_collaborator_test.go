package github

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccGithubRepositoryCollaborator(t *testing.T) {
	t.Parallel()

	if len(testAccConf.testExternalUser1) == 0 {
		t.Skip("No external user provided")
	}

	t.Run("creates invitations without error", func(t *testing.T) {
		t.Parallel()

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
		`, repoName, testAccConf.testExternalUser1)

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
		t.Parallel()

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
		`, repoName, testAccConf.owner, testAccConf.testExternalUser1)

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

	t.Run("can delete collaborators from archived repositories without error", func(t *testing.T) {
		t.Parallel()

		// Note: This test requires GH_TEST_COLLABORATOR to be set to a valid GitHub username and it won't work with `testExternalUser`
		testCollaborator := os.Getenv("GH_TEST_COLLABORATOR")
		if testCollaborator == "" {
			t.Skip("GH_TEST_COLLABORATOR not set, skipping archived repository collaborator test")
		}

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
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr(
							"github_repository_collaborator.test", "username",
							testAccConf.testExternalUser1,
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

func TestParseRepoName(t *testing.T) {
	t.Parallel()

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
			t.Parallel()

			gotOwner, gotRepoName := parseRepoName(tt.repoName, tt.defaultOwner)
			if gotOwner != tt.wantOwner || gotRepoName != tt.wantRepoName {
				t.Errorf("parseRepoName(%q, %q) = %q, %q, want %q, %q",
					tt.repoName, tt.defaultOwner, gotOwner, gotRepoName, tt.wantOwner, tt.wantRepoName)
			}
		})
	}
}

func TestGithubRepositoryCollaboratorExpiredInvitation(t *testing.T) {
	expiredInvitation := `[{"id": 7, "invitee": {"login": "user"}, "permissions": "write", "expired": true}]`

	t.Run("read removes collaborator with only an expired invitation from state", func(t *testing.T) {
		ts := githubApiMock([]*mockResponse{
			{
				ExpectedUri:    "/repos/owner/repo/invitations",
				ExpectedMethod: "GET",
				StatusCode:     200,
				ResponseBody:   expiredInvitation,
			},
			{
				ExpectedUri:    "/repos/owner/repo/collaborators",
				ExpectedMethod: "GET",
				StatusCode:     200,
				ResponseBody:   `[]`,
			},
		})
		defer ts.Close()

		meta := &Owner{name: "owner", v3client: mustCreateTestGitHubClient(t, ts.URL)}
		d := schema.TestResourceDataRaw(t, resourceGithubRepositoryCollaborator().Schema, map[string]any{
			"repository": "repo",
			"username":   "user",
		})
		d.SetId(buildTwoPartID("repo", "user"))

		if err := resourceGithubRepositoryCollaboratorRead(d, meta); err != nil {
			t.Fatalf("expected no error, got: %v", err)
		}
		if d.Id() != "" {
			t.Fatalf("expected collaborator to be removed from state, got id %q", d.Id())
		}
	})

	t.Run("create deletes an expired invitation before inviting again", func(t *testing.T) {
		ts := githubApiMock([]*mockResponse{
			{
				ExpectedUri:    "/repos/owner/repo/invitations",
				ExpectedMethod: "GET",
				StatusCode:     200,
				ResponseBody:   expiredInvitation,
			},
			{
				ExpectedUri:    "/repos/owner/repo/invitations/7",
				ExpectedMethod: "DELETE",
				StatusCode:     204,
			},
			{
				ExpectedUri:    "/repos/owner/repo/collaborators/user",
				ExpectedMethod: "PUT",
				StatusCode:     201,
				ResponseBody:   `{"id": 8, "invitee": {"login": "user"}, "permissions": "write"}`,
			},
			{
				ExpectedUri:    "/repos/owner/repo/invitations",
				ExpectedMethod: "GET",
				StatusCode:     200,
				ResponseBody:   `[{"id": 8, "invitee": {"login": "user"}, "permissions": "write", "expired": false}]`,
			},
		})
		defer ts.Close()

		meta := &Owner{name: "owner", v3client: mustCreateTestGitHubClient(t, ts.URL)}
		d := schema.TestResourceDataRaw(t, resourceGithubRepositoryCollaborator().Schema, map[string]any{
			"repository": "repo",
			"username":   "user",
		})

		if err := resourceGithubRepositoryCollaboratorCreate(d, meta); err != nil {
			t.Fatalf("expected no error, got: %v", err)
		}
		if got := d.Get("invitation_id").(string); got != "8" {
			t.Fatalf("expected new invitation id 8, got %q", got)
		}
	})
}
