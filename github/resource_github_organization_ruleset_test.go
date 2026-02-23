package github

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccGithubOrganizationRuleset(t *testing.T) {
	t.Run("create_branch_ruleset", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-org-ruleset-%s", testResourcePrefix, randomID)
		rulesetName := fmt.Sprintf("%s-branch-ruleset-%s", testResourcePrefix, randomID)

		workflowFilePath := ".github/workflows/echo.yaml"

		config := fmt.Sprintf(`
locals {
		workflow_content = <<EOT
name: Echo Workflow

on: [pull_request]

jobs:
	echo:
		runs-on: ubuntu-latest
		steps:
			- run: echo "Hello, world!"
EOT
}
resource "github_repository" "test" {
	name = "%s"
	visibility = "private"
	auto_init = true
}

resource "github_repository_file" "workflow_file" {
	repository          = github_repository.test.name
	branch              = "main"
	file                = "%[3]s"
	content             = replace(local.workflow_content, "\t", " ") # NOTE: 'content' must be indented with spaces, not tabs
	commit_message      = "Managed by Terraform"
	commit_author       = "Terraform User"
	commit_email        = "terraform@example.com"
}

resource "github_actions_repository_access_level" "test" {
	repository = github_repository.test.name
	access_level = "organization"
}

resource "github_organization_ruleset" "test" {
	name        = "%[2]s"
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

	bypass_actors {
		actor_id    = 1
		actor_type  = "OrganizationAdmin"
		bypass_mode = "always"
	}

	conditions {
		repository_name {
			include = ["~ALL"]
			exclude = []
		}

		ref_name {
			include = ["~ALL"]
			exclude = []
		}
	}

	rules {
		creation = true

		update = true

		deletion                = true
		required_linear_history = true

		required_signatures = false

		pull_request {
			required_approving_review_count   = 2
			required_review_thread_resolution = true
			require_code_owner_review         = true
			dismiss_stale_reviews_on_push     = true
			require_last_push_approval        = true
		}

		copilot_code_review {
			review_on_push             = true
			review_draft_pull_requests = false
		}

		required_status_checks {

			required_check {
				context = "ci"
			}

			strict_required_status_checks_policy = true
			do_not_enforce_on_create             = true
		}

		required_workflows {
			do_not_enforce_on_create = true
			required_workflow {
				path          = github_repository_file.workflow_file.file
				repository_id = github_repository.test.repo_id
				ref           = "main" # Default ref is master
			}
		}

		required_code_scanning {
			required_code_scanning_tool {
			alerts_threshold = "errors"
			security_alerts_threshold = "high_or_higher"
			tool = "CodeQL"
			}
		}

		branch_name_pattern {
			name     = "test"
			negate   = false
			operator = "starts_with"
			pattern  = "test"
		}

		non_fast_forward = true
	}
}
`, repoName, rulesetName, workflowFilePath)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasPaidOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "name", rulesetName),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "target", "branch"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "enforcement", "active"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "bypass_actors.#", "3"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "bypass_actors.0.actor_type", "DeployKey"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "bypass_actors.0.bypass_mode", "always"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "bypass_actors.1.actor_id", "5"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "bypass_actors.1.actor_type", "RepositoryRole"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "bypass_actors.1.bypass_mode", "always"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "bypass_actors.2.actor_id", "1"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "bypass_actors.2.actor_type", "OrganizationAdmin"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "bypass_actors.2.bypass_mode", "always"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "rules.0.pull_request.0.allowed_merge_methods.#", "3"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "rules.0.required_workflows.0.do_not_enforce_on_create", "true"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "rules.0.required_workflows.0.required_workflow.0.path", workflowFilePath),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "rules.0.required_code_scanning.0.required_code_scanning_tool.0.alerts_threshold", "errors"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "rules.0.required_code_scanning.0.required_code_scanning_tool.0.security_alerts_threshold", "high_or_higher"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "rules.0.required_code_scanning.0.required_code_scanning_tool.0.tool", "CodeQL"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "rules.0.copilot_code_review.0.review_on_push", "true"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "rules.0.copilot_code_review.0.review_draft_pull_requests", "false"),
					),
				},
			},
		})
	})

	t.Run("create_ruleset_with_repository_property", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		rulesetName := fmt.Sprintf("%s-repo-prop-ruleset-%s", testResourcePrefix, randomID)
		propName := fmt.Sprintf("e2e_test_team_%s", randomID)

		config := fmt.Sprintf(`
resource "github_organization_custom_properties" "team" {
	property_name  = "%[2]s"
	value_type     = "single_select"
	required       = false
	allowed_values = ["blue", "red", "backend", "platform"]
}

resource "github_organization_ruleset" "test" {
	name        = "%[1]s"
	target      = "branch"
	enforcement = "active"

	conditions {
		repository_property {
			include = [{
				name            = "%[2]s"
				source          = "custom"
				property_values = ["blue"]
			}]
			exclude = []
		}

		ref_name {
			include = ["~DEFAULT_BRANCH"]
			exclude = []
		}
	}

	rules {
		creation                = true
		update                  = true
		deletion                = true
		required_linear_history = true
	}

	depends_on = [github_organization_custom_properties.team]
}
`, rulesetName, propName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasPaidOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "name", rulesetName),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "target", "branch"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "enforcement", "active"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "conditions.0.repository_property.0.include.#", "1"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "conditions.0.repository_property.0.include.0.name", propName),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "conditions.0.repository_property.0.include.0.source", "custom"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "conditions.0.repository_property.0.include.0.property_values.#", "1"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "conditions.0.repository_property.0.include.0.property_values.0", "blue"),
					),
				},
			},
		})
	})

	t.Run("create_ruleset_with_repository_property_exclude", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		rulesetName := fmt.Sprintf("%s-repo-prop-exclude-ruleset-%s", testResourcePrefix, randomID)
		propName := fmt.Sprintf("e2e_test_team_%s", randomID)

		config := fmt.Sprintf(`
resource "github_organization_custom_properties" "team" {
	property_name  = "%[2]s"
	value_type     = "single_select"
	required       = false
	allowed_values = ["blue", "red", "backend", "platform"]
}

resource "github_organization_ruleset" "test" {
	name        = "%[1]s"
	target      = "branch"
	enforcement = "active"

	conditions {
		repository_property {
			include = []
			exclude = [{
				name            = "%[2]s"
				source          = "custom"
				property_values = ["red"]
			}]
		}

		ref_name {
			include = ["~DEFAULT_BRANCH"]
			exclude = []
		}
	}

	rules {
		required_linear_history = true
	}

	depends_on = [github_organization_custom_properties.team]
}
`, rulesetName, propName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasPaidOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "name", rulesetName),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "conditions.0.repository_property.0.include.#", "0"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "conditions.0.repository_property.0.exclude.#", "1"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "conditions.0.repository_property.0.exclude.0.name", propName),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "conditions.0.repository_property.0.exclude.0.source", "custom"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "conditions.0.repository_property.0.exclude.0.property_values.0", "red"),
					),
				},
			},
		})
	})

	t.Run("create_ruleset_with_multiple_repository_properties", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		rulesetName := fmt.Sprintf("%s-repo-prop-multiple-%s", testResourcePrefix, randomID)
		propEnvironmentName := fmt.Sprintf("e2e_test_environment_%s", randomID)
		propTierName := fmt.Sprintf("e2e_test_tier_%s", randomID)

		config := fmt.Sprintf(`
resource "github_organization_custom_properties" "environment" {
	property_name  = "%[2]s"
	value_type     = "single_select"
	required       = false
	allowed_values = ["production", "staging"]
}

resource "github_organization_custom_properties" "tier" {
	property_name  = "%[3]s"
	value_type     = "single_select"
	required       = false
	allowed_values = ["premium", "enterprise"]
}

resource "github_organization_ruleset" "test" {
	name        = "%[1]s"
	target      = "branch"
	enforcement = "active"

	conditions {
		repository_property {
			include = [
				{
					name            = "%[2]s"
					source          = "custom"
					property_values = ["production"]
				},
				{
					name            = "%[3]s"
					source          = "custom"
					property_values = ["premium", "enterprise"]
				}
			]
			exclude = []
		}

		ref_name {
			include = ["~DEFAULT_BRANCH"]
			exclude = []
		}
	}

	rules {
		required_signatures = true
	}

	depends_on = [
		github_organization_custom_properties.environment,
		github_organization_custom_properties.tier,
	]
}
`, rulesetName, propEnvironmentName, propTierName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasPaidOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "name", rulesetName),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "conditions.0.repository_property.0.include.#", "2"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "conditions.0.repository_property.0.include.0.name", propEnvironmentName),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "conditions.0.repository_property.0.include.0.source", "custom"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "conditions.0.repository_property.0.include.0.property_values.#", "1"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "conditions.0.repository_property.0.include.0.property_values.0", "production"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "conditions.0.repository_property.0.include.1.name", propTierName),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "conditions.0.repository_property.0.include.1.source", "custom"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "conditions.0.repository_property.0.include.1.property_values.#", "2"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "conditions.0.repository_property.0.include.1.property_values.0", "premium"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "conditions.0.repository_property.0.include.1.property_values.1", "enterprise"),
					),
				},
			},
		})
	})

	t.Run("update_repository_property", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		rulesetName := fmt.Sprintf("%s-repo-prop-update-%s", testResourcePrefix, randomID)
		propName := fmt.Sprintf("e2e_test_team_%s", randomID)

		config := fmt.Sprintf(`
resource "github_organization_custom_properties" "team" {
	property_name  = "%[2]s"
	value_type     = "single_select"
	required       = false
	allowed_values = ["blue", "red", "backend", "platform"]
}

resource "github_organization_ruleset" "test" {
	name        = "%[1]s"
	target      = "branch"
	enforcement = "active"

	conditions {
		repository_property {
			include = [{
				name            = "%[2]s"
				source          = "custom"
				property_values = ["blue"]
			}]
			exclude = []
		}

		ref_name {
			include = ["~DEFAULT_BRANCH"]
			exclude = []
		}
	}

	rules {
		creation = true
	}

	depends_on = [github_organization_custom_properties.team]
}
`, rulesetName, propName)

		configUpdated := fmt.Sprintf(`
resource "github_organization_custom_properties" "team" {
	property_name  = "%[2]s"
	value_type     = "single_select"
	required       = false
	allowed_values = ["blue", "red", "backend", "platform"]
}

resource "github_organization_ruleset" "test" {
	name        = "%[1]s"
	target      = "branch"
	enforcement = "active"

	conditions {
		repository_property {
			include = [{
				name            = "%[2]s"
				source          = "custom"
				property_values = ["backend", "platform"]
			}]
			exclude = []
		}

		ref_name {
			include = ["~DEFAULT_BRANCH"]
			exclude = []
		}
	}

	rules {
		creation = true
		update   = true
	}

	depends_on = [github_organization_custom_properties.team]
}
`, rulesetName, propName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasPaidOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "name", rulesetName),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "conditions.0.repository_property.0.include.0.property_values.#", "1"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "conditions.0.repository_property.0.include.0.property_values.0", "blue"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "conditions.0.repository_property.0.exclude.#", "0"),
					),
				},
				{
					Config: configUpdated,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "name", rulesetName),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "conditions.0.repository_property.0.include.0.property_values.#", "2"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "conditions.0.repository_property.0.include.0.property_values.0", "backend"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "conditions.0.repository_property.0.include.0.property_values.1", "platform"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "conditions.0.repository_property.0.exclude.#", "0"),
					),
				},
			},
		})
	})

	t.Run("create_push_ruleset", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		rulesetName := fmt.Sprintf("%s-push-ruleset-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
resource "github_organization_ruleset" "test" {
	name        = "%s"
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

	bypass_actors {
		actor_id    = 1
		actor_type  = "OrganizationAdmin"
		bypass_mode = "always"
	}

	conditions {
		repository_name {
			include = ["~ALL"]
			exclude = []
		}
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
`, rulesetName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasPaidOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "name", rulesetName),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "target", "push"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "enforcement", "active"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "bypass_actors.#", "3"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "bypass_actors.0.actor_type", "DeployKey"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "bypass_actors.0.bypass_mode", "always"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "bypass_actors.1.actor_id", "5"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "bypass_actors.1.actor_type", "RepositoryRole"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "bypass_actors.1.bypass_mode", "always"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "bypass_actors.2.actor_id", "1"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "bypass_actors.2.actor_type", "OrganizationAdmin"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "bypass_actors.2.bypass_mode", "always"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "rules.0.file_path_restriction.0.restricted_file_paths.0", "test.txt"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "rules.0.max_file_size.0.max_file_size", "99"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "rules.0.file_extension_restriction.0.restricted_file_extensions.0", "*.zip"),
					),
				},
			},
		})
	})

	t.Run("update_ruleset_name", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		name := fmt.Sprintf("test-acc-ruleset-%s", randomID)
		nameUpdated := fmt.Sprintf("test-acc-ruleset-updated-%s", randomID)

		config := `
resource "github_organization_ruleset" "test" {
	name        = "%s"
	target      = "branch"
	enforcement = "active"

	conditions {
		repository_name {
			include = ["~ALL"]
			exclude = []
		}

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
			PreCheck:          func() { skipUnlessHasPaidOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, name),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "name", name),
					),
				},
				{
					Config: fmt.Sprintf(config, nameUpdated),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "name", nameUpdated),
					),
				},
			},
		})
	})

	t.Run("update_clear_bypass_actors", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		rulesetName := fmt.Sprintf("%s-bypass-ruleset-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
resource "github_organization_ruleset" "test" {
	name        = "%s"
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

	bypass_actors {
		actor_id    = 1
		actor_type  = "OrganizationAdmin"
		bypass_mode = "always"
	}

	conditions {
		repository_name {
			include = ["~ALL"]
			exclude = []
		}

		ref_name {
			include = ["~ALL"]
			exclude = []
		}
	}

	rules {
		creation = true
	}
}
`, rulesetName)

		configUpdated := fmt.Sprintf(`
resource "github_organization_ruleset" "test" {
	name        = "%s"
	target      = "branch"
	enforcement = "active"

	conditions {
		repository_name {
			include = ["~ALL"]
			exclude = []
		}

		ref_name {
			include = ["~ALL"]
			exclude = []
		}
	}

	rules {
		creation = true
	}
}
`, rulesetName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasPaidOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "bypass_actors.#", "3"),
					),
				},
				{
					Config: configUpdated,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "bypass_actors.#", "0"),
					),
				},
			},
		})
	})

	t.Run("update_bypass_mode", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

		bypassMode := "always"
		bypassModeUpdated := "exempt"

		config := `
resource "github_organization_ruleset" "test" {
	name        = "test-bypass-update-%s"
	target      = "branch"
	enforcement = "active"

	bypass_actors {
		actor_id    = 1
		actor_type  = "OrganizationAdmin"
		bypass_mode = "%s"
	}

	conditions {

		repository_name {
			include = ["~ALL"]
			exclude = []
		}

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
			PreCheck:          func() { skipUnlessHasPaidOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, randomID, bypassMode),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "bypass_actors.0.bypass_mode", bypassMode),
					),
				},
				{
					Config: fmt.Sprintf(config, randomID, bypassModeUpdated),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "bypass_actors.0.bypass_mode", bypassModeUpdated),
					),
				},
			},
		})
	})

	t.Run("import", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

		config := fmt.Sprintf(`
	resource "github_organization_ruleset" "test" {
		name        = "test-%s"
		target      = "branch"
		enforcement = "active"

		conditions {
			repository_name {
				include = ["~ALL"]
				exclude = []
			}

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

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasPaidOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
				},
				{
					ResourceName:            "github_organization_ruleset.test",
					ImportState:             true,
					ImportStateVerify:       true,
					ImportStateVerifyIgnore: []string{"etag"},
				},
			},
		})
	})

	t.Run("validates_branch_target_requires_ref_name_condition", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		config := fmt.Sprintf(`
			resource "github_organization_ruleset" "test" {
				name        = "test-validation-%s"
				target      = "branch"
				enforcement = "active"

				conditions {
					repository_name {
						include = ["~ALL"]
						exclude = []
					}
				}

				rules {
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
					ExpectError: regexp.MustCompile("ref_name must be set for branch target"),
				},
			},
		})
	})

	t.Run("validates_tag_target_requires_ref_name_condition", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		config := fmt.Sprintf(`
			resource "github_organization_ruleset" "test" {
				name        = "test-tag-no-conditions-%s"
				target      = "tag"
				enforcement = "active"

				conditions {
					repository_name {
						include = ["~ALL"]
						exclude = []
					}
				}

				rules {
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
					ExpectError: regexp.MustCompile("ref_name must be set for tag target"),
				},
			},
		})
	})

	t.Run("validates_push_target_rejects_ref_name_condition", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		resourceName := "test-push-reject-ref-name"
		config := fmt.Sprintf(`
			resource "github_organization_ruleset" "%s" {
				name        = "test-push-with-ref-%s"
				target      = "push"
				enforcement = "active"

				conditions {
					ref_name {
						include = ["~ALL"]
						exclude = []
					}
					repository_name {
						include = ["~ALL"]
						exclude = []
					}
				}

				rules {
					# Push rulesets only support push-specific rules
					max_file_size {
						max_file_size = 100
					}
				}
			}
		`, resourceName, randomID)

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

	t.Run("validates_push_target_rejects_branch_or_tag_rules", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		resourceName := "test-push-reject-branch-rules"
		config := fmt.Sprintf(`
			resource "github_organization_ruleset" "%s" {
				name        = "test-push-branch-rule-%s"
				target      = "push"
				enforcement = "active"

				conditions {
					repository_name {
						include = ["~ALL"]
						exclude = []
					}
				}

				rules {
					# 'creation' is a branch/tag rule, not valid for push target
					creation = true
				}
			}
		`, resourceName, randomID)

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

	t.Run("validates_branch_target_rejects_push-only_rules", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		resourceName := "test-branch-reject-push-rules"
		config := fmt.Sprintf(`
			resource "github_organization_ruleset" "%s" {
				name        = "test-branch-push-rule-%s"
				target      = "branch"
				enforcement = "active"

				conditions {
					ref_name {
						include = ["~ALL"]
						exclude = []
					}
					repository_name {
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
		`, resourceName, randomID)

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

	t.Run("validates_conditions_require_exactly_one_repository_targeting", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		resourceName := "test-multiple-repo-targeting"
		config := fmt.Sprintf(`
			resource "github_organization_ruleset" "%s" {
				name        = "test-multiple-targeting-%s"
				target      = "branch"
				enforcement = "active"

				conditions {
					ref_name {
						include = ["~ALL"]
						exclude = []
					}
					repository_name {
						include = ["~ALL"]
						exclude = []
					}
					repository_id = [123]
				}

				rules {
					creation = true
				}
			}
		`, resourceName, randomID)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasPaidOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      config,
					ExpectError: regexp.MustCompile(`(?s)only one of.*conditions\.0\.repository_id.*conditions\.0\.repository_name.*conditions\.0\.repository_property.*can be specified`),
				},
			},
		})
	})

	t.Run("validates_conditions_require_at_least_one_repository_targeting", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		resourceName := "test-no-repo-targeting"
		config := fmt.Sprintf(`
			resource "github_organization_ruleset" "%s" {
				name        = "test-no-targeting-%s"
				target      = "branch"
				enforcement = "active"

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
		`, resourceName, randomID)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasPaidOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      config,
					ExpectError: regexp.MustCompile(`(?s)one of.*conditions\.0\.repository_id.*conditions\.0\.repository_name.*conditions\.0\.repository_property.*must be specified`),
				},
			},
		})
	})

	t.Run("validates_repository_property_works_as_single_targeting_option", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		rulesetName := fmt.Sprintf("%s-repo-prop-only-%s", testResourcePrefix, randomID)
		propName := fmt.Sprintf("e2e_test_environment_%s", randomID)

		config := fmt.Sprintf(`
resource "github_organization_custom_properties" "environment" {
	property_name  = "%[2]s"
	value_type     = "single_select"
	required       = false
	allowed_values = ["production", "staging"]
}

resource "github_organization_ruleset" "test" {
	name        = "%[1]s"
	target      = "branch"
	enforcement = "active"

	conditions {
		repository_property {
			include = [{
				name            = "%[2]s"
				source          = "custom"
				property_values = ["production", "staging"]
			}]
			exclude = []
		}

		ref_name {
			include = ["~DEFAULT_BRANCH"]
			exclude = []
		}
	}

	rules {
		creation = true
		update   = true
	}

	depends_on = [github_organization_custom_properties.environment]
}
`, rulesetName, propName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasPaidOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "name", rulesetName),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "target", "branch"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "enforcement", "active"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "conditions.0.repository_property.0.include.#", "1"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "conditions.0.repository_property.0.include.0.name", propName),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "conditions.0.repository_property.0.include.0.source", "custom"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "conditions.0.repository_property.0.include.0.property_values.#", "2"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "conditions.0.repository_property.0.include.0.property_values.0", "production"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "conditions.0.repository_property.0.include.0.property_values.1", "staging"),
					),
				},
			},
		})
	})

	t.Run("creates_push_ruleset", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		rulesetName := fmt.Sprintf("%stest-push-%s", testResourcePrefix, randomID)
		resourceName := "test-push-ruleset"
		resourceFullName := fmt.Sprintf("github_organization_ruleset.%s", resourceName)
		config := fmt.Sprintf(`
			resource "github_organization_ruleset" "%s" {
				name        = "%s"
				target      = "push"
				enforcement = "active"

				conditions {
					repository_name {
						include = ["~ALL"]
						exclude = []
					}
				}

				rules {
					# Push rulesets only support push-specific rules:
					# file_path_restriction, max_file_path_length, file_extension_restriction, max_file_size
					max_file_size {
						max_file_size = 100
					}
				}
			}
		`, resourceName, rulesetName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasPaidOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr(resourceFullName, "name", rulesetName),
						resource.TestCheckResourceAttr(resourceFullName, "target", "push"),
						resource.TestCheckResourceAttr(resourceFullName, "enforcement", "active"),
						resource.TestCheckResourceAttr(resourceFullName, "rules.0.max_file_size.0.max_file_size", "100"),
					),
				},
			},
		})
	})

	t.Run("validates_rules__required_status_checks_block", func(t *testing.T) {
		t.Run("required_check__context_block_should_not_be_empty", func(t *testing.T) {
			resourceName := "test-required-status-checks-context-is-not-empty"
			randomID := acctest.RandString(5)
			config := fmt.Sprintf(`
			resource "github_organization_ruleset" "%s" {
				name        = "test-context-is-not-empty-%s"
				target      = "branch"
				enforcement = "active"

				conditions {
					ref_name {
						include = ["~ALL"]
						exclude = []
					}
					repository_name {
						include = ["~ALL"]
						exclude = []
					}
				}

				rules {
					required_status_checks {
						required_check {
							context = ""
						}
					}
				}
			}
		`, resourceName, randomID)

			resource.Test(t, resource.TestCase{
				PreCheck:          func() { skipUnlessHasPaidOrgs(t) },
				ProviderFactories: providerFactories,
				Steps: []resource.TestStep{
					{
						Config:      config,
						ExpectError: regexp.MustCompile("expected \"context\" to not be an empty string"),
					},
				},
			})
		})
		t.Run("required_check_should_be_required_when_strict_required_status_checks_policy_is_set", func(t *testing.T) {
			resourceName := "test-required-check-is-required"
			randomID := acctest.RandString(5)
			config := fmt.Sprintf(`
			resource "github_organization_ruleset" "%s" {
				name        = "test-required-with-%s"
				target      = "branch"
				enforcement = "active"

				conditions {
					ref_name {
						include = ["~ALL"]
						exclude = []
					}
					repository_name {
						include = ["~ALL"]
						exclude = []
					}
				}

				rules {
					required_status_checks {
						strict_required_status_checks_policy = true
					}
				}
			}
		`, resourceName, randomID)

			resource.Test(t, resource.TestCase{
				PreCheck:          func() { skipUnlessHasPaidOrgs(t) },
				ProviderFactories: providerFactories,
				Steps: []resource.TestStep{
					{
						Config:      config,
						ExpectError: regexp.MustCompile("Insufficient required_check blocks"),
					},
				},
			})
		})
	})

	t.Run("updates_required_reviewers", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		teamName := fmt.Sprintf("%steam-req-rev-%s", testResourcePrefix, randomID)
		rulesetName := fmt.Sprintf("%s-ruleset-req-rev-%s", testResourcePrefix, randomID)

		config := `
resource "github_team" "test" {
	name = "%s"
}

resource "github_organization_ruleset" "test" {
	name        = "%s"
	target      = "branch"
	enforcement = "active"

	conditions {
		repository_name {
			include = ["~ALL"]
			exclude = []
		}

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
				file_patterns     = ["*.go", "src/**/*.ts"]
				minimum_approvals = %d
			}
		}
	}
}
`

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasPaidOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, teamName, rulesetName, 1),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "name", rulesetName),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "target", "branch"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "enforcement", "active"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "rules.0.pull_request.0.required_reviewers.#", "1"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "rules.0.pull_request.0.required_reviewers.0.minimum_approvals", "1"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "rules.0.pull_request.0.required_reviewers.0.file_patterns.#", "2"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "rules.0.pull_request.0.required_reviewers.0.file_patterns.0", "*.go"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "rules.0.pull_request.0.required_reviewers.0.file_patterns.1", "src/**/*.ts"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "rules.0.pull_request.0.required_reviewers.0.reviewer.0.type", "Team"),
					),
				},
				{
					Config: fmt.Sprintf(config, teamName, rulesetName, 2),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "rules.0.pull_request.0.required_reviewers.0.minimum_approvals", "2"),
					),
				},
				{
					ResourceName:            "github_organization_ruleset.test",
					ImportState:             true,
					ImportStateVerify:       true,
					ImportStateVerifyIgnore: []string{"etag"},
				},
			},
		})
	})
	t.Run("creates_rule_with_multiple_required_reviewers", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		teamName1 := fmt.Sprintf("%steam-req-rev-1-%s", testResourcePrefix, randomID)
		teamName2 := fmt.Sprintf("%steam-req-rev-2-%s", testResourcePrefix, randomID)
		rulesetName := fmt.Sprintf("%s-ruleset-multi-rev-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
resource "github_team" "test1" {
	name = "%s"
}

resource "github_team" "test2" {
	name = "%s"
}

resource "github_organization_ruleset" "test" {
	name        = "%s"
	target      = "branch"
	enforcement = "active"

	conditions {
		repository_name {
			include = ["~ALL"]
			exclude = []
		}

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
					id   = github_team.test1.id
					type = "Team"
				}
				file_patterns     = ["*.go"]
				minimum_approvals = 1
			}

			required_reviewers {
				reviewer {
					id   = github_team.test2.id
					type = "Team"
				}
				file_patterns     = ["*.md", "docs/**/*"]
				minimum_approvals = 1
			}
		}
	}
}
`, teamName1, teamName2, rulesetName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasPaidOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "name", rulesetName),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "target", "branch"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "enforcement", "active"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "rules.0.pull_request.0.required_reviewers.#", "2"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "rules.0.pull_request.0.required_reviewers.0.minimum_approvals", "1"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "rules.0.pull_request.0.required_reviewers.0.file_patterns.#", "1"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "rules.0.pull_request.0.required_reviewers.0.file_patterns.0", "*.go"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "rules.0.pull_request.0.required_reviewers.1.minimum_approvals", "1"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "rules.0.pull_request.0.required_reviewers.1.file_patterns.#", "2"),
					),
				},
			},
		})
	})
}

