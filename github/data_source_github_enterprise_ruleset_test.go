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

func TestAccGithubEnterpriseRulesetDataSource(t *testing.T) {
	t.Run("queries an enterprise ruleset", func(t *testing.T) {
		randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
		testRulesetName := fmt.Sprintf("%senterprise-ruleset-%s", testResourcePrefix, randomID)

		config := fmt.Sprintf(`
			resource "github_enterprise_ruleset" "test" {
				enterprise_slug = "%s"
				name            = "%s"
				target          = "branch"
				enforcement     = "active"

				conditions {
					organization_name {
						include = ["~ALL"]
						exclude = []
					}

					repository_name {
						include = ["~ALL"]
						exclude = []
					}

					ref_name {
						include = ["refs/heads/main"]
						exclude = []
					}
				}

				rules {
					creation = false
					deletion = false
				}
			}

			data "github_enterprise_ruleset" "test" {
				enterprise_slug = github_enterprise_ruleset.test.enterprise_slug
				ruleset_id      = github_enterprise_ruleset.test.ruleset_id

				depends_on = [github_enterprise_ruleset.test]
			}
		`, testAccConf.enterpriseSlug, testRulesetName)

		resource.Test(t, resource.TestCase{
			PreCheck: func() {
				skipUnlessEnterprise(t)
			},
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_enterprise_ruleset.test", tfjsonpath.New("name"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_enterprise_ruleset.test", tfjsonpath.New("name"), knownvalue.StringExact(testRulesetName)),
						statecheck.ExpectKnownValue("data.github_enterprise_ruleset.test", tfjsonpath.New("target"), knownvalue.StringExact("branch")),
						statecheck.ExpectKnownValue("data.github_enterprise_ruleset.test", tfjsonpath.New("enforcement"), knownvalue.StringExact("active")),
						statecheck.ExpectKnownValue("data.github_enterprise_ruleset.test", tfjsonpath.New("node_id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_enterprise_ruleset.test", tfjsonpath.New("etag"), knownvalue.NotNull()),
					},
				},
			},
		})
	})
}
