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

func TestAccGithubOrganizationRepositoryCustomProperty(t *testing.T) {
	const resourceAddr = "github_organization_repository_custom_property.test"

	t.Parallel()

	t.Run("creates a string property without error", func(t *testing.T) {
		t.Parallel()

		name := fmt.Sprintf("%s%s", testResourcePrefix, acctest.RandString(testRandomIDLength))
		config := fmt.Sprintf(`
		resource "github_organization_repository_custom_property" "test" {
			property_name = %[1]q
			value_type    = "string"
			description   = "tf-acc-test string property"
		}`, name)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceAddr, tfjsonpath.New("property_name"), knownvalue.StringExact(name)),
						statecheck.ExpectKnownValue(resourceAddr, tfjsonpath.New("value_type"), knownvalue.StringExact("string")),
						statecheck.ExpectKnownValue(resourceAddr, tfjsonpath.New("values_editable_by"), knownvalue.StringExact("org_actors")),
					},
				},
			},
		})
	})

	t.Run("creates a single_select property and grows allowed_values", func(t *testing.T) {
		t.Parallel()

		name := fmt.Sprintf("%s%s", testResourcePrefix, acctest.RandString(testRandomIDLength))
		configBefore := fmt.Sprintf(`
		resource "github_organization_repository_custom_property" "test" {
			property_name  = %[1]q
			value_type     = "single_select"
			description    = "tf-acc-test single_select property"
			allowed_values = ["one"]
		}`, name)
		configAfter := fmt.Sprintf(`
		resource "github_organization_repository_custom_property" "test" {
			property_name  = %[1]q
			value_type     = "single_select"
			description    = "tf-acc-test single_select property updated"
			allowed_values = ["one", "two"]
		}`, name)

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
		t.Parallel()

		name := fmt.Sprintf("%s%s", testResourcePrefix, acctest.RandString(testRandomIDLength))
		config := fmt.Sprintf(`
		resource "github_organization_repository_custom_property" "test" {
			property_name = %[1]q
			value_type    = "string"
			description   = "tf-acc-test import"
		}`, name)

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
		t.Parallel()

		nameBefore := fmt.Sprintf("%s%s", testResourcePrefix, acctest.RandString(testRandomIDLength))
		nameAfter := fmt.Sprintf("%s%s", testResourcePrefix, acctest.RandString(testRandomIDLength))
		before := fmt.Sprintf(`
		resource "github_organization_repository_custom_property" "test" {
			property_name = %[1]q
			value_type    = "string"
		}`, nameBefore)
		after := fmt.Sprintf(`
		resource "github_organization_repository_custom_property" "test" {
			property_name = %[1]q
			value_type    = "string"
		}`, nameAfter)

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
		t.Parallel()

		name := fmt.Sprintf("%s%s", testResourcePrefix, acctest.RandString(testRandomIDLength))
		before := fmt.Sprintf(`
		resource "github_organization_repository_custom_property" "test" {
			property_name = %[1]q
			value_type    = "string"
		}`, name)
		after := fmt.Sprintf(`
		resource "github_organization_repository_custom_property" "test" {
			property_name  = %[1]q
			value_type     = "single_select"
			allowed_values = ["x"]
		}`, name)

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
		t.Parallel()

		name := fmt.Sprintf("%s%s", testResourcePrefix, acctest.RandString(testRandomIDLength))
		config := fmt.Sprintf(`
		resource "github_organization_repository_custom_property" "test" {
			property_name  = %[1]q
			value_type     = "string"
			allowed_values = ["nope"]
		}`, name)

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
		t.Parallel()

		name := fmt.Sprintf("%s%s", testResourcePrefix, acctest.RandString(testRandomIDLength))
		config := fmt.Sprintf(`
		resource "github_organization_repository_custom_property" "test" {
			property_name = %[1]q
			value_type    = "single_select"
		}`, name)

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
		t.Parallel()

		name := fmt.Sprintf("%s%s", testResourcePrefix, acctest.RandString(testRandomIDLength))
		config := fmt.Sprintf(`
		resource "github_organization_repository_custom_property" "test" {
			property_name      = %[1]q
			value_type         = "string"
			values_editable_by = "nope"
		}`, name)

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
		t.Parallel()

		name := fmt.Sprintf("%s%s", testResourcePrefix, acctest.RandString(testRandomIDLength))
		before := fmt.Sprintf(`
		resource "github_organization_repository_custom_property" "test" {
			property_name      = %[1]q
			value_type         = "string"
			values_editable_by = "org_actors"
		}`, name)
		after := fmt.Sprintf(`
		resource "github_organization_repository_custom_property" "test" {
			property_name      = %[1]q
			value_type         = "string"
			values_editable_by = "org_and_repo_actors"
		}`, name)

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
		t.Parallel()

		// Mirrors the upstream behaviour where a value set via the UI before
		// Terraform managed the property is reflected back into state via the
		// Computed attribute even when the config omits it.
		propertyName := fmt.Sprintf("%s%s", testResourcePrefix, acctest.RandString(testRandomIDLength))
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
