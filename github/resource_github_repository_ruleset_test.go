package github

import (
	"fmt"
	"log"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccGithubRepositoryRuleset(t *testing.T) {
	baseRepoVisibility := "public"

	if testAccConf.authMode == enterprise {
		// This enables repos to be created even in GHEC EMU
		baseRepoVisibility = "private"
	}

	t.Run("create_branch_ruleset", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-ruleset-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
resource "github_repository" "test" {
	name = "%s"
	auto_init = true
	vulnerability_alerts = true
	visibility = "%s"
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

	bypass_actors {
		actor_type = "DeployKey"
		bypass_mode = "always"
	}

	bypass_actors {
		actor_id    = 5
		actor_type  = "RepositoryRole"
		bypass_mode = "always"
	}

	conditions {
		ref_name {
			include = ["refs/heads/main"]
			exclude = []
		}
	}

	rules {
		creation = true

		update = true

		deletion = true

		required_linear_history = true

		merge_queue {
			check_response_timeout_minutes    = 10
			grouping_strategy                 = "ALLGREEN"
			max_entries_to_build              = 5
			max_entries_to_merge              = 5
			merge_method                      = "SQUASH"
			min_entries_to_merge              = 1
			min_entries_to_merge_wait_minutes = 60
		}

		copilot_code_review {
			review_on_push             = true
			review_draft_pull_requests = false
		}

		required_deployments {
			required_deployment_environments = [github_repository_environment.example.environment]
		}

		required_signatures = false

		pull_request {
			allowed_merge_methods             = ["merge", "squash"]
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
`, repoName, baseRepoVisibility)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_repository_ruleset.test", "name", "test"),
						resource.TestCheckResourceAttr("github_repository_ruleset.test", "target", "branch"),
						resource.TestCheckResourceAttr("github_repository_ruleset.test", "enforcement", "active"),
						resource.TestCheckResourceAttr("github_repository_ruleset.test", "bypass_actors.#", "2"),
						resource.TestCheckResourceAttr("github_repository_ruleset.test", "bypass_actors.0.actor_type", "DeployKey"),
						resource.TestCheckResourceAttr("github_repository_ruleset.test", "bypass_actors.0.bypass_mode", "always"),
						resource.TestCheckResourceAttr("github_repository_ruleset.test", "bypass_actors.1.actor_id", "5"),
						resource.TestCheckResourceAttr("github_repository_ruleset.test", "bypass_actors.1.actor_type", "RepositoryRole"),
						resource.TestCheckResourceAttr("github_repository_ruleset.test", "bypass_actors.1.bypass_mode", "always"),
						resource.TestCheckResourceAttr("github_repository_ruleset.test", "rules.0.pull_request.0.allowed_merge_methods.#", "2"),
						resource.TestCheckResourceAttr("github_repository_ruleset.test", "rules.0.required_code_scanning.0.required_code_scanning_tool.0.alerts_threshold", "errors"),
						resource.TestCheckResourceAttr("github_repository_ruleset.test", "rules.0.required_code_scanning.0.required_code_scanning_tool.0.security_alerts_threshold", "high_or_higher"),
						resource.TestCheckResourceAttr("github_repository_ruleset.test", "rules.0.required_code_scanning.0.required_code_scanning_tool.0.tool", "CodeQL"),
						resource.TestCheckResourceAttr("github_repository_ruleset.test", "rules.0.copilot_code_review.0.review_on_push", "true"),
						resource.TestCheckResourceAttr("github_repository_ruleset.test", "rules.0.copilot_code_review.0.review_draft_pull_requests", "false"),
					),
				},
			},
		})
	})

	t.Run("create_branch_ruleset_with_enterprise_features", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-ruleset-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
	resource "github_repository" "test" {
		name = "%s"
		auto_init = false
		vulnerability_alerts = true
		visibility = "%s"
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
`, repoName, baseRepoVisibility)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_repository_ruleset.test", "name", "test"),
						resource.TestCheckResourceAttr("github_repository_ruleset.test", "enforcement", "active"),
					),
				},
			},
		})
	})

	t.Run("creates_push_ruleset", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-ruleset-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
resource "github_repository" "test" {
	name                 = "%s"
	auto_init            = false
	visibility           = "internal"
	vulnerability_alerts = true
}

resource "github_repository_ruleset" "test" {
	name        = "test-push"
	repository  = github_repository.test.id
	target      = "push"
	enforcement = "active"

	bypass_actors {
		actor_type = "DeployKey"
		bypass_mode = "always"
	}

	bypass_actors {
		actor_id    = 5
		actor_type  = "RepositoryRole"
		bypass_mode = "always"
	}

	rules {
		file_path_restriction {
			restricted_file_paths = ["test.txt"]
		}

		max_file_size {
			max_file_size = 99
		}

		file_extension_restriction {
			restricted_file_extensions = ["*.zip"]
		}
	}
}

`, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasPaidOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_repository_ruleset.test", "name", "test-push"),
						resource.TestCheckResourceAttr("github_repository_ruleset.test", "target", "push"),
						resource.TestCheckResourceAttr("github_repository_ruleset.test", "enforcement", "active"),
						resource.TestCheckResourceAttr("github_repository_ruleset.test", "bypass_actors.#", "2"),
						resource.TestCheckResourceAttr("github_repository_ruleset.test", "bypass_actors.0.actor_type", "DeployKey"),
						resource.TestCheckResourceAttr("github_repository_ruleset.test", "bypass_actors.0.bypass_mode", "always"),
						resource.TestCheckResourceAttr("github_repository_ruleset.test", "bypass_actors.1.actor_id", "5"),
						resource.TestCheckResourceAttr("github_repository_ruleset.test", "bypass_actors.1.actor_type", "RepositoryRole"),
						resource.TestCheckResourceAttr("github_repository_ruleset.test", "bypass_actors.1.bypass_mode", "always"),
						resource.TestCheckResourceAttr("github_repository_ruleset.test", "rules.0.file_path_restriction.0.restricted_file_paths.0", "test.txt"),
						resource.TestCheckResourceAttr("github_repository_ruleset.test", "rules.0.max_file_size.0.max_file_size", "99"),
						resource.TestCheckResourceAttr("github_repository_ruleset.test", "rules.0.file_extension_restriction.0.restricted_file_extensions.0", "*.zip"),
					),
				},
			},
		})
	})

	t.Run("update_ruleset_name", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-ruleset-rename-%s", testResourcePrefix, randomID)
		name := fmt.Sprintf("ruleset-%s", randomID)
		nameUpdated := fmt.Sprintf("%s-renamed", name)

		config := `
resource "github_repository" "test" {
	name         = "%s"
	description  = "Terraform acceptance tests %s"
	vulnerability_alerts = true
	visibility = "%s"
}

resource "github_repository_ruleset" "test" {
	name        = "%s"
	repository  = github_repository.test.id
	target      = "branch"
	enforcement = "active"

	rules {
		creation = true
	}
}
`

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, repoName, randomID, baseRepoVisibility, name),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_repository_ruleset.test", "name", name),
					),
				},
				{
					Config: fmt.Sprintf(config, repoName, randomID, baseRepoVisibility, nameUpdated),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_repository_ruleset.test", "name", nameUpdated),
					),
				},
			},
		})
	})

	t.Run("update_clear_bypass_actors", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-ruleset-bypass-%s", testResourcePrefix, randomID)

		bypassActorsConfig := `
bypass_actors {
  actor_type = "DeployKey"
  bypass_mode = "always"
}

bypass_actors {
  actor_id    = 5
  actor_type  = "RepositoryRole"
  bypass_mode = "always"
}
`
		baseConfig := `
resource "github_repository" "test" {
	name         = "%s"
	description  = "Terraform acceptance tests %[1]s"
	auto_init    = true
	visibility = "%s"
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

	%s

	rules {
		creation = true
	}
}
`
		config := fmt.Sprintf(baseConfig, repoName, baseRepoVisibility, bypassActorsConfig)

		configUpdated := fmt.Sprintf(baseConfig, repoName, baseRepoVisibility, "")
		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_repository_ruleset.test", "bypass_actors.#", "2"),
					),
				},
				{
					Config: configUpdated,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_repository_ruleset.test", "bypass_actors.#", "0"),
					),
				},
			},
		})
	})

	t.Run("update_bypass_mode", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-ruleset-bypass-update-%s", testResourcePrefix, randomID)

		bypassMode := "always"
		bypassModeUpdated := "exempt"

		config := `
resource "github_repository" "test" {
	name         = "%s"
	description  = "Terraform acceptance tests %s"
	auto_init    = true
	visibility = "%s"
}

resource "github_repository_ruleset" "test" {
	name        = "test-bypass-update"
	repository  = github_repository.test.id
	target      = "branch"
	enforcement = "active"

	bypass_actors {
		actor_id    = 5
		actor_type  = "RepositoryRole"
		bypass_mode = "%s"
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
`

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, repoName, randomID, baseRepoVisibility, bypassMode),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_repository_ruleset.test", "bypass_actors.0.bypass_mode", bypassMode),
					),
				},
				{
					Config: fmt.Sprintf(config, repoName, randomID, baseRepoVisibility, bypassModeUpdated),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_repository_ruleset.test", "bypass_actors.0.bypass_mode", bypassModeUpdated),
					),
				},
			},
		})
	})

	t.Run("import", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-ruleset-bpmod-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
        name         = "%s"
			  description  = "Terraform acceptance tests %s"
			  auto_init    = true
			  default_branch = "main"
				vulnerability_alerts = true
				visibility = "%s"
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
				}
			}
		`, repoName, randomID, baseRepoVisibility)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
				},
				{
					ResourceName:            "github_repository_ruleset.test",
					ImportState:             true,
					ImportStateVerify:       true,
					ImportStateIdFunc:       importRepositoryRulesetByResourcePaths("github_repository.test", "github_repository_ruleset.test"),
					ImportStateVerifyIgnore: []string{"etag"},
				},
			},
		})
	})
}

