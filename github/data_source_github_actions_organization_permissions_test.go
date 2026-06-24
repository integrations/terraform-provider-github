package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubActionsOrganizationPermissionsDataSource(t *testing.T) {
	t.Run("reads organization permissions without error", func(t *testing.T) {
		config := `
			resource "github_actions_organization_permissions" "test" {
				allowed_actions      = "local_only"
				enabled_repositories = "all"
			}

			data "github_actions_organization_permissions" "test" {
				depends_on = [github_actions_organization_permissions.test]
			}
		`

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(
							"data.github_actions_organization_permissions.test",
							tfjsonpath.New("allowed_actions"),
							knownvalue.StringExact("local_only"),
						),
						statecheck.ExpectKnownValue(
							"data.github_actions_organization_permissions.test",
							tfjsonpath.New("enabled_repositories"),
							knownvalue.StringExact("all"),
						),
					},
				},
			},
		})
	})

	t.Run("reads selected repository and action permissions", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%srepo-act-org-perm-ds-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
			resource "github_repository" "test" {
				name        = "%[1]s"
				description = "Terraform acceptance tests %[1]s"
			}

			resource "github_actions_organization_permissions" "test" {
				allowed_actions      = "selected"
				enabled_repositories = "selected"
				allowed_actions_config {
					github_owned_allowed = true
					verified_allowed     = true
					patterns_allowed     = ["actions/cache@*", "actions/checkout@*"]
				}
				enabled_repositories_config {
					repository_ids = [github_repository.test.repo_id]
				}
			}

			data "github_actions_organization_permissions" "test" {
				depends_on = [github_actions_organization_permissions.test]
			}
		`, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(
							"data.github_actions_organization_permissions.test",
							tfjsonpath.New("allowed_actions"),
							knownvalue.StringExact("selected"),
						),
						statecheck.ExpectKnownValue(
							"data.github_actions_organization_permissions.test",
							tfjsonpath.New("enabled_repositories"),
							knownvalue.StringExact("selected"),
						),
						statecheck.ExpectKnownValue(
							"data.github_actions_organization_permissions.test",
							tfjsonpath.New("allowed_actions_config").AtSliceIndex(0).AtMapKey("github_owned_allowed"),
							knownvalue.Bool(true),
						),
						statecheck.ExpectKnownValue(
							"data.github_actions_organization_permissions.test",
							tfjsonpath.New("allowed_actions_config").AtSliceIndex(0).AtMapKey("verified_allowed"),
							knownvalue.Bool(true),
						),
						statecheck.ExpectKnownValue(
							"data.github_actions_organization_permissions.test",
							tfjsonpath.New("enabled_repositories_config").AtSliceIndex(0).AtMapKey("repository_ids"),
							knownvalue.SetSizeExact(1),
						),
					},
				},
			},
		})
	})
}
