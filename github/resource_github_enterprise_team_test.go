package github

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubEnterpriseTeam(t *testing.T) {
	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

	if isEnterprise != "true" {
		t.Skip("Skipping because `ENTERPRISE_ACCOUNT` is not set or set to false")
	}
	if testEnterprise == "" {
		t.Skip("Skipping because `ENTERPRISE_SLUG` is not set")
	}

	config1 := fmt.Sprintf(`
		data "github_enterprise" "enterprise" {
			slug = "%s"
		}

		resource "github_enterprise_team" "test" {
			enterprise_slug             = data.github_enterprise.enterprise.slug
			name                        = "tf-acc-team-%s"
			description                 = "team for acceptance testing"
			organization_selection_type = "disabled"
		}
	`, testEnterprise, randomID)

	config2 := fmt.Sprintf(`
		data "github_enterprise" "enterprise" {
			slug = "%s"
		}

		resource "github_enterprise_team" "test" {
			enterprise_slug             = data.github_enterprise.enterprise.slug
			name                        = "tf-acc-team-%s"
			description                 = "updated description"
			organization_selection_type = "selected"
		}
	`, testEnterprise, randomID)

	check1 := resource.ComposeAggregateTestCheckFunc(
		resource.TestCheckResourceAttrSet("github_enterprise_team.test", "slug"),
		resource.TestCheckResourceAttrSet("github_enterprise_team.test", "team_id"),
		resource.TestCheckResourceAttr("github_enterprise_team.test", "organization_selection_type", "disabled"),
	)
	check2 := resource.ComposeAggregateTestCheckFunc(
		resource.TestCheckResourceAttr("github_enterprise_team.test", "description", "updated description"),
		resource.TestCheckResourceAttr("github_enterprise_team.test", "organization_selection_type", "selected"),
	)

	testCase := func(t *testing.T, mode string) {
		resource.Test(t, resource.TestCase{
			PreCheck:  func() { skipUnlessMode(t, mode) },
			Providers: testAccProviders,
			Steps: []resource.TestStep{
				{Config: config1, Check: check1},
				{Config: config2, Check: check2},
				{
					ResourceName:        "github_enterprise_team.test",
					ImportState:         true,
					ImportStateVerify:   true,
					ImportStateIdPrefix: fmt.Sprintf(`%s/`, testEnterprise),
				},
			},
		})
	}

	t.Run("with an enterprise account", func(t *testing.T) {
		testCase(t, enterprise)
	})
}

func TestAccGithubEnterpriseTeamOrganizations(t *testing.T) {
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

	config1 := fmt.Sprintf(`
		data "github_enterprise" "enterprise" {
			slug = "%s"
		}

		resource "github_enterprise_team" "test" {
			enterprise_slug             = data.github_enterprise.enterprise.slug
			name                        = "tf-acc-team-orgs-%s"
			organization_selection_type = "selected"
		}

		resource "github_enterprise_team_organizations" "test" {
			enterprise_slug = data.github_enterprise.enterprise.slug
			enterprise_team = github_enterprise_team.test.slug
			organization_slugs = ["%s"]
		}
	`, testEnterprise, randomID, testOrganization)

	config2 := fmt.Sprintf(`
		data "github_enterprise" "enterprise" {
			slug = "%s"
		}

		resource "github_enterprise_team" "test" {
			enterprise_slug             = data.github_enterprise.enterprise.slug
			name                        = "tf-acc-team-orgs-%s"
			organization_selection_type = "selected"
		}

		resource "github_enterprise_team_organizations" "test" {
			enterprise_slug = data.github_enterprise.enterprise.slug
			enterprise_team = github_enterprise_team.test.slug
			organization_slugs = []
		}
	`, testEnterprise, randomID)

	check1 := resource.ComposeAggregateTestCheckFunc(
		resource.TestCheckResourceAttr("github_enterprise_team_organizations.test", "organization_slugs.#", "1"),
		resource.TestCheckTypeSetElemAttr("github_enterprise_team_organizations.test", "organization_slugs.*", testOrganization),
	)
	check2 := resource.ComposeAggregateTestCheckFunc(
		resource.TestCheckResourceAttr("github_enterprise_team_organizations.test", "organization_slugs.#", "0"),
	)

	testCase := func(t *testing.T, mode string) {
		resource.Test(t, resource.TestCase{
			PreCheck:  func() { skipUnlessMode(t, mode) },
			Providers: testAccProviders,
			Steps: []resource.TestStep{
				{Config: config1, Check: check1},
				{Config: config2, Check: check2},
				{
					ResourceName:      "github_enterprise_team_organizations.test",
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	}

	t.Run("with an enterprise account", func(t *testing.T) {
		testCase(t, enterprise)
	})
}

func TestAccGithubEnterpriseTeamMembership(t *testing.T) {
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
			name            = "tf-acc-team-member-%s"
		}

		resource "github_enterprise_team_membership" "test" {
			enterprise_slug = data.github_enterprise.enterprise.slug
			enterprise_team = github_enterprise_team.test.slug
			username        = "%s"
		}
	`, testEnterprise, randomID, username)

	check := resource.ComposeAggregateTestCheckFunc(
		resource.TestCheckResourceAttr("github_enterprise_team_membership.test", "username", username),
	)

	testCase := func(t *testing.T, mode string) {
		resource.Test(t, resource.TestCase{
			PreCheck:  func() { skipUnlessMode(t, mode) },
			Providers: testAccProviders,
			Steps: []resource.TestStep{
				{Config: config, Check: check},
				{
					ResourceName:      "github_enterprise_team_membership.test",
					ImportState:       true,
					ImportStateVerify: true,
				},
			},
		})
	}

	t.Run("with an enterprise account", func(t *testing.T) {
		testCase(t, enterprise)
	})
}