func TestAccGithubRepositoryRulesetArchived(t *testing.T) {
	baseRepoVisibility := "public"

	if testAccConf.authMode == enterprise {
		// This enables repos to be created even in GHEC EMU
		baseRepoVisibility = "private"
	}

	t.Run("skips update and delete on archived repository", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-ruleset-arch-%s", testResourcePrefix, randomID)
		archivedBefore := false
		archivedAfter := true
		enforcementBefore := "active"
		enforcementAfter := "disabled"
		config := `
			resource "github_repository" "test" {
				name      = "%s"
				auto_init = true
				archived  = %t
				visibility = "%s"
			}

			resource "github_repository_ruleset" "test" {
				name        = "test"
				repository  = github_repository.test.name
				target      = "branch"
				enforcement = "%s"
				rules { creation = true }
			}
		`

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{Config: fmt.Sprintf(config, repoName, archivedBefore, baseRepoVisibility, enforcementBefore)},
				{Config: fmt.Sprintf(config, repoName, archivedAfter, baseRepoVisibility, enforcementBefore)},
				{Config: fmt.Sprintf(config, repoName, archivedAfter, baseRepoVisibility, enforcementAfter)},
			},
		})
	})

	t.Run("prevents creating ruleset on archived repository", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-ruleset-arch-cr-%s", testResourcePrefix, randomID)
		repoConfig := `
	resource "github_repository" "test" {
		name      = "%s"
		auto_init = true
		archived  = %t
		visibility = "%s"
	}
	%s
`
		rulesetConfig := `
resource "github_repository_ruleset" "test" {
	name       = "test"
	repository = github_repository.test.name
	target     = "branch"
	enforcement = "active"
	rules { creation = true }
}
`

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(repoConfig, repoName, false, baseRepoVisibility, ""),
				},
				{
					Config:      fmt.Sprintf(repoConfig, repoName, true, baseRepoVisibility, rulesetConfig),
					ExpectError: regexp.MustCompile("cannot create ruleset on archived repository"),
				},
			},
		})
	})
}

