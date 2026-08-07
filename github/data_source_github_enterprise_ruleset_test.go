package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/compare"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubEnterpriseRulesetDataSource(t *testing.T) {
	t.Parallel()

	t.Run("queries an enterprise ruleset without error", func(t *testing.T) {
		t.Parallel()

		name := fmt.Sprintf("%senterprise-ds-%s", testResourcePrefix, acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum))
		config := fmt.Sprintf(`
			%s

			data "github_enterprise_ruleset" "test" {
			  enterprise_slug = github_enterprise_ruleset.test.enterprise_slug
			  ruleset_id      = github_enterprise_ruleset.test.ruleset_id
			}
		`, fmt.Sprintf(enterpriseRulesetBasicConfig, testAccConf.enterpriseSlug, name))

		const dataSource = "data.github_enterprise_ruleset.test"

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(dataSource, tfjsonpath.New("name"), knownvalue.StringExact(name)),
						statecheck.ExpectKnownValue(dataSource, tfjsonpath.New("target"), knownvalue.StringExact("branch")),
						statecheck.ExpectKnownValue(dataSource, tfjsonpath.New("enforcement"), knownvalue.StringExact("active")),
						statecheck.ExpectKnownValue(dataSource, tfjsonpath.New("etag"), knownvalue.NotNull()),
						statecheck.CompareValuePairs(dataSource, tfjsonpath.New("node_id"), enterpriseRulesetResource, tfjsonpath.New("node_id"), compare.ValuesSame()),
						statecheck.CompareValuePairs(dataSource, tfjsonpath.New("ruleset_id"), enterpriseRulesetResource, tfjsonpath.New("ruleset_id"), compare.ValuesSame()),
					},
				},
			},
		})
	})
}
