package github

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

const enterpriseRulesetResource = "github_enterprise_ruleset.test"

// enterpriseRulesetBasicConfig is a minimal branch ruleset. Takes the enterprise slug
// and the ruleset name.
const enterpriseRulesetBasicConfig = `
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

// enterpriseRulesetBypassModeConfig takes the enterprise slug, the ruleset name and
// the bypass mode to apply to a single OrganizationAdmin actor.
const enterpriseRulesetBypassModeConfig = `
resource "github_enterprise_ruleset" "test" {
  enterprise_slug = "%s"
  name            = "%s"
  target          = "branch"
  enforcement     = "active"

  bypass_actors {
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
`

func TestAccGithubEnterpriseRuleset(t *testing.T) {
	t.Parallel()

	t.Run("creates a branch ruleset without error", func(t *testing.T) {
		t.Parallel()

		name := fmt.Sprintf("%senterprise-branch-%s", testResourcePrefix, acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum))

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(enterpriseRulesetBasicConfig, testAccConf.enterpriseSlug, name),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(enterpriseRulesetResource, tfjsonpath.New("ruleset_id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue(enterpriseRulesetResource, tfjsonpath.New("node_id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue(enterpriseRulesetResource, tfjsonpath.New("etag"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue(enterpriseRulesetResource, tfjsonpath.New("name"), knownvalue.StringExact(name)),
					},
				},
			},
		})
	})

	t.Run("creates a branch ruleset with all branch rules", func(t *testing.T) {
		t.Parallel()

		name := fmt.Sprintf("%senterprise-rules-%s", testResourcePrefix, acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum))
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
			    actor_type  = "EnterpriseOwner"
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
			    non_fast_forward        = true

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
			  }
			}
		`, testAccConf.enterpriseSlug, name)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(enterpriseRulesetResource, tfjsonpath.New("ruleset_id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue(enterpriseRulesetResource, tfjsonpath.New("bypass_actors"), knownvalue.ListSizeExact(2)),
						statecheck.ExpectKnownValue(enterpriseRulesetResource, tfjsonpath.New("bypass_actors").AtSliceIndex(0).AtMapKey("actor_type"), knownvalue.StringExact("DeployKey")),
						statecheck.ExpectKnownValue(enterpriseRulesetResource, tfjsonpath.New("bypass_actors").AtSliceIndex(1).AtMapKey("actor_type"), knownvalue.StringExact("EnterpriseOwner")),
						statecheck.ExpectKnownValue(enterpriseRulesetResource, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("required_code_scanning").AtSliceIndex(0).AtMapKey("required_code_scanning_tool").AtSliceIndex(0).AtMapKey("tool"), knownvalue.StringExact("CodeQL")),
						statecheck.ExpectKnownValue(enterpriseRulesetResource, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("copilot_code_review").AtSliceIndex(0).AtMapKey("review_on_push"), knownvalue.Bool(true)),
						statecheck.ExpectKnownValue(enterpriseRulesetResource, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("branch_name_pattern").AtSliceIndex(0).AtMapKey("operator"), knownvalue.StringExact("starts_with")),
					},
				},
			},
		})
	})

	t.Run("creates a tag ruleset without error", func(t *testing.T) {
		t.Parallel()

		name := fmt.Sprintf("%senterprise-tag-%s", testResourcePrefix, acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum))
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
			    tag_name_pattern {
			      operator = "starts_with"
			      pattern  = "v"
			    }
			  }
			}
		`, testAccConf.enterpriseSlug, name)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(enterpriseRulesetResource, tfjsonpath.New("target"), knownvalue.StringExact("tag")),
						statecheck.ExpectKnownValue(enterpriseRulesetResource, tfjsonpath.New("ruleset_id"), knownvalue.NotNull()),
					},
				},
			},
		})
	})

	t.Run("creates a push ruleset without error", func(t *testing.T) {
		t.Parallel()

		name := fmt.Sprintf("%senterprise-push-%s", testResourcePrefix, acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum))
		config := fmt.Sprintf(`
			resource "github_enterprise_ruleset" "test" {
			  enterprise_slug = "%s"
			  name            = "%s"
			  target          = "push"
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
		`, testAccConf.enterpriseSlug, name)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(enterpriseRulesetResource, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("file_path_restriction").AtSliceIndex(0).AtMapKey("restricted_file_paths").AtSliceIndex(0), knownvalue.StringExact("test.txt")),
						statecheck.ExpectKnownValue(enterpriseRulesetResource, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("max_file_size").AtSliceIndex(0).AtMapKey("max_file_size"), knownvalue.Int64Exact(99)),
					},
				},
			},
		})
	})

	t.Run("creates a repository ruleset without error", func(t *testing.T) {
		t.Parallel()

		name := fmt.Sprintf("%senterprise-repo-%s", testResourcePrefix, acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum))
		config := fmt.Sprintf(`
			resource "github_enterprise_ruleset" "test" {
			  enterprise_slug = "%s"
			  name            = "%s"
			  target          = "repository"
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
			  }

			  rules {
			    repository_delete   = true
			    repository_transfer = true

			    repository_name {
			      negate  = false
			      pattern = "^svc-"
			    }

			    repository_visibility {
			      internal = true
			      private  = true
			    }
			  }
			}
		`, testAccConf.enterpriseSlug, name)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(enterpriseRulesetResource, tfjsonpath.New("target"), knownvalue.StringExact("repository")),
						statecheck.ExpectKnownValue(enterpriseRulesetResource, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("repository_delete"), knownvalue.Bool(true)),
						statecheck.ExpectKnownValue(enterpriseRulesetResource, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("repository_transfer"), knownvalue.Bool(true)),
						statecheck.ExpectKnownValue(enterpriseRulesetResource, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("repository_name").AtSliceIndex(0).AtMapKey("pattern"), knownvalue.StringExact("^svc-")),
						statecheck.ExpectKnownValue(enterpriseRulesetResource, tfjsonpath.New("rules").AtSliceIndex(0).AtMapKey("repository_visibility").AtSliceIndex(0).AtMapKey("internal"), knownvalue.Bool(true)),
					},
				},
			},
		})
	})

	t.Run("targets organizations by id", func(t *testing.T) {
		t.Parallel()

		name := fmt.Sprintf("%senterprise-org-id-%s", testResourcePrefix, acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum))
		config := fmt.Sprintf(`
			data "github_organization" "test" {
			  name = "%s"
			}

			resource "github_enterprise_ruleset" "test" {
			  enterprise_slug = "%s"
			  name            = "%s"
			  target          = "branch"
			  enforcement     = "active"

			  conditions {
			    organization_id = [data.github_organization.test.id]

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
		`, testAccConf.owner, testAccConf.enterpriseSlug, name)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(enterpriseRulesetResource, tfjsonpath.New("conditions").AtSliceIndex(0).AtMapKey("organization_id"), knownvalue.ListSizeExact(1)),
					},
				},
			},
		})
	})

	t.Run("targets organizations by property", func(t *testing.T) {
		t.Parallel()

		name := fmt.Sprintf("%senterprise-org-prop-%s", testResourcePrefix, acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum))
		config := fmt.Sprintf(`
			resource "github_enterprise_ruleset" "test" {
			  enterprise_slug = "%s"
			  name            = "%s"
			  target          = "branch"
			  enforcement     = "active"

			  conditions {
			    organization_property {
			      include {
			        name            = "environment"
			        property_values = ["production"]
			      }

			      exclude {
			        name            = "environment"
			        property_values = ["sandbox"]
			      }
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
		`, testAccConf.enterpriseSlug, name)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(enterpriseRulesetResource, tfjsonpath.New("conditions").AtSliceIndex(0).AtMapKey("organization_property").AtSliceIndex(0).AtMapKey("include"), knownvalue.ListSizeExact(1)),
						statecheck.ExpectKnownValue(enterpriseRulesetResource, tfjsonpath.New("conditions").AtSliceIndex(0).AtMapKey("organization_property").AtSliceIndex(0).AtMapKey("exclude"), knownvalue.ListSizeExact(1)),
					},
				},
			},
		})
	})

	t.Run("rejects conflicting organization conditions", func(t *testing.T) {
		t.Parallel()

		name := fmt.Sprintf("%senterprise-org-conflict-%s", testResourcePrefix, acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum))
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

			    organization_property {
			      include {
			        name            = "environment"
			        property_values = ["production"]
			      }
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
		`, testAccConf.enterpriseSlug, name)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      config,
					ExpectError: regexp.MustCompile(`only one of ` + "`" + `conditions\.0\.organization_id,conditions\.0\.organization_name,conditions\.0\.organization_property` + "`" + ` can be specified`),
				},
			},
		})
	})

	t.Run("updates the bypass mode without error", func(t *testing.T) {
		t.Parallel()

		name := fmt.Sprintf("%senterprise-bypass-%s", testResourcePrefix, acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum))

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(enterpriseRulesetBypassModeConfig, testAccConf.enterpriseSlug, name, "always"),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(enterpriseRulesetResource, tfjsonpath.New("bypass_actors").AtSliceIndex(0).AtMapKey("bypass_mode"), knownvalue.StringExact("always")),
					},
				},
				{
					Config: fmt.Sprintf(enterpriseRulesetBypassModeConfig, testAccConf.enterpriseSlug, name, "exempt"),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(enterpriseRulesetResource, tfjsonpath.New("bypass_actors").AtSliceIndex(0).AtMapKey("bypass_mode"), knownvalue.StringExact("exempt")),
					},
				},
				{
					Config: fmt.Sprintf(enterpriseRulesetBasicConfig, testAccConf.enterpriseSlug, name),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(enterpriseRulesetResource, tfjsonpath.New("bypass_actors"), knownvalue.ListSizeExact(0)),
					},
				},
			},
		})
	})

	t.Run("imports an enterprise ruleset without error", func(t *testing.T) {
		t.Parallel()

		name := fmt.Sprintf("%senterprise-import-%s", testResourcePrefix, acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum))

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(enterpriseRulesetBasicConfig, testAccConf.enterpriseSlug, name),
				},
				{
					ResourceName:            enterpriseRulesetResource,
					ImportState:             true,
					ImportStateVerify:       true,
					ImportStateVerifyIgnore: []string{"etag"},
				},
			},
		})
	})

	t.Run("rejects push rules on a branch target", func(t *testing.T) {
		t.Parallel()

		name := fmt.Sprintf("%senterprise-badrule-%s", testResourcePrefix, acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum))
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
			    max_file_size {
			      max_file_size = 10
			    }
			  }
			}
		`, testAccConf.enterpriseSlug, name)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      config,
					ExpectError: regexp.MustCompile(`rule "max_file_size" is not valid for branch target`),
				},
			},
		})
	})

	t.Run("rejects ref_name on a repository target", func(t *testing.T) {
		t.Parallel()

		name := fmt.Sprintf("%senterprise-badcond-%s", testResourcePrefix, acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum))
		config := fmt.Sprintf(`
			resource "github_enterprise_ruleset" "test" {
			  enterprise_slug = "%s"
			  name            = "%s"
			  target          = "repository"
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
			    repository_delete = true
			  }
			}
		`, testAccConf.enterpriseSlug, name)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      config,
					ExpectError: regexp.MustCompile(`ref_name must not be set for repository target`),
				},
			},
		})
	})

	t.Run("rejects conflicting repository conditions", func(t *testing.T) {
		t.Parallel()

		name := fmt.Sprintf("%senterprise-conflict-%s", testResourcePrefix, acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum))
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
		`, testAccConf.enterpriseSlug, name)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      config,
					ExpectError: regexp.MustCompile(`only one of ` + "`" + `conditions\.0\.repository_name,conditions\.0\.repository_property` + "`" + ` can be specified`),
				},
			},
		})
	})
}
