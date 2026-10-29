package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubBranchProtectionV3_required_pull_request_reviews(t *testing.T) {
	t.Parallel()

	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
	teamName := fmt.Sprintf("%steam-%s", testResourcePrefix, randomID)
	repoName := fmt.Sprintf("%srepo-%s", testResourcePrefix, randomID)
	t.Run("configures_required_pull_request_reviews", func(t *testing.T) {
		t.Parallel()

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

			resource "github_branch_protection_v3" "test" {

				repository  = github_repository.test.name
				branch      = "main"

				required_pull_request_reviews {
					dismiss_stale_reviews      = true
					require_code_owner_reviews = true
					required_approving_review_count = 1
					require_last_push_approval = true
					dismissal_users = ["a"]
					dismissal_teams = ["b"]
					dismissal_apps = ["c"]
					bypass_pull_request_allowances {
						users = ["d"]
						teams = [github_team.test.slug]
						apps = ["e"]
					}
				}

				depends_on = [github_team_repository.test]
			}
	`, repoName, teamName)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(
							"github_branch_protection_v3.test", "required_pull_request_reviews.#", "1",
						),
						resource.TestCheckResourceAttr(
							"github_branch_protection_v3.test", "required_pull_request_reviews.0.dismiss_stale_reviews", "true",
						),
						resource.TestCheckResourceAttr(
							"github_branch_protection_v3.test", "required_pull_request_reviews.0.require_code_owner_reviews", "true",
						),
						resource.TestCheckResourceAttr(
							"github_branch_protection_v3.test", "required_pull_request_reviews.0.required_approving_review_count", "1",
						),
						resource.TestCheckResourceAttr(
							"github_branch_protection_v3.test", "required_pull_request_reviews.0.require_last_push_approval", "true",
						),
						resource.TestCheckResourceAttr(
							"github_branch_protection_v3.test", "required_pull_request_reviews.0.dismissal_users.#", "1",
						),
						resource.TestCheckResourceAttr(
							"github_branch_protection_v3.test", "required_pull_request_reviews.0.dismissal_teams.#", "1",
						),
						resource.TestCheckResourceAttr(
							"github_branch_protection_v3.test", "required_pull_request_reviews.0.dismissal_apps.#", "1",
						),
						resource.TestCheckResourceAttr(
							"github_branch_protection_v3.test", "required_pull_request_reviews.0.bypass_pull_request_allowances.#", "1",
						),
						resource.TestCheckResourceAttr(
							"github_branch_protection_v3.test", "required_pull_request_reviews.0.bypass_pull_request_allowances.0.users.#", "1",
						),
						resource.TestCheckResourceAttr(
							"github_branch_protection_v3.test", "required_pull_request_reviews.0.bypass_pull_request_allowances.0.teams.#", "1",
						),
						resource.TestCheckResourceAttr(
							"github_branch_protection_v3.test", "required_pull_request_reviews.0.bypass_pull_request_allowances.0.apps.#", "1",
						),
					),
				},
			},
		})
	})
}

func TestAccGithubBranchProtectionV3RequiredPullRequestReviewsBypassAllowances(t *testing.T) {
	t.Parallel()

	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
	teamName := fmt.Sprintf("%steam-%s", testResourcePrefix, randomID)
	repoName := fmt.Sprintf("%srepo-%s", testResourcePrefix, randomID)
	t.Run("configures_required_pull_request_reviews_with_bypass_allowances", func(t *testing.T) {
		t.Parallel()

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

			resource "github_branch_protection_v3" "test" {
				repository  = github_repository.test.name
				branch      = "main"

				required_pull_request_reviews {
					bypass_pull_request_allowances {
						teams = [github_team.test.slug]
					}
				}

				depends_on = [github_team_repository.test]
			}

	`, repoName, teamName)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(
							"github_branch_protection_v3.test", "required_pull_request_reviews.#", "1",
						),
						resource.TestCheckResourceAttr(
							"github_branch_protection_v3.test", "required_pull_request_reviews.0.bypass_pull_request_allowances.#", "1",
						),
					),
				},
			},
		})
	})
}