func TestAccGithubRepositoryRulesetValidation(t *testing.T) {
	t.Run("Validates push target rejects ref_name condition", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name         = "tf-acc-test-push-ref-%s"
				auto_init    = true
				visibility   = "private"
				vulnerability_alerts = true
			}

			resource "github_repository_ruleset" "test" {
				name        = "test-push-with-ref"
				repository  = github_repository.test.id
				target      = "push"
				enforcement = "active"

				conditions {
					ref_name {
						include = ["~ALL"]
						exclude = []
					}
				}

				rules {
					max_file_size {
						max_file_size = 100
					}
				}
			}
		`, randomID)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasPaidOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      config,
					ExpectError: regexp.MustCompile("ref_name must not be set for push target"),
				},
			},
		})
	})

	t.Run("Validates push target rejects branch/tag rules", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name         = "tf-acc-test-push-rules-%s"
				auto_init    = true
				visibility   = "private"
				vulnerability_alerts = true
			}

			resource "github_repository_ruleset" "test" {
				name        = "test-push-branch-rule"
				repository  = github_repository.test.id
				target      = "push"
				enforcement = "active"

				rules {
					# 'creation' is a branch/tag rule, not valid for push target
					creation = true
				}
			}
		`, randomID)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasPaidOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      config,
					ExpectError: regexp.MustCompile("rule .* is not valid for push target"),
				},
			},
		})
	})

	t.Run("Validates branch target rejects push-only rules", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name         = "tf-acc-test-branch-push-%s"
				auto_init    = true
				vulnerability_alerts = true

				visibility = "private"
			}

			resource "github_repository_ruleset" "test" {
				name        = "test-branch-push-rule"
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
					# 'max_file_size' is a push-only rule, not valid for branch target
					max_file_size {
						max_file_size = 100
					}
				}
			}
		`, randomID)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasPaidOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      config,
					ExpectError: regexp.MustCompile("rule .* is not valid for branch target"),
				},
			},
		})
	})

	t.Run("Validates tag target rejects push-only rules", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name         = "tf-acc-test-tag-push-%s"
				auto_init    = true
				vulnerability_alerts = true

				visibility = "private"
			}

			resource "github_repository_ruleset" "test" {
				name        = "test-tag-push-rule"
				repository  = github_repository.test.id
				target      = "tag"
				enforcement = "active"

				conditions {
					ref_name {
						include = ["~ALL"]
						exclude = []
					}
				}

				rules {
					# 'file_path_restriction' is a push-only rule, not valid for tag target
					file_path_restriction {
						restricted_file_paths = ["secrets/"]
					}
				}
			}
		`, randomID)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasPaidOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      config,
					ExpectError: regexp.MustCompile("rule .* is not valid for tag target"),
				},
			},
		})
	})
}

