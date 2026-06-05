package github

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
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

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)
		propertyName := fmt.Sprintf("%sproperty-%s", testResourcePrefix, randomID)
		option1 := "option1"
		option2 := "option2"

		config := fmt.Sprintf(`
resource "github_organization_custom_properties" "test" {
  allowed_values = ["%s", "%s"]
  description    = "Test Description"
  property_name  = "%s"
  value_type     = "single_select"
}

resource "github_repository" "test" {
  name      = "%s"
  auto_init = true
}

resource "github_repository_custom_property" "test" {
  repository     = github_repository.test.name
  property_name  = github_organization_custom_properties.test.property_name
  property_type  = github_organization_custom_properties.test.value_type
  property_value = %%s
}
`, option1, option2, propertyName, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, fmt.Sprintf(`["%s"]`, option1)),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_custom_property.test", tfjsonpath.New("repository_id"), knownvalue.NotNull()),
					},
				},
				{
					Config:      fmt.Sprintf(config, `["invalid_option"]`),
					ExpectError: regexp.MustCompile(`is not allowed for property`),
				},
				{
					Config: fmt.Sprintf(config, fmt.Sprintf(`["%s"]`, option2)),
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

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)
		propertyName := fmt.Sprintf("%sproperty-%s", testResourcePrefix, randomID)
		option1 := "option1"
		option2 := "option2"
		option3 := "option3"

		config := fmt.Sprintf(`
resource "github_organization_custom_properties" "test" {
  allowed_values = ["%s", "%s", "%s"]
  description    = "Test Description"
  property_name  = "%s"
  value_type     = "multi_select"
}

resource "github_repository" "test" {
  name      = "%s"
  auto_init = true
}

resource "github_repository_custom_property" "test" {
  repository     = github_repository.test.name
  property_name  = github_organization_custom_properties.test.property_name
  property_type  = github_organization_custom_properties.test.value_type
  property_value = %%s
}
`, option1, option2, option3, propertyName, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, fmt.Sprintf(`["%s", "%s"]`, option1, option2)),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_repository_custom_property.test", tfjsonpath.New("repository_id"), knownvalue.NotNull()),
					},
				},
				{
					Config:      fmt.Sprintf(config, `["invalid_option"]`),
					ExpectError: regexp.MustCompile(`is not allowed for property`),
				},
				{
					Config: fmt.Sprintf(config, fmt.Sprintf(`["%s"]`, option3)),
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

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)
		propertyName := fmt.Sprintf("%sproperty-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
resource "github_organization_custom_properties" "test" {
  description    = "Test Description"
  property_name  = "%s"
  value_type     = "true_false"
}

resource "github_repository" "test" {
  name      = "%s"
  auto_init = true
}

resource "github_repository_custom_property" "test" {
  repository    = github_repository.test.name
  property_name = github_organization_custom_properties.test.property_name
  property_type = github_organization_custom_properties.test.value_type
  property_value = %%s
}
`, propertyName, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
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

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)
		propertyName := fmt.Sprintf("%sproperty-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
resource "github_organization_custom_properties" "test" {
  description    = "Test Description"
  property_name  = "%s"
  value_type     = "url"
}

resource "github_repository" "test" {
  name      = "%s"
  auto_init = true
}

resource "github_repository_custom_property" "test" {
  repository     = github_repository.test.name
  property_name  = github_organization_custom_properties.test.property_name
  property_type  = github_organization_custom_properties.test.value_type
  property_value = %%s
}
`, propertyName, repoName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
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

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		repoName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)
		propertyName := fmt.Sprintf("%sproperty-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
resource "github_organization_custom_properties" "test" {
  description    = "Test Description"
  property_name  = "%s"
  value_type     = "string"
}

resource "github_repository" "test" {
  name      = "%s"
  auto_init = true
}

resource "github_repository_custom_property" "test" {
  repository     = github_repository.test.name
  property_name  = github_organization_custom_properties.test.property_name
  property_type  = github_organization_custom_properties.test.value_type
  property_value = %%s
}
`, propertyName, repoName)

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
