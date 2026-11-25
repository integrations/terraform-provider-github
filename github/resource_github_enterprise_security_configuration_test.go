package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubEnterpriseSecurityConfiguration(t *testing.T) {
	t.Parallel()

	skipUnlessEnterprise(t)

	// General lifecycle: create -> update -> import. Only the identifying/computed attributes
	// round-trip through import (the resource reconciles just the managed attributes in Read, per
	// the Optional-only design), so this scenario is name-only to keep ImportStateVerify clean.
	t.Run("basic", func(t *testing.T) {
		t.Parallel()

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		configName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)
		configNameUpdated := fmt.Sprintf("%supdated-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
resource "github_enterprise_security_configuration" "test" {
  enterprise_slug = %q
  name = "%%s"
}`, testAccConf.enterpriseSlug)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(config, configName),
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_enterprise_security_configuration.test", tfjsonpath.New("configuration_id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_enterprise_security_configuration.test", tfjsonpath.New("target_type"), knownvalue.NotNull()),
					},
				},
				{
					Config: fmt.Sprintf(config, configNameUpdated),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction("github_enterprise_security_configuration.test", plancheck.ResourceActionUpdate),
						},
					},
				},
				{
					ResourceName:      "github_enterprise_security_configuration.test",
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})

	t.Run("with settings", func(t *testing.T) {
		t.Parallel()

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		configName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
resource "github_enterprise_security_configuration" "test" {
  enterprise_slug = %q
  name = "%s"
  description = "Test configuration"
  advanced_security = "enabled"
  dependency_graph = "enabled"
  dependabot_alerts = "enabled"
  dependabot_security_updates = "enabled"
  code_scanning_default_setup = "enabled"
  secret_scanning = "enabled"
  secret_scanning_push_protection = "enabled"
  private_vulnerability_reporting = "enabled"
  enforcement = "enforced"
}`, testAccConf.enterpriseSlug, configName)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_enterprise_security_configuration.test", tfjsonpath.New("configuration_id"), knownvalue.NotNull()),
					},
				},
			},
		})
	})

	t.Run("with options", func(t *testing.T) {
		t.Parallel()

		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		configName := fmt.Sprintf("%s%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
resource "github_enterprise_security_configuration" "test" {
  enterprise_slug = %q
  name = "%s"
  description = "Test configuration with options"
  advanced_security = "enabled"
  dependency_graph = "enabled"
  dependency_graph_autosubmit_action = "enabled"
  dependency_graph_autosubmit_action_options {
    labeled_runners = true
  }
  code_scanning_default_setup = "enabled"
  code_scanning_default_setup_options {
    runner_type = "labeled"
    runner_label = "code-scanning"
  }
  secret_scanning = "enabled"
}`, testAccConf.enterpriseSlug, configName)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_enterprise_security_configuration.test", tfjsonpath.New("configuration_id"), knownvalue.NotNull()),
					},
				},
			},
		})
	})
}
