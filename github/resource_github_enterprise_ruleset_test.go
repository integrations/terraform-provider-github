package github

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
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
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: config,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue("github_enterprise_ruleset.test", tfjsonpath.New("enterprise_slug"), knownvalue.StringExact(testAccConf.enterpriseSlug)),
					statecheck.ExpectKnownValue("github_enterprise_ruleset.test", tfjsonpath.New("name"), knownvalue.StringExact(rulesetName)),
					statecheck.ExpectKnownValue("github_enterprise_ruleset.test", tfjsonpath.New("target"), knownvalue.StringExact("branch")),
					statecheck.ExpectKnownValue("github_enterprise_ruleset.test", tfjsonpath.New("enforcement"), knownvalue.StringExact("active")),
				},
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

	checks := []statecheck.StateCheck{
		statecheck.ExpectKnownValue("github_enterprise_ruleset.test", tfjsonpath.New("name"), knownvalue.StringExact(rulesetName)),
		statecheck.ExpectKnownValue("github_enterprise_ruleset.test", tfjsonpath.New("target"), knownvalue.StringExact("branch")),
		statecheck.ExpectKnownValue("github_enterprise_ruleset.test", tfjsonpath.New("enforcement"), knownvalue.StringExact("active")),
		statecheck.ExpectKnownValue("github_enterprise_ruleset.test", tfjsonpath.New("bypass_actors"), knownvalue.ListSizeExact(2)),
		statecheck.ExpectKnownValue("github_enterprise_ruleset.test", tfjsonpath.New("bypass_actors").AtSliceIndex(0).AtMapKey("actor_type"), knownvalue.StringExact("DeployKey")),
		statecheck.ExpectKnownValue("github_enterprise_ruleset.test", tfjsonpath.New("bypass_actors").AtSliceIndex(0).AtMapKey("bypass_mode"), knownvalue.StringExact("always")),
		statecheck.ExpectKnownValue("github_enterprise_ruleset.test", tfjsonpath.New("bypass_actors").AtSliceIndex(1).AtMapKey("actor_id"), knownvalue.Int64Exact(1)),
		statecheck.ExpectKnownValue("github_enterprise_ruleset.test", tfjsonpath.New("bypass_actors").AtSliceIndex(1).AtMapKey("actor_type"), knownvalue.StringExact("OrganizationAdmin")),
		statecheck.ExpectKnownValue("github_enterprise_ruleset.test", tfjsonpath.New("bypass_actors").AtSliceIndex(1).AtMapKey("bypass_mode"), knownvalue.StringExact("always")),
		statecheck.ExpectKnownValue("github_enterprise_ruleset.test", tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("required_code_scanning").AtSliceIndex(0).AtMapKey("required_code_scanning_tool").AtSliceIndex(0).AtMapKey("alerts_threshold"), knownvalue.StringExact("errors")),
		statecheck.ExpectKnownValue("github_enterprise_ruleset.test", tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("required_code_scanning").AtSliceIndex(0).AtMapKey("required_code_scanning_tool").AtSliceIndex(0).AtMapKey("security_alerts_threshold"), knownvalue.StringExact("high_or_higher")),
		statecheck.ExpectKnownValue("github_enterprise_ruleset.test", tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("required_code_scanning").AtSliceIndex(0).AtMapKey("required_code_scanning_tool").AtSliceIndex(0).AtMapKey("tool"), knownvalue.StringExact("CodeQL")),
		statecheck.ExpectKnownValue("github_enterprise_ruleset.test", tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("copilot_code_review").AtSliceIndex(0).AtMapKey("review_on_push"), knownvalue.Bool(true)),
		statecheck.ExpectKnownValue("github_enterprise_ruleset.test", tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("copilot_code_review").AtSliceIndex(0).AtMapKey("review_draft_pull_requests"), knownvalue.Bool(false)),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { skipUnlessEnterprise(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config:            config,
				ConfigStateChecks: checks,
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

	checks := []statecheck.StateCheck{
		statecheck.ExpectKnownValue("github_enterprise_ruleset.test", tfjsonpath.New("name"), knownvalue.StringExact(rulesetName)),
		statecheck.ExpectKnownValue("github_enterprise_ruleset.test", tfjsonpath.New("target"), knownvalue.StringExact("branch")),
		statecheck.ExpectKnownValue("github_enterprise_ruleset.test", tfjsonpath.New("enforcement"), knownvalue.StringExact("active")),
		statecheck.ExpectKnownValue("github_enterprise_ruleset.test", tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("required_workflows").AtSliceIndex(0).AtMapKey("do_not_enforce_on_create"), knownvalue.Bool(true)),
		statecheck.ExpectKnownValue("github_enterprise_ruleset.test", tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("required_workflows").AtSliceIndex(0).AtMapKey("required_workflow").AtSliceIndex(0).AtMapKey("path"), knownvalue.StringExact(workflowFilePath)),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { skipUnlessEnterprise(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config:            config,
				ConfigStateChecks: checks,
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

	checks := []statecheck.StateCheck{
		statecheck.ExpectKnownValue("github_enterprise_ruleset.test", tfjsonpath.New("name"), knownvalue.StringExact(rulesetName)),
		statecheck.ExpectKnownValue("github_enterprise_ruleset.test", tfjsonpath.New("target"), knownvalue.StringExact("tag")),
		statecheck.ExpectKnownValue("github_enterprise_ruleset.test", tfjsonpath.New("enforcement"), knownvalue.StringExact("active")),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { skipUnlessEnterprise(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config:            config,
				ConfigStateChecks: checks,
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

	checks := []statecheck.StateCheck{
		statecheck.ExpectKnownValue("github_enterprise_ruleset.test", tfjsonpath.New("name"), knownvalue.StringExact(rulesetName)),
		statecheck.ExpectKnownValue("github_enterprise_ruleset.test", tfjsonpath.New("target"), knownvalue.StringExact("push")),
		statecheck.ExpectKnownValue("github_enterprise_ruleset.test", tfjsonpath.New("enforcement"), knownvalue.StringExact("active")),
		statecheck.ExpectKnownValue("github_enterprise_ruleset.test", tfjsonpath.New("bypass_actors"), knownvalue.ListSizeExact(2)),
		statecheck.ExpectKnownValue("github_enterprise_ruleset.test", tfjsonpath.New("bypass_actors").AtSliceIndex(0).AtMapKey("actor_type"), knownvalue.StringExact("DeployKey")),
		statecheck.ExpectKnownValue("github_enterprise_ruleset.test", tfjsonpath.New("bypass_actors").AtSliceIndex(0).AtMapKey("bypass_mode"), knownvalue.StringExact("always")),
		statecheck.ExpectKnownValue("github_enterprise_ruleset.test", tfjsonpath.New("bypass_actors").AtSliceIndex(1).AtMapKey("actor_id"), knownvalue.Int64Exact(1)),
		statecheck.ExpectKnownValue("github_enterprise_ruleset.test", tfjsonpath.New("bypass_actors").AtSliceIndex(1).AtMapKey("actor_type"), knownvalue.StringExact("OrganizationAdmin")),
		statecheck.ExpectKnownValue("github_enterprise_ruleset.test", tfjsonpath.New("bypass_actors").AtSliceIndex(1).AtMapKey("bypass_mode"), knownvalue.StringExact("always")),
		statecheck.ExpectKnownValue("github_enterprise_ruleset.test", tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("file_path_restriction").AtSliceIndex(0).AtMapKey("restricted_file_paths").AtSliceIndex(0), knownvalue.StringExact("test.txt")),
		statecheck.ExpectKnownValue("github_enterprise_ruleset.test", tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("max_file_size").AtSliceIndex(0).AtMapKey("max_file_size"), knownvalue.Int64Exact(99)),
		statecheck.ExpectKnownValue("github_enterprise_ruleset.test", tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("file_extension_restriction").AtSliceIndex(0).AtMapKey("restricted_file_extensions").AtSliceIndex(0), knownvalue.StringExact("*.zip")),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { skipUnlessEnterprise(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config:            config,
				ConfigStateChecks: checks,
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

	checks := map[string][]statecheck.StateCheck{
		"before": {
			statecheck.ExpectKnownValue("github_enterprise_ruleset.test", tfjsonpath.New("name"), knownvalue.StringExact(name)),
		},
		"after": {
			statecheck.ExpectKnownValue("github_enterprise_ruleset.test", tfjsonpath.New("name"), knownvalue.StringExact(nameUpdated)),
		},
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { skipUnlessEnterprise(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config:            configs["before"],
				ConfigStateChecks: checks["before"],
			},
			{
				Config:            configs["after"],
				ConfigStateChecks: checks["after"],
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

	checks := map[string][]statecheck.StateCheck{
		"with_actors": {
			statecheck.ExpectKnownValue("github_enterprise_ruleset.test", tfjsonpath.New("bypass_actors"), knownvalue.ListSizeExact(2)),
		},
		"without_actors": {
			statecheck.ExpectKnownValue("github_enterprise_ruleset.test", tfjsonpath.New("bypass_actors"), knownvalue.ListSizeExact(0)),
		},
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { skipUnlessEnterprise(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config:            configs["with_actors"],
				ConfigStateChecks: checks["with_actors"],
			},
			{
				Config:            configs["without_actors"],
				ConfigStateChecks: checks["without_actors"],
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

	checks := map[string][]statecheck.StateCheck{
		"before": {
			statecheck.ExpectKnownValue("github_enterprise_ruleset.test", tfjsonpath.New("bypass_actors").AtSliceIndex(0).AtMapKey("bypass_mode"), knownvalue.StringExact(bypassMode)),
		},
		"after": {
			statecheck.ExpectKnownValue("github_enterprise_ruleset.test", tfjsonpath.New("bypass_actors").AtSliceIndex(0).AtMapKey("bypass_mode"), knownvalue.StringExact(bypassModeUpdated)),
		},
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { skipUnlessEnterprise(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config:            configs["before"],
				ConfigStateChecks: checks["before"],
			},
			{
				Config:            configs["after"],
				ConfigStateChecks: checks["after"],
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

	checks := []statecheck.StateCheck{
		statecheck.ExpectKnownValue("github_enterprise_ruleset.test", tfjsonpath.New("name"), knownvalue.StringExact(rulesetName)),
		statecheck.ExpectKnownValue("github_enterprise_ruleset.test", tfjsonpath.New("conditions").AtSliceIndex(0).AtMapKey("repository_name").AtSliceIndex(0).AtMapKey("include"), knownvalue.ListSizeExact(2)),
		statecheck.ExpectKnownValue("github_enterprise_ruleset.test", tfjsonpath.New("conditions").AtSliceIndex(0).AtMapKey("repository_name").AtSliceIndex(0).AtMapKey("include").AtSliceIndex(0), knownvalue.StringExact("prod-*")),
		statecheck.ExpectKnownValue("github_enterprise_ruleset.test", tfjsonpath.New("conditions").AtSliceIndex(0).AtMapKey("repository_name").AtSliceIndex(0).AtMapKey("include").AtSliceIndex(1), knownvalue.StringExact("production-*")),
		statecheck.ExpectKnownValue("github_enterprise_ruleset.test", tfjsonpath.New("conditions").AtSliceIndex(0).AtMapKey("repository_name").AtSliceIndex(0).AtMapKey("exclude").AtSliceIndex(0), knownvalue.StringExact("prod-test*")),
		statecheck.ExpectKnownValue("github_enterprise_ruleset.test", tfjsonpath.New("conditions").AtSliceIndex(0).AtMapKey("repository_name").AtSliceIndex(0).AtMapKey("protected"), knownvalue.Bool(true)),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { skipUnlessEnterprise(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config:            config,
				ConfigStateChecks: checks,
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

	checks := []statecheck.StateCheck{
		statecheck.ExpectKnownValue("github_enterprise_ruleset.test", tfjsonpath.New("name"), knownvalue.StringExact(rulesetName)),
		statecheck.ExpectKnownValue("github_enterprise_ruleset.test", tfjsonpath.New("target"), knownvalue.StringExact("branch")),
		statecheck.ExpectKnownValue("github_enterprise_ruleset.test", tfjsonpath.New("enforcement"), knownvalue.StringExact("active")),
		statecheck.ExpectKnownValue("github_enterprise_ruleset.test", tfjsonpath.New("conditions").AtSliceIndex(0).AtMapKey("organization_id"), knownvalue.ListSizeExact(1)),
		statecheck.ExpectKnownValue("github_enterprise_ruleset.test", tfjsonpath.New("conditions").AtSliceIndex(0).AtMapKey("organization_id").AtSliceIndex(0), knownvalue.Int64Exact(2284107)),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { skipUnlessEnterprise(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config:            config,
				ConfigStateChecks: checks,
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
		ProviderFactories: providerFactories,
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

func TestAccGithubEnterpriseRuleset_conflictingRepositoryConditions(t *testing.T) {
	config := fmt.Sprintf(`
		resource "github_enterprise_ruleset" "test" {
			enterprise_slug = "%s"
			name            = "%s-conflict-test"
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

				repository_property {
					include {
						name            = "language"
						property_values = ["Go"]
					}
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
	`, testAccConf.enterpriseSlug, testResourcePrefix)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { skipUnlessEnterprise(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config:      config,
				ExpectError: regexp.MustCompile(`"conditions.0.repository_name": only one of`),
			},
		},
	})
}
