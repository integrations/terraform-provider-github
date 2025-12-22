package github

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubEnterpriseTeamsDataSource(t *testing.T) {
	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

	if isEnterprise != "true" {
		t.Skip("Skipping because `ENTERPRISE_ACCOUNT` is not set or set to false")
	}
	if testEnterprise == "" {
		t.Skip("Skipping because `ENTERPRISE_SLUG` is not set")
	}

	config := fmt.Sprintf(`
		data "github_enterprise" "enterprise" {
			slug = "%s"
		}

		resource "github_enterprise_team" "test" {
			enterprise_slug = data.github_enterprise.enterprise.slug
			name            = "tf-acc-ds-enterprise-teams-%s"
		}

		data "github_enterprise_teams" "all" {
			enterprise_slug = data.github_enterprise.enterprise.slug
			depends_on      = [github_enterprise_team.test]
		}
	`, testEnterprise, randomID)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { skipUnlessMode(t, enterprise) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.github_enterprise_teams.all", "id"),
					resource.TestCheckResourceAttrSet("data.github_enterprise_teams.all", "teams.0.team_id"),
					resource.TestCheckResourceAttrSet("data.github_enterprise_teams.all", "teams.0.slug"),
					resource.TestCheckResourceAttrSet("data.github_enterprise_teams.all", "teams.0.name"),
				),
			},
		},
	})
}
