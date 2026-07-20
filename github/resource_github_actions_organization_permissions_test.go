package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubActionsOrganizationPermissions(t *testing.T) {
	// IMPORTANT: Do not run these tests in parallel as they modify the organization state.

	t.Run("full_lifecycle", func(t *testing.T) {
		repo := mustCreateTestRepository(t)

		configMinimal := `
resource "github_actions_organization_permissions" "test" {
  allowed_actions      = "all"
  enabled_repositories = "all"
}
`

		configFull := fmt.Sprintf(`
resource "github_actions_organization_permissions" "test" {
  allowed_actions      = "selected"
  enabled_repositories = "selected"

  allowed_actions_config {
    github_owned_allowed = true
    patterns_allowed     = ["actions/cache@*", "actions/checkout@*"]
    verified_allowed     = true
  }

  enabled_repositories_config {
    repository_ids = [%d]
  }

  sha_pinning_required = true
}
`, repo.GetID())

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: configMinimal,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_actions_organization_permissions.test", tfjsonpath.New("allowed_actions_config"), knownvalue.ListSizeExact(0)),
						statecheck.ExpectKnownValue("github_actions_organization_permissions.test", tfjsonpath.New("enabled_repositories_config"), knownvalue.ListSizeExact(0)),
						statecheck.ExpectKnownValue("github_actions_organization_permissions.test", tfjsonpath.New("sha_pinning_required"), knownvalue.NotNull()),
					},
				},
				{
					Config: configFull,
				},
				{
					ResourceName:      "github_actions_organization_permissions.test",
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})

	t.Run("test setting of organization allowed actions", func(t *testing.T) {
		allowedActions := "selected"
		enabledRepositories := "all"
		githubOwnedAllowed := true
		verifiedAllowed := true

		config := fmt.Sprintf(`

			resource "github_actions_organization_permissions" "test" {
				allowed_actions = "%s"
				enabled_repositories = "%s"
				allowed_actions_config {
					github_owned_allowed = %t
					patterns_allowed     = ["actions/cache@*", "actions/checkout@*"]
					verified_allowed     = %t
				}
			}
		`, allowedActions, enabledRepositories, githubOwnedAllowed, verifiedAllowed)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_actions_organization_permissions.test", "allowed_actions", allowedActions,
			),
			resource.TestCheckResourceAttr(
				"github_actions_organization_permissions.test", "enabled_repositories", enabledRepositories,
			),
			resource.TestCheckResourceAttr(
				"github_actions_organization_permissions.test", "allowed_actions_config.#", "1",
			),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  check,
				},
			},
		})
	})

	t.Run("test not setting of organization allowed actions without error", func(t *testing.T) {
		allowedActions := "selected"
		enabledRepositories := "all"

		config := fmt.Sprintf(`

			resource "github_actions_organization_permissions" "test" {
				allowed_actions = "%s"
				enabled_repositories = "%s"
			}
		`, allowedActions, enabledRepositories)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_actions_organization_permissions.test", "allowed_actions", allowedActions,
			),
			resource.TestCheckResourceAttr(
				"github_actions_organization_permissions.test", "enabled_repositories", enabledRepositories,
			),
			resource.TestCheckResourceAttr(
				"github_actions_organization_permissions.test", "allowed_actions_config.#", "0",
			),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  check,
				},
			},
		})
	})

	t.Run("test setting sha_pinning_required to true then updating to false", func(t *testing.T) {
		enabledRepositories := "all"

		configTrue := fmt.Sprintf(`
			resource "github_actions_organization_permissions" "test" {
				allowed_actions = "all"
				enabled_repositories = "%s"
				sha_pinning_required = true
			}
		`, enabledRepositories)

		configFalse := fmt.Sprintf(`
			resource "github_actions_organization_permissions" "test" {
				allowed_actions = "all"
				enabled_repositories = "%s"
				sha_pinning_required = false
			}
		`, enabledRepositories)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: configTrue,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr(
							"github_actions_organization_permissions.test", "sha_pinning_required", "true",
						),
					),
				},
				{
					Config: configFalse,
					Check: resource.ComposeTestCheckFunc(
						resource.TestCheckResourceAttr(
							"github_actions_organization_permissions.test", "sha_pinning_required", "false",
						),
					),
				},
			},
		})
	})

	t.Run("test setting of organization enabled repositories", func(t *testing.T) {
		allowedActions := "all"
		enabledRepositories := "selected"
		githubOwnedAllowed := true
		verifiedAllowed := true
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		randomID2 := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-act-org-perm-%s", testResourcePrefix, randomID)
		repoName2 := fmt.Sprintf("%srepo-act-org-perm-%s", testResourcePrefix, randomID2)

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name        = "%[1]s"
				description = "Terraform acceptance tests %[1]s"
				topics			= ["terraform", "testing"]
			}

			resource "github_repository" "test2" {
				name        = "%[2]s"
				description = "Terraform acceptance tests %[2]s"
				topics			= ["terraform", "testing"]
			}

			resource "github_actions_organization_permissions" "test" {
				allowed_actions = "%s"
				enabled_repositories = "%s"
				enabled_repositories_config {
					repository_ids       = [github_repository.test.repo_id, github_repository.test2.repo_id]
				}
			}
		`, repoName, repoName2, allowedActions, enabledRepositories, githubOwnedAllowed, verifiedAllowed)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(
				"github_actions_organization_permissions.test", "allowed_actions", allowedActions,
			),
			resource.TestCheckResourceAttr(
				"github_actions_organization_permissions.test", "enabled_repositories", enabledRepositories,
			),
			resource.TestCheckResourceAttr(
				"github_actions_organization_permissions.test", "enabled_repositories_config.#", "1",
			),
		)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  check,
				},
			},
		})
	})
}
