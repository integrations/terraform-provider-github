package github

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubRepositoryCustomProperty(t *testing.T) {
	t.Parallel()

	t.Run("single_select", func(t *testing.T) {
		t.Parallel()

		prop := mustCreateTestOrganizationRepositoryCustomProperty(t, "single_select", []string{"option1", "option2"})
		repo := mustCreateTestRepository(t)
		allowed := prop.GetAllowedValues()

		config := fmt.Sprintf(`
resource "github_repository_custom_property" "test" {
  repository     = "%s"
  property_name  = "%s"
  property_type  = "%s"
  property_value = %%s
}
`, repo.GetName(), prop.GetPropertyName(), prop.GetValueType())

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      fmt.Sprintf(config, `[]`),
					ExpectError: regexp.MustCompile(`Not enough list items`),
					PlanOnly:    true,
				},
				{
					Config: fmt.Sprintf(config, fmt.Sprintf(`["%s"]`, allowed[0])),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_custom_property.test", tfjsonpath.New("repository_id"), knownvalue.NotNull()),
					},
				},
				{
					Config:      fmt.Sprintf(config, `["invalid_option"]`),
					ExpectError: regexp.MustCompile(`is not allowed for property`),
				},
				{
					Config: fmt.Sprintf(config, fmt.Sprintf(`["%s"]`, allowed[1])),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_repository_custom_property.test", plancheck.ResourceActionUpdate),
						},
					},
				},
				{
					ResourceName:      "github_repository_custom_property.test",
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})

	t.Run("multi_select", func(t *testing.T) {
		t.Parallel()

		prop := mustCreateTestOrganizationRepositoryCustomProperty(t, "multi_select", []string{"option1", "option2", "option3"})
		repo := mustCreateTestRepository(t)
		allowed := prop.GetAllowedValues()

		config := fmt.Sprintf(`
resource "github_repository_custom_property" "test" {
  repository     = "%s"
  property_name  = "%s"
  property_type  = "%s"
  property_value = %%s
}
`, repo.GetName(), prop.GetPropertyName(), prop.GetValueType())

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      fmt.Sprintf(config, `[]`),
					ExpectError: regexp.MustCompile(`Not enough list items`),
					PlanOnly:    true,
				},
				{
					Config: fmt.Sprintf(config, fmt.Sprintf(`["%s", "%s"]`, allowed[0], allowed[1])),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_custom_property.test", tfjsonpath.New("repository_id"), knownvalue.NotNull()),
					},
				},
				{
					Config:      fmt.Sprintf(config, `["invalid_option"]`),
					ExpectError: regexp.MustCompile(`is not allowed for property`),
				},
				{
					Config: fmt.Sprintf(config, fmt.Sprintf(`["%s"]`, allowed[2])),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_repository_custom_property.test", plancheck.ResourceActionUpdate),
						},
					},
				},
				{
					ResourceName:      "github_repository_custom_property.test",
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})

	t.Run("true_false", func(t *testing.T) {
		t.Parallel()

		prop := mustCreateTestOrganizationRepositoryCustomProperty(t, "true_false", nil)
		repo := mustCreateTestRepository(t)

		config := fmt.Sprintf(`
resource "github_repository_custom_property" "test" {
  repository    = "%s"
  property_name = "%s"
  property_type = "%s"
  property_value = %%s
}
`, repo.GetName(), prop.GetPropertyName(), prop.GetValueType())

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      fmt.Sprintf(config, `[]`),
					ExpectError: regexp.MustCompile(`Not enough list items`),
					PlanOnly:    true,
				},
				{
					Config: fmt.Sprintf(config, `["true"]`),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_custom_property.test", tfjsonpath.New("repository_id"), knownvalue.NotNull()),
					},
				},
				{
					Config:      fmt.Sprintf(config, `["invalid_option"]`),
					ExpectError: regexp.MustCompile(`is not allowed for property`),
				},
				{
					Config: fmt.Sprintf(config, `["false"]`),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_repository_custom_property.test", plancheck.ResourceActionUpdate),
						},
					},
				},
				{
					ResourceName:      "github_repository_custom_property.test",
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})

	t.Run("url", func(t *testing.T) {
		t.Parallel()

		prop := mustCreateTestOrganizationRepositoryCustomProperty(t, "url", nil)
		repo := mustCreateTestRepository(t)

		config := fmt.Sprintf(`
resource "github_repository_custom_property" "test" {
  repository     = "%s"
  property_name  = "%s"
  property_type  = "%s"
  property_value = %%s
}
`, repo.GetName(), prop.GetPropertyName(), prop.GetValueType())

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      fmt.Sprintf(config, `[]`),
					ExpectError: regexp.MustCompile(`Not enough list items`),
					PlanOnly:    true,
				},
				{
					Config: fmt.Sprintf(config, `["https://example.com"]`),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_custom_property.test", tfjsonpath.New("repository_id"), knownvalue.NotNull()),
					},
				},
				{
					Config:      fmt.Sprintf(config, `["xxxx"]`),
					ExpectError: regexp.MustCompile(`URL must be absolute`),
				},
				{
					Config: fmt.Sprintf(config, `["https://example.com/test"]`),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_repository_custom_property.test", plancheck.ResourceActionUpdate),
						},
					},
				},
				{
					ResourceName:      "github_repository_custom_property.test",
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})

	t.Run("string", func(t *testing.T) {
		t.Parallel()

		prop := mustCreateTestOrganizationRepositoryCustomProperty(t, "string", nil)
		repo := mustCreateTestRepository(t)

		config := fmt.Sprintf(`
resource "github_repository_custom_property" "test" {
  repository     = "%s"
  property_name  = "%s"
  property_type  = "%s"
  property_value = %%s
}
`, repo.GetName(), prop.GetPropertyName(), prop.GetValueType())

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      fmt.Sprintf(config, `[]`),
					ExpectError: regexp.MustCompile(`Not enough list items`),
					PlanOnly:    true,
				},
				{
					Config:      fmt.Sprintf(config, `[""]`),
					ExpectError: regexp.MustCompile(`to not be an empty string`),
					PlanOnly:    true,
				},
				{
					Config: fmt.Sprintf(config, `["text"]`),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_custom_property.test", tfjsonpath.New("repository_id"), knownvalue.NotNull()),
					},
				},
				{
					Config: fmt.Sprintf(config, `["new text"]`),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_repository_custom_property.test", plancheck.ResourceActionUpdate),
						},
					},
				},
				{
					ResourceName:      "github_repository_custom_property.test",
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})
}
