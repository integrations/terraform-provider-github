package github

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubRepositoryEnvironmentDeploymentPolicy(t *testing.T) {
	t.Run("create_branch_policy", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)
		envName := "environment / test"
		config := fmt.Sprintf(`
data "github_user" "current" {
	username = ""
}

resource "github_repository" "test" {
	name      = "%s"
}

resource "github_repository_environment" "test" {
	repository 	= github_repository.test.name
	environment	= "%s"
	wait_timer	= 10000
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
`, repoName, envName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_repository_environment_deployment_policy.test", "repository", repoName),
						resource.TestCheckResourceAttr("github_repository_environment_deployment_policy.test", "environment", "environment / test"),
						resource.TestCheckResourceAttr("github_repository_environment_deployment_policy.test", "branch_pattern", "releases/*"),
						resource.TestCheckNoResourceAttr("github_repository_environment_deployment_policy.test", "tag_pattern"),
						resource.TestCheckResourceAttrSet("github_repository_environment_deployment_policy.test", "policy_id"),
					),
				},
			},
		})
	})

	t.Run("create_update_branch_policy", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-env-deploy-%s", testResourcePrefix, randomID)

		config1 := fmt.Sprintf(`
data "github_user" "current" {
	username = ""
}

resource "github_repository" "test" {
	name      = "%s"
}

resource "github_repository_environment" "test" {
	repository 	= github_repository.test.name
	environment	= "environment/test"
	wait_timer	= 10000
	reviewers {
		users = [data.github_user.current.id]
	}
	deployment_branch_policy {
		protected_branches     = false
		custom_branch_policies = true
	}
}

resource "github_repository_environment_deployment_policy" "test" {
	repository 	= github_repository.test.name
	environment	= github_repository_environment.test.environment
	branch_pattern = "main"
}
`, repoName)

		config2 := fmt.Sprintf(`
data "github_user" "current" {
	username = ""
}

resource "github_repository" "test" {
	name      = "%s"
}

resource "github_repository_environment" "test" {
	repository 	= github_repository.test.name
	environment	= "environment/test"
	wait_timer	= 10000
	reviewers {
		users = [data.github_user.current.id]
	}
	deployment_branch_policy {
		protected_branches     = false
		custom_branch_policies = true
	}
}

resource "github_repository_environment_deployment_policy" "test" {
	repository 	= github_repository.test.name
	environment	= github_repository_environment.test.environment
	branch_pattern = "release/*"
}
`, repoName)

		var policyID string
		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config1,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_repository_environment_deployment_policy.test", "repository", repoName),
						resource.TestCheckResourceAttr("github_repository_environment_deployment_policy.test", "environment", "environment/test"),
						resource.TestCheckResourceAttr("github_repository_environment_deployment_policy.test", "branch_pattern", "main"),
						resource.TestCheckNoResourceAttr("github_repository_environment_deployment_policy.test", "tag_pattern"),
						resource.TestCheckResourceAttrSet("github_repository_environment_deployment_policy.test", "policy_id"),
						getResourceAttr("github_repository_environment_deployment_policy.test", "policy_id", &policyID),
					),
				},
				{
					Config: config2,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_repository_environment_deployment_policy.test", "repository", repoName),
						resource.TestCheckResourceAttr("github_repository_environment_deployment_policy.test", "environment", "environment/test"),
						resource.TestCheckResourceAttr("github_repository_environment_deployment_policy.test", "branch_pattern", "release/*"),
						resource.TestCheckNoResourceAttr("github_repository_environment_deployment_policy.test", "tag_pattern"),
						resource.TestCheckResourceAttrPtr("github_repository_environment_deployment_policy.test", "policy_id", &policyID),
					),
				},
			},
		})
	})

	t.Run("create_tag_policy", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-env-deploy-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
data "github_user" "current" {
	username = ""
}

resource "github_repository" "test" {
	name      = "%s"
}

resource "github_repository_environment" "test" {
	repository 	= github_repository.test.name
	environment	= "environment/test"
	wait_timer	= 10000
	reviewers {
		users = [data.github_user.current.id]
	}
	deployment_branch_policy {
		protected_branches     = false
		custom_branch_policies = true
	}
}

resource "github_repository_environment_deployment_policy" "test" {
	repository 	= github_repository.test.name
	environment	= github_repository_environment.test.environment
	tag_pattern = "v*"
}
`, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_repository_environment_deployment_policy.test", "repository", repoName),
						resource.TestCheckResourceAttr("github_repository_environment_deployment_policy.test", "environment", "environment/test"),
						resource.TestCheckResourceAttr("github_repository_environment_deployment_policy.test", "tag_pattern", "v*"),
						resource.TestCheckNoResourceAttr("github_repository_environment_deployment_policy.test", "branch_pattern"),
					),
				},
			},
		})
	})

	t.Run("create_update_tag_policy", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-env-deploy-%s", testResourcePrefix, randomID)

		config1 := fmt.Sprintf(`
data "github_user" "current" {
	username = ""
}

resource "github_repository" "test" {
	name      = "%s"
}

resource "github_repository_environment" "test" {
	repository 	= github_repository.test.name
	environment	= "environment/test"
	wait_timer	= 10000
	reviewers {
		users = [data.github_user.current.id]
	}
	deployment_branch_policy {
		protected_branches     = false
		custom_branch_policies = true
	}
}

resource "github_repository_environment_deployment_policy" "test" {
	repository 	= github_repository.test.name
	environment	= github_repository_environment.test.environment
	tag_pattern = "v*"
}
`, repoName)

		config2 := fmt.Sprintf(`
data "github_user" "current" {
	username = ""
}

resource "github_repository" "test" {
	name      = "%s"
}

resource "github_repository_environment" "test" {
	repository 	= github_repository.test.name
	environment	= "environment/test"
	wait_timer	= 10000
	reviewers {
		users = [data.github_user.current.id]
	}
	deployment_branch_policy {
		protected_branches     = false
		custom_branch_policies = true
	}
}

resource "github_repository_environment_deployment_policy" "test" {
	repository 	= github_repository.test.name
	environment	= github_repository_environment.test.environment
	tag_pattern = "version*"
}
`, repoName)

		var policyID string
		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config1,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_repository_environment_deployment_policy.test", "repository", repoName),
						resource.TestCheckResourceAttr("github_repository_environment_deployment_policy.test", "environment", "environment/test"),
						resource.TestCheckResourceAttr("github_repository_environment_deployment_policy.test", "tag_pattern", "v*"),
						resource.TestCheckNoResourceAttr("github_repository_environment_deployment_policy.test", "branch_pattern"),
						getResourceAttr("github_repository_environment_deployment_policy.test", "policy_id", &policyID),
					),
				},
				{
					Config: config2,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_repository_environment_deployment_policy.test", "repository", repoName),
						resource.TestCheckResourceAttr("github_repository_environment_deployment_policy.test", "environment", "environment/test"),
						resource.TestCheckResourceAttr("github_repository_environment_deployment_policy.test", "tag_pattern", "version*"),
						resource.TestCheckNoResourceAttr("github_repository_environment_deployment_policy.test", "branch_pattern"),
						resource.TestCheckResourceAttrPtr("github_repository_environment_deployment_policy.test", "policy_id", &policyID),
					),
				},
			},
		})
	})

	t.Run("recreates_when_pattern_type_changes_from_branch_to_tag", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-env-deploy-%s", testResourcePrefix, randomID)

		config1 := fmt.Sprintf(`
data "github_user" "current" {
	username = ""
}

resource "github_repository" "test" {
	name      = "%s"
}

resource "github_repository_environment" "test" {
	repository 	= github_repository.test.name
	environment	= "environment/test"
	wait_timer	= 10000
	reviewers {
		users = [data.github_user.current.id]
	}
	deployment_branch_policy {
		protected_branches     = false
		custom_branch_policies = true
	}
}

resource "github_repository_environment_deployment_policy" "test" {
	repository 	= github_repository.test.name
	environment	= github_repository_environment.test.environment
	branch_pattern = "release/*"
}

`, repoName)

		config2 := fmt.Sprintf(`
data "github_user" "current" {
	username = ""
}

resource "github_repository" "test" {
	name      = "%s"
}

resource "github_repository_environment" "test" {
	repository 	= github_repository.test.name
	environment	= "environment/test"
	wait_timer	= 10000
	reviewers {
		users = [data.github_user.current.id]
	}
	deployment_branch_policy {
		protected_branches     = false
		custom_branch_policies = true
	}
}

resource "github_repository_environment_deployment_policy" "test" {
	repository 	= github_repository.test.name
	environment	= github_repository_environment.test.environment
	tag_pattern = "v*"
}
`, repoName)

		var policyID string
		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config1,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_repository_environment_deployment_policy.test", "repository", repoName),
						resource.TestCheckResourceAttr("github_repository_environment_deployment_policy.test", "environment", "environment/test"),
						resource.TestCheckResourceAttr("github_repository_environment_deployment_policy.test", "branch_pattern", "release/*"),
						resource.TestCheckNoResourceAttr("github_repository_environment_deployment_policy.test", "tag_pattern"),
						getResourceAttr("github_repository_environment_deployment_policy.test", "policy_id", &policyID),
					),
				},
				{
					Config: config2,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_repository_environment_deployment_policy.test", "repository", repoName),
						resource.TestCheckResourceAttr("github_repository_environment_deployment_policy.test", "environment", "environment/test"),
						resource.TestCheckResourceAttr("github_repository_environment_deployment_policy.test", "tag_pattern", "v*"),
						resource.TestCheckNoResourceAttr("github_repository_environment_deployment_policy.test", "branch_pattern"),
						resource.TestCheckResourceAttrWith("github_repository_environment_deployment_policy.test", "policy_id", func(id string) error {
							if id == policyID {
								return fmt.Errorf("expected policy_id to change when pattern type changes")
							}
							return nil
						}),
					),
				},
			},
		})
	})

	t.Run("recreates_when_pattern_type_changes_from_tag_to_branch", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-env-deploy-%s", testResourcePrefix, randomID)

		config1 := fmt.Sprintf(`
data "github_user" "current" {
	username = ""
}

resource "github_repository" "test" {
	name      = "%s"
}

resource "github_repository_environment" "test" {
	repository 	= github_repository.test.name
	environment	= "environment/test"
	wait_timer	= 10000
	reviewers {
		users = [data.github_user.current.id]
	}
	deployment_branch_policy {
		protected_branches     = false
		custom_branch_policies = true
	}
}

resource "github_repository_environment_deployment_policy" "test" {
	repository 	= github_repository.test.name
	environment	= github_repository_environment.test.environment
	tag_pattern = "v*"
}
`, repoName)

		config2 := fmt.Sprintf(`
data "github_user" "current" {
	username = ""
}

resource "github_repository" "test" {
	name      = "%s"
}

resource "github_repository_environment" "test" {
	repository 	= github_repository.test.name
	environment	= "environment/test"
	wait_timer	= 10000
	reviewers {
		users = [data.github_user.current.id]
	}
	deployment_branch_policy {
		protected_branches     = false
		custom_branch_policies = true
	}
}

resource "github_repository_environment_deployment_policy" "test" {
	repository 	= github_repository.test.name
	environment	= github_repository_environment.test.environment
	branch_pattern = "release/*"
}
`, repoName)

		var policyID string
		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config1,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_repository_environment_deployment_policy.test", "repository", repoName),
						resource.TestCheckResourceAttr("github_repository_environment_deployment_policy.test", "environment", "environment/test"),
						resource.TestCheckResourceAttr("github_repository_environment_deployment_policy.test", "tag_pattern", "v*"),
						resource.TestCheckNoResourceAttr("github_repository_environment_deployment_policy.test", "branch_pattern"),
						getResourceAttr("github_repository_environment_deployment_policy.test", "policy_id", &policyID),
					),
				},
				{
					Config: config2,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr("github_repository_environment_deployment_policy.test", "repository", repoName),
						resource.TestCheckResourceAttr("github_repository_environment_deployment_policy.test", "environment", "environment/test"),
						resource.TestCheckResourceAttr("github_repository_environment_deployment_policy.test", "branch_pattern", "release/*"),
						resource.TestCheckNoResourceAttr("github_repository_environment_deployment_policy.test", "tag_pattern"),
						resource.TestCheckResourceAttrWith("github_repository_environment_deployment_policy.test", "policy_id", func(id string) error {
							if id == policyID {
								return fmt.Errorf("expected policy_id to change when pattern type changes")
							}
							return nil
						}),
					),
				},
			},
		})
	})

	t.Run("import", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)
		envName := "environment / test"
		config := fmt.Sprintf(`
data "github_user" "current" {
	username = ""
}

resource "github_repository" "test" {
	name = "%s"
}

resource "github_repository_environment" "test" {
	repository  = github_repository.test.name
	environment = "%s"
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
`, repoName, envName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  resource.TestCheckResourceAttr("github_repository_environment_deployment_policy.test", "environment", envName),
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
		repoName := fmt.Sprintf("%srepo-env-deploy-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`
resource "github_repository" "test" {
	name      = "%s"
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
	repository 	= github_repository.test.name
	environment	= github_repository_environment.test.environment
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
		repoName := fmt.Sprintf("%srepo-env-deploy-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`
resource "github_repository" "test" {
	name      = "%s"
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
		repoName := fmt.Sprintf("%srepo-env-deploy-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`
resource "github_repository" "test" {
	name      = "%s"
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
		repoName := fmt.Sprintf("%srepo-env-deploy-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`
resource "github_repository" "test" {
	name      = "%s"
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
	tag_pattern = ""
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
