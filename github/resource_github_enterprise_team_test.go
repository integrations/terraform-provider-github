package github

import (
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubEnterpriseTeam(t *testing.T) {
	t.Run("creates and updates resource without error", func(t *testing.T) {
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
							description                 = "team for acceptance testing"
							organization_selection_type = "disabled"
						}
					`, testAccConf.enterpriseSlug, testResourcePrefix, randomID),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttrSet("github_enterprise_team.test", "slug"),
						resource.TestCheckResourceAttrSet("github_enterprise_team.test", "team_id"),
						resource.TestCheckResourceAttr("github_enterprise_team.test", "organization_selection_type", "disabled"),
					),
				},
				{
					Config: fmt.Sprintf(`
						data "github_enterprise" "enterprise" {
							slug = "%s"
						}

						resource "github_enterprise_team" "test" {
							enterprise_slug             = data.github_enterprise.enterprise.slug
							name                        = "%s%s"
							description                 = "updated description"
							organization_selection_type = "selected"
						}
					`, testAccConf.enterpriseSlug, testResourcePrefix, randomID),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("github_enterprise_team.test", "description", "updated description"),
						resource.TestCheckResourceAttr("github_enterprise_team.test", "organization_selection_type", "selected"),
					),
				},
			},
		})
	})

	t.Run("imports resource without error", func(t *testing.T) {
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
							description                 = "team for import testing"
							organization_selection_type = "disabled"
						}
					`, testAccConf.enterpriseSlug, testResourcePrefix, randomID),
				},
				{
					ResourceName:        "github_enterprise_team.test",
					ImportState:         true,
					ImportStateVerify:   true,
					ImportStateIdPrefix: fmt.Sprintf(`%s/`, testAccConf.enterpriseSlug),
				},
			},
		})
	})
}

func TestAccGithubEnterpriseTeamOrganizations(t *testing.T) {
	orgSlug := os.Getenv("ENTERPRISE_TEST_ORGANIZATION")
	if orgSlug == "" {
		t.Skip("ENTERPRISE_TEST_ORGANIZATION not set")
	}

	t.Run("assigns organizations to team without error", func(t *testing.T) {
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

						resource "github_enterprise_team_organizations" "test" {
							enterprise_slug    = data.github_enterprise.enterprise.slug
							team_slug          = github_enterprise_team.test.slug
							organization_slugs = [%q]
						}
					`, testAccConf.enterpriseSlug, testResourcePrefix, randomID, orgSlug),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("github_enterprise_team_organizations.test", "organization_slugs.#", "1"),
						resource.TestCheckTypeSetElemAttr("github_enterprise_team_organizations.test", "organization_slugs.*", orgSlug),
					),
				},
			},
		})
	})

	t.Run("imports resource without error", func(t *testing.T) {
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

						resource "github_enterprise_team_organizations" "test" {
							enterprise_slug    = data.github_enterprise.enterprise.slug
							team_slug          = github_enterprise_team.test.slug
							organization_slugs = [%q]
						}
					`, testAccConf.enterpriseSlug, testResourcePrefix, randomID, orgSlug),
				},
				{
					ResourceName:      "github_enterprise_team_organizations.test",
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})

	t.Run("errors on empty organizations", func(t *testing.T) {
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

						resource "github_enterprise_team_organizations" "test" {
							enterprise_slug    = data.github_enterprise.enterprise.slug
							team_slug          = github_enterprise_team.test.slug
							organization_slugs = []
						}
					`, testAccConf.enterpriseSlug, testResourcePrefix, randomID),
					ExpectError: regexp.MustCompile(`Attribute organization_slugs requires 1 item minimum`),
				},
			},
		})
	})
}

func TestAccGithubEnterpriseTeamMembership(t *testing.T) {
	username := os.Getenv("ENTERPRISE_TEST_USER")
	if username == "" {
		t.Skip("ENTERPRISE_TEST_USER not set")
	}

	t.Run("adds member to team without error", func(t *testing.T) {
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
					`, testAccConf.enterpriseSlug, testResourcePrefix, randomID, username),
					Check: resource.ComposeAggregateTestCheckFunc(
						resource.TestCheckResourceAttr("github_enterprise_team_membership.test", "username", username),
					),
				},
			},
		})
	})

	t.Run("imports resource without error", func(t *testing.T) {
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
					`, testAccConf.enterpriseSlug, testResourcePrefix, randomID, username),
				},
				{
					ResourceName:      "github_enterprise_team_membership.test",
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	})
}