func TestOrganizationPushRulesetSupport(t *testing.T) {
	// Test that organization push rulesets support all push-specific rules
	// This is a unit test since it only validates the expand/flatten functionality

	rulesMap := map[string]any{
		"file_path_restriction": []any{
			map[string]any{
				"restricted_file_paths": []any{"secrets/", "*.key", "private/"},
			},
		},
		"max_file_size": []any{
			map[string]any{
				"max_file_size": 100, // 100MB
			},
		},
		"max_file_path_length": []any{
			map[string]any{
				"max_file_path_length": 250,
			},
		},
		"file_extension_restriction": []any{
			map[string]any{
				"restricted_file_extensions": schema.NewSet(schema.HashString, []any{".exe", ".bat", ".sh", ".ps1"}),
			},
		},
	}

	input := []any{rulesMap}

	// Test expand functionality (organization rulesets use org=true)
	expandedRules := expandRules(input, true)

	if expandedRules == nil {
		t.Fatalf("expected expanded rules to not be nil")
	}

	// Verify we have all expected push rule types
	ruleCount := 0
	if expandedRules.FilePathRestriction != nil {
		ruleCount++
		filePathRule := rulesMap["file_path_restriction"].([]any)[0].(map[string]any)
		expectedFilePaths := len(filePathRule["restricted_file_paths"].([]any))
		if len(expandedRules.FilePathRestriction.RestrictedFilePaths) != expectedFilePaths {
			t.Errorf("Expected %d restricted file paths, got %d", expectedFilePaths, len(expandedRules.FilePathRestriction.RestrictedFilePaths))
		}
	}
	if expandedRules.MaxFileSize != nil {
		ruleCount++
		maxFileSizeRule := rulesMap["max_file_size"].([]any)[0].(map[string]any)
		expectedMaxFileSize := int64(maxFileSizeRule["max_file_size"].(int))
		if expandedRules.MaxFileSize.MaxFileSize != expectedMaxFileSize {
			t.Errorf("Expected max file size to be %d, got %d", expectedMaxFileSize, expandedRules.MaxFileSize.MaxFileSize)
		}
	}
	if expandedRules.MaxFilePathLength != nil {
		ruleCount++
		maxPathLengthRule := rulesMap["max_file_path_length"].([]any)[0].(map[string]any)
		expectedMaxPathLength := maxPathLengthRule["max_file_path_length"].(int)
		if expandedRules.MaxFilePathLength.MaxFilePathLength != expectedMaxPathLength {
			t.Errorf("Expected max file path length to be %d, got %d", expectedMaxPathLength, expandedRules.MaxFilePathLength.MaxFilePathLength)
		}
	}
	if expandedRules.FileExtensionRestriction != nil {
		ruleCount++
		fileExtRule := rulesMap["file_extension_restriction"].([]any)[0].(map[string]any)
		expectedExtensions := fileExtRule["restricted_file_extensions"].(*schema.Set).Len()
		if len(expandedRules.FileExtensionRestriction.RestrictedFileExtensions) != expectedExtensions {
			t.Errorf("Expected %d restricted file extensions, got %d", expectedExtensions, len(expandedRules.FileExtensionRestriction.RestrictedFileExtensions))
		}
	}

	expectedRuleCount := len(rulesMap)
	if ruleCount != expectedRuleCount {
		t.Fatalf("Expected %d expanded rules for organization push ruleset, got %d", expectedRuleCount, ruleCount)
	}

	// Test flatten functionality (organization rulesets use org=true)
	flattenedResult := flattenRules(t.Context(), expandedRules, true)

	if len(flattenedResult) != 1 {
		t.Fatalf("Expected 1 flattened result, got %d", len(flattenedResult))
	}

	flattenedRulesMap := flattenedResult[0].(map[string]any)

	// Verify file_path_restriction
	filePathRules := flattenedRulesMap["file_path_restriction"].([]map[string]any)
	if len(filePathRules) != 1 {
		t.Fatalf("Expected 1 file_path_restriction rule, got %d", len(filePathRules))
	}
	restrictedPaths := filePathRules[0]["restricted_file_paths"].([]string)
	if len(restrictedPaths) != 3 {
		t.Errorf("Expected 3 restricted file paths, got %d", len(restrictedPaths))
	}

	// Verify max_file_size
	maxFileSizeRules := flattenedRulesMap["max_file_size"].([]map[string]any)
	if len(maxFileSizeRules) != 1 {
		t.Fatalf("Expected 1 max_file_size rule, got %d", len(maxFileSizeRules))
	}
	if maxFileSizeRules[0]["max_file_size"] != int64(100) {
		t.Errorf("Expected max_file_size to be 100, got %v", maxFileSizeRules[0]["max_file_size"])
	}

	// Verify max_file_path_length
	maxFilePathLengthRules := flattenedRulesMap["max_file_path_length"].([]map[string]any)
	if len(maxFilePathLengthRules) != 1 {
		t.Fatalf("Expected 1 max_file_path_length rule, got %d", len(maxFilePathLengthRules))
	}
	if maxFilePathLengthRules[0]["max_file_path_length"] != 250 {
		t.Errorf("Expected max_file_path_length to be 250, got %v", maxFilePathLengthRules[0]["max_file_path_length"])
	}

	// Verify file_extension_restriction
	fileExtRules := flattenedRulesMap["file_extension_restriction"].([]map[string]any)
	if len(fileExtRules) != 1 {
		t.Fatalf("Expected 1 file_extension_restriction rule, got %d", len(fileExtRules))
	}
	restrictedExts := fileExtRules[0]["restricted_file_extensions"].([]string)
	if len(restrictedExts) != 4 {
		t.Errorf("Expected 4 restricted file extensions, got %d", len(restrictedExts))
	}
}
