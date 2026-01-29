package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccGithubEnterpriseActionsRunnerGroupOrgSettings(t *testing.T) {
	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

	if isEnterprise != "true" {
		t.Skip("Skipping because `ENTERPRISE_ACCOUNT` is not set or set to false")
	}

	if testEnterprise == "" {
		t.Skip("Skipping because `ENTERPRISE_SLUG` is not set")
	}

	if testOrganization == "" {
		t.Skip("Skipping because `GITHUB_ORGANIZATION` is not set")
	}

	t.Run("manages repository access for enterprise runner group", func(t *testing.T) {
		config := fmt.Sprintf(`
			data "github_enterprise" "enterprise" {
				slug = "%s"
			}

			data "github_organization" "org" {
				name = "%s"
			}

			# Create a test repository
			resource "github_repository" "test" {
				name        = "tf-acc-test-repo-%s"
				description = "Test repository for runner group access"
				visibility  = "private"
			}

			# Create an enterprise runner group
			resource "github_enterprise_actions_runner_group" "test" {
				enterprise_slug           = data.github_enterprise.enterprise.slug
				name                      = "tf-acc-test-rg-%s"
				visibility                = "selected"
				selected_organization_ids = [data.github_organization.org.id]
			}

			# Configure repository access for the enterprise runner group
			resource "github_enterprise_actions_runner_group_org_settings" "test" {
				organization                 = data.github_organization.org.name
				enterprise_runner_group_name = github_enterprise_actions_runner_group.test.name
				selected_repository_ids      = [github_repository.test.repo_id]
				allows_public_repositories   = true
			}
		`, testEnterprise, testOrganization, randomID, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet(
				"github_enterprise_actions_runner_group_org_settings.test", "runner_group_id",
			),
			resource.TestCheckResourceAttr(
				"github_enterprise_actions_runner_group_org_settings.test", "organization",
				testOrganization,
			),
			resource.TestCheckResourceAttr(
				"github_enterprise_actions_runner_group_org_settings.test", "allows_public_repositories",
				"true",
			),
			resource.TestCheckResourceAttr(
				"github_enterprise_actions_runner_group_org_settings.test", "selected_repository_ids.#",
				"1",
			),
		)

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: config,
						Check:  check,
					},
				},
			})
		}

		t.Run("with an enterprise account", func(t *testing.T) {
			testCase(t, enterprise)
		})
	})

	t.Run("updates repository access", func(t *testing.T) {
		configCreate := fmt.Sprintf(`
			data "github_enterprise" "enterprise" {
				slug = "%s"
			}

			data "github_organization" "org" {
				name = "%s"
			}

			resource "github_repository" "test1" {
				name        = "tf-acc-test-repo1-%s"
				description = "Test repository 1 for runner group access"
				visibility  = "private"
			}

			resource "github_repository" "test2" {
				name        = "tf-acc-test-repo2-%s"
				description = "Test repository 2 for runner group access"
				visibility  = "private"
			}

			resource "github_enterprise_actions_runner_group" "test" {
				enterprise_slug           = data.github_enterprise.enterprise.slug
				name                      = "tf-acc-test-rg-%s"
				visibility                = "selected"
				selected_organization_ids = [data.github_organization.org.id]
			}

			resource "github_enterprise_actions_runner_group_org_settings" "test" {
				organization                 = data.github_organization.org.name
				enterprise_runner_group_name = github_enterprise_actions_runner_group.test.name
				selected_repository_ids      = [github_repository.test1.repo_id]
			}
		`, testEnterprise, testOrganization, randomID, randomID, randomID)

		configUpdate := fmt.Sprintf(`
			data "github_enterprise" "enterprise" {
				slug = "%s"
			}

			data "github_organization" "org" {
				name = "%s"
			}

			resource "github_repository" "test1" {
				name        = "tf-acc-test-repo1-%s"
				description = "Test repository 1 for runner group access"
				visibility  = "private"
			}

			resource "github_repository" "test2" {
				name        = "tf-acc-test-repo2-%s"
				description = "Test repository 2 for runner group access"
				visibility  = "private"
			}

			resource "github_enterprise_actions_runner_group" "test" {
				enterprise_slug           = data.github_enterprise.enterprise.slug
				name                      = "tf-acc-test-rg-%s"
				visibility                = "selected"
				selected_organization_ids = [data.github_organization.org.id]
			}

			resource "github_enterprise_actions_runner_group_org_settings" "test" {
				organization                 = data.github_organization.org.name
				enterprise_runner_group_name = github_enterprise_actions_runner_group.test.name
				selected_repository_ids      = [github_repository.test1.repo_id, github_repository.test2.repo_id]
			}
		`, testEnterprise, testOrganization, randomID, randomID, randomID)

		checkCreate := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_enterprise_actions_runner_group_org_settings.test", "selected_repository_ids.#",
				"1",
			),
		)

		checkUpdate := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_enterprise_actions_runner_group_org_settings.test", "selected_repository_ids.#",
				"2",
			),
		)

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
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
		}

		t.Run("with an enterprise account", func(t *testing.T) {
			testCase(t, enterprise)
		})
	})

	t.Run("manages workflow restrictions", func(t *testing.T) {
		config := fmt.Sprintf(`
			data "github_enterprise" "enterprise" {
				slug = "%s"
			}

			data "github_organization" "org" {
				name = "%s"
			}

			resource "github_repository" "test" {
				name        = "tf-acc-test-repo-%s"
				description = "Test repository for runner group access"
				visibility  = "private"
			}

			resource "github_enterprise_actions_runner_group" "test" {
				enterprise_slug           = data.github_enterprise.enterprise.slug
				name                      = "tf-acc-test-rg-%s"
				visibility                = "selected"
				selected_organization_ids = [data.github_organization.org.id]
			}

			resource "github_enterprise_actions_runner_group_org_settings" "test" {
				organization                 = data.github_organization.org.name
				enterprise_runner_group_name = github_enterprise_actions_runner_group.test.name
				selected_repository_ids      = [github_repository.test.repo_id]
				restricted_to_workflows      = true
				selected_workflows           = ["${github_repository.test.full_name}/.github/workflows/test.yml@refs/heads/main"]
			}
		`, testEnterprise, testOrganization, randomID, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet(
				"github_enterprise_actions_runner_group_org_settings.test", "runner_group_id",
			),
			resource.TestCheckResourceAttr(
				"github_enterprise_actions_runner_group_org_settings.test", "restricted_to_workflows",
				"true",
			),
			resource.TestCheckResourceAttr(
				"github_enterprise_actions_runner_group_org_settings.test", "selected_workflows.#",
				"1",
			),
			func(state *terraform.State) error {
				githubRepository := state.RootModule().Resources["github_repository.test"].Primary
				fullName := githubRepository.Attributes["full_name"]

				runnerGroup := state.RootModule().Resources["github_enterprise_actions_runner_group_org_settings.test"].Primary
				workflowActual := runnerGroup.Attributes["selected_workflows.0"]

				workflowExpected := fmt.Sprintf("%s/.github/workflows/test.yml@refs/heads/main", fullName)

				if workflowActual != workflowExpected {
					return fmt.Errorf("expected selected_workflows.0 to be %s, got %s", workflowExpected, workflowActual)
				}
				return nil
			},
		)

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: config,
						Check:  check,
					},
				},
			})
		}

		t.Run("with an enterprise account", func(t *testing.T) {
			testCase(t, enterprise)
		})
	})

	t.Run("imports runner group access by ID", func(t *testing.T) {
		config := fmt.Sprintf(`
			data "github_enterprise" "enterprise" {
				slug = "%s"
			}

			data "github_organization" "org" {
				name = "%s"
			}

			resource "github_repository" "test" {
				name        = "tf-acc-test-repo-%s"
				description = "Test repository for runner group access"
				visibility  = "private"
			}

			resource "github_enterprise_actions_runner_group" "test" {
				enterprise_slug           = data.github_enterprise.enterprise.slug
				name                      = "tf-acc-test-rg-%s"
				visibility                = "selected"
				selected_organization_ids = [data.github_organization.org.id]
			}

			resource "github_enterprise_actions_runner_group_org_settings" "test" {
				organization                 = data.github_organization.org.name
				enterprise_runner_group_name = github_enterprise_actions_runner_group.test.name
				selected_repository_ids      = [github_repository.test.repo_id]
			}
		`, testEnterprise, testOrganization, randomID, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet(
				"github_enterprise_actions_runner_group_org_settings.test", "runner_group_id",
			),
			resource.TestCheckResourceAttr(
				"github_enterprise_actions_runner_group_org_settings.test", "organization",
				testOrganization,
			),
		)

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, mode) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: config,
						Check:  check,
					},
					{
						ResourceName:      "github_enterprise_actions_runner_group_org_settings.test",
						ImportState:       true,
						ImportStateVerify: true,
						ImportStateIdFunc: importEnterpriseActionsRunnerGroupOrgSettingsByID("github_enterprise_actions_runner_group_org_settings.test"),
					},
				},
			})
		}

		t.Run("with an enterprise account", func(t *testing.T) {
			testCase(t, enterprise)
		})
	})
}

func importEnterpriseActionsRunnerGroupOrgSettingsByID(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("resource not found: %s", resourceName)
		}
		return fmt.Sprintf("%s:%s", testOrganization, rs.Primary.Attributes["runner_group_id"]), nil
	}
}
