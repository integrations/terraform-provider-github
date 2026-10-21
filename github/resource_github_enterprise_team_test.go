package github

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/compare"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubEnterpriseTeam(t *testing.T) {
	t.Run("creates and updates resource without error", func(t *testing.T) {
		skipWithoutAccConf(t)
		randomID := acctest.RandString(5)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
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
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_enterprise_team.test", tfjsonpath.New("slug"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_enterprise_team.test", tfjsonpath.New("team_id"), knownvalue.NotNull()),
						statecheck.ExpectKnownValue("github_enterprise_team.test", tfjsonpath.New("organization_selection_type"), knownvalue.StringExact("disabled")),
					},
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
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_enterprise_team.test", tfjsonpath.New("description"), knownvalue.StringExact("updated description")),
						statecheck.ExpectKnownValue("github_enterprise_team.test", tfjsonpath.New("organization_selection_type"), knownvalue.StringExact("selected")),
					},
				},
			},
		})
	})

	t.Run("imports resource without error", func(t *testing.T) {
		skipWithoutAccConf(t)
		randomID := acctest.RandString(5)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
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
					ResourceName:            "github_enterprise_team.test",
					ImportState:             true,
					ImportStateVerify:       true,
					ImportStateVerifyIgnore: []string{"group_id"},
					ImportStateIdPrefix:     fmt.Sprintf(`%s/`, testAccConf.enterpriseSlug),
				},
			},
		})
	})
}

func TestAccGithubEnterpriseTeamOrganizations(t *testing.T) {
	skipWithoutAccConf(t)
	orgSlug := testAccConf.testEnterpriseOrg
	if orgSlug == "" {
		t.Skip("GH_TEST_ENTERPRISE_ORG not set")
	}

	t.Run("assigns organizations to team without error", func(t *testing.T) {
		skipWithoutAccConf(t)
		randomID := acctest.RandString(5)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
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
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_enterprise_team_organizations.test", tfjsonpath.New("organization_slugs"), knownvalue.SetSizeExact(1)),
						statecheck.ExpectKnownValue("github_enterprise_team_organizations.test", tfjsonpath.New("organization_slugs"), knownvalue.SetPartial([]knownvalue.Check{knownvalue.StringExact(strings.ToLower(orgSlug))})),
					},
				},
			},
		})
	})

	t.Run("imports resource without error", func(t *testing.T) {
		skipWithoutAccConf(t)
		randomID := acctest.RandString(5)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
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
}

func TestAccGithubEnterpriseTeamMembership(t *testing.T) {
	skipWithoutAccConf(t)
	username := testAccConf.testEnterpriseUser
	if username == "" {
		t.Skip("GH_TEST_ENTERPRISE_USER not set")
	}

	t.Run("adds member to team without error", func(t *testing.T) {
		skipWithoutAccConf(t)
		randomID := acctest.RandString(5)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
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
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("github_enterprise_team_membership.test", tfjsonpath.New("username"), knownvalue.StringExact(username)),
					},
				},
			},
		})
	})

	t.Run("imports resource without error", func(t *testing.T) {
		skipWithoutAccConf(t)
		randomID := acctest.RandString(5)

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessEnterprise(t) },
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

func TestAccGithubEnterpriseTeamRenameAndClearDescription(t *testing.T) {
	skipWithoutAccConf(t)
	name := testResourcePrefix + acctest.RandString(5)
	config := func(name, description string) string {
		return fmt.Sprintf(`
resource "github_enterprise_team" "test" {
  enterprise_slug = %q
  name = %q
  %s
}
`, testAccConf.enterpriseSlug, name, description)
	}
	sameTeamID := statecheck.CompareValue(compare.ValuesSame())
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { skipUnlessEnterprise(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{Config: config(name, `description = "Remove this description"`), ConfigStateChecks: []statecheck.StateCheck{sameTeamID.AddStateValue("github_enterprise_team.test", tfjsonpath.New("team_id"))}},
			{
				Config: config(name+"-renamed", ""),
				ConfigStateChecks: []statecheck.StateCheck{
					sameTeamID.AddStateValue("github_enterprise_team.test", tfjsonpath.New("team_id")),
					statecheck.ExpectKnownValue("github_enterprise_team.test", tfjsonpath.New("name"), knownvalue.StringExact(name+"-renamed")),
					statecheck.ExpectKnownValue("github_enterprise_team.test", tfjsonpath.New("description"), knownvalue.StringExact("")),
				},
			},
			{Config: config(name+"-renamed", ""), PlanOnly: true},
		},
	})
}

