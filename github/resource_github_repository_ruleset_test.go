package github

import (
	"fmt"
	"log"
	"regexp"
	"testing"

	"github.com/google/go-github/v89/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-log/tflogtest"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubRepositoryRuleset(t *testing.T) {
	t.Parallel()

	t.Run("create_branch_ruleset", func(t *testing.T) {
		t.Parallel()

		testRepo := mustCreateTestRepository(t)

		config := fmt.Sprintf(`
resource "github_repository_environment" "example" {
	environment  = "test"
	repository   = "%s"
}

resource "github_repository_ruleset" "test" {
	name        = "test"
	repository  = "%s"
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
`, testRepo.GetName(), testRepo.GetName())

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("name"), knownvalue.StringExact("test")),
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("target"), knownvalue.StringExact("branch")),
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("enforcement"), knownvalue.StringExact("active")),
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("bypass_actors"), knownvalue.ListSizeExact(2)),
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("bypass_actors").AtSliceIndex(0).AtMapKey("actor_type"), knownvalue.StringExact("DeployKey")),
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("bypass_actors").AtSliceIndex(0).AtMapKey("bypass_mode"), knownvalue.StringExact("always")),
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("bypass_actors").AtSliceIndex(1).AtMapKey("actor_id"), knownvalue.Int64Exact(5)),
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("bypass_actors").AtSliceIndex(1).AtMapKey("actor_type"), knownvalue.StringExact("RepositoryRole")),
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("bypass_actors").AtSliceIndex(1).AtMapKey("bypass_mode"), knownvalue.StringExact("always")),
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("creation"), knownvalue.Bool(true)),
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("update"), knownvalue.Bool(true)),
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("deletion"), knownvalue.Bool(true)),
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("required_linear_history"), knownvalue.Bool(true)),
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("required_signatures"), knownvalue.Bool(false)),
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("non_fast_forward"), knownvalue.Bool(true)),
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("merge_queue").AtSliceIndex(0).AtMapKey("check_response_timeout_minutes"), knownvalue.Int64Exact(10)),
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("merge_queue").AtSliceIndex(0).AtMapKey("grouping_strategy"), knownvalue.StringExact("ALLGREEN")),
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("merge_queue").AtSliceIndex(0).AtMapKey("max_entries_to_build"), knownvalue.Int64Exact(5)),
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("merge_queue").AtSliceIndex(0).AtMapKey("max_entries_to_merge"), knownvalue.Int64Exact(5)),
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("merge_queue").AtSliceIndex(0).AtMapKey("merge_method"), knownvalue.StringExact("SQUASH")),
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("merge_queue").AtSliceIndex(0).AtMapKey("min_entries_to_merge"), knownvalue.Int64Exact(1)),
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("merge_queue").AtSliceIndex(0).AtMapKey("min_entries_to_merge_wait_minutes"), knownvalue.Int64Exact(60)),
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("required_deployments").AtSliceIndex(0).AtMapKey("required_deployment_environments"), knownvalue.ListSizeExact(1)),
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("required_deployments").AtSliceIndex(0).AtMapKey("required_deployment_environments").AtSliceIndex(0), knownvalue.StringExact("test")),
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("pull_request").AtSliceIndex(0).AtMapKey("allowed_merge_methods"), knownvalue.ListSizeExact(2)),
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("pull_request").AtSliceIndex(0).AtMapKey("required_approving_review_count"), knownvalue.Int64Exact(2)),
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("pull_request").AtSliceIndex(0).AtMapKey("required_review_thread_resolution"), knownvalue.Bool(true)),
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("pull_request").AtSliceIndex(0).AtMapKey("require_code_owner_review"), knownvalue.Bool(true)),
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("pull_request").AtSliceIndex(0).AtMapKey("dismiss_stale_reviews_on_push"), knownvalue.Bool(true)),
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("pull_request").AtSliceIndex(0).AtMapKey("require_last_push_approval"), knownvalue.Bool(true)),
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("required_status_checks").AtSliceIndex(0).AtMapKey("required_check"), knownvalue.SetSizeExact(1)),
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("required_status_checks").AtSliceIndex(0).AtMapKey("strict_required_status_checks_policy"), knownvalue.Bool(true)),
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("required_status_checks").AtSliceIndex(0).AtMapKey("do_not_enforce_on_create"), knownvalue.Bool(true)),
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("required_code_scanning").AtSliceIndex(0).AtMapKey("required_code_scanning_tool"), knownvalue.SetSizeExact(1)),
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("copilot_code_review").AtSliceIndex(0).AtMapKey("review_on_push"), knownvalue.Bool(true)),
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("copilot_code_review").AtSliceIndex(0).AtMapKey("review_draft_pull_requests"), knownvalue.Bool(false)),
					},
				},
			},
		})
	})

	t.Run("validates_forked_repo_for_update_rule_parameter_branch_ruleset", func(t *testing.T) {
		t.Parallel()

		testRepo := mustCreateTestRepository(t)

		config := fmt.Sprintf(`
resource "github_repository_ruleset" "test" {
	name        = "test"
	repository  = "%s"
	target      = "branch"
	enforcement = "active"

	conditions {
		ref_name {
			include = ["refs/heads/main"]
			exclude = []
		}
	}

	rules {
		update = true
		update_allows_fetch_and_merge = %%t
	}
}
`, testRepo.GetName())

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      fmt.Sprintf(config, true),
					ExpectError: regexp.MustCompile(`cannot set update_allows_fetch_and_merge when repository is not a forked repository`),
				},
				{
					Config: fmt.Sprintf(config, false),
				},
				{
					Config:      fmt.Sprintf(config, true),
					ExpectError: regexp.MustCompile(`cannot set update_allows_fetch_and_merge when repository is not a forked repository`),
				},
			},
		})
	})

	t.Run("create_branch_ruleset_with_update_allows_fetch_and_merge_on_fork", func(t *testing.T) {
		t.Parallel()

		randomID := acctest.RandString(5)
		repoName := fmt.Sprintf("%sfork-ruleset-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
resource "github_repository" "forked" {
	name         = "%s"
	fork         = true
	source_owner = "integrations"
	source_repo  = "terraform-provider-github"
}

resource "github_repository_ruleset" "test" {
	name        = "test-fork-update-allows-fetch-and-merge"
	repository  = github_repository.forked.name
	target      = "branch"
	enforcement = "active"

	conditions {
		ref_name {
			include = ["~ALL"]
			exclude = []
		}
	}

	rules {
		update = true
		update_allows_fetch_and_merge = true
	}
}
`, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_repository_ruleset.test", plancheck.ResourceActionCreate),
						},
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("target"), knownvalue.StringExact("branch")),
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("update"), knownvalue.Bool(true)),
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("update_allows_fetch_and_merge"), knownvalue.Bool(true)),
					},
				},
			},
		})
	})

	t.Run("create_branch_ruleset_with_user_bypass_actor", func(t *testing.T) {
		t.Parallel()

		var username string
		if testAccConf.authMode == individual {
			username = testAccConf.owner
		} else {
			skipUnlessHasOrgUser1(t)
			username = testAccConf.testOrgUser1
		}

		repo := mustCreateTestRepository(t)

		config := fmt.Sprintf(`
data "github_user" "current" {
	username = "%s"
}

resource "github_repository_ruleset" "test" {
	name        = "test-user-bypass"
	repository  = "%s"
	target      = "branch"
	enforcement = "active"

	bypass_actors {
		actor_id    = tonumber(data.github_user.current.id)
		actor_type  = "User"
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
`, username, repo.GetName())

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("bypass_actors").AtSliceIndex(0).AtMapKey("actor_type"), knownvalue.StringExact("User")),
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("bypass_actors").AtSliceIndex(0).AtMapKey("bypass_mode"), knownvalue.StringExact("always")),
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("bypass_actors").AtSliceIndex(0).AtMapKey("actor_id"), knownvalue.NotNull()),
					},
				},
			},
		})
	})

	t.Run("create_branch_ruleset_with_enterprise_features", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)

		config := fmt.Sprintf(`
	resource "github_repository_ruleset" "test" {
		name        = "test"
		repository  = "%s"
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
`, repo.GetName())

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("enforcement"), knownvalue.StringExact("active")),
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("branch_name_pattern").AtSliceIndex(0).AtMapKey("name"), knownvalue.StringExact("test")),
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("branch_name_pattern").AtSliceIndex(0).AtMapKey("negate"), knownvalue.Bool(false)),
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("branch_name_pattern").AtSliceIndex(0).AtMapKey("operator"), knownvalue.StringExact("starts_with")),
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("branch_name_pattern").AtSliceIndex(0).AtMapKey("pattern"), knownvalue.StringExact("test")),
					},
				},
			},
		})
	})

	t.Run("creates_push_ruleset", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateInternalTestRepository(t)

		config := fmt.Sprintf(`
resource "github_repository_ruleset" "test" {
	name        = "test-push"
	repository  = "%s"
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

`, repo.GetName())

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasPaidOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_repository_ruleset.test", plancheck.ResourceActionCreate),
						},
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("name"), knownvalue.StringExact("test-push")),
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("target"), knownvalue.StringExact("push")),
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("enforcement"), knownvalue.StringExact("active")),
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("bypass_actors"), knownvalue.ListSizeExact(2)),
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("bypass_actors").AtSliceIndex(0).AtMapKey("actor_type"), knownvalue.StringExact("DeployKey")),
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("bypass_actors").AtSliceIndex(0).AtMapKey("bypass_mode"), knownvalue.StringExact("always")),
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("bypass_actors").AtSliceIndex(1).AtMapKey("actor_id"), knownvalue.Int64Exact(5)),
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("bypass_actors").AtSliceIndex(1).AtMapKey("actor_type"), knownvalue.StringExact("RepositoryRole")),
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("bypass_actors").AtSliceIndex(1).AtMapKey("bypass_mode"), knownvalue.StringExact("always")),
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("file_path_restriction").AtSliceIndex(0).AtMapKey("restricted_file_paths").AtSliceIndex(0), knownvalue.StringExact("test.txt")),
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("max_file_size").AtSliceIndex(0).AtMapKey("max_file_size"), knownvalue.Int64Exact(99)),
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("file_extension_restriction").AtSliceIndex(0).AtMapKey("restricted_file_extensions").AtSliceIndex(0), knownvalue.StringExact("*.zip")),
					},
				},
			},
		})
	})

	t.Run("update_ruleset_name", func(t *testing.T) {
		t.Parallel()

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repo := mustCreateTestRepository(t)

		name := fmt.Sprintf("ruleset-%s", randomID)
		nameUpdated := fmt.Sprintf("%s-renamed", name)

		config := fmt.Sprintf(`
resource "github_repository_ruleset" "test" {
	name        = "%%s"
	repository  = "%s"
	target      = "branch"
	enforcement = "active"

	rules {
		creation = true
	}
}
`, repo.GetName())

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, name),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("name"), knownvalue.StringExact(name)),
					},
				},
				{
					Config: fmt.Sprintf(config, nameUpdated),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_repository_ruleset.test", plancheck.ResourceActionUpdate),
						},
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("name"), knownvalue.StringExact(nameUpdated)),
					},
				},
			},
		})
	})

	t.Run("update_clear_bypass_actors", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)

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
		baseConfig := fmt.Sprintf(`
resource "github_repository_ruleset" "test" {
	name        = "test-bypass"
	repository  = "%s"
	target      = "branch"
	enforcement = "active"

	conditions {
		ref_name {
			include = ["~ALL"]
			exclude = []
		}
	}

	%%s

	rules {
		creation = true
	}
}
`, repo.GetName())
		config := fmt.Sprintf(baseConfig, bypassActorsConfig)

		configUpdated := fmt.Sprintf(baseConfig, "")
		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("bypass_actors"), knownvalue.ListSizeExact(2)),
					},
				},
				{
					Config: configUpdated,
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_repository_ruleset.test", plancheck.ResourceActionUpdate),
						},
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("bypass_actors"), knownvalue.ListSizeExact(0)),
					},
				},
			},
		})
	})

	t.Run("update_bypass_mode", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)

		bypassMode := "always"
		bypassModeUpdated := "exempt"

		config := `
resource "github_repository_ruleset" "test" {
	name        = "test-bypass-update"
	repository  = "%s"
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
					Config: fmt.Sprintf(config, repo.GetName(), bypassMode),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("bypass_actors").AtSliceIndex(0).AtMapKey("bypass_mode"), knownvalue.StringExact(bypassMode)),
					},
				},
				{
					Config: fmt.Sprintf(config, repo.GetName(), bypassModeUpdated),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_repository_ruleset.test", plancheck.ResourceActionUpdate),
						},
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("bypass_actors").AtSliceIndex(0).AtMapKey("bypass_mode"), knownvalue.StringExact(bypassModeUpdated)),
					},
				},
			},
		})
	})

	t.Run("import", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)

		config := fmt.Sprintf(`
			resource "github_repository_ruleset" "test" {
				name        = "test"
				repository  = "%s"
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
		`, repo.GetName())

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
					ImportStateIdFunc:       importRepositoryRulesetByResourcePaths(repo.GetName(), "github_repository_ruleset.test"),
					ImportStateVerifyIgnore: []string{"etag"},
				},
			},
		})
	})

	t.Run("create_branch_ruleset_with_enterprise_owner_bypass", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)

		config := fmt.Sprintf(`
resource "github_repository_ruleset" "test" {
	name        = "test"
	repository  = "%s"
	target      = "branch"
	enforcement = "active"

	conditions {
		ref_name {
			include = ["~ALL"]
			exclude = []
		}
	}

	bypass_actors {
		actor_type  = "EnterpriseOwner"
		bypass_mode = "always"
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
`, repo.GetName())

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("name"), knownvalue.StringExact("test")),
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("enforcement"), knownvalue.StringExact("active")),
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("bypass_actors").AtSliceIndex(0).AtMapKey("actor_type"), knownvalue.StringExact("EnterpriseOwner")),
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("bypass_actors").AtSliceIndex(0).AtMapKey("bypass_mode"), knownvalue.StringExact("always")),
					},
				},
			},
		})
	})

	t.Run("skips_update_and_delete_on_archived_repository", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)
		enforcementBefore := "active"
		enforcementAfter := "disabled"

		config := fmt.Sprintf(`
resource "github_repository_ruleset" "test" {
	name        = "test"
	repository  = "%s"
	target      = "branch"
	enforcement = "%%s"
	rules { creation = true }
}
`, repo.GetName())

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, enforcementBefore),
				},
				{
					PreConfig: func() { mustArchiveTestRepository(t, repo) },
					Config:    fmt.Sprintf(config, enforcementBefore),
				},
				{
					Config: fmt.Sprintf(config, enforcementAfter),
				},
			},
		})
	})

	t.Run("prevents_creating_ruleset_on_archived_repository", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)
		config := `
resource "github_repository_ruleset" "test" {
	name       = "test"
	repository = "%s"
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
					PreConfig:   func() { mustArchiveTestRepository(t, repo) },
					Config:      fmt.Sprintf(config, repo.GetName()),
					ExpectError: regexp.MustCompile("cannot create ruleset on archived repository"),
				},
			},
		})
	})

	t.Run("validates_push_target_rejects_ref_name_condition", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)

		config := fmt.Sprintf(`

			resource "github_repository_ruleset" "test" {
				name        = "test-push-with-ref"
				repository  = "%s"
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
		`, repo.GetName())

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

	t.Run("validates_push_target_rejects_branch_tag_rules", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)
		config := fmt.Sprintf(`
			resource "github_repository_ruleset" "test" {
				name        = "test-push-branch-rule"
				repository  = "%s"
				target      = "push"
				enforcement = "active"

				rules {
					# 'creation' is a branch/tag rule, not valid for push target
					creation = true
				}
			}
		`, repo.GetName())

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

	t.Run("validates_branch_target_rejects_push_only_rules", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)
		config := fmt.Sprintf(`
			resource "github_repository_ruleset" "test" {
				name        = "test-branch-push-rule"
				repository  = "%s"
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
		`, repo.GetName())

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

	t.Run("validates_tag_target_rejects_push_only_rules", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)
		config := fmt.Sprintf(`
			resource "github_repository_ruleset" "test" {
				name        = "test-tag-push-rule"
				repository  = "%s"
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
		`, repo.GetName())

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

	t.Run("updates_required_reviewers_successfully", func(t *testing.T) {
		t.Parallel()

		randomID := acctest.RandString(5)
		repo := mustCreateTestRepository(t)
		team := mustCreateTestTeam(t, nil)
		rulesetName := fmt.Sprintf("%s-ruleset-req-rev-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
resource "github_team_repository" "test" {
	team_id    = "%[1]d"
	repository = "%[2]s"
	permission = "push"
}

resource "github_repository_ruleset" "test" {
	name        = "%[3]s"
	repository  = "%[2]s"
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
					id   = "%[1]d"
					type = "Team"
				}
				file_patterns     = ["*.go"]
				minimum_approvals = %%d
			}
		}
	}
}
`, team.GetID(), repo.GetName(), rulesetName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, 1),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("name"), knownvalue.StringExact(rulesetName)),
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("target"), knownvalue.StringExact("branch")),
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("enforcement"), knownvalue.StringExact("active")),
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("pull_request").AtSliceIndex(0).AtMapKey("required_reviewers"), knownvalue.ListSizeExact(1)),
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("pull_request").AtSliceIndex(0).AtMapKey("required_reviewers").AtSliceIndex(0).AtMapKey("minimum_approvals"), knownvalue.Int64Exact(1)),
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("pull_request").AtSliceIndex(0).AtMapKey("required_reviewers").AtSliceIndex(0).AtMapKey("file_patterns"), knownvalue.ListSizeExact(1)),
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("pull_request").AtSliceIndex(0).AtMapKey("required_reviewers").AtSliceIndex(0).AtMapKey("file_patterns").AtSliceIndex(0), knownvalue.StringExact("*.go")),
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("pull_request").AtSliceIndex(0).AtMapKey("required_reviewers").AtSliceIndex(0).AtMapKey("reviewer").AtSliceIndex(0).AtMapKey("type"), knownvalue.StringExact("Team")),
					},
				},
				{
					Config: fmt.Sprintf(config, 2),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_repository_ruleset.test", plancheck.ResourceActionUpdate),
						},
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("pull_request").AtSliceIndex(0).AtMapKey("required_reviewers").AtSliceIndex(0).AtMapKey("minimum_approvals"), knownvalue.Int64Exact(2)),
					},
				},
				{
					ResourceName:            "github_repository_ruleset.test",
					ImportState:             true,
					ImportStateVerify:       true,
					ImportStateIdFunc:       importRepositoryRulesetByResourcePaths(repo.GetName(), "github_repository_ruleset.test"),
					ImportStateVerifyIgnore: []string{"etag"},
				},
			},
		})
	})

	// Verify possible regression found by https://github.com/integrations/terraform-provider-github/issues/3509
	t.Run("regression_shows_drift_when_bypass_actors_change_upstream", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)

		config := fmt.Sprintf(`
resource "github_repository_ruleset" "test" {
	name        = "test_regression_bypass_actors_drift"
	repository  = "%s"
	target      = "branch"
	enforcement = "active"


	bypass_actors {
		actor_id    = 946600 # Copilot code review
		actor_type  = "Integration"
		bypass_mode = "always"
	}

	bypass_actors {
		actor_id    = 1143301 # Copilot cloud agent
		actor_type  = "Integration"
		bypass_mode = "always"
	}

	conditions {
		ref_name {
			include = ["~DEFAULT_BRANCH"]
			exclude = []
		}
	}

	rules {
		creation = true
		update = true
		deletion = true
	}
}
`, repo.GetName())

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("bypass_actors"), knownvalue.ListSizeExact(2)),
					},
				},
				{
					Config: config,
					PreConfig: func() {
						client := testAccConf.meta.v3client
						ctx := tflogtest.RootLogger(t.Context(), log.Writer()) // This pattern can be used to capture logs during testing if needed
						// Update the ruleset to remove one of the bypass actors, simulating a change upstream
						rulesets, _, err := client.Repositories.GetAllRulesets(ctx, testAccConf.meta.name, repo.GetName(), &github.RepositoryListRulesetsOptions{})
						if err != nil {
							t.Fatalf("failed to list all repository rulesets: %v", err)
						}
						firstRuleset := rulesets[0]

						integrationActorType := github.BypassActorTypeIntegration
						bypassMode := github.BypassModeAlways

						firstRuleset.BypassActors = []*github.BypassActor{
							{
								ActorID:    new(int64(946600)),
								ActorType:  &integrationActorType,
								BypassMode: &bypassMode,
							},
						}
						tflog.Debug(ctx, "Removing 1 bypass actor from Repo Ruleset")
						ruleset, resp, err := client.Repositories.UpdateRuleset(ctx, testAccConf.meta.name, repo.GetName(), firstRuleset.GetID(), *firstRuleset)
						if err != nil {
							t.Fatalf("failed to remove 2nd bypass actor: %v", err)
						}
						tflog.Debug(ctx, "Successfully removed 1 bypass actor from Repo Ruleset", map[string]any{"bypass_actors": ruleset.GetBypassActors(), "etag": resp.Header.Get("ETag")})
					},
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_repository_ruleset.test", plancheck.ResourceActionUpdate),
						},
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_ruleset.test", tfjsonpath.New("bypass_actors"), knownvalue.ListSizeExact(2)),
					},
				},
			},
		})
	})
}

func importRepositoryRulesetByResourcePaths(repoNodeID, rulesetLogicalName string) resource.ImportStateIdFunc {
	// test importing using an ID of the form <repo-node-id>:<ruleset-id>
	return func(s *terraform.State) (string, error) {
		log.Printf("[DEBUG] Looking up tf state ")

		ruleset := s.RootModule().Resources[rulesetLogicalName]
		if ruleset == nil {
			return "", fmt.Errorf("Cannot find %s in terraform state", rulesetLogicalName)
		}
		rulesetID := ruleset.Primary.ID
		if rulesetID == "" {
			return "", fmt.Errorf("ruleset %s does not have an id in terraform state", rulesetLogicalName)
		}

		return fmt.Sprintf("%s:%s", repoNodeID, rulesetID), nil
	}
}
