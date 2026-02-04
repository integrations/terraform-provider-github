package github

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubEnterpriseTeamDataSource(t *testing.T) {
	t.Run("retrieves team by slug without error", func(t *testing.T) {
		randomID := acctest.RandString(5)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessMode(t, enterprise) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(`
						data "github_enterprise" "enterprise" {
							slug = "%s"
						}

						resource "github_enterprise_team" "test" {
							enterprise_slug = data.github_enterprise.enterprise.slug
							name            = "%s%s"
						}

						data "github_enterprise_team" "by_slug" {
							enterprise_slug = data.github_enterprise.enterprise.slug
							slug            = github_enterprise_team.test.slug
						}
					`, testAccConf.enterpriseSlug, testResourcePrefix, randomID),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttrSet("data.github_enterprise_team.by_slug", "id"),
						resource.TestCheckResourceAttrPair("data.github_enterprise_team.by_slug", "team_id", "github_enterprise_team.test", "team_id"),
						resource.TestCheckResourceAttrPair("data.github_enterprise_team.by_slug", "slug", "github_enterprise_team.test", "slug"),
						resource.TestCheckResourceAttrPair("data.github_enterprise_team.by_slug", "name", "github_enterprise_team.test", "name"),
					),
				},
			},
		})
	})

	t.Run("retrieves team by id without error", func(t *testing.T) {
		randomID := acctest.RandString(5)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessMode(t, enterprise) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(`
						data "github_enterprise" "enterprise" {
							slug = "%s"
						}

						resource "github_enterprise_team" "test" {
							enterprise_slug = data.github_enterprise.enterprise.slug
							name            = "%s%s"
						}

						data "github_enterprise_team" "by_id" {
							enterprise_slug = data.github_enterprise.enterprise.slug
							team_id         = github_enterprise_team.test.team_id
						}
					`, testAccConf.enterpriseSlug, testResourcePrefix, randomID),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttrSet("data.github_enterprise_team.by_id", "id"),
						resource.TestCheckResourceAttrPair("data.github_enterprise_team.by_id", "team_id", "github_enterprise_team.test", "team_id"),
						resource.TestCheckResourceAttrPair("data.github_enterprise_team.by_id", "slug", "github_enterprise_team.test", "slug"),
					),
				},
			},
		})
	})
}

func TestAccGithubEnterpriseTeamOrganizationsDataSource(t *testing.T) {
	orgSlug := os.Getenv("ENTERPRISE_TEST_ORGANIZATION")
	if orgSlug == "" {
		t.Skip("ENTERPRISE_TEST_ORGANIZATION not set")
	}

	t.Run("retrieves team organizations without error", func(t *testing.T) {
		randomID := acctest.RandString(5)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessMode(t, enterprise) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(`
						data "github_enterprise" "enterprise" {
							slug = "%s"
						}

						resource "github_enterprise_team" "test" {
							enterprise_slug             = data.github_enterprise.enterprise.slug
							name                        = "%s%s"
							organization_selection_type = "selected"
						}

						resource "github_enterprise_team_organizations" "assign" {
							enterprise_slug    = data.github_enterprise.enterprise.slug
							team_slug          = github_enterprise_team.test.slug
							organization_slugs = [%q]
						}

						data "github_enterprise_team_organizations" "test" {
							enterprise_slug = data.github_enterprise.enterprise.slug
							team_slug       = github_enterprise_team.test.slug
							depends_on      = [github_enterprise_team_organizations.assign]
						}
					`, testAccConf.enterpriseSlug, testResourcePrefix, randomID, orgSlug),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttrSet("data.github_enterprise_team_organizations.test", "id"),
						resource.TestCheckResourceAttr("data.github_enterprise_team_organizations.test", "organization_slugs.#", "1"),
						resource.TestCheckTypeSetElemAttr("data.github_enterprise_team_organizations.test", "organization_slugs.*", orgSlug),
					),
				},
			},
		})
	})
}

func TestAccGithubEnterpriseTeamMembershipDataSource(t *testing.T) {
	username := os.Getenv("ENTERPRISE_TEST_USER")
	if username == "" {
		t.Skip("ENTERPRISE_TEST_USER not set")
	}

	t.Run("retrieves team membership without error", func(t *testing.T) {
		randomID := acctest.RandString(5)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessMode(t, enterprise) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: fmt.Sprintf(`
						data "github_enterprise" "enterprise" {
							slug = "%s"
						}

						resource "github_enterprise_team" "test" {
							enterprise_slug = data.github_enterprise.enterprise.slug
							name            = "%s%s"
						}

						resource "github_enterprise_team_membership" "test" {
							enterprise_slug = data.github_enterprise.enterprise.slug
							team_slug       = github_enterprise_team.test.slug
							username        = %q
						}

						data "github_enterprise_team_membership" "test" {
							enterprise_slug = data.github_enterprise.enterprise.slug
							team_slug       = github_enterprise_team.test.slug
							username        = %q
							depends_on      = [github_enterprise_team_membership.test]
						}
					`, testAccConf.enterpriseSlug, testResourcePrefix, randomID, username, username),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttrSet("data.github_enterprise_team_membership.test", "id"),
						resource.TestCheckResourceAttr("data.github_enterprise_team_membership.test", "username", username),
					),
				},
			},
		})
	})
}
