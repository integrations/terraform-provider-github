package github

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestGithubRepositoryRulesets(t *testing.T) {
	t.Run("creates and updates repository rulesets without errors", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name = "tf-acc-test-%s"
				auto_init = true
				default_branch = "main"
				vulnerability_alerts = true
			}

			resource "github_repository_environment" "example" {
				environment  = "test"
				repository   = github_repository.test.name
			}

			resource "github_repository_ruleset" "test" {
				name        = "test"
				repository  = github_repository.test.id
				target      = "branch"
				enforcement = "active"

				conditions {
					ref_name {
						include = ["refs/heads/main"]
						exclude = []
					}
				}

				rules {
					creation = true

					update = true

					deletion                = true
					required_linear_history = true

					required_deployments {
						required_deployment_environments = ["test"]
					}

					required_signatures = false

					merge_queue {
						check_response_timeout_minutes    = 10
						grouping_strategy                 = "ALLGREEN"
						max_entries_to_build              = 5
						max_entries_to_merge              = 5
						merge_method                      = "MERGE"
						min_entries_to_merge              = 1
						min_entries_to_merge_wait_minutes = 60
					}

					pull_request {
						required_approving_review_count   = 2
						required_review_thread_resolution = true
						require_code_owner_review         = true
						dismiss_stale_reviews_on_push     = true
						require_last_push_approval        = true
					}

					required_status_checks {

						required_check {
							context = "ci"
						}

						strict_required_status_checks_policy = true
						do_not_enforce_on_create             = true
					}

					required_code_scanning {
					  required_code_scanning_tool {
						alerts_threshold = "errors"
						security_alerts_threshold = "high_or_higher"
						tool = "CodeQL"
					  }
					}

					non_fast_forward = true
				}
			}
		`, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_repository_ruleset.test",
				"name",
				"test",
			),
			resource.TestCheckResourceAttr(
				"github_repository_ruleset.test",
				"enforcement",
				"active",
			),
			resource.TestCheckResourceAttr(
				"github_repository_ruleset.test",
				"rules.0.required_code_scanning.0.required_code_scanning_tool.0.alerts_threshold",
				"errors",
			),
			resource.TestCheckResourceAttr(
				"github_repository_ruleset.test",
				"rules.0.required_code_scanning.0.required_code_scanning_tool.0.security_alerts_threshold",
				"high_or_higher",
			),
			resource.TestCheckResourceAttr(
				"github_repository_ruleset.test",
				"rules.0.required_code_scanning.0.required_code_scanning_tool.0.tool",
				"CodeQL",
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

	t.Run("creates and updates repository rulesets with enterprise features without errors", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name = "tf-acc-test-%s"
				auto_init = false
				vulnerability_alerts = true
			}

			resource "github_repository_environment" "example" {
				environment  = "test"
				repository   = github_repository.test.name
			}

			resource "github_repository_ruleset" "test" {
				name        = "test"
				repository  = github_repository.test.id
				target      = "branch"
				enforcement = "active"

				conditions {
					ref_name {
						include = ["~ALL"]
						exclude = []
					}
				}

				rules {
					branch_name_pattern {
						name     = "test"
						negate   = false
						operator = "starts_with"
						pattern  = "test"
					}
				}
			}
		`, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_repository_ruleset.test", "name",
				"test",
			),
			resource.TestCheckResourceAttr(
				"github_repository_ruleset.test", "enforcement",
				"active",
			),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessMode(t, enterprise) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  check,
				},
			},
		})
	})

	t.Run("updates a ruleset name without error", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf(`tf-acc-test-rename-%[1]s`, randomID)
		oldRSName := fmt.Sprintf(`ruleset-%[1]s`, randomID)
		newRSName := fmt.Sprintf(`%[1]s-renamed`, randomID)

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
			  name         = "%[1]s"
			  description  = "Terraform acceptance tests %[2]s"
			  vulnerability_alerts = true
			}

			resource "github_repository_ruleset" "test" {
				name        = "%[3]s"
				repository  = github_repository.test.id
				target      = "branch"
				enforcement = "active"

				rules {
					creation = true
				}
			}
		`, repoName, randomID, oldRSName)

		checks := map[string]resource.TestCheckFunc{
			"before": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_repository_ruleset.test", "name",
					oldRSName,
				),
			),
			"after": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_repository_ruleset.test", "name",
					newRSName,
				),
			),
		}

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  checks["before"],
				},
				{
					// Rename the ruleset to something else
					Config: strings.Replace(
						config,
						oldRSName,
						newRSName, 1),
					Check: checks["after"],
				},
			},
		})
	})

	t.Run("imports rulesets without error", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
			  name         = "tf-acc-test-import-%[1]s"
			  description  = "Terraform acceptance tests %[1]s"
			  auto_init    = true
			  default_branch = "main"
                          vulnerability_alerts = true
			}

			resource "github_repository_environment" "example" {
				environment  = "test"
				repository   = github_repository.test.name
			}

			resource "github_repository_ruleset" "test" {
				name        = "test"
				repository  = github_repository.test.id
				target      = "branch"
				enforcement = "active"

				conditions {
					ref_name {
						include = ["refs/heads/main"]
						exclude = []
					}
				}

				rules {
					creation = true

					update = true

					deletion                = true
					required_linear_history = true

					required_deployments {
						required_deployment_environments = ["test"]
					}

					required_signatures = false

					pull_request {
						required_approving_review_count   = 2
						required_review_thread_resolution = true
						require_code_owner_review         = true
						dismiss_stale_reviews_on_push     = true
						require_last_push_approval        = true
					}

					merge_queue {
						check_response_timeout_minutes    = 30
						grouping_strategy                 = "HEADGREEN"
						max_entries_to_build              = 4
						max_entries_to_merge              = 4
						merge_method                      = "SQUASH"
						min_entries_to_merge              = 2
						min_entries_to_merge_wait_minutes = 10
					}

					required_status_checks {

						required_check {
							context = "ci"
						}

						strict_required_status_checks_policy = true
						do_not_enforce_on_create             = true
					}

					non_fast_forward = true
				}
			}
		`, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet("github_repository_ruleset.test", "name"),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  check,
				},
				{
					ResourceName:      "github_repository_ruleset.test",
					ImportState:       true,
					ImportStateVerify: true,
					ImportStateIdFunc: importRepositoryRulesetByResourcePaths(
						"github_repository.test", "github_repository_ruleset.test"),
				},
			},
		})
	})

	t.Run("creates a push repository ruleset without errors", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

		config := fmt.Sprintf(`
			 resource "github_repository" "test" {
				 name                 = "tf-acc-test-%s"
				 auto_init            = false
				 visibility           = "internal"
				 vulnerability_alerts = true
			 }

			 resource "github_repository_ruleset" "test_push" {
				 name        = "test-push"
				 repository  = github_repository.test.id
				 target      = "push"
				 enforcement = "active"

				 rules {
					file_path_restriction {
					  restricted_file_paths = ["test.txt"]
					 }
					max_file_size {
					  max_file_size = 1048576
					}
					file_extension_restriction {
					   restricted_file_extensions = ["*.zip"]
					}
				 }
			 }

		`, randomID)
		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_repository_ruleset.test_push", "name",
				"test-push",
			),
			resource.TestCheckResourceAttr(
				"github_repository_ruleset.test_push", "target",
				"push",
			),
			resource.TestCheckResourceAttr(
				"github_repository_ruleset.test_push", "rules.0.file_path_restriction.0.restricted_file_paths.0",
				"test.txt",
			),
			resource.TestCheckResourceAttr(
				"github_repository_ruleset.test_push", "rules.0.max_file_size.0.max_file_size",
				"1048576",
			),
			resource.TestCheckResourceAttr(
				"github_repository_ruleset.test_push", "rules.0.file_extension_restriction.0.restricted_file_extensions.0",
				"*.zip",
			),
		)
		resource.Test(t, resource.TestCase{
			PreCheck:  func() { skipUnlessHasPaidOrgs(t) },
			Providers: testAccProviders,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  check,
				},
			},
		})
	})

	t.Run("creates repository ruleset with merge queue squash method", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name = "tf-acc-test-merge-queue-%s"
				auto_init = true
				default_branch = "main"
			}

			resource "github_repository_ruleset" "test" {
				name        = "merge-queue-test"
				repository  = github_repository.test.id
				target      = "branch"
				enforcement = "active"

				conditions {
					ref_name {
						include = ["refs/heads/main"]
						exclude = []
					}
				}

				rules {
					merge_queue {
						check_response_timeout_minutes    = 30
						grouping_strategy                 = "HEADGREEN"
						max_entries_to_build              = 4
						max_entries_to_merge              = 4
						merge_method                      = "SQUASH"
						min_entries_to_merge              = 2
						min_entries_to_merge_wait_minutes = 10
					}
				}
			}
		`, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_repository_ruleset.test", "name",
				"merge-queue-test",
			),
			resource.TestCheckResourceAttr(
				"github_repository_ruleset.test", "rules.0.merge_queue.0.merge_method",
				"SQUASH",
			),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:  func() { skipUnauthenticated(t) },
			Providers: testAccProviders,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  check,
				},
			},
		})
	})

	t.Run("removes bypass actors when removed from configuration", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name         = "tf-acc-test-bypass-%s"
				description  = "Terraform acceptance tests %[1]s"
				auto_init    = true
			}

			resource "github_team" "test" {
				name        = "tf-acc-test-team-%[1]s"
				description = "Terraform acc test team"
				privacy     = "closed"
			}

			resource "github_team_repository" "test" {
				team_id    = github_team.test.id
				repository = github_repository.test.name
				permission = "push"
			}

			resource "github_repository_ruleset" "test" {
				name        = "test-bypass"
				repository  = github_repository.test.id
				target      = "branch"
				enforcement = "active"

				bypass_actors {
					actor_id    = github_team_repository.test.team_id
					actor_type  = "Team"
					bypass_mode = "pull_request"
				}

				conditions {
					ref_name {
						include = ["~ALL"]
						exclude = []
					}
				}

				rules {
					pull_request {
						dismiss_stale_reviews_on_push     = false
						require_code_owner_review         = true
						require_last_push_approval        = false
						required_approving_review_count   = 1
						required_review_thread_resolution = false
					}
				}
			}
		`, randomID)

		configWithoutBypass := fmt.Sprintf(`
			resource "github_repository" "test" {
				name         = "tf-acc-test-bypass-%s"
				description  = "Terraform acceptance tests %[1]s"
				auto_init    = true
			}

			resource "github_team" "test" {
				name        = "tf-acc-test-team-%[1]s"
				description = "Terraform acc test team"
				privacy     = "closed"
			}

			resource "github_team_repository" "test" {
				team_id    = github_team.test.id
				repository = github_repository.test.name
				permission = "push"
			}

			resource "github_repository_ruleset" "test" {
				name        = "test-bypass"
				repository  = github_repository.test.id
				target      = "branch"
				enforcement = "active"

				conditions {
					ref_name {
						include = ["~ALL"]
						exclude = []
					}
				}

				rules {
					pull_request {
						dismiss_stale_reviews_on_push     = false
						require_code_owner_review         = true
						require_last_push_approval        = false
						required_approving_review_count   = 1
						required_review_thread_resolution = false
					}
				}
			}
		`, randomID)

		resource.Test(t, resource.TestCase{
			PreCheck:  func() { skipUnlessHasOrgs(t) },
			Providers: testAccProviders,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_repository_ruleset.test", "bypass_actors.#", "1"),
						resource.TestCheckResourceAttr("github_repository_ruleset.test", "bypass_actors.0.actor_type", "Team"),
					),
				},
				{
					Config: configWithoutBypass,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_repository_ruleset.test", "bypass_actors.#", "0"),
					),
				},
			},
		})
	})

	t.Run("updates ruleset without bypass actors defined", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name         = "tf-acc-test-no-bypass-%s"
				description  = "Terraform acceptance tests %[1]s"
				auto_init    = true
			}

			resource "github_repository_ruleset" "test" {
				name        = "test-no-bypass"
				repository  = github_repository.test.id
				target      = "branch"
				enforcement = "active"

				conditions {
					ref_name {
						include = ["~ALL"]
						exclude = []
					}
				}

				rules {
					deletion = true
				}
			}
		`, randomID)

		configUpdated := strings.Replace(
			config,
			"deletion = true",
			"deletion = false",
			1,
		)

		checks := map[string]resource.TestCheckFunc{
			"before": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_repository_ruleset.test", "rules.0.deletion",
					"true",
				),
				resource.TestCheckResourceAttr(
					"github_repository_ruleset.test", "bypass_actors.#",
					"0",
				),
			),
			"after": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_repository_ruleset.test", "rules.0.deletion",
					"false",
				),
				resource.TestCheckResourceAttr(
					"github_repository_ruleset.test", "bypass_actors.#",
					"0",
				),
			),
		}

		resource.Test(t, resource.TestCase{
			PreCheck:  func() { skipUnauthenticated(t) },
			Providers: testAccProviders,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  checks["before"],
				},
				{
					Config: configUpdated,
					Check:  checks["after"],
				},
			},
		})
	})

	t.Run("creates repository ruleset with all bypass_modes", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name         = "tf-acc-test-bypass-modes-%s"
				description  = "Terraform acceptance tests %[1]s"
				auto_init    = true
			}

			resource "github_team" "test_always" {
				name        = "tf-acc-test-team-always-%[1]s"
				description = "Terraform acc test team for always bypass"
				privacy     = "closed"
			}

			resource "github_team_repository" "test_always" {
				team_id    = github_team.test_always.id
				repository = github_repository.test.name
				permission = "push"
			}

			resource "github_team" "test_pull_request" {
				name        = "tf-acc-test-team-pr-%[1]s"
				description = "Terraform acc test team for pull_request bypass"
				privacy     = "closed"
				depends_on  = [github_team.test_always]
			}

			resource "github_team_repository" "test_pull_request" {
				team_id    = github_team.test_pull_request.id
				repository = github_repository.test.name
				permission = "push"
			}

			resource "github_team" "test_exempt" {
				name        = "tf-acc-test-team-exempt-%[1]s"
				description = "Terraform acc test team for exempt bypass"
				privacy     = "closed"
				depends_on  = [github_team.test_pull_request]
			}

			resource "github_team_repository" "test_exempt" {
				team_id    = github_team.test_exempt.id
				repository = github_repository.test.name
				permission = "push"
			}

			resource "github_repository_ruleset" "test" {
				name        = "test-bypass-modes"
				repository  = github_repository.test.id
				target      = "branch"
				enforcement = "active"

				bypass_actors {
					actor_id    = github_team_repository.test_always.team_id
					actor_type  = "Team"
					bypass_mode = "always"
				}

				bypass_actors {
					actor_id    = github_team_repository.test_pull_request.team_id
					actor_type  = "Team"
					bypass_mode = "pull_request"
				}

				bypass_actors {
					actor_id    = github_team_repository.test_exempt.team_id
					actor_type  = "Team"
					bypass_mode = "exempt"
				}

				conditions {
					ref_name {
						include = ["~ALL"]
						exclude = []
					}
				}

				rules {
					creation = true
				}
			}
		`, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_repository_ruleset.test", "bypass_actors.#",
				"3",
			),
			resource.TestCheckResourceAttrSet(
				"github_repository_ruleset.test", "bypass_actors.0.actor_id",
			),
			resource.TestCheckResourceAttr(
				"github_repository_ruleset.test", "bypass_actors.0.bypass_mode",
				"always",
			),
			resource.TestCheckResourceAttr(
				"github_repository_ruleset.test", "bypass_actors.0.actor_type",
				"Team",
			),
			resource.TestCheckResourceAttrSet(
				"github_repository_ruleset.test", "bypass_actors.1.actor_id",
			),
			resource.TestCheckResourceAttr(
				"github_repository_ruleset.test", "bypass_actors.1.bypass_mode",
				"pull_request",
			),
			resource.TestCheckResourceAttr(
				"github_repository_ruleset.test", "bypass_actors.1.actor_type",
				"Team",
			),
			resource.TestCheckResourceAttrSet(
				"github_repository_ruleset.test", "bypass_actors.2.actor_id",
			),
			resource.TestCheckResourceAttr(
				"github_repository_ruleset.test", "bypass_actors.2.bypass_mode",
				"exempt",
			),
			resource.TestCheckResourceAttr(
				"github_repository_ruleset.test", "bypass_actors.2.actor_type",
				"Team",
			),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:  func() { skipUnlessHasOrgs(t) },
			Providers: testAccProviders,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  check,
				},
			},
		})
	})

	t.Run("updates bypass_mode without error", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name         = "tf-acc-test-bypass-update-%s"
				description  = "Terraform acceptance tests %[1]s"
				auto_init    = true
			}

			resource "github_team" "test" {
				name        = "tf-acc-test-team-update-%[1]s"
				description = "Terraform acc test team"
				privacy     = "closed"
			}

			resource "github_team_repository" "test" {
				team_id    = github_team.test.id
				repository = github_repository.test.name
				permission = "push"
			}

			resource "github_repository_ruleset" "test" {
				name        = "test-bypass-update"
				repository  = github_repository.test.id
				target      = "branch"
				enforcement = "active"

				bypass_actors {
					actor_id    = github_team_repository.test.team_id
					actor_type  = "Team"
					bypass_mode = "always"
				}

				conditions {
					ref_name {
						include = ["~ALL"]
						exclude = []
					}
				}

				rules {
					creation = true
				}
			}
		`, randomID)

		configUpdated := strings.Replace(
			config,
			`bypass_mode = "always"`,
			`bypass_mode = "exempt"`,
			1,
		)

		checks := map[string]resource.TestCheckFunc{
			"before": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_repository_ruleset.test", "bypass_actors.0.bypass_mode",
					"always",
				),
			),
			"after": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_repository_ruleset.test", "bypass_actors.0.bypass_mode",
					"exempt",
				),
			),
		}

		resource.Test(t, resource.TestCase{
			PreCheck:  func() { skipUnlessHasOrgs(t) },
			Providers: testAccProviders,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  checks["before"],
				},
				{
					Config: configUpdated,
					Check:  checks["after"],
				},
			},
		})
	})

	t.Run("Creates repository ruleset with different actor types and bypass modes", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name         = "tf-acc-test-actor-types-%s"
				description  = "Terraform acceptance tests %[1]s"
				auto_init    = true
			}

			resource "github_team" "test" {
				name        = "tf-acc-test-team-actor-%[1]s"
				description = "Terraform acc test team"
				privacy     = "closed"
			}

			resource "github_team_repository" "test" {
				team_id    = github_team.test.id
				repository = github_repository.test.name
				permission = "push"
			}

			resource "github_repository_ruleset" "test" {
				name        = "test-actor-types"
				repository  = github_repository.test.id
				target      = "branch"
				enforcement = "active"

				bypass_actors {
					actor_id    = 0
					actor_type  = "OrganizationAdmin"
					bypass_mode = "exempt"
				}

				bypass_actors {
					actor_id    = 5
					actor_type  = "RepositoryRole"
					bypass_mode = "pull_request"
				}

				bypass_actors {
					actor_id    = github_team_repository.test.team_id
					actor_type  = "Team"
					bypass_mode = "always"
				}

				conditions {
					ref_name {
						include = ["~ALL"]
						exclude = []
					}
				}

				rules {
					creation = true
				}
			}
		`, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_repository_ruleset.test", "bypass_actors.#",
				"3",
			),
			resource.TestCheckResourceAttrSet(
				"github_repository_ruleset.test", "bypass_actors.2.actor_id",
			),
			resource.TestCheckResourceAttr(
				"github_repository_ruleset.test", "bypass_actors.2.actor_type",
				"Team",
			),
			resource.TestCheckResourceAttr(
				"github_repository_ruleset.test", "bypass_actors.2.bypass_mode",
				"always",
			),
			resource.TestCheckResourceAttr(
				"github_repository_ruleset.test", "bypass_actors.1.actor_id",
				"5",
			),
			resource.TestCheckResourceAttr(
				"github_repository_ruleset.test", "bypass_actors.1.actor_type",
				"RepositoryRole",
			),
			resource.TestCheckResourceAttr(
				"github_repository_ruleset.test", "bypass_actors.1.bypass_mode",
				"pull_request",
			),
			resource.TestCheckResourceAttr(
				"github_repository_ruleset.test", "bypass_actors.0.actor_id",
				"0",
			),
			resource.TestCheckResourceAttr(
				"github_repository_ruleset.test", "bypass_actors.0.actor_type",
				"OrganizationAdmin",
			),
			resource.TestCheckResourceAttr(
				"github_repository_ruleset.test", "bypass_actors.0.bypass_mode",
				"exempt",
			),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:  func() { skipUnlessHasOrgs(t) },
			Providers: testAccProviders,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  check,
				},
			},
		})
	})
}

