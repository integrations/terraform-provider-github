package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccGithubOrganizationInheritedRunnerGroupSettings(t *testing.T) {
	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

	t.Run("manages repository access for enterprise runner group", func(t *testing.T) {
		repoName := fmt.Sprintf("%srepo-%s", testResourcePrefix, randomID)
		rgName := fmt.Sprintf("%srg-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
			data "github_enterprise" "enterprise" {
				slug = "%s"
			}

			data "github_organization" "org" {
				name = "%s"
			}

			resource "github_repository" "test" {
				name        = "%s"
				description = "Test repository for runner group access"
				visibility  = "private"
			}

			resource "github_enterprise_actions_runner_group" "test" {
				enterprise_slug           = data.github_enterprise.enterprise.slug
				name                      = "%s"
				visibility                = "selected"
				selected_organization_ids = [data.github_organization.org.id]
			}

			resource "github_organization_inherited_runner_group_settings" "test" {
				organization                 = data.github_organization.org.name
				enterprise_runner_group_name = github_enterprise_actions_runner_group.test.name
				selected_repository_ids      = [github_repository.test.repo_id]
				allows_public_repositories   = true
			}
		`, testAccConf.enterpriseSlug, testAccConf.owner, repoName, rgName)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet(
				"github_organization_inherited_runner_group_settings.test", "runner_group_id",
			),
			resource.TestCheckResourceAttr(
				"github_organization_inherited_runner_group_settings.test", "organization",
				testAccConf.owner,
			),
			resource.TestCheckResourceAttr(
				"github_organization_inherited_runner_group_settings.test", "allows_public_repositories",
				"true",
			),
			resource.TestCheckResourceAttr(
				"github_organization_inherited_runner_group_settings.test", "selected_repository_ids.#",
				"1",
			),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  check,
				},
			},
		})
	})

	t.Run("updates repository access", func(t *testing.T) {
		repoName1 := fmt.Sprintf("%srepo1-%s", testResourcePrefix, randomID)
		repoName2 := fmt.Sprintf("%srepo2-%s", testResourcePrefix, randomID)
		rgName := fmt.Sprintf("%srg-update-%s", testResourcePrefix, randomID)

		configCreate := fmt.Sprintf(`
			data "github_enterprise" "enterprise" {
				slug = "%s"
			}

			data "github_organization" "org" {
				name = "%s"
			}

			resource "github_repository" "test1" {
				name        = "%s"
				description = "Test repository 1 for runner group access"
				visibility  = "private"
			}

			resource "github_repository" "test2" {
				name        = "%s"
				description = "Test repository 2 for runner group access"
				visibility  = "private"
			}

			resource "github_enterprise_actions_runner_group" "test" {
				enterprise_slug           = data.github_enterprise.enterprise.slug
				name                      = "%s"
				visibility                = "selected"
				selected_organization_ids = [data.github_organization.org.id]
			}

			resource "github_organization_inherited_runner_group_settings" "test" {
				organization                 = data.github_organization.org.name
				enterprise_runner_group_name = github_enterprise_actions_runner_group.test.name
				selected_repository_ids      = [github_repository.test1.repo_id]
			}
		`, testAccConf.enterpriseSlug, testAccConf.owner, repoName1, repoName2, rgName)

		configUpdate := fmt.Sprintf(`
			data "github_enterprise" "enterprise" {
				slug = "%s"
			}

			data "github_organization" "org" {
				name = "%s"
			}

			resource "github_repository" "test1" {
				name        = "%s"
				description = "Test repository 1 for runner group access"
				visibility  = "private"
			}

			resource "github_repository" "test2" {
				name        = "%s"
				description = "Test repository 2 for runner group access"
				visibility  = "private"
			}

			resource "github_enterprise_actions_runner_group" "test" {
				enterprise_slug           = data.github_enterprise.enterprise.slug
				name                      = "%s"
				visibility                = "selected"
				selected_organization_ids = [data.github_organization.org.id]
			}

			resource "github_organization_inherited_runner_group_settings" "test" {
				organization                 = data.github_organization.org.name
				enterprise_runner_group_name = github_enterprise_actions_runner_group.test.name
				selected_repository_ids      = [github_repository.test1.repo_id, github_repository.test2.repo_id]
			}
		`, testAccConf.enterpriseSlug, testAccConf.owner, repoName1, repoName2, rgName)

		checkCreate := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_organization_inherited_runner_group_settings.test", "selected_repository_ids.#",
				"1",
			),
		)

		checkUpdate := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_organization_inherited_runner_group_settings.test", "selected_repository_ids.#",
				"2",
			),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: configCreate,
					Check:  checkCreate,
				},
				{
					Config: configUpdate,
					Check:  checkUpdate,
				},
			},
		})
	})

	t.Run("manages workflow restrictions", func(t *testing.T) {
		repoName := fmt.Sprintf("%srepo-wf-%s", testResourcePrefix, randomID)
		rgName := fmt.Sprintf("%srg-wf-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
			data "github_enterprise" "enterprise" {
				slug = "%s"
			}

			data "github_organization" "org" {
				name = "%s"
			}

			resource "github_repository" "test" {
				name        = "%s"
				description = "Test repository for runner group access"
				visibility  = "private"
			}

			resource "github_enterprise_actions_runner_group" "test" {
				enterprise_slug           = data.github_enterprise.enterprise.slug
				name                      = "%s"
				visibility                = "selected"
				selected_organization_ids = [data.github_organization.org.id]
			}

			resource "github_organization_inherited_runner_group_settings" "test" {
				organization                 = data.github_organization.org.name
				enterprise_runner_group_name = github_enterprise_actions_runner_group.test.name
				selected_repository_ids      = [github_repository.test.repo_id]
				restricted_to_workflows      = true
				selected_workflows           = ["${github_repository.test.full_name}/.github/workflows/test.yml@refs/heads/main"]
			}
		`, testAccConf.enterpriseSlug, testAccConf.owner, repoName, rgName)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet(
				"github_organization_inherited_runner_group_settings.test", "runner_group_id",
			),
			resource.TestCheckResourceAttr(
				"github_organization_inherited_runner_group_settings.test", "restricted_to_workflows",
				"true",
			),
			resource.TestCheckResourceAttr(
				"github_organization_inherited_runner_group_settings.test", "selected_workflows.#",
				"1",
			),
			func(state *terraform.State) error {
				githubRepository := state.RootModule().Resources["github_repository.test"].Primary
				fullName := githubRepository.Attributes["full_name"]

				runnerGroup := state.RootModule().Resources["github_organization_inherited_runner_group_settings.test"].Primary
				workflowActual := runnerGroup.Attributes["selected_workflows.0"]

				workflowExpected := fmt.Sprintf("%s/.github/workflows/test.yml@refs/heads/main", fullName)

				if workflowActual != workflowExpected {
					return fmt.Errorf("expected selected_workflows.0 to be %s, got %s", workflowExpected, workflowActual)
				}
				return nil
			},
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  check,
				},
			},
		})
	})

	t.Run("imports runner group access by ID", func(t *testing.T) {
		repoName := fmt.Sprintf("%srepo-import-%s", testResourcePrefix, randomID)
		rgName := fmt.Sprintf("%srg-import-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
			data "github_enterprise" "enterprise" {
				slug = "%s"
			}

			data "github_organization" "org" {
				name = "%s"
			}

			resource "github_repository" "test" {
				name        = "%s"
				description = "Test repository for runner group access"
				visibility  = "private"
			}

			resource "github_enterprise_actions_runner_group" "test" {
				enterprise_slug           = data.github_enterprise.enterprise.slug
				name                      = "%s"
				visibility                = "selected"
				selected_organization_ids = [data.github_organization.org.id]
			}

			resource "github_organization_inherited_runner_group_settings" "test" {
				organization                 = data.github_organization.org.name
				enterprise_runner_group_name = github_enterprise_actions_runner_group.test.name
				selected_repository_ids      = [github_repository.test.repo_id]
			}
		`, testAccConf.enterpriseSlug, testAccConf.owner, repoName, rgName)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet(
				"github_organization_inherited_runner_group_settings.test", "runner_group_id",
			),
			resource.TestCheckResourceAttr(
				"github_organization_inherited_runner_group_settings.test", "organization",
				testAccConf.owner,
			),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  check,
				},
				{
					ResourceName:      "github_organization_inherited_runner_group_settings.test",
					ImportState:       true,
					ImportStateVerify: true,
					ImportStateIdFunc: importOrganizationInheritedRunnerGroupSettingsByID("github_organization_inherited_runner_group_settings.test"),
				},
			},
		})
	})
}

func importOrganizationInheritedRunnerGroupSettingsByID(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("resource not found: %s", resourceName)
		}
		return fmt.Sprintf("%s:%s", testAccConf.owner, rs.Primary.Attributes["runner_group_id"]), nil
	}
}
