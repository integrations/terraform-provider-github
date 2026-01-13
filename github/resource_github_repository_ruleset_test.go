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

func TestAccGithubRepositoryRuleset(t *testing.T) {
	t.Run("create_branch_ruleset", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-ruleset-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
resource "github_repository" "test" {
	name = "%s"
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

		required_deployments {
			required_deployment_environments = [github_repository_environment.example.environment]
		}

		required_signatures = false

		pull_request {
			dismiss_stale_reviews_on_push     = true
			require_code_owner_review         = true
			require_last_push_approval        = true
			required_approving_review_count   = 2
			required_review_thread_resolution = true
		}

		required_status_checks {
			do_not_enforce_on_create             = true
			strict_required_status_checks_policy = true

			required_check {
				context = "ci"
			}
		}

		non_fast_forward = true

		required_code_scanning {
			required_code_scanning_tool {
			alerts_threshold          = "errors"
			security_alerts_threshold = "high_or_higher"
			tool                      = "CodeQL"
			}
		}
	}
}
`, repoName)

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
						resource.TestCheckResourceAttr("github_repository_ruleset.test", "rules.0.required_code_scanning.0.required_code_scanning_tool.0.alerts_threshold", "errors"),
						resource.TestCheckResourceAttr("github_repository_ruleset.test", "rules.0.required_code_scanning.0.required_code_scanning_tool.0.security_alerts_threshold", "high_or_higher"),
						resource.TestCheckResourceAttr("github_repository_ruleset.test", "rules.0.required_code_scanning.0.required_code_scanning_tool.0.tool", "CodeQL"),
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
`, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessMode(t, enterprise) },
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
			max_file_size = 1048576
		}

		file_extension_restriction {
			restricted_file_extensions = ["*.zip"]
		}
	}
}

`, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:  func() { skipUnlessHasPaidOrgs(t) },
			Providers: testAccProviders,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_repository_ruleset.test", "name", "test-push"),
						resource.TestCheckResourceAttr("github_repository_ruleset.test", "target", "push"),
						resource.TestCheckResourceAttr("github_repository_ruleset.test", "enforcement", "active"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "bypass_actors.#", "2"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "bypass_actors.0.actor_type", "DeployKey"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "bypass_actors.0.bypass_mode", "always"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "bypass_actors.1.actor_id", "5"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "bypass_actors.1.actor_type", "RepositoryRole"),
						resource.TestCheckResourceAttr("github_organization_ruleset.test", "bypass_actors.1.bypass_mode", "always"),
						resource.TestCheckResourceAttr("github_repository_ruleset.test", "rules.0.file_path_restriction.0.restricted_file_paths.0", "test.txt"),
						resource.TestCheckResourceAttr("github_repository_ruleset.test", "rules.0.max_file_size.0.max_file_size", "1048576"),
						resource.TestCheckResourceAttr("github_repository_ruleset.test", "rules.0.file_extension_restriction.0.restricted_file_extensions.0", "*.zip"),
					),
				},
			},
		})
	})

	t.Run("update_ruleset_name", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-ruleset-rename-%s", testResourcePrefix, randomID)
		name := fmt.Sprintf(`ruleset-%[1]s`, randomID)
		nameUpdated := fmt.Sprintf(`%[1]s-renamed`, randomID)

		config := `
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
`

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, repoName, randomID, name),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_repository_ruleset.test", "name", name),
					),
				},
				{
					Config: fmt.Sprintf(config, repoName, randomID, nameUpdated),
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

		config := fmt.Sprintf(`
resource "github_repository" "test" {
	name         = "%s"
	description  = "Terraform acceptance tests %[1]s"
	auto_init    = true
}

resource "github_repository_ruleset" "test" {
	name        = "test-bypass"
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
			include = ["~ALL"]
			exclude = []
		}
	}

	rules {
		creation = true
	}
}
`, repoName)

		configUpdated := fmt.Sprintf(`
resource "github_repository" "test" {
	name         = "%s"
	description  = "Terraform acceptance tests %[1]s"
	auto_init    = true
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
		creation = true
	}
}
`, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:  func() { skipUnauthenticated(t) },
			Providers: testAccProviders,
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

		config := fmt.Sprintf(`
resource "github_repository" "test" {
	name         = "%s"
	description  = "Terraform acceptance tests %s"
	auto_init    = true
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
`, repoName, randomID, bypassMode)

		resource.Test(t, resource.TestCase{
			PreCheck:  func() { skipUnauthenticated(t) },
			Providers: testAccProviders,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, randomID, bypassMode),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_repository_ruleset.test", "bypass_actors.0.bypass_mode", bypassMode),
					),
				},
				{
					Config: fmt.Sprintf(config, randomID, bypassModeUpdated),
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
		`, repoName, randomID)

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
	t.Run("skips update and delete on archived repository", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-ruleset-arch-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name      = "%s"
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
		`, repoName)

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
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-ruleset-arch-cr-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name      = "%s"
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
		`, repoName)

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
