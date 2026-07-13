package github

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubActionsOrganizationVariable(t *testing.T) {
	t.Parallel()

	t.Run("with_visibility_all", func(t *testing.T) {
		t.Parallel()

		randomID := acctest.RandString(testRandomIDLength)
		varName := strings.ToUpper(fmt.Sprintf("%s%s", strings.ReplaceAll(testResourcePrefix, "-", "_"), randomID))

		config := fmt.Sprintf(`
resource "github_actions_organization_variable" "test" {
	variable_name = "%s"
	value         = "%%s"
	visibility    = "all"
}
`, varName)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, "my-value"),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_actions_organization_variable.test", tfjsonpath.New("created_at"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_actions_organization_variable.test", tfjsonpath.New("updated_at"), knownvalue.NotNull()),
					},
				},
				{
					Config: fmt.Sprintf(config, "my-value-2"),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_actions_organization_variable.test", plancheck.ResourceActionUpdate),
						},
					},
				},
				{
					ResourceName:      "github_actions_organization_variable.test",
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})

	t.Run("with_visibility_private", func(t *testing.T) {
		t.Parallel()

		randomID := acctest.RandString(testRandomIDLength)
		varName := strings.ToUpper(fmt.Sprintf("%s%s", strings.ReplaceAll(testResourcePrefix, "-", "_"), randomID))

		config := fmt.Sprintf(`
resource "github_actions_organization_variable" "test" {
	variable_name = "%s"
	value         = "%%s"
	visibility    = "private"
}
`, varName)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, "my-value"),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_actions_organization_variable.test", tfjsonpath.New("created_at"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_actions_organization_variable.test", tfjsonpath.New("updated_at"), knownvalue.NotNull()),
					},
				},
				{
					Config: fmt.Sprintf(config, "my-value-2"),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_actions_organization_variable.test", plancheck.ResourceActionUpdate),
						},
					},
				},
				{
					ResourceName:      "github_actions_organization_variable.test",
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})

	t.Run("with_visibility_selected_no_repos", func(t *testing.T) {
		t.Parallel()

		randomID := acctest.RandString(testRandomIDLength)
		varName := strings.ToUpper(fmt.Sprintf("%s%s", strings.ReplaceAll(testResourcePrefix, "-", "_"), randomID))

		config := fmt.Sprintf(`
resource "github_actions_organization_variable" "test" {
	variable_name = "%s"
	value         = "%%s"
	visibility    = "selected"
}
`, varName)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, "my-value"),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_actions_organization_variable.test", tfjsonpath.New("created_at"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_actions_organization_variable.test", tfjsonpath.New("updated_at"), knownvalue.NotNull()),
					},
				},
				{
					Config: fmt.Sprintf(config, "my-value-2"),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_actions_organization_variable.test", plancheck.ResourceActionUpdate),
						},
					},
				},
				{
					ResourceName:      "github_actions_organization_variable.test",
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})

	t.Run("with_visibility_selected_with_repos", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)
		repo2 := mustCreateTestRepository(t)

		randomID := acctest.RandString(testRandomIDLength)
		varName := strings.ToUpper(fmt.Sprintf("%s%s", strings.ReplaceAll(testResourcePrefix, "-", "_"), randomID))

		config := fmt.Sprintf(`
resource "github_actions_organization_variable" "test" {
	variable_name = "%s"
	value         = "my-value"
	visibility    = "selected"
  selected_repository_ids = [%%s]
}
`, varName)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, fmt.Sprintf(`"%v"`, repo.GetID())),
				},
				{
					Config: fmt.Sprintf(config, fmt.Sprintf(`"%v", "%v"`, repo.GetID(), repo2.GetID())),
				},
				{
					Config: fmt.Sprintf(config, ""),
				},
			},
		})
	})

	t.Run("errors_with_visibility_not_selected_and_selected_repository_ids", func(t *testing.T) {
		t.Parallel()

		randomID := acctest.RandString(testRandomIDLength)
		varName := strings.ToUpper(fmt.Sprintf("%s%s", strings.ReplaceAll(testResourcePrefix, "-", "_"), randomID))

		config := fmt.Sprintf(`
resource "github_actions_organization_variable" "test" {
  variable_name = "%s"
  value         = "foo"
  visibility    = "all"
  selected_repository_ids = [123456]
}
`, varName)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      config,
					ExpectError: regexp.MustCompile("cannot use selected_repository_ids without visibility being set to selected"),
				},
			},
		})
	})

	t.Run("errors_if_variable_already_exists", func(t *testing.T) {
		t.Parallel()

		randomID := acctest.RandString(testRandomIDLength)
		varName := strings.ToUpper(fmt.Sprintf("%s%s", strings.ReplaceAll(testResourcePrefix, "-", "_"), randomID))

		mustCreateTestOrganizationVariable(t, &varName, nil)

		config := fmt.Sprintf(`
resource "github_actions_organization_variable" "test" {
  variable_name = "%s"
  value         = "my-value"
  visibility    = "all"
}
`, varName)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      config,
					ExpectError: regexp.MustCompile(`Variable already exists`),
				},
			},
		})
	})
}
