package github

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestGithubOrganizationRulesets(t *testing.T) {
	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

	t.Run("Creates and updates organization rulesets without errors", func(t *testing.T) {
		resourceName := "test-create-and-update"
		config := fmt.Sprintf(`
			resource "github_repository" "%[1]s" {
				name = "test-%[2]s"
				visibility = "private"
				auto_init = true

        ignore_vulnerability_alerts_during_read = true

			}

			resource "github_repository_file" "%[1]s" {
				repository          = github_repository.%[1]s.name
				branch              = "main"
				file                = ".github/workflows/echo.yaml"
				content             = "name: Echo Workflow\n\non: [pull_request]\n\njobs:\n    echo:\n      runs-on: linux\n      steps:\n        - run: echo \"Hello, world!\"\n"
				commit_message      = "Managed by Terraform"
				commit_author       = "Terraform User"
				commit_email        = "terraform@example.com"
			}

			resource "github_actions_repository_access_level" "%[1]s" {
				repository = github_repository.%[1]s.name
				access_level = "organization"
			}

			resource "github_organization_ruleset" "%[1]s" {
				name        = "test-%[2]s"
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
					creation = true

					update = true

					deletion                = true
					required_linear_history = true

					required_signatures = false

					pull_request {
						allowed_merge_methods             = ["merge", "squash", "rebase"]
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
							path          = ".github/workflows/echo.yaml"
							repository_id = github_repository.%[1]s.repo_id
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
        depends_on = [github_repository_file.%[1]s]
			}
		`, resourceName, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				fmt.Sprintf("github_organization_ruleset.%s", resourceName),
				"name",
				fmt.Sprintf("test-%s", randomID),
			),
			resource.TestCheckResourceAttr(
				fmt.Sprintf("github_organization_ruleset.%s", resourceName),
				"enforcement",
				"active",
			),
			resource.TestCheckResourceAttr(
				fmt.Sprintf("github_organization_ruleset.%s", resourceName),
				"rules.0.required_workflows.0.do_not_enforce_on_create",
				"true",
			),
			resource.TestCheckResourceAttr(
				fmt.Sprintf("github_organization_ruleset.%s", resourceName),
				"rules.0.required_workflows.0.required_workflow.0.path",
				".github/workflows/echo.yaml",
			),
			resource.TestCheckResourceAttr(
				fmt.Sprintf("github_organization_ruleset.%s", resourceName),
				"rules.0.required_code_scanning.0.required_code_scanning_tool.0.alerts_threshold",
				"errors",
			),
			resource.TestCheckResourceAttr(
				fmt.Sprintf("github_organization_ruleset.%s", resourceName),
				"rules.0.required_code_scanning.0.required_code_scanning_tool.0.security_alerts_threshold",
				"high_or_higher",
			),
			resource.TestCheckResourceAttr(
				fmt.Sprintf("github_organization_ruleset.%s", resourceName),
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
		resourceName := "test-import-rulesets"
		config := fmt.Sprintf(`
			resource "github_organization_ruleset" "%s" {
				name        = "test-%s"
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
					creation = true

					update = true

					deletion                = true
					required_linear_history = true

					required_signatures = false

					pull_request {
            allowed_merge_methods = ["merge", "squash", "rebase"]
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
		`, resourceName, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet(fmt.Sprintf("github_organization_ruleset.%s", resourceName), "name"),
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
						ResourceName:      fmt.Sprintf("github_organization_ruleset.%s", resourceName),
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
		resourceName := "test-creates-and-updates-using-bypasses"
		config := fmt.Sprintf(`
			resource "github_organization_ruleset" "%s" {
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
					repository_name {
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
            allowed_merge_methods = ["merge", "squash", "rebase"]
						required_approving_review_count   = 2
						required_review_thread_resolution = true
						require_code_owner_review         = true
						dismiss_stale_reviews_on_push     = true
						require_last_push_approval        = true
					}
				}
			}
		`, resourceName, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				fmt.Sprintf("github_organization_ruleset.%s", resourceName), "bypass_actors.#",
				"3",
			),
			resource.TestCheckResourceAttr(
				fmt.Sprintf("github_organization_ruleset.%s", resourceName), "bypass_actors.0.actor_type",
				"DeployKey",
			),
			resource.TestCheckResourceAttr(
				fmt.Sprintf("github_organization_ruleset.%s", resourceName), "bypass_actors.0.bypass_mode",
				"always",
			),
			resource.TestCheckResourceAttr(
				fmt.Sprintf("github_organization_ruleset.%s", resourceName), "bypass_actors.2.actor_id",
				"5",
			),
			resource.TestCheckResourceAttr(
				fmt.Sprintf("github_organization_ruleset.%s", resourceName), "bypass_actors.2.actor_type",
				"RepositoryRole",
			),
			resource.TestCheckResourceAttr(
				fmt.Sprintf("github_organization_ruleset.%s", resourceName), "bypass_actors.2.bypass_mode",
				"always",
			),
			resource.TestCheckResourceAttr(
				fmt.Sprintf("github_organization_ruleset.%s", resourceName), "bypass_actors.1.actor_id",
				"1",
			),
			resource.TestCheckResourceAttr(
				fmt.Sprintf("github_organization_ruleset.%s", resourceName), "bypass_actors.1.actor_type",
				"OrganizationAdmin",
			),
			resource.TestCheckResourceAttr(
				fmt.Sprintf("github_organization_ruleset.%s", resourceName), "bypass_actors.1.bypass_mode",
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
		resourceName := "test-create-with-bypass-modes"
		config := fmt.Sprintf(`
			resource "github_organization_ruleset" "%s" {
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
					repository_name {
						include = ["~ALL"]
						exclude = []
					}
				}

				rules {
					creation = true
				}
			}
		`, resourceName, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				fmt.Sprintf("github_organization_ruleset.%s", resourceName), "bypass_actors.#",
				"3",
			),
			resource.TestCheckResourceAttr(
				fmt.Sprintf("github_organization_ruleset.%s", resourceName), "bypass_actors.0.actor_id",
				"1",
			),
			resource.TestCheckResourceAttr(
				fmt.Sprintf("github_organization_ruleset.%s", resourceName), "bypass_actors.0.actor_type",
				"OrganizationAdmin",
			),
			resource.TestCheckResourceAttr(
				fmt.Sprintf("github_organization_ruleset.%s", resourceName), "bypass_actors.0.bypass_mode",
				"always",
			),
			resource.TestCheckResourceAttr(
				fmt.Sprintf("github_organization_ruleset.%s", resourceName), "bypass_actors.2.actor_id",
				"5",
			),
			resource.TestCheckResourceAttr(
				fmt.Sprintf("github_organization_ruleset.%s", resourceName), "bypass_actors.2.actor_type",
				"RepositoryRole",
			),
			resource.TestCheckResourceAttr(
				fmt.Sprintf("github_organization_ruleset.%s", resourceName), "bypass_actors.2.bypass_mode",
				"pull_request",
			),
			resource.TestCheckResourceAttr(
				fmt.Sprintf("github_organization_ruleset.%s", resourceName), "bypass_actors.1.actor_id",
				"2",
			),
			resource.TestCheckResourceAttr(
				fmt.Sprintf("github_organization_ruleset.%s", resourceName), "bypass_actors.1.actor_type",
				"RepositoryRole",
			),
			resource.TestCheckResourceAttr(
				fmt.Sprintf("github_organization_ruleset.%s", resourceName), "bypass_actors.1.bypass_mode",
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

	t.Run("Validates branch target requires `ref_name` condition", func(t *testing.T) {
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

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config:      config,
						ExpectError: regexp.MustCompile("ref_name must be set for branch target"),
					},
				},
			})
		}

		t.Run("with an enterprise account", func(t *testing.T) {
			testCase(t, enterprise)
		})

		t.Run("with an organization account", func(t *testing.T) {
			t.Skip("organization account not supported for this operation, since it needs a paid Team plan.")
		})
	})

	t.Run("Validates tag target requires `ref_name` condition", func(t *testing.T) {
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

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config:      config,
						ExpectError: regexp.MustCompile("ref_name must be set for tag target"),
					},
				},
			})
		}

		t.Run("with an enterprise account", func(t *testing.T) {
			testCase(t, enterprise)
		})

		t.Run("with an organization account", func(t *testing.T) {
			t.Skip("organization account not supported for this operation, since it needs a paid Team plan.")
		})
	})

	t.Run("Validates push target rejects ref_name", func(t *testing.T) {
		config := fmt.Sprintf(`
			resource "github_organization_ruleset" "test" {
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
					creation = true
				}
			}
		`, randomID)

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config:      config,
						ExpectError: regexp.MustCompile("ref_name must not be set for push target"),
					},
				},
			})
		}

		t.Run("with an enterprise account", func(t *testing.T) {
			testCase(t, enterprise)
		})

		t.Run("with an organization account", func(t *testing.T) {
			t.Skip("organization account not supported for this operation, since it needs a paid Team plan.")
		})
	})

	t.Run("Validates repository target rejects ref_name", func(t *testing.T) {
		config := fmt.Sprintf(`
			resource "github_organization_ruleset" "test" {
				name        = "test-repo-with-ref-%s"
				target      = "repository"
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
					creation = true
				}
			}
		`, randomID)

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config:      config,
						ExpectError: regexp.MustCompile("ref_name must not be set for repository target"),
					},
				},
			})
		}

		t.Run("with an enterprise account", func(t *testing.T) {
			testCase(t, enterprise)
		})

		t.Run("with an organization account", func(t *testing.T) {
			t.Skip("organization account not supported for this operation, since it needs a paid Team plan.")
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
		t.Fatal("Expected expanded rules to not be nil")
		return
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