func TestAccGithubBranchProtectionV3_branch_push_restrictions(t *testing.T) {
	t.Parallel()

	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
	teamName := fmt.Sprintf("%steam-%s", testResourcePrefix, randomID)
	repoName := fmt.Sprintf("%srepo-%s", testResourcePrefix, randomID)
	t.Run("configures_branch_push_restrictions", func(t *testing.T) {
		t.Parallel()

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name      = "%s"
				auto_init = true
			}

			resource "github_team" "test" {
				name = "%s"
			}

			resource "github_team_repository" "test" {
				team_id    = "${github_team.test.id}"
				repository = "${github_repository.test.name}"
				permission = "pull"
			}

			resource "github_branch_protection_v3" "test" {

				repository   = github_repository.test.name
				branch       = "main"

				restrictions {
					teams = ["${github_team.test.slug}"]
				}

			}
			`, repoName, teamName)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  resource.TestCheckResourceAttr("github_branch_protection_v3.test", "restrictions.#", "1"),
				},
			},
		})
	})
}

func TestAccGithubBranchProtectionV3_computed_status_checks_no_churn(t *testing.T) {
	t.Parallel()
	skipUnlessHasOrgs(t)

	t.Run("handles_computed_status_checks_without_churn", func(t *testing.T) {
		t.Parallel()

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		testRepoName := fmt.Sprintf("%sbranch-protection-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name      = "%s"
				auto_init = true
			}

			resource "github_branch_protection_v3" "test" {
				repository = github_repository.test.name
				branch     = "main"

				required_status_checks {
					strict = true
					checks = [
						"ci/test",
						"ci/build"
					]
				}
			}
		`, testRepoName)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(
							"github_branch_protection_v3.test", "required_status_checks.#", "1",
						),
						resource.TestCheckResourceAttr(
							"github_branch_protection_v3.test", "required_status_checks.0.strict", "true",
						),
						resource.TestCheckResourceAttr(
							"github_branch_protection_v3.test", "required_status_checks.0.checks.#", "2",
						),
						resource.TestCheckTypeSetElemAttr(
							"github_branch_protection_v3.test", "required_status_checks.0.checks.*", "ci/test",
						),
						resource.TestCheckTypeSetElemAttr(
							"github_branch_protection_v3.test", "required_status_checks.0.checks.*", "ci/build",
						),
					),
				},
				// Re-apply the same config to test for churn
				{
					Config: config,
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("github_branch_protection_v3.test", "required_status_checks.#", "1"),
						resource.TestCheckResourceAttr("github_branch_protection_v3.test", "required_status_checks.0.strict", "true"),
						resource.TestCheckResourceAttr("github_branch_protection_v3.test", "required_status_checks.0.checks.#", "2"),
						resource.TestCheckTypeSetElemAttr("github_branch_protection_v3.test", "required_status_checks.0.checks.*", "ci/test"),
						resource.TestCheckTypeSetElemAttr("github_branch_protection_v3.test", "required_status_checks.0.checks.*", "ci/build"),
					),
					PlanOnly: true,
				},
			},
		})
	})
}

