package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccGithubEnterpriseRuleset_basic(t *testing.T) {
	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
	rulesetName := fmt.Sprintf("%s-enterprise-basic-%s", testResourcePrefix, randomID)

	rulesetHCL := `
		resource "github_enterprise_ruleset" "test" {
			enterprise_slug = "%s"
			name            = "%s"
			target          = "branch"
			enforcement     = "active"

			conditions {
				organization_name {
					include = ["~ALL"]
					exclude = []
				}

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
	config := fmt.Sprintf(rulesetHCL, testAccConf.enterpriseSlug, rulesetName)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { skipUnlessEnterprise(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("github_enterprise_ruleset.test", "enterprise_slug", testAccConf.enterpriseSlug),
					resource.TestCheckResourceAttr("github_enterprise_ruleset.test", "name", rulesetName),
					resource.TestCheckResourceAttr("github_enterprise_ruleset.test", "target", "branch"),
					resource.TestCheckResourceAttr("github_enterprise_ruleset.test", "enforcement", "active"),
				),
			},
		},
	})
}

func TestAccGithubEnterpriseRuleset_branch_rules(t *testing.T) {
	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
	rulesetName := fmt.Sprintf("%s-enterprise-branch-%s", testResourcePrefix, randomID)

	config := fmt.Sprintf(`
resource "github_enterprise_ruleset" "test" {
	enterprise_slug = "%s"
	name            = "%s"
	target          = "branch"
	enforcement     = "active"

	bypass_actors {
		actor_type  = "DeployKey"
		bypass_mode = "always"
	}

	bypass_actors {
		actor_id    = 1
		actor_type  = "OrganizationAdmin"
		bypass_mode = "always"
	}

	conditions {
		organization_name {
			include = ["~ALL"]
			exclude = []
		}

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
		creation                = true
		update                  = true
		deletion                = true
		required_linear_history = true
		required_signatures     = false

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

		required_code_scanning {
			required_code_scanning_tool {
				alerts_threshold          = "errors"
				security_alerts_threshold = "high_or_higher"
				tool                      = "CodeQL"
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
`, testAccConf.enterpriseSlug, rulesetName)

	check := resource.ComposeTestCheckFunc(
		resource.TestCheckResourceAttr("github_enterprise_ruleset.test", "name", rulesetName),
		resource.TestCheckResourceAttr("github_enterprise_ruleset.test", "target", "branch"),
		resource.TestCheckResourceAttr("github_enterprise_ruleset.test", "enforcement", "active"),
		resource.TestCheckResourceAttr("github_enterprise_ruleset.test", "bypass_actors.#", "2"),
		resource.TestCheckResourceAttr("github_enterprise_ruleset.test", "bypass_actors.0.actor_type", "DeployKey"),
		resource.TestCheckResourceAttr("github_enterprise_ruleset.test", "bypass_actors.0.bypass_mode", "always"),
		resource.TestCheckResourceAttr("github_enterprise_ruleset.test", "bypass_actors.1.actor_id", "1"),
		resource.TestCheckResourceAttr("github_enterprise_ruleset.test", "bypass_actors.1.actor_type", "OrganizationAdmin"),
		resource.TestCheckResourceAttr("github_enterprise_ruleset.test", "bypass_actors.1.bypass_mode", "always"),
		resource.TestCheckResourceAttr("github_enterprise_ruleset.test", "rules.0.required_code_scanning.0.required_code_scanning_tool.0.alerts_threshold", "errors"),
		resource.TestCheckResourceAttr("github_enterprise_ruleset.test", "rules.0.required_code_scanning.0.required_code_scanning_tool.0.security_alerts_threshold", "high_or_higher"),
		resource.TestCheckResourceAttr("github_enterprise_ruleset.test", "rules.0.required_code_scanning.0.required_code_scanning_tool.0.tool", "CodeQL"),
		resource.TestCheckResourceAttr("github_enterprise_ruleset.test", "rules.0.copilot_code_review.0.review_on_push", "true"),
		resource.TestCheckResourceAttr("github_enterprise_ruleset.test", "rules.0.copilot_code_review.0.review_draft_pull_requests", "false"),
	)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { skipUnlessEnterprise(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check:  check,
			},
		},
	})
}

func TestAccGithubEnterpriseRuleset_required_workflows(t *testing.T) {
	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
	rulesetName := fmt.Sprintf("%s-enterprise-wf-ruleset-%s", testResourcePrefix, randomID)
	workflowFilePath := ".github/workflows/echo.yaml"

	config := fmt.Sprintf(`
resource "github_enterprise_ruleset" "test" {
	enterprise_slug = "%s"
	name            = "%s"
	target          = "branch"
	enforcement     = "active"

	conditions {
		organization_name {
			include = ["~ALL"]
			exclude = []
		}

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
		required_workflows {
			do_not_enforce_on_create = true
			required_workflow {
				path          = "%s"
				repository_id = 1234567
				ref           = "main"
			}
		}
	}
}
`, testAccConf.enterpriseSlug, rulesetName, workflowFilePath)

	check := resource.ComposeTestCheckFunc(
		resource.TestCheckResourceAttr("github_enterprise_ruleset.test", "name", rulesetName),
		resource.TestCheckResourceAttr("github_enterprise_ruleset.test", "target", "branch"),
		resource.TestCheckResourceAttr("github_enterprise_ruleset.test", "enforcement", "active"),
		resource.TestCheckResourceAttr("github_enterprise_ruleset.test", "rules.0.required_workflows.0.do_not_enforce_on_create", "true"),
		resource.TestCheckResourceAttr("github_enterprise_ruleset.test", "rules.0.required_workflows.0.required_workflow.0.path", workflowFilePath),
	)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { skipUnlessEnterprise(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check:  check,
			},
		},
	})
}

func TestAccGithubEnterpriseRuleset_tag(t *testing.T) {
	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
	rulesetName := fmt.Sprintf("%s-enterprise-tag-%s", testResourcePrefix, randomID)

	config := fmt.Sprintf(`
resource "github_enterprise_ruleset" "test" {
	enterprise_slug = "%s"
	name            = "%s"
	target          = "tag"
	enforcement     = "active"

	conditions {
		organization_name {
			include = ["~ALL"]
			exclude = []
		}

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
		creation                = false
		deletion                = false
		required_linear_history = false
	}
}
`, testAccConf.enterpriseSlug, rulesetName)

	check := resource.ComposeTestCheckFunc(
		resource.TestCheckResourceAttr("github_enterprise_ruleset.test", "name", rulesetName),
		resource.TestCheckResourceAttr("github_enterprise_ruleset.test", "target", "tag"),
		resource.TestCheckResourceAttr("github_enterprise_ruleset.test", "enforcement", "active"),
	)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { skipUnlessEnterprise(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check:  check,
			},
		},
	})
}

func TestAccGithubEnterpriseRuleset_push(t *testing.T) {
	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
	rulesetName := fmt.Sprintf("%s-enterprise-push-%s", testResourcePrefix, randomID)

	config := fmt.Sprintf(`
resource "github_enterprise_ruleset" "test" {
	enterprise_slug = "%s"
	name            = "%s"
	target          = "push"
	enforcement     = "active"

	bypass_actors {
		actor_type  = "DeployKey"
		bypass_mode = "always"
	}

	bypass_actors {
		actor_id    = 1
		actor_type  = "OrganizationAdmin"
		bypass_mode = "always"
	}

	conditions {
		organization_name {
			include = ["~ALL"]
			exclude = []
		}

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
`, testAccConf.enterpriseSlug, rulesetName)

	check := resource.ComposeTestCheckFunc(
		resource.TestCheckResourceAttr("github_enterprise_ruleset.test", "name", rulesetName),
		resource.TestCheckResourceAttr("github_enterprise_ruleset.test", "target", "push"),
		resource.TestCheckResourceAttr("github_enterprise_ruleset.test", "enforcement", "active"),
		resource.TestCheckResourceAttr("github_enterprise_ruleset.test", "bypass_actors.#", "2"),
		resource.TestCheckResourceAttr("github_enterprise_ruleset.test", "bypass_actors.0.actor_type", "DeployKey"),
		resource.TestCheckResourceAttr("github_enterprise_ruleset.test", "bypass_actors.0.bypass_mode", "always"),
		resource.TestCheckResourceAttr("github_enterprise_ruleset.test", "bypass_actors.1.actor_id", "1"),
		resource.TestCheckResourceAttr("github_enterprise_ruleset.test", "bypass_actors.1.actor_type", "OrganizationAdmin"),
		resource.TestCheckResourceAttr("github_enterprise_ruleset.test", "bypass_actors.1.bypass_mode", "always"),
		resource.TestCheckResourceAttr("github_enterprise_ruleset.test", "rules.0.file_path_restriction.0.restricted_file_paths.0", "test.txt"),
		resource.TestCheckResourceAttr("github_enterprise_ruleset.test", "rules.0.max_file_size.0.max_file_size", "99"),
		resource.TestCheckResourceAttr("github_enterprise_ruleset.test", "rules.0.file_extension_restriction.0.restricted_file_extensions.0", "*.zip"),
	)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { skipUnlessEnterprise(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check:  check,
			},
		},
	})
}

func TestAccGithubEnterpriseRuleset_update_name(t *testing.T) {
	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
	name := fmt.Sprintf("test-enterprise-ruleset-%s", randomID)
	nameUpdated := fmt.Sprintf("test-enterprise-ruleset-updated-%s", randomID)

	configs := map[string]string{
		"before": fmt.Sprintf(`
resource "github_enterprise_ruleset" "test" {
	enterprise_slug = "%s"
	name            = "%s"
	target          = "branch"
	enforcement     = "active"

	conditions {
		organization_name {
			include = ["~ALL"]
			exclude = []
		}

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
`, testAccConf.enterpriseSlug, name),
		"after": fmt.Sprintf(`
resource "github_enterprise_ruleset" "test" {
	enterprise_slug = "%s"
	name            = "%s"
	target          = "branch"
	enforcement     = "active"

	conditions {
		organization_name {
			include = ["~ALL"]
			exclude = []
		}

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
`, testAccConf.enterpriseSlug, nameUpdated),
	}

	checks := map[string]resource.TestCheckFunc{
		"before": resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("github_enterprise_ruleset.test", "name", name),
		),
		"after": resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("github_enterprise_ruleset.test", "name", nameUpdated),
		),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { skipUnlessEnterprise(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: configs["before"],
				Check:  checks["before"],
			},
			{
				Config: configs["after"],
				Check:  checks["after"],
			},
		},
	})
}

func TestAccGithubEnterpriseRuleset_update_bypass_actors(t *testing.T) {
	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
	rulesetName := fmt.Sprintf("%s-enterprise-bypass-%s", testResourcePrefix, randomID)

	configs := map[string]string{
		"with_actors": fmt.Sprintf(`
resource "github_enterprise_ruleset" "test" {
	enterprise_slug = "%s"
	name            = "%s"
	target          = "branch"
	enforcement     = "active"

	bypass_actors {
		actor_type  = "DeployKey"
		bypass_mode = "always"
	}

	bypass_actors {
		actor_id    = 1
		actor_type  = "OrganizationAdmin"
		bypass_mode = "always"
	}

	conditions {
		organization_name {
			include = ["~ALL"]
			exclude = []
		}

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
`, testAccConf.enterpriseSlug, rulesetName),
		"without_actors": fmt.Sprintf(`
resource "github_enterprise_ruleset" "test" {
	enterprise_slug = "%s"
	name            = "%s"
	target          = "branch"
	enforcement     = "active"

	conditions {
		organization_name {
			include = ["~ALL"]
			exclude = []
		}

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
`, testAccConf.enterpriseSlug, rulesetName),
	}

	checks := map[string]resource.TestCheckFunc{
		"with_actors": resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("github_enterprise_ruleset.test", "bypass_actors.#", "2"),
		),
		"without_actors": resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("github_enterprise_ruleset.test", "bypass_actors.#", "0"),
		),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { skipUnlessEnterprise(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: configs["with_actors"],
				Check:  checks["with_actors"],
			},
			{
				Config: configs["without_actors"],
				Check:  checks["without_actors"],
			},
		},
	})
}

func TestAccGithubEnterpriseRuleset_update_bypass_mode(t *testing.T) {
	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
	rulesetName := fmt.Sprintf("test-enterprise-bypass-mode-%s", randomID)

	bypassMode := "always"
	bypassModeUpdated := "exempt"

	configs := map[string]string{
		"before": fmt.Sprintf(`
resource "github_enterprise_ruleset" "test" {
	enterprise_slug = "%s"
	name            = "%s"
	target          = "branch"
	enforcement     = "active"

	bypass_actors {
		actor_id    = 1
		actor_type  = "OrganizationAdmin"
		bypass_mode = "%s"
	}

	conditions {
		organization_name {
			include = ["~ALL"]
			exclude = []
		}

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
`, testAccConf.enterpriseSlug, rulesetName, bypassMode),
		"after": fmt.Sprintf(`
resource "github_enterprise_ruleset" "test" {
	enterprise_slug = "%s"
	name            = "%s"
	target          = "branch"
	enforcement     = "active"

	bypass_actors {
		actor_id    = 1
		actor_type  = "OrganizationAdmin"
		bypass_mode = "%s"
	}

	conditions {
		organization_name {
			include = ["~ALL"]
			exclude = []
		}

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
`, testAccConf.enterpriseSlug, rulesetName, bypassModeUpdated),
	}

	checks := map[string]resource.TestCheckFunc{
		"before": resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("github_enterprise_ruleset.test", "bypass_actors.0.bypass_mode", bypassMode),
		),
		"after": resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr("github_enterprise_ruleset.test", "bypass_actors.0.bypass_mode", bypassModeUpdated),
		),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { skipUnlessEnterprise(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: configs["before"],
				Check:  checks["before"],
			},
			{
				Config: configs["after"],
				Check:  checks["after"],
			},
		},
	})
}

func TestAccGithubEnterpriseRuleset_repository_targeting(t *testing.T) {
	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
	rulesetName := fmt.Sprintf("%s-enterprise-repo-targeting-%s", testResourcePrefix, randomID)

	config := fmt.Sprintf(`
resource "github_enterprise_ruleset" "test" {
	enterprise_slug = "%s"
	name            = "%s"
	target          = "branch"
	enforcement     = "active"

	conditions {
		organization_name {
			include = ["~ALL"]
			exclude = []
		}

		repository_name {
			include   = ["prod-*", "production-*"]
			exclude   = ["prod-test*"]
			protected = true
		}

		ref_name {
			include = ["refs/heads/main"]
			exclude = []
		}
	}

	rules {
		creation = false
		deletion = false
	}
}
`, testAccConf.enterpriseSlug, rulesetName)

	check := resource.ComposeTestCheckFunc(
		resource.TestCheckResourceAttr("github_enterprise_ruleset.test", "name", rulesetName),
		resource.TestCheckResourceAttr("github_enterprise_ruleset.test", "conditions.0.repository_name.0.include.#", "2"),
		resource.TestCheckResourceAttr("github_enterprise_ruleset.test", "conditions.0.repository_name.0.include.0", "prod-*"),
		resource.TestCheckResourceAttr("github_enterprise_ruleset.test", "conditions.0.repository_name.0.include.1", "production-*"),
		resource.TestCheckResourceAttr("github_enterprise_ruleset.test", "conditions.0.repository_name.0.exclude.0", "prod-test*"),
		resource.TestCheckResourceAttr("github_enterprise_ruleset.test", "conditions.0.repository_name.0.protected", "true"),
	)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { skipUnlessEnterprise(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check:  check,
			},
		},
	})
}

func TestAccGithubEnterpriseRuleset_organizationID(t *testing.T) {
	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
	rulesetName := fmt.Sprintf("%s-enterprise-org-id-%s", testResourcePrefix, randomID)

	config := fmt.Sprintf(`
resource "github_enterprise_ruleset" "test" {
	enterprise_slug = "%s"
	name            = "%s"
	target          = "branch"
	enforcement     = "active"

	conditions {
		organization_id = [2284107]

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
`, testAccConf.enterpriseSlug, rulesetName)

	check := resource.ComposeTestCheckFunc(
		resource.TestCheckResourceAttr("github_enterprise_ruleset.test", "name", rulesetName),
		resource.TestCheckResourceAttr("github_enterprise_ruleset.test", "target", "branch"),
		resource.TestCheckResourceAttr("github_enterprise_ruleset.test", "enforcement", "active"),
		resource.TestCheckResourceAttr("github_enterprise_ruleset.test", "conditions.0.organization_id.#", "1"),
		resource.TestCheckResourceAttr("github_enterprise_ruleset.test", "conditions.0.organization_id.0", "2284107"),
	)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { skipUnlessEnterprise(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check:  check,
			},
		},
	})
}


func TestAccGithubEnterpriseRuleset_import(t *testing.T) {
	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
	rulesetName := fmt.Sprintf("%s-enterprise-import-%s", testResourcePrefix, randomID)

	rulesetHCL := `
		resource "github_enterprise_ruleset" "test" {
			enterprise_slug = "%s"
			name            = "%s"
			target          = "branch"
			enforcement     = "active"

			conditions {
				organization_name {
					include = ["~ALL"]
					exclude = []
				}

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
	config := fmt.Sprintf(rulesetHCL, testAccConf.enterpriseSlug, rulesetName)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { skipUnlessEnterprise(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: config,
			},
			{
				ResourceName:            "github_enterprise_ruleset.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateIdFunc:       importEnterpriseRulesetByResourcePath("github_enterprise_ruleset.test"),
				ImportStateVerifyIgnore: []string{"etag"},
			},
		},
	})
}

func importEnterpriseRulesetByResourcePath(rulesetLogicalName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		ruleset := s.RootModule().Resources[rulesetLogicalName]
		if ruleset == nil {
			return "", fmt.Errorf("Cannot find %s in terraform state", rulesetLogicalName)
		}

		rulesetID := ruleset.Primary.ID
		if rulesetID == "" {
			return "", fmt.Errorf("ruleset %s does not have an id in terraform state", rulesetLogicalName)
		}

		enterpriseSlug := ruleset.Primary.Attributes["enterprise_slug"]
		if enterpriseSlug == "" {
			return "", fmt.Errorf("ruleset %s does not have enterprise_slug in terraform state", rulesetLogicalName)
		}

		return fmt.Sprintf("%s:%s", enterpriseSlug, rulesetID), nil
	}
}
