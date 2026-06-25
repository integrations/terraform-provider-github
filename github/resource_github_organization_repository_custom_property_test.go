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

func TestAccGithubOrganizationRepositoryCustomProperty(t *testing.T) {
	const resourceAddr = "github_organization_repository_custom_property.test"

	t.Run("creates a string property without error", func(t *testing.T) {
		config := `
		resource "github_organization_repository_custom_property" "test" {
			property_name = "tf-acc-test-string"
			value_type    = "string"
			description   = "tf-acc-test string property"
		}`

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceAddr, tfjsonpath.New("property_name"), knownvalue.StringExact("tf-acc-test-string")),
						statecheck.ExpectKnownValue(resourceAddr, tfjsonpath.New("value_type"), knownvalue.StringExact("string")),
						statecheck.ExpectKnownValue(resourceAddr, tfjsonpath.New("values_editable_by"), knownvalue.StringExact("org_actors")),
					},
				},
			},
		})
	})

	t.Run("creates a single_select property and grows allowed_values", func(t *testing.T) {
		configBefore := `
		resource "github_organization_repository_custom_property" "test" {
			property_name  = "tf-acc-test-single-select"
			value_type     = "single_select"
			description    = "tf-acc-test single_select property"
			allowed_values = ["one"]
		}`
		configAfter := `
		resource "github_organization_repository_custom_property" "test" {
			property_name  = "tf-acc-test-single-select"
			value_type     = "single_select"
			description    = "tf-acc-test single_select property updated"
			allowed_values = ["one", "two"]
		}`

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: configBefore,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceAddr, tfjsonpath.New("allowed_values"), knownvalue.ListExact([]knownvalue.Check{
							knownvalue.StringExact("one"),
						})),
					},
				},
				{
					Config: configAfter,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceAddr, tfjsonpath.New("allowed_values"), knownvalue.ListExact([]knownvalue.Check{
							knownvalue.StringExact("one"),
							knownvalue.StringExact("two"),
						})),
						statecheck.ExpectKnownValue(resourceAddr, tfjsonpath.New("description"), knownvalue.StringExact("tf-acc-test single_select property updated")),
					},
				},
			},
		})
	})

	t.Run("imports without error", func(t *testing.T) {
		config := `
		resource "github_organization_repository_custom_property" "test" {
			property_name = "tf-acc-test-import"
			value_type    = "string"
			description   = "tf-acc-test import"
		}`

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{Config: config},
				{
					ResourceName:      resourceAddr,
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})

	t.Run("forces new when property_name changes", func(t *testing.T) {
		before := `
		resource "github_organization_repository_custom_property" "test" {
			property_name = "tf-acc-test-rename-a"
			value_type    = "string"
		}`
		after := `
		resource "github_organization_repository_custom_property" "test" {
			property_name = "tf-acc-test-rename-b"
			value_type    = "string"
		}`

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{Config: before},
				{
					Config: after,
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(resourceAddr, plancheck.ResourceActionDestroyBeforeCreate),
						},
					},
				},
			},
		})
	})

	t.Run("forces new when value_type changes", func(t *testing.T) {
		before := `
		resource "github_organization_repository_custom_property" "test" {
			property_name = "tf-acc-test-retype"
			value_type    = "string"
		}`
		after := `
		resource "github_organization_repository_custom_property" "test" {
			property_name  = "tf-acc-test-retype"
			value_type     = "single_select"
			allowed_values = ["x"]
		}`

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{Config: before},
				{
					Config: after,
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(resourceAddr, plancheck.ResourceActionDestroyBeforeCreate),
						},
					},
				},
			},
		})
	})

	t.Run("rejects allowed_values on string type", func(t *testing.T) {
		config := `
		resource "github_organization_repository_custom_property" "test" {
			property_name  = "tf-acc-test-invalid"
			value_type     = "string"
			allowed_values = ["nope"]
		}`

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      config,
					ExpectError: regexp.MustCompile("allowed_values must not be set"),
				},
			},
		})
	})

	t.Run("requires allowed_values on single_select type", func(t *testing.T) {
		config := `
		resource "github_organization_repository_custom_property" "test" {
			property_name = "tf-acc-test-missing-allowed"
			value_type    = "single_select"
		}`

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      config,
					ExpectError: regexp.MustCompile("allowed_values is required"),
				},
			},
		})
	})

	t.Run("rejects invalid values_editable_by", func(t *testing.T) {
		config := `
		resource "github_organization_repository_custom_property" "test" {
			property_name      = "tf-acc-test-bad-editable"
			value_type         = "string"
			values_editable_by = "nope"
		}`

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      config,
					ExpectError: regexp.MustCompile("nope"),
				},
			},
		})
	})

	t.Run("updates values_editable_by from org_actors to org_and_repo_actors", func(t *testing.T) {
		before := `
		resource "github_organization_repository_custom_property" "test" {
			property_name      = "tf-acc-test-editable"
			value_type         = "string"
			values_editable_by = "org_actors"
		}`
		after := `
		resource "github_organization_repository_custom_property" "test" {
			property_name      = "tf-acc-test-editable"
			value_type         = "string"
			values_editable_by = "org_and_repo_actors"
		}`

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: before,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceAddr, tfjsonpath.New("values_editable_by"), knownvalue.StringExact("org_actors")),
					},
				},
				{
					Config: after,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceAddr, tfjsonpath.New("values_editable_by"), knownvalue.StringExact("org_and_repo_actors")),
					},
				},
			},
		})
	})

	t.Run("retains values_editable_by set out-of-band when omitted from config", func(t *testing.T) {
		// Mirrors the upstream behaviour where a value set via the UI before
		// Terraform managed the property is reflected back into state via the
		// Computed attribute even when the config omits it.
		propertyName := "tf-acc-test-ui-set"
		configWithField := fmt.Sprintf(`
		resource "github_organization_repository_custom_property" "test" {
			property_name      = %[1]q
			value_type         = "string"
			values_editable_by = "org_and_repo_actors"
		}`, propertyName)
		configWithoutField := fmt.Sprintf(`
		resource "github_organization_repository_custom_property" "test" {
			property_name = %[1]q
			value_type    = "string"
		}`, propertyName)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: configWithField,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceAddr, tfjsonpath.New("values_editable_by"), knownvalue.StringExact("org_and_repo_actors")),
					},
				},
				{
					Config: configWithoutField,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceAddr, tfjsonpath.New("values_editable_by"), knownvalue.StringExact("org_and_repo_actors")),
					},
				},
			},
		})
	})
}
