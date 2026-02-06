package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
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
		`, testAccConf.enterpriseSlug, testRulesetName)

		config2 := config + `
			data "github_enterprise_ruleset" "test" {
				enterprise_slug = github_enterprise_ruleset.test.enterprise_slug
				ruleset_id      = github_enterprise_ruleset.test.ruleset_id
			}
		`

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet(
				"data.github_enterprise_ruleset.test", "name",
			),
			resource.TestCheckResourceAttr(
				"data.github_enterprise_ruleset.test", "name",
				testRulesetName,
			),
			resource.TestCheckResourceAttr(
				"data.github_enterprise_ruleset.test", "target",
				"branch",
			),
			resource.TestCheckResourceAttr(
				"data.github_enterprise_ruleset.test", "enforcement",
				"active",
			),
			resource.TestCheckResourceAttrSet(
				"data.github_enterprise_ruleset.test", "node_id",
			),
			resource.TestCheckResourceAttrSet(
				"data.github_enterprise_ruleset.test", "etag",
			),
		)

		resource.Test(t, resource.TestCase{
			PreCheck: func() {
				skipUnlessEnterprise(t)
			},
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					Check:  resource.ComposeTestCheckFunc(),
				},
				{
					Config: config2,
					Check:  check,
				},
			},
		})
	})
}