func TestGithubRepositoryRulesetArchived(t *testing.T) {
	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

	t.Run("skips update and delete on archived repository", func(t *testing.T) {
		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name      = "tf-acc-test-archive-%s"
				auto_init = true
				archived  = false
			}

			resource "github_repository_ruleset" "test" {
				name        = "test"
				repository  = github_repository.test.name
				target      = "branch"
				enforcement = "active"
				rules { creation = true }
			}
		`, randomID)

		resource.Test(t, resource.TestCase{
			PreCheck:  func() { skipUnlessMode(t, individual) },
			Providers: testAccProviders,
			Steps: []resource.TestStep{
				{Config: config},
				{Config: strings.Replace(config, "archived  = false", "archived  = true", 1)},
				{Config: strings.Replace(strings.Replace(config, "archived  = false", "archived  = true", 1), `enforcement = "active"`, `enforcement = "disabled"`, 1)},
			},
		})
	})

	t.Run("prevents creating ruleset on archived repository", func(t *testing.T) {
		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name      = "tf-acc-test-archive-create-%s"
				auto_init = true
				archived  = true
			}
			resource "github_repository_ruleset" "test" {
				name       = "test"
				repository = github_repository.test.name
				target     = "branch"
				enforcement = "active"
				rules { creation = true }
			}
		`, randomID)

		resource.Test(t, resource.TestCase{
			PreCheck:  func() { skipUnlessMode(t, individual) },
			Providers: testAccProviders,
			Steps: []resource.TestStep{
				{Config: config, ExpectError: regexp.MustCompile("cannot create ruleset on archived repository")},
			},
		})
	})
}

func importRepositoryRulesetByResourcePaths(repoLogicalName, rulesetLogicalName string) resource.ImportStateIdFunc {
	// test importing using an ID of the form <repo-node-id>:<ruleset-id>
	return func(s *terraform.State) (string, error) {
		log.Printf("[DEBUG] Looking up tf state ")
		repo := s.RootModule().Resources[repoLogicalName]
		if repo == nil {
			return "", fmt.Errorf("Cannot find %s in terraform state", repoLogicalName)
		}
		repoID := repo.Primary.ID
		if repoID == "" {
			return "", fmt.Errorf("repository %s does not have an id in terraform state", repoLogicalName)
		}

		ruleset := s.RootModule().Resources[rulesetLogicalName]
		if ruleset == nil {
			return "", fmt.Errorf("Cannot find %s in terraform state", rulesetLogicalName)
		}
		rulesetID := ruleset.Primary.ID
		if rulesetID == "" {
			return "", fmt.Errorf("ruleset %s does not have an id in terraform state", rulesetLogicalName)
		}

		return fmt.Sprintf("%s:%s", repoID, rulesetID), nil
	}
}
