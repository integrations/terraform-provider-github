package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestAccGithubOrganizationRuleset(t *testing.T) {
	t.Run("create_branch_ruleset", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-org-ruleset-%s", testResourcePrefix, randomID)
		rulesetName := fmt.Sprintf("%s-branch-ruleset-%s", testResourcePrefix, randomID)

		workflowFilePath := ".github/workflows/echo.yaml"

		config := fmt.Sprintf(`
resource "github_repository" "test" {
	name = "%s"
	visibility = "private"
	auto_init = true
	ignore_vulnerability_alerts_during_read = true
}

resource "github_repository_file" "workflow_file" {
	repository          = github_repository.test.name
	branch              = "main"
	file                = "%[3]s"
	content             = <<EOT
name: Echo Workflow

on: [pull_request]

jobs:
  echo:
    runs-on: linux
    steps:
      - run: echo \"Hello, world!\"
EOT
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
			allowed_merge_methods = ["merge", "rebase", "squash"]
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
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "rules.0.required_workflows.0.do_not_enforce_on_create", "true"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "rules.0.required_workflows.0.required_workflow.0.path", workflowFilePath),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "rules.0.required_code_scanning.0.required_code_scanning_tool.0.alerts_threshold", "errors"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "rules.0.required_code_scanning.0.required_code_scanning_tool.0.security_alerts_threshold", "high_or_higher"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "rules.0.required_code_scanning.0.required_code_scanning_tool.0.tool", "CodeQL"),
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

		ref_name {
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
	flattenedResult := flattenRules(expandedRules, true)

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