func TestAccGithubRepositoryRuleset_requiredReviewers(t *testing.T) {
	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
	repoName := fmt.Sprintf("%srepo-ruleset-req-rev-%s", testResourcePrefix, randomID)
	teamName := fmt.Sprintf("%steam-req-rev-%s", testResourcePrefix, randomID)
	rulesetName := fmt.Sprintf("%s-ruleset-req-rev-%s", testResourcePrefix, randomID)
	baseRepoVisibility := "public"

	if testAccConf.authMode == enterprise {
		// This enables repos to be created even in GHEC EMU
		baseRepoVisibility = "private"
	}

	config := fmt.Sprintf(`
resource "github_repository" "test" {
	name      = "%s"
	auto_init = true
	visibility = "%s"

  ignore_vulnerability_alerts_during_read = true
}

resource "github_team" "test" {
	name = "%s"
}

resource "github_team_repository" "test" {
	team_id    = github_team.test.id
	repository = github_repository.test.name
	permission = "push"
}

resource "github_repository_ruleset" "test" {
	name        = "%s"
	repository  = github_repository.test.name
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
			allowed_merge_methods = ["merge", "squash"]
			required_approving_review_count = 1

			required_reviewers {
				reviewer {
					id   = github_team.test.id
					type = "Team"
				}
				file_patterns     = ["*.go"]
				minimum_approvals = 1
			}
		}
	}

	depends_on = [github_team_repository.test]
}
`, repoName, baseRepoVisibility, teamName, rulesetName)

	// Updated config: change minimum_approvals from 1 to 2
	configUpdated := fmt.Sprintf(`
resource "github_repository" "test" {
	name      = "%s"
	auto_init = true
	visibility = "%s"

  ignore_vulnerability_alerts_during_read = true
}

resource "github_team" "test" {
	name = "%s"
}

resource "github_team_repository" "test" {
	team_id    = github_team.test.id
	repository = github_repository.test.name
	permission = "push"
}

resource "github_repository_ruleset" "test" {
	name        = "%s"
	repository  = github_repository.test.name
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
			allowed_merge_methods = ["merge", "squash"]
			required_approving_review_count = 1

			required_reviewers {
				reviewer {
					id   = github_team.test.id
					type = "Team"
				}
				file_patterns     = ["*.go"]
				minimum_approvals = 2
			}
		}
	}

	depends_on = [github_team_repository.test]
}
`, repoName, baseRepoVisibility, teamName, rulesetName)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { skipUnlessHasOrgs(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("github_repository_ruleset.test", "name", rulesetName),
					resource.TestCheckResourceAttr("github_repository_ruleset.test", "target", "branch"),
					resource.TestCheckResourceAttr("github_repository_ruleset.test", "enforcement", "active"),
					resource.TestCheckResourceAttr("github_repository_ruleset.test", "rules.0.pull_request.0.required_reviewers.#", "1"),
					resource.TestCheckResourceAttr("github_repository_ruleset.test", "rules.0.pull_request.0.required_reviewers.0.minimum_approvals", "1"),
					resource.TestCheckResourceAttr("github_repository_ruleset.test", "rules.0.pull_request.0.required_reviewers.0.file_patterns.#", "1"),
					resource.TestCheckResourceAttr("github_repository_ruleset.test", "rules.0.pull_request.0.required_reviewers.0.file_patterns.0", "*.go"),
					resource.TestCheckResourceAttr("github_repository_ruleset.test", "rules.0.pull_request.0.required_reviewers.0.reviewer.0.type", "Team"),
				),
			},
			{
				Config: configUpdated,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("github_repository_ruleset.test", "rules.0.pull_request.0.required_reviewers.0.minimum_approvals", "2"),
				),
			},
			{
				ResourceName:            "github_repository_ruleset.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdFunc:       importRepositoryRulesetByResourcePaths("github_repository.test", "github_repository_ruleset.test"),
				ImportStateVerifyIgnore: []string{"etag"},
			},
		},
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
