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

func TestAccGithubRepositoryAutolinkReference(t *testing.T) {
	t.Run("creates repository autolink reference without error", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-autolink-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name        = "%s"
				description = "Test autolink creation"
			}

			resource "github_repository_autolink_reference" "autolink_default" {
				repository = github_repository.test.name

				key_prefix          = "TEST1-"
				target_url_template = "https://example.com/TEST-<num>"
			}

			resource "github_repository_autolink_reference" "autolink_alphanumeric" {
				repository = github_repository.test.name

				key_prefix          = "TEST2-"
				target_url_template = "https://example.com/TEST-<num>"
				is_alphanumeric     = true
			}

			resource "github_repository_autolink_reference" "autolink_numeric" {
				repository = github_repository.test.name

				key_prefix          = "TEST3-"
				target_url_template = "https://example.com/TEST-<num>"
				is_alphanumeric     = false
			}

			resource "github_repository_autolink_reference" "autolink_with_port" {
				repository = github_repository.test.name

				key_prefix          = "TEST4-"
				target_url_template = "https://example.com:8443/TEST-<num>"
			}
		`, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						// autolink_default
						statecheck.ExpectKnownValue("github_repository_autolink_reference.autolink_default", tfjsonpath.New("key_prefix"), knownvalue.StringExact("TEST1-")),
						statecheck.ExpectKnownValue("github_repository_autolink_reference.autolink_default", tfjsonpath.New("target_url_template"), knownvalue.StringExact("https://example.com/TEST-<num>")),
						statecheck.ExpectKnownValue("github_repository_autolink_reference.autolink_default", tfjsonpath.New("is_alphanumeric"), knownvalue.Bool(true)),
						statecheck.ExpectKnownValue("github_repository_autolink_reference.autolink_default", tfjsonpath.New("repository_id"), knownvalue.NotNull()),
						// autolink_alphanumeric
						statecheck.ExpectKnownValue("github_repository_autolink_reference.autolink_alphanumeric", tfjsonpath.New("key_prefix"), knownvalue.StringExact("TEST2-")),
						statecheck.ExpectKnownValue("github_repository_autolink_reference.autolink_alphanumeric", tfjsonpath.New("target_url_template"), knownvalue.StringExact("https://example.com/TEST-<num>")),
						statecheck.ExpectKnownValue("github_repository_autolink_reference.autolink_alphanumeric", tfjsonpath.New("is_alphanumeric"), knownvalue.Bool(true)),
						statecheck.ExpectKnownValue("github_repository_autolink_reference.autolink_alphanumeric", tfjsonpath.New("repository_id"), knownvalue.NotNull()),
						// autolink_numeric
						statecheck.ExpectKnownValue("github_repository_autolink_reference.autolink_numeric", tfjsonpath.New("key_prefix"), knownvalue.StringExact("TEST3-")),
						statecheck.ExpectKnownValue("github_repository_autolink_reference.autolink_numeric", tfjsonpath.New("target_url_template"), knownvalue.StringExact("https://example.com/TEST-<num>")),
						statecheck.ExpectKnownValue("github_repository_autolink_reference.autolink_numeric", tfjsonpath.New("is_alphanumeric"), knownvalue.Bool(false)),
						statecheck.ExpectKnownValue("github_repository_autolink_reference.autolink_numeric", tfjsonpath.New("repository_id"), knownvalue.NotNull()),
						// autolink_with_port
						statecheck.ExpectKnownValue("github_repository_autolink_reference.autolink_with_port", tfjsonpath.New("key_prefix"), knownvalue.StringExact("TEST4-")),
						statecheck.ExpectKnownValue("github_repository_autolink_reference.autolink_with_port", tfjsonpath.New("target_url_template"), knownvalue.StringExact("https://example.com:8443/TEST-<num>")),
						statecheck.ExpectKnownValue("github_repository_autolink_reference.autolink_with_port", tfjsonpath.New("is_alphanumeric"), knownvalue.Bool(true)),
						statecheck.ExpectKnownValue("github_repository_autolink_reference.autolink_with_port", tfjsonpath.New("repository_id"), knownvalue.NotNull()),
					},
				},
			},
		})
	})

	t.Run("imports repository autolink reference without error", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-autolink-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name        = "%s"
				description = "Test autolink creation"
			}

			resource "github_repository_autolink_reference" "autolink_default" {
				repository = github_repository.test.name

				key_prefix          = "TEST1-"
				target_url_template = "https://example.com/TEST-<num>"
			}

			resource "github_repository_autolink_reference" "autolink_alphanumeric" {
				repository = github_repository.test.name

				key_prefix          = "TEST2-"
				target_url_template = "https://example.com/TEST-<num>"
				is_alphanumeric     = true
			}

			resource "github_repository_autolink_reference" "autolink_numeric" {
				repository = github_repository.test.name

				key_prefix          = "TEST3-"
				target_url_template = "https://example.com/TEST-<num>"
				is_alphanumeric     = false
			}

			resource "github_repository_autolink_reference" "autolink_with_port" {
				repository = github_repository.test.name

				key_prefix          = "TEST4-"
				target_url_template = "https://example.com:8443/TEST-<num>"
			}
		`, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						// autolink_default
						statecheck.ExpectKnownValue("github_repository_autolink_reference.autolink_default", tfjsonpath.New("key_prefix"), knownvalue.StringExact("TEST1-")),
						statecheck.ExpectKnownValue("github_repository_autolink_reference.autolink_default", tfjsonpath.New("target_url_template"), knownvalue.StringExact("https://example.com/TEST-<num>")),
						statecheck.ExpectKnownValue("github_repository_autolink_reference.autolink_default", tfjsonpath.New("is_alphanumeric"), knownvalue.Bool(true)),
						statecheck.ExpectKnownValue("github_repository_autolink_reference.autolink_default", tfjsonpath.New("repository_id"), knownvalue.NotNull()),
						// autolink_alphanumeric
						statecheck.ExpectKnownValue("github_repository_autolink_reference.autolink_alphanumeric", tfjsonpath.New("key_prefix"), knownvalue.StringExact("TEST2-")),
						statecheck.ExpectKnownValue("github_repository_autolink_reference.autolink_alphanumeric", tfjsonpath.New("target_url_template"), knownvalue.StringExact("https://example.com/TEST-<num>")),
						statecheck.ExpectKnownValue("github_repository_autolink_reference.autolink_alphanumeric", tfjsonpath.New("is_alphanumeric"), knownvalue.Bool(true)),
						statecheck.ExpectKnownValue("github_repository_autolink_reference.autolink_alphanumeric", tfjsonpath.New("repository_id"), knownvalue.NotNull()),
						// autolink_numeric
						statecheck.ExpectKnownValue("github_repository_autolink_reference.autolink_numeric", tfjsonpath.New("key_prefix"), knownvalue.StringExact("TEST3-")),
						statecheck.ExpectKnownValue("github_repository_autolink_reference.autolink_numeric", tfjsonpath.New("target_url_template"), knownvalue.StringExact("https://example.com/TEST-<num>")),
						statecheck.ExpectKnownValue("github_repository_autolink_reference.autolink_numeric", tfjsonpath.New("is_alphanumeric"), knownvalue.Bool(false)),
						statecheck.ExpectKnownValue("github_repository_autolink_reference.autolink_numeric", tfjsonpath.New("repository_id"), knownvalue.NotNull()),
						// autolink_with_port
						statecheck.ExpectKnownValue("github_repository_autolink_reference.autolink_with_port", tfjsonpath.New("key_prefix"), knownvalue.StringExact("TEST4-")),
						statecheck.ExpectKnownValue("github_repository_autolink_reference.autolink_with_port", tfjsonpath.New("target_url_template"), knownvalue.StringExact("https://example.com:8443/TEST-<num>")),
						statecheck.ExpectKnownValue("github_repository_autolink_reference.autolink_with_port", tfjsonpath.New("is_alphanumeric"), knownvalue.Bool(true)),
						statecheck.ExpectKnownValue("github_repository_autolink_reference.autolink_with_port", tfjsonpath.New("repository_id"), knownvalue.NotNull()),
					},
				},
				// autolink_default
				{
					ResourceName:        "github_repository_autolink_reference.autolink_default",
					ImportState:         true,
					ImportStateVerify:   true,
					ImportStateIdPrefix: fmt.Sprintf("%s/", repoName),
				},
				// autolink_alphanumeric
				{
					ResourceName:        "github_repository_autolink_reference.autolink_alphanumeric",
					ImportState:         true,
					ImportStateVerify:   true,
					ImportStateIdPrefix: fmt.Sprintf("%s/", repoName),
				},
				// autolink_numeric
				{
					ResourceName:        "github_repository_autolink_reference.autolink_numeric",
					ImportState:         true,
					ImportStateVerify:   true,
					ImportStateIdPrefix: fmt.Sprintf("%s/", repoName),
				},
				// autolink_with_port
				{
					ResourceName:        "github_repository_autolink_reference.autolink_with_port",
					ImportState:         true,
					ImportStateVerify:   true,
					ImportStateIdPrefix: fmt.Sprintf("%s/", repoName),
				},
			},
		})
	})

	t.Run("imports repository autolink reference by key prefix without error", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-autolink-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`
			resource "github_repository" "oof" {
			  name         = "%s"
			  description  = "Test autolink creation"
			}

			resource "github_repository_autolink_reference" "autolink" {
			  repository = github_repository.oof.name

			  key_prefix 		  = "OOF-"
			  target_url_template = "https://awesome.com/find/OOF-<num>"
			}
		`, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
				},
				{
					ResourceName:      "github_repository_autolink_reference.autolink",
					ImportState:       true,
					ImportStateVerify: true,
					ImportStateId:     fmt.Sprintf("%s/OOF-", repoName),
				},
				{
					ResourceName:  "github_repository_autolink_reference.autolink",
					ImportState:   true,
					ImportStateId: fmt.Sprintf("%s/OCTOCAT-", repoName),
					ExpectError:   regexp.MustCompile(`cannot find autolink reference`),
				},
			},
		})
	})

	t.Run("deletes repository autolink reference without error", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-autolink-%s", testResourcePrefix, randomID)
		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name        = "%s"
				description = "Test autolink creation"
			}

			resource "github_repository_autolink_reference" "autolink_default" {
				repository = github_repository.test.name

				key_prefix          = "TEST1-"
				target_url_template = "https://example.com/TEST-<num>"
			}
		`, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:  config,
					Destroy: true,
				},
			},
		})
	})
	t.Run("should not recreate autolink reference when repository is renamed", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-rename-%s", testResourcePrefix, randomID)
		repoNameRenamed := fmt.Sprintf("%srepo-renamed-%s", testResourcePrefix, randomID)
		const configStr = `
	resource "github_repository" "test" {
		name        = "%s"
		description = "Test autolink creation"
	}

	resource "github_repository_autolink_reference" "autolink_default" {
		repository = github_repository.test.name

		key_prefix          = "TEST1-"
		target_url_template = "https://example.com/TEST-<num>"
	}
`

		repoIdChangeCheck := statecheck.CompareValue(compare.ValuesSame())
		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(configStr, repoName),
					ConfigStateChecks: []statecheck.StateCheck{
						repoIdChangeCheck.AddStateValue("github_repository_autolink_reference.autolink_default", tfjsonpath.New("repository_id")),
					},
				},
				{
					Config: fmt.Sprintf(configStr, repoNameRenamed),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_repository_autolink_reference.autolink_default", plancheck.ResourceActionUpdate),
						},
					},
					ConfigStateChecks: []statecheck.StateCheck{
						repoIdChangeCheck.AddStateValue("github_repository_autolink_reference.autolink_default", tfjsonpath.New("repository_id")),
						statecheck.ExpectKnownValue("github_repository_autolink_reference.autolink_default", tfjsonpath.New("repository"), knownvalue.StringExact(repoNameRenamed)),
					},
				},
			},
		})
	})
}
