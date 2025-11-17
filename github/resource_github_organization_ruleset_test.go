package github

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestGithubOrganizationRulesets(t *testing.T) {
	if isEnterprise != "true" {
		t.Skip("Skipping because `ENTERPRISE_ACCOUNT` is not set or set to false")
	}

	if testEnterprise == "" {
		t.Skip("Skipping because `ENTERPRISE_SLUG` is not set")
	}

	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

	t.Run("Creates and updates organization rulesets without errors", func(t *testing.T) {
		config := fmt.Sprintf(`
			resource "github_organization_ruleset" "test" {
				name        = "test-%s"
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
							path          = "path/to/workflow.yaml"
							repository_id = 1234
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
		`, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_organization_ruleset.test",
				"name",
				"test",
			),
			resource.TestCheckResourceAttr(
				"github_organization_ruleset.test",
				"enforcement",
				"active",
			),
			resource.TestCheckResourceAttr(
				"github_organization_ruleset.test",
				"rules.0.required_workflows.0.do_not_enforce_on_create",
				"true",
			),
			resource.TestCheckResourceAttr(
				"github_organization_ruleset.test",
				"rules.0.required_workflows.0.required_workflow.0.path",
				"path/to/workflow.yaml",
			),
			resource.TestCheckResourceAttr(
				"github_organization_ruleset.test",
				"rules.0.required_workflows.0.required_workflow.0.repository_id",
				"1234",
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

		t.Run("with an enterprise account", func(t *testing.T) {
			testCase(t, enterprise)
		})
	})

	t.Run("Updates a ruleset name without error", func(t *testing.T) {
		oldRSName := fmt.Sprintf(`ruleset-%[1]s`, randomID)
		newRSName := fmt.Sprintf(`%[1]s-renamed`, randomID)

		config := fmt.Sprintf(`
			resource "github_organization_ruleset" "test" {
				name        = "%s"
				target      = "branch"
				enforcement = "active"

				rules {
					creation = true
				}
			}
		`, oldRSName)

		checks := map[string]resource.TestCheckFunc{
			"before": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_organization_ruleset.test", "name",
					oldRSName,
				),
			),
			"after": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_organization_ruleset.test", "name",
					newRSName,
				),
			),
		}

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
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
		}

		t.Run("with an enterprise account", func(t *testing.T) {
			testCase(t, enterprise)
		})
	})

	t.Run("Imports rulesets without error", func(t *testing.T) {
		config := fmt.Sprintf(`
			resource "github_organization_ruleset" "test" {
				name        = "test-%s"
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

					required_status_checks {

						required_check {
							context = "ci"
						}

						strict_required_status_checks_policy = true
						do_not_enforce_on_create             = true
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
		`, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet("github_organization_ruleset.test", "name"),
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
					{
						ResourceName:      "github_organization_ruleset.test",
						ImportState:       true,
						ImportStateVerify: true,
					},
				},
			})
		}

		t.Run("with an enterprise account", func(t *testing.T) {
			testCase(t, enterprise)
		})
	})

	t.Run("Creates and updates organization using bypasses", func(t *testing.T) {
		config := fmt.Sprintf(`
			resource "github_organization_ruleset" "test" {
				name        = "test-%s"
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
				}
			}
		`, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_organization_ruleset.test", "bypass_actors.#",
				"3",
			),
			resource.TestCheckResourceAttr(
				"github_organization_ruleset.test", "bypass_actors.0.actor_type",
				"DeployKey",
			),
			resource.TestCheckResourceAttr(
				"github_organization_ruleset.test", "bypass_actors.0.bypass_mode",
				"always",
			),
			resource.TestCheckResourceAttr(
				"github_organization_ruleset.test", "bypass_actors.1.actor_id",
				"5",
			),
			resource.TestCheckResourceAttr(
				"github_organization_ruleset.test", "bypass_actors.1.actor_type",
				"RepositoryRole",
			),
			resource.TestCheckResourceAttr(
				"github_organization_ruleset.test", "bypass_actors.1.bypass_mode",
				"always",
			),
			resource.TestCheckResourceAttr(
				"github_organization_ruleset.test", "bypass_actors.2.actor_id",
				"1",
			),
			resource.TestCheckResourceAttr(
				"github_organization_ruleset.test", "bypass_actors.2.actor_type",
				"OrganizationAdmin",
			),
			resource.TestCheckResourceAttr(
				"github_organization_ruleset.test", "bypass_actors.2.bypass_mode",
				"always",
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

		t.Run("with an enterprise account", func(t *testing.T) {
			testCase(t, enterprise)
		})
	})

	t.Run("Creates organization ruleset with all bypass_modes", func(t *testing.T) {
		config := fmt.Sprintf(`
			resource "github_organization_ruleset" "test" {
				name        = "test-bypass-modes-%s"
				target      = "branch"
				enforcement = "active"

				bypass_actors {
					actor_id    = 1
					actor_type  = "OrganizationAdmin"
					bypass_mode = "always"
				}

				bypass_actors {
					actor_id    = 5
					actor_type  = "RepositoryRole"
					bypass_mode = "pull_request"
				}

				bypass_actors {
					actor_id    = 2
					actor_type  = "RepositoryRole"
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
				"github_organization_ruleset.test", "bypass_actors.#",
				"3",
			),
			resource.TestCheckResourceAttr(
				"github_organization_ruleset.test", "bypass_actors.0.actor_id",
				"1",
			),
			resource.TestCheckResourceAttr(
				"github_organization_ruleset.test", "bypass_actors.0.actor_type",
				"OrganizationAdmin",
			),
			resource.TestCheckResourceAttr(
				"github_organization_ruleset.test", "bypass_actors.0.bypass_mode",
				"always",
			),
			resource.TestCheckResourceAttr(
				"github_organization_ruleset.test", "bypass_actors.1.actor_id",
				"5",
			),
			resource.TestCheckResourceAttr(
				"github_organization_ruleset.test", "bypass_actors.1.actor_type",
				"RepositoryRole",
			),
			resource.TestCheckResourceAttr(
				"github_organization_ruleset.test", "bypass_actors.1.bypass_mode",
				"pull_request",
			),
			resource.TestCheckResourceAttr(
				"github_organization_ruleset.test", "bypass_actors.2.actor_id",
				"2",
			),
			resource.TestCheckResourceAttr(
				"github_organization_ruleset.test", "bypass_actors.2.actor_type",
				"RepositoryRole",
			),
			resource.TestCheckResourceAttr(
				"github_organization_ruleset.test", "bypass_actors.2.bypass_mode",
				"exempt",
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

		t.Run("with an enterprise account", func(t *testing.T) {
			testCase(t, enterprise)
		})
	})

	t.Run("Updates organization ruleset bypass_mode without error", func(t *testing.T) {
		config := fmt.Sprintf(`
			resource "github_organization_ruleset" "test" {
				name        = "test-bypass-update-%s"
				target      = "branch"
				enforcement = "active"

				bypass_actors {
					actor_id    = 1
					actor_type  = "OrganizationAdmin"
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
					"github_organization_ruleset.test", "bypass_actors.0.bypass_mode",
					"always",
				),
			),
			"after": resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"github_organization_ruleset.test", "bypass_actors.0.bypass_mode",
					"exempt",
				),
			),
		}

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
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
		}

		t.Run("with an enterprise account", func(t *testing.T) {
			testCase(t, enterprise)
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
				"max_file_size": float64(10485760), // 10MB
			},
		},
		"max_file_path_length": []any{
			map[string]any{
				"max_file_path_length": 250,
			},
		},
		"file_extension_restriction": []any{
			map[string]any{
				"restricted_file_extensions": []any{".exe", ".bat", ".sh", ".ps1"},
			},
		},
	}

	input := []any{rulesMap}

	// Test expand functionality (organization rulesets use org=true)
	expandedRules := expandRules(input, true)

	if len(expandedRules) != 4 {
		t.Fatalf("Expected 4 expanded rules for organization push ruleset, got %d", len(expandedRules))
	}

	// Verify we have all expected push rule types
	ruleTypes := make(map[string]bool)
	for _, rule := range expandedRules {
		ruleTypes[rule.Type] = true
	}

	expectedPushRules := []string{"file_path_restriction", "max_file_size", "max_file_path_length", "file_extension_restriction"}
	for _, expectedType := range expectedPushRules {
		if !ruleTypes[expectedType] {
			t.Errorf("Expected organization push rule type %s not found in expanded rules", expectedType)
		}
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
	if maxFileSizeRules[0]["max_file_size"] != int64(10485760) {
		t.Errorf("Expected max_file_size to be 10485760, got %v", maxFileSizeRules[0]["max_file_size"])
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
