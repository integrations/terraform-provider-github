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
  default_value = ["dev"]
}
`, name)

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
						statecheck.ExpectKnownValue(resourceAddr, tfjsonpath.New("default_value"), knownvalue.ListExact([]knownvalue.Check{
							knownvalue.StringExact("dev"),
						})),
					},
				},
			},
		})
	})

	t.Run("creates a true_false property with a default value", func(t *testing.T) {
		t.Parallel()

		name := fmt.Sprintf("%s%s", testResourcePrefix, acctest.RandString(testRandomIDLength))
		config := fmt.Sprintf(`
resource "github_organization_repository_custom_property" "test" {
  property_name = %[1]q
  value_type    = "true_false"
  description   = "tf-acc-test true_false property"
  default_value = [%%q]
}
`, name)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, "false"),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceAddr, tfjsonpath.New("default_value"), knownvalue.ListExact([]knownvalue.Check{
							knownvalue.StringExact("false"),
						})),
					},
				},
				{
					Config: fmt.Sprintf(config, "true"),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(resourceAddr, plancheck.ResourceActionUpdate),
						},
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceAddr, tfjsonpath.New("default_value"), knownvalue.ListExact([]knownvalue.Check{
							knownvalue.StringExact("true"),
						})),
					},
				},
				{
					ResourceName:      resourceAddr,
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})

	t.Run("creates a multi_select property with multiple default values", func(t *testing.T) {
		t.Parallel()

		name := fmt.Sprintf("%s%s", testResourcePrefix, acctest.RandString(testRandomIDLength))
		config := fmt.Sprintf(`
resource "github_organization_repository_custom_property" "test" {
  property_name  = %[1]q
  value_type     = "multi_select"
  description    = "tf-acc-test multi_select property"
  allowed_values = ["one", "two", "three"]
  default_value  = %%s
}
`, name)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, `["one", "two"]`),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceAddr, tfjsonpath.New("default_value"), knownvalue.ListExact([]knownvalue.Check{
							knownvalue.StringExact("one"),
							knownvalue.StringExact("two"),
						})),
					},
				},
				{
					Config: fmt.Sprintf(config, `["three"]`),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(resourceAddr, plancheck.ResourceActionUpdate),
						},
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceAddr, tfjsonpath.New("default_value"), knownvalue.ListExact([]knownvalue.Check{
							knownvalue.StringExact("three"),
						})),
					},
				},
				{
					ResourceName:      resourceAddr,
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})

	t.Run("creates a single_select property and grows allowed_values", func(t *testing.T) {
		t.Parallel()

		name := fmt.Sprintf("%s%s", testResourcePrefix, acctest.RandString(testRandomIDLength))
		config := fmt.Sprintf(`
