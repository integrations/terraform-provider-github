package github

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubCopilotEnterpriseSeatsDataSource(t *testing.T) {
	t.Parallel()

	skipUnlessEnterprise(t)

	t.Run("lists the Copilot seats of the enterprise", func(t *testing.T) {
		t.Parallel()

		config := fmt.Sprintf(`
data "github_copilot_enterprise_seats" "test" {
  enterprise_slug = "%s"
}
`, testAccConf.enterpriseSlug)

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_copilot_enterprise_seats.test", tfjsonpath.New("id"), knownvalue.StringExact(testAccConf.enterpriseSlug)),
						statecheck.ExpectKnownValue("data.github_copilot_enterprise_seats.test", tfjsonpath.New("total_seats"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_copilot_enterprise_seats.test", tfjsonpath.New("seats"), knownvalue.NotNull()),
					},
				},
			},
		})
	})

	t.Run("returns an error for an unknown enterprise", func(t *testing.T) {
		t.Parallel()

		config := fmt.Sprintf(`
data "github_copilot_enterprise_seats" "test" {
  enterprise_slug = "%senterprise-%s"
}
`, testResourcePrefix, acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum))

		resource.Test(t, resource.TestCase{
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config:      config,
					ExpectError: regexp.MustCompile(`Not Found`),
				},
			},
		})
	})
}
