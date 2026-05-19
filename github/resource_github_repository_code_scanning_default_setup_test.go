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

func testAccCodeScanningDefaultSetupConfig(repoName, extraAttrs string) string {
	return fmt.Sprintf(`
		resource "github_repository" "test" {
			name       = "%s"
			visibility = "public"
			auto_init  = true
		}

		resource "github_repository_code_scanning_default_setup" "test" {
			repository = github_repository.test.name
			%s
		}
	`, repoName, extraAttrs)
}

func TestAccGithubRepositoryCodeScanningDefaultSetup(t *testing.T) {
	t.Run("configures with explicit query suite and languages", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-code-scanning-%s", testResourcePrefix, randomID)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: testAccCodeScanningDefaultSetupConfig(repoName, `
						state       = "configured"
						query_suite = "extended"
						languages   = ["javascript-typescript", "python"]
					`),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_code_scanning_default_setup.test",
							tfjsonpath.New("state"), knownvalue.StringExact("configured")),
						statecheck.ExpectKnownValue("github_repository_code_scanning_default_setup.test",
							tfjsonpath.New("query_suite"), knownvalue.StringExact("extended")),
						statecheck.ExpectKnownValue("github_repository_code_scanning_default_setup.test",
							tfjsonpath.New("languages"), knownvalue.SetSizeExact(2)),
					},
				},
			},
		})
	})

	t.Run("is idempotent when already not-configured", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-code-scanning-%s", testResourcePrefix, randomID)
		config := testAccCodeScanningDefaultSetupConfig(repoName, `state = "not-configured"`)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_code_scanning_default_setup.test",
							tfjsonpath.New("state"), knownvalue.StringExact("not-configured")),
					},
				},
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_code_scanning_default_setup.test",
							tfjsonpath.New("state"), knownvalue.StringExact("not-configured")),
					},
				},
			},
		})
	})

	t.Run("imports code scanning default setup", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-code-scanning-%s", testResourcePrefix, randomID)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: testAccCodeScanningDefaultSetupConfig(repoName, `state = "configured"`),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_code_scanning_default_setup.test",
							tfjsonpath.New("repository"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_repository_code_scanning_default_setup.test",
							tfjsonpath.New("repository_id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_repository_code_scanning_default_setup.test",
							tfjsonpath.New("state"), knownvalue.StringExact("configured")),
						statecheck.ExpectKnownValue("github_repository_code_scanning_default_setup.test",
							tfjsonpath.New("query_suite"), knownvalue.NotNull()),
					},
				},
				{
					ResourceName:      "github_repository_code_scanning_default_setup.test",
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})

	t.Run("specifies languages not present in repo without error", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-code-scanning-%s", testResourcePrefix, randomID)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: testAccCodeScanningDefaultSetupConfig(repoName, `
						state     = "configured"
						languages = ["go", "java-kotlin"]
					`),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_code_scanning_default_setup.test",
							tfjsonpath.New("state"), knownvalue.StringExact("configured")),
					},
				},
			},
		})
	})

	t.Run("prevents configuring on archived repository", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-code-scanning-%s", testResourcePrefix, randomID)
		repoConfig := `
			resource "github_repository" "test" {
				name       = "%s"
				visibility = "public"
				auto_init  = true
				archived   = %t
			}
			%s
		`
		codeScanningConfig := `
			resource "github_repository_code_scanning_default_setup" "test" {
				repository = github_repository.test.name
				state      = "configured"
			}
		`

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(repoConfig, repoName, false, ""),
				},
				{
					Config:      fmt.Sprintf(repoConfig, repoName, true, codeScanningConfig),
					ExpectError: regexp.MustCompile("is archived"),
				},
			},
		})
	})

	t.Run("full lifecycle: configure, update query suite, unconfigure", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-code-scanning-%s", testResourcePrefix, randomID)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: testAccCodeScanningDefaultSetupConfig(repoName, `
						state       = "configured"
						query_suite = "default"
					`),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_code_scanning_default_setup.test",
							tfjsonpath.New("state"), knownvalue.StringExact("configured")),
						statecheck.ExpectKnownValue("github_repository_code_scanning_default_setup.test",
							tfjsonpath.New("query_suite"), knownvalue.StringExact("default")),
					},
				},
				{
					Config: testAccCodeScanningDefaultSetupConfig(repoName, `
						state       = "configured"
						query_suite = "extended"
					`),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_code_scanning_default_setup.test",
							tfjsonpath.New("state"), knownvalue.StringExact("configured")),
						statecheck.ExpectKnownValue("github_repository_code_scanning_default_setup.test",
							tfjsonpath.New("query_suite"), knownvalue.StringExact("extended")),
					},
				},
				{
					Config: testAccCodeScanningDefaultSetupConfig(repoName, `state = "not-configured"`),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_code_scanning_default_setup.test",
							tfjsonpath.New("state"), knownvalue.StringExact("not-configured")),
					},
				},
			},
		})
	})
}
