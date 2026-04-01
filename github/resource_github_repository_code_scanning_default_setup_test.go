package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
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
	t.Run("configures and unconfigures code scanning default setup", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-code-scanning-%s", testResourcePrefix, randomID)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: testAccCodeScanningDefaultSetupConfig(repoName, `state = "configured"`),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr(
							"github_repository_code_scanning_default_setup.test", "state", "configured"),
						resource.TestCheckResourceAttrSet(
							"github_repository_code_scanning_default_setup.test", "query_suite"),
					),
				},
				{
					Config: testAccCodeScanningDefaultSetupConfig(repoName, `state = "not-configured"`),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr(
							"github_repository_code_scanning_default_setup.test", "state", "not-configured"),
					),
				},
			},
		})
	})

	t.Run("configures with extended query suite", func(t *testing.T) {
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
					`),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr(
							"github_repository_code_scanning_default_setup.test", "state", "configured"),
						resource.TestCheckResourceAttr(
							"github_repository_code_scanning_default_setup.test", "query_suite", "extended"),
					),
				},
			},
		})
	})

	t.Run("configures with explicit languages", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-code-scanning-%s", testResourcePrefix, randomID)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: testAccCodeScanningDefaultSetupConfig(repoName, `
						state     = "configured"
						languages = ["javascript-typescript", "python"]
					`),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr(
							"github_repository_code_scanning_default_setup.test", "state", "configured"),
						resource.TestCheckResourceAttr(
							"github_repository_code_scanning_default_setup.test", "languages.#", "2"),
					),
				},
			},
		})
	})

	t.Run("updates query suite without changing state", func(t *testing.T) {
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
					Check: resource.TestCheckResourceAttr(
						"github_repository_code_scanning_default_setup.test", "query_suite", "default"),
				},
				{
					Config: testAccCodeScanningDefaultSetupConfig(repoName, `
						state       = "configured"
						query_suite = "extended"
					`),
					Check: resource.TestCheckResourceAttr(
						"github_repository_code_scanning_default_setup.test", "query_suite", "extended"),
				},
			},
		})
	})

	t.Run("is idempotent when already not-configured", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-code-scanning-%s", testResourcePrefix, randomID)
		config := testAccCodeScanningDefaultSetupConfig(repoName, `state = "not-configured"`)

		check := resource.TestCheckResourceAttr(
			"github_repository_code_scanning_default_setup.test", "state", "not-configured")

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{Config: config, Check: check},
				{Config: config, Check: check},
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
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttrSet(
							"github_repository_code_scanning_default_setup.test", "repository"),
						resource.TestCheckResourceAttr(
							"github_repository_code_scanning_default_setup.test", "state", "configured"),
						resource.TestCheckResourceAttrSet(
							"github_repository_code_scanning_default_setup.test", "query_suite"),
					),
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
					Check: resource.TestCheckResourceAttr(
						"github_repository_code_scanning_default_setup.test", "state", "configured"),
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
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr(
							"github_repository_code_scanning_default_setup.test", "state", "configured"),
						resource.TestCheckResourceAttr(
							"github_repository_code_scanning_default_setup.test", "query_suite", "default"),
					),
				},
				{
					Config: testAccCodeScanningDefaultSetupConfig(repoName, `
						state       = "configured"
						query_suite = "extended"
					`),
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr(
							"github_repository_code_scanning_default_setup.test", "state", "configured"),
						resource.TestCheckResourceAttr(
							"github_repository_code_scanning_default_setup.test", "query_suite", "extended"),
					),
				},
				{
					Config: testAccCodeScanningDefaultSetupConfig(repoName, `state = "not-configured"`),
					Check: resource.TestCheckResourceAttr(
						"github_repository_code_scanning_default_setup.test", "state", "not-configured"),
				},
			},
		})
	})
}
