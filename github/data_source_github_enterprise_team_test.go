package github

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/compare"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
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
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_enterprise_team.by_slug", tfjsonpath.New("id"), knownvalue.NotNull()),
						statecheck.CompareValuePairs("data.github_enterprise_team.by_slug", tfjsonpath.New("team_id"), "github_enterprise_team.test", tfjsonpath.New("team_id"), compare.ValuesSame()),
						statecheck.CompareValuePairs("data.github_enterprise_team.by_slug", tfjsonpath.New("slug"), "github_enterprise_team.test", tfjsonpath.New("slug"), compare.ValuesSame()),
						statecheck.CompareValuePairs("data.github_enterprise_team.by_slug", tfjsonpath.New("name"), "github_enterprise_team.test", tfjsonpath.New("name"), compare.ValuesSame()),
					},
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
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_enterprise_team.by_id", tfjsonpath.New("id"), knownvalue.NotNull()),
						statecheck.CompareValuePairs("data.github_enterprise_team.by_id", tfjsonpath.New("team_id"), "github_enterprise_team.test", tfjsonpath.New("team_id"), compare.ValuesSame()),
						statecheck.CompareValuePairs("data.github_enterprise_team.by_id", tfjsonpath.New("slug"), "github_enterprise_team.test", tfjsonpath.New("slug"), compare.ValuesSame()),
					},
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
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_enterprise_team_organizations.test", tfjsonpath.New("id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_enterprise_team_organizations.test", tfjsonpath.New("organization_slugs"), knownvalue.SetSizeExact(1)),
						statecheck.ExpectKnownValue("data.github_enterprise_team_organizations.test", tfjsonpath.New("organization_slugs"), knownvalue.SetPartial([]knownvalue.Check{knownvalue.StringExact(orgSlug)})),
					},
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
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_enterprise_team_membership.test", tfjsonpath.New("id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("data.github_enterprise_team_membership.test", tfjsonpath.New("username"), knownvalue.StringExact(username)),
					},
				},
			},
		})
	})
}
