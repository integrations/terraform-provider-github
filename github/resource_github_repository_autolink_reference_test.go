package github

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/compare"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubRepositoryAutolinkReference(t *testing.T) {
	t.Parallel()

	t.Run("creates_repository_autolink_reference_without_error", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)
		config := fmt.Sprintf(`
			resource "github_repository_autolink_reference" "autolink_default" {
				repository = "%[1]s"

				key_prefix          = "TEST1-"
				target_url_template = "https://example.com/TEST-<num>"
			}

			resource "github_repository_autolink_reference" "autolink_alphanumeric" {
				repository = "%[1]s"

				key_prefix          = "TEST2-"
				target_url_template = "https://example.com/TEST-<num>"
				is_alphanumeric     = true
			}

			resource "github_repository_autolink_reference" "autolink_numeric" {
				repository = "%[1]s"

				key_prefix          = "TEST3-"
				target_url_template = "https://example.com/TEST-<num>"
				is_alphanumeric     = false
			}

			resource "github_repository_autolink_reference" "autolink_with_port" {
				repository = "%[1]s"

				key_prefix          = "TEST4-"
				target_url_template = "https://example.com:8443/TEST-<num>"
			}
		`, repo.GetName())

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_autolink_reference.autolink_default", tfjsonpath.New("repository_id"), knownvalue.NotNull()),
					},
				},
			},
		})
	})

	t.Run("imports_repository_autolink_reference_without_error", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)
		config := fmt.Sprintf(`
			resource "github_repository_autolink_reference" "autolink_default" {
				repository = "%[1]s"

				key_prefix          = "TEST1-"
				target_url_template = "https://example.com/TEST-<num>"
			}

			resource "github_repository_autolink_reference" "autolink_alphanumeric" {
				repository = "%[1]s"

				key_prefix          = "TEST2-"
				target_url_template = "https://example.com/TEST-<num>"
				is_alphanumeric     = true
			}

			resource "github_repository_autolink_reference" "autolink_numeric" {
				repository = "%[1]s"

				key_prefix          = "TEST3-"
				target_url_template = "https://example.com/TEST-<num>"
				is_alphanumeric     = false
			}

			resource "github_repository_autolink_reference" "autolink_with_port" {
				repository = "%[1]s"

				key_prefix          = "TEST4-"
				target_url_template = "https://example.com:8443/TEST-<num>"
			}
		`, repo.GetName())

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_autolink_reference.autolink_default", tfjsonpath.New("repository_id"), knownvalue.NotNull()),
					},
				},
				// autolink_default
				{
					ResourceName:            "github_repository_autolink_reference.autolink_default",
					ImportState:             true,
					ImportStateVerify:       true,
					ImportStateVerifyIgnore: []string{"etag"},
					ImportStateIdPrefix:     fmt.Sprintf("%s"+idSeparator, repo.GetName()),
				},
				// autolink_alphanumeric
				{
					ResourceName:            "github_repository_autolink_reference.autolink_alphanumeric",
					ImportState:             true,
					ImportStateVerify:       true,
					ImportStateVerifyIgnore: []string{"etag"},
					ImportStateIdPrefix:     fmt.Sprintf("%s"+idSeparator, repo.GetName()),
				},
				// autolink_numeric
				{
					ResourceName:            "github_repository_autolink_reference.autolink_numeric",
					ImportState:             true,
					ImportStateVerify:       true,
					ImportStateVerifyIgnore: []string{"etag"},
					ImportStateIdPrefix:     fmt.Sprintf("%s"+idSeparator, repo.GetName()),
				},
				// autolink_with_port
				{
					ResourceName:            "github_repository_autolink_reference.autolink_with_port",
					ImportState:             true,
					ImportStateVerify:       true,
					ImportStateVerifyIgnore: []string{"etag"},
					ImportStateIdPrefix:     fmt.Sprintf("%s"+idSeparator, repo.GetName()),
				},
			},
		})
	})

	t.Run("imports_repository_autolink_reference_by_key_prefix_without_error", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)
		config := fmt.Sprintf(`
			resource "github_repository_autolink_reference" "autolink" {
			  repository = "%[1]s"

			  key_prefix 		  = "OOF-"
			  target_url_template = "https://awesome.com/find/OOF-<num>"
			}
		`, repo.GetName())

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
				},
				{
					ResourceName:            "github_repository_autolink_reference.autolink",
					ImportState:             true,
					ImportStateVerify:       true,
					ImportStateVerifyIgnore: []string{"etag"},
					ImportStateIdFunc: func(*terraform.State) (string, error) {
						return buildID(repo.GetName(), "OOF-")
					},
				},
				{
					ResourceName: "github_repository_autolink_reference.autolink",
					ImportState:  true,
					ImportStateIdFunc: func(*terraform.State) (string, error) {
						return buildID(repo.GetName(), "OCTOCAT-")
					},
					ImportStateVerify:       true,
					ImportStateVerifyIgnore: []string{"etag"},
					ExpectError:             regexp.MustCompile(`cannot find autolink reference`),
				},
			},
		})
	})

	t.Run("deletes_repository_autolink_reference_without_error", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)
		config := fmt.Sprintf(`
			resource "github_repository_autolink_reference" "autolink_default" {
				repository = "%s"

				key_prefix          = "TEST1-"
				target_url_template = "https://example.com/TEST-<num>"
			}
		`, repo.GetName())

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

	t.Run("rejects_key_prefix_ending_with_a_digit_at_plan_time", func(t *testing.T) {
		t.Parallel()

		config := `
			resource "github_repository_autolink_reference" "autolink_invalid" {
				repository          = "some-repo"
				key_prefix          = "PTFY25"
				target_url_template = "https://example.com/<num>"
			}
		`

		resource.UnitTest(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      config,
					ExpectError: regexp.MustCompile(`must not end with a number`),
				},
			},
		})
	})
	t.Run("should_not_recreate_autolink_reference_when_repository_is_renamed", func(t *testing.T) {
		repo := mustCreateTestRepository(t)
		repoNameRenamed := fmt.Sprintf("%s-renamed", repo.GetName())
		mustRenameTestRepository(t, repo, repoNameRenamed)
		const configStr = `
	resource "github_repository_autolink_reference" "autolink_default" {
		repository = "%s"

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
					Config: fmt.Sprintf(configStr, repo.GetName()),
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
