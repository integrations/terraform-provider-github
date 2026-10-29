package github

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccGithubBranchProtectionV4(t *testing.T) {
	t.Parallel()
	skipUnauthenticated(t)

	t.Run("configures_default_settings_when_empty", func(t *testing.T) {
		t.Parallel()

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

		testRepoName := fmt.Sprintf("%sbranch-protection-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`

		resource "github_repository" "test" {
		  name      = "%s"
		  auto_init = true
		}

		resource "github_branch_protection" "test" {

		  repository_id = github_repository.test.node_id
		  pattern       = "main"

		}

	`, testRepoName)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("github_branch_protection.test", "pattern", "main"),
						resource.TestCheckResourceAttr("github_branch_protection.test", "require_signed_commits", "false"),
						resource.TestCheckResourceAttr("github_branch_protection.test", "require_conversation_resolution", "false"),
						resource.TestCheckResourceAttr("github_branch_protection.test", "required_status_checks.#", "0"),
						resource.TestCheckResourceAttr("github_branch_protection.test", "required_pull_request_reviews.#", "0"),
						resource.TestCheckResourceAttr("github_branch_protection.test", "restrict_pushes.#", "0"),
						resource.TestCheckResourceAttr("github_branch_protection.test", "lock_branch", "false"),
					),
				},
				{
					ResourceName:      "github_branch_protection.test",
					ImportState:       true,
					ImportStateVerify: true,
					ImportStateIdFunc: importBranchProtectionByRepoName(
						testRepoName, "main",
					),
				},
				{
					ResourceName: "github_branch_protection.test",
					ImportState:  true,
					ExpectError: regexp.MustCompile(
						`could not find a branch protection rule with the pattern 'no-such-pattern'`,
					),
					ImportStateIdFunc: importBranchProtectionByRepoName(
						testRepoName, "no-such-pattern",
					),
				},
			},
		})
	})

	t.Run("configures_default_settings_when_conversation_resolution_is_true", func(t *testing.T) {
		t.Parallel()

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		testRepoName := fmt.Sprintf("%sbranch-protection-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`

		resource "github_repository" "test" {
		  name      = "%s"
		  auto_init = true
		}

		resource "github_branch_protection" "test" {

		  repository_id = github_repository.test.node_id
		  pattern       = "main"

		  require_conversation_resolution = true
		}

	`, testRepoName)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("github_branch_protection.test", "pattern", "main"),
						resource.TestCheckResourceAttr("github_branch_protection.test", "require_signed_commits", "false"),
						resource.TestCheckResourceAttr("github_branch_protection.test", "require_conversation_resolution", "true"),
						resource.TestCheckResourceAttr("github_branch_protection.test", "required_linear_history", "false"),
						resource.TestCheckResourceAttr("github_branch_protection.test", "required_status_checks.#", "0"),
						resource.TestCheckResourceAttr("github_branch_protection.test", "required_pull_request_reviews.#", "0"),
						resource.TestCheckResourceAttr("github_branch_protection.test", "restrict_pushes.#", "0"),
						resource.TestCheckResourceAttr("github_branch_protection.test", "lock_branch", "false"),
					),
				},
				{
					ResourceName:      "github_branch_protection.test",
					ImportState:       true,
					ImportStateVerify: true,
					ImportStateIdFunc: importBranchProtectionByRepoName(
						testRepoName, "main",
					),
				},
				{
					ResourceName: "github_branch_protection.test",
					ImportState:  true,
					ExpectError: regexp.MustCompile(
						`could not find a branch protection rule with the pattern 'no-such-pattern'`,
					),
					ImportStateIdFunc: importBranchProtectionByRepoName(
						testRepoName, "no-such-pattern",
					),
				},
			},
		})
	})

	t.Run("configures_required_status_checks", func(t *testing.T) {
		t.Parallel()

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		testRepoName := fmt.Sprintf("%sbranch-protection-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`

			resource "github_repository" "test" {
			  name      = "%s"
			  auto_init = true
			}

			resource "github_branch_protection" "test" {

			  repository_id = github_repository.test.node_id
			  pattern       = "main"

			  required_status_checks {
			    strict   = true
			    contexts = ["github/foo"]
			  }

			}

	`, testRepoName)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  resource.TestCheckResourceAttr("github_branch_protection.test", "required_status_checks.#", "1"),
				},
				{
					ResourceName:      "github_branch_protection.test",
					ImportState:       true,
					ImportStateVerify: true,
					ImportStateIdFunc: importBranchProtectionByRepoID(
						"github_repository.test", "main",
					),
				},
				{
					ResourceName: "github_branch_protection.test",
					ImportState:  true,
					ExpectError: regexp.MustCompile(
						`could not find a branch protection rule with the pattern 'no-such-pattern'`,
					),
					ImportStateIdFunc: importBranchProtectionByRepoID(
						"github_repository.test", "no-such-pattern",
					),
				},
			},
		})
	})

	t.Run("configures_required_pull_request_reviews", func(t *testing.T) {
		t.Parallel()

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		testRepoName := fmt.Sprintf("%sbranch-protection-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`

			resource "github_repository" "test" {
			  name      = "%s"
			  auto_init = true
			}

			resource "github_branch_protection" "test" {

			  repository_id = github_repository.test.node_id
			  pattern       = "main"

			  required_pull_request_reviews {
				dismiss_stale_reviews      = true
				require_code_owner_reviews = true
				require_last_push_approval = true
			  }

			}

	`, testRepoName)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("github_branch_protection.test", "required_pull_request_reviews.#", "1"),
						resource.TestCheckResourceAttr("github_branch_protection.test", "required_pull_request_reviews.0.dismiss_stale_reviews", "true"),
						resource.TestCheckResourceAttr("github_branch_protection.test", "required_pull_request_reviews.0.require_code_owner_reviews", "true"),
						resource.TestCheckResourceAttr("github_branch_protection.test", "required_pull_request_reviews.0.required_approving_review_count", "1"),
						resource.TestCheckResourceAttr("github_branch_protection.test", "required_pull_request_reviews.0.require_last_push_approval", "true"),
					),
				},
			},
		})
	})

	t.Run("configures_branch_push_restrictions", func(t *testing.T) {
		t.Parallel()

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		testRepoName := fmt.Sprintf("%sbranch-protection-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`
			resource "github_repository" "test" {
			  name      = "%s"
			  auto_init = true
			}

			resource "github_branch_protection" "test" {

			  repository_id = github_repository.test.node_id
			  pattern       = "main"

			  restrict_pushes {}
			}
	`, testRepoName)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("github_branch_protection.test", "restrict_pushes.#", "1"),
						resource.TestCheckResourceAttr("github_branch_protection.test", "restrict_pushes.0.blocks_creations", "true"),
						resource.TestCheckResourceAttr("github_branch_protection.test", "restrict_pushes.0.push_allowances.#", "0"),
					),
				},
			},
		})
	})

	t.Run("configures_branch_push_restrictions_with_node_id", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)

		var username string
		if testAccConf.authMode == individual {
			username = testAccConf.owner
		} else {
			skipUnlessHasOrgUser1(t)
			username = testAccConf.testOrgUser1
			mustAddRepositoryCollaborator(t, repo, username)
		}

		user := mustGetUser(t, username)

		config := fmt.Sprintf(`
resource "github_branch_protection" "test" {
  repository_id = "%s"
  pattern       = "main"

  restrict_pushes {
    push_allowances = ["%s"]
  }
}
`, repo.GetNodeID(), user.GetNodeID())

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
				},
			},
		})
	})

	t.Run("configures_branch_push_restrictions_with_username", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)

		var username string
		if testAccConf.authMode == individual {
			username = testAccConf.owner
		} else {
			skipUnlessHasOrgUser1(t)
			username = testAccConf.testOrgUser1
			mustAddRepositoryCollaborator(t, repo, username)
		}

		config := fmt.Sprintf(`
resource "github_branch_protection" "test" {
  repository_id = "%s"
  pattern       = "main"

  restrict_pushes {
    push_allowances = ["/%s"]
  }
}
`, repo.GetNodeID(), username)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
				},
			},
		})
	})

	t.Run("configures_branch_push_restrictions_with_blocks_creations_false", func(t *testing.T) {
		t.Parallel()

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		testRepoName := fmt.Sprintf("%sbranch-protection-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`
			resource "github_repository" "test" {
			  name      = "%s"
			  auto_init = true
			}

			resource "github_branch_protection" "test" {

			  repository_id = github_repository.test.node_id
			  pattern       = "main"

			  restrict_pushes {
				blocks_creations = false
			  }
			}
	`, testRepoName)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("github_branch_protection.test", "restrict_pushes.#", "1"),
						resource.TestCheckResourceAttr("github_branch_protection.test", "restrict_pushes.0.blocks_creations", "false"),
						resource.TestCheckResourceAttr("github_branch_protection.test", "restrict_pushes.0.push_allowances.#", "0"),
					),
				},
			},
		})
	})

	t.Run("configures_force_pushes_and_deletions", func(t *testing.T) {
		t.Parallel()

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		testRepoName := fmt.Sprintf("%sbranch-protection-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`
			resource "github_repository" "test" {
			  name      = "%s"
			  auto_init = true
			}

			resource "github_branch_protection" "test" {

			  repository_id = github_repository.test.node_id
			  pattern       = "main"

			  allows_deletions    = true
			  allows_force_pushes = true

			}
	`, testRepoName)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("github_branch_protection.test", "allows_deletions", "true"),
						resource.TestCheckResourceAttr("github_branch_protection.test", "allows_force_pushes", "true"),
					),
				},
			},
		})
	})

	t.Run("configures_non_empty_list_of_force_push_bypassers", func(t *testing.T) {
		t.Parallel()

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		testRepoName := fmt.Sprintf("%sbranch-protection-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`

			resource "github_repository" "test" {
			  name      = "%s"
			  auto_init = true
			}

			data "github_user" "test" {
			  username = "%s"
			}

			resource "github_branch_protection" "test" {

			  repository_id = github_repository.test.node_id
			  pattern       = "main"

			  force_push_bypassers = [
				data.github_user.test.node_id
			  ]

			}

	`, testRepoName, testAccConf.testOrgUser1)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  resource.TestCheckResourceAttr("github_branch_protection.test", "force_push_bypassers.#", "1"),
				},
			},
		})
	})

	t.Run("configures_allow_force_push_with_a_team_as_bypasser", func(t *testing.T) {
		t.Parallel()

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		teamName := fmt.Sprintf("%sbranch-protection-%s", testResourcePrefix, randomID)
		repoName := teamName
		config := fmt.Sprintf(`

			resource "github_repository" "test" {
			  name      = "%s"
			  auto_init = true
			}

			resource "github_team" "test" {
				name = "%s"
			}

			resource "github_team_repository" "test" {
				team_id    = github_team.test.id
				repository = github_repository.test.name
				permission = "admin"
			}

			resource "github_branch_protection" "test" {
			  repository_id  = github_repository.test.node_id
			  pattern        = "main"

			  force_push_bypassers = [
				  "%s/${github_team.test.slug}"
	]
			}

	`, repoName, teamName, testAccConf.owner)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("github_branch_protection.test", "force_push_bypassers.#", "1"),
						resource.TestCheckResourceAttr("github_branch_protection.test", "allows_force_pushes", "false"),
					),
				},
			},
		})
	})

	t.Run("configures_empty_list_of_force_push_bypassers", func(t *testing.T) {
		t.Parallel()

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		testRepoName := fmt.Sprintf("%sbranch-protection-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`

			resource "github_repository" "test" {
			  name      = "%s"
			  auto_init = true
			}

			resource "github_branch_protection" "test" {

			  repository_id = github_repository.test.node_id
			  pattern       = "main"

			  force_push_bypassers = []

			}

	`, testRepoName)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  resource.TestCheckResourceAttr("github_branch_protection.test", "force_push_bypassers.#", "0"),
				},
			},
		})
	})

	t.Run("configures_non_empty_list_of_pull_request_bypassers", func(t *testing.T) {
		t.Parallel()

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		testRepoName := fmt.Sprintf("%sbranch-protection-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`

			resource "github_repository" "test" {
			  name      = "%s"
			  auto_init = true
			}

			resource "github_branch_protection" "test" {

			  repository_id = github_repository.test.node_id
			  pattern       = "main"

			  required_pull_request_reviews {
				pull_request_bypassers = [
					"/%s",
				]
			  }

			}

			`, testRepoName, testAccConf.testOrgUser1)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  resource.TestCheckResourceAttr("github_branch_protection.test", "required_pull_request_reviews.0.pull_request_bypassers.#", "1"),
				},
			},
		})
	})

	t.Run("configures_empty_list_of_pull_request_bypassers", func(t *testing.T) {
		t.Parallel()

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		testRepoName := fmt.Sprintf("%sbranch-protection-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`

			resource "github_repository" "test" {
			  name      = "%s"
			  auto_init = true
			}

			resource "github_branch_protection" "test" {

			  repository_id  = github_repository.test.node_id
			  pattern        = "main"

			  required_pull_request_reviews {
				pull_request_bypassers = []
			  }

			}

	`, testRepoName)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  resource.TestCheckResourceAttr("github_branch_protection.test", "required_pull_request_reviews.0.pull_request_bypassers.#", "0"),
				},
			},
		})
	})
}

func importBranchProtectionByRepoName(repo, pattern string) resource.ImportStateIdFunc {
	// test importing using an ID of the form <repo-name>:<branch-protection-pattern>
	return func(s *terraform.State) (string, error) {
		return fmt.Sprintf("%s:%s", repo, pattern), nil
	}
}

func importBranchProtectionByRepoID(repoLogicalName, pattern string) resource.ImportStateIdFunc {
	// test importing using an ID of the form <repo-node-id>:<branch-protection-pattern>
	// by retrieving the GraphQL ID from the terraform.State
	return func(s *terraform.State) (string, error) {
		repo := s.RootModule().Resources[repoLogicalName]
		if repo == nil {
			return "", fmt.Errorf("cannot find %s in terraform state", repoLogicalName)
		}
		repoID, found := repo.Primary.Attributes["node_id"]
		if !found {
			return "", fmt.Errorf("repository %s does not have a node_id in terraform state", repo.Primary.ID)
		}
		return fmt.Sprintf("%s:%s", repoID, pattern), nil
	}
}
