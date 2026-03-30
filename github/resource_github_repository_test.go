package github

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-plugin-testing/compare"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubRepository(t *testing.T) {
	t.Run("creates and updates repositories without error", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		testRepoName := fmt.Sprintf("%screate-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`
			resource "github_repository" "test" {

				name                        = "%s"
				description                 = "Terraform acceptance tests %[1]s"
				has_discussions             = true
				has_issues                  = true
				has_wiki                    = true
				has_downloads               = true
				allow_merge_commit          = true
				allow_squash_merge          = false
				allow_rebase_merge          = false
				allow_auto_merge            = true
				merge_commit_title          = "MERGE_MESSAGE"
				merge_commit_message        = "PR_TITLE"
				auto_init                   = false
				web_commit_signoff_required = true
				visibility                  = "%s"
			}
		`, testRepoName, testAccConf.testRepositoryVisibility)

		resource.ParallelTest(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository.test", tfjsonpath.New("has_issues"), knownvalue.Bool(true)),
						statecheck.ExpectKnownValue("github_repository.test", tfjsonpath.New("has_discussions"), knownvalue.Bool(true)),
						statecheck.ExpectKnownValue("github_repository.test", tfjsonpath.New("allow_auto_merge"), knownvalue.Bool(true)),
						statecheck.ExpectKnownValue("github_repository.test", tfjsonpath.New("merge_commit_title"), knownvalue.StringExact("MERGE_MESSAGE")),
						statecheck.ExpectKnownValue("github_repository.test", tfjsonpath.New("web_commit_signoff_required"), knownvalue.Bool(true)),
					},
				},
			},
		})
	})

	t.Run("updates a repositories name without error", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		oldName := fmt.Sprintf(`%srename-%s`, testResourcePrefix, randomID)
		newName := fmt.Sprintf(`%s-renamed`, oldName)

		config := `
			resource "github_repository" "test" {
				name         = "%[1]s"
				description  = "Terraform acceptance tests %[2]s"
				visibility   = "%s"
			}
		`

		nameDiffer := statecheck.CompareValue(compare.ValuesDiffer())
		resource.ParallelTest(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, oldName, randomID, testAccConf.testRepositoryVisibility),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository.test", tfjsonpath.New("name"), knownvalue.StringExact(oldName)),
						statecheck.ExpectKnownValue("github_repository.test", tfjsonpath.New("full_name"), knownvalue.StringRegexp(regexp.MustCompile(regexp.QuoteMeta(oldName)))),
						nameDiffer.AddStateValue("github_repository.test", tfjsonpath.New("name")),
					},
				},
				{
					Config: fmt.Sprintf(config, newName, randomID, testAccConf.testRepositoryVisibility),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository.test", tfjsonpath.New("name"), knownvalue.StringExact(newName)),
						statecheck.ExpectKnownValue("github_repository.test", tfjsonpath.New("full_name"), knownvalue.StringRegexp(regexp.MustCompile(regexp.QuoteMeta(newName)))),
						nameDiffer.AddStateValue("github_repository.test", tfjsonpath.New("name")),
					},
				},
			},
		})
	})

	t.Run("imports repositories without error", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		testRepoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name         = "%s"
				description  = "Terraform acceptance tests %[1]s"
				auto_init 	 = false
				visibility   = "%s"
			}
		`, testRepoName, testAccConf.testRepositoryVisibility)

		resource.ParallelTest(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository.test", tfjsonpath.New("name"), knownvalue.NotNull()),
					},
				},
				{
					ResourceName:            "github_repository.test",
					ImportState:             true,
					ImportStateVerify:       true,
					ImportStateVerifyIgnore: []string{"auto_init", "vulnerability_alerts", "ignore_vulnerability_alerts_during_read", "etag"},
				},
			},
		})
	})

	t.Run("archives repositories without error", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		testRepoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)

		config := `
resource "github_repository" "test" {
	name         = "%s"
	archived     = %s
	visibility   = "%s"
}
`

		resource.ParallelTest(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, testRepoName, "false", testAccConf.testRepositoryVisibility),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository.test", tfjsonpath.New("archived"), knownvalue.Bool(false)),
					},
				},
				{
					Config: fmt.Sprintf(config, testRepoName, "true", testAccConf.testRepositoryVisibility),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository.test", tfjsonpath.New("archived"), knownvalue.Bool(true)),
					},
				},
			},
		})
	})

	t.Run("manages the project feature for a repository", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		testRepoName := fmt.Sprintf("%sproject-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name         = "%s"
				description  = "Terraform acceptance tests %[1]s"
				has_projects = false
				visibility   = "%s"
			}
		`, testRepoName, testAccConf.testRepositoryVisibility)

		resource.ParallelTest(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository.test", tfjsonpath.New("has_projects"), knownvalue.Bool(false)),
					},
				},
				{
					Config: strings.Replace(config,
						`has_projects = false`,
						`has_projects = true`, 1),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository.test", tfjsonpath.New("has_projects"), knownvalue.Bool(true)),
					},
				},
			},
		})
	})

	t.Run("manages the default branch feature for a repository", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		testRepoName := fmt.Sprintf("%sbranch-%s", testResourcePrefix, randomID)
		config := `
			resource "github_repository" "test" {
				name           = "%s"
				description    = "Terraform acceptance tests %[1]s"
				default_branch = "%s"
				auto_init      = true
				visibility     = "%s"
			}

			resource "github_branch" "default" {
				repository = github_repository.test.name
				branch     = "default"
			}
		`

		defaultBranchChangeCheck := statecheck.CompareValue(compare.ValuesDiffer())
		resource.ParallelTest(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, testRepoName, "main", testAccConf.testRepositoryVisibility),
					ConfigStateChecks: []statecheck.StateCheck{
						defaultBranchChangeCheck.AddStateValue("github_repository.test", tfjsonpath.New("default_branch")),
					},
				},
				{
					Config: fmt.Sprintf(config, testRepoName, "default", testAccConf.testRepositoryVisibility),
					ConfigStateChecks: []statecheck.StateCheck{
						defaultBranchChangeCheck.AddStateValue("github_repository.test", tfjsonpath.New("default_branch")),
					},
				},
				{
					Config: fmt.Sprintf(config, testRepoName, "main", testAccConf.testRepositoryVisibility),
					ConfigStateChecks: []statecheck.StateCheck{
						defaultBranchChangeCheck.AddStateValue("github_repository.test", tfjsonpath.New("default_branch")),
					},
				},
			},
		})
	})

	t.Run("updates_default_branchon_an_empty_repository_without_error", func(t *testing.T) {
		// Although default_branch is deprecated, for backwards compatibility
		// we allow setting it to "main".

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		testRepoName := fmt.Sprintf("%sempty-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name           = "%s"
				description    = "Terraform acceptance tests %[1]s"
				default_branch = "main"
				visibility     = "%s"
			}
		`, testRepoName, testAccConf.testRepositoryVisibility)

		defaultBranchChangeCheck := statecheck.CompareValue(compare.ValuesSame())
		resource.ParallelTest(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository.test", tfjsonpath.New("default_branch"), knownvalue.StringExact("main")),
						defaultBranchChangeCheck.AddStateValue("github_repository.test", tfjsonpath.New("default_branch")),
					},
				},
				{
					Config: strings.Replace(config,
						`acceptance tests`,
						`acceptance test`, 1),
					ConfigStateChecks: []statecheck.StateCheck{
						defaultBranchChangeCheck.AddStateValue("github_repository.test", tfjsonpath.New("default_branch")),
					},
				},
			},
		})
	})

	t.Run("manages the license and gitignore feature for a repository", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		testRepoName := fmt.Sprintf("%slicense-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name           = "%s"
				description    = "Terraform acceptance tests %[1]s"
				license_template   = "ms-pl"
				gitignore_template = "C++"
				visibility         = "%s"
			}
		`, testRepoName, testAccConf.testRepositoryVisibility)

		resource.ParallelTest(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository.test", tfjsonpath.New("license_template"), knownvalue.StringExact("ms-pl")),
						statecheck.ExpectKnownValue("github_repository.test", tfjsonpath.New("gitignore_template"), knownvalue.StringExact("C++")),
					},
				},
			},
		})
	})

	t.Run("configures_topics_for_a_repository", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		testRepoName := fmt.Sprintf("%stopic-%s", testResourcePrefix, randomID)
		topicsBefore := `["terraform", "testing"]`
		topicsAfter := `["terraform", "testing", "extra-topic"]`
		config := `
			resource "github_repository" "test" {
				name        = "%s"
				description = "Terraform acceptance tests %[1]s"
				topics			= %s
				visibility     = "%s"
			}
		`

		resource.ParallelTest(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, testRepoName, topicsBefore, testAccConf.testRepositoryVisibility),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository.test", tfjsonpath.New("topics"), knownvalue.SetSizeExact(2)),
					},
				},
				{
					Config: fmt.Sprintf(config, testRepoName, topicsAfter, testAccConf.testRepositoryVisibility),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository.test", tfjsonpath.New("topics"), knownvalue.SetSizeExact(3)),
					},
				},
			},
		})
	})

	t.Run("creates a repository using a public template", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		testRepoName := fmt.Sprintf("%stemplate-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name        = "%s"
				description = "Terraform acceptance tests %[1]s"
				visibility  = "%s"
				template {
					owner = "%s"
					repository = "%s"
				}

			}
		`, testRepoName, testAccConf.testRepositoryVisibility, testAccConf.testPublicTemplateRepositoryOwner, testAccConf.testPublicTemplateRepository)

		resource.ParallelTest(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t); skipIfEMUEnterprise(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository.test", tfjsonpath.New("is_template"), knownvalue.Bool(false)),
					},
				},
			},
		})
	})

	t.Run("creates a repository using an org template", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		testRepoName := fmt.Sprintf("%stemplate-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name        = "%s"
				description = "Terraform acceptance tests %[1]s"
				visibility  = "%s"
				template {
					owner = "%s"
					repository = "%s"
				}

			}
		`, testRepoName, testAccConf.testRepositoryVisibility, testAccConf.owner, testAccConf.testOrgTemplateRepository)

		resource.ParallelTest(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository.test", tfjsonpath.New("is_template"), knownvalue.Bool(false)),
					},
				},
			},
		})
	})

	t.Run("archives repositories on destroy", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		testRepoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)

		config := `
resource "github_repository" "test" {
	name               = "%s"
	auto_init          = true
	archive_on_destroy = true
	archived           = %s
	visibility         = "%s"
}
`

		resource.ParallelTest(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, testRepoName, "false", testAccConf.testRepositoryVisibility),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository.test", tfjsonpath.New("archived"), knownvalue.Bool(false)),
					},
				},
				{
					Config: fmt.Sprintf(config, testRepoName, "true", testAccConf.testRepositoryVisibility),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository.test", tfjsonpath.New("archived"), knownvalue.Bool(true)),
					},
				},
			},
		})
	})

	t.Run("create_private_with_forking", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
resource "github_repository" "test" {
	name       = "%s"
	visibility = "private"

	allow_forking = true
}
`, repoName)

		resource.ParallelTest(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository.test", tfjsonpath.New("visibility"), knownvalue.StringExact("private")),
						statecheck.ExpectKnownValue("github_repository.test", tfjsonpath.New("allow_forking"), knownvalue.Bool(true)),
					},
				},
			},
		})
	})

	t.Run("create_private_without_forking", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
		resource "github_repository" "test" {
			name       = "%s"
			visibility = "private"

			allow_forking = false
		}
		`, repoName)

		resource.ParallelTest(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository.test", tfjsonpath.New("visibility"), knownvalue.StringExact("private")),
						statecheck.ExpectKnownValue("github_repository.test", tfjsonpath.New("allow_forking"), knownvalue.Bool(false)),
					},
				},
			},
		})
	})

	t.Run("create_private_with_forking_unset", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
resource "github_repository" "test" {
	name       = "%s"
	visibility = "private"
}
`, repoName)

		resource.ParallelTest(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository.test", tfjsonpath.New("visibility"), knownvalue.StringExact("private")),
						statecheck.ExpectKnownValue("github_repository.test", tfjsonpath.New("allow_forking"), knownvalue.NotNull()),
					},
				},
			},
		})
	})

	t.Run("update_public_to_private_allow_forking", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		testRepoName := fmt.Sprintf("%svisibility-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name       = "%s"
				visibility = "public"
			}
		`, testRepoName)

		configPrivate := fmt.Sprintf(`
			resource "github_repository" "test" {
				name          = "%s"
				visibility    = "private"
				allow_forking = false
			}
		`, testRepoName)

		configPrivateForking := fmt.Sprintf(`
			resource "github_repository" "test" {
				name          = "%s"
				visibility    = "private"
				allow_forking = true
			}
		`, testRepoName)

		resource.ParallelTest(t, resource.TestCase{
			PreCheck: func() {
				skipUnauthenticated(t)
				skipIfEMUEnterprise(t)
			},
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository.test", tfjsonpath.New("visibility"), knownvalue.StringExact("public")),
						statecheck.ExpectKnownValue("github_repository.test", tfjsonpath.New("allow_forking"), knownvalue.Bool(true)),
					},
				},
				{
					Config: configPrivate,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository.test", tfjsonpath.New("visibility"), knownvalue.StringExact("private")),
						statecheck.ExpectKnownValue("github_repository.test", tfjsonpath.New("allow_forking"), knownvalue.Bool(false)),
					},
				},
				{
					Config: configPrivateForking,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository.test", tfjsonpath.New("visibility"), knownvalue.StringExact("private")),
						statecheck.ExpectKnownValue("github_repository.test", tfjsonpath.New("allow_forking"), knownvalue.Bool(true)),
					},
				},
			},
		})
	})

	t.Run("create_with_vulnerability_alerts", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
		resource "github_repository" "test" {
			name       = "%s"
			visibility = "%s"

			vulnerability_alerts = true
		}
		`, repoName, testAccConf.testRepositoryVisibility)

		resource.ParallelTest(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository.test", tfjsonpath.New("vulnerability_alerts"), knownvalue.Bool(true)),
					},
				},
			},
		})
	})

	t.Run("create_without_vulnerability_alerts", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
		resource "github_repository" "test" {
			name       = "%s"
			visibility = "%s"

			vulnerability_alerts = false
		}
		`, repoName, testAccConf.testRepositoryVisibility)

		resource.ParallelTest(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository.test", tfjsonpath.New("vulnerability_alerts"), knownvalue.Bool(false)),
					},
				},
			},
		})
	})

	t.Run("create_with_vulnerability_alerts_unset", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
		resource "github_repository" "test" {
			name       = "%s"
			visibility = "%s"
		}
		`, repoName, testAccConf.testRepositoryVisibility)

		// NOTE: terraform-plugin-testing does not support asserting the nonexistence of a value
		// (no equivalent to the legacy TestCheckNoResourceAttr). We only verify creation succeeds.
		resource.ParallelTest(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
				},
			},
		})
	})

	t.Run("update_vulnerability_alerts", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)

		config := `
		resource "github_repository" "test" {
			name       = "%s"
			visibility = "%s"

			vulnerability_alerts = %t
		}
		`

		resource.ParallelTest(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, repoName, testAccConf.testRepositoryVisibility, false),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository.test", tfjsonpath.New("vulnerability_alerts"), knownvalue.Bool(false)),
					},
				},
				{
					Config: fmt.Sprintf(config, repoName, testAccConf.testRepositoryVisibility, true),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository.test", tfjsonpath.New("vulnerability_alerts"), knownvalue.Bool(true)),
					},
				},
			},
		})
	})

	t.Run("create and modify merge commit strategy without error", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		testRepoName := fmt.Sprintf("%smodify-co-str-%s", testResourcePrefix, randomID)
		mergeCommitTitle := "PR_TITLE"
		mergeCommitMessage := "BLANK"
		updatedMergeCommitTitle := "MERGE_MESSAGE"
		updatedMergeCommitMessage := "PR_TITLE"
		config := `
resource "github_repository" "test" {

		name                 = "%[1]s"
		allow_merge_commit   = true
		merge_commit_title   = "%s"
		merge_commit_message = "%s"
		visibility           = "%s"
}
`

		resource.ParallelTest(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, testRepoName, mergeCommitTitle, mergeCommitMessage, testAccConf.testRepositoryVisibility),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository.test", tfjsonpath.New("merge_commit_title"), knownvalue.StringExact(mergeCommitTitle)),
						statecheck.ExpectKnownValue("github_repository.test", tfjsonpath.New("merge_commit_message"), knownvalue.StringExact(mergeCommitMessage)),
					},
				},
				{
					Config: fmt.Sprintf(config, testRepoName, updatedMergeCommitTitle, updatedMergeCommitMessage, testAccConf.testRepositoryVisibility),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository.test", tfjsonpath.New("merge_commit_title"), knownvalue.StringExact(updatedMergeCommitTitle)),
						statecheck.ExpectKnownValue("github_repository.test", tfjsonpath.New("merge_commit_message"), knownvalue.StringExact(updatedMergeCommitMessage)),
					},
				},
			},
		})
	})

	t.Run("validate_required_fields_for_squash_merge_commit_strategy", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		testRepoName := fmt.Sprintf("%smodify-sq-str-%s", testResourcePrefix, randomID)

		config := `
		resource "github_repository" "test" {
				name                        = "%s"
				squash_merge_commit_title   = "PR_TITLE"
				squash_merge_commit_message = "PR_BODY"
				visibility                  = "%s"
				%s
		}
`
		resource.ParallelTest(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      fmt.Sprintf(config, testRepoName, testAccConf.testRepositoryVisibility, "allow_squash_merge = false"),
					ExpectError: regexp.MustCompile("allow_squash_merge is required.*"),
				},
				{
					Config: fmt.Sprintf(config, testRepoName, testAccConf.testRepositoryVisibility, ""),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository.test", tfjsonpath.New("allow_squash_merge"), knownvalue.Bool(true)),
					},
				},
			},
		},
		)
	})

	t.Run("validate_required_fields_for_merge_commit_strategy", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		testRepoName := fmt.Sprintf("%smodify-sq-str-%s", testResourcePrefix, randomID)

		config := `
		resource "github_repository" "test" {
				name                        = "%s"
				merge_commit_title   = "PR_TITLE"
				merge_commit_message = "PR_BODY"
				visibility                  = "%s"
				%s
		}
`
		resource.ParallelTest(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      fmt.Sprintf(config, testRepoName, testAccConf.testRepositoryVisibility, "allow_merge_commit = false"),
					ExpectError: regexp.MustCompile("allow_merge_commit is required.*"),
				},
				{
					Config: fmt.Sprintf(config, testRepoName, testAccConf.testRepositoryVisibility, ""),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository.test", tfjsonpath.New("allow_merge_commit"), knownvalue.Bool(true)),
					},
				},
			},
		},
		)
	})

	t.Run("create and modify squash merge commit strategy without error", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		testRepoName := fmt.Sprintf("%smodify-sq-str-%s", testResourcePrefix, randomID)
		testRepoNameAfter := fmt.Sprintf("%s-modified", testRepoName)
		squashMergeCommitTitle := "PR_TITLE"
		squashMergeCommitMessage := "PR_BODY"
		updatedSquashMergeCommitTitle := "COMMIT_OR_PR_TITLE"
		updatedSquashMergeCommitMessage := "COMMIT_MESSAGES"

		config := `
		resource "github_repository" "test" {
				name                        = "%s"
				allow_squash_merge          = true
				squash_merge_commit_title   = "%s"
				squash_merge_commit_message = "%s"
				visibility                  = "%s"
		}
`

		resource.ParallelTest(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, testRepoName, squashMergeCommitTitle, squashMergeCommitMessage, testAccConf.testRepositoryVisibility),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository.test", tfjsonpath.New("squash_merge_commit_title"), knownvalue.StringExact(squashMergeCommitTitle)),
						statecheck.ExpectKnownValue("github_repository.test", tfjsonpath.New("squash_merge_commit_message"), knownvalue.StringExact(squashMergeCommitMessage)),
					},
				},
				{
					Config: fmt.Sprintf(config, testRepoNameAfter, updatedSquashMergeCommitTitle, updatedSquashMergeCommitMessage, testAccConf.testRepositoryVisibility),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository.test", tfjsonpath.New("squash_merge_commit_title"), knownvalue.StringExact(updatedSquashMergeCommitTitle)),
						statecheck.ExpectKnownValue("github_repository.test", tfjsonpath.New("squash_merge_commit_message"), knownvalue.StringExact(updatedSquashMergeCommitMessage)),
					},
				},
			},
		})
	})

	// t.Run("create a repository with go as primary_language", func(t *testing.T) {
	// 	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
	// 	testResourceName := fmt.Sprintf("%srepo-%s", testResourcePrefix, randomID)
	// 	config := fmt.Sprintf(`
	// 		resource "github_repository" "test" {
	// 			name = "%s"
	// 			auto_init = true
	// 		}
	// 		resource "github_repository_file" "test" {
	// 			repository     = github_repository.test.name
	// 			file           = "test.go"
	// 			content        = "package main"
	// 		}
	// 	`, testResourceName)

	// 	check := resource.ComposeTestCheckFunc(
	// 		resource.TestCheckResourceAttr("github_repository.test", "primary_language", "Go"),
	// 	)

	// 	resource.ParallelTest(t, resource.TestCase{
	// 		PreCheck:          func() { skipUnauthenticated(t) },
	// 		ProviderFactories: providerFactories,
	// 		Steps: []resource.TestStep{
	// 			{
	// 				// Not doing any checks since the file needs to be created before the language can be updated
	// 				Config: config,
	// 			},
	// 			{
	// 				// Re-running the terraform will refresh the language since the go-file has been created
	// 				Config: config,
	// 				Check:  check,
	// 			},
	// 		},
	// 	})
	// })

	t.Run("manages the legacy pages feature for a repository", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		testRepoName := fmt.Sprintf("%slegacy-pages-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name         = "%s"
				auto_init    = true
				visibility   = "%s"
				pages {
					build_type = "legacy"

					source {
						branch = "main"
					}
				}
			}
			`, testRepoName, testAccConf.testRepositoryVisibility)

		resource.ParallelTest(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository.test",
							tfjsonpath.New("pages").AtSliceIndex(0).AtMapKey("source").AtSliceIndex(0).AtMapKey("branch"),
							knownvalue.StringExact("main")),
						statecheck.ExpectKnownValue("github_repository.test",
							tfjsonpath.New("pages").AtSliceIndex(0).AtMapKey("source").AtSliceIndex(0).AtMapKey("path"),
							knownvalue.StringExact("/")),
					},
				},
			},
		})
	})

	t.Run("manages the pages from workflow feature for a repository", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		testRepoName := fmt.Sprintf("%sworkflow-pages-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`
		resource "github_repository" "test" {
			name         = "%s"
			auto_init    = true
			visibility   = "%s"
			pages {
				build_type = "workflow"
			}
		}
		`, testRepoName, testAccConf.testRepositoryVisibility)

		resource.ParallelTest(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					// NOTE: terraform-plugin-testing does not support asserting the nonexistence of a value;
					// TypeList nil is represented as empty slice in JSON state
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository.test",
							tfjsonpath.New("pages").AtSliceIndex(0).AtMapKey("source"),
							knownvalue.ListSizeExact(0)),
					},
				},
			},
		})
	})

	t.Run("manages the security feature for a private repository", func(t *testing.T) {
		if !testAccConf.testAdvancedSecurity {
			t.Skip("Advanced Security is not enabled for this account")
		}

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		testRepoName := fmt.Sprintf("%ssecurity-private-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name        = "%s"
				description = "A repository created by Terraform to test security features"
				visibility  = "private"
				security_and_analysis {
					advanced_security {
						status = "enabled"
					}
					code_security {
						status = "enabled"
					}
					secret_scanning {
						status = "enabled"
					}
					secret_scanning_push_protection {
						status = "enabled"
					}
					secret_scanning_ai_detection {
						status = "enabled"
					}
					secret_scanning_non_provider_patterns {
						status = "enabled"
					}
				}
			}
			`, testRepoName)

		securityPath := tfjsonpath.New("security_and_analysis").AtSliceIndex(0)
		resource.ParallelTest(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository.test", securityPath.AtMapKey("advanced_security").AtSliceIndex(0).AtMapKey("status"), knownvalue.StringExact("enabled")),
						statecheck.ExpectKnownValue("github_repository.test", securityPath.AtMapKey("code_security").AtSliceIndex(0).AtMapKey("status"), knownvalue.StringExact("enabled")),
						statecheck.ExpectKnownValue("github_repository.test", securityPath.AtMapKey("secret_scanning").AtSliceIndex(0).AtMapKey("status"), knownvalue.StringExact("enabled")),
						statecheck.ExpectKnownValue("github_repository.test", securityPath.AtMapKey("secret_scanning_push_protection").AtSliceIndex(0).AtMapKey("status"), knownvalue.StringExact("enabled")),
						statecheck.ExpectKnownValue("github_repository.test", securityPath.AtMapKey("secret_scanning_ai_detection").AtSliceIndex(0).AtMapKey("status"), knownvalue.StringExact("enabled")),
						statecheck.ExpectKnownValue("github_repository.test", securityPath.AtMapKey("secret_scanning_non_provider_patterns").AtSliceIndex(0).AtMapKey("status"), knownvalue.StringExact("enabled")),
					},
				},
			},
		})
	})

	t.Run("manages the security feature for a public repository", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		testRepoName := fmt.Sprintf("%ssecurity-public-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name        = "%s"
				description = "A repository created by Terraform to test security features"
				visibility  = "public"
				security_and_analysis {
					secret_scanning {
						status = "enabled"
					}
					# seems like it can only be "enabled" for an organization that has purchased GHAS
					secret_scanning_push_protection {
						 status = "disabled"
					}
				}
			}
			`, testRepoName)

		securityPath := tfjsonpath.New("security_and_analysis").AtSliceIndex(0)
		resource.ParallelTest(t, resource.TestCase{
			PreCheck: func() {
				skipUnauthenticated(t)
				skipIfEMUEnterprise(t)
			},
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository.test", securityPath.AtMapKey("secret_scanning").AtSliceIndex(0).AtMapKey("status"), knownvalue.StringExact("enabled")),
						statecheck.ExpectKnownValue("github_repository.test", securityPath.AtMapKey("secret_scanning_push_protection").AtSliceIndex(0).AtMapKey("status"), knownvalue.StringExact("disabled")),
					},
				},
			},
		})
	})

	t.Run("creates repos with private visibility", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		testRepoName := fmt.Sprintf("%svisibility-private-%s", testResourcePrefix, randomID)
		config := `
			resource "github_repository" "private" {
				name       = "%s"
				visibility = "private"
			}
		`

		resource.ParallelTest(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, testRepoName),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository.private", tfjsonpath.New("visibility"), knownvalue.StringExact("private")),
					},
				},
			},
		})
	})

	t.Run("creates repos with internal visibility", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		testRepoName := fmt.Sprintf("%svisibility-internal-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name       = "%s"
				visibility = "internal"
			}
		`, testRepoName)

		resource.ParallelTest(t, resource.TestCase{
			PreCheck:          func() { skipUnlessMode(t, enterprise) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository.test", tfjsonpath.New("visibility"), knownvalue.StringExact("internal")),
					},
				},
			},
		})
	})

	t.Run("updates repos to private visibility", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		testRepoName := fmt.Sprintf("%svisibility-public-%s", testResourcePrefix, randomID)
		config := `
			resource "github_repository" "public" {
				name       = "%s"
				visibility = "%s"
				vulnerability_alerts = false
			}
		`

		resource.ParallelTest(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t); skipIfEMUEnterprise(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, testRepoName, "public"),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository.public", tfjsonpath.New("visibility"), knownvalue.StringExact("public")),
					},
				},
				{
					Config: fmt.Sprintf(config, testRepoName, "private"),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository.public", tfjsonpath.New("visibility"), knownvalue.StringExact("private")),
					},
				},
			},
		})
	})

	t.Run("updates repos to public visibility", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		testRepoName := fmt.Sprintf("%spublic-vuln-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`
				resource "github_repository" "test" {
				name       = "%s"
				visibility = "private"
			}
		`, testRepoName)

		resource.ParallelTest(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository.test", tfjsonpath.New("vulnerability_alerts"), knownvalue.Null()),
					},
				},
				{
					Config: strings.Replace(config,
						`}`,
						"vulnerability_alerts = true\n}", 1),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository.test", tfjsonpath.New("vulnerability_alerts"), knownvalue.Bool(true)),
						statecheck.ExpectKnownValue("github_repository.test", tfjsonpath.New("visibility"), knownvalue.StringExact("private")),
					},
				},
			},
		})
	})

	t.Run("updates repos to internal visibility", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		testRepoName := fmt.Sprintf("%sinternal-vuln-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name       = "%s"
				visibility = "private"
			}
		`, testRepoName)

		resource.ParallelTest(t, resource.TestCase{
			PreCheck:          func() { skipUnlessMode(t, enterprise) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository.test", tfjsonpath.New("vulnerability_alerts"), knownvalue.Null()),
					},
				},
				{
					Config: strings.Replace(config,
						`}`,
						"vulnerability_alerts = true\n}", 1),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository.test", tfjsonpath.New("vulnerability_alerts"), knownvalue.Bool(true)),
						statecheck.ExpectKnownValue("github_repository.test", tfjsonpath.New("visibility"), knownvalue.StringExact("private")),
					},
				},
			},
		})
	})

	t.Run("sets private visibility for repositories created by a template", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		testRepoName := fmt.Sprintf("%stemplate-visibility-private-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`
			resource "github_repository" "private" {
				name       = "%s"
				visibility = "private"
				template {
					owner      = "%s"
					repository = "%s"
				}
			}
		`, testRepoName, testAccConf.testPublicTemplateRepositoryOwner, testAccConf.testPublicTemplateRepository)

		resource.ParallelTest(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t); skipIfEMUEnterprise(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository.private", tfjsonpath.New("visibility"), knownvalue.StringExact("private")),
						statecheck.ExpectKnownValue("github_repository.private", tfjsonpath.New("private"), knownvalue.Bool(true)),
					},
				},
			},
		})
	})

	t.Run("create_internal_repo_from_template", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		testRepoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`
            resource "github_repository" "test" {
				name       = "%s"
				visibility = "internal"
				template {
					owner      = "%s"
					repository = "%s"
				}
			}
		`, testRepoName, testAccConf.testPublicTemplateRepositoryOwner, testAccConf.testPublicTemplateRepository)

		resource.ParallelTest(t, resource.TestCase{
			PreCheck:          func() { skipUnlessMode(t, enterprise); skipIfEMUEnterprise(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository.test", tfjsonpath.New("visibility"), knownvalue.StringExact("internal")),
						statecheck.ExpectKnownValue("github_repository.test", tfjsonpath.New("private"), knownvalue.Bool(true)),
					},
				},
			},
		})
	})

	t.Run("check_web_commit_signoff_required_enabled", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		testRepoName := fmt.Sprintf("%scommit-signoff-%s", testResourcePrefix, randomID)
		config := `
			resource "github_repository" "test" {
				name                        = "%s"
				auto_init                   = true
				web_commit_signoff_required = %s
			}
		`

		resource.ParallelTest(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, testRepoName, "true"),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository.test", tfjsonpath.New("web_commit_signoff_required"), knownvalue.Bool(true)),
					},
				},
			},
		})
	})

	t.Run("check_web_commit_signoff_required_disabled", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		testRepoName := fmt.Sprintf("%scommit-signoff-%s", testResourcePrefix, randomID)
		config := `
			resource "github_repository" "test" {
				name                        = "%s"
				auto_init                   = true
				web_commit_signoff_required = %s
			}
		`

		resource.ParallelTest(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, testRepoName, "false"),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository.test", tfjsonpath.New("web_commit_signoff_required"), knownvalue.Bool(false)),
					},
				},
			},
		})
	})

	t.Run("check_web_commit_signoff_required_not_set", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		testRepoName := fmt.Sprintf("%scommit-signoff-%s", testResourcePrefix, randomID)
		config := `
			resource "github_repository" "test" {
				name                        = "%s"
				auto_init                   = true
			}
		`

		resource.ParallelTest(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, testRepoName),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository.test", tfjsonpath.New("web_commit_signoff_required"), knownvalue.Bool(false)),
					},
				},
			},
		})
	})

	t.Run("check_web_commit_signoff_required_organization_enabled_but_not_set", func(t *testing.T) {
		t.Skip("This test should be run manually after confirming that the test organization has 'Require contributors to sign off on web-based commits' enabled under Organizations -> Settings -> Repository -> Repository defaults.")

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		testRepoName := fmt.Sprintf("%scommit-signoff-%s", testResourcePrefix, randomID)

		config := `
			resource "github_repository" "test" {
				name        = "%s"
				description = "%s"
				visibility  = "private"
			}
		`

		resource.ParallelTest(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, testRepoName, "foo"),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository.test", tfjsonpath.New("web_commit_signoff_required"), knownvalue.Bool(true)),
					},
				},
				{
					Config: fmt.Sprintf(config, testRepoName, "bar"),
				},
			},
		})
	})

	t.Run("check_allow_forking_not_set", func(t *testing.T) {
		t.Skip("This test should be run manually after confirming that the test organization has been correctly configured to disable setting forking at the repo level.")

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		testRepoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)

		config := `
resource "github_repository" "private" {
	name        = "%s"
	description = "%s"
	visibility  = "private"
}
`

		resource.ParallelTest(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, testRepoName, "foo"),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository.private", tfjsonpath.New("allow_forking"), knownvalue.Bool(false)),
					},
				},
				{
					Config: fmt.Sprintf(config, testRepoName, "bar"),
				},
			},
		})
	})

	t.Run("check_vulnerability_alerts_not_set", func(t *testing.T) {
		t.Skip("This test should be run manually after confirming that the test organization has been correctly configured to disable setting vulnerability alerts at the repo level.")

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		testRepoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)

		config := `
resource "github_repository" "private" {
	name        = "%s"
	description = "%s"
	visibility  = "public"
}
`

		resource.ParallelTest(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, testRepoName, "foo"),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository.private", tfjsonpath.New("vulnerability_alerts"), knownvalue.Bool(true)),
					},
				},
				{
					Config: fmt.Sprintf(config, testRepoName, "bar"),
				},
			},
		})
	})
}

func Test_expandPages(t *testing.T) {
	t.Run("expand Pages configuration with workflow", func(t *testing.T) {
		input := []any{map[string]any{
			"build_type": "workflow",
			"source":     []any{map[string]any{}},
		}}

		pages := expandPages(input)
		if pages == nil {
			t.Fatal("pages is nil")
		}
		if pages.GetBuildType() != "workflow" {
			t.Errorf("got %q; want %q", pages.GetBuildType(), "workflow")
		}
		if pages.GetSource().GetBranch() != "main" {
			t.Errorf("got %q; want %q", pages.GetSource().GetBranch(), "main")
		}
	})

	t.Run("expand Pages configuration with source", func(t *testing.T) {
		input := []any{map[string]any{
			"build_type": "legacy",
			"source": []any{map[string]any{
				"branch": "main",
				"path":   "/docs",
			}},
		}}

		pages := expandPages(input)
		if pages == nil {
			t.Fatal("pages is nil")
		}
		if pages.GetBuildType() != "legacy" {
			t.Errorf("got %q; want %q", pages.GetBuildType(), "legacy")
		}
		if pages.GetSource().GetBranch() != "main" {
			t.Errorf("got %q; want %q", pages.GetSource().GetBranch(), "main")
		}
		if pages.GetSource().GetPath() != "/docs" {
			t.Errorf("got %q; want %q", pages.GetSource().GetPath(), "/docs")
		}
	})

	t.Run("forks a repository without error", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		testRepoName := fmt.Sprintf("%sfork-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`
				resource "github_repository" "forked" {
					name         = "%s"
					description  = "Terraform acceptance test - forked repository %[1]s"
					fork         = true
					source_owner = "integrations"
					source_repo  = "terraform-provider-github"
				}
		 	`, testRepoName)

		resource.ParallelTest(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository.forked", tfjsonpath.New("fork"), knownvalue.StringExact("true")),
						statecheck.ExpectKnownValue("github_repository.forked", tfjsonpath.New("html_url"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_repository.forked", tfjsonpath.New("ssh_clone_url"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_repository.forked", tfjsonpath.New("git_clone_url"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_repository.forked", tfjsonpath.New("http_clone_url"), knownvalue.NotNull()),
					},
				},
			},
		})
	})

	t.Run("can update forked repository properties", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		testRepoName := fmt.Sprintf("%sfork-update-%s", testResourcePrefix, randomID)
		initialConfig := fmt.Sprintf(`
				resource "github_repository" "forked_update" {
					name         = "%s"
					description  = "Initial description for forked repo"
					fork         = true
					source_owner = "integrations"
					source_repo  = "terraform-provider-github"
					has_wiki     = true
					has_issues   = false
				}
		 `, testRepoName)

		updatedConfig := fmt.Sprintf(`
				resource "github_repository" "forked_update" {
					name         = "%s"
					description  = "Updated description for forked repo"
					fork         = true
					source_owner = "integrations"
					source_repo  = "terraform-provider-github"
					has_wiki     = false
					has_issues   = true
				}
		 `, testRepoName)

		resource.ParallelTest(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: initialConfig,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository.forked_update", tfjsonpath.New("description"), knownvalue.StringExact("Initial description for forked repo")),
						statecheck.ExpectKnownValue("github_repository.forked_update", tfjsonpath.New("has_wiki"), knownvalue.Bool(true)),
						statecheck.ExpectKnownValue("github_repository.forked_update", tfjsonpath.New("has_issues"), knownvalue.Bool(false)),
					},
				},
				{
					Config: updatedConfig,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository.forked_update", tfjsonpath.New("description"), knownvalue.StringExact("Updated description for forked repo")),
						statecheck.ExpectKnownValue("github_repository.forked_update", tfjsonpath.New("has_wiki"), knownvalue.Bool(false)),
						statecheck.ExpectKnownValue("github_repository.forked_update", tfjsonpath.New("has_issues"), knownvalue.Bool(true)),
					},
				},
				{
					ResourceName:            "github_repository.forked_update",
					ImportState:             true,
					ImportStateVerify:       true,
					ImportStateVerifyIgnore: []string{"auto_init", "vulnerability_alerts", "ignore_vulnerability_alerts_during_read"},
				},
			},
		})
	})
}

func TestGithubRepositoryTopicPassesValidation(t *testing.T) {
	resource := resourceGithubRepository()
	schema := resource.Schema["topics"].Elem.(*schema.Schema)
	diags := schema.ValidateDiagFunc("ef69e1a3-66be-40ca-bb62-4f36186aa292", cty.Path{cty.GetAttrStep{Name: "topics"}})
	if diags.HasError() {
		t.Error(fmt.Errorf("unexpected topics validation failure: %s", diags[0].Summary))
	}
}

func TestGithubRepositoryTopicFailsValidationWhenOverMaxCharacters(t *testing.T) {
	resource := resourceGithubRepository()
	schema := resource.Schema["topics"].Elem.(*schema.Schema)

	diags := schema.ValidateDiagFunc(strings.Repeat("a", 51), cty.Path{cty.GetAttrStep{Name: "topics"}})
	if len(diags) != 1 {
		t.Error(fmt.Errorf("unexpected number of topic validation failures; expected=1; actual=%d", len(diags)))
	}
	expectedFailure := "invalid value for topics (must include only lowercase alphanumeric characters or hyphens and cannot start with a hyphen and consist of 50 characters or less)"
	actualFailure := diags[0].Summary
	if expectedFailure != actualFailure {
		t.Error(fmt.Errorf("unexpected topic validation failure; expected=%s; action=%s", expectedFailure, actualFailure))
	}
}

type resourceDataLike map[string]any

func (d resourceDataLike) GetOk(key string) (any, bool) {
	v, ok := d[key]
	return v, ok
}

func TestResourceGithubParseFullName(t *testing.T) {
	t.Run("parses valid full name", func(t *testing.T) {
		o := "moyorg"
		r := "myrepo"

		org, repo, ok := resourceGithubParseFullName(resourceDataLike(map[string]any{"full_name": fmt.Sprintf("%s/%s", o, r)}))
		if !ok {
			t.Error("expected ok to be true, got false")
		}
		if org != o {
			t.Errorf("unexpected org (wanted %s, got %s)", o, org)
		}
		if repo != r {
			t.Errorf("unexpected repo (wanted %s, got %s)", r, repo)
		}
	})

	t.Run("handles missing full name", func(t *testing.T) {
		_, _, ok := resourceGithubParseFullName(resourceDataLike(map[string]any{}))
		if ok {
			t.Fatal("expected ok to be false, got true")
		}
	})

	t.Run("handles malformed full name", func(t *testing.T) {
		_, _, ok := resourceGithubParseFullName(resourceDataLike(map[string]any{"full_name": "malformed"}))
		if ok {
			t.Fatal("expected ok to be false, got true")
		}
	})
}

func TestGithubRepositoryNameFailsValidationWhenOverMaxCharacters(t *testing.T) {
	resource := resourceGithubRepository()
	schema := resource.Schema["name"]

	diags := schema.ValidateDiagFunc(strings.Repeat("a", 101), cty.GetAttrPath("name"))
	if len(diags) != 1 {
		t.Error(fmt.Errorf("unexpected number of name validation failures; expected=1; actual=%d", len(diags)))
	}
	expectedFailure := "invalid value for name (must include only alphanumeric characters, underscores or hyphens and consist of 100 characters or less)"
	actualFailure := diags[0].Summary
	if expectedFailure != actualFailure {
		t.Error(fmt.Errorf("unexpected name validation failure; expected=%s; action=%s", expectedFailure, actualFailure))
	}
}

func TestGithubRepositoryNameFailsValidationWithSpace(t *testing.T) {
	resource := resourceGithubRepository()
	schema := resource.Schema["name"]

	diags := schema.ValidateDiagFunc("test space", cty.GetAttrPath("name"))
	if len(diags) != 1 {
		t.Error(fmt.Errorf("unexpected number of name validation failures; expected=1; actual=%d", len(diags)))
	}
	expectedFailure := "invalid value for name (must include only alphanumeric characters, underscores or hyphens and consist of 100 characters or less)"
	actualFailure := diags[0].Summary
	if expectedFailure != actualFailure {
		t.Error(fmt.Errorf("unexpected name validation failure; expected=%s; action=%s", expectedFailure, actualFailure))
	}
}