func TestAccGithubEnterpriseTeamOrganizationsUpdateAndRename(t *testing.T) {
	skipWithoutAccConf(t)
	orgA, orgB := testAccConf.testEnterpriseOrg, testAccConf.testEnterpriseOrg2
	if orgA == "" || orgB == "" || orgA == orgB {
		t.Skip("two distinct organizations in the test enterprise are required: GH_TEST_ENTERPRISE_ORG and GH_TEST_ENTERPRISE_ORG_2")
	}
	name := testResourcePrefix + acctest.RandString(5)
	config := func(teamName, org string) string {
		return fmt.Sprintf(`
resource "github_enterprise_team" "test" {
  enterprise_slug = %q
  name = %q
  organization_selection_type = "selected"
}
resource "github_enterprise_team_organizations" "test" {
  enterprise_slug = %q
  team_id = github_enterprise_team.test.team_id
  organization_slugs = [%q]
}
`, testAccConf.enterpriseSlug, teamName, testAccConf.enterpriseSlug, org)
	}
	sameTeamID := statecheck.CompareValue(compare.ValuesSame())
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { skipUnlessEnterprise(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{Config: config(name, orgA), ConfigStateChecks: []statecheck.StateCheck{sameTeamID.AddStateValue("github_enterprise_team_organizations.test", tfjsonpath.New("resolved_team_id"))}},
			{
				Config: config(name+"-renamed", orgB),
				ConfigStateChecks: []statecheck.StateCheck{
					sameTeamID.AddStateValue("github_enterprise_team_organizations.test", tfjsonpath.New("resolved_team_id")),
					statecheck.ExpectKnownValue("github_enterprise_team_organizations.test", tfjsonpath.New("organization_slugs"), knownvalue.SetExact([]knownvalue.Check{knownvalue.StringExact(strings.ToLower(orgB))})),
				},
			},
			{Config: config(name+"-renamed", strings.ToUpper(orgB)), PlanOnly: true},
		},
	})
}

func TestAccGithubEnterpriseTeamNumericImports(t *testing.T) {
	for _, kind := range []string{"membership", "organizations"} {
		t.Run(kind, func(t *testing.T) {
			skipWithoutAccConf(t)
			value := testAccConf.testEnterpriseUser
			argument := "username"
			if kind == "organizations" {
				value = testAccConf.testEnterpriseOrg
				argument = "organization_slugs"
			}
			if value == "" {
				t.Skip("the corresponding GH_TEST_ENTERPRISE_USER or GH_TEST_ENTERPRISE_ORG fixture is required")
			}
			assignment := fmt.Sprintf("%s = %q", argument, value)
			if kind == "organizations" {
				assignment = fmt.Sprintf("%s = [%q]", argument, value)
			}
			address := "github_enterprise_team_" + kind + ".test"
			config := fmt.Sprintf(`
resource "github_enterprise_team" "test" {
 enterprise_slug = %q
 name = %q
 organization_selection_type = "selected"
}
resource "github_enterprise_team_%s" "test" {
 enterprise_slug = %q
 team_id = github_enterprise_team.test.team_id
 %s
}
`, testAccConf.enterpriseSlug, testResourcePrefix+acctest.RandString(5), kind, testAccConf.enterpriseSlug, assignment)
			resource.Test(t, resource.TestCase{
				PreCheck: func() { skipUnlessEnterprise(t) }, ProviderFactories: providerFactories,
				Steps: []resource.TestStep{
					{Config: config},
					{
						// Import blocks verify that importing produces no additional changes.
						ResourceName: address, ImportState: true, ImportStateKind: resource.ImportBlockWithID,
						ImportStateIdFunc: func(s *terraform.State) (string, error) {
							team := s.RootModule().Resources["github_enterprise_team.test"]
							if team == nil || team.Primary == nil {
								return "", fmt.Errorf("enterprise team missing from test state")
							}
							id := testAccConf.enterpriseSlug + "/" + team.Primary.ID
							if kind == "membership" {
								id += "/" + value
							}
							return id, nil
						},
					},
				},
			})
		})
	}
}