func TestAccGithubBranchProtectionV3_computed_status_contexts_no_churn(t *testing.T) {
	t.Parallel()
	skipUnlessHasOrgs(t)

	t.Run("handles_computed_status_contexts_without_churn", func(t *testing.T) {
		t.Parallel()

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		testRepoName := fmt.Sprintf("%sbranch-protection-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name      = "%s"
				auto_init = true
			}

			resource "github_branch_protection_v3" "test" {
				repository = github_repository.test.name
				branch     = "main"

				required_status_checks {
					strict = true
					contexts = [
						"ci/test",
						"ci/build"
					]
				}
			}
		`, testRepoName)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(
							"github_branch_protection_v3.test", "required_status_checks.#", "1",
						),
						resource.TestCheckResourceAttr(
							"github_branch_protection_v3.test", "required_status_checks.0.strict", "true",
						),
						resource.TestCheckResourceAttr(
							"github_branch_protection_v3.test", "required_status_checks.0.contexts.#", "2",
						),
						resource.TestCheckTypeSetElemAttr(
							"github_branch_protection_v3.test", "required_status_checks.0.contexts.*", "ci/test",
						),
						resource.TestCheckTypeSetElemAttr(
							"github_branch_protection_v3.test", "required_status_checks.0.contexts.*", "ci/build",
						),
					),
				},
				// Re-apply the same config to test for churn
				{
					Config: config,
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("github_branch_protection_v3.test", "required_status_checks.#", "1"),
						resource.TestCheckResourceAttr("github_branch_protection_v3.test", "required_status_checks.0.strict", "true"),
						resource.TestCheckResourceAttr("github_branch_protection_v3.test", "required_status_checks.0.contexts.#", "2"),
						resource.TestCheckTypeSetElemAttr("github_branch_protection_v3.test", "required_status_checks.0.contexts.*", "ci/test"),
						resource.TestCheckTypeSetElemAttr("github_branch_protection_v3.test", "required_status_checks.0.contexts.*", "ci/build"),
					),
					PlanOnly: true,
				},
			},
		})
	})
}

func TestAccGithubBranchProtectionV3_update_with_status_checks(t *testing.T) {
	t.Parallel()

	// A read populates both `contexts` and `checks` from the API response, so updating
	// any other setting used to send every required check twice, which GitHub rejects
	// with "Context must be unique per branch protection".
	for _, statusChecksField := range []string{"contexts", "checks"} {
		t.Run(fmt.Sprintf("updates_other_settings_when_%s_is_set", statusChecksField), func(t *testing.T) {
			t.Parallel()

			repo := mustCreateTestRepository(t)

			config := func(enforceAdmins bool) string {
				return fmt.Sprintf(`
				resource "github_branch_protection_v3" "test" {
					repository     = "%s"
					branch         = "main"
					enforce_admins = %t

					required_status_checks {
						strict = true
						%s = [
							"ci/test",
							"ci/build"
						]
					}
				}
			`, repo.GetName(), enforceAdmins, statusChecksField)
			}

			resource.Test(t, resource.TestCase{
				ProviderFactories: providerFactories,
				Steps: []resource.TestStep{
					{
						Config: config(false),
					},
					{
						Config: config(true),
						ConfigStateChecks: []statecheck.StateCheck{
							statecheck.ExpectKnownValue("github_branch_protection_v3.test", tfjsonpath.New("enforce_admins"), knownvalue.Bool(true)),
							statecheck.ExpectKnownValue("github_branch_protection_v3.test", tfjsonpath.New("required_status_checks").AtSliceIndex(0).AtMapKey(statusChecksField), knownvalue.SetSizeExact(2)),
						},
					},
				},
			})
		})
	}
}

func TestAccGithubBranchProtectionV3(t *testing.T) {
	t.Parallel()
	skipUnlessHasOrgs(t)

	t.Run("configures_default_settings_when_empty", func(t *testing.T) {
		t.Parallel()

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		testRepoName := fmt.Sprintf("%sbranch-protection-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`
		resource "github_repository" "test" {
			name      = "%s"
			auto_init = true
		}

		resource "github_branch_protection_v3" "test" {

			repository  = github_repository.test.name
			branch      = "main"

		}
		`, testRepoName)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("github_branch_protection_v3.test", "branch", "main"),
						resource.TestCheckResourceAttr("github_branch_protection_v3.test", "require_signed_commits", "false"),
						resource.TestCheckResourceAttr("github_branch_protection_v3.test", "require_conversation_resolution", "false"),
						resource.TestCheckResourceAttr("github_branch_protection_v3.test", "required_status_checks.#", "0"),
						resource.TestCheckResourceAttr("github_branch_protection_v3.test", "required_pull_request_reviews.#", "0"),
						resource.TestCheckResourceAttr("github_branch_protection_v3.test", "push_restrictions.#", "0"),
					),
				},
			},
		})
	})

	t.Run("configures_conversation_resolution", func(t *testing.T) {
		t.Parallel()

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		testRepoName := fmt.Sprintf("%sbranch-protection-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`
		resource "github_repository" "test" {
			name      = "%s"
			auto_init = true
		}

		resource "github_branch_protection_v3" "test" {

			repository  = github_repository.test.name
			branch      = "main"

			require_conversation_resolution = true
		}
		`, testRepoName)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("github_branch_protection_v3.test", "branch", "main"),
						resource.TestCheckResourceAttr("github_branch_protection_v3.test", "require_signed_commits", "false"),
						resource.TestCheckResourceAttr("github_branch_protection_v3.test", "require_conversation_resolution", "true"),
						resource.TestCheckResourceAttr("github_branch_protection_v3.test", "required_status_checks.#", "0"),
						resource.TestCheckResourceAttr("github_branch_protection_v3.test", "required_pull_request_reviews.#", "0"),
						resource.TestCheckResourceAttr("github_branch_protection_v3.test", "push_restrictions.#", "0"),
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

		resource "github_branch_protection_v3" "test" {
			repository  = github_repository.test.name
			branch      = "main"

			required_status_checks {
				strict   = true
				checks = [
					"github/foo",
					"github/bar:-1",
					"github:foo:baz:1",
				]
			}

		}
	`, testRepoName)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttrSet("github_branch_protection_v3.test", "required_status_checks.#"),
						resource.TestCheckResourceAttr("github_branch_protection_v3.test", "required_status_checks.#", "1"),
						resource.TestCheckResourceAttr("github_branch_protection_v3.test", "required_status_checks.0.strict", "true"),
						resource.TestCheckResourceAttrSet("github_branch_protection_v3.test", "required_status_checks.0.checks.#"),
						resource.TestCheckResourceAttr("github_branch_protection_v3.test", "required_status_checks.0.checks.#", "3"),
					),
				},
			},
		})
	})

	t.Run("configures_required_status_checks_context", func(t *testing.T) {
		t.Parallel()

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		testRepoName := fmt.Sprintf("%sbranch-protection-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`
		resource "github_repository" "test" {
			name      = "%s"
			auto_init = true
		}

		resource "github_branch_protection_v3" "test" {
			repository  = github_repository.test.name
			branch      = "main"

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
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttrSet("github_branch_protection_v3.test", "required_status_checks.#"),
						resource.TestCheckResourceAttr("github_branch_protection_v3.test", "required_status_checks.#", "1"),
						resource.TestCheckResourceAttr("github_branch_protection_v3.test", "required_status_checks.0.strict", "true"),
						resource.TestCheckResourceAttrSet("github_branch_protection_v3.test", "required_status_checks.0.checks.#"),
						resource.TestCheckResourceAttr("github_branch_protection_v3.test", "required_status_checks.0.checks.#", "1"),
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

			resource "github_team" "test" {
				name = "%[1]s"
			}

			resource "github_team_repository" "test" {
				team_id    = github_team.test.id
				repository = github_repository.test.name
				permission = "admin"
			}

			resource "github_branch_protection_v3" "test" {

				repository  = github_repository.test.name
				branch      = "main"

				required_pull_request_reviews {
					dismiss_stale_reviews      = true
					require_code_owner_reviews = true
					required_approving_review_count = 1
					require_last_push_approval = true
					dismissal_users = ["a"]
					dismissal_teams = ["b"]
					dismissal_apps = ["c"]
					bypass_pull_request_allowances {
						users = ["d"]
						teams = [github_team.test.slug]
						apps = ["e"]
					}
				}

				depends_on = [github_team_repository.test]
			}
	`, testRepoName)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(
							"github_branch_protection_v3.test", "required_pull_request_reviews.#", "1",
						),
						resource.TestCheckResourceAttr(
							"github_branch_protection_v3.test", "required_pull_request_reviews.0.dismiss_stale_reviews", "true",
						),
						resource.TestCheckResourceAttr(
							"github_branch_protection_v3.test", "required_pull_request_reviews.0.require_code_owner_reviews", "true",
						),
						resource.TestCheckResourceAttr(
							"github_branch_protection_v3.test", "required_pull_request_reviews.0.required_approving_review_count", "1",
						),
						resource.TestCheckResourceAttr(
							"github_branch_protection_v3.test", "required_pull_request_reviews.0.require_last_push_approval", "true",
						),
						resource.TestCheckResourceAttr(
							"github_branch_protection_v3.test", "required_pull_request_reviews.0.dismissal_users.#", "1",
						),
						resource.TestCheckResourceAttr(
							"github_branch_protection_v3.test", "required_pull_request_reviews.0.dismissal_teams.#", "1",
						),
						resource.TestCheckResourceAttr(
							"github_branch_protection_v3.test", "required_pull_request_reviews.0.dismissal_apps.#", "1",
						),
						resource.TestCheckResourceAttr(
							"github_branch_protection_v3.test", "required_pull_request_reviews.0.bypass_pull_request_allowances.#", "1",
						),
						resource.TestCheckResourceAttr(
							"github_branch_protection_v3.test", "required_pull_request_reviews.0.bypass_pull_request_allowances.0.users.#", "1",
						),
						resource.TestCheckResourceAttr(
							"github_branch_protection_v3.test", "required_pull_request_reviews.0.bypass_pull_request_allowances.0.teams.#", "1",
						),
						resource.TestCheckResourceAttr(
							"github_branch_protection_v3.test", "required_pull_request_reviews.0.bypass_pull_request_allowances.0.apps.#", "1",
						),
					),
				},
			},
		})
	})

	t.Run("configures_required_pull_request_reviews_with_bypass_allowances", func(t *testing.T) {
		t.Parallel()

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		testResourceName := fmt.Sprintf("%sbranch-protection-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`

			resource "github_repository" "test" {
				name      = "%s"
				auto_init = true
			}

			resource "github_team" "test" {
				name = "%[1]s"
			}

			resource "github_team_repository" "test" {
				team_id    = github_team.test.id
				repository = github_repository.test.name
				permission = "admin"
			}

			resource "github_branch_protection_v3" "test" {
				repository  = github_repository.test.name
				branch      = "main"

				required_pull_request_reviews {
					bypass_pull_request_allowances {
						teams = [github_team.test.slug]
					}
				}

				depends_on = [github_team_repository.test]
			}

	`, testResourceName)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr(
							"github_branch_protection_v3.test", "required_pull_request_reviews.#", "1",
						),
						resource.TestCheckResourceAttr(
							"github_branch_protection_v3.test", "required_pull_request_reviews.0.bypass_pull_request_allowances.#", "1",
						),
					),
				},
			},
		})
	})

	t.Run("configures_branch_push_restrictions", func(t *testing.T) {
		t.Parallel()

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		testResourceName := fmt.Sprintf("%sbranch-protection-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name      = "%s"
				auto_init = true
			}

			resource "github_team" "test" {
				name = "%[1]s"
			}

			resource "github_team_repository" "test" {
				team_id    = "${github_team.test.id}"
				repository = "${github_repository.test.name}"
				permission = "pull"
			}

			resource "github_branch_protection_v3" "test" {

				repository   = github_repository.test.name
				branch       = "main"

				restrictions {
					teams = ["${github_team.test.slug}"]
				}

			}
			`, testResourceName)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  resource.TestCheckResourceAttr("github_branch_protection_v3.test", "restrictions.#", "1"),
				},
			},
		})
	})
}
