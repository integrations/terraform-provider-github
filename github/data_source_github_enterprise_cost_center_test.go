package github

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubEnterpriseCostCenterDataSource(t *testing.T) {
	skipUnlessEnterprise(t)
	skipUnlessHasOrgUser1(t)

	cc := mustCreateTestEnterpriseCostCenter(t)
	repo := mustCreateTestRepository(t)
	mustAddTestEnterpriseCostCenterResources(t, cc, github.CostCenterResourceRequest{
		Users:         []string{testAccConf.testOrgUser1},
		Organizations: []string{testAccConf.owner},
		Repositories:  []string{repo.GetFullName()},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { skipUnlessEnterprise(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{{
			Config: fmt.Sprintf(`
				data "github_enterprise_cost_center" "test" {
					enterprise_slug = "%s"
					cost_center_id  = "%s"
				}
			`, testAccConf.enterpriseSlug, cc.ID),
			ConfigStateChecks: []statecheck.StateCheck{
				statecheck.ExpectKnownValue("data.github_enterprise_cost_center.test", tfjsonpath.New("id"), knownvalue.StringExact(cc.ID)),
				statecheck.ExpectKnownValue("data.github_enterprise_cost_center.test", tfjsonpath.New("enterprise_slug"), knownvalue.StringExact(testAccConf.enterpriseSlug)),
				statecheck.ExpectKnownValue("data.github_enterprise_cost_center.test", tfjsonpath.New("cost_center_id"), knownvalue.StringExact(cc.ID)),
				statecheck.ExpectKnownValue("data.github_enterprise_cost_center.test", tfjsonpath.New("name"), knownvalue.StringExact(cc.Name)),
				statecheck.ExpectKnownValue("data.github_enterprise_cost_center.test", tfjsonpath.New("state"), knownvalue.StringExact("active")),
				statecheck.ExpectKnownValue("data.github_enterprise_cost_center.test", tfjsonpath.New("azure_subscription"), knownvalue.StringExact(cc.GetAzureSubscription())),
				statecheck.ExpectKnownValue("data.github_enterprise_cost_center.test", tfjsonpath.New("users"), knownvalue.SetExact([]knownvalue.Check{knownvalue.StringRegexp(caseInsensitiveExactRegexp(testAccConf.testOrgUser1))})),
				statecheck.ExpectKnownValue("data.github_enterprise_cost_center.test", tfjsonpath.New("organizations"), knownvalue.SetExact([]knownvalue.Check{knownvalue.StringRegexp(caseInsensitiveExactRegexp(testAccConf.owner))})),
				statecheck.ExpectKnownValue("data.github_enterprise_cost_center.test", tfjsonpath.New("repositories"), knownvalue.SetExact([]knownvalue.Check{knownvalue.StringRegexp(caseInsensitiveExactRegexp(repo.GetFullName()))})),
			},
		}},
	})
}

func caseInsensitiveExactRegexp(value string) *regexp.Regexp {
	return regexp.MustCompile("(?i)^" + regexp.QuoteMeta(value) + "$")
}
