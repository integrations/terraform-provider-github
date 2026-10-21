package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubCodeSecurityConfiguration(t *testing.T) {
	t.Parallel()

	const resourceName = "github_code_security_configuration.test"

	t.Run("creates, updates and imports an organization configuration without error", func(t *testing.T) {
		t.Parallel()

		name := fmt.Sprintf("%s%s", testResourcePrefix, acctest.RandString(5))

		config := fmt.Sprintf(`
resource "github_code_security_configuration" "test" {
  name        = "%s"
  description = "%%s"

  dependency_graph                = "enabled"
  dependabot_alerts               = "%%s"
  private_vulnerability_reporting = "%%s"
  enforcement                     = "%%s"
}
`, name)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, "Terraform acceptance test configuration", "disabled", "disabled", "unenforced"),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("name"), knownvalue.StringExact(name)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("target_type"), knownvalue.StringExact("organization")),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("configuration_id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("html_url"), knownvalue.NotNull()),
					},
				},
				{
					Config: fmt.Sprintf(config, "Terraform acceptance test configuration (updated)", "enabled", "enabled", "enforced"),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
						},
					},
				},
				{
					ResourceName:      resourceName,
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})

	t.Run("manages nested option blocks on an organization configuration without error", func(t *testing.T) {
		t.Parallel()

		name := fmt.Sprintf("%s%s", testResourcePrefix, acctest.RandString(5))

		config := fmt.Sprintf(`
resource "github_team" "test" {
  name = "%s"
}

resource "github_code_security_configuration" "test" {
  name        = "%s"
  description = "Terraform acceptance test nested options"

  advanced_security = "enabled"

  dependency_graph                   = "enabled"
  dependency_graph_autosubmit_action = "enabled"
  dependency_graph_autosubmit_action_options {
    labeled_runners = %%t
  }

  code_scanning_default_setup = "enabled"
  code_scanning_default_setup_options {
    runner_type  = "%%s"
    runner_label = "%%s"
  }
  code_scanning_options {
    allow_advanced = %%t
  }

  secret_scanning                  = "enabled"
  secret_scanning_push_protection  = "enabled"
  secret_scanning_delegated_bypass = "enabled"
  secret_scanning_delegated_bypass_options {
    reviewers {
      reviewer_id   = github_team.test.id
      reviewer_type = "TEAM"
    }
  }

  enforcement = "unenforced"
}
`, name, name)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, false, "standard", "", true),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("dependency_graph_autosubmit_action_options").AtSliceIndex(0).AtMapKey("labeled_runners"), knownvalue.Bool(false)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("code_scanning_default_setup_options").AtSliceIndex(0).AtMapKey("runner_type"), knownvalue.StringExact("standard")),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("code_scanning_options").AtSliceIndex(0).AtMapKey("allow_advanced"), knownvalue.Bool(true)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("secret_scanning_delegated_bypass_options").AtSliceIndex(0).AtMapKey("reviewers"), knownvalue.ListSizeExact(1)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("secret_scanning_delegated_bypass_options").AtSliceIndex(0).AtMapKey("reviewers").AtSliceIndex(0).AtMapKey("reviewer_type"), knownvalue.StringExact("TEAM")),
					},
				},
				{
					Config: fmt.Sprintf(config, true, "labeled", "linux", false),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
						},
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("dependency_graph_autosubmit_action_options").AtSliceIndex(0).AtMapKey("labeled_runners"), knownvalue.Bool(true)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("code_scanning_default_setup_options").AtSliceIndex(0).AtMapKey("runner_type"), knownvalue.StringExact("labeled")),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("code_scanning_default_setup_options").AtSliceIndex(0).AtMapKey("runner_label"), knownvalue.StringExact("linux")),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("code_scanning_options").AtSliceIndex(0).AtMapKey("allow_advanced"), knownvalue.Bool(false)),
					},
				},
				{
					Config: fmt.Sprintf(config, true, "labeled", "linux", false),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
			},
		})
	})

	t.Run("manages default_for_new_repos on an organization configuration without error", func(t *testing.T) {
		t.Parallel()

		name := fmt.Sprintf("%s%s", testResourcePrefix, acctest.RandString(5))

		config := fmt.Sprintf(`
resource "github_code_security_configuration" "test" {
  name        = "%s"
  description = "Terraform acceptance test default configuration"

  dependency_graph  = "enabled"
  dependabot_alerts = "enabled"
  enforcement       = "unenforced"
  %%s
}
`, name)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, `default_for_new_repos = "private_and_internal"`),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("default_for_new_repos"), knownvalue.StringExact("private_and_internal")),
					},
				},
				{
					Config: fmt.Sprintf(config, `default_for_new_repos = "all"`),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
						},
					},
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("default_for_new_repos"), knownvalue.StringExact("all")),
					},
				},
				{
					Config: fmt.Sprintf(config, ""),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("default_for_new_repos"), knownvalue.StringExact("")),
					},
				},
				{
					Config: fmt.Sprintf(config, ""),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectEmptyPlan(),
						},
					},
				},
			},
		})
	})

	t.Run("creates, updates and imports an enterprise configuration without error", func(t *testing.T) {
		t.Parallel()

		name := fmt.Sprintf("%s%s", testResourcePrefix, acctest.RandString(5))

		config := fmt.Sprintf(`
resource "github_code_security_configuration" "test" {
  enterprise_slug = "%s"
  name            = "%s"
  description     = "Terraform acceptance test enterprise configuration"

  dependency_graph  = "enabled"
  dependabot_alerts = "%%s"
  enforcement       = "unenforced"
}
`, testAccConf.enterpriseSlug, name)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, "disabled"),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("enterprise_slug"), knownvalue.StringExact(testAccConf.enterpriseSlug)),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("target_type"), knownvalue.StringExact("enterprise")),
						statecheck.ExpectKnownValue(resourceName, tfjsonpath.New("configuration_id"), knownvalue.NotNull()),
					},
				},
				{
					Config: fmt.Sprintf(config, "enabled"),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
						},
					},
				},
				{
					ResourceName:      resourceName,
					ImportState:       true,
					ImportStateVerify: true,
					ImportStateIdFunc: importCodeSecurityConfigurationByEnterprise(resourceName),
				},
			},
		})
	})
}

// importCodeSecurityConfigurationByEnterprise builds an import ID of the form
// <enterprise_slug>:<configuration_id> from the resource in state.
func importCodeSecurityConfigurationByEnterprise(logicalName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs := s.RootModule().Resources[logicalName]
		if rs == nil {
			return "", fmt.Errorf("cannot find %s in terraform state", logicalName)
		}
		return fmt.Sprintf("%s:%s", testAccConf.enterpriseSlug, rs.Primary.ID), nil
	}
}