resource "github_organization_repository_custom_property" "test" {
  property_name  = %[1]q
  value_type     = "single_select"
  description    = "tf-acc-test single_select property %%[1]s"
  allowed_values = %%[2]s
}
`, name)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, "initial", `["one"]`),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceAddr, tfjsonpath.New("allowed_values"), knownvalue.ListExact([]knownvalue.Check{
							knownvalue.StringExact("one"),
						})),
					},
				},
				{
					Config: fmt.Sprintf(config, "updated", `["one", "two"]`),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(resourceAddr, plancheck.ResourceActionUpdate),
						},
					},
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
}
`, name)

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

	t.Run("recreates a property deleted outside of terraform", func(t *testing.T) {
		t.Parallel()

		name := fmt.Sprintf("%s%s", testResourcePrefix, acctest.RandString(testRandomIDLength))
		config := fmt.Sprintf(`
resource "github_organization_repository_custom_property" "test" {
  property_name = %[1]q
  value_type    = "string"
  description   = "tf-acc-test out-of-band delete"
}
`, name)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{Config: config},
				{
					// Read must classify the resulting 404 as "gone" and drop the
					// resource from state, so the next plan recreates it rather
					// than erroring.
					PreConfig: func() {
						if _, err := testAccConf.meta.v3client.Organizations.RemoveCustomProperty(t.Context(), testAccConf.meta.name, name); err != nil {
							t.Fatalf("failed to delete organization custom property %s out of band: %v", name, err)
						}
					},
					Config: config,
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(resourceAddr, plancheck.ResourceActionCreate),
						},
					},
				},
			},
		})
	})

	t.Run("destroys cleanly when the property is already gone", func(t *testing.T) {
		t.Parallel()

		name := fmt.Sprintf("%s%s", testResourcePrefix, acctest.RandString(testRandomIDLength))
		config := fmt.Sprintf(`
resource "github_organization_repository_custom_property" "test" {
  property_name = %[1]q
  value_type    = "string"
  description   = "tf-acc-test delete of a missing property"
}
`, name)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{Config: config},
				{
					// Delete must treat a 404 as success. Removing the property
					// out of band immediately before destroy exercises that branch,
					// which the recreate test above never reaches.
					PreConfig: func() {
						if _, err := testAccConf.meta.v3client.Organizations.RemoveCustomProperty(t.Context(), testAccConf.meta.name, name); err != nil {
							t.Fatalf("failed to delete organization custom property %s out of band: %v", name, err)
						}
					},
					Config:  config,
					Destroy: true,
				},
			},
		})
	})

	t.Run("creates a url property with a default value", func(t *testing.T) {
		t.Parallel()

		name := fmt.Sprintf("%s%s", testResourcePrefix, acctest.RandString(testRandomIDLength))
		config := fmt.Sprintf(`
resource "github_organization_repository_custom_property" "test" {
  property_name = %[1]q
  value_type    = "url"
  description   = "tf-acc-test url property"
  default_value = ["https://example.com/runbook"]
}
`, name)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceAddr, tfjsonpath.New("value_type"), knownvalue.StringExact("url")),
						statecheck.ExpectKnownValue(resourceAddr, tfjsonpath.New("default_value"), knownvalue.ListExact([]knownvalue.Check{
							knownvalue.StringExact("https://example.com/runbook"),
						})),
					},
				},
				{
					ResourceName:      resourceAddr,
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})

	t.Run("rejects a non-boolean default_value on true_false", func(t *testing.T) {
		t.Parallel()

		name := fmt.Sprintf("%s%s", testResourcePrefix, acctest.RandString(testRandomIDLength))
		config := fmt.Sprintf(`
resource "github_organization_repository_custom_property" "test" {
  property_name = %[1]q
  value_type    = "true_false"
  default_value = ["True"]
}
`, name)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      config,
					ExpectError: regexp.MustCompile(`default_value must be "true" or "false"`),
				},
			},
		})
	})

	t.Run("rejects an empty string in allowed_values", func(t *testing.T) {
		t.Parallel()

		name := fmt.Sprintf("%s%s", testResourcePrefix, acctest.RandString(testRandomIDLength))
		config := fmt.Sprintf(`
resource "github_organization_repository_custom_property" "test" {
  property_name  = %[1]q
  value_type     = "single_select"
  allowed_values = [""]
}
`, name)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      config,
					ExpectError: regexp.MustCompile("expected .* to not be an empty string"),
				},
			},
		})
	})

	t.Run("forces new when property_name changes", func(t *testing.T) {
		t.Parallel()

		nameBefore := fmt.Sprintf("%s%s", testResourcePrefix, acctest.RandString(testRandomIDLength))
		nameAfter := fmt.Sprintf("%s%s", testResourcePrefix, acctest.RandString(testRandomIDLength))
		config := `
resource "github_organization_repository_custom_property" "test" {
  property_name = %[1]q
  value_type    = "string"
}
`

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{Config: fmt.Sprintf(config, nameBefore)},
				{
					Config: fmt.Sprintf(config, nameAfter),
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
		config := fmt.Sprintf(`
resource "github_organization_repository_custom_property" "test" {
  property_name = %[1]q
  %%s
}
`, name)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{Config: fmt.Sprintf(config, `value_type = "string"`)},
				{
					Config: fmt.Sprintf(config, "value_type    = \"single_select\"\n  allowed_values = [\"x\"]"),
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
}
`, name)

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
}
`, name)

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

	t.Run("rejects multiple default_value entries on a scalar type", func(t *testing.T) {
		t.Parallel()

		name := fmt.Sprintf("%s%s", testResourcePrefix, acctest.RandString(testRandomIDLength))
		config := fmt.Sprintf(`
resource "github_organization_repository_custom_property" "test" {
  property_name = %[1]q
  value_type    = "string"
  default_value = ["one", "two"]
}
`, name)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      config,
					ExpectError: regexp.MustCompile("default_value must contain at most one element"),
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
}
`, name)

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
		config := fmt.Sprintf(`
resource "github_organization_repository_custom_property" "test" {
  property_name      = %[1]q
  value_type         = "string"
  values_editable_by = %%q
}
`, name)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, "org_actors"),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceAddr, tfjsonpath.New("values_editable_by"), knownvalue.StringExact("org_actors")),
					},
				},
				{
					Config: fmt.Sprintf(config, "org_and_repo_actors"),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(resourceAddr, plancheck.ResourceActionUpdate),
						},
					},
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
		name := fmt.Sprintf("%s%s", testResourcePrefix, acctest.RandString(testRandomIDLength))
		config := fmt.Sprintf(`
resource "github_organization_repository_custom_property" "test" {
  property_name = %[1]q
  value_type    = "string"
  %%s
}
`, name)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, `values_editable_by = "org_and_repo_actors"`),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceAddr, tfjsonpath.New("values_editable_by"), knownvalue.StringExact("org_and_repo_actors")),
					},
				},
				{
					Config: fmt.Sprintf(config, ""),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceAddr, tfjsonpath.New("values_editable_by"), knownvalue.StringExact("org_and_repo_actors")),
					},
				},
			},
		})
	})
}
