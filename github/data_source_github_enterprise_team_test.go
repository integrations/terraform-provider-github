package github

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubEnterpriseTeamDataSource(t *testing.T) {
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
			name            = "tf-acc-ds-team-%s"
		}

		data "github_enterprise_team" "by_slug" {
			enterprise_slug = data.github_enterprise.enterprise.slug
			slug            = github_enterprise_team.test.slug
		}

		data "github_enterprise_team" "by_id" {
			enterprise_slug = data.github_enterprise.enterprise.slug
			team_id         = github_enterprise_team.test.team_id
		}
	`, testEnterprise, randomID)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { skipUnlessMode(t, enterprise) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.github_enterprise_team.by_slug", "id"),
					resource.TestCheckResourceAttrPair("data.github_enterprise_team.by_slug", "team_id", "github_enterprise_team.test", "team_id"),
					resource.TestCheckResourceAttrPair("data.github_enterprise_team.by_slug", "slug", "github_enterprise_team.test", "slug"),
					resource.TestCheckResourceAttrPair("data.github_enterprise_team.by_slug", "name", "github_enterprise_team.test", "name"),
					resource.TestCheckResourceAttrSet("data.github_enterprise_team.by_id", "id"),
					resource.TestCheckResourceAttrPair("data.github_enterprise_team.by_id", "team_id", "github_enterprise_team.test", "team_id"),
					resource.TestCheckResourceAttrPair("data.github_enterprise_team.by_id", "slug", "github_enterprise_team.test", "slug"),
				),
			},
		},
	})
}

func TestAccGithubEnterpriseTeamOrganizationsDataSource(t *testing.T) {
	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

	if isEnterprise != "true" {
		t.Skip("Skipping because `ENTERPRISE_ACCOUNT` is not set or set to false")
	}
	if testEnterprise == "" {
		t.Skip("Skipping because `ENTERPRISE_SLUG` is not set")
	}
	if testOrganization == "" {
		t.Skip("Skipping because `GITHUB_OWNER`/`GITHUB_ORGANIZATION` is not set")
	}

	config := fmt.Sprintf(`
		data "github_enterprise" "enterprise" {
			slug = "%s"
		}

		resource "github_enterprise_team" "test" {
			enterprise_slug             = data.github_enterprise.enterprise.slug
			name                        = "tf-acc-ds-team-orgs-%s"
			organization_selection_type = "selected"
		}

		resource "github_enterprise_team_organizations" "assign" {
			enterprise_slug    = data.github_enterprise.enterprise.slug
			enterprise_team    = github_enterprise_team.test.slug
			organization_slugs = ["%s"]
		}

		data "github_enterprise_team_organizations" "test" {
			enterprise_slug = data.github_enterprise.enterprise.slug
			enterprise_team = github_enterprise_team.test.slug
			depends_on      = [github_enterprise_team_organizations.assign]
		}
	`, testEnterprise, randomID, testOrganization)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { skipUnlessMode(t, enterprise) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.github_enterprise_team_organizations.test", "id"),
					resource.TestCheckResourceAttr("data.github_enterprise_team_organizations.test", "organization_slugs.#", "1"),
					resource.TestCheckTypeSetElemAttr("data.github_enterprise_team_organizations.test", "organization_slugs.*", testOrganization),
				),
			},
		},
	})
}

func TestAccGithubEnterpriseTeamMembershipDataSource(t *testing.T) {
	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)
	username := os.Getenv("GITHUB_TEST_USER")

	if isEnterprise != "true" {
		t.Skip("Skipping because `ENTERPRISE_ACCOUNT` is not set or set to false")
	}
	if testEnterprise == "" {
		t.Skip("Skipping because `ENTERPRISE_SLUG` is not set")
	}
	if username == "" {
		t.Skip("Skipping because `GITHUB_TEST_USER` is not set")
	}

	config := fmt.Sprintf(`
		data "github_enterprise" "enterprise" {
			slug = "%s"
		}

		resource "github_enterprise_team" "test" {
			enterprise_slug = data.github_enterprise.enterprise.slug
			name            = "tf-acc-ds-team-member-%s"
		}

		resource "github_enterprise_team_membership" "test" {
			enterprise_slug = data.github_enterprise.enterprise.slug
			enterprise_team = github_enterprise_team.test.slug
			username        = "%s"
		}

		data "github_enterprise_team_membership" "test" {
			enterprise_slug = data.github_enterprise.enterprise.slug
			enterprise_team = github_enterprise_team.test.slug
			username        = "%s"
			depends_on      = [github_enterprise_team_membership.test]
		}
	`, testEnterprise, randomID, username, username)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { skipUnlessMode(t, enterprise) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.github_enterprise_team_membership.test", "id"),
					resource.TestCheckResourceAttr("data.github_enterprise_team_membership.test", "username", username),
				),
			},
		},
	})
}
