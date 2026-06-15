package github

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/compare"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubRepositoryEnvironmentDeploymentPolicy(t *testing.T) {
	t.Run("create_branch_policy", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
data "github_user" "current" {
	username = ""
}

resource "github_repository" "test" {
	name = "%s"
}

resource "github_repository_environment" "test" {
	repository  = github_repository.test.name
	environment = "test"
	wait_timer = 10000
	reviewers {
		users = [data.github_user.current.id]
	}
	deployment_branch_policy {
		protected_branches     = false
		custom_branch_policies = true
	}
}

resource "github_repository_environment_deployment_policy" "test" {
	repository 	   = github_repository.test.name
	environment	   = github_repository_environment.test.environment
	branch_pattern = "releases/*"
}
`, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_environment_deployment_policy.test", tfjsonpath.New("repository_id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_repository_environment_deployment_policy.test", tfjsonpath.New("policy_id"), knownvalue.NotNull()),
					},
				},
			},
		})
	})

	t.Run("create_update_branch_policy", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)

		config := `
data "github_user" "current" {
	username = ""
}

resource "github_repository" "test" {
	name = "%s"
}

resource "github_repository_environment" "test" {
	repository  = github_repository.test.name
	environment = "test"
	wait_timer  = 10000
	reviewers {
		users = [data.github_user.current.id]
	}
	deployment_branch_policy {
		protected_branches     = false
		custom_branch_policies = true
	}
}

resource "github_repository_environment_deployment_policy" "test" {
	repository     = github_repository.test.name
	environment    = github_repository_environment.test.environment
	branch_pattern = "%s"
}
`

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, repoName, "main"),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_environment_deployment_policy.test", tfjsonpath.New("repository_id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_repository_environment_deployment_policy.test", tfjsonpath.New("policy_id"), knownvalue.NotNull()),
					},
				},
				{
					Config: fmt.Sprintf(config, repoName, "release/*"),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_repository_environment_deployment_policy.test", plancheck.ResourceActionUpdate),
						},
					},
				},
			},
		})
	})

	t.Run("create_tag_policy", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
data "github_user" "current" {
	username = ""
}

resource "github_repository" "test" {
	name = "%s"
}

resource "github_repository_environment" "test" {
	repository  = github_repository.test.name
	environment = "test"
	wait_timer  = 10000
	reviewers {
		users = [data.github_user.current.id]
	}
	deployment_branch_policy {
		protected_branches     = false
		custom_branch_policies = true
	}
}

resource "github_repository_environment_deployment_policy" "test" {
	repository  = github_repository.test.name
	environment = github_repository_environment.test.environment
	tag_pattern = "v*"
}
`, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_environment_deployment_policy.test", tfjsonpath.New("repository_id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_repository_environment_deployment_policy.test", tfjsonpath.New("policy_id"), knownvalue.NotNull()),
					},
				},
			},
		})
	})

	t.Run("create_update_tag_policy", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)

		config := `
data "github_user" "current" {
	username = ""
}

resource "github_repository" "test" {
	name = "%s"
}

resource "github_repository_environment" "test" {
	repository 	= github_repository.test.name
	environment = "environment/test"
	wait_timer  = 10000
	reviewers {
		users = [data.github_user.current.id]
	}
	deployment_branch_policy {
		protected_branches     = false
		custom_branch_policies = true
	}
}

resource "github_repository_environment_deployment_policy" "test" {
	repository  = github_repository.test.name
	environment = github_repository_environment.test.environment
	tag_pattern = "%s"
}
`

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, repoName, "v*"),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_environment_deployment_policy.test", tfjsonpath.New("repository_id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_repository_environment_deployment_policy.test", tfjsonpath.New("policy_id"), knownvalue.NotNull()),
					},
				},
				{
					Config: fmt.Sprintf(config, repoName, "version*"),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_repository_environment_deployment_policy.test", plancheck.ResourceActionUpdate),
						},
					},
				},
			},
		})
	})

	t.Run("recreates_when_pattern_type_changes_from_branch_to_tag", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)

		preConfig := fmt.Sprintf(`
data "github_user" "current" {
	username = ""
}

resource "github_repository" "test" {
	name = "%s"
}

resource "github_repository_environment" "test" {
	repository  = github_repository.test.name
	environment = "test"
	wait_timer  = 10000
	reviewers {
		users = [data.github_user.current.id]
	}
	deployment_branch_policy {
		protected_branches     = false
		custom_branch_policies = true
	}
}
`, repoName)

		config1 := fmt.Sprintf(`
%s

resource "github_repository_environment_deployment_policy" "test" {
	repository     = github_repository.test.name
	environment    = github_repository_environment.test.environment
	branch_pattern = "release/*"
}

`, preConfig)

		config2 := fmt.Sprintf(`
%s

resource "github_repository_environment_deployment_policy" "test" {
	repository  = github_repository.test.name
	environment = github_repository_environment.test.environment
	tag_pattern = "v*"
}
`, preConfig)

		comparePolicyIDChanages := statecheck.CompareValue(compare.ValuesDiffer())
		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config1,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_environment_deployment_policy.test", tfjsonpath.New("repository_id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_repository_environment_deployment_policy.test", tfjsonpath.New("policy_id"), knownvalue.NotNull()),
						comparePolicyIDChanages.AddStateValue("github_repository_environment_deployment_policy.test", tfjsonpath.New("policy_id")),
					},
				},
				{
					Config: config2,
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_repository_environment_deployment_policy.test", plancheck.ResourceActionReplace),
						},
					},
					ConfigStateChecks: []statecheck.StateCheck{
						comparePolicyIDChanages.AddStateValue("github_repository_environment_deployment_policy.test", tfjsonpath.New("policy_id")),
					},
				},
			},
		})
	})

	t.Run("recreates_when_pattern_type_changes_from_tag_to_branch", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)

		preConfig := fmt.Sprintf(`
data "github_user" "current" {
	username = ""
}

resource "github_repository" "test" {
	name = "%s"
}

resource "github_repository_environment" "test" {
	repository  = github_repository.test.name
	environment = "test"
	wait_timer  = 10000
	reviewers {
		users = [data.github_user.current.id]
	}
	deployment_branch_policy {
		protected_branches     = false
		custom_branch_policies = true
	}
}
`, repoName)

		config1 := fmt.Sprintf(`
%s

resource "github_repository_environment_deployment_policy" "test" {
	repository  = github_repository.test.name
	environment = github_repository_environment.test.environment
	tag_pattern = "v*"
}
`, preConfig)

		config2 := fmt.Sprintf(`
%s

resource "github_repository_environment_deployment_policy" "test" {
	repository     = github_repository.test.name
	environment    = github_repository_environment.test.environment
	branch_pattern = "release/*"
}

`, preConfig)

		comparePolicyIDChanages := statecheck.CompareValue(compare.ValuesDiffer())
		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config1,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_environment_deployment_policy.test", tfjsonpath.New("repository_id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_repository_environment_deployment_policy.test", tfjsonpath.New("policy_id"), knownvalue.NotNull()),
						comparePolicyIDChanages.AddStateValue("github_repository_environment_deployment_policy.test", tfjsonpath.New("policy_id")),
					},
				},
				{
					Config: config2,
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_repository_environment_deployment_policy.test", plancheck.ResourceActionReplace),
						},
					},
					ConfigStateChecks: []statecheck.StateCheck{
						comparePolicyIDChanages.AddStateValue("github_repository_environment_deployment_policy.test", tfjsonpath.New("policy_id")),
					},
				},
			},
		})
	})

	t.Run("import", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
data "github_user" "current" {
	username = ""
}

resource "github_repository" "test" {
	name = "%s"
}

resource "github_repository_environment" "test" {
	repository  = github_repository.test.name
	environment = "test"
	wait_timer  = 10000
	reviewers {
		users = [data.github_user.current.id]
	}
	deployment_branch_policy {
		protected_branches     = false
		custom_branch_policies = true
	}
}

resource "github_repository_environment_deployment_policy" "test" {
	repository     = github_repository.test.name
	environment    = github_repository_environment.test.environment
	branch_pattern = "main"
}
`, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_environment_deployment_policy.test", tfjsonpath.New("repository_id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_repository_environment_deployment_policy.test", tfjsonpath.New("policy_id"), knownvalue.NotNull()),
					},
				},
				{
					ResourceName:      "github_repository_environment_deployment_policy.test",
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})

	t.Run("errors when no patterns are set", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
resource "github_repository" "test" {
	name = "%s"
}

resource "github_repository_environment" "test" {
	repository  = github_repository.test.name
	environment = "environment/test"
	deployment_branch_policy {
		protected_branches     = false
		custom_branch_policies = true
	}
}

resource "github_repository_environment_deployment_policy" "test" {
	repository  = github_repository.test.name
	environment = github_repository_environment.test.environment
}
`, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      config,
					ExpectError: regexp.MustCompile("one of `branch_pattern,tag_pattern` must be specified"),
				},
			},
		})
	})

	t.Run("errors when both patterns are set", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
resource "github_repository" "test" {
	name = "%s"
}

resource "github_repository_environment" "test" {
	repository  = github_repository.test.name
	environment = "environment/test"
	deployment_branch_policy {
		protected_branches     = false
		custom_branch_policies = true
	}
}

resource "github_repository_environment_deployment_policy" "test" {
	repository     = github_repository.test.name
	environment    = github_repository_environment.test.environment
	branch_pattern = "main"
	tag_pattern    = "v*"
}
`, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      config,
					ExpectError: regexp.MustCompile("only one of `branch_pattern,tag_pattern` can be specified"),
				},
			},
		})
	})

	t.Run("errors when an empty branch pattern is set", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
resource "github_repository" "test" {
	name = "%s"
}

resource "github_repository_environment" "test" {
	repository 	= github_repository.test.name
	environment	= "environment/test"
	deployment_branch_policy {
		protected_branches     = false
		custom_branch_policies = true
	}
}

resource "github_repository_environment_deployment_policy" "test" {
	repository     = github_repository.test.name
	environment    = github_repository_environment.test.environment
	branch_pattern = ""
}
`, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      config,
					ExpectError: regexp.MustCompile(`expected "branch_pattern" to not be an empty string`),
				},
			},
		})
	})

	t.Run("errors when an empty tag pattern is set", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
resource "github_repository" "test" {
	name = "%s"
}

resource "github_repository_environment" "test" {
	repository  = github_repository.test.name
	environment = "environment/test"
	deployment_branch_policy {
		protected_branches     = false
		custom_branch_policies = true
	}
}

resource "github_repository_environment_deployment_policy" "test" {
	repository     = github_repository.test.name
	environment    = github_repository_environment.test.environment
	tag_pattern    = ""
}
`, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      config,
					ExpectError: regexp.MustCompile(`expected "tag_pattern" to not be an empty string`),
				},
			},
		})
	})
}
