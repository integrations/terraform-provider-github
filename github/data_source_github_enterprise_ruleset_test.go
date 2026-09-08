package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubEnterpriseRulesetDataSource(t *testing.T) {
	t.Parallel()

	t.Run("queries an enterprise ruleset without error", func(t *testing.T) {
		t.Parallel()

		skipUnlessEnterprise(t)

		ruleset := mustCreateTestEnterpriseRuleset(t)

		config := fmt.Sprintf(`
			data "github_enterprise_ruleset" "test" {
			  enterprise_slug = "%s"
			  ruleset_id      = %d
			}
		`, testAccConf.enterpriseSlug, ruleset.GetID())

		const dataSource = "data.github_enterprise_ruleset.test"

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue(dataSource, tfjsonpath.New("name"), knownvalue.StringExact(ruleset.Name)),
						statecheck.ExpectKnownValue(dataSource, tfjsonpath.New("target"), knownvalue.StringExact("branch")),
						statecheck.ExpectKnownValue(dataSource, tfjsonpath.New("enforcement"), knownvalue.StringExact("active")),
						statecheck.ExpectKnownValue(dataSource, tfjsonpath.New("etag"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue(dataSource, tfjsonpath.New("node_id"), knownvalue.StringExact(ruleset.GetNodeID())),
						statecheck.ExpectKnownValue(dataSource, tfjsonpath.New("ruleset_id"), knownvalue.Int64Exact(ruleset.GetID())),
					},
				},
			},
		})
	})
}
